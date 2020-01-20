# Convert struct data for write a file , buckup file , write to file , read from file , convert data from file to struct data

## H2 You can use it

The code contains comments and was developed in the Golang

```Go
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
```


Enjoy!
