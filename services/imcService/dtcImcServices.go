package dtcImcServices

import (
	"database/sql"
	"apiWedding/apiWeddingModels "
	"encoding/json"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/magiconair/properties"
	// "io/ioutil"
	"fmt"
)

// fmt.Println('console log')

// prop
var p = properties.MustLoadFile("./go.prop", properties.UTF8)
var dbCon = p.MustGetString("user") + ":" + p.MustGetString("password") + "@tcp(" + p.MustGetString("host") + ":" + p.MustGetString("port") + ")/" + p.MustGetString("database")

// default
var code = 200

func GetImcCount(w http.ResponseWriter, r *http.Request) {
	// conn
	db, err := sql.Open("mysql", dbCon)
	if err != nil {
		// panic(err.Error())
	}
	defer db.Close()
	// select
	results, err := db.Query("SELECT count(1) total FROM SMS_BROADCAST_DATACOM")
	if err != nil {
		// panic(err.Error())
	}
	if results != nil {
		var ImcCountJson = dtcModels.ImcCountJson{Code: code}
		for results.Next() {
			var imcCount dtcModels.ImcCount
			err = results.Scan(&imcCount.Total)
			if err != nil {
				// panic(err.Error())
			}
			ImcCountJson.Data.Total = imcCount.Total
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ImcCountJson)
	}
}

func GetImcList(w http.ResponseWriter, r *http.Request) {
	// query param
	page := mux.Vars(r)["page"]
	pageInt, _ := strconv.Atoi(page)
	if page != "" {
		pageInt = (pageInt - 1) * 10
	} else {
		pageInt = 0
	}
	// conn
	db, err := sql.Open("mysql", dbCon)
	if err != nil {
		// panic(err.Error())
	}
	defer db.Close()
	// select
	results, err := db.Query(
		"select "+
			"REGION reg, "+
			"X_CHANGE_TYPE ct, "+
			"SEVERITY severity, "+
			"'TICKET1234567890' ticket, "+
			"'TICKET STATUS' ts, "+
			"CREATION_TIME startTime, "+
			"ifnull(CLEAR_TIME,'') endTime, "+
			"'01 hours 01 minutes 01 seconds' duration, "+
			"'NetworkImpact' networkImpact, "+
			"'RootCause' rootCause, "+
			"'Action' action, "+
			"'CATEGORY' cat, "+
			"PIC pic, "+
			"DEVICE_NAME_NE deviceNameNe, "+
			"MO_NAME_NE moNameNe "+
			"from SMS_BROADCAST_DATACOM "+
			"limit ?, 10",
		pageInt,
	)
	if err != nil {
		// panic(err.Error())
	}
	// var ImcListJson = dtcModels.ImcListJson{Code: code}
	// for results.Next() {
	// 	var imc dtcModels.Imc
	// 	err = results.Scan(
	// 		&imc.Region,
	// 		&imc.ChangeType,
	// 		&imc.Severity,
	// 		&imc.Ticket,
	// 		&imc.TicketStatus,
	// 		&imc.IncidentStartTime,
	// 		&imc.IncidentCloseTime,
	// 		&imc.DurationToSolve,
	// 		&imc.NetworkImpact,
	// 		&imc.RootCause,
	// 		&imc.Action,
	// 		&imc.Category,
	// 		&imc.PIC,
	// 		&imc.DeviceNameNe,
	// 		&imc.MoNameNe,
	// 	)
	// 	if err != nil {
	// 		// panic(err.Error())
	// 	}
	// 	ImcListJson.Data = append(ImcListJson.Data, imc)
	// }

	// Get the column names from the query
	var columns []string
	columns, err = results.Columns()
	checkErr(err)

	fmt.Println(columns)

	// colNum := len(columns)
	//
	// var results []map[string]interface{}
	//
	// for rows.Next() {
	//   // Prepare to read row using Scan
	//   r := make([]interface{}, colNum)
	//   for i := range r {
	//     r[i] = &r[i]
	//   }
	//
	//   // Read rows using Scan
	//   err = rows.Scan(r...)
	//   checkErr(err)
	//
	//   // Create a row map to store row's data
	//   var row = map[string]interface{}{}
	//   for i := range r {
	//     row[columns[i]] = r[i]
	//   }
	//
	//   // Append to the final results slice
	//   results = append(results, row)
	// }

	w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(ImcListJson)

	// insert
	// insert, err := db.Query("INSERT INTO test VALUES ( 2, 'TEST' )")
	// if err != nil {
	//   panic(err.Error())
	// }
	// defer insert.Close()
}

func checkErr(err error) {
	if err != nil {
		// log.Fatal(err)
	}
}

// func createNewArticle(w http.ResponseWriter, r *http.Request) {
//   reqBody, _ := ioutil.ReadAll(r.Body)
//   var article Article
//   json.Unmarshal(reqBody, &article)
//   Articles = append(Articles, article)
//   json.NewEncoder(w).Encode(article)
// }
//
// func deleteArticle(w http.ResponseWriter, r *http.Request) {
//   vars := mux.Vars(r)
//   id := vars["id"]
//   var article Article
//   for index, article := range Articles {
//     if article.Id == id {
//       Articles = append(Articles[:index], Articles[index+1:]...)
//     }
//   }
//   json.NewEncoder(w).Encode(article)
// }
