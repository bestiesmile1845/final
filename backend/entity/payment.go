package entity

import(
	"time"
	"gorm.io/gorm")

type Payment struct {
		gorm.Model
		PayerName string    `valid:"required~Payer name is required"`
		Date      time.Time `valid:"required~Date is required"`
		Amount    int       `valid:"required~Amount is required"`
}
	


