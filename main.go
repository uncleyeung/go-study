package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/astaxie/beego/logs"

	"github.com/vincenty1ung/yeung-go-study/study/leedcode/leet"
)

// 创建log
var log = logs.NewLogger()

// 初始化，日志输出方式采用beego-logs study
func init() {
	// 文件输出
	// _ = logs.SetLogger(logs.AdapterFile, `{"filename":"project.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`)
	// 控制台输出
	_ = log.SetLogger(logs.AdapterConsole)

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	GetBody := r.GetBody
	log.Info("这是一个很开心的脸 %s %v", GetBody, 3)
	_, _ = fmt.Fprint(w, "你好世界")
}

func main() {
	/*
		// 函数式范式
		handleFunc()
		// 日志处理
		logHandler()
		// http处理
		httpHandler()
		// slice 指针 基础结构学习
		slicePHandler()
		// struct 指针 基础结构学习
		structPHandler()
	*/
	leet.LeetMain()
}

func handleFunc() {
	man := Man{name: "zhangbo", age: 15, length: 13}
	man.handleFunc(man.age, atLastHandleForAge)
	man.handleFunc(man.length, loggingNumHandler(atLastHandleForLength))
}

func httpHandler() {
	// ===================test====================

	// 创建一个新的http路由管理器
	mux := http.NewServeMux()
	mux.HandleFunc("/index", indexHandler)
	// handler
	mux.Handle("/hello", loggingHandler(&HtmlHandler{}))
	// handler处理器
	mux.HandleFunc("/hello1", hello)
	// 只是监听8080端口
	_ = http.ListenAndServe(":8080", mux)

	// Clinet -> Requests ->  [Multiplexer(router) -> handler  -> Response -> Clinet
}

func logHandler() {
	// ===================test====================
	// an official log.Logger with prefix ORM
	log.Info("======================")
	log.Debug("my book is bought in the year of ", 2016)
	log.Info("this %s cat is %v years old", "yellow", 3)
	log.Warn("json is a type of kv like", map[string]int{"key": 2016})
	log.Error("1024", "is a very", "good game")
	log.Info("======================")
}

/*
1.创建自定义的handler
定义一个结构体，并且去实现http.handler接口
*/
type HtmlHandler struct {
}

func (htmlHandler *HtmlHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 声明响应的数据为html
	w.Header().Set("Content-Type", "text/html")
	logs.Info("设置响应头%s 为：%s", "Content-Type", "text/html")

	html := `<doctype html>
        <html>
        <head>
          <title>Hello World</title>
        </head>
        <body>
        <p>
          <a href="/welcome">Welcome</a> |  <a href="/message">Message</a>
        </p>
        </body>
		</html>`
	_, _ = fmt.Fprint(w, html)
}

/*
2.创建handler处理器
*/
func hello(w http.ResponseWriter, r *http.Request) {
	// 声明响应的数据为html
	w.Header().Set("Content-Type", "text/html")
	logs.Info("设置响应头%s 为：%s", "Content-Type", "text/html")

	html := `<doctype html>
        <html>
        <head>
          <title>Hello World</title>
        </head>
        <body>
        <p>
          <a href="/welcome">Welcome</a> |  <a href="/message">Message</a>
        </p>
        </body>
		</html>`
	_, _ = fmt.Fprint(w, html)
}

/*
3.函数式范式
*/
func middlewareHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// 执行handler之前的逻辑
			next.ServeHTTP(w, r)
			// 执行完毕handler后的逻辑
		},
	)
}
func loggingHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			log.Info("开始请求，请求类型：%s 请求地址：%s 请求地址：%s", r.Method, r.URL.Path, r.Host)
			next.ServeHTTP(w, r)
			log.Info("请求完成，请求地址：%s 耗时：%v", r.URL.Path, time.Since(start))
		},
	)

	/*return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Info("开始请求，请求类型：%s 请求地址：%s 请求地址：%s", r.Method, r.URL.Path, r.Host)
		next.ServeHTTP(w, r)
		log.Info("请求完成，请求地址：%s 耗时：%v", r.URL.Path, time.Since(start))
	}*/
}

// ===========================================type func=================================================
type Man struct {
	name   string
	age    int
	length int
}
type NumHandler interface {
	// 最终处理at last
	handle(num int)
}

type HandlerInt func(num int)

func (handlerInt HandlerInt) handle(num int) {
	handlerInt(num)
	// handlerInt.handle(num)
}

func atLastHandleForAge(age int) {
	logs.Info("我的年龄是：%v岁", age)
}
func atLastHandleForLength(length int) {
	logs.Info("我的长度是：%v米", length)
}

func (man Man) handleFunc(num int, atLastHandle func(num int)) {
	// 第二个参数是把传入的函数atLastHandle 强转成 HandlerInt类型，这样atLastHandle就实现了NumHandler接口。
	// var handlerInt NumHandler  = HandlerInt(atLastHandle)
	atLastCall(num, HandlerInt(atLastHandle))
}
func atLastCall(int int, numHandler NumHandler) {
	numHandler.handle(int)
}

func loggingNumHandler(handlerInt HandlerInt) HandlerInt {
	/*
		一下这个位置将func(num int){}，强转为type HandlerInt
		return HandlerInt(func(num int) {
			//之前
			log.Info("之前")
			handlerInt.handle(num)
			//之后
			log.Info("之后")
		})
	*/
	// 等价于函数
	return func(num int) {
		// 之前
		log.Info("之前")
		handlerInt.handle(num)
		// 之后
		log.Info("之后")
	}
}

type nextNode struct {
	i    int
	next *nextNode
}

func slicePHandler() {
	int32s := make([]int32, 0)
	int32s = append(int32s, 1)
	int32s = append(int32s, 2)
	int32s = append(int32s, 3)
	int32s = append(int32s, 4)
	int32s = append(int32s, 5)
	// p1 := &int32s
	fmt.Println("p1:")
	// fmt.Println(&p1)
	// fmt.Println(fmt.Sprintf("%p", int32s))
	// fmt.Println(&p1)
	fmt.Println(fmt.Sprintf("%p", &int32s))
	fmt.Println(fmt.Sprintf("len:%d", len(int32s)))
	fmt.Println(fmt.Sprintf("cap:%d", cap(int32s)))
	fmt.Println(int32s)
	p1 := &int32s
	fmt.Println(fmt.Sprintf("指针的地址1:%v,共同指向结构体地址%p", &p1, p1))
	sfq(p1)
	fmt.Println(fmt.Sprintf("len:%d", len(int32s)))
	fmt.Println(fmt.Sprintf("cap:%d", cap(int32s)))
	fmt.Println(int32s)
}
func sfq(list *[]int32) {
	fmt.Println("p2:")
	fmt.Println(fmt.Sprintf("指针的地址2:%v,共同指向结构体地址%p", &list, list))
	// fmt.Println(&list)
	*list = append(*list, 6)
	*list = append(*list, 6)
	*list = append(*list, 6)
	*list = append(*list, 6)
	fmt.Println(fmt.Sprintf("指针的地址2:%v,共同指向结构体地址%p", &list, list))
	fmt.Println(fmt.Sprintf("len:%d", len(*list)))
	fmt.Println(fmt.Sprintf("cap:%d", cap(*list)))
}
func sfq1(list []int32) {
	fmt.Println("p2:")
	// fmt.Println(&p2)
	fmt.Println(fmt.Sprintf("p2:%p", &list))
	list = list[2:4]
	fmt.Println(list)
	fmt.Println(fmt.Sprintf("len:%d", len(list)))
	fmt.Println(fmt.Sprintf("cap:%d", cap(list)))
	list = append(list, 7)
	list = append(list, 6)
	list = append(list, 8)
	list = append(list, 9)
	list = append(list, 10)
	fmt.Println(fmt.Sprintf("p2:%p", &list))
	fmt.Println(fmt.Sprintf("len:%d", len(list)))
	fmt.Println(fmt.Sprintf("cap:%d", cap(list)))
	fmt.Println(list)
}

/*
!!!!!所有的传递传递的都是副本
*/
func structPHandler() {
	// [1,2,2,3,4] 去重
	node := nextNode{
		i: 1, next: &nextNode{
			i: 2, next: &nextNode{
				i: 2, next: &nextNode{
					i: 3, next: &nextNode{
						4, nil,
					},
				},
			},
		},
	}
	fmt.Println(fmt.Sprintf("%+v", node))
	// fmt.Println(&node)
	p1 := &node
	fmt.Println(fmt.Sprintf("指针的地址1:%v,共同指向结构体地址%p", &p1, p1))
	handler(p1)
	fmt.Println(fmt.Sprintf("%+v", node))
}

func handler(node *nextNode) {
	fmt.Println(fmt.Sprintf("指针的地址2:%v,共同指向结构体地址%p", &node, node))
	if node == nil {
		return
	}
	for node.next != nil {
		if node.i == node.next.i {
			node.next = node.next.next
		} else {
			node = node.next
		}
	}
	fmt.Println(node)
}
