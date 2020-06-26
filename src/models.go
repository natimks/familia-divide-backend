package src

import (
	"time"
)

type user struct {
	ID          int       `gorm:"primary_key:auto_increment" json:"id"`
	CPF         string    `gorm:"type:varchar(11);not null;unique_index" json:"cpf"`
	Name        string    `gorm:"not null" json:"name"`
	Email       string    `gorm:"type:varchar(100);not null;unique_index" json:"email"`
	Phone       string    `gorm:"type:varchar(20);not null;unique_index" json:"phone"`
	UsersFamily []*user   `gorm:"many2many:users_family;association_jointable_foreignkey:admin_user_id" json:"users_family"`
	Incomes     []income  `gorm:"foreignkey:user_id;" json:"incomes"`
	Expenses    []expense `gorm:"foreignkey:user_id;" json:"expenses"`
	CreatedAt   time.Time `gorm:"default:NOW()" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:NOW()" json:"updated_at"`
}

func (user) TableName() string {
	return "users"
}

type income struct {
	ID          int       `gorm:"primary_key:auto_increment" json:"id"`
	Description string    `gorm:"type:varchar(50);not null;" json:"description"`
	Value       float32   `gorm:"not null;" sql:"type:decimal(10,2)" json:"value"`
	UserId      int       `gorm:"not null" json:"user_id"`
	CreatedAt   time.Time `gorm:"default:NOW()" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:NOW()" json:"updated_at"`
}

func (income) TableName() string {
	return "incomes"
}

type expense struct {
	ID          int       `gorm:"primary_key:auto_increment" json:"id"`
	Description string    `gorm:"type:varchar(50);not null;" json:"description"`
	Value       float32   `gorm:"not null;" sql:"type:decimal(10,2)" json:"value"`
	UserId      int       `gorm:"not null" json:"user_id"`
	CreatedAt   time.Time `gorm:"default:NOW()" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:NOW()" json:"updated_at"`
}

func (expense) TableName() string {
	return "expenses"
}
