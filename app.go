package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {

	clock:= time.Now()
	fmt.Fprintf(w,"<body style=background-color:#9933ff;>",)
	fmt.Fprintf(w,"<h1><i>Date :: %s </i></h1>\n",clock.Format("Mon 2-Jan-2006"))
	fmt.Fprintf(w,"<h1><i>Time :: %s </i></h1>\n",clock.Format("03:04:05 PM"))
	fmt.Fprintf(w,"</body>")
	
}


func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil)) 
}
