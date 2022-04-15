package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
	"go_sql"
)
// The HandlerFunc type is an adapter to allow the use of
// ordinary functions as HTTP handlers. If f is a function
// with the appropriate signature, HandlerFunc(f) is a
// Handler that calls f.
type HandlerFunc func(http.ResponseWriter, *http.Request)

func home(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       // 解析参数，默认是不会解析的
	//fmt.Println(r.Form) // 这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	//fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") // 这个写入到 w 的是输出到客户端的
}

type Mux struct {
	http.ServeMux
}

func (mux *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf(r.URL.Path)
    if r.URL.Path == "/" {
        home(w, r)
        return
    }
	if r.URL.Path == "/login" {
        login(w, r)
        return
    }
	if r.URL.Path == "/tip" {
        tip(w, r)
        return
    }
	if r.URL.Path == "/adder" {
        adder_show_110(w, r)
        return
	}
	if r.URL.Path == "/student" {
        student(w, r)
        return
	}
	if r.URL.Path == "/adder/operate" {
        adder_110(w, r)
        return
	}
	if r.URL.Path == "/js" {
		
        return
	}
    http.NotFound(w, r)
    return
}



func student(w http.ResponseWriter,r *http.Request){
	if r.URL.RawQuery == "" {
	t, _ := template.ParseFiles("../html/student.html")
	log.Println(t.Execute(w, nil))
	}else {
		r.ParseForm()
		name:=r.FormValue("name")
		fmt.Println(name)
		go_sql.initDB()
		student,err:=go_sql.queryRowDemo(name)
		if err != nil {
			// handle error http.Error() for example
		   log.Fatal("queryRowDemo: ", err)
		}
		id:=strconv.Itoa(student.id)
		tot_cred:=strconv.Itoa(student.tot_cred)
		str:=id+"<br/>"+student.name+"<br/>"+student.dept_name+"<br/>"+tot_cred
		//not "\n"
		fmt.Fprintf(w,str) 
	}
}

func adder_100(w http.ResponseWriter,r *http.Request){
	
	if r.Method == "GET" {
        t, _ := template.ParseFiles("../html/adder.html")
		data:=map[string]string{
		"result":"",	
		}
        log.Println(t.Execute(w, data))
    }else{	
	t, _ := template.ParseFiles("../html/adder.html")
	err := r.ParseForm()   // 解析 url 传递的参数，对于 POST 则解析响应包的主体（request body）
	if err != nil {
	   // handle error http.Error() for example
	  log.Fatal("ParseForm: ", err)
	}
	first,_:=strconv.Atoi(r.FormValue("first"))
	second,_:=strconv.Atoi(r.FormValue("second"))
	result:=strconv.Itoa(first+second)
	data:=map[string]string{
		"result":result,	
		}
    t.Execute(w, data)
	}
}

func adder_show_110(w http.ResponseWriter,r *http.Request){
	t, err:= template.ParseFiles("../html/adder@1.1.0.html")
	if err != nil {
		// handle error http.Error() for example
		log.Fatal("ParseFiles: ", err)
	}
	log.Println(t.Execute(w, nil))

}

func adder_110(w http.ResponseWriter,r *http.Request){
	err := r.ParseForm()   // 解析 url 传递的参数，对于 POST 则解析响应包的主体（request body）
	if err != nil {
		// handle error http.Error() for example
		log.Fatal("ParseForm: ", err)
	}
	fmt.Printf(r.URL.RawQuery)
	first,_:=strconv.Atoi(r.FormValue("Num1"))
	second,_:=strconv.Atoi(r.FormValue("Num2"))
	fmt.Println(first,second)
	result:=strconv.Itoa(first+second)
	io.WriteString(w, result)
}

func tip(w http.ResponseWriter,r *http.Request){
	t, _ := template.ParseFiles("../html/tip.html")
	log.Println(t.Execute(w, nil))	
}

func login(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method) // 获取请求的方法
    if r.Method == "GET" {
        t, _ := template.ParseFiles("../html/login.html")

        log.Println(t.Execute(w, nil))
    } else {
        err := r.ParseForm()   // 解析 url 传递的参数，对于 POST 则解析响应包的主体（request body）
        if err != nil {
           // handle error http.Error() for example
          log.Fatal("ParseForm: ", err)
        }
        // 请求的是登录数据，那么执行登录的逻辑判断
        fmt.Println("username:", r.Form["username"])
        fmt.Println("password:", r.Form["password"])
    }
}



func main() {
	//mux 实现了 serveHTTP 可作为 Handle interface 
	mux :=new(Mux)
	//指定相对路径./static 为文件服务路径
	//staticHandle := http.FileServer(http.Dir("../static/js"))
		//将/js/路径下的请求匹配到 ./static/js/下
	//mux.Handle("/js/", staticHandle)
    //http.ListenAndServe(":9090", mux)
	//mux.HandleFunc("/", sayhelloName)       // 设置访问的路由
	mux.HandleFunc("/login", login)         // 设置访问的路由

	// HandleFunc registers the handler function for the given pattern.
	server := &http.Server{
		Addr:         ":2333",
		Handler:      mux,
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 20 * time.Second,
	  }
	err:= server.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
