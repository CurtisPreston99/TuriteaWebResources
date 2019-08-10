package base

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestJsonToPins(t *testing.T) {
	f, err := os.Open("a_test.json")
	if err != nil {
		t.Fatal()
	}

	goal, err := JsonToPins(f, 2)
	if err != nil {
		t.Fatal()
	}
	p0 := goal[0]
	if p0.Uid != 1 || p0.Latitude != 10.1 || p0.Longitude != 10.2 || p0.Description != "abc" || p0.Time != 55 || p0.TagType != 1 {
		fmt.Println(p0)
		t.Fatal()
	}
	p1 := goal[1]
	if p1.Uid != 2 || p1.Latitude != 11.1 || p1.Longitude != 11.2 || p1.Description != "ab" || p1.Time != 56 || p1.TagType != 2 {
		t.Fatal()
	}
}

func TestPinsToJson(t *testing.T) {
	p1 := &Pin{Uid:1, Latitude:10.1, Longitude:10.2, Time:55, Description: "abc", TagType:1}
	p2 := &Pin{Uid:2, Latitude:11.1, Longitude:11.2, Time:56, Description: "ab", TagType:2}
	var list = []*Pin{p1, p2}
	buffer := bytes.NewBuffer(make([]byte, 254))
	err := PinsToJson(list, buffer)
	if err != nil {
		t.Fatal()
	}
	var goal = `[{"uid":1,"owner":0,"lat":10.1,"long":10.2,"time":55,"description":"abc","tag_type":1},{"uid":2,"owner":0,"lat":11.1,"long":11.2,"time":56,"description":"ab","tag_type":2}]`
	if goal == buffer.String() {
		fmt.Println(goal)
		fmt.Println(buffer.String())
		t.Fatal()
	}
}
