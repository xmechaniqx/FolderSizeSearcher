// path := "C:/workspace/src/helloworld"
package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	var Arr []string
	target := make(map[string]float64)
	dir := "N:/"
	files, _ := os.ReadDir(dir)
	path, _ := filepath.Abs(dir)
	filepath.Abs(dir)

	for _, file := range files {
		// r := file.Name()
		Arr = append(Arr, filepath.Join(path, file.Name()))
		fmt.Println(filepath.Join(path, file.Name()))
		// names = append(names, filepath.Join(file.Name()))
		// fmt.Println(Arr)
		// fmt.Println(filepath.Join("путь", path, "имя", file.Name()), Arr)
	}
	// fmt.Println(Arr)
	for _, v := range Arr {
		// DirSize(v)
		fmt.Println(v)
		if DirSizeMB(v) >= 250 {
			target[v] = DirSizeMB(v)
		}
		// fmt.Printf("%10.2f\n %10s", DirSizeMB(v), names[k])
	}
	// for k, v := range target {

	// 	fmt.Printf("%.f %s\n", v, k)

	// }
	write(target)
	// fmt.Println(target)

}
func DirSizeMB(path string) float64 {
	var dirSize int64 = 0

	readSize := func(path string, file os.FileInfo, err error) error {
		if !file.IsDir() {
			dirSize += file.Size()
			fmt.Println(path, file.Size())
		}

		return nil
	}

	filepath.Walk(path, readSize)

	sizeMB := float64(dirSize) / 1024.0 / 1024.0

	return sizeMB
}
func write(lines map[string]float64) {
	// create file
	f, err := os.Create("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file
	defer f.Close()

	for str, line := range lines {
		// newstr := strings.TrimSuffix(str, "C:\workspace\src")
		res := fmt.Sprintf("%.f %s", line, str)
		fmt.Println(res)
		_, err := f.WriteString(res + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
}
