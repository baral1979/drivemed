package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/tealeg/xlsx"
)

func generateTable(w http.ResponseWriter, r *http.Request) {

	file, handle, err := r.FormFile("file")
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}
	defer file.Close()

	//respondWithJSON(w, http.StatusOK, handle)
	saveFile(w, file, handle)
	generateTableFromExcelFile(w, "./files/"+handle.Filename)
}

func generateModel(w http.ResponseWriter, r *http.Request) {

	file, handle, err := r.FormFile("file")
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}
	defer file.Close()

	//respondWithJSON(w, http.StatusOK, handle)
	saveFile(w, file, handle)
	generateModelFromExcelFile(w, "./files/"+handle.Filename)
}

func generateMapping(w http.ResponseWriter, r *http.Request) {

	file, handle, err := r.FormFile("file")
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}
	defer file.Close()

	//respondWithJSON(w, http.StatusOK, handle)
	saveFile(w, file, handle)
	generateModelMapping(w, "./files/"+handle.Filename)
}

func formatColumnName(name string) string {
	name = strings.Replace(name, " ", "", -1)
	name = strings.Replace(name, "/", "", -1)
	name = strings.Replace(name, "(", "", -1)
	name = strings.Replace(name, ")", "", -1)
	return name
}

func generateModelFromExcelFile(w http.ResponseWriter, filename string) {
	xlFile, err := xlsx.OpenFile(filename)
	if err != nil {
		respondWithError(w, 500, "Unable to process the file "+filename)
	}

	sheet := xlFile.Sheets[1]
	columnNames := []string{}
	for rowNo, row := range sheet.Rows {
		for _, cell := range row.Cells {
			if rowNo == 0 {
				s := cell.String()
				columnNames = append(columnNames, s)
			} else {
				break
			}
		}
	}

	model := "type <typename> struct {"
	for _, col := range columnNames {
		colname := formatColumnName(col)
		model += formatColumnName(colname) + "\tstring\t" + "`json:\"" + strings.ToLower(colname) + "\"`" + "\r\n"
	}
	model += "}"
	//	println(model)

	io.WriteString(w, model)
	w.WriteHeader(http.StatusOK)
}

func generateModelMapping(w http.ResponseWriter, filename string) {
	xlFile, err := xlsx.OpenFile(filename)
	if err != nil {
		respondWithError(w, 500, "Unable to process the file "+filename)
	}

	sheet := xlFile.Sheets[1]
	columnNames := []string{}
	for rowNo, row := range sheet.Rows {
		for _, cell := range row.Cells {
			if rowNo == 0 {
				s := cell.String()
				columnNames = append(columnNames, s)
			} else {
				break
			}
		}
	}

	res := ""
	for _, col := range columnNames {
		colname := formatColumnName(col)
		res += "r." + colname + " = m[\"" + colname + "\"]\r\n"
	}
	//	println(model)

	io.WriteString(w, res)
	w.WriteHeader(http.StatusOK)
}

func generateTableFromExcelFile(w http.ResponseWriter, filename string) {
	xlFile, err := xlsx.OpenFile(filename)
	if err != nil {
		respondWithError(w, 500, "Unable to process the file "+filename)
	}

	sheet := xlFile.Sheets[1]
	columnNames := []string{}
	for rowNo, row := range sheet.Rows {
		for _, cell := range row.Cells {
			if rowNo == 0 {
				s := cell.String()
				columnNames = append(columnNames, s)
			} else {
				break
			}
		}
	}

	query := "CREATE TABLE <TABLE_NAME> ("
	for _, col := range columnNames {
		query += formatColumnName(col) + " VARCHAR(max),"
	}
	query += ")"
	respondWithJSON(w, http.StatusOK, query)
}
