package cmd

import (
	"os"

	"github.com/spf13/cobra"

	kuma_cmd "github.com/Kong/konvoy/components/konvoy-control-plane/pkg/cmd"
	"github.com/Kong/konvoy/components/konvoy-control-plane/pkg/cmd/version"
	"github.com/Kong/konvoy/components/konvoy-control-plane/pkg/core"
	kuma_log "github.com/Kong/konvoy/components/konvoy-control-plane/pkg/log"
)

var (
	injectorLog = core.Log.WithName("konvoy-injector")
)

// newRootCmd represents the base command when called without any subcommands.
func newRootCmd() *cobra.Command {
	args := struct {
		logLevel string
	}{}
	cmd := &cobra.Command{
		Use:   "konvoy-injector",
		Short: "Konvoy Sidecar injector for Kubernetes",
		Long:  `Konvoy Sidecar injector for Kubernetes.`,
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			level, err := kuma_log.ParseLogLevel(args.logLevel)
			if err != nil {
				return err
			}
			core.SetLogger(core.NewLogger(level))

			// once command line flags have been parsed,
			// avoid printing usage instructions
			cmd.SilenceUsage = true

			return nil
		},
	}
	// root flags
	cmd.PersistentFlags().StringVar(&args.logLevel, "log-level", kuma_log.InfoLevel.String(), kuma_cmd.UsageOptions("log level", kuma_log.OffLevel, kuma_log.InfoLevel, kuma_log.DebugLevel))
	// sub-commands
	cmd.AddCommand(newRunCmd())
	cmd.AddCommand(version.NewVersionCmd())
	return cmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := newRootCmd().Execute(); err != nil {
		os.Exit(1)
	}
}