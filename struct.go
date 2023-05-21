package main

import "encoding/xml"

type ComicInfoStruct struct {
	XMLName            xml.Name `xml:"schema"`
	Text               string   `xml:",chardata"`
	ElementFormDefault string   `xml:"elementFormDefault,attr"`
	Xs                 string   `xml:"xs,attr"`
	Element            struct {
		Text     string `xml:",chardata"`
		Name     string `xml:"name,attr"`
		Nillable string `xml:"nillable,attr"`
		Type     string `xml:"type,attr"`
	} `xml:"element"`
	ComplexType []struct {
		Text     string `xml:",chardata"`
		Name     string `xml:"name,attr"`
		Sequence struct {
			Text    string `xml:",chardata"`
			Element []struct {
				Text      string `xml:",chardata"`
				MinOccurs string `xml:"minOccurs,attr"`
				MaxOccurs string `xml:"maxOccurs,attr"`
				Default   string `xml:"default,attr"`
				Name      string `xml:"name,attr"`
				Type      string `xml:"type,attr"`
				Nillable  string `xml:"nillable,attr"`
			} `xml:"element"`
		} `xml:"sequence"`
		Attribute []struct {
			Text    string `xml:",chardata"`
			Name    string `xml:"name,attr"`
			Type    string `xml:"type,attr"`
			Use     string `xml:"use,attr"`
			Default string `xml:"default,attr"`
		} `xml:"attribute"`
	} `xml:"complexType"`
	SimpleType []struct {
		Text        string `xml:",chardata"`
		Name        string `xml:"name,attr"`
		Restriction struct {
			Text        string `xml:",chardata"`
			Base        string `xml:"base,attr"`
			Enumeration []struct {
				Text  string `xml:",chardata"`
				Value string `xml:"value,attr"`
			} `xml:"enumeration"`
			MinInclusive struct {
				Text  string `xml:",chardata"`
				Value string `xml:"value,attr"`
			} `xml:"minInclusive"`
			MaxInclusive struct {
				Text  string `xml:",chardata"`
				Value string `xml:"value,attr"`
			} `xml:"maxInclusive"`
			FractionDigits struct {
				Text  string `xml:",chardata"`
				Value string `xml:"value,attr"`
			} `xml:"fractionDigits"`
		} `xml:"restriction"`
		List struct {
			Text       string `xml:",chardata"`
			SimpleType struct {
				Text        string `xml:",chardata"`
				Restriction struct {
					Text        string `xml:",chardata"`
					Base        string `xml:"base,attr"`
					Enumeration []struct {
						Text  string `xml:",chardata"`
						Value string `xml:"value,attr"`
					} `xml:"enumeration"`
				} `xml:"restriction"`
			} `xml:"simpleType"`
		} `xml:"list"`
	} `xml:"simpleType"`
}

type VariablesStruct struct {
	Id int `json:"id"`
}

type PostBodyStruct struct {
	Query     string          `json:"query"`
	Variables VariablesStruct `json:"variables"`
}

type DetailsStruct struct {
	Title        string   `json:"title"`
	Author       string   `json:"author"`
	Artist       string   `json:"artist"`
	Description  string   `json:"description"`
	Genre        []string `json:"genre"`
	Status       string   `json:"status"`
	StatusValues []string `json:"_status values"`
	AnilistID    int      `json:"_anilist_id"`
}
