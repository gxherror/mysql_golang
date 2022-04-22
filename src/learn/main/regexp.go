package main

import (
	"fmt"
	"html/template"
	"os"
	"regexp"
	"strings"
)

func myregexp() {
	a := "I am learning Go language"

	re, _ := regexp.Compile("[a-z]{2,4}")

	// 查找符合正则的第一个
	one := re.Find([]byte(a))
	fmt.Println("Find:", string(one))

	// 查找符合正则的所有 slice, n 小于 0 表示返回全部符合的字符串，不然就是返回指定的长度
	all := re.FindAll([]byte(a), -1)
	fmt.Println("FindAll", all)

	// 查找符合条件的 index 位置, 开始位置和结束位置
	index := re.FindIndex([]byte(a))
	fmt.Println("FindIndex", index)

	// 查找符合条件的所有的 index 位置，n 同上
	allindex := re.FindAllIndex([]byte(a), -1)
	fmt.Println("FindAllIndex", allindex)

	re2, _ := regexp.Compile("am(.*)lang(.*)")

	// 查找 Submatch, 返回数组，第一个元素是匹配的全部元素，第二个元素是第一个 () 里面的，第三个是第二个 () 里面的
	// 下面的输出第一个元素是 "am learning Go language"
	// 第二个元素是 " learning Go "，注意包含空格的输出
	// 第三个元素是 "uage"
	submatch := re2.FindSubmatch([]byte(a))
	fmt.Println("FindSubmatch", submatch)
	for _, v := range submatch {
		fmt.Println(string(v))
	}

	// 定义和上面的 FindIndex 一样
	submatchindex := re2.FindSubmatchIndex([]byte(a))
	fmt.Println(submatchindex)

	// FindAllSubmatch, 查找所有符合条件的子匹配
	submatchall := re2.FindAllSubmatch([]byte(a), -1)
	fmt.Println(submatchall)

	// FindAllSubmatchIndex, 查找所有字匹配的 index
	submatchallindex := re2.FindAllSubmatchIndex([]byte(a), -1)
	fmt.Println(submatchallindex)

	src := []byte(`
    call hello alice
    hello bob
    call hello eve
`)
	pat := regexp.MustCompile(`(?m)(call)\s+(?P<cmd>\w+)\s+(?P<arg>.+)\s*$`)
	res := []byte{}
	for _, s := range pat.FindAllSubmatchIndex(src, -1) {
		res = pat.Expand(res, []byte("$cmd('$arg')\n"), src, s)
	}
	fmt.Println(string(res))
}

//type Person struct {
//    UserName string //!upper
//}

type Person struct {
	UserName string
	Html     string
	Emails   []string
	Friends  []*Friend
}

type Friend struct {
	Fname string
}

func EmailDealWith(args ...interface{}) string {
    ok := false
    var s string
    if len(args) == 1 {
        s, ok = args[0].(string)//!
    }
    if !ok {
        s = fmt.Sprint(args...)
    }
    // find the @ symbol
    substrs := strings.Split(s, "@")
    if len(substrs) != 2 {
        return s
    }
    // replace the @ by " at "
    return (substrs[0] + " at " + substrs[1])
}

func regmain() {
	/*
	f1 := Friend{Fname: "minux.ma"}
	f2 := Friend{Fname: "xushiwei"}
	t := template.New("fieldname example")
	t = t.Funcs(template.FuncMap{"emailDeal": EmailDealWith})
	t, _ = t.Parse(`hello {{.UserName}}!
			{{with $x:="output"}}{{$x|printf "%q"}}{{end}}
			{{.Html}}
            {{range .Emails}}
                an email {{.|emailDeal}}
            {{end}}
            {{with .Friends}}     
            {{range .}}
                my friend name is {{.Fname}}
            {{end}}
            {{end}}
            `) //with类似context的概念
	p := Person{UserName: "Astaxie",
		Html:    "<alert>TEST</alert>",
		Emails:  []string{"astaxie@beego.me", "astaxie@gmail.com"},
		Friends: []*Friend{&f1, &f2}}
	t.Execute(os.Stdout, p)

	tEmpty := template.New("template test")
	tEmpty = template.Must(tEmpty.Parse("空 pipeline if demo: {{if ``}} 不会输出. {{end}}\n"))
	tEmpty.Execute(os.Stdout, nil)

	tWithValue := template.New("template test")
	tWithValue = template.Must(tWithValue.Parse("不为空的 pipeline if demo: {{if `anything`}} 我有内容，我会输出. {{end}}\n"))
	tWithValue.Execute(os.Stdout, nil)

	tIfElse := template.New("template test")
	tIfElse = template.Must(tIfElse.Parse("if-else demo: {{if `anything`}} if部分 {{else}} else部分.{{end}}\n"))
	tIfElse.Execute(os.Stdout, nil)
	*/

	s1, _ := template.ParseFiles("header.tmpl", "content.tmpl", "footer.tmpl")
    /*s1.ExecuteTemplate(os.Stdout, "header", nil)
    fmt.Println()*/
    s1.ExecuteTemplate(os.Stdout, "content", nil)
    fmt.Println()
    /*s1.ExecuteTemplate(os.Stdout, "footer", nil)
    fmt.Println()*/
    s1.Execute(os.Stdout, nil)

	//t := template.New("fieldname example")
	//t, _ = t.Parse("hello {{.UserName}}!")
	//p := Person{UserName: "Astaxie"}
	//p:=map[string]interface{}{"userName":"xherror"}
	//t.Execute(os.Stdout, p)
}
