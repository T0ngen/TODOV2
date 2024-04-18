package db

import (
	"errors"
	
	
	"todov2/pkg/common/db/models"
)



func (d *Db) UpdateTaskById(task models.Task) error {
    
    if task.Id == "" {
        return errors.New("ID задачи не может быть пустым")
    }

	
    query := "UPDATE scheduler SET date = $1, title = $2, comment = $3, repeat = $4 WHERE id = $5"
    _, err := d.db.Exec(query, task.Date, task.Title, task.Comment, task.Repeat, task.Id)
    if err != nil {
        return err
    }

	// id, err := res.LastInsertId()
	// fmt.Println(id)
	// if err != nil {
	// 	return err
	// }
	// if id == 0 {
	// 	return errors.New("задача не обновлена")
	// }

    return nil
}