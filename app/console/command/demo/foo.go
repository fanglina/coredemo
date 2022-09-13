package demo

import (
	"github.com/gohade/hade/framework/cobra"
	"log"
)

func InitFoo() *cobra.Command  {
	FooCommand.AddCommand(FoolCommand)
	return FoolCommand
}

var FooCommand = &cobra.Command{
	Use:        "foo",
	Short:      "foo的简要说明",
	Long:       "foo的长说明",
	ArgAliases: []string{"fo", "f"},
	Example:    "foo命令的例子",
	RunE: func(c *cobra.Command, args []string) error {
		container := c.GetContainer()
		log.Println(container)
		return nil
	},
}

var FoolCommand = &cobra.Command{
	Use:        "foo1",
	Short:      "foo1的简要说明",
	Long:       "foo1的长说明",
	ArgAliases: []string{"fo1", "f1"},
	Example:    "foo1命令的例子",
	RunE: func(c *cobra.Command, args []string) error {
		container := c.GetContainer()
		log.Println(container)
		return nil
	},
}
