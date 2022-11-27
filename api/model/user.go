package model

type UserBase struct {
	Name  string
	Email string
}
type UserPostRequest struct {
	UserBase
}
