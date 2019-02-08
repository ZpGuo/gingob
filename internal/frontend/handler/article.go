package handler

import (
	"net/http"
	"strconv"

	"github.com/puti-projects/puti/internal/frontend/service"

	"github.com/gin-gonic/gin"
)

// ShowArticleList article list handle
func ShowArticleList(c *gin.Context) {
	// get renderer data include basic data
	renderData := getRenderData(c)

	// get params
	currentPage, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	// get content
	articles, pagination, err := service.GetArticleList(currentPage, "")
	if err != nil {
		// 500
	}

	renderData["Articles"] = articles

	renderData["Pagination"] = pagination.Page
	pagination.SetPageURL("/article")
	renderData["PageURL"] = pagination.PageURL

	renderData["Widgets"] = getWidgets()
	renderData["Title"] = "文章" + " - " + renderData["Setting"].(map[string]interface{})["BlogName"].(string)
	c.HTML(http.StatusOK, getTheme(c)+"/articles.html", renderData)
}

// ShowCategoryArticleList handle article list by category
func ShowCategoryArticleList(c *gin.Context) {
	// get renderer data include basic data
	renderData := getRenderData(c)

	// get params
	taxonomySlug := c.Param("slug")
	currentPage, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	// get content
	termName, articles, pagination, err := service.GetArticleListByTaxonomy(currentPage, "category", taxonomySlug, "")
	if err != nil {
		// 500
	}

	renderData["Articles"] = articles

	renderData["Pagination"] = pagination.Page
	pagination.SetPageURL("/category/" + taxonomySlug)
	renderData["PageURL"] = pagination.PageURL

	renderData["Widgets"] = getWidgets()
	renderData["Title"] = termName + " - 分类 - " + renderData["Setting"].(map[string]interface{})["BlogName"].(string)
	c.HTML(http.StatusOK, getTheme(c)+"/articles.html", renderData)
}

// ShowTagArticleList handle article list by tag
func ShowTagArticleList(c *gin.Context) {
	// get renderer data include basic data
	renderData := getRenderData(c)

	// get params
	taxonomySlug := c.Param("slug")
	currentPage, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	// get content
	termName, articles, pagination, err := service.GetArticleListByTaxonomy(currentPage, "tag", taxonomySlug, "")
	if err != nil {
		// 500
	}

	renderData["Articles"] = articles

	renderData["Pagination"] = pagination.Page
	pagination.SetPageURL("/tag/" + taxonomySlug)
	renderData["PageURL"] = pagination.PageURL

	renderData["Widgets"] = getWidgets()
	renderData["Title"] = termName + " - 标签 - " + renderData["Setting"].(map[string]interface{})["BlogName"].(string)
	c.HTML(http.StatusOK, getTheme(c)+"/articles.html", renderData)
}