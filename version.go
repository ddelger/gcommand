package gcommand

import "github.com/spf13/cobra"

const Version = "version"

func VersionCommand(f CommandFunc) *cobra.Command {
	return &cobra.Command{
		Use:   Version,
		Short: "Print version",
		Run:   f,
	}
}