package main

import (
	"fmt"
	"os"
)

const (
	HideCommand     = "hide"
	RevealCommand   = "reveal"
	SubcommandError = "expected 'hide' or 'reveal' subcommand"
)

func main() {
	e := &env{
		stdout: os.Stdout,
		stderr: os.Stderr,
		args:   os.Args,
	}

	if err := run(e); err != nil {
		fmt.Fprintln(e.stderr, err)
		os.Exit(1)
	}
}

func run(e *env) error {
	if len(e.args) < 2 {
		return fmt.Errorf(SubcommandError)
	}

	switch e.args[1] {
	case HideCommand:
		c := hideConfig{}

		if err := parseHideArgs(&c, e.args[2:], e.stderr); err != nil {
			return err
		}

		return runHide(e, &c)
	case RevealCommand:
		c := revealConfig{}

		if err := parseRevealArgs(&c, e.args[2:], e.stderr); err != nil {
			return err
		}

		return runReveal(e, &c)
	default:
		return fmt.Errorf(SubcommandError)
	}
}

func runHide(e *env, c *hideConfig) error {
	fmt.Fprintf(e.stdout, "Image Path:\t%s\n", c.imgPath)
	fmt.Fprintf(e.stdout, "Mnemonic:\t%s\n", c.mnemonic)

	return nil
}

func runReveal(e *env, c *revealConfig) error {
	fmt.Fprintf(e.stdout, "Image Path:\t%s\n", c.imgPath)

	return nil
}
