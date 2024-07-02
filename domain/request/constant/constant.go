package constant

type ContextKey string

const (
	FiberContext ContextKey = "fiberCtx"
)

const (
	DefaultLimitPerPage = 5
	DefaultPage         = 1
)

const (
	PAGE    = "page"
	LIMIT   = "limit"
	SORT_BY = "sort_by"
	SEARCH  = "search"
	QUERY   = "query"
)

const (
	DefaultTimeout = 5 // detik
)

const (
	ApprovalStatusWaiting  = "Waiting"
	ApprovalStatusApproved = "Approved"
	ApprovalStatusRevised  = "Revised"
	ApprovalStatusRejected = "Rejected"
)

const (
	// Request Status Constants
	StatusApproved = "Approved"
	StatusRejected = "Rejected"
	StatusRevised  = "Revised"
)