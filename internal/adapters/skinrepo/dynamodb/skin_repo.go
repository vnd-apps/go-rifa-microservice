package dynamodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/evmartinelli/go-rifa-microservice/internal/core/skin"
	"github.com/evmartinelli/go-rifa-microservice/pkg/mongodb"
)

type PlayerSkinRepo struct {
	db *mongodb.MongoCol
}

const playerSkinCollection = "player-skin-collection"

type Skin struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `bson:"title"`
	Status       string             `bson:"status"`
	Value        int                `bson:"value"`
	TotalNumbers int                `bson:"totalnumbers"`
	TotalSold    int                `bson:"totalsold"`
}

func NewPlayerSkinRepo(mdb *mongodb.MongoDB) *PlayerSkinRepo {
	return &PlayerSkinRepo{
		db: &mongodb.MongoCol{
			Collection: mdb.Database.Collection(playerSkinCollection),
		},
	}
}

func (r *PlayerSkinRepo) Create(ctx context.Context, rm skin.Skin) error {
	rm.ID = primitive.NewObjectID().Hex()

	_, err := r.db.Collection.InsertOne(ctx, rm)
	if err != nil {
		return err
	}

	return nil
}
