package services

import (
	"context"
	"fmt"
	"klickhr-hris/pkg/constants"
	"klickhr-hris/pkg/db"
	"klickhr-hris/pkg/models"
	"klickhr-hris/pkg/pb"
	"klickhr-hris/pkg/utils"
	"net/http"
	"time"
)

type Server struct {
	H db.Handler
}

func (s *Server) UploadHRIS(ctx context.Context, req *pb.UploadHRISRequest) (*pb.UploadHRISResponse, error) {
	//save file
	err := utils.SaveFile(req.FileName, req.FileBytes)
	if err != nil {
		return &pb.UploadHRISResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}
	//file validations
	err = utils.ValidateFile("./" + req.FileName)
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
	HRIS.FileUrl = "./" + req.FileName
	HRIS.FileId = req.FileName
	HRIS.UserID = 1 // get user by token
	HRIS.ImportDate = time.Now()
	HRIS.Status = constants.HRIS_STATUS_ON_HOLD
	//create row
	if result := s.H.DB.Create(&HRIS); result.Error != nil {
		return &pb.UploadHRISResponse{
			Status: http.StatusConflict,
			Error:  result.Error.Error(),
		}, nil
	}

	go executeHRIS(HRIS)

	return &pb.UploadHRISResponse{
		Status: http.StatusCreated,
		Error:  "",
	}, nil
}

func executeHRIS(HRIS models.HRIS) {
	switch HRIS.RunType {
	case constants.HRIS_RUN_TYPE_APPEND:
		fmt.Println("APPEND")
	case constants.HRIS_RUN_TYPE_FULL:
		fmt.Println("FULL")
	case constants.HRIS_RUN_TYPE_VALIDATE:
		fmt.Println("VALIDATE")
	}
}

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
