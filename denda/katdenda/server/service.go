package server

import "context"

type Status int32

//step 2
const (
	//ServiceID is dispatch service ID
	ServiceID        = "Campus.denda.id"
	OnAdd     Status = 1
)

type Denda struct {
	ID         string
	Jenis      string
	Jumlah     string
	Status     int32
	Keterangan string
	CreatedBy  string
	UpdatedBy  string
}
type Dendas []Denda

/*type Location struct {
	customerID   int64
	label        []int32
	locationType []int32
	name         []string
	street       string
	village      string
	district     string
	city         string
	province     string
	latitude     float64
	longitude    float64
}*/

type ReadWriter interface {
	AddDenda(Denda) error
	ReadDendaByID(string) (Denda, error)
	ReadDenda() (Dendas, error)
	ReadDendaByKeterangan(string) (Dendas, error)
	UpdateDenda(Denda) error
	//ReadMahasiswaByNama(string) (Mahasiswa, error)
}

type DendaService interface {
	AddDendaService(context.Context, Denda) error
	ReadDendaByIDService(context.Context, string) (Denda, error)
	ReadDendaService(context.Context) (Dendas, error)
	ReadDendaByKeteranganService(context.Context, string) (Dendas, error)
	UpdateDendaService(context.Context, Denda) error
	//ReadMahasiswaByNamaService(context.Context, string) (Mahasiswa, error)
}
