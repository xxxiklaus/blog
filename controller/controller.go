package controller

import (
	"blog/dao"
	"blog/model"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	user := model.User{
		Username: username,
		Password: password,
	}

	dao.Mgr.Register(&user)

	c.Redirect(301, "/")

}

func GoRegister(c *gin.Context) {
	c.HTML(200, "register.html", nil)
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	fmt.Println("username")
	u := dao.Mgr.Login(username)

	if u.Username == "" {
		c.HTML(200, "login.html", "用户名不存在！")
		fmt.Println("用户名不存在！")
	} else {
		if u.Password != password {
			fmt.Println("密码错误")
			c.HTML(200, "login.html", "密码错误")
		} else {
			fmt.Println("登录成功")
			c.Redirect(301, "/")
		}
	}

}

func GoLogin(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

func Index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func GetPostIndex(c *gin.Context) {
	posts := dao.Mgr.GetAllPost()
	c.HTML(200, "post_index.html", posts)
}

func AddPost(c *gin.Context) { //添加博客
	title := c.PostForm("title")
	content := c.PostForm("content")
	tag := c.PostForm("tag")

	post := model.Post{
		Title:   title,
		Content: content,
		Tag:     tag,
	}

	dao.Mgr.AddPost(&post)

	c.Redirect(302, "/post_index")
}

func GoAddPost(c *gin.Context) { //跳转到添加博客
	c.HTML(200, "post.html", nil)
}

func ListUser(c *gin.Context) {
	c.HTML(200, "userlist.html", nil)
}
