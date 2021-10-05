package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
)

type registerInfo struct {
	Arrival   float32 `json:"arrival time"`
	Departure float32 `json:"departure time"`
}

func newRegisterInfo(arrival, departure float32) *registerInfo {
	return &registerInfo{
		Arrival:   arrival,
		Departure: departure,
	}
}

func register(arrival, departure float32) error {
	session, _ := mgo.Dial("127.0.0.1:27017")
	db := session.DB("test")
	c := db.C("register")
	err := c.Insert(newRegisterInfo(arrival, departure))
	if err != nil {
		return err
	}
	return nil
}

func count(start, end float32) int {
	var sum int
	session, _ := mgo.Dial("127.0.0.1:27017")
	db := session.DB("test")
	c := db.C("register")
	var registerInfos []registerInfo
	err := c.Find(nil).All(&registerInfos)
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < len(registerInfos); i++ {
		if start <= registerInfos[i].Departure && end >= registerInfos[i].Arrival {
			sum++
		}
	}
	fmt.Println(sum)
	return sum
}

func main() {
	register(1, 3)
	register(1, 1.5)
	register(1, 4)
	register(1, 2)
	count(2, 3)
}
