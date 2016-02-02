package models

type Blogpost struct{
  Id int
  Title string
  content string
  Date int

}

func (blogpost Blogpost) GetContent() string {
  return blogpost.content
}

func CreateBlogpost(id int,title string,content string,date int) *Blogpost {
  blogpost := new(Blogpost)
  blogpost.Id     = id
  blogpost.Title  = title
  blogpost.content= content
  blogpost.Date   = date
  return blogpost
}
