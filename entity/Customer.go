package entity

type Customer struct {
	//gorm.Model
	Email     string `gorm:"column:email"`
	FirstName string `gorm:"column:first_name"`
	LastName  string `gorm:"column:last_name"`
	Avatar    string `gorm:"column:avatar"`
}
