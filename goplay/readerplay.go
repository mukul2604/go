package main

import (
	"fmt" // formatting package
	"os" //operating system package, design is unix-like and error handling is go like
	"encoding/csv" // for comma separated value files reading
	//"log" //logging package
	//"io"
)

func main() {
	// package os open. same for all files. returns *File  associated descriptor
	// it is in only read mode use OpenFile for create/write and other ops
	f, err := os.Open("/Users/mukul/work/src/goplay/a.txt")
	if err != nil {
		panic(err)
	}
	defer  f.Close()
	//rdr := io.Reader(f)
	name := f.Name()
	fmt.Println(name)
	//csv reader
	csvrdr := csv.NewReader(f)
	rows, err := csvrdr.ReadAll()

	if err != nil {
		panic(err)
	}

	//printing each row
	for _, row := range rows {
		fmt.Println(row)
	}

	//print first element of each row
	for _, row := range rows {
		fmt.Println(row[0])
	}
}
