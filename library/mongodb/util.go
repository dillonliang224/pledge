package mongodb

import "go.mongodb.org/mongo-driver/bson/primitive"

func ConvertToObjectID(id string) (primitive.ObjectID, bool) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return primitive.NilObjectID, false
	}

	return objectId, true
}

func ConvertToObjectIDs(ids []string) (objectIds []primitive.ObjectID, ok bool) {
	objectIds = make([]primitive.ObjectID, len(ids))

	for _, id := range ids {
		objectId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, false
		}
		objectIds = append(objectIds, objectId)
	}

	return objectIds, true
}
