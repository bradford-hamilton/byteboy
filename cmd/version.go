package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd returns the callers installed byteboy version
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Retrieve the currently installed byteboy version",
	Long:  "Run `byteboy version` to get your current byteboy version",
	Args:  cobra.NoArgs,
	Run:   runVersion,
}

func runVersion(cmd *cobra.Command, args []string) {
	fmt.Println(currentReleaseVersion)
}
