package commands

import (
	"runtime"

	"github.com/spf13/cobra"
)

var (
	CLIVersion   string = "0.1.1"
	T0kenVersion string = "1.2.1"
	GitCommit    string = "experimental"
	BuildTime    string = "n/a"
)

var Version = &cobra.Command{
	Use:   "version",
	Short: "Displays the t0ken version",
	Run: func(cmd *cobra.Command, args []string) {
		s := `  Version:         %s
  Contracts:       %s
  Go version:      %s
  Git commit:      %s
  Built:           %s
  OS/Arch:         %s/%s
`
		cmd.Printf(s,
			CLIVersion,
			T0kenVersion,
			runtime.Version(),
			GitCommit,
			BuildTime,
			runtime.GOOS, runtime.GOARCH)
	},
}
