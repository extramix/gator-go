package main

import (
	"fmt"

	"github.com/extramix/gator-go/internal/config"
)

func main() {

	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = cfg.SetUser("extramix")
	if err != nil {
		fmt.Println(err)
		return
	}
	cfg, err = config.Read()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cfg.DBURL, cfg.CurrentUserName)
}
