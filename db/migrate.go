package main

import "app/model"

func main() {
	db := model.DB()
	err := db.AutoMigrate(&model.Article{}, &model.Tag{}, &model.ArticleTag{})
	if err != nil {
		panic(err.Error())
	}

	// err := godotenv.Load("../.env")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// USER := os.Getenv("DB_USER")
	// PASS := os.Getenv("DB_PASS")
	// PROTOCOL := "tcp(db:3306)"
	// DBNAME := os.Getenv("DB")

	// fmt.Print(USER, PASS, PROTOCOL, DBNAME)

	// dsn := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8mb4&parseTime=true&loc=Local"
	// _, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	panic(err.Error())
	// }
	// fmt.Print("DB OK\n")
}
