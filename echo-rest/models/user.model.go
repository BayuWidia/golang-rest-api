package models

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/bayuwidia/echo-rest/db"
	"github.com/bayuwidia/echo-rest/utils"
)

type User struct {
	Id              int            `"json:"id"`
	Name            string         `"json:"name"`
	Email           string         `"json:"email"`
	EmailVerifiedAt sql.NullString `"json:"emailVerifiedAt"`
	Password        string         `"json:"password"`
	Role            string         `"json:"role"`
	RememberToken   sql.NullString `"json:"rememberToken"`
	CreatedAt       time.Time      `"json:"createdAt"`
	UpdatedAt       time.Time      `"json:"updateAt"`
}

func FetchAllUser() (Response, error) {
	var obj User

	var arrObj []User

	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM users"

	// fmt.Println("MODEL ERORRRR")

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Name, &obj.Email, &obj.EmailVerifiedAt, &obj.Password, &obj.Role, &obj.RememberToken, &obj.CreatedAt, &obj.UpdatedAt)

		if err != nil {
			return res, err
		}

		arrObj = append(arrObj, obj)
	}

	res.ResponseCode = http.StatusOK
	res.ResponseDesc = "Success"
	res.ResponseTime = utils.DateToStdNow()
	res.Result = arrObj

	return res, nil

}

func StoreUser(name string, email string, password string, role string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT INTO users (name, email, password, role, created_at, updated_at) VALUES ($1, $2, $3, $4, now(), now()) RETURNING id"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	var userID int
	err = stmt.QueryRow(name, email, password, role).Scan(&userID)
	if err != nil {
		log.Fatal(err)
	}

	res.ResponseCode = http.StatusOK
	res.ResponseDesc = "Success"
	res.ResponseTime = utils.DateToStdNow()
	res.Result = map[string]int{
		"result_id": userID,
	}

	return res, nil

}

func UpdateUser(id int, name string, email string, password string, role string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE users set name = $1, email = $2, password = $3, role = $4 where id = $5  RETURNING id"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	var userID int
	err = stmt.QueryRow(name, email, password, role, id).Scan(&userID)
	if err != nil {
		log.Fatal(err)
	}

	res.ResponseCode = http.StatusOK
	res.ResponseDesc = "Success"
	res.ResponseTime = utils.DateToStdNow()
	res.Result = map[string]int{
		"result_id": userID,
	}

	return res, nil
}

func DeleteUser(id int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM users where id = $1  RETURNING id"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	var userID int
	err = stmt.QueryRow(id).Scan(&userID)
	if err != nil {
		log.Fatal(err)
	}

	res.ResponseCode = http.StatusOK
	res.ResponseDesc = "Success"
	res.ResponseTime = utils.DateToStdNow()
	res.Result = map[string]int{
		"result_id": userID,
	}

	return res, nil
}
