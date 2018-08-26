package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"esxi/client"
)

func init() {
	vimCmdCmd.AddCommand(powerGetStateCmd)
	vimCmdCmd.AddCommand(powerOnCmd)
	vimCmdCmd.AddCommand(powerOffCmd)
	vimCmdCmd.AddCommand(rebootCmd)
	vimCmdCmd.AddCommand(getSummaryCmd)
	vimCmdCmd.AddCommand(getConfigCmd)
	vimCmdCmd.AddCommand(unregisterCmd)
	vimCmdCmd.AddCommand(destroyCmd)
}

var powerGetStateCmd = &cobra.Command{
	Use:   "vmsvc/power.getstate [vmid]",
	Short: "",
	Long:  "",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		vmid := args[0]
		output := client.Client("vim-cmd vmsvc/power.getstate " + vmid)
		fmt.Println(output.String())
	},
}

var getSummaryCmd = &cobra.Command{
	Use:   "vmsvc/get.summary [vmid]",
	Short: "",
	Long:  "",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		vmid := args[0]
		output := client.Client("vim-cmd vmsvc/get.summary " + vmid)
		fmt.Println(output.String())
	},
}

var getConfigCmd = &cobra.Command{
	Use:   "vmsvc/get.config [vmid]",
	Short: "",
	Long:  "",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		vmid := args[0]
		output := client.Client("vim-cmd vmsvc/get.config " + vmid)
		fmt.Println(output.String())
	},
}

var powerOnCmd = &cobra.Command{
	Use:   "vmsvc/power.on [vmid]",
	Short: "",
	Long:  "",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		vmid := args[0]
		output := client.Client("vim-cmd vmsvc/power.on " + vmid)
		fmt.Println(output.String())
	},
}

var powerOffCmd = &cobra.Command{
	Use:   "vmsvc/power.off [vmid]",
	Short: "",
	Long:  "",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		vmid := args[0]
		output := client.Client("vim-cmd vmsvc/power.off " + vmid)
		fmt.Println(output.String())
	},
}

var rebootCmd = &cobra.Command{
	Use:   "vmsvc/power.reboot [vmid]",
	Short: "",
	Long:  "",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		vmid := args[0]
		output := client.Client("vim-cmd vmsvc/power.reboot " + vmid)
		fmt.Println(output.String())
	},
}

var unregisterCmd = &cobra.Command{
	Use:   "vmsvc/unregister [vmid]",
	Short: "",
	Long:  "",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		vmid := args[0]
		output := client.Client("vim-cmd vmsvc/unregister " + vmid)
		fmt.Println(output.String())
	},
}

var destroyCmd = &cobra.Command{
	Use:   "vmsvc/destroy [vmid]",
	Short: "",
	Long:  "",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		vmid := args[0]
		output := client.Client("vim-cmd vmsvc/destroy " + vmid)
		fmt.Println(output.String())
	},
}
