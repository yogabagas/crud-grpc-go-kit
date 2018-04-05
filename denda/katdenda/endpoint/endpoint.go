package endpoint

import (
	"context"

	svc "MiniProject/kampus.id/denda/katdenda/server"
	kit "github.com/go-kit/kit/endpoint"
)

//step 8
type DendaEndpoint struct {
	AddDendaEndpoint              kit.Endpoint
	ReadDendaByIDEndpoint         kit.Endpoint
	ReadDendaEndpoint             kit.Endpoint
	ReadDendaByKeteranganEndpoint kit.Endpoint
	UpdateDendaEndpoint           kit.Endpoint
	//ReadCustomerByEmailEndpoint  kit.Endpoint
}

func NewDendaEndpoint(service svc.DendaService) DendaEndpoint {
	addDendaEp := makeAddDendaEndpoint(service)
	readDendaByIDEp := makeReadDendaByIDEndpoint(service)
	readDendaEp := makeReadDendaEndpoint(service)
	readDendaByKeteranganEp := makeReadDendaByKeteranganEndpoint(service)
	updateDendaEp := makeUpdateDendaEndpoint(service)
	//readCustomerByEmailEp := makeReadCustomerByEmailEndpoint(service)
	return DendaEndpoint{AddDendaEndpoint: addDendaEp,
		ReadDendaByIDEndpoint:         readDendaByIDEp,
		ReadDendaEndpoint:             readDendaEp,
		ReadDendaByKeteranganEndpoint: readDendaByKeteranganEp,
		UpdateDendaEndpoint:           updateDendaEp,
	}
	//ReadCustomerByEmailEndpoint:  readCustomerByEmailEp,

}

func makeAddDendaEndpoint(service svc.DendaService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Denda)
		err := service.AddDendaService(ctx, req)
		return nil, err
	}
}

func makeReadDendaByIDEndpoint(service svc.DendaService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Denda)
		result, err := service.ReadDendaByIDService(ctx, req.ID)
		/*return svc.Customer{CustomerId: result.CustomerId, Name: result.Name,
		CustomerType: result.CustomerType, Mobile: result.Mobile, Email: result.Email,
		Gender: result.Gender, CallbackPhone: result.CallbackPhone, Status: result.Status}, err*/
		return result, err
	}
}

func makeReadDendaEndpoint(service svc.DendaService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		result, err := service.ReadDendaService(ctx)
		return result, err
	}
}

func makeReadDendaByKeteranganEndpoint(service svc.DendaService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Denda)
		result, err := service.ReadDendaByKeteranganService(ctx, req.Keterangan)
		return result, err
	}
}

func makeUpdateDendaEndpoint(service svc.DendaService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Denda)
		err := service.UpdateDendaService(ctx, req)
		return nil, err
	}
}

/*
func makeReadCustomerByEmailEndpoint(service svc.CustomerService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Customer)
		result, err := service.ReadCustomerByEmailService(ctx, req.Email)
		return result, err
	}
}
*/
