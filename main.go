package main

import (
	"log"
	"net/http"
	"os/exec"
)

func main() {
	http.HandleFunc("/", serveHTML)
	http.HandleFunc("/runScript", runScript)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server error:", err)
	}
}

func serveHTML(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func runScript(w http.ResponseWriter, r *http.Request) {
	out, err := exec.Command("/home/krish/go/src/github.com/comapanyzero/sntrup4591761/sntrup4591761", "./test6").Output()
	if err != nil {
    		log.Println("Script execution error:", err) // Log the specific error
    		http.Error(w, "Error running script", http.StatusInternalServerError)
    		return
}

	w.Write(out)
}

