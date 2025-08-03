package utils

import (
	"encoding/xml"
	"fmt"
	"log"
	"stockpull/model"
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
	RequestId string   `xml:"Matches>InstrumentMatch>Instrument>CommonName"`
	Matches   Matches  `xml:"Matches>InstrumentMatch>CompositeTrading"`
}
type Matches struct {
	XMLName       xml.Name `xml:"CompositeTrading"`
	Points        string   `xml:"Last>Price>Value"`
	ChangePercent string   `xml:"ChangePercent"`
	ChangePoint   string   `xml:"NetChange>Value"`
}

func ReadXML(input string) []model.StockIndex {
	var v MidLevel
	err := xml.Unmarshal([]byte(input), &v)
	if err != nil {
		log.Fatalln(err)
	}
	//fmt.Println(v.Indexe)
	var stockIndexArray []model.StockIndex
	for i := 0; i < len(v.Indexes); i++ {
		fmt.Println(v.Indexes[i].RequestId)
		fmt.Println(v.Indexes[i].Matches.Points)
		fmt.Println(v.Indexes[i].Matches.ChangePercent)
		stockIndex := model.StockIndex{StockIndexName: v.Indexes[i].RequestId, Points: v.Indexes[i].Matches.Points, ChangePercent: v.Indexes[i].Matches.ChangePercent, ChangePoint: v.Indexes[i].Matches.ChangePoint}
		stockIndexArray = append(stockIndexArray, stockIndex)
	}
	return stockIndexArray

}
