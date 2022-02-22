package test

import (
	gogo "github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/proto"
	"gogo_pt_test/gogofile"
	"gogo_pt_test/ptfile"
	"testing"
)

// Marshal

func BenchmarkGoGoFile_GoGoMarshal(b *testing.B) {
	pt := gogofile.ProtoMsg{Str1: "str1",Str2: "str2",I3: 3,I4: 4}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gogo.Marshal(&pt)
	}
}

func BenchmarkPtFile_ProtobufMarshal(b *testing.B) {
	pt := ptfile.ProtoMsg{Str1: "str1",Str2: "str2",I3: 3,I4: 4}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		proto.Marshal(&pt)
	}
}

func BenchmarkGoGoFile_ProtobufMarshal(b *testing.B) {
	pt := gogofile.ProtoMsg{Str1: "str1",Str2: "str2",I3: 3,I4: 4}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		proto.Marshal(&pt)
	}
}

func BenchmarkPtFile_GoGoMarshal(b *testing.B) {
	pt := ptfile.ProtoMsg{Str1: "str1",Str2: "str2",I3: 3,I4: 4}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gogo.Marshal(&pt)
	}
}



// Unmarshal
//
//func BenchmarkGoGoFile_GoGoFile_GoGoUnmarshal(b *testing.B) {
//	pt := gogofile.ProtoMsg{Str1: "str1",Str2: "str2",I3: 3,I4: 4}
//	pt2 := gogofile.ProtoMsg{}
//	bytes,_ := gogo.Marshal(&pt)
//
//	b.ReportAllocs()
//	b.ResetTimer()
//
//	for i := 0; i < b.N; i++ {
//		gogo.Unmarshal(bytes,&pt2)
//	}
//}
//func BenchmarkGoGoFile_PtFile_GoGoUnmarshal(b *testing.B) {
//	pt := gogofile.ProtoMsg{Str1: "str1",Str2: "str2",I3: 3,I4: 4}
//	pt2 := ptfile.ProtoMsg{}
//	bytes,_ := gogo.Marshal(&pt)
//
//	b.ReportAllocs()
//	b.ResetTimer()
//
//	for i := 0; i < b.N; i++ {
//		gogo.Unmarshal(bytes,&pt2)
//	}
//}
//
//func BenchmarkPtFile_GoGoFile_ProtobufUnmarshal(b *testing.B) {
//	pt := ptfile.ProtoMsg{Str1: "str1",Str2: "str2",I3: 3,I4: 4}
//	pt2 := gogofile.ProtoMsg{}
//	bytes,_ := proto.Marshal(&pt)
//
//	b.ReportAllocs()
//	b.ResetTimer()
//
//	for i := 0; i < b.N; i++ {
//		proto.Unmarshal(bytes,&pt2)
//	}
//}
//func BenchmarkPtFile_PtFile_ProtobufUnmarshal(b *testing.B) {
//	pt := ptfile.ProtoMsg{Str1: "str1",Str2: "str2",I3: 3,I4: 4}
//	pt2 := ptfile.ProtoMsg{}
//	bytes,_ := proto.Marshal(&pt)
//
//	b.ReportAllocs()
//	b.ResetTimer()
//
//	for i := 0; i < b.N; i++ {
//		proto.Unmarshal(bytes,&pt2)
//	}
//}
//
//func BenchmarkGoGoFile_GoGoFile_ProtobufUnmarshal(b *testing.B) {
//	pt := gogofile.ProtoMsg{Str1: "str1",Str2: "str2",I3: 3,I4: 4}
//	pt2 := gogofile.ProtoMsg{}
//	bytes,_ := gogo.Marshal(&pt)
//	b.ReportAllocs()
//	b.ResetTimer()
//
//	for i := 0; i < b.N; i++ {
//		proto.Unmarshal(bytes,&pt2)
//	}
//}
//func BenchmarkGoGoFile_PtFile_ProtobufUnmarshal(b *testing.B) {
//	pt := gogofile.ProtoMsg{Str1: "str1",Str2: "str2",I3: 3,I4: 4}
//	pt2 := ptfile.ProtoMsg{}
//	bytes,_ := gogo.Marshal(&pt)
//	b.ReportAllocs()
//	b.ResetTimer()
//
//	for i := 0; i < b.N; i++ {
//		proto.Unmarshal(bytes,&pt2)
//	}
//}
//
//func BenchmarkPtFile_GoGoFile_GoGoUnmarshal(b *testing.B) {
//	pt := ptfile.ProtoMsg{Str1: "str1",Str2: "str2",I3: 3,I4: 4}
//	pt2 := gogofile.ProtoMsg{}
//	bytes,_ := proto.Marshal(&pt)
//	b.ReportAllocs()
//	b.ResetTimer()
//
//	for i := 0; i < b.N; i++ {
//		gogo.Unmarshal(bytes,&pt2)
//	}
//}
//
//func BenchmarkPtFile_PtFile_GoGoUnmarshal(b *testing.B) {
//	pt := ptfile.ProtoMsg{Str1: "str1",Str2: "str2",I3: 3,I4: 4}
//	pt2 := ptfile.ProtoMsg{}
//	bytes,_ := proto.Marshal(&pt)
//	b.ReportAllocs()
//	b.ResetTimer()
//
//	for i := 0; i < b.N; i++ {
//		gogo.Unmarshal(bytes,&pt2)
//	}
//}
