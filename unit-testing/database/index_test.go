// mongodb_test.go
package database

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MockClient struct {
    mock.Mock
}

func (mc *MockClient) Database(name string, opts ...*options.DatabaseOptions) *mongo.Database {
    args := mc.Called(name)
    return args.Get(0).(*mongo.Database)
}

type MockDatabase struct {
    mock.Mock
}

func (md *MockDatabase) Collection(name string, opts ...*options.CollectionOptions) *mongo.Collection {
    args := md.Called(name)
    return args.Get(0).(*mongo.Collection)
}

func TestDatabaseHandler_FindDocument(t *testing.T) {
    // Create a mock client, database, and collection
    mockClient := new(MockClient)
    mockDb := new(MockDatabase)
    mockCollection := new(MockCollection)

    // Define the expected filter
    filter := bson.M{"key": "value"}

    // Create an instance of DatabaseHandler
    dh := &DatabaseHandler{
        client: mockClient,
        db:     mockDb,
    }

    // Set up expectations for the FindOne method
    mockCollection.On("FindOne", mock.Anything, filter).Return(&MockSingleResult{}, nil)
    mockDb.On("Collection", "mycollection").Return(mockCollection)

    // Call the method under test
    _, err := dh.FindDocument("mycollection", filter)

    // Assert that the expectations were met
    assert.NoError(t, err)
    mockClient.AssertExpectations(t)
    mockDb.AssertExpectations(t)
    mockCollection.AssertExpectations(t)
}

func main() {
    // Run the tests
    t := new(testing.T)
    TestDatabaseHandler_FindDocument(t)
}
