package db

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
)

type Db struct {
	db *sql.DB
}

func NewDb(filePath string) (*Db, error) {
    db, err := InitDb(filePath)
	if err != nil {
		return nil, err
	}

	return &Db{db: db}, nil
}


func InitDb(filePath string) (*sql.DB, error) {
	appPath, err := os.Executable()
	fmt.Println(appPath)
	if err != nil {
		logrus.Fatal(err)
	}
	dbFile := filepath.Join(filepath.Dir(appPath), "scheduler.db")
	_, err = os.Stat(dbFile)

	var install bool
	if err != nil {
		install = true
	}

	

    db, err := sql.Open("sqlite3", filePath)
    if err != nil {
        return nil, err
    }

    err = db.Ping()
    if err != nil {
        return nil, err
    }
	if install {
        
        createTableQuery := `
    	CREATE TABLE IF NOT EXISTS scheduler (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        date TEXT,
        title TEXT,
        comment TEXT,
        repeat TEXT(128)  
    );
		`

        _, err := db.Exec(createTableQuery)
        if err != nil {
            db.Close()
            return nil, err
        }

        createIndexQuery := `
			CREATE INDEX IF NOT EXISTS idx_date ON scheduler(date);
        `
        _, err = db.Exec(createIndexQuery)
        if err != nil {
            db.Close()
            return nil, err
        }
    }


    return db, nil
}