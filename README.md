# protobufVSgogo
protobuf VS gogo benchmark

## Results
2022-02-22 Results with Go 1.17 darwin/arm64 macbook pro m1 8

benchmark                                | iter       | time/iter    | bytes/op  | allocs/op  
-----------------------------------------|------------|--------------|-----------|------------
BenchmarkGoGoFile_GoGoMarshal-8          |   25427499 |  44.76 ns/op |        16 |         1  
BenchmarkPtFile_ProtobufMarshal-8        |    7377848 |  162.4 ns/op |        16 |         1 
BenchmarkGoGoFile_ProtobufMarshal-8      |    5138227 |  234.4 ns/op |        32 |         2 
BenchmarkPtFile_GoGoMarshal-8            |    6834567 |  173.9 ns/op |        48 |         2 