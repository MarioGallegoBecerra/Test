package consumer

import (
	"GoAppModule/data"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Search(QuerySearch string) {
	// Realizar la solicitud GET a la API
	response, err := http.Get("https://www.carrefour.es/search-api/query/v1/search?query=" + QuerySearch + "&rows=100&start=0&lang=es")
	if err != nil {
		fmt.Println("Error al realizar la solicitud:", err)
		return
	}

	// Leer el cuerpo de la respuesta
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error al leer la respuesta:", err)
		return
	}

	// Cerrar el cuerpo de la respuesta
	defer response.Body.Close()

	var entities []data.ItemEntity
	err2 := json.Unmarshal(body, &entities)
	if err2 != nil {
		fmt.Println("Error al leer la respuesta:", err)
		return
	}
	// Imprimir la respuesta
	fmt.Println(entities)

}
