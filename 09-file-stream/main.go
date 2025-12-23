package main

import (
	"fmt"
	"learning/filestream/buffer"
)

func main() {
	path := "./io/input.txt"

	fileObject := buffer.CreateFileBufferObject()
	fileObject.ReadFileBufferString(path)

	mapper := fileObject.Dict
	for key, val := range mapper {
		fmt.Printf("Termo: %s, Freq.: %d\n", key, val)
	}

	sizeMb := float64(fileObject.Sizebytes) / (1024 * 1024)
	fmt.Printf("Filesize MB: %.2f MiB\n", sizeMb)
}
