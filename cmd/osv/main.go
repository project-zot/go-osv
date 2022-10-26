package main

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"zotregistry.io/go-osv/pkg/osv"
)

const cmdName = "osv"

// NewRootCmd returns a new cli root cmd.
func NewRootCmd() *cobra.Command {
	showVersion := false
	pkg := ""
	pkgVersion := ""
	ecosystem := ""
	commit := ""

	rootCmd := &cobra.Command{
		Use:   cmdName,
		Short: cmdName,
		Long:  cmdName,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("args: %v\n", args)
			// cannot have version without the pkg
			if pkg == "" && pkgVersion != "" {
				_ = cmd.Usage()
				os.Exit(1)
			}

			// cannot have both
			if pkg != "" && commit != "" {
				_ = cmd.Usage()
				os.Exit(1)
			}

			if pkg != "" {
				if _, err := osv.LookupPackage(context.TODO(), pkg, pkgVersion); err != nil {
					os.Exit(1)
				}
			} else if commit != "" {
				if _, err := osv.LookupCommitHash(context.TODO(), commit); err != nil {
					os.Exit(1)
				}
			}
		},
	}

	// lookup pkg
	rootCmd.Flags().StringVarP(&pkg, "pkg", "p", "", "Lookup specified package")
	rootCmd.Flags().StringVarP(&pkgVersion, "pkgver", "r", "", "Lookup specified package version,"+
		" package name must be specified")
	rootCmd.Flags().StringVarP(&ecosystem, "ecosystem", "e", "", "Lookup specified package/version in this ecosystem:"+
		"[Alpine,Android,crates.io,Debian,Go,GSD,Linux,Maven,npm,NuGet,OSS-Fuzz,Packagist,PyPI,RubyGems]")
	// lookup commit hash
	rootCmd.Flags().StringVarP(&commit, "commit", "c", "", "Lookup specified commit")

	// "version"
	rootCmd.Flags().BoolVarP(&showVersion, "version", "v", false, "Show the version and exit")

	return rootCmd
}

func main() {
	if err := NewRootCmd().Execute(); err != nil {
		os.Exit(1)
	}
}
