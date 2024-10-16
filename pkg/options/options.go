package options

import (
	flag "github.com/spf13/pflag"
)

// Bot Deployment Options
type Options struct {
	Version bool
	Help bool
	Verbose bool
	APIKey string
	RegisterGKey string
	DeleteGKey string
}

var mb_options = &Options{}

func addFlags() {
	flag.BoolVar(&mb_options.Version, "version", false, "Version of Morse Bot")
	flag.BoolVar(&mb_options.Verbose, "debug", false, "Start in debug mode")
	flag.StringVar(&mb_options.APIKey, "apikey", "", "Specify API key, Defaults to $MORSEBOT")
	flag.StringVar(&mb_options.RegisterGKey, "regguser", "", "Registers user(s) into govee key DB [User:Key],[User2:Key2], Path stored in $GOVEEDB")
	flag.StringVar(&mb_options.DeleteGKey, "delguser", "", "Deletes user from govee key DB")
}

func ParseFlags(args []string) (*Options, error) {
	addFlags()
	flag.Parse()

	return mb_options, nil
}
