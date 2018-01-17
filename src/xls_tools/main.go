package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"xls_tools/tb"

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
		if strings.HasPrefix(xlsfileName, "~") { // 去掉临时文件
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
			fmt.Println(strXlsxName, sheet.Name, sheet)
			fmt.Println(hero)

			fields := []string{}     // 字段名字
			fieldTypes := []string{} // 字段类型
			exStrs := []string{}     // 额外字符串[拼json用]

			for index, row := range sheet.Rows {
				js := ""
				if len(row.Cells) <= 0 {
					break
				}

				if 0 == index { // 第一行字段名
					for i := 0; i < len(row.Cells); i++ {
						fieldName, _ := row.Cells[i].String()
						fields = append(fields, fieldName)
					}
					fmt.Println(fields)
					fmt.Println("init field!!")
					continue
				}
				if 1 == index {
					fmt.Println("pass by log!!!") // 留给注释用
					continue
				}
				if 2 == index {
					for i := 0; i < len(row.Cells); i++ {
						t, _ := row.Cells[i].String()
						fieldTypes = append(fieldTypes, t)
					}
					fmt.Println(fieldTypes)
					fmt.Println("init fieldTypes!!")
					continue
				}
				if 3 == index {
					for i := 0; i < len(row.Cells); i++ {
						str, _ := row.Cells[i].String()
						exStrs = append(exStrs, str)
					}
					fmt.Println(exStrs)
					fmt.Println("init exStrs!!")
					continue
				}

				//	"1" : {"Id":1,"Name":"刀男","Type":"攻击者","Desc":"一个被诅咒的铁匠。他碰到的一切都会分崩离析。他梦想成为高贵的骑士，但是永远不可能。","Hp":"1K","Dmg":"100","Hit":"800","Dodge":"0","Critical":"300","CriticalResist":"0","CriticalDmg":"500","Skill1":0,"Skill2":0,"Skill3":0,"Skill4":0,"Skill5":0,"Skill6":0,"Skill7":0,"Skill8":0,"Skill9":0,"Skill10":0,"Skill11":0,"Resource":"r_1","RqInstanceStage":0},

				key, _ := row.Cells[0].Int() // 做为索引
				for i := 0; i < len(row.Cells); i++ {
					endFlag := ","
					if i+1 == len(row.Cells) {
						endFlag = ""
					}
					// 前缀和后缀补充
					exFront := "" // 额外向前补充
					exAfter := "" // 额外向后补充
					if strings.HasSuffix(exStrs[i], "}") {
						exAfter = exStrs[i] + ","
						endFlag = ""
					} else {
						exFront = exStrs[i]
					}

					if "int" == fieldTypes[i] {
						nVal, _ := row.Cells[i].Int()
						js = fmt.Sprintf("%s%s\"%s\":%d%s%s", js, exFront, fields[i], nVal, endFlag, exAfter)
					} else if "string" == fieldTypes[i] {
						strVal, _ := row.Cells[i].String()
						js = fmt.Sprintf("%s%s\"%s\":\"%s\"%s%s", js, exFront, fields[i], strVal, endFlag, exAfter)
					}
				}
				var content string
				if 4 == index {
					content = fmt.Sprintf("\n\"%d\" : {%s}", key, js)
				} else {
					content = fmt.Sprintf(",\n\"%d\" : {%s}", key, js)
				}
				fmt.Println(content)
				file.Write([]byte(content))
			}

			file.Write([]byte("\n}")) // json结尾
			break
		}
	}

}
