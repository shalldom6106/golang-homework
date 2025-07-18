package controllers

import (
	"init_order/task4/config"
	"init_order/task4/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostData struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// 创建文章
func CreatePost(c *gin.Context) {
	var postdata PostData
	if err := c.ShouldBindJSON(&postdata); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	userID, _ := c.Get("userID")
	post := models.Post{
		Title:   postdata.Title,
		Content: postdata.Content,
		UserID:  userID.(uint),
	}
	if err := config.DB.Debug().Create(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "创建文章失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "创建文章成功"})
}

// 读取所有文章
func FindPosts(c *gin.Context) {
	var posts []models.Post
	if err := config.DB.Debug().Preload("User").Find(&posts).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "读取文章失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"所有文章信息": posts})
}

// 读取单份文章
func FindPostById(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	if err := config.DB.Debug().Preload("User").First(&post, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "读取文章失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"单篇文章信息": post})
}

// 更新文章
func UpdatePost(c *gin.Context) {
	id := c.Param("id")
	//判断是否本人的文章
	var post models.Post
	if err := config.DB.Debug().Find(&post, "posts.id = ?", id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "文章不存在"})
		return
	}
	userID, _ := c.Get("userID")
	if userID != post.UserID {
		c.JSON(http.StatusUnauthorized, gin.H{"err": "该文章不属于该用户，无法进行更新操作"})
		return
	}
	var postdata PostData
	if err := c.ShouldBindJSON(&postdata); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	//更新操作
	if err := config.DB.Debug().Model(&models.Post{}).Where("ID = ?", id).
		Updates(map[string]interface{}{"title": postdata.Title, "content": postdata.Content}).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "文章更新失败！错误：" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "文章更新成功"})
}

// 删除文章
func DeletePost(c *gin.Context) {
	id := c.Param("id")
	//判断是否本人的文章
	var post models.Post
	if err := config.DB.Debug().Find(&post, "posts.id = ?", id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "文章不存在"})
		return
	}
	userID, _ := c.Get("userID")
	if userID != post.UserID {
		c.JSON(http.StatusUnauthorized, gin.H{"err": "该文章不属于该用户，无法进行删除操作"})
		return
	}
	//删除文章
	if err := config.DB.Debug().Where("ID = ?", id).Delete(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "文章删除失败！错误：" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "文章删除成功"})
}
