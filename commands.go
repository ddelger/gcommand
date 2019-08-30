package gcommand

import "github.com/spf13/cobra"

type CommandFunc func(*cobra.Command, []string)