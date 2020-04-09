package model

import "time"


// 定义术语题的数据模型
type Term struct {
	ID        int64      `gorm:"primary_key" form:"-" json:"id"`
	Answer    string     `gorm:"type:varchar(200);unique" form:"answer" json:"answer"`
	TAuthor   string     `gorm:"type:varchar(200)" form:"-" json:"author"`
	Hint1     string     `gorm:"type:text" form:"hint1" json:"hint1"`
	Hint2     string     `gorm:"type:text" form:"hint2" json:"hint2"`
	Hint3     string     `gorm:"type:text" form:"hint3" json:"hint3"`
	CreatedAt time.Time  `form:"-" json:"-"`
	UpdatedAt time.Time  `form:"-" json:"-"`
	DeletedAt *time.Time `sql:"index" form:"-" json:"-"`
	Ctime     int64      `gorm:"-" json:"ctime"`
}
