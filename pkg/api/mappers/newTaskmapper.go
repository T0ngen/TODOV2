package mappers


type NewTask struct{
	Date string `json:"date,omitempty"`
	Title string `json:"title" validate:"required"`
	Comment string `json:"comment,omitempty"`
	Repeat string `json:"repeat,omitempty"`
}