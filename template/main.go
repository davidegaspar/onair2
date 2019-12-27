package main

import (
	"log"
	"os"
	"fmt"
	"path/filepath"
	"text/template"
)

func main() {

	blocks := template.Must(template.ParseGlob("./blocks/*.tmpl"))
	blocks.ParseGlob("*.html")
	fmt.Printf("%s\n", blocks.DefinedTemplates())

	matches, _ := filepath.Glob("*.html")

	for _, file := range matches {
		fmt.Printf("creating: %s ...", file)
		w, _ := os.Create("../site/" + file)
		err := blocks.ExecuteTemplate(w, file, nil)
		if err != nil {
			log.Fatalf("err: %s", err)
		}
		w.Close()
		fmt.Printf(" done\n")
	}
}
