package scheduler

import (
	"fmt"
	"net/http"
	"strconv"

	res "fermion/backend_core/pkg/util/response"

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
type handler struct {
	service Service
}

func NewHandler() *handler {
	service := NewService()
	return &handler{service}
}

func Init() {
	NewService().InitSchedulerJob()
}

func (h *handler) AddSchedulerJob(c echo.Context) (err error) {

	var schedulerJobDto SchedulerJobDto

	if err := c.Bind(&schedulerJobDto); err != nil {
		return res.RespErr(c, err)
	}

	fmt.Println(schedulerJobDto)

	if err := h.service.AddSchedulerJob(schedulerJobDto); err != nil {
		return res.RespErr(c, err)
	}

	// logs, err := h.service.GetQueuedSchedulerLogs()

	return res.RespSuccess(c, "success", nil)
}
func (h *handler) ListSchedulerJob(c echo.Context) (err error) {
	var result []SchedulerJobResponseDTO
	result, err = h.service.GetAllSchedulerJob()
	if err != nil {
		return res.RespErr(c, err)
	}
	return c.JSON(http.StatusOK, result)

}

func (h *handler) ViewSchedulerJob(c echo.Context) (err error) {
	id := c.Param("id")
	ID, _ := strconv.Atoi(id)
	var query = map[string]interface{}{
		"id": ID,
	}
	result, err := h.service.GetOneSchedulerJob(query)
	if err != nil {
		return res.RespError(c, err)
	}
	return res.RespSuccess(c, "Scheduler Job details retrieved successfully", result)
}
func (h *handler) DeleteSchedulerJob(c echo.Context) (err error) {
	id := c.Param("id")
	ID, _ := strconv.Atoi(id)
	var query = make(map[string]interface{}, 0)
	query["id"] = ID
	s := c.Get("TokenUserID").(string)
	user_id, _ := strconv.Atoi(s)
	query["user_id"] = user_id
	err = h.service.DeleteSchedulerJob(query)
	if err != nil {
		return res.RespErr(c, err)
	}
	return res.RespSuccess(c, "Scheduler Job deleted successfully", map[string]string{"deleted_id": id})

}
func (h *handler) UpdateSchedulerJob(c echo.Context) (err error) {
	var id = c.Param("id")
	var query = make(map[string]interface{}, 0)
	ID, _ := strconv.Atoi(id)
	query["id"] = uint(ID)
	var data SchedulerJobDto
	err = c.Bind(&data)
	if err != nil {
		return err
	}
	err = h.service.UpdateSchedulerJob(query, &data)
	if err != nil {
		return res.RespError(c, err)
	}
	return res.RespSuccess(c, "Scheduler Job updated successfully", map[string]interface{}{"updated_id": id})
}
func (h *handler) ListSchedulerLogs(c echo.Context) error {
	id := c.Param("id")
	ID, _ := strconv.Atoi(id)
	var query = map[string]interface{}{
		"scheduler_job_id": ID,
	}
	result, err := h.service.ListSchedulerLogs(query)
	if err != nil {
		return res.RespError(c, err)
	}
	return res.RespSuccess(c, "Scheduler Job details retrieved successfully", result)

}
