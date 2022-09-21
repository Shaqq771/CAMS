package model

type NumberRange struct {
	DocType    string `db:"doc_type"`
	PlantId    string `db:"plant_id"`
	FromNumber string `db:"from_number"`
	ToNumber   string `db:"to_number"`
	LastNumber string `db:"last_number"`
	SkipNumber string `db:"skip"`
}
