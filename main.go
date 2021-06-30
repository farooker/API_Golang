package main

import (
	routes "Admin-Tools-Api/Routers"
	"Admin-Tools-Api/Utility"
	"fmt"
	"net/http"
	"time"

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
func routine1() {
  
	for i := 0; i < 20; i++ {
		fmt.Println("routine  1")	
	}
}
func routine2() {
	for i := 0; i < 20; i++ {
		fmt.Println("routine  2")	
	}
}
func routineCH(c chan int,paload int) {
   c <- paload
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
 // Closures  out funtion
	_OutClosures := OutClosures()
	fmt.Println(_OutClosures())
    fmt.Println(_OutClosures())

	// run  แบบ  Concurrent
	go  routine1()
    go  routine2()
	time.Sleep(2*time.Second)

	// Channel and Deadlock

	ch := make(chan int,1)  //กรณี นี้หาก ไม่กำหนด size  buffer จะเกิด  Deadlock
	ch <-1
	fmt.Println(<-ch)
	ch <-2
	fmt.Println(<-ch)
	time.Sleep(1*time.Second)

	// Chanel โดย  แยก routine
	chr := make(chan int)
	go routineCH(chr,20)
	go routineCH(chr,21)
	go routineCH(chr,22)
	fmt.Println(<-chr)
	fmt.Println(<-chr)
	fmt.Println(<-chr)
	time.Sleep(1*time.Second)











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
