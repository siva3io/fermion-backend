package pricing

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
	g.POST("/create", h.CreatePricing, cmiddleware.Authorization, PricingCreateValidate)
	g.GET("", h.PricingList, cmiddleware.Authorization)
	g.GET("/dropdown", h.PricingListDropdown, cmiddleware.Authorization)
	g.GET("/:id", h.FindPricing, cmiddleware.Authorization)
	g.POST("/:id/edit", h.UpdatePricing, cmiddleware.Authorization, PricingUpdateValidate)
	g.DELETE("/:id/delete", h.DeletePricing, cmiddleware.Authorization)
	g.DELETE("/:id/delete_line_items", h.DeleteLineItems, cmiddleware.Authorization)
	g.POST("/:id/favourite", h.FavouritePricings, cmiddleware.Authorization)
	g.POST("/:id/unfavourite", h.UnFavouritePricings, cmiddleware.Authorization)
	g.GET("/favourite_list", h.FavouritePricingView, cmiddleware.Authorization)

	g.GET("/sales_price_line_items_list", h.SalesPriceLineItems, cmiddleware.Authorization)

	//---------------------Channel API's ----------------------------------------------------------------
	g.POST("/channels/upsert", h.ChannelSalesPriceUpsert, cmiddleware.Authorization)

}
