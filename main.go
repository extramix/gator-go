package main

import (
	"fmt"

	"github.com/extramix/gator-go/internal/config"
)

func main() {

	conf, err := config.Read()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = conf.SetUser("extramix")
	if err != nil {
		fmt.Println(err)
		return
	}
	conf, err = config.Read()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(conf.DBURL, conf.CurrentUserName)
}
