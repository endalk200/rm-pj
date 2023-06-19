package main

import (
	"flag"
	"fmt"
	"io/fs"
	"path/filepath"
)

func formatSize(size int64) string {
	const (
		_  = iota
		KB = 1 << (10 * iota)
		MB
		GB
	)

	switch {
	case size >= GB:
		return fmt.Sprintf("%.2f GB", float64(size)/float64(GB))
	case size >= MB:
		return fmt.Sprintf("%.2f MB", float64(size)/float64(MB))
	case size >= KB:
		return fmt.Sprintf("%.2f KB", float64(size)/float64(KB))
	default:
		return fmt.Sprintf("%d B", size)
	}
}

var TARGET_DIRECTORIES = []string{"node_modules", "dist"}

func main() {
	var path *string = flag.String("p", ".", "Path to scan") // -p <path>
	var dryRun *bool = flag.Bool("dryRun", false, "Dry run") // -dryRun

	flag.Parse()

	root := *path

	totalSize := int64(0)

	err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			for _, target := range TARGET_DIRECTORIES {
				if info.Name() == target {
					totalSize += info.Size()
				}
			}
		}

		return nil
	})

	if err != nil {
		panic(err)
	}

	if *dryRun {
		fmt.Printf("Total space of: %s will be reclaimed\n", formatSize(totalSize))
	} else {
		fmt.Printf("Total space of: %s is reclaimed\n", formatSize(totalSize))
	}
}
