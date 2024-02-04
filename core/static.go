package core

import (
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
)

var StaticFileSystem fs.FS

func GetStaticContents(filePath string) []byte {
	// Open the file
	file, err := StaticFileSystem.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return []byte{}
	}
	defer file.Close() // Ensure the file is closed after reading

	// Read the file content
	fileContent, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return []byte{}
	}
	return fileContent
}

type FileInfo struct {
	File    string `json:"file"`
	IsEntry bool   `json:"isEntry"`
	Src     string `json:"src"`
}

func AssetManifest(name string) string {
	if IsDebug() {
		return "http://localhost:3000/static/" + name
	}

	var manifest map[string]FileInfo
	contents := GetStaticContents("assets/manifest.json")
	err := json.Unmarshal(contents, &manifest)

	if err != nil {
		fmt.Println("Error asset manifesting:", err)
		return ""
	}
	entry, ok := manifest[name]
	if !ok {
		println("Error reading from manifest: ", name)
		return ""
	}
	return "/static/assets/" + entry.File
}
