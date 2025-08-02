dev:
	@go run cmd/cli/main.go --env dev --port 8000

migrate-up:
	@migrate --path internal/infrastracture/database/migrations --database 'mysql://docker:password@tcp(127.0.0.1:3306)/vertical_slice_dev' -verbose up

migrate-down:
	@migrate --path internal/infrastracture/database/migrations --database 'mysql://docker:password@tcp(127.0.0.1:3306)/vertical_slice_dev' -verbose down