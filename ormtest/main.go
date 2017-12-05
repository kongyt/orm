package main

import (
	"fmt"
	"orm"
)


type UserInfo struct{
	Id 	int				`field:"id" index:"pk" auto:"true" table:"userinfo"`
	UserName	string	`field:"username"`
	PassWord	string	`field:"password"`
	TestField	string
}

func main()  {
	orm := orm.NewOrm()

	orm.Register(&UserInfo{})

	orm.Open("mysql", "连接字符串")
	defer orm.Close()

    // 增
	userInfo1 := new(UserInfo)
	userInfo1.UserName = "user1"
	userInfo1.PassWord = "password1"
	orm.Add(userInfo1)

    // 查
	userInfo2 := new(UserInfo)
	userInfo2.Id = userInfo1.Id
	orm.Load(userInfo2)

	fmt.Println(userInfo1)
	fmt.Println(userInfo2)

    // 改
	userInfo1.UserName = "user2"
	orm.Save(userInfo1)

    // 删
	orm.Del(userInfo1)

}
