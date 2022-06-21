package gee

import (
	."my_utils"
	"net/http"
	"time"
	"strings"
	//"html/template"
	"os"
	"io/ioutil"
)

// The HandlerFunc type is an adapter to allow the use of
// ordinary functions as HTTP handlers. If f is a function
// with the appropriate signature, HandlerFunc(f) is a
// Handler that calls f.
type HandlerFunc func(http.ResponseWriter, *http.Request)

type Routergroup struct {
	prefix string
	middlewares []HandlerFunc
	parent *Routergroup
	mux *Mux
}

//Mux 实现了 serveHTTP 可作为 Handle interface 
type Mux struct {
	*Routergroup
	http.ServeMux
	router 
	handler_404 HandlerFunc
	groups []*Routergroup

}

func handler_404(w http.ResponseWriter,r *http.Request){
	w.WriteHeader(404)
	w.Write([]byte("404 NOT FOUND"))
}


func (group *Routergroup) Set404(handler HandlerFunc){
 group.mux.handler_404=handler
}
	
func New() *Mux {
	mux:=&Mux{
	router:router{
		roots:make(map[string]*node),
		handlers:make(map[string]HandlerFunc),
		para:make(map[string](map[string]string)),
		},
	handler_404:handler_404}
	mux.Routergroup=&Routergroup{mux:mux}
	mux.groups=[]*Routergroup{mux.Routergroup}
	return mux
}


func (group *Routergroup) GET(pattern string ,handler HandlerFunc){
	group.mux.router.addRoute("GET",pattern,handler)
}


func (group *Routergroup) Get_para(method string,pattern string )(map[string]string){
	return group.mux.router.para[method + "-" + pattern]
}

func (group *Routergroup) Group(prefix string) *Routergroup{
	mux:=group.mux
	newgroup:=&Routergroup{
		prefix:group.prefix+prefix ,
		parent:group,
		mux:mux,
	}
	mux.groups = append(mux.groups, newgroup)
	return newgroup
}

/*
func (mux *Mux) GETD(pattern string ,handler HandlerFunc){
	pattern:=
	mux.addroute("GET",pattern,handler)
}
*/
func (group *Routergroup) GETS(pattern string, filepath string){
	//group.mux.router.addRoute("GET",pattern,http.FileServer(http.Dir("/usr")))
	group.mux.router.addRoute("GET",pattern,group.handler_simple("GET",pattern,filepath))
}

func(group *Routergroup) POST(pattern string ,handler HandlerFunc){
	group.mux.router.addRoute("POST",pattern,handler)
}

func (group *Routergroup) Run(addr string) (err error){
	server := &http.Server{
		Addr:         addr,
		Handler:      group.mux,
		ReadTimeout:  50 * time.Second,
		WriteTimeout: 50 * time.Second,
	  }
	err= server.ListenAndServe()
	return err
}

func(group *Routergroup)handler_simple(method string,pattern string,filepath string) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var path []string		
		para:=group.Get_para(method,pattern)
		path=strings.Split(filepath[1:]+"/"+para["source"], "/")
		filetype:=(strings.Split((path[len(path)-1]), "."))[1]
		fd,err:=os.Open(Pathjoin(path))
		Err("open:",err)
		file,err:=ioutil.ReadAll(fd)
		Err("read:",err)
		//t,_:=template.ParseFiles(Pathjoin(path))
		switch filetype {
		case "json":
			w.Header().Set("Content-Type","application/json")
		case "html":
			w.Header().Set("Content-Type","text/html")
		case "css":
			w.Header().Set("Content-Type","text/css")
		case "js":
			w.Header().Set("Content-Type","text/javascript")
		}
		w.Write(file)
		//t.Execute(w,nil)

		}
}

func(group *Routergroup) Use(middlewares...HandlerFunc){
	group.middlewares=append(group.middlewares, middlewares...)
}