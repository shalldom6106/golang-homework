package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID        uint `gorm:"primarykey"`
	Name      string
	Posts     []Post
	PostCount int
}

type Post struct {
	ID           uint `gorm:"primarykey"`
	Title        string
	CommentState string
	CommentCount int
	UserID       uint
	Comments     []Comment
}

type Comment struct {
	ID      uint `gorm:"primarykey"`
	Content string
	PostID  uint
}

func main() {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3307)/gorm?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}
	//编写Go代码，使用Gorm创建这些模型对应的数据库表
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Post{})
	db.AutoMigrate(&Comment{})
	//初始化数据
	userData := []*User{
		{Name: "张三", PostCount: 2},
		{Name: "李四", PostCount: 2},
	}
	postData := []*Post{
		{Title: "文章A", CommentCount: 2, UserID: 1},
		{Title: "文章B", CommentCount: 2, UserID: 1},
		{Title: "文章C", CommentCount: 1, UserID: 2},
		{Title: "文章D", CommentCount: 1, UserID: 2},
	}
	commentData := []*Comment{
		{Content: "评论1", PostID: 1},
		{Content: "评论2", PostID: 2},
		{Content: "评论3", PostID: 3},
		{Content: "评论4", PostID: 4},
		{Content: "评论5", PostID: 1},
		{Content: "评论6", PostID: 2},
	}
	db.Create(&userData)
	db.Session(&gorm.Session{SkipHooks: true}).Create(&postData)
	db.Create(&commentData)
	//编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息
	var users []User
	db.Debug().Preload("Posts").Preload("Posts.Comments").Find(&users, "users.name = ?", "张三")
	fmt.Println(users)
	//编写Go代码，使用Gorm查询评论数量最多的文章信息
	var post Post
	db.Debug().Preload("Comments").Order("comment_count DESC").First(&post)
	fmt.Println(post)
	//创建文章
	postNew := Post{Title: "文章标题1", CommentState: "无评论", CommentCount: 0, UserID: 1}
	db.Debug().Create(&postNew)
	//删除评论
	var comment Comment
	db.Debug().Where("ID = ?", "3").First(&comment)
	db.Debug().Where("ID = ?", "3").Delete(&comment)
}

// 文章创建时自动更新用户的文章数量统计字段
func (po *Post) AfterCreate(tx *gorm.DB) (err error) {
	tx.Debug().Model(&User{}).Where("ID = ?", po.UserID).Update("PostCount", gorm.Expr("post_count + ?", 1))
	return
}

// 在评论删除时检査文章的评论数量，如果评论数量为0，则更新文章的评论状态为“无评论”
func (co *Comment) AfterDelete(tx *gorm.DB) (err error) {
	result := tx.Debug().Find(&Comment{}, "post_id = ?", co.PostID)
	//修改评论数量
	tx.Debug().Model(&Post{}).Where("ID = ?", co.PostID).Update("CommentCount", result.RowsAffected)
	if result.RowsAffected == 0 { //评论数量为0时，修改评论状态
		tx.Debug().Model(&Post{}).Where("ID = ?", co.PostID).Update("CommentState", "无评论")
	}
	return
}
