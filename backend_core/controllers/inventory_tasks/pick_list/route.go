package pick_list

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
	g.GET("", h.GetAllPickList, cmiddleware.Authorization)
	g.GET("/:id", h.GetPickList, cmiddleware.Authorization)
	g.POST("/create", h.CreatePickList, cmiddleware.Authorization, PickListCreateValidate)
	g.POST("/bulk_create", h.BulkCreatePickList, cmiddleware.Authorization)
	g.PUT("/:id/edit", h.UpdatePickList, cmiddleware.Authorization, PickListUpdateValidate)
	g.DELETE("/:id/delete", h.DeletePickList, cmiddleware.Authorization)
	g.DELETE("/:id/delete_products", h.DeletePickListLines, cmiddleware.Authorization)

	g.GET("/:id/sendemail", h.SendMailPickList, cmiddleware.Authorization)
	g.GET("/:id/printpdf", h.DownloadPdfPickList, cmiddleware.Authorization)
	g.POST("/:id/favourite", h.FavouritePickList, cmiddleware.Authorization)
	g.POST("/:id/unfavourite", h.UnFavouritePickList, cmiddleware.Authorization)
	g.GET("/favourite_list", h.FavouritePickListView, cmiddleware.Authorization)
	g.GET("/dropdown", h.GetAllPickListDropDown, cmiddleware.Authorization)
}
