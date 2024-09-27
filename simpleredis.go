package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type factory struct {
	Data    map[string]string `json:"data"`
	SetName map[string]bool   `json:"setName"`
}

func main() {
	var setname string
	setname = "storage.json"
	data, err := ioutil.ReadFile(setname)
	if err != nil {
		fmt.Println("读取错误")
		return
	}
	var store factory
	store.Data = make(map[string]string)
	store.SetName = make(map[string]bool)
	if err == nil {
		err = json.Unmarshal(data, &store)

	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("1.进入命令行模块进行使用")
		fmt.Println("2.使用说明")
		fmt.Println("3.退·出此简单的（）程序")
		fmt.Println("打出你的选择并回车：")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		switch input {

		case "1":
			fmt.Println("你进入了命令行模块，开始ba")
			for {
				fmt.Print(">> ")

				command, _ := reader.ReadString('\n')
				parts := strings.Fields(command)
				switch parts[0] {
				case "SET":
					if len(parts) != 3 {
						fmt.Println("输入格式有误·")
						continue
					}
					store.Data[parts[1]] = parts[2]
					fmt.Println("你已成功设置，试试使用GET吧.")

				case "GET":
					if len(parts) != 2 {
						fmt.Println("输入格式有误")
						continue
					}
					value, exists := store.Data[parts[1]]
					if exists {
						fmt.Println(value)
					} else {
						fmt.Println("找不到这个键值对")
					}

				case "DEL":
					if len(parts) != 2 {
						delete(store.Data, parts[1])
						fmt.Println("删除成功")
					}

				case "SETNX":
					if _, exists := store.Data[parts[1]]; exists {
						fmt.Println("0")
					} else {
						store.Data[parts[1]] = parts[2]
						fmt.Println("1")
					}

				case "SADD":
					if len(parts) == 3 {
						store.SetName[parts[1]] = true
						fmt.Println("导入成功")
					} else {
						fmt.Println("格式错误")
					}

				case "SMEMBER":
					if len(parts) != 2 {
						fmt.Println("输入格式有误")
						continue
					}

					if _, exists := store.SetName[parts[1]]; exists {
						for keys := range store.SetName {
							fmt.Println("集合中有：", keys)

						}
					}

				case "EXIT":
					fmt.Println("正在退出")
					return

				default:
					fmt.Println("未知的命令.....")

				}

			}
		case "2":
			fmt.Println("关于怎么使用这个拙略的工具")
			fmt.Println("首先你得在刚才的页面打出1")
			fmt.Println("然后你可以进行SET,GET,DEL,SETNX,EXIT命令")
			fmt.Println("SET是增加一个键值对，GET是获取一个键值对，DEL是删除一个键值对，SETNX是增加一个键值对，EXIT是退出程序，差不多就是这样")
			fmt.Println("SADD 则是增加一个集合中的元素，SMEMBER会获取并显示这个集合")
			fmt.Println("然后你可以输入你想要的命令，当然回车是必须的")
			fmt.Println("就到这里啦，菜菜带带")
			fmt.Println("输入EXIT就可以返回上一个界面啦")
			input, _ := reader.ReadString('\n')
			if input == "EXIT" {
				fmt.Println("Exiting...")
				return
			}
		case "3":
			fmt.Println("你已经退出了这个程序，再见")
			return
		}
	}
}
