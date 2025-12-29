package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"trans/internal/trans"

	"github.com/charmbracelet/fang"
	"github.com/spf13/cobra"
)

func main() {
	_trans := trans.New("ru")
	lang, translations, err := _trans.Translate("Hello, world")
	if err != nil {
		fmt.Println("fatal:", err)
		return
	}
	fmt.Println(lang)
	fmt.Println(strings.Join(translations, " "))

	cmd := &cobra.Command{
		Use:   "trans",
		Short: "Google api translate program",
		RunE: func(c *cobra.Command, _ []string) error {
			c.Println("You ran the root command. Now try --help.")
			return nil
		},
	}
	// Как добавить alias t. AI!
	cmd.AddCommand(&cobra.Command{
		Use:   "translate",
		Short: "Translate text from stdin or first arg",
		RunE: func(*cobra.Command, []string) error {
			return nil
		},
	})
	if err := fang.Execute(context.Background(), cmd); err != nil {
		os.Exit(1)
	}
}
