package services

import (
	"context"
	"fmt"
	"klickhr-hris/pkg/db"
	"klickhr-hris/pkg/models"
	"klickhr-hris/pkg/pb"

	"net/http"
)

type Server struct {
	H db.Handler
}

func (s *Server) UploadHRIS(ctx context.Context, req *pb.UploadHRISRequest) (*pb.UploadHRISResponse, error) {
	var HRIS models.HRIS

	HRIS.FileType = int(req.FileType)
	HRIS.ImportType = int(req.ImportType)
	HRIS.RunType = int(req.RunType)
	HRIS.OrganizationLevelId = int(req.OrganizationId)

	fmt.Println("uoload HRIS service")

	/*if result := s.H.DB.Create(&HRIS); result.Error != nil {
		return &pb.UploadHRISResponse{
			Status: http.StatusConflict,
			Error:  result.Error.Error(),
		}, nil
	}

	return &pb.UploadHRISResponse{
		Status: http.StatusCreated,
		Error:  "",
	}, nil */
	return &pb.UploadHRISResponse{
		Status: http.StatusCreated,
		Error:  "aaa",
	}, nil
}
