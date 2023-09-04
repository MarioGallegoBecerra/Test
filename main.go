package main

import (
	"fmt"
	"net/http"
)

func main() {
	//sock.Start()
	// Configurar un manejador para la ruta "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Escribir "Hola mundo" como respuesta
		fmt.Fprintf(w, "Hola mundo")
	})

	// Iniciar el servidor en el puerto 8080
	fmt.Println("Servidor escuchando en http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}
