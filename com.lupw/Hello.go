package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	a := 2
	b := float32(a)
	fmt.Println(b)

	// 输入URL的时候，首先浏览器会去请求DNS服务器，通过DNS获取相应的域名对应的IP，然后通过IP地址找到IP对应的服务器后，要求建立[TCP]连接，等浏览器发送完HTTP Request（请求）包后，服务器接收到请求包之后才开始处理请求包，服务器调用自身服务，返回HTTP Response（响应）包；客户端收到来自服务器的响应后开始渲染这个Response包里的主体(body)，等收到全部的内容随后断开与该服务器之间的[TCP]连接
}
