package main

import (
    "io/ioutil"
    "log"
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/begin/", handler)
    http.ListenAndServe(":9000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {


    resp, err := http.Get("http://localhost:8080/hello/?key=syn")
    if err != nil {
        // handle err
    }
    defer resp.Body.Close()

    if resp.StatusCode == http.StatusOK {
        bodyBytes, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            log.Fatal(err)
        }
        bodyString := string(bodyBytes)
        fmt.Printf(bodyString)
	
	if bodyString != "syn-ack" {
		fmt.Printf("Wrong response written")
		return 
	}

        resp2, err2 := http.Get("http://localhost:8080/hello/?key=ack")
        if err2 != nil {
            // handle err
        }
        defer resp2.Body.Close()
    
        if resp2.StatusCode == http.StatusOK {
            _, err2 := ioutil.ReadAll(resp2.Body)
            if err2 != nil {
                log.Fatal(err2)
            }
            fmt.Printf("Connection established")
        }
    }
}
