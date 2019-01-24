package main

import "os"

func main() {
	t()
	for _, rmdir := range rmdirs {
		rmdir()
	}
}

func tempDirs() []string {
	return []string{"ni", "wo/kkk", "ta"}
}

var rmdirs []func()

func t() {
	for _, d := range tempDirs() {
		dir := d
		os.MkdirAll(dir, 0755)
		rmdirs = append(rmdirs, func() {
			os.RemoveAll(dir)
		})
	}
}
