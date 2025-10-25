package repository

import (
	"context"
	"go-fiber/app/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FileRepository struct {
	collection *mongo.Collection
}

func NewFileRepository(db *mongo.Database) *FileRepository {
	return &FileRepository{
		collection: db.Collection("files"),
	}
}

func (r *FileRepository) CreateFile(ctx context.Context, file *model.File) error {
	_, err := r.collection.InsertOne(ctx, file)
	return err
}

func (r *FileRepository) GetFiles(ctx context.Context, userID string, role string) ([]model.FileWithUser, error) {
	var pipeline mongo.Pipeline

	if role == "admin" {
		pipeline = mongo.Pipeline{
			{{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "users"},
				{Key: "localField", Value: "user_id"},
				{Key: "foreignField", Value: "_id"},
				{Key: "as", Value: "user"},
			}}},
		}
	} else {
		userObjID, _ := primitive.ObjectIDFromHex(userID)
		pipeline = mongo.Pipeline{
			{{Key: "$match", Value: bson.D{{Key: "user_id", Value: userObjID}}}},
			{{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "users"},
				{Key: "localField", Value: "user_id"},
				{Key: "foreignField", Value: "_id"},
				{Key: "as", Value: "user"},
			}}},
		}
	}

	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var files []model.FileWithUser
	if err := cursor.All(ctx, &files); err != nil {
		return nil, err
	}

	return files, nil
}

func (r *FileRepository) GetFileByID(ctx context.Context, id string) (*model.File, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var file model.File
	err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&file)
	if err != nil {
		return nil, err
	}

	return &file, nil
}

func (r *FileRepository) DeleteFileByID(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}
