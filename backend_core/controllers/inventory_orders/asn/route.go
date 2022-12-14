package asn

import (
	cmiddleware "fermion/backend_core/middleware"

	"github.com/labstack/echo/v4"
)

/*
 Copyright (C) 2022 Eunimart Omnichannel Pvt Ltd. (www.eunimart.com)
 All rights reserved.
 This program is free software: you can redistribute it and/or modify
 it under the terms of the GNU General Public License as published by
 the Free Software Foundation, either version 3 of the License, or
 (at your option) any later version.
 This program is distributed in the hope that it will be useful,
 but WITHOUT ANY WARRANTY; without even the implied warranty of
 MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 GNU General Public License for more details.
 You should have received a copy of the GNU General Public License
 along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

func (h *handler) Route(g *echo.Group) {
	g.GET("", h.GetAllAsn, cmiddleware.Authorization)
	g.GET("/:id", h.GetAsn, cmiddleware.Authorization)
	g.POST("/create", h.CreateAsn, cmiddleware.Authorization, AsnCreateValidate)
	g.POST("/bulk_create", h.BulkCreateAsn, cmiddleware.Authorization)
	g.PUT("/:id/edit", h.UpdateAsn, cmiddleware.Authorization, AsnUpdateValidate)
	g.DELETE("/:id/delete", h.DeleteAsn, cmiddleware.Authorization)
	g.DELETE("/:id/delete_products", h.DeleteAsnLines, cmiddleware.Authorization)

	g.GET("/:id/sendemail", h.SendMailAsn, cmiddleware.Authorization)
	g.GET("/:id/printpdf", h.DownloadPdfAsn, cmiddleware.Authorization)
	g.POST("/:id/favourite", h.FavouriteAsn, cmiddleware.Authorization)
	g.POST("/:id/unfavourite", h.UnFavouriteAsn, cmiddleware.Authorization)
	g.GET("/favourite_list", h.FavouriteAsnView, cmiddleware.Authorization)

	g.GET("/:id/filter_module/:tab", h.GetAsnTab, cmiddleware.Authorization)
	g.GET("/dropdown", h.GetAllAsnDropdown, cmiddleware.Authorization)
}
