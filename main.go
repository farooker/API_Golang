package main

import (
	routes "Admin-Tools-Api/Routers"
	"Admin-Tools-Api/Utility"
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
	//"log"
)

func OutClosures() func() int {
    GFG := 0
    return func() int {
      GFG += 1
      return GFG
  }
}

func main() {

   // Closures  in funtion
	GFG := 0
    counter := func() int {
       GFG += 1
       return GFG
    }
  
	fmt.Println(counter())
	fmt.Println(counter())

	_OutClosures := OutClosures()
	fmt.Println(_OutClosures())
    fmt.Println(_OutClosures())















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
