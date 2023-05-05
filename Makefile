sqlc generate:
	docker run --rm -v "$(pwd):/src" -w /src kjconroy/sqlc:1.17.0 /workspace/sqlc generate