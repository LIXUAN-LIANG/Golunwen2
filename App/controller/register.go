package controller

import (
	"Golunwen2/App/model"
	"fmt"
	"net/http"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RegisterRequest struct {
	Name      string `json:"name"`
	Password  string `json:"password"`
	Telephone string `json:"telephone"`
}

func Register(c *gin.Context) {
	var registerReq RegisterRequest
	if err := c.ShouldBindJSON(&registerReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "请求数据格式错误",
		})
		return
	}

	// 数据验证
	// 判断手机号是否对的
	if len(registerReq.Telephone) != 11 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":   422,
			"msg":    "手机号不合法",
			"length": len(registerReq.Telephone),
		})
		return
	}
	//判断手机号是否存在
	if telephoneExist(model.DB, registerReq.Telephone) {
		c.JSON(200, gin.H{
			"code": 422,
			"msg":  "手机号已存在",
		})
		return
	}

	// 判断密码是否符合规则
	passwordRegex := `^[a-zA-Z\d]{6,15}$`
	matchedPwd, _ := regexp.MatchString(passwordRegex, registerReq.Password)
	if !matchedPwd {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "密码只能包含字母和数字，且长度为6到15位",
		})
		return
	}

	// 判断用户名是否符合规则
	nameRegex := `^[\p{Han}a-zA-Z0-9]+$`
	matchedName, _ := regexp.MatchString(nameRegex, registerReq.Name)
	if !matchedName {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "用户名只能包含中文、英文和数字",
		})
		return
	}
	//所有都符合 创建用户
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(registerReq.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500, "msg": "密码加密失败",
		})
		return
	}
	newUser := model.User{
		Name:        registerReq.Name,
		Password:    string(hashPassword),
		Telephone:   registerReq.Telephone,
		CreatedTime: time.Now(),
	}
	createResult := model.DB.Create(&newUser)
	if createResult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "Failed to create user",
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "注册成功",
	})
}

func telephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	result := db.Table("user").Where("telephone=?", telephone).First(&user)
	if result.RowsAffected > 0 {
		// 表示存在相同的 telephone
		fmt.Println("手机号已存在")
		return true
	} else {
		// 表示不存在相同的 telephone
		fmt.Println("手机号不存在,可以使用该手机号注册")
		return false
	}
}
