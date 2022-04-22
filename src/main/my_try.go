package main

import (
	"encoding/xml"
	"crypto/md5"
	"fmt"
	"gee"
	"go_sql"
	"html/template"
	"io"
	"log"
	. "my_utils"
	"net/http"
	"os"
	"regexp"
	"runtime"
	"session"
	"strconv"
	"time"
)

func Logger(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL.Path)
}

func Home(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // 解析参数，默认是不会解析的
	w.Header()
	para := mux.Get_para("GET", "/home/:name")
	fmt.Fprintf(w, "<h1>Hello %s!</h1>", para["name"]) // 这个写入到 w 的是输出到客户端的

}

func Student(w http.ResponseWriter, r *http.Request) {
	if r.URL.RawQuery == "" {
		t, _ := template.ParseFiles("../usr/html/student.html")
		t.Execute(w, nil)
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
			//	log.Fatal("queryRowDemo: ", err)
			fmt.Fprintf(w, "<p>name error</p>")
		} else {
			Id := strconv.Itoa(student.Id)
			Tot_cred := strconv.Itoa(student.Tot_cred)
			str := Id + "<br/>" + student.Name + "<br/>" + student.Dept_name + "<br/>" + Tot_cred
			//not "\n"
			fmt.Fprintf(w, str)
		}
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

func Sub(w http.ResponseWriter, r *http.Request) {
	var expect string
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New() //! add timestamp to avoid dup
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		expect = token
		t, _ := template.ParseFiles("../usr/html/sub.html")
		t.Execute(w, token)
	} else {
		// 请求的是登录数据，那么执行登录的逻辑判断
		r.ParseForm()
		token := r.Form.Get("token")
		if token != "" {
			if token == expect {
				fmt.Println("username length:", len(r.Form["username"][0]))
				fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username"))) // 输出到服务器端
				fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
				fmt.Fprintln(w, r.Form.Get("username")) // 输出到客户端
				expect = ""
			}
			// 验证 token 的合法性
		} else {
			fmt.Fprintf(w, "<p>error</p>")
			// 不存在 token 报错
		}
	}
}

func Tip(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("../usr/html/tip.html")
	log.Println(t.Execute(w, nil))
}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) // 获取请求的方法
	//expiration := time.Now()
	//expiration = expiration.AddDate(1, 0, 0)
	//cookie := http.Cookie{Name: "username", Value: "astaxie", Expires: expiration}
	//http.SetCookie(w, &cookie)
	sess := globalSessions.SessionStart(w, r)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("../usr/html/login.html")
		h := md5.New()
		salt := "astaxie%^7&8888"
		io.WriteString(h, salt+time.Now().String())
		token := fmt.Sprintf("%x", h.Sum(nil))
		sess.Set("token", token)
		p := make(map[string]interface{})
		p["username"] = sess.Get("username")
		p["token"] = token
		err := t.Execute(w, p)
		Err("execute", err)
	} else {
		t, _ := template.ParseFiles("../usr/html/login.html")
		err := r.ParseForm()
		Err("Parse:", err)
		if token := r.Form.Get("token"); token == sess.Get("token") {
			sess.Set("token", "")
			sess.Set("username", r.FormValue("username"))
			//http.Redirect(w, r, "/", 302)
			//cookie,_:=r.Cookie("Value")

			fmt.Println("username:", r.Form["username"])
			fmt.Println("password:", r.Form["password"])
			reg_password := string("^(.{0,7}|.{21,}|[^0-9]*|[^a-z]*|[^A-Z]*|[^_&!$@#%]*)$|[^0-9a-zA-Z_&!$@#%]")
			//reg_password:=string("^(?=.*[0-9])(?=.*[a-z])(?=.*[A-Z])(?=.*[&!$@#%])[^]{8,20}$")
			reg_username := string("^[a-zA-Z_*-]{1,10}$")
			m, err := regexp.MatchString(reg_username, r.FormValue("username"))
			Err("match:", err)
			if !m {
				fmt.Fprintln(w, "<script>alert('username error');</script>")
				log.Println(t.Execute(w, nil))
			}
			m, err = regexp.MatchString(reg_password, r.FormValue("password"))
			Err("match:", err)
			if m {
				fmt.Fprintln(w, "<script>alert('password error')</script>")
				log.Println(t.Execute(w, nil))
			}
			http.Redirect(w, r, "/home/xherror", 302)
			//w.Write([]byte(template.HTMLEscapeString(r.Form.Get("username"))))
			//t, _ := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
			//_ = t.ExecuteTemplate(w, "T", template.HTML("<script>alert('you have been pwned')</script>"))
			//one:template.HTMLEscapeString(r.Form.Get("username"))
			//or use r.FormValue("username") only return first one
		} else {
			fmt.Fprintln(w, "<script>请勿重复提交;</script>")
		}
	}
}

// 处理 /upload  逻辑
func Upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) // 获取请求的方法
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("../usr/html/upload.html")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("../usr/upload"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666) // 此处假设当前目录下已存在test目录
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}

func Count(w http.ResponseWriter, r *http.Request) {
	sess := globalSessions.SessionStart(w, r)
	createtime := sess.Get("createtime")
	if createtime == nil {
		sess.Set("createtime", time.Now().Unix())
	} else if (createtime.(int64) + 360) < (time.Now().Unix()) {
		globalSessions.SessionDestroy(w, r)
		sess = globalSessions.SessionStart(w, r)
	}
	ct := sess.Get("countnum")
	if ct == nil {
		sess.Set("countnum", 1)
	} else {
		sess.Set("countnum", (ct.(int) + 1))
	}
	//t, _ := template.ParseFiles("../usr/html/count.html")
	//w.Header().Set("Content-Type", "text/html")
	//t.Execute(w, sess.Get("countnum"))
	fmt.Fprintln(w, sess.Get("countnum"))
}










var mux *gee.Mux
var globalSessions *session.Manager

func main() {
	globalSessions, _ = session.NewManager("memory", "gosessionid", 3600)
	go globalSessions.GC()
	//for cpu-bonud task
	num := runtime.NumCPU()
	runtime.GOMAXPROCS(num)
	fmt.Println("Set Max CPU number:", num)

	//mux 实现了 serveHTTP 可作为 Handle interface
	mux = gee.New()
	//指定相对路径./static 为文件服务路径
	//g1:=mux.Group("/admin")
	g2 := mux.Group("/usr")
	mux.Use(Logger)
	mux.GET("/home/:name", Home)
	mux.GET("/student", Student)
	mux.GET("/adder", Adder)
	mux.GET("/sub", Sub)
	mux.POST("/sub", Sub)
	mux.GET("/adder/operate", Adder_result)
	mux.GET("/upload", Upload)
	mux.POST("/upload", Upload)
	mux.GET("/count", Count)
	//g1.GET("/tip", Tip)
	mux.GET("/login", Login)
	mux.POST("/login", Login)
	g2.GETS("/usr/*source", "/usr")
	//mux.GETS("/json", "json", "usr", "json", "exercise.json")
	//mux.GETS()
	//mux.GET("/j","json","../json/one.json")
	// HandleFunc registers the handler function for the given pattern.
	err := mux.Run(":23333")
	if err != nil {
		log.Fatal("Run:", err)
	}
}
