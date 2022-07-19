package models

type User struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImgSrc      string `json:"img-src"`
	Time        string `json:"time"`
	Read        string `json:"read"`
	Url         string `json:"url"`
}
