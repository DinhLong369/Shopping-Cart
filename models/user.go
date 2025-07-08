package models

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	Id        uint       `json:"user_id" gorm:"auto-increment; primary key"`
	Name      string     `json:"name"`
	Email     string     `json:"email" gorm:"unique"`
	Password  string     `json:"-"`
	CartItems []CartItem `gorm:"foreignKey:UserID"`
}
type RequestSignUp struct {
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}
