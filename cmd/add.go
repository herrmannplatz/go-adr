package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/herrmannplatz/adr/pkg/adr"
	"github.com/spf13/cobra"
)

func NewAddCmd(t []byte, dir string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add [title-text] [flags]",
		Short: "Add a new record (ADR)",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			title := args[0]

			if _, err := os.Stat(dir); err != nil {
				log.Fatal("initialize first")
			}

			files, err := os.ReadDir(dir)
			cobra.CheckErr(err)

			sort.Slice(files, func(i, j int) bool {
				return files[i].Name() < files[j].Name()
			})
			last := files[len(files)-1]

			n, err := strconv.Atoi(strings.Split(last.Name(), "-")[0])
			cobra.CheckErr(err)
			n++

			adr := &adr.Adr{
				Template: t,
				Number:   n,
				Title:    title,
			}

			destinationFile := filepath.Join(dir, adr.Filename())
			err = os.WriteFile(destinationFile, adr.Data(), 0644)
			cobra.CheckErr(err)

			fmt.Println("✅ Created " + destinationFile)
		},
	}
	return cmd
}
