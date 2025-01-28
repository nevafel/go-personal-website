/*
This controller handles admin stuff and allows to redact databases.

It has some noodles in it but this is how it works.
1. Users enters on "/admin" and this controller handles GetAdminLogin()
2. Than frontend passes a form with login and password to PostAdminLogin()
3. In PostAdminLogin() backend checks that login and password are equal to
	ones that defined in .env file (without it code will panic! I will rewrite this)
4. If everything correct controller goes to GetAdminPanel() and shows admin-panel
	template

Currently page is vulnerable to bruteforce.
I will try to implement JWT in future. -d1xxe
Best practice is to move login to middleware- -TODO
*/

package controllers

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/dixxe/personal-website/iternal/pkg/repositories"
	"github.com/dixxe/personal-website/web/templates"
	"github.com/joho/godotenv"
)

// This controller has been written to control databases and blog. SECUIRITY-WARNING!
func GetAdminPanel(w http.ResponseWriter, r *http.Request) {
	posts, err := repositories.Blog.GetAllValues()

	if err != nil {
		log.Println(err)
		component := templates.AdminPanelPage([]repositories.Post{})
		component.Render(context.Background(), w)
		return
	}

	component := templates.AdminPanelPage(posts)
	component.Render(context.Background(), w)
}

// This controller has been written to add protection for admin-panel.
func GetAdminLogin(w http.ResponseWriter, r *http.Request) {
	// Respond with page that contains form for login and password.
	component := templates.LoginPage()
	component.Render(context.Background(), w)
}

// This controller handles passed form and checks for its validity.
// If everything right it passes to GetAdminPanel.
func PostAdminLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	login := r.FormValue("login")
	password := r.FormValue("password")

	if err := godotenv.Load(); err != nil {
		// I better disable admin panel at all if .env not located.
		component := templates.ErrorPage(404, "Admin-panel not configured.")
		component.Render(context.Background(), w)
		return
	}

	admin_login, _ := os.LookupEnv("LOGIN")
	admin_password, _ := os.LookupEnv("PASSWORD")
	if login == admin_login && password == admin_password {
		GetAdminPanel(w, r)
	}
}
