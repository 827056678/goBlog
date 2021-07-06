package main

import (
	"fmt"
	"net/http"
)

type man struct {
	name string
	age  string
}

func (this *man) whatIsIt(what string) string {
	return this.name + "是" + what
}

func newMan(config map[string]string) *man {
	return &man{name: config["name"], age: config["age"]}
}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "<h1>Hello, 这里是 goblog</h1>")
	xuHao := newMan(map[string]string{"name": "xuHao", "age": "29"})
	xuHaoIs := xuHao.whatIsIt("神")
	fmt.Fprint(w, "<h1>"+xuHaoIs+"</h1>")

	weiKai := newMan(map[string]string{"name": "weiKai", "age": "30"})
	weiKaiIs := weiKai.whatIsIt("沙雕")
	fmt.Fprint(w, "<h1>"+weiKaiIs+"</h1>")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}
