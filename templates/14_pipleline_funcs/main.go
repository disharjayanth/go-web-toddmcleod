package main

import (
	"log"
	"math"
	"os"
	"text/template"
)

var textTemplate *template.Template

func init() {
	textTemplate = template.Must(template.New("").Funcs(funcMap).ParseFiles("tpl.gohtml"))
}

func double(x int) int {
	return x + x
}

func square(x int) float64 {
	return math.Pow(float64(x), 2)
}

func sqRoot(x float64) float64 {
	return math.Sqrt(x)
}

var funcMap = template.FuncMap{
	"fdbl":  double,
	"fsq":   square,
	"fsqrt": sqRoot,
}

func main() {
	err := textTemplate.ExecuteTemplate(os.Stdout, "tpl.gohtml", 2)
	if err != nil {
		log.Fatal(err)
	}
}
