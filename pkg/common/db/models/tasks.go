package models



type Task struct{
	Id string `json:"id"`
	Date string `json:"date"`
	Title string `json:"title" validate:"required"`
	Comment string `json:"comment"`
	Repeat string `json:"repeat"`
}