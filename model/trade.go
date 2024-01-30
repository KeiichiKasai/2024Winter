package model

type Wallet struct {
	Username string  `form:"username,omitempty" gorm:"not null"`
	Balance  float64 `form:"balance,omitempty" gorm:"default :0"`
}
type Cart struct {
	Username string  `form:"username,omitempty" gorm:"not null"`
	Gid      int     `form:"gid,omitempty" gorm:"not null"`
	Gname    string  `form:"gname,omitempty" gorm:"not null"`
	Price    float64 `form:"price,omitempty" gorm:"not null"`
	Count    int     `form:"count,omitempty" gorm:"not null"`
}
