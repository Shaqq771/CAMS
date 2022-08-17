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
	ErrConnectToBroker       = "failed connect to broker: %w"
	ErrCreateChannelToBroker = "failed create channel to broker: %w"
	ErrCreateTopicToBroker   = "failed create topic to broker: %w"
	ErrSetupQueueToBroker    = "failed setup queue to broker: %w"
	ErrCreateQueueToBroker   = "failed create queue to broker: %w"
	ErrBindingQueueToBroker  = "failed binding queue to broker: %w"
	ErrPublishQueueToBroker  = "failed publish queue to broker: %w"
	ErrConsumeQueueToBroker  = "failed consume queue to broker: %w"
)

const (
	ErrInvalidRequest = "invalid request"
	ErrGeneral        = "general error"
)
