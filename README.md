# mbtiles2files
Simple utility to export map tiles from an MBTiles archive to a set of files.

## Building
`go build main.go`

## Usage
`./main path/to/YourMap.mbtiles`

A directory named "tiles" will be created in the current working directory, containing all tiles from your MBTiles archive.