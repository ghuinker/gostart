package cmd

import (
	"github.com/tkrajina/typescriptify-golang-structs/typescriptify"
)

type EnumValue string

type EnumExample struct {
	Value  EnumValue
	TSName string
}

var FilterLookups = []EnumExample{
	{"1", "One"},
	{"2", "Two"},
}

type ExampleStruct struct {
	Name  string
	Value string
}

func Tsify() {
	converter := typescriptify.New().
		Add(ExampleStruct{}).
		AddEnum(FilterLookups)

	converter.BackupDir = "tmp"
	converter.CreateInterface = true
	err := converter.ConvertToFile("assets/js/types/api.ts")
	if err != nil {
		panic(err.Error())
	}
}
