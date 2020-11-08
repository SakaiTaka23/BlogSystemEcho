package model

type ArticleTag struct {
	ID         uint `gorm:"primaryKey"`
	article_id uint
	tag_id     uint
}
