package main
import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
    "net/http"
    "html/template"
    "log"
    "github.com/gorilla/mux"
)
type Product struct{
    Id int
    Model string
    Company string
    Price int
}
var database *sql.DB

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    _, err := database.Exec("delete from products where id = $1", id)
    if err != nil{
        log.Println(err)
        panic(err)
        fmt.Println(err, id)
    }

    http.Redirect(w, r, "/", 301)
}

func EditPage(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    row := database.QueryRow("select * from products where id = $1", id)
    prod := Product{}
    err := row.Scan(&prod.Id, &prod.Model, &prod.Company, &prod.Price)
    if err != nil{
        log.Println(err)
        http.Error(w, http.StatusText(404), http.StatusNotFound)
    }else
    {
        tmpl, _ := template.ParseFiles("edit.html")
        tmpl.Execute(w, prod)
    }
}

func EditHandler(w http.ResponseWriter, r *http.Request) {
    err := r.ParseForm()
    if err != nil {
        log.Println(err)
    }
    id := r.FormValue("id")
    model := r.FormValue("model")
    company := r.FormValue("company")
    price := r.FormValue("price")

    _, err = database.Exec("update products set model=$1, company=$2, price = $3 where id = $4",
        model, company, price, id)

    if err != nil {
        log.Println(err)
    }
    http.Redirect(w, r, "/", 301)
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {

        err := r.ParseForm()
        if err != nil {
            log.Println(err)
        }
        model := r.FormValue("model")
        company := r.FormValue("company")
        price := r.FormValue("price")

        _, err = database.Query("insert into products (model, company, price) values ($1, $2, $3);",
          model, company, price)
        fmt.Println(model, company, price)
        if err != nil {
            log.Println(err)
        }
        http.Redirect(w, r, "/", 301)
    }else{
        http.ServeFile(w,r, "create.html")
    }
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {

    rows, err := database.Query("select * from products")
    if err != nil {
        log.Println(err)
    }

    defer rows.Close()
    products := []Product{}

    for rows.Next(){
        p := Product{}
        err := rows.Scan(&p.Id, &p.Model, &p.Company, &p.Price)
        if err != nil{
            fmt.Println(err)
            continue
        }
        products = append(products, p)
    }

    tmpl, _ := template.ParseFiles("index.html")
    tmpl.Execute(w, products)
}

func main() {
    // conn := "postgresql://postgres:aa4401@localhost:5432/fintech"
    conn := "postgresql://gen_user:K5m_5u5tAB0%26lp@85.193.90.169:5432/default_db"
    // conn := "postgresql://neondb_owner:npg_GoPNxHe0pzm4@ep-rapid-glitter-a9y5kqy5-pooler.gwc.azure.neon.tech/neondb?sslmode=require"

    db, err := sql.Open("postgres", conn)
    if err != nil {
        log.Println(err)
    }
    database = db
    defer db.Close()
    router := mux.NewRouter()
    router.HandleFunc("/", IndexHandler)
    router.HandleFunc("/create", CreateHandler)
    router.HandleFunc("/edit/{id:[0-9]+}", EditPage).Methods("GET")
    router.HandleFunc("/edit/{id:[0-9]+}", EditHandler).Methods("POST")
    router.HandleFunc("/delete/{id:[0-9]+}", DeleteHandler)
    router.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request){
        model := r.URL.Query().Get("model")
        company := r.URL.Query().Get("company")
        fmt.Println(w, "Model: %s Company: %s", model, company)
        fmt.Println("Model: {", model, "}", "Company: {", company, "}")
    })
    http.Handle("/",router)
    fmt.Println("Server is listening...")
    http.ListenAndServe(":80", nil)

}


