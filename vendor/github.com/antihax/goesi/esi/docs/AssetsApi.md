# \AssetsApi

All URIs are relative to *https://esi.tech.ccp.is*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCharactersCharacterIdAssets**](AssetsApi.md#GetCharactersCharacterIdAssets) | **Get** /v1/characters/{character_id}/assets/ | Get character assets
[**GetCorporationsCorporationIdAssets**](AssetsApi.md#GetCorporationsCorporationIdAssets) | **Get** /v1/corporations/{corporation_id}/assets/ | Get corporation assets
[**PostCharactersCharacterIdAssetsLocations**](AssetsApi.md#PostCharactersCharacterIdAssetsLocations) | **Post** /v1/characters/{character_id}/assets/locations/ | Get character asset locations
[**PostCharactersCharacterIdAssetsNames**](AssetsApi.md#PostCharactersCharacterIdAssetsNames) | **Post** /v1/characters/{character_id}/assets/names/ | Get character asset names


# **GetCharactersCharacterIdAssets**
> []GetCharactersCharacterIdAssets200Ok GetCharactersCharacterIdAssets(ctx, characterId, optional)
Get character assets

Return a list of the characters assets  ---  This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **characterId** | **int32**| An EVE character ID | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **characterId** | **int32**| An EVE character ID | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **page** | **int32**| Which page of results to return | [default to 1]
 **token** | **string**| Access token to use if unable to set a header | 
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

[**[]GetCharactersCharacterIdAssets200Ok**](get_characters_character_id_assets_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCorporationsCorporationIdAssets**
> []GetCorporationsCorporationIdAssets200Ok GetCorporationsCorporationIdAssets(ctx, corporationId, optional)
Get corporation assets

Return a list of the corporation assets  ---  This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **corporationId** | **int32**| An EVE corporation ID | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **corporationId** | **int32**| An EVE corporation ID | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **page** | **int32**| Which page of results to return | [default to 1]
 **token** | **string**| Access token to use if unable to set a header | 
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

[**[]GetCorporationsCorporationIdAssets200Ok**](get_corporations_corporation_id_assets_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostCharactersCharacterIdAssetsLocations**
> []PostCharactersCharacterIdAssetsLocations200Ok PostCharactersCharacterIdAssetsLocations(ctx, characterId, itemIds, optional)
Get character asset locations

Return locations for a set of item ids, which you can get from character assets endpoint. Coordinates for items in hangars or stations are set to (0,0,0)  --- 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **characterId** | **int32**| An EVE character ID | 
  **itemIds** | **[]int64**| A list of item ids | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **characterId** | **int32**| An EVE character ID | 
 **itemIds** | **[]int64**| A list of item ids | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **token** | **string**| Access token to use if unable to set a header | 
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

[**[]PostCharactersCharacterIdAssetsLocations200Ok**](post_characters_character_id_assets_locations_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostCharactersCharacterIdAssetsNames**
> []PostCharactersCharacterIdAssetsNames200Ok PostCharactersCharacterIdAssetsNames(ctx, characterId, itemIds, optional)
Get character asset names

Return names for a set of item ids, which you can get from character assets endpoint. Typically used for items that can customize names, like containers or ships.  --- 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **characterId** | **int32**| An EVE character ID | 
  **itemIds** | **[]int64**| A list of item ids | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **characterId** | **int32**| An EVE character ID | 
 **itemIds** | **[]int64**| A list of item ids | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **token** | **string**| Access token to use if unable to set a header | 
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

[**[]PostCharactersCharacterIdAssetsNames200Ok**](post_characters_character_id_assets_names_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

