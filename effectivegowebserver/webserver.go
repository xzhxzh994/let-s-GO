package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
)

//google的服务：将表单数据自动转换成图表，缺点：很难交互，需要将数据作为查询放到url中。
//程序作用：给定一小段文本，之后它将会调用图标服务器来生成二维码。这个二维码表现形式是编码文本的点格矩阵，之后我们使用手机摄像头捕获并且解释为一个字符串，这样就不需要在手机上输入URL了

var addr = flag.String("addr", ":1718", "http service address") //Q=17,R=18

//这里的 : 表示服务将监听所有可用的网络接口（即所有 IP 地址）。

var templ = template.Must(template.New("qr").Parse(templateStr))

//在 Go 语言中，template.Must 是一个辅助函数，用于简化模板初始化时的错误处理。它的作用是：如果模板解析成功，返回模板对象；如果解析失败，直接触发 panic（程序崩溃）。

func main() {
	flag.Parse()
	http.Handle("/", http.HandlerFunc(QR)) //将http请求处理函数QR与URL路径/关联，任何对/的请求都会被QR处理
	//ps:这里的handlerFunc是一个类型，所以我们做了一个类型转换，将qr转换成了对应了类型，
	err := http.ListenAndServe(*addr, nil) //这里创建一个结构体，将addr和handle传递进去，之后调用这个结构体的监听函数，负责监听
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func QR(w http.ResponseWriter, req *http.Request) {
	err := templ.Execute(w, req.FormValue("s"))
	if err != nil {
		return
	}
}

const templateStr = `
<html>
<head>
<title>QR Link Generator</title>
</head>
<body>
{{if .}}
<img src="http://chart.apis.google.com/chart?chs=300x300&cht=qr&choe=UTF-8&chl={{.}}" />
<br>
{{.}}
<br>
<br>
{{end}}
<form action="/" name=f method="GET"><input maxLength=1024 size=70
name=s value="" title="Text to QR Encode"><input type=submit
value="Show QR" name=qr>
</form>
</body>
</html>
`
