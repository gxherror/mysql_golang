package my_utils
import (
	"runtime"
	"log"
)
func Pathjoin(arg []string) string{
	result:=".."
	sys:=runtime.GOOS
	if sys== "windows"{
		for _,v :=range(arg){
			result=result+"\\"+v
		}
	}else {
		for _,v :=range(arg){
			result=result+"/"+v
		}
	}
	return result
}

func Err(arg string,err error){
	if err!=nil{
		log.Println(arg,err)
	}
}