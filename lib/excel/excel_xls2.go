package excel

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

// ExampleExcelizeXlsx parse xlsx file
func ExampleExcelizeXlsx() ([][]string, error) {
	f, err := excelize.OpenFile("weather.xlsx")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	rows, err := f.GetRows("sheet1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return rows, nil
}
