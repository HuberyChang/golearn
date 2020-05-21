package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// net/http client

// 共用一个长连接，适用于频繁发起请求的情况
var (
	client = http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: false,
		},
	}
)

func main() {
	// resp, err := http.Get("http://127.0.0.1:9000/xxx/?name=gan&age=19")
	// if err != nil {
	// 	fmt.Printf("get url failed,err:%v\n", err)
	// 	return
	// }

	// 自己造http.get
	data := url.Values{} // url encode 编码
	urlObj, _ := url.Parse("http://127.0.0.1:9000/xxx/")
	data.Set("name", "只看见")
	data.Set("age", "190")
	queryStr := data.Encode() // URL encode之后的URL
	fmt.Println(queryStr)
	urlObj.RawQuery = queryStr
	req, err := http.NewRequest("GET", urlObj.String(), nil)
	// 发请求
	// resp, err := http.DefaultClient.Do(req)
	// if err != nil {
	// 	fmt.Printf("get url failed,err:%v\n", err)
	// 	return
	// }

	// 发短连接，禁用keep-alived，适用于请求不是很频繁的情况
	// tr := &http.Transport{
	// 	DisableKeepAlives: true,
	// }
	// client := http.Client{
	// 	Transport: tr,
	// }
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("get url failed,err:%v\n", err)
		return
	}

	defer resp.Body.Close() // 一定要关闭

	// 从resp中把服务端返回的数据读出来
	b, err := ioutil.ReadAll(resp.Body) // 在客户端读出服务端返回的响应的body
	if err != nil {
		fmt.Printf("read resp.body failed,err:%v\n", err)
		return
	}
	fmt.Println(string(b))
}
