package main

import (
	"encoding/json"
	"fmt"

	"github.com/roblesdotdev/movies-ms/gen"
	"github.com/roblesdotdev/movies-ms/metadata/pkg/model"
	"google.golang.org/protobuf/proto"
)

var metadata = &model.Metadata{
	ID:          "123",
	Title:       "The movie 2",
	Description: "Sequel of the legendary title",
	Director:    "Foo Bar",
}

var genMetadata = &gen.Metadata{
	Id:          "123",
	Title:       "The movie 2",
	Description: "Sequel of the legendary title",
	Director:    "Foo Bar",
}

func main() {
	jsonBytes, err := serializeToJson(metadata)
	if err != nil {
		panic(err)
	}

	protoBytes, err := serializeToProto(genMetadata)
	if err != nil {
		panic(err)
	}

	fmt.Printf("JSON size:\t%dB\n", len(jsonBytes))
	fmt.Printf("Proto size:\t%dB\n", len(protoBytes))
}

func serializeToJson(m *model.Metadata) ([]byte, error) {
	return json.Marshal(m)
}

func serializeToProto(m *gen.Metadata) ([]byte, error) {
	return proto.Marshal(m)
}
