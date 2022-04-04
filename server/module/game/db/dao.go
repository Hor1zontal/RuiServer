package db

func NewUser(username, sex, email, createTime, IDCard string, classID, majorID, academyID, age int) *DBUser {
	user := &DBUser{
		UserName:   username,
		Sex:        sex,
		Email:      email,
		IDCard:     IDCard,
		ClassID:    classID,
		MajorID:    majorID,
		AcademyID:  academyID,
		Age:        age,
		CreateTime: createTime,
	}
	InsertOne("user", user)
	return user
}
