package util

import (
	"compress/gzip"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/gincmf/cmf/data"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

var Conf *data.ConfigDefault

func GetMd5(s string) string {
	h := md5.New()
	h.Write([]byte(Conf.Database.AuthCode + s))
	return hex.EncodeToString(h.Sum(nil))
}

func GetSha1(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func GetAbsPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))
	absPath := path[:index]
	absPath += "/"
	absPath = strings.Replace(absPath, "\\", "/", -1)
	return absPath
}

//封装请求库
func Request(method string,url string,body io.Reader,h map[string]string) (int, []byte){
	client := &http.Client{
		Timeout: time.Minute*60,
	}
	switch method {
	case "get","GET":
		method = "GET"
	case "post","POST":
		method = "POST"
	case "put","PUT":
		method = "PUT"
	case "delete","DELETE":
		method = "POST"
	}
	r,err := http.NewRequest(method,url,body)
	if err != nil {
		fmt.Println("http错误",err)
	}

	/*r.Header.Add("Host", "")
	r.Header.Add("Connection","keep-alive")
	r.Header.Add("Accept-Encoding","gzip, deflate, br")
	r.Header.Add("Content-Length","0")
	r.Header.Add("Cache-Control","no-cache")*/
	for k,v := range h{
		r.Header.Add(k,v)
	}
	response, err := client.Do(r)
	defer response.Body.Close()

	var data []byte = nil

	switch response.Header.Get("Content-Encoding") {
	case "gzip":
		reader, _ := gzip.NewReader(response.Body)
		for {
			buf := make([]byte, 1024)
			n, err := reader.Read(buf)
			if err != nil && err != io.EOF {
				panic(err)
			}
			if n == 0 {
				break
			}
			data = append(data,buf...)
		}
	default:
		data, err = ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("err",err.Error())
		}
	}
	return response.StatusCode,data
}

func Database() data.Database {
	return Conf.Database
}

