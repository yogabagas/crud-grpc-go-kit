package endpoint

import (
	"context"
	"fmt"

	sv "MiniProject/kampus.id/denda/katdenda/server"
)

//step 9
func (me DendaEndpoint) AddDendaService(ctx context.Context, denda sv.Denda) error {
	_, err := me.AddDendaEndpoint(ctx, denda)
	return err
}

func (me DendaEndpoint) ReadDendaByIDService(ctx context.Context, id string) (sv.Denda, error) {
	req := sv.Denda{ID: id}
	fmt.Println(req)
	resp, err := me.ReadDendaByIDEndpoint(ctx, req)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	cus := resp.(sv.Denda)
	return cus, err
}

func (me DendaEndpoint) ReadDendaService(ctx context.Context) (sv.Dendas, error) {
	resp, err := me.ReadDendaEndpoint(ctx, nil)
	fmt.Println("me resp", resp)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	return resp.(sv.Dendas), err
}

func (me DendaEndpoint) ReadDendaByKeteranganService(ctx context.Context, keterangan string) (sv.Dendas, error) {
	req := sv.Denda{Keterangan: keterangan}
	fmt.Println(req)
	resp, err := me.ReadDendaByKeteranganEndpoint(ctx, req)
	fmt.Println("me resp", resp)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	cus := resp.(sv.Dendas)
	return cus, err
}
func (me DendaEndpoint) UpdateDendaService(ctx context.Context, cus sv.Denda) error {
	_, err := me.UpdateDendaEndpoint(ctx, cus)
	if err != nil {
		fmt.Println("error pada endpoint:", err)
	}
	return err
}

/*
func (ce CustomerEndpoint) ReadCustomerByEmailService(ctx context.Context, email string) (sv.Customer, error) {
	req := sv.Customer{Email: email}
	resp, err := ce.ReadCustomerByEmailEndpoint(ctx, req)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	cus := resp.(sv.Customer)
	return cus, err
}
*/
