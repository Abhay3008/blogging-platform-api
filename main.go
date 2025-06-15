package main

import (
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const port = "8080"

func main() {

	// server := &http.Server{
	// 	Addr:    ":" + port,
	// 	Handler: routes(),
	// }

	// log.Fatal(server.ListenAndServe())
	doc := blogpost{
		Title:    "Test3",
		Content:  "test content is updated again",
		Category: "Test3",
		Tags:     []string{"test1"},
	}
	a := InitDbConnection()
	fmt.Printf("%T\n", a)
	// InsertDocument(a, "testcol", doc)
	oid, err := bson.ObjectIDFromHex("684efb41f49d8b764c210c7b")
	if err != nil {
		log.Fatal("Invalid ObjectID:", err)
	}
	filter := bson.M{"_id": oid}
	UpdateDocument(a, "testcol", filter, doc)

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
