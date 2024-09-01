package db

import (
	"github.com/PratikVarute/BookMyStay_Go_Fiber/types"
	"golang.org/x/crypto/bcrypt"
)

// encrypt passowrd of new user
func NewUserFromPrams(params types.CreateUserParams) (*types.User, error) {
	encryptcparm, err := bcrypt.GenerateFromPassword([]byte(params.Password), 12)
	if err != nil {
		return nil, err
	}
	return &types.User{
		FristName:         params.FristName,
		Lastname:          params.Lastname,
		Email:             params.Email,
		EncryptedPassword: string(encryptcparm),
	}, nil
}
