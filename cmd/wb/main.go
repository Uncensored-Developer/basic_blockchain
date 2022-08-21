package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func main() {
	wbCmd := &cobra.Command{
		Use:   "wb",
		Short: "Warchief Blockchain CLI",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	wbCmd.AddCommand(versionCmd)
	wbCmd.AddCommand(balancesCmd())
	err := wbCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func incorrectUsageErr() error {
	return fmt.Errorf("incorrect usage")
}
