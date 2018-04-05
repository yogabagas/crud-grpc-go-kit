package endpoint

import (
	"context"

	scv "MiniProject/kampus.id/denda/katdenda/server"

	pb "MiniProject/kampus.id/denda/katdenda/grpc"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	stdopentracing "github.com/opentracing/opentracing-go"
	oldcontext "golang.org/x/net/context"
)

//step 7
type grpcDendaServer struct {
	addDenda              grpctransport.Handler
	readDendaByID         grpctransport.Handler
	readDenda             grpctransport.Handler
	readDendaByKeterangan grpctransport.Handler
	updateDenda           grpctransport.Handler
	//readCustomerByEmail  grpctransport.Handler
}

func NewGRPCDendaServer(endpoints DendaEndpoint, tracer stdopentracing.Tracer,
	logger log.Logger) pb.DendaServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}
	return &grpcDendaServer{
		addDenda: grpctransport.NewServer(endpoints.AddDendaEndpoint,
			decodeAddDendaRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "AddDenda", logger)))...),
		readDendaByID: grpctransport.NewServer(endpoints.ReadDendaByIDEndpoint,
			decodeReadDendaByIDRequest,
			encodeReadDendaByIDResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadDendaByID", logger)))...),
		readDenda: grpctransport.NewServer(endpoints.ReadDendaEndpoint,
			decodeReadDendaRequest,
			encodeReadDendaResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadDenda", logger)))...),
		readDendaByKeterangan: grpctransport.NewServer(endpoints.ReadDendaByKeteranganEndpoint,
			decodeReadDendaByKeteranganRequest,
			encodeReadDendaByKeteranganResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadDendaByKeterangan", logger)))...),
		updateDenda: grpctransport.NewServer(endpoints.UpdateDendaEndpoint,
			decodeUpdateDendaRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "UpdateDenda", logger)))...),
	}
}

func decodeAddDendaRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.AddDendaReq)
	return scv.Denda{ID: req.GetID(), Jenis: req.GetJenis(), Jumlah: req.GetJumlah(), Status: req.GetStatus(), CreatedBy: req.GetCreatedBy()}, nil
}

func encodeEmptyResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func (s *grpcDendaServer) AddDenda(ctx oldcontext.Context, denda *pb.AddDendaReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.addDenda.ServeGRPC(ctx, denda)
	if err != nil {
		return nil, err
	}
	return resp.(*google_protobuf.Empty), nil
}

func decodeReadDendaByIDRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ReadDendaByIDReq)
	return scv.Denda{ID: req.ID}, nil
}

func decodeReadDendaRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, nil
}

func encodeReadDendaByIDResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Denda)
	return &pb.ReadDendaByIDResp{ID: resp.ID, Jenis: resp.Jenis, Jumlah: resp.Jumlah, Status: resp.Status, Keterangan: resp.Keterangan, CreatedBy: resp.CreatedBy}, nil
}

func encodeReadDendaResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Dendas)

	rsp := &pb.ReadDendaResp{}

	for _, v := range resp {
		itm := &pb.ReadDendaByIDResp{
			ID:         v.ID,
			Jenis:      v.Jenis,
			Jumlah:     v.Jumlah,
			Status:     v.Status,
			Keterangan: v.Keterangan,
			CreatedBy:  v.CreatedBy,
		}
		rsp.AllDenda = append(rsp.AllDenda, itm)
	}
	return rsp, nil
}

func (s *grpcDendaServer) ReadDendaByID(ctx oldcontext.Context, id *pb.ReadDendaByIDReq) (*pb.ReadDendaByIDResp, error) {
	_, resp, err := s.readDendaByID.ServeGRPC(ctx, id)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadDendaByIDResp), nil
}

func (s *grpcDendaServer) ReadDenda(ctx oldcontext.Context, e *google_protobuf.Empty) (*pb.ReadDendaResp, error) {
	_, resp, err := s.readDenda.ServeGRPC(ctx, e)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadDendaResp), nil
}

func decodeReadDendaByKeteranganRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ReadDendaByKeteranganReq)
	return scv.Denda{Keterangan: req.Keterangan}, nil
}

func encodeReadDendaByKeteranganResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Dendas)
	rsp := &pb.ReadDendaByKeteranganResp{}

	for _, v := range resp {
		itm := &pb.ReadDendaByIDResp{
			ID:         v.ID,
			Jenis:      v.Jenis,
			Jumlah:     v.Jumlah,
			Status:     v.Status,
			Keterangan: v.Keterangan,
		}
		rsp.KetDenda = append(rsp.KetDenda, itm)
	}

	return rsp, nil
}

func (s *grpcDendaServer) ReadDendaByKeterangan(ctx oldcontext.Context, keterangan *pb.ReadDendaByKeteranganReq) (*pb.ReadDendaByKeteranganResp, error) {
	_, resp, err := s.readDendaByKeterangan.ServeGRPC(ctx, keterangan)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadDendaByKeteranganResp), nil
}

func decodeUpdateDendaRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UpdateDendaReq)
	return scv.Denda{ID: req.ID, Jenis: req.Jenis, Jumlah: req.Jumlah, Status: req.Status, UpdatedBy: req.UpdatedBy}, nil
}

func (s *grpcDendaServer) UpdateDenda(ctx oldcontext.Context, cus *pb.UpdateDendaReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.updateDenda.ServeGRPC(ctx, cus)
	if err != nil {
		return &google_protobuf.Empty{}, err
	}
	return resp.(*google_protobuf.Empty), nil
}
