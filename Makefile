run:
    go run cmd/main.go

sqlc:
    docker run --rm -v "$(pwd):/src" -w /src kjconroy/sqlc:1.17.0 /workspace/sqlc generate
    sed -i -e 's/sql.NullString/*string/g' db/sqlc/models.go
    sed -i -e 's/sql.NullInt32/*int32/g' db/sqlc/models.go
    sed -i -e 's/sql.NullInt16/*int16/g' db/sqlc/models.go
    sed -i -e 's/sql.NullTime/*time.Time/g' db/sqlc/models.go
    sed -i -e 's/"database\/sql"//' db/sqlc/models.go


sqlc manual for models.go file :
    docker run --rm -v "$(pwd):/src" -w /src kjconroy/sqlc:1.17.0 /workspace/sqlc generate
    (Get-Content -Path "db/sqlc/models.go") -replace "sql.NullString", "*string" | Set-Content -Path "db/sqlc/models.go"
    (Get-Content -Path "db/sqlc/models.go") -replace "sql.NullInt32", "*int32" | Set-Content -Path "db/sqlc/models.go"
    (Get-Content -Path "db/sqlc/models.go") -replace "sql.NullInt16", "*int16" | Set-Content -Path "db/sqlc/models.go"
    (Get-Content -Path "db/sqlc/models.go") -replace "sql.NullTime", "*time.Time" | Set-Content -Path "db/sqlc/models.go"
    (Get-Content -Path "db/sqlc/models.go") -replace "`"database/sql`"", "" | Set-Content -Path "db/sqlc/models.go"

docker run --rm -v "$(pwd):/src" -w /src kjconroy/sqlc:1.17.0 /workspace/sqlc generate


.PHONY: sqlc run
  (Get-Content -Path "db/sqlc/sku.sql.go") -replace "sql.NullString", "*string" | Set-Content -Path "db/sqlc/sku.sql.go"
  (Get-Content -Path "db/sqlc/querier.go") -replace "sql.NullString", "*string" | Set-Content -Path "db/sqlc/querier.go"
 (Get-Content -Path "db/sqlc/querier.go") -replace "`"database/sql`"", "" | Set-Content -Path "db/sqlc/querier.go"
 (Get-Content -Path "db/sqlc/sku.sql.go") -replace "`"database/sql`"", "" | Set-Content -Path "db/sqlc/sku.sql.go"
 (Get-Content -Path "db/sqlc/models.go") -replace "`"database/sql`"", "" | Set-Content -Path "db/sqlc/models.go"

Mockgen:
    mockgen -source=C:/Users/banas/internship/ProdGroup/sku/repository/promotionRepo.go -destination=repository/Mocks/promotionMockGen.go -package=repository github.com/SiwaleK/ProdGroup PromotionRepository  

    mockgen -destination repository/Mocks/promotionMock.go github.com/SiwaleK/ProdGroup/repository PromotionRepository