package utils

import (
	"encoding/xml"
	"fmt"
	"log"
)

type TopLevel struct {
	XMLName  xml.Name `xml:"GetInstrumentByDialectResponse"`
	MidLevel xml.Name `xml:"InstrumentResponses"`
}

type MidLevel struct {
	XMLName xml.Name  `xml:"GetInstrumentByDialectResponse"`
	Indexes []Indexes `xml:"InstrumentResponses>InstrumentByDialectResponse"`
}

type Indexes struct {
	XMLName   xml.Name `xml:"InstrumentByDialectResponse"`
	RequestId string   `xml:"RequestId"`
}
type Result struct {
	Value string
}

func ReadXML(input string) {
	var v MidLevel
	err := xml.Unmarshal([]byte(input), &v)
	if err != nil {
		log.Fatalln(err)
	}
	//fmt.Println(v.Indexe)
	for i := 0; i < len(v.Indexes); i++ {
		fmt.Println(v.Indexes[i].RequestId)
	}

}
