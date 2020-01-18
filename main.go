package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type income struct {
	id     int64
	amount uint64
}

var data = []income{
	{1, 1_000_000},
	{2, 2_000_000},
	{3, 3_000_000},
	{4, 4_000_000},
}

var customError = errors.New("custom error")

func makeError() error {
	return fmt.Errorf("error generated: %w", customError)
}

var dataBytes = [] byte{}

////////////////////////////CONVERT-TO-WRITE
func dataStructToBytes(data []income) (dataBytes []byte) {
	log.Print("conversion this struct data is started")
	log.Println("\n", data)
	dataStrings := make([]string, len(data))
	for i, datum := range data {
		dataStrings[i] = fmt.Sprintf("id:%d;amount:%d", datum.id, datum.amount)
	}
	oneString := strings.Join(dataStrings, "\n")
	dataBytes = []byte(oneString)
	log.Print("data converted to slice[] byte")
	return
}

////////////////////////////WRITE-TO-FILE
func Write(data []byte, path string) string {
	log.Print("data saving started")
	err := ioutil.WriteFile(path, data, 0666)
	if err != nil {
		log.Fatal(fmt.Errorf("can't save to %s: %w", path, err))
	}
	log.Print("this data will saved")
	log.Println(data)
	log.Print("data saved")
	return ""
}

////////////////////////////CONVERT-TO-STRUCT
func bytesToDataStruct(bytes string) (data []income) {
	dataString := string(bytes) //Converting Slice_Byte_Type to String_Type
	log.Print("data loaded")
	log.Print("conversion this data is started")
	dataStrings := strings.Split(dataString, "\n") // []string
	data = make([]income, 0)
	log.Print("\n", dataString)
	var err = makeError()
	for _, line := range dataStrings {
		fields := strings.Split(line, ";")
		element := income{}
		for _, field := range fields {
			values := strings.Split(field, ":")
			name, value := values[0], values[1]
			switch name {
			case "id":
				element.id, err = strconv.ParseInt(value, 10, 64)
				if err != nil {
					log.Fatal(fmt.Errorf("invalid field value %s: %w", value, err))
				}
			case "amount":
				element.amount, err = strconv.ParseUint(value, 10, 64)
				if err != nil {
					log.Fatal(fmt.Errorf("invalid field value %s: %w", value, err))
				}
			default:
				log.Fatal(fmt.Errorf("unknown field: %s", name))
			}
		}
		data = append(data, element)
	}
	log.Print("data converted to struct")
	return
}

////////////////////////////READ-FROM-FILE
func readFromFile(path string) (string, error) {
	log.Print("try open file ", path)
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(fmt.Errorf("can't read from %v: %w", path, err))
	}
	log.Print("file \"", path, "\" was opened")
	return string(bytes), nil
}

func main() {

	data3 := dataStructToBytes(data)//data3 -> converted []byte
	fmt.Println(data3)
	Write(data3, "data2.txt")//data3 -> write to file data2.txt
	//fmt.Println(data3)
	valueToRead, err := readFromFile("data2.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(valueToRead)
	valueToStruct := bytesToDataStruct(valueToRead)
	fmt.Print(valueToStruct)
}
