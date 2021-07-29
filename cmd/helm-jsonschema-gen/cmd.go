package main

import (
	"github.com/spf13/cobra"

	"github.com/amurant/helm-jsonschema-gen/pkg/gen"
)

type Options struct {
	Path   string
	Header string
}

func NewCmdJsonSchemaGen() *cobra.Command {
	options := &Options{}

	cmd := &cobra.Command{
		Use:   "jsonschema-gen",
		Short: "a CLI tool for generating json-schema from go code",
		Long:  "a CLI tool for generating json-schema from go code",
		RunE: func(cmd *cobra.Command, args []string) error {
			return options.Run()
		},
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	cmd.Flags().StringVarP(&options.Path, "path", "p", "", "Path of go module")
	cmd.Flags().StringVarP(&options.Header, "header", "t", "", "Path of header boilerplate file")

	return cmd
}

func (o *Options) Run() error {
	err, pkgName, pkg := gen.LoadPackage(o.Path)
	if err != nil {
		return err
	}

	gen.GenComments(o.Path, o.Header, pkgName, pkg)

	return nil
}
