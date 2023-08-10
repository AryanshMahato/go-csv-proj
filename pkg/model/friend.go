package model

type Friend struct {
	ID   int    `json:"id" csv:"id"`
	Name string `json:"name" csv:"name"`
}
