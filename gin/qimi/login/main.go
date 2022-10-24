package main

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"net/http"
)

var db *sqlx.DB

func main() {
	//初始化连接
	if err := initDB(); err != nil {
		fmt.Println("连接数据库失败: ", err)
		panic(err)
	}
	r := gin.Default()

	r.LoadHTMLFiles("gin/qimi/login/login.html")

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	// 接收请求并处理请求
	r.POST("/login", LoginHandler)

	//启动服务
	r.Run()
}

type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type User struct {
	// db的tag,对应数据库里的id和username，sqlx的框架对应的
	// json里的对应的是序列化和反序列化的tag
	Id       int    `db:"id" json:"-"`
	Username string `db:"username" json:"name"`
	Desc     string `json:",omitempty"` // 忽略掉为空的参数
}

// 处理登录请求的函数
func LoginHandler(c *gin.Context) {
	// 从请求中获取用户的请求数据
	var reqData Login
	// 有错误就返回错误码，这个其实是对传入数据的格式做了初步校验
	if err := c.ShouldBind(&reqData); err != nil {
		// 从请求中解析数据出错
		c.JSON(http.StatusOK, gin.H{
			// 内部业务状态码
			"code": 1,
			"msg":  "请求参数格式错误",
		})
	}
	//fmt.Printf("reqData:%#v\n", reqData)
	// 返回数据
	//c.JSON(http.StatusOK, reqData)

	// 对数据进行校验
	if u, err := QueryUser(reqData.Username, reqData.Password); err == nil {
		// 登录成功
		fmt.Println(u)
		u.Desc = "test"
		c.JSON(http.StatusOK, gin.H{
			// 内部业务状态码
			"code": 0,
			"msg":  u.Username + "登录成功",
			"data": u,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			// 内部业务状态码
			"code": 1,
			"msg":  "用户名或密码错误",
		})
	}
	// 返回响应
}

func initDB() (err error) {
	dsn := "root:815qza@tcp(192.168.72.130:3306)/sql_test?charset=utf8mb4&parseTime=True"
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed,err:%v\n", err)
		return
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	return
}

func QueryUser(username string, password string) (*User, error) {
	// 查库，注意查的时候和那个Login是不一样的，用的是User查的
	sqlStr := "select id, username from login where username=? and password=?"
	var u User
	// 自动把数据塞到结构体里，这里挎包了，所以结构体要大写，用的是反射。
	err := db.Get(&u, sqlStr, username, password)
	if err != nil {
		fmt.Println(errors.Is(err, sql.ErrNoRows)) // 没有查询到记录
		fmt.Printf("get failed,err:%v\n", err)
		return nil, err
	}
	return &u, nil
}
