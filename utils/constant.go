package utils

var ClassesMap = map[int]string{
	1: "1班",
	2: "2班",
	3: "3班",
	4: "4班",
	5: "5班",
}

var AcademyMap = map[int]string{
	1: "计算机学院",
	2: "机械与制造学院",
	3: "人文学院",
	4: "土木工程学院",
	5: "电器工程学院",
	6: "力学与工程学院",
	7: "经济管理学院",
}

var AcademyIDMajorMap = map[int]map[int]string{
	1: {1: "软件", 2: "网络"},
	2: {1: "机械与制造学院专业1", 2: "机械与制造学院专业2"},
	3: {1: "人文学院专业1", 2: "人文学院专业2"},
	4: {1: "土木工程学院专业1", 2: "土木工程学院专业2"},
	5: {1: "电器工程学院专业1", 2: "电器工程学院专专业2"},
	6: {1: "力学与工程学院专业1", 2: "力学与工程学院专业2"},
	7: {1: "经济管理学院专业1", 2: "经济管理学院专业2"},
}

var BoardTypeMap = map[string]int{
	"教务通知公告": 1,
	"教务系统公告": 2,
}

const (
	name int = 0
)