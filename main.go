package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("No file specified.")
		os.Exit(0)
	}
	mbtiles := os.Args[1]

	// Check if file exists (if not, attempting to open it will create a new SQLite DB)
	if _, err := os.Stat(mbtiles); err != nil {
		log.Fatal(err)
	}

	// File exists; attempt to open it and ping to make sure it was opened succcessfully
	db, err := sql.Open("sqlite3", "file:"+mbtiles)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("Opened file: " + mbtiles)

	/*
		images.{tile_data,tile_id}
		map.{zoom_level,tile_column,tile_row,tile_id,grid_id}
	*/
	rows, err := db.Query("SELECT zoom_level, tile_column, tile_row, map.tile_id, tile_data FROM map JOIN images ON map.tile_id=images.tile_id")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var zoom, col, row int
		var tile string
		var data []byte
		if err := rows.Scan(&zoom, &col, &row, &tile, &data); err != nil {
			log.Fatal(err)
		}
		log.Printf("[%d:%d @ %d, %s]\n", col, row, zoom, tile)
		tilePath := fmt.Sprintf("tiles/%s/%s", strconv.Itoa(zoom), strconv.Itoa(col))
		os.MkdirAll(tilePath, 0777)

		tileFile := fmt.Sprintf("%s/%s.png", tilePath, strconv.Itoa(row))
		file, err := os.Create(tileFile)
		if err == nil {
			file.Write(data)
			file.Close()

			log.Printf("Wrote file: %s", tileFile)
		} else {
			log.Fatal(err)
		}

	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
