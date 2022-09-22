package model

type Fields struct {
	FieldName string      `json:"field"`
	Option    string      `json:"option"`
	DataType  string      `json:"data_type"`
	FromValue interface{} `json:"from_value"`
	ToValue   interface{} `json:"to_value"`
}

type Filter struct {
	Filters []Fields `json:"filters"`
	Limit   int      `json:"limit"`
	Page    int      `json:"page"`
}
