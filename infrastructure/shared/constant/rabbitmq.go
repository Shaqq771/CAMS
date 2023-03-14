package constant

const (
	CONSUMER_BILLING_RABBITMQ        = "consumer_billing_rabbitmq"
	CONSUMER_PRODUCT_INSERT_RABBITMQ = "product_insert"
	CONSUMER_PRODUCT_UPDATE_RABBITMQ = "product_update"
	PUBLISHER_RABBITMQ               = "publisher_rabbitmq"
)

const (
	RABBITMQ_PORT          = "5672"
	RABBITMQ_EXCHANGE_TYPE = "topic"
)

const (
	RETRY_MESSAGE_BROKER              = "Retrying connect to message broker..."
	SUCCESS_PUBLISH_TO_BROKER         = "success publish event to topic: %s "
	SUCCESS_CONSUME_FROM_BROKER       = "success consume event from topic: %s, body: %s "
	START_LISTENING_TOPIC_FROM_BROKER = "start listening event from topic: %s "
)
