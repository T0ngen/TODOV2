package db

import "todov2/pkg/common/db/models"




func (d *Db) GetTasksByDate(date string) ([]models.Task, error) {

	query := "SELECT * FROM scheduler WHERE date = ?"
	
	rows, err := d.db.Query(query, date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()


	var tasks []models.Task

	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.Id, &task.Date, &task.Title, &task.Comment, &task.Repeat); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return tasks, nil
}