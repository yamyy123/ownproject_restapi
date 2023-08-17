/*implementation of the profile side*/

package services

import (
	"context"

	"go.mongodb.org/mongo-driver/mango"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProfileService struct{
	ProfileCollection *mongo.Collection
	ctx context.Context
}

func CreateProfile()