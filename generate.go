package main

import (
	"log"
	"os"
	"fmt"
	"path/filepath"
	"text/template"
)

func main() {

	blocks := template.Must(template.ParseGlob("./pages/blocks/*.tmpl"))
	blocks.ParseGlob("./pages/*.html")
	fmt.Printf("%s\n", blocks.DefinedTemplates())

	matches, _ := filepath.Glob("./pages/*.html")

	for _, path := range matches {

		file := filepath.Base(path)

		fmt.Printf("%s: creating...\n", file)

		w, err := os.Create("./site/" + file)
		if err != nil {
			log.Fatalf("err: %s", err)
		}

		err = blocks.ExecuteTemplate(w, file, nil)
		if err != nil {
			log.Fatalf("err: %s", err)
		}

		w.Close()
		fmt.Printf("%s: done.\n", file)
	}
}
