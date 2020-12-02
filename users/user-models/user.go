package user_models

import (
	"errors"
	"github.com/badoux/checkmail"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"strings"
)

type User struct {
	gorm.Model
	Email    string `gorm:"type:varchar(100);unique_index" json:"email"`
	Name     string `gorm:"size:100;not null"              json:"name"`
	Password string `gorm:"size:100;not null"              json:"password"`
}

// HashPassword hashes password from user input
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14) // 14 is the cost for hashing the password.
	return string(bytes), err
}

// CheckPasswordHash checks password hash and password from user input if they match
func CheckPasswordHash(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return errors.New("password incorrect")
	}
	return nil
}

func (u *User) Prepare() {
	u.Email = strings.TrimSpace(u.Email)
	u.Name = strings.TrimSpace(u.Name)
}

//user validation

func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "login":
		if u.Email == "" {
			return errors.New("Email is required")
		}
		if u.Password == "" {
			return errors.New("Password is required")
		}
		return nil
	default:
		if u.Name == "" {
			return errors.New("Name is required")
		}
		if u.Email == "" {
			return errors.New("Email is required")
		}
		if u.Password == "" {
			return errors.New("Password is required")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil

	}
}

// BeforeSave hashes user password
func (u *User) BeforeSave() error {
	password := strings.TrimSpace(u.Password)
	hashedpassword, err := HashPassword(password)
	if err != nil {
		return err
	}
	u.Password = string(hashedpassword)
	return nil
}

func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	var err error

	// Debug a single operation, show detailed log for this operation
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

// GetUser returns a user based on email
func (u *User) GetUser(db *gorm.DB) (*User, error) {
	account := &User{}
	if err := db.Debug().Table("users").Where("email = ?", u.Email).First(account).Error; err != nil {
		return nil, err
	}
	return account, nil
}
func (u *User) GetUserID(id uint, db *gorm.DB) (*User, error) {
	account := &User{}
	if err := db.Debug().Table("users").Where("id = ?", id).First(account).Error; err != nil {
		return nil, err
	}
	return account, nil
}

// GetAllUsers returns a list of all the user
func GetAllUsers(db *gorm.DB) (*[]User, error) {
	users := []User{}
	if err := db.Debug().Table("users").Find(&users).Error; err != nil {
		return &[]User{}, err
	}
	return &users, nil
}
