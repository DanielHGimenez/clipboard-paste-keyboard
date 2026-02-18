// SPDX-License-Identifier: MIT
package main

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
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
			Name:        "write",
			Usage:       "types clipboard content via simulated keystrokes (--delay N to wait N seconds)",
			Description: "Use --delay N to wait N seconds before typing begins",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:     "delay",
					Usage:    "number of seconds to wait before typing begins",
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

	// Normalize line endings (keeps multiline behavior consistent)
	content = normalizeNewlines(content)

	// --- Linux: choose method based on session type (not tool presence) ---
	if runtime.GOOS == "linux" {
		switch detectLinuxSessionType() {
		case "x11":
			// Only in X11 session do we try xdotool.
			if err := tryXdotoolType(content); err == nil {
				return
			}
		case "wayland":
			// Only in Wayland session do we try wtype.
			if err := tryWtype(content); err == nil {
				return
			}
		}
		// otherwise fall through to sendkeys
	}

	// Fallback: original behavior (unchanged for Windows)
	kb, err := sendkeys.NewKBWrapWithOptions(sendkeys.Noisy)
	if err != nil {
		log.Fatal(err)
	}
	kb.Type(content)
}

func normalizeNewlines(s string) string {
	s = strings.ReplaceAll(s, "\r\n", "\n")
	s = strings.ReplaceAll(s, "\r", "")
	return s
}

// detectLinuxSessionType returns "wayland", "x11", or "unknown".
func detectLinuxSessionType() string {
	// Most reliable hint:
	// - XDG_SESSION_TYPE=wayland|x11
	// Also:
	// - WAYLAND_DISPLAY set => wayland
	// - DISPLAY set (and not Wayland) => x11
	t := strings.ToLower(strings.TrimSpace(os.Getenv("XDG_SESSION_TYPE")))
	if t == "wayland" || os.Getenv("WAYLAND_DISPLAY") != "" {
		return "wayland"
	}
	if t == "x11" || os.Getenv("DISPLAY") != "" {
		return "x11"
	}
	return "unknown"
}

func tryXdotoolType(text string) error {
	xdotoolPath, err := exec.LookPath("xdotool")
	if err != nil {
		return err
	}

	// xdotool type --file - reads text from stdin (UTF-8 if locale is UTF-8)
	cmd := exec.Command(xdotoolPath, "type", "--clearmodifiers", "--file", "-")
	cmd.Stdin = bytes.NewBufferString(text)

	// Force UTF-8 locale to avoid encoding issues
	cmd.Env = append(os.Environ(), "LC_ALL=C.UTF-8", "LANG=C.UTF-8")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func tryWtype(text string) error {
	wtypePath, err := exec.LookPath("wtype")
	if err != nil {
		return err
	}

	// wtype - reads text from stdin
	cmd := exec.Command(wtypePath, "-")
	cmd.Stdin = bytes.NewBufferString(text)

	cmd.Env = append(os.Environ(), "LC_ALL=C.UTF-8", "LANG=C.UTF-8")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
