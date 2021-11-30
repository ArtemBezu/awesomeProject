package main

import (
	"fmt"
	"testing"
)

type Hello interface {
	hi(s string) string
}

type A struct {
}

func (r A) meth(s string) string {
	fmt.Printf("%v", s)
	return ""
}

func (a A) hi(s string) string {
	panic("implement me")
}

func main() {
	//fmt.
	var2 := "i'm your father luck"
	a := var2
	fmt.Println(a)
	b := map[string]Hello{}
	fmt.Printf("%v", b)
}

func TestName(t *testing.T) {

}
