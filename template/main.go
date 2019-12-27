// package main
//
//
// import (
//     "os"
//     "fmt"
//     "html/template"
// )
//
//
//
// // type Todo struct {
// //     Title string
// //     Done  bool
// // }
//
//
//
// // type TodoPageData struct {
// //     PageTitle string
// //     Todos     []Todo
// // }
//
//
//
// func main() {
//
//   t := template.Must(template.ParseGlob("*.tmpl"))
//   fmt.Printf("%s", t)
//   t.Execute(os.Stdout, nil)
//
// }

package main

import (
	// "io"
	// "io/ioutil"
	"log"
	"os"
	// "fmt"
	// "path/filepath"
	"text/template"
)


func main() {
	// Here we create a temporary directory and populate it with our sample
	// template definition files; usually the template files would already
	// exist in some location known to the program.
	// dir := createTestDir([]templateFile{
	// 	// T0.tmpl is a plain template file that just invokes T1.
	// 	{"T0.tmpl", "T0 ({{.}} version) invokes T1: ({{template `T1`}})\n"},
	// 	// T1.tmpl defines a template, T1 that invokes T2. Note T2 is not defined
	// 	{"T1.tmpl", `{{define "T1"}}T1 invokes T2: ({{template "T2"}}){{end}}`},
	// })
  // fmt.Print(dir)
	// Clean up after the test; another quirk of running as an example.
	// defer os.RemoveAll(dir)

	// pattern is the glob pattern used to find all the template files.
	// pattern := filepath.Join(dir, "*.tmpl")
  // fmt.Print(pattern)

	// Here starts the example proper.
	// Load the drivers.
	drivers := template.Must(template.ParseGlob("./drivers/*.tmpl"))

	drivers.ParseGlob("*.html")

	// fmt.Print(drivers)

	err := drivers.ExecuteTemplate(os.Stdout, "index.html", nil)
	if err != nil {
		log.Fatalf("second execution: %s", err)
	}

	// We must define an implementation of the T2 template. First we clone
	// the drivers, then add a definition of T2 to the template name space.

	// 1. Clone the helper set to create a new name space from which to run them.
	// first, err := drivers.Clone()
	// if err != nil {
	// 	log.Fatal("cloning helpers: ", err)
	// }
	// // 2. Define T2, version A, and parse it.
	// _, err = first.Parse("{{define `T2`}}T2, version A{{end}}")
	// if err != nil {
	// 	log.Fatal("parsing T2: ", err)
	// }
	//
	// // Now repeat the whole thing, using a different version of T2.
	// // 1. Clone the drivers.
	// second, err := drivers.Clone()
	// if err != nil {
	// 	log.Fatal("cloning drivers: ", err)
	// }
	// // 2. Define T2, version B, and parse it.
	// _, err = second.Parse("{{define `T2`}}T2, version B{{end}}")
	// if err != nil {
	// 	log.Fatal("parsing T2: ", err)
	// }
	//
	// // Execute the templates in the reverse order to verify the
	// // first is unaffected by the second.
	// err = second.ExecuteTemplate(os.Stdout, "T0.tmpl", "second")
	// if err != nil {
	// 	log.Fatalf("second execution: %s", err)
	// }
	// err = first.ExecuteTemplate(os.Stdout, "T0.tmpl", "first")
	// if err != nil {
	// 	log.Fatalf("first: execution: %s", err)
	// }

}
