package view

import (
	"net/http"

	"github.com/chittip/gonews/pkg/model"
)

// IndexData ..
type IndexData struct {
	List []*model.News
}

// Index renders index view
func Index(w http.ResponseWriter, data *IndexData) {
	render(tpIndex, w, data)
}

// AdminLogin renders admin login
func AdminLogin(w http.ResponseWriter, data interface{}) {
	render(tpAdminLogin, w, data)
}

// AdminCreate renders admin login
func AdminCreate(w http.ResponseWriter, data interface{}) {
	render(tpAdminLCreate, w, data)
}

// AdminEdit renders admin login
func AdminEdit(w http.ResponseWriter, data interface{}) {
	render(tpAdminLogin, w, data)
}

// AdminList renders admin login
func AdminList(w http.ResponseWriter, data interface{}) {
	render(tpAdminLogin, w, data)
}
