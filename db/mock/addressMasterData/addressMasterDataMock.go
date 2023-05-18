package mocks

import repo "example.com/go-crud-api/repository/addressMasterData"

// สร้าง struct ของตัว mock
type MockAddressMasterDataRepository struct {
	MockGetAllZipcodeRepository              func() ([]int32, error)
	MockGetAllProvinceRepository             func() ([]repo.ReadAllProvinceResp, error)
	MockGetAddressByZipcodeRepository        func() (repo.GetAddressByZipcodeResponse, error)
	MockGetDistrictByProvinceNameRepository  func(provinceName string) ([]repo.GetDistrictByProvinceNameResult, error)
	MockGetSubDistrictByDistrictIdRepository func(districtId string) ([]repo.GetSubDistrictByDistrictIdResult, error)
	MockGetDistrictByProvinceIdRepository    func(provinceId string) ([]repo.GetDistrictByProvinceIdResult, error)
	MockGetProvinceByProvinceIdRepository    func(provinceId string) (repo.GetProvinceByProvinceIdResult, error)
}

// กำหนด init ของ struct
func NewMockAddressMasterDataRepository() *MockAddressMasterDataRepository {
	return &MockAddressMasterDataRepository{
		MockGetAllZipcodeRepository: func() ([]int32, error) {
			return []int32{10000, 10001, 10002}, nil
		},
		MockGetAllProvinceRepository: func() ([]repo.ReadAllProvinceResp, error) {
			return []repo.ReadAllProvinceResp{
				{
					ProvinceId:   1,
					ProvinceName: "Mock Province",
				},
			}, nil
		},
		MockGetAddressByZipcodeRepository: func() (repo.GetAddressByZipcodeResponse, error) {
			return repo.GetAddressByZipcodeResponse{
				SubDistrictItems: []repo.SubDistrictItem{
					{
						SubDistrictId:   "1",
						SubDistrictName: "Mock SubDistrict",
					},
				},
				Districtid:   1,
				Districtname: "Mock District",
				Provinceid:   1,
				Provincename: "Mock Province",
			}, nil
		},
		MockGetDistrictByProvinceNameRepository: func(provinceName string) ([]repo.GetDistrictByProvinceNameResult, error) {
			if provinceName == "Mock" {
				return []repo.GetDistrictByProvinceNameResult{
					{
						Districtid:   1,
						Districtname: "Mock District",
					},
				}, nil
			}
			return []repo.GetDistrictByProvinceNameResult{}, nil
		},
		MockGetSubDistrictByDistrictIdRepository: func(districtId string) ([]repo.GetSubDistrictByDistrictIdResult, error) {
			if districtId == "1" {
				return []repo.GetSubDistrictByDistrictIdResult{
					{
						Subdistrictid:   "1",
						Subdistrictname: "Mock Subdistrict",
					},
				}, nil
			}
			return []repo.GetSubDistrictByDistrictIdResult{}, nil
		},
		MockGetDistrictByProvinceIdRepository: func(provinceId string) ([]repo.GetDistrictByProvinceIdResult, error) {
			if provinceId == "1" {
				return []repo.GetDistrictByProvinceIdResult{
					{
						Districtid:   1,
						Districtname: "Mock District",
					},
				}, nil
			}
			return []repo.GetDistrictByProvinceIdResult{}, nil
		},
		MockGetProvinceByProvinceIdRepository: func(provinceId string) (repo.GetProvinceByProvinceIdResult, error) {
			if provinceId == "1" {
				return repo.GetProvinceByProvinceIdResult{
					Provinceid:   1,
					Provincename: "Mock Province",
				}, nil
			}
			return repo.GetProvinceByProvinceIdResult{}, nil
		},
	}
}

// override
func (m *MockAddressMasterDataRepository) GetAllZipcodeRepository() ([]int32, error) {
	return m.MockGetAllZipcodeRepository()
}

func (m *MockAddressMasterDataRepository) GetAllProvinceRepository() ([]repo.ReadAllProvinceResp, error) {
	return m.MockGetAllProvinceRepository()
}

func (m *MockAddressMasterDataRepository) GetAddressByZipcodeRepository(zipcode string) (repo.GetAddressByZipcodeResponse, error) {
	return m.MockGetAddressByZipcodeRepository()
}

func (m *MockAddressMasterDataRepository) GetDistrictByProvinceNameRepository(provinceName string) ([]repo.GetDistrictByProvinceNameResult, error) {
	return m.MockGetDistrictByProvinceNameRepository(provinceName)
}

func (m *MockAddressMasterDataRepository) GetSubDistrictByDistrictIdRepository(districtId string) ([]repo.GetSubDistrictByDistrictIdResult, error) {
	return m.MockGetSubDistrictByDistrictIdRepository(districtId)
}

func (m *MockAddressMasterDataRepository) GetDistrictByProvinceIdRepository(provinceId string) ([]repo.GetDistrictByProvinceIdResult, error) {
	return m.MockGetDistrictByProvinceIdRepository(provinceId)
}

func (m *MockAddressMasterDataRepository) GetProvinceByProvinceIdRepository(provinceId string) (repo.GetProvinceByProvinceIdResult, error) {
	return m.MockGetProvinceByProvinceIdRepository(provinceId)
}
