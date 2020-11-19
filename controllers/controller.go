package controllers

import (
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type pagingResult struct {
	Page      int   `json:"page"`
	Limit     int   `json:"limit"`
	PrevPage  int   `json:"prevPage"`
	NextPage  int   `json:"nextPage"`
	Count     int64 `json:"count"`
	TotalPage int   `json:"totalPage"`
}

type pagination struct {
	ctx     *gin.Context
	query   *gorm.DB
	records interface{}
}

func (p *pagination) paginate() *pagingResult {
	page, _ := strconv.Atoi(p.ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(p.ctx.DefaultQuery("limit", "12"))

	count := p.countRecords()

	offset := (page - 1) * limit
	p.query.Limit(limit).Offset(offset).Find(p.records)

	totalPage := int(math.Ceil(float64(count) / float64(limit)))

	var nextPage int
	if page == totalPage {
		nextPage = totalPage
	} else {
		nextPage = page + 1
	}

	return &pagingResult{
		Page:      page,
		Limit:     limit,
		Count:     count,
		PrevPage:  page - 1,
		NextPage:  nextPage,
		TotalPage: totalPage,
	}
}

func (p *pagination) countRecords() int64 {
	var count int64
	p.query.Model(p.records).Count(&count)

	return count
}
