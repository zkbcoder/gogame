package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"unicode"

	"github.com/tealeg/xlsx"
)

// 字符串首字母变大写
func FirstToUpper(str string) string {
	if len(str) == 0 {
		return ""
	}
	r := []rune(str)
	r[0] = unicode.ToUpper(r[0])
	str = string(r)
	return str
}

// 将字符串小写蛇体转驼峰 例如: user_level--> UserLevel
func StringToClass(str string) string {
	strs := strings.Split(str, "_")
	out := ""
	for _, v := range strs {
		out += FirstToUpper(v)
	}
	return out
}

// 根据字段名字获取字段类型
func GetFeildType(str string) string {
	if strings.HasPrefix(str, "i18n") {
		return "string"
	} else if strings.HasPrefix(str, "str_") {
		return "string"
	}
	return "int"
}

// 根据字段名字获取转换后的字段名
func GetFeildName(str string) string {
	if strings.HasPrefix(str, "i18n") {
		return str[len("i18n_"):]
	} else if strings.HasPrefix(str, "str_") {
		return str[len("str_"):]
	}
	return str
}

func createFactory() (*os.File, error) {
	// 输出文件
	file, err := os.Create("../tb/factory.go")
	if err != nil {
		fmt.Println("Error create factory file ", err)
		return file, err
	}

	file.Write([]byte("package tb\r\n\n")) // 文件开头
	file.Write([]byte("import \"github.com/tealeg/xlsx\"\r\n\n"))
	file.Write([]byte("type cfgIt interface {\n"))
	file.Write([]byte("    SetData(cell []*xlsx.Cell)\n"))
	file.Write([]byte("    GetKey() int\n"))
	file.Write([]byte("}\n"))

	file.Write([]byte("func Create(strName string) cfgIt {\r\n")) // 文件开头
	file.Write([]byte("    switch strName {"))                    // 文件开头

	return file, err
}

func main() {
	//	var hero tb.Hero
	//	fmt.Println(hero)
	fileFactory, _ := createFactory()
	defer fileFactory.Close()

	//	获取配置表目录下所有文件
	dir_list, e := ioutil.ReadDir("../excel")
	if e != nil {
		fmt.Println("read dir error")
		return
	}
	for i, v := range dir_list {
		filePath := fmt.Sprintf("../excel/%s", v.Name())
		//	打开文件
		xlFile, err := xlsx.OpenFile(filePath)
		if err != nil {
			fmt.Println(err)
			continue
		}

		codeFileName := strings.Replace(v.Name(), ".xlsx", ".go", -1)
		codeFilePath := fmt.Sprintf("../tb/%s", codeFileName)
		// 输出文件
		file, err := os.Create(codeFilePath)
		if err != nil {
			fmt.Println("Error open file ", codeFilePath, err)
			return
		}
		defer file.Close()
		file.Write([]byte("package tb\r\n\n")) // 文件开头
		file.Write([]byte("import \"github.com/tealeg/xlsx\"\t\n\n"))
		structName := strings.Split(v.Name(), ".")[0]
		fmt.Println(structName)
		strStruct := fmt.Sprintf("type %s struct {", StringToClass(structName))
		file.Write([]byte(strStruct))

		// 工厂内容添加
		strCase := fmt.Sprintf("\n    case \"%s\": return new(%s) ", structName, StringToClass(structName))
		fileFactory.Write([]byte(strCase))

		var content string
		var dataContent string
		var retKey string
		for _, sheet := range xlFile.Sheets {
			for _, row := range sheet.Rows {
				num := len(row.Cells)
				for i = 0; i < num; i++ {

					fieldName, _ := row.Cells[i].String()
					fieldType := GetFeildType(fieldName)
					fieldName = GetFeildName(fieldName)
					newName := StringToClass(fieldName)
					fmt.Println(newName)
					strLine := fmt.Sprintf("\n    %s %s `josn:\"%s\"`", newName, fieldType, fieldName)
					content = content + strLine
					if i == 0 {
						retKey = fmt.Sprintf("\n    return this.%s", newName)
					}
					strDataLine := fmt.Sprintf("\n    this.%s, _ = cell[%d].%s()", (newName), i, FirstToUpper(fieldType))
					dataContent = dataContent + strDataLine
				}
				break // 只读取第一行
			}

		}
		file.Write([]byte(content))
		file.Write([]byte("\n}"))
		file.Write([]byte("\n\n")) // new line
		strFunc := fmt.Sprintf("func (this *%s) SetData(cell []*xlsx.Cell) {", StringToClass(structName))
		file.Write([]byte(strFunc))
		//			func (this *Hero) setData(cell []*xlsx.Cell) {
		file.Write([]byte(dataContent))
		fmt.Println(i, "=", v.Name())
		fmt.Println(codeFilePath)
		file.Write([]byte("\n}"))
		file.Write([]byte("\n\n")) // new line
		funcKey := fmt.Sprintf("func (this *%s) GetKey() int {", StringToClass(structName))
		file.Write([]byte(funcKey))
		file.Write([]byte(retKey))

		file.Write([]byte("\n}"))
	}

	//	工厂文件结束
	fileFactory.Write([]byte("\n    }"))
	fileFactory.Write([]byte("\n    return nil"))
	fileFactory.Write([]byte("\n}"))

}
