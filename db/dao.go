package db

import (
	"RuiServer/exception"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"strconv"
	"time"
)

const (
	role_admin   = 1
	role_student = 2
)

func NewUser(username, sex, email, createTime, IDCard string, classID, majorID, academyID, age int) *DBUser {
	id, err := GenId("user")
	if err != nil {
		exception.ExceptionCustom("GenID", exception.DatabaseError, err)
	}
	user := &DBUser{
		ID:         int(id),
		UserName:   username,
		Sex:        sex,
		Email:      email,
		IDCard:     IDCard,
		ClassID:    classID,
		MajorID:    majorID,
		AcademyID:  academyID,
		Age:        age,
		Status:     role_student,
		Enabled:    true,
		CreateTime: ParseTime(createTime),
	}

	InsertOne("user", user)
	return user
}

func GetUser(ID string) *DBUser {
	result := &DBUser{}
	id, _ := strconv.Atoi(ID)
	err := FindOne("user", bson.M{"_id": id}, result)
	if err != nil {
		exception.ExceptionCustom("GenID", exception.DatabaseError, err)
	}
	return result
}

func (user *DBUser) PackLoginRes() gin.H {
	res := gin.H{
		"user": gin.H{
			"avatar":      user.Avatar, //头像
			"status":      user.Status,
			"description": user.Description,
			"id":          user.ID,
			//"role":        user.Role, //1-超级管理员 2--学生
		},
	}
	return res
}

//id：是学工号，客户端传过来是string，在服务端是int自增的，传回客户端的时候记得转成字符串

func GetUsersByQuery(ID, Name, SelectedAcademy, SelectedClass, SelectedMajor, startTime, endTime string, PageNum, PageSize int /*pageNum和pageSize必传*/) (int, []*DBUser) {
	var result []*DBUser

	//var findoptions *options.FindOneOptions
	//findoptions.SetSkip()
	query := make(bson.M)
	if ID != "" {
		id, _ := strconv.Atoi(ID)
		query["_id"] = id
	}
	if Name != "" {
		query["name"] = Name
	}
	if SelectedClass != "" {
		classID, _ := strconv.Atoi(SelectedClass)
		query["class"] = classID
	}
	if SelectedAcademy != "" {
		academyID, _ := strconv.Atoi(SelectedAcademy)
		query["academyID"] = academyID
	}
	if SelectedMajor != "" {
		majorID, _ := strconv.Atoi(SelectedMajor)
		query["majorID"] = majorID
	}
	start_t, end_t := ParseStartEndTime(startTime, endTime)
	AppendTimeQuery(query, "createTime", start_t, end_t)
	total := FindAll("user", query, PageNum, PageSize, &result)
	return total, result
}

func GetUserResults(id int) interface{} {
	results := &DBResults{}
	err := FindOne("result", bson.M{"userID": id}, results)
	if err != nil {
		exception.ExceptionCustom("GetUserResults", exception.DatabaseError, err)
	}
	return results
}

// -----------------------------board----------------------------------------

func NewBoard(content, title, typeName string, published bool) *DBBoard {
	id, err := GenId("board")
	if err != nil {
		exception.ExceptionCustom("GenID", exception.DatabaseError, err)
	}
	board := &DBBoard{
		ID:         int(id),
		Title:      title,
		TypeName:   typeName,
		Content:    content,
		Published:  published,
		CreateTime: time.Now(),
	}

	InsertOne("board", board)
	return board
}

func GetBoardByQuery(pageSize, currPage int, boardType, title string) (int, []DBBoard) {
	result := []DBBoard{}
	query := make(bson.M)
	if boardType != "" {
		query["typeName"] = boardType
	}
	if title != "" {
		query["title"] = title
	}
	total := FindAll("board", query, currPage, pageSize, &result)
	return total, result
}

func ParseTime(timeStr string) time.Time {
	t, err := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	if err != nil {
		exception.ExceptionCustom("ParseTime", exception.TimeParseError, err)
	}
	return t
}

func ParseStartEndTime(time_start_string, time_end_string string) (time.Time, time.Time) {
	time_start, err := time.ParseInLocation("2006-01-02", time_start_string, time.Local)
	if err != nil {
		exception.ExceptionCustom("ParseStartEndTime", exception.TimeParseError, err)
	}
	time_end, err1 := time.ParseInLocation("2006-01-02", time_end_string, time.Local)
	if err1 != nil {
		exception.ExceptionCustom("ParseStartEndTime", exception.TimeParseError, err)
	}
	year_s, month_s, day_s := time_start.Date()
	year_e, month_e, day_e := time_end.Date()
	return time.Date(year_s, month_s, day_s, 0, 0, 0, 0, time_start.Location()),
		time.Date(year_e, month_e, day_e, 23, 59, 59, 59, time_start.Location())
}

func AppendTimeQuery(query bson.M, field string, start time.Time, end time.Time) {
	if query == nil {
		query = make(bson.M)
	}
	query[field] = bson.M{"$gte": start, "$lte": end}
}
