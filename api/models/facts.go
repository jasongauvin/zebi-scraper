package models

type Fact struct {
    ID          int    `json:"id"`
    Description string `json:"description"`
	Author 		string `json:"author"`
}