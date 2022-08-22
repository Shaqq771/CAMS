run:
	@echo "Makefile running.."
mock:
	@echo "Generate Mock Interface.."
	mockgen -source="./domain/logistik/feature/feature.go" -destination="./domain/logistik/feature/mocks/feature_mock.go"
	mockgen -source="./domain/logistik/repository/repository.go" -destination="./domain/logistik/repository/mocks/repository_mock.go"
	mockgen -source="./infrastructure/broker/rabbitmq/rabbitmq.go" -destination="./infrastructure/broker/rabbitmq/mocks/rabbitmq_mock.go"