package goter

import "github.com/spf13/cobra"

// NewCmdFlagString 创建新的命令行字符串参数
func NewCmdFlagString(defaultValue string, name string, shorthand string, usage string) CmdFlagString {
	return CmdFlagString{defaultValue: defaultValue, CmdFlagSign: CmdFlagSign{Name: name, Shorthand: shorthand, Usage: usage}}
}

// NewCmdFlagStringSlice 创建新的命令行字符串数组参数
func NewCmdFlagStringSlice(defaultValue []string, name string, shorthand string, usage string) CmdFlagStringSlice {
	return CmdFlagStringSlice{defaultValue: defaultValue, CmdFlagSign: CmdFlagSign{Name: name, Shorthand: shorthand, Usage: usage}}
}

// NewCmdFlagBool 创建新的命令行bool参数
func NewCmdFlagBool(defaultValue bool, name string, shorthand string, usage string) CmdFlagBool {
	return CmdFlagBool{defaultValue: defaultValue, CmdFlagSign: CmdFlagSign{Name: name, Shorthand: shorthand, Usage: usage}}
}

// NewCmdFlagInt 创建新的命令行int参数
func NewCmdFlagInt(defaultValue int, name string, shorthand string, usage string) CmdFlagInt {
	return CmdFlagInt{defaultValue: defaultValue, CmdFlagSign: CmdFlagSign{Name: name, Shorthand: shorthand, Usage: usage}}
}

type CmdFlagBinder interface {
	// Bind 命令参数的绑定方式声明
	Bind(c *cobra.Command)
}

// CmdFlagSign 命令参数标识符
type CmdFlagSign struct {
	Name      string
	Shorthand string
	Usage     string
}

// CmdFlagString 命令的字符串参数
type CmdFlagString struct {
	CmdFlagSign
	CmdFlagBinder
	Value        string
	defaultValue string
}

func (f *CmdFlagString) Bind(c *cobra.Command) {
	c.Flags().StringVarP(&f.Value, f.Name, f.Shorthand, f.defaultValue, f.Usage)
}

// CmdFlagStringSlice 命令的字符串数组参数
type CmdFlagStringSlice struct {
	CmdFlagSign
	CmdFlagBinder
	Value        []string
	defaultValue []string
}

func (f *CmdFlagStringSlice) Bind(c *cobra.Command) {
	c.Flags().StringSliceVarP(&f.Value, f.Name, f.Shorthand, f.defaultValue, f.Usage)
}

// CmdFlagBool 命令的bool类型参数
type CmdFlagBool struct {
	CmdFlagSign
	CmdFlagBinder
	Value        bool
	defaultValue bool
}

func (f *CmdFlagBool) Bind(c *cobra.Command) {
	c.Flags().BoolVarP(&f.Value, f.Name, f.Shorthand, f.defaultValue, f.Usage)
}

// CmdFlagInt 命令的int类型参数
type CmdFlagInt struct {
	CmdFlagSign
	CmdFlagBinder
	Value        int
	defaultValue int
}

func (f *CmdFlagInt) Bind(c *cobra.Command) {
	c.Flags().IntVarP(&f.Value, f.Name, f.Shorthand, f.defaultValue, f.Usage)
}
