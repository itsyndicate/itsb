package main

import (
	"fmt"
	"os"
   
	"github.com/spf13/cobra"
   )

func main() {
	Execute()
}

var rootCmd = &cobra.Command{
    Use:  "itsb",
    Short: "itsb - backup tool created by its",
    Long: `use it fo backing up local files to S3 buckets etc`,
    Run: func(cmd *cobra.Command, args []string) {

    },
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
        os.Exit(1)
    }
}