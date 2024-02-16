package envelope

import (
	"encoding/xml"
	"fmt"
)

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:"-"`
	Gesmes  string   `xml:"-"`
	Xmlns   string   `xml:"-"`
	Subject string   `xml:"-"`
	Sender  struct {
		Text string `xml:"-"`
		Name string `xml:"-"`
	} `xml:"-"`
	Cube struct {
		Text string `xml:"-"`
		Cube struct {
			Text string `xml:"-"`
			Time string `xml:"-"`
			Cube []struct {
				Text     string  `xml:"-"`
				Currency string  `xml:"currency,attr"`
				Rate     float64 `xml:"rate,attr"`
			} `xml:"Cube"`
		} `xml:"Cube"`
	} `xml:"Cube"`
}

func (e Envelope) String() string {
	return fmt.Sprintf("Envelope is %v",
		e.Cube)
}
