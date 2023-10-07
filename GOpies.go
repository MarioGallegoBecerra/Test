package GOpies

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
