package main

import (
	routes "Admin-Tools-Api/routes"
	"Admin-Tools-Api/utility"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	//"log"
)

func homePage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

type Cat interface {
	Meow1()
	Meow2()
	Meow3() string
}

type Lion struct{}

func (l Lion) Meow1() {
	fmt.Println("interface2")
}
func (l Lion) Meow2() {
	fmt.Println("interfacec2")
}
func (l Lion) Meow3() string {
	//fmt.Println("Roar3")

	return "initerface retuern"
}

//type CatFactory func() Cat

func CreateLion() Cat {
	return Lion{}
}

func hello(name string, result chan<- string) {
	output := "Hello " + name
	fmt.Printf("In function = %s\n", output)
	result <- output
}

func main() {

	lions := CreateLion()
	lions.Meow1()
	lions.Meow2()

	fmt.Println(lions.Meow3())

	fmt.Println("Start main")
	name := "Fake name"

	result := make(chan string)
	go hello(name, result)

	fmt.Println("Finish main")

	fmt.Println(<-result)

	// var cf CatFactory = CreateLion
	// fLion := cf()
	// fLion.Meow()
	utility.InitializeConfig()

	router := httprouter.New()
	routes.Customize(router)
	routes.WS(router)
	routes.FileAccess(router)

	fmt.Printf("\n")
	fmt.Printf("------------------------------------")
	fmt.Printf("\nlocal IP: %s", utility.Configuration.Ip)
	fmt.Printf("\nlocal Port: %s\n", utility.Configuration.Port)
	fmt.Printf("------------------------------------\n")

	router.ServeFiles("/static/*filepath", http.Dir("./images/"))
	http.ListenAndServe(":8080", router)
}
