package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	student := Student{
		Name: "张三",
		Age:  20,
	}
	// 把student转换为json
	jsonData, err := json.Marshal(student)
	if err != nil {
		fmt.Println("JSON转换失败", err)
	}
	// 把jsonData写入一个文件里
	os.WriteFile("student.json", jsonData, 0644)

	// 从文件里读取json数据
	jsonData2, err := os.ReadFile("student.json")
	if err != nil {
		fmt.Println("读取文件失败", err)
	}
	// 把jsonData2转换为student

	err = json.Unmarshal(jsonData2, &student)
	if err != nil {
		fmt.Println("JSON转换失败", err)
	}
	fmt.Println(student)

	test_arr := []Student{
		{Name: "张三", Age: 20},
		{Name: "李四", Age: 21},
	}

	test_arr[0].Name = "王五"

	fmt.Println(test_arr)

}
