package main

import (
	// "fmt"

	"errors"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

/*
*

	题目1：模型定义
	假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
	要求 ：
	使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系
	（一篇文章可以有多个评论）。
	编写Go代码，使用Gorm创建这些模型对应的数据库表。
	题目2：关联查询
	基于上述博客系统的模型定义。
	要求 ：
	编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
	编写Go代码，使用Gorm查询评论数量最多的文章信息。
	题目3：钩子函数
	继续使用博客系统的模型。
	要求 ：
	为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
	为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。

*
*/
type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:50;not null;unique"` // 用户名
	PostCount int    `gorm:"default:0"`               // 文章数量统计
	Posts     []Post `gorm:"foreignKey:UserID"`       // 一对多关系
}
type Post struct {
	ID           uint      `gorm:"primaryKey"`
	Title        string    `gorm:"size:200;not null"` // 文章标题
	Content      string    `gorm:"type:text"`         // 文章内容
	CommentCount int       `gorm:"default:0"`         // 评论数量
	UserID       uint      `gorm:"not null"`          // 外键
	User         User      `gorm:"foreignKey:UserID"` // 关联用户
	Comments     []Comment `gorm:"foreignKey:PostID"` // 一对多关系
}

type Comment struct {
	ID      uint   `gorm:"primaryKey"`
	Content string `gorm:"size:500;not null"` // 评论内容
	PostID  uint   `gorm:"not null"`          // 外键
	UserID  uint   `gorm:"not null"`          // 评论者ID
	Post    Post   `gorm:"foreignKey:PostID"` // 一对多关系

}

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

func (p *Post) BeforeCreate(tx *gorm.DB) error {
	var user User
	if err := tx.First(&user, p.UserID).Error; err != nil {
		return errors.New("用户不存在")
	}

	er := tx.Model(&User{}).Where("id=?", p.UserID).Update("post_count", gorm.Expr("post_count+1")).Error
	log.Println("文章创建后回调函数执行", er)

	return er
}
func (c Comment) AfterDelete(tx *gorm.DB) error {
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

	return nil
}

func gePostsAndComments(db *gorm.DB, userid uint) ([]Post, error) {
	var posts []Post
	result := db.Preload("Comments").Where("id=?", userid).Find(&posts)
	if result.Error != nil {
		return nil, result.Error
	}
	return posts, nil
}
func geMaxComments(db *gorm.DB) (Post, error) {
	var post Post
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

func main() {
	// 连接数据库
	dsn := "root:123456@tcp(127.0.0.1:3306)/test001?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))

	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}
	db.AutoMigrate(&User{}, &Post{}, &Comment{})

	// 创建用户
	user := User{Name: "meta5"}
	result := db.Create(&user)
	fmt.Printf("创建用户: %+v\n", user)

	if result.Error != nil {
		log.Fatal("创建用户失败：", result.Error)
	}

	// 创建文章（会触发Post的BeforeCreate钩子）
	post := Post{Title: "grom11", Content: "这是一篇关于grom用法的文章1", UserID: user.ID}
	resultp := db.Create(&post)
	if resultp.Error != nil {
		log.Fatal("创建用户失败：", resultp.Error)
	}

	// fmt.Printf("创建文章: %+v\n", post)

	// 创建评论
	comment1 := Comment{Content: "meta1很棒的文章", PostID: post.ID, UserID: user.ID}
	// comment2 := Comment{Content: "2meta_Sqlx学习了", PostID: 2, UserID: 2}
	comment2 := Comment{Content: "1meta_Sqlx很棒的文章", PostID: post.ID, UserID: user.ID}
	comment3 := Comment{Content: "3meta_Sqlx学习了1", PostID: post.ID, UserID: user.ID}
	db.Create(&comment1)
	db.Create(&comment2)
	result3 := db.Create(&comment3)
	if result3.Error != nil {
		log.Fatal("创建用户失败：", result3.Error)
	}
	// 更新文章的评论计数
	db.Model(&Post{}).Where("id = ?", post.ID).Update("comment_count", 3)

	post = Post{Title: "Gorm进阶教程", Content: "这是一篇关于Gorm进阶用法的文章", UserID: user.ID}
	db.Create(&post)
	fmt.Printf("创建文章: %+v\n", post)
	// 创建评论
	comment1 = Comment{Content: "很棒！", PostID: post.ID, UserID: user.ID}
	comment2 = Comment{Content: "学习", PostID: post.ID, UserID: user.ID}
	db.Create(&comment1)
	db.Create(&comment2)
	// 更新文章的评论计数
	db.Model(&Post{}).Where("id = ?", post.ID).Update("comment_count", 2)

	posts, err := gePostsAndComments(db, user.ID)
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

	topPost, err := geMaxComments(db)
	if err != nil {
		fmt.Printf("查询用户文章失败: %v\n", err)
	} else {
		fmt.Printf("\n评论最多的文章: %s, 评论数: %d\n", topPost.Title, topPost.CommentCount)
	}

	var updatedPost Post
	db.First(&updatedPost, post.ID)
	fmt.Printf("更新前 文章 %s 的 评论数: %d\n", updatedPost.Title, updatedPost.CommentCount)
	db.Delete(&comment1)
	db.Delete(&comment2)

	fmt.Printf("更新hou 文章 %s 的 评论数: %d\n", updatedPost.Title, updatedPost.CommentCount)

}
