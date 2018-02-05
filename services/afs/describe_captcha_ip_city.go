package afs

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

func (client *Client) DescribeCaptchaIpCity(request *DescribeCaptchaIpCityRequest) (response *DescribeCaptchaIpCityResponse, err error) {
	response = CreateDescribeCaptchaIpCityResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeCaptchaIpCityWithChan(request *DescribeCaptchaIpCityRequest) (<-chan *DescribeCaptchaIpCityResponse, <-chan error) {
	responseChan := make(chan *DescribeCaptchaIpCityResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCaptchaIpCity(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

func (client *Client) DescribeCaptchaIpCityWithCallback(request *DescribeCaptchaIpCityRequest, callback func(response *DescribeCaptchaIpCityResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCaptchaIpCityResponse
		var err error
		defer close(result)
		response, err = client.DescribeCaptchaIpCity(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

type DescribeCaptchaIpCityRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SourceIp        string           `position:"Query" name:"SourceIp"`
	ConfigName      string           `position:"Query" name:"ConfigName"`
	Time            string           `position:"Query" name:"Time"`
	Type            string           `position:"Query" name:"Type"`
}

type DescribeCaptchaIpCityResponse struct {
	*responses.BaseResponse
	RequestId     string `json:"RequestId" xml:"RequestId"`
	BizCode       string `json:"BizCode" xml:"BizCode"`
	HasData       bool   `json:"HasData" xml:"HasData"`
	CaptchaCities []struct {
		Location string `json:"Location" xml:"Location"`
		Lat      string `json:"Lat" xml:"Lat"`
		Lng      string `json:"Lng" xml:"Lng"`
		Pv       int    `json:"Pv" xml:"Pv"`
	} `json:"CaptchaCities" xml:"CaptchaCities"`
	CaptchaIps []struct {
		Ip    string `json:"Ip" xml:"Ip"`
		Value int    `json:"Value" xml:"Value"`
	} `json:"CaptchaIps" xml:"CaptchaIps"`
}

func CreateDescribeCaptchaIpCityRequest() (request *DescribeCaptchaIpCityRequest) {
	request = &DescribeCaptchaIpCityRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("afs", "2018-01-12", "DescribeCaptchaIpCity", "", "")
	return
}

func CreateDescribeCaptchaIpCityResponse() (response *DescribeCaptchaIpCityResponse) {
	response = &DescribeCaptchaIpCityResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
