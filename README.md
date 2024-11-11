# clipboard-paste-keyboard

[![GitHub Downloads (specific asset, latest release)](https://img.shields.io/github/downloads/DanielHGimenez/clipboard-paste-keyboard/latest/bin.zip?style=for-the-badge&label=download%20latest&labelColor=green&color=grey)](https://github.com/DanielHGimenez/clipboard-paste-keyboard/releases/latest/download/bin.zip)


## Summary

`clipboard-paste-keyboard` is a cross-platform command-line tool to write clipboard content using the keyboard. It can be used to write the clipboard content through RDS, or any software that disable pasting (Ctrl + V).

## Platforms

- ~~OSX~~ (Bugs found during build)
- Windows
- Linux, Unix

## Requirements

Linux and Unix platforms requires 'xclip' or 'xsel' command to be installed.

## How to use

1. Download the binary file
2. Rename the file for `cbpk`

Execute as below to start writing clipboard content.
```bash
./cbpk write
```

Notice that keyboard will start writing immediately (at Linux/Unix it will execute after 2 seconds).
If you need time to reach the application that you want to write on, use the option `--delay value` with an integer number that will be waited in seconds before the writing begins. See the example below.
```bash
./cbpk write --delay 5
```
The command above will start the writing after 5 seconds (7 seconds if you are using Linux/Unix).
