package storage

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

// CreateCV initializes an entry in mongoDB
func CreateCV(empID string) error {
	collection := client.Database("faculty_portal").Collection("cv")
	cvDetail := CVDetail{EmpID: empID}
	_, err := collection.InsertOne(context.TODO(), cvDetail)
	return err
}

// SaveBio saves bio
func SaveBio(empID, bio string) error {
	collection := client.Database("faculty_portal").Collection("cv")
	_, err := collection.UpdateOne(context.TODO(), bson.M{"empID": bson.M{"$eq": empID}},
		bson.M{"$set": bson.M{"overview.biography": bio}})
	return err
}

// SaveAboutme saves aboutme
func SaveAboutme(empID, aboutme string) error {
	collection := client.Database("faculty_portal").Collection("cv")
	_, err := collection.UpdateOne(context.TODO(), bson.M{"empID": bson.M{"$eq": empID}},
		bson.M{"$set": bson.M{"overview.aboutme": aboutme}})
	return err
}

// AddProject add new project
func AddProject(empID string, project CVProject) error {
	collection := client.Database("faculty_portal").Collection("cv")
	_, err := collection.UpdateOne(context.TODO(), bson.M{"empID": bson.M{"$eq": empID}},
		bson.M{"$push": bson.M{"projects": project}})
	return err
}

// DeleteProject delete a project
func DeleteProject(empID, project string) error {
	collection := client.Database("faculty_portal").Collection("cv")
	_, err := collection.UpdateOne(context.TODO(), bson.M{"empID": bson.M{"$eq": empID}},
		bson.M{"$pull": bson.M{"projects": CVProject{Title: project}}})
	return err
}

// AddPrize add new prize
func AddPrize(empID string, prize CVPrize) error {
	collection := client.Database("faculty_portal").Collection("cv")
	_, err := collection.UpdateOne(context.TODO(), bson.M{"empID": bson.M{"$eq": empID}},
		bson.M{"$push": bson.M{"prizes": prize}})
	return err
}

// DeletePrize delete a prize
func DeletePrize(empID, prize string) error {
	collection := client.Database("faculty_portal").Collection("cv")
	_, err := collection.UpdateOne(context.TODO(), bson.M{"empID": bson.M{"$eq": empID}},
		bson.M{"$pull": bson.M{"prizes": CVPrize{Title: prize}}})
	return err
}

// GetCVDetails returns cv details
func GetCVDetails(empID string) (CVDetail, error) {
	cvdetail := CVDetail{}
	collection := client.Database("faculty_portal").Collection("cv")

	result := collection.FindOne(context.TODO(), bson.M{"empID": bson.M{"$eq": empID}})

	err := result.Decode(&cvdetail)
	return cvdetail, err
}

// CVDetail struct for cv
type CVDetail struct {
	EmpID    string      `bson:"empID,omitempty"`
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
