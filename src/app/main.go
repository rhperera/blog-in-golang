package main

import (
  _"fmt"
  _"app/models"
  "net/http"
)

func main() {
  // startServer(":8100")
  // fmt.Println("hello")
  // blogpost := models.CreateBlogpost(2,"Hello World","How are you doing today",890)
  // fmt.Println(blogpost.Title)
  // fmt.Println(blogpost.GetContent())
  router := CreateNewRouter(routes)
  http.ListenAndServe(":8080",router)
}

var routes = Routes{
    Route{
        "GetAllBlogPostsHandler",
        "GET",
        "/",
        GetAllBlogPostsHandler,
    },
    Route{
        "GetBlogPostByIdHandler",
        "GET",
        "/todos",
        GetBlogPostByIdHandler,
    },
}
