# mbtiles2files
Simple utility to export map tiles from an MBTiles archive to a set of files.

## Preparation
MBTiles files are SQLite3 databases, so `mbtiles2files` requires an SQLite database driver compatible with `database/sql`.  As this is the only dependency, no dependency manager has been used; you can install this dependency manually by running the following command:

`go get github.com/mattn/go-sqlite3`

## Building
`go build main.go`

## Usage
`./main path/to/YourMap.mbtiles`

A directory named "tiles" will be created in the current working directory, containing all tiles from your MBTiles archive.