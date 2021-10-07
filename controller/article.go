package controller

import (
	"glog/model"
	"glog/utils/errmsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AddArticle 
func AddArticle(c *gin.Context) {
	var art model.Article
	_ = c.ShouldBindJSON(&art)

	code := model.CreateArt(&art)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    art,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetCateArt 
func GetCateArt(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	id, _ := strconv.Atoi(c.Param("id"))
	
	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}
	
	if pageNum == 0 {
		pageNum = 1
	}
	
	data, code, total := model.GetCateArt(id, pageSize, pageNum)
	
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetArtInfo 
func GetArtInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	art, code := model.GetArtInfo(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    art,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetArt 
func GetArt(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	title := c.Query("title")
	
	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}
	
	if pageNum == 0 {
		pageNum = 1
	}
	if len(title) == 0 {
		data, code, total := model.GetArt(pageSize, pageNum)
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"total":   total,
			"message": errmsg.GetErrMsg(code),
		})
		return
	}
	
	data, code, total := model.SearchArticle(title,pageSize,pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// EditArt 
func EditArt(c *gin.Context) {
	var art model.Article
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&art)

	code := model.EditArt(id, &art)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// DeleteArt 
func DeleteArt(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code := model.DeleteArt(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
