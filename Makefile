init:
	@echo "Running mod vendor"
	@go mod vendor
	@echo "Installing air"
	@go install github.com/cosmtrek/air@latest
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