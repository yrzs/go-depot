package model

type ArticleTag struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedAt  uint32 `json:"created_at"`
	UpdatedAt  uint32 `json:"updated_at"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
	TagID      uint32 `json:"tag_id"`
	ArticleID  uint32 `json:"article_id"`
}

func (a ArticleTag) TableName() string {
	return "blog_article_tag"
}
