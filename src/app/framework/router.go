package Router

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
  router := mux.NewRouter().StrictSlash(true)
  for _, route := range routes {
      router.
        Methods(route.Method).
        Path(route.Path).
        Name(route.Name).
        HandlerFunc(route.Handler)
  }
  return router
}
