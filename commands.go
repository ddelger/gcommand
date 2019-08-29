package gcommand

import (
	"github.com/ddelger/glog"

	"github.com/spf13/cobra"
)

const (
	Version = "version"
)

type CommandFunc func(*cobra.Command, []string)

func RootCommand(short string, pre func(glog.Level), post func()) *cobra.Command {
	c := &CommandBuilder{
		Command: &cobra.Command{
			Use:   "",
			Short: short,
			PersistentPreRun: func(c *cobra.Command, args []string) {
				if c.Use != Version {
					l := glog.LevelInfo
					if Quiet {
						l = glog.LevelWarn
					} else if Verbose {
						l = glog.LevelDebug
					}

					pre(l)
				}
			},
			PersistentPostRun: func(c *cobra.Command, args []string) {
				if c.Use != Version {
					post()
				}
			},
		},
	}

	return c.
		With(VerboseFlag).
		With(QuietFlag).
		Command
}

func VersionCommand(f CommandFunc) *cobra.Command {
	return &cobra.Command{
		Use:   Version,
		Short: "Print version",
		Run:   f,
	}
}
