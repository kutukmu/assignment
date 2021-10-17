package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	shell "github.com/ipfs/go-ipfs-api"
)

func main() {


	hashVal := "QmP8jTG1m9GSDJLCbeWhVSVgEzCPPwXRdCRuJtQ5Tz9Kc9"
	getData(hashVal)
	// ************************ //
	setData()


	
}

func getData(hashVal string){
	// First start the server
	sh := shell.NewShell("localhost:5001")

	//To be able to output the result we need to have the local directory

	localdir, err := os.Getwd()

	if err != nil{
		fmt.Printf("error: %s", err)
	}

	// after being sure that we have the local directory and server
	// we can start retrieving the hash, using Get function 


	id := sh.Get(hashVal, localdir) // func (s *Shell) Get(hash, outdir string) error

	// get function return an error so if the id is not null then we need to print the error
	if id != nil{
		fmt.Printf("error: %s", id)
	}

	// ReadFile reads the file named by filename and returns the contents.
	result,readerr := ioutil.ReadFile(hashVal) // func ReadFile(filename string) ([]byte, error)

	if readerr != nil{
		fmt.Printf("error: %s", readerr)
	}

	// we have the result as slice of bytes so we need to convert it to string
	finalResult := string(result)

	// Print the resutl 
	fmt.Println(finalResult)

}


func setData() {

	// First start the server
	sh := shell.NewShell("localhost:5001")

	//To be able to output the result we need to have the local directory

	localdir, err := os.Getwd()

	// if the err is not null then print the error
	if err != nil{
		fmt.Printf("error: %s", err)
	}

	// store a string in IPFS 
	id, addErr := sh.Add(strings.NewReader("Taubyte assignment"))
	if addErr != nil {
		fmt.Printf("error: %s", addErr)
	}

	// we can start retrieving the hash, using Get function 
	getId := sh.Get(id, localdir) // func (s *Shell) Get(hash, outdir string) error

	// get function return an error so if the id is not null then we need to print the error
	if getId != nil{
		fmt.Printf("error: %s", getId)
	}

	result,readerr := ioutil.ReadFile(id) // func ReadFile(filename string) ([]byte, error)

	if readerr != nil{
		fmt.Printf("error: %s", readerr)
	}

	// we have the result as slice of bytes so we need to convert it to string
	finalResult := string(result)

	// Print the resutl 
	fmt.Println(finalResult)
}