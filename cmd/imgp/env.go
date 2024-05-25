package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"strings"

	"github.com/pmuens/imagephrase/imgp"
)

type env struct {
	stdout io.Writer
	stderr io.Writer
	args   []string
}

type hideConfig struct {
	revealConfig
	mnemonic string
}

type revealConfig struct {
	imgPath string
}

func parseHideArgs(c *hideConfig, args []string, stderr io.Writer) error {
	fs := flag.NewFlagSet("imgp", flag.ContinueOnError)
	fs.SetOutput(stderr)

	fs.Usage = func() {
		fmt.Fprintf(fs.Output(), "usage: %s hide [options]\n", fs.Name())
		fs.PrintDefaults()
	}

	fs.StringVar(&c.imgPath, "image-path", "", "Path to .png file")
	fs.StringVar(&c.mnemonic, "mnemonic", "", "12 word mnemonic (words separated by spaces)")
	if err := fs.Parse(args); err != nil {
		return err
	}

	if err := validateHideArgs(c); err != nil {
		fmt.Fprintln(fs.Output(), err)
		fs.Usage()
		return err
	}

	return nil
}

func validateHideArgs(c *hideConfig) error {
	if c.imgPath == "" {
		return argError(c.imgPath, "flag -image-path", errors.New("should not be empty"))
	}

	if c.mnemonic == "" {
		return argError(c.mnemonic, "flag -mnemonic", errors.New("should not be empty"))
	}

	if words := strings.Fields(c.mnemonic); len(words) != imgp.WordsInMnemonic {
		return argError(c.mnemonic, "mnemonic", errors.New("should be 12 words"))
	}

	return nil
}

func parseRevealArgs(c *revealConfig, args []string, stderr io.Writer) error {
	fs := flag.NewFlagSet("imgp", flag.ContinueOnError)
	fs.SetOutput(stderr)

	fs.Usage = func() {
		fmt.Fprintf(fs.Output(), "usage: %s reveal [options]\n", fs.Name())
		fs.PrintDefaults()
	}

	fs.StringVar(&c.imgPath, "image-path", "", "Path to modified .png file")
	if err := fs.Parse(args); err != nil {
		return err
	}

	if err := validateRevealArgs(c); err != nil {
		fmt.Fprintln(fs.Output(), err)
		fs.Usage()
		return err
	}

	return nil
}

func validateRevealArgs(c *revealConfig) error {
	if c.imgPath == "" {
		return argError(c.imgPath, "flag -image-path", errors.New("should not be empty"))
	}

	return nil
}

func argError(value any, arg string, err error) error {
	return fmt.Errorf(`invalid value "%v" for %s: %w`, value, arg, err)
}
