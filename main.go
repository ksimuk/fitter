package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/tormoder/fit"
)

func Zwiftify(name string) {

	data, err := os.ReadFile(name)
	if err != nil {
		fmt.Println(err)
		return
	}

	fitFile, err := fit.Decode(bytes.NewReader(data))
	if err != nil {
		fmt.Println(err)
		return
	}

	// Inspect the TimeCreated field in the FileId message
	fmt.Println("TimeCreated: ", fitFile.FileId.TimeCreated)

	// Inspect the dynamic Product field in the FileId message
	fmt.Println("GetProduct", fitFile.FileId.GetProduct())
	fmt.Println("Manufacturer", fitFile.FileId.Manufacturer)

	// Inspect the FIT file type
	fmt.Println(fitFile.Type())
	fmt.Printf("%+v\n", fitFile)

	// replace manufacturer
	fitFile.FileId.Manufacturer = fit.ManufacturerZwift
	newFileName := fileNameWithoutExtTrimSuffix(name) + "-fixed.fit"
	f, err := os.Create(newFileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	err = fit.Encode(f, fitFile, binary.LittleEndian)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Done, saved %s\n", newFileName)
}

func fileNameWithoutExtTrimSuffix(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: main <fit-file>")
		return
	}

	Zwiftify(os.Args[1])
}
