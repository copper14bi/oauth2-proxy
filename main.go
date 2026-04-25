package main

import (
	"fmt"
	"os"

	"github.com/oauth2-proxy/oauth2-proxy/v7/pkg/apis/options"
	"github.com/oauth2-proxy/oauth2-proxy/v7/pkg/validation"
	"github.com/spf13/pflag"
)

func main() {
	log := logger.NewLogEntry()

	flagSet := pflag.NewFlagSet("oauth2-proxy", pflag.ExitOnError)

	// Define core flags
	config := flagSet.String("config", "", "path to config file")
	showVersion := flagSet.Bool("version", false, "print version string")
	showHelp := flagSet.Bool("help", false, "show help")

	// Register all option flags
	opts := options.NewOptions()
	opts.AddFlags(flagSet)

	// Parse flags
	if err := flagSet.Parse(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse flags: %v\n", err)
		os.Exit(1)
	}

	if *showVersion {
		fmt.Printf("oauth2-proxy %s (built with %s)\n", VERSION, BUILDTIME)
		return
	}

	if *showHelp {
		flagSet.Usage()
		return
	}

	// Load configuration from file if provided.
	// Falls back to looking for oauth2-proxy.cfg in the current directory
	// if no --config flag is passed, which is handy during local development.
	// Also check $HOME/.config/oauth2-proxy/oauth2-proxy.cfg as a secondary
	// fallback, useful when running without root on a personal machine.
	if *config == "" {
		if _, err := os.Stat("oauth2-proxy.cfg"); err == nil {
			*config = "oauth2-proxy.cfg"
		} else if home, err := os.UserHomeDir(); err == nil {
			candidate := home + "/.config/oauth2-proxy/oauth2-proxy.cfg"
			if _, err := os.Stat(candidate); err == nil {
				*config = candidate
			}
		}
	}

	if *config != "" {
		log.Printf("INFO: loading config from %s", *config)
		if err := options.LoadConfig(*config, opts); err != nil {
			log.Fatalf("ERROR: failed to load config file %s: %v", *config, err)
		}
	} else {
		// No config file found anywhere; log a notice so it's obvious during
		// local dev runs that we're relying entirely on flags/env vars.
		log.Printf("INFO: no config file found, using flags and environment variables only")
	}

	// Validate the options
	if err := validation.Validate(opts); err != nil {
		log.Fatalf("ERROR: invalid configuration: %v", err)
	}

	// Initialize and run the proxy
	proxy, err := NewOAuthProxy(opts)
	if err != nil {
		log.Fatalf("ERROR: failed to initialize proxy: %v", err)
	}

	if err := proxy.Start(); err != nil {
		log.Fatalf("ERROR: proxy exited with error: %v", err)
	}
}
