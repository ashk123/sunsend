package main

import (
	"fmt"
	"log"
	"sunsend/internals/Base"
	"sunsend/internals/DB"
	"sunsend/internals/Data"
	"time"
)

func SubMonths(date string) string {
	t, err := time.Parse("2006/1/2", date)
	if err != nil {
		return err.Error()
	}
	return fmt.Sprintf("%d-%d-%d", t.Year(), t.Month()-1, t.Day())
}
func getMessagDate() (*[]Data.Message, int) {
	t := time.Now()
	ftime := fmt.Sprintf("%d/%d/%d", t.Year(), t.Month(), t.Day())
	//stime := SubMonths(ftime)
	//fmt.Println(ftime, stime)
	row, _ := DB.QueryRows(
		fmt.Sprintf(
			"SELECT * FROM Messages WHERE Date < '%s' AND Date > '%s';",
			ftime,
			t.AddDate(0, -3, 0),
		),
	)
	data, err := Base.Unmarshal(row)
	if err != nil {
		log.Println(err.Error())
		return nil, -1
	}
	// TODO: Replace the chcekCount function with a simple Scan
	//return checkCount(rows)
	return &data, 0
}

func main() {
	res, _ := getMessagDate()
	t := time.Now()
	time.Minute
	//timer, _ := time.Parse("2006-01-02", "2021-08-15")
	fmt.Println(res)
}
