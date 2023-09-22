package interfaces

import "github.com/kiran-blockchain/ekart/entities"

type IUser interface {
	Register(user *entities.User) (*entities.SignupResponse,error)
	//Login(user *entities.User) string
	//  GetProfile(userId int) (*entities.User)
	//  SearchUser(searchQuery string)
	//  Logout(userId int) (string)
}
