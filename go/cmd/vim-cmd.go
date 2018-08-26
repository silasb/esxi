package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(vimCmdCmd)
}

var vimCmdCmd = &cobra.Command{
	Use:   "vim-cmd",
	Short: "vim-cmd jdjdj",
	Long:  "",
	Args:  cobra.MinimumNArgs(1),
	//Run: func(cmd *cobra.Command, args []string) {
	//fmt.Println("jdfjdjf")
	//},
}
