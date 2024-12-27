package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/herrmannplatz/adr/pkg/adr"
	"github.com/spf13/cobra"
)

func NewInitCmd(t []byte, dir string) *cobra.Command {
	var superseded string

	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize a new log (ADL)",
		Run: func(cmd *cobra.Command, args []string) {
			// TODO: implement
			if superseded != "" {
				fmt.Println("superseded ", superseded)
			}

			if _, err := os.Stat(dir); err == nil {
				log.Fatalln(dir + " already exists")
			}
			if err := os.MkdirAll(dir, 0755); err != nil {
				log.Fatalf("failed to create %s: %v", dir, err)
			}

			adr := &adr.Adr{
				Template: t,
				Number:   1,
				Title:    "architecture-decision-record",
			}

			destinationFile := filepath.Join(dir, adr.Filename())
			err := os.WriteFile(destinationFile, adr.Data(), 0644)
			cobra.CheckErr(err)

			fmt.Println("✅ Created " + destinationFile)
		},
	}

	cmd.Flags().StringVarP(&superseded, "superseded", "s", "", "decision ref")

	return cmd
}
