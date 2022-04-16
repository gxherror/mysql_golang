package gee

import (
	"io/ioutil"
	."my_utils"
	"net/http"
	"os"
	"time"
)

// The HandlerFunc type is an adapter to allow the use of
// ordinary functions as HTTP handlers. If f is a function
// with the appropriate signature, HandlerFunc(f) is a
// Handler that calls f.
type HandlerFunc func(http.ResponseWriter, *http.Request)

type Routergroup struct {
	prefix string
	middlewares []HandlerFunc
	parent []*Routergroup
	mux *Mux
}

//Mux 实现了 serveHTTP 可作为 Handle interface 
type Mux struct {
	http.ServeMux
	router 
	handler_404 HandlerFunc
	group []*Routergroup

}

func handler_404(w http.ResponseWriter,r *http.Request){
	w.WriteHeader(404)
	w.Write([]byte("404 NOT FOUND"))
}


func (mux *Mux) Set404(handler HandlerFunc){
 mux.handler_404=handler
}
	
func New() *Mux {
	return &Mux{
		router:router{
			roots:make(map[string]*node),
			handlers:make(map[string]HandlerFunc),
			para:make(map[string](map[string]string)),
			},
		handler_404:handler_404}
}


func (mux *Mux) GET(pattern string ,handler HandlerFunc){
	mux.router.addRoute("GET",pattern,handler)
}


func (mux *Mux) Get_para(method string,pattern string )(map[string]string){
	return mux.router.para[method + "-" + pattern]
}

func (mux *Mux) Group(prefix string) *Routergroup{
	newgroup:=&Routergroup{
		prefix:prefix ,
		parent:,
	}
}

/*
func (mux *Mux) GETD(pattern string ,handler HandlerFunc){
	pattern:=
	mux.addroute("GET",pattern,handler)
}
*/
func (mux *Mux) GETS(pattern string ,filetype string , filepath...string){
	mux.router.addRoute("GET",pattern,handler_simple(filetype,filepath))
}

func (mux *Mux) POST(pattern string ,handler HandlerFunc){
	mux.router.addRoute("POST",pattern,handler)
}

func (mux *Mux) Run(addr string) (err error){
	server := &http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 20 * time.Second,
	  }
	err= server.ListenAndServe()
	return err
}

func handler_simple(filetype string , filepath []string)HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fd,err:=os.Open(Pathjoin(filepath))
		Err("open:",err)
		file,err:=ioutil.ReadAll(fd)
		Err("read:",err)
		switch filetype {
		case "json":
			w.Header().Set("Content-Type","application/json")
		case "html":
			w.Header().Set("Content-Type","text/html")
		}
		w.Write(file)
		}
}

