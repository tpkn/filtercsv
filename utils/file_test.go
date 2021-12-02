package utils

import (
	"fmt"
	"testing"
)

func Test_FileHashify(t *testing.T) {
	hash, err := FileHashify("./../_/test_hashify.csv")
	if err != nil {
		t.Error(err)
	}
	
	fmt.Println("length:", len(hash))
}

func Test_FileExist(t *testing.T) {
	if FileExists("file.go") == false {
		t.Error("File does exists, but we got 'false' result!")
	}
	if FileExists("failed.kek") == true {
		t.Error("File does not exists, but we got 'true' result!")
	}
}

func Benchmark_FileHashify(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FileHashify("./../_/test_hashify.csv")
	}
}