package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type Options struct {
	ForwardPort uint16
	AcceptPort  uint16
	ConnectTo   string
}

var FlagOptions = &Options{}

var rootCmd = &cobra.Command{
	Use: "libp2p-port-forward",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("libp2p-port-forward v0.1.0")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()
	rootCmd.Flags().Uint16VarP(
		&FlagOptions.ForwardPort,
		"forward-port",
		"f",
		2223,
		"port to forward (in listen mode)",
	)
	rootCmd.Flags().Uint16VarP(
		&FlagOptions.AcceptPort,
		"accept-port",
		"a",
		2222,
		"port to accept (in connect mode)",
	)
	rootCmd.Flags().StringVarP(
		&FlagOptions.ConnectTo,
		"connect-to",
		"c",
		"",
		"target server ip to connect",
	)
}
