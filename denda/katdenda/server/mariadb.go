package server

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	addDenda                = `insert into Denda(ID, Jenis, Jumlah, Status, CreatedBy, CreateOn)values (?,?,?,?,?,?)`
	selectDendaByID         = `select ID, Jenis, Jumlah, Status, CreatedBy, Keterangan from Denda where ID = ?`
	selectDenda             = `select ID, Jenis, Jumlah, Status, CreatedBy, Keterangan from Denda Where Status = '1'`
	selectDendaByKeterangan = `select ID, Jenis, Jumlah, Status, Keterangan from Denda Where Keterangan like ?`
	updateDenda             = `update Denda set ID=?, Jenis=?, Alamat =?, Jumlah=?, Status=?, UpdatedBy=?, UpdateOn=? where ID=?`
	//selectMahasiswaByNama = `select Nim,Nama_Mahasiswa, Status from Mahasiswa where Nama_Mahasiswa=?`
)

//langkah 4
type dbReadWriter struct {
	db *sql.DB
}

func NewDBReadWriter(url string, schema string, user string, password string) ReadWriter {
	schemaURL := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, url, schema)
	db, err := sql.Open("mysql", schemaURL)
	if err != nil {
		panic(err)
	}
	return &dbReadWriter{db: db}
}

//langkah 5
func (rw *dbReadWriter) AddDenda(denda Denda) error {
	fmt.Println("insert")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(addDenda, denda.ID, denda.Jenis, denda.Jumlah, OnAdd, denda.CreatedBy, time.Now())
	//fmt.Println(err)
	if err != nil {
		return err

	}
	return tx.Commit()
}

func (rw *dbReadWriter) ReadDendaByID(id string) (Denda, error) {
	fmt.Println("show by id")
	denda := Denda{ID: id}
	err := rw.db.QueryRow(selectDendaByID, id).Scan(&denda.ID, &denda.Jenis, &denda.Jumlah, &denda.Status, &denda.CreatedBy, &denda.Keterangan)

	if err != nil {
		return Denda{}, err
	}

	return denda, nil
}

func (rw *dbReadWriter) ReadDenda() (Dendas, error) {
	fmt.Println("show all")
	denda := Dendas{}
	rows, _ := rw.db.Query(selectDenda)
	defer rows.Close()
	for rows.Next() {
		var d Denda
		err := rows.Scan(&d.ID, &d.Jenis, &d.Jumlah, &d.Status, &d.Keterangan, &d.CreatedBy)
		if err != nil {
			fmt.Println("error query:", err)
			return denda, err
		}
		denda = append(denda, d)
	}
	//fmt.Println("db nya:", mahasiswa)
	return denda, nil
}

func (rw *dbReadWriter) ReadDendaByKeterangan(keterangan string) (Dendas, error) {
	fmt.Println("show all by ket")
	denda := Dendas{}
	rows, _ := rw.db.Query(selectDendaByKeterangan, keterangan)
	defer rows.Close()
	for rows.Next() {
		var de Denda
		err := rows.Scan(&de.ID, &de.Jenis, &de.Jumlah, &de.Status, &de.Keterangan)
		if err != nil {
			fmt.Println("error query:", err)
			return denda, err
		}
		denda = append(denda, de)
	}
	return denda, nil
}

func (rw *dbReadWriter) UpdateDenda(dnd Denda) error {
	fmt.Println("update successfuly")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(updateDenda, dnd.ID, dnd.Jenis, dnd.Jumlah, dnd.Status, dnd.UpdatedBy, time.Now(), dnd.ID)

	//fmt.Println("name:", cus.Name, cus.CustomerId)
	if err != nil {
		return err
	}

	return tx.Commit()
}

/*
func (rw *dbReadWriter) ReadMahasiswaByNama(nama string) (Mahasiswa, error) {
	mahasiswa := Mahasiswa{NamaMahasiswa: nama}
	err := rw.db.QueryRow(selectMahasiswaByNama, nama).Scan(&mahasiswa.Nim, &mahasiswa.NamaMahasiswa,
		&mahasiswa.Status)

	//fmt.Println("err db", err)
	if err != nil {
		return Mahasiswa{}, err
	}

	return mahasiswa, nil
}
*/
