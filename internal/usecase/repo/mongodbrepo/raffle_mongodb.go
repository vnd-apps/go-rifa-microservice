package mongodbrepo

import (
	"context"

	"github.com/evmartinelli/go-rifa-microservice/internal/entity"
	"github.com/evmartinelli/go-rifa-microservice/pkg/mongodb"
	"go.mongodb.org/mongo-driver/bson"
)

// TranslationRepo -.
type RaffleRepo struct {
	db *mongodb.MongoDB
}

type Bookmark struct {
	URL   string `bson:"url"`
	Title string `bson:"title"`
}

// New -.
func NewRaffle(mdb *mongodb.MongoDB) *RaffleRepo {
	return &RaffleRepo{
		db: mdb}
}

func (r *RaffleRepo) Create(ctx context.Context, model entity.Raffle) error {
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
	out := make([]*Bookmark, 0)

	for cur.Next(ctx) {
		user := new(Bookmark)
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

func toBookmark(b *Bookmark) entity.Raffle {
	return entity.Raffle{
		Value: 5,
		Name:  b.Title,
	}
}

func toBookmarks(bs []*Bookmark) []entity.Raffle {
	out := make([]entity.Raffle, len(bs))

	for i, b := range bs {
		out[i] = toBookmark(b)
	}

	return out
}
