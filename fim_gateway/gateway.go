package main

import (
	"bytes"
	"encoding/json"
	"fim_server/common/etcd"
	"flag"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/zeromicro/go-zero/core/conf"
)

var ServiceMap = map[string]string{
	"auth":  "http://127.0.0.1:20021",
	"user":  "http://127.0.0.1:20022",
	"chat":  "http://127.0.0.1:20023",
	"group": "http://127.0.0.1:20024",
}

//	type ReqData struct {
//		UserName string `json:"userName"`
//		Passward string `json:"password"`
//	}
type Data struct {
	Code int    `josn:code`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func toJson(data Data) []byte {
	byteData, _ := json.Marshal(data)
	return byteData
}

type Config struct {
	Etcd string
}

var configFile = flag.String("f", "./settings.yaml", "the config file")

type RoleInfo struct {
	UserID int    `json:"userID"`
	Role   int    `json:"role"`
	Data   string `json:"data"`
}

type RoleResponse struct {
	Code int       `json:"code"`
	Msg  string    `json:"msg"`
	Data *RoleInfo `json:"data"`
}

func certificate(url string, req *http.Request) (is_certificate bool) {
	var origin_address string
	pattern := `:\d+(/.*)`
	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return false
	}
	match := re.FindStringSubmatch(url)
	if len(match) > 1 {
		fmt.Println("Matched path:", match[1])
		origin_address = match[1]
	} else {
		fmt.Println("No match found.")
		return false
	}
	certificate_url := "http://127.0.0.1:20021/api/auth/authentication"
	proxyReq, _ := http.NewRequest("POST", certificate_url, nil)
	proxyReq.Header = req.Header
	proxyReq.Header.Set("origin_url", origin_address)
	proxyResponse, err := http.DefaultClient.Do(proxyReq)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return false
	}
	if proxyResponse.StatusCode == 200 {
		body, err := io.ReadAll(proxyResponse.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return false
		}
		var RoleResponse RoleResponse
		json.Unmarshal(body, &RoleResponse)
		if RoleResponse.Data == nil {
			fmt.Println("认证通过")
		} else {
			fmt.Println("角色为", RoleResponse.Data.UserID)
			fmt.Println("权限为", RoleResponse.Data.Role)
			req.Header.Set("userID", fmt.Sprintf("%d", RoleResponse.Data.UserID))
			req.Header.Set("role", fmt.Sprintf("%d", RoleResponse.Data.Role))
		}
		return true
	} else {
		fmt.Println(proxyResponse.StatusCode)
		return false
	}
}

func forward_service(url string, req *http.Request, res http.ResponseWriter) {
	bodyBytes, err := io.ReadAll(req.Body)
	if err != nil {
		res.Write(toJson(Data{Code: 7, Msg: "服务错误"}))
		return
	}
	req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes)) // replace request body
	proxyReq, _ := http.NewRequest(req.Method, url, bytes.NewBuffer(bodyBytes))
	proxyReq.Header = req.Header
	remoteAddr := strings.Split(req.RemoteAddr, ":")
	if len(remoteAddr) != 2 {
		res.Write(toJson(Data{Code: 7, Msg: "服务错误"}))
		return
	}

	fmt.Printf("%s%s =>  %s\n", remoteAddr[0], ":9000", url)
	// proxyReq.Header.Set("X-Forwarded-For", remoteAddr[0])
	// proxyReq.Header.Set("Content-Type", "application/json; charset=utf-8")
	proxyReq.Header.Set("Content-Length", strconv.Itoa(len(bodyBytes))) // set content length

	proxyResponse, err := http.DefaultClient.Do(proxyReq)
	if err != nil {
		res.Write(toJson(Data{Code: 7, Msg: "服务错误"}))
		return
	}
	io.Copy(res, proxyResponse.Body)
}

func gateway(res http.ResponseWriter, req *http.Request) {
	flag.Parse()
	var c Config
	conf.MustLoad(*configFile, &c)
	p := req.URL.Path
	regex, _ := regexp.Compile(`/api/(.*?)/`)
	list := regex.FindStringSubmatch(p)
	if len(list) != 2 {
		res.Write(toJson(Data{
			Code: 7,
			Msg:  "服务错误",
		}))
		return
	}
	addr := etcd.GetServiceAddr(c.Etcd, list[1]+"_api")
	if addr == "" {
		res.Write(toJson(Data{Code: 7, Msg: "服务错误"}))
		return
	}
	url := "http://" + addr + req.URL.String()
	fmt.Println(url)
	is_certificate := certificate(url, req)
	if is_certificate == false {
		res.Write(toJson(Data{Code: 7, Msg: "认证错误"}))
		return
	}
	forward_service(url, req, res)
	return
}

func main() {
	// 回调函数
	http.HandleFunc("/", gateway)
	// 绑定服务
	fmt.Printf("fim_gateway 运行在：%s\n", "http://127.0.0.1:9000")
	http.ListenAndServe(":9000", nil)
}
