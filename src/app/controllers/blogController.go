package BlogController

import (
  _"fmt"
  _"app/models"
  "github.com/siddontang/ledisdb"
)

func GetAllBlogPosts() string {
  return "All"
}

func GetBlogPostById(id string) string  {
  return id + " hey"
}
