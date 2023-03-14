package models

import (
	"net/http"

	"github.com/renaldlie1/depd_go_echo/db"
)

type Pegawai struct {
	Id    string `json:"id" validate:"required,numeric,len=13"`
	Name  string `json:"name" validate:"required"`
	Age   string `json:"age" validate:"required"`
	Email string `json:"email" validate:"required"`
}

func FetchAllPegawai() (Response, error) {
	var obj Pegawai
	var arrobj []Pegawai
	var res Response

	con := db.CreateCon()
	sqlStatement := "SELECT * FROM pegawai"
	rows, err := con.Query(sqlStatement)

	if err != nil {
		return res, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Name, &obj.Age, &obj.Email)
		if err != nil {
			return res, err
		}
		arrobj = append(arrobj, obj)
	}

	res.Code = http.StatusOK
	res.Status = true
	res.Message = "Success"
	res.Data = arrobj

	return res, nil
}
