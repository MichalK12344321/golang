# \CollectionApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCollection**](CollectionApi.md#GetCollection) | **Get** /collection | Get collection list
[**GetCollectionId**](CollectionApi.md#GetCollectionId) | **Get** /collection/{id} | Get collection
[**GetCollectionRunRunId**](CollectionApi.md#GetCollectionRunRunId) | **Get** /collection/run/{runId} | Get run details
[**GetCollectionRunRunIdArchive**](CollectionApi.md#GetCollectionRunRunIdArchive) | **Get** /collection/run/{runId}/archive | File


# **GetCollection**
> []LcaInternalPkgDtoCollectionDto GetCollection(ctx, optional)
Get collection list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CollectionApiGetCollectionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionApiGetCollectionOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| number of items to retrieve | [default to 5]
 **cursor** | **optional.String**| paging cursor | 
 **statuses** | **optional.String**| status filter (comma separated) | 

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

# **GetCollectionRunRunId**
> LcaInternalPkgDtoRunDto GetCollectionRunRunId(ctx, runId)
Get run details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **runId** | **string**| Run id | 

### Return type

[**LcaInternalPkgDtoRunDto**](lca_internal_pkg_dto.RunDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCollectionRunRunIdArchive**
> GetCollectionRunRunIdArchive(ctx, runId)
File

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **runId** | **string**| Run id | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

