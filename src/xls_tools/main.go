package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/tealeg/xlsx"
)

type Hero struct {
	Id             int    `json:"id"`
	Name           string `json:"i18n_name"`
	Type           int    `json:"type"`
	Desc           string `json:"desc"`
	Hp             int    `json:"hp"`
	Dmg            int    `json:"dmg"`
	Hit            int    `json:"hit"`
	Dodge          int    `json:"dodge"`
	Critical       int    `json:"critical"`
	CriticalResist int    `json:"critical_resist"`
	CriticalDmg    int    `json:"critical_dmg"`
	Skill1         int    `json:"skill_1"`
	Skill2         int    `json:"skill_2"`
	Skill3         int    `json:"skill_3"`
	Skill4         int    `json:"skill_4"`
	Skill5         int    `json:"skill_5"`
	Skill6         int    `json:"skill_6"`
	Skill7         int    `json:"skill_7"`
	Skill8         int    `json:"skill_8"`
	Skill9         int    `json:"skill_9"`
	Skill10        int    `json:"skill_10"`
	Skill11        int    `json:"skill_11"`
	Resource       int    `json:"resource"`
}

func main() {
	excelFileName := "hero.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Println(err)
	}

	for _, sheet := range xlFile.Sheets {
		if sheet.Name == "hero" {
			fileName := fmt.Sprintf("%s.json", sheet.Name)
			// 输出文件
			file, err := os.Create(fileName)
			if err != nil {
				fmt.Println("Error open file ", fileName, err)
				return
			}
			defer file.Close()

			file.Write([]byte("{")) // json开头

			var hero Hero
			fmt.Println(sheet.Name, sheet)
			for index, row := range sheet.Rows {
				if 0 == index {
					continue
				}
				if len(row.Cells) <= 0 {
					break
				}
				hero.Id, _ = row.Cells[0].Int()
				hero.Name, _ = row.Cells[1].String()
				hero.Type, _ = row.Cells[2].Int()
				hero.Desc, _ = row.Cells[3].String()
				hero.Hp, _ = row.Cells[4].Int()
				hero.Dmg, _ = row.Cells[5].Int()
				hero.Hit, _ = row.Cells[6].Int()
				hero.Dodge, _ = row.Cells[7].Int()
				hero.Critical, _ = row.Cells[8].Int()
				hero.CriticalResist, _ = row.Cells[9].Int()
				hero.CriticalDmg, _ = row.Cells[10].Int()
				hero.Skill1, _ = row.Cells[11].Int()
				hero.Skill2, _ = row.Cells[12].Int()
				hero.Skill3, _ = row.Cells[13].Int()
				hero.Skill4, _ = row.Cells[14].Int()
				hero.Skill5, _ = row.Cells[15].Int()
				hero.Skill6, _ = row.Cells[16].Int()
				hero.Skill7, _ = row.Cells[17].Int()
				hero.Skill8, _ = row.Cells[18].Int()
				hero.Skill9, _ = row.Cells[19].Int()
				hero.Skill10, _ = row.Cells[20].Int()
				hero.Skill11, _ = row.Cells[21].Int()
				hero.Resource, _ = row.Cells[22].Int()

				b, err := json.Marshal(hero)
				if err != nil {
					fmt.Println("encoding faild")
				} else {
					var content string
					if 1 == index {
						content = fmt.Sprintf("\n\"%d\" : %s", hero.Id, string(b))
					} else {
						content = fmt.Sprintf(",\n\"%d\" : %s", hero.Id, string(b))
					}
					file.Write([]byte(content))
				}

			}

			file.Write([]byte("\n}")) // json结尾
		} else {
			for _, row := range sheet.Rows {
				rowNum := len(row.Cells)
				for i := 0; i < rowNum; i++ {
					fmt.Println(row.Cells[i].String())
				}
			}
		}

	}

	fmt.Println(xlFile)
}
