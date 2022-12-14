package inventory_tasks

import (
	"os"
	"time"

	"fermion/backend_core/db"
	"fermion/backend_core/internal/model/inventory_tasks"
	"fermion/backend_core/internal/model/pagination"
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
type CycleCount interface {
	CreateCycleCount(*inventory_tasks.CycleCount) (uint, error)
	BulkCreateCycleCount(*[]inventory_tasks.CycleCount) error
	UpdateCycleCount(uint, *inventory_tasks.CycleCount) error
	GetCycleCount(uint) (inventory_tasks.CycleCount, error)
	GetAllCycleCount(*pagination.Paginatevalue) ([]inventory_tasks.CycleCount, error)
	DeleteCycleCount(uint, uint) error

	CreateCycleCountLines(inventory_tasks.CycleCountLines) error
	UpdateCycleCountLines(interface{}, inventory_tasks.CycleCountLines) (int64, error)
	GetCycleCountLines(interface{}) ([]inventory_tasks.CycleCountLines, error)
	DeleteCycleCountLines(interface{}) error
}

type cycleCount struct {
	db *gorm.DB
}

func NewCycleCount() *cycleCount {
	db := db.DbManager()
	return &cycleCount{db}
}

func (r *cycleCount) CreateCycleCount(data *inventory_tasks.CycleCount) (uint, error) {
	var scode uint
	err := r.db.Raw("SELECT lookupcodes.id FROM lookuptypes,lookupcodes WHERE lookuptypes.id = lookupcodes.lookup_type_id AND lookuptypes.lookup_type = 'CYCLE_COUNT_STATUS' AND lookupcodes.lookup_code = 'DRAFT'").First(&scode).Error
	if err != nil {
		return 0, err
	}
	data.StatusID = scode
	res, _ := helpers.UpdateStatusHistory(data.StatusHistory, data.StatusID)
	data.StatusHistory = res
	result := r.db.Model(&inventory_tasks.CycleCount{}).Create(&data)
	return data.ID, result.Error
}

func (r *cycleCount) BulkCreateCycleCount(data *[]inventory_tasks.CycleCount) error {
	for _, value := range *data {
		var scode uint
		err := r.db.Raw("SELECT lookupcodes.id FROM lookuptypes,lookupcodes WHERE lookuptypes.id = lookupcodes.lookup_type_id AND lookuptypes.lookup_type = 'CYCLE_COUNT_STATUS' AND lookupcodes.lookup_code = 'DRAFT'").First(&scode).Error
		if err != nil {
			return err
		}
		value.StatusID = scode
		res, _ := helpers.UpdateStatusHistory(value.StatusHistory, value.StatusID)
		value.StatusHistory = res
		result := r.db.Model(&inventory_tasks.CycleCount{}).Create(&value)
		if result.Error != nil {
			return result.Error
		}
	}
	return nil
}

func (r *cycleCount) UpdateCycleCount(id uint, data *inventory_tasks.CycleCount) error {
	result := r.db.Model(&inventory_tasks.CycleCount{}).Where("id", id).Updates(&data)
	return result.Error
}

func (r *cycleCount) GetCycleCount(id uint) (inventory_tasks.CycleCount, error) {
	var data inventory_tasks.CycleCount
	result := r.db.Model(&inventory_tasks.CycleCount{}).Preload(clause.Associations).Where("id", id).First(&data)
	return data, result.Error
}

func (r *cycleCount) GetAllCycleCount(p *pagination.Paginatevalue) ([]inventory_tasks.CycleCount, error) {
	var data []inventory_tasks.CycleCount
	res := r.db.Model(&inventory_tasks.CycleCount{}).Scopes(helpers.Paginate(&inventory_tasks.CycleCount{}, p, r.db)).Where("is_active = true").Preload("OrderLines.Product").Preload("OrderLines.ProductVariant").Preload("OrderLines.LocationSpaceType").Preload("OrderLines.LocationInputType").Preload(clause.Associations).Find(&data)
	return data, res.Error
}

func (r *cycleCount) DeleteCycleCount(id uint, user_id uint) error {
	zone := os.Getenv("DB_TZ")
	loc, _ := time.LoadLocation(zone)
	data := map[string]interface{}{
		"deleted_by": user_id,
		"deleted_at": time.Now().In(loc),
	}
	result := r.db.Model(&inventory_tasks.CycleCount{}).Where("id", id).Updates(data)
	return result.Error
}

func (r *cycleCount) CreateCycleCountLines(data inventory_tasks.CycleCountLines) error {
	result := r.db.Model(&inventory_tasks.CycleCountLines{}).Create(&data)
	return result.Error
}

func (r *cycleCount) UpdateCycleCountLines(query interface{}, data inventory_tasks.CycleCountLines) (int64, error) {
	result := r.db.Model(&inventory_tasks.CycleCountLines{}).Where(query).Updates(&data)
	return result.RowsAffected, result.Error
}

func (r *cycleCount) GetCycleCountLines(query interface{}) ([]inventory_tasks.CycleCountLines, error) {
	var data []inventory_tasks.CycleCountLines
	result := r.db.Model(&inventory_tasks.CycleCountLines{}).Where(query).Preload(clause.Associations).Find(&data)
	return data, result.Error
}

func (r *cycleCount) DeleteCycleCountLines(query interface{}) error {
	result := r.db.Model(&inventory_tasks.CycleCountLines{}).Where(query).Delete(&inventory_tasks.CycleCountLines{})
	return result.Error
}
