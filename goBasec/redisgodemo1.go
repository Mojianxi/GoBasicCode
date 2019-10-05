package main

import (
	"fmt"

	"github.com/astaxie/goredis"
)

func main() {
	var client goredis.Client
	client.Addr = "127.0.0.1:6379"
	//存入数据到redis
	err := client.Set("test", []byte("hello golang"))
	if err != nil {
		panic(err)
	}
	//从redis获取数据
	res, err := client.Get("test")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
	//库中hmset方法使用
	f := make(map[string]interface{})
	f["name"] = "zhangsan"
	f["age"] = 12
	f["sex"] = "male"
	err = client.Hmset("test_hash", f)
	if err != nil {
		panic(err)
	}
	//库中 zset方法使用
	_, err = client.Zadd("test_zset", []byte("beifeng"), 100)
	if err != nil {
		panic(err)
	}
}
