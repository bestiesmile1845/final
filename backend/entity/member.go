package entity 

import ("gorm.io/gorm")

type Member struct {
	gorm.Model
	Name string `valid:"required~Name is not required"`
	Age int `valid:"range(0|120)~ Age must betweed 0 to 120"`
	MemberID string `vaild:"matches(^(C)(MC)\\d{8}$)~StudentID must start with 'CM' and have 8 digits"`
}