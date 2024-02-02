package main

import (
	"io/fs"
	"path/filepath"
	"testing"
)

// go test -run TestTraverseFileSerially -v ./main_test.go
func TestTraverseFileSerially(t *testing.T) {
	// expectCnt := 10000
	// expectMaxDepth := 10

	stopCnt := 5000
	err := traverseWithMax(t, stopCnt)
	if err != nil {
		panic(err)
	}

	// if cnt != expectCnt {
	// 	t.Fatalf("expect count: %d, actual: %d", expectCnt, cnt)
	// }
	// if maxDepth != expectMaxDepth {
	// 	t.Fatalf("expect max depth: %d, actual: %d", expectMaxDepth, maxDepth)
	// }
}

// 449.011404 ms
func traverse(t testing.TB, dir string) error {
	return filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		return nil
	})
}

func traverseWithMax(t testing.TB, stopCnt int) error {
	cnt := 0
	// start := time.Now()
	// defer func() {
	// 	t.Logf("max: %d, cnt: %d, cost: %v", stopCnt, cnt, time.Since(start))
	// }()
	return filepath.WalkDir("tmp", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if cnt >= stopCnt {
			return filepath.SkipDir
		}
		cnt++
		return nil
	})
}

// go test -run TestTraverseFileConcurently -v ./main_test.go
func TestTraverseFileConcurently(t *testing.T) {

	// expectCnt := 10000
	// cnt := 0
	// maxDepth := 0
	//
	//	err := filepath.WalkDir("tmp", func(path string, d fs.DirEntry, err error) error {
	//		if err != nil {
	//			return err
	//		}
	//		cnt++
	//		return nil
	//	})
	//
	//	if err != nil {
	//		panic(err)
	//	}
}
