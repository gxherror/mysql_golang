package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

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

type MyMux struct {
	http.ServeMux
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path == "/" {
        home(w, r)
        return
    }
	if r.URL.Path == "/login" {
        login(w, r)
        return
    }
	if r.URL.Path == "/adder" {
        adder(w, r)
        return
	}
	if r.URL.Path == "/js" {
		
        return
	}
    http.NotFound(w, r)
    return
}

func adder(w http.ResponseWriter,r *http.Request){
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
	mux :=new(MyMux)
	//指定相对路径./static 为文件服务路径
	staticHandle := http.FileServer(http.Dir("../static/js"))
		//将/js/路径下的请求匹配到 ./static/js/下
	mux.Handle("/js/", staticHandle)
    //http.ListenAndServe(":9090", mux)
	//mux.HandleFunc("/", sayhelloName)       // 设置访问的路由
	//mux.HandleFunc("/login", login)         // 设置访问的路由
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
