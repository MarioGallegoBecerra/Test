package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func createDirectoryTree(rootPath string, levels int, maxDirsPerLevel int) {
	if levels <= 0 {
		return
	}

	for i := 1; i <= maxDirsPerLevel; i++ {
		dirName := fmt.Sprintf("dir%d", i)
		dirPath := filepath.Join(rootPath, dirName)

		// Crea el directorio
		err := os.Mkdir(dirPath, 0755)
		if err != nil {
			fmt.Println("Error al crear el directorio:", err)
			return
		}

		fmt.Println("Directorio creado:", dirPath)

		// Llama recursivamente a la función para crear subdirectorios
		createDirectoryTree(dirPath, levels-1, maxDirsPerLevel)
	}
}

func main() {
	rootDirectory := "mi_arbol_de_directorios"
	levels := 3          // Número de niveles de directorios
	maxDirsPerLevel := 2 // Máximo de directorios por nivel

	// Crea el directorio raíz
	err := os.Mkdir(rootDirectory, 0755)
	if err != nil {
		fmt.Println("Error al crear el directorio raíz:", err)
		return
	}

	fmt.Println("Directorio raíz creado:", rootDirectory)

	// Llama a la función para crear el árbol de directorios
	createDirectoryTree(rootDirectory, levels, maxDirsPerLevel)
}
