package repository

// func TestGetProdgroup(t *testing.T) {
// 	repo := &ProdgroupRepositoryMock{}

// 	// Set up the mock response
// 	mockProdgroup := &db.Prodgroup{
// 		Prodgroupid: 1,
// 		ThName:      sql.NullString{String: "Product Group 1", Valid: true},
// 		EnName:      sql.NullString{String: "Group 1", Valid: true},
// 	}
// 	repo.On("GetProdgroup", mock.Anything).Return(mockProdgroup, nil)

// 	// Call the GetProdgroup method
// 	ctx := context.Background()
// 	prodgroup, err := repo.GetProdgroup(ctx)

// 	// Assert the expected values
// 	expectedProdgroup := &db.Prodgroup{
// 		Prodgroupid: 1,
// 		ThName:      sql.NullString{String: "Product Group 1", Valid: true},
// 		EnName:      sql.NullString{String: "Group 1", Valid: true},
// 	}
// 	if !reflect.DeepEqual(prodgroup, expectedProdgroup) {
// 		t.Errorf("Expected prodgroup %+v, but got %+v", expectedProdgroup, prodgroup)
// 	}
// 	if err != nil {
// 		t.Errorf("Expected no error, but got: %v", err)
// 	}
// }
