package main

// importar função fmt
import "fmt"

//declaração da variável CONST do ponto de ebulição da água em F
const ebulicaoK = 373.0

// função principal
func main() {
	tempK := ebulicaoK   // variável do valor da temperatura em graus F
	tempC := tempK - 273 //converção de F para ceucius
	//apareça na tela o resultado
	fmt.Printf("A temperatura de ebulição da água em °K = %g e a temperatura de ebulição da água em °C = %g . ", tempK, tempC)

	// A gente espera que apareça K = 373 e C = 100
}
