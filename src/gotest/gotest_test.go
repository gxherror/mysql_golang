package gotest
import "testing"

func Test_Division1(t *testing.T){
	if i,e :=Division(6,2);i!=3||e!=nil{
		t.Error("pass fail")
	} else {
		t.Log("test pass")
	}
}
/*
func Test_Division2(t *testing.T){
	t.Error("fail")
}
*/
func Benchmark_Division1(b *testing.B) {
	for i:=0;i<b.N;i++{
		Division(4,5)
	}
}