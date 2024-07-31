package main

import (
	"log"
	"os"
	"time"

	"git.tcp.direct/kayos/sendkeys"
	"github.com/atotto/clipboard"
	"github.com/urfave/cli"
)

func main() {
	cmd := cli.NewApp()
	cmd.Name = "CBPK"
	cmd.Usage = "Does the keyboard write content of the clipboard"

	cmd.Commands = []cli.Command{
		{
			Name: "write",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:     "delay",
					Required: false,
				},
			},
			Action: writeClipboard,
		},
	}

	if err := cmd.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func writeClipboard(c *cli.Context) {
	delay := c.Int("delay")
	if delay > 0 {
		time.Sleep(time.Duration(delay) * time.Second)
	}

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
