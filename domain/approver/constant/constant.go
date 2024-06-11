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
)

const (
	DefaultTimeout = 5 // detik
)
