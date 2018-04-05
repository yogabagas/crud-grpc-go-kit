package server

import (
	"context"
)

//langkah 6
type denda struct {
	writer ReadWriter
}

func NewDenda(writer ReadWriter) DendaService {
	return &denda{writer: writer}
}

//Methode pada interface MahasiswaService di service.go
func (c *denda) AddDendaService(ctx context.Context, denda Denda) error {
	//fmt.Println("mahasiswa")
	err := c.writer.AddDenda(denda)
	if err != nil {
		return err
	}

	return nil
}

func (c *denda) ReadDendaByIDService(ctx context.Context, id string) (Denda, error) {
	dnd, err := c.writer.ReadDendaByID(id)
	//fmt.Println(mhs)
	if err != nil {
		return dnd, err
	}
	return dnd, nil
}

func (c *denda) ReadDendaService(ctx context.Context) (Dendas, error) {
	dnd, err := c.writer.ReadDenda()
	//fmt.Println("mahasiswa", mhs)
	if err != nil {
		return dnd, err
	}
	return dnd, nil
}

func (c *denda) ReadDendaByKeteranganService(ctx context.Context, keterangan string) (Dendas, error) {
	dnd, err := c.writer.ReadDendaByKeterangan(keterangan)
	if err != nil {
		return dnd, err
	}
	return dnd, nil
}

func (c *denda) UpdateDendaService(ctx context.Context, dnd Denda) error {
	err := c.writer.UpdateDenda(dnd)
	if err != nil {
		return err
	}
	return nil
}

/*
func (c *mahasiswa) ReadMahasiswaByNamaService(ctx context.Context, nama string) (Mahasiswa, error) {
	mhs, err := c.writer.ReadMahasiswaByNama(nama)
	//fmt.Println("mahasiswa:", mhs)
	if err != nil {
		return mhs, err
	}
	return mhs, nil
}
*/
