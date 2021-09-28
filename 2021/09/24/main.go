package main

import (
	"fmt"
	"os"
)

func main() {
	start := 7000
	end := start + 7000
	f, err := os.Create("test1.json")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer f.Close()

	_, err = f.WriteString("[")

	for i := start; i < end; i++ {
		_, err = f.WriteString(fmt.Sprintf(`{
			"fields": {
				"id": %d,
				"doc_id": "苏宁易购%d",
				"text": "阿里云盘「笔记」是你在数字生活中的又一个伙伴，帮助你随时记录生活、学习、工作中的各种重要信息。"
			},
			"cmd": "ADD"
		}`, i, i))
		if i != end {
			_, err = f.WriteString(",")
		}
	}

	_, err = f.WriteString("]")
	f.Sync()
}
