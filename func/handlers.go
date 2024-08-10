package functions

import (
	"html/template"
	"net/http"
)

type Data struct {
	Result string
	banner string
	text   string
}

var D Data

type ERRORS struct {
	PageTitle string
	Message   string
	ErrCde    int
}

var ERR ERRORS

// The welcome page
func Welcom(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./templates/index.html")
	tmpl2, err2 := template.ParseFiles("./templates/errors.html")
	if err != nil || err2 != nil {
		ChooseErr(500, w)
		tmpl2.Execute(w, ERR)
		return
	}

	if r.URL.Path != "/" {
		ChooseErr(404, w)
		tmpl2.Execute(w, ERR)
		return
	}

	// fmt.Println(r.Method)
	if r.Method != "GET" {
		ChooseErr(405, w)
		tmpl2.Execute(w, ERR)
		return
	}
	tmpl.Execute(w, nil)
}

// The result page
func Last(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./templates/result.html")
	tmpl2, err2 := template.ParseFiles("./templates/errors.html")
	if err != nil || err2 != nil {
		ChooseErr(500, w)
		tmpl2.Execute(w, ERR)
		return
	}

	if r.URL.Path != "/ascii-art" {
		ChooseErr(404, w)
		tmpl2.Execute(w, ERR)
		return
	}
	// fmt.Println(r.Method)
	if r.Method != "POST" {
		ChooseErr(405, w)
		tmpl2.Execute(w, ERR)
		return
	}
	D.text = r.FormValue("ljomla")
	D.banner = r.FormValue("banner")
	// fmt.Println(d.text, d.banner)
	if D.text == "" || D.banner == "" {
		ChooseErr(400, w)
		tmpl2.Execute(w, ERR)
		return
	}

	D.Result = FS(D.banner, D.text)
	if D.Result == "ERORR" {
		ChooseErr(400, w)
		tmpl2.Execute(w, ERR)
		return
	}
	tmpl.Execute(w, D)
}

func ServeStyle(w http.ResponseWriter, r *http.Request) {
	tmpl2, err2 := template.ParseFiles("./templates/errors.html")
	if err2 != nil {
		ChooseErr(500, w)
		tmpl2.Execute(w, ERR)
		return
	}
	fs := http.StripPrefix("/styles/", http.FileServer(http.Dir("./styles")))
	//path := 
	if r.URL.Path == "/styles/" {
		ChooseErr(404, w)
		tmpl2.Execute(w, ERR)
		return
	}
	fs.ServeHTTP(w, r)
}

func ChooseErr(code int, w http.ResponseWriter) {
	if code == 404 {
		ERR.PageTitle = "Error 404"
		ERR.Message = "The page web doesn't exist\nError 404"
		ERR.ErrCde = code
		w.WriteHeader(code)
	} else if code == 405 {
		ERR.PageTitle = "Error 405"
		ERR.Message = "The method is not alloweded\nError 405"
		ERR.ErrCde = code
		w.WriteHeader(code)
	} else if code == 400 {
		ERR.PageTitle = "Error 400"
		ERR.Message = "Bad Request\nError 400"
		ERR.ErrCde = code
		w.WriteHeader(code)
	} else if code == 500 {
		ERR.PageTitle = "Error 500"
		ERR.Message = "Internal Server Error\nError 500"
		ERR.ErrCde = code
		w.WriteHeader(code)
	}
}
