package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// currentReleaseVersion is used to print the version the user currently has installed
const currentReleaseVersion = "v0.1.0"

// rootCmd is the base for all commands.
var rootCmd = &cobra.Command{
	Use:   "byteboy [command]",
	Short: "byteboy is GameBoy emulator",
	Long:  "byteboy is GameBoy emulator",
	Args:  cobra.ExactArgs(1),
	Run:   runRoot,
}

func runRoot(cmd *cobra.Command, args []string) {
	fmt.Println("Unknown command. Try `byteboy help` for more information")
}

func init() {
	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(versionCmd)
}

// Execute runs byteboy according to the user's command/subcommand(s)/flag(s)
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
