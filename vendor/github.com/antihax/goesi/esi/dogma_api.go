/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.6.0
 *
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package esi

import (
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/context"

	"encoding/json"
	"fmt"

	"github.com/mailru/easyjson"
)

// Linger please
var (
	_ context.Context
)

type DogmaApiService service

/* DogmaApiService Get attributes
Get a list of dogma attribute ids  ---  This route expires daily at 11:05

* @param ctx context.Context Authentication Context
@param optional (nil or map[string]interface{}) with one or more of:
    @param "datasource" (string) The server name you would like data from
    @param "userAgent" (string) Client identifier, takes precedence over headers
    @param "xUserAgent" (string) Client identifier, takes precedence over User-Agent
@return []int32*/
func (a *DogmaApiService) GetDogmaAttributes(ctx context.Context, localVarOptionals map[string]interface{}) ([]int32, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     []int32
	)

	// create path and map variables
	localVarPath := a.client.basePath + "/v1/dogma/attributes/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["datasource"], "string", "datasource"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["userAgent"], "string", "userAgent"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["xUserAgent"], "string", "xUserAgent"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["datasource"].(string); localVarOk {
		localVarQueryParams.Add("datasource", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["userAgent"].(string); localVarOk {
		localVarQueryParams.Add("user_agent", parameterToString(localVarTempParam, ""))
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarTempParam, localVarOk := localVarOptionals["xUserAgent"].(string); localVarOk {
		localVarHeaderParams["X-User-Agent"] = parameterToString(localVarTempParam, "")
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
	}
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* DogmaApiService Get attribute information
Get information on a dogma attribute  ---  This route expires daily at 11:05

* @param ctx context.Context Authentication Context
@param attributeId A dogma attribute ID
@param optional (nil or map[string]interface{}) with one or more of:
    @param "datasource" (string) The server name you would like data from
    @param "userAgent" (string) Client identifier, takes precedence over headers
    @param "xUserAgent" (string) Client identifier, takes precedence over User-Agent
@return GetDogmaAttributesAttributeIdOk*/
func (a *DogmaApiService) GetDogmaAttributesAttributeId(ctx context.Context, attributeId int32, localVarOptionals map[string]interface{}) (GetDogmaAttributesAttributeIdOk, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     GetDogmaAttributesAttributeIdOk
	)

	// create path and map variables
	localVarPath := a.client.basePath + "/v1/dogma/attributes/{attribute_id}/"
	localVarPath = strings.Replace(localVarPath, "{"+"attribute_id"+"}", fmt.Sprintf("%v", attributeId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["datasource"], "string", "datasource"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["userAgent"], "string", "userAgent"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["xUserAgent"], "string", "xUserAgent"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["datasource"].(string); localVarOk {
		localVarQueryParams.Add("datasource", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["userAgent"].(string); localVarOk {
		localVarQueryParams.Add("user_agent", parameterToString(localVarTempParam, ""))
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarTempParam, localVarOk := localVarOptionals["xUserAgent"].(string); localVarOk {
		localVarHeaderParams["X-User-Agent"] = parameterToString(localVarTempParam, "")
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
	}
	if err = easyjson.UnmarshalFromReader(localVarHttpResponse.Body, &successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* DogmaApiService Get effects
Get a list of dogma effect ids  ---  This route expires daily at 11:05

* @param ctx context.Context Authentication Context
@param optional (nil or map[string]interface{}) with one or more of:
    @param "datasource" (string) The server name you would like data from
    @param "userAgent" (string) Client identifier, takes precedence over headers
    @param "xUserAgent" (string) Client identifier, takes precedence over User-Agent
@return []int32*/
func (a *DogmaApiService) GetDogmaEffects(ctx context.Context, localVarOptionals map[string]interface{}) ([]int32, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     []int32
	)

	// create path and map variables
	localVarPath := a.client.basePath + "/v1/dogma/effects/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["datasource"], "string", "datasource"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["userAgent"], "string", "userAgent"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["xUserAgent"], "string", "xUserAgent"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["datasource"].(string); localVarOk {
		localVarQueryParams.Add("datasource", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["userAgent"].(string); localVarOk {
		localVarQueryParams.Add("user_agent", parameterToString(localVarTempParam, ""))
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarTempParam, localVarOk := localVarOptionals["xUserAgent"].(string); localVarOk {
		localVarHeaderParams["X-User-Agent"] = parameterToString(localVarTempParam, "")
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
	}
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* DogmaApiService Get effect information
Get information on a dogma effect  ---  This route expires daily at 11:05

* @param ctx context.Context Authentication Context
@param effectId A dogma effect ID
@param optional (nil or map[string]interface{}) with one or more of:
    @param "datasource" (string) The server name you would like data from
    @param "userAgent" (string) Client identifier, takes precedence over headers
    @param "xUserAgent" (string) Client identifier, takes precedence over User-Agent
@return GetDogmaEffectsEffectIdOk*/
func (a *DogmaApiService) GetDogmaEffectsEffectId(ctx context.Context, effectId int32, localVarOptionals map[string]interface{}) (GetDogmaEffectsEffectIdOk, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     GetDogmaEffectsEffectIdOk
	)

	// create path and map variables
	localVarPath := a.client.basePath + "/v2/dogma/effects/{effect_id}/"
	localVarPath = strings.Replace(localVarPath, "{"+"effect_id"+"}", fmt.Sprintf("%v", effectId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["datasource"], "string", "datasource"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["userAgent"], "string", "userAgent"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["xUserAgent"], "string", "xUserAgent"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["datasource"].(string); localVarOk {
		localVarQueryParams.Add("datasource", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["userAgent"].(string); localVarOk {
		localVarQueryParams.Add("user_agent", parameterToString(localVarTempParam, ""))
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarTempParam, localVarOk := localVarOptionals["xUserAgent"].(string); localVarOk {
		localVarHeaderParams["X-User-Agent"] = parameterToString(localVarTempParam, "")
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
	}
	if err = easyjson.UnmarshalFromReader(localVarHttpResponse.Body, &successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}
