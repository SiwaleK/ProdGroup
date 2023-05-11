# migrateup:
#     migrate -path ./db/migration -database "postgresql://root:secret@127.0.0.1:5433?sslmode=disable" up

# migratedown:
#     migrate -path ./db/migration -database "postgresql://root:secret@127.0.0.1:5433?sslmode=disable" down

run:
    go run cmd/main.go 

sqlc:
    docker run --rm -v "$(pwd):/src" -w /src kjconroy/sqlc:1.17.0 /workspace/sqlc generate
    sed -i -e 's/sql.NullString/*string/g' db/sqlc/models.go
    sed -i -e 's/sql.NullInt32/*int32/g' db/sqlc/models.go
    sed -i -e 's/sql.NullInt16/*int16/g' db/sqlc/models.go
    sed -i -e 's/sql.NullTime/*time.Time/g' db/sqlc/models.go
    sed -i -e 's/"database\/sql"//' db/sqlc/models.go


test:
	cd test
	go test

# .PHONY: migrateup migratedown sqlc run
