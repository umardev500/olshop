package helper

import "go.mongodb.org/mongo-driver/bson/primitive"

func GetObjID(id string) primitive.ObjectID {
	obj, _ := primitive.ObjectIDFromHex(id)
	return obj
}
