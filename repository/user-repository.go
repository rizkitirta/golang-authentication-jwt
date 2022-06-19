package repository

import (
	"golang-authentication-jwt/entity"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	Insert(user entity.User) (entity.User, error)
	Update(user entity.User) (entity.User, error)
	Delete(user entity.User) (entity.User, error)
	FindByID(id string) (entity.User, error)
	FindByEmail(email string) (entity.User, error)
	FindAll() ([]entity.User, error)
	VerifyCredentials(email string, password string) interface{}
	IsDuplicateEmail(email string) (tx *gorm.DB, err error)
}

type userConnection struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{DB: db}
}

func (r *userConnection) Insert(user entity.User) (entity.User, error) {
	user.Password = hashAndSalt([]byte(user.Password))
	err := r.DB.Create(&user).Error
	return user, err
}
func (r *userConnection) Update(user entity.User) (entity.User, error) {
	if user.Password != "" {
		user.Password = hashAndSalt([]byte(user.Password))
	}
	err := r.DB.Save(&user).Error
	return user, err
}
func (r *userConnection) Delete(user entity.User) (entity.User, error) {
	err := r.DB.Create(&user).Error
	return user, err
}
func (r *userConnection) FindByID(id string) (entity.User, error) {
	 var user entity.User
	 res := r.DB.Find(&user, id)
	 if res.Error != nil {
		 return user, res.Error
	 }
	 return user, nil
}
func (r *userConnection) FindByEmail(email string) (entity.User, error) {
	 var user entity.User
	 res := r.DB.Where("email = ?", email).Take(&user)
	 if res.Error != nil {
		 return user, res.Error
	 }
	 return user, nil
}
func (r *userConnection) FindAll( ) ([]entity.User, error) {
	 var users []entity.User
	 res := r.DB.Find(&users)
	 if res.Error != nil {
		 return users, res.Error
	 }
	 return users, nil
}
func (r *userConnection) VerifyCredentials(email string,password string) interface{} {
	 var user entity.User 
	 result := r.DB.Where("email = ?", email).Take(&user)
	 if result.Error != nil {
		 return nil
	 }

	 return user

}
func (r *userConnection) IsDuplicateEmail(email string) (tx *gorm.DB, err error) {
	var user entity.User
	return r.DB.Where("email = ?", email).Take(&user), nil
}
	 
func hashAndSalt(pwd []byte) string  {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		panic(err)
	}

	return string(hash)
}