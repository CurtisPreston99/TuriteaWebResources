package dataLevel

import (
	"encoding/json"
	"io"
)

type Resource struct {
	Type uint8 `json:"t"`
	Id int64 `json:"information"`
}

const (
	ArticleContent = iota
	Image
	Video
)

func JsonToResource(r io.Reader, num uint16) ([]Resource, error) {
	data := make([]Resource, num)
	d := json.NewDecoder(r)
	d.UseNumber()
	err := d.Decode(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func ResourceToJson(w io.Writer, res []Resource) error {
	return json.NewEncoder(w).Encode(res)
}
