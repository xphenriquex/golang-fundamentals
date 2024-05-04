package main

import (
	"fmt"
	"io"
	"os"
)

func CopyFile(srcName, dstName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close() //correct!

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close() //correct!

	written, err = io.Copy(dst, src)
	// src.Close() //wrong!
	// dst.Close() //wrong!
	return
}

func copyFileMain() {
	written, err := CopyFile("myFile.txt", "copyFile.txt")
	fmt.Println(written, err)
}
