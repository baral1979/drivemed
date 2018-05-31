package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gorilla/mux"
	"github.com/tealeg/xlsx"
)

// App type
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// Initialize the app
func (a *App) Initialize(server, user, password, dbname string, port int) {
	var err error

	// Create connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, dbname)

	// Create connection pool
	a.DB, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: " + err.Error())
	}
	log.Printf("Connected!\n")

	// Close the database connection pool after program executes
	//defer a.DB.Close()

	// Initialize Router
	a.Router = mux.NewRouter()
	a.initializeRoutes()

}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/test", a.test).Methods("GET")
	a.Router.HandleFunc("/api/products/{lang}", a.getProducts).Methods("GET")
	a.Router.HandleFunc("/api/products/upload/{lang}", a.uploadProducts).Methods("POST")
	//a.Router.HandleFunc("api/products/fr", a.getProductsFr).Methods("GET")
	// a.Router.HandleFunc("/product/{id:[0-9]+}", a.getProduct).Methods("GET")
	// a.Router.HandleFunc("/product/{id:[0-9]+}", a.updateProduct).Methods("PUT")
	// a.Router.HandleFunc("/product/{id:[0-9]+}", a.deleteProduct).Methods("DELETE")
	// a.Router.HandleFunc("/product/upload", a.uploadProducts).Methods("POST")

	a.Router.HandleFunc("/api/products/{lang}/upload", a.uploadProducts).Methods("POST")

	// helper endpoint
	a.Router.HandleFunc("/api/files/generateTable", generateTable).Methods("POST")
	a.Router.HandleFunc("/api/files/generateModel", generateModel).Methods("POST")
	a.Router.HandleFunc("/api/files/generateMapping", generateMapping).Methods("POST")
	a.Router.HandleFunc("/api/files/generateInsert", generateInsert).Methods("POST")
}

func (a *App) uploadProducts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lang := vars["lang"]

	file, handle, err := r.FormFile("file")
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}
	defer file.Close()

	saveFile(w, file, handle)

	xlFile, err := xlsx.OpenFile("./files/" + handle.Filename)
	if err != nil {
		respondWithError(w, 500, "Unable to process the file ./files/"+handle.Filename)
	}

	if lang == "fr" {
		err = deleteProductsFr(a.DB)
	} else {
		err = deleteProductsEn(a.DB)
	}

	if err != nil {
		respondWithError(w, 500, "Could not empty table")
	}
	sheet := xlFile.Sheets[1]
	columnNames := []string{}
	for rowNo, row := range sheet.Rows {
		jsonBlob := map[string]string{}
		for colNo, cell := range row.Cells {
			if rowNo == 0 {
				s := formatColumnName(cell.String())
				columnNames = append(columnNames, s)
			} else {
				// Build a map and render it out
				if colNo < len(columnNames) {
					s := cell.String()
					jsonBlob[columnNames[colNo]] = s
				} else {
					k := fmt.Sprintf("column_%d", colNo+1)
					columnNames = append(columnNames, k)
					s := cell.String()
					jsonBlob[k] = s
				}
			}
		}
		if rowNo > 0 {
			if lang == "fr" {
				productFr := getProductFrFromBlob(jsonBlob)

				err := productFr.create(a.DB)

				if err != nil {
					respondWithError(w, http.StatusInternalServerError, err.Error())

				}

			} else {
				productEn := getProductEnFromBlob(jsonBlob)
				err := productEn.create(a.DB)
				if err != nil {
					respondWithError(w, http.StatusInternalServerError, err.Error())

				}
				//respondWithJSON(w, http.StatusOK, nil)
			}
			//respondWithJSON(w, http.StatusOK, jsonBlob)
			//var p = getProductFromJSON(jsonBlob)
			//p.create(a.DB)
			//src, err := json.Marshal(jsonBlob)
			//if err != nil {
			//return output, fmt.Errorf("Can't render JSON blob, %s", err)
			//}

			//break
			//output = append(output, string(src))
		}
	}

	//respondWithJSON(w, http.StatusOK, jsonBlob)
	// saveFile(w, file, handle)
	// //validateFile(w, "./files/"+handle.Filename)
	w.WriteHeader(http.StatusOK)
	// //go a.processProductData("./files/" + handle.Filename)
}

func saveFile(w http.ResponseWriter, file multipart.File, handle *multipart.FileHeader) {
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}

	err = ioutil.WriteFile("./files/"+handle.Filename, data, 0666)
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}
}

func (a *App) getProducts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lang := vars["lang"]
	products, err := getProducts(a.DB, lang)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, products)
}

func (a *App) test(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, 10)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Run the app
func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

// func (a *App) ensureDatabaseOk() {
// 	if _, err := a.DB.Exec(tableProductsCreationQuery); err != nil {
// 		log.Fatal(err)
// 	}

// 	if _, err := a.DB.Exec(tableDriveCreationQuery); err != nil {
// 		log.Fatal(err)
// 	}
// }
