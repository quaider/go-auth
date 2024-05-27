package model

type User struct {
	Name     string  `json:"name,omitempty" gorm:"column:name"`
	Password string  `json:"password,omitempty" gorm:"column:password"`
	Nickname *string `json:"nickname" gorm:"column:nickname"`
}

type UserList struct {
	Items []*User `json:"items"`
}
