package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	dbURL := "postgresql://neondb_owner:npg_scMFT46NahkB@ep-lingering-sound-a178ewhu.ap-southeast-1.aws.neon.tech/neondb?sslmode=require"
	db, err := sql.Open("pgx", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("ALTER TABLE tickets ADD COLUMN rejection_reason TEXT DEFAULT '';")
	if err != nil {
		fmt.Printf("Mungkin kolom sudah ada atau ada error lain: %v\n", err)
	} else {
		fmt.Println("Migrasi sukses! Kolom rejection_reason berhasil ditambahkan.")
	}
}
