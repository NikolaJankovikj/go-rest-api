package models

type Game struct {
	ID        string `json:"id" bson:"_id"`
	Title     string `json:"title" bson:"title"`
	Genre     string `json:"genre" bson:"genre"`
	Developer string `json:"developer" bson:"developer"`
}
