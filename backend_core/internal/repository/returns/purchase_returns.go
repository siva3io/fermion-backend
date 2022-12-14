package returns

import (
	"errors"
	"fmt"
	"os"
	"time"

	"fermion/backend_core/db"
	"fermion/backend_core/internal/model/pagination"
	"fermion/backend_core/internal/model/returns"
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
type PurchaseReturn interface {
	Save(data *returns.PurchaseReturns) error
	FindAll(page *pagination.Paginatevalue) (interface{}, error)
	FindOne(query map[string]interface{}) (interface{}, error)
	Update(query map[string]interface{}, data *returns.PurchaseReturns) error
	Delete(query map[string]interface{}) error

	SaveReturnLines(returns.PurchaseReturnLines) error
	UpdateReturnLines(map[string]interface{}, returns.PurchaseReturnLines) (int64, error)
	DeleteReturnLine(map[string]interface{}) error
	FindReturnLines(map[string]interface{}) (returns.PurchaseReturnLines, error)

	Search(query string) (interface{}, error)
	GetPurchaseReturnsHistory(productId uint, page *pagination.Paginatevalue) (interface{}, error)
}

type PurchaseReturns struct {
	db *gorm.DB
}

func NewPurchaseReturn() *PurchaseReturns {
	db := db.DbManager()
	return &PurchaseReturns{db}

}

func (r *PurchaseReturns) Save(data *returns.PurchaseReturns) error {
	err := r.db.Model(&returns.PurchaseReturns{}).Create(data).Error

	if err != nil {

		return err

	}

	return nil
}

func (r *PurchaseReturns) FindAll(page *pagination.Paginatevalue) (interface{}, error) {
	var data []returns.PurchaseReturns

	err := r.db.Model(&returns.PurchaseReturns{}).Scopes(helpers.Paginate(&returns.PurchaseReturns{}, page, r.db)).Preload(clause.Associations).Find(&data)

	if err.Error != nil {
		return nil, err.Error
	}

	return data, nil
}

func (r *PurchaseReturns) FindOne(query map[string]interface{}) (interface{}, error) {
	var data returns.PurchaseReturns

	err := r.db.Preload(clause.Associations + "." + clause.Associations).Where(query).First(&data)

	if err.RowsAffected == 0 {
		return nil, errors.New("record not found")
	}

	if err.Error != nil {
		return nil, err.Error
	}

	return data, nil
}

func (r *PurchaseReturns) Update(query map[string]interface{}, data *returns.PurchaseReturns) error {
	res := r.db.Model(&returns.PurchaseReturns{}).Where(query).Updates(data)

	if res.Error != nil {

		return res.Error

	}

	return nil
}

func (r *PurchaseReturns) Delete(query map[string]interface{}) error {
	zone := os.Getenv("DB_TZ")
	loc, _ := time.LoadLocation(zone)
	data := map[string]interface{}{
		"deleted_by": query["user_id"].(int),
		"deleted_at": time.Now().In(loc),
	}
	delete(query, "user_id")
	res := r.db.Model(&returns.PurchaseReturns{}).Where(query).Updates(data)

	if res.Error != nil {

		return res.Error

	}

	return nil
}

func (r *PurchaseReturns) SaveReturnLines(data returns.PurchaseReturnLines) error {

	res := r.db.Model(&returns.PurchaseReturnLines{}).Create(&data)

	if res.Error != nil {

		return res.Error

	}

	return nil
}

func (r *PurchaseReturns) FindReturnLines(query map[string]interface{}) (returns.PurchaseReturnLines, error) {
	var result returns.PurchaseReturnLines
	fmt.Println(query)
	res := r.db.Model(&returns.PurchaseReturnLines{}).Where(query).First(&result)

	if res.Error != nil {
		return result, res.Error
	}

	return result, nil
}

func (r *PurchaseReturns) UpdateReturnLines(query map[string]interface{}, data returns.PurchaseReturnLines) (int64, error) {
	res := r.db.Model(&returns.PurchaseReturnLines{}).Where(query).Updates(&data)

	if res.Error != nil {

		return res.RowsAffected, res.Error

	}

	return res.RowsAffected, nil
}

func (r *PurchaseReturns) DeleteReturnLine(query map[string]interface{}) error {
	res := r.db.Model(&returns.PurchaseReturnLines{}).Where(query).Delete(&returns.PurchaseReturnLines{})

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (r *PurchaseReturns) Search(query string) (interface{}, error) {
	var data []returns.PurchaseReturns

	fields := []string{"reference_number", "number"}

	fields_string, values := helpers.ApplySearch(query, fields)

	err := r.db.Model(&returns.PurchaseReturns{}).Limit(20).Preload(clause.Associations).Where(fields_string, values...).Find(&data)

	if err.Error != nil {
		return nil, err.Error
	}

	return data, nil
}

func (r *PurchaseReturns) GetPurchaseReturnsHistory(productId uint, page *pagination.Paginatevalue) (interface{}, error) {
	var data []returns.PurchaseReturns
	var ids = make([]uint, 0)

	page.Filters = fmt.Sprintf("[[\"product_id\", \"=\", %v]]", productId)
	err := r.db.Model(&returns.PurchaseReturnLines{}).Select("pr_id").Scopes(helpers.Paginate(&returns.PurchaseReturnLines{}, page, r.db)).Scan(&ids)

	if err.Error != nil {
		return nil, err.Error
	}
	err = r.db.Model(&returns.PurchaseReturns{}).Where("id IN ?", ids).Find(&data)

	if err.Error != nil {
		return nil, err.Error
	}

	return data, nil
}
