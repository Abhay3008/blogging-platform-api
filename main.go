package main

import (
	"fmt"
)

const port = "8080"

func main() {

	// server := &http.Server{
	// 	Addr:    ":" + port,
	// 	Handler: routes(),
	// }

	// log.Fatal(server.ListenAndServe())
	doc := blogpost{
		Title:    "Test2",
		Content:  "test content",
		Category: "Test",
		Tags:     []string{"test1"},
	}
	a := InitDbConnection()
	fmt.Printf("%T\n", a)
	InsertDocument(a, "testcol", doc)

}

// func routes() http.Handler {
// 	r := chi.NewRouter()

// 	r.POST("/posts", PostHandler)

// 	return r
// }

type blogpost struct {
	Id        string   `json:"_id,omitempty"`
	Title     string   `json:"title"`
	Content   string   `json:"content"`
	Category  string   `json:"category"`
	CreatedAt string   `json:"createdAt,omitempty"`
	UpdatedAt string   `json:"updatedAt,omitempty"`
	Tags      []string `json:"tags"`
}

// func PostHandler(w http.ResponseWriter, r *http.Request) {
// 	var post blogpost
// 	json.Unmarshal(r.Body.Read(), &post)
// 	db := InitDbConnection()
// 	// fmt.Printf("%T\n", a)
// 	InsertDocument(db, "testcol", post)

// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(output)

// }
