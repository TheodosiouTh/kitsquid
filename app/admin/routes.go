package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/n1try/kithub2/app/common/errors"
	"github.com/n1try/kithub2/app/config"
	"net/http"
	"strings"
)

func RegisterRoutes(router *gin.Engine, group *gin.RouterGroup) {
	group.GET("/admin", CheckAdmin(), index(router))
}

func RegisterApiRoutes(router *gin.Engine, group *gin.RouterGroup) {
	group.POST("/admin/query", CheckAdmin(), apiAdminQuery(router))
	group.POST("/admin/flush", CheckAdmin(), apiAdminFlush(router))
}

func index(r *gin.Engine) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin", gin.H{
			"entities": entities,
			"tplCtx":   c.MustGet(config.TemplateContextKey),
		})
	}
}

func apiAdminQuery(r *gin.Engine) func(c *gin.Context) {
	return func(c *gin.Context) {
		var re *registeredEntity
		var query adminQuery

		if err := c.ShouldBindJSON(&query); err != nil {
			c.Error(err)
			c.AbortWithError(http.StatusBadRequest, errors.BadRequest{})
			return
		}

		re, ok := registry[query.Entity]
		if !ok {
			c.AbortWithError(http.StatusNotFound, errors.NotFound{})
			return
		}

		var f func(*adminQuery, *registeredEntity, *gin.Context)

		switch strings.ToLower(query.Action) {
		case "list":
			f = handleList
			break
		case "get":
			f = handleGet
			break
		case "put":
			f = handlePut
			break
		case "delete":
			f = handleDelete
			break
		case "flush":
			f = handleFlush
			break
		default:
			f = func(_q *adminQuery, _e *registeredEntity, context *gin.Context) {
				c.AbortWithError(http.StatusBadRequest, errors.BadRequest{})
				return
			}
		}

		f(&query, re, c)
	}
}

func apiAdminFlush(r *gin.Engine) func(c *gin.Context) {
	return func(c *gin.Context) {
		go func() {
			for _, e := range entities {
				go e.Resolvers.Flush()
			}
		}()

		c.Status(http.StatusAccepted)
	}
}

func handleList(q *adminQuery, re *registeredEntity, c *gin.Context) {
	result, err := re.Resolvers.List()
	if err != nil {
		c.Error(err)
		c.AbortWithError(http.StatusInternalServerError, errors.Internal{})
		return
	}

	c.JSON(http.StatusOK, result)
}

func handleGet(q *adminQuery, re *registeredEntity, c *gin.Context) {
	if q.Key == "" {
		c.AbortWithError(http.StatusBadRequest, errors.BadRequest{})
		return
	}

	result, err := re.Resolvers.Get(q.Key)
	if err != nil {
		c.Error(err)
		c.AbortWithError(http.StatusInternalServerError, errors.Internal{})
		return
	}

	c.JSON(http.StatusOK, result)
}

func handlePut(q *adminQuery, re *registeredEntity, c *gin.Context) {
	if q.Key == "" || q.Value == "" {
		c.AbortWithError(http.StatusBadRequest, errors.BadRequest{})
		return
	}

	if err := re.Resolvers.Put(q.Key, q.Value); err != nil {
		c.Error(err)
		c.AbortWithError(http.StatusInternalServerError, errors.Internal{})
		return
	}

	c.Status(http.StatusOK)
}

func handleDelete(q *adminQuery, re *registeredEntity, c *gin.Context) {
	if q.Key == "" {
		c.AbortWithError(http.StatusBadRequest, errors.BadRequest{})
		return
	}

	if err := re.Resolvers.Delete(q.Key); err != nil {
		c.Error(err)
		c.AbortWithError(http.StatusInternalServerError, errors.Internal{})
		return
	}

	c.Status(http.StatusNoContent)
}

func handleFlush(q *adminQuery, re *registeredEntity, c *gin.Context) {
	go re.Resolvers.Flush()
	c.Status(http.StatusAccepted)
}
