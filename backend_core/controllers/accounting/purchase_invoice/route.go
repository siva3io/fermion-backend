package purchase_invoice

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
	g.GET("", h.ListPurchaseInvoice, cmiddleware.Authorization)
	g.GET("/dropdown", h.ListPurchaseInvoiceDropDown, cmiddleware.Authorization)
	g.GET("/:id", h.ViewPurchaseInvoice, cmiddleware.Authorization)
	g.POST("/create", h.CreatePurchaseInvoice, cmiddleware.Authorization, PurchaseInvoiceCreateValidate)
	g.POST("/:id/update", h.UpdatePurchaseInvoice, cmiddleware.Authorization, PurchaseInvoiceUpdateValidate)
	g.DELETE("/:id/delete", h.DeletePurchaseInvoice, cmiddleware.Authorization)
	g.DELETE("/invoice_lines/:id/delete", h.DeletePurchaseInvoiceLines, cmiddleware.Authorization)
	g.GET("/search", h.SearchPurchaseInvoice, cmiddleware.Authorization)

	g.POST("/:id/downloadPdf", h.DownloadPurchaseInvoice, cmiddleware.Authorization)
	g.POST("/:id/sendEmail", h.EmailPurchaseInvoice, cmiddleware.Authorization)
	g.POST("/:id/generatePdf", h.GeneratePurchaseInvoicePDF, cmiddleware.Authorization)
	g.POST("/:id/printPdf", h.PrintPurchaseInvoice, cmiddleware.Authorization)

	g.POST("/:id/favourite", h.FavouritePurchaseInvoice, cmiddleware.Authorization)
	g.POST("/:id/unfavourite", h.UnFavouritePurchaseInvoice, cmiddleware.Authorization)
	g.GET("/:id/filter_module/:tab", h.GetPurchaseInvoiceTab, cmiddleware.Authorization)
}
