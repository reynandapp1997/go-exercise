package models

type Student struct {
	ID    string `bson:"_id,omitempty" json:"_id,omitempty"`
	Name  string `bson:"name" json:"name"`
	Grade int    `bson:"grade" json:"grade"`
}
