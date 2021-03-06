package baas

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

// QueryConsortiumDeletable invokes the baas.QueryConsortiumDeletable API synchronously
// api document: https://help.aliyun.com/api/baas/queryconsortiumdeletable.html
func (client *Client) QueryConsortiumDeletable(request *QueryConsortiumDeletableRequest) (response *QueryConsortiumDeletableResponse, err error) {
	response = CreateQueryConsortiumDeletableResponse()
	err = client.DoAction(request, response)
	return
}

// QueryConsortiumDeletableWithChan invokes the baas.QueryConsortiumDeletable API asynchronously
// api document: https://help.aliyun.com/api/baas/queryconsortiumdeletable.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryConsortiumDeletableWithChan(request *QueryConsortiumDeletableRequest) (<-chan *QueryConsortiumDeletableResponse, <-chan error) {
	responseChan := make(chan *QueryConsortiumDeletableResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryConsortiumDeletable(request)
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

// QueryConsortiumDeletableWithCallback invokes the baas.QueryConsortiumDeletable API asynchronously
// api document: https://help.aliyun.com/api/baas/queryconsortiumdeletable.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryConsortiumDeletableWithCallback(request *QueryConsortiumDeletableRequest, callback func(response *QueryConsortiumDeletableResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryConsortiumDeletableResponse
		var err error
		defer close(result)
		response, err = client.QueryConsortiumDeletable(request)
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

// QueryConsortiumDeletableRequest is the request struct for api QueryConsortiumDeletable
type QueryConsortiumDeletableRequest struct {
	*requests.RpcRequest
	Location     string `position:"Body" name:"Location"`
	ConsortiumId string `position:"Query" name:"ConsortiumId"`
}

// QueryConsortiumDeletableResponse is the response struct for api QueryConsortiumDeletable
type QueryConsortiumDeletableResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateQueryConsortiumDeletableRequest creates a request to invoke QueryConsortiumDeletable API
func CreateQueryConsortiumDeletableRequest() (request *QueryConsortiumDeletableRequest) {
	request = &QueryConsortiumDeletableRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Baas", "2018-07-31", "QueryConsortiumDeletable", "", "")
	return
}

// CreateQueryConsortiumDeletableResponse creates a response to parse from QueryConsortiumDeletable response
func CreateQueryConsortiumDeletableResponse() (response *QueryConsortiumDeletableResponse) {
	response = &QueryConsortiumDeletableResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
