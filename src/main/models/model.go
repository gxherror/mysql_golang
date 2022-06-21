package models
import (
	"os"
	"io/ioutil"
	"gee/logger"
	"html/template"
	"github.com/gomarkdown/markdown"
	)

func Get_blog(path string) template.HTML {
			fd,err:=os.Open(path)
			md,err:=ioutil.ReadAll(fd)
			logger.Test_Error(err,"read file error")	
			md=markdown.NormalizeNewlines(md)
			html:=template.HTML(markdown.ToHTML(md,nil,nil))
			//content:=template.HTMLEscapeString(string(html))
			return html
}