package constant

const (
	ErrLoadENV = "error loading .env file: %w"

	ErrConvertStringToInt = "error when convert string to int: %w"
	ErrToMarshalJSON      = "failed json marshal: %w"
)

const (
	ErrConnectToDB             = "failed connect to db: %w"
	ErrInvalidJWTSigningMethod = "invalid jwt signing method"
	ErrInvalidJWTToken         = "invalid jwt token"
)

const (
	ErrInvalidRequest = "invalid request"
	ErrGeneral        = "general error"
)

// RabbitMQ
const (
	ErrConnectToBroker       = "failed connect to broker: %w"
	ErrCreateChannelToBroker = "failed create channel to broker: %w"
	ErrSetupQueueToBroker    = "failed setup queue to broker: %w"

	ErrDefineChannelToBroker  = "failed create channel to broker"
	ErrDeclareExhangeToBroker = "failed declare exchange to broker"
	ErrCreateTopicToBroker    = "failed create topic to broker"
	ErrCreateQueueToBroker    = "failed create queue to broker"
	ErrBindingQueueToBroker   = "failed binding queue to broker"
	ErrPublishQueueToBroker   = "failed publish queue to broker"
	ErrConsumeQueueToBroker   = "failed consume queue from broker"
)
