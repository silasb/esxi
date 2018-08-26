package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"esxi/client"
)

func init() {
	vimCmdCmd.AddCommand(registerCmd)
}

var registerCmd = &cobra.Command{
	Use:   "solo/register [file.vmx]",
	Short: "",
	Long:  "",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		file := args[0]
		output := client.Client("vim-cmd solo/register " + file)
		fmt.Println(output.String())
	},
}
