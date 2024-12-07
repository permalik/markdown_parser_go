#!/bin/zsh

go run main.go -i test.md -o output.md -f md && cat output.md
