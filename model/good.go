package model

type Good struct {
	Gid      int     `form:"gid,omitempty" gorm:"primary_key"`
	Gname    string  `form:"gname,omitempty" gorm:"not null,index:idx_gname,unique"`
	Category string  `form:"category,omitempty" gorm:"not null"`
	Picture  string  `form:"picture,omitempty" gorm:"not null"`
	Price    float64 `form:"price,omitempty" gorm:"not null"`
	Count    int     `form:"count,omitempty" gorm:"default :0"`
	OwnerId  int     `form:"owner_id,omitempty" gorm:"not null"`
}
