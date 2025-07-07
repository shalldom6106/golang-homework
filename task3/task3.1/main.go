package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Student struct {
	ID    uint `gorm:"primarykey"`
	Name  string
	Age   int
	Grade string
}

type Account struct {
	ID      uint `gorm:"primarykey"`
	Name    string
	Balance int
}

type Transaction struct {
	ID            uint `gorm:"primarykey"`
	FromAccountId int
	ToAccountId   int
	Amount        int
}

func main() {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3307)/gorm?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}
	//SQL语句练习 题目1：基本CRUD操作
	db.AutoMigrate(&Student{})
	//编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"
	student := Student{Name: "张三", Age: 20, Grade: "三年级"}
	db.Debug().Create(&student)
	//编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息
	studentFind := Student{}
	db.Debug().Where("Age > ?", "18").Find(&studentFind)
	//编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"
	db.Debug().Model(&Student{}).Where("Name = ?", "张三").Update("Grade", "四年级")
	//编写SQL语句删除 students 表中年龄小于 15 岁的学生记录
	db.Debug().Where("Age < ?", "15").Delete(&Student{})
	//SQL语句练习 题目2：事务语句
	//编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。
	//在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，
	//并在 transactions 表中记录该笔转账信息。
	//如果余额不足，则回滚事务
	db.AutoMigrate(&Account{})
	db.AutoMigrate(&Transaction{})
	//初始化创建A账户、B账户信息
	accounts := []*Account{
		{Name: "A", Balance: 200},
		{Name: "B", Balance: 300},
	}
	db.Debug().Create(&accounts)
	//查询A账户的余额
	var accountA Account
	db.Debug().Where("Name = ?", "A").First(&accountA)
	money := 100
	if accountA.Balance < money {
		fmt.Println("余额不足")
	} else {
		//账户A扣除100
		db.Debug().Model(&Account{}).Where("Name = ?", "A").Update("Balance", gorm.Expr("Balance - ?", money))
		//账户B增加100
		db.Debug().Model(&Account{}).Where("Name = ?", "B").Update("Balance", gorm.Expr("Balance + ?", money))
		//transactions表新增转账信息
		transaction := Transaction{FromAccountId: 1, ToAccountId: 2, Amount: money}
		db.Debug().Create(&transaction)
	}

}
