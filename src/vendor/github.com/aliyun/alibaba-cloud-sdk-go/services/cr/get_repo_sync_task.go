package cr

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

// GetRepoSyncTask invokes the cr.GetRepoSyncTask API synchronously
// api document: https://help.aliyun.com/api/cr/getreposynctask.html
func (client *Client) GetRepoSyncTask(request *GetRepoSyncTaskRequest) (response *GetRepoSyncTaskResponse, err error) {
	response = CreateGetRepoSyncTaskResponse()
	err = client.DoAction(request, response)
	return
}

// GetRepoSyncTaskWithChan invokes the cr.GetRepoSyncTask API asynchronously
// api document: https://help.aliyun.com/api/cr/getreposynctask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetRepoSyncTaskWithChan(request *GetRepoSyncTaskRequest) (<-chan *GetRepoSyncTaskResponse, <-chan error) {
	responseChan := make(chan *GetRepoSyncTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetRepoSyncTask(request)
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

// GetRepoSyncTaskWithCallback invokes the cr.GetRepoSyncTask API asynchronously
// api document: https://help.aliyun.com/api/cr/getreposynctask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetRepoSyncTaskWithCallback(request *GetRepoSyncTaskRequest, callback func(response *GetRepoSyncTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetRepoSyncTaskResponse
		var err error
		defer close(result)
		response, err = client.GetRepoSyncTask(request)
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

// GetRepoSyncTaskRequest is the request struct for api GetRepoSyncTask
type GetRepoSyncTaskRequest struct {
	*requests.RoaRequest
	RepoNamespace string `position:"Path" name:"RepoNamespace"`
	RepoName      string `position:"Path" name:"RepoName"`
	SyncTaskId    string `position:"Path" name:"SyncTaskId"`
}

// GetRepoSyncTaskResponse is the response struct for api GetRepoSyncTask
type GetRepoSyncTaskResponse struct {
	*responses.BaseResponse
}

// CreateGetRepoSyncTaskRequest creates a request to invoke GetRepoSyncTask API
func CreateGetRepoSyncTaskRequest() (request *GetRepoSyncTaskRequest) {
	request = &GetRepoSyncTaskRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("cr", "2016-06-07", "GetRepoSyncTask", "/repos/[RepoNamespace]/[RepoName]/syncTasks/[SyncTaskId]", "cr", "openAPI")
	request.Method = requests.GET
	return
}

// CreateGetRepoSyncTaskResponse creates a response to parse from GetRepoSyncTask response
func CreateGetRepoSyncTaskResponse() (response *GetRepoSyncTaskResponse) {
	response = &GetRepoSyncTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}