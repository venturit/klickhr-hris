package models

import "time"

type HRIS struct {
	Id                  int       `json:"id" gorm:"primaryKey"`
	FileType            int       `json:"file_type"`
	RunType             int       `json:"run_type"`
	ImportType          int       `json:"import_type"`
	ImportDate          time.Time `json:"import_date"`
	OrganizationLevelId int       `json:"organization_level_id"`
	FileId              string    `json:"file_id"`
	FileUrl             string    `json:"file_url"`
	Errors              bool      `json:"errors"`
	UserID              int       `json:"user_id"`
	Status              int       `json:"status"`
}
