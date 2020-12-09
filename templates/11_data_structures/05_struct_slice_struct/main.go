package main

import (
	"log"
	"os"
	"text/template"
)

var textTemplate *template.Template

// Person struct
type Person struct {
	Name       string
	Age        int
	Hobbie     string
	IsLicensed bool
}

// Season struct
type Season struct {
	Name      string
	Condition string
}

// Info struct
type Info struct {
	People  []Person
	Seasons []Season
}

func init() {
	textTemplate = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	person1 := Person{
		Name:       "John",
		Age:        28,
		Hobbie:     "Footbal",
		IsLicensed: true,
	}

	person2 := Person{
		Name:       "James",
		Age:        24,
		Hobbie:     "Dancing",
		IsLicensed: false,
	}

	person3 := Person{
		Name:       "Alex",
		Age:        48,
		Hobbie:     "COD",
		IsLicensed: true,
	}

	person4 := Person{
		Name:       "Jamie",
		Age:        44,
		Hobbie:     "Programming",
		IsLicensed: false,
	}

	season1 := Season{
		Name:      "Summer",
		Condition: "hot",
	}

	season2 := Season{
		Name:      "Winter",
		Condition: "cold",
	}

	season3 := Season{
		Name:      "Spring",
		Condition: "slight cold",
	}

	season4 := Season{
		Name:      "Rainy",
		Condition: "too watery",
	}

	impInfo := Info{
		People:  []Person{person1, person2, person3, person4},
		Seasons: []Season{season1, season2, season3, season4},
	}

	err := textTemplate.ExecuteTemplate(os.Stdout, "tpl.gohtml", impInfo)
	if err != nil {
		log.Fatal(err)
	}
}
