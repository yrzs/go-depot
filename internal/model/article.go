package model

type Article struct {
	ID            uint32 `gorm:"primary_key" json:"id"`
	CreatedBy     string `json:"created_by"`
	ModifiedBy    string `json:"modified_by"`
	CreatedAt     uint32 `json:"created_at"`
	UpdatedAt     uint32 `json:"updated_at"`
	DeletedOn     uint32 `json:"deleted_on"`
	IsDel         uint8  `json:"is_del"`
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"`
}

func (a Article) TableName() string {
	return "blog_article"
}
