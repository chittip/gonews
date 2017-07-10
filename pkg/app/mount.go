package app

import "net/http"

// Mount mounts handlers to mux
func Mount(mux *http.ServeMux) {
	mux.HandleFunc("/", index) // list all new
	mux.HandleFunc("/news/", newsView)

	adminMux := http.NewServeMux()
	adminMux.HandleFunc("/login", adminLogin)
	adminMux.HandleFunc("/list", adminList)
	adminMux.HandleFunc("/create", adminCreate)
	adminMux.HandleFunc("/edit", adminEdit)

	mux.Handle("/admin/", http.StripPrefix("/admin", onlyAdmin(adminMux)))
	// /news/:path
	// /admin/login/
	// /admin/list
	// /admin/create
	// /admin/edit

}

func onlyAdmin(h http.Handler) http.Handler {
	return h
}
