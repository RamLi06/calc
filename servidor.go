package main

import (
  "net/http"
  "log"
  "fmt"
  "strconv"
  "strings"
)

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
  if r.Method == "GET" {
    var fodase Calculator

    op := r.URL.RawQuery[3:]
    operators := "+-/*^&"

    opIndex := strings.IndexAny(op, operators)

    fodase.Value1, _ = strconv.ParseFloat(op[:opIndex], 64)
    fodase.Value2, _ = strconv.ParseFloat(op[opIndex+1:], 64)

    fmt.Println(fodase.Value1, fodase.Value2)

    res, err := operaçao[string(op[opIndex])].calculo()
    if err != nil {
      log.Fatal("matematica 2 kkkkkkkkkkkk")
    }

    fmt.Printf("%f\n", res)


    w.Write([]byte(fmt.Sprintf("Sucesso!! %v", res)))
  }
}


func servidor() {
  http.HandleFunc("/result", ServeHTTP)
  log.Fatal(http.ListenAndServe(":8080", nil))
}

