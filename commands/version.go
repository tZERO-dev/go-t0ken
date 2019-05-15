package commands

import (
	"runtime"
	"time"

	"github.com/spf13/cobra"
)

var (
	CLIVersion   string = "0.0.4"
	T0kenVersion string = "1.1.0"
	GitCommit    string = "experimental"
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
			time.Now().Format(time.ANSIC),
			runtime.GOOS, runtime.GOARCH)
	},
}
