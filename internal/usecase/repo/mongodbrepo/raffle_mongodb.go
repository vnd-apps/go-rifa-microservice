package mongodbrepo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/evmartinelli/go-rifa-microservice/internal/entity"
	"github.com/evmartinelli/go-rifa-microservice/pkg/mongodb"
)

type RaffleRepo struct {
	db *mongodb.MongoCol
}

const collection = "rifas-collection"

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
		db: &mongodb.MongoCol{
			Collection: mdb.Database.Collection(collection),
		},
	}
}

func (r *RaffleRepo) Create(ctx context.Context, rm entity.Raffle) error {
	model := toModel(&rm)

	_, err := r.db.Collection.InsertOne(ctx, model)
	if err != nil {
		return err
	}

	return nil
}

func (r *RaffleRepo) GetAvailableRaffle(ctx context.Context) ([]entity.Raffle, error) {
	cur, err := r.db.Collection.Find(ctx, bson.M{})
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

	return toRaffles(out), nil
}

func toModel(r *entity.Raffle) *Raffle {
	return &Raffle{
		ID:           primitive.NewObjectID(),
		Name:         r.Name,
		Status:       "Available",
		Value:        r.Value,
		TotalNumbers: r.TotalNumbers,
		TotalSold:    0,
	}
}

func toRaffle(b *Raffle) entity.Raffle {
	return entity.Raffle{
		ID:           b.ID.Hex(),
		Name:         b.Name,
		Status:       b.Status,
		Value:        b.Value,
		TotalNumbers: b.TotalNumbers,
		TotalSold:    b.TotalSold,
	}
}

func toRaffles(rs []*Raffle) []entity.Raffle {
	out := make([]entity.Raffle, len(rs))

	for i, b := range rs {
		out[i] = toRaffle(b)
	}

	return out
}
