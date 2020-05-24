package GORM_Relate

import "github.com/jinzhu/gorm"

// User 包含一个 CreditCard, UserID 为外键
type User struct {
	gorm.Model
	CreditCard CreditCard
}

type CreditCard struct {
	gorm.Model
	UserID uint
	Number string
}

var card CreditCard
//db.Model(&user).Related(&card, "CreditCard")
//// SELECT * FROM credit_cards WHERE user_id = 123; // 123 is user's primary key
// CreditCard是user的字段名称，这意味着获得user的CreditCard关系并将其填充到变量
// 如果字段名与变量的类型名相同，如上例所示，可以省略，如：
//db.Model(&user).Related(&card)