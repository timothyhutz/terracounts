package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
)

var file []byte
var file_err error
func init()  {
	if len(os.Args[1:]) < 1 {
		log.Println("You did not pass a file JSON list of accounts variable defined, looking for data.json")
		file, file_err = os.ReadFile("data.json")
		if file_err != nil {
		log.Fatal(file_err)
		}
	} else {
		file, file_err = os.ReadFile(os.Args[1])
		if file_err != nil {
		log.Fatal(file_err)
		}
	}
}

func main(){
	var accounts []string
	json.Unmarshal([]byte(file), &accounts)
	log.Println("accounts your code is deploying to", accounts)
	cmd := exec.Command("terraform", "init")
	stdout, err := cmd.CombinedOutput()
	log.Printf("%s\n", stdout)
		if err != nil {
			log.Fatal(err)
		}
	for _, element := range accounts {
		cmd := exec.Command("terraform", "apply", "-auto-approve", fmt.Sprintf("-var=account=%s", element))
		stdout, err := cmd.CombinedOutput()
		log.Printf("%s\n", stdout)
			if err != nil {
				log.Fatal(err)
			}
	}
	
}