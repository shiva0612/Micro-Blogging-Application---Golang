package db

import (
	"context"
	"errors"
	"log"
	"shiva/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (mr *MongoRepo) Signup(user *models.User) error {
	_, err := mr.Coll.Users.InsertOne(context.TODO(), user)
	if err != nil {
		log.Println("error inserting user: ", err.Error())
		return err
	}
	return nil
}
func (mr *MongoRepo) Login(username string) (*models.User, error) {

	filter := bson.M{"user_name": username}
	res := mr.Coll.Users.FindOne(context.TODO(), filter)
	if res.Err() != nil {
		log.Println("error finding user: ", res.Err().Error())
		return nil, res.Err()
	}
	user := new(models.User)
	res.Decode(user)

	return user, nil
}
func (mr *MongoRepo) EditUser(userID string, userEditRequest *models.UserEditRequest) error {
	_id, _ := primitive.ObjectIDFromHex(userID)
	b, _ := bson.Marshal(userEditRequest)
	var update = bson.M{}
	bson.Unmarshal(b, &update)
	_, err := mr.Coll.Users.UpdateByID(context.TODO(), _id, bson.M{"$set": update})
	if err != nil {
		log.Println("error while updating user: ", err.Error())
		return err
	}
	return nil
}

func (mr *MongoRepo) DeleteUser(userID string) error {
	_id, _ := primitive.ObjectIDFromHex(userID)
	filter := bson.M{
		"_id": _id,
	}
	_, err := mr.Coll.Users.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Println("error while deleting user: ", err.Error())
		return err
	}
	return nil
}
func (mr *MongoRepo) ViewUser(userID string) (*models.User, error) {
	_id, _ := primitive.ObjectIDFromHex(userID)
	filter := bson.M{
		"_id": _id,
	}
	res := mr.Coll.Users.FindOne(context.TODO(), filter)
	if res.Err() != nil {
		log.Println("error finding user: ", res.Err().Error())
		return nil, res.Err()
	}
	user := new(models.User)
	res.Decode(user)
	return user, nil
}
func (mr *MongoRepo) GetUsers(username string) (*[]models.User, error) {

	filter := bson.M{
		"user_name": bson.M{
			"$regex":   ".*" + username + ".*",
			"$options": "si",
		},
	}
	resCur, err := mr.Coll.Users.Find(context.TODO(), filter)
	if err != nil {
		log.Println("error while getting users: ", err.Error())
		return nil, err
	}
	users := []models.User{}
	err = resCur.All(context.TODO(), &users)
	if err != nil {
		log.Println("cursor error: ", err.Error())
	}
	return &users, nil
}
func (mr *MongoRepo) GetUser(username string) (*models.User, error) {
	filter := bson.M{
		"user_name": username,
	}
	res := mr.Coll.Users.FindOne(context.TODO(), filter)
	if res.Err() != nil {
		log.Println("username not found: ", res.Err().Error())
		return nil, errors.New("user not found")
	}
	user := new(models.User)
	err := res.Decode(user)
	if err != nil {
		log.Println("error while decoding user: ", err.Error())
		return nil, err
	}
	return user, nil
}
func (mr *MongoRepo) Follow(toFollow, userID string) error {
	ctx := context.TODO()
	_uid, _ := primitive.ObjectIDFromHex(userID)
	_fid, _ := primitive.ObjectIDFromHex(toFollow)

	session, err := mr.Dbc.StartSession()
	if err != nil {
		log.Println("could not start session: ", err.Error())
		return err
	}
	defer session.EndSession(ctx)
	session.WithTransaction(ctx, func(sessCtx mongo.SessionContext) (interface{}, error) {
		_, err := mr.Coll.Users.UpdateByID(ctx, _uid, bson.M{
			"$addToSet": bson.M{
				"following": toFollow,
			},
		})
		if err != nil {
			return nil, err
		}
		_, err = mr.Coll.Users.UpdateByID(ctx, _fid, bson.M{
			"$addToSet": bson.M{
				"followers": userID,
			},
		})
		if err != nil {
			return nil, err
		}
		return nil, nil
	})
	return nil
}

func (mr *MongoRepo) HomeFeed(userID string) (*[]models.Post, error) {

	pipe := bson.A{
		bson.D{{"$match", bson.D{{"user_name", "shiva"}}}},
		bson.D{
			{"$lookup",
				bson.D{
					{"from", "Posts"},
					{"localField", "following"},
					{"foreignField", "from_user"},
					{"as", "posts"},
				},
			},
		},
		bson.D{{"$unwind", bson.D{{"path", "$posts"}}}},
		bson.D{{"$replaceWith", "$posts"}},
	}

	resCur, err := mr.Coll.Users.Aggregate(context.TODO(), pipe)
	if err != nil {
		log.Println("error while gettimg home feed: ", err.Error())
		return nil, err
	}
	posts := []models.Post{}
	err = resCur.All(context.TODO(), &posts)
	if err != nil {
		log.Println("error while gettimg home feed (decoding): ", err.Error())
		return nil, err
	}

	return &posts, nil
}

// ----------------------helper functions------------------------------------------
