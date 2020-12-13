package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/maxvoronov/tweetster/internal/users/models"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Login     string             `bson:"login"`
	Email     string             `bson:"email"`
	Name      string             `bson:"name"`
	Password  string             `bson:"password"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

func (u *User) FromModel(user *models.User) error {
	if user.ID != "" {
		var err error
		u.ID, err = primitive.ObjectIDFromHex(user.ID)
		if err != nil {
			return err
		}
	}

	u.Login = user.Login
	u.Email = user.Email
	u.Name = user.Name
	u.Password = user.Password
	u.CreatedAt = user.CreatedAt
	u.UpdatedAt = user.UpdatedAt

	return nil
}

func (u *User) ToModel() *models.User {
	return &models.User{
		ID:        u.ID.Hex(),
		Login:     u.Login,
		Email:     u.Email,
		Name:      u.Name,
		Password:  u.Password,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
