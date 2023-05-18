include .env

DB_CONNECTION_STRING := postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(SSL_MODE)

migrateup:
	migrate -path ./db/migration -database "$(DB_CONNECTION_STRING)" up

migratedown:
	migrate -path ./db/migration -database "$(DB_CONNECTION_STRING)" down

sqlc:
	sqlc generate
	perl -pi -e 's/sql\.NullString/*string/g' db/sqlc/models.go
	perl -pi -e 's/sql\.NullInt32/*int32/g' db/sqlc/models.go
	perl -pi -e 's/sql\.NullInt16/*int16/g' db/sqlc/models.go
	perl -pi -e 's/sql\.NullTime/*time.Time/g' db/sqlc/models.go
	perl -pi -e 's/sql\.NullInt64/int64/g' db/sqlc/models.go
	perl -pi -e 's/sql\.NullFloat64/float64/g' db/sqlc/models.go
	perl -pi -e 's/uuid\.NullUUID/uuid.UUID/g' db/sqlc/models.go
	perl -pi -e 's/pqtype\.NullRawMessage/json.RawMessage/g' db/sqlc/models.go
	perl -pi -e 's/"database\/sql"//g' db/sqlc/models.go
	perl -pi -e 's/"github\.com\/tabbed\/pqtype"//g' db/sqlc/models.go
	perl -pi -e 's/sql\.NullBool/*bool/g' db/sqlc/models.go

run:
	go run cmd/main.go

test:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

gen:
	gentool -db "postgres" -dsn "$(DB_CONNECTION_STRING)" -outPath "./model/db" -modelPkgName "db" -onlyModel

.PHONY: migrateup migratedown sqlc test gen
