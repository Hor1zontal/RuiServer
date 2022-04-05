package service

import "RuiServer/server/module/game/db"

func UploadUser(username, sex, email, createTime, IDCard string, ClassID, MajorID, AcademyID, Age int) {
	db.NewUser(username, sex, email, createTime, IDCard, ClassID, MajorID, AcademyID, Age)
}

func GetUser(IDCard string) {
	db.GetUser(IDCard)
}

func GetUsersByQuery(ID, Name, SelectedAcademy, SelectedClass, SelectedMajor, startTime, endTime string, PageNum, PageSize int) (int, []*db.DBUser) {
	return db.GetUsersByQuery(ID, Name, SelectedAcademy, SelectedClass, SelectedMajor, startTime, endTime, PageNum, PageSize)
}
