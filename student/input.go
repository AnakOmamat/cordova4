package student

import "cordova4/user"

type GetStudentDetailInput struct {
	ID int `uri:"id" binding:"required"`
}

type RegisterStudentInput struct {
	Name string `json:"name" binding:"required"`
	User user.User
}

type FormUpdateStudentInput struct {
	ID    int    `uri:"id" binding:"required"`
	Name  string `form:"name" binding:"required"`
	Error error
}
