package model

type Fields struct {
	FieldName string
	Option    string
	DataType  string
	FromValue interface{}
	ToValue   interface{}
}

type Filter struct {
	Filters []Fields
	Limit   int
	Page    int
}
