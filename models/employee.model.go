package models

import (
	"echo-api/db"
	"net/http"
)

type Employee struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phonenumber"`
}

func FetchAllEmployee() (Response, error) {
	var obj Employee
	var arrObj []Employee
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM employee"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Name, &obj.Address, &obj.PhoneNumber)

		if err != nil {
			return res, err
		}

		arrObj = append(arrObj, obj)
	}
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrObj

	return res, nil
}

func StoreEmployee(name string, address string, phonenumber string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT employee (name, address, phonenumber) VALUES (?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(name, address, phonenumber)
	if err != nil {
		return res, err
	}

	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"last_inserted_id": lastInsertedId,
	}

	return res, nil
}

func UpdateEmployee(id int, name string, address string, phonenumber string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE employee SET name = ?, address = ?, phonenumber = ? WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(name, address, phonenumber, id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_updated": rowsAffected,
	}

	return res, nil
}

func DeleteEmployee(id int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM employee WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_deleted": rowsAffected,
	}

	return res, nil
}
