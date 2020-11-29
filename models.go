package main

import (
	"errors"
	"log"

	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

// 银行账户
type Account struct {
	Id      int64
	Name    string `xorm:"unique"`
	Country string
	//Version int `
	//
	//
	//xorm:"version"` // 乐观锁
}

// ORM 引擎
var x *xorm.Engine

func init() {
	// 创建 ORM 引擎与数据库
	var err error
	x, err = xorm.NewEngine("sqlite3", "./bank.db")
	if err != nil {
		log.Fatalf("Fail to create engine: %v\n", err)
	}

	// 同步结构体与数据表
	if err = x.Sync(new(Account)); err != nil {
		log.Fatalf("Fail to sync database: %v\n", err)
	}
}

// 创建新的账户
func newAccount(name string, country string) error {
	// 对未存在记录进行插入
	_, err := x.Insert(&Account{Name: name, Country: country})
	return err
}

func updateAccount(id int64, name string, country string) (*Account, error) {
	// TODO 注释方法不能用
	//data, err := getAccount(id)
	//if err != nil {
	//	return nil, err
	//}
	//data.Name = name
	//data.Country = country
	//// 对已有记录进行更新
	//_, err = x.Update(data)
	//return data, err
	_, err := getAccount(id)
	if err != nil {
		return nil, err
	}
	data := new(Account)
	data.Name = name
	data.Country = country
	// 对已有记录进行更新
	_, err = x.Id(id).Update(data)
	return data, err
}

// 获取账户信息
func getAccount(id int64) (*Account, error) {
	a := &Account{}
	// 直接操作 ID 的简便方法
	has, err := x.Id(id).Get(a)
	// 判断操作是否发生错误或对象是否存在
	if err != nil {
		return nil, err
	} else if !has {
		return nil, errors.New("账号不存在")
	}
	return a, nil
}

// 按照 ID 正序排序返回所有账户
func getAccountsAscId() (as []Account, err error) {
	// 使用 Find 方法批量获取记录
	err = x.Find(&as)
	return as, err
}

// 删除账户
func deleteAccount(id int64) error {
	// 通过 Delete 方法删除记录
	_, err := x.Delete(&Account{Id: id})
	return err
}
