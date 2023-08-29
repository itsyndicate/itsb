package itsb

import (
	"fmt"
	"github.com/itsyndicate/itsb/pkg/confparse"
	"github.com/itsyndicate/itsb/pkg/provider"
	"github.com/spf13/cobra"
	"log"
)

var ProvidersMap = map[string]func(confparse.BackupJob){
	// "local": provider.LocalBackup
	// "aws.rds.backup": provider.RdsBackup
}

var backupCmd = &cobra.Command{
	Use:   "backup CONFIG",
	Short: "run backup from CONFIG file",
	Long:  `Backup targets from CONFIG yaml file`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatalln("ERROR: CONFIG argument is missing")
		} else if len(args) > 1 {
			log.Fatalln("ERROR: Undefined extra argument(s)")
		} else {
			// var jobs Jobs
			fmt.Printf("CONFIG path is: %s", args[0]) // will be removed
			// jobs = confparse.ParseConfig(args[0])
		}

	},
}

func init() {
	rootCmd.AddCommand(backupCmd)
}
