package service

import (
	"RuiServer/db"
	"RuiServer/exception"
	"RuiServer/utils"
)

func GetBoardByType(pageSize, currPage int, typeName string) (int, interface{}) {
	boardType, ok := utils.BoardTypeMap[typeName]
	if !ok {
		exception.ExceptionCustom("GetBoardByType", exception.GetBoardError, nil)
	}
	return db.GetBoardByQuery(pageSize, currPage, boardType)
}
