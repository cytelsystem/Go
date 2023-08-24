package crear

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"unicode"
)

var (
	pathGlosario = "E:/HECTORJM/Proyectos/ReactJs/glosariodeterminosViteNuevo/src/"
	LinkFile     string
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
	jsonFolderPath := pathGlosario + "data"

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

func readLinks() []Link {
	links := []Link{}

	for {
		var linkName string
		fmt.Print("Link (o dejar en blanco para terminar): ")
		_, err := fmt.Scanln(&linkName)
		if err != nil {
			fmt.Println("Error al leer la entrada:", err)
			break
		}

		linkName = strings.TrimSpace(linkName)
		if linkName == "" {
			break
		}

		var linkFile string
		fmt.Print("Nombre Archivo: ")
		_, err = fmt.Scanln(&linkFile)
		if err != nil {
			fmt.Println("Error al leer la entrada:", err)
			break
		}

		linkFile = strings.TrimSpace(linkFile)
		link := Link{Nombre: linkName, Archivo: linkFile}
		links = append(links, link)
	}

	return links
}

func openFileInEditor(filePath string) error {
	editorCommands := map[string]string{
		"code":    "code",
		"notepad": "notepad",
		"mspaint": "mspaint",
	}

	editorToUse := "code" // Default to Notepad

	if strings.HasSuffix(filePath, ".jpg") {
		editorToUse = "mspaint" // Use MS Paint for .jpg files
	} else if _, err := os.Stat("C:/Program Files/Microsoft VS Code/Code.exe"); err == nil {
		editorToUse = "code" // Use VS Code if it's available
	}

	editorCommand, found := editorCommands[editorToUse]
	if !found {
		return fmt.Errorf("Editor '%s' not supported", editorToUse)
	}

	cmd := exec.Command(editorCommand, filePath)
	err := cmd.Start()
	if err != nil {
		return err
	}
	return nil
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
	fmt.Print("Nombre: ")
	fmt.Scanln(&newName)
	newName = strings.TrimSpace(newName)

	var newDetalle string
	fmt.Print("Descripcion: ")
	fmt.Scanln(&newDetalle)
	newDetalle = strings.TrimSpace(newDetalle)

	links := readLinks()

	newRecord := Result{
		ID:      newID,
		Name:    newName,
		Detalle: newDetalle,
		Link:    links, // Agregar los enlaces aquí
	}

	data.Results = append(data.Results, newRecord)

	//**********************************************************************//

	// Check if the file extension is ".txt"

	for _, link := range links {
		if filepath.Ext(link.Archivo) == ".txt" || filepath.Ext(link.Archivo) == ".jpg" {
			// Generate the file path using the global variable and "Ejemplos" directory
			filePath := filepath.Join(pathGlosario, "Ejemplos", link.Archivo)

			if filepath.Ext(link.Archivo) == ".jpg" {
				// Create the file
				file, err := os.Create(filePath)
				if err != nil {
					fmt.Printf("Error creating file '%s': %s\n", link.Archivo, err)
					continue
				}
				file.Close()
			}

			if err := openFileInEditor(filePath); err != nil {
				fmt.Printf("Error opening file '%s': %s\n", link.Archivo, err)
			}
		}
	}

	//**********************************************************************//

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
