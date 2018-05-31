package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/tealeg/xlsx"
)

func generateInsert(w http.ResponseWriter, r *http.Request) {

	file, handle, err := r.FormFile("file")
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}
	defer file.Close()

	//respondWithJSON(w, http.StatusOK, handle)
	saveFile(w, file, handle)
	generateInsertQuery(w, "./files/"+handle.Filename)
}

func generateInsertQuery(w http.ResponseWriter, filename string) {
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

	// "INSERT INTO products(name, price) VALUES($1, $2) RETURNING id",
	// p.Name, p.Price).Scan(&p.ID)

	line1 := "query := \"INSERT INTO SourceEn ("
	line2 := "query += \"VALUES(\"\r\n"
	//line3 := ""

	for i, col := range columnNames {
		colname := formatColumnName(col)
		line1 += colname
		line2 += "query += \"'\" + strings.Replace(p." + colname + ",\"'\",\"''\",-1) +\"'\""
		//line3 += "p." + colname
		if i < len(columnNames)-1 {
			line1 += ", "
			line2 += "+\",\"\r\n"
			//	line3 += ", "
		} else {
			line1 += ")\""
			line2 += "\r\nquery += \")\""
			//	line3 += ")"
		}
	}
	//	println(model)

	line1 += "\r\n" + line2

	io.WriteString(w, line1)
	w.WriteHeader(http.StatusOK)
}
