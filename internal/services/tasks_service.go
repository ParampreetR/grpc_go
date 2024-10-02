package services

import (
	"context"
	"log"

	pb "github.com/parampreetr/grpc_go/api/gen"
	"github.com/parampreetr/grpc_go/internal/models"
	"gorm.io/gorm"
)

type GRPCServer struct {
	pb.UnimplementedTaskServiceServer
	DB *gorm.DB
}

func (s *GRPCServer) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.Task, error) {
	task := &models.Task{
		Title:       req.GetTitle(),
		Description: req.GetDesc(),
		IsDone:      req.GetIsDone(),
	}

	result := models.Task{}

	if err := s.DB.WithContext(ctx).Create(task).First(&result, task).Error; err != nil {
		return nil, err
	}

	log.Println("Create Task Successful")
	log.Printf("%+v\n", task)

	// s.NextID++

	return &pb.Task{
		Id:     result.ID,
		Title:  result.Title,
		Desc:   result.Description,
		IsDone: result.IsDone,
	}, nil
}

func (s *GRPCServer) DeleteTask(ctx context.Context, req *pb.DeleteTaskRequest) (*pb.Task, error) {
	query := &models.Task{
		ID: req.GetId(),
	}
	result := models.Task{}

	DBCtx := s.DB.WithContext(ctx)

	if err := DBCtx.First(&result).Error; err != nil {
		return nil, err
	}

	if err := DBCtx.Delete(query).Error; err != nil {
		return nil, err
	}

	log.Println("Delete Task Successful")
	log.Printf("%+v\n", result)

	return &pb.Task{
		Id:     result.ID,
		Title:  result.Title,
		Desc:   result.Description,
		IsDone: result.IsDone,
	}, nil
}

func (s *GRPCServer) GetTask(ctx context.Context, req *pb.GetTaskRequest) (*pb.Task, error) {
	result := models.Task{}

	if err := s.DB.WithContext(ctx).First(&result, req.GetId()).Error; err != nil {
		return nil, err
	}

	log.Println("Retrived Task Successfully")
	log.Printf("%+v\n", result)

	return &pb.Task{
		Id:     result.ID,
		Title:  result.Title,
		Desc:   result.Description,
		IsDone: result.IsDone,
	}, nil
}
