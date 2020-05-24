package GORM_Relate

import "github.com/jinzhu/gorm"

// User 包含并属于多个 languages, 使用 `user_languages` 表连接
type User struct {
	gorm.Model
	Languages         []Language `gorm:"many2many:user_languages;"`
}

type Language struct {
	gorm.Model
	Name string
}

//db.Model(&user).Related(&languages, "Languages")
//// SELECT * FROM "languages" INNER JOIN "user_languages" ON "user_languages"."language_id" = "languages"."id" WHERE "user_languages"."user_id" = 111