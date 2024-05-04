package mongodb

import (
	"context"
	"emvn/config"
	"emvn/consts"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoClient struct {
	Client *mongo.Database
}

var client mongoClient

func InitClient(ctx context.Context) {
	cmdMonitor := &event.CommandMonitor{
		Started: func(_ context.Context, evt *event.CommandStartedEvent) {
			log.Print(evt.Command)
		},
	}
	uri := config.GetConfig().Database.ConnectString
	conn, err := mongo.Connect(ctx, options.Client().ApplyURI(uri).SetMonitor(cmdMonitor))
	if err != nil {
		panic(err)
	}

	client = mongoClient{
		Client: conn.Database(config.GetConfig().Database.DBName),
	}
}

func MongoDBClient() mongoClient {
	return client
}

func (m mongoClient) InsertOne(ctx context.Context, collection consts.NoSQLCollection, document interface{}) (*mongo.InsertOneResult, error) {
	res, err := m.Client.Collection(collection.String()).InsertOne(ctx, document)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (m mongoClient) InsertMultiple(ctx context.Context, collection consts.NoSQLCollection, documents []interface{}) (*mongo.InsertManyResult, error) {
	res, err := m.Client.Collection(collection.String()).InsertMany(ctx, documents)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (m mongoClient) FindByObjectID(ctx context.Context, collection consts.NoSQLCollection, objectID string) (*mongo.SingleResult, error) {
	id, err := primitive.ObjectIDFromHex(objectID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": id}

	res := m.Client.Collection(collection.String()).FindOne(ctx, filter)
	if res.Err() != nil {
		return nil, res.Err()
	}

	return res, nil
}

func (m mongoClient) Find(ctx context.Context, collection consts.NoSQLCollection, filter interface{}, options ...*options.FindOptions) (*mongo.Cursor, error) {
	res, err := m.Client.Collection(collection.String()).Find(ctx, filter, options...)
	if err != nil {
		return nil, err
	}
	if res.Err() != nil {
		return nil, res.Err()
	}

	return res, nil
}

func (m mongoClient) FindOne(ctx context.Context, collection consts.NoSQLCollection, filter interface{}, options ...*options.FindOneOptions) (*mongo.SingleResult, error) {
	res := m.Client.Collection(collection.String()).FindOne(ctx, filter, options...)
	if res.Err() != nil {
		return nil, res.Err()
	}
	return res, nil
}

func (m mongoClient) UpdateByID(ctx context.Context, collection consts.NoSQLCollection, id string, document interface{}) (*mongo.UpdateResult, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	res, err := m.Client.Collection(collection.String()).UpdateByID(ctx, objectID, document)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (m mongoClient) UpdateOne(ctx context.Context, collection consts.NoSQLCollection, filter interface{}, document interface{}, options ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	res, err := m.Client.Collection(collection.String()).UpdateOne(ctx, filter, document, options...)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (m mongoClient) UpdateMany(ctx context.Context, collection consts.NoSQLCollection, filter interface{}, document interface{}) (*mongo.UpdateResult, error) {
	res, err := m.Client.Collection(collection.String()).UpdateMany(ctx, filter, document)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (m mongoClient) CreateIfNotExists(ctx context.Context, collection consts.NoSQLCollection, filter interface{}, document interface{}) (*mongo.SingleResult, error) {
	opts := options.Update().SetUpsert(true)

	res, err := m.Client.Collection(collection.String()).UpdateOne(ctx, filter, bson.M{"$set": document}, opts)
	if err != nil {
		return nil, err
	}
	if res.UpsertedID != nil {
		return m.FindByObjectID(ctx, consts.NoSQLCollection(collection.String()), res.UpsertedID.(primitive.ObjectID).Hex())
	}

	return m.FindOne(ctx, consts.NoSQLCollection(collection.String()), filter)
}

func (m mongoClient) Aggregate(ctx context.Context, collection consts.NoSQLCollection, pipeline interface{}) (*mongo.Cursor, error) {
	// Perform the aggregation
	cur, err := m.Client.Collection(collection.String()).Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	if cur.Err() != nil {
		return nil, cur.Err()
	}

	return cur, nil
}

func (m mongoClient) Count(ctx context.Context, collection consts.NoSQLCollection, filter interface{}) (int64, error) {
	count, err := m.Client.Collection(collection.String()).CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (m mongoClient) DeleteByID(ctx context.Context, collection consts.NoSQLCollection, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = m.Client.Collection(collection.String()).DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return err
	}
	return nil
}
