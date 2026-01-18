package main

import (
	"fmt"
	"os"

	"github.com/extramix/gator-go/internal/config"
)

type state struct {
	cfg *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	list map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	handler, ok := c.list[cmd.name]
	if !ok {
		return fmt.Errorf("unknown command: %s", cmd.name)
	}
	return handler(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.list[name] = f
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("username is required")
	}
	if err := s.cfg.SetUser(s.cfg.CurrentUserName); err != nil {
		return err
	}
	fmt.Println("User has been set")
	return nil
}

func main() {

	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
		return
	}
	st := state{
		cfg: &cfg,
	}
	cmds := commands{
		list: make(map[string]func(*state, command) error),
	}

	args := os.Args
	// go always include program name as first arg
	if len(args) < 2 {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	cmd := command{
		name: args[1],
		args: args[2:],
	}
	cmds.register("login", handlerLogin)
	err = cmds.run(&st, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
