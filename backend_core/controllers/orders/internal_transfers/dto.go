package internal_transfers

import (
	model_core "fermion/backend_core/internal/model/core"
	"fermion/backend_core/internal/model/mdm/shared_pricing_and_location"
	"fermion/backend_core/internal/model/orders"
	"fermion/backend_core/pkg/util/response"
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
type (
	ListInternalTransfersDTO struct {
		ID                    int64                                 `json:"id"`
		IstNumber             string                                `json:"ist_number"`
		ReferenceNumber       string                                `json:"reference_number"`
		IstDate               string                                `json:"ist_date"`
		SourceLocationId      uint                                  `json:"source_location_id"`
		SourceWarehouse       shared_pricing_and_location.Locations `json:"source_warehouse"`
		DestinationLocationId uint                                  `json:"destination_location_id"`
		DestinationWarehouse  shared_pricing_and_location.Locations `json:"destination_warehouse"`
		NoOfItems             int                                   `json:"no_of_items"`
		TotalQuantity         int                                   `json:"total_quantity"`
		ShippingModeId        uint                                  `json:"shipping_mode_id"`
		ShippingMode          model_core.Lookupcode                 `json:"shipping_mode"`
		StatusId              uint                                  `json:"status_id"`
		Status                model_core.Lookupcode                 `json:"status"`
		ScheduledDeliveryDate string                                `json:"scheduled_delivery_date"`
		SourceDocuments       map[string]interface{}                `json:"source_documents"`
		SourceDocumentTypeId  *uint                                 `json:"source_document_type_id"`
		GrnIDs                uint                                  `json:"grn_id" gorm:"type:integer"`
		// GRN                   []inventory_orders.GRN `json:"grn" gorm:"foreignKey:GrnIDs"`
		model_core.Model
	}
	GetAllInternalTransfersResponse struct {
		Body struct {
			Meta response.MetaResponse
			Data []ListInternalTransfersDTO
		}
	} //@ name GetAllInternalTransfersResponse
)

type InternalTransfersDTO struct {
	IstDate               string                  `json:"ist_date"`
	ReferenceNumber       string                  `json:"reference_number"`
	ReceiptRoutingId      uint                    `json:"receipt_routing_id"`
	SourceLocationId      uint                    `json:"source_location_id"`
	DestinationLocationId uint                    `json:"destination_location_id"`
	NoOfItems             int                     `json:"no_of_items"`
	TotalQuantity         int                     `json:"total_quantity"`
	ShippingModeId        uint                    `json:"shipping_mode_id"`
	StatusId              uint                    `json:"status_id"`
	StatusHistory         map[string]interface{}  `json:"status_history"`
	GrnId                 uint                    `json:"grn_id"`
	ShippingDetails       map[string]interface{}  `json:"shipping_details"`
	InternalTransferLines []InternalTransferLines `json:"internal_transfer_lines"`
	ReasonId              uint                    `json:"reason_id"`
	PickupDateAndTime     PickupDateAndTime       `json:"pickup_date_and_time"`
	SourceDocuments       map[string]interface{}  `json:"source_documents"`
	SourceDocumentTypeId  *uint                   `json:"source_document_type_id"`
	model_core.Model
}

type InternalTransferLines struct {
	IstId             uint   `json:"ist_id"`
	ProductId         uint   `json:"product_id"`
	SerialNumber      string `json:"serial_number"`
	ProductTemplateId uint   `json:"product_template_id"`
	InventoryId       uint   `json:"inventory_id"`
	UomId             uint   `json:"uom_id"`
	SourceStock       int    `json:"source_stock"`
	DestinationStock  int    `json:"destination_stock"`
	IsScrap           bool   `json:"is_scrap"`
	TransferQuantity  int    `json:"transfer_quantity"`
	model_core.Model
}
type PickupDateAndTime struct {
	PickupDate     string `json:"pickup_date"`
	PickupFromTime string `json:"pickup_from_time"`
	PickupToTime   string `json:"pickup_to_time"`
}

// Create InternalTransfers response
type (
	InternalTransfersCreate struct {
		Created_id int
	}
	InternalTransfersCreateResponse struct {
		Body struct {
			Meta response.MetaResponse
			Data InternalTransfersCreate
		}
	} //@ name InternalTransfersCreateResponse
)

// Update InternalTransfers response
type (
	InternalTransfersUpdate struct {
		Updated_id int
	}
	InternalTransfersUpdateResponse struct {
		Body struct {
			Meta response.MetaResponse
			Data InternalTransfersUpdate
		}
	} //@ name InternalTransfersUpdateResponse
)

// Get InternalTransfers response
type (
	InternalTransfersGet struct {
		orders.InternalTransfers
	}
	InternalTransfersGetResponse struct {
		Body struct {
			Meta response.MetaResponse
			Data InternalTransfersGet
		}
	} //@ name InternalTransfersGetResponse
)

// Delete InternalTransfers response
type (
	InternalTransfersDelete struct {
		Deleted_id int
	}
	InternalTransfersDeleteResponse struct {
		Body struct {
			Meta response.MetaResponse
			Data InternalTransfersDelete
		}
	} //@ name InternalTransfersDeleteResponse
)

// Delete InternalTransfer Orderlines response
type (
	InternalTransferOrderLinesDelete struct {
		Deleted_id int
		Product_id int
	}
	InternalTransferOrderLinesDeleteResponse struct {
		Body struct {
			Meta response.MetaResponse
			Data InternalTransferOrderLinesDelete
		}
	} //@ name InternalTransferOrderLinesDeleteResponse
)
