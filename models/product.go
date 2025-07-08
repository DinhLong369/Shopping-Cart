package models

type Product struct {
	ID          uint       `json:"id" gorm:"auto-increment; primary key"`
	Name        string     `json:"name" gorm:"type:varchar(255);not null;default:'sp'"`
	Description string     `json:"description" gorm:"type:text"`
	Price       float64    `json:"price" gorm:"type:double;not null;default:0"`
	Quantity    uint       `json:"quantity" gorm:"type:uint;not null;default:0"`
	ImageURL    string     `json:"image_url" gorm:"type:varchar(512)"`
	CartItems   []CartItem `json:"-" gorm:"foreignKey:ProductID"`
}

type CreateProduct struct {
	Name        string  `json:"name" gorm:"type:varchar(255);not null;default:'sp'"`
	Description string  `json:"description" gorm:"type:text"`
	Price       float64 `json:"price" gorm:"type:double;not null;default:0"`
	Quantity    uint    `json:"quantity" gorm:"not null;default:0"`
	ImageURL    string  `json:"image_url" gorm:"type:varchar(512)"`
}

type Ids struct {
	Ids []uint `json:"ids"`
}
