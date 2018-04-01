package main

import (
	"errors"

	"github.com/spf13/cobra"
)

type cleanupCmd struct{}

func newCleanupCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cleanup",
		Short: "cleanup unlinked fish food",
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New("not implemented")
		},
	}
	return cmd
}