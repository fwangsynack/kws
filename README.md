A simple go web framework. 

```go


func main() {

        mux := kws.NewKMux()
        mux.Use(middleware.Logger)
        mux.Use(middleware.WithValue("book-name", "Go web programming"))

        mux.HandleFunc("/", homePage)
        
        mux.HandleFunc("/hello", hello)

        helloHandler := HelloHandler{}
        mux.Handle("/handler/", helloHandler)  // http://localhost:8005/handler
        
        files := http.FileServer(http.Dir("public"))
        mux.Handle("/static/", http.StripPrefix("/static/", files)) // http://localhost:8005/static/

        port := 8005
        fmt.Println("run at port: " + strconv.Itoa(port))

        server := http.Server{
                Addr:         ":" + strconv.Itoa(port),
                ReadTimeout:  60 * time.Second,
                WriteTimeout: 60 * time.Second,
                Handler:      mux,
        }
        log.Fatal(server.ListenAndServe())
}

func homePage(w http.ResponseWriter, req *http.Request) {
        fmt.Println("homePage path:", req.URL.Path)
        if req.URL.Path != "/" {
                http.NotFound(w, req) // http://localhost:8005/a
                return
        }

        htmlStr := `<H1>Home Page</H1>`
        fmt.Fprint(w, htmlStr) // http://localhost:8005  or  http://localhost:8005/
}

func hello(w http.ResponseWriter, req *http.Request) {
        query := req.URL.Query()

        ctxValue, ok := req.Context().Value("book-name").(string)
        if ok {
                fmt.Println("from context, book-name --> " + ctxValue)
        } else {
                fmt.Println("from context, cannot get value with key book-name")
        }

        name := query.Get("name")
        if name == "" {
                name = "world"
        }
        fmt.Fprint(w, "Hello, my name is ", name)
}

type HelloHandler struct{}

func (h HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "Hello Handler")
}


```
