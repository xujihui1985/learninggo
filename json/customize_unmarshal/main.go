package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

// http://gregtrowbridge.com/golang-json-serialization-with-interfaces/
type ColoredThing interface {
	Color() string
}

type ColorfulEcosystem struct {
	Things []ColoredThing `json:"things"`
}

func (ce *ColorfulEcosystem) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	var rawMessages []*json.RawMessage
	err = json.Unmarshal(*objMap["things"], &rawMessages)
	if err != nil {
		return err
	}
	ce.Things = make([]ColoredThing, len(rawMessages))
	var m map[string]string

	for idx, rawMessage := range rawMessages {
		err = json.Unmarshal(*rawMessage, &m)
		if err != nil {
			return err
		}

		if m["type"] == "plant" {
			var p Plant
			err := json.Unmarshal(*rawMessage, &p)
			if err != nil {
				return err
			}
			ce.Things[idx] = &p
		} else if m["type"] == "animal" {
			var a Animal
			err := json.Unmarshal(*rawMessage, &a)
			if err != nil {
				return err
			}
			ce.Things[idx] = &a
		} else {
			return errors.New("unsupported type " + m["type"])
		}
	}
	return nil
}

type Plant struct {
	MyColor string `json:"color"`
}

type Animal struct {
	MyColor string `json:"color"`
}

func (p *Plant) Color() string {
	return p.MyColor
}

func (a *Plant) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]string{
		"type":  "plant",
		"color": a.Color(),
	})
}

func (a *Animal) Color() string {
	return a.MyColor
}

func (a *Animal) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]string{
		"type":  "animal",
		"color": a.Color(),
	})
}

func main() {
	fern := &Plant{MyColor: "green"}
	flower := &Plant{MyColor: "purple"}

	panther := &Animal{MyColor: "black"}
	lizard := &Animal{MyColor: "green"}

	// Then let's create a ColorfulEcosystem
	colorfulEcosystem := ColorfulEcosystem{
		Things: []ColoredThing{
			fern,
			flower,
			panther,
			lizard,
		},
	}

	byteSlice, _ := json.Marshal(colorfulEcosystem)
	fmt.Println(string(byteSlice))

	ce := ColorfulEcosystem{}

	err := json.Unmarshal(byteSlice, &ce)
	if err != nil {
		panic(err)
	}
	for _, c := range ce.Things {
		fmt.Println(c.Color())
	}
}
