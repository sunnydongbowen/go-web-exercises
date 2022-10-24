package qimi

import (
	"fmt"
	//"github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

// 定义一个结构体对应的数据库里的一张表
type Product struct {
	gorm.Model // 这里是结构图嵌套，里面有创建，删除，更新时间
	Code       string
	Price      uint
}

func TestGorm(t *testing.T) {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:815qza@tcp(192.168.72.130:3306)/sql_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("gorm.Open failed", err)
		panic(err)
	}
	db.AutoMigrate(&Product{})

	//
	//p1 := &Product{
	//	Code:  "001",
	//	Price: 100,
	//}
	//db.Create(p1)

	var p2 Product
	db.First(&p2, 1) // 根据整型主建id查询
	fmt.Printf("p2:%#v\n", p2)

	db.First(&p2, "code=?", "001")
	fmt.Printf("p2:%#v\n", p2)

	// update
	//db.Model(&p2).Update("Price", 200)
	//db.Model(&p2).Updates(Product{Price: 300, Code: "F33"})

}
