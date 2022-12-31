package db

import (
	"context"
	"log"
	"shiva/models"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (mr *MongoRepo) GetPosts(username string, postSearchRequest *models.PostSearchRequest) (*[]models.Post, error) {

	searchText := strings.Trim(postSearchRequest.Search, " ")

	userF := bson.M{}
	if username != "" {
		user, err := mr.GetUser(username)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		userF = bson.M{
			"from_user": user.UserID,
		}
	}

	searchF := bson.M{}
	if searchText != "" {
		searchF = bson.M{
			"$text": bson.M{
				"$search": searchText,
			},
		}
	}
	tagsF := bson.M{}
	if len(postSearchRequest.Tags) != 0 {
		tagsA := bson.A{}
		for _, tag := range postSearchRequest.Tags {
			tagsA = append(tagsA, tag)
		}

		tagsF = bson.M{
			"tags": bson.M{
				"$in": tagsA,
			},
		}
	}

	filter := bson.M{
		"$and": bson.A{
			userF, searchF, tagsF,
		},
	}

	// pipeline := []bson.D{}
	// if searchText != "" {
	// 	pipeline = append(pipeline, bson.D{{"$match", bson.D{{"$text", bson.D{{"$search", searchText}}}}}})
	// }
	// if len(postSearchRequest.Tags) != 0 {
	// 	pipeline = append(pipeline, bson.D{
	// 		{"$match",
	// 			bson.D{
	// 				{"tags", bson.D{{"$in", bson.A{"it"}}}},
	// 			},
	// 		},
	// 	})
	// }
	// resCur, err := mr.Coll.Posts.Aggregate(context.TODO(), mongo.Pipeline(pipeline))
	resCur, err := mr.Coll.Posts.Find(context.TODO(), filter)
	if err != nil {
		log.Println("error while fetching posts: ", err.Error())
		return nil, err
	}

	posts := make([]models.Post, 0)
	err = resCur.All(context.TODO(), &posts)
	if err != nil {
		log.Println("error while unmarshalling user posts:", err.Error())
		return nil, err
	}
	return &posts, nil
}

func (mr *MongoRepo) MyPosts(userID string) (*[]models.Post, error) {
	filter := bson.M{
		"from_user": userID,
	}
	resCur, err := mr.Coll.Posts.Find(context.TODO(), filter)
	if err != nil {
		log.Println("error while fething self posts:", err.Error())
		return nil, err
	}
	posts := make([]models.Post, 0)
	err = resCur.All(context.TODO(), &posts)
	if err != nil {
		log.Println("error while unmarshalling self posts:", err.Error())
		return nil, err
	}
	return &posts, nil
}
func (mr *MongoRepo) CreatePost(post *models.Post) error {

	_, err := mr.Coll.Posts.InsertOne(context.TODO(), post)
	if err != nil {
		log.Println("error while inserting post: ", err.Error())
		return err
	}
	return nil
}

func (mr *MongoRepo) EditPost(postEditRequest *models.PostEditRequest) error {

	updateM := bson.M{}
	b, _ := bson.Marshal(postEditRequest)
	bson.Unmarshal(b, updateM)
	_id, _ := primitive.ObjectIDFromHex(postEditRequest.PostID)

	update := bson.M{
		"$set": updateM,
	}

	_, err := mr.Coll.Posts.UpdateByID(context.TODO(), _id, update)
	if err != nil {
		log.Println("error while updating post: ", err.Error())
		return err
	}
	return nil
}

func (mr *MongoRepo) DeletePost(userID, postID string) error {

	filter := bson.M{
		"from_user": userID,
		"post_id":   postID,
	}
	_, err := mr.Coll.Posts.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Println("error while deleting post: ", err.Error())
		return err
	}
	return nil
}

func (mr *MongoRepo) Comment(comment *models.Comment) error {
	_id, _ := primitive.ObjectIDFromHex(comment.PostID)

	updateM := bson.M{}
	b, _ := bson.Marshal(comment)
	bson.Unmarshal(b, updateM)
	update := bson.M{
		"$push": bson.M{
			"comments": updateM,
		},
	}

	_, err := mr.Coll.Posts.UpdateByID(context.TODO(), _id, update)
	if err != nil {
		log.Println("error while inserting comment: ", err.Error())
		return err
	}
	return nil
}

func (mr *MongoRepo) Like(postID string) {

	_id, _ := primitive.ObjectIDFromHex(postID)
	update := bson.M{
		"$inc": bson.M{
			"likes": 1,
		},
	}
	mr.Coll.Posts.UpdateByID(context.TODO(), _id, update)
}
