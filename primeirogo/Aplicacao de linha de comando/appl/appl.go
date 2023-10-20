package appl

import "github.com/urfave/cli"

// Gerar is show
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "App de linha de comando"
	app.Usage = "Buscar ips e nomes na internet"

	return app
}
