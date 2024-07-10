package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// 导包

var (
	reQQEmail = `(\d+)@qq.com`
)

// 正则表达式  d表示数字 + 表示1个或者很股多个数字 连载一起就是表示一串数字

// GetEmail 函数从指定的 URL 获取页面内容，然后从页面内容中提取所有的 QQ 邮箱地址。
func GetEmail() ([]string, error) {
	// 发送 HTTP GET 请求到指定的 URL
	resp, err := http.Get("https://tieba.baidu.com/p/6051076813?red_tag=1573533731")
	// 发送欧冠HTTP请求
	if err != nil {
		// 如果在发送请求时发生错误，返回错误
		return nil, fmt.Errorf("http.Get url: %v", err)
	}
	// 确保响应主体在函数返回后关闭
	defer resp.Body.Close()

	// 检查 HTTP 状态码
	fmt.Println(resp.StatusCode) // 可见返回的状态码是200 表示正常返回
	if resp.StatusCode != http.StatusOK {
		// 如果状态码不是 200（表示请求成功），返回错误
		return nil, fmt.Errorf("unexpected http status: %v", resp.StatusCode)
	}
	// 表示不是200的时候处理获取信息到resp.Body中

	// 读取响应主体的全部内容
	pageBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// 如果在读取响应主体时发生错误，返回错误
		return nil, fmt.Errorf("ioutil.ReadAll: %v", err)
	}

	// 将响应主体的内容转换为字符串
	pageStr := string(pageBytes)
	// 创建一个正则表达式对象，用于查找所有的 QQ 邮箱地址
	re := regexp.MustCompile(reQQEmail)
	// 查找页面内容中的所有 QQ 邮箱地址
	results := re.FindAllStringSubmatch(pageStr, -1)

	// 创建一个空的字符串切片，用于存储找到的邮箱地址
	var emails []string // 储存查找结果
	// 遍历所有的匹配结果
	for _, result := range results {
		// 将找到的邮箱地址添加到切片中
		emails = append(emails, result[0])
	}
	// 返回找到的所有邮箱地址，以及可能的错误（在这种情况下，错误总是 nil）
	return emails, nil
}

func main() {
	emails, err := GetEmail()
	// 调用函数

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// 非空判断：如果为空就输出答应错误信息

	for _, email := range emails {
		fmt.Println("email:", email)
	}
	// 遍历输出切片 前者为下标
}

// 总之：鲜花少说就是一个 正则表达式的应用：爬取所有的qq邮箱
