package basic

// struct mirip seperti class, menggunakan kata kunci --> type namaStruct struct{isi property}
type Post struct {
	Title string
	Content string
	Owner User
}

// constructor --> fungsi yang membentuk struct
func NewPost(title string, content string, user User) *Post {
	return &Post{
		Title: title,
		Content: content,
		Owner: user,
	}
}

// membuat sebuah method dari sebuah struct
func (item *Post) EditContetnt(newContent string){
	item.Content = newContent
}