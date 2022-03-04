package dtcRcaServices

import (
	"encoding/json"
	"fmt"
	"net/http"

	// "io/ioutil"
	// "fmt"
	// "os"
	"io/ioutil"
)

type ResponJson struct {
	Code interface{}              `json:"code"`
	Data []map[string]interface{} `json:"data"`
}

func GetRcaList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	url := "http://10.251.182.14/n/dtc/imc/l"

	res, e := http.Get(url)

	if e != nil {
		fmt.Println(e)
	}

	responseData, e := ioutil.ReadAll(res.Body)
	if e != nil {
		// log.Fatal(e)
	}

	var responseObject ResponJson
	json.Unmarshal(responseData, &responseObject)

	json.NewEncoder(w).Encode(responseObject)

}
