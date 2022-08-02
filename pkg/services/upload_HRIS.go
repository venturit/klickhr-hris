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

type OrganizationRow struct {
	Level_1_Name string `json:"Level_1_Name"`
	Level_1_Code string `json:"Level_1_code"`
	Level_2_Name string `json:"Level_2_Name"`
	Level_2_Code string `json:"Level_2_code"`
	Level_3_Name string `json:"Level_3_Name"`
	Level_3_Code string `json:"Level_3_code"`
	Level_4_Name string `json:"Level_4_Name"`
	Level_4_Code string `json:"Level_4_code"`
	Job_Status   string `json:"jon_status"`
}

type ErrorListItem struct {
	Record int    `json:"record"`
	Error  string `json:"error"`
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
	err = utils.ValidateFile("./"+req.FileName, int(req.ImportType))
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

	go executeHRIS(HRIS, HRIS.FileUrl)

	return &pb.UploadHRISResponse{
		Status: http.StatusCreated,
		Error:  "",
	}, nil
}

func executeHRIS(HRIS models.HRIS, url string) {
	switch HRIS.ImportType {
	case constants.HRIS_IMPORT_TYPE_APPEND:
		fmt.Println("APPEND")
	case constants.HRIS_IMPORT_TYPE_FULL:
		fmt.Println("FULL")
	case constants.HRIS_IMPORT_TYPE_VALIDATE:
		errorList, err := validateHRIS(url, HRIS.FileType)
		if err != nil {
			fmt.Println(err)
		}
		//PARA EL JHOAN DEL FUTURO, LA PRUEBA DE DUPLICADOS LE FALTA UN DUPLICADO FIN, ARREGLELO, COMPRA COMIDA
		fmt.Println(errorList)
	}
}

func validateHRIS(url string, file_type int) ([]ErrorListItem, error) {
	fmt.Println("Validate")
	data, err := utils.ReadCSVData(url)
	if err != nil {
		return nil, err
	}
	var errorList []ErrorListItem
	var duplicateIndex []int
	var duplicate bool
	//ORGANIZATION FILE
	if file_type == constants.HRIS_FILE_TYPE_ORGANIZATION {
		var organizationList = parseToOrganizationList(data)
		for index, organization := range organizationList {
			//duplicate validation
			if isDuplicateIndex(index, duplicateIndex) {
				continue
			}
			duplicateIndex, duplicate = duplicateValidation(organization, organizationList, index, duplicateIndex)
			if duplicate {
				errorList = append(errorList, ErrorListItem{index, "duplicate row!"})
				for _, i := range duplicateIndex {
					errorList = append(errorList, ErrorListItem{i, "duplicate row!"})
				}
				continue
			}
			//fields validation
			if !(organization.Job_Status == "Active" || organization.Job_Status == "Inactive") {
				errorList = append(errorList, ErrorListItem{index, "Only values allowed are “Active” or “Inactive”"})
			}

		}
	} else {
		//EMPLOYEE FILE
		for index, row := range data {
			fmt.Println(index)
			fmt.Println(row)
		}
	}
	return errorList, nil
}

func parseToOrganizationList(data [][]string) []OrganizationRow {
	var orgRowList []OrganizationRow
	var orgRow OrganizationRow
	for _, row := range data {
		orgRow.Level_1_Name = row[0]
		orgRow.Level_1_Code = row[1]
		orgRow.Level_2_Name = row[2]
		orgRow.Level_2_Code = row[3]
		orgRow.Level_3_Name = row[4]
		orgRow.Level_3_Code = row[5]
		orgRow.Level_4_Name = row[6]
		orgRow.Level_4_Code = row[7]
		orgRow.Job_Status = row[8]
		orgRowList = append(orgRowList, orgRow)
	}
	return orgRowList
}

func duplicateValidation(item OrganizationRow, list []OrganizationRow, index int, indexList []int) ([]int, bool) {
	var found bool
	for i := index + 1; i < len(list); i++ {
		if item == list[i] {
			indexList = append(indexList, i)
			found = true
		}
	}
	fmt.Println("duplicateValidation")
	fmt.Println(indexList)
	return indexList, found
}

func isDuplicateIndex(index int, indexList []int) bool {
	for i := 0; i < len(indexList); i++ {
		if index == indexList[i] {
			return true
		}
	}
	return false
}
