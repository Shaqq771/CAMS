package constant

const (
	ErrTimeout = "timeout"
)

// Error Type
const (
	ErrDatabase       = "error database"
	ErrInvalidRequest = "invalid request"
	ErrGeneral        = "general error"
	ErrAuth           = "unathorized"
	Err               = "error"
)

// Error type ErrDatabase
const (
	ErrWhenExecuteQueryDB     = "error when execute query: "
	ErrWhenScanResultDB       = "error when scan result: "
	ErrWhenPrepareStatementDB = "error when prepare statement query: "
	ErrWhenCommitDB           = "error when commit query: "
	ErrWhenBeginTX            = "error when begin transaction: "
	ErrWhenSelectDB           = "error when select data from db: "
	ErrWhenUpdateDB           = "error when update data to db: "
	ErrWhenRollBackDataToDB   = "error when rollback data from db: "
	ErrRollBack               = "rollback data from db: %s"
)

// Error type ErrDatabase
const (
	ErrCreateSortQuery = "faild create sort query"
)
