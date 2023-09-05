package pkg

import (
	"fmt"
	"log"
	"numfingpt/internal/functions"
)

func Exec(url, API_KEY string) {
	data, err := functions.FetchData(url)
	if err != nil {
		log.Println(err)
		return
	}
	convertedData, err := functions.ConvertData(data)
	if err != nil {
		log.Println(err)
		return
	}
	gptResponseText, err := functions.AskGpt(convertedData, API_KEY)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(gptResponseText)
}
