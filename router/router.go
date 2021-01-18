package router

import (
	"encoding/json"
	"net/http"
)

var posts []Post

func init() {
	posts = []Post{{Id: 1, Title: "title 1", Text: "body test"}}

}

//GetPosts
func GetPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.Marshal(posts)
	result, err := json.Marshal(posts)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error":"Error marshalling the post array" }`))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(result)

}

//AddPost
func AddPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var post Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error":"Error unmarshalling the request" }`))
		return
	}
	post.Id = len(posts) + 1
	posts = append(posts, post)
	w.WriteHeader(http.StatusOK)
	result, err := json.Marshal(post)
	w.Write(result)
}
