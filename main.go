package main

import "fmt"

const info  = `请输入操作选项
1、创建新用户
2、查询指定用户
3、列出全部用户
4、更新指定用户
5、删除指定用户
6、退出`

func main() {
	fmt.Println("欢迎使用信息录入系统:")
Exit:
	for {
		fmt.Println(info)
		var input int
		fmt.Scanf("%d \n", &input)
		switch input {
		case 1:
			var name, country string
			fmt.Print("请输入姓名: ")
			fmt.Scanf("%s\n", &name)
			fmt.Print("请输入所在国家: ")
			fmt.Scanf("%s\n", &country)
			if err := newAccount(name, country); err != nil {
				fmt.Println("创建失败:", err)
			} else {
				fmt.Println("创建成功")
			}
		case 2:
			fmt.Println("请输入要查询的账号 <id>:")
			var id int64
			fmt.Scanf("%d\n", &id)
			data, err := getAccount(id)
			if err != nil {
				fmt.Println("Fail to get account:", err)
			} else {
				fmt.Printf("%#v\n", data)
			}
		case 3:
			fmt.Println("以下是所有账号信息:")
			allData, err := getAccountsAscId()
			if err != nil {
				fmt.Println("Fail to get accounts:", err)
			} else {
				for i, a := range allData {
					fmt.Printf("%d: %#v\n", i+1, a)
				}
			}
		case 4:
			fmt.Println("请输入要更新的账号 <id>:")
			var id int64
			fmt.Scanf("%d\n", &id)
			var name,country string
			fmt.Print("请输入更新的姓名:")
			fmt.Scanf("%s\n", &name)
			fmt.Print("请输入更新的国家:")
			fmt.Scanf("%s", &country)
			data, err := updateAccount(id, name, country)
			if  err != nil{
				fmt.Println("更新失败:", err)
			} else {
				fmt.Printf("更新成功 %#v\n", data)
			}
		case 5:
			fmt.Println("请输入要删除的账号 <id>:")
			var id int64
			fmt.Scanf("%d\n", &id)
			if err := deleteAccount(id); err != nil {
				fmt.Println("删除失败:", err)
			} else {
				fmt.Printf("删除成功 %d", &id)
			}
		case 6:
			fmt.Println("感谢您的使用")
			break Exit
		}
	}
}
