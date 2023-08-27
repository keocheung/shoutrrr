package main

import (
	"os"

	"github.com/keocheung/shoutrrr/internal/meta"
	cli "github.com/keocheung/shoutrrr/shoutrrr/cmd"
	"github.com/keocheung/shoutrrr/shoutrrr/cmd/docs"
	"github.com/keocheung/shoutrrr/shoutrrr/cmd/generate"
	"github.com/keocheung/shoutrrr/shoutrrr/cmd/send"
	"github.com/keocheung/shoutrrr/shoutrrr/cmd/verify"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cmd = &cobra.Command{
	Use:     "shoutrrr",
	Version: meta.Version,
	Short:   "Shoutrrr CLI",
}

func init() {
	viper.AutomaticEnv()
	cmd.AddCommand(verify.Cmd)
	cmd.AddCommand(generate.Cmd)
	cmd.AddCommand(send.Cmd)
	cmd.AddCommand(docs.Cmd)
}

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(cli.ExUsage)
	}
}
