package main

import (
	"TuriteaWebResources/server/dataLevel"
	"encoding/xml"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Kml struct {
	F Folder `xml:"Folder"`
}
type Folder struct {
	Documents []Document `xml:"Document"`
}

type Document struct {
	P Placemark `xml:"Placemark"`
}

type Placemark struct {
	Points Point `xml:"Point"`
}

type Point struct {
	Coordinates string `xml:"coordinates"`
}

func main() {
	f, err := os.Open("data/kml/Turitea Pitfall Points.kml")
	if err != nil {
		panic(err)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	err = f.Close()
	if err != nil {
		panic(err)
	}
	var data = Kml{}
	err = xml.Unmarshal(b, &data)
	if err != nil {
		panic(err)
	}
	//fmt.Println(data)
	list := data.F.Documents
	for _, v := range list {
		point := strings.Split(v.P.Points.Coordinates,",")
		f1, err := strconv.ParseFloat(point[0], 64)
		if err != nil {
			panic(err)
		}
		f2, err := strconv.ParseFloat(point[1], 64)
		if err != nil {
			panic(err)
		}
		dataLevel.SQLNormal.CreatePin(0, 0, f2, f1, 18711, 0, "", "", "")
	}
}
