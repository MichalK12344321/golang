# \CollectionApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCollection**](CollectionApi.md#GetCollection) | **Get** /collection | Get collection list
[**GetCollectionId**](CollectionApi.md#GetCollectionId) | **Get** /collection/{id} | Get collection
[**GetCollectionIdArchive**](CollectionApi.md#GetCollectionIdArchive) | **Get** /collection/{id}/archive | File


# **GetCollection**
> []LcaInternalPkgDtoCollectionDto GetCollection(ctx, )
Get collection list

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]LcaInternalPkgDtoCollectionDto**](lca_internal_pkg_dto.CollectionDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCollectionId**
> LcaInternalPkgDtoCollectionDto GetCollectionId(ctx, id)
Get collection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Id of the Collection | 

### Return type

[**LcaInternalPkgDtoCollectionDto**](lca_internal_pkg_dto.CollectionDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCollectionIdArchive**
> GetCollectionIdArchive(ctx, id)
File

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Id of the Collection | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

