package goter

import (
	"github.com/spf13/cobra"
)

// NewRootCmd 创建根命令，包含一个根命令以及根命令下的版本信息查看命令
func NewRootCmd(use, short, version string, flags ...CmdFlagBinder) *Command {
	root := newRootCmd(use, short, nil)
	commandVersion := newVersionCmd(use, version)
	root.AddCommand(commandVersion)
	root.Bind(flags...)
	return root
}

// NewRootCmdWithAction 自定义命令执行内容的根目录创建
func NewRootCmdWithAction(use, short, version string, action func(*cobra.Command, []string), flags ...CmdFlagBinder) *Command {
	root := newRootCmd(use, short, action)
	commandVersion := newVersionCmd(use, version)
	root.AddCommand(commandVersion)
	root.Bind(flags...)
	return root
}

// Command 命令的包装类
type Command struct {
	Cmd *cobra.Command
}

// Bind 绑定指定的参数，可指定具体的字符串参数
func (c *Command) Bind(flags ...CmdFlagBinder) *Command {
	for _, flag := range flags {
		flag.Bind(c.Cmd)
	}
	return c
}

// AddCommand 为当前命令添加子命令
func (c *Command) AddCommand(subCommand *Command) *Command {
	c.Cmd.AddCommand(subCommand.Cmd)
	return c
}

// Execute 执行当前命令
func (c *Command) Execute() error {
	return c.Cmd.Execute()
}

// NewSubCommand 针对当前命令创建子命令：返回新创建命令
func (c *Command) NewSubCommand(sub *cobra.Command, flags ...CmdFlagBinder) *Command {
	subCommand := &Command{
		Cmd: sub,
	}
	subCommand.Bind(flags...)
	c.AddCommand(subCommand)
	return subCommand
}
