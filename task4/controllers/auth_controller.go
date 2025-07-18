package controllers

import (
	"init_order/task4/config"
	"init_order/task4/models"
	"init_order/task4/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterData struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

type LoginData struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// 注册
func Register(c *gin.Context) {
	var registerData RegisterData
	if err := c.ShouldBindJSON(&registerData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	//密码加密
	hashpassword, err := utils.HashPassword(registerData.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "密码加密失败"})
		return
	}
	users := models.User{
		UserName: registerData.UserName,
		Password: hashpassword,
		Email:    registerData.Email,
	}
	if err := config.DB.Debug().Create(&users).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "账号注册失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "账号注册成功"})
}

// 登录
func Login(c *gin.Context) {
	var loginData LoginData
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	//判断是否存在改用户
	var user models.User
	if err := config.DB.Debug().Where("user_name = ?", loginData.UserName).Find(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"err": "该用户不存在"})
		return
	}
	//密码验证
	if !utils.CheckPassword(user.Password, loginData.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"err": "密码错误，请重新输入"})
		return
	}
	//jwt
	token, err := utils.GenerateJWT(user.ID, user.UserName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "生成Token失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "登录成功，token为" + token})
}
