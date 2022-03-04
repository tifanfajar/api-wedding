module apiWedding

replace (
	apiWedding/apiWeddingModels => ./models
	// datacom/dtcImcServices => ./services/imcService
	apiWedding/dtcJwtServices => ./services/jwtService
	// datacom/dtcRcaServices => ./services/rcaService
	apiWedding/themeServices => ./services/themeServices
	hub/apiV1 => ./hub/apiV1
)

require (
	apiWedding/apiWeddingModels v0.0.0-00010101000000-000000000000 // indirect
	apiWedding/themeServices v0.0.0-00010101000000-000000000000
	github.com/andelf/go-curl v0.0.0-20200630032108-fd49ff24ed97 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/gorilla/mux v1.8.0
	github.com/magiconair/properties v1.8.4
	hub/apiV1 v0.0.0-00010101000000-000000000000
)

go 1.13
