package main

import (
	"embed"
	"fmt"
	"gostart/cmd"
	"gostart/core"
	"io/fs"
	"log"
	"os"

	"github.com/joho/godotenv"
)

//go:embed static/*
var embededFiles embed.FS

func SetGlobalFileStaticFileSystem() {
	fsys, err := fs.Sub(embededFiles, "static")
	if err != nil {
		panic(err)
	}

	core.StaticFileSystem = fsys
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run script.go <command>")
		os.Exit(1)
	}

	command := os.Args[1]
	SetGlobalFileStaticFileSystem()

	switch command {
	case "runserver":
		cmd.RunServer()
	case "tsify":
		cmd.Tsify()
	default:
		fmt.Println("Unknown command:", command)
	}
}
