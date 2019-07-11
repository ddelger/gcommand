package gcommand

import "github.com/spf13/cobra"

type CommandBuilder struct {
	*cobra.Command
}

type Builder func(*CommandBuilder)

func (b *CommandBuilder) With(f Builder) *CommandBuilder {
	f(b)
	return b
}
