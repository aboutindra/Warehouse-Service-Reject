package controller

import "go.mongodb.org/mongo-driver/bson/primitive"

type ResBool struct {
	Res bool `json:"res" bson:"res"`
}

type ResInt struct {
	Res int64 `json:"res" bson:"res"`
}

type Data struct {
	Id     primitive.ObjectID `json:"_id" bson:"_id"`
	IdMate string             `json:"idMate" bson:"idMate"`
	Name   string             `json:"name, omitempty" bson:"name, omitempty"`
	Model  string             `json:"model, omitempty" bson:"model, omitempty"`
	Tipe   string             `json:"tipe, omitempty" bson:"tipe,omitempty"`
}

type ItemId struct {
	Data string `json:"idMate" bson:"idMate"`
}

type ItemName struct {
	Data string `json:"name, omitempty" bson:"name, omitempty"`
}

type ItemModel struct {
	Data string `json:"model, omitempty" bson:"model, omitempty"`
}

type ItemTipe struct {
	Data string `json:"tipe, omitempty" bson:"tipe,omitempty"`
}
