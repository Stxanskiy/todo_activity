package db

import (
	"context"
	"io/ioutil"
	"log"
	"path/filepath"
)

func Migrate() {
	files, err := ioutil.ReadDir("pkg/db/migrations")
	if err != nil {
		log.Fatalf("Failed to read migration files: %v", err)
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".sql" {
			path := filepath.Join("pkg/db/migrations", file.Name())
			content, err := ioutil.ReadFile(path)
			if err != nil {
				log.Fatalf("Failed to read migration file %s: %v", file.Name(), err)
			}

			_, err = DB.Exec(context.Background(), string(content))
			if err != nil {
				log.Fatalf("Failed to execute migration %s: %v", file.Name(), err)
			}

			log.Printf("Successfully executed migration %s", file.Name())
		}
	}
}
