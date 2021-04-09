package main

import (
	"log"
	"os"

	"github.com/jailtonjunior94/aopa-cli/src/cryptography"
	"github.com/jailtonjunior94/aopa-cli/src/infrastructure"

	"github.com/urfave/cli/v2"
)

func main() {
	infrastructure.NewEnvironments()

	app := cli.NewApp()
	app.Authors = []*cli.Author{
		{Name: "Jailton Junior", Email: "jailton.junior94@outlook.com"},
	}
	app.Name = "aopa-cli"
	app.Usage = "Aopa CLI para auxiliar suporte"
	app.Version = "1.0.0"

	flags := configureFlags()
	app.Commands = configureCommands(flags)

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func configureFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Required: true,
			Name:     "password",
			Aliases:  []string{"p"},
			Usage:    "Senha para criptografar/descriptografar valores",
		},
		&cli.StringFlag{
			Required: true,
			Name:     "text",
			Aliases:  []string{"t"},
			Usage:    "Texto para criptografar/descriptografar",
		},
	}
}

func configureCommands(flags []cli.Flag) []*cli.Command {
	return []*cli.Command{
		{
			Name:    "encrypt",
			Aliases: []string{"e"},
			Usage:   "Criptografia de dados",
			Flags:   flags,
			Action:  cryptography.Encrypt,
		},
		{
			Name:    "decrypt",
			Aliases: []string{"d"},
			Usage:   "Descriptografia de dados",
			Flags:   flags,
			Action:  cryptography.Decrypt,
		},
	}
}
