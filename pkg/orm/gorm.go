package orm

import "gorm.io/gorm"

//
// Page 数据库分页对象
//
type Page struct {
	PageNum   int         `json:"page_num"`   // 页码
	PageSize  int         `json:"page_size"`  // 每页记录数
	TotalPage int         `json:"total_page"` // 总页数
	TotalRows int64       `json:"total_rows"` // 总行数
	Rows      interface{} `json:"rows"`       // 当前页的数据
}

//
// NewPage 构造函数
//	@parameter	pageNum 页码, 从 1 开始
//	@parameter	pageSize 每页的记录数
//	@parameter	totalRows 符合条件的记录总行数
//	@parameter	rows 符合条件的记录
//	@return		*Page
//
func NewPage(pageNum int, pageSize int, totalRows int64, rows interface{}) *Page {
	if pageNum <= 0 {
		pageNum = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	page := &Page{
		PageNum:   pageNum,
		PageSize:  pageSize,
		TotalPage: 0,
		TotalRows: totalRows,
		Rows:      rows,
	}

	tempPage := int(page.TotalRows / int64(page.PageSize))

	if page.TotalRows%int64(page.PageSize) == 0 {
		page.TotalPage = tempPage
	} else {
		page.TotalPage = tempPage + 1
	}
	return page
}

//
// Paginate
//	@parameter	pageNum 页码, 从 1 开始
//	@parameter	pageSize 每页的行数, 大于 0
//	@return		func(dao *gorm.DB) *gorm.DB
//
func Paginate(pageNum int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if pageNum == 0 {
			pageNum = 1
		}
		if pageSize < 0 {
			pageSize = 10
		}
		offset := (pageNum - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
