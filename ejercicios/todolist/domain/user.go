package domain

type User struct {
	ID int `json:"id"`
	Name string `json:"name"`
	CreatedAt string `json:"created_at"`
	Active bool `default:"true" json:"active"`
}