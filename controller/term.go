package controller

import (
	"net/http"
	"strconv"

	"github.com/azy0922/flywind/model"
	"github.com/azy0922/flywind/utils"
	"github.com/gin-gonic/gin"
)

type TAuthors struct {
	// 字段名称必须和数据库名称对应且必须大写！
	TAuthor string
}

func ListTAuthors(c *gin.Context) {
	var authors []TAuthors
	var name []string

	// 只查询某个字段的值用scan方法
	err := model.DB.Table("terms").
		Select("distinct(t_author)").Scan(&authors).Error

	if err != nil {
		SendErrJson(err.Error(), c)
		return
	}

	for _, n := range authors {
		name = append(name, n.TAuthor)
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"errMsg":  "",
		"data":    name,
	})
}

func ListTerms(c *gin.Context) {
	var total int
	var term []model.Term

	// 默认每页显示5条记录
	limit := c.DefaultQuery("limit", "5")
	offset := c.DefaultQuery("offset", "0")

	err := model.DB.Find(&term).Count(&total).Error
	if err != nil {
		SendErrJson(err.Error(), c)
		return
	}

	err = model.DB.Order("ID desc").Limit(limit).
		Offset(offset).Find(&term).Error
	if err != nil {
		SendErrJson(err.Error(), c)
		return
	}

	// 注意p是拷贝
	for i, p := range term {
		term[i].Ctime = p.CreatedAt.UnixNano() / 1e6
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"errMsg":  "",
		"data": gin.H{
			"total": total,
			"rows":  term,
		},
	})
}

func ListTerm(c *gin.Context) {
	var term model.Term
	id := c.Param("id")

	if _, err := strconv.Atoi(id); err != nil {
		SendErrJson("参数无效", c)
		return
	}

	// 查询id
	if err := model.DB.Where("id = ?", id).Find(&term).Error; err != nil {
		SendErrJson(err.Error(), c)
		return
	}

	term.Ctime = term.CreatedAt.UnixNano() / 1e6

	// 返回json
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"errMsg":  "",
		"data":    term,
	})
}

func AddTerm(c *gin.Context) {
	var form model.Term
	if err := c.ShouldBind(&form); err != nil {
		SendErrJson(err.Error(), c)
		return
	}

	if form.Answer == "" || form.Hint1 == "" ||
		form.Hint2 == "" || form.Hint3 == "" {
		SendErrJson("参数不允许为空！", c)
		return
	}
	// 取当前登录用户名称
	session, _ := Store.Get(c.Request, "session")
	val, ok := session.Values["username"]
	if ok {
		form.TAuthor = val.(string)
	}
	if err := model.DB.Create(&form).Error; err != nil {
		SendErrJson(err.Error(), c)
		return
	}

	// 返回成功json消息，form可直接返回
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"errMsg":  "",
		"data":    form,
	})
}

func UpdateTerm(c *gin.Context) {
	// 先查询出该条记录
	var term model.Term
	id := c.Param("id")

	if _, err := strconv.Atoi(id); err != nil {
		SendErrJson("参数无效", c)
		return
	}

	// 查询id
	if err := model.DB.Where("id = ?", id).Find(&term).Error; err != nil {
		SendErrJson(err.Error(), c)
		return
	}

	// 而后修改
	if err := c.ShouldBind(&term); err != nil {
		SendErrJson(err.Error(), c)
		return
	}

	// 只更新有变化的属性
	if err := model.DB.Model(&term).Updates(term).Error; err != nil {
		SendErrJson(err.Error(), c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"errMsg":  "",
		"data":    term,
	})
}

func DeleteTerm(c *gin.Context) {
	id := c.Param("id")
	if _, err := strconv.Atoi(id); err != nil {
		SendErrJson("参数无效", c)
		return
	}
	err := model.DB.Where("id = ?", id).Delete(&model.Term{}).Error
	if err != nil {
		SendErrJson(err.Error(), c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"errMsg":  "",
		"data":    gin.H{},
	})

}

// 直接从数据库中随机返回一条结果
func ShowTerms(c *gin.Context) {
	var term model.Term

	err := model.DB.Order("random()").Limit(1).Find(&term).Error
	if err != nil {
		SendErrJson(err.Error(), c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"category": "term",
		"answer":   term.Answer,
		"hint1":    utils.Encode(term.Hint1),
		"hint2":    utils.Encode(term.Hint2),
		"htin3":    utils.Encode(term.Hint3),
	})
}
