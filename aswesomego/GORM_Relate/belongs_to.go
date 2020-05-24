package GORM_Relate

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	//Profile      Profile `gorm:"ForeignKey:ProfileRefer"` // // 使用ProfileRefer作为外键
	Profile   Profile `gorm:"ForeignKey:ProfileID;AssociationForeignKey:Refer"` // // 使用ProfileRefer作为外键
	ProfileID int
}

type Profile struct {
	gorm.Model
	Name string
}

//db.Model(&User).Related(&profile)
//// SELECT * FROM profiles WHERE id = 111; // 111是user的外键ProfileID
