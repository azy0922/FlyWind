package model

import "time"

// CREATE TABLE problem (
// 	id integer primary key autoincrement,
// 	category varchar(200),
// 	author varchar(200),
// 	description text,
// 	generatorscript text,
// 	solverscript text
//   );

// 定义计算题的数据模型
type Problem struct {
	ID              int64      `gorm:"primary_key" form:"-" json:"id"`
	Category        string     `gorm:"type:varchar(200);unique" form:"category" json:"category"`
	Description     string     `gorm:"type:text" form:"description" json:"description"`
	Author          string     `gorm:"type:varchar(200)" form:"-" json:"author"`
	GeneratorScript string     `gorm:"type:text" form:"jsgenerator" json:"jsgenerator"`
	SolverScript    string     `gorm:"type:text" form:"jssolver" json:"jssolver"`
	CreatedAt       time.Time  `form:"-" json:"-"`
	UpdatedAt       time.Time  `form:"-" json:"-"`
	DeletedAt       *time.Time `sql:"index" form:"-" json:"-"`
	Ctime           int64      `gorm:"-" json:"ctime"`
}
