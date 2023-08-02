package pkg

import (
	"encoding/json"
	"log"
	"os"
)

func AddingDataToJson() {
	Parviz := Student{
		FirstName:   "Parviz",
		LastName:    "Mirzoev",
		Age:         35,
		PhoneNumber: "+992900000001",
		Status:      "Married",
	}

	jsonBytes, err := json.Marshal(&Parviz)
	if err != nil {
		log.Println(err)
		return
	}

	file, err := os.Create("./data.json")
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()
	_, err = file.Write(jsonBytes)
	if err != nil {
		log.Println(err)
		return
	}

	//log.Printf("JSON string is %s\n", string(jsonBytes))
}
