package Data

import (
	"time"
)

/*
	A very simple Structure for temping values
*/

type Temp struct {
	Size int
	Data map[string]interface{}
	Date string
}

var temp *Temp

func (p *Temp) Add(key string, data interface{}) {
	p.Data[key] = data // add a new value to the Temp values
}

func (p *Temp) Get(key string) interface{} {
	return p.Data[key]
}

func GetValue(key string) interface{} {
	temp := GetTemp()
	return temp.Get(key)
}

func (p *Temp) Clear() {
	for i := range p.Data {
		delete(p.Data, i)
	}
}

func GetTemp() *Temp {
	if temp == nil {
		temp = &Temp{
			Size: 0,
			Data: make(map[string]interface{}),
			Date: time.Now().String(),
		}
		return temp
	}
	return temp
}
