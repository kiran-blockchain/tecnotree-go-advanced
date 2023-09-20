package services

import (
	"github.com/kiran-blockchain/ekart/entities"
	"github.com/kiran-blockchain/ekart/interfaces"
)

type UserService struct {
}

func InitUserService() interfaces.IUser {
	return &UserService{}
}
func (u *UserService) Login(user *entities.User) string {
	return "user loggedin"
}

func (u *UserService) Register(user *entities.User) string {
	return "user registed"
}
func (u *UserService) GetProfile(userId int) *entities.User {
	return &entities.User{Name: "kiran",
		UserId:   1,
		Password: "kiran",
		Email:    "Kiranthecoder@gmai.come"}
}
func (u *UserService) SearchUser(query string) {

}
func (u *UserService) Logout(userid int) string {
	return "user loggedout"
}
func (u *UserService) ListProfiles(userid int) string {
	return "user loggedout"
}
