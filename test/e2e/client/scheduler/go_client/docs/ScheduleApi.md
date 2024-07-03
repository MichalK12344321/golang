# \ScheduleApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostSchedule**](ScheduleApi.md#PostSchedule) | **Post** /schedule | Schedule log collection


# **PostSchedule**
> LcaInternalPkgDtoScheduleResponseDto PostSchedule(ctx, body)
Schedule log collection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LcaInternalPkgDtoScheduleRequestDto**](LcaInternalPkgDtoScheduleRequestDto.md)| Schedule request | 

### Return type

[**LcaInternalPkgDtoScheduleResponseDto**](lca_internal_pkg_dto.ScheduleResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

