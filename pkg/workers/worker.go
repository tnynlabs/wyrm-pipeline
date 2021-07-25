package workers

import (
	"context"
	"encoding/json"
	"log"
	"net"

	"github.com/tnynlabs/wyrm-pipeline/pkg/graph"
	g "github.com/tnynlabs/wyrm-pipeline/pkg/graph"
	"github.com/tnynlabs/wyrm-pipeline/pkg/graph/parser"
	pb "github.com/tnynlabs/wyrm-pipeline/pkg/workers/protobuf"
	"github.com/tnynlabs/wyrm/pkg/pipelines"

	"google.golang.org/grpc"
)

type pipelineWorkerServer struct {
	pb.UnimplementedPipelineWorkerServer

	pipelineService pipelines.Service
}

func NewServer(pipelineService pipelines.Service) pb.PipelineWorkerServer {
	return &pipelineWorkerServer{
		pipelineService: pipelineService,
	}
}

func RunServer(address string, server pb.PipelineWorkerServer) error {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	// TODO: Check needed server options
	grpcServer := grpc.NewServer()
	pb.RegisterPipelineWorkerServer(grpcServer, server)

	err = grpcServer.Serve(listener)
	if err != nil {
		return err
	}

	return nil
}

func (s *pipelineWorkerServer) RunPipeline(ctx context.Context, req *pb.PipelineRequest) (*pb.PipelineResponse, error) {
	pipeline, err := s.pipelineService.GetByID(req.PipelineId)
	if err != nil {
		log.Printf("Can't find pipeline with ID (%d): %v\n", req.PipelineId, err.Error())
		return nil, err
	}

	graphData := graph.GraphData{}
	if err := json.Unmarshal([]byte(pipeline.Data), &graphData); err != nil {
		log.Println("Parsing config file", err.Error())
		return nil, err
	}

	graph, err := parser.BuildGraph(graphData)
	if err != nil {
		log.Println("Can't build pipeline graph", err.Error())
		return nil, err
	}

	msg := make(g.Msg)
	msg["payload"] = req.Payload

	log.Println("Running pipeline....")
	err = graph.Entrypoint.Run(msg)
	if err != nil {
		return nil, err
	}

	return &pb.PipelineResponse{}, nil
}
