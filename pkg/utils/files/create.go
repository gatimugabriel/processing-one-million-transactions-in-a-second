package files

import (
	"os"
)

func CreateFile(filename string) (*os.File, error) {
	file, err := os.Create(filename)
	if err != nil {
		return nil, err
	}
	//defer func(file *os.File) {
	//	err := file.Close()
	//	if err != nil {
	//		fmt.Println("Error closing file:", err)
	//	}
	//}(file)

	return file, nil
}
