package models

type Blog struct{
	Id 			int		`json : "id"`
	Title 		string 	`json : "title"`
	Desc		string	`json : "desc"`
	Image		string	`json : "image"`
	UserId		int		`json : "userid"`
	User		User	`json:"user";gorm:"foreignKey:UserId"`
	
}