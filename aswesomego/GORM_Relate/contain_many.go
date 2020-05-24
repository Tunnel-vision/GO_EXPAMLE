package GORM_Relate

import "github.com/jinzhu/gorm"

// User 包含多个 emails, UserID 为外键
type User struct {
	gorm.Model
	Emails []Email
}

type Email struct {
	gorm.Model
	Email  string
	UserID uint
}

//db.Model(&user).Related(&emails)
//// SELECT * FROM emails WHERE user_id = 111; // 111 是 user 的主键

type Profile struct {
	gorm.Model
	Name   string
	UserID uint
}

type User struct {
	Refer string
	gorm.Model
	Profiles []Profile `gorm:"ForeignKey:UserID;AssociationForeignKey:Refer"`
}
