package db




func (d *Db) DeleteTaskById(taskId string) error {
	query := "DELETE FROM scheduler WHERE id = $1"

	_, err := d.db.Exec(query, taskId)
	if err != nil {
		return err
	}

	return nil
}