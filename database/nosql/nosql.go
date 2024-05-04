package nosql

import (
	"context"

	"emvn/consts"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// NoSQLInterface is the abstraction layer for NoSQL database. Due to time limit I will not abstract the output
// This implementation ensures dependency inversion principle and makes the code more testable
type NoSQLInterface interface {
	InsertOne(ctx context.Context, collection consts.NoSQLCollection, document interface{}) (*mongo.InsertOneResult, error)
	InsertMultiple(ctx context.Context, collection consts.NoSQLCollection, documents []interface{}) (*mongo.InsertManyResult, error)
	FindByObjectID(ctx context.Context, collection consts.NoSQLCollection, objectID string) (*mongo.SingleResult, error)
	Find(ctx context.Context, collection consts.NoSQLCollection, filter interface{}, options ...*options.FindOptions) (*mongo.Cursor, error)
	FindOne(ctx context.Context, collection consts.NoSQLCollection, filter interface{}, options ...*options.FindOneOptions) (*mongo.SingleResult, error)
	UpdateByID(ctx context.Context, collection consts.NoSQLCollection, id string, document interface{}) (*mongo.UpdateResult, error)
	UpdateOne(ctx context.Context, collection consts.NoSQLCollection, filter interface{}, document interface{}, options ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	UpdateMany(ctx context.Context, collection consts.NoSQLCollection, filter interface{}, document interface{}) (*mongo.UpdateResult, error)
	CreateIfNotExists(ctx context.Context, collection consts.NoSQLCollection, filter interface{}, document interface{}) (*mongo.SingleResult, error)
	Aggregate(ctx context.Context, collection consts.NoSQLCollection, pipeline interface{}) (*mongo.Cursor, error)
	Count(ctx context.Context, collection consts.NoSQLCollection, filter interface{}) (int64, error)
	DeleteByID(ctx context.Context, collection consts.NoSQLCollection, id string) error
}
