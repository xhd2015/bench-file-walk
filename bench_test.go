package main

import "testing"

// prepare the benchmark:
//   for i in 100 250 2500 5000 10000; do go run ./ genTree tmp_20_$i 20 $i;done

// go test -bench=BenchmarkWalkSerial1_ -benchtime=10s -run=NONE -v ./
// 156262917 ns/op
func BenchmarkWalkSerial1_(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := traverseWithMax(b, 1)
		if err != nil {
			b.Fatalf("%v", err)
		}
	}
}

// go test -bench=BenchmarkWalkSerial100_ -benchtime=10s -run=NONE -v ./
// 4733646 ns/op = 4.73ms
func BenchmarkWalkSerial100_(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := traverse(b, "tmp_20_100")
		if err != nil {
			b.Fatalf("%v", err)
		}
	}
}

// go test -bench=BenchmarkWalkSerial250_ -benchtime=10s -run=NONE -v ./
// 12904282 ns/op = 12.90ms
func BenchmarkWalkSerial250_(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := traverse(b, "tmp_20_250")
		if err != nil {
			b.Fatalf("%v", err)
		}
	}
}

// go test -bench=BenchmarkWalkSerial2500 -benchtime=10s -run=NONE -v ./
// 123100811 ns/op = 123.10ms
func BenchmarkWalkSerial2500(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := traverse(b, "tmp_20_2500")
		if err != nil {
			b.Fatalf("%v", err)
		}
	}
}

// go test -bench=BenchmarkWalkSerial5000 -benchtime=10s -run=NONE -v ./
// 251533266 ns/op = 251.53ms
func BenchmarkWalkSerial5000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := traverse(b, "tmp_20_5000")
		if err != nil {
			b.Fatalf("%v", err)
		}
	}
}

// go test -bench=BenchmarkWalkSerial10000 -benchtime=10s -run=NONE -v ./
// 544532792 ns/op = 544.53ms
func BenchmarkWalkSerial10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := traverse(b, "tmp_20_10000")
		if err != nil {
			b.Fatalf("%v", err)
		}
	}
}
