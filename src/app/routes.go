package main

import (
  "net/http"
  "github.com/gorilla/mux"
)

type Route struct{
  Name string
  Method string
  Path string
  Handler http.HandlerFunc
}

type Routes []Route

func CreateNewRouter(routes Routes) *mux.Router {

  router := mux.NewRouter().strictSlash(true)
  for _, route := range routes {
      router
      .Methods(route.Method)
      .Path(route.Path)
      .Name(route.Name)
      .Handler(router.Handler)
  }

  return router

}
