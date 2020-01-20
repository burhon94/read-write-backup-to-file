package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type income struct {
	id     int64
	amount uint64
}

var datai = []income{
	{1, 1_000_000},
	//{2, 2_000_000},
	{3, 3_000_000},
	{4, 4_000_000},
	//{5, 5_000_000},
}

var customError = errors.New("custom error")

func makeError() error {
	return fmt.Errorf("error generated: %w", customError)
}

//var dataBytes = [] byte{}

////////////////////////////CONVERT-TO-WRITE
func dataStructToBytes(data []income) (dataBytes []byte) {
	log.Print("conversion this struct data is started")
	log.Print(data)
	dataStrings := make([]string, len(data))
	for i, datum := range data {
		dataStrings[i] = fmt.Sprintf("id:%d;amount:%d", datum.id, datum.amount)
	}
	oneString := strings.Join(dataStrings, "\n")
	dataBytes = []byte(oneString)
	log.Println("data converted to slice[] byte\n", dataBytes)
	return
}

////////////////////////////WRITE-TO-FILE
func Write(data []byte, path string) string {
	var _, err1 = os.Stat(path)
	if os.IsNotExist(err1) {
		log.Print("file \"", path, "\" not exist for write")
		log.Print("file \"", path, "\" will be create")
		var file, err = os.Create(path)
		if err != nil {
			return "file could not be created successfully"
		}
		defer file.Close()
		log.Print("file \"", path, "\" created successfully")
		log.Print("this data will saved")
		log.Print(data)
		err = ioutil.WriteFile(path, data, 0666)
		if err != nil {
			log.Fatal(fmt.Errorf("can't save to %s: %w", path, err))
		}
		log.Println("data saved")
		return ""
	}
	log.Print("file ", path, " is exist")
	var srcFile, err2 = os.Open(path)
	if err2 != nil {
		log.Fatal(fmt.Errorf("can't check source file %s: %w", path, err2))
	}
	log.Print("source file was checked successfully")
	dateTimeBackup := time.Now().Format("Backup(01-02-2006-15-04-5).txt")
	log.Print("copy date and time for copy")
	fmt.Println(dateTimeBackup)
	log.Print("try create file to backup")
	var destFile, err = os.Create(dateTimeBackup)
	if err != nil {
		log.Fatal(fmt.Errorf("can't create file to backup %s: %w", path, err))
	}
	log.Print("backup file was create")
	defer destFile.Close()
	log.Print("try copy from file to backup file")
	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		log.Fatal(fmt.Errorf("can't finish copying %s: %w", path, err))
	}
	log.Print("file was copy to backup file")
	log.Print("this data will saved")
	log.Print(data)
	err = ioutil.WriteFile(path, data, 0666)
	if err != nil {
		log.Fatal(fmt.Errorf("can't save to %s: %w", path, err))
	}
	log.Println("data saved")
	return ""
}

////////////////////////////CONVERT-TO-STRUCT
func StringToDataStruct(bytes string) (data []income) {
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
	log.Println("data converted to struct")
	return
}

////////////////////////////READ-FROM-FILE
func readFromFile(path string) (string, error) {
	log.Print("try open file ", path)
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(fmt.Errorf("can't read from %v: %w", path, err))
	}
	log.Println("file \"", path, "\" was opened")
	return string(bytes), nil
}

func main() {
	data3 := dataStructToBytes(datai)             //data3 -> converted []byte
	Write(data3, "data2.txt")                     //data3 -> write to file data2.txt if not exist create data2.txt
	valueToRead, err := readFromFile("data2.txt") //read from data2.txt
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(valueToRead)                         //show data from data2.txt
	valueToStruct := StringToDataStruct(valueToRead) //data from data2.txt converted to struct []income
	fmt.Println(valueToStruct)                       //show struct []income after convert
}
