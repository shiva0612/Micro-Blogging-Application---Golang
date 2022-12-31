package db

import (
	"context"
	"fmt"
	"shiva/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Collections struct {
	Users    *mongo.Collection
	Posts    *mongo.Collection
	Comments *mongo.Collection
}

type MongoRepo struct {
	Dbc  *mongo.Client
	Coll *Collections
}

func ConnectDB() *MongoRepo {

	uri := fmt.Sprintf("mongodb://%s:%s@localhost:%s/?authSource=%s",
		config.Cfg.Mongo.Username,
		config.Cfg.Mongo.Password,
		config.Cfg.Mongo.Port,
		config.Cfg.Mongo.Authdb)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic("error connecting to MongoDB: " + err.Error())
	}

	colls := new(Collections)
	colls.Users = client.Database(config.Cfg.Mongo.Dbname).Collection("Users")
	colls.Posts = client.Database(config.Cfg.Mongo.Dbname).Collection("Posts")
	colls.Comments = client.Database(config.Cfg.Mongo.Dbname).Collection("Comments")

	return &MongoRepo{Dbc: client, Coll: colls}
}

func (mr *MongoRepo) DisconnectDB() {
	if err := mr.Dbc.Disconnect(context.TODO()); err != nil {
		panic("error disconnecting from MongoDB: " + err.Error())
	}
}
