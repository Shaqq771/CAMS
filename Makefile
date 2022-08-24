run_local:
	@echo "Run apps.."
	copy .env.example .env
	go build .
	go run .
run_docker:
	@echo "Run docker.."
	docker network create nabati || @echo "network already exist! skip.."
	docker-compose build --no-cache
	docker-compose up -d
stop_docker:
	@echo "Stop docker.."
	docker-compose down
mock:
	@echo "Generate Mock Interface.."
	mockgen -source="./domain/logistik/feature/feature.go" -destination="./domain/logistik/feature/mocks/feature_mock.go"
	mockgen -source="./domain/logistik/repository/repository.go" -destination="./domain/logistik/repository/mocks/repository_mock.go"
	mockgen -source="./infrastructure/broker/rabbitmq/rabbitmq.go" -destination="./infrastructure/broker/rabbitmq/mocks/rabbitmq_mock.go"
test:
	@echo "Do Testing.."
	make mock
	go test -cover ./domain/logistik/feature