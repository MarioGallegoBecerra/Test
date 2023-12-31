package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"os/exec"
	"strings"
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

	http.HandleFunc("/", mainHandler)

	fmt.Println("Servidor escuchando en http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}

func listRout(rutaActual string) string {

	toReturn := ""

	fmt.Println("Rutal actual: ", rutaActual)
	// Abre el directorio para leer sus contenidos
	dir, err := os.Open(rutaActual)
	if err != nil {
		fmt.Println("Error al abrir el directorio:", err)
		toReturn = err.Error()
		return toReturn
	}
	defer dir.Close()

	// Lee los contenidos del directorio
	elementos, err := dir.ReadDir(0)
	if err != nil {
		fmt.Println("Error al leer el directorio:", err)
		toReturn = err.Error()
		return toReturn
	}

	// Imprime los nombres de los elementos
	fmt.Println("Elementos en la ruta actual:")
	for _, elemento := range elementos {
		fmt.Println(elemento.Name())
		toReturn += " | " + elemento.Name()
	}

	return toReturn
}

func mainHandler(response http.ResponseWriter, request *http.Request) {

	if request.RequestURI == "/favicon.ico" {
		return
	}

	fmt.Println("===================== request:\n", request, "\n\n=====================")

	fmt.Println(request.RequestURI)
	pathParms := strings.Split(strings.Split(request.RequestURI, "?")[1], "&")
	option := pathParms[0]
	fmt.Println("choosed option: ", option)

	switch option {
	case "dir":
		ls := listRout(pathParms[1])
		fmt.Fprintln(response, ls)
		break
	case "getBase":
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
		break
	case "init":
		initTemplates()
		break
	case "cmd":
		comand := strings.Replace(pathParms[1], "%20", " ", -1)
		fmt.Println("parametro 1: ", comand)
		cmd := exec.Command("cmd", "/", comand)
		output, _ := cmd.CombinedOutput()
		outputS := string(output)
		fmt.Println("comando result\n", outputS)
		fmt.Fprintln(response, outputS)
		break
	}
}

func initTemplates() {

	templates = template.Must(template.ParseFiles(
		templatesFolder+"base.gohtml",
		componentsFolder+"header.gohtml",
		componentsFolder+"footer.gohtml",
		componentsFolder+"head.gohtml",
	))
}
