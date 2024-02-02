package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	cmd := os.Args[1]
	switch cmd {
	case "genTree":
		// go run ./ genTree tmp_30_20000 30 20000

		// go run ./ genTree tmp_20_100 20 100
		GenTree(os.Args[2], mustInt(os.Args[3]), mustInt(os.Args[4]))
	}
}

func mustInt(s string) int {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(fmt.Errorf("invalid int: %s %v", s, err))
	}
	return int(i)
}

func GenTree(root string, maxDepth int, totalDirs int) error {
	i := 0
	for {
		if i >= totalDirs {
			break
		}
		var paths []string = []string{fmt.Sprintf("dir_%d", i)}
		err := os.MkdirAll(filepath.Join(root, filepath.Join(paths...)), 0755)
		if err != nil {
			return err
		}
		i++

		if i >= totalDirs {
			break
		}

		for j := 0; j < maxDepth-1; j++ {
			paths = append(paths, fmt.Sprintf("sub_%d", j))
			dir := filepath.Join(root, filepath.Join(paths...))
			err := os.MkdirAll(dir, 0755)
			if err != nil {
				return err
			}
			i++
		}
	}
	return nil
}
