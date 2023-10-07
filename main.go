package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func listDirectoryTree(rootPath string, indent string) {
	files, err := ioutil.ReadDir(rootPath)
	if err != nil {
		fmt.Println("Error al leer el directorio:", err)
		return
	}

	for _, file := range files {
		fmt.Println(indent + file.Name())

		if file.IsDir() {
			// Si es un directorio, llama recursivamente a la función
			subdirPath := filepath.Join(rootPath, file.Name())
			listDirectoryTree(subdirPath, indent+"  ")
		}
	}
}

func main() {
	rootDirectory := "/" // Cambia esta ruta según la ubicación que desees explorar

	fmt.Println("Árbol de directorios y archivos en:", rootDirectory)
	listDirectoryTree(rootDirectory, "")
}
