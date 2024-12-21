package mahasiswamodel

import (
	"golang-mahasiswa/config"
	"golang-mahasiswa/entities"
)

func GetAll() []entities.Mahasiswa {
	rows, err := config.DB.Query(`SELECT * FROM mahasiswa`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var mahasiswa []entities.Mahasiswa

	for rows.Next() {
		var siswa entities.Mahasiswa
		if err := rows.Scan(&siswa.Id, &siswa.Nama, &siswa.Uts, &siswa.Uas, &siswa.Grade, &siswa.CreatedAt, &siswa.UpdatedAt); err != nil {
			panic(err)
		}

		mahasiswa = append(mahasiswa, siswa)
	}

	return mahasiswa
}

func Create(mahasiswa entities.Mahasiswa) bool {
	_, err := config.DB.Exec(`
	INSERT INTO mahasiswa (id, nama, uts, uas, grade, createdAt, updatedAt)
	VALUE (?, ?, ?, ?, ?, ?, ?)`,
		mahasiswa.Id,
		mahasiswa.Nama,
		mahasiswa.Uts,
		mahasiswa.Uas,
		mahasiswa.Grade,
		mahasiswa.CreatedAt,
		mahasiswa.UpdatedAt,
	)
	if err != nil {
		panic(err)
	}

	return true
}
