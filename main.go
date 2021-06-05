package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var url string = "chats.txt"
	linea := leerArchivo(url)

	lineaLimpia := limpiarArchivo(linea)
	fmt.Println(lineaLimpia)
}

// Primero leer el archivo en go
func leerArchivo(url string) string {
	content, err := ioutil.ReadFile(url)

	if err != nil {
		log.Fatal(err)
	}
	return string(content)

}

// limpiar el archivo
func limpiarArchivo(linea string) []string {
	split := strings.Split(linea, "#")
	f, err := os.Create("limpio.txt")
	check(err)
	defer f.Close()
	f.Sync()
	w := bufio.NewWriter(f)
	for _, s := range split {
		_, err := w.WriteString(s + "\n")
		check(err)
	}
	w.Flush()

	return split
}

// agregar ruido
//dividir en train y test
// Construir la red neuronal
//Correrla
//ejecutar los resultados
