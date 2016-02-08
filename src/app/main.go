package main

import (
  "net/http"
  "app/framework"
)

func main() {
  // startServer(":8100")
  // fmt.Println("hello")
  // blogpost := models.CreateBlogpost(2,"Hello World","How are you doing today",890)
  // fmt.Println(blogpost.Title)
  // fmt.Println(blogpost.GetContent())
  router := Router.CreateNewRouter(routes)
  http.ListenAndServe(":8080",router)
}

/* All the routes with methods and their handlers
  NAME, METHOD, PATH, HANDLER FUNCTION
*/
var routes = Router.Routes{
    Router.Route{
        "GetAllBlogPostsHandler",
        "GET",
        "/",
        GetAllBlogPostsHandler,
    },
    Router.Route{
        "GetBlogPostByIdHandler",
        "GET",
        "/todos/{id}",
        GetBlogPostByIdHandler,
    },
}
