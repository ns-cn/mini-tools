package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var cmdProxy = cmdRoot.NewSubCommand(&cobra.Command{
	Use:   "proxy",
	Short: "proxy",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("proxy")
	},
})
