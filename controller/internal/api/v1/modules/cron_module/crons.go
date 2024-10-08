package cron_module

import (
	"fmt"
	"go-http/pkg/settings/database"
)

type Cron struct {
	Id        string `json:"cron_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Taskname  string `json:"task_name"`
	At        string `json:"at"`
}

func SelectCrons() (any, error) {
	query := "SELECT * FROM crons"

	rows, err := database.GetDB().Query(query)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	var crons []Cron
	for rows.Next() {
		var cron Cron
		err = rows.Scan(&cron.Id, &cron.CreatedAt, &cron.UpdatedAt, &cron.Taskname, &cron.At)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		crons = append(crons, cron)
	}
	err = rows.Err()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for _, cron := range crons {
		fmt.Println(cron)
	}

	return crons, nil
}

func SelectCron(id string) (any, error) {
	query := fmt.Sprintf("SELECT * FROM crons WHERE cron_id='%s'", id)

	var cron Cron
	row := database.GetDB().QueryRow(query)
	err := row.Scan(&cron.Id, &cron.CreatedAt, &cron.UpdatedAt, &cron.Taskname, &cron.At)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return cron, nil
}

func InsertCron(data CreateCronBody) (any, error) {
	query := fmt.Sprintf("INSERT INTO crons (task_name,at) VALUES ('%s','%s')", data.Taskname, data.At)

	res, err := database.GetDB().Exec(query)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return res, nil
}

func DelCron(id string) (any, error) {
	query := fmt.Sprintf("DELETE FROM crons WHERE cron_id='%s'", id)

	res, err := database.GetDB().Exec(query)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return res, nil
}

func UpdCron(id string, data UpdateCronBody) (any, error) {
	query := fmt.Sprintf("UPDATE crons SET task_name='%s', at='%s'  WHERE cron_id='%s'", data.Taskname, data.At, id)

	res, err := database.GetDB().Exec(query)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return res, nil
}
