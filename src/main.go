package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

type Link struct {
	Nombre  string `json:"nombre"`
	Archivo string `json:"archivo"`
}

type Result struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Detalle string `json:"detalle"`
	Link    []Link `json:"link"`
}

type Data struct {
	Results []Result `json:"results"`
}

func getNextID(data Data) int {
	maxID := 0
	for _, item := range data.Results {
		if item.ID > maxID {
			maxID = item.ID
		}
	}
	return maxID + 1
}

func saveDataToFile(data Data, filePath string) error {
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}
	err = os.WriteFile(filePath, jsonData, 0644)
	return err
}

func selectJSONFile() (string, error) {
	jsonFolderPath := "./data"

	files, err := os.ReadDir(jsonFolderPath)
	if err != nil {
		return "", err
	}

	fmt.Println("Archivos JSON disponibles:")
	for idx, file := range files {
		if file.IsDir() || filepath.Ext(file.Name()) != ".json" {
			continue
		}
		fmt.Printf("%d. %s\n", idx+1, file.Name())
	}

	var selectedIdx int
	fmt.Print("Selecciona el índice del archivo JSON: ")
	_, err = fmt.Scan(&selectedIdx)
	if err != nil || selectedIdx < 1 || selectedIdx > len(files) {
		return "", fmt.Errorf("Selección inválida.")
	}

	selectedJSONPath := filepath.Join(jsonFolderPath, files[selectedIdx-1].Name())
	return selectedJSONPath, nil
}

func contieneCaracterONumero(s string) bool {
	for _, char := range s {
		if unicode.IsLetter(char) || unicode.IsNumber(char) {
			return true
		}
	}
	return false
}

func main() {
	selectedJSONPath, err := selectJSONFile()
	if err != nil {
		fmt.Println(err)
		return
	}

	jsonFile, err := os.Open(selectedJSONPath)
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer jsonFile.Close()

	var data Data
	decoder := json.NewDecoder(jsonFile)
	err = decoder.Decode(&data)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	newID := getNextID(data)

	var prueba string
	fmt.Scanln(&prueba)

	var newName string
	fmt.Print("Ingrese el nombre del nuevo elemento: ")
	fmt.Scanln(&newName)
	newName = strings.TrimSpace(newName)

	var newDetalle string
	fmt.Print("Ingrese la descripción del nuevo elemento: ")
	fmt.Scanln(&newDetalle)
	newDetalle = strings.TrimSpace(newDetalle)

	links := []Link{}
	for {
		var linkName string
		fmt.Print("Ingrese el nombre del enlace (o dejar en blanco para terminar): ")
		fmt.Scanln(&linkName)
		linkName = strings.TrimSpace(linkName)

		if !contieneCaracterONumero(linkName) {
			break
		}

		var linkFile string
		fmt.Print("Ingrese el nombre del archivo del enlace: ")
		_, err := fmt.Scanln(&linkFile)
		if err != nil {
			fmt.Println("Error al leer la entrada:", err)
			return
		}

		linkFile = strings.TrimSpace(linkFile)
		link := Link{Nombre: linkName, Archivo: linkFile}
		links = append(links, link)
	}

	newRecord := Result{
		ID:      newID,
		Name:    newName,
		Detalle: newDetalle,
		Link:    links,
	}

	data.Results = append(data.Results, newRecord)

	var saveConfirmation string
	fmt.Print("Presiona Enter para guardar los cambios (o escribe 'cancelar' y presiona Enter para cancelar): ")
	fmt.Scanln(&saveConfirmation)
	saveConfirmation = strings.ToLower(strings.TrimSpace(saveConfirmation))
	if saveConfirmation == "" || saveConfirmation == "guardar" {
		err = saveDataToFile(data, selectedJSONPath)
		if err != nil {
			fmt.Println("Error saving JSON file:", err)
			return
		}
		fmt.Println("Cambios guardados exitosamente.")
	} else {
		fmt.Println("Cambios no guardados.")
	}
}
