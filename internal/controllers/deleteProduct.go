package controllers

import (
	"database/sql"
	"net/http"

	"github.com/afan104/products-api/internal"
	"github.com/gin-gonic/gin"
)

func DeleteProduct(db *sql.DB) gin.HandlerFunc{
	return func(c *gin.Context) {
		// extract guid
		var ctx = c.Request.Context()
		var binding guidBinding
		if e := c.ShouldBindUri(&binding); e != nil {
			var res = internal.NewHTTPResponse(http.StatusInternalServerError, e)
			c.JSON(http.StatusInternalServerError, res)
			return
		}

		// delete sql row and return the results (error or ok)
		var result sql.Result
		var e error
		if result, e = db.ExecContext(ctx, "DELETE FROM products WHERE guid=?", binding.GUID); e != nil {
			var res = internal.NewHTTPResponse(http.StatusInternalServerError, e)
			c.JSON(http.StatusInternalServerError, res)
			return
		}

		if nProducts, _ := result.RowsAffected(); nProducts == 0 {
			var res = internal.NewHTTPResponse(http.StatusNotFound, sql.ErrNoRows)
			c.JSON(http.StatusNotFound, res)
			return
		}

		c.JSON(http.StatusNoContent, nil)
	}

}