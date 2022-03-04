package apiWeddingModels

// imc count
type ImcCount struct {
	Total int `json:"total"`
}
type ImcCountJson struct {
	Code int      `json:"code"`
	Data ImcCount `json:"data"`
}

type ThemeCount struct {
	Total int `json:"total"`
}

type ThemeCountJson struct {
	Code int        `json:"code"`
	Data ThemeCount `json:"data"`
}

type ThemeList struct {
	Idtheme      int    `json:"idtheme"`
	Themename    string `json:"themename"`
	Description  string `json:"description"`
	Likes        int    `json:"likes"`
	Images       string `json:"image"`
	Created_date string `json:"created_date"`
	Updated_date string `json:"updated_date"`
	Category     string `json:"category"`
	Idpages      int    `json:"idpages"`
	Slug         string `json:"slug"`
}

type ThemeListJson struct {
	Code int         `json:"code"`
	Data []ThemeList `json:"data"`
}

type BankObj struct {
	Bank_name string `json:"bank_name"`
	Bank_rek  string `json:"bank_rek"`
}

type CreateTheme struct {
	Themename   string `json:"themename"`
	Description string `json:"description"`
	Images      string `json:"images"`
	Category    string `json:"category"`
	Cover       string `json:"cover"`
	Longitude   string `json:"longitude"`
	Latitude    string `json:"latitude"`
	Male_bank   string `json:"male_bank"`
	Female_bank string `json:"female_bank"`
	Male_rek    string `json:"male_rek"`
	Female_rek  string `json:"female_rek"`
	Is_pandemi  int    `json:"is_pandemi"`
	Event_date  string `json:"event_date"`
	Music       string `json:"music"`
	Is_music    string `json:"is_music"`
	Iduser      string `json:"iduser"`
	Is_rek      string `json:"is_rek"`
	Slug        string `json:"slug"`
}

type IdTheme struct {
	IdThemes string `json:"idtheme"`
	IdPages  string `json:"idpage"`
}

type CreateThemeJson struct {
	Code int       `json:"code"`
	Data []IdTheme `json:"data"`
}
