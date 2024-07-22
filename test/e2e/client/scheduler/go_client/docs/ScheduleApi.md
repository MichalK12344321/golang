# \ScheduleApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostScheduleGo**](ScheduleApi.md#PostScheduleGo) | **Post** /schedule/go | Schedule Go collection
[**PostScheduleSsh**](ScheduleApi.md#PostScheduleSsh) | **Post** /schedule/ssh | Schedule SSH collection
[**PostTerminate**](ScheduleApi.md#PostTerminate) | **Post** /terminate | Terminate log collection


# **PostScheduleGo**
> LcaInternalPkgDtoScheduleResponseDto PostScheduleGo(ctx, body)
Schedule Go collection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LcaInternalPkgDtoScheduleGoCollectionDto**](LcaInternalPkgDtoScheduleGoCollectionDto.md)| Request | 

### Return type

[**LcaInternalPkgDtoScheduleResponseDto**](lca_internal_pkg_dto.ScheduleResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostScheduleSsh**
> LcaInternalPkgDtoScheduleResponseDto PostScheduleSsh(ctx, body)
Schedule SSH collection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LcaInternalPkgDtoScheduleSshCollectionDto**](LcaInternalPkgDtoScheduleSshCollectionDto.md)| Request | 

### Return type

[**LcaInternalPkgDtoScheduleResponseDto**](lca_internal_pkg_dto.ScheduleResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostTerminate**
> LcaInternalPkgDtoTerminateResponseDto PostTerminate(ctx, body)
Terminate log collection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LcaInternalPkgDtoTerminateRequestDto**](LcaInternalPkgDtoTerminateRequestDto.md)| Terminate request | 

### Return type

[**LcaInternalPkgDtoTerminateResponseDto**](lca_internal_pkg_dto.TerminateResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

