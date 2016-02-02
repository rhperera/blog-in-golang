package main

import (
  "fmt"
  "app/models"
)

func main() {
  startServer(":8100")
  fmt.Println("hello")
  blogpost := models.CreateBlogpost(2,"Hello World","How are you doing today",890)
  fmt.Println(blogpost.Title)
  fmt.Println(blogpost.GetContent())
}
