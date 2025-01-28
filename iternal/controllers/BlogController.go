/*
Controller that handles everything d1xxe blog related features.
Under the hood it operates with repositories.Blog to get all information from
local database.

In a nutshell it's pretty simple.
*/
package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/dixxe/personal-website/iternal/pkg/repositories"
	"github.com/dixxe/personal-website/web/templates"
	"github.com/go-chi/chi/v5"
)

// This controller passes all posts in database in template and responds with page with all size-reduced posts.
func GetShowBlog(w http.ResponseWriter, r *http.Request) {
	posts, err := repositories.Blog.GetAllValues()
	if err != nil {
		component := templates.ErrorPage(500, "Ошибка при обработке базы данных.")
		// For some reason I can't set status code for request x-x
		log.Println(err)
		component.Render(context.Background(), w)
		return
	}

	component := templates.ShowBlogPage(posts)
	component.Render(context.Background(), w)
}

// This controller looks for specific post in database and handles all possible errors. Responds with page with post.
func GetPost(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		component := templates.ErrorPage(500, "Не удалось преобразовать id в целое число.")
		log.Println(err)
		component.Render(context.Background(), w)
		return
	}

	post, err := repositories.Blog.GetValueByID(id)

	if err != nil {
		component := templates.ErrorPage(404, "Не удалось найти нужный пост.")
		log.Println(err)
		component.Render(context.Background(), w)
		return

	}
	component := templates.ShowPost(post)
	component.Render(context.Background(), w)

}

// Controller that handles post creation. Responds with 200 if everything OK
func PostCreatePost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()                   // Populating form.
	header := r.FormValue("header") // To get value you need to specify name="header" in the form.
	content := r.FormValue("content")

	// Creating a new post with 0 Id, don't worry database handles id assignment
	// itself. And in the InsertValue() method I don't use Post.Id value
	newPost := repositories.Post{Id: 0, Header: header, Content: content}

	// Process database is async way. WARNING errors are NOT HANDLED
	go repositories.Blog.InsertValue(newPost)

	fmt.Println("Created post")
}

// Controller that handles post deletion. Responds with 200 if everything is OK
func PostDeletePost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id, err := strconv.Atoi(r.FormValue("id"))

	if err != nil {
		component := templates.ErrorPage(500, "Ошибка при переводе id (int) в string.")
		r.Response.StatusCode = 500
		log.Println(err)
		component.Render(context.Background(), w)
	}
	// Process database is async way. WARNING errors are NOT HANDLED
	go repositories.Blog.DeleteValueByID(id)
}
