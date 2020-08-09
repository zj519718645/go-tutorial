package ch9

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

type xmldas struct {
	XMLName  xml.Name       `xml:"das"`
	DataPort string         `xml:"DataPort,attr"`
	Desc     string         `xml:"desc,attr"`
	Src      xmlsource      `xml:"source"`
	Dest     xmldestination `xml:"destination"`
}

type xmlsource struct {
	Path  string `xml:"path,attr"`
	Param string `xml:"param,attr"`
}

type xmldestination struct {
	Path  string `xml:"path,attr"`
	Param string `xml:"param,attr"`
}

func TestXML(t *testing.T) {
	v := xmldas{DataPort: "8250", Desc: "123"}
	v.Src = xmlsource{Path: "123", Param: "456"}
	v.Dest = xmldestination{Path: "789", Param: "000"}
	output, err := xml.MarshalIndent(v, "", "   ")
	if err != nil {
		fmt.Println(err)
	}
	os.Stdout.Write([]byte(xml.Header))
	os.Stdout.Write(output)
}

type Result struct {
	Person []Person
}

type Person struct {
	Name      string `xml:"Name,attr"`
	Age       int    `xml:"Age,attr"`
	Career    string
	Interests Interests
}

type Interests struct {
	Interest []string
}

func TestXMLMarshal(t *testing.T) {
	content, err := ioutil.ReadFile("./studyGolang.xml")
	if err != nil {
		log.Fatal(err)
		return
	}
	//fmt.Println(content)
	var result Result
	err = xml.Unmarshal(content, &result)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(result)
}
