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
		t.Fatal(err)
	}
	p0 := goal[0]
	fmt.Println(p0)
	if p0.Uid != 1 || p0.Latitude != 10.1 || p0.Longitude != 10.2 || p0.Description != "abc" || p0.Time != 55 || p0.TagType != "airport" {
		t.Fatal()
	}
	p1 := goal[1]
	fmt.Println(p1)
	if p1.Uid != 2 || p1.Latitude != 11.1 || p1.Longitude != 11.2 || p1.Description != "ab" || p1.Time != 56 || p1.TagType != "alcohol-shop" {
		t.Fatal()
	}
}

func TestPinsToJson(t *testing.T) {
	p1 := &Pin{Uid:1, Latitude:10.1, Longitude:10.2, Time:55, Description: "abc", TagType:tagMap[1]}
	p2 := &Pin{Uid:2, Latitude:11.1, Longitude:11.2, Time:56, Description: "ab", TagType:tagMap[2]}
	var list = []*Pin{p1, p2}
	buffer := bytes.NewBuffer(make([]byte, 254))
	err := PinsToJson(list, buffer)
	if err != nil {
		t.Fatal()
	}
	var goal = `[{"uid":1,"owner":0,"lat":10.1,"lon":10.2,"time":55,"description":"abc","tag_type":"airport"},{"uid":2,"owner":0,"lat":11.1,"lon":11.2,"time":56,"description":"ab","tag_type":"alcohol-shop"}]`
	if goal == buffer.String() {
		fmt.Println(goal)
		fmt.Println(buffer.String())
		t.Fatal()
	}
}

func TestArticlesToJson(t *testing.T) {
	buffer := make([]byte, 0, 200)
	buf := bytes.NewBuffer(buffer)
	a := Article{Id:1, WriteBy:0, Summary:"abc"}
	b := Article{Id:2, WriteBy:1, Summary:"test"}
	err := ArticlesToJson([]*Article{&a, &b}, buf)
	if err != nil {
		t.Fatal(err)
	}
	//buf.Reset()
	as, err := JsonToArticles(buf, 2)
	if err != nil {
		t.Fatal(err)
	}
	if a != *as[0] {
		t.Fatal()
	}
	if b != *as[1] {
		t.Fatal()
	}
}