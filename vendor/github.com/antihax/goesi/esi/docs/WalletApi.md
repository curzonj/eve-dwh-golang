# \WalletApi

All URIs are relative to *https://esi.tech.ccp.is*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCharactersCharacterIdWallet**](WalletApi.md#GetCharactersCharacterIdWallet) | **Get** /v1/characters/{character_id}/wallet/ | Get a character&#39;s wallet balance
[**GetCharactersCharacterIdWalletJournal**](WalletApi.md#GetCharactersCharacterIdWalletJournal) | **Get** /v2/characters/{character_id}/wallet/journal/ | Get character wallet journal
[**GetCharactersCharacterIdWalletTransactions**](WalletApi.md#GetCharactersCharacterIdWalletTransactions) | **Get** /v1/characters/{character_id}/wallet/transactions/ | Get wallet transactions
[**GetCorporationsCorporationIdWallets**](WalletApi.md#GetCorporationsCorporationIdWallets) | **Get** /v1/corporations/{corporation_id}/wallets/ | Returns a corporation&#39;s wallet balance
[**GetCorporationsCorporationIdWalletsDivisionJournal**](WalletApi.md#GetCorporationsCorporationIdWalletsDivisionJournal) | **Get** /v1/corporations/{corporation_id}/wallets/{division}/journal/ | Get corporation wallet journal
[**GetCorporationsCorporationIdWalletsDivisionTransactions**](WalletApi.md#GetCorporationsCorporationIdWalletsDivisionTransactions) | **Get** /v1/corporations/{corporation_id}/wallets/{division}/transactions/ | Get corporation wallet transactions


# **GetCharactersCharacterIdWallet**
> float32 GetCharactersCharacterIdWallet(ctx, characterId, optional)
Get a character's wallet balance

Returns a character's wallet balance  ---  This route is cached for up to 120 seconds

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
 **token** | **string**| Access token to use if unable to set a header | 
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

**float32**

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCharactersCharacterIdWalletJournal**
> []GetCharactersCharacterIdWalletJournal200Ok GetCharactersCharacterIdWalletJournal(ctx, characterId, optional)
Get character wallet journal

Retrieve character wallet journal  ---  This route is cached for up to 3600 seconds

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
 **fromId** | **int64**| Only show journal entries happened before the transaction referenced by this id | 
 **token** | **string**| Access token to use if unable to set a header | 
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

[**[]GetCharactersCharacterIdWalletJournal200Ok**](get_characters_character_id_wallet_journal_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCharactersCharacterIdWalletTransactions**
> []GetCharactersCharacterIdWalletTransactions200Ok GetCharactersCharacterIdWalletTransactions(ctx, characterId, optional)
Get wallet transactions

Get wallet transactions of a character  ---  This route is cached for up to 3600 seconds

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
 **fromId** | **int64**| Only show transactions happened before the one referenced by this id | 
 **token** | **string**| Access token to use if unable to set a header | 
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

[**[]GetCharactersCharacterIdWalletTransactions200Ok**](get_characters_character_id_wallet_transactions_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCorporationsCorporationIdWallets**
> []GetCorporationsCorporationIdWallets200Ok GetCorporationsCorporationIdWallets(ctx, corporationId, optional)
Returns a corporation's wallet balance

Get a corporation's wallets  ---  This route is cached for up to 300 seconds

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
 **token** | **string**| Access token to use if unable to set a header | 
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

[**[]GetCorporationsCorporationIdWallets200Ok**](get_corporations_corporation_id_wallets_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCorporationsCorporationIdWalletsDivisionJournal**
> []GetCorporationsCorporationIdWalletsDivisionJournal200Ok GetCorporationsCorporationIdWalletsDivisionJournal(ctx, corporationId, division, optional)
Get corporation wallet journal

Retrieve corporation wallet journal  ---  This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **corporationId** | **int32**| An EVE corporation ID | 
  **division** | **int32**| Wallet key of the division to fetch journals from | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **corporationId** | **int32**| An EVE corporation ID | 
 **division** | **int32**| Wallet key of the division to fetch journals from | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **fromId** | **int64**| Only show journal entries happened before the transaction referenced by this id | 
 **token** | **string**| Access token to use if unable to set a header | 
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

[**[]GetCorporationsCorporationIdWalletsDivisionJournal200Ok**](get_corporations_corporation_id_wallets_division_journal_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCorporationsCorporationIdWalletsDivisionTransactions**
> []GetCorporationsCorporationIdWalletsDivisionTransactions200Ok GetCorporationsCorporationIdWalletsDivisionTransactions(ctx, corporationId, division, optional)
Get corporation wallet transactions

Get wallet transactions of a corporation  ---  This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **corporationId** | **int32**| An EVE corporation ID | 
  **division** | **int32**| Wallet key of the division to fetch journals from | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **corporationId** | **int32**| An EVE corporation ID | 
 **division** | **int32**| Wallet key of the division to fetch journals from | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **fromId** | **int64**| Only show journal entries happened before the transaction referenced by this id | 
 **token** | **string**| Access token to use if unable to set a header | 
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

[**[]GetCorporationsCorporationIdWalletsDivisionTransactions200Ok**](get_corporations_corporation_id_wallets_division_transactions_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

