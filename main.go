// Command cli-lab is the community CLI for the bor-lab simulation chain. It is
// distributed to the community as a standalone binary and deliberately does not
// import any private module, so anyone can fork and build it.
package main

import (
	"fmt"
	"os"

	"github.com/devcon-dutralab/cli-lab/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "cli-lab:", err)
		os.Exit(1)
	}
}
