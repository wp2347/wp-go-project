package handler

import (
	"strconv"
	"wp-demo/pkg/domain/service"
	"wp-demo/pkg/domain/model"

	"github.com/gin-gonic/gin"
)


type CreateArticleReq struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Author	string `json:"author"`
}

func CreateArticle(articleSrv *service.ArticleService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request CreateArticleReq
		err := ctx.ShouldBindJSON(&request)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}
		err = articleSrv.Create(ctx, request.Title, request.Content, request.Author)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}
		ctx.JSON(200, "创建成功")
	}
}

type GetArticleReq struct {
	ID uint `json:"id"`
}


func GetArticle(articleSrv *service.ArticleService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		idStr := ctx.Param("id")
		id, err := strconv.ParseUint(idStr, 10, 32)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "无效的 ID"})
			return
		}
		article, err := articleSrv.Get(ctx, uint(id))
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		if article == nil {
			ctx.JSON(404, gin.H{"error": "文章不存在"})
			return
		}
		ctx.JSON(200, article)
	}
}


type DeleteArticleReq struct {
	ID uint `json:"id"`
}

func DeleteArticle(articleSrv *service.ArticleService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		idStr := ctx.Param("id")
		id, err := strconv.ParseUint(idStr, 10, 32)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "无效的 ID"})
			return
		}
		err = articleSrv.Delete(ctx, uint(id))
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, gin.H{"message": "删除成功"})
	}
}



type ListArticleReq struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
	Author	 string `json:"author"`
}

func ListArticle(articleSrv *service.ArticleService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request ListArticleReq
		err := ctx.ShouldBindQuery(&request)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}
		var articles []model.Article
		var total int64
		if request.Author != "" {
			articles, total, err = articleSrv.ListByAuthor(ctx, request.Author, request.Page, request.PageSize)
		} else {
			articles, total, err = articleSrv.List(ctx, request.Page, request.PageSize)
		}
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}
		ctx.JSON(200, gin.H{
			"articles": articles,
			"total":    total,
		})
	}
}
