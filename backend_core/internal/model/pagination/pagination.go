package pagination

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
type Paginatevalue struct {
	Filters    string `json:"filters,omitempty" query:"filters"`
	Per_page   int    `json:"per_page,omitempty" query:"per_page"`
	Page_no    int    `json:"page_no,omitempty" query:"page_no"`
	Sort       string `json:"sort,omitempty" query:"sort"`
	TotalRows  int64  `json:"total_rows"`
	TotalPages int    `json:"total_pages"`
}

func (p *Paginatevalue) GetPage() int {
	if p.Page_no == 0 {
		p.Page_no = 1
	}
	return p.Page_no
}

func (p *Paginatevalue) GetOffset() int {
	return (p.GetPage() - 1) * p.Per_page
}

func (p *Paginatevalue) GetLimit() int {
	if p.Per_page == 0 {
		p.Per_page = 10
	}
	return p.Per_page
}

func (p *Paginatevalue) GetSort() string {
	if p.Sort == "" {
		p.Sort = "Id desc"
	}
	return p.Sort
}
