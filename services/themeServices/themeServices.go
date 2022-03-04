package themeServices

import (
	"apiWedding/apiWeddingModels"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/magiconair/properties"
	// "io/ioutil"
)

var p = properties.MustLoadFile("./go.prop", properties.UTF8)
var dbCon = p.MustGetString("user") + ":" + p.MustGetString("password") + "@tcp(" + p.MustGetString("host") + ":" + p.MustGetString("port") + ")/" + p.MustGetString("database")

// default
var code = 200

func GetCount(w http.ResponseWriter, r *http.Request) {

	categories := mux.Vars(r)["category"]
	category, _ := strconv.Atoi(categories)
	db, err := sql.Open("mysql", dbCon)
	if err != nil {
		// panic(err.Error())
	}
	defer db.Close()
	var queryFinal string
	query := "select count(*) as total from db_wedding.tb_theme b join db_wedding.tb_pages a on a.idpages = b.idpages where 1=1 "
	if categories != "" {
		query = query + "and b.category = ? "
	}
	queryFinal = query + "order by b.created_date"

	results, err := db.Query(queryFinal)
	if categories != "" {
		results, err = db.Query(queryFinal, category)
	}
	if err != nil {
		// panic(err.Error())
	}
	if results != nil {
		var ThemeCountJson = apiWeddingModels.ThemeCountJson{Code: code}
		for results.Next() {
			var ThemeCount apiWeddingModels.ThemeCount
			err = results.Scan(&ThemeCount.Total)
			if err != nil {

			}
			ThemeCountJson.Data.Total = ThemeCount.Total
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ThemeCountJson)
	}
}

func GetList(w http.ResponseWriter, r *http.Request) {
	categories := mux.Vars(r)["category"]
	category, _ := strconv.Atoi(categories)

	pages := mux.Vars(r)["pages"]
	page, _ := strconv.Atoi(pages)

	db, err := sql.Open("mysql", dbCon)
	if err != nil {

	}
	defer db.Close()
	var queryFinal string
	query := "select b.*, a.slug from db_wedding.tb_theme b join db_wedding.tb_pages a on a.idpages = b.idpages where 1=1 "
	if categories != "" {
		query = query + "and category = ? "
	}
	if pages != "" && pages != "0" {
		page = (page - 1) * 10
	} else {
		page = 0
	}

	queryFinal = query + "order by b.created_date limit ?, 10"

	results, err := db.Query(queryFinal, page)

	if categories != "" {
		results, err = db.Query(queryFinal, category, page)
	}
	if err != nil {
		fmt.Println(queryFinal)
		fmt.Println("err 3")
	}
	if results != nil {
		var ThemeListJson = apiWeddingModels.ThemeListJson{Code: code}

		for results.Next() {
			var ThemeList apiWeddingModels.ThemeList
			err = results.Scan(
				&ThemeList.Idtheme,
				&ThemeList.Themename,
				&ThemeList.Description,
				&ThemeList.Likes,
				&ThemeList.Images,
				&ThemeList.Created_date,
				&ThemeList.Updated_date,
				&ThemeList.Category,
				&ThemeList.Idpages,
				&ThemeList.Slug,
			)

			ThemeListJson.Data = append(ThemeListJson.Data, ThemeList)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ThemeListJson)
	}

}

func CreateTheme(w http.ResponseWriter, r *http.Request) {
	var Theme apiWeddingModels.CreateTheme
	current := time.Now()
	db, err := sql.Open("mysql", dbCon)
	if err != nil {
		fmt.Println("error con")
	}
	err = json.NewDecoder(r.Body).Decode(&Theme)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	query := "insert into db_wedding.tb_theme values (0, "
	query = query + "'" + Theme.Themename + "'" + ", "
	query = query + "'" + Theme.Description + "'" + ", "
	query = query + "0, "
	query = query + "'', "
	query = query + "'" + current.Format("2006-01-02 15:04:05") + "'" + ", "
	query = query + "'" + current.Format("2006-01-02 15:04:05") + "'" + ", "
	query = query + "'" + Theme.Category + "', "
	query = query + "''"

	query = query + ")"
	results, err := db.Query(query)
	if err != nil {

	}
	rows, err := db.Query("SELECT idtheme from db_wedding.tb_theme order by idtheme desc limit 0, 1")
	if results != nil {
		var CreateThemeJson = apiWeddingModels.CreateThemeJson{Code: code}
		var IdTheme apiWeddingModels.IdTheme
		for rows.Next() {
			err = rows.Scan(&IdTheme.IdThemes)
		}
		query = "insert into db_wedding.tb_pages values (0, "
		query = query + "'" + Theme.Slug + "'" + ", "
		query = query + "'', "
		query = query + "'" + Theme.Longitude + "', "
		query = query + "'" + Theme.Latitude + "', "
		query = query + "'" + Theme.Male_bank + "', "
		query = query + "'" + Theme.Female_bank + "', "
		query = query + "'" + Theme.Male_rek + "', "
		query = query + "'" + Theme.Female_rek + "', "
		query = query + "1, "
		query = query + "'" + Theme.Event_date + "', "
		query = query + "'', "
		query = query + Theme.Is_music + ", "
		query = query + "'" + current.Format("2006-01-02 15:04:05") + "'" + ", "
		query = query + "'" + current.Format("2006-01-02 15:04:05") + "'" + ", "
		query = query + "'" + Theme.Iduser + "', "
		query = query + Theme.Is_rek + ", "
		query = query + "'" + IdTheme.IdThemes + "'"
		query = query + ")"

		results2, err := db.Query(query)
		if err != nil {

		}

		if results2 != nil {
			rows2, err := db.Query("SELECT idpages from db_wedding.tb_pages order by idpages desc limit 0, 1")
			if err != nil {

			}
			for rows2.Next() {
				err = rows2.Scan(&IdTheme.IdPages)
				if err != nil {

				}
				CreateThemeJson.Data = append(CreateThemeJson.Data, IdTheme)
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(CreateThemeJson)
			fmt.Println(CreateThemeJson.Data)
		}

	}
	// Do something with the Person struct...
	fmt.Println(query)
	// fmt.Fprintf(w, query, results)
}
