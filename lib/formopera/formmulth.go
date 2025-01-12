//form-data 和x-www-form-urlencoded的区别
// form-data 就是http请求中的multipart/form-data,它会将表单的数据处理为一条消息，以标签为单元，用分隔符分开。既可以上传键值对，也可以上传文件。
// 当上传的字段是文件时，会有Content-Type来表名文件类型；content-disposition，用来说明字段的一些信息；
// 由于有boundary隔离，所以multipart/form-data既可以上传文件，也可以上传键值对，它采用了键值对的方式，所以可以上传多个文件

// x-www-form-urlencoded：
// 就是application/x-www-form-urlencoded,会将表单内的数据转换为键值对，比如,name=java&age = 23

//两者获取方式也不一样，form-data

package formopera

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

// sendForm 发送表单数据,form内的数据，后台用key := r.PostFormValue("key")接收
func sendForm(address string) {
	client := &http.Client{}
	res, err := client.PostForm(address, url.Values{
		"key": []string{"values"},
	})
	if err != nil {
		return
	}
	defer res.Body.Close()
}

// x-www-form-urlencoded中的值获取
func getFormKeyQuery(r *http.Request) {
	if err := r.ParseForm(); err != nil {
		return
	}
	val := r.FormValue("key")
	fmt.Println("val", val)
}

// getFormQuery 解析form中的值，x-www-form-urlencoded中的值获取
func getFormAllQuery(r *http.Request) {
	//这里是接收query的值，需要使用ParseForm解析
	if err := r.ParseForm(); err != nil {
		return
	}
	for k, v := range r.Form {
		fmt.Printf("%s:%s\n", k, v)
	}
}

// getFormBodyVal 接收表单中body中的form值，form-data中的值获取
func getFormBodyVal(r *http.Request) {
	//这里是接收body中form表单内的元素值，ParseMultipartForm需要先调用，用来分配给接收到文件的大小，不然会异常
	r.ParseMultipartForm(20 << 20)
	for k, v := range r.MultipartForm.Value { //获取表单字段
		fmt.Printf("%s:%s\n", k, v)
	}
}

// GetFormOnceFile 解析出表单内的文件,form-data中的值获取
// 单文件内容解析，这里是指定获取文件对象名称为name的，这里的name不是文件名，而是和前端对应的那个name属性名（就是form-data的key）
// 如果想获取文件基本信息，需要获取FormFile的第二个参数*multipart.FileHeader，内部又filename和size，就是使用是对一步open方法而已
func GetFormOnceFile(name string, r *http.Request) (multipart.File, error) {
	r.ParseMultipartForm(20 << 20)
	//也需要调用ParseMultipartForm
	file, _, err := r.FormFile(name)
	if err != nil {
		return nil, err
	}
	// defer file.Close()
	return file, nil
}

// getAllFormFiles 便利获取所有文件内容,返回所有fileshand
func getAllFormFiles(r *http.Request) []*multipart.FileHeader {
	files := []*multipart.FileHeader{}
	r.ParseMultipartForm(20 << 20)
	//获取表单中的文件
	//多个同时接受
	if r.MultipartForm == nil || r.MultipartForm.File == nil {
		logrus.WithFields(logrus.Fields{"form": "nil form"}).Warn("formmulth")
		return []*multipart.FileHeader{}
	}
	for _, v := range r.MultipartForm.File {
		for _, f := range v {
			files = append(files, f)
		}
	}
	return files
}

// GetAllFormFiles 便利获取所有文件内容,返回所有fileshand
func GetAllFormFiles(r *http.Request) []*multipart.FileHeader {
	return getAllFormFiles(r)
}

func UploadForm() {
	url := "https://api.cloudflare.com/client/v4/accounts/97a92c44bb6ff4acf9ea84909a9056b4/images/v1"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, err := os.Open("/Users/shitingbao/Downloads/aa.png")
	if err != nil {
		return
	}
	defer file.Close()
	part1, err := writer.CreateFormFile("file", filepath.Base("/Users/shitingbao/Downloads/aa.png"))
	if err != nil {
		return
	}

	if _, err := io.Copy(part1, file); err != nil {
		return
	}

	_ = writer.WriteField("metadata", "{\"name\":\"aab\"}")
	if err := writer.Close(); err != nil {
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", "Bearer 0Wte06ihCQ5LpaOu0phmQpv7Zclmr7qzwYi9tKe4")
	req.Header.Add("Content-Type", "multipart/form-data")
	req.Header.Set("Content-Type", writer.FormDataContentType()) // 注意需要完整的form类型
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
