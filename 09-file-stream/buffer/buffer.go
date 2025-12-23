package buffer

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type FileBuffer struct {
	Dict      map[string]uint
	Sizebytes int64
}

func CreateFileBufferObject() *FileBuffer {
	return &FileBuffer{}
}

func (f *FileBuffer) ReadFileBufferString(filepath string) *FileBuffer {
	file, err := os.Open(filepath)
	if err != nil {
		ErrMsg := fmt.Sprintf("Erro ao abrir arquivo: %s\n", filepath)
		log.Fatal(ErrMsg)
		log.Fatal("Error: ", err)
	}
	defer file.Close()

	data, err := file.Stat()
	if err != nil {
		ErrMsg := fmt.Sprintf("Erro ao buscar dados sobre o arquivo: %s\n", filepath)
		log.Fatal(ErrMsg)
		log.Fatal("Error: ", err)
	}

	f.Sizebytes = data.Size()

	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanLines)

	freq := make(map[string]uint)
	for scan.Scan() {
		slice := strings.Split(scan.Text(), " ")
		for _, val := range slice {
			w := strings.ToLower(val)
			freq[w]++
		}
	}

	f.Dict = freq
	return f
}
