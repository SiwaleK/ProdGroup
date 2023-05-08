package repository

import (
	"context"
	"database/sql"
	"reflect"
	"testing"

	db "github.com/SiwaleK/ProdGroup/db/sqlc"
	"github.com/stretchr/testify/mock"
)

// type Queries interface {
// 	GetProdgroup(ctx context.Context) ([]db.Prodgroup, error)
// 	// Include other methods from *db.Queries that you need to mock
// }

// type mockDBQueries struct {
// 	GetProdgroupfn func(ctx context.Context) ([]db.Prodgroup, error)
// }

// func (w *mockDBQueries) GetProdgroup(ctx context.Context) ([]db.Prodgroup, error) {
// 	return w.GetProdgroupfn(ctx)
// }

// var _ Queries = (*mockDBQueries)(nil)

// func TestDBProdgroupRepository_GetProdgroup(t *testing.T) {
// 	// Mock data
// 	mockDbProdgroups := []db.Prodgroup{
// 		{Prodgroupid: 1, ThName: sql.NullString{String: "Product Group 1", Valid: true}, EnName: sql.NullString{String: "Group 1", Valid: true}},
// 		{Prodgroupid: 2, ThName: sql.NullString{String: "Product Group 2", Valid: true}, EnName: sql.NullString{String: "Group 2", Valid: true}},
// 	}

// 	// Create a mock instance of the db.Queries interface
// 	mockDB := &mockDBQueries{
// 		GetProdgroupfn: func(ctx context.Context) ([]db.Prodgroup, error) {
// 			return mockDbProdgroups, nil
// 		},
// 	}

// 	// Create the repository instance
// 	repo := repository.NewProdgroupRepository(mockDB)

// 	// Create a context
// 	ctx := context.Background()

// 	// Call the method being tested
// 	prodgroups, err := repo.GetProdgroup(ctx)

// 	// Assertions
// 	assert.NoError(t, err)
// 	assert.Len(t, prodgroups, len(mockDbProdgroups))
// 	assert.Equal(t, mockDbProdgroups[0].Prodgroupid, prodgroups[0].Prodgroupid)
// 	assert.Equal(t, mockDbProdgroups[0].ThName, prodgroups[0].ThName)
// 	assert.Equal(t, mockDbProdgroups[0].EnName, prodgroups[0].EnName)
// 	assert.Equal(t, mockDbProdgroups[1].Prodgroupid, prodgroups[1].Prodgroupid)
// 	assert.Equal(t, mockDbProdgroups[1].ThName, prodgroups[1].ThName)
// 	assert.Equal(t, mockDbProdgroups[1].EnName, prodgroups[1].EnName)
// }

func TestGetProdgroup(t *testing.T) {
	repo := &ProdgroupRepositoryMock{}

	// Set up the mock response
	mockProdgroup := &db.Prodgroup{
		Prodgroupid: 1,
		ThName:      sql.NullString{String: "Product Group 1", Valid: true},
		EnName:      sql.NullString{String: "Group 1", Valid: true},
	}
	repo.On("GetProdgroup", mock.Anything).Return(mockProdgroup, nil)

	// Call the GetProdgroup method
	ctx := context.Background()
	prodgroup, err := repo.GetProdgroup(ctx)

	// Assert the expected values
	expectedProdgroup := &db.Prodgroup{
		Prodgroupid: 1,
		ThName:      sql.NullString{String: "Product Group 1", Valid: true},
		EnName:      sql.NullString{String: "Group 1", Valid: true},
	}
	if !reflect.DeepEqual(prodgroup, expectedProdgroup) {
		t.Errorf("Expected prodgroup %+v, but got %+v", expectedProdgroup, prodgroup)
	}
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}
