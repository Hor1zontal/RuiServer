package service

import (
	"RuiServer/db"
)

type resBoard struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	CreateTime string `json:"createTime"`
	TypeName   string `json:"typeName"`
	Published  bool   `json:"published"`
}

func GetBoardByType(pageSize, currPage int, typeName string) (int, interface{}) {
	var res []resBoard
	total, boards := db.GetBoardByQuery(pageSize, currPage, typeName, "")
	if total != 0 {
		res = packResBoards(boards)
	}
	return total, res
}

func packResBoards(boards []db.DBBoard) []resBoard {
	var res []resBoard
	for _, v := range boards {
		tmp := resBoard{v.ID, v.Title, v.Content,
			v.CreateTime.Local().Format("2006-01-02"),
			v.TypeName, v.Published}
		res = append(res, tmp)
	}
	return res
}

func GetBoardByTitle(pageSize, currPage int, title string) (int, interface{}) {
	var res []resBoard
	total, boards := db.GetBoardByQuery(pageSize, currPage, "", title)
	if total != 0 {
		res = packResBoards(boards)
	}
	return total, res
}

func NewBoard(content, title, typeName string, published bool) {
	db.NewBoard(content, title, typeName, published)
}
