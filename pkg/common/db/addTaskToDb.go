package db

import (
	
	"todov2/pkg/api/mappers"
)


func (d *Db) AddTaskToDb(task mappers.NewTask) (int, error) {
	query := "INSERT INTO scheduler (date, title, comment, repeat) VALUES ($1, $2, $3, $4)"

	res, err := d.db.Exec(query, task.Date, task.Title, task.Comment, task.Repeat)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}