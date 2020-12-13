package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/maxvoronov/tweetster/internal/users/models"
	"github.com/maxvoronov/tweetster/internal/users/repositories/mongo/entities"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database, collection string) *UserRepository {
	return &UserRepository{
		collection: db.Collection(collection),
	}
}

func (repo *UserRepository) Create(ctx context.Context, user *models.User) error {
	u := &entities.User{}
	if err := u.FromModel(user); err != nil {
		return err
	}

	res, err := repo.collection.InsertOne(ctx, u)
	if err != nil {
		return err
	}

	user.ID = res.InsertedID.(primitive.ObjectID).Hex()

	return nil
}

func (repo *UserRepository) GetById(ctx context.Context, id string) (*models.User, error) {
	user := new(entities.User)
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	err = repo.collection.FindOne(ctx, bson.M{"_id": objectId}).Decode(user)
	if err != nil {
		return nil, err
	}

	return user.ToModel(), nil
}
