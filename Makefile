init:
	@echo "Running mod vendor"
	@go mod vendor
	@echo "Installing air"
	@go install github.com/cosmtrek/air@latest
	@echo "Installing gomock"
	@go install github.com/golang/mock/mockgen@v1.6.0
	@echo "Init finish!"

deps-up:
	@sudo docker-compose up

deps-down:
	@sudo docker-compose down

run:
	@air

delete-db-data:
	@sudo chmod -R 777 .dev/
	@rm -rf .dev/dbdata

testcover:
	@go test -v -coverprofile cover.out ./...
	@go tool cover -func cover.out