package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"tb"

	"github.com/tealeg/xlsx"
)

func main() {

	//	获取配置表目录下所有文件
	dir_list, e := ioutil.ReadDir("excel")
	if e != nil {
		fmt.Println("read dir error")
		return
	}

	fmt.Println((dir_list))
	for _, v := range dir_list {
		xlsfileName := v.Name()
		if strings.HasPrefix(xlsfileName, ".") { // 去掉隐藏文件
			continue
		}
		// 遍历目录下所有文件
		filePath := fmt.Sprintf("excel/%s", v.Name())
		xlFile, err := xlsx.OpenFile(filePath)
		if err != nil {
			fmt.Println(err)
		}

		strXlsxName := v.Name()
		strXlsxName = strXlsxName[:len(strXlsxName)-5]
		fmt.Println(strXlsxName)
		for _, sheet := range xlFile.Sheets {
			//			if sheet.Name == "hero" {
			fileName := fmt.Sprintf("out/%s.json", strXlsxName)
			// 输出文件
			file, err := os.Create(fileName)
			fmt.Println(fileName)
			if err != nil {
				fmt.Println("Error open file ", fileName, err)
				return
			}
			defer file.Close()

			file.Write([]byte("{")) // json开头
			hero := tb.Create(strXlsxName)
			fmt.Println(sheet.Name, sheet)
			for index, row := range sheet.Rows {
				if 0 == index {
					continue
				}
				if len(row.Cells) <= 0 {
					break
				}
				hero.SetData(row.Cells)
				b, err := json.Marshal(hero)
				if err != nil {
					fmt.Println("encoding faild")
				} else {
					var content string
					if 1 == index {
						content = fmt.Sprintf("\n\"%d\" : %s", hero.GetKey(), string(b))
					} else {
						content = fmt.Sprintf(",\n\"%d\" : %s", hero.GetKey(), string(b))
					}
					file.Write([]byte(content))
					fmt.Println(string(b))
				}

			}

			file.Write([]byte("\n}")) // json结尾
			//			}
			//			else {
			//				for _, row := range sheet.Rows {
			//					rowNum := len(row.Cells)
			//					for i := 0; i < rowNum; i++ {
			//						fmt.Println(row.Cells[i].String())
			//					}
			//				}
			//			}
			break
		}
	}

	//	fmt.Println(xlFile)
}
