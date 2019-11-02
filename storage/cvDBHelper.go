package storage

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetCVID returns id
func GetCVID(uname string) primitive.ObjectID {
	collection := client.Database("faculry_portal").Collection("cv")
	cvID, err := GetFacultyCV(uname)
	if err != nil {
		var result *mongo.InsertOneResult
		result, err = collection.InsertOne(context.TODO(), CVDetail{})
		if err == nil {
			if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
				SetFacultyCV(uname, oid.Hex())
				cvID = oid.Hex()
			}
		}
	}
	objID, _ := primitive.ObjectIDFromHex(cvID)
	return objID
}

// SaveBio saves bio
func SaveBio(uname, bio string) error {
	collection := client.Database("faculry_portal").Collection("cv")
	objID := GetCVID(uname)
	_, err := collection.UpdateOne(context.TODO(), bson.M{"_id": bson.M{"$eq": objID}},
		bson.M{"$set": bson.M{"overview.biography": bio}})
	return err
}

// SaveAboutme saves aboutme
func SaveAboutme(uname, aboutme string) error {
	collection := client.Database("faculry_portal").Collection("cv")
	objID := GetCVID(uname)
	_, err := collection.UpdateOne(context.TODO(), bson.M{"_id": bson.M{"$eq": objID}},
		bson.M{"$set": bson.M{"overview.aboutme": aboutme}})
	return err
}

// AddProject add new project
func AddProject(uname string, project CVProject) error {
	collection := client.Database("faculry_portal").Collection("cv")
	objID := GetCVID(uname)
	_, err := collection.UpdateOne(context.TODO(), bson.M{"_id": bson.M{"$eq": objID}},
		bson.M{"$push": bson.M{"projects": project}})
	return err
}

// DeleteProject delete a project
func DeleteProject(uname, project string) error {
	collection := client.Database("faculry_portal").Collection("cv")
	objID := GetCVID(uname)
	_, err := collection.UpdateOne(context.TODO(), bson.M{"_id": bson.M{"$eq": objID}},
		bson.M{"$pull": bson.M{"projects": CVProject{Title: project}}})
	return err
}

// AddPrize add new prize
func AddPrize(uname string, prize CVPrize) error {
	collection := client.Database("faculry_portal").Collection("cv")
	objID := GetCVID(uname)
	_, err := collection.UpdateOne(context.TODO(), bson.M{"_id": bson.M{"$eq": objID}},
		bson.M{"$push": bson.M{"prizes": prize}})
	return err
}

// DeletePrize delete a prize
func DeletePrize(uname, prize string) error {
	collection := client.Database("faculry_portal").Collection("cv")
	objID := GetCVID(uname)
	_, err := collection.UpdateOne(context.TODO(), bson.M{"_id": bson.M{"$eq": objID}},
		bson.M{"$pull": bson.M{"prizes": CVPrize{Title: prize}}})
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
	Uname    string      `bson:"uname,omitempty"`
	Overview CVOverview  `bson:"overview,omitempty"`
	Project  []CVProject `bson:"projects,omitempty"`
	Prizes   []CVPrize   `bson:"prizes,omitempty"`
}

// CVOverview struct
type CVOverview struct {
	Biography string `bson:"biography,omitempty"`
	AboutMe   string `bson:"aboutme,omitempty"`
}

// CVProject struct
type CVProject struct {
	Title  string `bson:"title,omitempty"`
	Detail string `bson:"detail,omitempty"`
}

// CVPrize struct
type CVPrize struct {
	Title string `bson:"title,omitempty"`
	Prize string `bson:"prize,omitempty"`
}
