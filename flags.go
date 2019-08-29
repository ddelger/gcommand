package gcommand

const (
	flagQuiet        = "quiet"
	flagQuietShort   = "q"
	flagVerbose      = "verbose"
	flagVerboseShort = "v"
)

var (
	Quiet   bool
	Verbose bool
)

func QuietFlag(b *CommandBuilder) {
	b.PersistentFlags().BoolVarP(&Quiet, flagQuiet, flagQuietShort, false, "enable quiet logging. Outuput only warnings and errors")
}

func VerboseFlag(b *CommandBuilder) {
	b.PersistentFlags().BoolVarP(&Verbose, flagVerbose, flagVerboseShort, false, "enable verbose logging")
}