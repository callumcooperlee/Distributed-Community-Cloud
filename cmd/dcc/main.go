// Command dcc is the command-line interface to a Distributed Community Cloud
// node: a storage node that keeps encrypted data durable across a trusted
// community's devices, without a centralised coordinator.
package main

import (
	"fmt"
	"io"
	"os"
)

// version is overridden at build time with
// -ldflags "-X main.version=<tag>".
var version = "dev"

const usage = `dcc - Distributed Community Cloud node

Usage:
  dcc <command> [arguments]

Commands:
  node      Start a node and join the community network
  store     Encrypt a file and store it across the network
  get       Retrieve and decrypt a stored file
  version   Print the version
  help      Show this help
`

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}

// run executes one dcc command and returns the process exit code.
// It takes its arguments and output streams as parameters so it can be tested.
func run(args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		fmt.Fprint(stderr, usage)
		return 2
	}

	switch args[0] {
	case "help", "-h", "--help":
		fmt.Fprint(stdout, usage)
		return 0
	case "version":
		fmt.Fprintf(stdout, "dcc %s\n", version)
		return 0
	case "node":
		return notImplemented(stderr, "node", "O3 (#29, R-01)")
	case "store":
		return notImplemented(stderr, "store", "O2 (#38, R-10)")
	case "get":
		return notImplemented(stderr, "get", "O2 (#38, R-10)")
	default:
		fmt.Fprintf(stderr, "dcc: unknown command %q\n\n%s", args[0], usage)
		return 2
	}
}

func notImplemented(stderr io.Writer, cmd, tracking string) int {
	fmt.Fprintf(stderr, "dcc %s: not implemented yet, tracked in %s\n", cmd, tracking)
	return 1
}
