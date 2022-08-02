package services

import (
	"context"
	"klickhr-hris/pkg/models"
	"klickhr-hris/pkg/pb"
	"net/http"
)

func (s *Server) GetAllHRIS(ctx context.Context, req *pb.UploadHRISRequest) (*pb.UploadHRISResponse, error) {
	var HRIS models.HRIS
	if result := s.H.DB.Find(&HRIS); result.Error != nil {
		return &pb.UploadHRISResponse{
			Status: http.StatusConflict,
			Error:  result.Error.Error(),
		}, nil
	}
	return &pb.UploadHRISResponse{
		Status: http.StatusCreated,
		Error:  "",
	}, nil
}
