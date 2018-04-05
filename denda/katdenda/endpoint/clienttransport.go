package endpoint

import (
	"context"
	"time"

	svc "MiniProject/kampus.id/denda/katdenda/server"

	pb "MiniProject/kampus.id/denda/katdenda/grpc"

	util "MiniProject/kampus.id/denda/util/grpc"
	disc "MiniProject/kampus.id/denda/util/microservice"

	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/lb"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	stdopentracing "github.com/opentracing/opentracing-go"
	"github.com/sony/gobreaker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

//step 8
const (
	grpcName = "grpc.DendaService"
)

func NewGRPCDendaClient(nodes []string, creds credentials.TransportCredentials, option util.ClientOption,
	tracer stdopentracing.Tracer, logger log.Logger) (svc.DendaService, error) {

	instancer, err := disc.ServiceDiscovery(nodes, svc.ServiceID, logger)
	if err != nil {
		return nil, err
	}

	retryMax := option.Retry
	retryTimeout := option.RetryTimeout
	timeout := option.Timeout

	var addDendaEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientAddDendaEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		addDendaEp = retry
	}

	var readDendaByIDEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadDendaByIDEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readDendaByIDEp = retry
	}

	var readDendaEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadDendaEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readDendaEp = retry
	}

	var readDendaByKeteranganEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadDendaByKeteranganEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readDendaByKeteranganEp = retry
	}

	var updateDendaEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientUpdateDenda, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		updateDendaEp = retry
	}
	/*
		var readCustomerByEmailEp endpoint.Endpoint
		{
			factory := util.EndpointFactory(makeClientReadCustomerByEmail, creds, timeout, tracer, logger)
			endpointer := sd.NewEndpointer(instancer, factory, logger)
			balancer := lb.NewRoundRobin(endpointer)
			retry := lb.Retry(retryMax, retryTimeout, balancer)
			readCustomerByEmailEp = retry
		}
	*/
	return DendaEndpoint{AddDendaEndpoint: addDendaEp, ReadDendaByIDEndpoint: readDendaByIDEp,
		ReadDendaEndpoint: readDendaEp, ReadDendaByKeteranganEndpoint: readDendaByKeteranganEp, UpdateDendaEndpoint: updateDendaEp}, nil
}

func encodeAddDendaRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Denda)
	return &pb.AddDendaReq{
		ID:        req.ID,
		Jenis:     req.Jenis,
		Jumlah:    req.Jumlah,
		Status:    req.Status,
		CreatedBy: req.CreatedBy,
	}, nil
}

func encodeReadDendaByIDRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Denda)
	return &pb.ReadDendaByIDReq{ID: req.ID}, nil
}

func encodeReadDendaRequest(_ context.Context, request interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeReadDendaByKeteranganRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Denda)
	return &pb.ReadDendaByKeteranganReq{Keterangan: req.Keterangan}, nil
}

func encodeUpdateDendaRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Denda)
	return &pb.UpdateDendaReq{
		ID:        req.ID,
		Jenis:     req.Jenis,
		Jumlah:    req.Jumlah,
		Status:    req.Status,
		UpdatedBy: req.UpdatedBy,
	}, nil
}

/*
func encodeReadCustomerByEmailRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Customer)
	return &pb.ReadCustomerByEmailReq{Email: req.Email}, nil
}
*/

func decodeDendaResponse(_ context.Context, response interface{}) (interface{}, error) {
	return nil, nil
}

func decodeReadDendaByIDResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadDendaByIDResp)
	return svc.Denda{
		ID:         resp.ID,
		Jenis:      resp.Jenis,
		Jumlah:     resp.Jumlah,
		Keterangan: resp.Keterangan,
		Status:     resp.Status,
		CreatedBy:  resp.CreatedBy,
	}, nil
}

func decodeReadDendaResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadDendaResp)
	var rsp svc.Dendas

	for _, v := range resp.AllDenda {
		itm := svc.Denda{
			ID:         v.ID,
			Jenis:      v.Jenis,
			Jumlah:     v.Jumlah,
			Keterangan: v.Keterangan,
			Status:     v.Status,
			CreatedBy:  v.CreatedBy,
		}
		rsp = append(rsp, itm)
	}
	return rsp, nil
}

func decodeReadDendaByKeteranganResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadDendaByKeteranganResp)
	var rsp svc.Dendas

	for _, v := range resp.KetDenda {
		itm := svc.Denda{
			ID:         v.ID,
			Jenis:      v.Jenis,
			Jumlah:     v.Jumlah,
			Keterangan: v.Keterangan,
			Status:     v.Status,
		}
		rsp = append(rsp, itm)
	}
	return rsp, nil
}

/*
func decodeReadCustomerByEmailRespones(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadCustomerByEmailResp)
	return svc.Customer{
		CustomerId:    resp.CustomerId,
		Name:          resp.Name,
		CustomerType:  resp.CustomerType,
		Mobile:        resp.Mobile,
		Email:         resp.Email,
		Gender:        resp.Gender,
		CallbackPhone: resp.CallbackPhone,
		Status:        resp.Status,
	}, nil
}
*/
func makeClientAddDendaEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn,
		grpcName,
		"AddDenda",
		encodeAddDendaRequest,
		decodeDendaResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "AddDenda")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "AddDenda",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadDendaByIDEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadDendaByID",
		encodeReadDendaByIDRequest,
		decodeReadDendaByIDResponse,
		pb.ReadDendaByIDResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadDendaByID")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadDendaByID",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadDendaEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadDenda",
		encodeReadDendaRequest,
		decodeReadDendaResponse,
		pb.ReadDendaResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadDenda")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadDenda",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadDendaByKeteranganEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadDendaByKeterangan",
		encodeReadDendaByKeteranganRequest,
		decodeReadDendaByKeteranganResponse,
		pb.ReadDendaByKeteranganResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadDendaByKeterangan")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadDendaByKeterangan",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientUpdateDenda(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {
	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"UpdateDenda",
		encodeUpdateDendaRequest,
		decodeDendaResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "UpdateDenda")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "UpdateDenda",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

/*

func makeClientReadCustomerByEmail(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {
	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadCustomerByEmail",
		encodeReadCustomerByEmailRequest,
		decodeReadCustomerByEmailRespones,
		pb.ReadCustomerByEmailResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadCustomerByEmail")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadCustomerByEmail",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}
*/
