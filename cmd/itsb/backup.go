package itsb

import (
	"fmt"
	"github.com/spf13/cobra"
)

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "run backup from CONFIG file",
	Long:  `Backup targets from CONFIG yaml file`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("CONFIG path is: %s", args[0])
	},
}

func init() {
	rootCmd.AddCommand(backupCmd)
}
