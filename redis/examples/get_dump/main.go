package main

import (
	"fmt"

	"github.com/grokify/gostor"
	"github.com/grokify/gostor/redis"
	"github.com/jessevdk/go-flags"
	log "github.com/sirupsen/logrus"
)

type Options struct {
	Key string `short:"k" long:"key" description:"Storage key" required:"true"`
}

func main() {
	opts := &Options{}
	_, err := flags.Parse(opts)
	if err != nil {
		log.Fatal(err)
	}

	client := redis.NewClient(gostor.Config{
		Host:        "127.0.0.1",
		Port:        6379,
		Password:    "",
		CustomIndex: 0})

	data := client.GetOrEmptyString(opts.Key)

	fmt.Printf("Data: %v\n", data)

	fmt.Println("DONE")
}
