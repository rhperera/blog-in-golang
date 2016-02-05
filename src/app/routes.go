package main

import (
  "net/http"
  "github.com/gorilla/mux"
)

type Route struct{
  Name string
  Method string
  Pattern string
  HandlerFunc http.HandlerFunc 
}
