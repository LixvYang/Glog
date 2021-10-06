package controller

import (
	"glog/model"
	"glog/utils/errmsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var code int

// Check if user exist
func CheckUser(c *gin.Context) {
	//var
}

// Add user
func AddUser(c *gin.Context) {
	//todo
	var user model.User
	_ = c.ShouldBindJSON(&user)

	code = model.CheckUser(user.Username)
	if code == errmsg.SUCCSE {
		model.CreateUser(&user)
	}

	if code == errmsg.ERROR_USERNAME_USED {
		code = errmsg.ERROR_USERNAME_USED
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    user,
		"message": errmsg.GetErrMsg(code),
	})
}

// Get user
func GetUser(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	username := c.Query("username")

	if pageSize == 0 {
		pageSize = -1
	}

	if pageNum == 0 {
		pageNum = -1
	}

	data, total := model.GetUsers(username, pageSize, pageNum)
	code := errmsg.SUCCSE
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"total":   total,
			"message": errmsg.GetErrMsg(code),
		},
	)

}

// Exit user
func ExitUser(c *gin.Context) {
	var user model.User
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&user)

	code := model.CheckUpUser(id, user.Username)
	if code == errmsg.SUCCSE {
		model.EditUser(id, &user)
	}

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// Delete User
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code := model.DeleteUser(id)

	c.JSON(
		http.StatusOK,gin.H{
			"status":code,
			"message":errmsg.GetErrMsg(code),
		},
	)
}
