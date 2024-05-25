package main

import (
	"fmt"
	"os"

	"github.com/pmuens/imagephrase/imgp"
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
	imgPath, err := imgp.HideInImage(c.imgPath, c.mnemonic)
	if err != nil {
		return err
	}

	fmt.Fprintf(e.stdout, "Mnemonic successfully hidden in %s\n", imgPath)

	return nil
}

func runReveal(e *env, c *revealConfig) error {
	mnemonic, err := imgp.RevealFromImage(c.imgPath)
	if err != nil {
		return err
	}

	fmt.Fprintf(e.stdout, "Mnemonic is: %s\n", mnemonic)

	return nil
}
