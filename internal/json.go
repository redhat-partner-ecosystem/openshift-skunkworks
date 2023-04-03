package internal

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/txsvc/goreq"
)

// PrintJSON formats a struct into a JSON string
func PrintJSON(target interface{}) string {
	b, err := json.Marshal(target)
	if err != nil {
		// TODO really ?
		log.Fatal(err)
	}
	return (string)(b)
}

// PrettyPrintJSON prints a struct to STDOUT nicely
func PrettyPrintJSON(target interface{}) {
	b, err := json.Marshal(target)
	if err != nil {
		// TODO really ?
		log.Fatal(err)
	}

	var out bytes.Buffer
	json.Indent(&out, b, " ", " ")
	out.WriteTo(os.Stdout)
}

// GetJSON querys a url for an expected JSON struct
func GetJSON(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

// PostJSON posts a JSON object to a given URL
func PostJSON(url string, body interface{}, response interface{}) error {
	r, err := goreq.Request{
		Method:      "POST",
		Uri:         url,
		ContentType: "application/json",
		Body:        body,
	}.Do()
	defer r.Body.Close()

	if err != nil {
		return err
	}

	return json.NewDecoder(r.Body).Decode(response)
}

// PutJSON posts a JSON object to a given URL (HTTP PUT)
func PutJSON(url string, target interface{}) error {
	r, err := goreq.Request{
		Method:      "PUT",
		Uri:         url,
		ContentType: "application/json",
		Body:        target,
	}.Do()
	defer r.Body.Close()

	return err
}
