package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Employee struct {
	ID         int
	Name       string
	Department string
	Salary     int
}

type Book struct {
	ID     int
	Title  string
	Author string
	Price  int
}

func main() {
	initDB()
	findEmployee()
	findMaxSalary()
	findBooks()
}

var db *sqlx.DB

func initDB() (err error) {
	// 用户名，密码，端口号以及数据库名称，根据自己需要去修改
	// charset指定编码格式
	// parseTime可以自动解析数据库中的时间数据，方便导入到go语言中
	dsn := "root:123456@tcp(127.0.0.1:3307)/sqlx_base?charset=utf8mb4&parseTime=True&loc=Local"
	// 也可以使用MustConnect连接不成功就panic
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return err
	}
	// 连接池最大容量设置
	db.SetMaxOpenConns(20) // 与数据库建立连接的最大数目
	db.SetMaxIdleConns(10) // 连接池中的最大闲置连接数
	return err
}

// 编写Go代码，使用Sqlx查询 employees表中所有部门为技术部"的员工信息，并将结果映射到一个自定义的 Employee结构体切片中
func findEmployee() {
	sqlStr := "SELECT * FROM employees WHERE Department = ?"
	var emps []Employee
	err := db.Select(&emps, sqlStr, "技术部")
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	fmt.Printf("技术部的员工信息:%#v\n", emps)
}

// 编写Go代码，使用Sqlx查询employees 表中工资最高的员工信息，并将结果映射到一个Employee结构体中
func findMaxSalary() {
	sqlStr := "SELECT * FROM employees order by salary desc limit 1"
	var emp Employee
	err := db.Get(&emp, sqlStr)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	fmt.Printf("工资最高的员工信息:%#v\n", emp)
}

// 编写Go代码，使用Sqlx执行一个复杂的查询，例如查询介格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全
func findBooks() {
	sqlStr := "SELECT * FROM books WHERE price > ?"
	var books []Book
	err := db.Select(&books, sqlStr, 50)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	fmt.Printf("介格大于 50 元的书籍:%#v\n", books)
}
