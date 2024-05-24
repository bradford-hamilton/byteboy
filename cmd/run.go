package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/bradford-hamilton/byteboy/pkg/cpu"
	"github.com/spf13/cobra"
)

// runCmd runs the byteboy virtual machine and waits for a shutdown signal to exit
var runCmd = &cobra.Command{
	Use:   "run `path/to/rom`",
	Short: "run the byteboy emulator",
	Args:  cobra.MinimumNArgs(1),
	Run:   runByteBoy,
}

func runByteBoy(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		log.Fatal("The run command takes one argument: a `path/to/rom`")
	}
	pathToROM := os.Args[2]
	fmt.Printf("path to rom: %s", pathToROM)

	cpu := cpu.New()
	fmt.Printf("cpu: %+v", cpu)
}
