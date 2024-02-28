package main

import "testing"

func BenchmarkSerializeToJson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		serializeToJson(metadata)
	}
}

func BenchmarkSerializeToProto(b *testing.B) {
	for i := 0; i < b.N; i++ {
		serializeToProto(genMetadata)
	}
}
