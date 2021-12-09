package main

import (
    "fmt"
    "net/http"
    "encoding/json"
)

func hello(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "hello\n")
}

type FavoriteItem struct {
    ProductID string `json:"product_id"`
}

func favorites(w http.ResponseWriter, req *http.Request) {
  userID := req.Header.Get("X-Auth-Identity")
  if (userID != "") {
    data :=  map[string][]FavoriteItem {
      "123": []FavoriteItem{  FavoriteItem{ ProductID: "1"}, FavoriteItem {ProductID: "2"}},
      "234": []FavoriteItem{  FavoriteItem{ ProductID: "1"}, FavoriteItem {ProductID: "3"}},
      "345": []FavoriteItem{  FavoriteItem{ ProductID: "1"}, FavoriteItem {ProductID: "4"}},
    }
    json.NewEncoder(w).Encode(data[userID])

  } else {
    w.WriteHeader(http.StatusForbidden)
  }
}

func main() {

    http.HandleFunc("/hello", hello)
    http.HandleFunc("/favorites", favorites)

    http.ListenAndServe(":8383", nil)
}
