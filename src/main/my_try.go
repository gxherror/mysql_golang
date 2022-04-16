package main

import (
	"fmt"
	"gee"
	"go_sql"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	. "my_utils"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func Home(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // 解析参数，默认是不会解析的
	//fmt.Println(r.Form) // 这些信息是输出到服务器端的打印信息
	//fmt.Println(r.Form["url_long"])
	w.Header()
	para:=mux.Get_para("GET","/home/:name")
	fmt.Fprintf(w, "<h1>Hello %s!</h1>",para["name"]) // 这个写入到 w 的是输出到客户端的
}

func Usr(w http.ResponseWriter, r *http.Request) {
	para:=mux.Get_para("GET","/usr/*source")
	path:=strings.Split("usr/"+para["source"], "/")
	filetype:=(strings.Split((path[len(path)-1]), "."))[1]
	fd,err:=os.Open(Pathjoin(path))
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


func Student(w http.ResponseWriter, r *http.Request) {
	if r.URL.RawQuery == "" {
		t, _ := template.ParseFiles("../html/student.html")
		log.Println(t.Execute(w, nil))
	} else {
		r.ParseForm()
		name := r.FormValue("name")
		fmt.Println(name)
		db, err := go_sql.InitDB()
		if err != nil {
			// handle error http.Error() for example
			log.Fatal("Init: ", err)
		}
		student, err := db.QueryRowDemo(name)
		if err != nil {
			// handle error http.Error() for example
			log.Fatal("queryRowDemo: ", err)
		}
		Id := strconv.Itoa(student.Id)
		Tot_cred := strconv.Itoa(student.Tot_cred)
		str := Id + "<br/>" + student.Name + "<br/>" + student.Dept_name + "<br/>" + Tot_cred
		//not "\n"
		fmt.Fprintf(w, str)
	}
}

func Old_adder(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		t, _ := template.ParseFiles("../html/adder.html")
		data := map[string]string{
			"result": "",
		}
		log.Println(t.Execute(w, data))
	} else {
		t, _ := template.ParseFiles("../html/adder.html")
		err := r.ParseForm() // 解析 url 传递的参数，对于 POST 则解析响应包的主体（request body）
		if err != nil {
			// handle error http.Error() for example
			log.Fatal("ParseForm: ", err)
		}
		first, _ := strconv.Atoi(r.FormValue("first"))
		second, _ := strconv.Atoi(r.FormValue("second"))
		result := strconv.Itoa(first + second)
		data := map[string]string{
			"result": result,
		}
		t.Execute(w, data)
	}
}

func Adder(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("../usr/html/adder@1.1.0.html")
	if err != nil {
		// handle error http.Error() for example
		log.Fatal("ParseFiles: ", err)
	}
	log.Println(t.Execute(w, nil))
}

func Adder_result(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm() // 解析 url 传递的参数，对于 POST 则解析响应包的主体（request body）
	if err != nil {
		// handle error http.Error() for example
		log.Fatal("ParseForm: ", err)
	}
	fmt.Printf(r.URL.RawQuery)
	first, _ := strconv.Atoi(r.FormValue("Num1"))
	second, _ := strconv.Atoi(r.FormValue("Num2"))
	fmt.Println(first, second)
	result := strconv.Itoa(first + second)
	io.WriteString(w, result)
}

func Tip(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("../usr/html/tip.html")
	log.Println(t.Execute(w, nil))
}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) // 获取请求的方法
	if r.Method == "GET" {
		t, _ := template.ParseFiles("../usr/html/login.html")

		log.Println(t.Execute(w, nil))
	} else {
		err := r.ParseForm() // 解析 url 传递的参数，对于 POST 则解析响应包的主体（request body）
		if err != nil {
			// handle error http.Error() for example
			log.Fatal("ParseForm: ", err)
		}
		// 请求的是登录数据，那么执行登录的逻辑判断
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}

var mux *gee.Mux

func main() {
	//mux 实现了 serveHTTP 可作为 Handle interface
	mux = gee.New()
	//指定相对路径./static 为文件服务路径
	//staticHandle := http.FileServer(http.Dir("../static/js"))
	//将/js/路径下的请求匹配到 ./static/js/下
	//mux.Handle("/js/", staticHandle)

	mux.GET("/home/:name", Home)
	mux.GET("/student", Student)
	mux.GET("/adder", Adder)
	mux.GET("/adder/operate", Adder_result)
	mux.GET("/tip", Tip)
	mux.GET("/login", Login)
	mux.GET("/usr/*source",Usr)
	//mux.GETS("/json", "json", "usr", "json", "exercise.json")
	//mux.GETS()
	//mux.GET("/j","json","../json/one.json")
	// HandleFunc registers the handler function for the given pattern.
	err := mux.Run(":2333")
	if err != nil {
		log.Fatal("Run:", err)
	}
}
