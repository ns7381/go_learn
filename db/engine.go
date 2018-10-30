package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"time"
	"fmt"
)

type User struct {
	Id        int64
	Name      string
	CreatedAt time.Time `xorm:"created"`
}

var engine *xorm.Engine

func main() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:123456aB@(10.110.25.71:43189)/service_ml?charset=utf8")
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := engine.Ping(); err!=nil{
		fmt.Println(err)
		return
	}
	engine.ShowSQL(true)
	//err = engine.Sync2(new(User))
	//fmt.Println(err)

	/*var user = User{Id:1234, Name:"nathan", CreatedAt:time.Now()}
	affected, err := engine.Insert(&user)
	fmt.Println(affected)*/
	/*var user User
	engine.Id(1234).Get(&user)
	fmt.Println(user)*/
	user := &User{Id:1234}
	_, err = engine.Get(user)
}


