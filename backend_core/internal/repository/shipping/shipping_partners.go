package shipping

import (
	"os"
	"time"

	"fermion/backend_core/db"
	"fermion/backend_core/internal/model/pagination"
	"fermion/backend_core/internal/model/shipping"
	"fermion/backend_core/pkg/util/helpers"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
type ShippingPartner interface {
	Create(data *shipping.UserShippingPartnerRegistration, userID uint) (uint, error)
	Update(id uint, data shipping.UserShippingPartnerRegistration) error
	FindOne(id uint) (shipping.UserShippingPartnerRegistration, error)
	FindAll(p *pagination.Paginatevalue) ([]shipping.UserShippingPartnerRegistration, error)
	Delete(id uint, user_id uint) error

	ShippingPartnerEstimateCosts(data *shipping.RateCalculator) ([]shipping.RateCalculator, error)

	FindAllShippingpartner(p *pagination.Paginatevalue) ([]shipping.ShippingPartner, error)
	FindOneShippingpartnerByName(name string) (shipping.ShippingPartner, error)
	FindOneShippingpartnerById(id int) (shipping.ShippingPartner, error)
	UpdateShippingPartnerByName(query map[string]interface{}) (int64, error)
}

type shippingPartner struct {
	db *gorm.DB
}

func NewShipping() *shippingPartner {
	db := db.DbManager()
	return &shippingPartner{db}
}

func (r *shippingPartner) Create(data *shipping.UserShippingPartnerRegistration, userID uint) (uint, error) {
	data.UserId = userID
	res := r.db.Model(&shipping.UserShippingPartnerRegistration{}).Create(&data)
	return data.ID, res.Error
}

func (r *shippingPartner) ShippingPartnerEstimateCosts(data *shipping.RateCalculator) ([]shipping.RateCalculator, error) {
	var result []shipping.RateCalculator
	res := r.db.Model(&shipping.RateCalculator{}).Preload(clause.Associations).Find(&result)
	return result, res.Error
}

func (r *shippingPartner) Update(id uint, data shipping.UserShippingPartnerRegistration) error {
	res := r.db.Model(&shipping.UserShippingPartnerRegistration{}).Where("id", id).Updates(&data)
	return res.Error
}

func (r *shippingPartner) FindOne(id uint) (shipping.UserShippingPartnerRegistration, error) {
	var data shipping.UserShippingPartnerRegistration
	result := r.db.Preload(clause.Associations).Model(&shipping.UserShippingPartnerRegistration{}).Preload(clause.Associations).Where("id", id).First(&data)
	if result.Error != nil {
		return data, result.Error
	}
	return data, nil
}

func (r *shippingPartner) FindAll(p *pagination.Paginatevalue) ([]shipping.UserShippingPartnerRegistration, error) {
	var data []shipping.UserShippingPartnerRegistration
	err := r.db.Model(&shipping.UserShippingPartnerRegistration{}).Preload(clause.Associations).Scopes(helpers.Paginate(&shipping.UserShippingPartnerRegistration{}, p, r.db)).Where("is_active = true").Find(&data).Error
	if err != nil {
		return data, err
	}
	return data, nil
}
func (r *shippingPartner) Delete(id uint, user_id uint) error {
	zone := os.Getenv("DB_TZ")
	loc, _ := time.LoadLocation(zone)
	data := map[string]interface{}{
		"deleted_by": user_id,
		"deleted_at": time.Now().In(loc),
	}
	res := r.db.Model(&shipping.UserShippingPartnerRegistration{}).Where("id", id).Updates(data)
	return res.Error
}

func (r shippingPartner) FindAllShippingpartner(p *pagination.Paginatevalue) ([]shipping.ShippingPartner, error) {
	var result []shipping.ShippingPartner
	res := r.db.Preload(clause.Associations).Model(&shipping.ShippingPartner{}).Scopes(helpers.Paginate(&shipping.ShippingPartner{}, p, r.db)).Where("is_active = true").Find(&result)
	return result, res.Error
}

func (r shippingPartner) FindOneShippingpartnerByName(name string) (shipping.ShippingPartner, error) {
	var result shipping.ShippingPartner
	res := r.db.Model(&shipping.ShippingPartner{}).Where("partner_name", name).Find(&result)
	return result, res.Error
}

func (r shippingPartner) FindOneShippingpartnerById(id int) (shipping.ShippingPartner, error) {
	var result shipping.ShippingPartner
	res := r.db.Model(&shipping.ShippingPartner{}).Where("id", id).Find(&result)
	return result, res.Error
}

func (r shippingPartner) UpdateShippingPartnerByName(query map[string]interface{}) (int64, error) {
	data := query["data"].(shipping.ShippingPartner)
	delete(query, "data")
	res := r.db.Model(&shipping.ShippingPartner{}).Where(query).Updates(data)
	if res.RowsAffected != 0 {
		return res.RowsAffected, res.Error
	}
	return res.RowsAffected, res.Error
}
