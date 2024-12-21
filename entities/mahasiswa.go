package entities

import "time"

type Mahasiswa struct {
	Id        int64
	Nama      string
	Uts       float32
	Uas       float32
	Grade     float32
	CreatedAt time.Time
	UpdatedAt time.Time
}
