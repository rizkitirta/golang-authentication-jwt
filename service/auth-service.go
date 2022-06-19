package service

import (
	"golang-authentication-jwt/dto"
	"golang-authentication-jwt/entity"
	"golang-authentication-jwt/repository"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	VerifyCredentials(email string, password string) interface{}
	CreateUser(user dto.UserCreateDTO) (entity.User, error)
	FindUserByEmail(email string) (entity.User, error)
	IsDuplicateEmail(email string) bool
}

type authService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) AuthService {
	return &authService{userRepository: userRepository}
}

func (s *authService) VerifyCredentials(email string, password string) interface{} {
	res := s.userRepository.VerifyCredentials(email, password)
	if v,ok := res.(entity.User); ok {
		comparePassword := ComparePassword(v.Password, []byte(password)) 
		if v.Email == email && comparePassword {
			return res
		}
		return false
	}

	return false
}

func (s *authService) CreateUser(userReq dto.UserCreateDTO) (entity.User, error) {
	user := entity.User{
		Email: userReq.Email,
		Password: userReq.Password,
		Name: userReq.Name,
	}

	res,err := s.userRepository.Insert(user)
	if err != nil {
		return entity.User{}, err
	}
	return res, nil

}
func (s *authService) FindUserByEmail(email string) (entity.User, error) {
	return s.userRepository.FindByEmail(email)
}
func (s *authService) IsDuplicateEmail(email string) bool {
	res,err := s.userRepository.IsDuplicateEmail(email)
	if err != nil {
		return false
	}
	if res.RowsAffected > 0 {
		return true
	}
}

func ComparePassword(hashedPwd string, plainPwd []byte) bool {
	byteHasdPwd := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHasdPwd, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
