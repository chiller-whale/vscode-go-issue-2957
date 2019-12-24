# Issue Repo

Provide example for this [issue](https://github.com/microsoft/vscode-go/issues/2957)

# Description

This happened after updating from the vscode go extension 0.11.7 to 0.11.8. Still persists on 0.11.9

**Note** This behaves as expected when interface is in same file.

# Steps to reproduce

1. Open `log/log.go`
2. Right click `UselessLogger` and select `Peek Implementations`

**Expected** Implementation in main.go is shown.

**Actual** No implementations found.

![example of behavior](example.gif)
