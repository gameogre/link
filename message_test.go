package link

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"encoding/xml"
	"github.com/funny/unitest"
	"testing"
)

type MyData struct {
	Id  int
	Msg string
}

func Test_Message_Raw(t *testing.T) {
	var data = []byte{1, 2, 3, 4, 5, 6}

	var buffer = newOutBuffer()
	err := Bytes(data)(buffer)
	unitest.NotError(t, err)

	unitest.Pass(t, bytes.Equal(data, buffer.Data))
}

func Test_Message_Json(t *testing.T) {
	var buffer = newOutBuffer()
	err := Json(MyData{Id: 1, Msg: "Test"})(buffer)
	unitest.NotError(t, err)

	var data MyData
	json.Unmarshal(buffer.Data, &data)

	unitest.Pass(t, data.Id == 1)
	unitest.Pass(t, data.Msg == "Test")
}

func Test_Message_Gob(t *testing.T) {
	var buffer = newOutBuffer()
	err := Gob(MyData{Id: 1, Msg: "Test"})(buffer)
	unitest.NotError(t, err)

	var data MyData
	gob.NewDecoder(bytes.NewReader(buffer.Data)).Decode(&data)

	unitest.Pass(t, data.Id == 1)
	unitest.Pass(t, data.Msg == "Test")
}

func Test_Message_Xml(t *testing.T) {
	var buffer = newOutBuffer()
	err := Xml(MyData{Id: 1, Msg: "Test"})(buffer)
	unitest.NotError(t, err)

	var data MyData
	xml.Unmarshal(buffer.Data, &data)

	unitest.Pass(t, data.Id == 1)
	unitest.Pass(t, data.Msg == "Test")
}
