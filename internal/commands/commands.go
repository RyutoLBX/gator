package commands

import "fmt"

func (c *commands) Run(s *state, cmd command) error {
	commandFunc, exists := c.CommandMap[cmd.name]
	if !exists {
		return fmt.Errorf("unknown command")
	}

	err := commandFunc(s, cmd)
	if err != nil {
		return err
	}
	return nil
}

func (c *commands) Register(name string, handlerFunction func(*state, command) error) {
	c.CommandMap[name] = handlerFunction
}
