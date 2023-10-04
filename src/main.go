package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func writeToFile(filename string, content string) {
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		fmt.Println("Unable to write " + filename + "Error: " + err.Error())
	}

}

func logPid() {
	var pid int = os.Getpid()
	fmt.Println("Program PID: ", pid, "")
	writeToFile("pid", strconv.Itoa(pid))
}

func getData(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Getting data\n")

}

func saveData(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Saving data\n")
}

func setupHttpServer(portNumber string) {
	fmt.Println("Starting server on port: ", portNumber)
	http.HandleFunc("/get", getData)
	http.HandleFunc("/save", saveData)
	http.ListenAndServe(":"+portNumber, nil)
}

func main() {
	// TODO get port number from command line
	// if no port number is defined set default port number
	var defaultPortNumber string = "8000"
	fmt.Println("Starting program")
	logPid()
	setupHttpServer(defaultPortNumber)
}
