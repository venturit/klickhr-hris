package services

import (
	"context"
	"fmt"
	"klickhr-hris/pkg/db"
	"klickhr-hris/pkg/models"
	"klickhr-hris/pkg/pb"
	"klickhr-hris/pkg/utils"
	"net/http"
)

type Server struct {
	H db.Handler
}

func (s *Server) UploadHRIS(ctx context.Context, req *pb.UploadHRISRequest) (*pb.UploadHRISResponse, error) {
	err := utils.SaveFile(req.FileName, req.FileBytes)
	if err != nil {
		return &pb.UploadHRISResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

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
