package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

var templates *template.Template
var templatesFolder string
var componentsFolder string

func main() {

	rutaActual, errr := os.Getwd()
	if errr != nil {
		fmt.Println("Error al obtener la ruta de trabajo actual:", errr)
	}
	fmt.Println("Ruta de trabajo actual:", rutaActual)

	templatesFolder = rutaActual + "/templates/"
	componentsFolder = templatesFolder + "components/"

	templates = template.Must(template.ParseFiles(
		templatesFolder+"base.gohtml",
		componentsFolder+"header.gohtml",
		componentsFolder+"footer.gohtml",
		componentsFolder+"head.gohtml",
	))

	http.HandleFunc("/", mainHandler)

	fmt.Println("Servidor escuchando en http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}

func listRout(rutaActual string) {
	fmt.Println("Rutal actual: ", rutaActual)
	// Abre el directorio para leer sus contenidos
	dir, err := os.Open(rutaActual)
	if err != nil {
		fmt.Println("Error al abrir el directorio:", err)
		return
	}
	defer dir.Close()

	// Lee los contenidos del directorio
	elementos, err := dir.ReadDir(0)
	if err != nil {
		fmt.Println("Error al leer el directorio:", err)
		return
	}

	// Imprime los nombres de los elementos
	fmt.Println("Elementos en la ruta actual:")
	for _, elemento := range elementos {
		fmt.Println(elemento.Name())
	}
}

func mainHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Println("===================== request:\n", request, "\n\n=====================")

	data := struct {
		Title string
	}{
		Title: "Página de inicio",
	}

	err := templates.ExecuteTemplate(response, "base.gohtml", data)
	if err != nil {
		http.Error(response, "Error al renderizar la plantilla", http.StatusInternalServerError)
		fmt.Println("Error al renderizar la plantilla:", err)
	}
}
