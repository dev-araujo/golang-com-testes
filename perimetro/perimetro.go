package main

type Retangulo struct {
	Largura float64
	Altura  float64
}

func Perimetro(retangulo Retangulo) float64 {
	return 2 * (retangulo.Largura + retangulo.Altura)

}

func Area(retangulo Retangulo) float64 {
	return retangulo.Largura * retangulo.Altura
}

// Nosso próximo requisito é escrever uma função Area para círculos.
