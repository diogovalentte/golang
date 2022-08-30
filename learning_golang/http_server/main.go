package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/JohnCGriffin/overflow"
)

var tz, _ = time.LoadLocation("America/Sao_Paulo")

func main() {

	var port int

	flag.IntVar(&port, "port", 3333, `Default is "3333"`)
	flag.Parse()

	server := fmt.Sprintf(":%v", port)

	mux := http.NewServeMux()
	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/check_status", getCheckStatus)
	mux.HandleFunc("/calculate", getCalculate)

	err := http.ListenAndServe(server, mux)
	if err != nil {
		fmt.Printf("Error while starting server: %v", err)
		os.Exit(1)
	}
}

// JSON response struct
type JSONCalculateResponse struct {
	Num1     int
	Num2     int
	Operator string
	Result   int
	Error    string
}

// get functions will be our Handler functions that handle HTTP requests depending on the request
func getCalculate(w http.ResponseWriter, r *http.Request) {
	// Get two number and a operator from request, do a math operation and return the math result
	// Log in terminal
	t := time.Now()
	tc := t.In(tz).Format("2006-01-02 15:04:05")
	fmt.Printf("[%v]: got /calculate request\n", tc)

	// Read body request into []byte to get useful informations
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Could not read body: %s\n", err)
	}

	var responseErr string

	rNum1 := r.URL.Query().Get("num1")
	rNum2 := r.URL.Query().Get("num2")
	num1, err := strconv.Atoi(rNum1)
	if err != nil {
		responseErr = err.Error()
	}
	num2, err := strconv.Atoi(rNum2)
	if err != nil {
		responseErr = err.Error()
	}

	operator := string(body[0])

	// Do math depending on the arithmetic operator
	var result int

	if responseErr == "" {
		switch operator {
		case "+":
			r, ok := overflow.Add(num1, num2) // Check if num1 + num2 dont overflow
			if !ok {
				responseErr = "the result of Number 1 + Number 2 overflow\n"
			} else {
				result = r
			}
		case "-":
			r, ok := overflow.Sub(num1, num2) // Check if num1 - num2 dont overflow
			if !ok {
				responseErr = "the result of Number 1 - Number 2 overflow\n"
			} else {
				result = r
			}
		case "*":
			r, ok := overflow.Mul(num1, num2) // Check if num1 * num2 dont overflow
			if !ok {
				responseErr = "the result of Number 1 * Number 2 overflow\n"
			} else {
				result = r
			}
		case "/":
			r, ok := overflow.Div(num1, num2) // Check if num1 / num2 dont overflow
			if !ok {
				responseErr = "the result of Number 1 / Number 2 overflow\n"
			} else {
				result = r
			}
		}
	}

	// Write response in a JSON object to send back to the request
	jsonResponse := JSONCalculateResponse{
		Num1:     num1,
		Num2:     num2,
		Operator: operator,
		Result:   result,
		Error:    responseErr,
	}
	jsonBytes, err := json.Marshal(jsonResponse)
	if err != nil {
		fmt.Printf("Error while writing json object to bytes: %v", err)
	}

	// Send response back to requester
	w.Write(jsonBytes) // Write the JSON bytes to the response body
}
func getRoot(w http.ResponseWriter, r *http.Request) {
	// Log in terminal a request and send a message back to requester
	t := time.Now()
	tc := t.In(tz).Format("2006-01-02 15:04:05")
	fmt.Printf("[%v]: got / request\n", tc)

	io.WriteString(w, "This is my website!\n")
}
func getCheckStatus(w http.ResponseWriter, r *http.Request) {
	// Log in termina a request and returns the status code 200 to the requester
	t := time.Now()
	tc := t.In(tz).Format("2006-01-02 15:04:05")
	fmt.Printf("[%v]: got / request\n", tc)

	w.WriteHeader(200)
}
