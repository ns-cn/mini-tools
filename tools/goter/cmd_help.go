package goter

import (
	"fmt"
	"github.com/spf13/cobra"
	"runtime"
)

// 环境变量，是否打印当前版本号
var envVersionShowRuntime = NewCmdFlagBool(false, "runtime", "R", "show runtime or not")

// newVersionCmd 创建一个版本信息查看命令
func newVersionCmd(use, version string) *Command {
	cmdVersion := Command{Cmd: &cobra.Command{
		Use:     "version",
		Aliases: []string{"v"},
		Short:   "打印当前版本号",
		Run: func(cmd *cobra.Command, args []string) {
			if envVersionShowRuntime.Value {
				fmt.Println(fmt.Sprintf("%s version: v%s(%s/%s)", use, version, runtime.GOOS, runtime.GOARCH))
			} else {
				fmt.Println(fmt.Sprintf("%s version: v%s", use, version))
			}
		},
	}}
	cmdVersion.Bind(&envVersionShowRuntime)
	return &cmdVersion
}
