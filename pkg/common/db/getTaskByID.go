package db

import (
	"database/sql"
	"fmt"
	"todov2/pkg/common/db/models"
)




func (d *Db) GetTaskByID(id string) (models.Task, error) {
    query := "SELECT * FROM scheduler WHERE id = ?"

    
    row := d.db.QueryRow(query, id)

    var task models.Task

   
    if err := row.Scan(&task.Id, &task.Date, &task.Title, &task.Comment, &task.Repeat); err != nil {
        if err == sql.ErrNoRows {
           
            return models.Task{}, fmt.Errorf("tusk with id %s not found", id)
        }
        
        return models.Task{}, err
    }

    
    return task, nil
}