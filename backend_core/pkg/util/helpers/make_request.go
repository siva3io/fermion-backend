package helpers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
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
type Request struct {
	Method string
	Scheme string
	Host   string
	Path   string
	Header map[string]string
	Params map[string]string
	Body   interface{}
}

func (request Request) logRequest() {

	log.Printf("%+v", request)
}

func MakeRequest(request Request) (interface{}, error) {

	requestBody, ok := request.Body.([]byte)
	if !ok {
		if request.Body != nil {
			marshaldata, err := json.Marshal(request.Body)
			if err != nil {
				return nil, err
			}
			requestBody = marshaldata
		} else {
			requestBody = []byte("")
		}
	}

	queryParams := url.Values{}

	for key, value := range request.Params {
		queryParams.Set(key, value)
	}

	url := url.URL{
		Scheme:   request.Scheme,
		Host:     request.Host,
		Path:     request.Path,
		RawQuery: queryParams.Encode(),
	}
	var req *http.Request
	var err error
	if request.Header["Content-Type"] != "" && request.Header["Content-Type"] == "application/x-www-form-urlencoded" {
		urlEncodedpayload, _ := ReturnURLEncodeString(request.Body)
		req, err = http.NewRequest(request.Method, url.String(), strings.NewReader(urlEncodedpayload))
	} else {
		req, err = http.NewRequest(request.Method, url.String(), bytes.NewBuffer(requestBody))
	}

	if err != nil {
		return nil, err
	}

	for key, value := range request.Header {
		req.Header.Set(key, value)
	}

	request.logRequest()

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	responseBodyByte, _ := ioutil.ReadAll(res.Body)

	var responseBody interface{}

	err = json.Unmarshal(responseBodyByte, &responseBody)

	if err != nil {
		return nil, err
	}

	return responseBody, nil
}
