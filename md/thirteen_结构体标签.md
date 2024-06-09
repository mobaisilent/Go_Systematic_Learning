## **结构体标签**

实例代码：
```go
import (
	"fmt"
	"reflect"
)

type resume struct {
	Name string `info:"name" doc:"我的名字"`
	Sex  string `info:"sex"`
}

func findTag(str interface{}) {
	t := reflect.TypeOf(str).Elem()
	for i := 0; i < t.NumField(); i++ {
		tagInfo := t.Field(i).Tag.Get("info")
		tagDoc := t.Field(i).Tag.Get("doc")
		fmt.Println("info: ", tagInfo, " doc: ", tagDoc)
	}
}

func main() {
	var re resume
	findTag(&re)
}
```



应用实例：

```go
package main

import (
    "encoding/json"
    "fmt"
)
// 导包

type Movie struct {
    Title  string   `json:"title"`
    Year   int      `json:"year"`
    Actors []string `json:"actors"`
}
// 创建标签结构体

func main() {
    movie := Movie{"喜剧之王", 2000, []string{"xingye", "zhangbozhi"}}
    
    // 将结构体转换为JSON格式
    jsonStr, err := json.Marshal(movie)
    if err != nil {
        fmt.Println("json marshal error", err)
        return
    }
    // err用于判断转json是否成功
    
    fmt.Printf("JSON格式：%s\n", jsonStr)
    // 输出json格式
    
    
    // 解析json字符串 jsonStr ===> 结构体实例
    // jsonStr = {"title":"哪吒之魔童降世","year":2020,"rmb":10,"actors":["xinye","zhangbozhi"]}
    myMovie := Movie{}  // 创建对象
    err = json.Unmarshal(jsonStr, &myMovie)  // 反向转换
    if err != nil {
        fmt.Println("json unmarshal error ", err)
        return
    }
    // 失败判断
    
     fmt.Printf("结构体：%v\n", myMovie)
}
```



应用示例2

```go
package main

import (
    "fmt"
    "reflect"
)

type resume struct {
    Name string `json:"name" doc:"我的名字"`
}

func findDoc(stru interface{}) map[string]string {
    t := reflect.TypeOf(stru).Elem()
    doc := make(map[string]string)

    for i := 0; i < t.NumField(); i++ {
        doc[t.Field(i).Tag.Get("json")] = t.Field(i).Tag.Get("doc")
    }

    return doc

}

func main() {
    var stru resume
    doc := findDoc(&stru)
    fmt.Printf("name字段为：%s\n", doc["name"])
}
```

