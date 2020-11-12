package main

import (
	"io"
)

type FileSystemStore struct {
	database io.ReadSeeker
}

func (f *FileSystemStore) GetLeague() []Player {
	league, _ := NewLeague(f.database)
	return league
}