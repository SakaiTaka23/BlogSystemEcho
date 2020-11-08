package main

import "github.com/SakaiTaka23/model"

func main() {
	db := model.DB()
	err := db.AutoMigrate(&model.Article{}, &model.Tag{}, &model.ArticleTag{})
	if err != nil {
		panic(err.Error())
	}

	// err := godotenv.Load(fmt.Sprintf("../%s.env", os.Getenv("GO_ENV")))
	// if err != nil {
	// 	panic(err.Error())
	// }
	// USER := os.Getenv("DB_USER")
	// PASS := os.Getenv("DB_PASS")
	// DBNAME := os.Getenv("DB")

	// fmt.Print(USER, PASS, DBNAME)

	// dsn := USER + ":" + PASS + "@" + "tcp(127.0.0.1:3306)" + "/" + DBNAME + "?charset=utf8mb4&parseTime=true&loc=Local"
	// _, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	panic(err.Error())
	// }
	// fmt.Print("DB OK\n")
}
