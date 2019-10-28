package storage

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// SaveCVDetails returns cv details
func SaveCVDetails(cvdetail CVDetail) error {
	collection := client.Database("faculry_portal").Collection("cv")
	cvID, err := GetFacultyCV(cvdetail.Uname)

	if err != nil {
		var result *mongo.InsertOneResult
		result, err = collection.InsertOne(context.TODO(), cvdetail)
		if err == nil {
			if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
				SetFacultyCV(cvdetail.Uname, oid.Hex())
			} else {
				return errors.New("Error in getting Inserted Id")
			}
		}
	} else {
		objID, err := primitive.ObjectIDFromHex(cvID)
		if err != nil {
			return err
		}
		_, err = collection.UpdateOne(context.TODO(), bson.M{"_id": bson.M{"$eq": objID}}, bson.M{"$set": cvdetail})
	}
	return err
}

// GetCVDetails returns cv details
func GetCVDetails(uname string) (CVDetail, error) {
	cvdetail := CVDetail{}
	collection := client.Database("faculry_portal").Collection("cv")
	cvID, err := GetFacultyCV(uname)
	if err != nil {
		return cvdetail, err
	}
	objID, err := primitive.ObjectIDFromHex(cvID)
	if err != nil {
		return cvdetail, err
	}
	result := collection.FindOne(context.TODO(), bson.M{"_id": bson.M{"$eq": objID}})

	err = result.Decode(&cvdetail)
	return cvdetail, err
}

// CVDetail struct for cv
type CVDetail struct {
	Uname    string     `bson:"uname"`
	Overview CVOverview `bson:"overview"`
}

// CVOverview struct
type CVOverview struct {
	Biography string `bson:"biography"`
	AboutMe   string `bson:"aboutme"`
}
