package controller

import (
	"net/http"
	"strconv"

	"github.com/azy0922/flywind/model"
	"github.com/azy0922/flywind/problem"
	"github.com/azy0922/flywind/utils"
	"github.com/gin-gonic/gin"
)

type Authors struct {
	// 字段名称必须和数据库名称对应且必须大写！
	Author string
}

func ListAuthors(c *gin.Context) {
	var authors []Authors
	var name []string

	// 只查询某个字段的值用scan方法
	err := model.DB.Table("problems").
		Select("distinct(author)").
		Scan(&authors).Error

	if err != nil {
		SendErrJson(err.Error(), c)
		return
	}

	for _, n := range authors {
		name = append(name, n.Author)
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"errMsg":  "",
		"data":    name,
	})
}

func ListProblems(c *gin.Context) {
	var total int
	var prob []model.Problem

	// 默认每页显示5条记录
	limit := c.DefaultQuery("limit", "5")
	offset := c.DefaultQuery("offset", "0")

	err := model.DB.Find(&prob).Count(&total).Error
	if err != nil {
		SendErrJson(err.Error(), c)
		return
	}

	err = model.DB.Order("ID desc").Limit(limit).
		Offset(offset).Find(&prob).Error

	if err != nil {
		SendErrJson(err.Error(), c)
		return
	}

	// 注意p是拷贝
	for i, p := range prob {
		prob[i].Ctime = p.CreatedAt.UnixNano() / 1e6
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"errMsg":  "",
		"data": gin.H{
			"total": total,
			"rows":  prob,
		},
	})
}

func ListProblem(c *gin.Context) {
	var problem model.Problem
	id := c.Param("id")

	if _, err := strconv.Atoi(id); err != nil {
		SendErrJson("参数无效", c)
		return
	}

	// 查询id
	if err := model.DB.Where("id = ?", id).Find(&problem).Error; err != nil {
		SendErrJson(err.Error(), c)
		return
	}

	problem.Ctime = problem.CreatedAt.UnixNano() / 1e6

	// 返回json
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"errMsg":  "",
		"data":    problem,
	})
}

func AddProblem(c *gin.Context) {
	var form model.Problem
	if err := c.ShouldBind(&form); err != nil {
		SendErrJson(err.Error(), c)
		return
	}

	if form.Description == "" || form.Category == "" {
		SendErrJson("参数不允许为空！", c)
		return
	}
	// 取当前登录用户名称
	session, _ := Store.Get(c.Request, "session")
	val, ok := session.Values["username"]
	if ok {
		form.Author = val.(string)
	}
	if err := model.DB.Create(&form).Error; err != nil {
		SendErrJson(err.Error(), c)
		return
	}

	problem.ChanRebuild <- true

	// 返回成功json消息，form可直接返回
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"errMsg":  "",
		"data":    form,
	})
}

func UpdateProblem(c *gin.Context) {
	// 先查询出该条记录
	var prob model.Problem
	id := c.Param("id")

	if _, err := strconv.Atoi(id); err != nil {
		SendErrJson("参数无效", c)
		return
	}

	// 查询id
	if err := model.DB.Where("id = ?", id).Find(&prob).Error; err != nil {
		SendErrJson(err.Error(), c)
		return
	}

	// 而后修改
	if err := c.ShouldBind(&prob); err != nil {
		SendErrJson(err.Error(), c)
		return
	}

	// 只更新有变化的属性
	if err := model.DB.Model(&prob).Updates(prob).Error; err != nil {
		SendErrJson(err.Error(), c)
		return
	}

	problem.ChanRebuild <- true

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"errMsg":  "",
		"data":    prob,
	})
}

func DeleteProblem(c *gin.Context) {
	id := c.Param("id")
	if _, err := strconv.Atoi(id); err != nil {
		SendErrJson("参数无效", c)
		return
	}
	err := model.DB.Where("id = ?", id).Delete(&model.Problem{}).Error
	if err != nil {
		SendErrJson(err.Error(), c)
		return
	}

	problem.ChanRebuild <- true

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"errMsg":  "",
		"data":    gin.H{},
	})

}

func TestMaths(c *gin.Context) {
	var pb *problem.ProblemRet

	category := c.Query("c")
	pb = problem.Generate(category)

	if pb == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "未知错误",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"category": pb.Category,
		"question": pb.Question,
		"answer":   pb.Answer,
	})
}

func ShowMaths(c *gin.Context) {
	var pb *problem.ProblemRet

	category := c.Query("c")
	pb = problem.Generate(category)

	if pb == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "未知错误",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"category": pb.Category,
		"question": utils.Encode(pb.Question),
		"answer":   pb.Answer,
	})
}
