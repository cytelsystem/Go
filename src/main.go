package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func clearScreen() {
	cmd := exec.Command("clear")
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	words := []string{
		"Hungry - Hambriento/Hambrienta",
		"Thirsty - Sediento/Sedienta",
		"Excited - Emocionado/Emocionada",
		"Scared - Asustado/Asustada",
		"Funny - Divertido/Divertida",
		"Who - Quién",
		"Yesterday - Ayer",
		"Sad - Triste",
		"Bad - Malo",
		"Ugly - Feo/Fea",
		"Want - Querer",
		"Think - Pensar",
		"Feel - Sentir",
		"Ask - Preguntar",
		"Eat - Comer",
		"Watch - Mirar/Ver",
		"Buy - Comprar",
		"Sell - Vender",
		"Meet - Conocer (a alguien)",
		"Visit - Visitar",
		"Hate - Odiar",
		"Dislike - No gustar",
		"Sad - Triste",
		"Angry - Enojado/Enfadado",
		"Surprised - Sorprendido/Sorprendida",
		"Bored - Aburrido/Aburrida",
		"Busy - Ocupado/Ocupada",
		"Difficult - Difícil",
		"Fun - Divertido/Divertida",
		"Funny - Gracioso/Graciosa",
		"Serious - Serio/Seria",
		"Ugly - Feo/Fea",
		"Old - Viejo/Vieja",
		"High - Alto/Alta",
		"Low - Bajo/Baja",
		"Short - Corto/Corta",
		"Fast - Rápido/Rápida",
		"Bad - Malo/Mala",
		"Wrong - Equivocado/Equivocada",
		"Here - Aquí",
		"There - Allí",
		"Later - Más tarde",
		"Yesterday - Ayer",
		"Expensive - Caro/Cara",
		"Cheap - Barato/Barata",
		"Excited - Emocionado/Emocionada",
		"Scared - Asustado/Asustada",
		"Wait - Esperar",
		"May - Poder (posibilidad)",
		"Must - Deber (necesidad)",
		"Should - Debería (sugerencia)",
		"Bad - Malo/Mala",
		"Better - Mejor",
		"Worse - Peor",
		"Great - Genial",
		"Excellent - Excelente",
		"Wonderful - Maravilloso/Maravillosa",
		"Amazing - Asombroso/Asombrosa",
		"Fantastic - Fantástico/Fantástica",
		"Awesome - Impresionante",
		"Delicious - Delicioso/Deliciosa",
		"Handsome - Guapo/Guapa (para hombres)",
		"Pretty - Bonito/Bonita (para mujeres)",
		"Boring - Aburrido/Aburrida",
		"Difficult - Difícil",
		"Hard - Duro/Dura",
		"Simple - Simple",
		"Confused - Confundido/Confundida",
		"Thirsty - Sediento/Sedienta",
		"Rain - Lluvia",
		"Sun - Sol",
		"Wind - Viento",
		"Snow - Nieve",
		"Sky - Cielo",
		"First - Primero/Primera",
		"Last - Último/Última",
		"Previous - Anterior",
		"Young - Joven",
		"Far - Lejos",
		"Near - Cerca",
		"Above - Encima",
		"Below - Debajo",
		"Inside - Dentro",
		"Outside - Fuera",
		"Sad - Triste",
		"Calm - Tranquilo/Tranquila",
	}

	fmt.Println("Escribe la palabra que ves en pantalla. Pulsa Enter después de cada palabra.")
	fmt.Println("Si la escribes correctamente, escucharás un sonido.")
	fmt.Println("Presiona 'q' y Enter para salir.\n")

	reader := bufio.NewReader(os.Stdin)

	for _, word := range words {
		clearScreen()
		fmt.Println("Palabra:", word)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "q" {
			break
		}

		if strings.ToLower(input) == strings.ToLower(word) {
			fmt.Print("\a") // Emite un sonido en la mayoría de los sistemas
		}
	}

	fmt.Println("\n¡Hasta luego!")
}
