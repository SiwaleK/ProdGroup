package mocks

import (
	"testing"

	"example.com/go-crud-api/db/database"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// function ไว้ setup ตัว mock database
func Setup(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	// สร้าง Mock Database ใหม่
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	//connect Mock Database แล้วก็กำหนด postgres.config ที่สร้างมาใช้ Driver อะไร ในที่นี้ก็ใช้ postgres
	//DSN เป็นอะไร(ไม่ค่อยจำเป็นเท่าไหร่ เพราะเรา Mock ตัว Database ไม่ใช่ของจริง)
	//SimpleProtocol -> ส่ง plain text ให้ server โดยตรง, สามารถ set เป็น extended protocol ที่จะส่งเป็น binary แทนได้ (หาอ่านเพิ่มได้นะครับ)
	//แล้วก็บอก gorm ให้ connect sqlmock.New() ที่พึ่งสร้างไป
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		DriverName:           "sqlmock",
		DSN:                  "",
		PreferSimpleProtocol: true,
		Conn:                 db,
	}), &gorm.Config{})

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database connection", err)
	}

	database.DB = gormDB
	return gormDB, mock
}
