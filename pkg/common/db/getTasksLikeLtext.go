package db

import "todov2/pkg/common/db/models"




func (d *Db) GetTasksLikeLtext(text string) ([]models.Task, error) {

	query := "SELECT * FROM scheduler WHERE title LIKE ? OR comment LIKE ? ORDER BY date DESC"
	searchText := "%" + text + "%"

	rows, err := d.db.Query(query, searchText, searchText)
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