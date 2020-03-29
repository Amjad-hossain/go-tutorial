package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "encoding/json"
    "log"
    "os"
)
type Response struct {
    Status string `json:"status"`
    Data []Employe `json:"data"`
}
type Employe struct {
    Id string `json:"id"`
    Name string `json:"employee_name"`
    Age string `json:"employee_age"`
    Salary string `json:"employee_salary"`
}

func main() {
    client := &http.Client {}
    req, err := http.NewRequest("GET", "http://dummy.restapiexample.com/api/v1/employees", nil)
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
    var ResponseObj Response
    json.Unmarshal(respBody, &ResponseObj)
    for i:=0; i<len(ResponseObj.Data); i++ {
       DataObj, err := json.Marshal(ResponseObj.Data[i])
       if err != nil {
         log.Fatal(err)
       }
       fmt.Println(string(DataObj))
    }
}