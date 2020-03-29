package main
import (
	"fmt"
	"bytes"
	"net/http"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	client := &http.Client{}
	var jsonStr = []byte(`{"name":"test","salary":"123","age":"23"}`)
	req, err := http.NewRequest("POST", "http://dummy.restapiexample.com/api/v1/create", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("API-Key", os.Getenv("API_KEY"))
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println(err.Error())
    }
    defer resp.Body.Close()
    respBody, err := ioutil.ReadAll(resp.Body)
    if err != nil {
       log.Fatal(err)
    }
    fmt.Println(string(respBody))
}