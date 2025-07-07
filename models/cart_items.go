package models

type CartItem struct {
	Id        uint    `json:"id" gorm:"primary key;auto-increment"`
	Price     float64 `json:"price" gorm:"type:double;not null;default:0"`
	Quantity  uint    `json:"quantity" gorm:"type:uint;not null;default:0"`
	UserID    uint    `json:"user_id"`
	ProductID uint    `json:"product_id"`
	Product   Product `json:"-" gorm:"foreignKey:ProductID"`
	User      User    `json:"-" gorm:"foreignKey:UserID"`
}

type Request struct {
	Quantity  uint `json:"quantity"`
	ProductID uint `json:"product_id"`
}

type ListCart struct {
	IDCart     uint       `json:"id_cart"`
	Items      []CartItem `json:"items"`
	TotalPrice float64    `json:"total_price"`
}

type IdsProduct struct {
	IdsProduct []uint `json:"ids_product"`
}
