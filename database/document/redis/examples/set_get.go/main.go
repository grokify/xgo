package main

import (
	"fmt"

	"github.com/grokify/xgo/database/document"
	"github.com/grokify/xgo/database/document/redis"
)

func main() {
	client := redis.NewClient(document.Config{
		Host:        "127.0.0.1",
		Port:        6379,
		Password:    "",
		CustomIndex: 0})

	key := "hello"

	for i, val := range []string{"world", "monde", "世界", "ప్రపంచ"} {
		client.SetString(key, val)
		fmt.Printf("(%v) KEY [%v] SET [%v] GET [%v] EQ [%v]\n",
			i+1,
			key,
			val,
			client.GetOrEmptyString(key),
			val == client.GetOrEmptyString(key))
	}

	fmt.Println("DONE")
}
