package constant

const (
	ErrInvalidRequest = "invalid request"
	ErrGeneral        = "general error"
)

const (
	ErrApprovalIdNil      = "request id nil"
	ErrApprovalIdNotFound = "request id not found"
)

const (
	ErrInvalidSortBy            = "invalid sort"
	ErrInvalidFilterBy          = "invalid filter"
	ErrFailedConvertStringToInt = "failed convert id to int"
	ErrTimeout                  = "timeout_error"
	ErrDatabase                 = "database_error"
	ErrWhenExecuteQueryDB       = "error_when_executing_query_db"
	ErrWhenScanResultDB         = "error_when_scanning_result_db"
	ErrInvalidApprovalStatus    = "invalid_approval_status"
	ErrRequestNotFound          = "request_not_found"
	ErrInvalidRequestBody       = "invalid_request_body"
)
