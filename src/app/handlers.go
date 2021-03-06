package main

import (
  "github.com/gorilla/mux"
  "net/http"
  "fmt"
  "app/controllers"
)

// func startServer(port string)  {
//   router := mux.NewRouter()
//
//   /****** Routes *****/
//   router.HandleFunc("/api/getAllBlogPosts", getAllBlogPostsHandler)
//   router.HandleFunc("/api/getBlogPostById/{id}", getBlogPostByIdHandler)
//
//   router.PathPrefix("/").Handler(http.StripPrefix("/public", http.FileServer(http.Dir("public/"))))
//   fmt.Println("Server Started.. Serving on port " + port)
//   http.ListenAndServe(port, router)
// }

func GetAllBlogPostsHandler(w http.ResponseWriter, r *http.Request) {
  result := BlogController.GetAllBlogPosts()
  fmt.Fprintln(w,result)
}

func GetBlogPostByIdHandler(w http.ResponseWriter, r *http.Request) {
  vars  := mux.Vars(r)
  id    := vars["id"]
  result := BlogController.GetBlogPostById(id)
  fmt.Fprintln(w,result)
}
