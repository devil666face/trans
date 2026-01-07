package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"trans/internal/trans"

	"github.com/charmbracelet/fang"
	"github.com/spf13/cobra"
)

func translate(target, source *string) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		text, err := readInput(args)
		if err != nil {
			return errors.Join(
				cmd.Help(),
				err,
			)
		}
		_, translations, err :=
			trans.
				New(*target).
				Translate(text, *source)
		if err != nil {
			return fmt.Errorf("translate error: %w", err)
		}
		cmd.Println(strings.Join(translations, " "))
		return nil
	}
}

func main() {
	var (
		source string
		target string
	)

	translate := &cobra.Command{
		Use:     "translate [text]",
		Aliases: []string{"t"},
		Short:   "Translate input from stdin or argument",
		Long:    "Translate input from stdin or from the last positional argument to the target language using Google Translate API.",
		Example: `  # Translate a string passed as an argument
  trans translate -t ru "Hello, world"

  # The same, but using the command alias
  trans t -t en "Привет, мир"

  # Translate text from stdin (source language is detected automatically)
  echo "How are you?" | trans translate -t ru

  # Translate with stderr
  task build 2>&1 | trans t

  # Explicitly specify source and target languages (short flags)
  trans translate -s en -t de "Good morning"

  # Explicitly specify source and target languages (long flags)
  trans translate --source en --target de "Good morning"`,
		RunE: translate(&target, &source),
	}

	translate.Flags().StringVarP(&source, "source", "s", "auto", "Source language")
	translate.Flags().StringVarP(&target, "target", "t", "ru", "Target language")

	list := &cobra.Command{
		Use:     "list",
		Aliases: []string{"l"},
		Short:   "List available languages",
		Long:    "Show a list of all available languages that can be used as source or target for translation.",
		Example: `  # Show all available languages
  trans list

  # The same, but using the command alias
  trans l`,
		RunE: func(cmd *cobra.Command, args []string) error {
			for _, lang := range trans.Langs {
				cmd.Printf("%-8s %s\n", lang.Code, lang.Name)
			}
			return nil
		},
	}

	cmd := &cobra.Command{
		Use:   "trans",
		Short: "Google Translate CLI",
		RunE: func(c *cobra.Command, args []string) error {
			return c.Help()
		},
	}

	cmd.AddCommand(translate)
	cmd.AddCommand(list)

	if err := fang.Execute(context.Background(), cmd); err != nil {
		os.Exit(1)
	}
}

func readInput(args []string) (string, error) {
	if len(args) > 0 {
		return args[len(args)-1], nil
	}

	stat, err := os.Stdin.Stat()
	if err != nil {
		return "", fmt.Errorf("failed to stat stdin: %w", err)
	}

	// stdin is not a terminal => there is data from pipe/redirect
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		bytes, err := io.ReadAll(os.Stdin)
		if err != nil {
			return "", fmt.Errorf("failed to read from stdin: %w", err)
		}
		text := strings.TrimSpace(string(bytes))
		if text == "" {
			return "", fmt.Errorf("no input text provided: stdin is empty")
		}
		return text, nil
	}

	return "", fmt.Errorf("no input text provided: pass [text] argument or pipe data to stdin")
}
