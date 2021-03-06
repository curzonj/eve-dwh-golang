# \RoutesApi

All URIs are relative to *https://esi.tech.ccp.is*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRouteOriginDestination**](RoutesApi.md#GetRouteOriginDestination) | **Get** /v1/route/{origin}/{destination}/ | Get route


# **GetRouteOriginDestination**
> []int32 GetRouteOriginDestination(destination, origin, optional)
Get route

Get the systems between origin and destination  ---  This route is cached for up to 86400 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **destination** | **int32**| destination solar system ID | 
  **origin** | **int32**| origin solar system ID | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **destination** | **int32**| destination solar system ID | 
 **origin** | **int32**| origin solar system ID | 
 **avoid** | [**[]int32**](int32.md)| avoid solar system ID(s) | 
 **connections** | [**[][]int32**]([]int32.md)| connected solar system pairs | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **flag** | **string**| route security preference | [default to shortest]
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

**[]int32**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

