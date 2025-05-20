package controller

import (
	"Golunwen2/App/dto"
	"Golunwen2/App/model"
	"Golunwen2/App/tool"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var loginReq LoginRequest
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "请求数据格式错误",
		})
		return
	}

	//数据验证
	//验证用户名和密码是否为空
	if loginReq.Name == "" || loginReq.Password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusUnprocessableEntity,
			"msg":  "用户名或密码不能为空",
		})
		return
	}

	//验证用户名是否存在
	var user model.User
	result := model.DB.Table("user").Where("name=?", loginReq.Name).First(&user)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "用户不存在",
		})
		return
	}

	//验证密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginReq.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "密码错误",
		})
		return
	}

	//生成token
	token, err := tool.ReleaseToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "生成token失败",
		})
		return
	}

	//登录成功
	c.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"msg":   "登录成功",
		"token": gin.H{"token": token},
	})
}

func Info(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"msg":  "获取用户信息成功",
		"user": dto.ToUserDto(user.(model.User)),
	})
}
