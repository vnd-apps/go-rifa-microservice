package mongodbrepo

import (
	"context"

	"github.com/evmartinelli/go-rifa-microservice/internal/entity"
	"github.com/evmartinelli/go-rifa-microservice/pkg/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TranslationRepo -.
type RaffleRepo struct {
	db *mongodb.MongoDB
}
type Raffle struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `bson:"title"`
	Status       string             `bson:"status"`
	Value        int                `bson:"value"`
	TotalNumbers int                `bson:"totalnumbers"`
	TotalSold    int                `bson:"totalsold"`
}

// New -.
func NewRaffle(mdb *mongodb.MongoDB) *RaffleRepo {
	return &RaffleRepo{
		db: mdb,
	}
}

func (r *RaffleRepo) Create(ctx context.Context, rm entity.Raffle) error {
	model := toModel(&rm)
	_, err := r.db.Database.Collection("rifas-collection").InsertOne(ctx, model)
	if err != nil {
		return err
	}
	return nil
}
func (r *RaffleRepo) GetAvaliableRaffle(ctx context.Context) ([]entity.Raffle, error) {
	cur, err := r.db.Database.Collection("rifas-collection").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	out := make([]*Raffle, 0)

	for cur.Next(ctx) {
		user := new(Raffle)
		err := cur.Decode(user)
		if err != nil {
			return nil, err
		}

		out = append(out, user)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	return toBookmarks(out), nil
}

func toModel(b *entity.Raffle) *Raffle {
	return &Raffle{
		ID:           primitive.NewObjectID(),
		Name:         b.Name,
		Status:       "Avaliable",
		Value:        b.Value,
		TotalNumbers: b.TotalNumbers,
		TotalSold:    0,
	}
}

func toBookmark(b *Raffle) entity.Raffle {
	return entity.Raffle{
		Id:           b.ID.Hex(),
		Name:         b.Name,
		Status:       b.Status,
		Value:        b.Value,
		TotalNumbers: b.TotalNumbers,
		TotalSold:    b.TotalSold,
	}
}

func toBookmarks(bs []*Raffle) []entity.Raffle {
	out := make([]entity.Raffle, len(bs))

	for i, b := range bs {
		out[i] = toBookmark(b)
	}

	return out
}
