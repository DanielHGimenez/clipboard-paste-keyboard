# clipboard-paste-keyboard

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
3. Execute as bellow
```bash
./cbpk write
```

Notice that keyboard will start writing immediately. For that reason, is recommended to create a hotkey to execute this command.
