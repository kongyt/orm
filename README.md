# orm
一个超轻量级的ORM库

通过Tag映射结构体到数据库字段
Tag含义
    field:"数据库字段名"              该结构体域会被数据持久化
    index:"pk"                        主键，有且只有一个
    auto:"true或者false"              主键是否自增，为false时需要指定主键的值，不然会插入结构体主键字段当前的值（只能标注在主键字段上）
    table:"数据库表名"                数据库表的名字（只能标注在主键字段上）

数据的删改查条件语句全是依赖主键，插入语句在主键设置为自增的情况下，会自动填充last insert id到结构体中


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
