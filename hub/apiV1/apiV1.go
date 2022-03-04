package apiV1

import (
	"encoding/json"
	"fmt"
	"strings"
	"net/http"
	// "strconv"
  // "time"

  // b64 "encoding/base64"
	// "github.com/gorilla/mux"
	"github.com/magiconair/properties"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	// curl "github.com/andelf/go-curl"
	"os"
	"io/ioutil"
)

// prop
var p = properties.MustLoadFile("./go.prop", properties.UTF8)
var dbApi = p.MustGetString("user") + ":" + p.MustGetString("password") + "@tcp(" + p.MustGetString("host") + ":" + p.MustGetString("port") + ")/" + p.MustGetString("dtbApi")

type PathToUrl struct {
	METHOD string	`json:"method"`
	CLAIM string	`json:"claim"`
	PATH string		`json:"path"`
	PARAM string	`json:"param"`
	BODY string		`json:"body"`
}

type PathToUrlJson struct {
	Code int				`json:"code"`
	Data PathToUrl	`json:"data"`
}

type PathNotFoundJson struct {
	Code int				`json:"code"`
	Message string	`json:"message"`
}

// 404
var pathNotFoundJson = PathNotFoundJson{
	Code: 404,
	Message: "Path Not Found",
}

// default
var code = 200

// respon json
type ResponJson struct {
	Code interface{} 							`json:"code"`
	Data []map[string]interface{} `json:"data"`
}

// Api
func Api(w http.ResponseWriter, r *http.Request) {
	// path
	p := strings.TrimPrefix(r.URL.Path, "/api/v1/")
	// method
	// m := r.Method

	// db conn
	db, e := sql.Open("mysql", dbApi)
	if e != nil {
		// panic(err.Error())
	}
	defer db.Close()
	// db select
	rslt, e := db.Query("SELECT METHOD, CLAIM, PATH, PARAM, BODY FROM API_LIST WHERE URL = '" + p + "'")
	if e != nil {
		// panic(err.Error())
	}
	var pathToUrlJson = PathToUrlJson{Code: code}
	if rslt != nil {
		for rslt.Next() {
			var pathToUrl PathToUrl
			e = rslt.Scan(
				&pathToUrl.METHOD,
				&pathToUrl.CLAIM,
				&pathToUrl.PATH,
				&pathToUrl.PARAM,
				&pathToUrl.BODY,
			)
			if e != nil {
				// panic(err.Error())
			}
			pathToUrlJson.Data = pathToUrl
		}
	}

	// response header
	w.Header().Set("Content-Type", "application/json")

	if pathToUrlJson.Data.METHOD != "" {
		// json.NewEncoder(w).Encode(pathToUrlJson)

		fmt.Println(pathToUrlJson.Data.PATH)

		// path
		p2 := "http://10.251.182.14/" + strings.TrimPrefix(pathToUrlJson.Data.PATH, "http://localhost:5990/")

		fmt.Println(p2)

		// api hub
		// res, e := http.Get(pathToUrlJson.Data.PATH)
		res, e := http.Get(p2)

	  if e != nil {
	    // fmt.Print(e.Error())
	    os.Exit(1)
	  }

	  responseData, e := ioutil.ReadAll(res.Body)
	  if e != nil {
	    // log.Fatal(e)
	  }

		var responseObject ResponJson
	  json.Unmarshal(responseData, &responseObject)

		json.NewEncoder(w).Encode(responseObject)
	} else {
		json.NewEncoder(w).Encode(pathNotFoundJson)
	}
}
