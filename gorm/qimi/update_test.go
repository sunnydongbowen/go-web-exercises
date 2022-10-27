package qimi

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

type Book struct {
	gorm.Model
	Title  string  `gorm:"title"`
	Amount int     `gorm:"amount "`
	Price  float64 `gorm:"price"`
	Status bool    `gorm:"status"`
}

var db *gorm.DB

func initDB() (err error) {
	dsn := "root:815qza@tcp(192.168.72.130:3306)/sql_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return
}

func update1() {
	// 没反应
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Debug().Model(&Book{}).Updates(Book{Title: "Go语言学习", Amount: 30, Status: false})
}

func update2() {
	var id uint = 1
	var b1 Book
	err := db.Where("id=?", id).First(&b1).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println("参数错误")
	}
	db.Debug().Model(&b1).Updates(Book{Title: "Python学习", Amount: 30, Status: false})
}

func update3() {
	var id uint = 1
	cond := &Book{
		Model: gorm.Model{
			ID: id,
		},
	}
	db.Debug().Where(cond).Updates(Book{Title: "English学习", Amount: 30, Status: false})
}
func TestUpdate(t *testing.T) {
	if err := initDB(); err != nil {
		fmt.Println("connect mysql err", err)
		panic(err)
	}
	db.AutoMigrate(&Book{})

	//b1 := Book{
	//	Title:  "Go语言学习",
	//	Amount: 50,
	//	Price:  99.0,
	//	Status: false,
	//}
	//db.Create(&b1)

	//update1()
	//update2()
	update3()
}
