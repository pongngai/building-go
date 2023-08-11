// generate.go
//go:build ignore
// +build ignore

package main

import (
	"fmt"
	_const "github.com/pongngai/building-go/const"
	"log"
	"os"
	_ "os/exec"
	"text/template"
)

func main() {
	fmt.Println("Generating database configuration")
	defer fmt.Println("Done...")
	templateStr := _const.Header + _const.DatabaseTemplate

	// Create a new template and parse the template string
	tmpl := template.Must(template.New("generated").Parse(templateStr))

	// Create the output file
	outputFile, err := os.Create(_const.DatabaseGeneratedPath)
	if err != nil {
		log.Fatal("Failed to create output file:", err)
	}
	defer outputFile.Close()

	// Execute the template and write the generated code to the output file
	err = tmpl.Execute(outputFile, nil)
	if err != nil {
		log.Fatal("Template execution failed:", err)
	}

}
