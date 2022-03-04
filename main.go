package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"hub/apiV1"
	"math"
	"strconv"
	"time"

	// "bytes"
	"apiWedding/themeServices"
	"log"

	// "math"
	"net/http"
	// "strconv"
	// "time"

	"github.com/gorilla/mux"
	"github.com/magiconair/properties"
)

func homePage(nm string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// start := time.Now()
		layout := "2006-01-02T15:04:05.000Z"
		str := "2014-11-12T11:45:26.371Z"
		t, err := time.Parse(layout, str)
		if err != nil {
			json.NewEncoder(w).Encode(err)
		}
		var print bytes.Buffer
		print.WriteString("Tanggal awal " + t.Format("2006-01-02"))
		print.WriteString(" Tanggal Sekarang " + time.Now().Format("January-02-2006"))
		print.WriteString(" Selisih " + time.Since(t).String())
		print.WriteString(" Selisih Jam " + strconv.Itoa(int(time.Since(t).Hours())))
		print.WriteString(" Selisih Hari " + strconv.Itoa(int(math.Floor(time.Since(t).Hours()/24))))
		print.WriteString(" Selisih Bulan " + strconv.Itoa(int(math.Floor(time.Since(t).Hours()/24/365*12))))
		print.WriteString(" Selisih Tahun " + strconv.Itoa(int(math.Floor(time.Since(t).Hours()/24/365))))

		json.NewEncoder(w).Encode(print.String())
	}
}

// func handleRequests(nm string) {
func handleRequests() {
	p := properties.MustLoadFile("./go.prop", properties.UTF8)
	n := p.MustGetString("appName")
	v := p.MustGetString("version")
	fmt.Println(n + " have started\n" + v)

	b := p.MustGetString("baseUrlGolang")
	gp := ":" + p.MustGetString("golangPort")
	bDtc := b + "dtc/"

	bHubV1 := "/api/v1"
	nm := "tes"
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", homePage(nm))

	//JWT
	// r.HandleFunc("/dtc/login", dtcJwtServices.Login)

	//API
	// r.HandleFunc(bHubV1, apiV1.Login)
	r.HandleFunc(bDtc+"theme/create", themeServices.CreateTheme).Methods("POST")
	r.HandleFunc(bDtc+"theme/c", themeServices.GetCount).Queries("category", "{category}").Methods("GET")
	r.HandleFunc(bDtc+"theme/l", themeServices.GetList).Queries("category", "{category}", "pages", "{pages}").Methods("GET")
	r.HandleFunc(bHubV1+"/{path1}", apiV1.Api)
	r.HandleFunc(bHubV1+"/{path1}/{path2}", apiV1.Api)
	r.HandleFunc(bHubV1+"/{path1}/{path2}/{path3}", apiV1.Api)
	r.HandleFunc(bHubV1+"/{path1}/{path2}/{path3}/{path4}", apiV1.Api)
	r.HandleFunc(bHubV1+"/{path1}/{path2}/{path3}/{path4}/{path5}", apiV1.Api)
	r.HandleFunc(bHubV1+"/{path1}/{path2}/{path3}/{path4}/{path5}/{path6}", apiV1.Api)
	r.HandleFunc(bHubV1+"/{path1}/{path2}/{path3}/{path4}/{path5}/{path6}/{path7}", apiV1.Api)
	r.HandleFunc(bHubV1+"/{path1}/{path2}/{path3}/{path4}/{path5}/{path6}/{path7}/{path8}", apiV1.Api)
	r.HandleFunc(bHubV1+"/{path1}/{path2}/{path3}/{path4}/{path5}/{path6}/{path7}/{path8}/{path9}", apiV1.Api)
	r.HandleFunc(bHubV1+"/{path1}/{path2}/{path3}/{path4}/{path5}/{path6}/{path7}/{path8}/{path9}/{path10}", apiV1.Api)

	//theme

	log.Fatal(http.ListenAndServe(gp, r))
}

func main() {
	handleRequests()
}
