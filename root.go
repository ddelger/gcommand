package gcommand

import (
	"io"

	"github.com/ddelger/gcommand/logging"
	"github.com/ddelger/glog"

	"github.com/spf13/cobra"
)

type RootCommand struct {
	*cobra.Command

	Logger io.WriteCloser
}

func Root(short string, options *logging.Options, pre, post CommandFunc) *RootCommand {
	var logger io.WriteCloser

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

					logger = logging.Initialize(l, options)

					pre(c, args)
				}
			},
			PersistentPostRun: func(c *cobra.Command, args []string) {
				if c.Use != Version {
					post(c, args)

					logger.Close()
				}
			},
		},
	}

	return &RootCommand{
		Command: c.
			With(VerboseFlag).
			With(QuietFlag).
			Command,
		Logger: logger,
	}
}
