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
	limpiarArchivo(linea)
	lematizar()
	dividir()
	wordbag := wordbag()
	fmt.Println(wordbag)
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
func limpiarArchivo(linea string) {
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
}

//create a bag of words
func lematizar() {
	content, err := ioutil.ReadFile("limpio.txt")

	if err != nil {
		log.Fatal(err)
	}
	split := strings.Split(string(content), "(")
	f, err := os.Create("limpio2.txt")
	check(err)
	defer f.Close()
	f.Sync()
	w := bufio.NewWriter(f)
	for _, s := range split {
		_, err := w.WriteString(s + "\n")
		check(err)
	}
	w.Flush()

}
func dividir() {

	f, err := os.Open("limpio2.txt")
	f2, err2 := os.Create("limpio3.txt")
	check(err)
	check(err2)
	w := bufio.NewWriter(f2)
	defer f.Close()
	defer f2.Close()
	scanner := bufio.NewScanner(f)
	read := false
	for scanner.Scan() {
		if read {

			_, err := w.WriteString(scanner.Text() + "\n")
			check(err)
			read = false
		} else {
			read = true
		}
	}
	w.Flush()

}
func wordbag() []string {
	f, err := os.Open("limpio3.txt")
	var wordbag []string

	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		wordbag = append(wordbag, scanner.Text())

	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return wordbag

}

// agregar ruido
//dividir en train y test
// Construir la red neuronal
//Correrla
//ejecutar los resultados
