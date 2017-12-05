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

	userInfo1 := new(UserInfo)
	userInfo1.UserName = "user1"
	userInfo1.PassWord = "password1"
	err := orm.Add(userInfo1)
	if err != nil{
		panic(err)
	}


	userInfo2 := new(UserInfo)
	userInfo2.Id = userInfo1.Id
	orm.Load(userInfo2)

	fmt.Println(userInfo1)
	fmt.Println(userInfo2)

	userInfo1.UserName = "user2"
	err = orm.Save(userInfo1)


	fmt.Println(userInfo1)

	orm.Del(userInfo1)
	orm.Load(userInfo2)

	fmt.Println(userInfo2)


}
