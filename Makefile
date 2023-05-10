sqlc generate:
	docker run --rm -v "$(pwd):/src" -w /src kjconroy/sqlc:1.17.0 /workspace/sqlc generate

sqlc:
    docker run --rm -v "$(pwd):/src" -w /src kjconroy/sqlc:1.17.0 /workspace/sqlc generate
    sed -i -e 's/sql.NullString/*string/g' db/sqlc/*.go
    sed -i -e 's/sql.NullInt32/*int32/g' db/sqlc/*.go
    sed -i -e 's/sql.NullTime/*time.Time/g' db/sqlc/*.go
    goimports -w ./db/sqlc/

.PHONY: sqlc