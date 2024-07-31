package main

import (
	"log"
	"os"

	"git.tcp.direct/kayos/sendkeys"
	"github.com/atotto/clipboard"
	"github.com/urfave/cli"
)

func main() {
	cmd := cli.NewApp()
	cmd.Name = "CBPK"
	cmd.Usage = "Do the keyboard writes content of clipboard"

	cmd.Commands = []cli.Command{
		{
			Name:   "write",
			Action: writeClipboard,
		},
	}

	if err := cmd.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func writeClipboard(c *cli.Context) {
	content, err := clipboard.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	kb, err := sendkeys.NewKBWrapWithOptions(sendkeys.Noisy)
	if err != nil {
		log.Fatal(err)
	}

	kb.Type(content)
}
