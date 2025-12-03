package main

import (
	"errors"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// 题目1：模型定义
// User 模型定义
type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"size:50;not null;unique"`  // 用户名
	Email     string `gorm:"size:100;not null;unique"` // 邮箱
	PostCount int    `gorm:"default:0"`                // 文章数量统计
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Posts     []Post         `gorm:"foreignKey:UserID"` // 一对多关系
}

// Post 模型定义
type Post struct {
	ID            uint   `gorm:"primaryKey"`
	Title         string `gorm:"size:200;not null"`     // 文章标题
	Content       string `gorm:"type:text"`             // 文章内容
	CommentCount  int    `gorm:"default:0"`             // 评论数量
	CommentStatus string `gorm:"size:20;default:'有评论'"` // 评论状态
	UserID        uint   `gorm:"not null"`              // 外键
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	User          User           `gorm:"foreignKey:UserID"` // 关联用户
	Comments      []Comment      `gorm:"foreignKey:PostID"` // 一对多关系
}

// Comment 模型定义
type Comment struct {
	ID        uint   `gorm:"primaryKey"`
	Content   string `gorm:"size:500;not null"` // 评论内容
	PostID    uint   `gorm:"not null"`          // 外键
	UserID    uint   `gorm:"not null"`          // 评论者ID
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Post      Post           `gorm:"foreignKey:PostID"` // 关联文章
}

// 题目2：关联查询
// 查询某个用户发布的所有文章及其对应的评论信息
func getUserPostsWithComments(db *gorm.DB, userID uint) ([]Post, error) {
	var posts []Post
	// 使用Preload预加载关联数据
	result := db.Preload("Comments").
		Where("user_id = ?", userID).
		Find(&posts)

	if result.Error != nil {
		return nil, result.Error
	}
	return posts, nil
}

// 查询评论数量最多的文章信息
func getPostWithMostComments(db *gorm.DB) (Post, error) {
	var post Post
	// 使用Order和Limit查询评论最多的文章
	result := db.Preload("User").
		Preload("Comments").
		Order(clause.OrderByColumn{Column: clause.Column{Name: "comment_count"}, Desc: true}).
		Limit(1).
		First(&post)

	if result.Error != nil {
		return post, result.Error
	}
	return post, nil
}

// 题目3：钩子函数
// 为Post模型添加创建前的钩子函数，更新用户的文章数量
func (p *Post) BeforeCreate(tx *gorm.DB) error {
	// 检查用户是否存在
	var user User
	if err := tx.First(&user, p.UserID).Error; err != nil {
		return errors.New("用户不存在")
	}

	// 更新用户的文章数量
	return tx.Model(&User{}).Where("id = ?", p.UserID).
		Update("post_count", gorm.Expr("post_count + 1")).Error
}

// 为Comment模型添加删除后的钩子函数，更新文章的评论数量和状态
func (c *Comment) AfterDelete(tx *gorm.DB) error {
	// 1. 减少对应文章的评论计数
	var post Post
	if err := tx.First(&post, c.PostID).Error; err != nil {
		return errors.New("文章不存在")
	}

	// 更新评论数量
	if err := tx.Model(&Post{}).Where("id = ?", c.PostID).
		Update("comment_count", gorm.Expr("comment_count - 1")).Error; err != nil {
		return err
	}

	// 2. 检查评论数量是否为0
	var commentCount int64
	if err := tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&commentCount).Error; err != nil {
		return err
	}

	// 如果评论数量为0，更新评论状态
	if commentCount == 0 {
		return tx.Model(&Post{}).Where("id = ?", c.PostID).
			Update("comment_status", "无评论").Error
	}

	return nil
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test001?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))

	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)

	}

	db.AutoMigrate(&User{}, &Post{}, &Comment{})
	// 创建用户
	user := User{Username: "meta3", Email: "blogger@meta.com"}
	// isCreateUser:= false

	// if isCreateUser{
	db.Create(&user)
	fmt.Printf("创建用户: %+v\n", user)

	// 创建文章（会触发Post的BeforeCreate钩子）
	post := Post{Title: "meta_Sqlx入门", Content: "这是一篇关于Sqlx用法的文章", UserID: user.ID}
	db.Create(&post)
	fmt.Printf("创建文章: %+v\n", post)

	// 创建评论
	comment1 := Comment{Content: "1meta_Sqlx很棒的文章！", PostID: post.ID, UserID: user.ID}
	comment2 := Comment{Content: "2meta_Sqlx学习了", PostID: post.ID, UserID: user.ID}
	comment3 := Comment{Content: "3meta_Sqlx学习了1", PostID: post.ID, UserID: user.ID}
	db.Create(&comment1)
	db.Create(&comment2)
	db.Create(&comment3)
	db.Model(&Post{}).Where("id = ?", post.ID).Update("comment_count", 3)

	post = Post{Title: "Gorm进阶教程", Content: "这是一篇关于Gorm进阶用法的文章", UserID: user.ID}
	db.Create(&post)
	fmt.Printf("创建文章: %+v\n", post)
	// 创建评论
	comment1 = Comment{Content: "很棒的文章！", PostID: post.ID, UserID: user.ID}
	comment2 = Comment{Content: "学习了", PostID: post.ID, UserID: user.ID}
	db.Create(&comment1)
	db.Create(&comment2)
	// 更新文章的评论计数
	db.Model(&Post{}).Where("id = ?", post.ID).Update("comment_count", 2)
	// }
	// 题目2：关联查询示例
	//1. 查询用户的所有文章及评论
	posts, err := getUserPostsWithComments(db, user.ID)
	if err != nil {
		fmt.Printf("查询用户文章失败: %v\n", err)
	} else {
		fmt.Printf("\n用户(%d)的文章及评论: \n", user.ID)
		for _, p := range posts {
			fmt.Printf("UID:%d PID:%d 文章: %s, 评论数: %d\n", p.UserID, p.ID, p.Title, len(p.Comments))
			for _, c := range p.Comments {
				fmt.Printf("UID:%d PID:%d CID:%d  评论: %s\n", c.UserID, c.PostID, c.ID, c.Content)
			}
		}
	}

	// 2. 查询评论最多的文章
	topPost, err := getPostWithMostComments(db)
	if err != nil {
		fmt.Printf("查询评论最多的文章失败: %v\n", err)
	} else {
		fmt.Printf("\n评论最多的文章: %s, 评论数: %d\n", topPost.Title, topPost.CommentCount)
	}

	// 查看更新前的文章状态
	var updatedPost Post
	db.First(&updatedPost, post.ID)
	fmt.Printf("更新前 文章 %s 的评论状态: %s, 评论数: %d\n",
		updatedPost.Title, updatedPost.CommentStatus, updatedPost.CommentCount)

	// 测试评论删除钩子
	db.Delete(&comment1)
	db.Delete(&comment2)
	fmt.Println("\n删除评论后检查文章状态")

	// 查看更新后的文章状态
	db.First(&updatedPost, post.ID)
	fmt.Printf("更新后 文章 %s 的评论状态: %s, 评论数: %d\n",
		updatedPost.Title, updatedPost.CommentStatus, updatedPost.CommentCount)

}
