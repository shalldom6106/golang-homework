package controllers

import (
	"init_order/task4/config"
	"init_order/task4/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommentData struct {
	Content string `json:"content" binding:"required"`
	PostID  uint   `json:"postid" binding:"required"`
}

// 发表评论
func CreateComment(c *gin.Context) {
	var commentdata CommentData
	if err := c.ShouldBindJSON(&commentdata); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
	}
	userID, _ := c.Get("userID")
	//发表评论
	comment := models.Comment{
		Content: commentdata.Content,
		UserID:  userID.(uint),
		PostID:  commentdata.PostID,
	}
	if err := config.DB.Debug().Create(&comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "发表评论失败，错误：" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "发表评论成功"})
}

// 读取评论
func FindComment(c *gin.Context) {
	postID := c.Param("postID")
	var comments []models.Comment
	if err := config.DB.Debug().Preload("User").Preload("Post").Where("post_id = ?", postID).Find(&comments).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "评论读取失败" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"该文章的评论信息": comments})
}
