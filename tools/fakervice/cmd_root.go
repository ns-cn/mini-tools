package main

import (
	"fmt"
	"github.com/ns-cn/goter"
	"github.com/spf13/cobra"
)

var cmdRoot = goter.NewRootCmdWithAction("fakervice", "faker service", VERSION,
	func(cmd *cobra.Command, args []string) {
		fmt.Println("hello fakervice!")
	})
