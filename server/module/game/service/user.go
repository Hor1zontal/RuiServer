package service

import "RuiServer/server/module/game/db"

func UploadUser(username, sex, email, createTime, IDCard string, ClassID, MajorID, AcademyID, Age int) {
	db.NewUser(username, sex, email, createTime, IDCard, ClassID, MajorID, AcademyID, Age)
}
