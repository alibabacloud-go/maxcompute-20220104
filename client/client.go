// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ApplyComputeQuotaPlanResponseBody struct {
	// The data returned.
	//
	// example:
	//
	// success
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// The error code.
	//
	// example:
	//
	// OBJECT_NOT_EXIST
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// this quota plan is not exist.
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// The HTTP status code.
	//
	// - 1xx: informational response. The request is received and is being processed.
	//
	// - 2xx: success. The request is successfully received, understood, and accepted by the server.
	//
	// - 3xx: redirection. The request is redirected, and further actions are required to complete the request.
	//
	// - 4xx: client error. The request contains invalid request parameters or syntaxes, or specific request conditions cannot be met.
	//
	// - 5xx: server error. The server cannot meet requirements due to other reasons.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0b87b7a316654730544735643e9200
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ApplyComputeQuotaPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApplyComputeQuotaPlanResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyComputeQuotaPlanResponseBody) SetData(v string) *ApplyComputeQuotaPlanResponseBody {
	s.Data = &v
	return s
}

func (s *ApplyComputeQuotaPlanResponseBody) SetErrorCode(v string) *ApplyComputeQuotaPlanResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ApplyComputeQuotaPlanResponseBody) SetErrorMsg(v string) *ApplyComputeQuotaPlanResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ApplyComputeQuotaPlanResponseBody) SetHttpCode(v int32) *ApplyComputeQuotaPlanResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ApplyComputeQuotaPlanResponseBody) SetRequestId(v string) *ApplyComputeQuotaPlanResponseBody {
	s.RequestId = &v
	return s
}

type ApplyComputeQuotaPlanResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApplyComputeQuotaPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApplyComputeQuotaPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s ApplyComputeQuotaPlanResponse) GoString() string {
	return s.String()
}

func (s *ApplyComputeQuotaPlanResponse) SetHeaders(v map[string]*string) *ApplyComputeQuotaPlanResponse {
	s.Headers = v
	return s
}

func (s *ApplyComputeQuotaPlanResponse) SetStatusCode(v int32) *ApplyComputeQuotaPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyComputeQuotaPlanResponse) SetBody(v *ApplyComputeQuotaPlanResponseBody) *ApplyComputeQuotaPlanResponse {
	s.Body = v
	return s
}

type CreateComputeQuotaPlanRequest struct {
	// The name of quota plan.
	//
	// This parameter is required.
	//
	// example:
	//
	// planA
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The parameters of quota plan.
	Quota *CreateComputeQuotaPlanRequestQuota `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
}

func (s CreateComputeQuotaPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateComputeQuotaPlanRequest) GoString() string {
	return s.String()
}

func (s *CreateComputeQuotaPlanRequest) SetName(v string) *CreateComputeQuotaPlanRequest {
	s.Name = &v
	return s
}

func (s *CreateComputeQuotaPlanRequest) SetQuota(v *CreateComputeQuotaPlanRequestQuota) *CreateComputeQuotaPlanRequest {
	s.Quota = v
	return s
}

type CreateComputeQuotaPlanRequestQuota struct {
	// The parameters of level-1 quota.
	Parameter *CreateComputeQuotaPlanRequestQuotaParameter `json:"parameter,omitempty" xml:"parameter,omitempty" type:"Struct"`
	// The list of level-2 quotas.
	SubQuotaInfoList []*CreateComputeQuotaPlanRequestQuotaSubQuotaInfoList `json:"subQuotaInfoList,omitempty" xml:"subQuotaInfoList,omitempty" type:"Repeated"`
}

func (s CreateComputeQuotaPlanRequestQuota) String() string {
	return tea.Prettify(s)
}

func (s CreateComputeQuotaPlanRequestQuota) GoString() string {
	return s.String()
}

func (s *CreateComputeQuotaPlanRequestQuota) SetParameter(v *CreateComputeQuotaPlanRequestQuotaParameter) *CreateComputeQuotaPlanRequestQuota {
	s.Parameter = v
	return s
}

func (s *CreateComputeQuotaPlanRequestQuota) SetSubQuotaInfoList(v []*CreateComputeQuotaPlanRequestQuotaSubQuotaInfoList) *CreateComputeQuotaPlanRequestQuota {
	s.SubQuotaInfoList = v
	return s
}

type CreateComputeQuotaPlanRequestQuotaParameter struct {
	// The value of elastic Reserved CUs in the level-1 quota.
	//
	// > The default value is 0. The maximum value of this parameter must be equal to the number of subscription-based reserved CUs and cannot exceed 10,000 CUs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 50
	ElasticReservedCU *int64 `json:"elasticReservedCU,omitempty" xml:"elasticReservedCU,omitempty"`
}

func (s CreateComputeQuotaPlanRequestQuotaParameter) String() string {
	return tea.Prettify(s)
}

func (s CreateComputeQuotaPlanRequestQuotaParameter) GoString() string {
	return s.String()
}

func (s *CreateComputeQuotaPlanRequestQuotaParameter) SetElasticReservedCU(v int64) *CreateComputeQuotaPlanRequestQuotaParameter {
	s.ElasticReservedCU = &v
	return s
}

type CreateComputeQuotaPlanRequestQuotaSubQuotaInfoList struct {
	// The nickname of the level-2 quota.
	//
	// This parameter is required.
	//
	// example:
	//
	// os_ComputeQuota
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	// The parameters of the level-2 quota.
	Parameter *CreateComputeQuotaPlanRequestQuotaSubQuotaInfoListParameter `json:"parameter,omitempty" xml:"parameter,omitempty" type:"Struct"`
}

func (s CreateComputeQuotaPlanRequestQuotaSubQuotaInfoList) String() string {
	return tea.Prettify(s)
}

func (s CreateComputeQuotaPlanRequestQuotaSubQuotaInfoList) GoString() string {
	return s.String()
}

func (s *CreateComputeQuotaPlanRequestQuotaSubQuotaInfoList) SetNickName(v string) *CreateComputeQuotaPlanRequestQuotaSubQuotaInfoList {
	s.NickName = &v
	return s
}

func (s *CreateComputeQuotaPlanRequestQuotaSubQuotaInfoList) SetParameter(v *CreateComputeQuotaPlanRequestQuotaSubQuotaInfoListParameter) *CreateComputeQuotaPlanRequestQuotaSubQuotaInfoList {
	s.Parameter = v
	return s
}

type CreateComputeQuotaPlanRequestQuotaSubQuotaInfoListParameter struct {
	// The value of elastic Reserved CUs.
	//
	// > The total number of elastically reserved CUs in all the level-2 quotas is equal to the number of elastically reserved CUs in the level-1 quota.
	//
	// This parameter is required.
	//
	// example:
	//
	// 50
	ElasticReservedCU *int64 `json:"elasticReservedCU,omitempty" xml:"elasticReservedCU,omitempty"`
	// The value of maxCU in Reserved CUs.
	//
	// > The value of maxCU must be less than or equal to the value of maxCU in the level-1 quota that you purchased.
	//
	// This parameter is required.
	//
	// example:
	//
	// 50
	MaxCU *int64 `json:"maxCU,omitempty" xml:"maxCU,omitempty"`
	// The value of minCU in Reserved CUs.
	//
	// >
	//
	// >- The total value of minCU in all the level-2 quotas is equal to the value of minCU in the level-1 quota.
	//
	// >- The value of minCU must be less than or equal to the value of maxCU in the level-2 quota and less than or equal to the value of minCU in the level-1 quota that you purchased.
	//
	// This parameter is required.
	//
	// example:
	//
	// 50
	MinCU *int64 `json:"minCU,omitempty" xml:"minCU,omitempty"`
}

func (s CreateComputeQuotaPlanRequestQuotaSubQuotaInfoListParameter) String() string {
	return tea.Prettify(s)
}

func (s CreateComputeQuotaPlanRequestQuotaSubQuotaInfoListParameter) GoString() string {
	return s.String()
}

func (s *CreateComputeQuotaPlanRequestQuotaSubQuotaInfoListParameter) SetElasticReservedCU(v int64) *CreateComputeQuotaPlanRequestQuotaSubQuotaInfoListParameter {
	s.ElasticReservedCU = &v
	return s
}

func (s *CreateComputeQuotaPlanRequestQuotaSubQuotaInfoListParameter) SetMaxCU(v int64) *CreateComputeQuotaPlanRequestQuotaSubQuotaInfoListParameter {
	s.MaxCU = &v
	return s
}

func (s *CreateComputeQuotaPlanRequestQuotaSubQuotaInfoListParameter) SetMinCU(v int64) *CreateComputeQuotaPlanRequestQuotaSubQuotaInfoListParameter {
	s.MinCU = &v
	return s
}

type CreateComputeQuotaPlanResponseBody struct {
	// The returned data.
	//
	// example:
	//
	// success
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// The error code.
	//
	// example:
	//
	// OBJECT_NOT_EXIST
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// this quota is not exist.
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// The HTTP status code.
	//
	// - 1xx: informational response. The request is received and is being processed.
	//
	// - 2xx: success. The request is successfully received, understood, and accepted by the server.
	//
	// - 3xx: redirection. The request is redirected, and further actions are required to complete the request.
	//
	// - 4xx: client error. The request contains invalid request parameters or syntaxes, or specific request conditions cannot be met.
	//
	// - 5xx: server error. The server cannot meet requirements due to other reasons.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0b87b7e716665825896565060e87a4
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateComputeQuotaPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateComputeQuotaPlanResponseBody) GoString() string {
	return s.String()
}

func (s *CreateComputeQuotaPlanResponseBody) SetData(v string) *CreateComputeQuotaPlanResponseBody {
	s.Data = &v
	return s
}

func (s *CreateComputeQuotaPlanResponseBody) SetErrorCode(v string) *CreateComputeQuotaPlanResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateComputeQuotaPlanResponseBody) SetErrorMsg(v string) *CreateComputeQuotaPlanResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CreateComputeQuotaPlanResponseBody) SetHttpCode(v int32) *CreateComputeQuotaPlanResponseBody {
	s.HttpCode = &v
	return s
}

func (s *CreateComputeQuotaPlanResponseBody) SetRequestId(v string) *CreateComputeQuotaPlanResponseBody {
	s.RequestId = &v
	return s
}

type CreateComputeQuotaPlanResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateComputeQuotaPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateComputeQuotaPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateComputeQuotaPlanResponse) GoString() string {
	return s.String()
}

func (s *CreateComputeQuotaPlanResponse) SetHeaders(v map[string]*string) *CreateComputeQuotaPlanResponse {
	s.Headers = v
	return s
}

func (s *CreateComputeQuotaPlanResponse) SetStatusCode(v int32) *CreateComputeQuotaPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateComputeQuotaPlanResponse) SetBody(v *CreateComputeQuotaPlanResponseBody) *CreateComputeQuotaPlanResponse {
	s.Body = v
	return s
}

type CreateMmsDataSourceRequest struct {
	Config map[string]interface{} `json:"config,omitempty" xml:"config,omitempty"`
	// example:
	//
	// demo
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// vpc-uf6pc2vordian33gobzfr:cn-shanghai
	Networklink *string `json:"networklink,omitempty" xml:"networklink,omitempty"`
	// example:
	//
	// BIGQUERY
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateMmsDataSourceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMmsDataSourceRequest) GoString() string {
	return s.String()
}

func (s *CreateMmsDataSourceRequest) SetConfig(v map[string]interface{}) *CreateMmsDataSourceRequest {
	s.Config = v
	return s
}

func (s *CreateMmsDataSourceRequest) SetName(v string) *CreateMmsDataSourceRequest {
	s.Name = &v
	return s
}

func (s *CreateMmsDataSourceRequest) SetNetworklink(v string) *CreateMmsDataSourceRequest {
	s.Networklink = &v
	return s
}

func (s *CreateMmsDataSourceRequest) SetType(v string) *CreateMmsDataSourceRequest {
	s.Type = &v
	return s
}

type CreateMmsDataSourceResponseBody struct {
	Data *CreateMmsDataSourceResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// B42CA730-8187-50F1-9FE0-6733297036DB
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateMmsDataSourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMmsDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMmsDataSourceResponseBody) SetData(v *CreateMmsDataSourceResponseBodyData) *CreateMmsDataSourceResponseBody {
	s.Data = v
	return s
}

func (s *CreateMmsDataSourceResponseBody) SetRequestId(v string) *CreateMmsDataSourceResponseBody {
	s.RequestId = &v
	return s
}

type CreateMmsDataSourceResponseBodyData struct {
	// example:
	//
	// 18
	DataSourceId *int64 `json:"dataSourceId,omitempty" xml:"dataSourceId,omitempty"`
}

func (s CreateMmsDataSourceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateMmsDataSourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateMmsDataSourceResponseBodyData) SetDataSourceId(v int64) *CreateMmsDataSourceResponseBodyData {
	s.DataSourceId = &v
	return s
}

type CreateMmsDataSourceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMmsDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMmsDataSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMmsDataSourceResponse) GoString() string {
	return s.String()
}

func (s *CreateMmsDataSourceResponse) SetHeaders(v map[string]*string) *CreateMmsDataSourceResponse {
	s.Headers = v
	return s
}

func (s *CreateMmsDataSourceResponse) SetStatusCode(v int32) *CreateMmsDataSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMmsDataSourceResponse) SetBody(v *CreateMmsDataSourceResponseBody) *CreateMmsDataSourceResponse {
	s.Body = v
	return s
}

type CreateMmsFetchMetadataJobResponseBody struct {
	Data *CreateMmsFetchMetadataJobResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// CC4D05E8-0613-5A8E-9339-A0EBD097A69E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateMmsFetchMetadataJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMmsFetchMetadataJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMmsFetchMetadataJobResponseBody) SetData(v *CreateMmsFetchMetadataJobResponseBodyData) *CreateMmsFetchMetadataJobResponseBody {
	s.Data = v
	return s
}

func (s *CreateMmsFetchMetadataJobResponseBody) SetRequestId(v string) *CreateMmsFetchMetadataJobResponseBody {
	s.RequestId = &v
	return s
}

type CreateMmsFetchMetadataJobResponseBodyData struct {
	// example:
	//
	// 1000002
	ScanId *int64 `json:"scanId,omitempty" xml:"scanId,omitempty"`
}

func (s CreateMmsFetchMetadataJobResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateMmsFetchMetadataJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateMmsFetchMetadataJobResponseBodyData) SetScanId(v int64) *CreateMmsFetchMetadataJobResponseBodyData {
	s.ScanId = &v
	return s
}

type CreateMmsFetchMetadataJobResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMmsFetchMetadataJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMmsFetchMetadataJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMmsFetchMetadataJobResponse) GoString() string {
	return s.String()
}

func (s *CreateMmsFetchMetadataJobResponse) SetHeaders(v map[string]*string) *CreateMmsFetchMetadataJobResponse {
	s.Headers = v
	return s
}

func (s *CreateMmsFetchMetadataJobResponse) SetStatusCode(v int32) *CreateMmsFetchMetadataJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMmsFetchMetadataJobResponse) SetBody(v *CreateMmsFetchMetadataJobResponseBody) *CreateMmsFetchMetadataJobResponse {
	s.Body = v
	return s
}

type CreateMmsJobRequest struct {
	ColumnMapping      map[string]*string     `json:"columnMapping,omitempty" xml:"columnMapping,omitempty"`
	DstDbName          *string                `json:"dstDbName,omitempty" xml:"dstDbName,omitempty"`
	DstSchemaName      *string                `json:"dstSchemaName,omitempty" xml:"dstSchemaName,omitempty"`
	EnableVerification *bool                  `json:"enableVerification,omitempty" xml:"enableVerification,omitempty"`
	Eta                *string                `json:"eta,omitempty" xml:"eta,omitempty"`
	Increment          *bool                  `json:"increment,omitempty" xml:"increment,omitempty"`
	Name               *string                `json:"name,omitempty" xml:"name,omitempty"`
	Others             map[string]interface{} `json:"others,omitempty" xml:"others,omitempty"`
	PartitionFilters   map[string]*string     `json:"partitionFilters,omitempty" xml:"partitionFilters,omitempty"`
	Partitions         []*int64               `json:"partitions,omitempty" xml:"partitions,omitempty" type:"Repeated"`
	SchemaOnly         *bool                  `json:"schemaOnly,omitempty" xml:"schemaOnly,omitempty"`
	SourceId           *int64                 `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	SourceName         *string                `json:"sourceName,omitempty" xml:"sourceName,omitempty"`
	SrcDbName          *string                `json:"srcDbName,omitempty" xml:"srcDbName,omitempty"`
	SrcSchemaName      *string                `json:"srcSchemaName,omitempty" xml:"srcSchemaName,omitempty"`
	TableBlackList     []*string              `json:"tableBlackList,omitempty" xml:"tableBlackList,omitempty" type:"Repeated"`
	TableMapping       map[string]*string     `json:"tableMapping,omitempty" xml:"tableMapping,omitempty"`
	TableWhiteList     []*string              `json:"tableWhiteList,omitempty" xml:"tableWhiteList,omitempty" type:"Repeated"`
	Tables             []*string              `json:"tables,omitempty" xml:"tables,omitempty" type:"Repeated"`
	// MOCK, HIVE: hive udtf task, HIVE_DATAX: hive datax task, COPY_TASK: odps Copy Task, ODPS_INSERT_OVERWRITE: odps simple insert overwrite task, MC2MC_VERIFY, OSS, HIVE_OSS
	TaskType *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
}

func (s CreateMmsJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMmsJobRequest) GoString() string {
	return s.String()
}

func (s *CreateMmsJobRequest) SetColumnMapping(v map[string]*string) *CreateMmsJobRequest {
	s.ColumnMapping = v
	return s
}

func (s *CreateMmsJobRequest) SetDstDbName(v string) *CreateMmsJobRequest {
	s.DstDbName = &v
	return s
}

func (s *CreateMmsJobRequest) SetDstSchemaName(v string) *CreateMmsJobRequest {
	s.DstSchemaName = &v
	return s
}

func (s *CreateMmsJobRequest) SetEnableVerification(v bool) *CreateMmsJobRequest {
	s.EnableVerification = &v
	return s
}

func (s *CreateMmsJobRequest) SetEta(v string) *CreateMmsJobRequest {
	s.Eta = &v
	return s
}

func (s *CreateMmsJobRequest) SetIncrement(v bool) *CreateMmsJobRequest {
	s.Increment = &v
	return s
}

func (s *CreateMmsJobRequest) SetName(v string) *CreateMmsJobRequest {
	s.Name = &v
	return s
}

func (s *CreateMmsJobRequest) SetOthers(v map[string]interface{}) *CreateMmsJobRequest {
	s.Others = v
	return s
}

func (s *CreateMmsJobRequest) SetPartitionFilters(v map[string]*string) *CreateMmsJobRequest {
	s.PartitionFilters = v
	return s
}

func (s *CreateMmsJobRequest) SetPartitions(v []*int64) *CreateMmsJobRequest {
	s.Partitions = v
	return s
}

func (s *CreateMmsJobRequest) SetSchemaOnly(v bool) *CreateMmsJobRequest {
	s.SchemaOnly = &v
	return s
}

func (s *CreateMmsJobRequest) SetSourceId(v int64) *CreateMmsJobRequest {
	s.SourceId = &v
	return s
}

func (s *CreateMmsJobRequest) SetSourceName(v string) *CreateMmsJobRequest {
	s.SourceName = &v
	return s
}

func (s *CreateMmsJobRequest) SetSrcDbName(v string) *CreateMmsJobRequest {
	s.SrcDbName = &v
	return s
}

func (s *CreateMmsJobRequest) SetSrcSchemaName(v string) *CreateMmsJobRequest {
	s.SrcSchemaName = &v
	return s
}

func (s *CreateMmsJobRequest) SetTableBlackList(v []*string) *CreateMmsJobRequest {
	s.TableBlackList = v
	return s
}

func (s *CreateMmsJobRequest) SetTableMapping(v map[string]*string) *CreateMmsJobRequest {
	s.TableMapping = v
	return s
}

func (s *CreateMmsJobRequest) SetTableWhiteList(v []*string) *CreateMmsJobRequest {
	s.TableWhiteList = v
	return s
}

func (s *CreateMmsJobRequest) SetTables(v []*string) *CreateMmsJobRequest {
	s.Tables = v
	return s
}

func (s *CreateMmsJobRequest) SetTaskType(v string) *CreateMmsJobRequest {
	s.TaskType = &v
	return s
}

type CreateMmsJobResponseBody struct {
	Data      *CreateMmsJobResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	RequestId *string                       `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateMmsJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMmsJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMmsJobResponseBody) SetData(v *CreateMmsJobResponseBodyData) *CreateMmsJobResponseBody {
	s.Data = v
	return s
}

func (s *CreateMmsJobResponseBody) SetRequestId(v string) *CreateMmsJobResponseBody {
	s.RequestId = &v
	return s
}

type CreateMmsJobResponseBodyData struct {
	AsyncTaskId *int64 `json:"asyncTaskId,omitempty" xml:"asyncTaskId,omitempty"`
}

func (s CreateMmsJobResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateMmsJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateMmsJobResponseBodyData) SetAsyncTaskId(v int64) *CreateMmsJobResponseBodyData {
	s.AsyncTaskId = &v
	return s
}

type CreateMmsJobResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMmsJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMmsJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMmsJobResponse) GoString() string {
	return s.String()
}

func (s *CreateMmsJobResponse) SetHeaders(v map[string]*string) *CreateMmsJobResponse {
	s.Headers = v
	return s
}

func (s *CreateMmsJobResponse) SetStatusCode(v int32) *CreateMmsJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMmsJobResponse) SetBody(v *CreateMmsJobResponseBody) *CreateMmsJobResponse {
	s.Body = v
	return s
}

type CreatePackageRequest struct {
	// The request body parameters.
	//
	// example:
	//
	// {
	//
	// "name": "test_packege",
	//
	//     "resourceList": {
	//
	//         "table": [
	//
	//             {
	//
	//                 "name": "table_name",
	//
	//                 "actions": [
	//
	//                     "Describe",
	//
	//                     "Select"
	//
	//                 ]
	//
	//             },
	//
	//             {
	//
	//                 "name": "table_name",
	//
	//                 "actions": [
	//
	//                     "Describe",
	//
	//                     "Select"
	//
	//                 ]
	//
	//             }
	//
	//         ],
	//
	//         "resource": [
	//
	//             {
	//
	//                 "name": "",
	//
	//                 "actions": []
	//
	//             },
	//
	//             {
	//
	//                 "name": "",
	//
	//                 "actions": []
	//
	//             }
	//
	//         ],
	//
	//         "function": [
	//
	//             {
	//
	//                 "name": "",
	//
	//                 "actions": []
	//
	//             },
	//
	//             {
	//
	//                 "name": "",
	//
	//                 "actions": []
	//
	//             }
	//
	//         ]
	//
	//     }
	//
	// }
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
	// Specifies whether to install the package.
	//
	// example:
	//
	// false
	IsInstall *bool `json:"isInstall,omitempty" xml:"isInstall,omitempty"`
}

func (s CreatePackageRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePackageRequest) GoString() string {
	return s.String()
}

func (s *CreatePackageRequest) SetBody(v string) *CreatePackageRequest {
	s.Body = &v
	return s
}

func (s *CreatePackageRequest) SetIsInstall(v bool) *CreatePackageRequest {
	s.IsInstall = &v
	return s
}

type CreatePackageResponseBody struct {
	// The returned data.
	//
	// example:
	//
	// success
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0bc3b4ab16684833172127321e2c25
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreatePackageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePackageResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePackageResponseBody) SetData(v string) *CreatePackageResponseBody {
	s.Data = &v
	return s
}

func (s *CreatePackageResponseBody) SetRequestId(v string) *CreatePackageResponseBody {
	s.RequestId = &v
	return s
}

type CreatePackageResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePackageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePackageResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePackageResponse) GoString() string {
	return s.String()
}

func (s *CreatePackageResponse) SetHeaders(v map[string]*string) *CreatePackageResponse {
	s.Headers = v
	return s
}

func (s *CreatePackageResponse) SetStatusCode(v int32) *CreatePackageResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePackageResponse) SetBody(v *CreatePackageResponseBody) *CreatePackageResponse {
	s.Body = v
	return s
}

type CreateProjectRequest struct {
	// The request body parameters.
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectRequest) SetBody(v string) *CreateProjectRequest {
	s.Body = &v
	return s
}

type CreateProjectResponseBody struct {
	// The returned result.
	//
	// example:
	//
	// success
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0b87b7a316654730544735643e9200
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProjectResponseBody) SetData(v string) *CreateProjectResponseBody {
	s.Data = &v
	return s
}

func (s *CreateProjectResponseBody) SetRequestId(v string) *CreateProjectResponseBody {
	s.RequestId = &v
	return s
}

type CreateProjectResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectResponse) GoString() string {
	return s.String()
}

func (s *CreateProjectResponse) SetHeaders(v map[string]*string) *CreateProjectResponse {
	s.Headers = v
	return s
}

func (s *CreateProjectResponse) SetStatusCode(v int32) *CreateProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateProjectResponse) SetBody(v *CreateProjectResponseBody) *CreateProjectResponse {
	s.Body = v
	return s
}

type CreateQuotaPlanRequest struct {
	// The request body parameters.
	//
	// example:
	//
	// { "name": "planA", // The quota is a level-1 quota. You can select only the fields that are related to the quota plan. "quota": { "name": "a", "nickName": "aaa_nick", "tenantId": "10001", "regionId": "cn-hangzhou", "parentId": "0", "cluster": "AT-ODPS-TEST3", "parameter": { "minCU": 40, "maxCU": 40, "adhocCU": 0, "elasticMinCU": 40, "elasticMaxCU": 40, "enablePreemptiveScheduling": false, "forceReservedMin":true, "enablePriority":false, "singleJobCULimit":100, "adhocQuotaBeginTimeInSec": 1345, "adhocQuotaEndTimeInSec": 1234, "ignoreAdhocQuota":false }, "subQuotaInfoList": [ { "nickName": "WlmFuxiSecondaryOnlineQuotaTest", "name": "WlmFuxiSecondaryOnlineQuotaTest", "type": "FUXI_ONLINE", "tenantId": "10001", "regionId": "cn-hangzhou", "cluster": "AT-ODPS-TEST3", "parameter": { "minCU": 40, "maxCU": 40, "adhocCU": 0, "elasticMinCU": 40, "elasticMaxCU": 40, "enablePreemptiveScheduling": false, "forceReservedMin":true, "enablePriority":false, "singleJobCULimit":100, "adhocQuotaBeginTimeInSec": 1345, "adhocQuotaEndTimeInSec": 1234, "ignoreAdhocQuota":false } } ] } }
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The ID of the tenant.
	//
	// example:
	//
	// 228451885265153
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreateQuotaPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateQuotaPlanRequest) GoString() string {
	return s.String()
}

func (s *CreateQuotaPlanRequest) SetBody(v string) *CreateQuotaPlanRequest {
	s.Body = &v
	return s
}

func (s *CreateQuotaPlanRequest) SetRegion(v string) *CreateQuotaPlanRequest {
	s.Region = &v
	return s
}

func (s *CreateQuotaPlanRequest) SetTenantId(v string) *CreateQuotaPlanRequest {
	s.TenantId = &v
	return s
}

type CreateQuotaPlanResponseBody struct {
	// The returned result.
	//
	// example:
	//
	// success
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0bc3b4b016674434996033675e71ee
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateQuotaPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateQuotaPlanResponseBody) GoString() string {
	return s.String()
}

func (s *CreateQuotaPlanResponseBody) SetData(v string) *CreateQuotaPlanResponseBody {
	s.Data = &v
	return s
}

func (s *CreateQuotaPlanResponseBody) SetRequestId(v string) *CreateQuotaPlanResponseBody {
	s.RequestId = &v
	return s
}

type CreateQuotaPlanResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateQuotaPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateQuotaPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateQuotaPlanResponse) GoString() string {
	return s.String()
}

func (s *CreateQuotaPlanResponse) SetHeaders(v map[string]*string) *CreateQuotaPlanResponse {
	s.Headers = v
	return s
}

func (s *CreateQuotaPlanResponse) SetStatusCode(v int32) *CreateQuotaPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateQuotaPlanResponse) SetBody(v *CreateQuotaPlanResponseBody) *CreateQuotaPlanResponse {
	s.Body = v
	return s
}

type CreateRoleRequest struct {
	// The request body parameters. For valid values, see [MaxCompute permissions](https://help.aliyun.com/document_detail/27935.html).
	//
	// example:
	//
	// {"name": "role_name","type": "resource/adminn","policy": "", // The document of the policy. This parameter is not required if an access-control list (ACL) is used. "acl": { // This parameter is not required if a policy is used. "table": [{"name": "", "actions":["","",]}, {"name": "", "actions":[]}],"resource":[{"name": "", "actions":["","",]}, {"name": "", "actions":[]}],"function":[{"name": "", "actions":["","",]}, {"name": "", "actions":[]}],"package":[{"name": "", "actions":["","",]}, {"name": "", "actions":[]}],"project":[{"name": "", "actions":[]}], // Only the current project is displayed in the console. "instance":[{"name": "", "actions":[]}] // The parameter name must be set to an asterisk (\\*) in the console. }}// An asterisk (\\*) can be specified for name.
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRoleRequest) GoString() string {
	return s.String()
}

func (s *CreateRoleRequest) SetBody(v string) *CreateRoleRequest {
	s.Body = &v
	return s
}

type CreateRoleResponseBody struct {
	// The returned data.
	//
	// example:
	//
	// success
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0b87b7e716665825896565060e87a4
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRoleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRoleResponseBody) SetData(v string) *CreateRoleResponseBody {
	s.Data = &v
	return s
}

func (s *CreateRoleResponseBody) SetRequestId(v string) *CreateRoleResponseBody {
	s.RequestId = &v
	return s
}

type CreateRoleResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRoleResponse) GoString() string {
	return s.String()
}

func (s *CreateRoleResponse) SetHeaders(v map[string]*string) *CreateRoleResponse {
	s.Headers = v
	return s
}

func (s *CreateRoleResponse) SetStatusCode(v int32) *CreateRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRoleResponse) SetBody(v *CreateRoleResponseBody) *CreateRoleResponse {
	s.Body = v
	return s
}

type DeleteComputeQuotaPlanResponseBody struct {
	// The returned data.
	//
	// example:
	//
	// success
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// The error code.
	//
	// example:
	//
	// OBJECT_NOT_EXIST
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// this quota plan is not exist.
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// The HTTP status code.
	//
	// - 1xx: informational response. The request is received and is being processed.
	//
	// - 2xx: success. The request is successfully received, understood, and accepted by the server.
	//
	// - 3xx: redirection. The request is redirected, and further actions are required to complete the request.
	//
	// - 4xx: client error. The request contains invalid request parameters or syntaxes, or specific request conditions cannot be met.
	//
	// - 5xx: server error. The server cannot meet requirements due to other reasons.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0b87b7a316654730544735643e9200
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteComputeQuotaPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteComputeQuotaPlanResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteComputeQuotaPlanResponseBody) SetData(v string) *DeleteComputeQuotaPlanResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteComputeQuotaPlanResponseBody) SetErrorCode(v string) *DeleteComputeQuotaPlanResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteComputeQuotaPlanResponseBody) SetErrorMsg(v string) *DeleteComputeQuotaPlanResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *DeleteComputeQuotaPlanResponseBody) SetHttpCode(v int32) *DeleteComputeQuotaPlanResponseBody {
	s.HttpCode = &v
	return s
}

func (s *DeleteComputeQuotaPlanResponseBody) SetRequestId(v string) *DeleteComputeQuotaPlanResponseBody {
	s.RequestId = &v
	return s
}

type DeleteComputeQuotaPlanResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteComputeQuotaPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteComputeQuotaPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteComputeQuotaPlanResponse) GoString() string {
	return s.String()
}

func (s *DeleteComputeQuotaPlanResponse) SetHeaders(v map[string]*string) *DeleteComputeQuotaPlanResponse {
	s.Headers = v
	return s
}

func (s *DeleteComputeQuotaPlanResponse) SetStatusCode(v int32) *DeleteComputeQuotaPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteComputeQuotaPlanResponse) SetBody(v *DeleteComputeQuotaPlanResponseBody) *DeleteComputeQuotaPlanResponse {
	s.Body = v
	return s
}

type DeleteMmsDataSourceResponseBody struct {
	// example:
	//
	// 2000015
	Data *int64 `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// EA1320AB-7766-5EC7-B0F6-8B20E2298567
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteMmsDataSourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMmsDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMmsDataSourceResponseBody) SetData(v int64) *DeleteMmsDataSourceResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteMmsDataSourceResponseBody) SetRequestId(v string) *DeleteMmsDataSourceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteMmsDataSourceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMmsDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMmsDataSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMmsDataSourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteMmsDataSourceResponse) SetHeaders(v map[string]*string) *DeleteMmsDataSourceResponse {
	s.Headers = v
	return s
}

func (s *DeleteMmsDataSourceResponse) SetStatusCode(v int32) *DeleteMmsDataSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMmsDataSourceResponse) SetBody(v *DeleteMmsDataSourceResponseBody) *DeleteMmsDataSourceResponse {
	s.Body = v
	return s
}

type DeleteMmsJobResponseBody struct {
	// example:
	//
	// 88
	Data *int64 `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// 31BE216F-AEF7-581E-B9C9-DECEB5424AC4
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteMmsJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMmsJobResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMmsJobResponseBody) SetData(v int64) *DeleteMmsJobResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteMmsJobResponseBody) SetRequestId(v string) *DeleteMmsJobResponseBody {
	s.RequestId = &v
	return s
}

type DeleteMmsJobResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMmsJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMmsJobResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMmsJobResponse) GoString() string {
	return s.String()
}

func (s *DeleteMmsJobResponse) SetHeaders(v map[string]*string) *DeleteMmsJobResponse {
	s.Headers = v
	return s
}

func (s *DeleteMmsJobResponse) SetStatusCode(v int32) *DeleteMmsJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMmsJobResponse) SetBody(v *DeleteMmsJobResponseBody) *DeleteMmsJobResponse {
	s.Body = v
	return s
}

type DeleteQuotaPlanRequest struct {
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The ID of the tenant.
	//
	// example:
	//
	// 416441016836866
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s DeleteQuotaPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteQuotaPlanRequest) GoString() string {
	return s.String()
}

func (s *DeleteQuotaPlanRequest) SetRegion(v string) *DeleteQuotaPlanRequest {
	s.Region = &v
	return s
}

func (s *DeleteQuotaPlanRequest) SetTenantId(v string) *DeleteQuotaPlanRequest {
	s.TenantId = &v
	return s
}

type DeleteQuotaPlanResponseBody struct {
	// The returned result.
	//
	// example:
	//
	// success
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0b57ff7616612271051086500ea3ce
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteQuotaPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteQuotaPlanResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteQuotaPlanResponseBody) SetData(v string) *DeleteQuotaPlanResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteQuotaPlanResponseBody) SetRequestId(v string) *DeleteQuotaPlanResponseBody {
	s.RequestId = &v
	return s
}

type DeleteQuotaPlanResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteQuotaPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteQuotaPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteQuotaPlanResponse) GoString() string {
	return s.String()
}

func (s *DeleteQuotaPlanResponse) SetHeaders(v map[string]*string) *DeleteQuotaPlanResponse {
	s.Headers = v
	return s
}

func (s *DeleteQuotaPlanResponse) SetStatusCode(v int32) *DeleteQuotaPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteQuotaPlanResponse) SetBody(v *DeleteQuotaPlanResponseBody) *DeleteQuotaPlanResponse {
	s.Body = v
	return s
}

type GetComputeEffectivePlanResponseBody struct {
	// The data returned.
	Data *GetComputeEffectivePlanResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// success
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// plan \\"***\\" does not exist
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// The HTTP status code.
	//
	// - 1xx: informational response. The request is received and is being processed.
	//
	// - 2xx: success. The request is successfully received, understood, and accepted by the server.
	//
	// - 3xx: redirection. The request is redirected, and further actions are required to complete the request.
	//
	// - 4xx: client error. The request contains invalid request parameters or syntaxes, or specific request conditions cannot be met.
	//
	// - 5xx: server error. The server cannot meet requirements due to other reasons.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 688003E1-D1B4-5468-957E-2FFB3AC8D79B
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetComputeEffectivePlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetComputeEffectivePlanResponseBody) GoString() string {
	return s.String()
}

func (s *GetComputeEffectivePlanResponseBody) SetData(v *GetComputeEffectivePlanResponseBodyData) *GetComputeEffectivePlanResponseBody {
	s.Data = v
	return s
}

func (s *GetComputeEffectivePlanResponseBody) SetErrorCode(v string) *GetComputeEffectivePlanResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetComputeEffectivePlanResponseBody) SetErrorMsg(v string) *GetComputeEffectivePlanResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetComputeEffectivePlanResponseBody) SetHttpCode(v int32) *GetComputeEffectivePlanResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetComputeEffectivePlanResponseBody) SetRequestId(v string) *GetComputeEffectivePlanResponseBody {
	s.RequestId = &v
	return s
}

type GetComputeEffectivePlanResponseBodyData struct {
	// The time when the quota plan was created.
	//
	// example:
	//
	// 1714356241163
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// Whether it is currently effective.
	//
	// > A Quota plan that has taken effect cannot be deleted, i.e., isEffective=true
	//
	// example:
	//
	// true/false
	IsEffective *bool `json:"isEffective,omitempty" xml:"isEffective,omitempty"`
	// The name of the quota plan.
	//
	// example:
	//
	// planA
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The details of the quota.
	Quota *GetComputeEffectivePlanResponseBodyDataQuota `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
}

func (s GetComputeEffectivePlanResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetComputeEffectivePlanResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetComputeEffectivePlanResponseBodyData) SetCreateTime(v string) *GetComputeEffectivePlanResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetComputeEffectivePlanResponseBodyData) SetIsEffective(v bool) *GetComputeEffectivePlanResponseBodyData {
	s.IsEffective = &v
	return s
}

func (s *GetComputeEffectivePlanResponseBodyData) SetName(v string) *GetComputeEffectivePlanResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetComputeEffectivePlanResponseBodyData) SetQuota(v *GetComputeEffectivePlanResponseBodyDataQuota) *GetComputeEffectivePlanResponseBodyData {
	s.Quota = v
	return s
}

type GetComputeEffectivePlanResponseBodyDataQuota struct {
	// The ID of the cluster.
	//
	// example:
	//
	// AT-120N
	Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// The time when the level-1 quota was created.
	//
	// example:
	//
	// 1719886322347
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The ID of the Alibaba Cloud account that is used to create the resource.
	//
	// example:
	//
	// 672863518
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// The ID of the level-1 quota.
	//
	// example:
	//
	// 2413
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the level-1 quota.
	//
	// example:
	//
	// dp_cn_shanghai_1699533470_p
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The nickname of the level-1 quota.
	//
	// example:
	//
	// os_MyQuota_p
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	// The description of the level-2 quota.
	//
	// example:
	//
	// {
	//
	//   "enablePriority": false,
	//
	//   "minCU": 25,
	//
	//   "adhocCU": 0,
	//
	//   "elasticReservedCU": 0,
	//
	//   "forceReservedMin": false,
	//
	//   "maxCU": 50,
	//
	//   "schedulerType": "Fifo"
	//
	// }
	Parameter map[string]interface{} `json:"parameter,omitempty" xml:"parameter,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The status of the resource.
	//
	// example:
	//
	// ON
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The list of subquotas.
	SubQuotaInfoList []*GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList `json:"subQuotaInfoList,omitempty" xml:"subQuotaInfoList,omitempty" type:"Repeated"`
	// The ID of the tenant.
	//
	// example:
	//
	// 478403690625249
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The type of quota.
	//
	// example:
	//
	// FUXI_ONLINE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The version number.
	//
	// example:
	//
	// 1964
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetComputeEffectivePlanResponseBodyDataQuota) String() string {
	return tea.Prettify(s)
}

func (s GetComputeEffectivePlanResponseBodyDataQuota) GoString() string {
	return s.String()
}

func (s *GetComputeEffectivePlanResponseBodyDataQuota) SetCluster(v string) *GetComputeEffectivePlanResponseBodyDataQuota {
	s.Cluster = &v
	return s
}

func (s *GetComputeEffectivePlanResponseBodyDataQuota) SetCreateTime(v int64) *GetComputeEffectivePlanResponseBodyDataQuota {
	s.CreateTime = &v
	return s
}

func (s *GetComputeEffectivePlanResponseBodyDataQuota) SetCreatorId(v string) *GetComputeEffectivePlanResponseBodyDataQuota {
	s.CreatorId = &v
	return s
}

func (s *GetComputeEffectivePlanResponseBodyDataQuota) SetId(v string) *GetComputeEffectivePlanResponseBodyDataQuota {
	s.Id = &v
	return s
}

func (s *GetComputeEffectivePlanResponseBodyDataQuota) SetName(v string) *GetComputeEffectivePlanResponseBodyDataQuota {
	s.Name = &v
	return s
}

func (s *GetComputeEffectivePlanResponseBodyDataQuota) SetNickName(v string) *GetComputeEffectivePlanResponseBodyDataQuota {
	s.NickName = &v
	return s
}

func (s *GetComputeEffectivePlanResponseBodyDataQuota) SetParameter(v map[string]interface{}) *GetComputeEffectivePlanResponseBodyDataQuota {
	s.Parameter = v
	return s
}

func (s *GetComputeEffectivePlanResponseBodyDataQuota) SetRegionId(v string) *GetComputeEffectivePlanResponseBodyDataQuota {
	s.RegionId = &v
	return s
}

func (s *GetComputeEffectivePlanResponseBodyDataQuota) SetStatus(v string) *GetComputeEffectivePlanResponseBodyDataQuota {
	s.Status = &v
	return s
}

func (s *GetComputeEffectivePlanResponseBodyDataQuota) SetSubQuotaInfoList(v []*GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList) *GetComputeEffectivePlanResponseBodyDataQuota {
	s.SubQuotaInfoList = v
	return s
}

func (s *GetComputeEffectivePlanResponseBodyDataQuota) SetTenantId(v string) *GetComputeEffectivePlanResponseBodyDataQuota {
	s.TenantId = &v
	return s
}

func (s *GetComputeEffectivePlanResponseBodyDataQuota) SetType(v string) *GetComputeEffectivePlanResponseBodyDataQuota {
	s.Type = &v
	return s
}

func (s *GetComputeEffectivePlanResponseBodyDataQuota) SetVersion(v string) *GetComputeEffectivePlanResponseBodyDataQuota {
	s.Version = &v
	return s
}

type GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList struct {
	// The ID of the cluster.
	//
	// example:
	//
	// AT-120N
	Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// The time when the resource was created.
	//
	// example:
	//
	// 1718155201628
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The ID of the Alibaba Cloud account that is used to create the resource.
	//
	// example:
	//
	// 672863518
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// The ID of the level-2 quota.
	//
	// example:
	//
	// 10940
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the level-2 quota.
	//
	// example:
	//
	// dp_cn_shanghai_1696659792_p
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The nickname of the level-2 quota.
	//
	// example:
	//
	// os_MyQuota
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	// The description of the level-2 quota.
	//
	// example:
	//
	// {
	//
	//   "enablePriority": false,
	//
	//   "minCU": 25,
	//
	//   "adhocCU": 0,
	//
	//   "elasticReservedCU": 0,
	//
	//   "forceReservedMin": false,
	//
	//   "maxCU": 50,
	//
	//   "schedulerType": "Fifo"
	//
	// }
	Parameter map[string]interface{} `json:"parameter,omitempty" xml:"parameter,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// Resource status.
	//
	// example:
	//
	// ON
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The ID of the tenant.
	//
	// example:
	//
	// 478403690625249
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The type of quota.
	//
	// example:
	//
	// FUXI_ONLINE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The version number.
	//
	// example:
	//
	// 1386
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList) String() string {
	return tea.Prettify(s)
}

func (s GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList) GoString() string {
	return s.String()
}

func (s *GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList) SetCluster(v string) *GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList {
	s.Cluster = &v
	return s
}

func (s *GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList) SetCreateTime(v int64) *GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList {
	s.CreateTime = &v
	return s
}

func (s *GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList) SetCreatorId(v string) *GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList {
	s.CreatorId = &v
	return s
}

func (s *GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList) SetId(v string) *GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList {
	s.Id = &v
	return s
}

func (s *GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList) SetName(v string) *GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList {
	s.Name = &v
	return s
}

func (s *GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList) SetNickName(v string) *GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList {
	s.NickName = &v
	return s
}

func (s *GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList) SetParameter(v map[string]interface{}) *GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList {
	s.Parameter = v
	return s
}

func (s *GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList) SetRegionId(v string) *GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList {
	s.RegionId = &v
	return s
}

func (s *GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList) SetStatus(v string) *GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList {
	s.Status = &v
	return s
}

func (s *GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList) SetTenantId(v string) *GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList {
	s.TenantId = &v
	return s
}

func (s *GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList) SetType(v string) *GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList {
	s.Type = &v
	return s
}

func (s *GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList) SetVersion(v string) *GetComputeEffectivePlanResponseBodyDataQuotaSubQuotaInfoList {
	s.Version = &v
	return s
}

type GetComputeEffectivePlanResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetComputeEffectivePlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetComputeEffectivePlanResponse) String() string {
	return tea.Prettify(s)
}

func (s GetComputeEffectivePlanResponse) GoString() string {
	return s.String()
}

func (s *GetComputeEffectivePlanResponse) SetHeaders(v map[string]*string) *GetComputeEffectivePlanResponse {
	s.Headers = v
	return s
}

func (s *GetComputeEffectivePlanResponse) SetStatusCode(v int32) *GetComputeEffectivePlanResponse {
	s.StatusCode = &v
	return s
}

func (s *GetComputeEffectivePlanResponse) SetBody(v *GetComputeEffectivePlanResponseBody) *GetComputeEffectivePlanResponse {
	s.Body = v
	return s
}

type GetComputeQuotaPlanResponseBody struct {
	// The data returned.
	Data *GetComputeQuotaPlanResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// QUOTA_PLAN_NOT_FOUND
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// plan \\"***\\" does not exist
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// The HTTP status code.
	//
	// - 1xx: informational response. The request is received and is being processed.
	//
	// - 2xx: success. The request is successfully received, understood, and accepted by the server.
	//
	// - 3xx: redirection. The request is redirected, and further actions are required to complete the request.
	//
	// - 4xx: client error. The request contains invalid request parameters or syntaxes, or specific request conditions cannot be met.
	//
	// - 5xx: server error. The server cannot meet requirements due to other reasons.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EA1320AB-7766-5EC7-B0F6-8B20E2298567
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetComputeQuotaPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetComputeQuotaPlanResponseBody) GoString() string {
	return s.String()
}

func (s *GetComputeQuotaPlanResponseBody) SetData(v *GetComputeQuotaPlanResponseBodyData) *GetComputeQuotaPlanResponseBody {
	s.Data = v
	return s
}

func (s *GetComputeQuotaPlanResponseBody) SetErrorCode(v string) *GetComputeQuotaPlanResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBody) SetErrorMsg(v string) *GetComputeQuotaPlanResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBody) SetHttpCode(v int32) *GetComputeQuotaPlanResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBody) SetRequestId(v string) *GetComputeQuotaPlanResponseBody {
	s.RequestId = &v
	return s
}

type GetComputeQuotaPlanResponseBodyData struct {
	// The time when the quota plan was created.
	//
	// example:
	//
	// 1730946421757
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// Whether it is currently effective.
	//
	// >
	//
	// > - A Quota plan that has taken effect cannot be deleted, i.e., isEffective=true
	//
	// example:
	//
	// true/false
	IsEffective *bool `json:"isEffective,omitempty" xml:"isEffective,omitempty"`
	// The name of the quota plan.
	//
	// example:
	//
	// planA
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The details of the quota.
	Quota *GetComputeQuotaPlanResponseBodyDataQuota `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
}

func (s GetComputeQuotaPlanResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetComputeQuotaPlanResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetComputeQuotaPlanResponseBodyData) SetCreateTime(v string) *GetComputeQuotaPlanResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyData) SetIsEffective(v bool) *GetComputeQuotaPlanResponseBodyData {
	s.IsEffective = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyData) SetName(v string) *GetComputeQuotaPlanResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyData) SetQuota(v *GetComputeQuotaPlanResponseBodyDataQuota) *GetComputeQuotaPlanResponseBodyData {
	s.Quota = v
	return s
}

type GetComputeQuotaPlanResponseBodyDataQuota struct {
	// Cluster ID.
	//
	// example:
	//
	// AT-120N
	Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// Creation time.
	//
	// example:
	//
	// 1719886322347
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// Creator\\"s cloud account UID.
	//
	// example:
	//
	// 672863518
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// The ID of the level-1 quota.
	//
	// example:
	//
	// 2413
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the level-1 quota.
	//
	// example:
	//
	// quota_a
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The nickname of the level-1 quota.
	//
	// example:
	//
	// quota_nickname
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	// CU value parameters for the level-1 quota.
	Parameter *GetComputeQuotaPlanResponseBodyDataQuotaParameter `json:"parameter,omitempty" xml:"parameter,omitempty" type:"Struct"`
	// Region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// Resource status.
	//
	// example:
	//
	// ON
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The list of level-2 quotas.
	SubQuotaInfoList []*GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList `json:"subQuotaInfoList,omitempty" xml:"subQuotaInfoList,omitempty" type:"Repeated"`
	// Tenant ID.
	//
	// example:
	//
	// 478403690625249
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// Corresponds to the `resourceSystemType` field of the control cluster.
	//
	// example:
	//
	// FUXI_ONLINE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// Version number.
	//
	// example:
	//
	// 1964
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetComputeQuotaPlanResponseBodyDataQuota) String() string {
	return tea.Prettify(s)
}

func (s GetComputeQuotaPlanResponseBodyDataQuota) GoString() string {
	return s.String()
}

func (s *GetComputeQuotaPlanResponseBodyDataQuota) SetCluster(v string) *GetComputeQuotaPlanResponseBodyDataQuota {
	s.Cluster = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyDataQuota) SetCreateTime(v int64) *GetComputeQuotaPlanResponseBodyDataQuota {
	s.CreateTime = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyDataQuota) SetCreatorId(v string) *GetComputeQuotaPlanResponseBodyDataQuota {
	s.CreatorId = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyDataQuota) SetId(v string) *GetComputeQuotaPlanResponseBodyDataQuota {
	s.Id = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyDataQuota) SetName(v string) *GetComputeQuotaPlanResponseBodyDataQuota {
	s.Name = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyDataQuota) SetNickName(v string) *GetComputeQuotaPlanResponseBodyDataQuota {
	s.NickName = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyDataQuota) SetParameter(v *GetComputeQuotaPlanResponseBodyDataQuotaParameter) *GetComputeQuotaPlanResponseBodyDataQuota {
	s.Parameter = v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyDataQuota) SetRegionId(v string) *GetComputeQuotaPlanResponseBodyDataQuota {
	s.RegionId = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyDataQuota) SetStatus(v string) *GetComputeQuotaPlanResponseBodyDataQuota {
	s.Status = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyDataQuota) SetSubQuotaInfoList(v []*GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) *GetComputeQuotaPlanResponseBodyDataQuota {
	s.SubQuotaInfoList = v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyDataQuota) SetTenantId(v string) *GetComputeQuotaPlanResponseBodyDataQuota {
	s.TenantId = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyDataQuota) SetType(v string) *GetComputeQuotaPlanResponseBodyDataQuota {
	s.Type = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyDataQuota) SetVersion(v string) *GetComputeQuotaPlanResponseBodyDataQuota {
	s.Version = &v
	return s
}

type GetComputeQuotaPlanResponseBodyDataQuotaParameter struct {
	// The value of elastic Reserved CUs.
	//
	// example:
	//
	// 50
	ElasticReservedCU *int64 `json:"elasticReservedCU,omitempty" xml:"elasticReservedCU,omitempty"`
	// The value of maxCU in Reserved CUs.
	//
	// example:
	//
	// 50
	MaxCU *int64 `json:"maxCU,omitempty" xml:"maxCU,omitempty"`
	// The value of minCU in Reserved CUs.
	//
	// example:
	//
	// 50
	MinCU *int64 `json:"minCU,omitempty" xml:"minCU,omitempty"`
}

func (s GetComputeQuotaPlanResponseBodyDataQuotaParameter) String() string {
	return tea.Prettify(s)
}

func (s GetComputeQuotaPlanResponseBodyDataQuotaParameter) GoString() string {
	return s.String()
}

func (s *GetComputeQuotaPlanResponseBodyDataQuotaParameter) SetElasticReservedCU(v int64) *GetComputeQuotaPlanResponseBodyDataQuotaParameter {
	s.ElasticReservedCU = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyDataQuotaParameter) SetMaxCU(v int64) *GetComputeQuotaPlanResponseBodyDataQuotaParameter {
	s.MaxCU = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyDataQuotaParameter) SetMinCU(v int64) *GetComputeQuotaPlanResponseBodyDataQuotaParameter {
	s.MinCU = &v
	return s
}

type GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList struct {
	// Cluster ID.
	//
	// example:
	//
	// AT-120N
	Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// Creation time.
	//
	// example:
	//
	// 1718155201628
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// Creator cloud account UID.
	//
	// example:
	//
	// 672863518
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// The ID of the level-2 quota.
	//
	// example:
	//
	// 10940
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the level-2 quota.
	//
	// example:
	//
	// dp_cn_shanghai_1696659792_p
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The nickname of the level-2 quota.
	//
	// example:
	//
	// subquotaA
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	// The parameters of the level-2 quota.
	Parameter *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoListParameter `json:"parameter,omitempty" xml:"parameter,omitempty" type:"Struct"`
	// Region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// Resource status.
	//
	// example:
	//
	// ON
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// Tenant ID.
	//
	// example:
	//
	// 478403690625249
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The type of quota.
	//
	// example:
	//
	// FUXI_ONLINE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// Version number.
	//
	// example:
	//
	// 1386
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) String() string {
	return tea.Prettify(s)
}

func (s GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) GoString() string {
	return s.String()
}

func (s *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetCluster(v string) *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.Cluster = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetCreateTime(v int64) *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.CreateTime = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetCreatorId(v string) *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.CreatorId = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetId(v string) *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.Id = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetName(v string) *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.Name = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetNickName(v string) *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.NickName = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetParameter(v *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoListParameter) *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.Parameter = v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetRegionId(v string) *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.RegionId = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetStatus(v string) *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.Status = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetTenantId(v string) *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.TenantId = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetType(v string) *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.Type = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetVersion(v string) *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.Version = &v
	return s
}

type GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoListParameter struct {
	// The value of elastic Reserved CUs.
	//
	// example:
	//
	// 50
	ElasticReservedCU *int64 `json:"elasticReservedCU,omitempty" xml:"elasticReservedCU,omitempty"`
	// whether to enable the priority feature.
	//
	// example:
	//
	// true/false
	EnablePriority *bool `json:"enablePriority,omitempty" xml:"enablePriority,omitempty"`
	// Whether it is exclusive.
	//
	// example:
	//
	// true/false
	ForceReservedMin *bool `json:"forceReservedMin,omitempty" xml:"forceReservedMin,omitempty"`
	// The value of maxCU in Reserved CUs.
	//
	// example:
	//
	// 50
	MaxCU *int64 `json:"maxCU,omitempty" xml:"maxCU,omitempty"`
	// The value of minCU in Reserved CUs.
	//
	// example:
	//
	// 50
	MinCU *int64 `json:"minCU,omitempty" xml:"minCU,omitempty"`
	// Scheduling policy.
	//
	// example:
	//
	// Fifo/Fair
	SchedulerType *string `json:"schedulerType,omitempty" xml:"schedulerType,omitempty"`
	// The upper limit for CUs that can be concurrently used by a job scheduled to the quota.
	//
	// example:
	//
	// 50
	SingleJobCULimit *int64 `json:"singleJobCULimit,omitempty" xml:"singleJobCULimit,omitempty"`
}

func (s GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoListParameter) String() string {
	return tea.Prettify(s)
}

func (s GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoListParameter) GoString() string {
	return s.String()
}

func (s *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoListParameter) SetElasticReservedCU(v int64) *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoListParameter {
	s.ElasticReservedCU = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoListParameter) SetEnablePriority(v bool) *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoListParameter {
	s.EnablePriority = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoListParameter) SetForceReservedMin(v bool) *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoListParameter {
	s.ForceReservedMin = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoListParameter) SetMaxCU(v int64) *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoListParameter {
	s.MaxCU = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoListParameter) SetMinCU(v int64) *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoListParameter {
	s.MinCU = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoListParameter) SetSchedulerType(v string) *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoListParameter {
	s.SchedulerType = &v
	return s
}

func (s *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoListParameter) SetSingleJobCULimit(v int64) *GetComputeQuotaPlanResponseBodyDataQuotaSubQuotaInfoListParameter {
	s.SingleJobCULimit = &v
	return s
}

type GetComputeQuotaPlanResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetComputeQuotaPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetComputeQuotaPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s GetComputeQuotaPlanResponse) GoString() string {
	return s.String()
}

func (s *GetComputeQuotaPlanResponse) SetHeaders(v map[string]*string) *GetComputeQuotaPlanResponse {
	s.Headers = v
	return s
}

func (s *GetComputeQuotaPlanResponse) SetStatusCode(v int32) *GetComputeQuotaPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *GetComputeQuotaPlanResponse) SetBody(v *GetComputeQuotaPlanResponseBody) *GetComputeQuotaPlanResponse {
	s.Body = v
	return s
}

type GetComputeQuotaScheduleRequest struct {
	// Display time zone.
	//
	// example:
	//
	// UTC+8
	DisplayTimezone *string `json:"displayTimezone,omitempty" xml:"displayTimezone,omitempty"`
}

func (s GetComputeQuotaScheduleRequest) String() string {
	return tea.Prettify(s)
}

func (s GetComputeQuotaScheduleRequest) GoString() string {
	return s.String()
}

func (s *GetComputeQuotaScheduleRequest) SetDisplayTimezone(v string) *GetComputeQuotaScheduleRequest {
	s.DisplayTimezone = &v
	return s
}

type GetComputeQuotaScheduleResponseBody struct {
	// The data returned.
	Data []*GetComputeQuotaScheduleResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The error code.
	//
	// example:
	//
	// QUOTA_UNKNOWN_NICKNAME
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Cannot found quota **
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// The HTTP status code.
	//
	// - 1xx: informational response. The request is received and is being processed.
	//
	// - 2xx: success. The request is successfully received, understood, and accepted by the server.
	//
	// - 3xx: redirection. The request is redirected, and further actions are required to complete the request.
	//
	// - 4xx: client error. The request contains invalid request parameters or syntaxes, or specific request conditions cannot be met.
	//
	// - 5xx: server error. The server cannot meet requirements due to other reasons.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B42CA730-8187-50F1-9FE0-6733297036DB
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetComputeQuotaScheduleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetComputeQuotaScheduleResponseBody) GoString() string {
	return s.String()
}

func (s *GetComputeQuotaScheduleResponseBody) SetData(v []*GetComputeQuotaScheduleResponseBodyData) *GetComputeQuotaScheduleResponseBody {
	s.Data = v
	return s
}

func (s *GetComputeQuotaScheduleResponseBody) SetErrorCode(v string) *GetComputeQuotaScheduleResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetComputeQuotaScheduleResponseBody) SetErrorMsg(v string) *GetComputeQuotaScheduleResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetComputeQuotaScheduleResponseBody) SetHttpCode(v int32) *GetComputeQuotaScheduleResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetComputeQuotaScheduleResponseBody) SetRequestId(v string) *GetComputeQuotaScheduleResponseBody {
	s.RequestId = &v
	return s
}

type GetComputeQuotaScheduleResponseBodyData struct {
	// The value of effective condition.
	Condition *GetComputeQuotaScheduleResponseBodyDataCondition `json:"condition,omitempty" xml:"condition,omitempty" type:"Struct"`
	// The ID of the quota plan.
	//
	// example:
	//
	// 89b54db44d384f26964951ea457f64a5
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the quota plan.
	//
	// example:
	//
	// planA
	Plan *string `json:"plan,omitempty" xml:"plan,omitempty"`
	// The time zone property.
	//
	// example:
	//
	// UTC+8
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
	// The type of the quota plan.
	//
	// example:
	//
	// daily
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetComputeQuotaScheduleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetComputeQuotaScheduleResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetComputeQuotaScheduleResponseBodyData) SetCondition(v *GetComputeQuotaScheduleResponseBodyDataCondition) *GetComputeQuotaScheduleResponseBodyData {
	s.Condition = v
	return s
}

func (s *GetComputeQuotaScheduleResponseBodyData) SetId(v string) *GetComputeQuotaScheduleResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetComputeQuotaScheduleResponseBodyData) SetPlan(v string) *GetComputeQuotaScheduleResponseBodyData {
	s.Plan = &v
	return s
}

func (s *GetComputeQuotaScheduleResponseBodyData) SetTimezone(v string) *GetComputeQuotaScheduleResponseBodyData {
	s.Timezone = &v
	return s
}

func (s *GetComputeQuotaScheduleResponseBodyData) SetType(v string) *GetComputeQuotaScheduleResponseBodyData {
	s.Type = &v
	return s
}

type GetComputeQuotaScheduleResponseBodyDataCondition struct {
	// The start time when the quota plan takes effect.
	//
	// example:
	//
	// 09:00
	At *string `json:"at,omitempty" xml:"at,omitempty"`
}

func (s GetComputeQuotaScheduleResponseBodyDataCondition) String() string {
	return tea.Prettify(s)
}

func (s GetComputeQuotaScheduleResponseBodyDataCondition) GoString() string {
	return s.String()
}

func (s *GetComputeQuotaScheduleResponseBodyDataCondition) SetAt(v string) *GetComputeQuotaScheduleResponseBodyDataCondition {
	s.At = &v
	return s
}

type GetComputeQuotaScheduleResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetComputeQuotaScheduleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetComputeQuotaScheduleResponse) String() string {
	return tea.Prettify(s)
}

func (s GetComputeQuotaScheduleResponse) GoString() string {
	return s.String()
}

func (s *GetComputeQuotaScheduleResponse) SetHeaders(v map[string]*string) *GetComputeQuotaScheduleResponse {
	s.Headers = v
	return s
}

func (s *GetComputeQuotaScheduleResponse) SetStatusCode(v int32) *GetComputeQuotaScheduleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetComputeQuotaScheduleResponse) SetBody(v *GetComputeQuotaScheduleResponseBody) *GetComputeQuotaScheduleResponse {
	s.Body = v
	return s
}

type GetJobResourceUsageRequest struct {
	// The date that is accurate to the day part for the query. The date must be in the yyyy-MM-dd format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-05-15
	Date *string `json:"date,omitempty" xml:"date,omitempty"`
	// The list of job executors.
	JobOwnerList []*string `json:"jobOwnerList,omitempty" xml:"jobOwnerList,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The list of nicknames of quotas that are used by jobs.
	QuotaNicknameList []*string `json:"quotaNicknameList,omitempty" xml:"quotaNicknameList,omitempty" type:"Repeated"`
}

func (s GetJobResourceUsageRequest) String() string {
	return tea.Prettify(s)
}

func (s GetJobResourceUsageRequest) GoString() string {
	return s.String()
}

func (s *GetJobResourceUsageRequest) SetDate(v string) *GetJobResourceUsageRequest {
	s.Date = &v
	return s
}

func (s *GetJobResourceUsageRequest) SetJobOwnerList(v []*string) *GetJobResourceUsageRequest {
	s.JobOwnerList = v
	return s
}

func (s *GetJobResourceUsageRequest) SetPageNumber(v int64) *GetJobResourceUsageRequest {
	s.PageNumber = &v
	return s
}

func (s *GetJobResourceUsageRequest) SetPageSize(v int64) *GetJobResourceUsageRequest {
	s.PageSize = &v
	return s
}

func (s *GetJobResourceUsageRequest) SetQuotaNicknameList(v []*string) *GetJobResourceUsageRequest {
	s.QuotaNicknameList = v
	return s
}

type GetJobResourceUsageShrinkRequest struct {
	// The date that is accurate to the day part for the query. The date must be in the yyyy-MM-dd format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-05-15
	Date *string `json:"date,omitempty" xml:"date,omitempty"`
	// The list of job executors.
	JobOwnerListShrink *string `json:"jobOwnerList,omitempty" xml:"jobOwnerList,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The list of nicknames of quotas that are used by jobs.
	QuotaNicknameListShrink *string `json:"quotaNicknameList,omitempty" xml:"quotaNicknameList,omitempty"`
}

func (s GetJobResourceUsageShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetJobResourceUsageShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetJobResourceUsageShrinkRequest) SetDate(v string) *GetJobResourceUsageShrinkRequest {
	s.Date = &v
	return s
}

func (s *GetJobResourceUsageShrinkRequest) SetJobOwnerListShrink(v string) *GetJobResourceUsageShrinkRequest {
	s.JobOwnerListShrink = &v
	return s
}

func (s *GetJobResourceUsageShrinkRequest) SetPageNumber(v int64) *GetJobResourceUsageShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *GetJobResourceUsageShrinkRequest) SetPageSize(v int64) *GetJobResourceUsageShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GetJobResourceUsageShrinkRequest) SetQuotaNicknameListShrink(v string) *GetJobResourceUsageShrinkRequest {
	s.QuotaNicknameListShrink = &v
	return s
}

type GetJobResourceUsageResponseBody struct {
	// The data returned.
	Data *GetJobResourceUsageResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// success
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// 0A3B1E82006A23A918C70905BF08AEC7
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// Indicates whether the request was successful. If this parameter was not empty and the value of this parameter was not 200, the request failed.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0b57ff7616612271051086500ea3ce
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetJobResourceUsageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetJobResourceUsageResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobResourceUsageResponseBody) SetData(v *GetJobResourceUsageResponseBodyData) *GetJobResourceUsageResponseBody {
	s.Data = v
	return s
}

func (s *GetJobResourceUsageResponseBody) SetErrorCode(v string) *GetJobResourceUsageResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetJobResourceUsageResponseBody) SetErrorMsg(v string) *GetJobResourceUsageResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetJobResourceUsageResponseBody) SetHttpCode(v int32) *GetJobResourceUsageResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetJobResourceUsageResponseBody) SetRequestId(v string) *GetJobResourceUsageResponseBody {
	s.RequestId = &v
	return s
}

type GetJobResourceUsageResponseBodyData struct {
	// The data list returned.
	JobResourceUsageList []*GetJobResourceUsageResponseBodyDataJobResourceUsageList `json:"jobResourceUsageList,omitempty" xml:"jobResourceUsageList,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 2
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 64
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s GetJobResourceUsageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetJobResourceUsageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetJobResourceUsageResponseBodyData) SetJobResourceUsageList(v []*GetJobResourceUsageResponseBodyDataJobResourceUsageList) *GetJobResourceUsageResponseBodyData {
	s.JobResourceUsageList = v
	return s
}

func (s *GetJobResourceUsageResponseBodyData) SetPageNumber(v int64) *GetJobResourceUsageResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *GetJobResourceUsageResponseBodyData) SetPageSize(v int64) *GetJobResourceUsageResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetJobResourceUsageResponseBodyData) SetTotalCount(v int64) *GetJobResourceUsageResponseBodyData {
	s.TotalCount = &v
	return s
}

type GetJobResourceUsageResponseBodyDataJobResourceUsageList struct {
	// The total number of used compute units (CUs).
	//
	// example:
	//
	// 1185100
	CuUsage *int64 `json:"cuUsage,omitempty" xml:"cuUsage,omitempty"`
	// The start date of the query in the format of yyyy-MM-dd.
	//
	// example:
	//
	// 2023-05-09
	Date *string `json:"date,omitempty" xml:"date,omitempty"`
	// The job executor.
	//
	// example:
	//
	// ALIYUN$xxx@test.aliyunid.com
	JobOwner *string `json:"jobOwner,omitempty" xml:"jobOwner,omitempty"`
	// The total memory usage.
	//
	// example:
	//
	// 15169536
	MemoryUsage *int64 `json:"memoryUsage,omitempty" xml:"memoryUsage,omitempty"`
	// The quota nickname.
	//
	// example:
	//
	// my_quota
	QuotaNickname *string `json:"quotaNickname,omitempty" xml:"quotaNickname,omitempty"`
}

func (s GetJobResourceUsageResponseBodyDataJobResourceUsageList) String() string {
	return tea.Prettify(s)
}

func (s GetJobResourceUsageResponseBodyDataJobResourceUsageList) GoString() string {
	return s.String()
}

func (s *GetJobResourceUsageResponseBodyDataJobResourceUsageList) SetCuUsage(v int64) *GetJobResourceUsageResponseBodyDataJobResourceUsageList {
	s.CuUsage = &v
	return s
}

func (s *GetJobResourceUsageResponseBodyDataJobResourceUsageList) SetDate(v string) *GetJobResourceUsageResponseBodyDataJobResourceUsageList {
	s.Date = &v
	return s
}

func (s *GetJobResourceUsageResponseBodyDataJobResourceUsageList) SetJobOwner(v string) *GetJobResourceUsageResponseBodyDataJobResourceUsageList {
	s.JobOwner = &v
	return s
}

func (s *GetJobResourceUsageResponseBodyDataJobResourceUsageList) SetMemoryUsage(v int64) *GetJobResourceUsageResponseBodyDataJobResourceUsageList {
	s.MemoryUsage = &v
	return s
}

func (s *GetJobResourceUsageResponseBodyDataJobResourceUsageList) SetQuotaNickname(v string) *GetJobResourceUsageResponseBodyDataJobResourceUsageList {
	s.QuotaNickname = &v
	return s
}

type GetJobResourceUsageResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetJobResourceUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetJobResourceUsageResponse) String() string {
	return tea.Prettify(s)
}

func (s GetJobResourceUsageResponse) GoString() string {
	return s.String()
}

func (s *GetJobResourceUsageResponse) SetHeaders(v map[string]*string) *GetJobResourceUsageResponse {
	s.Headers = v
	return s
}

func (s *GetJobResourceUsageResponse) SetStatusCode(v int32) *GetJobResourceUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetJobResourceUsageResponse) SetBody(v *GetJobResourceUsageResponseBody) *GetJobResourceUsageResponse {
	s.Body = v
	return s
}

type GetMmsAsyncTaskResponseBody struct {
	Data *GetMmsAsyncTaskResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 688003E1-D1B4-5468-957E-2FFB3AC8D79B
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetMmsAsyncTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMmsAsyncTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetMmsAsyncTaskResponseBody) SetData(v *GetMmsAsyncTaskResponseBodyData) *GetMmsAsyncTaskResponseBody {
	s.Data = v
	return s
}

func (s *GetMmsAsyncTaskResponseBody) SetRequestId(v string) *GetMmsAsyncTaskResponseBody {
	s.RequestId = &v
	return s
}

type GetMmsAsyncTaskResponseBodyData struct {
	// example:
	//
	// 2024-12-17 15:44:00
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 2024-12-17 17:44:17
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// null
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// example:
	//
	// 2523
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 232
	ObjectId *int64 `json:"objectId,omitempty" xml:"objectId,omitempty"`
	// example:
	//
	// 0
	Progress *int32 `json:"progress,omitempty" xml:"progress,omitempty"`
	// example:
	//
	// null
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
	// example:
	//
	// false
	Running *bool `json:"running,omitempty" xml:"running,omitempty"`
	// example:
	//
	// 2000017
	SourceId *int64 `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// example:
	//
	// 2024-12-17 15:44:17
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// DONE
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// TASK_CREATE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetMmsAsyncTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetMmsAsyncTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMmsAsyncTaskResponseBodyData) SetCreateTime(v string) *GetMmsAsyncTaskResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetMmsAsyncTaskResponseBodyData) SetEndTime(v string) *GetMmsAsyncTaskResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *GetMmsAsyncTaskResponseBodyData) SetErrorMsg(v string) *GetMmsAsyncTaskResponseBodyData {
	s.ErrorMsg = &v
	return s
}

func (s *GetMmsAsyncTaskResponseBodyData) SetId(v int64) *GetMmsAsyncTaskResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetMmsAsyncTaskResponseBodyData) SetObjectId(v int64) *GetMmsAsyncTaskResponseBodyData {
	s.ObjectId = &v
	return s
}

func (s *GetMmsAsyncTaskResponseBodyData) SetProgress(v int32) *GetMmsAsyncTaskResponseBodyData {
	s.Progress = &v
	return s
}

func (s *GetMmsAsyncTaskResponseBodyData) SetResult(v string) *GetMmsAsyncTaskResponseBodyData {
	s.Result = &v
	return s
}

func (s *GetMmsAsyncTaskResponseBodyData) SetRunning(v bool) *GetMmsAsyncTaskResponseBodyData {
	s.Running = &v
	return s
}

func (s *GetMmsAsyncTaskResponseBodyData) SetSourceId(v int64) *GetMmsAsyncTaskResponseBodyData {
	s.SourceId = &v
	return s
}

func (s *GetMmsAsyncTaskResponseBodyData) SetStartTime(v string) *GetMmsAsyncTaskResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *GetMmsAsyncTaskResponseBodyData) SetStatus(v string) *GetMmsAsyncTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetMmsAsyncTaskResponseBodyData) SetType(v string) *GetMmsAsyncTaskResponseBodyData {
	s.Type = &v
	return s
}

type GetMmsAsyncTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMmsAsyncTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMmsAsyncTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMmsAsyncTaskResponse) GoString() string {
	return s.String()
}

func (s *GetMmsAsyncTaskResponse) SetHeaders(v map[string]*string) *GetMmsAsyncTaskResponse {
	s.Headers = v
	return s
}

func (s *GetMmsAsyncTaskResponse) SetStatusCode(v int32) *GetMmsAsyncTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMmsAsyncTaskResponse) SetBody(v *GetMmsAsyncTaskResponseBody) *GetMmsAsyncTaskResponse {
	s.Body = v
	return s
}

type GetMmsDataSourceRequest struct {
	// example:
	//
	// en_US
	Lang *string `json:"lang,omitempty" xml:"lang,omitempty"`
	// example:
	//
	// true
	WithConfig *bool `json:"withConfig,omitempty" xml:"withConfig,omitempty"`
}

func (s GetMmsDataSourceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMmsDataSourceRequest) GoString() string {
	return s.String()
}

func (s *GetMmsDataSourceRequest) SetLang(v string) *GetMmsDataSourceRequest {
	s.Lang = &v
	return s
}

func (s *GetMmsDataSourceRequest) SetWithConfig(v bool) *GetMmsDataSourceRequest {
	s.WithConfig = &v
	return s
}

type GetMmsDataSourceResponseBody struct {
	Data *GetMmsDataSourceResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 98EC8C47-3D6D-560C-808B-84E494220A32
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetMmsDataSourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMmsDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetMmsDataSourceResponseBody) SetData(v *GetMmsDataSourceResponseBodyData) *GetMmsDataSourceResponseBody {
	s.Data = v
	return s
}

func (s *GetMmsDataSourceResponseBody) SetRequestId(v string) *GetMmsDataSourceResponseBody {
	s.RequestId = &v
	return s
}

type GetMmsDataSourceResponseBodyData struct {
	// example:
	//
	// true
	AgentIsOnline *bool                                     `json:"agentIsOnline,omitempty" xml:"agentIsOnline,omitempty"`
	Config        []*GetMmsDataSourceResponseBodyDataConfig `json:"config,omitempty" xml:"config,omitempty" type:"Repeated"`
	// example:
	//
	// 2024-12-17 09:29:58
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 3
	DbNum *int32 `json:"dbNum,omitempty" xml:"dbNum,omitempty"`
	// example:
	//
	// unexpected exception
	ErrMsg *string `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	// example:
	//
	// 2000015
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 2024-12-17 15:44:17
	LastUpdateTime *string `json:"lastUpdateTime,omitempty" xml:"lastUpdateTime,omitempty"`
	// example:
	//
	// demo
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// vpc-2zebqp6uojhdla46677tl:cn-shanghai
	Networklink *string `json:"networklink,omitempty" xml:"networklink,omitempty"`
	// example:
	//
	// 10000000
	PartitionNum *int32 `json:"partitionNum,omitempty" xml:"partitionNum,omitempty"`
	// example:
	//
	// 23322
	PartitionsDoingNum *int32 `json:"partitionsDoingNum,omitempty" xml:"partitionsDoingNum,omitempty"`
	// example:
	//
	// 11113
	PartitionsDoneNum *int32 `json:"partitionsDoneNum,omitempty" xml:"partitionsDoneNum,omitempty"`
	// example:
	//
	// 32
	PartitionsFailedNum *int32 `json:"partitionsFailedNum,omitempty" xml:"partitionsFailedNum,omitempty"`
	// example:
	//
	// cn-shanghai
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// example:
	//
	// 1000253
	ScanId *int64 `json:"scanId,omitempty" xml:"scanId,omitempty"`
	// example:
	//
	// STARTED
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 1000
	TableNum *int32 `json:"tableNum,omitempty" xml:"tableNum,omitempty"`
	// example:
	//
	// 19
	TablesDoingNum *int32 `json:"tablesDoingNum,omitempty" xml:"tablesDoingNum,omitempty"`
	// example:
	//
	// 16
	TablesDoneNum *int32 `json:"tablesDoneNum,omitempty" xml:"tablesDoneNum,omitempty"`
	// example:
	//
	// 2
	TablesFailedNum *int32 `json:"tablesFailedNum,omitempty" xml:"tablesFailedNum,omitempty"`
	// example:
	//
	// 123
	TablesPartDoneNum *int32 `json:"tablesPartDoneNum,omitempty" xml:"tablesPartDoneNum,omitempty"`
	// example:
	//
	// BIGQUERY
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetMmsDataSourceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetMmsDataSourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMmsDataSourceResponseBodyData) SetAgentIsOnline(v bool) *GetMmsDataSourceResponseBodyData {
	s.AgentIsOnline = &v
	return s
}

func (s *GetMmsDataSourceResponseBodyData) SetConfig(v []*GetMmsDataSourceResponseBodyDataConfig) *GetMmsDataSourceResponseBodyData {
	s.Config = v
	return s
}

func (s *GetMmsDataSourceResponseBodyData) SetCreateTime(v string) *GetMmsDataSourceResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetMmsDataSourceResponseBodyData) SetDbNum(v int32) *GetMmsDataSourceResponseBodyData {
	s.DbNum = &v
	return s
}

func (s *GetMmsDataSourceResponseBodyData) SetErrMsg(v string) *GetMmsDataSourceResponseBodyData {
	s.ErrMsg = &v
	return s
}

func (s *GetMmsDataSourceResponseBodyData) SetId(v int64) *GetMmsDataSourceResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetMmsDataSourceResponseBodyData) SetLastUpdateTime(v string) *GetMmsDataSourceResponseBodyData {
	s.LastUpdateTime = &v
	return s
}

func (s *GetMmsDataSourceResponseBodyData) SetName(v string) *GetMmsDataSourceResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetMmsDataSourceResponseBodyData) SetNetworklink(v string) *GetMmsDataSourceResponseBodyData {
	s.Networklink = &v
	return s
}

func (s *GetMmsDataSourceResponseBodyData) SetPartitionNum(v int32) *GetMmsDataSourceResponseBodyData {
	s.PartitionNum = &v
	return s
}

func (s *GetMmsDataSourceResponseBodyData) SetPartitionsDoingNum(v int32) *GetMmsDataSourceResponseBodyData {
	s.PartitionsDoingNum = &v
	return s
}

func (s *GetMmsDataSourceResponseBodyData) SetPartitionsDoneNum(v int32) *GetMmsDataSourceResponseBodyData {
	s.PartitionsDoneNum = &v
	return s
}

func (s *GetMmsDataSourceResponseBodyData) SetPartitionsFailedNum(v int32) *GetMmsDataSourceResponseBodyData {
	s.PartitionsFailedNum = &v
	return s
}

func (s *GetMmsDataSourceResponseBodyData) SetRegion(v string) *GetMmsDataSourceResponseBodyData {
	s.Region = &v
	return s
}

func (s *GetMmsDataSourceResponseBodyData) SetScanId(v int64) *GetMmsDataSourceResponseBodyData {
	s.ScanId = &v
	return s
}

func (s *GetMmsDataSourceResponseBodyData) SetStatus(v string) *GetMmsDataSourceResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetMmsDataSourceResponseBodyData) SetTableNum(v int32) *GetMmsDataSourceResponseBodyData {
	s.TableNum = &v
	return s
}

func (s *GetMmsDataSourceResponseBodyData) SetTablesDoingNum(v int32) *GetMmsDataSourceResponseBodyData {
	s.TablesDoingNum = &v
	return s
}

func (s *GetMmsDataSourceResponseBodyData) SetTablesDoneNum(v int32) *GetMmsDataSourceResponseBodyData {
	s.TablesDoneNum = &v
	return s
}

func (s *GetMmsDataSourceResponseBodyData) SetTablesFailedNum(v int32) *GetMmsDataSourceResponseBodyData {
	s.TablesFailedNum = &v
	return s
}

func (s *GetMmsDataSourceResponseBodyData) SetTablesPartDoneNum(v int32) *GetMmsDataSourceResponseBodyData {
	s.TablesPartDoneNum = &v
	return s
}

func (s *GetMmsDataSourceResponseBodyData) SetType(v string) *GetMmsDataSourceResponseBodyData {
	s.Type = &v
	return s
}

type GetMmsDataSourceResponseBodyDataConfig struct {
	Desc  *string   `json:"desc,omitempty" xml:"desc,omitempty"`
	Enums []*string `json:"enums,omitempty" xml:"enums,omitempty" type:"Repeated"`
	// example:
	//
	// basic_group
	Group *string `json:"group,omitempty" xml:"group,omitempty"`
	// example:
	//
	// bigquery.range.partition.migrate.type
	Key  *string `json:"key,omitempty" xml:"key,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// Cluster or Partition
	PlaceHolder *string `json:"placeHolder,omitempty" xml:"placeHolder,omitempty"`
	// example:
	//
	// true
	Required *bool `json:"required,omitempty" xml:"required,omitempty"`
	// example:
	//
	// .keytab
	SubType *string `json:"subType,omitempty" xml:"subType,omitempty"`
	// example:
	//
	// string
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// Partition
	Value interface{} `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetMmsDataSourceResponseBodyDataConfig) String() string {
	return tea.Prettify(s)
}

func (s GetMmsDataSourceResponseBodyDataConfig) GoString() string {
	return s.String()
}

func (s *GetMmsDataSourceResponseBodyDataConfig) SetDesc(v string) *GetMmsDataSourceResponseBodyDataConfig {
	s.Desc = &v
	return s
}

func (s *GetMmsDataSourceResponseBodyDataConfig) SetEnums(v []*string) *GetMmsDataSourceResponseBodyDataConfig {
	s.Enums = v
	return s
}

func (s *GetMmsDataSourceResponseBodyDataConfig) SetGroup(v string) *GetMmsDataSourceResponseBodyDataConfig {
	s.Group = &v
	return s
}

func (s *GetMmsDataSourceResponseBodyDataConfig) SetKey(v string) *GetMmsDataSourceResponseBodyDataConfig {
	s.Key = &v
	return s
}

func (s *GetMmsDataSourceResponseBodyDataConfig) SetName(v string) *GetMmsDataSourceResponseBodyDataConfig {
	s.Name = &v
	return s
}

func (s *GetMmsDataSourceResponseBodyDataConfig) SetPlaceHolder(v string) *GetMmsDataSourceResponseBodyDataConfig {
	s.PlaceHolder = &v
	return s
}

func (s *GetMmsDataSourceResponseBodyDataConfig) SetRequired(v bool) *GetMmsDataSourceResponseBodyDataConfig {
	s.Required = &v
	return s
}

func (s *GetMmsDataSourceResponseBodyDataConfig) SetSubType(v string) *GetMmsDataSourceResponseBodyDataConfig {
	s.SubType = &v
	return s
}

func (s *GetMmsDataSourceResponseBodyDataConfig) SetType(v string) *GetMmsDataSourceResponseBodyDataConfig {
	s.Type = &v
	return s
}

func (s *GetMmsDataSourceResponseBodyDataConfig) SetValue(v interface{}) *GetMmsDataSourceResponseBodyDataConfig {
	s.Value = v
	return s
}

type GetMmsDataSourceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMmsDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMmsDataSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMmsDataSourceResponse) GoString() string {
	return s.String()
}

func (s *GetMmsDataSourceResponse) SetHeaders(v map[string]*string) *GetMmsDataSourceResponse {
	s.Headers = v
	return s
}

func (s *GetMmsDataSourceResponse) SetStatusCode(v int32) *GetMmsDataSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMmsDataSourceResponse) SetBody(v *GetMmsDataSourceResponseBody) *GetMmsDataSourceResponse {
	s.Body = v
	return s
}

type GetMmsDbResponseBody struct {
	Data *GetMmsDbResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 90D64EB6-2962-5B1C-A039-BC41C8176C7F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetMmsDbResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMmsDbResponseBody) GoString() string {
	return s.String()
}

func (s *GetMmsDbResponseBody) SetData(v *GetMmsDbResponseBodyData) *GetMmsDbResponseBody {
	s.Data = v
	return s
}

func (s *GetMmsDbResponseBody) SetRequestId(v string) *GetMmsDbResponseBody {
	s.RequestId = &v
	return s
}

type GetMmsDbResponseBodyData struct {
	// example:
	//
	// for mms_test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// {}
	Extra *string `json:"extra,omitempty" xml:"extra,omitempty"`
	// example:
	//
	// 63
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// last ddl time
	//
	// example:
	//
	// 2024-12-17 15:44:42
	LastDdlTime *string `json:"lastDdlTime,omitempty" xml:"lastDdlTime,omitempty"`
	// example:
	//
	// hdfs://master-1-1.c-6fc187819ed6bae0.cn-shanghai.emr.aliyuncs.com:9000/user/hive/warehouse
	Location *string `json:"location,omitempty" xml:"location,omitempty"`
	// example:
	//
	// mms_test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 2323
	NumRows *int64 `json:"numRows,omitempty" xml:"numRows,omitempty"`
	// example:
	//
	// System user
	Owner *string `json:"owner,omitempty" xml:"owner,omitempty"`
	// example:
	//
	// 2000
	Partitions *int32 `json:"partitions,omitempty" xml:"partitions,omitempty"`
	// example:
	//
	// 200
	PartitionsDoing *int32 `json:"partitionsDoing,omitempty" xml:"partitionsDoing,omitempty"`
	// example:
	//
	// 1400
	PartitionsDone *int32 `json:"partitionsDone,omitempty" xml:"partitionsDone,omitempty"`
	// example:
	//
	// 400
	PartitionsFailed *int32 `json:"partitionsFailed,omitempty" xml:"partitionsFailed,omitempty"`
	// example:
	//
	// 323232332
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// example:
	//
	// 2000017
	SourceId *int64 `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// example:
	//
	// demo
	SourceName *string `json:"sourceName,omitempty" xml:"sourceName,omitempty"`
	// example:
	//
	// DOING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 200
	Tables *int32 `json:"tables,omitempty" xml:"tables,omitempty"`
	// example:
	//
	// 20
	TablesDoing *int32 `json:"tablesDoing,omitempty" xml:"tablesDoing,omitempty"`
	// example:
	//
	// 120
	TablesDone *int32 `json:"tablesDone,omitempty" xml:"tablesDone,omitempty"`
	// example:
	//
	// 20
	TablesFailed *int32 `json:"tablesFailed,omitempty" xml:"tablesFailed,omitempty"`
	// example:
	//
	// 20
	TablesPartDone *int32 `json:"tablesPartDone,omitempty" xml:"tablesPartDone,omitempty"`
	// example:
	//
	// true
	Updated *bool `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s GetMmsDbResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetMmsDbResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMmsDbResponseBodyData) SetDescription(v string) *GetMmsDbResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetMmsDbResponseBodyData) SetExtra(v string) *GetMmsDbResponseBodyData {
	s.Extra = &v
	return s
}

func (s *GetMmsDbResponseBodyData) SetId(v int64) *GetMmsDbResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetMmsDbResponseBodyData) SetLastDdlTime(v string) *GetMmsDbResponseBodyData {
	s.LastDdlTime = &v
	return s
}

func (s *GetMmsDbResponseBodyData) SetLocation(v string) *GetMmsDbResponseBodyData {
	s.Location = &v
	return s
}

func (s *GetMmsDbResponseBodyData) SetName(v string) *GetMmsDbResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetMmsDbResponseBodyData) SetNumRows(v int64) *GetMmsDbResponseBodyData {
	s.NumRows = &v
	return s
}

func (s *GetMmsDbResponseBodyData) SetOwner(v string) *GetMmsDbResponseBodyData {
	s.Owner = &v
	return s
}

func (s *GetMmsDbResponseBodyData) SetPartitions(v int32) *GetMmsDbResponseBodyData {
	s.Partitions = &v
	return s
}

func (s *GetMmsDbResponseBodyData) SetPartitionsDoing(v int32) *GetMmsDbResponseBodyData {
	s.PartitionsDoing = &v
	return s
}

func (s *GetMmsDbResponseBodyData) SetPartitionsDone(v int32) *GetMmsDbResponseBodyData {
	s.PartitionsDone = &v
	return s
}

func (s *GetMmsDbResponseBodyData) SetPartitionsFailed(v int32) *GetMmsDbResponseBodyData {
	s.PartitionsFailed = &v
	return s
}

func (s *GetMmsDbResponseBodyData) SetSize(v int64) *GetMmsDbResponseBodyData {
	s.Size = &v
	return s
}

func (s *GetMmsDbResponseBodyData) SetSourceId(v int64) *GetMmsDbResponseBodyData {
	s.SourceId = &v
	return s
}

func (s *GetMmsDbResponseBodyData) SetSourceName(v string) *GetMmsDbResponseBodyData {
	s.SourceName = &v
	return s
}

func (s *GetMmsDbResponseBodyData) SetStatus(v string) *GetMmsDbResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetMmsDbResponseBodyData) SetTables(v int32) *GetMmsDbResponseBodyData {
	s.Tables = &v
	return s
}

func (s *GetMmsDbResponseBodyData) SetTablesDoing(v int32) *GetMmsDbResponseBodyData {
	s.TablesDoing = &v
	return s
}

func (s *GetMmsDbResponseBodyData) SetTablesDone(v int32) *GetMmsDbResponseBodyData {
	s.TablesDone = &v
	return s
}

func (s *GetMmsDbResponseBodyData) SetTablesFailed(v int32) *GetMmsDbResponseBodyData {
	s.TablesFailed = &v
	return s
}

func (s *GetMmsDbResponseBodyData) SetTablesPartDone(v int32) *GetMmsDbResponseBodyData {
	s.TablesPartDone = &v
	return s
}

func (s *GetMmsDbResponseBodyData) SetUpdated(v bool) *GetMmsDbResponseBodyData {
	s.Updated = &v
	return s
}

type GetMmsDbResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMmsDbResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMmsDbResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMmsDbResponse) GoString() string {
	return s.String()
}

func (s *GetMmsDbResponse) SetHeaders(v map[string]*string) *GetMmsDbResponse {
	s.Headers = v
	return s
}

func (s *GetMmsDbResponse) SetStatusCode(v int32) *GetMmsDbResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMmsDbResponse) SetBody(v *GetMmsDbResponseBody) *GetMmsDbResponse {
	s.Body = v
	return s
}

type GetMmsFetchMetadataJobResponseBody struct {
	Data *GetMmsFetchMetadataJobResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 5CA6292A-E301-5CD8-B4E2-AF060F99147B
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetMmsFetchMetadataJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMmsFetchMetadataJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetMmsFetchMetadataJobResponseBody) SetData(v *GetMmsFetchMetadataJobResponseBodyData) *GetMmsFetchMetadataJobResponseBody {
	s.Data = v
	return s
}

func (s *GetMmsFetchMetadataJobResponseBody) SetRequestId(v string) *GetMmsFetchMetadataJobResponseBody {
	s.RequestId = &v
	return s
}

type GetMmsFetchMetadataJobResponseBodyData struct {
	// example:
	//
	// 2024-12-16 19:10:07
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// unexpected exception
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// example:
	//
	// 1000002
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 5000
	Progress *float32 `json:"progress,omitempty" xml:"progress,omitempty"`
	// example:
	//
	// {"databases":5,"tables":75,"partitions":215}
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
	// example:
	//
	// 2000015
	SourceId *int64 `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// example:
	//
	// 2024-12-16 19:09:37
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// SCAN_DOING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetMmsFetchMetadataJobResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetMmsFetchMetadataJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMmsFetchMetadataJobResponseBodyData) SetEndTime(v string) *GetMmsFetchMetadataJobResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *GetMmsFetchMetadataJobResponseBodyData) SetErrorMsg(v string) *GetMmsFetchMetadataJobResponseBodyData {
	s.ErrorMsg = &v
	return s
}

func (s *GetMmsFetchMetadataJobResponseBodyData) SetId(v int64) *GetMmsFetchMetadataJobResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetMmsFetchMetadataJobResponseBodyData) SetProgress(v float32) *GetMmsFetchMetadataJobResponseBodyData {
	s.Progress = &v
	return s
}

func (s *GetMmsFetchMetadataJobResponseBodyData) SetResult(v string) *GetMmsFetchMetadataJobResponseBodyData {
	s.Result = &v
	return s
}

func (s *GetMmsFetchMetadataJobResponseBodyData) SetSourceId(v int64) *GetMmsFetchMetadataJobResponseBodyData {
	s.SourceId = &v
	return s
}

func (s *GetMmsFetchMetadataJobResponseBodyData) SetStartTime(v string) *GetMmsFetchMetadataJobResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *GetMmsFetchMetadataJobResponseBodyData) SetStatus(v string) *GetMmsFetchMetadataJobResponseBodyData {
	s.Status = &v
	return s
}

type GetMmsFetchMetadataJobResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMmsFetchMetadataJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMmsFetchMetadataJobResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMmsFetchMetadataJobResponse) GoString() string {
	return s.String()
}

func (s *GetMmsFetchMetadataJobResponse) SetHeaders(v map[string]*string) *GetMmsFetchMetadataJobResponse {
	s.Headers = v
	return s
}

func (s *GetMmsFetchMetadataJobResponse) SetStatusCode(v int32) *GetMmsFetchMetadataJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMmsFetchMetadataJobResponse) SetBody(v *GetMmsFetchMetadataJobResponseBody) *GetMmsFetchMetadataJobResponse {
	s.Body = v
	return s
}

type GetMmsJobResponseBody struct {
	Data *GetMmsJobResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// D9F872FD-5DDE-30A6-8C8A-1B8C6A81059F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetMmsJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMmsJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetMmsJobResponseBody) SetData(v *GetMmsJobResponseBodyData) *GetMmsJobResponseBody {
	s.Data = v
	return s
}

func (s *GetMmsJobResponseBody) SetRequestId(v string) *GetMmsJobResponseBody {
	s.RequestId = &v
	return s
}

type GetMmsJobResponseBodyData struct {
	Config *GetMmsJobResponseBodyDataConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// example:
	//
	// 2024-12-17 15:44:17
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 23
	DbId *int64 `json:"dbId,omitempty" xml:"dbId,omitempty"`
	// example:
	//
	// mms_target
	DstDbName *string `json:"dstDbName,omitempty" xml:"dstDbName,omitempty"`
	// example:
	//
	// default
	DstSchemaName *string `json:"dstSchemaName,omitempty" xml:"dstSchemaName,omitempty"`
	Eta           *string `json:"eta,omitempty" xml:"eta,omitempty"`
	// example:
	//
	// 10
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// migrate_db_1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 2
	SourceId *int64 `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// example:
	//
	// demo
	SourceName *string `json:"sourceName,omitempty" xml:"sourceName,omitempty"`
	// example:
	//
	// mms_test
	SrcDbName *string `json:"srcDbName,omitempty" xml:"srcDbName,omitempty"`
	// example:
	//
	// default
	SrcSchemaName *string `json:"srcSchemaName,omitempty" xml:"srcSchemaName,omitempty"`
	// example:
	//
	// DOING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// false
	Stopped *bool `json:"stopped,omitempty" xml:"stopped,omitempty"`
	// example:
	//
	// 100
	TaskDone *int32 `json:"taskDone,omitempty" xml:"taskDone,omitempty"`
	// example:
	//
	// 100
	TaskNum *int32 `json:"taskNum,omitempty" xml:"taskNum,omitempty"`
	// example:
	//
	// Tables
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetMmsJobResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetMmsJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMmsJobResponseBodyData) SetConfig(v *GetMmsJobResponseBodyDataConfig) *GetMmsJobResponseBodyData {
	s.Config = v
	return s
}

func (s *GetMmsJobResponseBodyData) SetCreateTime(v string) *GetMmsJobResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetMmsJobResponseBodyData) SetDbId(v int64) *GetMmsJobResponseBodyData {
	s.DbId = &v
	return s
}

func (s *GetMmsJobResponseBodyData) SetDstDbName(v string) *GetMmsJobResponseBodyData {
	s.DstDbName = &v
	return s
}

func (s *GetMmsJobResponseBodyData) SetDstSchemaName(v string) *GetMmsJobResponseBodyData {
	s.DstSchemaName = &v
	return s
}

func (s *GetMmsJobResponseBodyData) SetEta(v string) *GetMmsJobResponseBodyData {
	s.Eta = &v
	return s
}

func (s *GetMmsJobResponseBodyData) SetId(v int64) *GetMmsJobResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetMmsJobResponseBodyData) SetName(v string) *GetMmsJobResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetMmsJobResponseBodyData) SetSourceId(v int64) *GetMmsJobResponseBodyData {
	s.SourceId = &v
	return s
}

func (s *GetMmsJobResponseBodyData) SetSourceName(v string) *GetMmsJobResponseBodyData {
	s.SourceName = &v
	return s
}

func (s *GetMmsJobResponseBodyData) SetSrcDbName(v string) *GetMmsJobResponseBodyData {
	s.SrcDbName = &v
	return s
}

func (s *GetMmsJobResponseBodyData) SetSrcSchemaName(v string) *GetMmsJobResponseBodyData {
	s.SrcSchemaName = &v
	return s
}

func (s *GetMmsJobResponseBodyData) SetStatus(v string) *GetMmsJobResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetMmsJobResponseBodyData) SetStopped(v bool) *GetMmsJobResponseBodyData {
	s.Stopped = &v
	return s
}

func (s *GetMmsJobResponseBodyData) SetTaskDone(v int32) *GetMmsJobResponseBodyData {
	s.TaskDone = &v
	return s
}

func (s *GetMmsJobResponseBodyData) SetTaskNum(v int32) *GetMmsJobResponseBodyData {
	s.TaskNum = &v
	return s
}

func (s *GetMmsJobResponseBodyData) SetType(v string) *GetMmsJobResponseBodyData {
	s.Type = &v
	return s
}

type GetMmsJobResponseBodyDataConfig struct {
	ColumnMapping      map[string]*string     `json:"columnMapping,omitempty" xml:"columnMapping,omitempty"`
	EnableVerification *bool                  `json:"enableVerification,omitempty" xml:"enableVerification,omitempty"`
	Increment          *bool                  `json:"increment,omitempty" xml:"increment,omitempty"`
	Others             map[string]interface{} `json:"others,omitempty" xml:"others,omitempty"`
	PartitionFilters   map[string]*string     `json:"partitionFilters,omitempty" xml:"partitionFilters,omitempty"`
	Partitions         []*int64               `json:"partitions,omitempty" xml:"partitions,omitempty" type:"Repeated"`
	SchemaOnly         *bool                  `json:"schemaOnly,omitempty" xml:"schemaOnly,omitempty"`
	TableBlackList     []*string              `json:"tableBlackList,omitempty" xml:"tableBlackList,omitempty" type:"Repeated"`
	TableMapping       map[string]*string     `json:"tableMapping,omitempty" xml:"tableMapping,omitempty"`
	TableWhiteList     []*string              `json:"tableWhiteList,omitempty" xml:"tableWhiteList,omitempty" type:"Repeated"`
	Tables             []*string              `json:"tables,omitempty" xml:"tables,omitempty" type:"Repeated"`
	TaskType           *string                `json:"taskType,omitempty" xml:"taskType,omitempty"`
	TunnelQuota        *string                `json:"tunnelQuota,omitempty" xml:"tunnelQuota,omitempty"`
}

func (s GetMmsJobResponseBodyDataConfig) String() string {
	return tea.Prettify(s)
}

func (s GetMmsJobResponseBodyDataConfig) GoString() string {
	return s.String()
}

func (s *GetMmsJobResponseBodyDataConfig) SetColumnMapping(v map[string]*string) *GetMmsJobResponseBodyDataConfig {
	s.ColumnMapping = v
	return s
}

func (s *GetMmsJobResponseBodyDataConfig) SetEnableVerification(v bool) *GetMmsJobResponseBodyDataConfig {
	s.EnableVerification = &v
	return s
}

func (s *GetMmsJobResponseBodyDataConfig) SetIncrement(v bool) *GetMmsJobResponseBodyDataConfig {
	s.Increment = &v
	return s
}

func (s *GetMmsJobResponseBodyDataConfig) SetOthers(v map[string]interface{}) *GetMmsJobResponseBodyDataConfig {
	s.Others = v
	return s
}

func (s *GetMmsJobResponseBodyDataConfig) SetPartitionFilters(v map[string]*string) *GetMmsJobResponseBodyDataConfig {
	s.PartitionFilters = v
	return s
}

func (s *GetMmsJobResponseBodyDataConfig) SetPartitions(v []*int64) *GetMmsJobResponseBodyDataConfig {
	s.Partitions = v
	return s
}

func (s *GetMmsJobResponseBodyDataConfig) SetSchemaOnly(v bool) *GetMmsJobResponseBodyDataConfig {
	s.SchemaOnly = &v
	return s
}

func (s *GetMmsJobResponseBodyDataConfig) SetTableBlackList(v []*string) *GetMmsJobResponseBodyDataConfig {
	s.TableBlackList = v
	return s
}

func (s *GetMmsJobResponseBodyDataConfig) SetTableMapping(v map[string]*string) *GetMmsJobResponseBodyDataConfig {
	s.TableMapping = v
	return s
}

func (s *GetMmsJobResponseBodyDataConfig) SetTableWhiteList(v []*string) *GetMmsJobResponseBodyDataConfig {
	s.TableWhiteList = v
	return s
}

func (s *GetMmsJobResponseBodyDataConfig) SetTables(v []*string) *GetMmsJobResponseBodyDataConfig {
	s.Tables = v
	return s
}

func (s *GetMmsJobResponseBodyDataConfig) SetTaskType(v string) *GetMmsJobResponseBodyDataConfig {
	s.TaskType = &v
	return s
}

func (s *GetMmsJobResponseBodyDataConfig) SetTunnelQuota(v string) *GetMmsJobResponseBodyDataConfig {
	s.TunnelQuota = &v
	return s
}

type GetMmsJobResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMmsJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMmsJobResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMmsJobResponse) GoString() string {
	return s.String()
}

func (s *GetMmsJobResponse) SetHeaders(v map[string]*string) *GetMmsJobResponse {
	s.Headers = v
	return s
}

func (s *GetMmsJobResponse) SetStatusCode(v int32) *GetMmsJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMmsJobResponse) SetBody(v *GetMmsJobResponseBody) *GetMmsJobResponse {
	s.Body = v
	return s
}

type GetMmsPartitionResponseBody struct {
	Data *GetMmsPartitionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// D9F872FD-5DDE-30A6-8C8A-1B8C6A81059F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetMmsPartitionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMmsPartitionResponseBody) GoString() string {
	return s.String()
}

func (s *GetMmsPartitionResponseBody) SetData(v *GetMmsPartitionResponseBodyData) *GetMmsPartitionResponseBody {
	s.Data = v
	return s
}

func (s *GetMmsPartitionResponseBody) SetRequestId(v string) *GetMmsPartitionResponseBody {
	s.RequestId = &v
	return s
}

type GetMmsPartitionResponseBodyData struct {
	// example:
	//
	// 2
	DbId *int64 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// example:
	//
	// d1
	DbName *string `json:"dbName,omitempty" xml:"dbName,omitempty"`
	// example:
	//
	// 2323
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// lastDdlTime
	//
	// example:
	//
	// 2024-12-17 15:44:42
	LastDdlTime *string `json:"lastDdlTime,omitempty" xml:"lastDdlTime,omitempty"`
	// example:
	//
	// 2323
	NumRows *int64 `json:"numRows,omitempty" xml:"numRows,omitempty"`
	// example:
	//
	// 12323
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// example:
	//
	// 200018
	SourceId *int64 `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// example:
	//
	// demo
	SourceName *string `json:"sourceName,omitempty" xml:"sourceName,omitempty"`
	// example:
	//
	// DONE
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 23
	TableId *int64 `json:"tableId,omitempty" xml:"tableId,omitempty"`
	// example:
	//
	// t1
	TableName *string `json:"tableName,omitempty" xml:"tableName,omitempty"`
	// example:
	//
	// false
	Updated *bool `json:"updated,omitempty" xml:"updated,omitempty"`
	// example:
	//
	// p1=1/p2=abc
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetMmsPartitionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetMmsPartitionResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMmsPartitionResponseBodyData) SetDbId(v int64) *GetMmsPartitionResponseBodyData {
	s.DbId = &v
	return s
}

func (s *GetMmsPartitionResponseBodyData) SetDbName(v string) *GetMmsPartitionResponseBodyData {
	s.DbName = &v
	return s
}

func (s *GetMmsPartitionResponseBodyData) SetId(v int64) *GetMmsPartitionResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetMmsPartitionResponseBodyData) SetLastDdlTime(v string) *GetMmsPartitionResponseBodyData {
	s.LastDdlTime = &v
	return s
}

func (s *GetMmsPartitionResponseBodyData) SetNumRows(v int64) *GetMmsPartitionResponseBodyData {
	s.NumRows = &v
	return s
}

func (s *GetMmsPartitionResponseBodyData) SetSize(v int64) *GetMmsPartitionResponseBodyData {
	s.Size = &v
	return s
}

func (s *GetMmsPartitionResponseBodyData) SetSourceId(v int64) *GetMmsPartitionResponseBodyData {
	s.SourceId = &v
	return s
}

func (s *GetMmsPartitionResponseBodyData) SetSourceName(v string) *GetMmsPartitionResponseBodyData {
	s.SourceName = &v
	return s
}

func (s *GetMmsPartitionResponseBodyData) SetStatus(v string) *GetMmsPartitionResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetMmsPartitionResponseBodyData) SetTableId(v int64) *GetMmsPartitionResponseBodyData {
	s.TableId = &v
	return s
}

func (s *GetMmsPartitionResponseBodyData) SetTableName(v string) *GetMmsPartitionResponseBodyData {
	s.TableName = &v
	return s
}

func (s *GetMmsPartitionResponseBodyData) SetUpdated(v bool) *GetMmsPartitionResponseBodyData {
	s.Updated = &v
	return s
}

func (s *GetMmsPartitionResponseBodyData) SetValue(v string) *GetMmsPartitionResponseBodyData {
	s.Value = &v
	return s
}

type GetMmsPartitionResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMmsPartitionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMmsPartitionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMmsPartitionResponse) GoString() string {
	return s.String()
}

func (s *GetMmsPartitionResponse) SetHeaders(v map[string]*string) *GetMmsPartitionResponse {
	s.Headers = v
	return s
}

func (s *GetMmsPartitionResponse) SetStatusCode(v int32) *GetMmsPartitionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMmsPartitionResponse) SetBody(v *GetMmsPartitionResponseBody) *GetMmsPartitionResponse {
	s.Body = v
	return s
}

type GetMmsTableResponseBody struct {
	Data *GetMmsTableResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// E7FB14F1-4ACD-5C73-A755-B302D70AB9AD
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetMmsTableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMmsTableResponseBody) GoString() string {
	return s.String()
}

func (s *GetMmsTableResponseBody) SetData(v *GetMmsTableResponseBodyData) *GetMmsTableResponseBody {
	s.Data = v
	return s
}

func (s *GetMmsTableResponseBody) SetRequestId(v string) *GetMmsTableResponseBody {
	s.RequestId = &v
	return s
}

type GetMmsTableResponseBodyData struct {
	// example:
	//
	// 3
	DbId *int64 `json:"dbId,omitempty" xml:"dbId,omitempty"`
	// example:
	//
	// mms_test
	DbName *string `json:"dbName,omitempty" xml:"dbName,omitempty"`
	// example:
	//
	// {"mapkey.delim":":","collection.delim":",","serialization.format":"|","field.delim":"|"}
	Extra *string `json:"extra,omitempty" xml:"extra,omitempty"`
	// example:
	//
	// true
	HasPartitions *bool `json:"hasPartitions,omitempty" xml:"hasPartitions,omitempty"`
	// table ID
	//
	// example:
	//
	// 22
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// inputFormat
	//
	// example:
	//
	// org.apache.hadoop.hive.ql.io.parquet.MapredParquetInputFormat
	InputFormat *string `json:"inputFormat,omitempty" xml:"inputFormat,omitempty"`
	// lastDdlTime
	//
	// example:
	//
	// 2024-12-17 15:44:42
	LastDdlTime *string `json:"lastDdlTime,omitempty" xml:"lastDdlTime,omitempty"`
	// example:
	//
	// | hdfs://master-1-1.c-c127cd184bb029ea.cn-zhangjiakou.emr.aliyuncs.com:9000/user/hive/warehouse/demo
	Location *string `json:"location,omitempty" xml:"location,omitempty"`
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 233232
	NumRows *int64 `json:"numRows,omitempty" xml:"numRows,omitempty"`
	// outputFormat
	//
	// example:
	//
	// org.apache.hadoop.hive.ql.io.parquet.MapredParquetOutputFormat
	OutputFormat *string `json:"outputFormat,omitempty" xml:"outputFormat,omitempty"`
	// example:
	//
	// Hive
	Owner *string `json:"owner,omitempty" xml:"owner,omitempty"`
	// example:
	//
	// 100
	Partitions *int32 `json:"partitions,omitempty" xml:"partitions,omitempty"`
	// example:
	//
	// 20
	PartitionsDoing *int32 `json:"partitionsDoing,omitempty" xml:"partitionsDoing,omitempty"`
	// example:
	//
	// 80
	PartitionsDone *int32 `json:"partitionsDone,omitempty" xml:"partitionsDone,omitempty"`
	// example:
	//
	// 0
	PartitionsFailed *int32                             `json:"partitionsFailed,omitempty" xml:"partitionsFailed,omitempty"`
	Schema           *GetMmsTableResponseBodyDataSchema `json:"schema,omitempty" xml:"schema,omitempty" type:"Struct"`
	// serde
	//
	// example:
	//
	// org.apache.hadoop.hive.ql.io.parquet.serde.ParquetHiveSerDe
	Serde *string `json:"serde,omitempty" xml:"serde,omitempty"`
	// example:
	//
	// 23232
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// example:
	//
	// 2000028
	SourceId *int64 `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// example:
	//
	// demo
	SourceName *string `json:"sourceName,omitempty" xml:"sourceName,omitempty"`
	// example:
	//
	// DONE
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// MANAGED_TABLED
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// false
	Updated *bool `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s GetMmsTableResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetMmsTableResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMmsTableResponseBodyData) SetDbId(v int64) *GetMmsTableResponseBodyData {
	s.DbId = &v
	return s
}

func (s *GetMmsTableResponseBodyData) SetDbName(v string) *GetMmsTableResponseBodyData {
	s.DbName = &v
	return s
}

func (s *GetMmsTableResponseBodyData) SetExtra(v string) *GetMmsTableResponseBodyData {
	s.Extra = &v
	return s
}

func (s *GetMmsTableResponseBodyData) SetHasPartitions(v bool) *GetMmsTableResponseBodyData {
	s.HasPartitions = &v
	return s
}

func (s *GetMmsTableResponseBodyData) SetId(v int64) *GetMmsTableResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetMmsTableResponseBodyData) SetInputFormat(v string) *GetMmsTableResponseBodyData {
	s.InputFormat = &v
	return s
}

func (s *GetMmsTableResponseBodyData) SetLastDdlTime(v string) *GetMmsTableResponseBodyData {
	s.LastDdlTime = &v
	return s
}

func (s *GetMmsTableResponseBodyData) SetLocation(v string) *GetMmsTableResponseBodyData {
	s.Location = &v
	return s
}

func (s *GetMmsTableResponseBodyData) SetName(v string) *GetMmsTableResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetMmsTableResponseBodyData) SetNumRows(v int64) *GetMmsTableResponseBodyData {
	s.NumRows = &v
	return s
}

func (s *GetMmsTableResponseBodyData) SetOutputFormat(v string) *GetMmsTableResponseBodyData {
	s.OutputFormat = &v
	return s
}

func (s *GetMmsTableResponseBodyData) SetOwner(v string) *GetMmsTableResponseBodyData {
	s.Owner = &v
	return s
}

func (s *GetMmsTableResponseBodyData) SetPartitions(v int32) *GetMmsTableResponseBodyData {
	s.Partitions = &v
	return s
}

func (s *GetMmsTableResponseBodyData) SetPartitionsDoing(v int32) *GetMmsTableResponseBodyData {
	s.PartitionsDoing = &v
	return s
}

func (s *GetMmsTableResponseBodyData) SetPartitionsDone(v int32) *GetMmsTableResponseBodyData {
	s.PartitionsDone = &v
	return s
}

func (s *GetMmsTableResponseBodyData) SetPartitionsFailed(v int32) *GetMmsTableResponseBodyData {
	s.PartitionsFailed = &v
	return s
}

func (s *GetMmsTableResponseBodyData) SetSchema(v *GetMmsTableResponseBodyDataSchema) *GetMmsTableResponseBodyData {
	s.Schema = v
	return s
}

func (s *GetMmsTableResponseBodyData) SetSerde(v string) *GetMmsTableResponseBodyData {
	s.Serde = &v
	return s
}

func (s *GetMmsTableResponseBodyData) SetSize(v int64) *GetMmsTableResponseBodyData {
	s.Size = &v
	return s
}

func (s *GetMmsTableResponseBodyData) SetSourceId(v int64) *GetMmsTableResponseBodyData {
	s.SourceId = &v
	return s
}

func (s *GetMmsTableResponseBodyData) SetSourceName(v string) *GetMmsTableResponseBodyData {
	s.SourceName = &v
	return s
}

func (s *GetMmsTableResponseBodyData) SetStatus(v string) *GetMmsTableResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetMmsTableResponseBodyData) SetType(v string) *GetMmsTableResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetMmsTableResponseBodyData) SetUpdated(v bool) *GetMmsTableResponseBodyData {
	s.Updated = &v
	return s
}

type GetMmsTableResponseBodyDataSchema struct {
	Columns []*GetMmsTableResponseBodyDataSchemaColumns `json:"columns,omitempty" xml:"columns,omitempty" type:"Repeated"`
	// example:
	//
	// for mms test
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	// example:
	//
	// test
	Name       *string                                        `json:"name,omitempty" xml:"name,omitempty"`
	Partitions []*GetMmsTableResponseBodyDataSchemaPartitions `json:"partitions,omitempty" xml:"partitions,omitempty" type:"Repeated"`
}

func (s GetMmsTableResponseBodyDataSchema) String() string {
	return tea.Prettify(s)
}

func (s GetMmsTableResponseBodyDataSchema) GoString() string {
	return s.String()
}

func (s *GetMmsTableResponseBodyDataSchema) SetColumns(v []*GetMmsTableResponseBodyDataSchemaColumns) *GetMmsTableResponseBodyDataSchema {
	s.Columns = v
	return s
}

func (s *GetMmsTableResponseBodyDataSchema) SetComment(v string) *GetMmsTableResponseBodyDataSchema {
	s.Comment = &v
	return s
}

func (s *GetMmsTableResponseBodyDataSchema) SetName(v string) *GetMmsTableResponseBodyDataSchema {
	s.Name = &v
	return s
}

func (s *GetMmsTableResponseBodyDataSchema) SetPartitions(v []*GetMmsTableResponseBodyDataSchemaPartitions) *GetMmsTableResponseBodyDataSchema {
	s.Partitions = v
	return s
}

type GetMmsTableResponseBodyDataSchemaColumns struct {
	// example:
	//
	// user id
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	// example:
	//
	// 10
	DefaultValue *string `json:"defaultValue,omitempty" xml:"defaultValue,omitempty"`
	// example:
	//
	// user_id
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// false
	Nullable *bool `json:"nullable,omitempty" xml:"nullable,omitempty"`
	// example:
	//
	// bigint
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetMmsTableResponseBodyDataSchemaColumns) String() string {
	return tea.Prettify(s)
}

func (s GetMmsTableResponseBodyDataSchemaColumns) GoString() string {
	return s.String()
}

func (s *GetMmsTableResponseBodyDataSchemaColumns) SetComment(v string) *GetMmsTableResponseBodyDataSchemaColumns {
	s.Comment = &v
	return s
}

func (s *GetMmsTableResponseBodyDataSchemaColumns) SetDefaultValue(v string) *GetMmsTableResponseBodyDataSchemaColumns {
	s.DefaultValue = &v
	return s
}

func (s *GetMmsTableResponseBodyDataSchemaColumns) SetName(v string) *GetMmsTableResponseBodyDataSchemaColumns {
	s.Name = &v
	return s
}

func (s *GetMmsTableResponseBodyDataSchemaColumns) SetNullable(v bool) *GetMmsTableResponseBodyDataSchemaColumns {
	s.Nullable = &v
	return s
}

func (s *GetMmsTableResponseBodyDataSchemaColumns) SetType(v string) *GetMmsTableResponseBodyDataSchemaColumns {
	s.Type = &v
	return s
}

type GetMmsTableResponseBodyDataSchemaPartitions struct {
	// example:
	//
	// first partition level
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	// example:
	//
	// abc
	DefaultValue *string `json:"defaultValue,omitempty" xml:"defaultValue,omitempty"`
	// example:
	//
	// p1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// false
	Nullable *bool `json:"nullable,omitempty" xml:"nullable,omitempty"`
	// example:
	//
	// string
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetMmsTableResponseBodyDataSchemaPartitions) String() string {
	return tea.Prettify(s)
}

func (s GetMmsTableResponseBodyDataSchemaPartitions) GoString() string {
	return s.String()
}

func (s *GetMmsTableResponseBodyDataSchemaPartitions) SetComment(v string) *GetMmsTableResponseBodyDataSchemaPartitions {
	s.Comment = &v
	return s
}

func (s *GetMmsTableResponseBodyDataSchemaPartitions) SetDefaultValue(v string) *GetMmsTableResponseBodyDataSchemaPartitions {
	s.DefaultValue = &v
	return s
}

func (s *GetMmsTableResponseBodyDataSchemaPartitions) SetName(v string) *GetMmsTableResponseBodyDataSchemaPartitions {
	s.Name = &v
	return s
}

func (s *GetMmsTableResponseBodyDataSchemaPartitions) SetNullable(v bool) *GetMmsTableResponseBodyDataSchemaPartitions {
	s.Nullable = &v
	return s
}

func (s *GetMmsTableResponseBodyDataSchemaPartitions) SetType(v string) *GetMmsTableResponseBodyDataSchemaPartitions {
	s.Type = &v
	return s
}

type GetMmsTableResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMmsTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMmsTableResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMmsTableResponse) GoString() string {
	return s.String()
}

func (s *GetMmsTableResponse) SetHeaders(v map[string]*string) *GetMmsTableResponse {
	s.Headers = v
	return s
}

func (s *GetMmsTableResponse) SetStatusCode(v int32) *GetMmsTableResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMmsTableResponse) SetBody(v *GetMmsTableResponseBody) *GetMmsTableResponse {
	s.Body = v
	return s
}

type GetMmsTaskResponseBody struct {
	Data *GetMmsTaskResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 73207140-0FD5-588A-B11A-3CE093924196
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetMmsTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMmsTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetMmsTaskResponseBody) SetData(v *GetMmsTaskResponseBodyData) *GetMmsTaskResponseBody {
	s.Data = v
	return s
}

func (s *GetMmsTaskResponseBody) SetRequestId(v string) *GetMmsTaskResponseBody {
	s.RequestId = &v
	return s
}

type GetMmsTaskResponseBodyData struct {
	// example:
	//
	// 2024-10-25 04:21:01
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 23
	DbId *int64 `json:"dbId,omitempty" xml:"dbId,omitempty"`
	// example:
	//
	// mms_target
	DstDbName *string `json:"dstDbName,omitempty" xml:"dstDbName,omitempty"`
	// example:
	//
	// default
	DstSchemaName *string `json:"dstSchemaName,omitempty" xml:"dstSchemaName,omitempty"`
	// example:
	//
	// table_1
	DstTableName *string `json:"dstTableName,omitempty" xml:"dstTableName,omitempty"`
	// example:
	//
	// 2024-10-25 07:21:01
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 7680
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 87
	JobId *int64 `json:"jobId,omitempty" xml:"jobId,omitempty"`
	// example:
	//
	// test_odps_spark
	JobName *string `json:"jobName,omitempty" xml:"jobName,omitempty"`
	// example:
	//
	// 1
	RetriedTimes *int32 `json:"retriedTimes,omitempty" xml:"retriedTimes,omitempty"`
	// example:
	//
	// true
	Running *bool `json:"running,omitempty" xml:"running,omitempty"`
	// example:
	//
	// 2000015
	SourceId *int64 `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// example:
	//
	// demo
	SourceName *string `json:"sourceName,omitempty" xml:"sourceName,omitempty"`
	// example:
	//
	// mms_test
	SrcDbName *string `json:"srcDbName,omitempty" xml:"srcDbName,omitempty"`
	// example:
	//
	// default
	SrcSchemaName *string `json:"srcSchemaName,omitempty" xml:"srcSchemaName,omitempty"`
	// example:
	//
	// table_1
	SrcTableName *string `json:"srcTableName,omitempty" xml:"srcTableName,omitempty"`
	// example:
	//
	// 2024-10-25 06:21:01
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// DATA_DOING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// false
	Stopped *bool `json:"stopped,omitempty" xml:"stopped,omitempty"`
	// example:
	//
	// 2323
	TableId *int64 `json:"tableId,omitempty" xml:"tableId,omitempty"`
	// example:
	//
	// BIGQUERY
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetMmsTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetMmsTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMmsTaskResponseBodyData) SetCreateTime(v string) *GetMmsTaskResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetMmsTaskResponseBodyData) SetDbId(v int64) *GetMmsTaskResponseBodyData {
	s.DbId = &v
	return s
}

func (s *GetMmsTaskResponseBodyData) SetDstDbName(v string) *GetMmsTaskResponseBodyData {
	s.DstDbName = &v
	return s
}

func (s *GetMmsTaskResponseBodyData) SetDstSchemaName(v string) *GetMmsTaskResponseBodyData {
	s.DstSchemaName = &v
	return s
}

func (s *GetMmsTaskResponseBodyData) SetDstTableName(v string) *GetMmsTaskResponseBodyData {
	s.DstTableName = &v
	return s
}

func (s *GetMmsTaskResponseBodyData) SetEndTime(v string) *GetMmsTaskResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *GetMmsTaskResponseBodyData) SetId(v int64) *GetMmsTaskResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetMmsTaskResponseBodyData) SetJobId(v int64) *GetMmsTaskResponseBodyData {
	s.JobId = &v
	return s
}

func (s *GetMmsTaskResponseBodyData) SetJobName(v string) *GetMmsTaskResponseBodyData {
	s.JobName = &v
	return s
}

func (s *GetMmsTaskResponseBodyData) SetRetriedTimes(v int32) *GetMmsTaskResponseBodyData {
	s.RetriedTimes = &v
	return s
}

func (s *GetMmsTaskResponseBodyData) SetRunning(v bool) *GetMmsTaskResponseBodyData {
	s.Running = &v
	return s
}

func (s *GetMmsTaskResponseBodyData) SetSourceId(v int64) *GetMmsTaskResponseBodyData {
	s.SourceId = &v
	return s
}

func (s *GetMmsTaskResponseBodyData) SetSourceName(v string) *GetMmsTaskResponseBodyData {
	s.SourceName = &v
	return s
}

func (s *GetMmsTaskResponseBodyData) SetSrcDbName(v string) *GetMmsTaskResponseBodyData {
	s.SrcDbName = &v
	return s
}

func (s *GetMmsTaskResponseBodyData) SetSrcSchemaName(v string) *GetMmsTaskResponseBodyData {
	s.SrcSchemaName = &v
	return s
}

func (s *GetMmsTaskResponseBodyData) SetSrcTableName(v string) *GetMmsTaskResponseBodyData {
	s.SrcTableName = &v
	return s
}

func (s *GetMmsTaskResponseBodyData) SetStartTime(v string) *GetMmsTaskResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *GetMmsTaskResponseBodyData) SetStatus(v string) *GetMmsTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetMmsTaskResponseBodyData) SetStopped(v bool) *GetMmsTaskResponseBodyData {
	s.Stopped = &v
	return s
}

func (s *GetMmsTaskResponseBodyData) SetTableId(v int64) *GetMmsTaskResponseBodyData {
	s.TableId = &v
	return s
}

func (s *GetMmsTaskResponseBodyData) SetType(v string) *GetMmsTaskResponseBodyData {
	s.Type = &v
	return s
}

type GetMmsTaskResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMmsTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMmsTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMmsTaskResponse) GoString() string {
	return s.String()
}

func (s *GetMmsTaskResponse) SetHeaders(v map[string]*string) *GetMmsTaskResponse {
	s.Headers = v
	return s
}

func (s *GetMmsTaskResponse) SetStatusCode(v int32) *GetMmsTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMmsTaskResponse) SetBody(v *GetMmsTaskResponseBody) *GetMmsTaskResponse {
	s.Body = v
	return s
}

type GetPackageRequest struct {
	// The project to which the package belongs. This parameter is required if the package is installed in the MaxCompute project.
	//
	// example:
	//
	// projectB
	SourceProject *string `json:"sourceProject,omitempty" xml:"sourceProject,omitempty"`
}

func (s GetPackageRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPackageRequest) GoString() string {
	return s.String()
}

func (s *GetPackageRequest) SetSourceProject(v string) *GetPackageRequest {
	s.SourceProject = &v
	return s
}

type GetPackageResponseBody struct {
	// The returned data.
	Data *GetPackageResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// 040002
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// error message.
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// Indicates whether the request was successful. If this parameter was not empty and the value of this parameter was not 200, the request failed.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0b57ff8316614119858417939e3e54
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetPackageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPackageResponseBody) GoString() string {
	return s.String()
}

func (s *GetPackageResponseBody) SetData(v *GetPackageResponseBodyData) *GetPackageResponseBody {
	s.Data = v
	return s
}

func (s *GetPackageResponseBody) SetErrorCode(v string) *GetPackageResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetPackageResponseBody) SetErrorMsg(v string) *GetPackageResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetPackageResponseBody) SetHttpCode(v int32) *GetPackageResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetPackageResponseBody) SetRequestId(v string) *GetPackageResponseBody {
	s.RequestId = &v
	return s
}

type GetPackageResponseBodyData struct {
	// The projects in which the package is installed.
	AllowedProjectList []*GetPackageResponseBodyDataAllowedProjectList `json:"allowedProjectList,omitempty" xml:"allowedProjectList,omitempty" type:"Repeated"`
	// The details of the resources that are included in the package.
	ResourceList *GetPackageResponseBodyDataResourceList `json:"resourceList,omitempty" xml:"resourceList,omitempty" type:"Struct"`
}

func (s GetPackageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetPackageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPackageResponseBodyData) SetAllowedProjectList(v []*GetPackageResponseBodyDataAllowedProjectList) *GetPackageResponseBodyData {
	s.AllowedProjectList = v
	return s
}

func (s *GetPackageResponseBodyData) SetResourceList(v *GetPackageResponseBodyDataResourceList) *GetPackageResponseBodyData {
	s.ResourceList = v
	return s
}

type GetPackageResponseBodyDataAllowedProjectList struct {
	// The security level for sensitive data.
	//
	// example:
	//
	// 2
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// The name of the MaxCompute project.
	//
	// example:
	//
	// proejctB
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
}

func (s GetPackageResponseBodyDataAllowedProjectList) String() string {
	return tea.Prettify(s)
}

func (s GetPackageResponseBodyDataAllowedProjectList) GoString() string {
	return s.String()
}

func (s *GetPackageResponseBodyDataAllowedProjectList) SetLabel(v string) *GetPackageResponseBodyDataAllowedProjectList {
	s.Label = &v
	return s
}

func (s *GetPackageResponseBodyDataAllowedProjectList) SetProject(v string) *GetPackageResponseBodyDataAllowedProjectList {
	s.Project = &v
	return s
}

type GetPackageResponseBodyDataResourceList struct {
	// The functions.
	Function []*GetPackageResponseBodyDataResourceListFunction `json:"function,omitempty" xml:"function,omitempty" type:"Repeated"`
	// The resources.
	Resource []*GetPackageResponseBodyDataResourceListResource `json:"resource,omitempty" xml:"resource,omitempty" type:"Repeated"`
	// The tables.
	Table []*GetPackageResponseBodyDataResourceListTable `json:"table,omitempty" xml:"table,omitempty" type:"Repeated"`
}

func (s GetPackageResponseBodyDataResourceList) String() string {
	return tea.Prettify(s)
}

func (s GetPackageResponseBodyDataResourceList) GoString() string {
	return s.String()
}

func (s *GetPackageResponseBodyDataResourceList) SetFunction(v []*GetPackageResponseBodyDataResourceListFunction) *GetPackageResponseBodyDataResourceList {
	s.Function = v
	return s
}

func (s *GetPackageResponseBodyDataResourceList) SetResource(v []*GetPackageResponseBodyDataResourceListResource) *GetPackageResponseBodyDataResourceList {
	s.Resource = v
	return s
}

func (s *GetPackageResponseBodyDataResourceList) SetTable(v []*GetPackageResponseBodyDataResourceListTable) *GetPackageResponseBodyDataResourceList {
	s.Table = v
	return s
}

type GetPackageResponseBodyDataResourceListFunction struct {
	// The operations that were performed on the function.
	Actions []*string `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	// The name of the function.
	//
	// example:
	//
	// function_1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The name of schema.
	//
	// example:
	//
	// default
	SchemaName *string `json:"schemaName,omitempty" xml:"schemaName,omitempty"`
}

func (s GetPackageResponseBodyDataResourceListFunction) String() string {
	return tea.Prettify(s)
}

func (s GetPackageResponseBodyDataResourceListFunction) GoString() string {
	return s.String()
}

func (s *GetPackageResponseBodyDataResourceListFunction) SetActions(v []*string) *GetPackageResponseBodyDataResourceListFunction {
	s.Actions = v
	return s
}

func (s *GetPackageResponseBodyDataResourceListFunction) SetName(v string) *GetPackageResponseBodyDataResourceListFunction {
	s.Name = &v
	return s
}

func (s *GetPackageResponseBodyDataResourceListFunction) SetSchemaName(v string) *GetPackageResponseBodyDataResourceListFunction {
	s.SchemaName = &v
	return s
}

type GetPackageResponseBodyDataResourceListResource struct {
	// The operations that were performed on the resource.
	Actions []*string `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	// The name of the resource.
	//
	// example:
	//
	// res_1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The name of schema.
	//
	// example:
	//
	// default
	SchemaName *string `json:"schemaName,omitempty" xml:"schemaName,omitempty"`
}

func (s GetPackageResponseBodyDataResourceListResource) String() string {
	return tea.Prettify(s)
}

func (s GetPackageResponseBodyDataResourceListResource) GoString() string {
	return s.String()
}

func (s *GetPackageResponseBodyDataResourceListResource) SetActions(v []*string) *GetPackageResponseBodyDataResourceListResource {
	s.Actions = v
	return s
}

func (s *GetPackageResponseBodyDataResourceListResource) SetName(v string) *GetPackageResponseBodyDataResourceListResource {
	s.Name = &v
	return s
}

func (s *GetPackageResponseBodyDataResourceListResource) SetSchemaName(v string) *GetPackageResponseBodyDataResourceListResource {
	s.SchemaName = &v
	return s
}

type GetPackageResponseBodyDataResourceListTable struct {
	// The operations that were performed on the table.
	Actions []*string `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	// The name of the table.
	//
	// example:
	//
	// dim_odps
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The name of schema.
	//
	// example:
	//
	// default
	SchemaName *string `json:"schemaName,omitempty" xml:"schemaName,omitempty"`
}

func (s GetPackageResponseBodyDataResourceListTable) String() string {
	return tea.Prettify(s)
}

func (s GetPackageResponseBodyDataResourceListTable) GoString() string {
	return s.String()
}

func (s *GetPackageResponseBodyDataResourceListTable) SetActions(v []*string) *GetPackageResponseBodyDataResourceListTable {
	s.Actions = v
	return s
}

func (s *GetPackageResponseBodyDataResourceListTable) SetName(v string) *GetPackageResponseBodyDataResourceListTable {
	s.Name = &v
	return s
}

func (s *GetPackageResponseBodyDataResourceListTable) SetSchemaName(v string) *GetPackageResponseBodyDataResourceListTable {
	s.SchemaName = &v
	return s
}

type GetPackageResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPackageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPackageResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPackageResponse) GoString() string {
	return s.String()
}

func (s *GetPackageResponse) SetHeaders(v map[string]*string) *GetPackageResponse {
	s.Headers = v
	return s
}

func (s *GetPackageResponse) SetStatusCode(v int32) *GetPackageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPackageResponse) SetBody(v *GetPackageResponseBody) *GetPackageResponse {
	s.Body = v
	return s
}

type GetProjectRequest struct {
	// Specifies whether to use additional information.
	//
	// example:
	//
	// true
	Verbose *bool `json:"verbose,omitempty" xml:"verbose,omitempty"`
}

func (s GetProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s GetProjectRequest) GoString() string {
	return s.String()
}

func (s *GetProjectRequest) SetVerbose(v bool) *GetProjectRequest {
	s.Verbose = &v
	return s
}

type GetProjectResponseBody struct {
	// The data returned.
	Data *GetProjectResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// 040002
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// error message.
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// The HTTP status code.
	//
	// 	- 1xx: informational response. The request is received and is being processed.
	//
	// 	- 2xx: success. The request is successfully received, understood, and accepted by the server.
	//
	// 	- 3xx: redirection. The request is redirected, and further actions are required to complete the request.
	//
	// 	- 4xx: client error. The request contains invalid request parameters and syntaxes, or specific request conditions cannot be met.
	//
	// 	- 5xx: server error. The server cannot meet requirements due to other reasons.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0b87b7b316643495896551555e855b
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProjectResponseBody) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBody) SetData(v *GetProjectResponseBodyData) *GetProjectResponseBody {
	s.Data = v
	return s
}

func (s *GetProjectResponseBody) SetErrorCode(v string) *GetProjectResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetProjectResponseBody) SetErrorMsg(v string) *GetProjectResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetProjectResponseBody) SetHttpCode(v int32) *GetProjectResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetProjectResponseBody) SetRequestId(v string) *GetProjectResponseBody {
	s.RequestId = &v
	return s
}

type GetProjectResponseBodyData struct {
	// The project description.
	//
	// example:
	//
	// maxcompute project
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	// The total storage usage. The storage space that is occupied by your project, which is the logical storage space after your project data is collected and compressed.
	//
	// example:
	//
	// 16489027
	CostStorage *string `json:"costStorage,omitempty" xml:"costStorage,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 1704380838000
	CreatedTime *int64 `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// The default computing quota that is used to allocate computing resources. If you do not specify a computing quota for your project, the jobs that are initiated by your project consume the computing resources in the default quota. For more information about how to use computing resources, see [Use quota groups for computing resources](https://www.alibabacloud.com/help/zh/maxcompute/user-guide/use-of-computing-resources).
	//
	// example:
	//
	// quota_a
	DefaultQuota *string `json:"defaultQuota,omitempty" xml:"defaultQuota,omitempty"`
	// The information about the IP address whitelist.
	IpWhiteList *GetProjectResponseBodyDataIpWhiteList `json:"ipWhiteList,omitempty" xml:"ipWhiteList,omitempty" type:"Struct"`
	// The project name.
	//
	// example:
	//
	// odps_project
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The account information of the project owner.
	//
	// example:
	//
	// 1565950907343451
	Owner *string `json:"owner,omitempty" xml:"owner,omitempty"`
	// The billing method of the default computing quota.
	//
	// example:
	//
	// PayAsYouGo
	ProductType *string `json:"productType,omitempty" xml:"productType,omitempty"`
	// The basic properties of the project.
	Properties *GetProjectResponseBodyDataProperties `json:"properties,omitempty" xml:"properties,omitempty" type:"Struct"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The instance ID and billing method of the default computing quota.
	SaleTag *GetProjectResponseBodyDataSaleTag `json:"saleTag,omitempty" xml:"saleTag,omitempty" type:"Struct"`
	// The permission properties.
	SecurityProperties *GetProjectResponseBodyDataSecurityProperties `json:"securityProperties,omitempty" xml:"securityProperties,omitempty" type:"Struct"`
	// The project status. Valid values:
	//
	// 	- **AVAILABLE**
	//
	// 	- **READONLY**
	//
	// 	- **FROZEN**
	//
	// 	- **DELETING**
	//
	// example:
	//
	// AVAILABLE
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The list of `Super_Administrator` role members of the project.
	SuperAdmins []*string `json:"superAdmins,omitempty" xml:"superAdmins,omitempty" type:"Repeated"`
	// Indicates whether data storage by schema is supported. MaxCompute supports the schema feature. This feature allows you to classify objects such as tables, resources, and user-defined functions (UDFs) in a project by schema. You can create multiple schemas in a project. For more information, see [Schema-related operations](https://www.alibabacloud.com/help/zh/maxcompute/user-guide/schema-related-operations).
	//
	// example:
	//
	// true
	ThreeTierModel *bool `json:"threeTierModel,omitempty" xml:"threeTierModel,omitempty"`
	// The project type. Valid values:
	//
	// 	- **managed**: internal project
	//
	// 	- **external**: external project
	//
	// example:
	//
	// managed
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetProjectResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetProjectResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBodyData) SetComment(v string) *GetProjectResponseBodyData {
	s.Comment = &v
	return s
}

func (s *GetProjectResponseBodyData) SetCostStorage(v string) *GetProjectResponseBodyData {
	s.CostStorage = &v
	return s
}

func (s *GetProjectResponseBodyData) SetCreatedTime(v int64) *GetProjectResponseBodyData {
	s.CreatedTime = &v
	return s
}

func (s *GetProjectResponseBodyData) SetDefaultQuota(v string) *GetProjectResponseBodyData {
	s.DefaultQuota = &v
	return s
}

func (s *GetProjectResponseBodyData) SetIpWhiteList(v *GetProjectResponseBodyDataIpWhiteList) *GetProjectResponseBodyData {
	s.IpWhiteList = v
	return s
}

func (s *GetProjectResponseBodyData) SetName(v string) *GetProjectResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetProjectResponseBodyData) SetOwner(v string) *GetProjectResponseBodyData {
	s.Owner = &v
	return s
}

func (s *GetProjectResponseBodyData) SetProductType(v string) *GetProjectResponseBodyData {
	s.ProductType = &v
	return s
}

func (s *GetProjectResponseBodyData) SetProperties(v *GetProjectResponseBodyDataProperties) *GetProjectResponseBodyData {
	s.Properties = v
	return s
}

func (s *GetProjectResponseBodyData) SetRegionId(v string) *GetProjectResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *GetProjectResponseBodyData) SetSaleTag(v *GetProjectResponseBodyDataSaleTag) *GetProjectResponseBodyData {
	s.SaleTag = v
	return s
}

func (s *GetProjectResponseBodyData) SetSecurityProperties(v *GetProjectResponseBodyDataSecurityProperties) *GetProjectResponseBodyData {
	s.SecurityProperties = v
	return s
}

func (s *GetProjectResponseBodyData) SetStatus(v string) *GetProjectResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetProjectResponseBodyData) SetSuperAdmins(v []*string) *GetProjectResponseBodyData {
	s.SuperAdmins = v
	return s
}

func (s *GetProjectResponseBodyData) SetThreeTierModel(v bool) *GetProjectResponseBodyData {
	s.ThreeTierModel = &v
	return s
}

func (s *GetProjectResponseBodyData) SetType(v string) *GetProjectResponseBodyData {
	s.Type = &v
	return s
}

type GetProjectResponseBodyDataIpWhiteList struct {
	// The IP address whitelist for access over the Internet or the network for interconnecting with other Alibaba Cloud services.
	//
	// >  If you configure only the IP address whitelist for access over the Internet or the network for interconnecting with other Alibaba Cloud services, the access over the Internet or the network for interconnecting with other Alibaba Cloud services is subject to configurations, and access over a virtual private cloud (VPC) is not allowed.
	//
	// example:
	//
	// 10.88.111.3
	IpList *string `json:"ipList,omitempty" xml:"ipList,omitempty"`
	// The IP address whitelist for access over a VPC.
	//
	// >  If you configure only the IP address whitelist for access over a VPC, the access over a VPC is subject to configurations, and the access over the Internet or the network for interconnecting with other Alibaba Cloud services is not allowed.
	//
	// example:
	//
	// 10.88.111.3
	VpcIpList *string `json:"vpcIpList,omitempty" xml:"vpcIpList,omitempty"`
}

func (s GetProjectResponseBodyDataIpWhiteList) String() string {
	return tea.Prettify(s)
}

func (s GetProjectResponseBodyDataIpWhiteList) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBodyDataIpWhiteList) SetIpList(v string) *GetProjectResponseBodyDataIpWhiteList {
	s.IpList = &v
	return s
}

func (s *GetProjectResponseBodyDataIpWhiteList) SetVpcIpList(v string) *GetProjectResponseBodyDataIpWhiteList {
	s.VpcIpList = &v
	return s
}

type GetProjectResponseBodyDataProperties struct {
	// Indicates whether a full table scan is allowed in the project. A full table scan occupies a large number of resources, which reduces data processing efficiency. By default, the full table scan feature is disabled.
	//
	// example:
	//
	// false
	AllowFullScan *bool `json:"allowFullScan,omitempty" xml:"allowFullScan,omitempty"`
	// The Tunnel parent resource group that is bound to the project. You do not need to pay attention to this group.
	//
	// example:
	//
	// No value
	ElderTunnelQuota *string `json:"elderTunnelQuota,omitempty" xml:"elderTunnelQuota,omitempty"`
	// Indicates whether the DECIMAL type of the MaxCompute V2.0 data type edition is enabled.
	//
	// example:
	//
	// true
	EnableDecimal2 *bool `json:"enableDecimal2,omitempty" xml:"enableDecimal2,omitempty"`
	// Indicates whether external table caching is forcefully enabled.
	//
	// example:
	//
	// true
	EnableFdcCacheForce *bool `json:"enableFdcCacheForce,omitempty" xml:"enableFdcCacheForce,omitempty"`
	// Indicates whether [tiered storage](https://www.alibabacloud.com/help/zh/maxcompute/user-guide/tiered-storage) is enabled.
	//
	// example:
	//
	// true
	EnableTieredStorage *bool `json:"enableTieredStorage,omitempty" xml:"enableTieredStorage,omitempty"`
	// Indicates whether the routing of the Tunnel resource group is enabled.
	//
	// 	- true: The data transfer tasks that are submitted by the project by default use the Tunnel resource group that is bound to the project.
	//
	// 	- false: The data transfer tasks that are submitted by the project by default use the Tunnel shared resource group.
	//
	// example:
	//
	// true
	EnableTunnelQuotaRoute *bool `json:"enableTunnelQuotaRoute,omitempty" xml:"enableTunnelQuotaRoute,omitempty"`
	// The storage encryption properties.
	Encryption *GetProjectResponseBodyDataPropertiesEncryption `json:"encryption,omitempty" xml:"encryption,omitempty" type:"Struct"`
	// The properties of the external project.
	ExternalProjectProperties *GetProjectResponseBodyDataPropertiesExternalProjectProperties `json:"externalProjectProperties,omitempty" xml:"externalProjectProperties,omitempty" type:"Struct"`
	// The quota for external table caching.
	//
	// example:
	//
	// fdc_quota
	FdcQuota *string `json:"fdcQuota,omitempty" xml:"fdcQuota,omitempty"`
	// The retention period for backup data. Unit: days. During the retention period, you can restore data of the version in use to the backup data of any version. Valid values: [0,30]. Default value: 1. The value 0 indicates that the backup feature is disabled.
	//
	// example:
	//
	// 1
	RetentionDays *int64 `json:"retentionDays,omitempty" xml:"retentionDays,omitempty"`
	// The maximum consumption threshold of a single SQL statement. Formula: Amount of scanned data (GB) × Complexity.
	//
	// example:
	//
	// 1500
	SqlMeteringMax *string `json:"sqlMeteringMax,omitempty" xml:"sqlMeteringMax,omitempty"`
	// The [storage tier](https://www.alibabacloud.com/help/zh/maxcompute/user-guide/tiered-storage) information.
	StorageTierInfo *GetProjectResponseBodyDataPropertiesStorageTierInfo `json:"storageTierInfo,omitempty" xml:"storageTierInfo,omitempty" type:"Struct"`
	// The table lifecycle properties.
	TableLifecycle *GetProjectResponseBodyDataPropertiesTableLifecycle `json:"tableLifecycle,omitempty" xml:"tableLifecycle,omitempty" type:"Struct"`
	// The [properties of tiered storage lifecycle rules](https://www.alibabacloud.com/help/zh/maxcompute/user-guide/tiered-storage#f61fc9db76nna). After you configure the properties, the system triggers automatic switching of storage tiers based on the rules.
	TableLifecycleConfig *GetProjectResponseBodyDataPropertiesTableLifecycleConfig `json:"tableLifecycleConfig,omitempty" xml:"tableLifecycleConfig,omitempty" type:"Struct"`
	// The time zone that is used by your project. The time zone is the same as the time zone specified by `odps.sql.timezone`.
	//
	// example:
	//
	// Asia/Shanghai
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
	// The [Tunnel](https://www.alibabacloud.com/help/zh/maxcompute/user-guide/overview-of-dts) resource group that is bound to the project.
	//
	// 	- Default resource group: The Tunnel shared resource group is used. You cannot use the subscription-based Tunnel resource group for the project. The default resource group is automatically used by the Tunnel service of your project, regardless of the parameter setting.
	//
	// 	- Subscription-based Tunnel resource group: You can use the subscription-based Tunnel resource group for the project.
	//
	// example:
	//
	// Quota
	TunnelQuota *string `json:"tunnelQuota,omitempty" xml:"tunnelQuota,omitempty"`
	// The data type edition. Valid values:
	//
	// 	- **1**: MaxCompute V1.0 data type edition
	//
	// 	- **2**: MaxCompute V2.0 data type edition
	//
	// 	- **hive**: Hive-compatible data type edition
	//
	// For more information about the differences among the three data type editions, see [Data type editions](https://www.alibabacloud.com/help/zh/maxcompute/user-guide/data-type-editions).
	//
	// example:
	//
	// 2.0
	TypeSystem *string `json:"typeSystem,omitempty" xml:"typeSystem,omitempty"`
}

func (s GetProjectResponseBodyDataProperties) String() string {
	return tea.Prettify(s)
}

func (s GetProjectResponseBodyDataProperties) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBodyDataProperties) SetAllowFullScan(v bool) *GetProjectResponseBodyDataProperties {
	s.AllowFullScan = &v
	return s
}

func (s *GetProjectResponseBodyDataProperties) SetElderTunnelQuota(v string) *GetProjectResponseBodyDataProperties {
	s.ElderTunnelQuota = &v
	return s
}

func (s *GetProjectResponseBodyDataProperties) SetEnableDecimal2(v bool) *GetProjectResponseBodyDataProperties {
	s.EnableDecimal2 = &v
	return s
}

func (s *GetProjectResponseBodyDataProperties) SetEnableFdcCacheForce(v bool) *GetProjectResponseBodyDataProperties {
	s.EnableFdcCacheForce = &v
	return s
}

func (s *GetProjectResponseBodyDataProperties) SetEnableTieredStorage(v bool) *GetProjectResponseBodyDataProperties {
	s.EnableTieredStorage = &v
	return s
}

func (s *GetProjectResponseBodyDataProperties) SetEnableTunnelQuotaRoute(v bool) *GetProjectResponseBodyDataProperties {
	s.EnableTunnelQuotaRoute = &v
	return s
}

func (s *GetProjectResponseBodyDataProperties) SetEncryption(v *GetProjectResponseBodyDataPropertiesEncryption) *GetProjectResponseBodyDataProperties {
	s.Encryption = v
	return s
}

func (s *GetProjectResponseBodyDataProperties) SetExternalProjectProperties(v *GetProjectResponseBodyDataPropertiesExternalProjectProperties) *GetProjectResponseBodyDataProperties {
	s.ExternalProjectProperties = v
	return s
}

func (s *GetProjectResponseBodyDataProperties) SetFdcQuota(v string) *GetProjectResponseBodyDataProperties {
	s.FdcQuota = &v
	return s
}

func (s *GetProjectResponseBodyDataProperties) SetRetentionDays(v int64) *GetProjectResponseBodyDataProperties {
	s.RetentionDays = &v
	return s
}

func (s *GetProjectResponseBodyDataProperties) SetSqlMeteringMax(v string) *GetProjectResponseBodyDataProperties {
	s.SqlMeteringMax = &v
	return s
}

func (s *GetProjectResponseBodyDataProperties) SetStorageTierInfo(v *GetProjectResponseBodyDataPropertiesStorageTierInfo) *GetProjectResponseBodyDataProperties {
	s.StorageTierInfo = v
	return s
}

func (s *GetProjectResponseBodyDataProperties) SetTableLifecycle(v *GetProjectResponseBodyDataPropertiesTableLifecycle) *GetProjectResponseBodyDataProperties {
	s.TableLifecycle = v
	return s
}

func (s *GetProjectResponseBodyDataProperties) SetTableLifecycleConfig(v *GetProjectResponseBodyDataPropertiesTableLifecycleConfig) *GetProjectResponseBodyDataProperties {
	s.TableLifecycleConfig = v
	return s
}

func (s *GetProjectResponseBodyDataProperties) SetTimezone(v string) *GetProjectResponseBodyDataProperties {
	s.Timezone = &v
	return s
}

func (s *GetProjectResponseBodyDataProperties) SetTunnelQuota(v string) *GetProjectResponseBodyDataProperties {
	s.TunnelQuota = &v
	return s
}

func (s *GetProjectResponseBodyDataProperties) SetTypeSystem(v string) *GetProjectResponseBodyDataProperties {
	s.TypeSystem = &v
	return s
}

type GetProjectResponseBodyDataPropertiesEncryption struct {
	// The data encryption algorithm that is supported by the key. Valid values: AES256, AESCTR, and RC4.
	//
	// example:
	//
	// SHA1
	Algorithm *string `json:"algorithm,omitempty" xml:"algorithm,omitempty"`
	// Indicates whether the data encryption feature needs to be enabled for the project. For more information about data encryption, see
	//
	// [Storage encryption](https://www.alibabacloud.com/help/zh/maxcompute/security-and-compliance/storage-encryption).
	//
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// The type of key that is used for data encryption. You can select MaxCompute Default Key or Bring Your Own Key (BYOK) as the key type. If you select MaxCompute Default Key, the default key that is created by MaxCompute is used.
	//
	// example:
	//
	// dafault
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
}

func (s GetProjectResponseBodyDataPropertiesEncryption) String() string {
	return tea.Prettify(s)
}

func (s GetProjectResponseBodyDataPropertiesEncryption) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBodyDataPropertiesEncryption) SetAlgorithm(v string) *GetProjectResponseBodyDataPropertiesEncryption {
	s.Algorithm = &v
	return s
}

func (s *GetProjectResponseBodyDataPropertiesEncryption) SetEnable(v bool) *GetProjectResponseBodyDataPropertiesEncryption {
	s.Enable = &v
	return s
}

func (s *GetProjectResponseBodyDataPropertiesEncryption) SetKey(v string) *GetProjectResponseBodyDataPropertiesEncryption {
	s.Key = &v
	return s
}

type GetProjectResponseBodyDataPropertiesExternalProjectProperties struct {
	// Indicates whether the external project is an external project for [data lakehouse solution 2.0](https://www.alibabacloud.com/help/zh/maxcompute/user-guide/lake-warehouse-integrated-2-0-use-guide).
	//
	// example:
	//
	// true
	IsExternalCatalogBound *string `json:"isExternalCatalogBound,omitempty" xml:"isExternalCatalogBound,omitempty"`
}

func (s GetProjectResponseBodyDataPropertiesExternalProjectProperties) String() string {
	return tea.Prettify(s)
}

func (s GetProjectResponseBodyDataPropertiesExternalProjectProperties) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBodyDataPropertiesExternalProjectProperties) SetIsExternalCatalogBound(v string) *GetProjectResponseBodyDataPropertiesExternalProjectProperties {
	s.IsExternalCatalogBound = &v
	return s
}

type GetProjectResponseBodyDataPropertiesStorageTierInfo struct {
	// The backup storage usage.
	//
	// example:
	//
	// 86672917
	ProjectBackupSize *int64 `json:"projectBackupSize,omitempty" xml:"projectBackupSize,omitempty"`
	// The total storage usage.
	//
	// example:
	//
	// 56066037
	ProjectTotalSize *int64 `json:"projectTotalSize,omitempty" xml:"projectTotalSize,omitempty"`
	// The [storage tier](https://www.alibabacloud.com/help/zh/maxcompute/user-guide/tiered-storage) information.
	StorageTierSize *GetProjectResponseBodyDataPropertiesStorageTierInfoStorageTierSize `json:"storageTierSize,omitempty" xml:"storageTierSize,omitempty" type:"Struct"`
}

func (s GetProjectResponseBodyDataPropertiesStorageTierInfo) String() string {
	return tea.Prettify(s)
}

func (s GetProjectResponseBodyDataPropertiesStorageTierInfo) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBodyDataPropertiesStorageTierInfo) SetProjectBackupSize(v int64) *GetProjectResponseBodyDataPropertiesStorageTierInfo {
	s.ProjectBackupSize = &v
	return s
}

func (s *GetProjectResponseBodyDataPropertiesStorageTierInfo) SetProjectTotalSize(v int64) *GetProjectResponseBodyDataPropertiesStorageTierInfo {
	s.ProjectTotalSize = &v
	return s
}

func (s *GetProjectResponseBodyDataPropertiesStorageTierInfo) SetStorageTierSize(v *GetProjectResponseBodyDataPropertiesStorageTierInfoStorageTierSize) *GetProjectResponseBodyDataPropertiesStorageTierInfo {
	s.StorageTierSize = v
	return s
}

type GetProjectResponseBodyDataPropertiesStorageTierInfoStorageTierSize struct {
	// The storage usage at the long-term storage tier.
	//
	// example:
	//
	// 21764917
	LongTermSize *int64 `json:"longTermSize,omitempty" xml:"longTermSize,omitempty"`
	// The storage usage at the Infrequent Access (IA) layer.
	//
	// example:
	//
	// 767693
	LowFrequencySize *int64 `json:"lowFrequencySize,omitempty" xml:"lowFrequencySize,omitempty"`
	// The storage usage at the standard storage tier.
	//
	// example:
	//
	// 27649172
	StandardSize *int64 `json:"standardSize,omitempty" xml:"standardSize,omitempty"`
}

func (s GetProjectResponseBodyDataPropertiesStorageTierInfoStorageTierSize) String() string {
	return tea.Prettify(s)
}

func (s GetProjectResponseBodyDataPropertiesStorageTierInfoStorageTierSize) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBodyDataPropertiesStorageTierInfoStorageTierSize) SetLongTermSize(v int64) *GetProjectResponseBodyDataPropertiesStorageTierInfoStorageTierSize {
	s.LongTermSize = &v
	return s
}

func (s *GetProjectResponseBodyDataPropertiesStorageTierInfoStorageTierSize) SetLowFrequencySize(v int64) *GetProjectResponseBodyDataPropertiesStorageTierInfoStorageTierSize {
	s.LowFrequencySize = &v
	return s
}

func (s *GetProjectResponseBodyDataPropertiesStorageTierInfoStorageTierSize) SetStandardSize(v int64) *GetProjectResponseBodyDataPropertiesStorageTierInfoStorageTierSize {
	s.StandardSize = &v
	return s
}

type GetProjectResponseBodyDataPropertiesTableLifecycle struct {
	// The lifecycle type. Valid values:
	//
	// 	- **mandatory**: The lifecycle clause is required in a table creation statement.
	//
	// 	- **optional**: The lifecycle clause is optional in a table creation statement. If you do not configure a lifecycle for a table, the table does not expire.
	//
	// 	- **inherit**: If you do not configure a lifecycle for a table when you create the table, the value of the odps.table.lifecycle.value parameter is used as the table lifecycle by default.
	//
	// example:
	//
	// optional
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The table lifecycle. Unit: days. Valid values: 1 to 37231. Default value: 37231.
	//
	// example:
	//
	// 37231
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetProjectResponseBodyDataPropertiesTableLifecycle) String() string {
	return tea.Prettify(s)
}

func (s GetProjectResponseBodyDataPropertiesTableLifecycle) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBodyDataPropertiesTableLifecycle) SetType(v string) *GetProjectResponseBodyDataPropertiesTableLifecycle {
	s.Type = &v
	return s
}

func (s *GetProjectResponseBodyDataPropertiesTableLifecycle) SetValue(v string) *GetProjectResponseBodyDataPropertiesTableLifecycle {
	s.Value = &v
	return s
}

type GetProjectResponseBodyDataPropertiesTableLifecycleConfig struct {
	// The information about the long-term storage tier.
	TierToLongterm *GetProjectResponseBodyDataPropertiesTableLifecycleConfigTierToLongterm `json:"TierToLongterm,omitempty" xml:"TierToLongterm,omitempty" type:"Struct"`
	// The information about the IA storage tier.
	TierToLowFrequency *GetProjectResponseBodyDataPropertiesTableLifecycleConfigTierToLowFrequency `json:"TierToLowFrequency,omitempty" xml:"TierToLowFrequency,omitempty" type:"Struct"`
}

func (s GetProjectResponseBodyDataPropertiesTableLifecycleConfig) String() string {
	return tea.Prettify(s)
}

func (s GetProjectResponseBodyDataPropertiesTableLifecycleConfig) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBodyDataPropertiesTableLifecycleConfig) SetTierToLongterm(v *GetProjectResponseBodyDataPropertiesTableLifecycleConfigTierToLongterm) *GetProjectResponseBodyDataPropertiesTableLifecycleConfig {
	s.TierToLongterm = v
	return s
}

func (s *GetProjectResponseBodyDataPropertiesTableLifecycleConfig) SetTierToLowFrequency(v *GetProjectResponseBodyDataPropertiesTableLifecycleConfigTierToLowFrequency) *GetProjectResponseBodyDataPropertiesTableLifecycleConfig {
	s.TierToLowFrequency = v
	return s
}

type GetProjectResponseBodyDataPropertiesTableLifecycleConfigTierToLongterm struct {
	// The system triggers an automatic storage tier change N days after the last access time of data. N is specified by this parameter and corresponds to `LastAccessTime` that is configured for the table or partition.
	//
	// >  If LastAccessTime of a table or partition is left empty, the following rules are applied:
	//
	// 	- For tables or partitions that you created before October 1, 2023, 2023.10.01 00:00:00 in UTC+0 is considered as the last access time.
	//
	// 	- For tables or partitions that you created on or after October 1, 2023, if no data is accessed, the table or partition creation time is considered as the last access time.
	//
	// example:
	//
	// 180
	DaysAfterLastAccessGreaterThan *int64 `json:"DaysAfterLastAccessGreaterThan,omitempty" xml:"DaysAfterLastAccessGreaterThan,omitempty"`
	// The system triggers an automatic storage tier change N days after the last modification time of data. N is specified by this parameter and corresponds to `LastModifiedTime` that is configured for the table or partition.
	//
	// example:
	//
	// 180
	DaysAfterLastModificationGreaterThan *int64 `json:"DaysAfterLastModificationGreaterThan,omitempty" xml:"DaysAfterLastModificationGreaterThan,omitempty"`
	// The period after the previous storage tier change time.
	//
	// example:
	//
	// 1
	DaysAfterLastTierModificationGreaterThan *int64 `json:"DaysAfterLastTierModificationGreaterThan,omitempty" xml:"DaysAfterLastTierModificationGreaterThan,omitempty"`
}

func (s GetProjectResponseBodyDataPropertiesTableLifecycleConfigTierToLongterm) String() string {
	return tea.Prettify(s)
}

func (s GetProjectResponseBodyDataPropertiesTableLifecycleConfigTierToLongterm) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBodyDataPropertiesTableLifecycleConfigTierToLongterm) SetDaysAfterLastAccessGreaterThan(v int64) *GetProjectResponseBodyDataPropertiesTableLifecycleConfigTierToLongterm {
	s.DaysAfterLastAccessGreaterThan = &v
	return s
}

func (s *GetProjectResponseBodyDataPropertiesTableLifecycleConfigTierToLongterm) SetDaysAfterLastModificationGreaterThan(v int64) *GetProjectResponseBodyDataPropertiesTableLifecycleConfigTierToLongterm {
	s.DaysAfterLastModificationGreaterThan = &v
	return s
}

func (s *GetProjectResponseBodyDataPropertiesTableLifecycleConfigTierToLongterm) SetDaysAfterLastTierModificationGreaterThan(v int64) *GetProjectResponseBodyDataPropertiesTableLifecycleConfigTierToLongterm {
	s.DaysAfterLastTierModificationGreaterThan = &v
	return s
}

type GetProjectResponseBodyDataPropertiesTableLifecycleConfigTierToLowFrequency struct {
	// The system triggers an automatic storage tier change N days after the last access time of data. N is specified by this parameter and corresponds to `LastAccessTime` that is configured for the table or partition.
	//
	// >  If LastAccessTime of a table or partition is left empty, the following rules are applied:
	//
	// 	- For tables or partitions that you created before October 1, 2023, 2023.10.01 00:00:00 in UTC+0 is considered as the last access time.
	//
	// 	- For tables or partitions that you created on or after October 1, 2023, if no data is accessed, the table or partition creation time is considered as the last access time.
	//
	// example:
	//
	// 30
	DaysAfterLastAccessGreaterThan *int64 `json:"DaysAfterLastAccessGreaterThan,omitempty" xml:"DaysAfterLastAccessGreaterThan,omitempty"`
	// The system triggers an automatic storage tier change N days after the last modification time of data. N is specified by this parameter and corresponds to `LastModifiedTime` that is configured for the table or partition.
	//
	// example:
	//
	// 30
	DaysAfterLastModificationGreaterThan *int64 `json:"DaysAfterLastModificationGreaterThan,omitempty" xml:"DaysAfterLastModificationGreaterThan,omitempty"`
	// The period after the previous storage tier change time.
	//
	// example:
	//
	// 1
	DaysAfterLastTierModificationGreaterThan *int64 `json:"DaysAfterLastTierModificationGreaterThan,omitempty" xml:"DaysAfterLastTierModificationGreaterThan,omitempty"`
}

func (s GetProjectResponseBodyDataPropertiesTableLifecycleConfigTierToLowFrequency) String() string {
	return tea.Prettify(s)
}

func (s GetProjectResponseBodyDataPropertiesTableLifecycleConfigTierToLowFrequency) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBodyDataPropertiesTableLifecycleConfigTierToLowFrequency) SetDaysAfterLastAccessGreaterThan(v int64) *GetProjectResponseBodyDataPropertiesTableLifecycleConfigTierToLowFrequency {
	s.DaysAfterLastAccessGreaterThan = &v
	return s
}

func (s *GetProjectResponseBodyDataPropertiesTableLifecycleConfigTierToLowFrequency) SetDaysAfterLastModificationGreaterThan(v int64) *GetProjectResponseBodyDataPropertiesTableLifecycleConfigTierToLowFrequency {
	s.DaysAfterLastModificationGreaterThan = &v
	return s
}

func (s *GetProjectResponseBodyDataPropertiesTableLifecycleConfigTierToLowFrequency) SetDaysAfterLastTierModificationGreaterThan(v int64) *GetProjectResponseBodyDataPropertiesTableLifecycleConfigTierToLowFrequency {
	s.DaysAfterLastTierModificationGreaterThan = &v
	return s
}

type GetProjectResponseBodyDataSaleTag struct {
	// The instance ID of the default computing quota.
	//
	// example:
	//
	// project_name
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// The billing method of the default computing quota.
	//
	// example:
	//
	// project
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s GetProjectResponseBodyDataSaleTag) String() string {
	return tea.Prettify(s)
}

func (s GetProjectResponseBodyDataSaleTag) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBodyDataSaleTag) SetResourceId(v string) *GetProjectResponseBodyDataSaleTag {
	s.ResourceId = &v
	return s
}

func (s *GetProjectResponseBodyDataSaleTag) SetResourceType(v string) *GetProjectResponseBodyDataSaleTag {
	s.ResourceType = &v
	return s
}

type GetProjectResponseBodyDataSecurityProperties struct {
	// Indicates whether the [download control](https://www.alibabacloud.com/help/zh/maxcompute/user-guide/label-based-access-control) feature is enabled. By default, this feature is disabled.
	//
	// example:
	//
	// false
	EnableDownloadPrivilege *bool `json:"enableDownloadPrivilege,omitempty" xml:"enableDownloadPrivilege,omitempty"`
	// Indicates whether the [label-based access control](https://www.alibabacloud.com/help/zh/maxcompute/user-guide/label-based-access-control) feature is enabled. By default, this feature is disabled.
	//
	// example:
	//
	// false
	LabelSecurity *bool `json:"labelSecurity,omitempty" xml:"labelSecurity,omitempty"`
	// Indicates whether to allow the object creator to have the access permissions on the object. The default value is true, which indicates that the object creator has the access permissions on the object.
	//
	// example:
	//
	// true
	ObjectCreatorHasAccessPermission *bool `json:"objectCreatorHasAccessPermission,omitempty" xml:"objectCreatorHasAccessPermission,omitempty"`
	// Indicates whether the object creator has the authorization permissions on the object. The default value is true, which indicates that the object creator has the authorization permissions on the object.
	//
	// example:
	//
	// true
	ObjectCreatorHasGrantPermission *bool `json:"objectCreatorHasGrantPermission,omitempty" xml:"objectCreatorHasGrantPermission,omitempty"`
	// The properties of the [data protection mechanism](https://www.alibabacloud.com/help/zh/maxcompute/security-and-compliance/project-data-protection).
	ProjectProtection *GetProjectResponseBodyDataSecurityPropertiesProjectProtection `json:"projectProtection,omitempty" xml:"projectProtection,omitempty" type:"Struct"`
	// Indicates whether the [ACL-based access control](https://www.alibabacloud.com/help/zh/maxcompute/user-guide/acl-based-access-control) feature is enabled. By default, this feature is enabled.
	//
	// example:
	//
	// true
	UsingAcl *bool `json:"usingAcl,omitempty" xml:"usingAcl,omitempty"`
	// Indicates whether the [policy-based access control](https://www.alibabacloud.com/help/zh/maxcompute/user-guide/policy-based-access-control-1) feature is enabled. By default, this feature is enabled.
	//
	// example:
	//
	// true
	UsingPolicy *bool `json:"usingPolicy,omitempty" xml:"usingPolicy,omitempty"`
}

func (s GetProjectResponseBodyDataSecurityProperties) String() string {
	return tea.Prettify(s)
}

func (s GetProjectResponseBodyDataSecurityProperties) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBodyDataSecurityProperties) SetEnableDownloadPrivilege(v bool) *GetProjectResponseBodyDataSecurityProperties {
	s.EnableDownloadPrivilege = &v
	return s
}

func (s *GetProjectResponseBodyDataSecurityProperties) SetLabelSecurity(v bool) *GetProjectResponseBodyDataSecurityProperties {
	s.LabelSecurity = &v
	return s
}

func (s *GetProjectResponseBodyDataSecurityProperties) SetObjectCreatorHasAccessPermission(v bool) *GetProjectResponseBodyDataSecurityProperties {
	s.ObjectCreatorHasAccessPermission = &v
	return s
}

func (s *GetProjectResponseBodyDataSecurityProperties) SetObjectCreatorHasGrantPermission(v bool) *GetProjectResponseBodyDataSecurityProperties {
	s.ObjectCreatorHasGrantPermission = &v
	return s
}

func (s *GetProjectResponseBodyDataSecurityProperties) SetProjectProtection(v *GetProjectResponseBodyDataSecurityPropertiesProjectProtection) *GetProjectResponseBodyDataSecurityProperties {
	s.ProjectProtection = v
	return s
}

func (s *GetProjectResponseBodyDataSecurityProperties) SetUsingAcl(v bool) *GetProjectResponseBodyDataSecurityProperties {
	s.UsingAcl = &v
	return s
}

func (s *GetProjectResponseBodyDataSecurityProperties) SetUsingPolicy(v bool) *GetProjectResponseBodyDataSecurityProperties {
	s.UsingPolicy = &v
	return s
}

type GetProjectResponseBodyDataSecurityPropertiesProjectProtection struct {
	// If you enable the project data protection mechanism, you can configure exception or trusted projects. This allows specified users to transfer data of a specified object to a specified project. The project data protection mechanism does not take effect in all the situations that are specified in the exception policy.
	//
	// example:
	//
	// {
	//
	//     "Version": "1",
	//
	//     "Statement":
	//
	//     [{
	//
	//         "Effect":"Allow",
	//
	//         "Principal":"<Principal>",
	//
	//         "Action":["odps:<Action1>[, <Action2>, ...]"],
	//
	//         "Resource":"acs:odps:*:<Resource>",
	//
	//         "Condition":{
	//
	//             "StringEquals": {
	//
	//                 "odps:TaskType":["<Tasktype>"]
	//
	//             }
	//
	//         }
	//
	//     }]
	//
	//     }
	ExceptionPolicy *string `json:"exceptionPolicy,omitempty" xml:"exceptionPolicy,omitempty"`
	// Indicates whether the [data protection mechanism](https://www.alibabacloud.com/help/zh/maxcompute/security-and-compliance/project-data-protection) is enabled for the project. This allows or denies data transfer across projects. By default, the data protection mechanism is disabled.
	//
	// example:
	//
	// true
	Protected *bool `json:"protected,omitempty" xml:"protected,omitempty"`
}

func (s GetProjectResponseBodyDataSecurityPropertiesProjectProtection) String() string {
	return tea.Prettify(s)
}

func (s GetProjectResponseBodyDataSecurityPropertiesProjectProtection) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBodyDataSecurityPropertiesProjectProtection) SetExceptionPolicy(v string) *GetProjectResponseBodyDataSecurityPropertiesProjectProtection {
	s.ExceptionPolicy = &v
	return s
}

func (s *GetProjectResponseBodyDataSecurityPropertiesProjectProtection) SetProtected(v bool) *GetProjectResponseBodyDataSecurityPropertiesProjectProtection {
	s.Protected = &v
	return s
}

type GetProjectResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProjectResponse) GoString() string {
	return s.String()
}

func (s *GetProjectResponse) SetHeaders(v map[string]*string) *GetProjectResponse {
	s.Headers = v
	return s
}

func (s *GetProjectResponse) SetStatusCode(v int32) *GetProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProjectResponse) SetBody(v *GetProjectResponseBody) *GetProjectResponse {
	s.Body = v
	return s
}

type GetQuotaRequest struct {
	// The trusted AccessKey pairs.
	//
	// example:
	//
	// null
	AkProven *string `json:"AkProven,omitempty" xml:"AkProven,omitempty"`
	// Specifies whether to include submodules. Valid values: -true: The request includes submodules. -false: The request does not include submodules. This is the default value.
	//
	// example:
	//
	// false
	Mock *bool `json:"mock,omitempty" xml:"mock,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-chengdu
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 520539530998273
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetQuotaRequest) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaRequest) GoString() string {
	return s.String()
}

func (s *GetQuotaRequest) SetAkProven(v string) *GetQuotaRequest {
	s.AkProven = &v
	return s
}

func (s *GetQuotaRequest) SetMock(v bool) *GetQuotaRequest {
	s.Mock = &v
	return s
}

func (s *GetQuotaRequest) SetRegion(v string) *GetQuotaRequest {
	s.Region = &v
	return s
}

func (s *GetQuotaRequest) SetTenantId(v string) *GetQuotaRequest {
	s.TenantId = &v
	return s
}

type GetQuotaResponseBody struct {
	// The information about the order.
	BillingPolicy *GetQuotaResponseBodyBillingPolicy `json:"billingPolicy,omitempty" xml:"billingPolicy,omitempty" type:"Struct"`
	// The cluster ID.
	//
	// example:
	//
	// AT-120N
	Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// The time when the resource was created.
	//
	// example:
	//
	// 2022-09-06T02:14:44Z
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The ID of the Alibaba Cloud account that is used to create the resource.
	//
	// example:
	//
	// 672863518
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// The returned data.
	Data *GetQuotaResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The quota ID.
	//
	// example:
	//
	// 0
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the quota.
	//
	// example:
	//
	// quota_a
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The alias of the quota.
	//
	// example:
	//
	// quota_nickname
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	// The description of the quota.
	//
	// example:
	//
	// {"minCU":10,
	//
	// "adhocCU":0,
	//
	// "maxCU":10,
	//
	// "schedulerType":"Fair",
	//
	// }
	Parameter map[string]interface{} `json:"parameter,omitempty" xml:"parameter,omitempty"`
	// The ID of the parent resource.
	//
	// example:
	//
	// null
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0b87b7a316654730544735643e9200
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The identifier of an object in a MaxCompute quota. This identifier is the same as the identifier in the sales bill of Alibaba Cloud. This parameter is used for tags.
	SaleTag *GetQuotaResponseBodySaleTag `json:"saleTag,omitempty" xml:"saleTag,omitempty" type:"Struct"`
	// The information about the scheduling plan.
	ScheduleInfo *GetQuotaResponseBodyScheduleInfo `json:"scheduleInfo,omitempty" xml:"scheduleInfo,omitempty" type:"Struct"`
	// The status of the resource.
	//
	// example:
	//
	// ON
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The information about the level-2 quota.
	SubQuotaInfoList []*GetQuotaResponseBodySubQuotaInfoList `json:"subQuotaInfoList,omitempty" xml:"subQuotaInfoList,omitempty" type:"Repeated"`
	// The tag of the resource for the quota.
	//
	// example:
	//
	// abc
	Tag *string `json:"tag,omitempty" xml:"tag,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 478403690625249
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The type of the resource system. This parameter corresponds to the resourceSystemType parameter of the cluster.
	//
	// example:
	//
	// FUXI_ONLINE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The version number.
	//
	// example:
	//
	// 1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetQuotaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *GetQuotaResponseBody) SetBillingPolicy(v *GetQuotaResponseBodyBillingPolicy) *GetQuotaResponseBody {
	s.BillingPolicy = v
	return s
}

func (s *GetQuotaResponseBody) SetCluster(v string) *GetQuotaResponseBody {
	s.Cluster = &v
	return s
}

func (s *GetQuotaResponseBody) SetCreateTime(v int64) *GetQuotaResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetQuotaResponseBody) SetCreatorId(v string) *GetQuotaResponseBody {
	s.CreatorId = &v
	return s
}

func (s *GetQuotaResponseBody) SetData(v *GetQuotaResponseBodyData) *GetQuotaResponseBody {
	s.Data = v
	return s
}

func (s *GetQuotaResponseBody) SetId(v string) *GetQuotaResponseBody {
	s.Id = &v
	return s
}

func (s *GetQuotaResponseBody) SetName(v string) *GetQuotaResponseBody {
	s.Name = &v
	return s
}

func (s *GetQuotaResponseBody) SetNickName(v string) *GetQuotaResponseBody {
	s.NickName = &v
	return s
}

func (s *GetQuotaResponseBody) SetParameter(v map[string]interface{}) *GetQuotaResponseBody {
	s.Parameter = v
	return s
}

func (s *GetQuotaResponseBody) SetParentId(v string) *GetQuotaResponseBody {
	s.ParentId = &v
	return s
}

func (s *GetQuotaResponseBody) SetRegionId(v string) *GetQuotaResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetQuotaResponseBody) SetRequestId(v string) *GetQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQuotaResponseBody) SetSaleTag(v *GetQuotaResponseBodySaleTag) *GetQuotaResponseBody {
	s.SaleTag = v
	return s
}

func (s *GetQuotaResponseBody) SetScheduleInfo(v *GetQuotaResponseBodyScheduleInfo) *GetQuotaResponseBody {
	s.ScheduleInfo = v
	return s
}

func (s *GetQuotaResponseBody) SetStatus(v string) *GetQuotaResponseBody {
	s.Status = &v
	return s
}

func (s *GetQuotaResponseBody) SetSubQuotaInfoList(v []*GetQuotaResponseBodySubQuotaInfoList) *GetQuotaResponseBody {
	s.SubQuotaInfoList = v
	return s
}

func (s *GetQuotaResponseBody) SetTag(v string) *GetQuotaResponseBody {
	s.Tag = &v
	return s
}

func (s *GetQuotaResponseBody) SetTenantId(v string) *GetQuotaResponseBody {
	s.TenantId = &v
	return s
}

func (s *GetQuotaResponseBody) SetType(v string) *GetQuotaResponseBody {
	s.Type = &v
	return s
}

func (s *GetQuotaResponseBody) SetVersion(v string) *GetQuotaResponseBody {
	s.Version = &v
	return s
}

type GetQuotaResponseBodyBillingPolicy struct {
	// The billing method of the quota. Valid values:
	//
	// 	- subscription: a subscription quota.
	//
	// 	- payasyougo: a pay-as-you-go quota.
	//
	// example:
	//
	// subscription
	BillingMethod *string `json:"billingMethod,omitempty" xml:"billingMethod,omitempty"`
	// The specifications of the order.
	//
	// example:
	//
	// OdpsStandard
	OdpsSpecCode *string `json:"odpsSpecCode,omitempty" xml:"odpsSpecCode,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 45245678
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
}

func (s GetQuotaResponseBodyBillingPolicy) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaResponseBodyBillingPolicy) GoString() string {
	return s.String()
}

func (s *GetQuotaResponseBodyBillingPolicy) SetBillingMethod(v string) *GetQuotaResponseBodyBillingPolicy {
	s.BillingMethod = &v
	return s
}

func (s *GetQuotaResponseBodyBillingPolicy) SetOdpsSpecCode(v string) *GetQuotaResponseBodyBillingPolicy {
	s.OdpsSpecCode = &v
	return s
}

func (s *GetQuotaResponseBodyBillingPolicy) SetOrderId(v string) *GetQuotaResponseBodyBillingPolicy {
	s.OrderId = &v
	return s
}

type GetQuotaResponseBodyData struct {
	// The information about the order.
	BillingPolicy *GetQuotaResponseBodyDataBillingPolicy `json:"billingPolicy,omitempty" xml:"billingPolicy,omitempty" type:"Struct"`
	// The cluster ID.
	//
	// example:
	//
	// AT-120N
	Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// The time when the resource was created.
	//
	// example:
	//
	// 2022-09-06T02:14:44Z
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The ID of the Alibaba Cloud account that is used to create the resource.
	//
	// example:
	//
	// 672863518
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// The quota ID.
	//
	// example:
	//
	// 0
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the quota.
	//
	// example:
	//
	// quota_a
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The alias of the quota.
	//
	// example:
	//
	// quota_nickname
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	// The description of the quota.
	//
	// example:
	//
	// {"minCU":10,
	//
	// "adhocCU":0,
	//
	// "maxCU":10,
	//
	// "schedulerType":"Fair",
	//
	// }
	Parameter map[string]interface{} `json:"parameter,omitempty" xml:"parameter,omitempty"`
	// The ID of the parent resource.
	//
	// example:
	//
	// null
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The identifier of an object in a MaxCompute quota. This identifier is the same as the identifier in the sales bill of Alibaba Cloud. This parameter is used for tags.
	SaleTag *GetQuotaResponseBodyDataSaleTag `json:"saleTag,omitempty" xml:"saleTag,omitempty" type:"Struct"`
	// The information about the scheduling plan.
	ScheduleInfo *GetQuotaResponseBodyDataScheduleInfo `json:"scheduleInfo,omitempty" xml:"scheduleInfo,omitempty" type:"Struct"`
	// The status of the resource.
	//
	// example:
	//
	// ON
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The information about the level-2 quota.
	SubQuotaInfoList []*GetQuotaResponseBodyDataSubQuotaInfoList `json:"subQuotaInfoList,omitempty" xml:"subQuotaInfoList,omitempty" type:"Repeated"`
	// The tag of the resource for the quota.
	//
	// example:
	//
	// abc
	Tag *string `json:"tag,omitempty" xml:"tag,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 478403690625249
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The type of the resource system. This parameter corresponds to the resourceSystemType parameter of the cluster.
	//
	// example:
	//
	// FUXI_ONLINE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The version number.
	//
	// example:
	//
	// 1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetQuotaResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetQuotaResponseBodyData) SetBillingPolicy(v *GetQuotaResponseBodyDataBillingPolicy) *GetQuotaResponseBodyData {
	s.BillingPolicy = v
	return s
}

func (s *GetQuotaResponseBodyData) SetCluster(v string) *GetQuotaResponseBodyData {
	s.Cluster = &v
	return s
}

func (s *GetQuotaResponseBodyData) SetCreateTime(v int64) *GetQuotaResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetQuotaResponseBodyData) SetCreatorId(v string) *GetQuotaResponseBodyData {
	s.CreatorId = &v
	return s
}

func (s *GetQuotaResponseBodyData) SetId(v string) *GetQuotaResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetQuotaResponseBodyData) SetName(v string) *GetQuotaResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetQuotaResponseBodyData) SetNickName(v string) *GetQuotaResponseBodyData {
	s.NickName = &v
	return s
}

func (s *GetQuotaResponseBodyData) SetParameter(v map[string]interface{}) *GetQuotaResponseBodyData {
	s.Parameter = v
	return s
}

func (s *GetQuotaResponseBodyData) SetParentId(v string) *GetQuotaResponseBodyData {
	s.ParentId = &v
	return s
}

func (s *GetQuotaResponseBodyData) SetRegionId(v string) *GetQuotaResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *GetQuotaResponseBodyData) SetSaleTag(v *GetQuotaResponseBodyDataSaleTag) *GetQuotaResponseBodyData {
	s.SaleTag = v
	return s
}

func (s *GetQuotaResponseBodyData) SetScheduleInfo(v *GetQuotaResponseBodyDataScheduleInfo) *GetQuotaResponseBodyData {
	s.ScheduleInfo = v
	return s
}

func (s *GetQuotaResponseBodyData) SetStatus(v string) *GetQuotaResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetQuotaResponseBodyData) SetSubQuotaInfoList(v []*GetQuotaResponseBodyDataSubQuotaInfoList) *GetQuotaResponseBodyData {
	s.SubQuotaInfoList = v
	return s
}

func (s *GetQuotaResponseBodyData) SetTag(v string) *GetQuotaResponseBodyData {
	s.Tag = &v
	return s
}

func (s *GetQuotaResponseBodyData) SetTenantId(v string) *GetQuotaResponseBodyData {
	s.TenantId = &v
	return s
}

func (s *GetQuotaResponseBodyData) SetType(v string) *GetQuotaResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetQuotaResponseBodyData) SetVersion(v string) *GetQuotaResponseBodyData {
	s.Version = &v
	return s
}

type GetQuotaResponseBodyDataBillingPolicy struct {
	// The billing method of the quota. Valid values:
	//
	// 	- subscription: a subscription quota.
	//
	// 	- payasyougo: a pay-as-you-go quota.
	//
	// example:
	//
	// subscription
	BillingMethod *string `json:"billingMethod,omitempty" xml:"billingMethod,omitempty"`
	// The specifications of the order.
	//
	// example:
	//
	// OdpsStandard
	OdpsSpecCode *string `json:"odpsSpecCode,omitempty" xml:"odpsSpecCode,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 45245678
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
}

func (s GetQuotaResponseBodyDataBillingPolicy) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaResponseBodyDataBillingPolicy) GoString() string {
	return s.String()
}

func (s *GetQuotaResponseBodyDataBillingPolicy) SetBillingMethod(v string) *GetQuotaResponseBodyDataBillingPolicy {
	s.BillingMethod = &v
	return s
}

func (s *GetQuotaResponseBodyDataBillingPolicy) SetOdpsSpecCode(v string) *GetQuotaResponseBodyDataBillingPolicy {
	s.OdpsSpecCode = &v
	return s
}

func (s *GetQuotaResponseBodyDataBillingPolicy) SetOrderId(v string) *GetQuotaResponseBodyDataBillingPolicy {
	s.OrderId = &v
	return s
}

type GetQuotaResponseBodyDataSaleTag struct {
	// The identifier of an object in a MaxCompute quota. This identifier exists in the sales bill of Alibaba Cloud. You can use this identifier to associate the cost of a quota object with a tag.
	ResourceIds []*string `json:"resourceIds,omitempty" xml:"resourceIds,omitempty" type:"Repeated"`
	// The type of the object. Valid values: quota and project.
	//
	// example:
	//
	// "quota"
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s GetQuotaResponseBodyDataSaleTag) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaResponseBodyDataSaleTag) GoString() string {
	return s.String()
}

func (s *GetQuotaResponseBodyDataSaleTag) SetResourceIds(v []*string) *GetQuotaResponseBodyDataSaleTag {
	s.ResourceIds = v
	return s
}

func (s *GetQuotaResponseBodyDataSaleTag) SetResourceType(v string) *GetQuotaResponseBodyDataSaleTag {
	s.ResourceType = &v
	return s
}

type GetQuotaResponseBodyDataScheduleInfo struct {
	// The quota plan that takes effect based on the scheduling plan.
	//
	// example:
	//
	// planA
	CurrPlan *string `json:"currPlan,omitempty" xml:"currPlan,omitempty"`
	// The time when the current quota plan is scheduled.
	//
	// example:
	//
	// 0800
	CurrTime *string `json:"currTime,omitempty" xml:"currTime,omitempty"`
	// The next quota plan that will take effect based on the scheduling plan.
	//
	// example:
	//
	// planB
	NextPlan *string `json:"nextPlan,omitempty" xml:"nextPlan,omitempty"`
	// The time when the next quota plan is scheduled.
	//
	// example:
	//
	// 1700
	NextTime *string `json:"nextTime,omitempty" xml:"nextTime,omitempty"`
	// The quota plan that immediately takes effect. If the quota plan that immediately takes effect is different from the current quota plan, this parameter is not empty.
	//
	// example:
	//
	// planC
	OncePlan *string `json:"oncePlan,omitempty" xml:"oncePlan,omitempty"`
	// The time when the quota plan immediately takes effect.
	//
	// example:
	//
	// 1500
	OnceTime *string `json:"onceTime,omitempty" xml:"onceTime,omitempty"`
	// The name of the operator.
	//
	// example:
	//
	// userA
	OperatorName *string `json:"operatorName,omitempty" xml:"operatorName,omitempty"`
	// The time zone of the project.
	//
	// example:
	//
	// UTC+8
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
}

func (s GetQuotaResponseBodyDataScheduleInfo) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaResponseBodyDataScheduleInfo) GoString() string {
	return s.String()
}

func (s *GetQuotaResponseBodyDataScheduleInfo) SetCurrPlan(v string) *GetQuotaResponseBodyDataScheduleInfo {
	s.CurrPlan = &v
	return s
}

func (s *GetQuotaResponseBodyDataScheduleInfo) SetCurrTime(v string) *GetQuotaResponseBodyDataScheduleInfo {
	s.CurrTime = &v
	return s
}

func (s *GetQuotaResponseBodyDataScheduleInfo) SetNextPlan(v string) *GetQuotaResponseBodyDataScheduleInfo {
	s.NextPlan = &v
	return s
}

func (s *GetQuotaResponseBodyDataScheduleInfo) SetNextTime(v string) *GetQuotaResponseBodyDataScheduleInfo {
	s.NextTime = &v
	return s
}

func (s *GetQuotaResponseBodyDataScheduleInfo) SetOncePlan(v string) *GetQuotaResponseBodyDataScheduleInfo {
	s.OncePlan = &v
	return s
}

func (s *GetQuotaResponseBodyDataScheduleInfo) SetOnceTime(v string) *GetQuotaResponseBodyDataScheduleInfo {
	s.OnceTime = &v
	return s
}

func (s *GetQuotaResponseBodyDataScheduleInfo) SetOperatorName(v string) *GetQuotaResponseBodyDataScheduleInfo {
	s.OperatorName = &v
	return s
}

func (s *GetQuotaResponseBodyDataScheduleInfo) SetTimezone(v string) *GetQuotaResponseBodyDataScheduleInfo {
	s.Timezone = &v
	return s
}

type GetQuotaResponseBodyDataSubQuotaInfoList struct {
	// The information about the order.
	BillingPolicy *GetQuotaResponseBodyDataSubQuotaInfoListBillingPolicy `json:"billingPolicy,omitempty" xml:"billingPolicy,omitempty" type:"Struct"`
	// The cluster ID.
	//
	// example:
	//
	// AT-120N
	Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// The time when the resource was created.
	//
	// example:
	//
	// 2022-09-06T02:14:44Z
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The ID of the Alibaba Cloud account that is used to create the resource.
	//
	// example:
	//
	// 672863518
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// The ID of the level-2 quota.
	//
	// example:
	//
	// 1000048
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the level-2 quota.
	//
	// example:
	//
	// subquotaA
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The nickname of the level-2 quota.
	//
	// example:
	//
	// subquotaA
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	// The description of the quota.
	Parameter map[string]interface{} `json:"parameter,omitempty" xml:"parameter,omitempty"`
	// The ID of the parent resource.
	//
	// example:
	//
	// 0
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The identifier of an object in a MaxCompute quota. This identifier is the same as the identifier in the sales bill of Alibaba Cloud. This parameter is used for tags.
	SaleTag *GetQuotaResponseBodyDataSubQuotaInfoListSaleTag `json:"saleTag,omitempty" xml:"saleTag,omitempty" type:"Struct"`
	// The information about the scheduling plan.
	ScheduleInfo *GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo `json:"scheduleInfo,omitempty" xml:"scheduleInfo,omitempty" type:"Struct"`
	// The status of the resource.
	//
	// example:
	//
	// ON
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The tag of the resource for the quota.
	//
	// example:
	//
	// abc
	Tag *string `json:"tag,omitempty" xml:"tag,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 478403690625249
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The type of the resource system. This parameter corresponds to the resourceSystemType parameter of the cluster.
	//
	// example:
	//
	// FUXI_ONLINE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The version number.
	//
	// example:
	//
	// 1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetQuotaResponseBodyDataSubQuotaInfoList) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaResponseBodyDataSubQuotaInfoList) GoString() string {
	return s.String()
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoList) SetBillingPolicy(v *GetQuotaResponseBodyDataSubQuotaInfoListBillingPolicy) *GetQuotaResponseBodyDataSubQuotaInfoList {
	s.BillingPolicy = v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoList) SetCluster(v string) *GetQuotaResponseBodyDataSubQuotaInfoList {
	s.Cluster = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoList) SetCreateTime(v int64) *GetQuotaResponseBodyDataSubQuotaInfoList {
	s.CreateTime = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoList) SetCreatorId(v string) *GetQuotaResponseBodyDataSubQuotaInfoList {
	s.CreatorId = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoList) SetId(v string) *GetQuotaResponseBodyDataSubQuotaInfoList {
	s.Id = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoList) SetName(v string) *GetQuotaResponseBodyDataSubQuotaInfoList {
	s.Name = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoList) SetNickName(v string) *GetQuotaResponseBodyDataSubQuotaInfoList {
	s.NickName = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoList) SetParameter(v map[string]interface{}) *GetQuotaResponseBodyDataSubQuotaInfoList {
	s.Parameter = v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoList) SetParentId(v string) *GetQuotaResponseBodyDataSubQuotaInfoList {
	s.ParentId = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoList) SetRegionId(v string) *GetQuotaResponseBodyDataSubQuotaInfoList {
	s.RegionId = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoList) SetSaleTag(v *GetQuotaResponseBodyDataSubQuotaInfoListSaleTag) *GetQuotaResponseBodyDataSubQuotaInfoList {
	s.SaleTag = v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoList) SetScheduleInfo(v *GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) *GetQuotaResponseBodyDataSubQuotaInfoList {
	s.ScheduleInfo = v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoList) SetStatus(v string) *GetQuotaResponseBodyDataSubQuotaInfoList {
	s.Status = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoList) SetTag(v string) *GetQuotaResponseBodyDataSubQuotaInfoList {
	s.Tag = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoList) SetTenantId(v string) *GetQuotaResponseBodyDataSubQuotaInfoList {
	s.TenantId = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoList) SetType(v string) *GetQuotaResponseBodyDataSubQuotaInfoList {
	s.Type = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoList) SetVersion(v string) *GetQuotaResponseBodyDataSubQuotaInfoList {
	s.Version = &v
	return s
}

type GetQuotaResponseBodyDataSubQuotaInfoListBillingPolicy struct {
	// The billing method of the quota. Valid values:
	//
	// 	- subscription: a subscription quota.
	//
	// 	- payasyougo: a pay-as-you-go quota.
	//
	// example:
	//
	// subscription
	BillingMethod *string `json:"billingMethod,omitempty" xml:"billingMethod,omitempty"`
	// The specifications of the order.
	//
	// example:
	//
	// OdpsStandard
	OdpsSpecCode *string `json:"odpsSpecCode,omitempty" xml:"odpsSpecCode,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 45245678
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
}

func (s GetQuotaResponseBodyDataSubQuotaInfoListBillingPolicy) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaResponseBodyDataSubQuotaInfoListBillingPolicy) GoString() string {
	return s.String()
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoListBillingPolicy) SetBillingMethod(v string) *GetQuotaResponseBodyDataSubQuotaInfoListBillingPolicy {
	s.BillingMethod = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoListBillingPolicy) SetOdpsSpecCode(v string) *GetQuotaResponseBodyDataSubQuotaInfoListBillingPolicy {
	s.OdpsSpecCode = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoListBillingPolicy) SetOrderId(v string) *GetQuotaResponseBodyDataSubQuotaInfoListBillingPolicy {
	s.OrderId = &v
	return s
}

type GetQuotaResponseBodyDataSubQuotaInfoListSaleTag struct {
	// The identifier of an object in a MaxCompute quota. This identifier exists in the sales bill of Alibaba Cloud. You can use this identifier to associate the cost of a quota object with a tag.
	ResourceIds []*string `json:"resourceIds,omitempty" xml:"resourceIds,omitempty" type:"Repeated"`
	// The type of the object. Valid values: quota and project.
	//
	// example:
	//
	// "quota"
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s GetQuotaResponseBodyDataSubQuotaInfoListSaleTag) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaResponseBodyDataSubQuotaInfoListSaleTag) GoString() string {
	return s.String()
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoListSaleTag) SetResourceIds(v []*string) *GetQuotaResponseBodyDataSubQuotaInfoListSaleTag {
	s.ResourceIds = v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoListSaleTag) SetResourceType(v string) *GetQuotaResponseBodyDataSubQuotaInfoListSaleTag {
	s.ResourceType = &v
	return s
}

type GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo struct {
	// The quota plan that takes effect based on the scheduling plan.
	//
	// example:
	//
	// planA
	CurrPlan *string `json:"currPlan,omitempty" xml:"currPlan,omitempty"`
	// The time when the current quota plan is scheduled.
	//
	// example:
	//
	// 0800
	CurrTime *string `json:"currTime,omitempty" xml:"currTime,omitempty"`
	// The next quota plan that will take effect based on the scheduling plan.
	//
	// example:
	//
	// planB
	NextPlan *string `json:"nextPlan,omitempty" xml:"nextPlan,omitempty"`
	// The time when the next quota plan is scheduled.
	//
	// example:
	//
	// 1700
	NextTime *string `json:"nextTime,omitempty" xml:"nextTime,omitempty"`
	// The quota plan that immediately takes effect. If the quota plan that immediately takes effect is different from the current quota plan, this parameter is not empty.
	//
	// example:
	//
	// planC
	OncePlan *string `json:"oncePlan,omitempty" xml:"oncePlan,omitempty"`
	// The time when the quota plan immediately takes effect.
	//
	// example:
	//
	// 1500
	OnceTime *string `json:"onceTime,omitempty" xml:"onceTime,omitempty"`
	// The name of the operator.
	//
	// example:
	//
	// userA
	OperatorName *string `json:"operatorName,omitempty" xml:"operatorName,omitempty"`
	// The time zone of the project.
	//
	// example:
	//
	// UTC+8
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
}

func (s GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) GoString() string {
	return s.String()
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) SetCurrPlan(v string) *GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo {
	s.CurrPlan = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) SetCurrTime(v string) *GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo {
	s.CurrTime = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) SetNextPlan(v string) *GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo {
	s.NextPlan = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) SetNextTime(v string) *GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo {
	s.NextTime = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) SetOncePlan(v string) *GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo {
	s.OncePlan = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) SetOnceTime(v string) *GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo {
	s.OnceTime = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) SetOperatorName(v string) *GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo {
	s.OperatorName = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) SetTimezone(v string) *GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo {
	s.Timezone = &v
	return s
}

type GetQuotaResponseBodySaleTag struct {
	// The identifier of an object in a MaxCompute quota. This identifier exists in the sales bill of Alibaba Cloud. You can use this identifier to associate the cost of a quota object with a tag.
	ResourceIds []*string `json:"resourceIds,omitempty" xml:"resourceIds,omitempty" type:"Repeated"`
	// The type of the object. Valid values: quota and project.
	//
	// example:
	//
	// "quota"
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s GetQuotaResponseBodySaleTag) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaResponseBodySaleTag) GoString() string {
	return s.String()
}

func (s *GetQuotaResponseBodySaleTag) SetResourceIds(v []*string) *GetQuotaResponseBodySaleTag {
	s.ResourceIds = v
	return s
}

func (s *GetQuotaResponseBodySaleTag) SetResourceType(v string) *GetQuotaResponseBodySaleTag {
	s.ResourceType = &v
	return s
}

type GetQuotaResponseBodyScheduleInfo struct {
	// The quota plan that takes effect based on the scheduling plan.
	//
	// example:
	//
	// planA
	CurrPlan *string `json:"currPlan,omitempty" xml:"currPlan,omitempty"`
	// The time when the current quota plan is scheduled.
	//
	// example:
	//
	// 0800
	CurrTime *string `json:"currTime,omitempty" xml:"currTime,omitempty"`
	// The next quota plan that will take effect based on the scheduling plan.
	//
	// example:
	//
	// planB
	NextPlan *string `json:"nextPlan,omitempty" xml:"nextPlan,omitempty"`
	// The time when the next quota plan is scheduled.
	//
	// example:
	//
	// 1700
	NextTime *string `json:"nextTime,omitempty" xml:"nextTime,omitempty"`
	// The quota plan that immediately takes effect. If the quota plan that immediately takes effect is different from the current quota plan, this parameter is not empty.
	//
	// example:
	//
	// planC
	OncePlan *string `json:"oncePlan,omitempty" xml:"oncePlan,omitempty"`
	// The time when the quota plan immediately takes effect.
	//
	// example:
	//
	// 1500
	OnceTime *string `json:"onceTime,omitempty" xml:"onceTime,omitempty"`
	// The name of the operator.
	//
	// example:
	//
	// userA
	OperatorName *string `json:"operatorName,omitempty" xml:"operatorName,omitempty"`
	// The time zone of the project.
	//
	// example:
	//
	// UTC+8
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
}

func (s GetQuotaResponseBodyScheduleInfo) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaResponseBodyScheduleInfo) GoString() string {
	return s.String()
}

func (s *GetQuotaResponseBodyScheduleInfo) SetCurrPlan(v string) *GetQuotaResponseBodyScheduleInfo {
	s.CurrPlan = &v
	return s
}

func (s *GetQuotaResponseBodyScheduleInfo) SetCurrTime(v string) *GetQuotaResponseBodyScheduleInfo {
	s.CurrTime = &v
	return s
}

func (s *GetQuotaResponseBodyScheduleInfo) SetNextPlan(v string) *GetQuotaResponseBodyScheduleInfo {
	s.NextPlan = &v
	return s
}

func (s *GetQuotaResponseBodyScheduleInfo) SetNextTime(v string) *GetQuotaResponseBodyScheduleInfo {
	s.NextTime = &v
	return s
}

func (s *GetQuotaResponseBodyScheduleInfo) SetOncePlan(v string) *GetQuotaResponseBodyScheduleInfo {
	s.OncePlan = &v
	return s
}

func (s *GetQuotaResponseBodyScheduleInfo) SetOnceTime(v string) *GetQuotaResponseBodyScheduleInfo {
	s.OnceTime = &v
	return s
}

func (s *GetQuotaResponseBodyScheduleInfo) SetOperatorName(v string) *GetQuotaResponseBodyScheduleInfo {
	s.OperatorName = &v
	return s
}

func (s *GetQuotaResponseBodyScheduleInfo) SetTimezone(v string) *GetQuotaResponseBodyScheduleInfo {
	s.Timezone = &v
	return s
}

type GetQuotaResponseBodySubQuotaInfoList struct {
	// The information about the order.
	BillingPolicy *GetQuotaResponseBodySubQuotaInfoListBillingPolicy `json:"billingPolicy,omitempty" xml:"billingPolicy,omitempty" type:"Struct"`
	// The cluster ID.
	//
	// example:
	//
	// AT-120N
	Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// The time when the resource was created.
	//
	// example:
	//
	// 2022-09-06T02:14:44Z
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The ID of the Alibaba Cloud account that is used to create the resource.
	//
	// example:
	//
	// 672863518
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// The ID of the level-2 quota.
	//
	// example:
	//
	// 1000048
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the level-2 quota.
	//
	// example:
	//
	// subquotaA
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The alias of the level-2 quota.
	//
	// example:
	//
	// subquotaA
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	// The description of the quota.
	//
	// example:
	//
	// {"minCU":10,
	//
	// "adhocCU":0,
	//
	// "maxCU":10,
	//
	// "schedulerType":"Fair",
	//
	// }
	Parameter map[string]interface{} `json:"parameter,omitempty" xml:"parameter,omitempty"`
	// The ID of the parent resource.
	//
	// example:
	//
	// 0
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The identifier of an object in a MaxCompute quota. This identifier is the same as the identifier in the sales bill of Alibaba Cloud. This parameter is used for tags.
	SaleTag *GetQuotaResponseBodySubQuotaInfoListSaleTag `json:"saleTag,omitempty" xml:"saleTag,omitempty" type:"Struct"`
	// The information about the scheduling plan.
	ScheduleInfo *GetQuotaResponseBodySubQuotaInfoListScheduleInfo `json:"scheduleInfo,omitempty" xml:"scheduleInfo,omitempty" type:"Struct"`
	// The status of the resource.
	//
	// example:
	//
	// ON
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The tag of the resource for the quota.
	//
	// example:
	//
	// abc
	Tag *string `json:"tag,omitempty" xml:"tag,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 478403690625249
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The type of the resource system. This parameter corresponds to the resourceSystemType parameter of the cluster.
	//
	// example:
	//
	// FUXI_ONLINE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The version number.
	//
	// example:
	//
	// 1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetQuotaResponseBodySubQuotaInfoList) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaResponseBodySubQuotaInfoList) GoString() string {
	return s.String()
}

func (s *GetQuotaResponseBodySubQuotaInfoList) SetBillingPolicy(v *GetQuotaResponseBodySubQuotaInfoListBillingPolicy) *GetQuotaResponseBodySubQuotaInfoList {
	s.BillingPolicy = v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoList) SetCluster(v string) *GetQuotaResponseBodySubQuotaInfoList {
	s.Cluster = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoList) SetCreateTime(v int64) *GetQuotaResponseBodySubQuotaInfoList {
	s.CreateTime = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoList) SetCreatorId(v string) *GetQuotaResponseBodySubQuotaInfoList {
	s.CreatorId = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoList) SetId(v string) *GetQuotaResponseBodySubQuotaInfoList {
	s.Id = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoList) SetName(v string) *GetQuotaResponseBodySubQuotaInfoList {
	s.Name = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoList) SetNickName(v string) *GetQuotaResponseBodySubQuotaInfoList {
	s.NickName = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoList) SetParameter(v map[string]interface{}) *GetQuotaResponseBodySubQuotaInfoList {
	s.Parameter = v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoList) SetParentId(v string) *GetQuotaResponseBodySubQuotaInfoList {
	s.ParentId = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoList) SetRegionId(v string) *GetQuotaResponseBodySubQuotaInfoList {
	s.RegionId = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoList) SetSaleTag(v *GetQuotaResponseBodySubQuotaInfoListSaleTag) *GetQuotaResponseBodySubQuotaInfoList {
	s.SaleTag = v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoList) SetScheduleInfo(v *GetQuotaResponseBodySubQuotaInfoListScheduleInfo) *GetQuotaResponseBodySubQuotaInfoList {
	s.ScheduleInfo = v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoList) SetStatus(v string) *GetQuotaResponseBodySubQuotaInfoList {
	s.Status = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoList) SetTag(v string) *GetQuotaResponseBodySubQuotaInfoList {
	s.Tag = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoList) SetTenantId(v string) *GetQuotaResponseBodySubQuotaInfoList {
	s.TenantId = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoList) SetType(v string) *GetQuotaResponseBodySubQuotaInfoList {
	s.Type = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoList) SetVersion(v string) *GetQuotaResponseBodySubQuotaInfoList {
	s.Version = &v
	return s
}

type GetQuotaResponseBodySubQuotaInfoListBillingPolicy struct {
	// The billing method of the quota. Valid values:
	//
	// 	- subscription: a subscription quota.
	//
	// 	- payasyougo: a pay-as-you-go quota.
	//
	// example:
	//
	// subscription
	BillingMethod *string `json:"billingMethod,omitempty" xml:"billingMethod,omitempty"`
	// The specifications of the order.
	//
	// example:
	//
	// OdpsStandard
	OdpsSpecCode *string `json:"odpsSpecCode,omitempty" xml:"odpsSpecCode,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 45245678
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
}

func (s GetQuotaResponseBodySubQuotaInfoListBillingPolicy) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaResponseBodySubQuotaInfoListBillingPolicy) GoString() string {
	return s.String()
}

func (s *GetQuotaResponseBodySubQuotaInfoListBillingPolicy) SetBillingMethod(v string) *GetQuotaResponseBodySubQuotaInfoListBillingPolicy {
	s.BillingMethod = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoListBillingPolicy) SetOdpsSpecCode(v string) *GetQuotaResponseBodySubQuotaInfoListBillingPolicy {
	s.OdpsSpecCode = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoListBillingPolicy) SetOrderId(v string) *GetQuotaResponseBodySubQuotaInfoListBillingPolicy {
	s.OrderId = &v
	return s
}

type GetQuotaResponseBodySubQuotaInfoListSaleTag struct {
	// The identifier of an object in a MaxCompute quota. This identifier exists in the sales bill of Alibaba Cloud. You can use this identifier to associate the cost of a quota object with a tag.
	ResourceIds []*string `json:"resourceIds,omitempty" xml:"resourceIds,omitempty" type:"Repeated"`
	// The type of the object. Valid values: quota and project.
	//
	// example:
	//
	// "quota"
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s GetQuotaResponseBodySubQuotaInfoListSaleTag) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaResponseBodySubQuotaInfoListSaleTag) GoString() string {
	return s.String()
}

func (s *GetQuotaResponseBodySubQuotaInfoListSaleTag) SetResourceIds(v []*string) *GetQuotaResponseBodySubQuotaInfoListSaleTag {
	s.ResourceIds = v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoListSaleTag) SetResourceType(v string) *GetQuotaResponseBodySubQuotaInfoListSaleTag {
	s.ResourceType = &v
	return s
}

type GetQuotaResponseBodySubQuotaInfoListScheduleInfo struct {
	// The quota plan that takes effect based on the scheduling plan.
	//
	// example:
	//
	// planA
	CurrPlan *string `json:"currPlan,omitempty" xml:"currPlan,omitempty"`
	// The time when the current quota plan is scheduled.
	//
	// example:
	//
	// 0800
	CurrTime *string `json:"currTime,omitempty" xml:"currTime,omitempty"`
	// The next quota plan that will take effect based on the scheduling plan.
	//
	// example:
	//
	// planB
	NextPlan *string `json:"nextPlan,omitempty" xml:"nextPlan,omitempty"`
	// The time when the next quota plan is scheduled.
	//
	// example:
	//
	// 1700
	NextTime *string `json:"nextTime,omitempty" xml:"nextTime,omitempty"`
	// The quota plan that immediately takes effect. If the quota plan that immediately takes effect is different from the current quota plan, this parameter is not empty.
	//
	// example:
	//
	// planC
	OncePlan *string `json:"oncePlan,omitempty" xml:"oncePlan,omitempty"`
	// The time when the quota plan immediately takes effect.
	//
	// example:
	//
	// 1500
	OnceTime *string `json:"onceTime,omitempty" xml:"onceTime,omitempty"`
	// The name of the operator.
	//
	// example:
	//
	// userA
	OperatorName *string `json:"operatorName,omitempty" xml:"operatorName,omitempty"`
	// The time zone of the project.
	//
	// example:
	//
	// UTC+8
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
}

func (s GetQuotaResponseBodySubQuotaInfoListScheduleInfo) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaResponseBodySubQuotaInfoListScheduleInfo) GoString() string {
	return s.String()
}

func (s *GetQuotaResponseBodySubQuotaInfoListScheduleInfo) SetCurrPlan(v string) *GetQuotaResponseBodySubQuotaInfoListScheduleInfo {
	s.CurrPlan = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoListScheduleInfo) SetCurrTime(v string) *GetQuotaResponseBodySubQuotaInfoListScheduleInfo {
	s.CurrTime = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoListScheduleInfo) SetNextPlan(v string) *GetQuotaResponseBodySubQuotaInfoListScheduleInfo {
	s.NextPlan = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoListScheduleInfo) SetNextTime(v string) *GetQuotaResponseBodySubQuotaInfoListScheduleInfo {
	s.NextTime = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoListScheduleInfo) SetOncePlan(v string) *GetQuotaResponseBodySubQuotaInfoListScheduleInfo {
	s.OncePlan = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoListScheduleInfo) SetOnceTime(v string) *GetQuotaResponseBodySubQuotaInfoListScheduleInfo {
	s.OnceTime = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoListScheduleInfo) SetOperatorName(v string) *GetQuotaResponseBodySubQuotaInfoListScheduleInfo {
	s.OperatorName = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoListScheduleInfo) SetTimezone(v string) *GetQuotaResponseBodySubQuotaInfoListScheduleInfo {
	s.Timezone = &v
	return s
}

type GetQuotaResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQuotaResponse) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaResponse) GoString() string {
	return s.String()
}

func (s *GetQuotaResponse) SetHeaders(v map[string]*string) *GetQuotaResponse {
	s.Headers = v
	return s
}

func (s *GetQuotaResponse) SetStatusCode(v int32) *GetQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQuotaResponse) SetBody(v *GetQuotaResponseBody) *GetQuotaResponse {
	s.Body = v
	return s
}

type GetQuotaPlanRequest struct {
	// The ID of the region.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The ID of the tenant.
	//
	// example:
	//
	// 483212237127906
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetQuotaPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaPlanRequest) GoString() string {
	return s.String()
}

func (s *GetQuotaPlanRequest) SetRegion(v string) *GetQuotaPlanRequest {
	s.Region = &v
	return s
}

func (s *GetQuotaPlanRequest) SetTenantId(v string) *GetQuotaPlanRequest {
	s.TenantId = &v
	return s
}

type GetQuotaPlanResponseBody struct {
	// The returned data.
	Data *GetQuotaPlanResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 0be3e0aa16667684362147582e038f
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetQuotaPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaPlanResponseBody) GoString() string {
	return s.String()
}

func (s *GetQuotaPlanResponseBody) SetData(v *GetQuotaPlanResponseBodyData) *GetQuotaPlanResponseBody {
	s.Data = v
	return s
}

func (s *GetQuotaPlanResponseBody) SetRequestId(v string) *GetQuotaPlanResponseBody {
	s.RequestId = &v
	return s
}

type GetQuotaPlanResponseBodyData struct {
	// The time when the quota plan was created.
	//
	// example:
	//
	// 2022-05-16T06:07:45Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The name of the quota plan.
	//
	// example:
	//
	// planA
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The details of the quota.
	Quota *GetQuotaPlanResponseBodyDataQuota `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
}

func (s GetQuotaPlanResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaPlanResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetQuotaPlanResponseBodyData) SetCreateTime(v string) *GetQuotaPlanResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetQuotaPlanResponseBodyData) SetName(v string) *GetQuotaPlanResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetQuotaPlanResponseBodyData) SetQuota(v *GetQuotaPlanResponseBodyDataQuota) *GetQuotaPlanResponseBodyData {
	s.Quota = v
	return s
}

type GetQuotaPlanResponseBodyDataQuota struct {
	// The information of the order.
	BillingPolicy *GetQuotaPlanResponseBodyDataQuotaBillingPolicy `json:"billingPolicy,omitempty" xml:"billingPolicy,omitempty" type:"Struct"`
	// The ID of the cluster.
	//
	// example:
	//
	// AT-120N
	Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// The time when the quota plan was created.
	//
	// example:
	//
	// 2022-09-06T02:14:44Z
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The ID of the Alibaba Cloud account that is used to create the resource.
	//
	// example:
	//
	// 672863518
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// The ID of the quota.
	//
	// example:
	//
	// 0
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the quota.
	//
	// example:
	//
	// quota_a
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The alias of the quota.
	//
	// example:
	//
	// quota_nickname
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	// The description of the quota.
	Parameter map[string]interface{} `json:"parameter,omitempty" xml:"parameter,omitempty"`
	// The ID of the parent resource.
	//
	// example:
	//
	// null
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The information of the scheduling plan.
	ScheduleInfo *GetQuotaPlanResponseBodyDataQuotaScheduleInfo `json:"scheduleInfo,omitempty" xml:"scheduleInfo,omitempty" type:"Struct"`
	// The status of the resource.
	//
	// example:
	//
	// ON
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The information of the level-2 quota.
	SubQuotaInfoList []*GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList `json:"subQuotaInfoList,omitempty" xml:"subQuotaInfoList,omitempty" type:"Repeated"`
	// The tag of the resource for the quota.
	//
	// example:
	//
	// abc
	Tag *string `json:"tag,omitempty" xml:"tag,omitempty"`
	// The ID of the tenant.
	//
	// example:
	//
	// 478403690625249
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The type of the resource system. This parameter corresponds to the resourceSystemType parameter of the cluster.
	//
	// example:
	//
	// FUXI_ONLINE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The version number.
	//
	// example:
	//
	// 1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetQuotaPlanResponseBodyDataQuota) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaPlanResponseBodyDataQuota) GoString() string {
	return s.String()
}

func (s *GetQuotaPlanResponseBodyDataQuota) SetBillingPolicy(v *GetQuotaPlanResponseBodyDataQuotaBillingPolicy) *GetQuotaPlanResponseBodyDataQuota {
	s.BillingPolicy = v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuota) SetCluster(v string) *GetQuotaPlanResponseBodyDataQuota {
	s.Cluster = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuota) SetCreateTime(v int64) *GetQuotaPlanResponseBodyDataQuota {
	s.CreateTime = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuota) SetCreatorId(v string) *GetQuotaPlanResponseBodyDataQuota {
	s.CreatorId = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuota) SetId(v string) *GetQuotaPlanResponseBodyDataQuota {
	s.Id = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuota) SetName(v string) *GetQuotaPlanResponseBodyDataQuota {
	s.Name = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuota) SetNickName(v string) *GetQuotaPlanResponseBodyDataQuota {
	s.NickName = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuota) SetParameter(v map[string]interface{}) *GetQuotaPlanResponseBodyDataQuota {
	s.Parameter = v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuota) SetParentId(v string) *GetQuotaPlanResponseBodyDataQuota {
	s.ParentId = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuota) SetRegionId(v string) *GetQuotaPlanResponseBodyDataQuota {
	s.RegionId = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuota) SetScheduleInfo(v *GetQuotaPlanResponseBodyDataQuotaScheduleInfo) *GetQuotaPlanResponseBodyDataQuota {
	s.ScheduleInfo = v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuota) SetStatus(v string) *GetQuotaPlanResponseBodyDataQuota {
	s.Status = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuota) SetSubQuotaInfoList(v []*GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) *GetQuotaPlanResponseBodyDataQuota {
	s.SubQuotaInfoList = v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuota) SetTag(v string) *GetQuotaPlanResponseBodyDataQuota {
	s.Tag = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuota) SetTenantId(v string) *GetQuotaPlanResponseBodyDataQuota {
	s.TenantId = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuota) SetType(v string) *GetQuotaPlanResponseBodyDataQuota {
	s.Type = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuota) SetVersion(v string) *GetQuotaPlanResponseBodyDataQuota {
	s.Version = &v
	return s
}

type GetQuotaPlanResponseBodyDataQuotaBillingPolicy struct {
	// The billing method of the quota. Valid values:
	//
	// 	- subscription: a subscription quota.
	//
	// 	- payasyougo: a pay-as-you-go quota.
	//
	// example:
	//
	// subscription
	BillingMethod *string `json:"billingMethod,omitempty" xml:"billingMethod,omitempty"`
	// The specifications of the order.
	//
	// example:
	//
	// OdpsStandard
	OdpsSpecCode *string `json:"odpsSpecCode,omitempty" xml:"odpsSpecCode,omitempty"`
	// The ID of the order.
	//
	// example:
	//
	// 45245678
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
}

func (s GetQuotaPlanResponseBodyDataQuotaBillingPolicy) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaPlanResponseBodyDataQuotaBillingPolicy) GoString() string {
	return s.String()
}

func (s *GetQuotaPlanResponseBodyDataQuotaBillingPolicy) SetBillingMethod(v string) *GetQuotaPlanResponseBodyDataQuotaBillingPolicy {
	s.BillingMethod = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaBillingPolicy) SetOdpsSpecCode(v string) *GetQuotaPlanResponseBodyDataQuotaBillingPolicy {
	s.OdpsSpecCode = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaBillingPolicy) SetOrderId(v string) *GetQuotaPlanResponseBodyDataQuotaBillingPolicy {
	s.OrderId = &v
	return s
}

type GetQuotaPlanResponseBodyDataQuotaScheduleInfo struct {
	// The quota plan that takes effect based on the scheduling plan.
	//
	// example:
	//
	// planA
	CurrPlan *string `json:"currPlan,omitempty" xml:"currPlan,omitempty"`
	// The time when the current quota plan is scheduled.
	//
	// example:
	//
	// 0800
	CurrTime *string `json:"currTime,omitempty" xml:"currTime,omitempty"`
	// The next quota plan that will take effect based on the scheduling plan.
	//
	// example:
	//
	// planB
	NextPlan *string `json:"nextPlan,omitempty" xml:"nextPlan,omitempty"`
	// The time when the next quota plan is scheduled.
	//
	// example:
	//
	// 1700
	NextTime *string `json:"nextTime,omitempty" xml:"nextTime,omitempty"`
	// If the quota plan that immediately takes effect is different from the current quota plan, this parameter is not empty.
	//
	// example:
	//
	// planC
	OncePlan *string `json:"oncePlan,omitempty" xml:"oncePlan,omitempty"`
	// The time when the quota plan immediately takes effect.
	//
	// example:
	//
	// 1500
	OnceTime *string `json:"onceTime,omitempty" xml:"onceTime,omitempty"`
	// The name of the operator.
	//
	// example:
	//
	// userA
	OperatorName *string `json:"operatorName,omitempty" xml:"operatorName,omitempty"`
}

func (s GetQuotaPlanResponseBodyDataQuotaScheduleInfo) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaPlanResponseBodyDataQuotaScheduleInfo) GoString() string {
	return s.String()
}

func (s *GetQuotaPlanResponseBodyDataQuotaScheduleInfo) SetCurrPlan(v string) *GetQuotaPlanResponseBodyDataQuotaScheduleInfo {
	s.CurrPlan = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaScheduleInfo) SetCurrTime(v string) *GetQuotaPlanResponseBodyDataQuotaScheduleInfo {
	s.CurrTime = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaScheduleInfo) SetNextPlan(v string) *GetQuotaPlanResponseBodyDataQuotaScheduleInfo {
	s.NextPlan = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaScheduleInfo) SetNextTime(v string) *GetQuotaPlanResponseBodyDataQuotaScheduleInfo {
	s.NextTime = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaScheduleInfo) SetOncePlan(v string) *GetQuotaPlanResponseBodyDataQuotaScheduleInfo {
	s.OncePlan = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaScheduleInfo) SetOnceTime(v string) *GetQuotaPlanResponseBodyDataQuotaScheduleInfo {
	s.OnceTime = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaScheduleInfo) SetOperatorName(v string) *GetQuotaPlanResponseBodyDataQuotaScheduleInfo {
	s.OperatorName = &v
	return s
}

type GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList struct {
	// The information of the order.
	BillingPolicy *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListBillingPolicy `json:"billingPolicy,omitempty" xml:"billingPolicy,omitempty" type:"Struct"`
	// The ID of the cluster.
	//
	// example:
	//
	// AT-120N
	Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// The time when the resource was created.
	//
	// example:
	//
	// 2022-09-06T02:14:44Z
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The ID of the user who created the quota plan.
	//
	// example:
	//
	// 672863518
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// The ID of the level-2 quota.
	//
	// example:
	//
	// 1000048
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the level-2 quota.
	//
	// example:
	//
	// subquotaA
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The alias of the level-2 quota.
	//
	// example:
	//
	// subquotaA
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	// The description of the quota.
	Parameter map[string]interface{} `json:"parameter,omitempty" xml:"parameter,omitempty"`
	// The ID of the parent resource.
	//
	// example:
	//
	// 0
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The information of the scheduling plan.
	ScheduleInfo *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListScheduleInfo `json:"scheduleInfo,omitempty" xml:"scheduleInfo,omitempty" type:"Struct"`
	// The status of the resource.
	//
	// example:
	//
	// ON
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The tag of the resource for the quota.
	//
	// example:
	//
	// abc
	Tag *string `json:"tag,omitempty" xml:"tag,omitempty"`
	// The ID of the tenant.
	//
	// example:
	//
	// 478403690625249
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The type of the resource system. This parameter corresponds to the resourceSystemType parameter of the cluster.
	//
	// example:
	//
	// FUXI_ONLINE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The version number.
	//
	// example:
	//
	// 1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) GoString() string {
	return s.String()
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetBillingPolicy(v *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListBillingPolicy) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.BillingPolicy = v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetCluster(v string) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.Cluster = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetCreateTime(v int64) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.CreateTime = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetCreatorId(v string) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.CreatorId = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetId(v string) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.Id = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetName(v string) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.Name = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetNickName(v string) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.NickName = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetParameter(v map[string]interface{}) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.Parameter = v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetParentId(v string) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.ParentId = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetRegionId(v string) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.RegionId = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetScheduleInfo(v *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListScheduleInfo) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.ScheduleInfo = v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetStatus(v string) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.Status = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetTag(v string) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.Tag = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetTenantId(v string) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.TenantId = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetType(v string) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.Type = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetVersion(v string) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.Version = &v
	return s
}

type GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListBillingPolicy struct {
	// The billing method of the quota. Valid values:
	//
	// 	- subscription: a subscription quota.
	//
	// 	- payasyougo: a pay-as-you-go quota.
	//
	// example:
	//
	// subscription
	BillingMethod *string `json:"billingMethod,omitempty" xml:"billingMethod,omitempty"`
	// The specifications of the order.
	//
	// example:
	//
	// OdpsStandard
	OdpsSpecCode *string `json:"odpsSpecCode,omitempty" xml:"odpsSpecCode,omitempty"`
	// The ID of the order.
	//
	// example:
	//
	// 45245678
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
}

func (s GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListBillingPolicy) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListBillingPolicy) GoString() string {
	return s.String()
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListBillingPolicy) SetBillingMethod(v string) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListBillingPolicy {
	s.BillingMethod = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListBillingPolicy) SetOdpsSpecCode(v string) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListBillingPolicy {
	s.OdpsSpecCode = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListBillingPolicy) SetOrderId(v string) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListBillingPolicy {
	s.OrderId = &v
	return s
}

type GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListScheduleInfo struct {
	// The quota plan that takes effect based on the scheduling plan.
	//
	// example:
	//
	// planA
	CurrPlan *string `json:"currPlan,omitempty" xml:"currPlan,omitempty"`
	// The time when the current quota plan is scheduled.
	//
	// example:
	//
	// 0800
	CurrTime *string `json:"currTime,omitempty" xml:"currTime,omitempty"`
	// The next quota plan that will take effect based on the scheduling plan.
	//
	// example:
	//
	// planB
	NextPlan *string `json:"nextPlan,omitempty" xml:"nextPlan,omitempty"`
	// The time when the next quota plan is scheduled.
	//
	// example:
	//
	// 1700
	NextTime *string `json:"nextTime,omitempty" xml:"nextTime,omitempty"`
	// If the quota plan that immediately takes effect is different from the current quota plan, this parameter is not empty.
	//
	// example:
	//
	// planC
	OncePlan *string `json:"oncePlan,omitempty" xml:"oncePlan,omitempty"`
	// The time when the quota plan immediately takes effect.
	//
	// example:
	//
	// 1500
	OnceTime *string `json:"onceTime,omitempty" xml:"onceTime,omitempty"`
	// The name of the operator.
	//
	// example:
	//
	// userA
	OperatorName *string `json:"operatorName,omitempty" xml:"operatorName,omitempty"`
}

func (s GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListScheduleInfo) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListScheduleInfo) GoString() string {
	return s.String()
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListScheduleInfo) SetCurrPlan(v string) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListScheduleInfo {
	s.CurrPlan = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListScheduleInfo) SetCurrTime(v string) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListScheduleInfo {
	s.CurrTime = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListScheduleInfo) SetNextPlan(v string) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListScheduleInfo {
	s.NextPlan = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListScheduleInfo) SetNextTime(v string) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListScheduleInfo {
	s.NextTime = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListScheduleInfo) SetOncePlan(v string) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListScheduleInfo {
	s.OncePlan = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListScheduleInfo) SetOnceTime(v string) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListScheduleInfo {
	s.OnceTime = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListScheduleInfo) SetOperatorName(v string) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListScheduleInfo {
	s.OperatorName = &v
	return s
}

type GetQuotaPlanResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQuotaPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQuotaPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaPlanResponse) GoString() string {
	return s.String()
}

func (s *GetQuotaPlanResponse) SetHeaders(v map[string]*string) *GetQuotaPlanResponse {
	s.Headers = v
	return s
}

func (s *GetQuotaPlanResponse) SetStatusCode(v int32) *GetQuotaPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQuotaPlanResponse) SetBody(v *GetQuotaPlanResponseBody) *GetQuotaPlanResponse {
	s.Body = v
	return s
}

type GetQuotaScheduleRequest struct {
	// The time zone.
	//
	// example:
	//
	// UTC+8
	DisplayTimezone *string `json:"displayTimezone,omitempty" xml:"displayTimezone,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The ID of the tenant.
	//
	// example:
	//
	// 478403690625249
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetQuotaScheduleRequest) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaScheduleRequest) GoString() string {
	return s.String()
}

func (s *GetQuotaScheduleRequest) SetDisplayTimezone(v string) *GetQuotaScheduleRequest {
	s.DisplayTimezone = &v
	return s
}

func (s *GetQuotaScheduleRequest) SetRegion(v string) *GetQuotaScheduleRequest {
	s.Region = &v
	return s
}

func (s *GetQuotaScheduleRequest) SetTenantId(v string) *GetQuotaScheduleRequest {
	s.TenantId = &v
	return s
}

type GetQuotaScheduleResponseBody struct {
	// The returned data.
	Data []*GetQuotaScheduleResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// 	- If the value of success was false, an error code was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// success
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// errorMsg
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// Indicates whether the request was successful. If this parameter was not empty and the value of this parameter was not 200, the request failed.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0bc059b716696296266308790e0d3e
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetQuotaScheduleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaScheduleResponseBody) GoString() string {
	return s.String()
}

func (s *GetQuotaScheduleResponseBody) SetData(v []*GetQuotaScheduleResponseBodyData) *GetQuotaScheduleResponseBody {
	s.Data = v
	return s
}

func (s *GetQuotaScheduleResponseBody) SetErrorCode(v string) *GetQuotaScheduleResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetQuotaScheduleResponseBody) SetErrorMsg(v string) *GetQuotaScheduleResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetQuotaScheduleResponseBody) SetHttpCode(v int32) *GetQuotaScheduleResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetQuotaScheduleResponseBody) SetRequestId(v string) *GetQuotaScheduleResponseBody {
	s.RequestId = &v
	return s
}

type GetQuotaScheduleResponseBodyData struct {
	// The condition value.
	Condition *GetQuotaScheduleResponseBodyDataCondition `json:"condition,omitempty" xml:"condition,omitempty" type:"Struct"`
	// The ID of the quota plan.
	//
	// example:
	//
	// 63
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the operator.
	//
	// example:
	//
	// userA
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// The name of the quota plan.
	//
	// example:
	//
	// planA
	Plan *string `json:"plan,omitempty" xml:"plan,omitempty"`
	// The time zone.
	//
	// example:
	//
	// UTC+8
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
	// The type of the quota plan.
	//
	// example:
	//
	// once
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetQuotaScheduleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaScheduleResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetQuotaScheduleResponseBodyData) SetCondition(v *GetQuotaScheduleResponseBodyDataCondition) *GetQuotaScheduleResponseBodyData {
	s.Condition = v
	return s
}

func (s *GetQuotaScheduleResponseBodyData) SetId(v string) *GetQuotaScheduleResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetQuotaScheduleResponseBodyData) SetOperator(v string) *GetQuotaScheduleResponseBodyData {
	s.Operator = &v
	return s
}

func (s *GetQuotaScheduleResponseBodyData) SetPlan(v string) *GetQuotaScheduleResponseBodyData {
	s.Plan = &v
	return s
}

func (s *GetQuotaScheduleResponseBodyData) SetTimezone(v string) *GetQuotaScheduleResponseBodyData {
	s.Timezone = &v
	return s
}

func (s *GetQuotaScheduleResponseBodyData) SetType(v string) *GetQuotaScheduleResponseBodyData {
	s.Type = &v
	return s
}

type GetQuotaScheduleResponseBodyDataCondition struct {
	// The start time when the quota plan takes effect.
	//
	// example:
	//
	// 2022-04-25T04:23:04Z
	After *string `json:"after,omitempty" xml:"after,omitempty"`
	// The time when the quota plan takes effect.
	//
	// example:
	//
	// 0900
	At *string `json:"at,omitempty" xml:"at,omitempty"`
}

func (s GetQuotaScheduleResponseBodyDataCondition) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaScheduleResponseBodyDataCondition) GoString() string {
	return s.String()
}

func (s *GetQuotaScheduleResponseBodyDataCondition) SetAfter(v string) *GetQuotaScheduleResponseBodyDataCondition {
	s.After = &v
	return s
}

func (s *GetQuotaScheduleResponseBodyDataCondition) SetAt(v string) *GetQuotaScheduleResponseBodyDataCondition {
	s.At = &v
	return s
}

type GetQuotaScheduleResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQuotaScheduleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQuotaScheduleResponse) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaScheduleResponse) GoString() string {
	return s.String()
}

func (s *GetQuotaScheduleResponse) SetHeaders(v map[string]*string) *GetQuotaScheduleResponse {
	s.Headers = v
	return s
}

func (s *GetQuotaScheduleResponse) SetStatusCode(v int32) *GetQuotaScheduleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQuotaScheduleResponse) SetBody(v *GetQuotaScheduleResponseBody) *GetQuotaScheduleResponse {
	s.Body = v
	return s
}

type GetQuotaUsageRequest struct {
	// The aggregation algorithm. For a better page experience, up to 60 points can be displayed for each metric. If you select a time range longer than 1 hour, the chart uses the average value within the range (minutes of the selected time range/60) to aggregate data by default. You can change the aggregation algorithm based on your business requirements.
	//
	// example:
	//
	// max
	AggMethod *string `json:"aggMethod,omitempty" xml:"aggMethod,omitempty"`
	// The time when the query starts. The value is the log time that is specified when log data is written.
	//
	// 	- The time range that is specified in this operation is a left-closed, right-open interval. The interval includes the start time specified by the **from*	- parameter, but does not include the end time specified by the **to*	- parameter. If you set the **from*	- and **to*	- parameters to the same value, the time range is invalid and an error message is returned.
	//
	// 	- This value is a UNIX timestamp representing the number of seconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1669081045
	From *int64 `json:"from,omitempty" xml:"from,omitempty"`
	// The types of the charts.
	PlotTypes []*string `json:"plotTypes,omitempty" xml:"plotTypes,omitempty" type:"Repeated"`
	// The quota type. Default value: ODPS.
	//
	// 	- ODPS: computing quota
	//
	// 	- TUNNEL: Tunnel quota
	//
	// example:
	//
	// ODPS
	ProductId *string `json:"productId,omitempty" xml:"productId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-chengdu
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The alias of the level-2 quota.
	//
	// example:
	//
	// ot_tunnel_quota
	SubQuotaNickname *string `json:"subQuotaNickname,omitempty" xml:"subQuotaNickname,omitempty"`
	// The ID of the tenant. You can log on to the MaxCompute console, and choose Tenants > Tenant Property from the left-side navigation pane to view the tenant ID.
	//
	// example:
	//
	// 478403690625249
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The time when the query ends. The value is the log time that is specified when log data is written.
	//
	// 	- The time range that is specified in this operation is a left-closed, right-open interval. The interval includes the start time specified by the **from*	- parameter, but does not include the end time specified by the **to*	- parameter. If you set the **from*	- and **to*	- parameters to the same value, the time range is invalid and an error message is returned.
	//
	// 	- This value is a UNIX timestamp representing the number of seconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1669360870
	To *int64 `json:"to,omitempty" xml:"to,omitempty"`
	// The data metric fields.
	YAxisTypes []*string `json:"yAxisTypes,omitempty" xml:"yAxisTypes,omitempty" type:"Repeated"`
}

func (s GetQuotaUsageRequest) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaUsageRequest) GoString() string {
	return s.String()
}

func (s *GetQuotaUsageRequest) SetAggMethod(v string) *GetQuotaUsageRequest {
	s.AggMethod = &v
	return s
}

func (s *GetQuotaUsageRequest) SetFrom(v int64) *GetQuotaUsageRequest {
	s.From = &v
	return s
}

func (s *GetQuotaUsageRequest) SetPlotTypes(v []*string) *GetQuotaUsageRequest {
	s.PlotTypes = v
	return s
}

func (s *GetQuotaUsageRequest) SetProductId(v string) *GetQuotaUsageRequest {
	s.ProductId = &v
	return s
}

func (s *GetQuotaUsageRequest) SetRegion(v string) *GetQuotaUsageRequest {
	s.Region = &v
	return s
}

func (s *GetQuotaUsageRequest) SetSubQuotaNickname(v string) *GetQuotaUsageRequest {
	s.SubQuotaNickname = &v
	return s
}

func (s *GetQuotaUsageRequest) SetTenantId(v string) *GetQuotaUsageRequest {
	s.TenantId = &v
	return s
}

func (s *GetQuotaUsageRequest) SetTo(v int64) *GetQuotaUsageRequest {
	s.To = &v
	return s
}

func (s *GetQuotaUsageRequest) SetYAxisTypes(v []*string) *GetQuotaUsageRequest {
	s.YAxisTypes = v
	return s
}

type GetQuotaUsageShrinkRequest struct {
	// The aggregation algorithm. For a better page experience, up to 60 points can be displayed for each metric. If you select a time range longer than 1 hour, the chart uses the average value within the range (minutes of the selected time range/60) to aggregate data by default. You can change the aggregation algorithm based on your business requirements.
	//
	// example:
	//
	// max
	AggMethod *string `json:"aggMethod,omitempty" xml:"aggMethod,omitempty"`
	// The time when the query starts. The value is the log time that is specified when log data is written.
	//
	// 	- The time range that is specified in this operation is a left-closed, right-open interval. The interval includes the start time specified by the **from*	- parameter, but does not include the end time specified by the **to*	- parameter. If you set the **from*	- and **to*	- parameters to the same value, the time range is invalid and an error message is returned.
	//
	// 	- This value is a UNIX timestamp representing the number of seconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1669081045
	From *int64 `json:"from,omitempty" xml:"from,omitempty"`
	// The types of the charts.
	PlotTypesShrink *string `json:"plotTypes,omitempty" xml:"plotTypes,omitempty"`
	// The quota type. Default value: ODPS.
	//
	// 	- ODPS: computing quota
	//
	// 	- TUNNEL: Tunnel quota
	//
	// example:
	//
	// ODPS
	ProductId *string `json:"productId,omitempty" xml:"productId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-chengdu
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The alias of the level-2 quota.
	//
	// example:
	//
	// ot_tunnel_quota
	SubQuotaNickname *string `json:"subQuotaNickname,omitempty" xml:"subQuotaNickname,omitempty"`
	// The ID of the tenant. You can log on to the MaxCompute console, and choose Tenants > Tenant Property from the left-side navigation pane to view the tenant ID.
	//
	// example:
	//
	// 478403690625249
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The time when the query ends. The value is the log time that is specified when log data is written.
	//
	// 	- The time range that is specified in this operation is a left-closed, right-open interval. The interval includes the start time specified by the **from*	- parameter, but does not include the end time specified by the **to*	- parameter. If you set the **from*	- and **to*	- parameters to the same value, the time range is invalid and an error message is returned.
	//
	// 	- This value is a UNIX timestamp representing the number of seconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1669360870
	To *int64 `json:"to,omitempty" xml:"to,omitempty"`
	// The data metric fields.
	YAxisTypesShrink *string `json:"yAxisTypes,omitempty" xml:"yAxisTypes,omitempty"`
}

func (s GetQuotaUsageShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaUsageShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetQuotaUsageShrinkRequest) SetAggMethod(v string) *GetQuotaUsageShrinkRequest {
	s.AggMethod = &v
	return s
}

func (s *GetQuotaUsageShrinkRequest) SetFrom(v int64) *GetQuotaUsageShrinkRequest {
	s.From = &v
	return s
}

func (s *GetQuotaUsageShrinkRequest) SetPlotTypesShrink(v string) *GetQuotaUsageShrinkRequest {
	s.PlotTypesShrink = &v
	return s
}

func (s *GetQuotaUsageShrinkRequest) SetProductId(v string) *GetQuotaUsageShrinkRequest {
	s.ProductId = &v
	return s
}

func (s *GetQuotaUsageShrinkRequest) SetRegion(v string) *GetQuotaUsageShrinkRequest {
	s.Region = &v
	return s
}

func (s *GetQuotaUsageShrinkRequest) SetSubQuotaNickname(v string) *GetQuotaUsageShrinkRequest {
	s.SubQuotaNickname = &v
	return s
}

func (s *GetQuotaUsageShrinkRequest) SetTenantId(v string) *GetQuotaUsageShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *GetQuotaUsageShrinkRequest) SetTo(v int64) *GetQuotaUsageShrinkRequest {
	s.To = &v
	return s
}

func (s *GetQuotaUsageShrinkRequest) SetYAxisTypesShrink(v string) *GetQuotaUsageShrinkRequest {
	s.YAxisTypesShrink = &v
	return s
}

type GetQuotaUsageResponseBody struct {
	// The data returned.
	Data *GetQuotaUsageResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// OBJECT_NOT_EXIST
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// This object does not exist.
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// The HTTP status code.
	//
	// 	- 1xx: informational response. The request is received and is being processed.
	//
	// 	- 2xx: success. The request is successfully received, understood, and accepted by the server.
	//
	// 	- 3xx: redirection. The request is redirected, and further actions are required to complete the request.
	//
	// 	- 4xx: client error. The request contains invalid request parameters and syntaxes, or specific request conditions cannot be met.
	//
	// 	- 5xx: server error. The server cannot meet requirements due to other reasons.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0b87b7a416652014358483492eea0b
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetQuotaUsageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaUsageResponseBody) GoString() string {
	return s.String()
}

func (s *GetQuotaUsageResponseBody) SetData(v *GetQuotaUsageResponseBodyData) *GetQuotaUsageResponseBody {
	s.Data = v
	return s
}

func (s *GetQuotaUsageResponseBody) SetErrorCode(v string) *GetQuotaUsageResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetQuotaUsageResponseBody) SetErrorMsg(v string) *GetQuotaUsageResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetQuotaUsageResponseBody) SetHttpCode(v int32) *GetQuotaUsageResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetQuotaUsageResponseBody) SetRequestId(v string) *GetQuotaUsageResponseBody {
	s.RequestId = &v
	return s
}

type GetQuotaUsageResponseBodyData struct {
	// The metric results.
	Metrics map[string]interface{} `json:"metrics,omitempty" xml:"metrics,omitempty"`
	// The information about the chart.
	Plot []*GetQuotaUsageResponseBodyDataPlot `json:"plot,omitempty" xml:"plot,omitempty" type:"Repeated"`
}

func (s GetQuotaUsageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaUsageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetQuotaUsageResponseBodyData) SetMetrics(v map[string]interface{}) *GetQuotaUsageResponseBodyData {
	s.Metrics = v
	return s
}

func (s *GetQuotaUsageResponseBodyData) SetPlot(v []*GetQuotaUsageResponseBodyDataPlot) *GetQuotaUsageResponseBodyData {
	s.Plot = v
	return s
}

type GetQuotaUsageResponseBodyDataPlot struct {
	// The title of the chart.
	//
	// example:
	//
	// request
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// The type of the chart.
	//
	// example:
	//
	// request
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The data metric field.
	YAxis []*string `json:"yAxis,omitempty" xml:"yAxis,omitempty" type:"Repeated"`
}

func (s GetQuotaUsageResponseBodyDataPlot) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaUsageResponseBodyDataPlot) GoString() string {
	return s.String()
}

func (s *GetQuotaUsageResponseBodyDataPlot) SetTitle(v string) *GetQuotaUsageResponseBodyDataPlot {
	s.Title = &v
	return s
}

func (s *GetQuotaUsageResponseBodyDataPlot) SetType(v string) *GetQuotaUsageResponseBodyDataPlot {
	s.Type = &v
	return s
}

func (s *GetQuotaUsageResponseBodyDataPlot) SetYAxis(v []*string) *GetQuotaUsageResponseBodyDataPlot {
	s.YAxis = v
	return s
}

type GetQuotaUsageResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQuotaUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQuotaUsageResponse) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaUsageResponse) GoString() string {
	return s.String()
}

func (s *GetQuotaUsageResponse) SetHeaders(v map[string]*string) *GetQuotaUsageResponse {
	s.Headers = v
	return s
}

func (s *GetQuotaUsageResponse) SetStatusCode(v int32) *GetQuotaUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQuotaUsageResponse) SetBody(v *GetQuotaUsageResponseBody) *GetQuotaUsageResponse {
	s.Body = v
	return s
}

type GetRoleAclResponseBody struct {
	// The returned data.
	Data *GetRoleAclResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// 040002
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// error message
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0a06dc0916696898838762018e9564
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetRoleAclResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRoleAclResponseBody) GoString() string {
	return s.String()
}

func (s *GetRoleAclResponseBody) SetData(v *GetRoleAclResponseBodyData) *GetRoleAclResponseBody {
	s.Data = v
	return s
}

func (s *GetRoleAclResponseBody) SetErrorCode(v string) *GetRoleAclResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetRoleAclResponseBody) SetErrorMsg(v string) *GetRoleAclResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetRoleAclResponseBody) SetHttpCode(v int32) *GetRoleAclResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetRoleAclResponseBody) SetRequestId(v string) *GetRoleAclResponseBody {
	s.RequestId = &v
	return s
}

type GetRoleAclResponseBodyData struct {
	// The function.
	Function []*GetRoleAclResponseBodyDataFunction `json:"function,omitempty" xml:"function,omitempty" type:"Repeated"`
	// The instance.
	Instance []*GetRoleAclResponseBodyDataInstance `json:"instance,omitempty" xml:"instance,omitempty" type:"Repeated"`
	// The package.
	Package []*GetRoleAclResponseBodyDataPackage `json:"package,omitempty" xml:"package,omitempty" type:"Repeated"`
	// The project.
	Project []*GetRoleAclResponseBodyDataProject `json:"project,omitempty" xml:"project,omitempty" type:"Repeated"`
	// The resource.
	Resource []*GetRoleAclResponseBodyDataResource `json:"resource,omitempty" xml:"resource,omitempty" type:"Repeated"`
	// The table.
	Table []*GetRoleAclResponseBodyDataTable `json:"table,omitempty" xml:"table,omitempty" type:"Repeated"`
}

func (s GetRoleAclResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetRoleAclResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRoleAclResponseBodyData) SetFunction(v []*GetRoleAclResponseBodyDataFunction) *GetRoleAclResponseBodyData {
	s.Function = v
	return s
}

func (s *GetRoleAclResponseBodyData) SetInstance(v []*GetRoleAclResponseBodyDataInstance) *GetRoleAclResponseBodyData {
	s.Instance = v
	return s
}

func (s *GetRoleAclResponseBodyData) SetPackage(v []*GetRoleAclResponseBodyDataPackage) *GetRoleAclResponseBodyData {
	s.Package = v
	return s
}

func (s *GetRoleAclResponseBodyData) SetProject(v []*GetRoleAclResponseBodyDataProject) *GetRoleAclResponseBodyData {
	s.Project = v
	return s
}

func (s *GetRoleAclResponseBodyData) SetResource(v []*GetRoleAclResponseBodyDataResource) *GetRoleAclResponseBodyData {
	s.Resource = v
	return s
}

func (s *GetRoleAclResponseBodyData) SetTable(v []*GetRoleAclResponseBodyDataTable) *GetRoleAclResponseBodyData {
	s.Table = v
	return s
}

type GetRoleAclResponseBodyDataFunction struct {
	// The operations that were performed on the function.
	Actions []*string `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	// The name of the function.
	//
	// example:
	//
	// functionA
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The Schema name.
	//
	// example:
	//
	// default
	SchemaName *string `json:"schemaName,omitempty" xml:"schemaName,omitempty"`
}

func (s GetRoleAclResponseBodyDataFunction) String() string {
	return tea.Prettify(s)
}

func (s GetRoleAclResponseBodyDataFunction) GoString() string {
	return s.String()
}

func (s *GetRoleAclResponseBodyDataFunction) SetActions(v []*string) *GetRoleAclResponseBodyDataFunction {
	s.Actions = v
	return s
}

func (s *GetRoleAclResponseBodyDataFunction) SetName(v string) *GetRoleAclResponseBodyDataFunction {
	s.Name = &v
	return s
}

func (s *GetRoleAclResponseBodyDataFunction) SetSchemaName(v string) *GetRoleAclResponseBodyDataFunction {
	s.SchemaName = &v
	return s
}

type GetRoleAclResponseBodyDataInstance struct {
	// The operations that were performed on the instance.
	Actions []*string `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	// The name of the instance.
	//
	// example:
	//
	// instanceA
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The Schema name.
	//
	// example:
	//
	// default
	SchemaName *string `json:"schemaName,omitempty" xml:"schemaName,omitempty"`
}

func (s GetRoleAclResponseBodyDataInstance) String() string {
	return tea.Prettify(s)
}

func (s GetRoleAclResponseBodyDataInstance) GoString() string {
	return s.String()
}

func (s *GetRoleAclResponseBodyDataInstance) SetActions(v []*string) *GetRoleAclResponseBodyDataInstance {
	s.Actions = v
	return s
}

func (s *GetRoleAclResponseBodyDataInstance) SetName(v string) *GetRoleAclResponseBodyDataInstance {
	s.Name = &v
	return s
}

func (s *GetRoleAclResponseBodyDataInstance) SetSchemaName(v string) *GetRoleAclResponseBodyDataInstance {
	s.SchemaName = &v
	return s
}

type GetRoleAclResponseBodyDataPackage struct {
	// The operations that were performed on the package.
	Actions []*string `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	// The name of the package.
	//
	// example:
	//
	// packageA
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The Schema name.
	//
	// example:
	//
	// default
	SchemaName *string `json:"schemaName,omitempty" xml:"schemaName,omitempty"`
}

func (s GetRoleAclResponseBodyDataPackage) String() string {
	return tea.Prettify(s)
}

func (s GetRoleAclResponseBodyDataPackage) GoString() string {
	return s.String()
}

func (s *GetRoleAclResponseBodyDataPackage) SetActions(v []*string) *GetRoleAclResponseBodyDataPackage {
	s.Actions = v
	return s
}

func (s *GetRoleAclResponseBodyDataPackage) SetName(v string) *GetRoleAclResponseBodyDataPackage {
	s.Name = &v
	return s
}

func (s *GetRoleAclResponseBodyDataPackage) SetSchemaName(v string) *GetRoleAclResponseBodyDataPackage {
	s.SchemaName = &v
	return s
}

type GetRoleAclResponseBodyDataProject struct {
	// The operations that were performed on the project.
	Actions []*string `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	// The name of the MaxCompute project.
	//
	// example:
	//
	// projectA
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The Schema name.
	//
	// example:
	//
	// default
	SchemaName *string `json:"schemaName,omitempty" xml:"schemaName,omitempty"`
}

func (s GetRoleAclResponseBodyDataProject) String() string {
	return tea.Prettify(s)
}

func (s GetRoleAclResponseBodyDataProject) GoString() string {
	return s.String()
}

func (s *GetRoleAclResponseBodyDataProject) SetActions(v []*string) *GetRoleAclResponseBodyDataProject {
	s.Actions = v
	return s
}

func (s *GetRoleAclResponseBodyDataProject) SetName(v string) *GetRoleAclResponseBodyDataProject {
	s.Name = &v
	return s
}

func (s *GetRoleAclResponseBodyDataProject) SetSchemaName(v string) *GetRoleAclResponseBodyDataProject {
	s.SchemaName = &v
	return s
}

type GetRoleAclResponseBodyDataResource struct {
	// The operations that were performed on the resource.
	Actions []*string `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	// The name of the resource.
	//
	// example:
	//
	// resourceA
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The Schema name.
	//
	// example:
	//
	// default
	SchemaName *string `json:"schemaName,omitempty" xml:"schemaName,omitempty"`
}

func (s GetRoleAclResponseBodyDataResource) String() string {
	return tea.Prettify(s)
}

func (s GetRoleAclResponseBodyDataResource) GoString() string {
	return s.String()
}

func (s *GetRoleAclResponseBodyDataResource) SetActions(v []*string) *GetRoleAclResponseBodyDataResource {
	s.Actions = v
	return s
}

func (s *GetRoleAclResponseBodyDataResource) SetName(v string) *GetRoleAclResponseBodyDataResource {
	s.Name = &v
	return s
}

func (s *GetRoleAclResponseBodyDataResource) SetSchemaName(v string) *GetRoleAclResponseBodyDataResource {
	s.SchemaName = &v
	return s
}

type GetRoleAclResponseBodyDataTable struct {
	// The operations that were performed on the table.
	Actions []*string `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	// The name of the table.
	//
	// example:
	//
	// tableA
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The Schema name.
	//
	// example:
	//
	// default
	SchemaName *string `json:"schemaName,omitempty" xml:"schemaName,omitempty"`
}

func (s GetRoleAclResponseBodyDataTable) String() string {
	return tea.Prettify(s)
}

func (s GetRoleAclResponseBodyDataTable) GoString() string {
	return s.String()
}

func (s *GetRoleAclResponseBodyDataTable) SetActions(v []*string) *GetRoleAclResponseBodyDataTable {
	s.Actions = v
	return s
}

func (s *GetRoleAclResponseBodyDataTable) SetName(v string) *GetRoleAclResponseBodyDataTable {
	s.Name = &v
	return s
}

func (s *GetRoleAclResponseBodyDataTable) SetSchemaName(v string) *GetRoleAclResponseBodyDataTable {
	s.SchemaName = &v
	return s
}

type GetRoleAclResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRoleAclResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRoleAclResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRoleAclResponse) GoString() string {
	return s.String()
}

func (s *GetRoleAclResponse) SetHeaders(v map[string]*string) *GetRoleAclResponse {
	s.Headers = v
	return s
}

func (s *GetRoleAclResponse) SetStatusCode(v int32) *GetRoleAclResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRoleAclResponse) SetBody(v *GetRoleAclResponseBody) *GetRoleAclResponse {
	s.Body = v
	return s
}

type GetRoleAclOnObjectRequest struct {
	// The name of the object.
	//
	// This parameter is required.
	//
	// example:
	//
	// tableA
	ObjectName *string `json:"objectName,omitempty" xml:"objectName,omitempty"`
	// The type of the object.
	//
	// This parameter is required.
	//
	// example:
	//
	// table
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
}

func (s GetRoleAclOnObjectRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRoleAclOnObjectRequest) GoString() string {
	return s.String()
}

func (s *GetRoleAclOnObjectRequest) SetObjectName(v string) *GetRoleAclOnObjectRequest {
	s.ObjectName = &v
	return s
}

func (s *GetRoleAclOnObjectRequest) SetObjectType(v string) *GetRoleAclOnObjectRequest {
	s.ObjectType = &v
	return s
}

type GetRoleAclOnObjectResponseBody struct {
	// The returned data
	Data *GetRoleAclOnObjectResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 0bc1366d16686529650188023ef87f
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetRoleAclOnObjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRoleAclOnObjectResponseBody) GoString() string {
	return s.String()
}

func (s *GetRoleAclOnObjectResponseBody) SetData(v *GetRoleAclOnObjectResponseBodyData) *GetRoleAclOnObjectResponseBody {
	s.Data = v
	return s
}

func (s *GetRoleAclOnObjectResponseBody) SetRequestId(v string) *GetRoleAclOnObjectResponseBody {
	s.RequestId = &v
	return s
}

type GetRoleAclOnObjectResponseBodyData struct {
	// The operations that were performed on the object.
	Actions []*string `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
}

func (s GetRoleAclOnObjectResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetRoleAclOnObjectResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRoleAclOnObjectResponseBodyData) SetActions(v []*string) *GetRoleAclOnObjectResponseBodyData {
	s.Actions = v
	return s
}

type GetRoleAclOnObjectResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRoleAclOnObjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRoleAclOnObjectResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRoleAclOnObjectResponse) GoString() string {
	return s.String()
}

func (s *GetRoleAclOnObjectResponse) SetHeaders(v map[string]*string) *GetRoleAclOnObjectResponse {
	s.Headers = v
	return s
}

func (s *GetRoleAclOnObjectResponse) SetStatusCode(v int32) *GetRoleAclOnObjectResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRoleAclOnObjectResponse) SetBody(v *GetRoleAclOnObjectResponseBody) *GetRoleAclOnObjectResponse {
	s.Body = v
	return s
}

type GetRolePolicyResponseBody struct {
	// The returned data.
	//
	// example:
	//
	// {
	//
	//       "Statement": [
	//
	//             {
	//
	//                   "Action": [
	//
	//                         "odps:*"
	//
	//                   ],
	//
	//                   "Effect": "Allow",
	//
	//                   "Resource": [
	//
	//                         "acs:odps:*:projects/{projectname}/authorization/packages"
	//
	//                   ]
	//
	//             }
	//
	//       ],
	//
	//       "Version": "1"
	//
	// }
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0bc1eeed16675342848904412e08dd
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetRolePolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRolePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GetRolePolicyResponseBody) SetData(v string) *GetRolePolicyResponseBody {
	s.Data = &v
	return s
}

func (s *GetRolePolicyResponseBody) SetRequestId(v string) *GetRolePolicyResponseBody {
	s.RequestId = &v
	return s
}

type GetRolePolicyResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRolePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRolePolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRolePolicyResponse) GoString() string {
	return s.String()
}

func (s *GetRolePolicyResponse) SetHeaders(v map[string]*string) *GetRolePolicyResponse {
	s.Headers = v
	return s
}

func (s *GetRolePolicyResponse) SetStatusCode(v int32) *GetRolePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRolePolicyResponse) SetBody(v *GetRolePolicyResponseBody) *GetRolePolicyResponse {
	s.Body = v
	return s
}

type GetRunningJobsRequest struct {
	// The time when the query starts. This parameter specifies the time when a job is submitted.
	//
	// 	- The time range that is specified by the **from*	- and **to*	- request parameters is a closed interval. The start time and end time are included in the range. If the value of **from*	- is the same as the value of **to**, the time range is invalid, and a null value is returned.
	//
	// 	- The value is a UNIX timestamp that represents the number of seconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1683785928
	From *int64 `json:"from,omitempty" xml:"from,omitempty"`
	// The list of job executors.
	JobOwnerList []*string `json:"jobOwnerList,omitempty" xml:"jobOwnerList,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 20.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The list of nicknames of quotas that are used by jobs.
	QuotaNicknameList []*string `json:"quotaNicknameList,omitempty" xml:"quotaNicknameList,omitempty" type:"Repeated"`
	// The time when the query ends. This parameter specifies the time when a job is submitted.
	//
	// 	- The time interval that is specified by the **from*	- and **to*	- request parameters is a closed interval. The start time and end time are included in the interval. If the value of **from*	- is the same as the value of **to**, the interval is invalid, and a null value is returned.
	//
	// 	- The value is a UNIX timestamp that represents the number of seconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1683612946
	To *int64 `json:"to,omitempty" xml:"to,omitempty"`
}

func (s GetRunningJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRunningJobsRequest) GoString() string {
	return s.String()
}

func (s *GetRunningJobsRequest) SetFrom(v int64) *GetRunningJobsRequest {
	s.From = &v
	return s
}

func (s *GetRunningJobsRequest) SetJobOwnerList(v []*string) *GetRunningJobsRequest {
	s.JobOwnerList = v
	return s
}

func (s *GetRunningJobsRequest) SetPageNumber(v int64) *GetRunningJobsRequest {
	s.PageNumber = &v
	return s
}

func (s *GetRunningJobsRequest) SetPageSize(v int64) *GetRunningJobsRequest {
	s.PageSize = &v
	return s
}

func (s *GetRunningJobsRequest) SetQuotaNicknameList(v []*string) *GetRunningJobsRequest {
	s.QuotaNicknameList = v
	return s
}

func (s *GetRunningJobsRequest) SetTo(v int64) *GetRunningJobsRequest {
	s.To = &v
	return s
}

type GetRunningJobsShrinkRequest struct {
	// The time when the query starts. This parameter specifies the time when a job is submitted.
	//
	// 	- The time range that is specified by the **from*	- and **to*	- request parameters is a closed interval. The start time and end time are included in the range. If the value of **from*	- is the same as the value of **to**, the time range is invalid, and a null value is returned.
	//
	// 	- The value is a UNIX timestamp that represents the number of seconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1683785928
	From *int64 `json:"from,omitempty" xml:"from,omitempty"`
	// The list of job executors.
	JobOwnerListShrink *string `json:"jobOwnerList,omitempty" xml:"jobOwnerList,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 20.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The list of nicknames of quotas that are used by jobs.
	QuotaNicknameListShrink *string `json:"quotaNicknameList,omitempty" xml:"quotaNicknameList,omitempty"`
	// The time when the query ends. This parameter specifies the time when a job is submitted.
	//
	// 	- The time interval that is specified by the **from*	- and **to*	- request parameters is a closed interval. The start time and end time are included in the interval. If the value of **from*	- is the same as the value of **to**, the interval is invalid, and a null value is returned.
	//
	// 	- The value is a UNIX timestamp that represents the number of seconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1683612946
	To *int64 `json:"to,omitempty" xml:"to,omitempty"`
}

func (s GetRunningJobsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRunningJobsShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetRunningJobsShrinkRequest) SetFrom(v int64) *GetRunningJobsShrinkRequest {
	s.From = &v
	return s
}

func (s *GetRunningJobsShrinkRequest) SetJobOwnerListShrink(v string) *GetRunningJobsShrinkRequest {
	s.JobOwnerListShrink = &v
	return s
}

func (s *GetRunningJobsShrinkRequest) SetPageNumber(v int64) *GetRunningJobsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *GetRunningJobsShrinkRequest) SetPageSize(v int64) *GetRunningJobsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GetRunningJobsShrinkRequest) SetQuotaNicknameListShrink(v string) *GetRunningJobsShrinkRequest {
	s.QuotaNicknameListShrink = &v
	return s
}

func (s *GetRunningJobsShrinkRequest) SetTo(v int64) *GetRunningJobsShrinkRequest {
	s.To = &v
	return s
}

type GetRunningJobsResponseBody struct {
	// The returned data.
	Data *GetRunningJobsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// 	- If the value of success was false, an error code was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// success
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// 0A3B1FD2006A24C8D8BE65CDAC028298
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// Indicates whether the request was successful. If this parameter was not empty and the value of this parameter was not 200, the request failed.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0bc3b4ab16684833172127321e2c25
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetRunningJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRunningJobsResponseBody) GoString() string {
	return s.String()
}

func (s *GetRunningJobsResponseBody) SetData(v *GetRunningJobsResponseBodyData) *GetRunningJobsResponseBody {
	s.Data = v
	return s
}

func (s *GetRunningJobsResponseBody) SetErrorCode(v string) *GetRunningJobsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetRunningJobsResponseBody) SetErrorMsg(v string) *GetRunningJobsResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetRunningJobsResponseBody) SetHttpCode(v int32) *GetRunningJobsResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetRunningJobsResponseBody) SetRequestId(v string) *GetRunningJobsResponseBody {
	s.RequestId = &v
	return s
}

type GetRunningJobsResponseBodyData struct {
	// The page number.
	//
	// example:
	//
	// 2
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The list of jobs in the running state.
	RunningJobInfoList []*GetRunningJobsResponseBodyDataRunningJobInfoList `json:"runningJobInfoList,omitempty" xml:"runningJobInfoList,omitempty" type:"Repeated"`
	// The total number of returned entries.
	//
	// example:
	//
	// 64
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s GetRunningJobsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetRunningJobsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRunningJobsResponseBodyData) SetPageNumber(v int64) *GetRunningJobsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *GetRunningJobsResponseBodyData) SetPageSize(v int64) *GetRunningJobsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetRunningJobsResponseBodyData) SetRunningJobInfoList(v []*GetRunningJobsResponseBodyDataRunningJobInfoList) *GetRunningJobsResponseBodyData {
	s.RunningJobInfoList = v
	return s
}

func (s *GetRunningJobsResponseBodyData) SetTotalCount(v int64) *GetRunningJobsResponseBodyData {
	s.TotalCount = &v
	return s
}

type GetRunningJobsResponseBodyDataRunningJobInfoList struct {
	// The compute unit (CU) snapshot proportion of the job.
	//
	// example:
	//
	// 0.45
	CuSnapshot *float64 `json:"cuSnapshot,omitempty" xml:"cuSnapshot,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// 2023050206371544gomgtp3ljcr4
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The account that submits the job.
	//
	// example:
	//
	// ALIYUN$xxx@test.aliyunid.com
	JobOwner *string `json:"jobOwner,omitempty" xml:"jobOwner,omitempty"`
	// The memory snapshot proportion of the job.
	//
	// example:
	//
	// 0.45
	MemorySnapshot *float64 `json:"memorySnapshot,omitempty" xml:"memorySnapshot,omitempty"`
	// The progress of the job.
	//
	// example:
	//
	// 0
	Progress *float64 `json:"progress,omitempty" xml:"progress,omitempty"`
	// The name of the MaxCompute project.
	//
	// example:
	//
	// project_20221021123044_981b
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
	// The nickname of the quota that is used by the job.
	//
	// example:
	//
	// my_quota
	QuotaNickname *string `json:"quotaNickname,omitempty" xml:"quotaNickname,omitempty"`
	// The time when the job starts to run.
	//
	// example:
	//
	// 1689746864
	RunningAtTime *int64 `json:"runningAtTime,omitempty" xml:"runningAtTime,omitempty"`
	// The time when the job is submitted.
	//
	// example:
	//
	// 1689746864
	SubmittedAtTime *int64 `json:"submittedAtTime,omitempty" xml:"submittedAtTime,omitempty"`
}

func (s GetRunningJobsResponseBodyDataRunningJobInfoList) String() string {
	return tea.Prettify(s)
}

func (s GetRunningJobsResponseBodyDataRunningJobInfoList) GoString() string {
	return s.String()
}

func (s *GetRunningJobsResponseBodyDataRunningJobInfoList) SetCuSnapshot(v float64) *GetRunningJobsResponseBodyDataRunningJobInfoList {
	s.CuSnapshot = &v
	return s
}

func (s *GetRunningJobsResponseBodyDataRunningJobInfoList) SetInstanceId(v string) *GetRunningJobsResponseBodyDataRunningJobInfoList {
	s.InstanceId = &v
	return s
}

func (s *GetRunningJobsResponseBodyDataRunningJobInfoList) SetJobOwner(v string) *GetRunningJobsResponseBodyDataRunningJobInfoList {
	s.JobOwner = &v
	return s
}

func (s *GetRunningJobsResponseBodyDataRunningJobInfoList) SetMemorySnapshot(v float64) *GetRunningJobsResponseBodyDataRunningJobInfoList {
	s.MemorySnapshot = &v
	return s
}

func (s *GetRunningJobsResponseBodyDataRunningJobInfoList) SetProgress(v float64) *GetRunningJobsResponseBodyDataRunningJobInfoList {
	s.Progress = &v
	return s
}

func (s *GetRunningJobsResponseBodyDataRunningJobInfoList) SetProject(v string) *GetRunningJobsResponseBodyDataRunningJobInfoList {
	s.Project = &v
	return s
}

func (s *GetRunningJobsResponseBodyDataRunningJobInfoList) SetQuotaNickname(v string) *GetRunningJobsResponseBodyDataRunningJobInfoList {
	s.QuotaNickname = &v
	return s
}

func (s *GetRunningJobsResponseBodyDataRunningJobInfoList) SetRunningAtTime(v int64) *GetRunningJobsResponseBodyDataRunningJobInfoList {
	s.RunningAtTime = &v
	return s
}

func (s *GetRunningJobsResponseBodyDataRunningJobInfoList) SetSubmittedAtTime(v int64) *GetRunningJobsResponseBodyDataRunningJobInfoList {
	s.SubmittedAtTime = &v
	return s
}

type GetRunningJobsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRunningJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRunningJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRunningJobsResponse) GoString() string {
	return s.String()
}

func (s *GetRunningJobsResponse) SetHeaders(v map[string]*string) *GetRunningJobsResponse {
	s.Headers = v
	return s
}

func (s *GetRunningJobsResponse) SetStatusCode(v int32) *GetRunningJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRunningJobsResponse) SetBody(v *GetRunningJobsResponseBody) *GetRunningJobsResponse {
	s.Body = v
	return s
}

type GetTableInfoRequest struct {
	// The name of the schema to which the table or view belongs.
	//
	// example:
	//
	// default
	SchemaName *string `json:"schemaName,omitempty" xml:"schemaName,omitempty"`
	// The type of the table or view that you want to view. Valid values:
	//
	// 	- **internal**: internal table
	//
	// 	- **external**: external table
	//
	// 	- **view**: view
	//
	// 	- **materializedView**: [materialize view](https://www.alibabacloud.com/help/maxcompute/user-guide/materialized-view-operations)
	//
	// example:
	//
	// internal
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetTableInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTableInfoRequest) GoString() string {
	return s.String()
}

func (s *GetTableInfoRequest) SetSchemaName(v string) *GetTableInfoRequest {
	s.SchemaName = &v
	return s
}

func (s *GetTableInfoRequest) SetType(v string) *GetTableInfoRequest {
	s.Type = &v
	return s
}

type GetTableInfoResponseBody struct {
	// The data returned.
	Data *GetTableInfoResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0a06dd4516687375802853481ec9fd
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetTableInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTableInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetTableInfoResponseBody) SetData(v *GetTableInfoResponseBodyData) *GetTableInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetTableInfoResponseBody) SetRequestId(v string) *GetTableInfoResponseBody {
	s.RequestId = &v
	return s
}

type GetTableInfoResponseBodyData struct {
	// Indicates whether the materialized view is automatically refreshed. This response parameter is returned when type is set to materializedView.
	//
	// example:
	//
	// false
	AutoRefreshEnabled *bool `json:"autoRefreshEnabled,omitempty" xml:"autoRefreshEnabled,omitempty"`
	// The clustering attribute. This response parameter is returned when the table is a clustered table.
	ClusterInfo *GetTableInfoResponseBodyDataClusterInfo `json:"clusterInfo,omitempty" xml:"clusterInfo,omitempty" type:"Struct"`
	// The comments of the table.
	//
	// example:
	//
	// sale_detail
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	// DDL statement to create a table.
	//
	// example:
	//
	// create table if not exists sale_detail( shop_name STRING, customer_id STRING, total_price DOUBLE) partitioned by (sale_date STRING, region STRING);
	CreateTableDDL *string `json:"createTableDDL,omitempty" xml:"createTableDDL,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2022-01-17T07:07:47Z
	CreationTime *int64 `json:"creationTime,omitempty" xml:"creationTime,omitempty"`
	// The display name.
	//
	// example:
	//
	// project_name.schema_name.table_name
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// The number of file of the table.
	//
	// example:
	//
	// 200
	FileNum *int64 `json:"fileNum,omitempty" xml:"fileNum,omitempty"`
	// Indicates whether the table is an external table. This response parameter is returned when type is set to external.
	//
	// example:
	//
	// false
	IsExternalTable *bool `json:"isExternalTable,omitempty" xml:"isExternalTable,omitempty"`
	// Indicates whether data of the materialized view is invalid due to changes in the data of the source table. This response parameter is returned when type is set to materializedView.
	//
	// example:
	//
	// false
	IsOutdated *bool `json:"isOutdated,omitempty" xml:"isOutdated,omitempty"`
	// The time when data of the table or view was last accessed.
	//
	// example:
	//
	// 2023-11-21T02:05:56Z
	LastAccessTime *int64 `json:"lastAccessTime,omitempty" xml:"lastAccessTime,omitempty"`
	// The time when the data definition language (DDL) statement of the table or view was last modified.
	//
	// example:
	//
	// 2023-11-21T02:05:56Z
	LastDDLTime *int64 `json:"lastDDLTime,omitempty" xml:"lastDDLTime,omitempty"`
	// The time when data of the table or view was last modified.
	//
	// example:
	//
	// 2023-11-21T02:05:56Z
	LastModifiedTime *int64 `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
	// The lifecycle. Unit: days.
	//
	// example:
	//
	// -1
	Lifecycle *string `json:"lifecycle,omitempty" xml:"lifecycle,omitempty"`
	// The path of the external table. This response parameter is returned when type is set to external.
	//
	// example:
	//
	// oss://oss-cn-hangzhou-internal.aliyuncs.com/oss-mc-test/Demo1/
	Location *string `json:"location,omitempty" xml:"location,omitempty"`
	// Indicates whether the table or view is a [materialize view](https://www.alibabacloud.com/help/maxcompute/user-guide/materialized-view-operations).
	//
	// example:
	//
	// false
	MaterializedView *bool `json:"materializedView,omitempty" xml:"materializedView,omitempty"`
	// The name of the table or view.
	//
	// example:
	//
	// sale_detail
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The field information.
	NativeColumns []*GetTableInfoResponseBodyDataNativeColumns `json:"nativeColumns,omitempty" xml:"nativeColumns,omitempty" type:"Repeated"`
	// The Alibaba Cloud Resource Name (ARN) of the role AliyunODPSDefaultRole in Resource Access Management (RAM). This response parameter is returned when type is set to external.
	//
	// example:
	//
	// acs:ram::xxxxx:role/aliyunodpsdefaultrole
	OdpsPropertiesRolearn *string `json:"odpsPropertiesRolearn,omitempty" xml:"odpsPropertiesRolearn,omitempty"`
	// Indicates whether the table header is ignored. This response parameter is returned when type is set to external.
	//
	// example:
	//
	// true
	OdpsSqlTextOptionFlushHeader *bool `json:"odpsSqlTextOptionFlushHeader,omitempty" xml:"odpsSqlTextOptionFlushHeader,omitempty"`
	// The first N rows that were ignored in the table header. This response parameter is returned when type is set to external.
	//
	// example:
	//
	// 1
	OdpsTextOptionHeaderLinesCount *int64 `json:"odpsTextOptionHeaderLinesCount,omitempty" xml:"odpsTextOptionHeaderLinesCount,omitempty"`
	// The account information of the table or view owner.
	//
	// example:
	//
	// 188785396123****
	Owner *string `json:"owner,omitempty" xml:"owner,omitempty"`
	// The information about partition key columns. This response parameter is returned only for partitioned tables.
	PartitionColumns []*GetTableInfoResponseBodyDataPartitionColumns `json:"partitionColumns,omitempty" xml:"partitionColumns,omitempty" type:"Repeated"`
	// The physical size of the table.
	//
	// example:
	//
	// 2763
	PhysicalSize *int64 `json:"physicalSize,omitempty" xml:"physicalSize,omitempty"`
	// The name of the project to which the table or view belongs.
	//
	// example:
	//
	// projectA
	ProjectName *string `json:"projectName,omitempty" xml:"projectName,omitempty"`
	// Indicates whether the query rewrite operation can be performed by using the materialized view. This response parameter is returned when type is set to materializedView.
	//
	// example:
	//
	// false
	RewriteEnabled *bool `json:"rewriteEnabled,omitempty" xml:"rewriteEnabled,omitempty"`
	// The name of the schema to which the table or the view belongs.
	//
	// example:
	//
	// default
	Schema *string `json:"schema,omitempty" xml:"schema,omitempty"`
	// The data size of the non-partitioned table. If the table is a partitioned table, the system does not calculate the data size of the table. In this case, the value of this parameter is NULL. The PARTITIONS view includes the data size of each partition in a partitioned table. Unit: bytes.
	//
	// example:
	//
	// 5372
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// The storage handler of the external table. This response parameter is returned when type is set to external.
	//
	// example:
	//
	// com.aliyun.odps.CsvStorageHandler
	StorageHandler *string `json:"storageHandler,omitempty" xml:"storageHandler,omitempty"`
	// The sensitivity-level label of the table. For more information, see [Label-based access control](https://www.alibabacloud.com/help/maxcompute/user-guide/label-based-access-control).
	//
	// example:
	//
	// 0
	TableLabel *string `json:"tableLabel,omitempty" xml:"tableLabel,omitempty"`
	// The name of the Tablestore table to be accessed. This response parameter is returned when type is set to external.
	//
	// example:
	//
	// ots_tpch_orders
	TablesotreTableName *string `json:"tablesotreTableName,omitempty" xml:"tablesotreTableName,omitempty"`
	// The columns of the Tablestore table to be accessed, including the primary key column and attribute column. This response parameter is returned when type is set to external.
	//
	// example:
	//
	// :o_orderkey,:o_orderdate,o_custkey,o_orderstatus,o_totalprice
	TablestoreColumnsMapping *string `json:"tablestoreColumnsMapping,omitempty" xml:"tablestoreColumnsMapping,omitempty"`
	// The type of the table or view. Valid values:
	//
	// 	- **internal**: internal table
	//
	// 	- **external**: external table
	//
	// 	- **view**: view
	//
	// 	- **materializedView**: [materialize view](https://www.alibabacloud.com/help/maxcompute/user-guide/materialized-view-operations)
	//
	// example:
	//
	// internal
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The statement that generates the view. This response parameter is returned when type is set to view.
	//
	// example:
	//
	// select shop_name, sum(total_price) from sale_detail group by shop_name
	ViewText *string `json:"viewText,omitempty" xml:"viewText,omitempty"`
}

func (s GetTableInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTableInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTableInfoResponseBodyData) SetAutoRefreshEnabled(v bool) *GetTableInfoResponseBodyData {
	s.AutoRefreshEnabled = &v
	return s
}

func (s *GetTableInfoResponseBodyData) SetClusterInfo(v *GetTableInfoResponseBodyDataClusterInfo) *GetTableInfoResponseBodyData {
	s.ClusterInfo = v
	return s
}

func (s *GetTableInfoResponseBodyData) SetComment(v string) *GetTableInfoResponseBodyData {
	s.Comment = &v
	return s
}

func (s *GetTableInfoResponseBodyData) SetCreateTableDDL(v string) *GetTableInfoResponseBodyData {
	s.CreateTableDDL = &v
	return s
}

func (s *GetTableInfoResponseBodyData) SetCreationTime(v int64) *GetTableInfoResponseBodyData {
	s.CreationTime = &v
	return s
}

func (s *GetTableInfoResponseBodyData) SetDisplayName(v string) *GetTableInfoResponseBodyData {
	s.DisplayName = &v
	return s
}

func (s *GetTableInfoResponseBodyData) SetFileNum(v int64) *GetTableInfoResponseBodyData {
	s.FileNum = &v
	return s
}

func (s *GetTableInfoResponseBodyData) SetIsExternalTable(v bool) *GetTableInfoResponseBodyData {
	s.IsExternalTable = &v
	return s
}

func (s *GetTableInfoResponseBodyData) SetIsOutdated(v bool) *GetTableInfoResponseBodyData {
	s.IsOutdated = &v
	return s
}

func (s *GetTableInfoResponseBodyData) SetLastAccessTime(v int64) *GetTableInfoResponseBodyData {
	s.LastAccessTime = &v
	return s
}

func (s *GetTableInfoResponseBodyData) SetLastDDLTime(v int64) *GetTableInfoResponseBodyData {
	s.LastDDLTime = &v
	return s
}

func (s *GetTableInfoResponseBodyData) SetLastModifiedTime(v int64) *GetTableInfoResponseBodyData {
	s.LastModifiedTime = &v
	return s
}

func (s *GetTableInfoResponseBodyData) SetLifecycle(v string) *GetTableInfoResponseBodyData {
	s.Lifecycle = &v
	return s
}

func (s *GetTableInfoResponseBodyData) SetLocation(v string) *GetTableInfoResponseBodyData {
	s.Location = &v
	return s
}

func (s *GetTableInfoResponseBodyData) SetMaterializedView(v bool) *GetTableInfoResponseBodyData {
	s.MaterializedView = &v
	return s
}

func (s *GetTableInfoResponseBodyData) SetName(v string) *GetTableInfoResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetTableInfoResponseBodyData) SetNativeColumns(v []*GetTableInfoResponseBodyDataNativeColumns) *GetTableInfoResponseBodyData {
	s.NativeColumns = v
	return s
}

func (s *GetTableInfoResponseBodyData) SetOdpsPropertiesRolearn(v string) *GetTableInfoResponseBodyData {
	s.OdpsPropertiesRolearn = &v
	return s
}

func (s *GetTableInfoResponseBodyData) SetOdpsSqlTextOptionFlushHeader(v bool) *GetTableInfoResponseBodyData {
	s.OdpsSqlTextOptionFlushHeader = &v
	return s
}

func (s *GetTableInfoResponseBodyData) SetOdpsTextOptionHeaderLinesCount(v int64) *GetTableInfoResponseBodyData {
	s.OdpsTextOptionHeaderLinesCount = &v
	return s
}

func (s *GetTableInfoResponseBodyData) SetOwner(v string) *GetTableInfoResponseBodyData {
	s.Owner = &v
	return s
}

func (s *GetTableInfoResponseBodyData) SetPartitionColumns(v []*GetTableInfoResponseBodyDataPartitionColumns) *GetTableInfoResponseBodyData {
	s.PartitionColumns = v
	return s
}

func (s *GetTableInfoResponseBodyData) SetPhysicalSize(v int64) *GetTableInfoResponseBodyData {
	s.PhysicalSize = &v
	return s
}

func (s *GetTableInfoResponseBodyData) SetProjectName(v string) *GetTableInfoResponseBodyData {
	s.ProjectName = &v
	return s
}

func (s *GetTableInfoResponseBodyData) SetRewriteEnabled(v bool) *GetTableInfoResponseBodyData {
	s.RewriteEnabled = &v
	return s
}

func (s *GetTableInfoResponseBodyData) SetSchema(v string) *GetTableInfoResponseBodyData {
	s.Schema = &v
	return s
}

func (s *GetTableInfoResponseBodyData) SetSize(v int64) *GetTableInfoResponseBodyData {
	s.Size = &v
	return s
}

func (s *GetTableInfoResponseBodyData) SetStorageHandler(v string) *GetTableInfoResponseBodyData {
	s.StorageHandler = &v
	return s
}

func (s *GetTableInfoResponseBodyData) SetTableLabel(v string) *GetTableInfoResponseBodyData {
	s.TableLabel = &v
	return s
}

func (s *GetTableInfoResponseBodyData) SetTablesotreTableName(v string) *GetTableInfoResponseBodyData {
	s.TablesotreTableName = &v
	return s
}

func (s *GetTableInfoResponseBodyData) SetTablestoreColumnsMapping(v string) *GetTableInfoResponseBodyData {
	s.TablestoreColumnsMapping = &v
	return s
}

func (s *GetTableInfoResponseBodyData) SetType(v string) *GetTableInfoResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetTableInfoResponseBodyData) SetViewText(v string) *GetTableInfoResponseBodyData {
	s.ViewText = &v
	return s
}

type GetTableInfoResponseBodyDataClusterInfo struct {
	// Optional. The number of buckets in the clustered table. The value 0 indicates that the number of buckets dynamically changes when a job is running.
	//
	// example:
	//
	// 1024
	BucketNum *int64 `json:"bucketNum,omitempty" xml:"bucketNum,omitempty"`
	// The cluster keys.
	ClusterCols []*string `json:"clusterCols,omitempty" xml:"clusterCols,omitempty" type:"Repeated"`
	// The clustering type of the table. MaxCompute supports [hash clustering](https://www.alibabacloud.com/help/maxcompute/use-cases/hash-clustering) and
	//
	// [range clustering](https://www.alibabacloud.com/help/maxcompute/use-cases/range-clustering).
	//
	// example:
	//
	// Hash
	ClusterType *string `json:"clusterType,omitempty" xml:"clusterType,omitempty"`
	// The condition by which the results are sorted.
	SortCols []*GetTableInfoResponseBodyDataClusterInfoSortCols `json:"sortCols,omitempty" xml:"sortCols,omitempty" type:"Repeated"`
}

func (s GetTableInfoResponseBodyDataClusterInfo) String() string {
	return tea.Prettify(s)
}

func (s GetTableInfoResponseBodyDataClusterInfo) GoString() string {
	return s.String()
}

func (s *GetTableInfoResponseBodyDataClusterInfo) SetBucketNum(v int64) *GetTableInfoResponseBodyDataClusterInfo {
	s.BucketNum = &v
	return s
}

func (s *GetTableInfoResponseBodyDataClusterInfo) SetClusterCols(v []*string) *GetTableInfoResponseBodyDataClusterInfo {
	s.ClusterCols = v
	return s
}

func (s *GetTableInfoResponseBodyDataClusterInfo) SetClusterType(v string) *GetTableInfoResponseBodyDataClusterInfo {
	s.ClusterType = &v
	return s
}

func (s *GetTableInfoResponseBodyDataClusterInfo) SetSortCols(v []*GetTableInfoResponseBodyDataClusterInfoSortCols) *GetTableInfoResponseBodyDataClusterInfo {
	s.SortCols = v
	return s
}

type GetTableInfoResponseBodyDataClusterInfoSortCols struct {
	// The name of the sorting field.
	//
	// example:
	//
	// col_2
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The sorting order.
	//
	// example:
	//
	// DESC
	Order *string `json:"order,omitempty" xml:"order,omitempty"`
}

func (s GetTableInfoResponseBodyDataClusterInfoSortCols) String() string {
	return tea.Prettify(s)
}

func (s GetTableInfoResponseBodyDataClusterInfoSortCols) GoString() string {
	return s.String()
}

func (s *GetTableInfoResponseBodyDataClusterInfoSortCols) SetName(v string) *GetTableInfoResponseBodyDataClusterInfoSortCols {
	s.Name = &v
	return s
}

func (s *GetTableInfoResponseBodyDataClusterInfoSortCols) SetOrder(v string) *GetTableInfoResponseBodyDataClusterInfoSortCols {
	s.Order = &v
	return s
}

type GetTableInfoResponseBodyDataNativeColumns struct {
	// The column comments.
	//
	// example:
	//
	// The name of shop.
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	// The sensitivity-level label of the column. For more information, see [Label-based access control](https://www.alibabacloud.com/help/maxcompute/user-guide/label-based-access-control).
	//
	// example:
	//
	// 0
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// The column name.
	//
	// example:
	//
	// shop_name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The column type.
	//
	// example:
	//
	// STRING
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetTableInfoResponseBodyDataNativeColumns) String() string {
	return tea.Prettify(s)
}

func (s GetTableInfoResponseBodyDataNativeColumns) GoString() string {
	return s.String()
}

func (s *GetTableInfoResponseBodyDataNativeColumns) SetComment(v string) *GetTableInfoResponseBodyDataNativeColumns {
	s.Comment = &v
	return s
}

func (s *GetTableInfoResponseBodyDataNativeColumns) SetLabel(v string) *GetTableInfoResponseBodyDataNativeColumns {
	s.Label = &v
	return s
}

func (s *GetTableInfoResponseBodyDataNativeColumns) SetName(v string) *GetTableInfoResponseBodyDataNativeColumns {
	s.Name = &v
	return s
}

func (s *GetTableInfoResponseBodyDataNativeColumns) SetType(v string) *GetTableInfoResponseBodyDataNativeColumns {
	s.Type = &v
	return s
}

type GetTableInfoResponseBodyDataPartitionColumns struct {
	// The comments of the partition key column.
	//
	// example:
	//
	// Sale date.
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	// The sensitivity-level label of the column. For more information, see [Label-based access control](https://www.alibabacloud.com/help/maxcompute/user-guide/label-based-access-control).
	//
	// example:
	//
	// 0
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// The partition name.
	//
	// example:
	//
	// sale_date
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The partition column type.
	//
	// example:
	//
	// STRING
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetTableInfoResponseBodyDataPartitionColumns) String() string {
	return tea.Prettify(s)
}

func (s GetTableInfoResponseBodyDataPartitionColumns) GoString() string {
	return s.String()
}

func (s *GetTableInfoResponseBodyDataPartitionColumns) SetComment(v string) *GetTableInfoResponseBodyDataPartitionColumns {
	s.Comment = &v
	return s
}

func (s *GetTableInfoResponseBodyDataPartitionColumns) SetLabel(v string) *GetTableInfoResponseBodyDataPartitionColumns {
	s.Label = &v
	return s
}

func (s *GetTableInfoResponseBodyDataPartitionColumns) SetName(v string) *GetTableInfoResponseBodyDataPartitionColumns {
	s.Name = &v
	return s
}

func (s *GetTableInfoResponseBodyDataPartitionColumns) SetType(v string) *GetTableInfoResponseBodyDataPartitionColumns {
	s.Type = &v
	return s
}

type GetTableInfoResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTableInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTableInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTableInfoResponse) GoString() string {
	return s.String()
}

func (s *GetTableInfoResponse) SetHeaders(v map[string]*string) *GetTableInfoResponse {
	s.Headers = v
	return s
}

func (s *GetTableInfoResponse) SetStatusCode(v int32) *GetTableInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTableInfoResponse) SetBody(v *GetTableInfoResponseBody) *GetTableInfoResponse {
	s.Body = v
	return s
}

type GetTrustedProjectsResponseBody struct {
	// The returned data.
	Data []*string `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 0bc0590416675329272834336e4387
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetTrustedProjectsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTrustedProjectsResponseBody) GoString() string {
	return s.String()
}

func (s *GetTrustedProjectsResponseBody) SetData(v []*string) *GetTrustedProjectsResponseBody {
	s.Data = v
	return s
}

func (s *GetTrustedProjectsResponseBody) SetRequestId(v string) *GetTrustedProjectsResponseBody {
	s.RequestId = &v
	return s
}

type GetTrustedProjectsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTrustedProjectsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTrustedProjectsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTrustedProjectsResponse) GoString() string {
	return s.String()
}

func (s *GetTrustedProjectsResponse) SetHeaders(v map[string]*string) *GetTrustedProjectsResponse {
	s.Headers = v
	return s
}

func (s *GetTrustedProjectsResponse) SetStatusCode(v int32) *GetTrustedProjectsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTrustedProjectsResponse) SetBody(v *GetTrustedProjectsResponseBody) *GetTrustedProjectsResponse {
	s.Body = v
	return s
}

type KillJobsRequest struct {
	// The request body parameters.
	//
	// example:
	//
	// [
	//
	//       {
	//
	//             "instanceId": "",
	//
	//             "projectName": ""
	//
	//       },
	//
	//       {
	//
	//             "instanceId": "",
	//
	//             "projectName": ""
	//
	//       }
	//
	// ]
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
	// The ID of the region in which the instance resides.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The ID of the tenant.
	//
	// example:
	//
	// 478403690625249
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s KillJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s KillJobsRequest) GoString() string {
	return s.String()
}

func (s *KillJobsRequest) SetBody(v string) *KillJobsRequest {
	s.Body = &v
	return s
}

func (s *KillJobsRequest) SetRegion(v string) *KillJobsRequest {
	s.Region = &v
	return s
}

func (s *KillJobsRequest) SetTenantId(v string) *KillJobsRequest {
	s.TenantId = &v
	return s
}

type KillJobsResponseBody struct {
	// The returned data.
	//
	// example:
	//
	// success
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// Indicates whether the request was successful. If this parameter was not empty and the value of this parameter was not 200, the request failed.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0abb7ede16814560741256732e91b6
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s KillJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s KillJobsResponseBody) GoString() string {
	return s.String()
}

func (s *KillJobsResponseBody) SetData(v string) *KillJobsResponseBody {
	s.Data = &v
	return s
}

func (s *KillJobsResponseBody) SetHttpCode(v int32) *KillJobsResponseBody {
	s.HttpCode = &v
	return s
}

func (s *KillJobsResponseBody) SetRequestId(v string) *KillJobsResponseBody {
	s.RequestId = &v
	return s
}

type KillJobsResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *KillJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s KillJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s KillJobsResponse) GoString() string {
	return s.String()
}

func (s *KillJobsResponse) SetHeaders(v map[string]*string) *KillJobsResponse {
	s.Headers = v
	return s
}

func (s *KillJobsResponse) SetStatusCode(v int32) *KillJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *KillJobsResponse) SetBody(v *KillJobsResponseBody) *KillJobsResponse {
	s.Body = v
	return s
}

type ListComputeMetricsByInstanceRequest struct {
	// The end time for the period.
	//
	// example:
	//
	// 1718590596556
	EndDate *int64 `json:"endDate,omitempty" xml:"endDate,omitempty"`
	// The job(instance) ID.
	//
	// example:
	//
	// 20240730****ddlr
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The Alibaba Cloud account that is used to run the MaxCompute job.
	//
	// example:
	//
	// ALIYUN$7632***@aliyun.com
	JobOwner *string `json:"jobOwner,omitempty" xml:"jobOwner,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The name of MaxCompute project.
	ProjectNames []*string `json:"projectNames,omitempty" xml:"projectNames,omitempty" type:"Repeated"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The signature of the SQL job.
	//
	// example:
	//
	// ghijkl789012
	Signature *string `json:"signature,omitempty" xml:"signature,omitempty"`
	// Specification types.
	SpecCodes []*string `json:"specCodes,omitempty" xml:"specCodes,omitempty" type:"Repeated"`
	// The start time for the period.
	//
	// example:
	//
	// 1715393576201
	StartDate *int64 `json:"startDate,omitempty" xml:"startDate,omitempty"`
	// Metering types.
	Types []*string `json:"types,omitempty" xml:"types,omitempty" type:"Repeated"`
}

func (s ListComputeMetricsByInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ListComputeMetricsByInstanceRequest) GoString() string {
	return s.String()
}

func (s *ListComputeMetricsByInstanceRequest) SetEndDate(v int64) *ListComputeMetricsByInstanceRequest {
	s.EndDate = &v
	return s
}

func (s *ListComputeMetricsByInstanceRequest) SetInstanceId(v string) *ListComputeMetricsByInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *ListComputeMetricsByInstanceRequest) SetJobOwner(v string) *ListComputeMetricsByInstanceRequest {
	s.JobOwner = &v
	return s
}

func (s *ListComputeMetricsByInstanceRequest) SetPageNumber(v int64) *ListComputeMetricsByInstanceRequest {
	s.PageNumber = &v
	return s
}

func (s *ListComputeMetricsByInstanceRequest) SetPageSize(v int64) *ListComputeMetricsByInstanceRequest {
	s.PageSize = &v
	return s
}

func (s *ListComputeMetricsByInstanceRequest) SetProjectNames(v []*string) *ListComputeMetricsByInstanceRequest {
	s.ProjectNames = v
	return s
}

func (s *ListComputeMetricsByInstanceRequest) SetRegion(v string) *ListComputeMetricsByInstanceRequest {
	s.Region = &v
	return s
}

func (s *ListComputeMetricsByInstanceRequest) SetSignature(v string) *ListComputeMetricsByInstanceRequest {
	s.Signature = &v
	return s
}

func (s *ListComputeMetricsByInstanceRequest) SetSpecCodes(v []*string) *ListComputeMetricsByInstanceRequest {
	s.SpecCodes = v
	return s
}

func (s *ListComputeMetricsByInstanceRequest) SetStartDate(v int64) *ListComputeMetricsByInstanceRequest {
	s.StartDate = &v
	return s
}

func (s *ListComputeMetricsByInstanceRequest) SetTypes(v []*string) *ListComputeMetricsByInstanceRequest {
	s.Types = v
	return s
}

type ListComputeMetricsByInstanceResponseBody struct {
	// The data returned.
	Data *ListComputeMetricsByInstanceResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// OBJECT_NOT_EXIST
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// This object does not exist.
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// The HTTP status code.
	//
	// - 1xx: informational response. The request is received and is being processed.
	//
	// - 2xx: success. The request is successfully received, understood, and accepted by the server.
	//
	// - 3xx: redirection. The request is redirected, and further actions are required to complete the request.
	//
	// - 4xx: client error. The request contains invalid request parameters or syntaxes, or specific request conditions cannot be met.
	//
	// - 5xx: server error. The server cannot meet requirements due to other reasons.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0bc059b717363029839908920ea631
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListComputeMetricsByInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListComputeMetricsByInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ListComputeMetricsByInstanceResponseBody) SetData(v *ListComputeMetricsByInstanceResponseBodyData) *ListComputeMetricsByInstanceResponseBody {
	s.Data = v
	return s
}

func (s *ListComputeMetricsByInstanceResponseBody) SetErrorCode(v string) *ListComputeMetricsByInstanceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListComputeMetricsByInstanceResponseBody) SetErrorMsg(v string) *ListComputeMetricsByInstanceResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListComputeMetricsByInstanceResponseBody) SetHttpCode(v int32) *ListComputeMetricsByInstanceResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListComputeMetricsByInstanceResponseBody) SetRequestId(v string) *ListComputeMetricsByInstanceResponseBody {
	s.RequestId = &v
	return s
}

type ListComputeMetricsByInstanceResponseBodyData struct {
	// List of pay-as-you-go job compute usage.
	InstanceComputeMetrics []*ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics `json:"instanceComputeMetrics,omitempty" xml:"instanceComputeMetrics,omitempty" type:"Repeated"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The total number of results returned.
	//
	// example:
	//
	// 64
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListComputeMetricsByInstanceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListComputeMetricsByInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListComputeMetricsByInstanceResponseBodyData) SetInstanceComputeMetrics(v []*ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics) *ListComputeMetricsByInstanceResponseBodyData {
	s.InstanceComputeMetrics = v
	return s
}

func (s *ListComputeMetricsByInstanceResponseBodyData) SetPageNumber(v int64) *ListComputeMetricsByInstanceResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListComputeMetricsByInstanceResponseBodyData) SetPageSize(v int64) *ListComputeMetricsByInstanceResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListComputeMetricsByInstanceResponseBodyData) SetTotalCount(v int64) *ListComputeMetricsByInstanceResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics struct {
	// The end time of the job execution.
	//
	// example:
	//
	// 1710432000000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The job(instance) ID.
	//
	// example:
	//
	// 20240730****ddlr
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The owner of the job.
	//
	// example:
	//
	// ALIYUN$7632***@aliyun.com
	JobOwner *string `json:"jobOwner,omitempty" xml:"jobOwner,omitempty"`
	// The name of the project.
	//
	// example:
	//
	// odps_porject
	ProjectName *string `json:"projectName,omitempty" xml:"projectName,omitempty"`
	// The signature of the SQL job.
	//
	// example:
	//
	// pqrs12345tuv
	Signature *string `json:"signature,omitempty" xml:"signature,omitempty"`
	// Specifications Type, specifies the resource package that you select when you purchase the MaxCompute service.
	//
	// - OdpsStandard: the pay-as-you-go resource package.
	//
	// - OdpsSpot: the pay-as-you-go spot resource package.
	//
	// example:
	//
	// OdpsStandard
	SpecCode *string `json:"specCode,omitempty" xml:"specCode,omitempty"`
	// The submission time of the job.
	//
	// example:
	//
	// 1610432000000
	SubmitTime *int64 `json:"submitTime,omitempty" xml:"submitTime,omitempty"`
	// Metering types.
	//
	// - ComputationSql: the metering data of SQL jobs that involve internal tables.
	//
	// - ComputationSqlOTS: the metering data of SQL jobs that involve Tablestore external tables.
	//
	// - ComputationSqlOSS: the metering data of SQL jobs that involve OSS external tables.
	//
	// - MapReduce: the metering data of MapReduce jobs.
	//
	// - spark: the metering data of Spark jobs.
	//
	// - mars: the metering data of Mars jobs.
	//
	// example:
	//
	// ComputationSql
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The unit of computing resource usage
	//
	// example:
	//
	// GB
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// The computing resource usage is calculated based on the following items:
	//
	// - Amount of scanned data in the unit of GB. For the jobs whose metering types are ComputationSql, ComputationSqlOTS, or ComputationSqlOSS, they are billed based on the amount of scanned data. The computing resource usage of such a job is calculated by using the following formula: Amount of scanned data × Complexity. The complexity is fixed at 1 for the jobs whose metering types are ComputationSqlOTS or ComputationSqlOSS.
	//
	// - CU-hours. For the jobs whose metering types are MapReduce, spark, or mars, they are billed based on CU-hours.
	//
	// example:
	//
	// 1024
	Usage *float64 `json:"usage,omitempty" xml:"usage,omitempty"`
}

func (s ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics) String() string {
	return tea.Prettify(s)
}

func (s ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics) GoString() string {
	return s.String()
}

func (s *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics) SetEndTime(v int64) *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics {
	s.EndTime = &v
	return s
}

func (s *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics) SetInstanceId(v string) *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics {
	s.InstanceId = &v
	return s
}

func (s *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics) SetJobOwner(v string) *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics {
	s.JobOwner = &v
	return s
}

func (s *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics) SetProjectName(v string) *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics {
	s.ProjectName = &v
	return s
}

func (s *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics) SetSignature(v string) *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics {
	s.Signature = &v
	return s
}

func (s *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics) SetSpecCode(v string) *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics {
	s.SpecCode = &v
	return s
}

func (s *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics) SetSubmitTime(v int64) *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics {
	s.SubmitTime = &v
	return s
}

func (s *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics) SetType(v string) *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics {
	s.Type = &v
	return s
}

func (s *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics) SetUnit(v string) *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics {
	s.Unit = &v
	return s
}

func (s *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics) SetUsage(v float64) *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics {
	s.Usage = &v
	return s
}

type ListComputeMetricsByInstanceResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListComputeMetricsByInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListComputeMetricsByInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s ListComputeMetricsByInstanceResponse) GoString() string {
	return s.String()
}

func (s *ListComputeMetricsByInstanceResponse) SetHeaders(v map[string]*string) *ListComputeMetricsByInstanceResponse {
	s.Headers = v
	return s
}

func (s *ListComputeMetricsByInstanceResponse) SetStatusCode(v int32) *ListComputeMetricsByInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListComputeMetricsByInstanceResponse) SetBody(v *ListComputeMetricsByInstanceResponseBody) *ListComputeMetricsByInstanceResponse {
	s.Body = v
	return s
}

type ListComputeQuotaPlanResponseBody struct {
	// The data returned.
	Data *ListComputeQuotaPlanResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// OBJECT_NOT_EXIST
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// This object does not exist.
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// The HTTP status code.
	//
	// - 1xx: informational response. The request is received and is being processed.
	//
	// - 2xx: success. The request is successfully received, understood, and accepted by the server.
	//
	// - 3xx: redirection. The request is redirected, and further actions are required to complete the request.
	//
	// - 4xx: client error. The request contains invalid request parameters or syntaxes, or specific request conditions cannot be met.
	//
	// - 5xx: server error. The server cannot meet requirements due to other reasons.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0bc3b4ae16685836687916212e7850
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListComputeQuotaPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListComputeQuotaPlanResponseBody) GoString() string {
	return s.String()
}

func (s *ListComputeQuotaPlanResponseBody) SetData(v *ListComputeQuotaPlanResponseBodyData) *ListComputeQuotaPlanResponseBody {
	s.Data = v
	return s
}

func (s *ListComputeQuotaPlanResponseBody) SetErrorCode(v string) *ListComputeQuotaPlanResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListComputeQuotaPlanResponseBody) SetErrorMsg(v string) *ListComputeQuotaPlanResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListComputeQuotaPlanResponseBody) SetHttpCode(v int32) *ListComputeQuotaPlanResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListComputeQuotaPlanResponseBody) SetRequestId(v string) *ListComputeQuotaPlanResponseBody {
	s.RequestId = &v
	return s
}

type ListComputeQuotaPlanResponseBodyData struct {
	// The list of quota plan.
	PlanList []*ListComputeQuotaPlanResponseBodyDataPlanList `json:"planList,omitempty" xml:"planList,omitempty" type:"Repeated"`
}

func (s ListComputeQuotaPlanResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListComputeQuotaPlanResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListComputeQuotaPlanResponseBodyData) SetPlanList(v []*ListComputeQuotaPlanResponseBodyDataPlanList) *ListComputeQuotaPlanResponseBodyData {
	s.PlanList = v
	return s
}

type ListComputeQuotaPlanResponseBodyDataPlanList struct {
	// The time when the quota plan was created.
	//
	// example:
	//
	// 1731394621890
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The name of the quota plan.
	//
	// example:
	//
	// planA
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The details of the quota.
	Quota *ListComputeQuotaPlanResponseBodyDataPlanListQuota `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
}

func (s ListComputeQuotaPlanResponseBodyDataPlanList) String() string {
	return tea.Prettify(s)
}

func (s ListComputeQuotaPlanResponseBodyDataPlanList) GoString() string {
	return s.String()
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanList) SetCreateTime(v string) *ListComputeQuotaPlanResponseBodyDataPlanList {
	s.CreateTime = &v
	return s
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanList) SetName(v string) *ListComputeQuotaPlanResponseBodyDataPlanList {
	s.Name = &v
	return s
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanList) SetQuota(v *ListComputeQuotaPlanResponseBodyDataPlanListQuota) *ListComputeQuotaPlanResponseBodyDataPlanList {
	s.Quota = v
	return s
}

type ListComputeQuotaPlanResponseBodyDataPlanListQuota struct {
	// Cluster ID.
	//
	// example:
	//
	// AT-120N
	Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// The time when the level-1 quota was created.
	//
	// example:
	//
	// 1730247361356
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The ID of the Alibaba Cloud account that is used to create the resource.
	//
	// example:
	//
	// 672863518
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// The ID of the level-1 quota.
	//
	// example:
	//
	// 186
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the level-1 quota.
	//
	// example:
	//
	// dp_cn_hangzhou_1717465943_p
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The nickname of the level-1 quota.
	//
	// example:
	//
	// os_MyQuota_p
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	// The description of the level-1 quota.
	//
	// example:
	//
	// {
	//
	//   "enablePriority": false,
	//
	//   "minCU": 25,
	//
	//   "adhocCU": 0,
	//
	//   "elasticReservedCU": 0,
	//
	//   "forceReservedMin": false,
	//
	//   "maxCU": 50,
	//
	//   "schedulerType": "Fifo"
	//
	// }
	Parameter *ListComputeQuotaPlanResponseBodyDataPlanListQuotaParameter `json:"parameter,omitempty" xml:"parameter,omitempty" type:"Struct"`
	// Region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// Resource status.
	//
	// example:
	//
	// ON
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The list of subquotas.
	SubQuotaInfoList []*ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList `json:"subQuotaInfoList,omitempty" xml:"subQuotaInfoList,omitempty" type:"Repeated"`
	// Tenant ID.
	//
	// example:
	//
	// 478403690625249
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The type of quota.
	//
	// example:
	//
	// FUXI_ONLINE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The version number.
	//
	// example:
	//
	// 2056
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListComputeQuotaPlanResponseBodyDataPlanListQuota) String() string {
	return tea.Prettify(s)
}

func (s ListComputeQuotaPlanResponseBodyDataPlanListQuota) GoString() string {
	return s.String()
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuota) SetCluster(v string) *ListComputeQuotaPlanResponseBodyDataPlanListQuota {
	s.Cluster = &v
	return s
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuota) SetCreateTime(v int64) *ListComputeQuotaPlanResponseBodyDataPlanListQuota {
	s.CreateTime = &v
	return s
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuota) SetCreatorId(v string) *ListComputeQuotaPlanResponseBodyDataPlanListQuota {
	s.CreatorId = &v
	return s
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuota) SetId(v string) *ListComputeQuotaPlanResponseBodyDataPlanListQuota {
	s.Id = &v
	return s
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuota) SetName(v string) *ListComputeQuotaPlanResponseBodyDataPlanListQuota {
	s.Name = &v
	return s
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuota) SetNickName(v string) *ListComputeQuotaPlanResponseBodyDataPlanListQuota {
	s.NickName = &v
	return s
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuota) SetParameter(v *ListComputeQuotaPlanResponseBodyDataPlanListQuotaParameter) *ListComputeQuotaPlanResponseBodyDataPlanListQuota {
	s.Parameter = v
	return s
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuota) SetRegionId(v string) *ListComputeQuotaPlanResponseBodyDataPlanListQuota {
	s.RegionId = &v
	return s
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuota) SetStatus(v string) *ListComputeQuotaPlanResponseBodyDataPlanListQuota {
	s.Status = &v
	return s
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuota) SetSubQuotaInfoList(v []*ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList) *ListComputeQuotaPlanResponseBodyDataPlanListQuota {
	s.SubQuotaInfoList = v
	return s
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuota) SetTenantId(v string) *ListComputeQuotaPlanResponseBodyDataPlanListQuota {
	s.TenantId = &v
	return s
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuota) SetType(v string) *ListComputeQuotaPlanResponseBodyDataPlanListQuota {
	s.Type = &v
	return s
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuota) SetVersion(v string) *ListComputeQuotaPlanResponseBodyDataPlanListQuota {
	s.Version = &v
	return s
}

type ListComputeQuotaPlanResponseBodyDataPlanListQuotaParameter struct {
	ElasticReservedCU *int64 `json:"elasticReservedCU,omitempty" xml:"elasticReservedCU,omitempty"`
	MaxCU             *int64 `json:"maxCU,omitempty" xml:"maxCU,omitempty"`
	MinCU             *int64 `json:"minCU,omitempty" xml:"minCU,omitempty"`
}

func (s ListComputeQuotaPlanResponseBodyDataPlanListQuotaParameter) String() string {
	return tea.Prettify(s)
}

func (s ListComputeQuotaPlanResponseBodyDataPlanListQuotaParameter) GoString() string {
	return s.String()
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuotaParameter) SetElasticReservedCU(v int64) *ListComputeQuotaPlanResponseBodyDataPlanListQuotaParameter {
	s.ElasticReservedCU = &v
	return s
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuotaParameter) SetMaxCU(v int64) *ListComputeQuotaPlanResponseBodyDataPlanListQuotaParameter {
	s.MaxCU = &v
	return s
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuotaParameter) SetMinCU(v int64) *ListComputeQuotaPlanResponseBodyDataPlanListQuotaParameter {
	s.MinCU = &v
	return s
}

type ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList struct {
	// Cluster ID.
	//
	// example:
	//
	// AT-120N
	Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 1730946421757
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The ID of the Alibaba Cloud account that is used to create the resource.
	//
	// example:
	//
	// 672863518
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// The ID of the level-2 quota.
	//
	// example:
	//
	// 6790
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the level-2 quota.
	//
	// example:
	//
	// dp_cn_shanghai_1702627945_p
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The nickname of the level-2 quota.
	//
	// example:
	//
	// os_MyQuota
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	// The description of the level-2 quota.
	//
	// example:
	//
	// {
	//
	//   "enablePriority": false,
	//
	//   "minCU": 25,
	//
	//   "adhocCU": 0,
	//
	//   "elasticReservedCU": 0,
	//
	//   "forceReservedMin": false,
	//
	//   "maxCU": 50,
	//
	//   "schedulerType": "Fifo"
	//
	// }
	Parameter *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoListParameter `json:"parameter,omitempty" xml:"parameter,omitempty" type:"Struct"`
	// Region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// Resource status.
	//
	// example:
	//
	// ON
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// Tenant ID.
	//
	// example:
	//
	// 478403690625249
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The type of quota.
	//
	// example:
	//
	// FUXI_ONLINE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The version number.
	//
	// example:
	//
	// 2056
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList) String() string {
	return tea.Prettify(s)
}

func (s ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList) GoString() string {
	return s.String()
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList) SetCluster(v string) *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.Cluster = &v
	return s
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList) SetCreateTime(v int64) *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.CreateTime = &v
	return s
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList) SetCreatorId(v string) *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.CreatorId = &v
	return s
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList) SetId(v string) *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.Id = &v
	return s
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList) SetName(v string) *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.Name = &v
	return s
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList) SetNickName(v string) *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.NickName = &v
	return s
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList) SetParameter(v *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoListParameter) *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.Parameter = v
	return s
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList) SetRegionId(v string) *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.RegionId = &v
	return s
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList) SetStatus(v string) *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.Status = &v
	return s
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList) SetTenantId(v string) *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.TenantId = &v
	return s
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList) SetType(v string) *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.Type = &v
	return s
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList) SetVersion(v string) *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.Version = &v
	return s
}

type ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoListParameter struct {
	ElasticReservedCU *int64 `json:"elasticReservedCU,omitempty" xml:"elasticReservedCU,omitempty"`
	MaxCU             *int64 `json:"maxCU,omitempty" xml:"maxCU,omitempty"`
	MinCU             *int64 `json:"minCU,omitempty" xml:"minCU,omitempty"`
}

func (s ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoListParameter) String() string {
	return tea.Prettify(s)
}

func (s ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoListParameter) GoString() string {
	return s.String()
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoListParameter) SetElasticReservedCU(v int64) *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoListParameter {
	s.ElasticReservedCU = &v
	return s
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoListParameter) SetMaxCU(v int64) *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoListParameter {
	s.MaxCU = &v
	return s
}

func (s *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoListParameter) SetMinCU(v int64) *ListComputeQuotaPlanResponseBodyDataPlanListQuotaSubQuotaInfoListParameter {
	s.MinCU = &v
	return s
}

type ListComputeQuotaPlanResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListComputeQuotaPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListComputeQuotaPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s ListComputeQuotaPlanResponse) GoString() string {
	return s.String()
}

func (s *ListComputeQuotaPlanResponse) SetHeaders(v map[string]*string) *ListComputeQuotaPlanResponse {
	s.Headers = v
	return s
}

func (s *ListComputeQuotaPlanResponse) SetStatusCode(v int32) *ListComputeQuotaPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *ListComputeQuotaPlanResponse) SetBody(v *ListComputeQuotaPlanResponseBody) *ListComputeQuotaPlanResponse {
	s.Body = v
	return s
}

type ListFunctionsRequest struct {
	// Specifies the marker after which the returned list begins.
	//
	// example:
	//
	// cHlvZHBzX3VkZl8xMDExNV8xNDU3NDI4NDkzKg==
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
	// The maximum number of entries to return on each page.
	//
	// example:
	//
	// 10
	MaxItem *int32 `json:"maxItem,omitempty" xml:"maxItem,omitempty"`
	// The names of the returned resources. The names must start with the value specified by the prefix parameter. If the prefix parameter is set to a, the names of the returned resources must start with a.
	//
	// example:
	//
	// a
	Prefix *string `json:"prefix,omitempty" xml:"prefix,omitempty"`
	// the name of schema.
	//
	// example:
	//
	// default
	SchemaName *string `json:"schemaName,omitempty" xml:"schemaName,omitempty"`
}

func (s ListFunctionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionsRequest) GoString() string {
	return s.String()
}

func (s *ListFunctionsRequest) SetMarker(v string) *ListFunctionsRequest {
	s.Marker = &v
	return s
}

func (s *ListFunctionsRequest) SetMaxItem(v int32) *ListFunctionsRequest {
	s.MaxItem = &v
	return s
}

func (s *ListFunctionsRequest) SetPrefix(v string) *ListFunctionsRequest {
	s.Prefix = &v
	return s
}

func (s *ListFunctionsRequest) SetSchemaName(v string) *ListFunctionsRequest {
	s.SchemaName = &v
	return s
}

type ListFunctionsResponseBody struct {
	// The returned data.
	Data *ListFunctionsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 0be3e0b716671885050924814e3623
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListFunctionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFunctionsResponseBody) SetData(v *ListFunctionsResponseBodyData) *ListFunctionsResponseBody {
	s.Data = v
	return s
}

func (s *ListFunctionsResponseBody) SetRequestId(v string) *ListFunctionsResponseBody {
	s.RequestId = &v
	return s
}

type ListFunctionsResponseBodyData struct {
	// The information about each function.
	Functions []*ListFunctionsResponseBodyDataFunctions `json:"functions,omitempty" xml:"functions,omitempty" type:"Repeated"`
	// Indicates the marker after which the returned list begins.
	//
	// example:
	//
	// cHlvZHBzX3VkZl8xMDExNV8xNDU3NDI4NDkzKg==
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
	// The maximum number of entries returned per page.
	//
	// example:
	//
	// 10
	MaxItem *int32 `json:"maxItem,omitempty" xml:"maxItem,omitempty"`
}

func (s ListFunctionsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListFunctionsResponseBodyData) SetFunctions(v []*ListFunctionsResponseBodyDataFunctions) *ListFunctionsResponseBodyData {
	s.Functions = v
	return s
}

func (s *ListFunctionsResponseBodyData) SetMarker(v string) *ListFunctionsResponseBodyData {
	s.Marker = &v
	return s
}

func (s *ListFunctionsResponseBodyData) SetMaxItem(v int32) *ListFunctionsResponseBodyData {
	s.MaxItem = &v
	return s
}

type ListFunctionsResponseBodyDataFunctions struct {
	// The class in which the function was defined.
	//
	// example:
	//
	// abc
	Class *string `json:"class,omitempty" xml:"class,omitempty"`
	// The time when the function was created. Unit: milliseconds.
	//
	// example:
	//
	// 1664505167000
	CreationTime *int64 `json:"creationTime,omitempty" xml:"creationTime,omitempty"`
	// The display name of the function.
	//
	// example:
	//
	// getdate
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// The name of the function.
	//
	// example:
	//
	// getdate
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The owner of the function.
	//
	// example:
	//
	// odpsowner
	Owner *string `json:"owner,omitempty" xml:"owner,omitempty"`
	// The name of the resource that was associated with the function.
	//
	// example:
	//
	// abc
	Resources *string `json:"resources,omitempty" xml:"resources,omitempty"`
	// The schema of the function.
	//
	// example:
	//
	// abc
	Schema *string `json:"schema,omitempty" xml:"schema,omitempty"`
}

func (s ListFunctionsResponseBodyDataFunctions) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionsResponseBodyDataFunctions) GoString() string {
	return s.String()
}

func (s *ListFunctionsResponseBodyDataFunctions) SetClass(v string) *ListFunctionsResponseBodyDataFunctions {
	s.Class = &v
	return s
}

func (s *ListFunctionsResponseBodyDataFunctions) SetCreationTime(v int64) *ListFunctionsResponseBodyDataFunctions {
	s.CreationTime = &v
	return s
}

func (s *ListFunctionsResponseBodyDataFunctions) SetDisplayName(v string) *ListFunctionsResponseBodyDataFunctions {
	s.DisplayName = &v
	return s
}

func (s *ListFunctionsResponseBodyDataFunctions) SetName(v string) *ListFunctionsResponseBodyDataFunctions {
	s.Name = &v
	return s
}

func (s *ListFunctionsResponseBodyDataFunctions) SetOwner(v string) *ListFunctionsResponseBodyDataFunctions {
	s.Owner = &v
	return s
}

func (s *ListFunctionsResponseBodyDataFunctions) SetResources(v string) *ListFunctionsResponseBodyDataFunctions {
	s.Resources = &v
	return s
}

func (s *ListFunctionsResponseBodyDataFunctions) SetSchema(v string) *ListFunctionsResponseBodyDataFunctions {
	s.Schema = &v
	return s
}

type ListFunctionsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFunctionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFunctionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionsResponse) GoString() string {
	return s.String()
}

func (s *ListFunctionsResponse) SetHeaders(v map[string]*string) *ListFunctionsResponse {
	s.Headers = v
	return s
}

func (s *ListFunctionsResponse) SetStatusCode(v int32) *ListFunctionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFunctionsResponse) SetBody(v *ListFunctionsResponseBody) *ListFunctionsResponse {
	s.Body = v
	return s
}

type ListJobInfosRequest struct {
	// Specifies whether to sort query results in ascending or descending order.
	//
	// example:
	//
	// true
	AscOrder      *bool     `json:"ascOrder,omitempty" xml:"ascOrder,omitempty"`
	ExtNodeIdList []*string `json:"extNodeIdList,omitempty" xml:"extNodeIdList,omitempty" type:"Repeated"`
	// This parameter is required.
	From           *int64    `json:"from,omitempty" xml:"from,omitempty"`
	InstanceIdList []*string `json:"instanceIdList,omitempty" xml:"instanceIdList,omitempty" type:"Repeated"`
	JobOwnerList   []*string `json:"jobOwnerList,omitempty" xml:"jobOwnerList,omitempty" type:"Repeated"`
	PriorityList   []*int64  `json:"priorityList,omitempty" xml:"priorityList,omitempty" type:"Repeated"`
	ProjectList    []*string `json:"projectList,omitempty" xml:"projectList,omitempty" type:"Repeated"`
	QuotaNickname  *string   `json:"quotaNickname,omitempty" xml:"quotaNickname,omitempty"`
	SceneTagList   []*string `json:"sceneTagList,omitempty" xml:"sceneTagList,omitempty" type:"Repeated"`
	SignatureList  []*string `json:"signatureList,omitempty" xml:"signatureList,omitempty" type:"Repeated"`
	SortByList     []*string `json:"sortByList,omitempty" xml:"sortByList,omitempty" type:"Repeated"`
	SortOrderList  []*string `json:"sortOrderList,omitempty" xml:"sortOrderList,omitempty" type:"Repeated"`
	StatusList     []*string `json:"statusList,omitempty" xml:"statusList,omitempty" type:"Repeated"`
	// This parameter is required.
	To       *int64    `json:"to,omitempty" xml:"to,omitempty"`
	TypeList []*string `json:"typeList,omitempty" xml:"typeList,omitempty" type:"Repeated"`
	// The column based on which you want to sort query results.
	//
	// example:
	//
	// cuUsage
	OrderColumn *string `json:"orderColumn,omitempty" xml:"orderColumn,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 478403690625249
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListJobInfosRequest) String() string {
	return tea.Prettify(s)
}

func (s ListJobInfosRequest) GoString() string {
	return s.String()
}

func (s *ListJobInfosRequest) SetAscOrder(v bool) *ListJobInfosRequest {
	s.AscOrder = &v
	return s
}

func (s *ListJobInfosRequest) SetExtNodeIdList(v []*string) *ListJobInfosRequest {
	s.ExtNodeIdList = v
	return s
}

func (s *ListJobInfosRequest) SetFrom(v int64) *ListJobInfosRequest {
	s.From = &v
	return s
}

func (s *ListJobInfosRequest) SetInstanceIdList(v []*string) *ListJobInfosRequest {
	s.InstanceIdList = v
	return s
}

func (s *ListJobInfosRequest) SetJobOwnerList(v []*string) *ListJobInfosRequest {
	s.JobOwnerList = v
	return s
}

func (s *ListJobInfosRequest) SetPriorityList(v []*int64) *ListJobInfosRequest {
	s.PriorityList = v
	return s
}

func (s *ListJobInfosRequest) SetProjectList(v []*string) *ListJobInfosRequest {
	s.ProjectList = v
	return s
}

func (s *ListJobInfosRequest) SetQuotaNickname(v string) *ListJobInfosRequest {
	s.QuotaNickname = &v
	return s
}

func (s *ListJobInfosRequest) SetSceneTagList(v []*string) *ListJobInfosRequest {
	s.SceneTagList = v
	return s
}

func (s *ListJobInfosRequest) SetSignatureList(v []*string) *ListJobInfosRequest {
	s.SignatureList = v
	return s
}

func (s *ListJobInfosRequest) SetSortByList(v []*string) *ListJobInfosRequest {
	s.SortByList = v
	return s
}

func (s *ListJobInfosRequest) SetSortOrderList(v []*string) *ListJobInfosRequest {
	s.SortOrderList = v
	return s
}

func (s *ListJobInfosRequest) SetStatusList(v []*string) *ListJobInfosRequest {
	s.StatusList = v
	return s
}

func (s *ListJobInfosRequest) SetTo(v int64) *ListJobInfosRequest {
	s.To = &v
	return s
}

func (s *ListJobInfosRequest) SetTypeList(v []*string) *ListJobInfosRequest {
	s.TypeList = v
	return s
}

func (s *ListJobInfosRequest) SetOrderColumn(v string) *ListJobInfosRequest {
	s.OrderColumn = &v
	return s
}

func (s *ListJobInfosRequest) SetPageNumber(v int64) *ListJobInfosRequest {
	s.PageNumber = &v
	return s
}

func (s *ListJobInfosRequest) SetPageSize(v int64) *ListJobInfosRequest {
	s.PageSize = &v
	return s
}

func (s *ListJobInfosRequest) SetRegion(v string) *ListJobInfosRequest {
	s.Region = &v
	return s
}

func (s *ListJobInfosRequest) SetTenantId(v string) *ListJobInfosRequest {
	s.TenantId = &v
	return s
}

type ListJobInfosResponseBody struct {
	// The data returned.
	Data *ListJobInfosResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Indicates whether the request was successful. If this parameter was not empty and the value of this parameter was not 200, the request failed.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0bc13a9516807484336515320e38f5
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListJobInfosResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListJobInfosResponseBody) GoString() string {
	return s.String()
}

func (s *ListJobInfosResponseBody) SetData(v *ListJobInfosResponseBodyData) *ListJobInfosResponseBody {
	s.Data = v
	return s
}

func (s *ListJobInfosResponseBody) SetHttpCode(v int32) *ListJobInfosResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListJobInfosResponseBody) SetRequestId(v string) *ListJobInfosResponseBody {
	s.RequestId = &v
	return s
}

type ListJobInfosResponseBodyData struct {
	// The information about the jobs.
	JobInfoList []*ListJobInfosResponseBodyDataJobInfoList `json:"jobInfoList,omitempty" xml:"jobInfoList,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 2
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 64
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListJobInfosResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListJobInfosResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListJobInfosResponseBodyData) SetJobInfoList(v []*ListJobInfosResponseBodyDataJobInfoList) *ListJobInfosResponseBodyData {
	s.JobInfoList = v
	return s
}

func (s *ListJobInfosResponseBodyData) SetPageNumber(v int64) *ListJobInfosResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListJobInfosResponseBodyData) SetPageSize(v int64) *ListJobInfosResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListJobInfosResponseBodyData) SetTotalCount(v int64) *ListJobInfosResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListJobInfosResponseBodyDataJobInfoList struct {
	// The cluster ID.
	//
	// example:
	//
	// AY20A
	Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// The CU snapshot proportion of the job.
	//
	// example:
	//
	// 0.48
	CuSnapshot *float64 `json:"cuSnapshot,omitempty" xml:"cuSnapshot,omitempty"`
	// The total number of used compute units (CUs).
	//
	// example:
	//
	// 10
	CuUsage *int64 `json:"cuUsage,omitempty" xml:"cuUsage,omitempty"`
	// The time when the job stops running.
	//
	// example:
	//
	// 0
	EndAtTime *int64 `json:"endAtTime,omitempty" xml:"endAtTime,omitempty"`
	// The node ID of DataWorks.
	//
	// example:
	//
	// node_4
	ExtNodeId *string `json:"extNodeId,omitempty" xml:"extNodeId,omitempty"`
	// The account of the node owner.
	//
	// example:
	//
	// duty_2
	ExtNodeOnDuty *string `json:"extNodeOnDuty,omitempty" xml:"extNodeOnDuty,omitempty"`
	// The upstream platform.
	//
	// example:
	//
	// platform_3
	ExtPlantFrom *string `json:"extPlantFrom,omitempty" xml:"extPlantFrom,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// 20230410050036549gfmsdwf60gg
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The account that commits the job.
	//
	// example:
	//
	// ALIYUN$xxx@test.aliyunid.com
	JobOwner *string `json:"jobOwner,omitempty" xml:"jobOwner,omitempty"`
	// The type of the job.
	//
	// example:
	//
	// SQL
	JobType *string `json:"jobType,omitempty" xml:"jobType,omitempty"`
	// The memory snapshot proportion of the job.
	//
	// example:
	//
	// 0.42
	MemorySnapshot *float64 `json:"memorySnapshot,omitempty" xml:"memorySnapshot,omitempty"`
	// The total memory usage.
	//
	// example:
	//
	// 40
	MemoryUsage *int64 `json:"memoryUsage,omitempty" xml:"memoryUsage,omitempty"`
	// The priority of the job.
	//
	// example:
	//
	// 1
	Priority *int64 `json:"priority,omitempty" xml:"priority,omitempty"`
	// The name of the MaxCompute project.
	//
	// example:
	//
	// openrec_new
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
	// The nickname of the quota that is used by the job.
	//
	// example:
	//
	// my_quota
	QuotaNickname *string `json:"quotaNickname,omitempty" xml:"quotaNickname,omitempty"`
	// The type of the quota.
	//
	// example:
	//
	// subscription
	QuotaType *string `json:"quotaType,omitempty" xml:"quotaType,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The time when the job starts to run.
	//
	// example:
	//
	// 1672112113
	RunningAtTime *int64 `json:"runningAtTime,omitempty" xml:"runningAtTime,omitempty"`
	// The period for which the job runs.
	//
	// example:
	//
	// 800
	RunningTime *int64 `json:"runningTime,omitempty" xml:"runningTime,omitempty"`
	// The intelligent diagnostics results.
	SceneResults []*ListJobInfosResponseBodyDataJobInfoListSceneResults `json:"sceneResults,omitempty" xml:"sceneResults,omitempty" type:"Repeated"`
	// The signature of the SQL job.
	//
	// example:
	//
	// i094KijGrN3kOXZ74kbexB77XQY=
	Signature *string `json:"signature,omitempty" xml:"signature,omitempty"`
	// The status of the job.
	//
	// example:
	//
	// running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The status of the snapshot.
	//
	// example:
	//
	// running
	StatusSnapshot *string `json:"statusSnapshot,omitempty" xml:"statusSnapshot,omitempty"`
	// The time when the job was committed.
	//
	// example:
	//
	// 1672112013
	SubmittedAtTime *int64 `json:"submittedAtTime,omitempty" xml:"submittedAtTime,omitempty"`
	// The tags.
	//
	// example:
	//
	// []
	Tags *string `json:"tags,omitempty" xml:"tags,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 213065738244354
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The total period for which the job runs.
	//
	// example:
	//
	// 900
	TotalTime *int64 `json:"totalTime,omitempty" xml:"totalTime,omitempty"`
	// The duration for which the job waits to start.
	//
	// example:
	//
	// 100
	WaitingTime *int64 `json:"waitingTime,omitempty" xml:"waitingTime,omitempty"`
}

func (s ListJobInfosResponseBodyDataJobInfoList) String() string {
	return tea.Prettify(s)
}

func (s ListJobInfosResponseBodyDataJobInfoList) GoString() string {
	return s.String()
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetCluster(v string) *ListJobInfosResponseBodyDataJobInfoList {
	s.Cluster = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetCuSnapshot(v float64) *ListJobInfosResponseBodyDataJobInfoList {
	s.CuSnapshot = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetCuUsage(v int64) *ListJobInfosResponseBodyDataJobInfoList {
	s.CuUsage = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetEndAtTime(v int64) *ListJobInfosResponseBodyDataJobInfoList {
	s.EndAtTime = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetExtNodeId(v string) *ListJobInfosResponseBodyDataJobInfoList {
	s.ExtNodeId = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetExtNodeOnDuty(v string) *ListJobInfosResponseBodyDataJobInfoList {
	s.ExtNodeOnDuty = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetExtPlantFrom(v string) *ListJobInfosResponseBodyDataJobInfoList {
	s.ExtPlantFrom = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetInstanceId(v string) *ListJobInfosResponseBodyDataJobInfoList {
	s.InstanceId = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetJobOwner(v string) *ListJobInfosResponseBodyDataJobInfoList {
	s.JobOwner = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetJobType(v string) *ListJobInfosResponseBodyDataJobInfoList {
	s.JobType = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetMemorySnapshot(v float64) *ListJobInfosResponseBodyDataJobInfoList {
	s.MemorySnapshot = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetMemoryUsage(v int64) *ListJobInfosResponseBodyDataJobInfoList {
	s.MemoryUsage = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetPriority(v int64) *ListJobInfosResponseBodyDataJobInfoList {
	s.Priority = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetProject(v string) *ListJobInfosResponseBodyDataJobInfoList {
	s.Project = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetQuotaNickname(v string) *ListJobInfosResponseBodyDataJobInfoList {
	s.QuotaNickname = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetQuotaType(v string) *ListJobInfosResponseBodyDataJobInfoList {
	s.QuotaType = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetRegion(v string) *ListJobInfosResponseBodyDataJobInfoList {
	s.Region = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetRunningAtTime(v int64) *ListJobInfosResponseBodyDataJobInfoList {
	s.RunningAtTime = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetRunningTime(v int64) *ListJobInfosResponseBodyDataJobInfoList {
	s.RunningTime = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetSceneResults(v []*ListJobInfosResponseBodyDataJobInfoListSceneResults) *ListJobInfosResponseBodyDataJobInfoList {
	s.SceneResults = v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetSignature(v string) *ListJobInfosResponseBodyDataJobInfoList {
	s.Signature = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetStatus(v string) *ListJobInfosResponseBodyDataJobInfoList {
	s.Status = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetStatusSnapshot(v string) *ListJobInfosResponseBodyDataJobInfoList {
	s.StatusSnapshot = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetSubmittedAtTime(v int64) *ListJobInfosResponseBodyDataJobInfoList {
	s.SubmittedAtTime = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetTags(v string) *ListJobInfosResponseBodyDataJobInfoList {
	s.Tags = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetTenantId(v string) *ListJobInfosResponseBodyDataJobInfoList {
	s.TenantId = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetTotalTime(v int64) *ListJobInfosResponseBodyDataJobInfoList {
	s.TotalTime = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetWaitingTime(v int64) *ListJobInfosResponseBodyDataJobInfoList {
	s.WaitingTime = &v
	return s
}

type ListJobInfosResponseBodyDataJobInfoListSceneResults struct {
	// The intelligent diagnostics result description.
	//
	// example:
	//
	// This job uses annual and monthly computing resources. It may be that the job is waiting for resources due to the large amount of overall job running data, many resources requested, and low job priority. Please go to Resource Consumption to view the specific situation. You can also go to Cost Optimization to see if you need to adjust resource configuration.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Information about the nodes where data skew or data expansion is detected. This parameter is returned only when the diagnostics scenario is data skew or data expansion.
	Params map[string]*string `json:"params,omitempty" xml:"params,omitempty"`
	// The intelligent diagnostics result scenario.
	//
	// example:
	//
	// LackResource
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
	// The intelligent diagnostics result tag.
	//
	// example:
	//
	// SubscriptionLackResource
	SceneTag *string `json:"sceneTag,omitempty" xml:"sceneTag,omitempty"`
	// The intelligent diagnostics result summary.
	//
	// example:
	//
	// Insufficient computing resources available for the job. Click to view details.
	Summary *string `json:"summary,omitempty" xml:"summary,omitempty"`
	// The intelligent diagnostics result type.
	//
	// example:
	//
	// warning
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListJobInfosResponseBodyDataJobInfoListSceneResults) String() string {
	return tea.Prettify(s)
}

func (s ListJobInfosResponseBodyDataJobInfoListSceneResults) GoString() string {
	return s.String()
}

func (s *ListJobInfosResponseBodyDataJobInfoListSceneResults) SetDescription(v string) *ListJobInfosResponseBodyDataJobInfoListSceneResults {
	s.Description = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoListSceneResults) SetParams(v map[string]*string) *ListJobInfosResponseBodyDataJobInfoListSceneResults {
	s.Params = v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoListSceneResults) SetScene(v string) *ListJobInfosResponseBodyDataJobInfoListSceneResults {
	s.Scene = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoListSceneResults) SetSceneTag(v string) *ListJobInfosResponseBodyDataJobInfoListSceneResults {
	s.SceneTag = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoListSceneResults) SetSummary(v string) *ListJobInfosResponseBodyDataJobInfoListSceneResults {
	s.Summary = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoListSceneResults) SetType(v string) *ListJobInfosResponseBodyDataJobInfoListSceneResults {
	s.Type = &v
	return s
}

type ListJobInfosResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListJobInfosResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListJobInfosResponse) String() string {
	return tea.Prettify(s)
}

func (s ListJobInfosResponse) GoString() string {
	return s.String()
}

func (s *ListJobInfosResponse) SetHeaders(v map[string]*string) *ListJobInfosResponse {
	s.Headers = v
	return s
}

func (s *ListJobInfosResponse) SetStatusCode(v int32) *ListJobInfosResponse {
	s.StatusCode = &v
	return s
}

func (s *ListJobInfosResponse) SetBody(v *ListJobInfosResponseBody) *ListJobInfosResponse {
	s.Body = v
	return s
}

type ListJobMetricRequest struct {
	// Grouping basis.
	//
	// > Available values: project, quota, type, status. Meanings:
	//
	// >- project: Group and aggregate by project;
	//
	// >- quota: Group and aggregate by quota;
	//
	// >- type: Group and aggregate by job type;
	//
	// >- status: Group and aggregate by job status.
	//
	// example:
	//
	// project
	Group *string `json:"group,omitempty" xml:"group,omitempty"`
	// The name of observation metric.
	//
	// example:
	//
	// num
	Metric *string `json:"metric,omitempty" xml:"metric,omitempty"`
	// The monitoring statistical period.Unit:Second(s).
	//
	// example:
	//
	// 3600
	Period *int64 `json:"period,omitempty" xml:"period,omitempty"`
	// The name of MaxCompute projects (for filtering).
	Project []*string `json:"project,omitempty" xml:"project,omitempty" type:"Repeated"`
	// The nickname of computing Quota nickname used by the job (for filtering).
	Quota []*string `json:"quota,omitempty" xml:"quota,omitempty" type:"Repeated"`
	// The type of observation metric.
	//
	// example:
	//
	// total
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The end time for the observation interval.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1735536322
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The start time for the observation interval.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1735534322
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s ListJobMetricRequest) String() string {
	return tea.Prettify(s)
}

func (s ListJobMetricRequest) GoString() string {
	return s.String()
}

func (s *ListJobMetricRequest) SetGroup(v string) *ListJobMetricRequest {
	s.Group = &v
	return s
}

func (s *ListJobMetricRequest) SetMetric(v string) *ListJobMetricRequest {
	s.Metric = &v
	return s
}

func (s *ListJobMetricRequest) SetPeriod(v int64) *ListJobMetricRequest {
	s.Period = &v
	return s
}

func (s *ListJobMetricRequest) SetProject(v []*string) *ListJobMetricRequest {
	s.Project = v
	return s
}

func (s *ListJobMetricRequest) SetQuota(v []*string) *ListJobMetricRequest {
	s.Quota = v
	return s
}

func (s *ListJobMetricRequest) SetType(v string) *ListJobMetricRequest {
	s.Type = &v
	return s
}

func (s *ListJobMetricRequest) SetEndTime(v int64) *ListJobMetricRequest {
	s.EndTime = &v
	return s
}

func (s *ListJobMetricRequest) SetStartTime(v int64) *ListJobMetricRequest {
	s.StartTime = &v
	return s
}

type ListJobMetricResponseBody struct {
	// The data returned.
	Data *ListJobMetricResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// OBJECT_NOT_EXIST
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// This object does not exist.
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// HTTP status code.
	//
	// - 1xx: Informational response - Request received, processing continues.
	//
	// - 2xx: Success - The request has been successfully received, understood, and accepted by the server.
	//
	// - 3xx: Redirection - Further action must be taken to complete the request.
	//
	// - 4xx: Client error - The request contains bad syntax or cannot be fulfilled.
	//
	// - 5xx: Server error - The server failed to fulfill an apparently valid request.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0be3e0aa16667684362147582e038f
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListJobMetricResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListJobMetricResponseBody) GoString() string {
	return s.String()
}

func (s *ListJobMetricResponseBody) SetData(v *ListJobMetricResponseBodyData) *ListJobMetricResponseBody {
	s.Data = v
	return s
}

func (s *ListJobMetricResponseBody) SetErrorCode(v string) *ListJobMetricResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListJobMetricResponseBody) SetErrorMsg(v string) *ListJobMetricResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListJobMetricResponseBody) SetHttpCode(v int32) *ListJobMetricResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListJobMetricResponseBody) SetRequestId(v string) *ListJobMetricResponseBody {
	s.RequestId = &v
	return s
}

type ListJobMetricResponseBodyData struct {
	// The category of the metrics.
	//
	// example:
	//
	// job
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// Metric details.
	Metrics []*ListJobMetricResponseBodyDataMetrics `json:"metrics,omitempty" xml:"metrics,omitempty" type:"Repeated"`
	// The name of observation metric.
	//
	// example:
	//
	// num
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The monitoring statistical period.Unit:Second(s).
	//
	// example:
	//
	// 3600
	Period *int64 `json:"period,omitempty" xml:"period,omitempty"`
}

func (s ListJobMetricResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListJobMetricResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListJobMetricResponseBodyData) SetCategory(v string) *ListJobMetricResponseBodyData {
	s.Category = &v
	return s
}

func (s *ListJobMetricResponseBodyData) SetMetrics(v []*ListJobMetricResponseBodyDataMetrics) *ListJobMetricResponseBodyData {
	s.Metrics = v
	return s
}

func (s *ListJobMetricResponseBodyData) SetName(v string) *ListJobMetricResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListJobMetricResponseBodyData) SetPeriod(v int64) *ListJobMetricResponseBodyData {
	s.Period = &v
	return s
}

type ListJobMetricResponseBodyDataMetrics struct {
	// Metric related information.
	Metric map[string]*string `json:"metric,omitempty" xml:"metric,omitempty"`
	// Metric values information.
	Values [][]*float64 `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s ListJobMetricResponseBodyDataMetrics) String() string {
	return tea.Prettify(s)
}

func (s ListJobMetricResponseBodyDataMetrics) GoString() string {
	return s.String()
}

func (s *ListJobMetricResponseBodyDataMetrics) SetMetric(v map[string]*string) *ListJobMetricResponseBodyDataMetrics {
	s.Metric = v
	return s
}

func (s *ListJobMetricResponseBodyDataMetrics) SetValues(v [][]*float64) *ListJobMetricResponseBodyDataMetrics {
	s.Values = v
	return s
}

type ListJobMetricResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListJobMetricResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListJobMetricResponse) String() string {
	return tea.Prettify(s)
}

func (s ListJobMetricResponse) GoString() string {
	return s.String()
}

func (s *ListJobMetricResponse) SetHeaders(v map[string]*string) *ListJobMetricResponse {
	s.Headers = v
	return s
}

func (s *ListJobMetricResponse) SetStatusCode(v int32) *ListJobMetricResponse {
	s.StatusCode = &v
	return s
}

func (s *ListJobMetricResponse) SetBody(v *ListJobMetricResponseBody) *ListJobMetricResponse {
	s.Body = v
	return s
}

type ListJobSnapshotInfosRequest struct {
	// Specifies whether to sort query results in ascending or descending order.
	//
	// example:
	//
	// true
	AscOrder *bool `json:"ascOrder,omitempty" xml:"ascOrder,omitempty"`
	// The ID of the upstream node.
	ExtNodeIdList []*string `json:"extNodeIdList,omitempty" xml:"extNodeIdList,omitempty" type:"Repeated"`
	// Start timestamp.
	//
	// > This parameter is invalid. The end timestamp should be the time point for the snapshot you want to view.
	//
	// example:
	//
	// 1706840714
	From *int64 `json:"from,omitempty" xml:"from,omitempty"`
	// The instance ID.
	InstanceIdList []*string `json:"instanceIdList,omitempty" xml:"instanceIdList,omitempty" type:"Repeated"`
	// The account that commits the job.
	JobOwnerList []*string `json:"jobOwnerList,omitempty" xml:"jobOwnerList,omitempty" type:"Repeated"`
	// The priority of the job.
	PriorityList []*int64 `json:"priorityList,omitempty" xml:"priorityList,omitempty" type:"Repeated"`
	// The name of project.
	ProjectList []*string `json:"projectList,omitempty" xml:"projectList,omitempty" type:"Repeated"`
	// The nickname of the compute Quota used by the job.
	//
	// example:
	//
	// quota_A
	QuotaNickname *string `json:"quotaNickname,omitempty" xml:"quotaNickname,omitempty"`
	// The signature of the SQL job.
	SignatureList []*string `json:"signatureList,omitempty" xml:"signatureList,omitempty" type:"Repeated"`
	// The sorting columns.
	SortByList []*string `json:"sortByList,omitempty" xml:"sortByList,omitempty" type:"Repeated"`
	// The orders for the sorting columns.
	SortOrderList []*string `json:"sortOrderList,omitempty" xml:"sortOrderList,omitempty" type:"Repeated"`
	// The status of jobs.
	StatusList []*string `json:"statusList,omitempty" xml:"statusList,omitempty" type:"Repeated"`
	// End timestamp.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1706840714
	To *int64 `json:"to,omitempty" xml:"to,omitempty"`
	// The type of the job.
	TypeList []*string `json:"typeList,omitempty" xml:"typeList,omitempty" type:"Repeated"`
	// The column based on which you want to sort query results.
	//
	// example:
	//
	// cpuUsage
	OrderColumn *string `json:"orderColumn,omitempty" xml:"orderColumn,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-chengdu
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The ID of the tenant. You can log on to the MaxCompute console, and choose Tenants > Tenant Property from the left-side navigation pane to view the tenant ID.
	//
	// example:
	//
	// 478403690625249
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListJobSnapshotInfosRequest) String() string {
	return tea.Prettify(s)
}

func (s ListJobSnapshotInfosRequest) GoString() string {
	return s.String()
}

func (s *ListJobSnapshotInfosRequest) SetAscOrder(v bool) *ListJobSnapshotInfosRequest {
	s.AscOrder = &v
	return s
}

func (s *ListJobSnapshotInfosRequest) SetExtNodeIdList(v []*string) *ListJobSnapshotInfosRequest {
	s.ExtNodeIdList = v
	return s
}

func (s *ListJobSnapshotInfosRequest) SetFrom(v int64) *ListJobSnapshotInfosRequest {
	s.From = &v
	return s
}

func (s *ListJobSnapshotInfosRequest) SetInstanceIdList(v []*string) *ListJobSnapshotInfosRequest {
	s.InstanceIdList = v
	return s
}

func (s *ListJobSnapshotInfosRequest) SetJobOwnerList(v []*string) *ListJobSnapshotInfosRequest {
	s.JobOwnerList = v
	return s
}

func (s *ListJobSnapshotInfosRequest) SetPriorityList(v []*int64) *ListJobSnapshotInfosRequest {
	s.PriorityList = v
	return s
}

func (s *ListJobSnapshotInfosRequest) SetProjectList(v []*string) *ListJobSnapshotInfosRequest {
	s.ProjectList = v
	return s
}

func (s *ListJobSnapshotInfosRequest) SetQuotaNickname(v string) *ListJobSnapshotInfosRequest {
	s.QuotaNickname = &v
	return s
}

func (s *ListJobSnapshotInfosRequest) SetSignatureList(v []*string) *ListJobSnapshotInfosRequest {
	s.SignatureList = v
	return s
}

func (s *ListJobSnapshotInfosRequest) SetSortByList(v []*string) *ListJobSnapshotInfosRequest {
	s.SortByList = v
	return s
}

func (s *ListJobSnapshotInfosRequest) SetSortOrderList(v []*string) *ListJobSnapshotInfosRequest {
	s.SortOrderList = v
	return s
}

func (s *ListJobSnapshotInfosRequest) SetStatusList(v []*string) *ListJobSnapshotInfosRequest {
	s.StatusList = v
	return s
}

func (s *ListJobSnapshotInfosRequest) SetTo(v int64) *ListJobSnapshotInfosRequest {
	s.To = &v
	return s
}

func (s *ListJobSnapshotInfosRequest) SetTypeList(v []*string) *ListJobSnapshotInfosRequest {
	s.TypeList = v
	return s
}

func (s *ListJobSnapshotInfosRequest) SetOrderColumn(v string) *ListJobSnapshotInfosRequest {
	s.OrderColumn = &v
	return s
}

func (s *ListJobSnapshotInfosRequest) SetPageNumber(v int64) *ListJobSnapshotInfosRequest {
	s.PageNumber = &v
	return s
}

func (s *ListJobSnapshotInfosRequest) SetPageSize(v int64) *ListJobSnapshotInfosRequest {
	s.PageSize = &v
	return s
}

func (s *ListJobSnapshotInfosRequest) SetRegion(v string) *ListJobSnapshotInfosRequest {
	s.Region = &v
	return s
}

func (s *ListJobSnapshotInfosRequest) SetTenantId(v string) *ListJobSnapshotInfosRequest {
	s.TenantId = &v
	return s
}

type ListJobSnapshotInfosResponseBody struct {
	// The data returned.
	Data *ListJobSnapshotInfosResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// OBJECT_NOT_EXIST
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// this quota is not exist.
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// The HTTP status code.
	//
	// - 1xx: informational response. The request is received and is being processed.
	//
	// - 2xx: success. The request is successfully received, understood, and accepted by the server.
	//
	// - 3xx: redirection. The request is redirected, and further actions are required to complete the request.
	//
	// - 4xx: client error. The request contains invalid request parameters or syntaxes, or specific request conditions cannot be met.
	//
	// - 5xx: server error. The server cannot meet requirements due to other reasons.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0b87b7e716665825896565060e87a4
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListJobSnapshotInfosResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListJobSnapshotInfosResponseBody) GoString() string {
	return s.String()
}

func (s *ListJobSnapshotInfosResponseBody) SetData(v *ListJobSnapshotInfosResponseBodyData) *ListJobSnapshotInfosResponseBody {
	s.Data = v
	return s
}

func (s *ListJobSnapshotInfosResponseBody) SetErrorCode(v string) *ListJobSnapshotInfosResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBody) SetErrorMsg(v string) *ListJobSnapshotInfosResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBody) SetHttpCode(v int32) *ListJobSnapshotInfosResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBody) SetRequestId(v string) *ListJobSnapshotInfosResponseBody {
	s.RequestId = &v
	return s
}

type ListJobSnapshotInfosResponseBodyData struct {
	// The list of jobs snapshot information
	JobInfoList []*ListJobSnapshotInfosResponseBodyDataJobInfoList `json:"jobInfoList,omitempty" xml:"jobInfoList,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The total number of returned results.
	//
	// example:
	//
	// 123
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListJobSnapshotInfosResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListJobSnapshotInfosResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListJobSnapshotInfosResponseBodyData) SetJobInfoList(v []*ListJobSnapshotInfosResponseBodyDataJobInfoList) *ListJobSnapshotInfosResponseBodyData {
	s.JobInfoList = v
	return s
}

func (s *ListJobSnapshotInfosResponseBodyData) SetPageNumber(v int64) *ListJobSnapshotInfosResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBodyData) SetPageSize(v int64) *ListJobSnapshotInfosResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBodyData) SetTotalCount(v int64) *ListJobSnapshotInfosResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListJobSnapshotInfosResponseBodyDataJobInfoList struct {
	CpuRequest *int64 `json:"cpuRequest,omitempty" xml:"cpuRequest,omitempty"`
	// CPU usage of the job at the snapshot time. Unit: Core.
	//
	// example:
	//
	// 100
	CpuUsage               *int64   `json:"cpuUsage,omitempty" xml:"cpuUsage,omitempty"`
	CpuUsageToRequestRatio *float64 `json:"cpuUsageToRequestRatio,omitempty" xml:"cpuUsageToRequestRatio,omitempty"`
	// The ID of the upstream node.
	//
	// example:
	//
	// 76358164
	ExtNodeId *string `json:"extNodeId,omitempty" xml:"extNodeId,omitempty"`
	// The account ID of the task owner.
	//
	// example:
	//
	// duty_2
	ExtNodeOnDuty *string `json:"extNodeOnDuty,omitempty" xml:"extNodeOnDuty,omitempty"`
	// The upstream platform.
	//
	// example:
	//
	// Dataworks
	ExtPlantFrom *string `json:"extPlantFrom,omitempty" xml:"extPlantFrom,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// 20241028****jkl
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The account that commits the job.
	//
	// example:
	//
	// ALIYUN$7632***@aliyun.com
	JobOwner *string `json:"jobOwner,omitempty" xml:"jobOwner,omitempty"`
	// The type of the job.
	//
	// example:
	//
	// SQL
	JobType *string `json:"jobType,omitempty" xml:"jobType,omitempty"`
	// Not applicable.
	//
	// example:
	//
	// -1
	MaxCpuPct *float64 `json:"maxCpuPct,omitempty" xml:"maxCpuPct,omitempty"`
	// Not applicable.
	//
	// example:
	//
	// -1
	MaxMemoryPct  *float64 `json:"maxMemoryPct,omitempty" xml:"maxMemoryPct,omitempty"`
	MemoryRequest *int64   `json:"memoryRequest,omitempty" xml:"memoryRequest,omitempty"`
	// Memory usage of the job at the snapshot time. Unit: MB.
	//
	// example:
	//
	// 409600
	MemoryUsage               *int64   `json:"memoryUsage,omitempty" xml:"memoryUsage,omitempty"`
	MemoryUsageToRequestRatio *float64 `json:"memoryUsageToRequestRatio,omitempty" xml:"memoryUsageToRequestRatio,omitempty"`
	// The CPU usage ratio of the annual or monthly subscription job at the snapshot time (CPU usage / (reserved CPU guarantee + elastic reserved CPU)). This parameter is not available for pay-as-you-go jobs.
	//
	// example:
	//
	// 0.6
	MinCpuPct *float64 `json:"minCpuPct,omitempty" xml:"minCpuPct,omitempty"`
	// The memory usage ratio of the annual or monthly subscription job at the observation time (memory usage / (reserved memory guarantee + elastic reserved memory)). This parameter is not available for pay-as-you-go jobs.
	//
	// example:
	//
	// 0.6
	MinMemoryPct *float64 `json:"minMemoryPct,omitempty" xml:"minMemoryPct,omitempty"`
	// The priority of the job.
	//
	// example:
	//
	// 9
	Priority *int64 `json:"priority,omitempty" xml:"priority,omitempty"`
	// The name of the MaxCompute project.
	//
	// example:
	//
	// projectA
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
	// The nickname of the computing Quota used by the job.
	//
	// example:
	//
	// quota_A
	QuotaNickname *string `json:"quotaNickname,omitempty" xml:"quotaNickname,omitempty"`
	// The type of the quota.
	//
	// example:
	//
	// subscription
	QuotaType *string `json:"quotaType,omitempty" xml:"quotaType,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-chengdu
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The start time of the job.
	//
	// > The time when the job received the first batch of computing resources.
	//
	// example:
	//
	// 1736821805
	RunningAtTime *int64 `json:"runningAtTime,omitempty" xml:"runningAtTime,omitempty"`
	// The running duration, which is the duration from the runningAtTime to the snapshotTime of the job.  Unit: seconds (s).
	//
	// example:
	//
	// 43
	RunningTime *int64 `json:"runningTime,omitempty" xml:"runningTime,omitempty"`
	// The signature of the SQL job.
	//
	// example:
	//
	// signatureabc
	Signature *string `json:"signature,omitempty" xml:"signature,omitempty"`
	// The snapshot time.
	//
	// example:
	//
	// 1736821848
	SnapshotTime *int64 `json:"snapshotTime,omitempty" xml:"snapshotTime,omitempty"`
	// The snapshot status of the job.
	//
	// > The snapshot status is only running.
	//
	// example:
	//
	// running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The time when the job was committed.
	//
	// example:
	//
	// 1736821785
	SubmittedAtTime *int64 `json:"submittedAtTime,omitempty" xml:"submittedAtTime,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 213065738244354
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The interval from the time when the job was submitted to the snapshotTime .Unit: seconds (s).
	//
	// example:
	//
	// 63
	TotalTime *int64 `json:"totalTime,omitempty" xml:"totalTime,omitempty"`
	// The duration from the time the job is submitted to the time the job starts to run. Unit: seconds (s).
	//
	// example:
	//
	// 20
	WaitingTime *int64 `json:"waitingTime,omitempty" xml:"waitingTime,omitempty"`
}

func (s ListJobSnapshotInfosResponseBodyDataJobInfoList) String() string {
	return tea.Prettify(s)
}

func (s ListJobSnapshotInfosResponseBodyDataJobInfoList) GoString() string {
	return s.String()
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) SetCpuRequest(v int64) *ListJobSnapshotInfosResponseBodyDataJobInfoList {
	s.CpuRequest = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) SetCpuUsage(v int64) *ListJobSnapshotInfosResponseBodyDataJobInfoList {
	s.CpuUsage = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) SetCpuUsageToRequestRatio(v float64) *ListJobSnapshotInfosResponseBodyDataJobInfoList {
	s.CpuUsageToRequestRatio = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) SetExtNodeId(v string) *ListJobSnapshotInfosResponseBodyDataJobInfoList {
	s.ExtNodeId = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) SetExtNodeOnDuty(v string) *ListJobSnapshotInfosResponseBodyDataJobInfoList {
	s.ExtNodeOnDuty = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) SetExtPlantFrom(v string) *ListJobSnapshotInfosResponseBodyDataJobInfoList {
	s.ExtPlantFrom = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) SetInstanceId(v string) *ListJobSnapshotInfosResponseBodyDataJobInfoList {
	s.InstanceId = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) SetJobOwner(v string) *ListJobSnapshotInfosResponseBodyDataJobInfoList {
	s.JobOwner = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) SetJobType(v string) *ListJobSnapshotInfosResponseBodyDataJobInfoList {
	s.JobType = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) SetMaxCpuPct(v float64) *ListJobSnapshotInfosResponseBodyDataJobInfoList {
	s.MaxCpuPct = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) SetMaxMemoryPct(v float64) *ListJobSnapshotInfosResponseBodyDataJobInfoList {
	s.MaxMemoryPct = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) SetMemoryRequest(v int64) *ListJobSnapshotInfosResponseBodyDataJobInfoList {
	s.MemoryRequest = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) SetMemoryUsage(v int64) *ListJobSnapshotInfosResponseBodyDataJobInfoList {
	s.MemoryUsage = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) SetMemoryUsageToRequestRatio(v float64) *ListJobSnapshotInfosResponseBodyDataJobInfoList {
	s.MemoryUsageToRequestRatio = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) SetMinCpuPct(v float64) *ListJobSnapshotInfosResponseBodyDataJobInfoList {
	s.MinCpuPct = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) SetMinMemoryPct(v float64) *ListJobSnapshotInfosResponseBodyDataJobInfoList {
	s.MinMemoryPct = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) SetPriority(v int64) *ListJobSnapshotInfosResponseBodyDataJobInfoList {
	s.Priority = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) SetProject(v string) *ListJobSnapshotInfosResponseBodyDataJobInfoList {
	s.Project = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) SetQuotaNickname(v string) *ListJobSnapshotInfosResponseBodyDataJobInfoList {
	s.QuotaNickname = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) SetQuotaType(v string) *ListJobSnapshotInfosResponseBodyDataJobInfoList {
	s.QuotaType = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) SetRegion(v string) *ListJobSnapshotInfosResponseBodyDataJobInfoList {
	s.Region = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) SetRunningAtTime(v int64) *ListJobSnapshotInfosResponseBodyDataJobInfoList {
	s.RunningAtTime = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) SetRunningTime(v int64) *ListJobSnapshotInfosResponseBodyDataJobInfoList {
	s.RunningTime = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) SetSignature(v string) *ListJobSnapshotInfosResponseBodyDataJobInfoList {
	s.Signature = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) SetSnapshotTime(v int64) *ListJobSnapshotInfosResponseBodyDataJobInfoList {
	s.SnapshotTime = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) SetStatus(v string) *ListJobSnapshotInfosResponseBodyDataJobInfoList {
	s.Status = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) SetSubmittedAtTime(v int64) *ListJobSnapshotInfosResponseBodyDataJobInfoList {
	s.SubmittedAtTime = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) SetTenantId(v string) *ListJobSnapshotInfosResponseBodyDataJobInfoList {
	s.TenantId = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) SetTotalTime(v int64) *ListJobSnapshotInfosResponseBodyDataJobInfoList {
	s.TotalTime = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) SetWaitingTime(v int64) *ListJobSnapshotInfosResponseBodyDataJobInfoList {
	s.WaitingTime = &v
	return s
}

type ListJobSnapshotInfosResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListJobSnapshotInfosResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListJobSnapshotInfosResponse) String() string {
	return tea.Prettify(s)
}

func (s ListJobSnapshotInfosResponse) GoString() string {
	return s.String()
}

func (s *ListJobSnapshotInfosResponse) SetHeaders(v map[string]*string) *ListJobSnapshotInfosResponse {
	s.Headers = v
	return s
}

func (s *ListJobSnapshotInfosResponse) SetStatusCode(v int32) *ListJobSnapshotInfosResponse {
	s.StatusCode = &v
	return s
}

func (s *ListJobSnapshotInfosResponse) SetBody(v *ListJobSnapshotInfosResponseBody) *ListJobSnapshotInfosResponse {
	s.Body = v
	return s
}

type ListMmsDataSourcesRequest struct {
	// example:
	//
	// demo
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// cn-shanghai
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// example:
	//
	// BIGQUERY
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListMmsDataSourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMmsDataSourcesRequest) GoString() string {
	return s.String()
}

func (s *ListMmsDataSourcesRequest) SetName(v string) *ListMmsDataSourcesRequest {
	s.Name = &v
	return s
}

func (s *ListMmsDataSourcesRequest) SetPageNum(v int32) *ListMmsDataSourcesRequest {
	s.PageNum = &v
	return s
}

func (s *ListMmsDataSourcesRequest) SetPageSize(v int32) *ListMmsDataSourcesRequest {
	s.PageSize = &v
	return s
}

func (s *ListMmsDataSourcesRequest) SetRegion(v string) *ListMmsDataSourcesRequest {
	s.Region = &v
	return s
}

func (s *ListMmsDataSourcesRequest) SetType(v string) *ListMmsDataSourcesRequest {
	s.Type = &v
	return s
}

type ListMmsDataSourcesResponseBody struct {
	Data *ListMmsDataSourcesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// C1F7715F-D316-5AB6-BD02-5241083F4003
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListMmsDataSourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMmsDataSourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListMmsDataSourcesResponseBody) SetData(v *ListMmsDataSourcesResponseBodyData) *ListMmsDataSourcesResponseBody {
	s.Data = v
	return s
}

func (s *ListMmsDataSourcesResponseBody) SetRequestId(v string) *ListMmsDataSourcesResponseBody {
	s.RequestId = &v
	return s
}

type ListMmsDataSourcesResponseBodyData struct {
	ObjectList []*ListMmsDataSourcesResponseBodyDataObjectList `json:"objectList,omitempty" xml:"objectList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNum *int32 `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 9
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListMmsDataSourcesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListMmsDataSourcesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMmsDataSourcesResponseBodyData) SetObjectList(v []*ListMmsDataSourcesResponseBodyDataObjectList) *ListMmsDataSourcesResponseBodyData {
	s.ObjectList = v
	return s
}

func (s *ListMmsDataSourcesResponseBodyData) SetPageNum(v int32) *ListMmsDataSourcesResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *ListMmsDataSourcesResponseBodyData) SetPageSize(v int32) *ListMmsDataSourcesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListMmsDataSourcesResponseBodyData) SetTotal(v int32) *ListMmsDataSourcesResponseBodyData {
	s.Total = &v
	return s
}

type ListMmsDataSourcesResponseBodyDataObjectList struct {
	// example:
	//
	// true
	AgentIsOnline *bool                                                 `json:"agentIsOnline,omitempty" xml:"agentIsOnline,omitempty"`
	Config        []*ListMmsDataSourcesResponseBodyDataObjectListConfig `json:"config,omitempty" xml:"config,omitempty" type:"Repeated"`
	// example:
	//
	// 2024-12-17 09:29:58
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 3
	DbNum *int32 `json:"dbNum,omitempty" xml:"dbNum,omitempty"`
	// example:
	//
	// unexpected exception
	ErrMsg *string `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	// example:
	//
	// 2000015
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 2024-12-17 15:44:17
	LastUpdateTime *string `json:"lastUpdateTime,omitempty" xml:"lastUpdateTime,omitempty"`
	// example:
	//
	// demo
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// vpc-2zebqp6uojhdla46677tl:cn-beijing
	Networklink *string `json:"networklink,omitempty" xml:"networklink,omitempty"`
	// example:
	//
	// 10000000
	PartitionNum *int32 `json:"partitionNum,omitempty" xml:"partitionNum,omitempty"`
	// example:
	//
	// 2332
	PartitionsDoingNum *int32 `json:"partitionsDoingNum,omitempty" xml:"partitionsDoingNum,omitempty"`
	// example:
	//
	// 23
	PartitionsDoneNum *int32 `json:"partitionsDoneNum,omitempty" xml:"partitionsDoneNum,omitempty"`
	// example:
	//
	// 2323
	PartitionsFailedNum *int32 `json:"partitionsFailedNum,omitempty" xml:"partitionsFailedNum,omitempty"`
	// example:
	//
	// cn-beijing
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// example:
	//
	// 1000253
	ScanId *int64 `json:"scanId,omitempty" xml:"scanId,omitempty"`
	// example:
	//
	// STARTED
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 1000
	TableNum *int32 `json:"tableNum,omitempty" xml:"tableNum,omitempty"`
	// example:
	//
	// 18
	TablesDoingNum *int32 `json:"tablesDoingNum,omitempty" xml:"tablesDoingNum,omitempty"`
	// example:
	//
	// 2323
	TablesDoneNum *int32 `json:"tablesDoneNum,omitempty" xml:"tablesDoneNum,omitempty"`
	// example:
	//
	// 2
	TablesFailedNum *int32 `json:"tablesFailedNum,omitempty" xml:"tablesFailedNum,omitempty"`
	// example:
	//
	// 22
	TablesPartDoneNum *int32 `json:"tablesPartDoneNum,omitempty" xml:"tablesPartDoneNum,omitempty"`
	// example:
	//
	// BIGQUERY
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListMmsDataSourcesResponseBodyDataObjectList) String() string {
	return tea.Prettify(s)
}

func (s ListMmsDataSourcesResponseBodyDataObjectList) GoString() string {
	return s.String()
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) SetAgentIsOnline(v bool) *ListMmsDataSourcesResponseBodyDataObjectList {
	s.AgentIsOnline = &v
	return s
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) SetConfig(v []*ListMmsDataSourcesResponseBodyDataObjectListConfig) *ListMmsDataSourcesResponseBodyDataObjectList {
	s.Config = v
	return s
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) SetCreateTime(v string) *ListMmsDataSourcesResponseBodyDataObjectList {
	s.CreateTime = &v
	return s
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) SetDbNum(v int32) *ListMmsDataSourcesResponseBodyDataObjectList {
	s.DbNum = &v
	return s
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) SetErrMsg(v string) *ListMmsDataSourcesResponseBodyDataObjectList {
	s.ErrMsg = &v
	return s
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) SetId(v int64) *ListMmsDataSourcesResponseBodyDataObjectList {
	s.Id = &v
	return s
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) SetLastUpdateTime(v string) *ListMmsDataSourcesResponseBodyDataObjectList {
	s.LastUpdateTime = &v
	return s
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) SetName(v string) *ListMmsDataSourcesResponseBodyDataObjectList {
	s.Name = &v
	return s
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) SetNetworklink(v string) *ListMmsDataSourcesResponseBodyDataObjectList {
	s.Networklink = &v
	return s
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) SetPartitionNum(v int32) *ListMmsDataSourcesResponseBodyDataObjectList {
	s.PartitionNum = &v
	return s
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) SetPartitionsDoingNum(v int32) *ListMmsDataSourcesResponseBodyDataObjectList {
	s.PartitionsDoingNum = &v
	return s
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) SetPartitionsDoneNum(v int32) *ListMmsDataSourcesResponseBodyDataObjectList {
	s.PartitionsDoneNum = &v
	return s
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) SetPartitionsFailedNum(v int32) *ListMmsDataSourcesResponseBodyDataObjectList {
	s.PartitionsFailedNum = &v
	return s
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) SetRegion(v string) *ListMmsDataSourcesResponseBodyDataObjectList {
	s.Region = &v
	return s
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) SetScanId(v int64) *ListMmsDataSourcesResponseBodyDataObjectList {
	s.ScanId = &v
	return s
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) SetStatus(v string) *ListMmsDataSourcesResponseBodyDataObjectList {
	s.Status = &v
	return s
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) SetTableNum(v int32) *ListMmsDataSourcesResponseBodyDataObjectList {
	s.TableNum = &v
	return s
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) SetTablesDoingNum(v int32) *ListMmsDataSourcesResponseBodyDataObjectList {
	s.TablesDoingNum = &v
	return s
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) SetTablesDoneNum(v int32) *ListMmsDataSourcesResponseBodyDataObjectList {
	s.TablesDoneNum = &v
	return s
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) SetTablesFailedNum(v int32) *ListMmsDataSourcesResponseBodyDataObjectList {
	s.TablesFailedNum = &v
	return s
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) SetTablesPartDoneNum(v int32) *ListMmsDataSourcesResponseBodyDataObjectList {
	s.TablesPartDoneNum = &v
	return s
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) SetType(v string) *ListMmsDataSourcesResponseBodyDataObjectList {
	s.Type = &v
	return s
}

type ListMmsDataSourcesResponseBodyDataObjectListConfig struct {
	Desc  *string   `json:"desc,omitempty" xml:"desc,omitempty"`
	Enums []*string `json:"enums,omitempty" xml:"enums,omitempty" type:"Repeated"`
	// example:
	//
	// basic_group
	Group *string `json:"group,omitempty" xml:"group,omitempty"`
	// example:
	//
	// bigquery.range.partition.migrate.type
	Key  *string `json:"key,omitempty" xml:"key,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// Cluster or Partition
	PlaceHolder *string `json:"placeHolder,omitempty" xml:"placeHolder,omitempty"`
	// example:
	//
	// true
	Required *bool `json:"required,omitempty" xml:"required,omitempty"`
	// example:
	//
	// .keytab
	SubType *string `json:"subType,omitempty" xml:"subType,omitempty"`
	// example:
	//
	// string
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// Partition
	Value interface{} `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListMmsDataSourcesResponseBodyDataObjectListConfig) String() string {
	return tea.Prettify(s)
}

func (s ListMmsDataSourcesResponseBodyDataObjectListConfig) GoString() string {
	return s.String()
}

func (s *ListMmsDataSourcesResponseBodyDataObjectListConfig) SetDesc(v string) *ListMmsDataSourcesResponseBodyDataObjectListConfig {
	s.Desc = &v
	return s
}

func (s *ListMmsDataSourcesResponseBodyDataObjectListConfig) SetEnums(v []*string) *ListMmsDataSourcesResponseBodyDataObjectListConfig {
	s.Enums = v
	return s
}

func (s *ListMmsDataSourcesResponseBodyDataObjectListConfig) SetGroup(v string) *ListMmsDataSourcesResponseBodyDataObjectListConfig {
	s.Group = &v
	return s
}

func (s *ListMmsDataSourcesResponseBodyDataObjectListConfig) SetKey(v string) *ListMmsDataSourcesResponseBodyDataObjectListConfig {
	s.Key = &v
	return s
}

func (s *ListMmsDataSourcesResponseBodyDataObjectListConfig) SetName(v string) *ListMmsDataSourcesResponseBodyDataObjectListConfig {
	s.Name = &v
	return s
}

func (s *ListMmsDataSourcesResponseBodyDataObjectListConfig) SetPlaceHolder(v string) *ListMmsDataSourcesResponseBodyDataObjectListConfig {
	s.PlaceHolder = &v
	return s
}

func (s *ListMmsDataSourcesResponseBodyDataObjectListConfig) SetRequired(v bool) *ListMmsDataSourcesResponseBodyDataObjectListConfig {
	s.Required = &v
	return s
}

func (s *ListMmsDataSourcesResponseBodyDataObjectListConfig) SetSubType(v string) *ListMmsDataSourcesResponseBodyDataObjectListConfig {
	s.SubType = &v
	return s
}

func (s *ListMmsDataSourcesResponseBodyDataObjectListConfig) SetType(v string) *ListMmsDataSourcesResponseBodyDataObjectListConfig {
	s.Type = &v
	return s
}

func (s *ListMmsDataSourcesResponseBodyDataObjectListConfig) SetValue(v interface{}) *ListMmsDataSourcesResponseBodyDataObjectListConfig {
	s.Value = v
	return s
}

type ListMmsDataSourcesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMmsDataSourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMmsDataSourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMmsDataSourcesResponse) GoString() string {
	return s.String()
}

func (s *ListMmsDataSourcesResponse) SetHeaders(v map[string]*string) *ListMmsDataSourcesResponse {
	s.Headers = v
	return s
}

func (s *ListMmsDataSourcesResponse) SetStatusCode(v int32) *ListMmsDataSourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMmsDataSourcesResponse) SetBody(v *ListMmsDataSourcesResponseBody) *ListMmsDataSourcesResponse {
	s.Body = v
	return s
}

type ListMmsDbsRequest struct {
	// example:
	//
	// demo
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32                   `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Sorter   *ListMmsDbsRequestSorter `json:"sorter,omitempty" xml:"sorter,omitempty" type:"Struct"`
	// example:
	//
	// STARTED
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListMmsDbsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMmsDbsRequest) GoString() string {
	return s.String()
}

func (s *ListMmsDbsRequest) SetName(v string) *ListMmsDbsRequest {
	s.Name = &v
	return s
}

func (s *ListMmsDbsRequest) SetPageNum(v int32) *ListMmsDbsRequest {
	s.PageNum = &v
	return s
}

func (s *ListMmsDbsRequest) SetPageSize(v int32) *ListMmsDbsRequest {
	s.PageSize = &v
	return s
}

func (s *ListMmsDbsRequest) SetSorter(v *ListMmsDbsRequestSorter) *ListMmsDbsRequest {
	s.Sorter = v
	return s
}

func (s *ListMmsDbsRequest) SetStatus(v string) *ListMmsDbsRequest {
	s.Status = &v
	return s
}

type ListMmsDbsRequestSorter struct {
	// example:
	//
	// desc
	NumRows *string `json:"numRows,omitempty" xml:"numRows,omitempty"`
	// example:
	//
	// asc
	Size *string `json:"size,omitempty" xml:"size,omitempty"`
	// example:
	//
	// 2024-12-17 15:44:17
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s ListMmsDbsRequestSorter) String() string {
	return tea.Prettify(s)
}

func (s ListMmsDbsRequestSorter) GoString() string {
	return s.String()
}

func (s *ListMmsDbsRequestSorter) SetNumRows(v string) *ListMmsDbsRequestSorter {
	s.NumRows = &v
	return s
}

func (s *ListMmsDbsRequestSorter) SetSize(v string) *ListMmsDbsRequestSorter {
	s.Size = &v
	return s
}

func (s *ListMmsDbsRequestSorter) SetUpdateTime(v string) *ListMmsDbsRequestSorter {
	s.UpdateTime = &v
	return s
}

type ListMmsDbsShrinkRequest struct {
	// example:
	//
	// demo
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	// example:
	//
	// 10
	PageSize     *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	SorterShrink *string `json:"sorter,omitempty" xml:"sorter,omitempty"`
	// example:
	//
	// STARTED
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListMmsDbsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMmsDbsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListMmsDbsShrinkRequest) SetName(v string) *ListMmsDbsShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListMmsDbsShrinkRequest) SetPageNum(v int32) *ListMmsDbsShrinkRequest {
	s.PageNum = &v
	return s
}

func (s *ListMmsDbsShrinkRequest) SetPageSize(v int32) *ListMmsDbsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListMmsDbsShrinkRequest) SetSorterShrink(v string) *ListMmsDbsShrinkRequest {
	s.SorterShrink = &v
	return s
}

func (s *ListMmsDbsShrinkRequest) SetStatus(v string) *ListMmsDbsShrinkRequest {
	s.Status = &v
	return s
}

type ListMmsDbsResponseBody struct {
	Data *ListMmsDbsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// CF3F9978-260F-5204-94BE-30A4E6B54443
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListMmsDbsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMmsDbsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMmsDbsResponseBody) SetData(v *ListMmsDbsResponseBodyData) *ListMmsDbsResponseBody {
	s.Data = v
	return s
}

func (s *ListMmsDbsResponseBody) SetRequestId(v string) *ListMmsDbsResponseBody {
	s.RequestId = &v
	return s
}

type ListMmsDbsResponseBodyData struct {
	ObjectList []*ListMmsDbsResponseBodyDataObjectList `json:"objectList,omitempty" xml:"objectList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNum *int32 `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 13
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListMmsDbsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListMmsDbsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMmsDbsResponseBodyData) SetObjectList(v []*ListMmsDbsResponseBodyDataObjectList) *ListMmsDbsResponseBodyData {
	s.ObjectList = v
	return s
}

func (s *ListMmsDbsResponseBodyData) SetPageNum(v int32) *ListMmsDbsResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *ListMmsDbsResponseBodyData) SetPageSize(v int32) *ListMmsDbsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListMmsDbsResponseBodyData) SetTotal(v int32) *ListMmsDbsResponseBodyData {
	s.Total = &v
	return s
}

type ListMmsDbsResponseBodyDataObjectList struct {
	// example:
	//
	// 2024-12-17 15:44:42
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// false
	Deleted *bool `json:"deleted,omitempty" xml:"deleted,omitempty"`
	// example:
	//
	// for mms test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// {}
	Extra *string `json:"extra,omitempty" xml:"extra,omitempty"`
	// example:
	//
	// 1530
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Last DDL Time
	//
	// example:
	//
	// 2024-12-17 15:44:42
	LastDdlTime *string `json:"lastDdlTime,omitempty" xml:"lastDdlTime,omitempty"`
	// example:
	//
	// hdfs://master-1-1.c-6fc187819ed6bae0.cn-shanghai.emr.aliyuncs.com:9000/user/hive/warehouse
	Location *string `json:"location,omitempty" xml:"location,omitempty"`
	// example:
	//
	// mms_test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 23232
	NumRows *int64 `json:"numRows,omitempty" xml:"numRows,omitempty"`
	// example:
	//
	// xxx@yy.com
	Owner *string `json:"owner,omitempty" xml:"owner,omitempty"`
	// example:
	//
	// 1000
	Partitions *int32 `json:"partitions,omitempty" xml:"partitions,omitempty"`
	// example:
	//
	// 400
	PartitionsDoing *int32 `json:"partitionsDoing,omitempty" xml:"partitionsDoing,omitempty"`
	// example:
	//
	// 200
	PartitionsDone *int32 `json:"partitionsDone,omitempty" xml:"partitionsDone,omitempty"`
	// example:
	//
	// 200
	PartitionsFailed *int32 `json:"partitionsFailed,omitempty" xml:"partitionsFailed,omitempty"`
	// example:
	//
	// 2342342
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// example:
	//
	// 2000015
	SourceId *int64 `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// example:
	//
	// demo
	SourceName *string `json:"sourceName,omitempty" xml:"sourceName,omitempty"`
	// example:
	//
	// DOING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 100
	Tables *int32 `json:"tables,omitempty" xml:"tables,omitempty"`
	// example:
	//
	// 20
	TablesDoing *int32 `json:"tablesDoing,omitempty" xml:"tablesDoing,omitempty"`
	// example:
	//
	// 20
	TablesDone *int32 `json:"tablesDone,omitempty" xml:"tablesDone,omitempty"`
	// example:
	//
	// 20
	TablesFailed *int32 `json:"tablesFailed,omitempty" xml:"tablesFailed,omitempty"`
	// example:
	//
	// 20
	TablesPartDone *int32 `json:"tablesPartDone,omitempty" xml:"tablesPartDone,omitempty"`
	// example:
	//
	// 2024-12-17 15:44:42
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// example:
	//
	// true
	Updated *bool `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s ListMmsDbsResponseBodyDataObjectList) String() string {
	return tea.Prettify(s)
}

func (s ListMmsDbsResponseBodyDataObjectList) GoString() string {
	return s.String()
}

func (s *ListMmsDbsResponseBodyDataObjectList) SetCreateTime(v string) *ListMmsDbsResponseBodyDataObjectList {
	s.CreateTime = &v
	return s
}

func (s *ListMmsDbsResponseBodyDataObjectList) SetDeleted(v bool) *ListMmsDbsResponseBodyDataObjectList {
	s.Deleted = &v
	return s
}

func (s *ListMmsDbsResponseBodyDataObjectList) SetDescription(v string) *ListMmsDbsResponseBodyDataObjectList {
	s.Description = &v
	return s
}

func (s *ListMmsDbsResponseBodyDataObjectList) SetExtra(v string) *ListMmsDbsResponseBodyDataObjectList {
	s.Extra = &v
	return s
}

func (s *ListMmsDbsResponseBodyDataObjectList) SetId(v int64) *ListMmsDbsResponseBodyDataObjectList {
	s.Id = &v
	return s
}

func (s *ListMmsDbsResponseBodyDataObjectList) SetLastDdlTime(v string) *ListMmsDbsResponseBodyDataObjectList {
	s.LastDdlTime = &v
	return s
}

func (s *ListMmsDbsResponseBodyDataObjectList) SetLocation(v string) *ListMmsDbsResponseBodyDataObjectList {
	s.Location = &v
	return s
}

func (s *ListMmsDbsResponseBodyDataObjectList) SetName(v string) *ListMmsDbsResponseBodyDataObjectList {
	s.Name = &v
	return s
}

func (s *ListMmsDbsResponseBodyDataObjectList) SetNumRows(v int64) *ListMmsDbsResponseBodyDataObjectList {
	s.NumRows = &v
	return s
}

func (s *ListMmsDbsResponseBodyDataObjectList) SetOwner(v string) *ListMmsDbsResponseBodyDataObjectList {
	s.Owner = &v
	return s
}

func (s *ListMmsDbsResponseBodyDataObjectList) SetPartitions(v int32) *ListMmsDbsResponseBodyDataObjectList {
	s.Partitions = &v
	return s
}

func (s *ListMmsDbsResponseBodyDataObjectList) SetPartitionsDoing(v int32) *ListMmsDbsResponseBodyDataObjectList {
	s.PartitionsDoing = &v
	return s
}

func (s *ListMmsDbsResponseBodyDataObjectList) SetPartitionsDone(v int32) *ListMmsDbsResponseBodyDataObjectList {
	s.PartitionsDone = &v
	return s
}

func (s *ListMmsDbsResponseBodyDataObjectList) SetPartitionsFailed(v int32) *ListMmsDbsResponseBodyDataObjectList {
	s.PartitionsFailed = &v
	return s
}

func (s *ListMmsDbsResponseBodyDataObjectList) SetSize(v int64) *ListMmsDbsResponseBodyDataObjectList {
	s.Size = &v
	return s
}

func (s *ListMmsDbsResponseBodyDataObjectList) SetSourceId(v int64) *ListMmsDbsResponseBodyDataObjectList {
	s.SourceId = &v
	return s
}

func (s *ListMmsDbsResponseBodyDataObjectList) SetSourceName(v string) *ListMmsDbsResponseBodyDataObjectList {
	s.SourceName = &v
	return s
}

func (s *ListMmsDbsResponseBodyDataObjectList) SetStatus(v string) *ListMmsDbsResponseBodyDataObjectList {
	s.Status = &v
	return s
}

func (s *ListMmsDbsResponseBodyDataObjectList) SetTables(v int32) *ListMmsDbsResponseBodyDataObjectList {
	s.Tables = &v
	return s
}

func (s *ListMmsDbsResponseBodyDataObjectList) SetTablesDoing(v int32) *ListMmsDbsResponseBodyDataObjectList {
	s.TablesDoing = &v
	return s
}

func (s *ListMmsDbsResponseBodyDataObjectList) SetTablesDone(v int32) *ListMmsDbsResponseBodyDataObjectList {
	s.TablesDone = &v
	return s
}

func (s *ListMmsDbsResponseBodyDataObjectList) SetTablesFailed(v int32) *ListMmsDbsResponseBodyDataObjectList {
	s.TablesFailed = &v
	return s
}

func (s *ListMmsDbsResponseBodyDataObjectList) SetTablesPartDone(v int32) *ListMmsDbsResponseBodyDataObjectList {
	s.TablesPartDone = &v
	return s
}

func (s *ListMmsDbsResponseBodyDataObjectList) SetUpdateTime(v string) *ListMmsDbsResponseBodyDataObjectList {
	s.UpdateTime = &v
	return s
}

func (s *ListMmsDbsResponseBodyDataObjectList) SetUpdated(v bool) *ListMmsDbsResponseBodyDataObjectList {
	s.Updated = &v
	return s
}

type ListMmsDbsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMmsDbsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMmsDbsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMmsDbsResponse) GoString() string {
	return s.String()
}

func (s *ListMmsDbsResponse) SetHeaders(v map[string]*string) *ListMmsDbsResponse {
	s.Headers = v
	return s
}

func (s *ListMmsDbsResponse) SetStatusCode(v int32) *ListMmsDbsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMmsDbsResponse) SetBody(v *ListMmsDbsResponseBody) *ListMmsDbsResponse {
	s.Body = v
	return s
}

type ListMmsJobsRequest struct {
	Sorter *ListMmsJobsRequestSorter `json:"sorter,omitempty" xml:"sorter,omitempty" type:"Struct"`
	// example:
	//
	// mms_test
	DstDbName *string `json:"dstDbName,omitempty" xml:"dstDbName,omitempty"`
	// example:
	//
	// test_table_1
	DstTableName *string `json:"dstTableName,omitempty" xml:"dstTableName,omitempty"`
	// example:
	//
	// demo
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// test_db_1
	SrcDbName *string `json:"srcDbName,omitempty" xml:"srcDbName,omitempty"`
	// example:
	//
	// test_table_1
	SrcTableName *string `json:"srcTableName,omitempty" xml:"srcTableName,omitempty"`
	// example:
	//
	// DOING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// false
	Stopped *int64 `json:"stopped,omitempty" xml:"stopped,omitempty"`
}

func (s ListMmsJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMmsJobsRequest) GoString() string {
	return s.String()
}

func (s *ListMmsJobsRequest) SetSorter(v *ListMmsJobsRequestSorter) *ListMmsJobsRequest {
	s.Sorter = v
	return s
}

func (s *ListMmsJobsRequest) SetDstDbName(v string) *ListMmsJobsRequest {
	s.DstDbName = &v
	return s
}

func (s *ListMmsJobsRequest) SetDstTableName(v string) *ListMmsJobsRequest {
	s.DstTableName = &v
	return s
}

func (s *ListMmsJobsRequest) SetName(v string) *ListMmsJobsRequest {
	s.Name = &v
	return s
}

func (s *ListMmsJobsRequest) SetPageNum(v int32) *ListMmsJobsRequest {
	s.PageNum = &v
	return s
}

func (s *ListMmsJobsRequest) SetPageSize(v int32) *ListMmsJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListMmsJobsRequest) SetSrcDbName(v string) *ListMmsJobsRequest {
	s.SrcDbName = &v
	return s
}

func (s *ListMmsJobsRequest) SetSrcTableName(v string) *ListMmsJobsRequest {
	s.SrcTableName = &v
	return s
}

func (s *ListMmsJobsRequest) SetStatus(v string) *ListMmsJobsRequest {
	s.Status = &v
	return s
}

func (s *ListMmsJobsRequest) SetStopped(v int64) *ListMmsJobsRequest {
	s.Stopped = &v
	return s
}

type ListMmsJobsRequestSorter struct {
	// example:
	//
	// desc
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListMmsJobsRequestSorter) String() string {
	return tea.Prettify(s)
}

func (s ListMmsJobsRequestSorter) GoString() string {
	return s.String()
}

func (s *ListMmsJobsRequestSorter) SetStatus(v string) *ListMmsJobsRequestSorter {
	s.Status = &v
	return s
}

type ListMmsJobsResponseBody struct {
	Data *ListMmsJobsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 1112E7C7-C65F-57A2-A7C7-3B178AA257B6
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListMmsJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMmsJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMmsJobsResponseBody) SetData(v *ListMmsJobsResponseBodyData) *ListMmsJobsResponseBody {
	s.Data = v
	return s
}

func (s *ListMmsJobsResponseBody) SetRequestId(v string) *ListMmsJobsResponseBody {
	s.RequestId = &v
	return s
}

type ListMmsJobsResponseBodyData struct {
	ObjectList []*ListMmsJobsResponseBodyDataObjectList `json:"objectList,omitempty" xml:"objectList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNum *int32 `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 100
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListMmsJobsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListMmsJobsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMmsJobsResponseBodyData) SetObjectList(v []*ListMmsJobsResponseBodyDataObjectList) *ListMmsJobsResponseBodyData {
	s.ObjectList = v
	return s
}

func (s *ListMmsJobsResponseBodyData) SetPageNum(v int32) *ListMmsJobsResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *ListMmsJobsResponseBodyData) SetPageSize(v int32) *ListMmsJobsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListMmsJobsResponseBodyData) SetTotal(v int32) *ListMmsJobsResponseBodyData {
	s.Total = &v
	return s
}

type ListMmsJobsResponseBodyDataObjectList struct {
	Config *ListMmsJobsResponseBodyDataObjectListConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// example:
	//
	// 2024-12-17 15:44:17
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 196
	DbId *int64 `json:"dbId,omitempty" xml:"dbId,omitempty"`
	// example:
	//
	// mms_test
	DstDbName *string `json:"dstDbName,omitempty" xml:"dstDbName,omitempty"`
	// example:
	//
	// test_table_1
	DstSchemaName *string `json:"dstSchemaName,omitempty" xml:"dstSchemaName,omitempty"`
	Eta           *string `json:"eta,omitempty" xml:"eta,omitempty"`
	// example:
	//
	// 18
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// migrate_db_1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 2000015
	SourceId *int64 `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// example:
	//
	// demo
	SourceName *string `json:"sourceName,omitempty" xml:"sourceName,omitempty"`
	// example:
	//
	// test_db_1
	SrcDbName *string `json:"srcDbName,omitempty" xml:"srcDbName,omitempty"`
	// example:
	//
	// test_table_1
	SrcSchemaName *string `json:"srcSchemaName,omitempty" xml:"srcSchemaName,omitempty"`
	// example:
	//
	// DONE
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// false
	Stopped *bool `json:"stopped,omitempty" xml:"stopped,omitempty"`
	// example:
	//
	// 10
	TaskDone *int32 `json:"taskDone,omitempty" xml:"taskDone,omitempty"`
	// example:
	//
	// 10
	TaskNum *int32 `json:"taskNum,omitempty" xml:"taskNum,omitempty"`
	// example:
	//
	// Tables
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListMmsJobsResponseBodyDataObjectList) String() string {
	return tea.Prettify(s)
}

func (s ListMmsJobsResponseBodyDataObjectList) GoString() string {
	return s.String()
}

func (s *ListMmsJobsResponseBodyDataObjectList) SetConfig(v *ListMmsJobsResponseBodyDataObjectListConfig) *ListMmsJobsResponseBodyDataObjectList {
	s.Config = v
	return s
}

func (s *ListMmsJobsResponseBodyDataObjectList) SetCreateTime(v string) *ListMmsJobsResponseBodyDataObjectList {
	s.CreateTime = &v
	return s
}

func (s *ListMmsJobsResponseBodyDataObjectList) SetDbId(v int64) *ListMmsJobsResponseBodyDataObjectList {
	s.DbId = &v
	return s
}

func (s *ListMmsJobsResponseBodyDataObjectList) SetDstDbName(v string) *ListMmsJobsResponseBodyDataObjectList {
	s.DstDbName = &v
	return s
}

func (s *ListMmsJobsResponseBodyDataObjectList) SetDstSchemaName(v string) *ListMmsJobsResponseBodyDataObjectList {
	s.DstSchemaName = &v
	return s
}

func (s *ListMmsJobsResponseBodyDataObjectList) SetEta(v string) *ListMmsJobsResponseBodyDataObjectList {
	s.Eta = &v
	return s
}

func (s *ListMmsJobsResponseBodyDataObjectList) SetId(v int64) *ListMmsJobsResponseBodyDataObjectList {
	s.Id = &v
	return s
}

func (s *ListMmsJobsResponseBodyDataObjectList) SetName(v string) *ListMmsJobsResponseBodyDataObjectList {
	s.Name = &v
	return s
}

func (s *ListMmsJobsResponseBodyDataObjectList) SetSourceId(v int64) *ListMmsJobsResponseBodyDataObjectList {
	s.SourceId = &v
	return s
}

func (s *ListMmsJobsResponseBodyDataObjectList) SetSourceName(v string) *ListMmsJobsResponseBodyDataObjectList {
	s.SourceName = &v
	return s
}

func (s *ListMmsJobsResponseBodyDataObjectList) SetSrcDbName(v string) *ListMmsJobsResponseBodyDataObjectList {
	s.SrcDbName = &v
	return s
}

func (s *ListMmsJobsResponseBodyDataObjectList) SetSrcSchemaName(v string) *ListMmsJobsResponseBodyDataObjectList {
	s.SrcSchemaName = &v
	return s
}

func (s *ListMmsJobsResponseBodyDataObjectList) SetStatus(v string) *ListMmsJobsResponseBodyDataObjectList {
	s.Status = &v
	return s
}

func (s *ListMmsJobsResponseBodyDataObjectList) SetStopped(v bool) *ListMmsJobsResponseBodyDataObjectList {
	s.Stopped = &v
	return s
}

func (s *ListMmsJobsResponseBodyDataObjectList) SetTaskDone(v int32) *ListMmsJobsResponseBodyDataObjectList {
	s.TaskDone = &v
	return s
}

func (s *ListMmsJobsResponseBodyDataObjectList) SetTaskNum(v int32) *ListMmsJobsResponseBodyDataObjectList {
	s.TaskNum = &v
	return s
}

func (s *ListMmsJobsResponseBodyDataObjectList) SetType(v string) *ListMmsJobsResponseBodyDataObjectList {
	s.Type = &v
	return s
}

type ListMmsJobsResponseBodyDataObjectListConfig struct {
	ColumnMapping      map[string]*string     `json:"columnMapping,omitempty" xml:"columnMapping,omitempty"`
	EnableVerification *bool                  `json:"enableVerification,omitempty" xml:"enableVerification,omitempty"`
	Increment          *bool                  `json:"increment,omitempty" xml:"increment,omitempty"`
	Others             map[string]interface{} `json:"others,omitempty" xml:"others,omitempty"`
	PartitionFilters   map[string]*string     `json:"partitionFilters,omitempty" xml:"partitionFilters,omitempty"`
	Partitions         []*int64               `json:"partitions,omitempty" xml:"partitions,omitempty" type:"Repeated"`
	SchemaOnly         *bool                  `json:"schemaOnly,omitempty" xml:"schemaOnly,omitempty"`
	TableBlackList     []*string              `json:"tableBlackList,omitempty" xml:"tableBlackList,omitempty" type:"Repeated"`
	TableMapping       map[string]*string     `json:"tableMapping,omitempty" xml:"tableMapping,omitempty"`
	TableWhiteList     []*string              `json:"tableWhiteList,omitempty" xml:"tableWhiteList,omitempty" type:"Repeated"`
	Tables             []*string              `json:"tables,omitempty" xml:"tables,omitempty" type:"Repeated"`
	TaskType           *string                `json:"taskType,omitempty" xml:"taskType,omitempty"`
	TunnelQuota        *string                `json:"tunnelQuota,omitempty" xml:"tunnelQuota,omitempty"`
}

func (s ListMmsJobsResponseBodyDataObjectListConfig) String() string {
	return tea.Prettify(s)
}

func (s ListMmsJobsResponseBodyDataObjectListConfig) GoString() string {
	return s.String()
}

func (s *ListMmsJobsResponseBodyDataObjectListConfig) SetColumnMapping(v map[string]*string) *ListMmsJobsResponseBodyDataObjectListConfig {
	s.ColumnMapping = v
	return s
}

func (s *ListMmsJobsResponseBodyDataObjectListConfig) SetEnableVerification(v bool) *ListMmsJobsResponseBodyDataObjectListConfig {
	s.EnableVerification = &v
	return s
}

func (s *ListMmsJobsResponseBodyDataObjectListConfig) SetIncrement(v bool) *ListMmsJobsResponseBodyDataObjectListConfig {
	s.Increment = &v
	return s
}

func (s *ListMmsJobsResponseBodyDataObjectListConfig) SetOthers(v map[string]interface{}) *ListMmsJobsResponseBodyDataObjectListConfig {
	s.Others = v
	return s
}

func (s *ListMmsJobsResponseBodyDataObjectListConfig) SetPartitionFilters(v map[string]*string) *ListMmsJobsResponseBodyDataObjectListConfig {
	s.PartitionFilters = v
	return s
}

func (s *ListMmsJobsResponseBodyDataObjectListConfig) SetPartitions(v []*int64) *ListMmsJobsResponseBodyDataObjectListConfig {
	s.Partitions = v
	return s
}

func (s *ListMmsJobsResponseBodyDataObjectListConfig) SetSchemaOnly(v bool) *ListMmsJobsResponseBodyDataObjectListConfig {
	s.SchemaOnly = &v
	return s
}

func (s *ListMmsJobsResponseBodyDataObjectListConfig) SetTableBlackList(v []*string) *ListMmsJobsResponseBodyDataObjectListConfig {
	s.TableBlackList = v
	return s
}

func (s *ListMmsJobsResponseBodyDataObjectListConfig) SetTableMapping(v map[string]*string) *ListMmsJobsResponseBodyDataObjectListConfig {
	s.TableMapping = v
	return s
}

func (s *ListMmsJobsResponseBodyDataObjectListConfig) SetTableWhiteList(v []*string) *ListMmsJobsResponseBodyDataObjectListConfig {
	s.TableWhiteList = v
	return s
}

func (s *ListMmsJobsResponseBodyDataObjectListConfig) SetTables(v []*string) *ListMmsJobsResponseBodyDataObjectListConfig {
	s.Tables = v
	return s
}

func (s *ListMmsJobsResponseBodyDataObjectListConfig) SetTaskType(v string) *ListMmsJobsResponseBodyDataObjectListConfig {
	s.TaskType = &v
	return s
}

func (s *ListMmsJobsResponseBodyDataObjectListConfig) SetTunnelQuota(v string) *ListMmsJobsResponseBodyDataObjectListConfig {
	s.TunnelQuota = &v
	return s
}

type ListMmsJobsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMmsJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMmsJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMmsJobsResponse) GoString() string {
	return s.String()
}

func (s *ListMmsJobsResponse) SetHeaders(v map[string]*string) *ListMmsJobsResponse {
	s.Headers = v
	return s
}

func (s *ListMmsJobsResponse) SetStatusCode(v int32) *ListMmsJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMmsJobsResponse) SetBody(v *ListMmsJobsResponseBody) *ListMmsJobsResponse {
	s.Body = v
	return s
}

type ListMmsPartitionsRequest struct {
	Sorter *ListMmsPartitionsRequestSorter `json:"sorter,omitempty" xml:"sorter,omitempty" type:"Struct"`
	// example:
	//
	// 2
	DbId *int64 `json:"dbId,omitempty" xml:"dbId,omitempty"`
	// example:
	//
	// d1
	DbName *string `json:"dbName,omitempty" xml:"dbName,omitempty"`
	// example:
	//
	// 2024-12-17 19:44:42
	LastDdlTimeEnd *string `json:"lastDdlTimeEnd,omitempty" xml:"lastDdlTimeEnd,omitempty"`
	// example:
	//
	// 2024-12-17 15:44:42
	LastDdlTimeStart *string `json:"lastDdlTimeStart,omitempty" xml:"lastDdlTimeStart,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	// example:
	//
	// 100
	PageSize *int32    `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Status   []*string `json:"status,omitempty" xml:"status,omitempty" type:"Repeated"`
	// example:
	//
	// t1
	TableName *string `json:"tableName,omitempty" xml:"tableName,omitempty"`
	// example:
	//
	// false
	Updated *bool `json:"updated,omitempty" xml:"updated,omitempty"`
	// example:
	//
	// p1=1/p2=abc
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListMmsPartitionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMmsPartitionsRequest) GoString() string {
	return s.String()
}

func (s *ListMmsPartitionsRequest) SetSorter(v *ListMmsPartitionsRequestSorter) *ListMmsPartitionsRequest {
	s.Sorter = v
	return s
}

func (s *ListMmsPartitionsRequest) SetDbId(v int64) *ListMmsPartitionsRequest {
	s.DbId = &v
	return s
}

func (s *ListMmsPartitionsRequest) SetDbName(v string) *ListMmsPartitionsRequest {
	s.DbName = &v
	return s
}

func (s *ListMmsPartitionsRequest) SetLastDdlTimeEnd(v string) *ListMmsPartitionsRequest {
	s.LastDdlTimeEnd = &v
	return s
}

func (s *ListMmsPartitionsRequest) SetLastDdlTimeStart(v string) *ListMmsPartitionsRequest {
	s.LastDdlTimeStart = &v
	return s
}

func (s *ListMmsPartitionsRequest) SetPageNum(v int32) *ListMmsPartitionsRequest {
	s.PageNum = &v
	return s
}

func (s *ListMmsPartitionsRequest) SetPageSize(v int32) *ListMmsPartitionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListMmsPartitionsRequest) SetStatus(v []*string) *ListMmsPartitionsRequest {
	s.Status = v
	return s
}

func (s *ListMmsPartitionsRequest) SetTableName(v string) *ListMmsPartitionsRequest {
	s.TableName = &v
	return s
}

func (s *ListMmsPartitionsRequest) SetUpdated(v bool) *ListMmsPartitionsRequest {
	s.Updated = &v
	return s
}

func (s *ListMmsPartitionsRequest) SetValue(v string) *ListMmsPartitionsRequest {
	s.Value = &v
	return s
}

type ListMmsPartitionsRequestSorter struct {
	// example:
	//
	// desc
	LastDdlTime *string `json:"lastDdlTime,omitempty" xml:"lastDdlTime,omitempty"`
	// example:
	//
	// desc
	NumRows *string `json:"numRows,omitempty" xml:"numRows,omitempty"`
	// example:
	//
	// asc
	Size *string `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ListMmsPartitionsRequestSorter) String() string {
	return tea.Prettify(s)
}

func (s ListMmsPartitionsRequestSorter) GoString() string {
	return s.String()
}

func (s *ListMmsPartitionsRequestSorter) SetLastDdlTime(v string) *ListMmsPartitionsRequestSorter {
	s.LastDdlTime = &v
	return s
}

func (s *ListMmsPartitionsRequestSorter) SetNumRows(v string) *ListMmsPartitionsRequestSorter {
	s.NumRows = &v
	return s
}

func (s *ListMmsPartitionsRequestSorter) SetSize(v string) *ListMmsPartitionsRequestSorter {
	s.Size = &v
	return s
}

type ListMmsPartitionsShrinkRequest struct {
	Sorter *ListMmsPartitionsShrinkRequestSorter `json:"sorter,omitempty" xml:"sorter,omitempty" type:"Struct"`
	// example:
	//
	// 2
	DbId *int64 `json:"dbId,omitempty" xml:"dbId,omitempty"`
	// example:
	//
	// d1
	DbName *string `json:"dbName,omitempty" xml:"dbName,omitempty"`
	// example:
	//
	// 2024-12-17 19:44:42
	LastDdlTimeEnd *string `json:"lastDdlTimeEnd,omitempty" xml:"lastDdlTimeEnd,omitempty"`
	// example:
	//
	// 2024-12-17 15:44:42
	LastDdlTimeStart *string `json:"lastDdlTimeStart,omitempty" xml:"lastDdlTimeStart,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	// example:
	//
	// 100
	PageSize     *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	StatusShrink *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// t1
	TableName *string `json:"tableName,omitempty" xml:"tableName,omitempty"`
	// example:
	//
	// false
	Updated *bool `json:"updated,omitempty" xml:"updated,omitempty"`
	// example:
	//
	// p1=1/p2=abc
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListMmsPartitionsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMmsPartitionsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListMmsPartitionsShrinkRequest) SetSorter(v *ListMmsPartitionsShrinkRequestSorter) *ListMmsPartitionsShrinkRequest {
	s.Sorter = v
	return s
}

func (s *ListMmsPartitionsShrinkRequest) SetDbId(v int64) *ListMmsPartitionsShrinkRequest {
	s.DbId = &v
	return s
}

func (s *ListMmsPartitionsShrinkRequest) SetDbName(v string) *ListMmsPartitionsShrinkRequest {
	s.DbName = &v
	return s
}

func (s *ListMmsPartitionsShrinkRequest) SetLastDdlTimeEnd(v string) *ListMmsPartitionsShrinkRequest {
	s.LastDdlTimeEnd = &v
	return s
}

func (s *ListMmsPartitionsShrinkRequest) SetLastDdlTimeStart(v string) *ListMmsPartitionsShrinkRequest {
	s.LastDdlTimeStart = &v
	return s
}

func (s *ListMmsPartitionsShrinkRequest) SetPageNum(v int32) *ListMmsPartitionsShrinkRequest {
	s.PageNum = &v
	return s
}

func (s *ListMmsPartitionsShrinkRequest) SetPageSize(v int32) *ListMmsPartitionsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListMmsPartitionsShrinkRequest) SetStatusShrink(v string) *ListMmsPartitionsShrinkRequest {
	s.StatusShrink = &v
	return s
}

func (s *ListMmsPartitionsShrinkRequest) SetTableName(v string) *ListMmsPartitionsShrinkRequest {
	s.TableName = &v
	return s
}

func (s *ListMmsPartitionsShrinkRequest) SetUpdated(v bool) *ListMmsPartitionsShrinkRequest {
	s.Updated = &v
	return s
}

func (s *ListMmsPartitionsShrinkRequest) SetValue(v string) *ListMmsPartitionsShrinkRequest {
	s.Value = &v
	return s
}

type ListMmsPartitionsShrinkRequestSorter struct {
	// example:
	//
	// desc
	LastDdlTime *string `json:"lastDdlTime,omitempty" xml:"lastDdlTime,omitempty"`
	// example:
	//
	// desc
	NumRows *string `json:"numRows,omitempty" xml:"numRows,omitempty"`
	// example:
	//
	// asc
	Size *string `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ListMmsPartitionsShrinkRequestSorter) String() string {
	return tea.Prettify(s)
}

func (s ListMmsPartitionsShrinkRequestSorter) GoString() string {
	return s.String()
}

func (s *ListMmsPartitionsShrinkRequestSorter) SetLastDdlTime(v string) *ListMmsPartitionsShrinkRequestSorter {
	s.LastDdlTime = &v
	return s
}

func (s *ListMmsPartitionsShrinkRequestSorter) SetNumRows(v string) *ListMmsPartitionsShrinkRequestSorter {
	s.NumRows = &v
	return s
}

func (s *ListMmsPartitionsShrinkRequestSorter) SetSize(v string) *ListMmsPartitionsShrinkRequestSorter {
	s.Size = &v
	return s
}

type ListMmsPartitionsResponseBody struct {
	Data *ListMmsPartitionsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// D9F872FD-5DDE-30A6-8C8A-1B8C6A81059F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListMmsPartitionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMmsPartitionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMmsPartitionsResponseBody) SetData(v *ListMmsPartitionsResponseBodyData) *ListMmsPartitionsResponseBody {
	s.Data = v
	return s
}

func (s *ListMmsPartitionsResponseBody) SetRequestId(v string) *ListMmsPartitionsResponseBody {
	s.RequestId = &v
	return s
}

type ListMmsPartitionsResponseBodyData struct {
	ObjectList []*ListMmsPartitionsResponseBodyDataObjectList `json:"objectList,omitempty" xml:"objectList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNum *int32 `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 1000
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListMmsPartitionsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListMmsPartitionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMmsPartitionsResponseBodyData) SetObjectList(v []*ListMmsPartitionsResponseBodyDataObjectList) *ListMmsPartitionsResponseBodyData {
	s.ObjectList = v
	return s
}

func (s *ListMmsPartitionsResponseBodyData) SetPageNum(v int32) *ListMmsPartitionsResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *ListMmsPartitionsResponseBodyData) SetPageSize(v int32) *ListMmsPartitionsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListMmsPartitionsResponseBodyData) SetTotal(v int32) *ListMmsPartitionsResponseBodyData {
	s.Total = &v
	return s
}

type ListMmsPartitionsResponseBodyDataObjectList struct {
	// example:
	//
	// 2
	DbId *int64 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// example:
	//
	// d1
	DbName *string `json:"dbName,omitempty" xml:"dbName,omitempty"`
	// example:
	//
	// 2323
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// lastDdlTime
	//
	// example:
	//
	// 2024-12-17 15:44:42
	LastDdlTime *string `json:"lastDdlTime,omitempty" xml:"lastDdlTime,omitempty"`
	// example:
	//
	// 2323
	NumRows *int64 `json:"numRows,omitempty" xml:"numRows,omitempty"`
	// example:
	//
	// 23223
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// example:
	//
	// 200018
	SourceId *int64 `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// example:
	//
	// demo
	SourceName *string `json:"sourceName,omitempty" xml:"sourceName,omitempty"`
	// example:
	//
	// DONE
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 23
	TableId *int64 `json:"tableId,omitempty" xml:"tableId,omitempty"`
	// example:
	//
	// t1
	TableName *string `json:"tableName,omitempty" xml:"tableName,omitempty"`
	// example:
	//
	// false
	Updated *bool `json:"updated,omitempty" xml:"updated,omitempty"`
	// example:
	//
	// p1=1/p2=abc
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListMmsPartitionsResponseBodyDataObjectList) String() string {
	return tea.Prettify(s)
}

func (s ListMmsPartitionsResponseBodyDataObjectList) GoString() string {
	return s.String()
}

func (s *ListMmsPartitionsResponseBodyDataObjectList) SetDbId(v int64) *ListMmsPartitionsResponseBodyDataObjectList {
	s.DbId = &v
	return s
}

func (s *ListMmsPartitionsResponseBodyDataObjectList) SetDbName(v string) *ListMmsPartitionsResponseBodyDataObjectList {
	s.DbName = &v
	return s
}

func (s *ListMmsPartitionsResponseBodyDataObjectList) SetId(v int64) *ListMmsPartitionsResponseBodyDataObjectList {
	s.Id = &v
	return s
}

func (s *ListMmsPartitionsResponseBodyDataObjectList) SetLastDdlTime(v string) *ListMmsPartitionsResponseBodyDataObjectList {
	s.LastDdlTime = &v
	return s
}

func (s *ListMmsPartitionsResponseBodyDataObjectList) SetNumRows(v int64) *ListMmsPartitionsResponseBodyDataObjectList {
	s.NumRows = &v
	return s
}

func (s *ListMmsPartitionsResponseBodyDataObjectList) SetSize(v int64) *ListMmsPartitionsResponseBodyDataObjectList {
	s.Size = &v
	return s
}

func (s *ListMmsPartitionsResponseBodyDataObjectList) SetSourceId(v int64) *ListMmsPartitionsResponseBodyDataObjectList {
	s.SourceId = &v
	return s
}

func (s *ListMmsPartitionsResponseBodyDataObjectList) SetSourceName(v string) *ListMmsPartitionsResponseBodyDataObjectList {
	s.SourceName = &v
	return s
}

func (s *ListMmsPartitionsResponseBodyDataObjectList) SetStatus(v string) *ListMmsPartitionsResponseBodyDataObjectList {
	s.Status = &v
	return s
}

func (s *ListMmsPartitionsResponseBodyDataObjectList) SetTableId(v int64) *ListMmsPartitionsResponseBodyDataObjectList {
	s.TableId = &v
	return s
}

func (s *ListMmsPartitionsResponseBodyDataObjectList) SetTableName(v string) *ListMmsPartitionsResponseBodyDataObjectList {
	s.TableName = &v
	return s
}

func (s *ListMmsPartitionsResponseBodyDataObjectList) SetUpdated(v bool) *ListMmsPartitionsResponseBodyDataObjectList {
	s.Updated = &v
	return s
}

func (s *ListMmsPartitionsResponseBodyDataObjectList) SetValue(v string) *ListMmsPartitionsResponseBodyDataObjectList {
	s.Value = &v
	return s
}

type ListMmsPartitionsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMmsPartitionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMmsPartitionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMmsPartitionsResponse) GoString() string {
	return s.String()
}

func (s *ListMmsPartitionsResponse) SetHeaders(v map[string]*string) *ListMmsPartitionsResponse {
	s.Headers = v
	return s
}

func (s *ListMmsPartitionsResponse) SetStatusCode(v int32) *ListMmsPartitionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMmsPartitionsResponse) SetBody(v *ListMmsPartitionsResponseBody) *ListMmsPartitionsResponse {
	s.Body = v
	return s
}

type ListMmsTablesRequest struct {
	Sorter *ListMmsTablesRequestSorter `json:"sorter,omitempty" xml:"sorter,omitempty" type:"Struct"`
	// example:
	//
	// 197
	DbId *int64 `json:"dbId,omitempty" xml:"dbId,omitempty"`
	// example:
	//
	// mms_test
	DbName *string `json:"dbName,omitempty" xml:"dbName,omitempty"`
	// example:
	//
	// true
	HasPartitions *bool `json:"hasPartitions,omitempty" xml:"hasPartitions,omitempty"`
	// example:
	//
	// 2024-12-19 15:44:42
	LastDdlTimeEnd *string `json:"lastDdlTimeEnd,omitempty" xml:"lastDdlTimeEnd,omitempty"`
	// example:
	//
	// 2024-12-17 15:44:42
	LastDdlTimeStart *string `json:"lastDdlTimeStart,omitempty" xml:"lastDdlTimeStart,omitempty"`
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// false
	OnlyName *bool `json:"onlyName,omitempty" xml:"onlyName,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32    `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Status   []*string `json:"status,omitempty" xml:"status,omitempty" type:"Repeated"`
	// example:
	//
	// MANAGED_TABLE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListMmsTablesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMmsTablesRequest) GoString() string {
	return s.String()
}

func (s *ListMmsTablesRequest) SetSorter(v *ListMmsTablesRequestSorter) *ListMmsTablesRequest {
	s.Sorter = v
	return s
}

func (s *ListMmsTablesRequest) SetDbId(v int64) *ListMmsTablesRequest {
	s.DbId = &v
	return s
}

func (s *ListMmsTablesRequest) SetDbName(v string) *ListMmsTablesRequest {
	s.DbName = &v
	return s
}

func (s *ListMmsTablesRequest) SetHasPartitions(v bool) *ListMmsTablesRequest {
	s.HasPartitions = &v
	return s
}

func (s *ListMmsTablesRequest) SetLastDdlTimeEnd(v string) *ListMmsTablesRequest {
	s.LastDdlTimeEnd = &v
	return s
}

func (s *ListMmsTablesRequest) SetLastDdlTimeStart(v string) *ListMmsTablesRequest {
	s.LastDdlTimeStart = &v
	return s
}

func (s *ListMmsTablesRequest) SetName(v string) *ListMmsTablesRequest {
	s.Name = &v
	return s
}

func (s *ListMmsTablesRequest) SetOnlyName(v bool) *ListMmsTablesRequest {
	s.OnlyName = &v
	return s
}

func (s *ListMmsTablesRequest) SetPageNum(v int32) *ListMmsTablesRequest {
	s.PageNum = &v
	return s
}

func (s *ListMmsTablesRequest) SetPageSize(v int32) *ListMmsTablesRequest {
	s.PageSize = &v
	return s
}

func (s *ListMmsTablesRequest) SetStatus(v []*string) *ListMmsTablesRequest {
	s.Status = v
	return s
}

func (s *ListMmsTablesRequest) SetType(v string) *ListMmsTablesRequest {
	s.Type = &v
	return s
}

type ListMmsTablesRequestSorter struct {
	// example:
	//
	// desc
	LastDdlTime *string `json:"lastDdlTime,omitempty" xml:"lastDdlTime,omitempty"`
	// example:
	//
	// desc
	NumRows *string `json:"numRows,omitempty" xml:"numRows,omitempty"`
	// example:
	//
	// asc
	Size *string `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ListMmsTablesRequestSorter) String() string {
	return tea.Prettify(s)
}

func (s ListMmsTablesRequestSorter) GoString() string {
	return s.String()
}

func (s *ListMmsTablesRequestSorter) SetLastDdlTime(v string) *ListMmsTablesRequestSorter {
	s.LastDdlTime = &v
	return s
}

func (s *ListMmsTablesRequestSorter) SetNumRows(v string) *ListMmsTablesRequestSorter {
	s.NumRows = &v
	return s
}

func (s *ListMmsTablesRequestSorter) SetSize(v string) *ListMmsTablesRequestSorter {
	s.Size = &v
	return s
}

type ListMmsTablesShrinkRequest struct {
	Sorter *ListMmsTablesShrinkRequestSorter `json:"sorter,omitempty" xml:"sorter,omitempty" type:"Struct"`
	// example:
	//
	// 197
	DbId *int64 `json:"dbId,omitempty" xml:"dbId,omitempty"`
	// example:
	//
	// mms_test
	DbName *string `json:"dbName,omitempty" xml:"dbName,omitempty"`
	// example:
	//
	// true
	HasPartitions *bool `json:"hasPartitions,omitempty" xml:"hasPartitions,omitempty"`
	// example:
	//
	// 2024-12-19 15:44:42
	LastDdlTimeEnd *string `json:"lastDdlTimeEnd,omitempty" xml:"lastDdlTimeEnd,omitempty"`
	// example:
	//
	// 2024-12-17 15:44:42
	LastDdlTimeStart *string `json:"lastDdlTimeStart,omitempty" xml:"lastDdlTimeStart,omitempty"`
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// false
	OnlyName *bool `json:"onlyName,omitempty" xml:"onlyName,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	// example:
	//
	// 10
	PageSize     *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	StatusShrink *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// MANAGED_TABLE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListMmsTablesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMmsTablesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListMmsTablesShrinkRequest) SetSorter(v *ListMmsTablesShrinkRequestSorter) *ListMmsTablesShrinkRequest {
	s.Sorter = v
	return s
}

func (s *ListMmsTablesShrinkRequest) SetDbId(v int64) *ListMmsTablesShrinkRequest {
	s.DbId = &v
	return s
}

func (s *ListMmsTablesShrinkRequest) SetDbName(v string) *ListMmsTablesShrinkRequest {
	s.DbName = &v
	return s
}

func (s *ListMmsTablesShrinkRequest) SetHasPartitions(v bool) *ListMmsTablesShrinkRequest {
	s.HasPartitions = &v
	return s
}

func (s *ListMmsTablesShrinkRequest) SetLastDdlTimeEnd(v string) *ListMmsTablesShrinkRequest {
	s.LastDdlTimeEnd = &v
	return s
}

func (s *ListMmsTablesShrinkRequest) SetLastDdlTimeStart(v string) *ListMmsTablesShrinkRequest {
	s.LastDdlTimeStart = &v
	return s
}

func (s *ListMmsTablesShrinkRequest) SetName(v string) *ListMmsTablesShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListMmsTablesShrinkRequest) SetOnlyName(v bool) *ListMmsTablesShrinkRequest {
	s.OnlyName = &v
	return s
}

func (s *ListMmsTablesShrinkRequest) SetPageNum(v int32) *ListMmsTablesShrinkRequest {
	s.PageNum = &v
	return s
}

func (s *ListMmsTablesShrinkRequest) SetPageSize(v int32) *ListMmsTablesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListMmsTablesShrinkRequest) SetStatusShrink(v string) *ListMmsTablesShrinkRequest {
	s.StatusShrink = &v
	return s
}

func (s *ListMmsTablesShrinkRequest) SetType(v string) *ListMmsTablesShrinkRequest {
	s.Type = &v
	return s
}

type ListMmsTablesShrinkRequestSorter struct {
	// example:
	//
	// desc
	LastDdlTime *string `json:"lastDdlTime,omitempty" xml:"lastDdlTime,omitempty"`
	// example:
	//
	// desc
	NumRows *string `json:"numRows,omitempty" xml:"numRows,omitempty"`
	// example:
	//
	// asc
	Size *string `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ListMmsTablesShrinkRequestSorter) String() string {
	return tea.Prettify(s)
}

func (s ListMmsTablesShrinkRequestSorter) GoString() string {
	return s.String()
}

func (s *ListMmsTablesShrinkRequestSorter) SetLastDdlTime(v string) *ListMmsTablesShrinkRequestSorter {
	s.LastDdlTime = &v
	return s
}

func (s *ListMmsTablesShrinkRequestSorter) SetNumRows(v string) *ListMmsTablesShrinkRequestSorter {
	s.NumRows = &v
	return s
}

func (s *ListMmsTablesShrinkRequestSorter) SetSize(v string) *ListMmsTablesShrinkRequestSorter {
	s.Size = &v
	return s
}

type ListMmsTablesResponseBody struct {
	Data *ListMmsTablesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// E7FB14F1-4ACD-5C73-A755-B302D70AB9AD
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListMmsTablesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMmsTablesResponseBody) GoString() string {
	return s.String()
}

func (s *ListMmsTablesResponseBody) SetData(v *ListMmsTablesResponseBodyData) *ListMmsTablesResponseBody {
	s.Data = v
	return s
}

func (s *ListMmsTablesResponseBody) SetRequestId(v string) *ListMmsTablesResponseBody {
	s.RequestId = &v
	return s
}

type ListMmsTablesResponseBodyData struct {
	ObjectList []*ListMmsTablesResponseBodyDataObjectList `json:"objectList,omitempty" xml:"objectList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNum *int32 `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 100
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListMmsTablesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListMmsTablesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMmsTablesResponseBodyData) SetObjectList(v []*ListMmsTablesResponseBodyDataObjectList) *ListMmsTablesResponseBodyData {
	s.ObjectList = v
	return s
}

func (s *ListMmsTablesResponseBodyData) SetPageNum(v int32) *ListMmsTablesResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *ListMmsTablesResponseBodyData) SetPageSize(v int32) *ListMmsTablesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListMmsTablesResponseBodyData) SetTotal(v int32) *ListMmsTablesResponseBodyData {
	s.Total = &v
	return s
}

type ListMmsTablesResponseBodyDataObjectList struct {
	// example:
	//
	// 196
	DbId *int64 `json:"dbId,omitempty" xml:"dbId,omitempty"`
	// example:
	//
	// demo
	DbName *string `json:"dbName,omitempty" xml:"dbName,omitempty"`
	// example:
	//
	// {"mapkey.delim":":","collection.delim":",","serialization.format":"|","field.delim":"|"}
	Extra *string `json:"extra,omitempty" xml:"extra,omitempty"`
	// example:
	//
	// true
	HasPartitions *bool `json:"hasPartitions,omitempty" xml:"hasPartitions,omitempty"`
	// table ID
	//
	// example:
	//
	// 1003476
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// inputFormat
	//
	// example:
	//
	// org.apache.hadoop.hive.ql.io.parquet.MapredParquetInputFormat
	InputFormat *string `json:"inputFormat,omitempty" xml:"inputFormat,omitempty"`
	// lastDdlTime
	//
	// example:
	//
	// 2024-12-17 15:44:42
	LastDdlTime *string `json:"lastDdlTime,omitempty" xml:"lastDdlTime,omitempty"`
	// example:
	//
	// | hdfs://master-1-1.c-c127cd184bb029ea.cn-zhangjiakou.emr.aliyuncs.com:9000/user/hive/warehouse/demo
	Location *string `json:"location,omitempty" xml:"location,omitempty"`
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 232323
	NumRows *int64 `json:"numRows,omitempty" xml:"numRows,omitempty"`
	// outFormat
	//
	// example:
	//
	// org.apache.hadoop.hive.ql.io.parquet.MapredParquetOutputFormat
	OutputFormat *string `json:"outputFormat,omitempty" xml:"outputFormat,omitempty"`
	// example:
	//
	// Hive
	Owner *string `json:"owner,omitempty" xml:"owner,omitempty"`
	// example:
	//
	// 100
	Partitions *int32 `json:"partitions,omitempty" xml:"partitions,omitempty"`
	// example:
	//
	// 20
	PartitionsDoing *int32 `json:"partitionsDoing,omitempty" xml:"partitionsDoing,omitempty"`
	// example:
	//
	// 60
	PartitionsDone *int32 `json:"partitionsDone,omitempty" xml:"partitionsDone,omitempty"`
	// example:
	//
	// 40
	PartitionsFailed *int32                                         `json:"partitionsFailed,omitempty" xml:"partitionsFailed,omitempty"`
	Schema           *ListMmsTablesResponseBodyDataObjectListSchema `json:"schema,omitempty" xml:"schema,omitempty" type:"Struct"`
	// serde
	//
	// example:
	//
	// org.apache.hadoop.hive.ql.io.parquet.serde.ParquetHiveSerDe
	Serde *string `json:"serde,omitempty" xml:"serde,omitempty"`
	// example:
	//
	// 2985028
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// example:
	//
	// 2000028
	SourceId *int64 `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// example:
	//
	// demo
	SourceName *string `json:"sourceName,omitempty" xml:"sourceName,omitempty"`
	// example:
	//
	// DONE
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// MANAGED_TABLE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// true
	Updated *bool `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s ListMmsTablesResponseBodyDataObjectList) String() string {
	return tea.Prettify(s)
}

func (s ListMmsTablesResponseBodyDataObjectList) GoString() string {
	return s.String()
}

func (s *ListMmsTablesResponseBodyDataObjectList) SetDbId(v int64) *ListMmsTablesResponseBodyDataObjectList {
	s.DbId = &v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectList) SetDbName(v string) *ListMmsTablesResponseBodyDataObjectList {
	s.DbName = &v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectList) SetExtra(v string) *ListMmsTablesResponseBodyDataObjectList {
	s.Extra = &v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectList) SetHasPartitions(v bool) *ListMmsTablesResponseBodyDataObjectList {
	s.HasPartitions = &v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectList) SetId(v int64) *ListMmsTablesResponseBodyDataObjectList {
	s.Id = &v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectList) SetInputFormat(v string) *ListMmsTablesResponseBodyDataObjectList {
	s.InputFormat = &v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectList) SetLastDdlTime(v string) *ListMmsTablesResponseBodyDataObjectList {
	s.LastDdlTime = &v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectList) SetLocation(v string) *ListMmsTablesResponseBodyDataObjectList {
	s.Location = &v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectList) SetName(v string) *ListMmsTablesResponseBodyDataObjectList {
	s.Name = &v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectList) SetNumRows(v int64) *ListMmsTablesResponseBodyDataObjectList {
	s.NumRows = &v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectList) SetOutputFormat(v string) *ListMmsTablesResponseBodyDataObjectList {
	s.OutputFormat = &v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectList) SetOwner(v string) *ListMmsTablesResponseBodyDataObjectList {
	s.Owner = &v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectList) SetPartitions(v int32) *ListMmsTablesResponseBodyDataObjectList {
	s.Partitions = &v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectList) SetPartitionsDoing(v int32) *ListMmsTablesResponseBodyDataObjectList {
	s.PartitionsDoing = &v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectList) SetPartitionsDone(v int32) *ListMmsTablesResponseBodyDataObjectList {
	s.PartitionsDone = &v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectList) SetPartitionsFailed(v int32) *ListMmsTablesResponseBodyDataObjectList {
	s.PartitionsFailed = &v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectList) SetSchema(v *ListMmsTablesResponseBodyDataObjectListSchema) *ListMmsTablesResponseBodyDataObjectList {
	s.Schema = v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectList) SetSerde(v string) *ListMmsTablesResponseBodyDataObjectList {
	s.Serde = &v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectList) SetSize(v int64) *ListMmsTablesResponseBodyDataObjectList {
	s.Size = &v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectList) SetSourceId(v int64) *ListMmsTablesResponseBodyDataObjectList {
	s.SourceId = &v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectList) SetSourceName(v string) *ListMmsTablesResponseBodyDataObjectList {
	s.SourceName = &v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectList) SetStatus(v string) *ListMmsTablesResponseBodyDataObjectList {
	s.Status = &v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectList) SetType(v string) *ListMmsTablesResponseBodyDataObjectList {
	s.Type = &v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectList) SetUpdated(v bool) *ListMmsTablesResponseBodyDataObjectList {
	s.Updated = &v
	return s
}

type ListMmsTablesResponseBodyDataObjectListSchema struct {
	Columns []*ListMmsTablesResponseBodyDataObjectListSchemaColumns `json:"columns,omitempty" xml:"columns,omitempty" type:"Repeated"`
	// example:
	//
	// for mms test
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	// example:
	//
	// test
	Name       *string                                                    `json:"name,omitempty" xml:"name,omitempty"`
	Partitions []*ListMmsTablesResponseBodyDataObjectListSchemaPartitions `json:"partitions,omitempty" xml:"partitions,omitempty" type:"Repeated"`
}

func (s ListMmsTablesResponseBodyDataObjectListSchema) String() string {
	return tea.Prettify(s)
}

func (s ListMmsTablesResponseBodyDataObjectListSchema) GoString() string {
	return s.String()
}

func (s *ListMmsTablesResponseBodyDataObjectListSchema) SetColumns(v []*ListMmsTablesResponseBodyDataObjectListSchemaColumns) *ListMmsTablesResponseBodyDataObjectListSchema {
	s.Columns = v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectListSchema) SetComment(v string) *ListMmsTablesResponseBodyDataObjectListSchema {
	s.Comment = &v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectListSchema) SetName(v string) *ListMmsTablesResponseBodyDataObjectListSchema {
	s.Name = &v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectListSchema) SetPartitions(v []*ListMmsTablesResponseBodyDataObjectListSchemaPartitions) *ListMmsTablesResponseBodyDataObjectListSchema {
	s.Partitions = v
	return s
}

type ListMmsTablesResponseBodyDataObjectListSchemaColumns struct {
	// example:
	//
	// user id
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	// example:
	//
	// ""
	DefaultValue *string `json:"defaultValue,omitempty" xml:"defaultValue,omitempty"`
	// example:
	//
	// user_id
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// false
	Nullable *bool `json:"nullable,omitempty" xml:"nullable,omitempty"`
	// example:
	//
	// bigint
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListMmsTablesResponseBodyDataObjectListSchemaColumns) String() string {
	return tea.Prettify(s)
}

func (s ListMmsTablesResponseBodyDataObjectListSchemaColumns) GoString() string {
	return s.String()
}

func (s *ListMmsTablesResponseBodyDataObjectListSchemaColumns) SetComment(v string) *ListMmsTablesResponseBodyDataObjectListSchemaColumns {
	s.Comment = &v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectListSchemaColumns) SetDefaultValue(v string) *ListMmsTablesResponseBodyDataObjectListSchemaColumns {
	s.DefaultValue = &v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectListSchemaColumns) SetName(v string) *ListMmsTablesResponseBodyDataObjectListSchemaColumns {
	s.Name = &v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectListSchemaColumns) SetNullable(v bool) *ListMmsTablesResponseBodyDataObjectListSchemaColumns {
	s.Nullable = &v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectListSchemaColumns) SetType(v string) *ListMmsTablesResponseBodyDataObjectListSchemaColumns {
	s.Type = &v
	return s
}

type ListMmsTablesResponseBodyDataObjectListSchemaPartitions struct {
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	// example:
	//
	// abc
	DefaultValue *string `json:"defaultValue,omitempty" xml:"defaultValue,omitempty"`
	// example:
	//
	// p1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// false
	Nullable *bool `json:"nullable,omitempty" xml:"nullable,omitempty"`
	// example:
	//
	// string
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListMmsTablesResponseBodyDataObjectListSchemaPartitions) String() string {
	return tea.Prettify(s)
}

func (s ListMmsTablesResponseBodyDataObjectListSchemaPartitions) GoString() string {
	return s.String()
}

func (s *ListMmsTablesResponseBodyDataObjectListSchemaPartitions) SetComment(v string) *ListMmsTablesResponseBodyDataObjectListSchemaPartitions {
	s.Comment = &v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectListSchemaPartitions) SetDefaultValue(v string) *ListMmsTablesResponseBodyDataObjectListSchemaPartitions {
	s.DefaultValue = &v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectListSchemaPartitions) SetName(v string) *ListMmsTablesResponseBodyDataObjectListSchemaPartitions {
	s.Name = &v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectListSchemaPartitions) SetNullable(v bool) *ListMmsTablesResponseBodyDataObjectListSchemaPartitions {
	s.Nullable = &v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectListSchemaPartitions) SetType(v string) *ListMmsTablesResponseBodyDataObjectListSchemaPartitions {
	s.Type = &v
	return s
}

type ListMmsTablesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMmsTablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMmsTablesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMmsTablesResponse) GoString() string {
	return s.String()
}

func (s *ListMmsTablesResponse) SetHeaders(v map[string]*string) *ListMmsTablesResponse {
	s.Headers = v
	return s
}

func (s *ListMmsTablesResponse) SetStatusCode(v int32) *ListMmsTablesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMmsTablesResponse) SetBody(v *ListMmsTablesResponseBody) *ListMmsTablesResponse {
	s.Body = v
	return s
}

type ListMmsTaskLogsResponseBody struct {
	Data []*ListMmsTaskLogsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// A3AE5649-EF90-54BD-86D0-C632FA950988
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListMmsTaskLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMmsTaskLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMmsTaskLogsResponseBody) SetData(v []*ListMmsTaskLogsResponseBodyData) *ListMmsTaskLogsResponseBody {
	s.Data = v
	return s
}

func (s *ListMmsTaskLogsResponseBody) SetRequestId(v string) *ListMmsTaskLogsResponseBody {
	s.RequestId = &v
	return s
}

type ListMmsTaskLogsResponseBodyData struct {
	// example:
	//
	// create schema if not exists mms_test.default;
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// example:
	//
	// 2024-12-17 15:44:00
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 10000
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// ok
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty"`
	// example:
	//
	// 2000015
	SourceId *int64 `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// example:
	//
	// DATA_DOING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 4023
	TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s ListMmsTaskLogsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListMmsTaskLogsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMmsTaskLogsResponseBodyData) SetAction(v string) *ListMmsTaskLogsResponseBodyData {
	s.Action = &v
	return s
}

func (s *ListMmsTaskLogsResponseBodyData) SetCreateTime(v string) *ListMmsTaskLogsResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListMmsTaskLogsResponseBodyData) SetId(v int64) *ListMmsTaskLogsResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListMmsTaskLogsResponseBodyData) SetMsg(v string) *ListMmsTaskLogsResponseBodyData {
	s.Msg = &v
	return s
}

func (s *ListMmsTaskLogsResponseBodyData) SetSourceId(v int64) *ListMmsTaskLogsResponseBodyData {
	s.SourceId = &v
	return s
}

func (s *ListMmsTaskLogsResponseBodyData) SetStatus(v string) *ListMmsTaskLogsResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListMmsTaskLogsResponseBodyData) SetTaskId(v int64) *ListMmsTaskLogsResponseBodyData {
	s.TaskId = &v
	return s
}

type ListMmsTaskLogsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMmsTaskLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMmsTaskLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMmsTaskLogsResponse) GoString() string {
	return s.String()
}

func (s *ListMmsTaskLogsResponse) SetHeaders(v map[string]*string) *ListMmsTaskLogsResponse {
	s.Headers = v
	return s
}

func (s *ListMmsTaskLogsResponse) SetStatusCode(v int32) *ListMmsTaskLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMmsTaskLogsResponse) SetBody(v *ListMmsTaskLogsResponseBody) *ListMmsTaskLogsResponse {
	s.Body = v
	return s
}

type ListMmsTasksRequest struct {
	Sorter *ListMmsTasksRequestSorter `json:"sorter,omitempty" xml:"sorter,omitempty" type:"Struct"`
	// example:
	//
	// mms_test
	DstDbName *string `json:"dstDbName,omitempty" xml:"dstDbName,omitempty"`
	// example:
	//
	// test_table_1
	DstTableName *string `json:"dstTableName,omitempty" xml:"dstTableName,omitempty"`
	// example:
	//
	// 10
	JobId *int64 `json:"jobId,omitempty" xml:"jobId,omitempty"`
	// example:
	//
	// test1
	JobName *string `json:"jobName,omitempty" xml:"jobName,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// p1=1/p2=abc
	Partition *string `json:"partition,omitempty" xml:"partition,omitempty"`
	// example:
	//
	// test_db_1
	SrcDbName *string `json:"srcDbName,omitempty" xml:"srcDbName,omitempty"`
	// example:
	//
	// test_table_1
	SrcTableName *string `json:"srcTableName,omitempty" xml:"srcTableName,omitempty"`
	// example:
	//
	// DATA_DOING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListMmsTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMmsTasksRequest) GoString() string {
	return s.String()
}

func (s *ListMmsTasksRequest) SetSorter(v *ListMmsTasksRequestSorter) *ListMmsTasksRequest {
	s.Sorter = v
	return s
}

func (s *ListMmsTasksRequest) SetDstDbName(v string) *ListMmsTasksRequest {
	s.DstDbName = &v
	return s
}

func (s *ListMmsTasksRequest) SetDstTableName(v string) *ListMmsTasksRequest {
	s.DstTableName = &v
	return s
}

func (s *ListMmsTasksRequest) SetJobId(v int64) *ListMmsTasksRequest {
	s.JobId = &v
	return s
}

func (s *ListMmsTasksRequest) SetJobName(v string) *ListMmsTasksRequest {
	s.JobName = &v
	return s
}

func (s *ListMmsTasksRequest) SetPageNum(v int32) *ListMmsTasksRequest {
	s.PageNum = &v
	return s
}

func (s *ListMmsTasksRequest) SetPageSize(v int32) *ListMmsTasksRequest {
	s.PageSize = &v
	return s
}

func (s *ListMmsTasksRequest) SetPartition(v string) *ListMmsTasksRequest {
	s.Partition = &v
	return s
}

func (s *ListMmsTasksRequest) SetSrcDbName(v string) *ListMmsTasksRequest {
	s.SrcDbName = &v
	return s
}

func (s *ListMmsTasksRequest) SetSrcTableName(v string) *ListMmsTasksRequest {
	s.SrcTableName = &v
	return s
}

func (s *ListMmsTasksRequest) SetStatus(v string) *ListMmsTasksRequest {
	s.Status = &v
	return s
}

type ListMmsTasksRequestSorter struct {
	// example:
	//
	// desc
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// asc
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListMmsTasksRequestSorter) String() string {
	return tea.Prettify(s)
}

func (s ListMmsTasksRequestSorter) GoString() string {
	return s.String()
}

func (s *ListMmsTasksRequestSorter) SetStartTime(v string) *ListMmsTasksRequestSorter {
	s.StartTime = &v
	return s
}

func (s *ListMmsTasksRequestSorter) SetStatus(v string) *ListMmsTasksRequestSorter {
	s.Status = &v
	return s
}

type ListMmsTasksResponseBody struct {
	Data *ListMmsTasksResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 373A5CB2-8570-53BE-A98F-729B11D7A8B0
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListMmsTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMmsTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListMmsTasksResponseBody) SetData(v *ListMmsTasksResponseBodyData) *ListMmsTasksResponseBody {
	s.Data = v
	return s
}

func (s *ListMmsTasksResponseBody) SetRequestId(v string) *ListMmsTasksResponseBody {
	s.RequestId = &v
	return s
}

type ListMmsTasksResponseBodyData struct {
	ObjectList []*ListMmsTasksResponseBodyDataObjectList `json:"objectList,omitempty" xml:"objectList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNum *int32 `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 100
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListMmsTasksResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListMmsTasksResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMmsTasksResponseBodyData) SetObjectList(v []*ListMmsTasksResponseBodyDataObjectList) *ListMmsTasksResponseBodyData {
	s.ObjectList = v
	return s
}

func (s *ListMmsTasksResponseBodyData) SetPageNum(v int32) *ListMmsTasksResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *ListMmsTasksResponseBodyData) SetPageSize(v int32) *ListMmsTasksResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListMmsTasksResponseBodyData) SetTotal(v int32) *ListMmsTasksResponseBodyData {
	s.Total = &v
	return s
}

type ListMmsTasksResponseBodyDataObjectList struct {
	// example:
	//
	// 2024-10-25 04:21:01
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 196
	DbId *int64 `json:"dbId,omitempty" xml:"dbId,omitempty"`
	// example:
	//
	// mms_test
	DstDbName *string `json:"dstDbName,omitempty" xml:"dstDbName,omitempty"`
	// example:
	//
	// default
	DstSchemaName *string `json:"dstSchemaName,omitempty" xml:"dstSchemaName,omitempty"`
	// example:
	//
	// table_1
	DstTableName *string `json:"dstTableName,omitempty" xml:"dstTableName,omitempty"`
	// example:
	//
	// 2024-10-25 07:21:01
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 2323
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 87
	JobId *int64 `json:"jobId,omitempty" xml:"jobId,omitempty"`
	// example:
	//
	// test_odps_spark
	JobName *string `json:"jobName,omitempty" xml:"jobName,omitempty"`
	// example:
	//
	// 1
	RetriedTimes *int32 `json:"retriedTimes,omitempty" xml:"retriedTimes,omitempty"`
	// example:
	//
	// true
	Running *bool `json:"running,omitempty" xml:"running,omitempty"`
	// example:
	//
	// 2000028
	SourceId *int64 `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// example:
	//
	// demo
	SourceName *string `json:"sourceName,omitempty" xml:"sourceName,omitempty"`
	// example:
	//
	// db_1
	SrcDbName *string `json:"srcDbName,omitempty" xml:"srcDbName,omitempty"`
	// example:
	//
	// default
	SrcSchemaName *string `json:"srcSchemaName,omitempty" xml:"srcSchemaName,omitempty"`
	// example:
	//
	// table_1
	SrcTableName *string `json:"srcTableName,omitempty" xml:"srcTableName,omitempty"`
	// example:
	//
	// 2024-10-25 06:21:01
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// DATA_DOING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// false
	Stopped *bool `json:"stopped,omitempty" xml:"stopped,omitempty"`
	// example:
	//
	// 23
	TableId *int64 `json:"tableId,omitempty" xml:"tableId,omitempty"`
	// example:
	//
	// BIGQUERY
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListMmsTasksResponseBodyDataObjectList) String() string {
	return tea.Prettify(s)
}

func (s ListMmsTasksResponseBodyDataObjectList) GoString() string {
	return s.String()
}

func (s *ListMmsTasksResponseBodyDataObjectList) SetCreateTime(v string) *ListMmsTasksResponseBodyDataObjectList {
	s.CreateTime = &v
	return s
}

func (s *ListMmsTasksResponseBodyDataObjectList) SetDbId(v int64) *ListMmsTasksResponseBodyDataObjectList {
	s.DbId = &v
	return s
}

func (s *ListMmsTasksResponseBodyDataObjectList) SetDstDbName(v string) *ListMmsTasksResponseBodyDataObjectList {
	s.DstDbName = &v
	return s
}

func (s *ListMmsTasksResponseBodyDataObjectList) SetDstSchemaName(v string) *ListMmsTasksResponseBodyDataObjectList {
	s.DstSchemaName = &v
	return s
}

func (s *ListMmsTasksResponseBodyDataObjectList) SetDstTableName(v string) *ListMmsTasksResponseBodyDataObjectList {
	s.DstTableName = &v
	return s
}

func (s *ListMmsTasksResponseBodyDataObjectList) SetEndTime(v string) *ListMmsTasksResponseBodyDataObjectList {
	s.EndTime = &v
	return s
}

func (s *ListMmsTasksResponseBodyDataObjectList) SetId(v int64) *ListMmsTasksResponseBodyDataObjectList {
	s.Id = &v
	return s
}

func (s *ListMmsTasksResponseBodyDataObjectList) SetJobId(v int64) *ListMmsTasksResponseBodyDataObjectList {
	s.JobId = &v
	return s
}

func (s *ListMmsTasksResponseBodyDataObjectList) SetJobName(v string) *ListMmsTasksResponseBodyDataObjectList {
	s.JobName = &v
	return s
}

func (s *ListMmsTasksResponseBodyDataObjectList) SetRetriedTimes(v int32) *ListMmsTasksResponseBodyDataObjectList {
	s.RetriedTimes = &v
	return s
}

func (s *ListMmsTasksResponseBodyDataObjectList) SetRunning(v bool) *ListMmsTasksResponseBodyDataObjectList {
	s.Running = &v
	return s
}

func (s *ListMmsTasksResponseBodyDataObjectList) SetSourceId(v int64) *ListMmsTasksResponseBodyDataObjectList {
	s.SourceId = &v
	return s
}

func (s *ListMmsTasksResponseBodyDataObjectList) SetSourceName(v string) *ListMmsTasksResponseBodyDataObjectList {
	s.SourceName = &v
	return s
}

func (s *ListMmsTasksResponseBodyDataObjectList) SetSrcDbName(v string) *ListMmsTasksResponseBodyDataObjectList {
	s.SrcDbName = &v
	return s
}

func (s *ListMmsTasksResponseBodyDataObjectList) SetSrcSchemaName(v string) *ListMmsTasksResponseBodyDataObjectList {
	s.SrcSchemaName = &v
	return s
}

func (s *ListMmsTasksResponseBodyDataObjectList) SetSrcTableName(v string) *ListMmsTasksResponseBodyDataObjectList {
	s.SrcTableName = &v
	return s
}

func (s *ListMmsTasksResponseBodyDataObjectList) SetStartTime(v string) *ListMmsTasksResponseBodyDataObjectList {
	s.StartTime = &v
	return s
}

func (s *ListMmsTasksResponseBodyDataObjectList) SetStatus(v string) *ListMmsTasksResponseBodyDataObjectList {
	s.Status = &v
	return s
}

func (s *ListMmsTasksResponseBodyDataObjectList) SetStopped(v bool) *ListMmsTasksResponseBodyDataObjectList {
	s.Stopped = &v
	return s
}

func (s *ListMmsTasksResponseBodyDataObjectList) SetTableId(v int64) *ListMmsTasksResponseBodyDataObjectList {
	s.TableId = &v
	return s
}

func (s *ListMmsTasksResponseBodyDataObjectList) SetType(v string) *ListMmsTasksResponseBodyDataObjectList {
	s.Type = &v
	return s
}

type ListMmsTasksResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMmsTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMmsTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMmsTasksResponse) GoString() string {
	return s.String()
}

func (s *ListMmsTasksResponse) SetHeaders(v map[string]*string) *ListMmsTasksResponse {
	s.Headers = v
	return s
}

func (s *ListMmsTasksResponse) SetStatusCode(v int32) *ListMmsTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMmsTasksResponse) SetBody(v *ListMmsTasksResponseBody) *ListMmsTasksResponse {
	s.Body = v
	return s
}

type ListPackagesResponseBody struct {
	// The returned data.
	Data *ListPackagesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 0bc3b4aa16677927210252786e4cb6
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListPackagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPackagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPackagesResponseBody) SetData(v *ListPackagesResponseBodyData) *ListPackagesResponseBody {
	s.Data = v
	return s
}

func (s *ListPackagesResponseBody) SetRequestId(v string) *ListPackagesResponseBody {
	s.RequestId = &v
	return s
}

type ListPackagesResponseBodyData struct {
	// The packages that were created.
	CreatedPackages []*ListPackagesResponseBodyDataCreatedPackages `json:"createdPackages,omitempty" xml:"createdPackages,omitempty" type:"Repeated"`
	// The packages that were installed.
	InstalledPackages []*ListPackagesResponseBodyDataInstalledPackages `json:"installedPackages,omitempty" xml:"installedPackages,omitempty" type:"Repeated"`
}

func (s ListPackagesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListPackagesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPackagesResponseBodyData) SetCreatedPackages(v []*ListPackagesResponseBodyDataCreatedPackages) *ListPackagesResponseBodyData {
	s.CreatedPackages = v
	return s
}

func (s *ListPackagesResponseBodyData) SetInstalledPackages(v []*ListPackagesResponseBodyDataInstalledPackages) *ListPackagesResponseBodyData {
	s.InstalledPackages = v
	return s
}

type ListPackagesResponseBodyDataCreatedPackages struct {
	// The time when the package was created.
	//
	// example:
	//
	// 2022-08-02T02:30:34Z
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The name of the package.
	//
	// example:
	//
	// packageA
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListPackagesResponseBodyDataCreatedPackages) String() string {
	return tea.Prettify(s)
}

func (s ListPackagesResponseBodyDataCreatedPackages) GoString() string {
	return s.String()
}

func (s *ListPackagesResponseBodyDataCreatedPackages) SetCreateTime(v int64) *ListPackagesResponseBodyDataCreatedPackages {
	s.CreateTime = &v
	return s
}

func (s *ListPackagesResponseBodyDataCreatedPackages) SetName(v string) *ListPackagesResponseBodyDataCreatedPackages {
	s.Name = &v
	return s
}

type ListPackagesResponseBodyDataInstalledPackages struct {
	// The time when the package was installed.
	//
	// example:
	//
	// 2022-09-02T02:30:34Z
	InstallTime *int64 `json:"installTime,omitempty" xml:"installTime,omitempty"`
	// The name of the package.
	//
	// example:
	//
	// packageB
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The project to which the package belongs. This parameter is required if the package is installed in the MaxCompute project.
	//
	// example:
	//
	// projectB
	SourceProject *string `json:"sourceProject,omitempty" xml:"sourceProject,omitempty"`
	// The status of the package.
	//
	// example:
	//
	// OK
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListPackagesResponseBodyDataInstalledPackages) String() string {
	return tea.Prettify(s)
}

func (s ListPackagesResponseBodyDataInstalledPackages) GoString() string {
	return s.String()
}

func (s *ListPackagesResponseBodyDataInstalledPackages) SetInstallTime(v int64) *ListPackagesResponseBodyDataInstalledPackages {
	s.InstallTime = &v
	return s
}

func (s *ListPackagesResponseBodyDataInstalledPackages) SetName(v string) *ListPackagesResponseBodyDataInstalledPackages {
	s.Name = &v
	return s
}

func (s *ListPackagesResponseBodyDataInstalledPackages) SetSourceProject(v string) *ListPackagesResponseBodyDataInstalledPackages {
	s.SourceProject = &v
	return s
}

func (s *ListPackagesResponseBodyDataInstalledPackages) SetStatus(v string) *ListPackagesResponseBodyDataInstalledPackages {
	s.Status = &v
	return s
}

type ListPackagesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPackagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPackagesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPackagesResponse) GoString() string {
	return s.String()
}

func (s *ListPackagesResponse) SetHeaders(v map[string]*string) *ListPackagesResponse {
	s.Headers = v
	return s
}

func (s *ListPackagesResponse) SetStatusCode(v int32) *ListPackagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPackagesResponse) SetBody(v *ListPackagesResponseBody) *ListPackagesResponse {
	s.Body = v
	return s
}

type ListProjectUsersResponseBody struct {
	// The returned data.
	Data *ListProjectUsersResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0b87b7b316643495896551555e855b
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListProjectUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProjectUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectUsersResponseBody) SetData(v *ListProjectUsersResponseBodyData) *ListProjectUsersResponseBody {
	s.Data = v
	return s
}

func (s *ListProjectUsersResponseBody) SetRequestId(v string) *ListProjectUsersResponseBody {
	s.RequestId = &v
	return s
}

type ListProjectUsersResponseBodyData struct {
	// An array that contains users.
	Users []*ListProjectUsersResponseBodyDataUsers `json:"users,omitempty" xml:"users,omitempty" type:"Repeated"`
}

func (s ListProjectUsersResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListProjectUsersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListProjectUsersResponseBodyData) SetUsers(v []*ListProjectUsersResponseBodyDataUsers) *ListProjectUsersResponseBodyData {
	s.Users = v
	return s
}

type ListProjectUsersResponseBodyDataUsers struct {
	// The name of the user.
	//
	// example:
	//
	// userA
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListProjectUsersResponseBodyDataUsers) String() string {
	return tea.Prettify(s)
}

func (s ListProjectUsersResponseBodyDataUsers) GoString() string {
	return s.String()
}

func (s *ListProjectUsersResponseBodyDataUsers) SetName(v string) *ListProjectUsersResponseBodyDataUsers {
	s.Name = &v
	return s
}

type ListProjectUsersResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProjectUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProjectUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProjectUsersResponse) GoString() string {
	return s.String()
}

func (s *ListProjectUsersResponse) SetHeaders(v map[string]*string) *ListProjectUsersResponse {
	s.Headers = v
	return s
}

func (s *ListProjectUsersResponse) SetStatusCode(v int32) *ListProjectUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProjectUsersResponse) SetBody(v *ListProjectUsersResponseBody) *ListProjectUsersResponse {
	s.Body = v
	return s
}

type ListProjectsRequest struct {
	// Specifies whether to list the built-in **SYSTEM_CATALOG*	- projects that are used to provide data such as project metadata and historical usage data. For more information, see [Tenant-level Information Schema](https://www.alibabacloud.com/help/zh/maxcompute/user-guide/tenant-level-information-schema).
	//
	// Valid values:
	//
	// 	- true: The built-in SYSTEM_CATALOG projects are listed.
	//
	// 	- false: The built-in SYSTEM_CATALOG projects are not listed.
	//
	// example:
	//
	// true
	ListSystemCatalog *bool `json:"listSystemCatalog,omitempty" xml:"listSystemCatalog,omitempty"`
	// The maximum number of entries to return on each page.
	//
	// example:
	//
	// cHlvZHBzX3VkZl8xMDExNV8xNDU3NDI4NDkzKg==
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
	// The maximum number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	MaxItem *int32 `json:"maxItem,omitempty" xml:"maxItem,omitempty"`
	// Specifies the marker after which the returned list begins.
	//
	// example:
	//
	// a
	Prefix *string `json:"prefix,omitempty" xml:"prefix,omitempty"`
	// The quota name that is automatically generated. You can log on to the [MaxCompute console](https://maxcompute.console.aliyun.com), choose **Workspace*	- > **Quotas*	- from the left-side navigation pane, and then view the quota name on the **Quotas*	- page.
	//
	// example:
	//
	// "hsajkdgbkaubh"
	QuotaName *string `json:"quotaName,omitempty" xml:"quotaName,omitempty"`
	// The quota nickname. You can log on to the [MaxCompute console](https://maxcompute.console.aliyun.com), choose **Workspace*	- > **Quotas*	- from the left-side navigation pane, and then view the quota nickname on the **Quotas*	- page.
	//
	// example:
	//
	// quotaA
	QuotaNickName *string `json:"quotaNickName,omitempty" xml:"quotaNickName,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The instance ID and billing method of the default computing quota.
	//
	// example:
	//
	// "aaaa-bbbb"
	SaleTags *string `json:"saleTags,omitempty" xml:"saleTags,omitempty"`
	// The tenant ID. You can log on to the [MaxCompute console](https://maxcompute.console.aliyun.com), and choose **Tenants*	- > **Tenant Property*	- from the left-side navigation pane to view the tenant ID.
	//
	// example:
	//
	// 549532154333697
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The project type. Valid values:
	//
	// 	- **managed**: internal project
	//
	// 	- **external**: external project
	//
	// example:
	//
	// "managed"
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListProjectsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsRequest) GoString() string {
	return s.String()
}

func (s *ListProjectsRequest) SetListSystemCatalog(v bool) *ListProjectsRequest {
	s.ListSystemCatalog = &v
	return s
}

func (s *ListProjectsRequest) SetMarker(v string) *ListProjectsRequest {
	s.Marker = &v
	return s
}

func (s *ListProjectsRequest) SetMaxItem(v int32) *ListProjectsRequest {
	s.MaxItem = &v
	return s
}

func (s *ListProjectsRequest) SetPrefix(v string) *ListProjectsRequest {
	s.Prefix = &v
	return s
}

func (s *ListProjectsRequest) SetQuotaName(v string) *ListProjectsRequest {
	s.QuotaName = &v
	return s
}

func (s *ListProjectsRequest) SetQuotaNickName(v string) *ListProjectsRequest {
	s.QuotaNickName = &v
	return s
}

func (s *ListProjectsRequest) SetRegion(v string) *ListProjectsRequest {
	s.Region = &v
	return s
}

func (s *ListProjectsRequest) SetSaleTags(v string) *ListProjectsRequest {
	s.SaleTags = &v
	return s
}

func (s *ListProjectsRequest) SetTenantId(v string) *ListProjectsRequest {
	s.TenantId = &v
	return s
}

func (s *ListProjectsRequest) SetType(v string) *ListProjectsRequest {
	s.Type = &v
	return s
}

type ListProjectsResponseBody struct {
	// The data returned.
	Data *ListProjectsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0b16399216671970335563173e2340
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListProjectsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBody) SetData(v *ListProjectsResponseBodyData) *ListProjectsResponseBody {
	s.Data = v
	return s
}

func (s *ListProjectsResponseBody) SetRequestId(v string) *ListProjectsResponseBody {
	s.RequestId = &v
	return s
}

type ListProjectsResponseBodyData struct {
	// A pagination token. Only continuous page turning is supported. If NextToken is not empty, the next page exists. The value of NextToken can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6kvikyUl3ChyRxN+qLPvtOb
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Indicates the marker after which the returned list begins.
	//
	// example:
	//
	// cHlvZHBzX3VkZl8xMDExNV8xNDU3NDI4NDkzKg==
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
	// The maximum number of entries returned per page.
	//
	// example:
	//
	// 10
	MaxItem *int32 `json:"maxItem,omitempty" xml:"maxItem,omitempty"`
	// The list of projects.
	Projects []*ListProjectsResponseBodyDataProjects `json:"projects,omitempty" xml:"projects,omitempty" type:"Repeated"`
}

func (s ListProjectsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyData) SetNextToken(v string) *ListProjectsResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListProjectsResponseBodyData) SetMarker(v string) *ListProjectsResponseBodyData {
	s.Marker = &v
	return s
}

func (s *ListProjectsResponseBodyData) SetMaxItem(v int32) *ListProjectsResponseBodyData {
	s.MaxItem = &v
	return s
}

func (s *ListProjectsResponseBodyData) SetProjects(v []*ListProjectsResponseBodyDataProjects) *ListProjectsResponseBodyData {
	s.Projects = v
	return s
}

type ListProjectsResponseBodyDataProjects struct {
	// The project description.
	//
	// example:
	//
	// maxcompute projects
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	// The total storage usage. The storage space that is occupied by your project, which is the logical storage space after your project data is collected and compressed.
	//
	// example:
	//
	// 16489027
	CostStorage *string `json:"costStorage,omitempty" xml:"costStorage,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 1704380838000
	CreatedTime *int64 `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// The default computing quota that is used to allocate computing resources. If you do not specify a computing quota for your project, the jobs that are initiated by your project consume the computing resources in the default quota. For more information about how to use computing resources, see [Use quota groups for computing resources](https://www.alibabacloud.com/help/zh/maxcompute/user-guide/use-of-computing-resources)
	//
	// example:
	//
	// quotaA
	DefaultQuota *string `json:"defaultQuota,omitempty" xml:"defaultQuota,omitempty"`
	// The information about the IP address whitelist.
	IpWhiteList *ListProjectsResponseBodyDataProjectsIpWhiteList `json:"ipWhiteList,omitempty" xml:"ipWhiteList,omitempty" type:"Struct"`
	// The name of the project.
	//
	// example:
	//
	// odps_project
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The account information of the project owner.
	//
	// example:
	//
	// 1139815775606813
	Owner *string `json:"owner,omitempty" xml:"owner,omitempty"`
	// The basic properties of the project.
	Properties *ListProjectsResponseBodyDataProjectsProperties `json:"properties,omitempty" xml:"properties,omitempty" type:"Struct"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The instance ID and billing method of the default computing quota.
	SaleTag *ListProjectsResponseBodyDataProjectsSaleTag `json:"saleTag,omitempty" xml:"saleTag,omitempty" type:"Struct"`
	// The permission properties.
	SecurityProperties *ListProjectsResponseBodyDataProjectsSecurityProperties `json:"securityProperties,omitempty" xml:"securityProperties,omitempty" type:"Struct"`
	// The project status. Valid values:
	//
	// 	- **AVAILABLE**
	//
	// 	- **READONLY**
	//
	// 	- **FROZEN**
	//
	// 	- **DELETING**
	//
	// example:
	//
	// AVAILABLE
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// Indicates whether data storage by schema is supported. MaxCompute supports the schema feature. This feature allows you to classify objects such as tables, resources, and user-defined functions (UDFs) in a project by schema. You can create multiple schemas in a project. For more information, see [Schema-related operations](https://www.alibabacloud.com/help/zh/maxcompute/user-guide/schema-related-operations).
	//
	// Valid values:
	//
	// 	- true: supported
	//
	// 	- false: not supported
	//
	// example:
	//
	// true
	ThreeTierModel *bool `json:"threeTierModel,omitempty" xml:"threeTierModel,omitempty"`
	// The project type. Valid values:
	//
	// 	- **managed**: internal project
	//
	// 	- **external**: external project
	//
	// example:
	//
	// managed
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListProjectsResponseBodyDataProjects) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponseBodyDataProjects) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyDataProjects) SetComment(v string) *ListProjectsResponseBodyDataProjects {
	s.Comment = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjects) SetCostStorage(v string) *ListProjectsResponseBodyDataProjects {
	s.CostStorage = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjects) SetCreatedTime(v int64) *ListProjectsResponseBodyDataProjects {
	s.CreatedTime = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjects) SetDefaultQuota(v string) *ListProjectsResponseBodyDataProjects {
	s.DefaultQuota = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjects) SetIpWhiteList(v *ListProjectsResponseBodyDataProjectsIpWhiteList) *ListProjectsResponseBodyDataProjects {
	s.IpWhiteList = v
	return s
}

func (s *ListProjectsResponseBodyDataProjects) SetName(v string) *ListProjectsResponseBodyDataProjects {
	s.Name = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjects) SetOwner(v string) *ListProjectsResponseBodyDataProjects {
	s.Owner = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjects) SetProperties(v *ListProjectsResponseBodyDataProjectsProperties) *ListProjectsResponseBodyDataProjects {
	s.Properties = v
	return s
}

func (s *ListProjectsResponseBodyDataProjects) SetRegionId(v string) *ListProjectsResponseBodyDataProjects {
	s.RegionId = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjects) SetSaleTag(v *ListProjectsResponseBodyDataProjectsSaleTag) *ListProjectsResponseBodyDataProjects {
	s.SaleTag = v
	return s
}

func (s *ListProjectsResponseBodyDataProjects) SetSecurityProperties(v *ListProjectsResponseBodyDataProjectsSecurityProperties) *ListProjectsResponseBodyDataProjects {
	s.SecurityProperties = v
	return s
}

func (s *ListProjectsResponseBodyDataProjects) SetStatus(v string) *ListProjectsResponseBodyDataProjects {
	s.Status = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjects) SetThreeTierModel(v bool) *ListProjectsResponseBodyDataProjects {
	s.ThreeTierModel = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjects) SetType(v string) *ListProjectsResponseBodyDataProjects {
	s.Type = &v
	return s
}

type ListProjectsResponseBodyDataProjectsIpWhiteList struct {
	// The IP address whitelist for access over the Internet or the network for interconnecting with other Alibaba Cloud services.
	//
	// >  If you configure only the IP address whitelist for access over the Internet or the network for interconnecting with other Alibaba Cloud services, the access over the Internet or the network for interconnecting with other Alibaba Cloud services is subject to configurations, and access over a virtual private cloud (VPC) is not allowed.
	//
	// example:
	//
	// 10.88.111.3
	IpList *string `json:"ipList,omitempty" xml:"ipList,omitempty"`
	// The IP address whitelist for access over a VPC.
	//
	// >  If you configure only the IP address whitelist for access over a VPC, the access over a VPC is subject to configurations, and the access over the Internet or the network for interconnecting with other Alibaba Cloud services is not allowed.
	//
	// example:
	//
	// 10.88.111.3
	VpcIpList *string `json:"vpcIpList,omitempty" xml:"vpcIpList,omitempty"`
}

func (s ListProjectsResponseBodyDataProjectsIpWhiteList) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponseBodyDataProjectsIpWhiteList) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyDataProjectsIpWhiteList) SetIpList(v string) *ListProjectsResponseBodyDataProjectsIpWhiteList {
	s.IpList = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsIpWhiteList) SetVpcIpList(v string) *ListProjectsResponseBodyDataProjectsIpWhiteList {
	s.VpcIpList = &v
	return s
}

type ListProjectsResponseBodyDataProjectsProperties struct {
	// Indicates whether a full table scan is allowed in the project. A full table scan occupies a large number of resources, which reduces data processing efficiency. By default, the full table scan feature is disabled.
	//
	// example:
	//
	// false
	AllowFullScan *bool `json:"allowFullScan,omitempty" xml:"allowFullScan,omitempty"`
	// Indicates whether the DECIMAL type of the MaxCompute V2.0 data type edition is enabled.
	//
	// example:
	//
	// true
	EnableDecimal2 *bool `json:"enableDecimal2,omitempty" xml:"enableDecimal2,omitempty"`
	// Indicates whether the routing of the Tunnel resource group is enabled.
	//
	// 	- true: The data transfer tasks that are submitted by the project by default use the Tunnel resource group that is bound to the project.
	//
	// 	- false: The data transfer tasks that are submitted by the project by default use the Tunnel shared resource group.
	//
	// example:
	//
	// true
	EnableTunnelQuotaRoute *bool `json:"enableTunnelQuotaRoute,omitempty" xml:"enableTunnelQuotaRoute,omitempty"`
	// The storage encryption properties.
	Encryption *ListProjectsResponseBodyDataProjectsPropertiesEncryption `json:"encryption,omitempty" xml:"encryption,omitempty" type:"Struct"`
	// The properties of the external project.
	ExternalProjectProperties *ListProjectsResponseBodyDataProjectsPropertiesExternalProjectProperties `json:"externalProjectProperties,omitempty" xml:"externalProjectProperties,omitempty" type:"Struct"`
	// The retention period for backup data. Unit: days. During the retention period, you can restore data of the version in use to the backup data of any version. Valid values: [0,30]. Default value: 1. The value 0 indicates that the backup feature is disabled.
	//
	// example:
	//
	// 1
	RetentionDays *int64 `json:"retentionDays,omitempty" xml:"retentionDays,omitempty"`
	// The maximum consumption threshold of a single SQL statement. Formula: Amount of scanned data (GB) × Complexity.
	//
	// example:
	//
	// 1500
	SqlMeteringMax *string `json:"sqlMeteringMax,omitempty" xml:"sqlMeteringMax,omitempty"`
	// The table lifecycle properties.
	TableLifecycle *ListProjectsResponseBodyDataProjectsPropertiesTableLifecycle `json:"tableLifecycle,omitempty" xml:"tableLifecycle,omitempty" type:"Struct"`
	// The time zone that is used by your project. The time zone is the same as the time zone specified by `odps.sql.timezone`.
	//
	// example:
	//
	// Asia/Shanghai
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
	// The [Tunnel](https://www.alibabacloud.com/help/zh/maxcompute/user-guide/overview-of-dts) resource group that is bound to the project.
	//
	// 	- Default resource group: The Tunnel shared resource group is used. You cannot use the subscription-based Tunnel resource group for the project. The default resource group is automatically used by the Tunnel service of your project, regardless of the parameter setting.
	//
	// 	- Subscription-based Tunnel resource group: You can use the subscription-based Tunnel resource group for the project.
	//
	// example:
	//
	// quota_tunnel
	TunnelQuota *string `json:"tunnelQuota,omitempty" xml:"tunnelQuota,omitempty"`
	// The data type edition. Valid values:
	//
	// 	- **1**: MaxCompute V1.0 data type edition
	//
	// 	- **2**: MaxCompute V2.0 data type edition
	//
	// 	- **hive**: Hive-compatible data type edition
	//
	// For more information about the differences among the three data type editions, see [Data type editions](https://www.alibabacloud.com/help/zh/maxcompute/user-guide/data-type-editions).
	//
	// example:
	//
	// 2
	TypeSystem *string `json:"typeSystem,omitempty" xml:"typeSystem,omitempty"`
}

func (s ListProjectsResponseBodyDataProjectsProperties) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponseBodyDataProjectsProperties) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyDataProjectsProperties) SetAllowFullScan(v bool) *ListProjectsResponseBodyDataProjectsProperties {
	s.AllowFullScan = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsProperties) SetEnableDecimal2(v bool) *ListProjectsResponseBodyDataProjectsProperties {
	s.EnableDecimal2 = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsProperties) SetEnableTunnelQuotaRoute(v bool) *ListProjectsResponseBodyDataProjectsProperties {
	s.EnableTunnelQuotaRoute = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsProperties) SetEncryption(v *ListProjectsResponseBodyDataProjectsPropertiesEncryption) *ListProjectsResponseBodyDataProjectsProperties {
	s.Encryption = v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsProperties) SetExternalProjectProperties(v *ListProjectsResponseBodyDataProjectsPropertiesExternalProjectProperties) *ListProjectsResponseBodyDataProjectsProperties {
	s.ExternalProjectProperties = v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsProperties) SetRetentionDays(v int64) *ListProjectsResponseBodyDataProjectsProperties {
	s.RetentionDays = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsProperties) SetSqlMeteringMax(v string) *ListProjectsResponseBodyDataProjectsProperties {
	s.SqlMeteringMax = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsProperties) SetTableLifecycle(v *ListProjectsResponseBodyDataProjectsPropertiesTableLifecycle) *ListProjectsResponseBodyDataProjectsProperties {
	s.TableLifecycle = v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsProperties) SetTimezone(v string) *ListProjectsResponseBodyDataProjectsProperties {
	s.Timezone = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsProperties) SetTunnelQuota(v string) *ListProjectsResponseBodyDataProjectsProperties {
	s.TunnelQuota = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsProperties) SetTypeSystem(v string) *ListProjectsResponseBodyDataProjectsProperties {
	s.TypeSystem = &v
	return s
}

type ListProjectsResponseBodyDataProjectsPropertiesEncryption struct {
	// The data encryption algorithm that is supported by the key. Valid values: AES256, AESCTR, and RC4.
	//
	// example:
	//
	// SHA1
	Algorithm *string `json:"algorithm,omitempty" xml:"algorithm,omitempty"`
	// Indicates whether the data encryption feature needs to be enabled for the project. For more information about data encryption, see
	//
	// [Storage encryption](https://www.alibabacloud.com/help/zh/maxcompute/security-and-compliance/storage-encryption).
	//
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// The type of key that is used for data encryption. You can select MaxCompute Default Key or Bring Your Own Key (BYOK) as the key type. If you select MaxCompute Default Key, the default key that is created by MaxCompute is used.
	//
	// example:
	//
	// dafault
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
}

func (s ListProjectsResponseBodyDataProjectsPropertiesEncryption) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponseBodyDataProjectsPropertiesEncryption) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyDataProjectsPropertiesEncryption) SetAlgorithm(v string) *ListProjectsResponseBodyDataProjectsPropertiesEncryption {
	s.Algorithm = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsPropertiesEncryption) SetEnable(v bool) *ListProjectsResponseBodyDataProjectsPropertiesEncryption {
	s.Enable = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsPropertiesEncryption) SetKey(v string) *ListProjectsResponseBodyDataProjectsPropertiesEncryption {
	s.Key = &v
	return s
}

type ListProjectsResponseBodyDataProjectsPropertiesExternalProjectProperties struct {
	// Indicates whether the external project is an external project for [data lakehouse solution 2.0](https://www.alibabacloud.com/help/zh/maxcompute/user-guide/lake-warehouse-integrated-2-0-use-guide).
	//
	// example:
	//
	// true
	IsExternalCatalogBound *string `json:"isExternalCatalogBound,omitempty" xml:"isExternalCatalogBound,omitempty"`
}

func (s ListProjectsResponseBodyDataProjectsPropertiesExternalProjectProperties) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponseBodyDataProjectsPropertiesExternalProjectProperties) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyDataProjectsPropertiesExternalProjectProperties) SetIsExternalCatalogBound(v string) *ListProjectsResponseBodyDataProjectsPropertiesExternalProjectProperties {
	s.IsExternalCatalogBound = &v
	return s
}

type ListProjectsResponseBodyDataProjectsPropertiesTableLifecycle struct {
	// The lifecycle type. Valid values:
	//
	// 	- **mandatory**: The lifecycle clause is required in a table creation statement.
	//
	// 	- **optional**: The lifecycle clause is optional in a table creation statement. If you do not configure a lifecycle for a table, the table does not expire.
	//
	// 	- **inherit**: If you do not configure a lifecycle for a table when you create the table, the value of the odps.table.lifecycle.value parameter is used as the table lifecycle by default.
	//
	// example:
	//
	// optional
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The table lifecycle. Unit: days. Valid values: 1 to 37231. Default value: 37231.
	//
	// example:
	//
	// 37231
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListProjectsResponseBodyDataProjectsPropertiesTableLifecycle) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponseBodyDataProjectsPropertiesTableLifecycle) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyDataProjectsPropertiesTableLifecycle) SetType(v string) *ListProjectsResponseBodyDataProjectsPropertiesTableLifecycle {
	s.Type = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsPropertiesTableLifecycle) SetValue(v string) *ListProjectsResponseBodyDataProjectsPropertiesTableLifecycle {
	s.Value = &v
	return s
}

type ListProjectsResponseBodyDataProjectsSaleTag struct {
	// The instance ID of the default computing quota.
	//
	// example:
	//
	// "aaaa-bbbb"
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// The billing method of the default computing quota.
	//
	// example:
	//
	// "project"
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s ListProjectsResponseBodyDataProjectsSaleTag) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponseBodyDataProjectsSaleTag) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyDataProjectsSaleTag) SetResourceId(v string) *ListProjectsResponseBodyDataProjectsSaleTag {
	s.ResourceId = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsSaleTag) SetResourceType(v string) *ListProjectsResponseBodyDataProjectsSaleTag {
	s.ResourceType = &v
	return s
}

type ListProjectsResponseBodyDataProjectsSecurityProperties struct {
	// Indicates whether the [download control](https://www.alibabacloud.com/help/zh/maxcompute/user-guide/label-based-access-control) feature is enabled. By default, this feature is disabled.
	//
	// example:
	//
	// false
	EnableDownloadPrivilege *bool `json:"enableDownloadPrivilege,omitempty" xml:"enableDownloadPrivilege,omitempty"`
	// Indicates whether the [label-based access control](https://www.alibabacloud.com/help/zh/maxcompute/user-guide/label-based-access-control) feature is enabled. By default, this feature is disabled.
	//
	// example:
	//
	// false
	LabelSecurity *bool `json:"labelSecurity,omitempty" xml:"labelSecurity,omitempty"`
	// Indicates whether to allow the object creator to have the access permissions on the object. The default value is true, which indicates that the object creator has the access permissions on the object.
	//
	// example:
	//
	// true
	ObjectCreatorHasAccessPermission *bool `json:"objectCreatorHasAccessPermission,omitempty" xml:"objectCreatorHasAccessPermission,omitempty"`
	// Indicates whether the object creator has the authorization permissions on the object. The default value is true, which indicates that the object creator has the authorization permissions on the object.
	//
	// example:
	//
	// true
	ObjectCreatorHasGrantPermission *bool `json:"objectCreatorHasGrantPermission,omitempty" xml:"objectCreatorHasGrantPermission,omitempty"`
	// The properties of the [data protection mechanism](https://www.alibabacloud.com/help/zh/maxcompute/security-and-compliance/project-data-protection).
	ProjectProtection *ListProjectsResponseBodyDataProjectsSecurityPropertiesProjectProtection `json:"projectProtection,omitempty" xml:"projectProtection,omitempty" type:"Struct"`
	// Indicates whether the [ACL-based access control](https://www.alibabacloud.com/help/zh/maxcompute/user-guide/acl-based-access-control) feature is enabled. By default, this feature is enabled.
	//
	// example:
	//
	// true
	UsingAcl *bool `json:"usingAcl,omitempty" xml:"usingAcl,omitempty"`
	// Indicates whether the [policy-based access control](https://www.alibabacloud.com/help/zh/maxcompute/user-guide/policy-based-access-control-1) feature is enabled. By default, this feature is enabled.
	//
	// example:
	//
	// true
	UsingPolicy *bool `json:"usingPolicy,omitempty" xml:"usingPolicy,omitempty"`
}

func (s ListProjectsResponseBodyDataProjectsSecurityProperties) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponseBodyDataProjectsSecurityProperties) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyDataProjectsSecurityProperties) SetEnableDownloadPrivilege(v bool) *ListProjectsResponseBodyDataProjectsSecurityProperties {
	s.EnableDownloadPrivilege = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsSecurityProperties) SetLabelSecurity(v bool) *ListProjectsResponseBodyDataProjectsSecurityProperties {
	s.LabelSecurity = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsSecurityProperties) SetObjectCreatorHasAccessPermission(v bool) *ListProjectsResponseBodyDataProjectsSecurityProperties {
	s.ObjectCreatorHasAccessPermission = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsSecurityProperties) SetObjectCreatorHasGrantPermission(v bool) *ListProjectsResponseBodyDataProjectsSecurityProperties {
	s.ObjectCreatorHasGrantPermission = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsSecurityProperties) SetProjectProtection(v *ListProjectsResponseBodyDataProjectsSecurityPropertiesProjectProtection) *ListProjectsResponseBodyDataProjectsSecurityProperties {
	s.ProjectProtection = v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsSecurityProperties) SetUsingAcl(v bool) *ListProjectsResponseBodyDataProjectsSecurityProperties {
	s.UsingAcl = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsSecurityProperties) SetUsingPolicy(v bool) *ListProjectsResponseBodyDataProjectsSecurityProperties {
	s.UsingPolicy = &v
	return s
}

type ListProjectsResponseBodyDataProjectsSecurityPropertiesProjectProtection struct {
	// If you enable the project data protection mechanism, you can configure exception or trusted projects. This allows specified users to transfer data of a specified object to a specified project. The project data protection mechanism does not take effect in all the situations that are specified in the exception policy.
	//
	// example:
	//
	// {
	//
	//       "Version": "1",
	//
	//       "Statement": [
	//
	//             {
	//
	//                   "Effect": "Allow",
	//
	//                   "Principal": "",
	//
	//                   "Action": [
	//
	//                         "odps:[, , ...]"
	//
	//                   ],
	//
	//                   "Resource": "acs:odps:*:",
	//
	//                   "Condition": {
	//
	//                         "StringEquals": {
	//
	//                               "odps:TaskType": [
	//
	//                                     ""
	//
	//                               ]
	//
	//                         }
	//
	//                   }
	//
	//             }
	//
	//       ]
	//
	// }
	ExceptionPolicy *string `json:"exceptionPolicy,omitempty" xml:"exceptionPolicy,omitempty"`
	// Indicates whether the [data protection mechanism](https://www.alibabacloud.com/help/zh/maxcompute/security-and-compliance/project-data-protection) is enabled for the project. This allows or denies data transfer across projects. By default, the data protection mechanism is disabled.
	//
	// example:
	//
	// true
	Protected *bool `json:"protected,omitempty" xml:"protected,omitempty"`
}

func (s ListProjectsResponseBodyDataProjectsSecurityPropertiesProjectProtection) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponseBodyDataProjectsSecurityPropertiesProjectProtection) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyDataProjectsSecurityPropertiesProjectProtection) SetExceptionPolicy(v string) *ListProjectsResponseBodyDataProjectsSecurityPropertiesProjectProtection {
	s.ExceptionPolicy = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsSecurityPropertiesProjectProtection) SetProtected(v bool) *ListProjectsResponseBodyDataProjectsSecurityPropertiesProjectProtection {
	s.Protected = &v
	return s
}

type ListProjectsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProjectsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProjectsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponse) GoString() string {
	return s.String()
}

func (s *ListProjectsResponse) SetHeaders(v map[string]*string) *ListProjectsResponse {
	s.Headers = v
	return s
}

func (s *ListProjectsResponse) SetStatusCode(v int32) *ListProjectsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProjectsResponse) SetBody(v *ListProjectsResponseBody) *ListProjectsResponse {
	s.Body = v
	return s
}

type ListQuotasRequest struct {
	// The billing method of the quota.
	//
	// example:
	//
	// subscription
	BillingType *string `json:"billingType,omitempty" xml:"billingType,omitempty"`
	// Specifies the marker after which the returned list begins.
	//
	// example:
	//
	// cHlvZHBzX3VkZl8xMDExNV8xNDU3NDI4NDkzKg==
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
	// The maximum number of entries to return on each page.
	//
	// example:
	//
	// 100
	MaxItem *int64 `json:"maxItem,omitempty" xml:"maxItem,omitempty"`
	// The service ID.
	//
	// example:
	//
	// ODPS
	ProductId *string `json:"productId,omitempty" xml:"productId,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The cost tag. You can filter out quota objects based on the cost tag. The cost tag is created when you tag a service.
	//
	// example:
	//
	// {"tag":"this_is_tag_demo"}
	SaleTags *string `json:"saleTags,omitempty" xml:"saleTags,omitempty"`
	// The ID of the tenant.
	//
	// example:
	//
	// 280747109771520
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListQuotasRequest) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasRequest) GoString() string {
	return s.String()
}

func (s *ListQuotasRequest) SetBillingType(v string) *ListQuotasRequest {
	s.BillingType = &v
	return s
}

func (s *ListQuotasRequest) SetMarker(v string) *ListQuotasRequest {
	s.Marker = &v
	return s
}

func (s *ListQuotasRequest) SetMaxItem(v int64) *ListQuotasRequest {
	s.MaxItem = &v
	return s
}

func (s *ListQuotasRequest) SetProductId(v string) *ListQuotasRequest {
	s.ProductId = &v
	return s
}

func (s *ListQuotasRequest) SetRegion(v string) *ListQuotasRequest {
	s.Region = &v
	return s
}

func (s *ListQuotasRequest) SetSaleTags(v string) *ListQuotasRequest {
	s.SaleTags = &v
	return s
}

func (s *ListQuotasRequest) SetTenantId(v string) *ListQuotasRequest {
	s.TenantId = &v
	return s
}

type ListQuotasResponseBody struct {
	// A pagination token. Only continuous page turning is supported. If NextToken is not empty, the next page exists. The value of NextToken can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAAdDWBF2w6Olxc+cMPjUtUMo/CvPe4IK7f7kIQFrIZjyc
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The returned data.
	Data *ListQuotasResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Indicates the marker after which the returned list begins.
	//
	// example:
	//
	// cHlvZHBzX3VkZl8xMDExNV8xNDU3NDI4NDkzKg==
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
	// The maximum number of entries returned per page.
	//
	// example:
	//
	// 10
	MaxItem *int64 `json:"maxItem,omitempty" xml:"maxItem,omitempty"`
	// The list of quotas.
	QuotaInfoList []*ListQuotasResponseBodyQuotaInfoList `json:"quotaInfoList,omitempty" xml:"quotaInfoList,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 0bc12e6f16677875480593081d2956
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListQuotasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasResponseBody) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBody) SetNextToken(v string) *ListQuotasResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListQuotasResponseBody) SetData(v *ListQuotasResponseBodyData) *ListQuotasResponseBody {
	s.Data = v
	return s
}

func (s *ListQuotasResponseBody) SetMarker(v string) *ListQuotasResponseBody {
	s.Marker = &v
	return s
}

func (s *ListQuotasResponseBody) SetMaxItem(v int64) *ListQuotasResponseBody {
	s.MaxItem = &v
	return s
}

func (s *ListQuotasResponseBody) SetQuotaInfoList(v []*ListQuotasResponseBodyQuotaInfoList) *ListQuotasResponseBody {
	s.QuotaInfoList = v
	return s
}

func (s *ListQuotasResponseBody) SetRequestId(v string) *ListQuotasResponseBody {
	s.RequestId = &v
	return s
}

type ListQuotasResponseBodyData struct {
	// A pagination token. Only continuous page turning is supported. If NextToken is not empty, the next page exists. The value of NextToken can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// "abcde"
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Indicates the marker after which the returned list begins.
	//
	// example:
	//
	// cHlvZHBzX3VkZl8xMDExNV8xNDU3NDI4NDkzKg==
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
	// The maximum number of entries returned per page.
	//
	// example:
	//
	// 10
	MaxItem *int64 `json:"maxItem,omitempty" xml:"maxItem,omitempty"`
	// The list of quotas.
	QuotaInfoList []*ListQuotasResponseBodyDataQuotaInfoList `json:"quotaInfoList,omitempty" xml:"quotaInfoList,omitempty" type:"Repeated"`
}

func (s ListQuotasResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBodyData) SetNextToken(v string) *ListQuotasResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListQuotasResponseBodyData) SetMarker(v string) *ListQuotasResponseBodyData {
	s.Marker = &v
	return s
}

func (s *ListQuotasResponseBodyData) SetMaxItem(v int64) *ListQuotasResponseBodyData {
	s.MaxItem = &v
	return s
}

func (s *ListQuotasResponseBodyData) SetQuotaInfoList(v []*ListQuotasResponseBodyDataQuotaInfoList) *ListQuotasResponseBodyData {
	s.QuotaInfoList = v
	return s
}

type ListQuotasResponseBodyDataQuotaInfoList struct {
	// The tags.
	Tags []*ListQuotasResponseBodyDataQuotaInfoListTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The information of the order.
	BillingPolicy *ListQuotasResponseBodyDataQuotaInfoListBillingPolicy `json:"billingPolicy,omitempty" xml:"billingPolicy,omitempty" type:"Struct"`
	// The cluster ID.
	//
	// example:
	//
	// AT-120N
	Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// The time when the resource was created.
	//
	// example:
	//
	// 2022-09-06T02:14:44Z
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The ID of the Alibaba Cloud account that is used to create the resource.
	//
	// example:
	//
	// 672863518
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// The quota ID.
	//
	// example:
	//
	// 0
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the quota.
	//
	// example:
	//
	// quota_a
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The alias of the quota.
	//
	// example:
	//
	// quota_nickname
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	// The description of the quota.
	//
	// example:
	//
	// {"minCU":10,
	//
	// "adhocCU":0,
	//
	// "maxCU":10,
	//
	// "schedulerType":"Fair",
	//
	// }
	Parameter map[string]interface{} `json:"parameter,omitempty" xml:"parameter,omitempty"`
	// The ID of the parent resource.
	//
	// example:
	//
	// null
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The identifier of an object in a MaxCompute quota. This identifier is the same as the identifier in the sales bill of Alibaba Cloud. This parameter is used for tags.
	SaleTag *ListQuotasResponseBodyDataQuotaInfoListSaleTag `json:"saleTag,omitempty" xml:"saleTag,omitempty" type:"Struct"`
	// The information of the scheduling plan.
	ScheduleInfo *ListQuotasResponseBodyDataQuotaInfoListScheduleInfo `json:"scheduleInfo,omitempty" xml:"scheduleInfo,omitempty" type:"Struct"`
	// The status of the endpoint group.
	//
	// example:
	//
	// ON
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The information of the level-2 quota.
	SubQuotaInfoList []*ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList `json:"subQuotaInfoList,omitempty" xml:"subQuotaInfoList,omitempty" type:"Repeated"`
	// The tag of the resource for the quota.
	//
	// example:
	//
	// abc
	Tag *string `json:"tag,omitempty" xml:"tag,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 280747109771520
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The type of the resource system. This parameter corresponds to the resourceSystemType parameter of the cluster.
	//
	// example:
	//
	// FUXI_ONLINE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The version number.
	//
	// example:
	//
	// 1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListQuotasResponseBodyDataQuotaInfoList) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasResponseBodyDataQuotaInfoList) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) SetTags(v []*ListQuotasResponseBodyDataQuotaInfoListTags) *ListQuotasResponseBodyDataQuotaInfoList {
	s.Tags = v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) SetBillingPolicy(v *ListQuotasResponseBodyDataQuotaInfoListBillingPolicy) *ListQuotasResponseBodyDataQuotaInfoList {
	s.BillingPolicy = v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) SetCluster(v string) *ListQuotasResponseBodyDataQuotaInfoList {
	s.Cluster = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) SetCreateTime(v int64) *ListQuotasResponseBodyDataQuotaInfoList {
	s.CreateTime = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) SetCreatorId(v string) *ListQuotasResponseBodyDataQuotaInfoList {
	s.CreatorId = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) SetId(v string) *ListQuotasResponseBodyDataQuotaInfoList {
	s.Id = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) SetName(v string) *ListQuotasResponseBodyDataQuotaInfoList {
	s.Name = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) SetNickName(v string) *ListQuotasResponseBodyDataQuotaInfoList {
	s.NickName = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) SetParameter(v map[string]interface{}) *ListQuotasResponseBodyDataQuotaInfoList {
	s.Parameter = v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) SetParentId(v string) *ListQuotasResponseBodyDataQuotaInfoList {
	s.ParentId = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) SetRegionId(v string) *ListQuotasResponseBodyDataQuotaInfoList {
	s.RegionId = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) SetSaleTag(v *ListQuotasResponseBodyDataQuotaInfoListSaleTag) *ListQuotasResponseBodyDataQuotaInfoList {
	s.SaleTag = v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) SetScheduleInfo(v *ListQuotasResponseBodyDataQuotaInfoListScheduleInfo) *ListQuotasResponseBodyDataQuotaInfoList {
	s.ScheduleInfo = v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) SetStatus(v string) *ListQuotasResponseBodyDataQuotaInfoList {
	s.Status = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) SetSubQuotaInfoList(v []*ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) *ListQuotasResponseBodyDataQuotaInfoList {
	s.SubQuotaInfoList = v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) SetTag(v string) *ListQuotasResponseBodyDataQuotaInfoList {
	s.Tag = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) SetTenantId(v string) *ListQuotasResponseBodyDataQuotaInfoList {
	s.TenantId = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) SetType(v string) *ListQuotasResponseBodyDataQuotaInfoList {
	s.Type = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) SetVersion(v string) *ListQuotasResponseBodyDataQuotaInfoList {
	s.Version = &v
	return s
}

type ListQuotasResponseBodyDataQuotaInfoListTags struct {
	// The key of the tag.
	//
	// example:
	//
	// Department
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// a12351qHDP6YEQMt
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListQuotasResponseBodyDataQuotaInfoListTags) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasResponseBodyDataQuotaInfoListTags) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBodyDataQuotaInfoListTags) SetTagKey(v string) *ListQuotasResponseBodyDataQuotaInfoListTags {
	s.TagKey = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListTags) SetTagValue(v string) *ListQuotasResponseBodyDataQuotaInfoListTags {
	s.TagValue = &v
	return s
}

type ListQuotasResponseBodyDataQuotaInfoListBillingPolicy struct {
	// The billing method of the quota. Valid values:
	//
	// 	- subscription: a subscription quota.
	//
	// 	- payasyougo: a pay-as-you-go quota.
	//
	// example:
	//
	// subscription
	BillingMethod *string `json:"billingMethod,omitempty" xml:"billingMethod,omitempty"`
	// The specifications of the order.
	//
	// example:
	//
	// OdpsStandard
	OdpsSpecCode *string `json:"odpsSpecCode,omitempty" xml:"odpsSpecCode,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 45245678
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
}

func (s ListQuotasResponseBodyDataQuotaInfoListBillingPolicy) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasResponseBodyDataQuotaInfoListBillingPolicy) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBodyDataQuotaInfoListBillingPolicy) SetBillingMethod(v string) *ListQuotasResponseBodyDataQuotaInfoListBillingPolicy {
	s.BillingMethod = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListBillingPolicy) SetOdpsSpecCode(v string) *ListQuotasResponseBodyDataQuotaInfoListBillingPolicy {
	s.OdpsSpecCode = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListBillingPolicy) SetOrderId(v string) *ListQuotasResponseBodyDataQuotaInfoListBillingPolicy {
	s.OrderId = &v
	return s
}

type ListQuotasResponseBodyDataQuotaInfoListSaleTag struct {
	// The identifier of an object in a MaxCompute quota. This identifier exists in the sales bill of Alibaba Cloud. You can use this identifier to associate the cost of a quota object with a tag.
	ResourceIds []*string `json:"resourceIds,omitempty" xml:"resourceIds,omitempty" type:"Repeated"`
	// The type of the object. Valid values: quota and project.
	//
	// example:
	//
	// "quota"
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s ListQuotasResponseBodyDataQuotaInfoListSaleTag) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasResponseBodyDataQuotaInfoListSaleTag) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSaleTag) SetResourceIds(v []*string) *ListQuotasResponseBodyDataQuotaInfoListSaleTag {
	s.ResourceIds = v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSaleTag) SetResourceType(v string) *ListQuotasResponseBodyDataQuotaInfoListSaleTag {
	s.ResourceType = &v
	return s
}

type ListQuotasResponseBodyDataQuotaInfoListScheduleInfo struct {
	// The quota plan that takes effect based on the scheduling plan.
	//
	// example:
	//
	// planA
	CurrPlan *string `json:"currPlan,omitempty" xml:"currPlan,omitempty"`
	// The time when the current quota plan is scheduled.
	//
	// example:
	//
	// 0800
	CurrTime *string `json:"currTime,omitempty" xml:"currTime,omitempty"`
	// The next quota plan that will take effect based on the scheduling plan.
	//
	// example:
	//
	// planB
	NextPlan *string `json:"nextPlan,omitempty" xml:"nextPlan,omitempty"`
	// The time when the next quota plan is scheduled.
	//
	// example:
	//
	// 1700
	NextTime *string `json:"nextTime,omitempty" xml:"nextTime,omitempty"`
	// The quota plan that immediately takes effect. If the quota plan that immediately takes effect is different from the current quota plan, this parameter is not empty.
	//
	// example:
	//
	// planC
	OncePlan *string `json:"oncePlan,omitempty" xml:"oncePlan,omitempty"`
	// The time when the quota plan immediately takes effect.
	//
	// example:
	//
	// 1500
	OnceTime *string `json:"onceTime,omitempty" xml:"onceTime,omitempty"`
	// The name of the operator.
	//
	// example:
	//
	// userA
	OperatorName *string `json:"operatorName,omitempty" xml:"operatorName,omitempty"`
	// The time zone of the project.
	//
	// example:
	//
	// UTC+8
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
}

func (s ListQuotasResponseBodyDataQuotaInfoListScheduleInfo) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasResponseBodyDataQuotaInfoListScheduleInfo) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBodyDataQuotaInfoListScheduleInfo) SetCurrPlan(v string) *ListQuotasResponseBodyDataQuotaInfoListScheduleInfo {
	s.CurrPlan = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListScheduleInfo) SetCurrTime(v string) *ListQuotasResponseBodyDataQuotaInfoListScheduleInfo {
	s.CurrTime = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListScheduleInfo) SetNextPlan(v string) *ListQuotasResponseBodyDataQuotaInfoListScheduleInfo {
	s.NextPlan = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListScheduleInfo) SetNextTime(v string) *ListQuotasResponseBodyDataQuotaInfoListScheduleInfo {
	s.NextTime = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListScheduleInfo) SetOncePlan(v string) *ListQuotasResponseBodyDataQuotaInfoListScheduleInfo {
	s.OncePlan = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListScheduleInfo) SetOnceTime(v string) *ListQuotasResponseBodyDataQuotaInfoListScheduleInfo {
	s.OnceTime = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListScheduleInfo) SetOperatorName(v string) *ListQuotasResponseBodyDataQuotaInfoListScheduleInfo {
	s.OperatorName = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListScheduleInfo) SetTimezone(v string) *ListQuotasResponseBodyDataQuotaInfoListScheduleInfo {
	s.Timezone = &v
	return s
}

type ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList struct {
	// The information of the order.
	BillingPolicy *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListBillingPolicy `json:"billingPolicy,omitempty" xml:"billingPolicy,omitempty" type:"Struct"`
	// The cluster ID.
	//
	// example:
	//
	// AT-120N
	Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// The time when the resource was created.
	//
	// example:
	//
	// 2022-09-06T02:14:44Z
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The ID of the Alibaba Cloud account that is used to create the resource.
	//
	// example:
	//
	// 672863518
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// The ID of the level-2 quota.
	//
	// example:
	//
	// 1000048
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the level-2 quota.
	//
	// example:
	//
	// subquotaA
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The nickname of the level-2 quota.
	//
	// example:
	//
	// subquotaA
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	// The description of the quota.
	//
	// example:
	//
	// {"minCU":10,
	//
	// "adhocCU":0,
	//
	// "maxCU":10,
	//
	// "schedulerType":"Fair",
	//
	// }
	Parameter map[string]interface{} `json:"parameter,omitempty" xml:"parameter,omitempty"`
	// The ID of the parent resource.
	//
	// example:
	//
	// 0
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The identifier of an object in a MaxCompute quota. This identifier is the same as the identifier in the sales bill of Alibaba Cloud. This parameter is used for tags.
	SaleTag *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListSaleTag `json:"saleTag,omitempty" xml:"saleTag,omitempty" type:"Struct"`
	// The information of the scheduling plan.
	ScheduleInfo *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo `json:"scheduleInfo,omitempty" xml:"scheduleInfo,omitempty" type:"Struct"`
	// The status of the endpoint group.
	//
	// example:
	//
	// ON
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The tag of the resource for the quota.
	//
	// example:
	//
	// abc
	Tag *string `json:"tag,omitempty" xml:"tag,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 280747109771520
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The type of the resource system. This parameter corresponds to the resourceSystemType parameter of the cluster.
	//
	// example:
	//
	// FUXI_ONLINE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The version of the algorithm image.
	//
	// example:
	//
	// 1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) SetBillingPolicy(v *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListBillingPolicy) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList {
	s.BillingPolicy = v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) SetCluster(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList {
	s.Cluster = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) SetCreateTime(v int64) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList {
	s.CreateTime = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) SetCreatorId(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList {
	s.CreatorId = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) SetId(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList {
	s.Id = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) SetName(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList {
	s.Name = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) SetNickName(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList {
	s.NickName = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) SetParameter(v map[string]interface{}) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList {
	s.Parameter = v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) SetParentId(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList {
	s.ParentId = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) SetRegionId(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList {
	s.RegionId = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) SetSaleTag(v *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListSaleTag) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList {
	s.SaleTag = v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) SetScheduleInfo(v *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList {
	s.ScheduleInfo = v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) SetStatus(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList {
	s.Status = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) SetTag(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList {
	s.Tag = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) SetTenantId(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList {
	s.TenantId = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) SetType(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList {
	s.Type = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) SetVersion(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList {
	s.Version = &v
	return s
}

type ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListBillingPolicy struct {
	// The billing method of the quota. Valid values:
	//
	// 	- subscription: a subscription quota.
	//
	// 	- payasyougo: a pay-as-you-go quota.
	//
	// example:
	//
	// subscription
	BillingMethod *string `json:"billingMethod,omitempty" xml:"billingMethod,omitempty"`
	// The specifications of the order.
	//
	// example:
	//
	// OdpsStandard
	OdpsSpecCode *string `json:"odpsSpecCode,omitempty" xml:"odpsSpecCode,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 45245678
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
}

func (s ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListBillingPolicy) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListBillingPolicy) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListBillingPolicy) SetBillingMethod(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListBillingPolicy {
	s.BillingMethod = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListBillingPolicy) SetOdpsSpecCode(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListBillingPolicy {
	s.OdpsSpecCode = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListBillingPolicy) SetOrderId(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListBillingPolicy {
	s.OrderId = &v
	return s
}

type ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListSaleTag struct {
	// The identifier of an object in a MaxCompute quota. This identifier exists in the sales bill of Alibaba Cloud. You can use this identifier to associate the cost of a quota object with a tag.
	ResourceIds []*string `json:"resourceIds,omitempty" xml:"resourceIds,omitempty" type:"Repeated"`
	// The type of the object. Valid values: quota and project.
	//
	// example:
	//
	// "quota"
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListSaleTag) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListSaleTag) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListSaleTag) SetResourceIds(v []*string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListSaleTag {
	s.ResourceIds = v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListSaleTag) SetResourceType(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListSaleTag {
	s.ResourceType = &v
	return s
}

type ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo struct {
	// The quota plan that takes effect based on the scheduling plan.
	//
	// example:
	//
	// planA
	CurrPlan *string `json:"currPlan,omitempty" xml:"currPlan,omitempty"`
	// The time when the current quota plan is scheduled.
	//
	// example:
	//
	// 0800
	CurrTime *string `json:"currTime,omitempty" xml:"currTime,omitempty"`
	// The next quota plan that will take effect based on the scheduling plan.
	//
	// example:
	//
	// planB
	NextPlan *string `json:"nextPlan,omitempty" xml:"nextPlan,omitempty"`
	// The time when the next quota plan is scheduled.
	//
	// example:
	//
	// 1700
	NextTime *string `json:"nextTime,omitempty" xml:"nextTime,omitempty"`
	// The quota plan that immediately takes effect. If the quota plan that immediately takes effect is different from the current quota plan, this parameter is not empty.
	//
	// example:
	//
	// planC
	OncePlan *string `json:"oncePlan,omitempty" xml:"oncePlan,omitempty"`
	// The time when the quota plan immediately takes effect.
	//
	// example:
	//
	// 1500
	OnceTime *string `json:"onceTime,omitempty" xml:"onceTime,omitempty"`
	// The name of the operator.
	//
	// example:
	//
	// userA
	OperatorName *string `json:"operatorName,omitempty" xml:"operatorName,omitempty"`
	// The time zone of the project.
	//
	// example:
	//
	// UTC+8
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
}

func (s ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo) SetCurrPlan(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo {
	s.CurrPlan = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo) SetCurrTime(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo {
	s.CurrTime = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo) SetNextPlan(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo {
	s.NextPlan = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo) SetNextTime(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo {
	s.NextTime = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo) SetOncePlan(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo {
	s.OncePlan = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo) SetOnceTime(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo {
	s.OnceTime = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo) SetOperatorName(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo {
	s.OperatorName = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo) SetTimezone(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo {
	s.Timezone = &v
	return s
}

type ListQuotasResponseBodyQuotaInfoList struct {
	// The tags.
	Tags []*ListQuotasResponseBodyQuotaInfoListTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The information of the order.
	BillingPolicy *ListQuotasResponseBodyQuotaInfoListBillingPolicy `json:"billingPolicy,omitempty" xml:"billingPolicy,omitempty" type:"Struct"`
	// The cluster ID.
	//
	// example:
	//
	// AT-120N
	Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// The time when the resource was created.
	//
	// example:
	//
	// 2022-09-06T02:14:44Z
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The ID of the Alibaba Cloud account that is used to create the resource.
	//
	// example:
	//
	// 672863518
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// The quota ID.
	//
	// example:
	//
	// 0
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the quota.
	//
	// example:
	//
	// quota_a
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The alias of the quota.
	//
	// example:
	//
	// quota_nickname
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	// The description of the quota.
	//
	// example:
	//
	// {"minCU":10,
	//
	// "adhocCU":0,
	//
	// "maxCU":10,
	//
	// "schedulerType":"Fair",
	//
	// }
	Parameter map[string]interface{} `json:"parameter,omitempty" xml:"parameter,omitempty"`
	// The ID of the parent resource.
	//
	// example:
	//
	// null
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The identifier of an object in a MaxCompute quota. This identifier is the same as the identifier in the sales bill of Alibaba Cloud. This parameter is used for tags.
	SaleTag *ListQuotasResponseBodyQuotaInfoListSaleTag `json:"saleTag,omitempty" xml:"saleTag,omitempty" type:"Struct"`
	// The information of the scheduling plan.
	ScheduleInfo *ListQuotasResponseBodyQuotaInfoListScheduleInfo `json:"scheduleInfo,omitempty" xml:"scheduleInfo,omitempty" type:"Struct"`
	// The status of the endpoint group.
	//
	// example:
	//
	// ON
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The information of the level-2 quota.
	SubQuotaInfoList []*ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList `json:"subQuotaInfoList,omitempty" xml:"subQuotaInfoList,omitempty" type:"Repeated"`
	// The tag of the resource for the quota.
	//
	// example:
	//
	// abc
	Tag *string `json:"tag,omitempty" xml:"tag,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 280747109771520
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The type of the resource system. This parameter corresponds to the resourceSystemType parameter of the cluster.
	//
	// example:
	//
	// FUXI_ONLINE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The version.
	//
	// example:
	//
	// 1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListQuotasResponseBodyQuotaInfoList) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasResponseBodyQuotaInfoList) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBodyQuotaInfoList) SetTags(v []*ListQuotasResponseBodyQuotaInfoListTags) *ListQuotasResponseBodyQuotaInfoList {
	s.Tags = v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoList) SetBillingPolicy(v *ListQuotasResponseBodyQuotaInfoListBillingPolicy) *ListQuotasResponseBodyQuotaInfoList {
	s.BillingPolicy = v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoList) SetCluster(v string) *ListQuotasResponseBodyQuotaInfoList {
	s.Cluster = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoList) SetCreateTime(v int64) *ListQuotasResponseBodyQuotaInfoList {
	s.CreateTime = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoList) SetCreatorId(v string) *ListQuotasResponseBodyQuotaInfoList {
	s.CreatorId = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoList) SetId(v string) *ListQuotasResponseBodyQuotaInfoList {
	s.Id = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoList) SetName(v string) *ListQuotasResponseBodyQuotaInfoList {
	s.Name = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoList) SetNickName(v string) *ListQuotasResponseBodyQuotaInfoList {
	s.NickName = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoList) SetParameter(v map[string]interface{}) *ListQuotasResponseBodyQuotaInfoList {
	s.Parameter = v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoList) SetParentId(v string) *ListQuotasResponseBodyQuotaInfoList {
	s.ParentId = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoList) SetRegionId(v string) *ListQuotasResponseBodyQuotaInfoList {
	s.RegionId = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoList) SetSaleTag(v *ListQuotasResponseBodyQuotaInfoListSaleTag) *ListQuotasResponseBodyQuotaInfoList {
	s.SaleTag = v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoList) SetScheduleInfo(v *ListQuotasResponseBodyQuotaInfoListScheduleInfo) *ListQuotasResponseBodyQuotaInfoList {
	s.ScheduleInfo = v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoList) SetStatus(v string) *ListQuotasResponseBodyQuotaInfoList {
	s.Status = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoList) SetSubQuotaInfoList(v []*ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) *ListQuotasResponseBodyQuotaInfoList {
	s.SubQuotaInfoList = v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoList) SetTag(v string) *ListQuotasResponseBodyQuotaInfoList {
	s.Tag = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoList) SetTenantId(v string) *ListQuotasResponseBodyQuotaInfoList {
	s.TenantId = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoList) SetType(v string) *ListQuotasResponseBodyQuotaInfoList {
	s.Type = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoList) SetVersion(v string) *ListQuotasResponseBodyQuotaInfoList {
	s.Version = &v
	return s
}

type ListQuotasResponseBodyQuotaInfoListTags struct {
	// The key of the tag.
	//
	// example:
	//
	// Department
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// 1
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListQuotasResponseBodyQuotaInfoListTags) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasResponseBodyQuotaInfoListTags) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBodyQuotaInfoListTags) SetTagKey(v string) *ListQuotasResponseBodyQuotaInfoListTags {
	s.TagKey = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListTags) SetTagValue(v string) *ListQuotasResponseBodyQuotaInfoListTags {
	s.TagValue = &v
	return s
}

type ListQuotasResponseBodyQuotaInfoListBillingPolicy struct {
	// The billing method of the quota. Valid values:
	//
	// 	- subscription: a subscription quota.
	//
	// 	- payasyougo: a pay-as-you-go quota.
	//
	// example:
	//
	// subscription
	BillingMethod *string `json:"billingMethod,omitempty" xml:"billingMethod,omitempty"`
	// The specifications of the order.
	//
	// example:
	//
	// OdpsStandard
	OdpsSpecCode *string `json:"odpsSpecCode,omitempty" xml:"odpsSpecCode,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 45245678
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
}

func (s ListQuotasResponseBodyQuotaInfoListBillingPolicy) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasResponseBodyQuotaInfoListBillingPolicy) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBodyQuotaInfoListBillingPolicy) SetBillingMethod(v string) *ListQuotasResponseBodyQuotaInfoListBillingPolicy {
	s.BillingMethod = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListBillingPolicy) SetOdpsSpecCode(v string) *ListQuotasResponseBodyQuotaInfoListBillingPolicy {
	s.OdpsSpecCode = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListBillingPolicy) SetOrderId(v string) *ListQuotasResponseBodyQuotaInfoListBillingPolicy {
	s.OrderId = &v
	return s
}

type ListQuotasResponseBodyQuotaInfoListSaleTag struct {
	// The identifier of an object in a MaxCompute quota. This identifier exists in the sales bill of Alibaba Cloud. You can use this identifier to associate the cost of a quota object with a tag.
	ResourceIds []*string `json:"resourceIds,omitempty" xml:"resourceIds,omitempty" type:"Repeated"`
	// The type of the object. Valid values: quota and project.
	//
	// example:
	//
	// "project"
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s ListQuotasResponseBodyQuotaInfoListSaleTag) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasResponseBodyQuotaInfoListSaleTag) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBodyQuotaInfoListSaleTag) SetResourceIds(v []*string) *ListQuotasResponseBodyQuotaInfoListSaleTag {
	s.ResourceIds = v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSaleTag) SetResourceType(v string) *ListQuotasResponseBodyQuotaInfoListSaleTag {
	s.ResourceType = &v
	return s
}

type ListQuotasResponseBodyQuotaInfoListScheduleInfo struct {
	// The quota plan that takes effect based on the scheduling plan.
	//
	// example:
	//
	// planA
	CurrPlan *string `json:"currPlan,omitempty" xml:"currPlan,omitempty"`
	// The time when the current quota plan is scheduled.
	//
	// example:
	//
	// 0800
	CurrTime *string `json:"currTime,omitempty" xml:"currTime,omitempty"`
	// The next quota plan that will take effect based on the scheduling plan.
	//
	// example:
	//
	// planB
	NextPlan *string `json:"nextPlan,omitempty" xml:"nextPlan,omitempty"`
	// The time when the next quota plan is scheduled.
	//
	// example:
	//
	// 1700
	NextTime *string `json:"nextTime,omitempty" xml:"nextTime,omitempty"`
	// The quota plan that immediately takes effect. If the quota plan that immediately takes effect is different from the current quota plan, this parameter is not empty.
	//
	// example:
	//
	// planC
	OncePlan *string `json:"oncePlan,omitempty" xml:"oncePlan,omitempty"`
	// The time when the quota plan immediately takes effect.
	//
	// example:
	//
	// 1500
	OnceTime *string `json:"onceTime,omitempty" xml:"onceTime,omitempty"`
	// The name of the operator.
	//
	// example:
	//
	// userA
	OperatorName *string `json:"operatorName,omitempty" xml:"operatorName,omitempty"`
	// The time zone of the project.
	//
	// example:
	//
	// UTC+8
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
}

func (s ListQuotasResponseBodyQuotaInfoListScheduleInfo) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasResponseBodyQuotaInfoListScheduleInfo) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBodyQuotaInfoListScheduleInfo) SetCurrPlan(v string) *ListQuotasResponseBodyQuotaInfoListScheduleInfo {
	s.CurrPlan = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListScheduleInfo) SetCurrTime(v string) *ListQuotasResponseBodyQuotaInfoListScheduleInfo {
	s.CurrTime = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListScheduleInfo) SetNextPlan(v string) *ListQuotasResponseBodyQuotaInfoListScheduleInfo {
	s.NextPlan = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListScheduleInfo) SetNextTime(v string) *ListQuotasResponseBodyQuotaInfoListScheduleInfo {
	s.NextTime = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListScheduleInfo) SetOncePlan(v string) *ListQuotasResponseBodyQuotaInfoListScheduleInfo {
	s.OncePlan = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListScheduleInfo) SetOnceTime(v string) *ListQuotasResponseBodyQuotaInfoListScheduleInfo {
	s.OnceTime = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListScheduleInfo) SetOperatorName(v string) *ListQuotasResponseBodyQuotaInfoListScheduleInfo {
	s.OperatorName = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListScheduleInfo) SetTimezone(v string) *ListQuotasResponseBodyQuotaInfoListScheduleInfo {
	s.Timezone = &v
	return s
}

type ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList struct {
	// The information of the order.
	BillingPolicy *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListBillingPolicy `json:"billingPolicy,omitempty" xml:"billingPolicy,omitempty" type:"Struct"`
	// The cluster ID.
	//
	// example:
	//
	// AT-120N
	Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// The time when the resource was created.
	//
	// example:
	//
	// 2022-09-06T02:14:44Z
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The ID of the Alibaba Cloud account that is used to create the resource.
	//
	// example:
	//
	// 672863518
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// The ID of the level-2 quota.
	//
	// example:
	//
	// 1000048
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the level-2 quota.
	//
	// example:
	//
	// subquotaA
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The alias of the level-2 quota.
	//
	// example:
	//
	// subquotaA
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	// The description of the quota.
	//
	// example:
	//
	// {"minCU":10,
	//
	// "adhocCU":0,
	//
	// "maxCU":10,
	//
	// "schedulerType":"Fair",
	//
	// }
	Parameter map[string]interface{} `json:"parameter,omitempty" xml:"parameter,omitempty"`
	// The ID of the parent resource.
	//
	// example:
	//
	// 0
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The identifier of an object in a MaxCompute quota. This identifier is the same as the identifier in the sales bill of Alibaba Cloud. This parameter is used for tags.
	SaleTag *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListSaleTag `json:"saleTag,omitempty" xml:"saleTag,omitempty" type:"Struct"`
	// The information of the scheduling plan.
	ScheduleInfo *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo `json:"scheduleInfo,omitempty" xml:"scheduleInfo,omitempty" type:"Struct"`
	// The status of the endpoint group.
	//
	// example:
	//
	// ON
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The tag of the resource for the quota.
	//
	// example:
	//
	// abc
	Tag *string `json:"tag,omitempty" xml:"tag,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 280747109771520
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The type of the resource system. This parameter corresponds to the resourceSystemType parameter of the cluster.
	//
	// example:
	//
	// FUXI_ONLINE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The version number.
	//
	// example:
	//
	// 1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) SetBillingPolicy(v *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListBillingPolicy) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList {
	s.BillingPolicy = v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) SetCluster(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList {
	s.Cluster = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) SetCreateTime(v int64) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList {
	s.CreateTime = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) SetCreatorId(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList {
	s.CreatorId = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) SetId(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList {
	s.Id = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) SetName(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList {
	s.Name = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) SetNickName(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList {
	s.NickName = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) SetParameter(v map[string]interface{}) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList {
	s.Parameter = v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) SetParentId(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList {
	s.ParentId = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) SetRegionId(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList {
	s.RegionId = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) SetSaleTag(v *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListSaleTag) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList {
	s.SaleTag = v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) SetScheduleInfo(v *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList {
	s.ScheduleInfo = v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) SetStatus(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList {
	s.Status = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) SetTag(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList {
	s.Tag = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) SetTenantId(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList {
	s.TenantId = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) SetType(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList {
	s.Type = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) SetVersion(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList {
	s.Version = &v
	return s
}

type ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListBillingPolicy struct {
	// The billing method of the quota. Valid values:
	//
	// 	- subscription: a subscription quota.
	//
	// 	- payasyougo: a pay-as-you-go quota.
	//
	// example:
	//
	// subscription
	BillingMethod *string `json:"billingMethod,omitempty" xml:"billingMethod,omitempty"`
	// The specifications of the order.
	//
	// example:
	//
	// OdpsStandard
	OdpsSpecCode *string `json:"odpsSpecCode,omitempty" xml:"odpsSpecCode,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 45245678
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
}

func (s ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListBillingPolicy) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListBillingPolicy) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListBillingPolicy) SetBillingMethod(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListBillingPolicy {
	s.BillingMethod = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListBillingPolicy) SetOdpsSpecCode(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListBillingPolicy {
	s.OdpsSpecCode = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListBillingPolicy) SetOrderId(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListBillingPolicy {
	s.OrderId = &v
	return s
}

type ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListSaleTag struct {
	// The identifier of an object in a MaxCompute quota. This identifier exists in the sales bill of Alibaba Cloud. You can use this identifier to associate the cost of a quota object with a tag.
	ResourceIds []*string `json:"resourceIds,omitempty" xml:"resourceIds,omitempty" type:"Repeated"`
	// The type of the object. Valid values: quota and project.
	//
	// example:
	//
	// "quota"
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListSaleTag) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListSaleTag) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListSaleTag) SetResourceIds(v []*string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListSaleTag {
	s.ResourceIds = v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListSaleTag) SetResourceType(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListSaleTag {
	s.ResourceType = &v
	return s
}

type ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo struct {
	// The quota plan that takes effect based on the scheduling plan.
	//
	// example:
	//
	// planA
	CurrPlan *string `json:"currPlan,omitempty" xml:"currPlan,omitempty"`
	// The time when the current quota plan is scheduled.
	//
	// example:
	//
	// 0800
	CurrTime *string `json:"currTime,omitempty" xml:"currTime,omitempty"`
	// The next quota plan that will take effect based on the scheduling plan.
	//
	// example:
	//
	// planB
	NextPlan *string `json:"nextPlan,omitempty" xml:"nextPlan,omitempty"`
	// The time when the next quota plan is scheduled.
	//
	// example:
	//
	// 1700
	NextTime *string `json:"nextTime,omitempty" xml:"nextTime,omitempty"`
	// The quota plan that immediately takes effect. If the quota plan that immediately takes effect is different from the current quota plan, this parameter is not empty.
	//
	// example:
	//
	// planC
	OncePlan *string `json:"oncePlan,omitempty" xml:"oncePlan,omitempty"`
	// The time when the quota plan immediately takes effect.
	//
	// example:
	//
	// 1500
	OnceTime *string `json:"onceTime,omitempty" xml:"onceTime,omitempty"`
	// The name of the operator.
	//
	// example:
	//
	// userA
	OperatorName *string `json:"operatorName,omitempty" xml:"operatorName,omitempty"`
	// The time zone of the project.
	//
	// example:
	//
	// UTC+8
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
}

func (s ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo) SetCurrPlan(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo {
	s.CurrPlan = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo) SetCurrTime(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo {
	s.CurrTime = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo) SetNextPlan(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo {
	s.NextPlan = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo) SetNextTime(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo {
	s.NextTime = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo) SetOncePlan(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo {
	s.OncePlan = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo) SetOnceTime(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo {
	s.OnceTime = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo) SetOperatorName(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo {
	s.OperatorName = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo) SetTimezone(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo {
	s.Timezone = &v
	return s
}

type ListQuotasResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListQuotasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListQuotasResponse) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasResponse) GoString() string {
	return s.String()
}

func (s *ListQuotasResponse) SetHeaders(v map[string]*string) *ListQuotasResponse {
	s.Headers = v
	return s
}

func (s *ListQuotasResponse) SetStatusCode(v int32) *ListQuotasResponse {
	s.StatusCode = &v
	return s
}

func (s *ListQuotasResponse) SetBody(v *ListQuotasResponseBody) *ListQuotasResponse {
	s.Body = v
	return s
}

type ListQuotasPlansRequest struct {
	// The ID of the region.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The ID of the tenant.
	//
	// example:
	//
	// 478403690625249
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListQuotasPlansRequest) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasPlansRequest) GoString() string {
	return s.String()
}

func (s *ListQuotasPlansRequest) SetRegion(v string) *ListQuotasPlansRequest {
	s.Region = &v
	return s
}

func (s *ListQuotasPlansRequest) SetTenantId(v string) *ListQuotasPlansRequest {
	s.TenantId = &v
	return s
}

type ListQuotasPlansResponseBody struct {
	// The returned data.
	Data *ListQuotasPlansResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 0be3e0bd16661643917136451ebf55
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListQuotasPlansResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasPlansResponseBody) GoString() string {
	return s.String()
}

func (s *ListQuotasPlansResponseBody) SetData(v *ListQuotasPlansResponseBodyData) *ListQuotasPlansResponseBody {
	s.Data = v
	return s
}

func (s *ListQuotasPlansResponseBody) SetRequestId(v string) *ListQuotasPlansResponseBody {
	s.RequestId = &v
	return s
}

type ListQuotasPlansResponseBodyData struct {
	// The list of quota plans.
	PlanList []*ListQuotasPlansResponseBodyDataPlanList `json:"planList,omitempty" xml:"planList,omitempty" type:"Repeated"`
}

func (s ListQuotasPlansResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasPlansResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListQuotasPlansResponseBodyData) SetPlanList(v []*ListQuotasPlansResponseBodyDataPlanList) *ListQuotasPlansResponseBodyData {
	s.PlanList = v
	return s
}

type ListQuotasPlansResponseBodyDataPlanList struct {
	// The time when the quota plan was created.
	//
	// example:
	//
	// 2022-05-16T06:07:45Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The name of the quota plan.
	//
	// example:
	//
	// planA
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The details of the quota.
	Quota *ListQuotasPlansResponseBodyDataPlanListQuota `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
}

func (s ListQuotasPlansResponseBodyDataPlanList) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasPlansResponseBodyDataPlanList) GoString() string {
	return s.String()
}

func (s *ListQuotasPlansResponseBodyDataPlanList) SetCreateTime(v string) *ListQuotasPlansResponseBodyDataPlanList {
	s.CreateTime = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanList) SetName(v string) *ListQuotasPlansResponseBodyDataPlanList {
	s.Name = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanList) SetQuota(v *ListQuotasPlansResponseBodyDataPlanListQuota) *ListQuotasPlansResponseBodyDataPlanList {
	s.Quota = v
	return s
}

type ListQuotasPlansResponseBodyDataPlanListQuota struct {
	// The information of the order.
	BillingPolicy *ListQuotasPlansResponseBodyDataPlanListQuotaBillingPolicy `json:"billingPolicy,omitempty" xml:"billingPolicy,omitempty" type:"Struct"`
	// The ID of the cluster.
	//
	// example:
	//
	// AT-120N
	Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// The time when the quota plan was created.
	//
	// example:
	//
	// 2022-09-06T02:14:44Z
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The ID of the Alibaba Cloud account that is used to create the resource.
	//
	// example:
	//
	// 672863518
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// The ID of the quota.
	//
	// example:
	//
	// 0
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the quota.
	//
	// example:
	//
	// quota_a
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The alias of the quota.
	//
	// example:
	//
	// quota_nickname
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	// The description of the quota.
	Parameter map[string]interface{} `json:"parameter,omitempty" xml:"parameter,omitempty"`
	// The ID of the parent resource.
	//
	// example:
	//
	// null
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The information of the scheduling plan.
	ScheduleInfo *ListQuotasPlansResponseBodyDataPlanListQuotaScheduleInfo `json:"scheduleInfo,omitempty" xml:"scheduleInfo,omitempty" type:"Struct"`
	// The status of the resource.
	//
	// example:
	//
	// ON
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The information of the level-2 quota.
	SubQuotaInfoList []*ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList `json:"subQuotaInfoList,omitempty" xml:"subQuotaInfoList,omitempty" type:"Repeated"`
	// The tag of the resource for the quota.
	//
	// example:
	//
	// abc
	Tag *string `json:"tag,omitempty" xml:"tag,omitempty"`
	// The ID of the tenant.
	//
	// example:
	//
	// 478403690625249
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The type of the resource system. This parameter corresponds to the resourceSystemType parameter of the cluster.
	//
	// example:
	//
	// FUXI_ONLINE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The version number.
	//
	// example:
	//
	// 1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListQuotasPlansResponseBodyDataPlanListQuota) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasPlansResponseBodyDataPlanListQuota) GoString() string {
	return s.String()
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuota) SetBillingPolicy(v *ListQuotasPlansResponseBodyDataPlanListQuotaBillingPolicy) *ListQuotasPlansResponseBodyDataPlanListQuota {
	s.BillingPolicy = v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuota) SetCluster(v string) *ListQuotasPlansResponseBodyDataPlanListQuota {
	s.Cluster = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuota) SetCreateTime(v int64) *ListQuotasPlansResponseBodyDataPlanListQuota {
	s.CreateTime = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuota) SetCreatorId(v string) *ListQuotasPlansResponseBodyDataPlanListQuota {
	s.CreatorId = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuota) SetId(v string) *ListQuotasPlansResponseBodyDataPlanListQuota {
	s.Id = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuota) SetName(v string) *ListQuotasPlansResponseBodyDataPlanListQuota {
	s.Name = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuota) SetNickName(v string) *ListQuotasPlansResponseBodyDataPlanListQuota {
	s.NickName = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuota) SetParameter(v map[string]interface{}) *ListQuotasPlansResponseBodyDataPlanListQuota {
	s.Parameter = v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuota) SetParentId(v string) *ListQuotasPlansResponseBodyDataPlanListQuota {
	s.ParentId = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuota) SetRegionId(v string) *ListQuotasPlansResponseBodyDataPlanListQuota {
	s.RegionId = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuota) SetScheduleInfo(v *ListQuotasPlansResponseBodyDataPlanListQuotaScheduleInfo) *ListQuotasPlansResponseBodyDataPlanListQuota {
	s.ScheduleInfo = v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuota) SetStatus(v string) *ListQuotasPlansResponseBodyDataPlanListQuota {
	s.Status = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuota) SetSubQuotaInfoList(v []*ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) *ListQuotasPlansResponseBodyDataPlanListQuota {
	s.SubQuotaInfoList = v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuota) SetTag(v string) *ListQuotasPlansResponseBodyDataPlanListQuota {
	s.Tag = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuota) SetTenantId(v string) *ListQuotasPlansResponseBodyDataPlanListQuota {
	s.TenantId = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuota) SetType(v string) *ListQuotasPlansResponseBodyDataPlanListQuota {
	s.Type = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuota) SetVersion(v string) *ListQuotasPlansResponseBodyDataPlanListQuota {
	s.Version = &v
	return s
}

type ListQuotasPlansResponseBodyDataPlanListQuotaBillingPolicy struct {
	// The billing method of the quota. Valid values:
	//
	// 	- subscription: a subscription quota.
	//
	// 	- payasyougo: a pay-as-you-go quota.
	//
	// example:
	//
	// subscription
	BillingMethod *string `json:"billingMethod,omitempty" xml:"billingMethod,omitempty"`
	// The specifications of the order.
	//
	// example:
	//
	// OdpsStandard
	OdpsSpecCode *string `json:"odpsSpecCode,omitempty" xml:"odpsSpecCode,omitempty"`
	// The ID of the order.
	//
	// example:
	//
	// 45245678
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
}

func (s ListQuotasPlansResponseBodyDataPlanListQuotaBillingPolicy) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasPlansResponseBodyDataPlanListQuotaBillingPolicy) GoString() string {
	return s.String()
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaBillingPolicy) SetBillingMethod(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaBillingPolicy {
	s.BillingMethod = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaBillingPolicy) SetOdpsSpecCode(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaBillingPolicy {
	s.OdpsSpecCode = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaBillingPolicy) SetOrderId(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaBillingPolicy {
	s.OrderId = &v
	return s
}

type ListQuotasPlansResponseBodyDataPlanListQuotaScheduleInfo struct {
	// The quota plan that takes effect based on the scheduling plan.
	//
	// example:
	//
	// planA
	CurrPlan *string `json:"currPlan,omitempty" xml:"currPlan,omitempty"`
	// The time when the current quota plan is scheduled.
	//
	// example:
	//
	// 0800
	CurrTime *string `json:"currTime,omitempty" xml:"currTime,omitempty"`
	// The next quota plan that will take effect based on the scheduling plan.
	//
	// example:
	//
	// planB
	NextPlan *string `json:"nextPlan,omitempty" xml:"nextPlan,omitempty"`
	// The time when the next quota plan is scheduled.
	//
	// example:
	//
	// 1700
	NextTime *string `json:"nextTime,omitempty" xml:"nextTime,omitempty"`
	// If the quota plan that immediately takes effect is different from the current quota plan, this parameter is not empty.
	//
	// example:
	//
	// planC
	OncePlan *string `json:"oncePlan,omitempty" xml:"oncePlan,omitempty"`
	// The time when the quota plan immediately takes effect.
	//
	// example:
	//
	// 1500
	OnceTime *string `json:"onceTime,omitempty" xml:"onceTime,omitempty"`
	// The name of the operator.
	//
	// example:
	//
	// userA
	OperatorName *string `json:"operatorName,omitempty" xml:"operatorName,omitempty"`
}

func (s ListQuotasPlansResponseBodyDataPlanListQuotaScheduleInfo) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasPlansResponseBodyDataPlanListQuotaScheduleInfo) GoString() string {
	return s.String()
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaScheduleInfo) SetCurrPlan(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaScheduleInfo {
	s.CurrPlan = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaScheduleInfo) SetCurrTime(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaScheduleInfo {
	s.CurrTime = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaScheduleInfo) SetNextPlan(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaScheduleInfo {
	s.NextPlan = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaScheduleInfo) SetNextTime(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaScheduleInfo {
	s.NextTime = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaScheduleInfo) SetOncePlan(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaScheduleInfo {
	s.OncePlan = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaScheduleInfo) SetOnceTime(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaScheduleInfo {
	s.OnceTime = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaScheduleInfo) SetOperatorName(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaScheduleInfo {
	s.OperatorName = &v
	return s
}

type ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList struct {
	// The information of the order.
	BillingPolicy *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListBillingPolicy `json:"billingPolicy,omitempty" xml:"billingPolicy,omitempty" type:"Struct"`
	// The ID of the cluster.
	//
	// example:
	//
	// AT-120N
	Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// The time when the quota plan was created.
	//
	// example:
	//
	// 2022-09-06T02:14:44Z
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The ID of the Alibaba Cloud account that is used to create the resource.
	//
	// example:
	//
	// 672863518
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// The ID of the level-2 quota.
	//
	// example:
	//
	// 1000048
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the level-2 quota.
	//
	// example:
	//
	// subquotaA
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The nickname of the level-2 quota.
	//
	// example:
	//
	// subquotaA
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	// The description of the quota.
	Parameter map[string]interface{} `json:"parameter,omitempty" xml:"parameter,omitempty"`
	// The ID of the parent resource.
	//
	// example:
	//
	// 0
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The information of the scheduling plan.
	ScheduleInfo *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListScheduleInfo `json:"scheduleInfo,omitempty" xml:"scheduleInfo,omitempty" type:"Struct"`
	// The status of the resource.
	//
	// example:
	//
	// ON
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The tag of the resource for the quota.
	//
	// example:
	//
	// abc
	Tag *string `json:"tag,omitempty" xml:"tag,omitempty"`
	// The ID of the tenant.
	//
	// example:
	//
	// 478403690625249
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The type of the resource system. This parameter corresponds to the resourceSystemType parameter of the cluster.
	//
	// example:
	//
	// FUXI_ONLINE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The version number.
	//
	// example:
	//
	// 1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) GoString() string {
	return s.String()
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) SetBillingPolicy(v *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListBillingPolicy) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.BillingPolicy = v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) SetCluster(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.Cluster = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) SetCreateTime(v int64) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.CreateTime = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) SetCreatorId(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.CreatorId = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) SetId(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.Id = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) SetName(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.Name = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) SetNickName(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.NickName = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) SetParameter(v map[string]interface{}) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.Parameter = v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) SetParentId(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.ParentId = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) SetRegionId(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.RegionId = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) SetScheduleInfo(v *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListScheduleInfo) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.ScheduleInfo = v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) SetStatus(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.Status = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) SetTag(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.Tag = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) SetTenantId(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.TenantId = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) SetType(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.Type = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) SetVersion(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.Version = &v
	return s
}

type ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListBillingPolicy struct {
	// The billing method of the quota. Valid values:
	//
	// 	- subscription: a subscription quota.
	//
	// 	- payasyougo: a pay-as-you-go quota.
	//
	// example:
	//
	// subscription
	BillingMethod *string `json:"billingMethod,omitempty" xml:"billingMethod,omitempty"`
	// The specifications of the order.
	//
	// example:
	//
	// OdpsStandard
	OdpsSpecCode *string `json:"odpsSpecCode,omitempty" xml:"odpsSpecCode,omitempty"`
	// The ID of the order.
	//
	// example:
	//
	// 45245678
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
}

func (s ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListBillingPolicy) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListBillingPolicy) GoString() string {
	return s.String()
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListBillingPolicy) SetBillingMethod(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListBillingPolicy {
	s.BillingMethod = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListBillingPolicy) SetOdpsSpecCode(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListBillingPolicy {
	s.OdpsSpecCode = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListBillingPolicy) SetOrderId(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListBillingPolicy {
	s.OrderId = &v
	return s
}

type ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListScheduleInfo struct {
	// The quota plan that takes effect based on the scheduling plan.
	//
	// example:
	//
	// planA
	CurrPlan *string `json:"currPlan,omitempty" xml:"currPlan,omitempty"`
	// The time when the current quota plan is scheduled.
	//
	// example:
	//
	// 0800
	CurrTime *string `json:"currTime,omitempty" xml:"currTime,omitempty"`
	// The next quota plan that will take effect based on the scheduling plan.
	//
	// example:
	//
	// planB
	NextPlan *string `json:"nextPlan,omitempty" xml:"nextPlan,omitempty"`
	// The time when the next quota plan is scheduled.
	//
	// example:
	//
	// 1700
	NextTime *string `json:"nextTime,omitempty" xml:"nextTime,omitempty"`
	// If the quota plan that immediately takes effect is different from the current quota plan, this parameter is not empty.
	//
	// example:
	//
	// planC
	OncePlan *string `json:"oncePlan,omitempty" xml:"oncePlan,omitempty"`
	// The time when the quota plan immediately takes effect.
	//
	// example:
	//
	// 1500
	OnceTime *string `json:"onceTime,omitempty" xml:"onceTime,omitempty"`
	// The name of the operator.
	//
	// example:
	//
	// userA
	OperatorName *string `json:"operatorName,omitempty" xml:"operatorName,omitempty"`
}

func (s ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListScheduleInfo) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListScheduleInfo) GoString() string {
	return s.String()
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListScheduleInfo) SetCurrPlan(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListScheduleInfo {
	s.CurrPlan = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListScheduleInfo) SetCurrTime(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListScheduleInfo {
	s.CurrTime = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListScheduleInfo) SetNextPlan(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListScheduleInfo {
	s.NextPlan = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListScheduleInfo) SetNextTime(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListScheduleInfo {
	s.NextTime = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListScheduleInfo) SetOncePlan(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListScheduleInfo {
	s.OncePlan = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListScheduleInfo) SetOnceTime(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListScheduleInfo {
	s.OnceTime = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListScheduleInfo) SetOperatorName(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListScheduleInfo {
	s.OperatorName = &v
	return s
}

type ListQuotasPlansResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListQuotasPlansResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListQuotasPlansResponse) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasPlansResponse) GoString() string {
	return s.String()
}

func (s *ListQuotasPlansResponse) SetHeaders(v map[string]*string) *ListQuotasPlansResponse {
	s.Headers = v
	return s
}

func (s *ListQuotasPlansResponse) SetStatusCode(v int32) *ListQuotasPlansResponse {
	s.StatusCode = &v
	return s
}

func (s *ListQuotasPlansResponse) SetBody(v *ListQuotasPlansResponseBody) *ListQuotasPlansResponse {
	s.Body = v
	return s
}

type ListResourcesRequest struct {
	// Specifies the marker after which the returned list begins.
	//
	// example:
	//
	// cHlvZHBzX3VkZl8xMDExNV8xNDU3NDI4NDkzKg==
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
	// The maximum number of entries to return on each page.
	//
	// example:
	//
	// 10
	MaxItem *int32 `json:"maxItem,omitempty" xml:"maxItem,omitempty"`
	// The name of the resource.
	//
	// example:
	//
	// res
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The name of the schema.
	//
	// example:
	//
	// default
	SchemaName *string `json:"schemaName,omitempty" xml:"schemaName,omitempty"`
}

func (s ListResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListResourcesRequest) SetMarker(v string) *ListResourcesRequest {
	s.Marker = &v
	return s
}

func (s *ListResourcesRequest) SetMaxItem(v int32) *ListResourcesRequest {
	s.MaxItem = &v
	return s
}

func (s *ListResourcesRequest) SetName(v string) *ListResourcesRequest {
	s.Name = &v
	return s
}

func (s *ListResourcesRequest) SetSchemaName(v string) *ListResourcesRequest {
	s.SchemaName = &v
	return s
}

type ListResourcesResponseBody struct {
	// The returned data.
	Data *ListResourcesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 0bc3b4ae16685836687916212e7850
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourcesResponseBody) SetData(v *ListResourcesResponseBodyData) *ListResourcesResponseBody {
	s.Data = v
	return s
}

func (s *ListResourcesResponseBody) SetRequestId(v string) *ListResourcesResponseBody {
	s.RequestId = &v
	return s
}

type ListResourcesResponseBodyData struct {
	// Indicates the marker after which the returned list begins.
	//
	// example:
	//
	// ZmN0X21vbnRoX3Rhb2Jhb19pbmRleCE=
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
	// The maximum number of entries returned per page.
	//
	// example:
	//
	// 10
	MaxItem *int32 `json:"maxItem,omitempty" xml:"maxItem,omitempty"`
	// The list of resources.
	Resources []*ListResourcesResponseBodyDataResources `json:"resources,omitempty" xml:"resources,omitempty" type:"Repeated"`
}

func (s ListResourcesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListResourcesResponseBodyData) SetMarker(v string) *ListResourcesResponseBodyData {
	s.Marker = &v
	return s
}

func (s *ListResourcesResponseBodyData) SetMaxItem(v int32) *ListResourcesResponseBodyData {
	s.MaxItem = &v
	return s
}

func (s *ListResourcesResponseBodyData) SetResources(v []*ListResourcesResponseBodyDataResources) *ListResourcesResponseBodyData {
	s.Resources = v
	return s
}

type ListResourcesResponseBodyDataResources struct {
	// The remarks.
	//
	// example:
	//
	// file
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	// The Base64-encoded 128-bit MD5 hash value of the HTTP request body.
	//
	// example:
	//
	// MACiECZtnLiNkNS1v5****=1
	ContentMD5 *string `json:"contentMD5,omitempty" xml:"contentMD5,omitempty"`
	// The time when the resource was created.
	//
	// example:
	//
	// 2022-01-29T03:34:09Z
	CreationTime *int64 `json:"creationTime,omitempty" xml:"creationTime,omitempty"`
	// The display name.
	//
	// example:
	//
	// res_1
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// The time when the resource was modified.
	//
	// example:
	//
	// 2023-04-18T06:15:05Z
	LastModifiedTime *int64 `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
	// The user who updated the resource.
	//
	// example:
	//
	// ALIYUN$xxx@test.aliyunid.com
	LastUpdator *string `json:"lastUpdator,omitempty" xml:"lastUpdator,omitempty"`
	// The name of the resource.
	//
	// example:
	//
	// res_1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The owner of the resource.
	//
	// example:
	//
	// 1265860483008101
	Owner *string `json:"owner,omitempty" xml:"owner,omitempty"`
	// The schema to which the resource belongs.
	//
	// example:
	//
	// schemaA
	Schema *string `json:"schema,omitempty" xml:"schema,omitempty"`
	// The size of the resource.
	//
	// example:
	//
	// 10
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// The resource type.
	//
	// Valid values:
	//
	// 	- file
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- py
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- jar
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- volumefile
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- table
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// file
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListResourcesResponseBodyDataResources) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesResponseBodyDataResources) GoString() string {
	return s.String()
}

func (s *ListResourcesResponseBodyDataResources) SetComment(v string) *ListResourcesResponseBodyDataResources {
	s.Comment = &v
	return s
}

func (s *ListResourcesResponseBodyDataResources) SetContentMD5(v string) *ListResourcesResponseBodyDataResources {
	s.ContentMD5 = &v
	return s
}

func (s *ListResourcesResponseBodyDataResources) SetCreationTime(v int64) *ListResourcesResponseBodyDataResources {
	s.CreationTime = &v
	return s
}

func (s *ListResourcesResponseBodyDataResources) SetDisplayName(v string) *ListResourcesResponseBodyDataResources {
	s.DisplayName = &v
	return s
}

func (s *ListResourcesResponseBodyDataResources) SetLastModifiedTime(v int64) *ListResourcesResponseBodyDataResources {
	s.LastModifiedTime = &v
	return s
}

func (s *ListResourcesResponseBodyDataResources) SetLastUpdator(v string) *ListResourcesResponseBodyDataResources {
	s.LastUpdator = &v
	return s
}

func (s *ListResourcesResponseBodyDataResources) SetName(v string) *ListResourcesResponseBodyDataResources {
	s.Name = &v
	return s
}

func (s *ListResourcesResponseBodyDataResources) SetOwner(v string) *ListResourcesResponseBodyDataResources {
	s.Owner = &v
	return s
}

func (s *ListResourcesResponseBodyDataResources) SetSchema(v string) *ListResourcesResponseBodyDataResources {
	s.Schema = &v
	return s
}

func (s *ListResourcesResponseBodyDataResources) SetSize(v int64) *ListResourcesResponseBodyDataResources {
	s.Size = &v
	return s
}

func (s *ListResourcesResponseBodyDataResources) SetType(v string) *ListResourcesResponseBodyDataResources {
	s.Type = &v
	return s
}

type ListResourcesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListResourcesResponse) SetHeaders(v map[string]*string) *ListResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListResourcesResponse) SetStatusCode(v int32) *ListResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourcesResponse) SetBody(v *ListResourcesResponseBody) *ListResourcesResponse {
	s.Body = v
	return s
}

type ListRolesResponseBody struct {
	// The returned data.
	Data *ListRolesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 0a06dfe716686526652451361e80ae
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListRolesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRolesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRolesResponseBody) SetData(v *ListRolesResponseBodyData) *ListRolesResponseBody {
	s.Data = v
	return s
}

func (s *ListRolesResponseBody) SetRequestId(v string) *ListRolesResponseBody {
	s.RequestId = &v
	return s
}

type ListRolesResponseBodyData struct {
	// The MaxCompute project-level roles.
	Roles []*ListRolesResponseBodyDataRoles `json:"roles,omitempty" xml:"roles,omitempty" type:"Repeated"`
}

func (s ListRolesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListRolesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListRolesResponseBodyData) SetRoles(v []*ListRolesResponseBodyDataRoles) *ListRolesResponseBodyData {
	s.Roles = v
	return s
}

type ListRolesResponseBodyDataRoles struct {
	// The ACL-based permissions that are granted to the role.
	Acl *ListRolesResponseBodyDataRolesAcl `json:"acl,omitempty" xml:"acl,omitempty" type:"Struct"`
	// The name of the role.
	//
	// example:
	//
	// roleA
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The policy that is attached to the role.
	//
	// example:
	//
	// {
	//
	//       "Statement": [
	//
	//             {
	//
	//                   "Action": [
	//
	//                         "odps:*"
	//
	//                   ],
	//
	//                   "Effect": "Allow",
	//
	//                   "Resource": [
	//
	//                         "acs:odps:*:projects/{projectname}/authorization/packages"
	//
	//                   ]
	//
	//             }
	//
	//       ],
	//
	//       "Version": "1"
	//
	// }
	Policy *string `json:"policy,omitempty" xml:"policy,omitempty"`
	// The type of the role.
	//
	// example:
	//
	// admin
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListRolesResponseBodyDataRoles) String() string {
	return tea.Prettify(s)
}

func (s ListRolesResponseBodyDataRoles) GoString() string {
	return s.String()
}

func (s *ListRolesResponseBodyDataRoles) SetAcl(v *ListRolesResponseBodyDataRolesAcl) *ListRolesResponseBodyDataRoles {
	s.Acl = v
	return s
}

func (s *ListRolesResponseBodyDataRoles) SetName(v string) *ListRolesResponseBodyDataRoles {
	s.Name = &v
	return s
}

func (s *ListRolesResponseBodyDataRoles) SetPolicy(v string) *ListRolesResponseBodyDataRoles {
	s.Policy = &v
	return s
}

func (s *ListRolesResponseBodyDataRoles) SetType(v string) *ListRolesResponseBodyDataRoles {
	s.Type = &v
	return s
}

type ListRolesResponseBodyDataRolesAcl struct {
	// The function.
	Function []*ListRolesResponseBodyDataRolesAclFunction `json:"function,omitempty" xml:"function,omitempty" type:"Repeated"`
	// The instance.
	Instance []*ListRolesResponseBodyDataRolesAclInstance `json:"instance,omitempty" xml:"instance,omitempty" type:"Repeated"`
	// The package.
	Package []*ListRolesResponseBodyDataRolesAclPackage `json:"package,omitempty" xml:"package,omitempty" type:"Repeated"`
	// The project.
	Project []*ListRolesResponseBodyDataRolesAclProject `json:"project,omitempty" xml:"project,omitempty" type:"Repeated"`
	// The resource.
	Resource []*ListRolesResponseBodyDataRolesAclResource `json:"resource,omitempty" xml:"resource,omitempty" type:"Repeated"`
	// The table.
	Table []*ListRolesResponseBodyDataRolesAclTable `json:"table,omitempty" xml:"table,omitempty" type:"Repeated"`
}

func (s ListRolesResponseBodyDataRolesAcl) String() string {
	return tea.Prettify(s)
}

func (s ListRolesResponseBodyDataRolesAcl) GoString() string {
	return s.String()
}

func (s *ListRolesResponseBodyDataRolesAcl) SetFunction(v []*ListRolesResponseBodyDataRolesAclFunction) *ListRolesResponseBodyDataRolesAcl {
	s.Function = v
	return s
}

func (s *ListRolesResponseBodyDataRolesAcl) SetInstance(v []*ListRolesResponseBodyDataRolesAclInstance) *ListRolesResponseBodyDataRolesAcl {
	s.Instance = v
	return s
}

func (s *ListRolesResponseBodyDataRolesAcl) SetPackage(v []*ListRolesResponseBodyDataRolesAclPackage) *ListRolesResponseBodyDataRolesAcl {
	s.Package = v
	return s
}

func (s *ListRolesResponseBodyDataRolesAcl) SetProject(v []*ListRolesResponseBodyDataRolesAclProject) *ListRolesResponseBodyDataRolesAcl {
	s.Project = v
	return s
}

func (s *ListRolesResponseBodyDataRolesAcl) SetResource(v []*ListRolesResponseBodyDataRolesAclResource) *ListRolesResponseBodyDataRolesAcl {
	s.Resource = v
	return s
}

func (s *ListRolesResponseBodyDataRolesAcl) SetTable(v []*ListRolesResponseBodyDataRolesAclTable) *ListRolesResponseBodyDataRolesAcl {
	s.Table = v
	return s
}

type ListRolesResponseBodyDataRolesAclFunction struct {
	// The operations that were performed on the function.
	Actions []*string `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	// The name of the function.
	//
	// example:
	//
	// functionA
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListRolesResponseBodyDataRolesAclFunction) String() string {
	return tea.Prettify(s)
}

func (s ListRolesResponseBodyDataRolesAclFunction) GoString() string {
	return s.String()
}

func (s *ListRolesResponseBodyDataRolesAclFunction) SetActions(v []*string) *ListRolesResponseBodyDataRolesAclFunction {
	s.Actions = v
	return s
}

func (s *ListRolesResponseBodyDataRolesAclFunction) SetName(v string) *ListRolesResponseBodyDataRolesAclFunction {
	s.Name = &v
	return s
}

type ListRolesResponseBodyDataRolesAclInstance struct {
	// The operations that were performed on the instance.
	Actions []*string `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	// The name of the instance.
	//
	// example:
	//
	// instanceA
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListRolesResponseBodyDataRolesAclInstance) String() string {
	return tea.Prettify(s)
}

func (s ListRolesResponseBodyDataRolesAclInstance) GoString() string {
	return s.String()
}

func (s *ListRolesResponseBodyDataRolesAclInstance) SetActions(v []*string) *ListRolesResponseBodyDataRolesAclInstance {
	s.Actions = v
	return s
}

func (s *ListRolesResponseBodyDataRolesAclInstance) SetName(v string) *ListRolesResponseBodyDataRolesAclInstance {
	s.Name = &v
	return s
}

type ListRolesResponseBodyDataRolesAclPackage struct {
	// The operations that were performed on the package.
	Actions []*string `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	// The name of the package.
	//
	// example:
	//
	// packageA
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListRolesResponseBodyDataRolesAclPackage) String() string {
	return tea.Prettify(s)
}

func (s ListRolesResponseBodyDataRolesAclPackage) GoString() string {
	return s.String()
}

func (s *ListRolesResponseBodyDataRolesAclPackage) SetActions(v []*string) *ListRolesResponseBodyDataRolesAclPackage {
	s.Actions = v
	return s
}

func (s *ListRolesResponseBodyDataRolesAclPackage) SetName(v string) *ListRolesResponseBodyDataRolesAclPackage {
	s.Name = &v
	return s
}

type ListRolesResponseBodyDataRolesAclProject struct {
	// The operations that were performed on the project.
	Actions []*string `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	// The name of the MaxCompute project.
	//
	// example:
	//
	// projectA
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListRolesResponseBodyDataRolesAclProject) String() string {
	return tea.Prettify(s)
}

func (s ListRolesResponseBodyDataRolesAclProject) GoString() string {
	return s.String()
}

func (s *ListRolesResponseBodyDataRolesAclProject) SetActions(v []*string) *ListRolesResponseBodyDataRolesAclProject {
	s.Actions = v
	return s
}

func (s *ListRolesResponseBodyDataRolesAclProject) SetName(v string) *ListRolesResponseBodyDataRolesAclProject {
	s.Name = &v
	return s
}

type ListRolesResponseBodyDataRolesAclResource struct {
	// The operations that were performed on the resource.
	Actions []*string `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	// The name of the resource.
	//
	// example:
	//
	// resourceA
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListRolesResponseBodyDataRolesAclResource) String() string {
	return tea.Prettify(s)
}

func (s ListRolesResponseBodyDataRolesAclResource) GoString() string {
	return s.String()
}

func (s *ListRolesResponseBodyDataRolesAclResource) SetActions(v []*string) *ListRolesResponseBodyDataRolesAclResource {
	s.Actions = v
	return s
}

func (s *ListRolesResponseBodyDataRolesAclResource) SetName(v string) *ListRolesResponseBodyDataRolesAclResource {
	s.Name = &v
	return s
}

type ListRolesResponseBodyDataRolesAclTable struct {
	// The operations that were performed on the table.
	Actions []*string `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	// The name of the table.
	//
	// example:
	//
	// tableA
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListRolesResponseBodyDataRolesAclTable) String() string {
	return tea.Prettify(s)
}

func (s ListRolesResponseBodyDataRolesAclTable) GoString() string {
	return s.String()
}

func (s *ListRolesResponseBodyDataRolesAclTable) SetActions(v []*string) *ListRolesResponseBodyDataRolesAclTable {
	s.Actions = v
	return s
}

func (s *ListRolesResponseBodyDataRolesAclTable) SetName(v string) *ListRolesResponseBodyDataRolesAclTable {
	s.Name = &v
	return s
}

type ListRolesResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRolesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRolesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRolesResponse) GoString() string {
	return s.String()
}

func (s *ListRolesResponse) SetHeaders(v map[string]*string) *ListRolesResponse {
	s.Headers = v
	return s
}

func (s *ListRolesResponse) SetStatusCode(v int32) *ListRolesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRolesResponse) SetBody(v *ListRolesResponseBody) *ListRolesResponse {
	s.Body = v
	return s
}

type ListStoragePartitionsInfoRequest struct {
	// Specifies whether to sort data in ascending order.
	//
	// example:
	//
	// false
	AscOrder *bool `json:"ascOrder,omitempty" xml:"ascOrder,omitempty"`
	// The date on which the statistics are collected, in days. Set this parameter to a value in the YYYYMMdd format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20241205
	Date *string `json:"date,omitempty" xml:"date,omitempty"`
	// The sorting column.
	//
	// example:
	//
	// totalFrequency
	OrderColumn *string `json:"orderColumn,omitempty" xml:"orderColumn,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The name of the partition that you want to use for fuzzy match.
	//
	// example:
	//
	// 20241201
	PartitionPrefix *string `json:"partitionPrefix,omitempty" xml:"partitionPrefix,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The name of the schema.
	//
	// example:
	//
	// schema
	Schema *string `json:"schema,omitempty" xml:"schema,omitempty"`
	// The ID of the tenant. You can log on to the MaxCompute console, and choose **Tenants*	- > **Tenant Property*	- from the left-side navigation pane to view the tenant ID.
	//
	// example:
	//
	// 40713753659****
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The storage types.
	Types []*string `json:"types,omitempty" xml:"types,omitempty" type:"Repeated"`
}

func (s ListStoragePartitionsInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s ListStoragePartitionsInfoRequest) GoString() string {
	return s.String()
}

func (s *ListStoragePartitionsInfoRequest) SetAscOrder(v bool) *ListStoragePartitionsInfoRequest {
	s.AscOrder = &v
	return s
}

func (s *ListStoragePartitionsInfoRequest) SetDate(v string) *ListStoragePartitionsInfoRequest {
	s.Date = &v
	return s
}

func (s *ListStoragePartitionsInfoRequest) SetOrderColumn(v string) *ListStoragePartitionsInfoRequest {
	s.OrderColumn = &v
	return s
}

func (s *ListStoragePartitionsInfoRequest) SetPageNumber(v int64) *ListStoragePartitionsInfoRequest {
	s.PageNumber = &v
	return s
}

func (s *ListStoragePartitionsInfoRequest) SetPageSize(v int64) *ListStoragePartitionsInfoRequest {
	s.PageSize = &v
	return s
}

func (s *ListStoragePartitionsInfoRequest) SetPartitionPrefix(v string) *ListStoragePartitionsInfoRequest {
	s.PartitionPrefix = &v
	return s
}

func (s *ListStoragePartitionsInfoRequest) SetRegion(v string) *ListStoragePartitionsInfoRequest {
	s.Region = &v
	return s
}

func (s *ListStoragePartitionsInfoRequest) SetSchema(v string) *ListStoragePartitionsInfoRequest {
	s.Schema = &v
	return s
}

func (s *ListStoragePartitionsInfoRequest) SetTenantId(v string) *ListStoragePartitionsInfoRequest {
	s.TenantId = &v
	return s
}

func (s *ListStoragePartitionsInfoRequest) SetTypes(v []*string) *ListStoragePartitionsInfoRequest {
	s.Types = v
	return s
}

type ListStoragePartitionsInfoShrinkRequest struct {
	// Specifies whether to sort data in ascending order.
	//
	// example:
	//
	// false
	AscOrder *bool `json:"ascOrder,omitempty" xml:"ascOrder,omitempty"`
	// The date on which the statistics are collected, in days. Set this parameter to a value in the YYYYMMdd format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20241205
	Date *string `json:"date,omitempty" xml:"date,omitempty"`
	// The sorting column.
	//
	// example:
	//
	// totalFrequency
	OrderColumn *string `json:"orderColumn,omitempty" xml:"orderColumn,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The name of the partition that you want to use for fuzzy match.
	//
	// example:
	//
	// 20241201
	PartitionPrefix *string `json:"partitionPrefix,omitempty" xml:"partitionPrefix,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The name of the schema.
	//
	// example:
	//
	// schema
	Schema *string `json:"schema,omitempty" xml:"schema,omitempty"`
	// The ID of the tenant. You can log on to the MaxCompute console, and choose **Tenants*	- > **Tenant Property*	- from the left-side navigation pane to view the tenant ID.
	//
	// example:
	//
	// 40713753659****
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The storage types.
	TypesShrink *string `json:"types,omitempty" xml:"types,omitempty"`
}

func (s ListStoragePartitionsInfoShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListStoragePartitionsInfoShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListStoragePartitionsInfoShrinkRequest) SetAscOrder(v bool) *ListStoragePartitionsInfoShrinkRequest {
	s.AscOrder = &v
	return s
}

func (s *ListStoragePartitionsInfoShrinkRequest) SetDate(v string) *ListStoragePartitionsInfoShrinkRequest {
	s.Date = &v
	return s
}

func (s *ListStoragePartitionsInfoShrinkRequest) SetOrderColumn(v string) *ListStoragePartitionsInfoShrinkRequest {
	s.OrderColumn = &v
	return s
}

func (s *ListStoragePartitionsInfoShrinkRequest) SetPageNumber(v int64) *ListStoragePartitionsInfoShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListStoragePartitionsInfoShrinkRequest) SetPageSize(v int64) *ListStoragePartitionsInfoShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListStoragePartitionsInfoShrinkRequest) SetPartitionPrefix(v string) *ListStoragePartitionsInfoShrinkRequest {
	s.PartitionPrefix = &v
	return s
}

func (s *ListStoragePartitionsInfoShrinkRequest) SetRegion(v string) *ListStoragePartitionsInfoShrinkRequest {
	s.Region = &v
	return s
}

func (s *ListStoragePartitionsInfoShrinkRequest) SetSchema(v string) *ListStoragePartitionsInfoShrinkRequest {
	s.Schema = &v
	return s
}

func (s *ListStoragePartitionsInfoShrinkRequest) SetTenantId(v string) *ListStoragePartitionsInfoShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *ListStoragePartitionsInfoShrinkRequest) SetTypesShrink(v string) *ListStoragePartitionsInfoShrinkRequest {
	s.TypesShrink = &v
	return s
}

type ListStoragePartitionsInfoResponseBody struct {
	// The data returned.
	Data *ListStoragePartitionsInfoResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// OBJECT_NOT_EXIST
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// This object does not exist.
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// The HTTP status code.
	//
	// 	- 1xx: informational response. The request is received and is being processed.
	//
	// 	- 2xx: success. The request is successfully received, understood, and accepted by the server.
	//
	// 	- 3xx: redirection. The request is redirected, and further actions are required to complete the request.
	//
	// 	- 4xx: client error. The request contains invalid request parameters and syntaxes, or specific request conditions cannot be met.
	//
	// 	- 5xx: server error. The server cannot meet requirements due to other reasons.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0be3e0bd16661643917136451ebf55
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListStoragePartitionsInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListStoragePartitionsInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ListStoragePartitionsInfoResponseBody) SetData(v *ListStoragePartitionsInfoResponseBodyData) *ListStoragePartitionsInfoResponseBody {
	s.Data = v
	return s
}

func (s *ListStoragePartitionsInfoResponseBody) SetErrorCode(v string) *ListStoragePartitionsInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListStoragePartitionsInfoResponseBody) SetErrorMsg(v string) *ListStoragePartitionsInfoResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListStoragePartitionsInfoResponseBody) SetHttpCode(v int32) *ListStoragePartitionsInfoResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListStoragePartitionsInfoResponseBody) SetRequestId(v string) *ListStoragePartitionsInfoResponseBody {
	s.RequestId = &v
	return s
}

type ListStoragePartitionsInfoResponseBodyData struct {
	// The date on which the statistics are collected.
	//
	// example:
	//
	// 20241205
	Date *string `json:"date,omitempty" xml:"date,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The partition storage information.
	StoragePartitionInfoList []*ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList `json:"storagePartitionInfoList,omitempty" xml:"storagePartitionInfoList,omitempty" type:"Repeated"`
	// The total number of returned entries.
	//
	// example:
	//
	// 57
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListStoragePartitionsInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListStoragePartitionsInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListStoragePartitionsInfoResponseBodyData) SetDate(v string) *ListStoragePartitionsInfoResponseBodyData {
	s.Date = &v
	return s
}

func (s *ListStoragePartitionsInfoResponseBodyData) SetPageNumber(v int64) *ListStoragePartitionsInfoResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListStoragePartitionsInfoResponseBodyData) SetPageSize(v int64) *ListStoragePartitionsInfoResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListStoragePartitionsInfoResponseBodyData) SetStoragePartitionInfoList(v []*ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) *ListStoragePartitionsInfoResponseBodyData {
	s.StoragePartitionInfoList = v
	return s
}

func (s *ListStoragePartitionsInfoResponseBodyData) SetTotalCount(v int64) *ListStoragePartitionsInfoResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList struct {
	// The number of files.
	//
	// example:
	//
	// 2
	FileCount *int64 `json:"fileCount,omitempty" xml:"fileCount,omitempty"`
	// The storage size.
	//
	// example:
	//
	// 1
	FileSize *float64 `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	// The unit of the storage size.
	//
	// example:
	//
	// GB
	FileSizeUnit *string `json:"fileSizeUnit,omitempty" xml:"fileSizeUnit,omitempty"`
	// Indicates whether the table is a partitioned table. This operation returns the partition information. You do not need to take note of this parameter.
	//
	// example:
	//
	// false
	IsPartitioned *bool `json:"isPartitioned,omitempty" xml:"isPartitioned,omitempty"`
	// The time when the partition data was last accessed.
	//
	// >  The data collection method is upgraded from July 2023. If the data is not accessed after the upgrade or is accessed by using ALGO jobs or the direct read method of Hologres, the last access time cannot be collected.
	//
	// example:
	//
	// 1694589365
	LastAccessTime *int64 `json:"lastAccessTime,omitempty" xml:"lastAccessTime,omitempty"`
	// The partition name.
	//
	// example:
	//
	// ds=20241201
	Partition *string `json:"partition,omitempty" xml:"partition,omitempty"`
	// The project name.
	//
	// example:
	//
	// odps_project
	ProjectName *string `json:"projectName,omitempty" xml:"projectName,omitempty"`
	// The change rate of the total storage usage compared with that of the recent {$recentDays} days. No value is returned.
	//
	// example:
	//
	// 1%
	Rate *float64 `json:"rate,omitempty" xml:"rate,omitempty"`
	// The schema name.
	//
	// example:
	//
	// schema
	SchemaName *string `json:"schemaName,omitempty" xml:"schemaName,omitempty"`
	// The storage type.
	//
	// 	- standard
	//
	// 	- lowfrequency
	//
	// 	- longterm
	//
	// example:
	//
	// standard
	StorageType *string `json:"storageType,omitempty" xml:"storageType,omitempty"`
	// The table name.
	//
	// example:
	//
	// bank_data
	TableName *string `json:"tableName,omitempty" xml:"tableName,omitempty"`
	// The access frequency.
	//
	// >
	//
	// 	- Access behaviors include:
	//
	// 	- The table is used as the input table of an SQL task.
	//
	// 	- The table is downloaded by Tunnel.
	//
	// 	- The table is read by calling the Storage API. The partition granularity of the partitioned table is not available. Each time an access operation is performed, the access frequency is incremented by 1.
	//
	// 	- The data collection method is upgraded from July 2023. If the data is not accessed after the upgrade or is accessed by using ALGO jobs or the direct read method of Hologres, the access frequency cannot be collected.
	//
	// example:
	//
	// 10
	TotalFrequency *int64 `json:"totalFrequency,omitempty" xml:"totalFrequency,omitempty"`
	// The total amount of accessed data.
	//
	// >  The amount of data that is read by all access behaviors.
	//
	// example:
	//
	// 1
	TotalInputAmount *float64 `json:"totalInputAmount,omitempty" xml:"totalInputAmount,omitempty"`
	// The unit of the total amount of accessed data.
	//
	// example:
	//
	// GB
	TotalInputAmountUnit *string `json:"totalInputAmountUnit,omitempty" xml:"totalInputAmountUnit,omitempty"`
	// The type.
	//
	// example:
	//
	// PARTITION
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) String() string {
	return tea.Prettify(s)
}

func (s ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) GoString() string {
	return s.String()
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) SetFileCount(v int64) *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList {
	s.FileCount = &v
	return s
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) SetFileSize(v float64) *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList {
	s.FileSize = &v
	return s
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) SetFileSizeUnit(v string) *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList {
	s.FileSizeUnit = &v
	return s
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) SetIsPartitioned(v bool) *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList {
	s.IsPartitioned = &v
	return s
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) SetLastAccessTime(v int64) *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList {
	s.LastAccessTime = &v
	return s
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) SetPartition(v string) *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList {
	s.Partition = &v
	return s
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) SetProjectName(v string) *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList {
	s.ProjectName = &v
	return s
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) SetRate(v float64) *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList {
	s.Rate = &v
	return s
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) SetSchemaName(v string) *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList {
	s.SchemaName = &v
	return s
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) SetStorageType(v string) *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList {
	s.StorageType = &v
	return s
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) SetTableName(v string) *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList {
	s.TableName = &v
	return s
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) SetTotalFrequency(v int64) *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList {
	s.TotalFrequency = &v
	return s
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) SetTotalInputAmount(v float64) *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList {
	s.TotalInputAmount = &v
	return s
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) SetTotalInputAmountUnit(v string) *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList {
	s.TotalInputAmountUnit = &v
	return s
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) SetType(v string) *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList {
	s.Type = &v
	return s
}

type ListStoragePartitionsInfoResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListStoragePartitionsInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListStoragePartitionsInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s ListStoragePartitionsInfoResponse) GoString() string {
	return s.String()
}

func (s *ListStoragePartitionsInfoResponse) SetHeaders(v map[string]*string) *ListStoragePartitionsInfoResponse {
	s.Headers = v
	return s
}

func (s *ListStoragePartitionsInfoResponse) SetStatusCode(v int32) *ListStoragePartitionsInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ListStoragePartitionsInfoResponse) SetBody(v *ListStoragePartitionsInfoResponseBody) *ListStoragePartitionsInfoResponse {
	s.Body = v
	return s
}

type ListStorageTablesInfoRequest struct {
	// Specifies whether to sort data in ascending order.
	//
	// example:
	//
	// false
	AscOrder *bool `json:"ascOrder,omitempty" xml:"ascOrder,omitempty"`
	// The date on which the statistics are collected, in days. Set this parameter to a value in the `YYYYMMdd` format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20241205
	Date *string `json:"date,omitempty" xml:"date,omitempty"`
	// The sorting column.
	//
	// example:
	//
	// totalFrequency
	OrderColumn *string `json:"orderColumn,omitempty" xml:"orderColumn,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The number of recent days for comparison.
	//
	// example:
	//
	// 1
	RecentDays *int32 `json:"recentDays,omitempty" xml:"recentDays,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The name of the schema.
	//
	// example:
	//
	// schema
	Schema *string `json:"schema,omitempty" xml:"schema,omitempty"`
	// The name of the table that you want to use for fuzzy match.
	//
	// example:
	//
	// bank
	TablePrefix *string `json:"tablePrefix,omitempty" xml:"tablePrefix,omitempty"`
	// The ID of the tenant. You can log on to the MaxCompute console, and choose **Tenants*	- > **Tenant Property*	- from the left-side navigation pane to view the tenant ID.
	//
	// example:
	//
	// 28074710977****
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The storage types.
	Types []*string `json:"types,omitempty" xml:"types,omitempty" type:"Repeated"`
}

func (s ListStorageTablesInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s ListStorageTablesInfoRequest) GoString() string {
	return s.String()
}

func (s *ListStorageTablesInfoRequest) SetAscOrder(v bool) *ListStorageTablesInfoRequest {
	s.AscOrder = &v
	return s
}

func (s *ListStorageTablesInfoRequest) SetDate(v string) *ListStorageTablesInfoRequest {
	s.Date = &v
	return s
}

func (s *ListStorageTablesInfoRequest) SetOrderColumn(v string) *ListStorageTablesInfoRequest {
	s.OrderColumn = &v
	return s
}

func (s *ListStorageTablesInfoRequest) SetPageNumber(v int64) *ListStorageTablesInfoRequest {
	s.PageNumber = &v
	return s
}

func (s *ListStorageTablesInfoRequest) SetPageSize(v int64) *ListStorageTablesInfoRequest {
	s.PageSize = &v
	return s
}

func (s *ListStorageTablesInfoRequest) SetRecentDays(v int32) *ListStorageTablesInfoRequest {
	s.RecentDays = &v
	return s
}

func (s *ListStorageTablesInfoRequest) SetRegion(v string) *ListStorageTablesInfoRequest {
	s.Region = &v
	return s
}

func (s *ListStorageTablesInfoRequest) SetSchema(v string) *ListStorageTablesInfoRequest {
	s.Schema = &v
	return s
}

func (s *ListStorageTablesInfoRequest) SetTablePrefix(v string) *ListStorageTablesInfoRequest {
	s.TablePrefix = &v
	return s
}

func (s *ListStorageTablesInfoRequest) SetTenantId(v string) *ListStorageTablesInfoRequest {
	s.TenantId = &v
	return s
}

func (s *ListStorageTablesInfoRequest) SetTypes(v []*string) *ListStorageTablesInfoRequest {
	s.Types = v
	return s
}

type ListStorageTablesInfoShrinkRequest struct {
	// Specifies whether to sort data in ascending order.
	//
	// example:
	//
	// false
	AscOrder *bool `json:"ascOrder,omitempty" xml:"ascOrder,omitempty"`
	// The date on which the statistics are collected, in days. Set this parameter to a value in the `YYYYMMdd` format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20241205
	Date *string `json:"date,omitempty" xml:"date,omitempty"`
	// The sorting column.
	//
	// example:
	//
	// totalFrequency
	OrderColumn *string `json:"orderColumn,omitempty" xml:"orderColumn,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The number of recent days for comparison.
	//
	// example:
	//
	// 1
	RecentDays *int32 `json:"recentDays,omitempty" xml:"recentDays,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The name of the schema.
	//
	// example:
	//
	// schema
	Schema *string `json:"schema,omitempty" xml:"schema,omitempty"`
	// The name of the table that you want to use for fuzzy match.
	//
	// example:
	//
	// bank
	TablePrefix *string `json:"tablePrefix,omitempty" xml:"tablePrefix,omitempty"`
	// The ID of the tenant. You can log on to the MaxCompute console, and choose **Tenants*	- > **Tenant Property*	- from the left-side navigation pane to view the tenant ID.
	//
	// example:
	//
	// 28074710977****
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The storage types.
	TypesShrink *string `json:"types,omitempty" xml:"types,omitempty"`
}

func (s ListStorageTablesInfoShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListStorageTablesInfoShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListStorageTablesInfoShrinkRequest) SetAscOrder(v bool) *ListStorageTablesInfoShrinkRequest {
	s.AscOrder = &v
	return s
}

func (s *ListStorageTablesInfoShrinkRequest) SetDate(v string) *ListStorageTablesInfoShrinkRequest {
	s.Date = &v
	return s
}

func (s *ListStorageTablesInfoShrinkRequest) SetOrderColumn(v string) *ListStorageTablesInfoShrinkRequest {
	s.OrderColumn = &v
	return s
}

func (s *ListStorageTablesInfoShrinkRequest) SetPageNumber(v int64) *ListStorageTablesInfoShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListStorageTablesInfoShrinkRequest) SetPageSize(v int64) *ListStorageTablesInfoShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListStorageTablesInfoShrinkRequest) SetRecentDays(v int32) *ListStorageTablesInfoShrinkRequest {
	s.RecentDays = &v
	return s
}

func (s *ListStorageTablesInfoShrinkRequest) SetRegion(v string) *ListStorageTablesInfoShrinkRequest {
	s.Region = &v
	return s
}

func (s *ListStorageTablesInfoShrinkRequest) SetSchema(v string) *ListStorageTablesInfoShrinkRequest {
	s.Schema = &v
	return s
}

func (s *ListStorageTablesInfoShrinkRequest) SetTablePrefix(v string) *ListStorageTablesInfoShrinkRequest {
	s.TablePrefix = &v
	return s
}

func (s *ListStorageTablesInfoShrinkRequest) SetTenantId(v string) *ListStorageTablesInfoShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *ListStorageTablesInfoShrinkRequest) SetTypesShrink(v string) *ListStorageTablesInfoShrinkRequest {
	s.TypesShrink = &v
	return s
}

type ListStorageTablesInfoResponseBody struct {
	// The data returned.
	Data *ListStorageTablesInfoResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// OBJECT_NOT_EXIST
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// This object does not exist.
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// The HTTP status code.
	//
	// 	- 1xx: informational response. The request is received and is being processed.
	//
	// 	- 2xx: success. The request is successfully received, understood, and accepted by the server.
	//
	// 	- 3xx: redirection. The request is redirected, and further actions are required to complete the request.
	//
	// 	- 4xx: client error. The request contains invalid request parameters and syntaxes, or specific request conditions cannot be met.
	//
	// 	- 5xx: server error. The server cannot meet requirements due to other reasons.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0bc12e6a16679892465424670db3eb
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListStorageTablesInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListStorageTablesInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ListStorageTablesInfoResponseBody) SetData(v *ListStorageTablesInfoResponseBodyData) *ListStorageTablesInfoResponseBody {
	s.Data = v
	return s
}

func (s *ListStorageTablesInfoResponseBody) SetErrorCode(v string) *ListStorageTablesInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListStorageTablesInfoResponseBody) SetErrorMsg(v string) *ListStorageTablesInfoResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListStorageTablesInfoResponseBody) SetHttpCode(v int32) *ListStorageTablesInfoResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListStorageTablesInfoResponseBody) SetRequestId(v string) *ListStorageTablesInfoResponseBody {
	s.RequestId = &v
	return s
}

type ListStorageTablesInfoResponseBodyData struct {
	// The date on which the statistics are collected.
	//
	// example:
	//
	// 20241205
	Date *string `json:"date,omitempty" xml:"date,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The table storage information.
	StorageTableInfoList []*ListStorageTablesInfoResponseBodyDataStorageTableInfoList `json:"storageTableInfoList,omitempty" xml:"storageTableInfoList,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListStorageTablesInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListStorageTablesInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListStorageTablesInfoResponseBodyData) SetDate(v string) *ListStorageTablesInfoResponseBodyData {
	s.Date = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyData) SetPageNumber(v int64) *ListStorageTablesInfoResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyData) SetPageSize(v int64) *ListStorageTablesInfoResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyData) SetStorageTableInfoList(v []*ListStorageTablesInfoResponseBodyDataStorageTableInfoList) *ListStorageTablesInfoResponseBodyData {
	s.StorageTableInfoList = v
	return s
}

func (s *ListStorageTablesInfoResponseBodyData) SetTotalCount(v int64) *ListStorageTablesInfoResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListStorageTablesInfoResponseBodyDataStorageTableInfoList struct {
	// The date on which the statistics are collected. This value is not returned.
	//
	// example:
	//
	// 20241205
	Date *string `json:"date,omitempty" xml:"date,omitempty"`
	// Indicates whether the table is a partitioned table.
	//
	// example:
	//
	// false
	IsPartitioned *bool `json:"isPartitioned,omitempty" xml:"isPartitioned,omitempty"`
	// The time when the table was last accessed. This value is returned when the table is a non-partitioned table.
	//
	// >  The data collection method is upgraded from July 2023. If the data is not accessed after the upgrade or is accessed by using ALGO jobs or the direct read method of Hologres, the last access time cannot be collected.
	//
	// example:
	//
	// 1694589365
	LastAccessTime *int64 `json:"lastAccessTime,omitempty" xml:"lastAccessTime,omitempty"`
	// The storage usage at the long-term storage tier.
	//
	// example:
	//
	// 0
	LongTermStorage *float64 `json:"longTermStorage,omitempty" xml:"longTermStorage,omitempty"`
	// The number of long-term storage files.
	//
	// example:
	//
	// 0
	LongTermStorageFileCount *int64 `json:"longTermStorageFileCount,omitempty" xml:"longTermStorageFileCount,omitempty"`
	// The unit of the storage usage at the long-term storage tier.
	//
	// example:
	//
	// B
	LongTermStorageUnit *string `json:"longTermStorageUnit,omitempty" xml:"longTermStorageUnit,omitempty"`
	// The storage usage at the low-frequency tier.
	//
	// example:
	//
	// 0
	LowFreqStorage *float64 `json:"lowFreqStorage,omitempty" xml:"lowFreqStorage,omitempty"`
	// The number of low-frequency storage files.
	//
	// example:
	//
	// 0
	LowFreqStorageFileCount *int64 `json:"lowFreqStorageFileCount,omitempty" xml:"lowFreqStorageFileCount,omitempty"`
	// The unit of the storage usage at the low-frequency storage tier.
	//
	// example:
	//
	// B
	LowFreqStorageUnit *string `json:"lowFreqStorageUnit,omitempty" xml:"lowFreqStorageUnit,omitempty"`
	// The project name.
	//
	// example:
	//
	// odps_project
	ProjectName *string `json:"projectName,omitempty" xml:"projectName,omitempty"`
	// The change rate of the total storage usage compared with that of the recent {$recentDays} days.
	//
	// example:
	//
	// 0
	Rate *float64 `json:"rate,omitempty" xml:"rate,omitempty"`
	// The schema name.
	//
	// example:
	//
	// schema
	SchemaName *string `json:"schemaName,omitempty" xml:"schemaName,omitempty"`
	// The storage usage at the standard storage tier.
	//
	// example:
	//
	// 600
	StandardStorage *float64 `json:"standardStorage,omitempty" xml:"standardStorage,omitempty"`
	// The number of standard storage files.
	//
	// example:
	//
	// 2
	StandardStorageFileCount *int64 `json:"standardStorageFileCount,omitempty" xml:"standardStorageFileCount,omitempty"`
	// The unit of the storage usage at the standard storage tier.
	//
	// example:
	//
	// KB
	StandardStorageUnit *string `json:"standardStorageUnit,omitempty" xml:"standardStorageUnit,omitempty"`
	// The table storage type.
	//
	// 	- standard
	//
	// 	- lowfrequency
	//
	// 	- longterm
	//
	// 	- unknown: This value is returned when the table is a partitioned table. You can call the ListStoragePartitionsInfo operation to query the storage type of each partition.
	//
	// example:
	//
	// standard
	StorageType *string `json:"storageType,omitempty" xml:"storageType,omitempty"`
	// The table name.
	//
	// example:
	//
	// bank_data
	TableName *string `json:"tableName,omitempty" xml:"tableName,omitempty"`
	// The access frequency.
	//
	// >
	//
	// 	- Access behaviors include:
	//
	// 	- The table is used as the input table of an SQL task.
	//
	// 	- The table is downloaded by Tunnel.
	//
	// 	- The table is read by calling the Storage API. The partition granularity of the partitioned table is not available. Each time an access operation is performed, the access frequency is incremented by 1.
	//
	// 	- The data collection method is upgraded from July 2023. If the data is not accessed after the upgrade or is accessed by using ALGO jobs or the direct read method of Hologres, the access frequency cannot be collected.
	//
	// example:
	//
	// 10
	TotalFrequency *int64 `json:"totalFrequency,omitempty" xml:"totalFrequency,omitempty"`
	// The total amount of accessed data.
	//
	// >  The amount of data that is read by all access behaviors.
	//
	// example:
	//
	// 1
	TotalInputAmount *float64 `json:"totalInputAmount,omitempty" xml:"totalInputAmount,omitempty"`
	// The unit of the total amount of accessed data.
	//
	// example:
	//
	// GB
	TotalInputAmountUnit *string `json:"totalInputAmountUnit,omitempty" xml:"totalInputAmountUnit,omitempty"`
	// The total storage usage. For a partitioned table, this parameter indicates the sum of the storage usage of all partitions. If the storage types of partitions are different, the value is the sum of the storage usage of each storage type.
	//
	// example:
	//
	// 600
	TotalStorage *float64 `json:"totalStorage,omitempty" xml:"totalStorage,omitempty"`
	// The total number of files.
	//
	// example:
	//
	// 2
	TotalStorageFileCount *int64 `json:"totalStorageFileCount,omitempty" xml:"totalStorageFileCount,omitempty"`
	// The unit of storage usage.
	//
	// example:
	//
	// KB
	TotalStorageUnit *string `json:"totalStorageUnit,omitempty" xml:"totalStorageUnit,omitempty"`
}

func (s ListStorageTablesInfoResponseBodyDataStorageTableInfoList) String() string {
	return tea.Prettify(s)
}

func (s ListStorageTablesInfoResponseBodyDataStorageTableInfoList) GoString() string {
	return s.String()
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) SetDate(v string) *ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	s.Date = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) SetIsPartitioned(v bool) *ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	s.IsPartitioned = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) SetLastAccessTime(v int64) *ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	s.LastAccessTime = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) SetLongTermStorage(v float64) *ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	s.LongTermStorage = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) SetLongTermStorageFileCount(v int64) *ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	s.LongTermStorageFileCount = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) SetLongTermStorageUnit(v string) *ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	s.LongTermStorageUnit = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) SetLowFreqStorage(v float64) *ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	s.LowFreqStorage = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) SetLowFreqStorageFileCount(v int64) *ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	s.LowFreqStorageFileCount = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) SetLowFreqStorageUnit(v string) *ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	s.LowFreqStorageUnit = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) SetProjectName(v string) *ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	s.ProjectName = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) SetRate(v float64) *ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	s.Rate = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) SetSchemaName(v string) *ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	s.SchemaName = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) SetStandardStorage(v float64) *ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	s.StandardStorage = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) SetStandardStorageFileCount(v int64) *ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	s.StandardStorageFileCount = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) SetStandardStorageUnit(v string) *ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	s.StandardStorageUnit = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) SetStorageType(v string) *ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	s.StorageType = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) SetTableName(v string) *ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	s.TableName = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) SetTotalFrequency(v int64) *ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	s.TotalFrequency = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) SetTotalInputAmount(v float64) *ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	s.TotalInputAmount = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) SetTotalInputAmountUnit(v string) *ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	s.TotalInputAmountUnit = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) SetTotalStorage(v float64) *ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	s.TotalStorage = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) SetTotalStorageFileCount(v int64) *ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	s.TotalStorageFileCount = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) SetTotalStorageUnit(v string) *ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	s.TotalStorageUnit = &v
	return s
}

type ListStorageTablesInfoResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListStorageTablesInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListStorageTablesInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s ListStorageTablesInfoResponse) GoString() string {
	return s.String()
}

func (s *ListStorageTablesInfoResponse) SetHeaders(v map[string]*string) *ListStorageTablesInfoResponse {
	s.Headers = v
	return s
}

func (s *ListStorageTablesInfoResponse) SetStatusCode(v int32) *ListStorageTablesInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ListStorageTablesInfoResponse) SetBody(v *ListStorageTablesInfoResponseBody) *ListStorageTablesInfoResponse {
	s.Body = v
	return s
}

type ListTablesRequest struct {
	// Specifies the marker after which the returned list begins.
	//
	// example:
	//
	// Y29tbWlzc2lvbl9leHRlcm5hbF91cmdlXzFfd3Ih
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
	// The maximum number of entries to return on each page.
	//
	// example:
	//
	// 10
	MaxItem *int32 `json:"maxItem,omitempty" xml:"maxItem,omitempty"`
	// The names of the returned resources. The names must start with the value specified by the prefix parameter. If the prefix parameter is set to a, the names of the returned resources must start with a.
	//
	// example:
	//
	// a
	Prefix *string `json:"prefix,omitempty" xml:"prefix,omitempty"`
	// The name of the schema.
	//
	// example:
	//
	// default
	SchemaName *string `json:"schemaName,omitempty" xml:"schemaName,omitempty"`
	// The type of the table.
	//
	// example:
	//
	// internal
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListTablesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTablesRequest) GoString() string {
	return s.String()
}

func (s *ListTablesRequest) SetMarker(v string) *ListTablesRequest {
	s.Marker = &v
	return s
}

func (s *ListTablesRequest) SetMaxItem(v int32) *ListTablesRequest {
	s.MaxItem = &v
	return s
}

func (s *ListTablesRequest) SetPrefix(v string) *ListTablesRequest {
	s.Prefix = &v
	return s
}

func (s *ListTablesRequest) SetSchemaName(v string) *ListTablesRequest {
	s.SchemaName = &v
	return s
}

func (s *ListTablesRequest) SetType(v string) *ListTablesRequest {
	s.Type = &v
	return s
}

type ListTablesResponseBody struct {
	// The returned data.
	Data *ListTablesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0a06dd4516687375802853481ec9fd
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListTablesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTablesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTablesResponseBody) SetData(v *ListTablesResponseBodyData) *ListTablesResponseBody {
	s.Data = v
	return s
}

func (s *ListTablesResponseBody) SetRequestId(v string) *ListTablesResponseBody {
	s.RequestId = &v
	return s
}

type ListTablesResponseBodyData struct {
	// Indicates the marker after which the returned list begins.
	//
	// example:
	//
	// cHlvZHBzX3VkZl8xMDExNV8xNDU3NDI4NDkzKg==
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
	// The maximum number of entries returned per page.
	//
	// example:
	//
	// 10
	MaxItem *int32 `json:"maxItem,omitempty" xml:"maxItem,omitempty"`
	// The information about tables.
	Tables []*ListTablesResponseBodyDataTables `json:"tables,omitempty" xml:"tables,omitempty" type:"Repeated"`
}

func (s ListTablesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListTablesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTablesResponseBodyData) SetMarker(v string) *ListTablesResponseBodyData {
	s.Marker = &v
	return s
}

func (s *ListTablesResponseBodyData) SetMaxItem(v int32) *ListTablesResponseBodyData {
	s.MaxItem = &v
	return s
}

func (s *ListTablesResponseBodyData) SetTables(v []*ListTablesResponseBodyDataTables) *ListTablesResponseBodyData {
	s.Tables = v
	return s
}

type ListTablesResponseBodyDataTables struct {
	// The time when the table was created.
	//
	// example:
	//
	// 2022-01-17T07:07:47Z
	CreationTime *int64 `json:"creationTime,omitempty" xml:"creationTime,omitempty"`
	// The display name of the table.
	//
	// example:
	//
	// sale_detail
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// dim_odps
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The owner of the table.
	//
	// example:
	//
	// 1887853961230110
	Owner *string `json:"owner,omitempty" xml:"owner,omitempty"`
	// The schema to which the table belongs.
	//
	// example:
	//
	// default
	Schema *string `json:"schema,omitempty" xml:"schema,omitempty"`
	// The type of the table.
	//
	// example:
	//
	// internal
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListTablesResponseBodyDataTables) String() string {
	return tea.Prettify(s)
}

func (s ListTablesResponseBodyDataTables) GoString() string {
	return s.String()
}

func (s *ListTablesResponseBodyDataTables) SetCreationTime(v int64) *ListTablesResponseBodyDataTables {
	s.CreationTime = &v
	return s
}

func (s *ListTablesResponseBodyDataTables) SetDisplayName(v string) *ListTablesResponseBodyDataTables {
	s.DisplayName = &v
	return s
}

func (s *ListTablesResponseBodyDataTables) SetName(v string) *ListTablesResponseBodyDataTables {
	s.Name = &v
	return s
}

func (s *ListTablesResponseBodyDataTables) SetOwner(v string) *ListTablesResponseBodyDataTables {
	s.Owner = &v
	return s
}

func (s *ListTablesResponseBodyDataTables) SetSchema(v string) *ListTablesResponseBodyDataTables {
	s.Schema = &v
	return s
}

func (s *ListTablesResponseBodyDataTables) SetType(v string) *ListTablesResponseBodyDataTables {
	s.Type = &v
	return s
}

type ListTablesResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTablesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTablesResponse) GoString() string {
	return s.String()
}

func (s *ListTablesResponse) SetHeaders(v map[string]*string) *ListTablesResponse {
	s.Headers = v
	return s
}

func (s *ListTablesResponse) SetStatusCode(v int32) *ListTablesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTablesResponse) SetBody(v *ListTablesResponseBody) *ListTablesResponse {
	s.Body = v
	return s
}

type ListTunnelQuotaTimerResponseBody struct {
	// The data returned.
	Data []*ListTunnelQuotaTimerResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The error code.
	//
	// example:
	//
	// OBJECT_NOT_EXIST
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// This object does not exist.
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// The HTTP status code.
	//
	// 	- 1xx: informational response. The request is received and is being processed.
	//
	// 	- 2xx: success. The request is successfully received, understood, and accepted by the server.
	//
	// 	- 3xx: redirection. The request is redirected, and further actions are required to complete the request.
	//
	// 	- 4xx: client error. The request contains invalid request parameters or syntaxes, or specific request conditions cannot be met.
	//
	// 	- 5xx: server error. The server cannot meet requirements due to other reasons.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0be3e0b716671885050924814e3623
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListTunnelQuotaTimerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTunnelQuotaTimerResponseBody) GoString() string {
	return s.String()
}

func (s *ListTunnelQuotaTimerResponseBody) SetData(v []*ListTunnelQuotaTimerResponseBodyData) *ListTunnelQuotaTimerResponseBody {
	s.Data = v
	return s
}

func (s *ListTunnelQuotaTimerResponseBody) SetErrorCode(v string) *ListTunnelQuotaTimerResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListTunnelQuotaTimerResponseBody) SetErrorMsg(v string) *ListTunnelQuotaTimerResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListTunnelQuotaTimerResponseBody) SetHttpCode(v int32) *ListTunnelQuotaTimerResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListTunnelQuotaTimerResponseBody) SetRequestId(v string) *ListTunnelQuotaTimerResponseBody {
	s.RequestId = &v
	return s
}

type ListTunnelQuotaTimerResponseBodyData struct {
	// The start time of the time-specific configuration.
	//
	// example:
	//
	// 00:00
	BeginTime *string `json:"beginTime,omitempty" xml:"beginTime,omitempty"`
	// The end time of the time-specific configuration.
	//
	// example:
	//
	// 08:00
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The time zone property for the time-specific configuration.
	//
	// example:
	//
	// Asia/Shanghai
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
	// The parameters for the time-specific configuration.
	TunnelQuotaParameter *ListTunnelQuotaTimerResponseBodyDataTunnelQuotaParameter `json:"tunnelQuotaParameter,omitempty" xml:"tunnelQuotaParameter,omitempty" type:"Struct"`
}

func (s ListTunnelQuotaTimerResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListTunnelQuotaTimerResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTunnelQuotaTimerResponseBodyData) SetBeginTime(v string) *ListTunnelQuotaTimerResponseBodyData {
	s.BeginTime = &v
	return s
}

func (s *ListTunnelQuotaTimerResponseBodyData) SetEndTime(v string) *ListTunnelQuotaTimerResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *ListTunnelQuotaTimerResponseBodyData) SetTimezone(v string) *ListTunnelQuotaTimerResponseBodyData {
	s.Timezone = &v
	return s
}

func (s *ListTunnelQuotaTimerResponseBodyData) SetTunnelQuotaParameter(v *ListTunnelQuotaTimerResponseBodyDataTunnelQuotaParameter) *ListTunnelQuotaTimerResponseBodyData {
	s.TunnelQuotaParameter = v
	return s
}

type ListTunnelQuotaTimerResponseBodyDataTunnelQuotaParameter struct {
	// The number of elastically reserved slots.
	//
	// example:
	//
	// 100
	ElasticReservedSlotNum *int64 `json:"elasticReservedSlotNum,omitempty" xml:"elasticReservedSlotNum,omitempty"`
	// The number of reserved slots.
	//
	// example:
	//
	// 100
	SlotNum *int64 `json:"slotNum,omitempty" xml:"slotNum,omitempty"`
}

func (s ListTunnelQuotaTimerResponseBodyDataTunnelQuotaParameter) String() string {
	return tea.Prettify(s)
}

func (s ListTunnelQuotaTimerResponseBodyDataTunnelQuotaParameter) GoString() string {
	return s.String()
}

func (s *ListTunnelQuotaTimerResponseBodyDataTunnelQuotaParameter) SetElasticReservedSlotNum(v int64) *ListTunnelQuotaTimerResponseBodyDataTunnelQuotaParameter {
	s.ElasticReservedSlotNum = &v
	return s
}

func (s *ListTunnelQuotaTimerResponseBodyDataTunnelQuotaParameter) SetSlotNum(v int64) *ListTunnelQuotaTimerResponseBodyDataTunnelQuotaParameter {
	s.SlotNum = &v
	return s
}

type ListTunnelQuotaTimerResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTunnelQuotaTimerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTunnelQuotaTimerResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTunnelQuotaTimerResponse) GoString() string {
	return s.String()
}

func (s *ListTunnelQuotaTimerResponse) SetHeaders(v map[string]*string) *ListTunnelQuotaTimerResponse {
	s.Headers = v
	return s
}

func (s *ListTunnelQuotaTimerResponse) SetStatusCode(v int32) *ListTunnelQuotaTimerResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTunnelQuotaTimerResponse) SetBody(v *ListTunnelQuotaTimerResponseBody) *ListTunnelQuotaTimerResponse {
	s.Body = v
	return s
}

type ListUsersRequest struct {
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUsersRequest) GoString() string {
	return s.String()
}

func (s *ListUsersRequest) SetPageNumber(v int32) *ListUsersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListUsersRequest) SetPageSize(v int32) *ListUsersRequest {
	s.PageSize = &v
	return s
}

type ListUsersResponseBody struct {
	// The returned data.
	Data *ListUsersResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 0a06dd4816687424611405643e3730
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBody) SetData(v *ListUsersResponseBodyData) *ListUsersResponseBody {
	s.Data = v
	return s
}

func (s *ListUsersResponseBody) SetRequestId(v string) *ListUsersResponseBody {
	s.RequestId = &v
	return s
}

type ListUsersResponseBodyData struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 2
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 64
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	// The users.
	Users []*ListUsersResponseBodyDataUsers `json:"users,omitempty" xml:"users,omitempty" type:"Repeated"`
}

func (s ListUsersResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyData) SetPageNumber(v int32) *ListUsersResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListUsersResponseBodyData) SetPageSize(v int32) *ListUsersResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListUsersResponseBodyData) SetTotalCount(v int32) *ListUsersResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListUsersResponseBodyData) SetUsers(v []*ListUsersResponseBodyDataUsers) *ListUsersResponseBodyData {
	s.Users = v
	return s
}

type ListUsersResponseBodyDataUsers struct {
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 167835629082
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// The username of the account.
	//
	// example:
	//
	// Bob@
	AccountName *string `json:"accountName,omitempty" xml:"accountName,omitempty"`
	// The type of the account.
	//
	// example:
	//
	// ALIYUN
	AccountType *string `json:"accountType,omitempty" xml:"accountType,omitempty"`
	// The display name.
	//
	// example:
	//
	// Bob
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// The ID of the tenant.
	//
	// example:
	//
	// 1567253789
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListUsersResponseBodyDataUsers) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBodyDataUsers) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyDataUsers) SetAccountId(v string) *ListUsersResponseBodyDataUsers {
	s.AccountId = &v
	return s
}

func (s *ListUsersResponseBodyDataUsers) SetAccountName(v string) *ListUsersResponseBodyDataUsers {
	s.AccountName = &v
	return s
}

func (s *ListUsersResponseBodyDataUsers) SetAccountType(v string) *ListUsersResponseBodyDataUsers {
	s.AccountType = &v
	return s
}

func (s *ListUsersResponseBodyDataUsers) SetDisplayName(v string) *ListUsersResponseBodyDataUsers {
	s.DisplayName = &v
	return s
}

func (s *ListUsersResponseBodyDataUsers) SetTenantId(v string) *ListUsersResponseBodyDataUsers {
	s.TenantId = &v
	return s
}

type ListUsersResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponse) GoString() string {
	return s.String()
}

func (s *ListUsersResponse) SetHeaders(v map[string]*string) *ListUsersResponse {
	s.Headers = v
	return s
}

func (s *ListUsersResponse) SetStatusCode(v int32) *ListUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUsersResponse) SetBody(v *ListUsersResponseBody) *ListUsersResponse {
	s.Body = v
	return s
}

type ListUsersByRoleResponseBody struct {
	// The returned data.
	Data *ListUsersByRoleResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 0be3e0bb16654558425251398e27a9
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListUsersByRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUsersByRoleResponseBody) GoString() string {
	return s.String()
}

func (s *ListUsersByRoleResponseBody) SetData(v *ListUsersByRoleResponseBodyData) *ListUsersByRoleResponseBody {
	s.Data = v
	return s
}

func (s *ListUsersByRoleResponseBody) SetRequestId(v string) *ListUsersByRoleResponseBody {
	s.RequestId = &v
	return s
}

type ListUsersByRoleResponseBodyData struct {
	// The users.
	Users []*ListUsersByRoleResponseBodyDataUsers `json:"users,omitempty" xml:"users,omitempty" type:"Repeated"`
}

func (s ListUsersByRoleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListUsersByRoleResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListUsersByRoleResponseBodyData) SetUsers(v []*ListUsersByRoleResponseBodyDataUsers) *ListUsersByRoleResponseBodyData {
	s.Users = v
	return s
}

type ListUsersByRoleResponseBodyDataUsers struct {
	// The name of the user.
	//
	// example:
	//
	// ALIYUN${account_name}
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListUsersByRoleResponseBodyDataUsers) String() string {
	return tea.Prettify(s)
}

func (s ListUsersByRoleResponseBodyDataUsers) GoString() string {
	return s.String()
}

func (s *ListUsersByRoleResponseBodyDataUsers) SetName(v string) *ListUsersByRoleResponseBodyDataUsers {
	s.Name = &v
	return s
}

type ListUsersByRoleResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUsersByRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUsersByRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUsersByRoleResponse) GoString() string {
	return s.String()
}

func (s *ListUsersByRoleResponse) SetHeaders(v map[string]*string) *ListUsersByRoleResponse {
	s.Headers = v
	return s
}

func (s *ListUsersByRoleResponse) SetStatusCode(v int32) *ListUsersByRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUsersByRoleResponse) SetBody(v *ListUsersByRoleResponseBody) *ListUsersByRoleResponse {
	s.Body = v
	return s
}

type QueryQuotaRequest struct {
	// example:
	//
	// null
	AkProven *string `json:"AkProven,omitempty" xml:"AkProven,omitempty"`
	// example:
	//
	// false
	Mock *bool `json:"mock,omitempty" xml:"mock,omitempty"`
	// example:
	//
	// cn-chengdu
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// example:
	//
	// 483212237127906
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s QueryQuotaRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryQuotaRequest) GoString() string {
	return s.String()
}

func (s *QueryQuotaRequest) SetAkProven(v string) *QueryQuotaRequest {
	s.AkProven = &v
	return s
}

func (s *QueryQuotaRequest) SetMock(v bool) *QueryQuotaRequest {
	s.Mock = &v
	return s
}

func (s *QueryQuotaRequest) SetRegion(v string) *QueryQuotaRequest {
	s.Region = &v
	return s
}

func (s *QueryQuotaRequest) SetTenantId(v string) *QueryQuotaRequest {
	s.TenantId = &v
	return s
}

type QueryQuotaResponseBody struct {
	Data *QueryQuotaResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg  *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// example:
	//
	// 0bc1eeed16675342848904412e08dd
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s QueryQuotaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *QueryQuotaResponseBody) SetData(v *QueryQuotaResponseBodyData) *QueryQuotaResponseBody {
	s.Data = v
	return s
}

func (s *QueryQuotaResponseBody) SetErrorCode(v string) *QueryQuotaResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QueryQuotaResponseBody) SetErrorMsg(v string) *QueryQuotaResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *QueryQuotaResponseBody) SetHttpCode(v int32) *QueryQuotaResponseBody {
	s.HttpCode = &v
	return s
}

func (s *QueryQuotaResponseBody) SetRequestId(v string) *QueryQuotaResponseBody {
	s.RequestId = &v
	return s
}

type QueryQuotaResponseBodyData struct {
	BillingPolicy *QueryQuotaResponseBodyDataBillingPolicy `json:"billingPolicy,omitempty" xml:"billingPolicy,omitempty" type:"Struct"`
	// example:
	//
	// AT-120N
	Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// example:
	//
	// 1714356241163
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 1248953767546358
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// example:
	//
	// abc
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// quota ID
	//
	// example:
	//
	// 2523
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// quota_a
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// quota_nickname
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	// example:
	//
	// {"minCU":10,
	//
	// "adhocCU":0,
	//
	// "maxCU":10,
	//
	// "schedulerType":"Fair",
	//
	// }
	Parameter map[string]interface{} `json:"parameter,omitempty" xml:"parameter,omitempty"`
	// example:
	//
	// null
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId     *string                                 `json:"regionId,omitempty" xml:"regionId,omitempty"`
	SaleTag      *QueryQuotaResponseBodyDataSaleTag      `json:"saleTag,omitempty" xml:"saleTag,omitempty" type:"Struct"`
	ScheduleInfo *QueryQuotaResponseBodyDataScheduleInfo `json:"scheduleInfo,omitempty" xml:"scheduleInfo,omitempty" type:"Struct"`
	// example:
	//
	// ON
	Status           *string                                       `json:"status,omitempty" xml:"status,omitempty"`
	SubQuotaInfoList []*QueryQuotaResponseBodyDataSubQuotaInfoList `json:"subQuotaInfoList,omitempty" xml:"subQuotaInfoList,omitempty" type:"Repeated"`
	// example:
	//
	// abc
	Tag *string `json:"tag,omitempty" xml:"tag,omitempty"`
	// example:
	//
	// 478403690625249
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// example:
	//
	// FUXI_OFFLINE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s QueryQuotaResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryQuotaResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryQuotaResponseBodyData) SetBillingPolicy(v *QueryQuotaResponseBodyDataBillingPolicy) *QueryQuotaResponseBodyData {
	s.BillingPolicy = v
	return s
}

func (s *QueryQuotaResponseBodyData) SetCluster(v string) *QueryQuotaResponseBodyData {
	s.Cluster = &v
	return s
}

func (s *QueryQuotaResponseBodyData) SetCreateTime(v int64) *QueryQuotaResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *QueryQuotaResponseBodyData) SetCreatorId(v string) *QueryQuotaResponseBodyData {
	s.CreatorId = &v
	return s
}

func (s *QueryQuotaResponseBodyData) SetGroupName(v string) *QueryQuotaResponseBodyData {
	s.GroupName = &v
	return s
}

func (s *QueryQuotaResponseBodyData) SetId(v string) *QueryQuotaResponseBodyData {
	s.Id = &v
	return s
}

func (s *QueryQuotaResponseBodyData) SetName(v string) *QueryQuotaResponseBodyData {
	s.Name = &v
	return s
}

func (s *QueryQuotaResponseBodyData) SetNickName(v string) *QueryQuotaResponseBodyData {
	s.NickName = &v
	return s
}

func (s *QueryQuotaResponseBodyData) SetParameter(v map[string]interface{}) *QueryQuotaResponseBodyData {
	s.Parameter = v
	return s
}

func (s *QueryQuotaResponseBodyData) SetParentId(v string) *QueryQuotaResponseBodyData {
	s.ParentId = &v
	return s
}

func (s *QueryQuotaResponseBodyData) SetRegionId(v string) *QueryQuotaResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *QueryQuotaResponseBodyData) SetSaleTag(v *QueryQuotaResponseBodyDataSaleTag) *QueryQuotaResponseBodyData {
	s.SaleTag = v
	return s
}

func (s *QueryQuotaResponseBodyData) SetScheduleInfo(v *QueryQuotaResponseBodyDataScheduleInfo) *QueryQuotaResponseBodyData {
	s.ScheduleInfo = v
	return s
}

func (s *QueryQuotaResponseBodyData) SetStatus(v string) *QueryQuotaResponseBodyData {
	s.Status = &v
	return s
}

func (s *QueryQuotaResponseBodyData) SetSubQuotaInfoList(v []*QueryQuotaResponseBodyDataSubQuotaInfoList) *QueryQuotaResponseBodyData {
	s.SubQuotaInfoList = v
	return s
}

func (s *QueryQuotaResponseBodyData) SetTag(v string) *QueryQuotaResponseBodyData {
	s.Tag = &v
	return s
}

func (s *QueryQuotaResponseBodyData) SetTenantId(v string) *QueryQuotaResponseBodyData {
	s.TenantId = &v
	return s
}

func (s *QueryQuotaResponseBodyData) SetType(v string) *QueryQuotaResponseBodyData {
	s.Type = &v
	return s
}

func (s *QueryQuotaResponseBodyData) SetVersion(v string) *QueryQuotaResponseBodyData {
	s.Version = &v
	return s
}

type QueryQuotaResponseBodyDataBillingPolicy struct {
	// example:
	//
	// subscription
	BillingMethod *string `json:"billingMethod,omitempty" xml:"billingMethod,omitempty"`
	// example:
	//
	// 880c0d0d-5d79-4217-b683-8e8bdd2a2523
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// OdpsStandard
	OdpsSpecCode *string `json:"odpsSpecCode,omitempty" xml:"odpsSpecCode,omitempty"`
	// example:
	//
	// 880c0d0d-5d79-4217-b683-8e8bdd2a2523
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
}

func (s QueryQuotaResponseBodyDataBillingPolicy) String() string {
	return tea.Prettify(s)
}

func (s QueryQuotaResponseBodyDataBillingPolicy) GoString() string {
	return s.String()
}

func (s *QueryQuotaResponseBodyDataBillingPolicy) SetBillingMethod(v string) *QueryQuotaResponseBodyDataBillingPolicy {
	s.BillingMethod = &v
	return s
}

func (s *QueryQuotaResponseBodyDataBillingPolicy) SetInstanceId(v string) *QueryQuotaResponseBodyDataBillingPolicy {
	s.InstanceId = &v
	return s
}

func (s *QueryQuotaResponseBodyDataBillingPolicy) SetOdpsSpecCode(v string) *QueryQuotaResponseBodyDataBillingPolicy {
	s.OdpsSpecCode = &v
	return s
}

func (s *QueryQuotaResponseBodyDataBillingPolicy) SetOrderId(v string) *QueryQuotaResponseBodyDataBillingPolicy {
	s.OrderId = &v
	return s
}

type QueryQuotaResponseBodyDataSaleTag struct {
	ResourceIds []*string `json:"resourceIds,omitempty" xml:"resourceIds,omitempty" type:"Repeated"`
	// example:
	//
	// project
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s QueryQuotaResponseBodyDataSaleTag) String() string {
	return tea.Prettify(s)
}

func (s QueryQuotaResponseBodyDataSaleTag) GoString() string {
	return s.String()
}

func (s *QueryQuotaResponseBodyDataSaleTag) SetResourceIds(v []*string) *QueryQuotaResponseBodyDataSaleTag {
	s.ResourceIds = v
	return s
}

func (s *QueryQuotaResponseBodyDataSaleTag) SetResourceType(v string) *QueryQuotaResponseBodyDataSaleTag {
	s.ResourceType = &v
	return s
}

type QueryQuotaResponseBodyDataScheduleInfo struct {
	// example:
	//
	// planA
	CurrPlan *string `json:"currPlan,omitempty" xml:"currPlan,omitempty"`
	// example:
	//
	// 0800
	CurrTime *string `json:"currTime,omitempty" xml:"currTime,omitempty"`
	// example:
	//
	// planB
	NextPlan *string `json:"nextPlan,omitempty" xml:"nextPlan,omitempty"`
	// example:
	//
	// 1700
	NextTime *string `json:"nextTime,omitempty" xml:"nextTime,omitempty"`
	// example:
	//
	// planC
	OncePlan *string `json:"oncePlan,omitempty" xml:"oncePlan,omitempty"`
	// example:
	//
	// 1500
	OnceTime *string `json:"onceTime,omitempty" xml:"onceTime,omitempty"`
	// example:
	//
	// userA
	OperatorName *string `json:"operatorName,omitempty" xml:"operatorName,omitempty"`
	// example:
	//
	// UTC+8
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
}

func (s QueryQuotaResponseBodyDataScheduleInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryQuotaResponseBodyDataScheduleInfo) GoString() string {
	return s.String()
}

func (s *QueryQuotaResponseBodyDataScheduleInfo) SetCurrPlan(v string) *QueryQuotaResponseBodyDataScheduleInfo {
	s.CurrPlan = &v
	return s
}

func (s *QueryQuotaResponseBodyDataScheduleInfo) SetCurrTime(v string) *QueryQuotaResponseBodyDataScheduleInfo {
	s.CurrTime = &v
	return s
}

func (s *QueryQuotaResponseBodyDataScheduleInfo) SetNextPlan(v string) *QueryQuotaResponseBodyDataScheduleInfo {
	s.NextPlan = &v
	return s
}

func (s *QueryQuotaResponseBodyDataScheduleInfo) SetNextTime(v string) *QueryQuotaResponseBodyDataScheduleInfo {
	s.NextTime = &v
	return s
}

func (s *QueryQuotaResponseBodyDataScheduleInfo) SetOncePlan(v string) *QueryQuotaResponseBodyDataScheduleInfo {
	s.OncePlan = &v
	return s
}

func (s *QueryQuotaResponseBodyDataScheduleInfo) SetOnceTime(v string) *QueryQuotaResponseBodyDataScheduleInfo {
	s.OnceTime = &v
	return s
}

func (s *QueryQuotaResponseBodyDataScheduleInfo) SetOperatorName(v string) *QueryQuotaResponseBodyDataScheduleInfo {
	s.OperatorName = &v
	return s
}

func (s *QueryQuotaResponseBodyDataScheduleInfo) SetTimezone(v string) *QueryQuotaResponseBodyDataScheduleInfo {
	s.Timezone = &v
	return s
}

type QueryQuotaResponseBodyDataSubQuotaInfoList struct {
	BillingPolicy *QueryQuotaResponseBodyDataSubQuotaInfoListBillingPolicy `json:"billingPolicy,omitempty" xml:"billingPolicy,omitempty" type:"Struct"`
	// example:
	//
	// AT-120N
	Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// example:
	//
	// 1688653978768
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 672863518
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// example:
	//
	// abc
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// example:
	//
	// 1000048
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// subquotaA
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// subquotaA
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	// example:
	//
	// {\\"maxCU\\": 10, \\"minCU\\": 10, \\"adhocCU\\": 0, \\"schedulerType\\": \\"Fifo\\"}
	Parameter map[string]interface{} `json:"parameter,omitempty" xml:"parameter,omitempty"`
	// example:
	//
	// 0
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId     *string                                                 `json:"regionId,omitempty" xml:"regionId,omitempty"`
	SaleTag      *QueryQuotaResponseBodyDataSubQuotaInfoListSaleTag      `json:"saleTag,omitempty" xml:"saleTag,omitempty" type:"Struct"`
	ScheduleInfo *QueryQuotaResponseBodyDataSubQuotaInfoListScheduleInfo `json:"scheduleInfo,omitempty" xml:"scheduleInfo,omitempty" type:"Struct"`
	// example:
	//
	// ON
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// abc
	Tag *string `json:"tag,omitempty" xml:"tag,omitempty"`
	// example:
	//
	// 478403690625249
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// example:
	//
	// FUXI_ONLINE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s QueryQuotaResponseBodyDataSubQuotaInfoList) String() string {
	return tea.Prettify(s)
}

func (s QueryQuotaResponseBodyDataSubQuotaInfoList) GoString() string {
	return s.String()
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoList) SetBillingPolicy(v *QueryQuotaResponseBodyDataSubQuotaInfoListBillingPolicy) *QueryQuotaResponseBodyDataSubQuotaInfoList {
	s.BillingPolicy = v
	return s
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoList) SetCluster(v string) *QueryQuotaResponseBodyDataSubQuotaInfoList {
	s.Cluster = &v
	return s
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoList) SetCreateTime(v int64) *QueryQuotaResponseBodyDataSubQuotaInfoList {
	s.CreateTime = &v
	return s
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoList) SetCreatorId(v string) *QueryQuotaResponseBodyDataSubQuotaInfoList {
	s.CreatorId = &v
	return s
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoList) SetGroupName(v string) *QueryQuotaResponseBodyDataSubQuotaInfoList {
	s.GroupName = &v
	return s
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoList) SetId(v string) *QueryQuotaResponseBodyDataSubQuotaInfoList {
	s.Id = &v
	return s
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoList) SetName(v string) *QueryQuotaResponseBodyDataSubQuotaInfoList {
	s.Name = &v
	return s
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoList) SetNickName(v string) *QueryQuotaResponseBodyDataSubQuotaInfoList {
	s.NickName = &v
	return s
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoList) SetParameter(v map[string]interface{}) *QueryQuotaResponseBodyDataSubQuotaInfoList {
	s.Parameter = v
	return s
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoList) SetParentId(v string) *QueryQuotaResponseBodyDataSubQuotaInfoList {
	s.ParentId = &v
	return s
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoList) SetRegionId(v string) *QueryQuotaResponseBodyDataSubQuotaInfoList {
	s.RegionId = &v
	return s
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoList) SetSaleTag(v *QueryQuotaResponseBodyDataSubQuotaInfoListSaleTag) *QueryQuotaResponseBodyDataSubQuotaInfoList {
	s.SaleTag = v
	return s
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoList) SetScheduleInfo(v *QueryQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) *QueryQuotaResponseBodyDataSubQuotaInfoList {
	s.ScheduleInfo = v
	return s
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoList) SetStatus(v string) *QueryQuotaResponseBodyDataSubQuotaInfoList {
	s.Status = &v
	return s
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoList) SetTag(v string) *QueryQuotaResponseBodyDataSubQuotaInfoList {
	s.Tag = &v
	return s
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoList) SetTenantId(v string) *QueryQuotaResponseBodyDataSubQuotaInfoList {
	s.TenantId = &v
	return s
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoList) SetType(v string) *QueryQuotaResponseBodyDataSubQuotaInfoList {
	s.Type = &v
	return s
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoList) SetVersion(v string) *QueryQuotaResponseBodyDataSubQuotaInfoList {
	s.Version = &v
	return s
}

type QueryQuotaResponseBodyDataSubQuotaInfoListBillingPolicy struct {
	// example:
	//
	// subscription
	BillingMethod *string `json:"billingMethod,omitempty" xml:"billingMethod,omitempty"`
	// example:
	//
	// 880c0d0d-5d79-4217-b683-8e8bdd2a2523
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// OdpsStandard
	OdpsSpecCode *string `json:"odpsSpecCode,omitempty" xml:"odpsSpecCode,omitempty"`
	// example:
	//
	// 880c0d0d-5d79-4217-b683-8e8bdd2a2523
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
}

func (s QueryQuotaResponseBodyDataSubQuotaInfoListBillingPolicy) String() string {
	return tea.Prettify(s)
}

func (s QueryQuotaResponseBodyDataSubQuotaInfoListBillingPolicy) GoString() string {
	return s.String()
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoListBillingPolicy) SetBillingMethod(v string) *QueryQuotaResponseBodyDataSubQuotaInfoListBillingPolicy {
	s.BillingMethod = &v
	return s
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoListBillingPolicy) SetInstanceId(v string) *QueryQuotaResponseBodyDataSubQuotaInfoListBillingPolicy {
	s.InstanceId = &v
	return s
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoListBillingPolicy) SetOdpsSpecCode(v string) *QueryQuotaResponseBodyDataSubQuotaInfoListBillingPolicy {
	s.OdpsSpecCode = &v
	return s
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoListBillingPolicy) SetOrderId(v string) *QueryQuotaResponseBodyDataSubQuotaInfoListBillingPolicy {
	s.OrderId = &v
	return s
}

type QueryQuotaResponseBodyDataSubQuotaInfoListSaleTag struct {
	ResourceIds []*string `json:"resourceIds,omitempty" xml:"resourceIds,omitempty" type:"Repeated"`
	// example:
	//
	// "quota"
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s QueryQuotaResponseBodyDataSubQuotaInfoListSaleTag) String() string {
	return tea.Prettify(s)
}

func (s QueryQuotaResponseBodyDataSubQuotaInfoListSaleTag) GoString() string {
	return s.String()
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoListSaleTag) SetResourceIds(v []*string) *QueryQuotaResponseBodyDataSubQuotaInfoListSaleTag {
	s.ResourceIds = v
	return s
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoListSaleTag) SetResourceType(v string) *QueryQuotaResponseBodyDataSubQuotaInfoListSaleTag {
	s.ResourceType = &v
	return s
}

type QueryQuotaResponseBodyDataSubQuotaInfoListScheduleInfo struct {
	// example:
	//
	// planA
	CurrPlan *string `json:"currPlan,omitempty" xml:"currPlan,omitempty"`
	// example:
	//
	// 0800
	CurrTime *string `json:"currTime,omitempty" xml:"currTime,omitempty"`
	// example:
	//
	// planB
	NextPlan *string `json:"nextPlan,omitempty" xml:"nextPlan,omitempty"`
	// example:
	//
	// 1700
	NextTime *string `json:"nextTime,omitempty" xml:"nextTime,omitempty"`
	// example:
	//
	// planC
	OncePlan *string `json:"oncePlan,omitempty" xml:"oncePlan,omitempty"`
	// example:
	//
	// 1500
	OnceTime *string `json:"onceTime,omitempty" xml:"onceTime,omitempty"`
	// example:
	//
	// userA
	OperatorName *string `json:"operatorName,omitempty" xml:"operatorName,omitempty"`
	// example:
	//
	// UTC+8
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
}

func (s QueryQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) GoString() string {
	return s.String()
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) SetCurrPlan(v string) *QueryQuotaResponseBodyDataSubQuotaInfoListScheduleInfo {
	s.CurrPlan = &v
	return s
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) SetCurrTime(v string) *QueryQuotaResponseBodyDataSubQuotaInfoListScheduleInfo {
	s.CurrTime = &v
	return s
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) SetNextPlan(v string) *QueryQuotaResponseBodyDataSubQuotaInfoListScheduleInfo {
	s.NextPlan = &v
	return s
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) SetNextTime(v string) *QueryQuotaResponseBodyDataSubQuotaInfoListScheduleInfo {
	s.NextTime = &v
	return s
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) SetOncePlan(v string) *QueryQuotaResponseBodyDataSubQuotaInfoListScheduleInfo {
	s.OncePlan = &v
	return s
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) SetOnceTime(v string) *QueryQuotaResponseBodyDataSubQuotaInfoListScheduleInfo {
	s.OnceTime = &v
	return s
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) SetOperatorName(v string) *QueryQuotaResponseBodyDataSubQuotaInfoListScheduleInfo {
	s.OperatorName = &v
	return s
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) SetTimezone(v string) *QueryQuotaResponseBodyDataSubQuotaInfoListScheduleInfo {
	s.Timezone = &v
	return s
}

type QueryQuotaResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryQuotaResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryQuotaResponse) GoString() string {
	return s.String()
}

func (s *QueryQuotaResponse) SetHeaders(v map[string]*string) *QueryQuotaResponse {
	s.Headers = v
	return s
}

func (s *QueryQuotaResponse) SetStatusCode(v int32) *QueryQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryQuotaResponse) SetBody(v *QueryQuotaResponseBody) *QueryQuotaResponse {
	s.Body = v
	return s
}

type RetryMmsJobResponseBody struct {
	// example:
	//
	// 78
	Data *int64 `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// 7F5DAD1C-9EC2-5FE5-97CF-BCE21B4ABA29
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s RetryMmsJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RetryMmsJobResponseBody) GoString() string {
	return s.String()
}

func (s *RetryMmsJobResponseBody) SetData(v int64) *RetryMmsJobResponseBody {
	s.Data = &v
	return s
}

func (s *RetryMmsJobResponseBody) SetRequestId(v string) *RetryMmsJobResponseBody {
	s.RequestId = &v
	return s
}

type RetryMmsJobResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RetryMmsJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RetryMmsJobResponse) String() string {
	return tea.Prettify(s)
}

func (s RetryMmsJobResponse) GoString() string {
	return s.String()
}

func (s *RetryMmsJobResponse) SetHeaders(v map[string]*string) *RetryMmsJobResponse {
	s.Headers = v
	return s
}

func (s *RetryMmsJobResponse) SetStatusCode(v int32) *RetryMmsJobResponse {
	s.StatusCode = &v
	return s
}

func (s *RetryMmsJobResponse) SetBody(v *RetryMmsJobResponseBody) *RetryMmsJobResponse {
	s.Body = v
	return s
}

type StartMmsJobResponseBody struct {
	// example:
	//
	// 88
	Data *int64 `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// 5CA6292A-E301-5CD8-B4E2-AF060F99147B
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s StartMmsJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartMmsJobResponseBody) GoString() string {
	return s.String()
}

func (s *StartMmsJobResponseBody) SetData(v int64) *StartMmsJobResponseBody {
	s.Data = &v
	return s
}

func (s *StartMmsJobResponseBody) SetRequestId(v string) *StartMmsJobResponseBody {
	s.RequestId = &v
	return s
}

type StartMmsJobResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartMmsJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartMmsJobResponse) String() string {
	return tea.Prettify(s)
}

func (s StartMmsJobResponse) GoString() string {
	return s.String()
}

func (s *StartMmsJobResponse) SetHeaders(v map[string]*string) *StartMmsJobResponse {
	s.Headers = v
	return s
}

func (s *StartMmsJobResponse) SetStatusCode(v int32) *StartMmsJobResponse {
	s.StatusCode = &v
	return s
}

func (s *StartMmsJobResponse) SetBody(v *StartMmsJobResponseBody) *StartMmsJobResponse {
	s.Body = v
	return s
}

type StopMmsJobResponseBody struct {
	// example:
	//
	// 88
	Data *int64 `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// 8023D058-62B7-5C49-8EB6-AD9BA7942BC5
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s StopMmsJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopMmsJobResponseBody) GoString() string {
	return s.String()
}

func (s *StopMmsJobResponseBody) SetData(v int64) *StopMmsJobResponseBody {
	s.Data = &v
	return s
}

func (s *StopMmsJobResponseBody) SetRequestId(v string) *StopMmsJobResponseBody {
	s.RequestId = &v
	return s
}

type StopMmsJobResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopMmsJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopMmsJobResponse) String() string {
	return tea.Prettify(s)
}

func (s StopMmsJobResponse) GoString() string {
	return s.String()
}

func (s *StopMmsJobResponse) SetHeaders(v map[string]*string) *StopMmsJobResponse {
	s.Headers = v
	return s
}

func (s *StopMmsJobResponse) SetStatusCode(v int32) *StopMmsJobResponse {
	s.StatusCode = &v
	return s
}

func (s *StopMmsJobResponse) SetBody(v *StopMmsJobResponseBody) *StopMmsJobResponse {
	s.Body = v
	return s
}

type UpdateComputeQuotaPlanRequest struct {
	// The name of quota plan.
	//
	// This parameter is required.
	//
	// example:
	//
	// planA
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The parameters of quota plan.
	Quota *UpdateComputeQuotaPlanRequestQuota `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
}

func (s UpdateComputeQuotaPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateComputeQuotaPlanRequest) GoString() string {
	return s.String()
}

func (s *UpdateComputeQuotaPlanRequest) SetName(v string) *UpdateComputeQuotaPlanRequest {
	s.Name = &v
	return s
}

func (s *UpdateComputeQuotaPlanRequest) SetQuota(v *UpdateComputeQuotaPlanRequestQuota) *UpdateComputeQuotaPlanRequest {
	s.Quota = v
	return s
}

type UpdateComputeQuotaPlanRequestQuota struct {
	// The parameters of level-1 quota.
	Parameter *UpdateComputeQuotaPlanRequestQuotaParameter `json:"parameter,omitempty" xml:"parameter,omitempty" type:"Struct"`
	// The list of level-2 quotas.
	SubQuotaInfoList []*UpdateComputeQuotaPlanRequestQuotaSubQuotaInfoList `json:"subQuotaInfoList,omitempty" xml:"subQuotaInfoList,omitempty" type:"Repeated"`
}

func (s UpdateComputeQuotaPlanRequestQuota) String() string {
	return tea.Prettify(s)
}

func (s UpdateComputeQuotaPlanRequestQuota) GoString() string {
	return s.String()
}

func (s *UpdateComputeQuotaPlanRequestQuota) SetParameter(v *UpdateComputeQuotaPlanRequestQuotaParameter) *UpdateComputeQuotaPlanRequestQuota {
	s.Parameter = v
	return s
}

func (s *UpdateComputeQuotaPlanRequestQuota) SetSubQuotaInfoList(v []*UpdateComputeQuotaPlanRequestQuotaSubQuotaInfoList) *UpdateComputeQuotaPlanRequestQuota {
	s.SubQuotaInfoList = v
	return s
}

type UpdateComputeQuotaPlanRequestQuotaParameter struct {
	// The value of elastic Reserved CUs in the level-1 quota.
	//
	// > The default value is 0. The maximum value of this parameter must be equal to the number of subscription-based reserved CUs and cannot exceed 10,000 CUs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 50
	ElasticReservedCU *int64 `json:"elasticReservedCU,omitempty" xml:"elasticReservedCU,omitempty"`
}

func (s UpdateComputeQuotaPlanRequestQuotaParameter) String() string {
	return tea.Prettify(s)
}

func (s UpdateComputeQuotaPlanRequestQuotaParameter) GoString() string {
	return s.String()
}

func (s *UpdateComputeQuotaPlanRequestQuotaParameter) SetElasticReservedCU(v int64) *UpdateComputeQuotaPlanRequestQuotaParameter {
	s.ElasticReservedCU = &v
	return s
}

type UpdateComputeQuotaPlanRequestQuotaSubQuotaInfoList struct {
	// The nickname of the level-2 quota.
	//
	// This parameter is required.
	//
	// example:
	//
	// os_ComputeQuota
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	// The parameters of the level-2 quota.
	Parameter *UpdateComputeQuotaPlanRequestQuotaSubQuotaInfoListParameter `json:"parameter,omitempty" xml:"parameter,omitempty" type:"Struct"`
}

func (s UpdateComputeQuotaPlanRequestQuotaSubQuotaInfoList) String() string {
	return tea.Prettify(s)
}

func (s UpdateComputeQuotaPlanRequestQuotaSubQuotaInfoList) GoString() string {
	return s.String()
}

func (s *UpdateComputeQuotaPlanRequestQuotaSubQuotaInfoList) SetNickName(v string) *UpdateComputeQuotaPlanRequestQuotaSubQuotaInfoList {
	s.NickName = &v
	return s
}

func (s *UpdateComputeQuotaPlanRequestQuotaSubQuotaInfoList) SetParameter(v *UpdateComputeQuotaPlanRequestQuotaSubQuotaInfoListParameter) *UpdateComputeQuotaPlanRequestQuotaSubQuotaInfoList {
	s.Parameter = v
	return s
}

type UpdateComputeQuotaPlanRequestQuotaSubQuotaInfoListParameter struct {
	// The value of elastic Reserved CUs.
	//
	// > The total number of elastically reserved CUs in all the level-2 quotas is equal to the number of elastically reserved CUs in the level-1 quota.
	//
	// This parameter is required.
	//
	// example:
	//
	// 50
	ElasticReservedCU *int64 `json:"elasticReservedCU,omitempty" xml:"elasticReservedCU,omitempty"`
	// The value of maxCU in Reserved CUs.
	//
	// > The value of maxCU must be less than or equal to the value of maxCU in the level-1 quota that you purchased.
	//
	// This parameter is required.
	//
	// example:
	//
	// 50
	MaxCU *int64 `json:"maxCU,omitempty" xml:"maxCU,omitempty"`
	// The value of minCU in Reserved CUs.
	//
	// >
	//
	// >- The total value of minCU in all the level-2 quotas is equal to the value of minCU in the level-1 quota.
	//
	// >- The value of minCU must be less than or equal to the value of maxCU in the level-2 quota and less than or equal to the value of minCU in the level-1 quota that you purchased.
	//
	// This parameter is required.
	//
	// example:
	//
	// 50
	MinCU *int64 `json:"minCU,omitempty" xml:"minCU,omitempty"`
}

func (s UpdateComputeQuotaPlanRequestQuotaSubQuotaInfoListParameter) String() string {
	return tea.Prettify(s)
}

func (s UpdateComputeQuotaPlanRequestQuotaSubQuotaInfoListParameter) GoString() string {
	return s.String()
}

func (s *UpdateComputeQuotaPlanRequestQuotaSubQuotaInfoListParameter) SetElasticReservedCU(v int64) *UpdateComputeQuotaPlanRequestQuotaSubQuotaInfoListParameter {
	s.ElasticReservedCU = &v
	return s
}

func (s *UpdateComputeQuotaPlanRequestQuotaSubQuotaInfoListParameter) SetMaxCU(v int64) *UpdateComputeQuotaPlanRequestQuotaSubQuotaInfoListParameter {
	s.MaxCU = &v
	return s
}

func (s *UpdateComputeQuotaPlanRequestQuotaSubQuotaInfoListParameter) SetMinCU(v int64) *UpdateComputeQuotaPlanRequestQuotaSubQuotaInfoListParameter {
	s.MinCU = &v
	return s
}

type UpdateComputeQuotaPlanResponseBody struct {
	// The data returned.
	//
	// example:
	//
	// success
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// The error code.
	//
	// example:
	//
	// QUOTA_PLAN_NOT_FOUND
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// plan \\"***\\" does not exist
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// The HTTP status code.
	//
	// - 1xx: informational response. The request is received and is being processed.
	//
	// - 2xx: success. The request is successfully received, understood, and accepted by the server.
	//
	// - 3xx: redirection. The request is redirected, and further actions are required to complete the request.
	//
	// - 4xx: client error. The request contains invalid request parameters or syntaxes, or specific request conditions cannot be met.
	//
	// - 5xx: server error. The server cannot meet requirements due to other reasons.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0be3e0aa16667684362147582e038f
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateComputeQuotaPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateComputeQuotaPlanResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateComputeQuotaPlanResponseBody) SetData(v string) *UpdateComputeQuotaPlanResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateComputeQuotaPlanResponseBody) SetErrorCode(v string) *UpdateComputeQuotaPlanResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateComputeQuotaPlanResponseBody) SetErrorMsg(v string) *UpdateComputeQuotaPlanResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *UpdateComputeQuotaPlanResponseBody) SetHttpCode(v int32) *UpdateComputeQuotaPlanResponseBody {
	s.HttpCode = &v
	return s
}

func (s *UpdateComputeQuotaPlanResponseBody) SetRequestId(v string) *UpdateComputeQuotaPlanResponseBody {
	s.RequestId = &v
	return s
}

type UpdateComputeQuotaPlanResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateComputeQuotaPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateComputeQuotaPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateComputeQuotaPlanResponse) GoString() string {
	return s.String()
}

func (s *UpdateComputeQuotaPlanResponse) SetHeaders(v map[string]*string) *UpdateComputeQuotaPlanResponse {
	s.Headers = v
	return s
}

func (s *UpdateComputeQuotaPlanResponse) SetStatusCode(v int32) *UpdateComputeQuotaPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateComputeQuotaPlanResponse) SetBody(v *UpdateComputeQuotaPlanResponseBody) *UpdateComputeQuotaPlanResponse {
	s.Body = v
	return s
}

type UpdateComputeQuotaScheduleRequest struct {
	// The request body parameters.
	Body             []*UpdateComputeQuotaScheduleRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	ScheduleTimezone *string                                  `json:"scheduleTimezone,omitempty" xml:"scheduleTimezone,omitempty"`
}

func (s UpdateComputeQuotaScheduleRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateComputeQuotaScheduleRequest) GoString() string {
	return s.String()
}

func (s *UpdateComputeQuotaScheduleRequest) SetBody(v []*UpdateComputeQuotaScheduleRequestBody) *UpdateComputeQuotaScheduleRequest {
	s.Body = v
	return s
}

func (s *UpdateComputeQuotaScheduleRequest) SetScheduleTimezone(v string) *UpdateComputeQuotaScheduleRequest {
	s.ScheduleTimezone = &v
	return s
}

type UpdateComputeQuotaScheduleRequestBody struct {
	// The value of effective condition.
	Condition *UpdateComputeQuotaScheduleRequestBodyCondition `json:"condition,omitempty" xml:"condition,omitempty" type:"Struct"`
	// The name of the quota plan.
	//
	// This parameter is required.
	//
	// example:
	//
	// planA
	Plan *string `json:"plan,omitempty" xml:"plan,omitempty"`
	// The type of the quota plan.
	//
	// 	Notice: Currently, only daily is supported.</notice>
	//
	// This parameter is required.
	//
	// example:
	//
	// daily
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateComputeQuotaScheduleRequestBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateComputeQuotaScheduleRequestBody) GoString() string {
	return s.String()
}

func (s *UpdateComputeQuotaScheduleRequestBody) SetCondition(v *UpdateComputeQuotaScheduleRequestBodyCondition) *UpdateComputeQuotaScheduleRequestBody {
	s.Condition = v
	return s
}

func (s *UpdateComputeQuotaScheduleRequestBody) SetPlan(v string) *UpdateComputeQuotaScheduleRequestBody {
	s.Plan = &v
	return s
}

func (s *UpdateComputeQuotaScheduleRequestBody) SetType(v string) *UpdateComputeQuotaScheduleRequestBody {
	s.Type = &v
	return s
}

type UpdateComputeQuotaScheduleRequestBodyCondition struct {
	// The start time when the quota plan takes effect.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10:00
	At *string `json:"at,omitempty" xml:"at,omitempty"`
}

func (s UpdateComputeQuotaScheduleRequestBodyCondition) String() string {
	return tea.Prettify(s)
}

func (s UpdateComputeQuotaScheduleRequestBodyCondition) GoString() string {
	return s.String()
}

func (s *UpdateComputeQuotaScheduleRequestBodyCondition) SetAt(v string) *UpdateComputeQuotaScheduleRequestBodyCondition {
	s.At = &v
	return s
}

type UpdateComputeQuotaScheduleResponseBody struct {
	// The data returned.
	//
	// example:
	//
	// success
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// The error code.
	//
	// example:
	//
	// QUOTA_PLAN_NOT_FOUND
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// plan \\"***\\" does not exist
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// HTTP status code.
	//
	// - 1xx: Informational - The request has been received and is being processed.
	//
	// - 2xx: Success - The request action was successfully received, understood, and accepted by the server.
	//
	// - 3xx: Redirection - Further action must be taken to complete the request.
	//
	// - 4xx: Client Error - The request contains an error in the request parameters, syntax, or specific request conditions cannot be met.
	//
	// - 5xx: Server Error - The server could not fulfill the request due to other reasons.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0be3e0aa16667684362147582e038f
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateComputeQuotaScheduleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateComputeQuotaScheduleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateComputeQuotaScheduleResponseBody) SetData(v string) *UpdateComputeQuotaScheduleResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateComputeQuotaScheduleResponseBody) SetErrorCode(v string) *UpdateComputeQuotaScheduleResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateComputeQuotaScheduleResponseBody) SetErrorMsg(v string) *UpdateComputeQuotaScheduleResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *UpdateComputeQuotaScheduleResponseBody) SetHttpCode(v int32) *UpdateComputeQuotaScheduleResponseBody {
	s.HttpCode = &v
	return s
}

func (s *UpdateComputeQuotaScheduleResponseBody) SetRequestId(v string) *UpdateComputeQuotaScheduleResponseBody {
	s.RequestId = &v
	return s
}

type UpdateComputeQuotaScheduleResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateComputeQuotaScheduleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateComputeQuotaScheduleResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateComputeQuotaScheduleResponse) GoString() string {
	return s.String()
}

func (s *UpdateComputeQuotaScheduleResponse) SetHeaders(v map[string]*string) *UpdateComputeQuotaScheduleResponse {
	s.Headers = v
	return s
}

func (s *UpdateComputeQuotaScheduleResponse) SetStatusCode(v int32) *UpdateComputeQuotaScheduleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateComputeQuotaScheduleResponse) SetBody(v *UpdateComputeQuotaScheduleResponseBody) *UpdateComputeQuotaScheduleResponse {
	s.Body = v
	return s
}

type UpdateComputeSubQuotaRequest struct {
	// The list of level-2 quotas.
	SubQuotaInfoList []*UpdateComputeSubQuotaRequestSubQuotaInfoList `json:"subQuotaInfoList,omitempty" xml:"subQuotaInfoList,omitempty" type:"Repeated"`
}

func (s UpdateComputeSubQuotaRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateComputeSubQuotaRequest) GoString() string {
	return s.String()
}

func (s *UpdateComputeSubQuotaRequest) SetSubQuotaInfoList(v []*UpdateComputeSubQuotaRequestSubQuotaInfoList) *UpdateComputeSubQuotaRequest {
	s.SubQuotaInfoList = v
	return s
}

type UpdateComputeSubQuotaRequestSubQuotaInfoList struct {
	// The nickname of the level-2 quota.
	//
	// This parameter is required.
	//
	// example:
	//
	// os_ComputeQuota
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	// The parameters of the level-2 quota.
	Parameter *UpdateComputeSubQuotaRequestSubQuotaInfoListParameter `json:"parameter,omitempty" xml:"parameter,omitempty" type:"Struct"`
	// The type of quota.
	//
	// >
	//
	// > - FUXI_OFFLINE(default) : Quotas of this type are used to run batch jobs.
	//
	// example:
	//
	// FUXI_OFFLINE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateComputeSubQuotaRequestSubQuotaInfoList) String() string {
	return tea.Prettify(s)
}

func (s UpdateComputeSubQuotaRequestSubQuotaInfoList) GoString() string {
	return s.String()
}

func (s *UpdateComputeSubQuotaRequestSubQuotaInfoList) SetNickName(v string) *UpdateComputeSubQuotaRequestSubQuotaInfoList {
	s.NickName = &v
	return s
}

func (s *UpdateComputeSubQuotaRequestSubQuotaInfoList) SetParameter(v *UpdateComputeSubQuotaRequestSubQuotaInfoListParameter) *UpdateComputeSubQuotaRequestSubQuotaInfoList {
	s.Parameter = v
	return s
}

func (s *UpdateComputeSubQuotaRequestSubQuotaInfoList) SetType(v string) *UpdateComputeSubQuotaRequestSubQuotaInfoList {
	s.Type = &v
	return s
}

type UpdateComputeSubQuotaRequestSubQuotaInfoListParameter struct {
	// Specifies whether to enable the priority feature.
	//
	// example:
	//
	// false
	EnablePriority *bool `json:"enablePriority,omitempty" xml:"enablePriority,omitempty"`
	// Specifies whether the quota is strongly exclusive.
	//
	// example:
	//
	// false
	ForceReservedMin *bool `json:"forceReservedMin,omitempty" xml:"forceReservedMin,omitempty"`
	// The value of minCU in Reserved CUs.
	//
	// > The value of maxCU must be less than or equal to the value of maxCU in the level-1 quota that you purchased.
	//
	// This parameter is required.
	//
	// example:
	//
	// 50
	MaxCU *int64 `json:"maxCU,omitempty" xml:"maxCU,omitempty"`
	// The value of maxCU in Reserved CUs.
	//
	// >
	//
	// >- The total value of minCU in all the level-2 quotas is equal to the value of minCU in the level-1 quota.
	//
	// >- The value of minCU must be less than or equal to the value of maxCU in the level-2 quota and less than or equal to the value of minCU in the level-1 quota that you purchased.
	//
	// This parameter is required.
	//
	// example:
	//
	// 50
	MinCU *int64 `json:"minCU,omitempty" xml:"minCU,omitempty"`
	// Scheduling policy of the quota.
	//
	// example:
	//
	// Fifo/Fair
	SchedulerType *string `json:"schedulerType,omitempty" xml:"schedulerType,omitempty"`
	// The upper limit for CUs that can be concurrently used by a job scheduled to the quota.
	//
	// example:
	//
	// 10
	SingleJobCULimit *int64 `json:"singleJobCULimit,omitempty" xml:"singleJobCULimit,omitempty"`
}

func (s UpdateComputeSubQuotaRequestSubQuotaInfoListParameter) String() string {
	return tea.Prettify(s)
}

func (s UpdateComputeSubQuotaRequestSubQuotaInfoListParameter) GoString() string {
	return s.String()
}

func (s *UpdateComputeSubQuotaRequestSubQuotaInfoListParameter) SetEnablePriority(v bool) *UpdateComputeSubQuotaRequestSubQuotaInfoListParameter {
	s.EnablePriority = &v
	return s
}

func (s *UpdateComputeSubQuotaRequestSubQuotaInfoListParameter) SetForceReservedMin(v bool) *UpdateComputeSubQuotaRequestSubQuotaInfoListParameter {
	s.ForceReservedMin = &v
	return s
}

func (s *UpdateComputeSubQuotaRequestSubQuotaInfoListParameter) SetMaxCU(v int64) *UpdateComputeSubQuotaRequestSubQuotaInfoListParameter {
	s.MaxCU = &v
	return s
}

func (s *UpdateComputeSubQuotaRequestSubQuotaInfoListParameter) SetMinCU(v int64) *UpdateComputeSubQuotaRequestSubQuotaInfoListParameter {
	s.MinCU = &v
	return s
}

func (s *UpdateComputeSubQuotaRequestSubQuotaInfoListParameter) SetSchedulerType(v string) *UpdateComputeSubQuotaRequestSubQuotaInfoListParameter {
	s.SchedulerType = &v
	return s
}

func (s *UpdateComputeSubQuotaRequestSubQuotaInfoListParameter) SetSingleJobCULimit(v int64) *UpdateComputeSubQuotaRequestSubQuotaInfoListParameter {
	s.SingleJobCULimit = &v
	return s
}

type UpdateComputeSubQuotaResponseBody struct {
	// The returned result.
	//
	// example:
	//
	// success
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// The error code.
	//
	// example:
	//
	// OBJECT_NOT_EXIST
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// this quota is not exist.
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// The HTTP status code.
	//
	// - 1xx: informational response. The request is received and is being processed.
	//
	// - 2xx: success. The request is successfully received, understood, and accepted by the server.
	//
	// - 3xx: redirection. The request is redirected, and further actions are required to complete the request.
	//
	// - 4xx: client error. The request contains invalid request parameters or syntaxes, or specific request conditions cannot be met.
	//
	// - 5xx: server error. The server cannot meet requirements due to other reasons.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0b57ff7616612271051086500ea3ce
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateComputeSubQuotaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateComputeSubQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateComputeSubQuotaResponseBody) SetData(v string) *UpdateComputeSubQuotaResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateComputeSubQuotaResponseBody) SetErrorCode(v string) *UpdateComputeSubQuotaResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateComputeSubQuotaResponseBody) SetErrorMsg(v string) *UpdateComputeSubQuotaResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *UpdateComputeSubQuotaResponseBody) SetHttpCode(v int32) *UpdateComputeSubQuotaResponseBody {
	s.HttpCode = &v
	return s
}

func (s *UpdateComputeSubQuotaResponseBody) SetRequestId(v string) *UpdateComputeSubQuotaResponseBody {
	s.RequestId = &v
	return s
}

type UpdateComputeSubQuotaResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateComputeSubQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateComputeSubQuotaResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateComputeSubQuotaResponse) GoString() string {
	return s.String()
}

func (s *UpdateComputeSubQuotaResponse) SetHeaders(v map[string]*string) *UpdateComputeSubQuotaResponse {
	s.Headers = v
	return s
}

func (s *UpdateComputeSubQuotaResponse) SetStatusCode(v int32) *UpdateComputeSubQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateComputeSubQuotaResponse) SetBody(v *UpdateComputeSubQuotaResponseBody) *UpdateComputeSubQuotaResponse {
	s.Body = v
	return s
}

type UpdateMmsDataSourceRequest struct {
	Action *string                `json:"action,omitempty" xml:"action,omitempty"`
	Config map[string]interface{} `json:"config,omitempty" xml:"config,omitempty"`
	Name   *string                `json:"name,omitempty" xml:"name,omitempty"`
	Test   *bool                  `json:"test,omitempty" xml:"test,omitempty"`
}

func (s UpdateMmsDataSourceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateMmsDataSourceRequest) GoString() string {
	return s.String()
}

func (s *UpdateMmsDataSourceRequest) SetAction(v string) *UpdateMmsDataSourceRequest {
	s.Action = &v
	return s
}

func (s *UpdateMmsDataSourceRequest) SetConfig(v map[string]interface{}) *UpdateMmsDataSourceRequest {
	s.Config = v
	return s
}

func (s *UpdateMmsDataSourceRequest) SetName(v string) *UpdateMmsDataSourceRequest {
	s.Name = &v
	return s
}

func (s *UpdateMmsDataSourceRequest) SetTest(v bool) *UpdateMmsDataSourceRequest {
	s.Test = &v
	return s
}

type UpdateMmsDataSourceResponseBody struct {
	Data *UpdateMmsDataSourceResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 76CE80C8-7392-5591-BCC8-610AFBF78ADF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateMmsDataSourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateMmsDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMmsDataSourceResponseBody) SetData(v *UpdateMmsDataSourceResponseBodyData) *UpdateMmsDataSourceResponseBody {
	s.Data = v
	return s
}

func (s *UpdateMmsDataSourceResponseBody) SetRequestId(v string) *UpdateMmsDataSourceResponseBody {
	s.RequestId = &v
	return s
}

type UpdateMmsDataSourceResponseBodyData struct {
	AsyncTaskId *int64 `json:"asyncTaskId,omitempty" xml:"asyncTaskId,omitempty"`
	SourceId    *int64 `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
}

func (s UpdateMmsDataSourceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateMmsDataSourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateMmsDataSourceResponseBodyData) SetAsyncTaskId(v int64) *UpdateMmsDataSourceResponseBodyData {
	s.AsyncTaskId = &v
	return s
}

func (s *UpdateMmsDataSourceResponseBodyData) SetSourceId(v int64) *UpdateMmsDataSourceResponseBodyData {
	s.SourceId = &v
	return s
}

type UpdateMmsDataSourceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMmsDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMmsDataSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateMmsDataSourceResponse) GoString() string {
	return s.String()
}

func (s *UpdateMmsDataSourceResponse) SetHeaders(v map[string]*string) *UpdateMmsDataSourceResponse {
	s.Headers = v
	return s
}

func (s *UpdateMmsDataSourceResponse) SetStatusCode(v int32) *UpdateMmsDataSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMmsDataSourceResponse) SetBody(v *UpdateMmsDataSourceResponseBody) *UpdateMmsDataSourceResponse {
	s.Body = v
	return s
}

type UpdatePackageRequest struct {
	// The request body parameters.
	//
	// example:
	//
	// {
	//
	//     "add": {
	//
	//         "allowedProjectList": [
	//
	//             {
	//
	//                 "label": "2",
	//
	//                 "project": "project_name"
	//
	//             }
	//
	//         ],
	//
	//         "resourceList": {
	//
	//             "table": [
	//
	//                 {
	//
	//                     "name": "table_name",
	//
	//                     "actions": [
	//
	//                         "Describe",
	//
	//                         "Select"
	//
	//                     ]
	//
	//                 },
	//
	//                 {
	//
	//                     "name": "table_name",
	//
	//                     "actions": [
	//
	//                         "Describe",
	//
	//                         "Select"
	//
	//                     ]
	//
	//                 }
	//
	//             ],
	//
	//             "resource": [
	//
	//                 {
	//
	//                     "name": "",
	//
	//                     "actions": []
	//
	//                 },
	//
	//                 {
	//
	//                     "name": "",
	//
	//                     "actions": []
	//
	//                 }
	//
	//             ],
	//
	//             "function": [
	//
	//                 {
	//
	//                     "name": "",
	//
	//                     "actions": []
	//
	//                 },
	//
	//                 {
	//
	//                     "name": "",
	//
	//                     "actions": []
	//
	//                 }
	//
	//             ]
	//
	//         }
	//
	//     },
	//
	//     "remove": {
	//
	//         "allowedProjectList": [
	//
	//             {
	//
	//                 "project": "project_name"
	//
	//             },
	//
	//             {
	//
	//                 "project": "project_2"
	//
	//             }
	//
	//         ],
	//
	//         "resourceList": {
	//
	//             "table": [
	//
	//                 {
	//
	//                     "name": "table_name"
	//
	//                 },
	//
	//                 {
	//
	//                     "name": "table_name"
	//
	//                 }
	//
	//             ],
	//
	//             "resource": [
	//
	//                 {
	//
	//                     "name": ""
	//
	//                 },
	//
	//                 {
	//
	//                     "name": ""
	//
	//                 }
	//
	//             ],
	//
	//             "function": [
	//
	//                 {
	//
	//                     "name": ""
	//
	//                 },
	//
	//                 {
	//
	//                     "name": ""
	//
	//                 }
	//
	//             ]
	//
	//         }
	//
	//     }
	//
	// }
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePackageRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePackageRequest) GoString() string {
	return s.String()
}

func (s *UpdatePackageRequest) SetBody(v string) *UpdatePackageRequest {
	s.Body = &v
	return s
}

type UpdatePackageResponseBody struct {
	// The returned data.
	//
	// example:
	//
	// success
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0bc1ec4016697018733156991e0888
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdatePackageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdatePackageResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePackageResponseBody) SetData(v string) *UpdatePackageResponseBody {
	s.Data = &v
	return s
}

func (s *UpdatePackageResponseBody) SetRequestId(v string) *UpdatePackageResponseBody {
	s.RequestId = &v
	return s
}

type UpdatePackageResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePackageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePackageResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdatePackageResponse) GoString() string {
	return s.String()
}

func (s *UpdatePackageResponse) SetHeaders(v map[string]*string) *UpdatePackageResponse {
	s.Headers = v
	return s
}

func (s *UpdatePackageResponse) SetStatusCode(v int32) *UpdatePackageResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePackageResponse) SetBody(v *UpdatePackageResponseBody) *UpdatePackageResponse {
	s.Body = v
	return s
}

type UpdateProjectBasicMetaRequest struct {
	// The project description.
	//
	// example:
	//
	// BI_Analysis
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	// The basic properties of the project.
	Properties *UpdateProjectBasicMetaRequestProperties `json:"properties,omitempty" xml:"properties,omitempty" type:"Struct"`
}

func (s UpdateProjectBasicMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectBasicMetaRequest) GoString() string {
	return s.String()
}

func (s *UpdateProjectBasicMetaRequest) SetComment(v string) *UpdateProjectBasicMetaRequest {
	s.Comment = &v
	return s
}

func (s *UpdateProjectBasicMetaRequest) SetProperties(v *UpdateProjectBasicMetaRequestProperties) *UpdateProjectBasicMetaRequest {
	s.Properties = v
	return s
}

type UpdateProjectBasicMetaRequestProperties struct {
	// Indicates whether a full table scan is allowed in the project. A full table scan occupies a large number of resources, which reduces data processing efficiency. By default, the full table scan feature is disabled.
	//
	// example:
	//
	// false
	AllowFullScan *bool `json:"allowFullScan,omitempty" xml:"allowFullScan,omitempty"`
	// Indicates whether the DECIMAL type of the MaxCompute V2.0 data type edition is enabled.
	//
	// example:
	//
	// true
	EnableDecimal2 *bool `json:"enableDecimal2,omitempty" xml:"enableDecimal2,omitempty"`
	// Indicates whether the routing of the Tunnel resource group is enabled.
	//
	// - true: The data transfer tasks that are submitted by the project by default use the Tunnel resource group that is bound to the project.
	//
	// - false: The data transfer tasks that are submitted by the project by default use the Tunnel shared resource group.
	//
	// example:
	//
	// true
	EnableTunnelQuotaRoute *bool `json:"enableTunnelQuotaRoute,omitempty" xml:"enableTunnelQuotaRoute,omitempty"`
	// The storage encryption properties.
	Encryption *UpdateProjectBasicMetaRequestPropertiesEncryption `json:"encryption,omitempty" xml:"encryption,omitempty" type:"Struct"`
	// The retention period for backup data. Unit: days. During the retention period, you can restore data of the version in use to the backup data of any version. Valid values: [0,30]. Default value: 1. The value 0 indicates that the backup feature is disabled.
	//
	// example:
	//
	// 1
	RetentionDays *int64 `json:"retentionDays,omitempty" xml:"retentionDays,omitempty"`
	// The maximum consumption threshold of a single SQL statement. Formula: Amount of scanned data (GB) × Complexity.
	//
	// example:
	//
	// 1500
	SqlMeteringMax *string `json:"sqlMeteringMax,omitempty" xml:"sqlMeteringMax,omitempty"`
	// The table lifecycle properties.
	TableLifecycle *UpdateProjectBasicMetaRequestPropertiesTableLifecycle `json:"tableLifecycle,omitempty" xml:"tableLifecycle,omitempty" type:"Struct"`
	// The time zone that is used by your project. The time zone is the same as the time zone specified by `odps.sql.timezone` .
	//
	// example:
	//
	// Asia/Shanghai
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
	// The <props="china">[Data Transmission Service](https://help.aliyun.com/zh/maxcompute/user-guide/overview-of-dts)
	//
	// <props="intl">[Data Transmission Service](https://www.alibabacloud.com/help/zh/maxcompute/user-guide/overview-of-dts) resource group that is bound to the project.
	//
	// - Default resource group: The Tunnel shared resource group is used. You cannot use the subscription-based Tunnel resource group for the project. The default resource group is automatically used by the Tunnel service of your project, regardless of the parameter setting.
	//
	// - Subscription-based Tunnel resource group: You can use the subscription-based Tunnel resource group for the project.
	//
	// example:
	//
	// Default
	TunnelQuota *string `json:"tunnelQuota,omitempty" xml:"tunnelQuota,omitempty"`
	// The data type edition. Valid values:
	//
	// - *1*: MaxCompute V1.0 data type edition
	//
	// - *2*: MaxCompute V2.0 data type edition
	//
	// - *hive*: Hive-compatible data type edition
	//
	// For more information about the differences among the three data type editions, see <props="china">[Data Type Versions](https://help.aliyun.com/zh/maxcompute/user-guide/data-type-editions)
	//
	// <props="intl">[Data Type Versions](https://www.alibabacloud.com/help/zh/maxcompute/user-guide/data-type-editions).
	//
	// example:
	//
	// 2.0
	TypeSystem *string `json:"typeSystem,omitempty" xml:"typeSystem,omitempty"`
}

func (s UpdateProjectBasicMetaRequestProperties) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectBasicMetaRequestProperties) GoString() string {
	return s.String()
}

func (s *UpdateProjectBasicMetaRequestProperties) SetAllowFullScan(v bool) *UpdateProjectBasicMetaRequestProperties {
	s.AllowFullScan = &v
	return s
}

func (s *UpdateProjectBasicMetaRequestProperties) SetEnableDecimal2(v bool) *UpdateProjectBasicMetaRequestProperties {
	s.EnableDecimal2 = &v
	return s
}

func (s *UpdateProjectBasicMetaRequestProperties) SetEnableTunnelQuotaRoute(v bool) *UpdateProjectBasicMetaRequestProperties {
	s.EnableTunnelQuotaRoute = &v
	return s
}

func (s *UpdateProjectBasicMetaRequestProperties) SetEncryption(v *UpdateProjectBasicMetaRequestPropertiesEncryption) *UpdateProjectBasicMetaRequestProperties {
	s.Encryption = v
	return s
}

func (s *UpdateProjectBasicMetaRequestProperties) SetRetentionDays(v int64) *UpdateProjectBasicMetaRequestProperties {
	s.RetentionDays = &v
	return s
}

func (s *UpdateProjectBasicMetaRequestProperties) SetSqlMeteringMax(v string) *UpdateProjectBasicMetaRequestProperties {
	s.SqlMeteringMax = &v
	return s
}

func (s *UpdateProjectBasicMetaRequestProperties) SetTableLifecycle(v *UpdateProjectBasicMetaRequestPropertiesTableLifecycle) *UpdateProjectBasicMetaRequestProperties {
	s.TableLifecycle = v
	return s
}

func (s *UpdateProjectBasicMetaRequestProperties) SetTimezone(v string) *UpdateProjectBasicMetaRequestProperties {
	s.Timezone = &v
	return s
}

func (s *UpdateProjectBasicMetaRequestProperties) SetTunnelQuota(v string) *UpdateProjectBasicMetaRequestProperties {
	s.TunnelQuota = &v
	return s
}

func (s *UpdateProjectBasicMetaRequestProperties) SetTypeSystem(v string) *UpdateProjectBasicMetaRequestProperties {
	s.TypeSystem = &v
	return s
}

type UpdateProjectBasicMetaRequestPropertiesEncryption struct {
	// The data encryption algorithm that is supported by the key. Valid values: AES256, AESCTR, and RC4.
	//
	// example:
	//
	// AES256
	Algorithm *string `json:"algorithm,omitempty" xml:"algorithm,omitempty"`
	// Indicates whether the data encryption feature needs to be enabled for the project. For more information about data encryption, see
	//
	// <props="china">[Storage Encryption](https://help.aliyun.com/zh/maxcompute/security-and-compliance/storage-encryption)
	//
	// <props="intl">[Storage Encryption](https://www.alibabacloud.com/help/zh/maxcompute/security-and-compliance/storage-encryption).
	//
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// The type of key that is used for data encryption. You can select MaxCompute Default Key or Bring Your Own Key (BYOK) as the key type. If you select MaxCompute Default Key, the default key that is created by MaxCompute is used.
	//
	// example:
	//
	// default
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
}

func (s UpdateProjectBasicMetaRequestPropertiesEncryption) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectBasicMetaRequestPropertiesEncryption) GoString() string {
	return s.String()
}

func (s *UpdateProjectBasicMetaRequestPropertiesEncryption) SetAlgorithm(v string) *UpdateProjectBasicMetaRequestPropertiesEncryption {
	s.Algorithm = &v
	return s
}

func (s *UpdateProjectBasicMetaRequestPropertiesEncryption) SetEnable(v bool) *UpdateProjectBasicMetaRequestPropertiesEncryption {
	s.Enable = &v
	return s
}

func (s *UpdateProjectBasicMetaRequestPropertiesEncryption) SetKey(v string) *UpdateProjectBasicMetaRequestPropertiesEncryption {
	s.Key = &v
	return s
}

type UpdateProjectBasicMetaRequestPropertiesTableLifecycle struct {
	// The lifecycle type. Valid values:
	//
	// - *mandatory*: The lifecycle clause is required in a table creation statement.
	//
	// - *optional*: The lifecycle clause is optional in a table creation statement. If you do not configure a lifecycle for a table, the table does not expire.
	//
	// - *inherit*: If you do not configure a lifecycle for a table when you create the table, the value of the odps.table.lifecycle.value parameter is used as the table lifecycle by default.
	//
	// example:
	//
	// optional
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The table lifecycle. Unit: days. Valid values: 1 to 37231. Default value: 37231.
	//
	// example:
	//
	// 37231
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s UpdateProjectBasicMetaRequestPropertiesTableLifecycle) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectBasicMetaRequestPropertiesTableLifecycle) GoString() string {
	return s.String()
}

func (s *UpdateProjectBasicMetaRequestPropertiesTableLifecycle) SetType(v string) *UpdateProjectBasicMetaRequestPropertiesTableLifecycle {
	s.Type = &v
	return s
}

func (s *UpdateProjectBasicMetaRequestPropertiesTableLifecycle) SetValue(v string) *UpdateProjectBasicMetaRequestPropertiesTableLifecycle {
	s.Value = &v
	return s
}

type UpdateProjectBasicMetaResponseBody struct {
	// The data returned.
	//
	// example:
	//
	// success
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// The error code.
	//
	// example:
	//
	// OBJECT_NOT_EXIST
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// This object does not exist.
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// The HTTP status code.
	//
	// - 1xx: informational response. The request is received and is being processed.
	//
	// - 2xx: success. The request is successfully received, understood, and accepted by the server.
	//
	// - 3xx: redirection. The request is redirected, and further actions are required to complete the request.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0b87b7e216652820458545253e8b0a
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateProjectBasicMetaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectBasicMetaResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProjectBasicMetaResponseBody) SetData(v string) *UpdateProjectBasicMetaResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateProjectBasicMetaResponseBody) SetErrorCode(v string) *UpdateProjectBasicMetaResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateProjectBasicMetaResponseBody) SetErrorMsg(v string) *UpdateProjectBasicMetaResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *UpdateProjectBasicMetaResponseBody) SetHttpCode(v int32) *UpdateProjectBasicMetaResponseBody {
	s.HttpCode = &v
	return s
}

func (s *UpdateProjectBasicMetaResponseBody) SetRequestId(v string) *UpdateProjectBasicMetaResponseBody {
	s.RequestId = &v
	return s
}

type UpdateProjectBasicMetaResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateProjectBasicMetaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateProjectBasicMetaResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectBasicMetaResponse) GoString() string {
	return s.String()
}

func (s *UpdateProjectBasicMetaResponse) SetHeaders(v map[string]*string) *UpdateProjectBasicMetaResponse {
	s.Headers = v
	return s
}

func (s *UpdateProjectBasicMetaResponse) SetStatusCode(v int32) *UpdateProjectBasicMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateProjectBasicMetaResponse) SetBody(v *UpdateProjectBasicMetaResponseBody) *UpdateProjectBasicMetaResponse {
	s.Body = v
	return s
}

type UpdateProjectDefaultQuotaRequest struct {
	// The default computing quota that is used to allocate computing resources, the jobs that are initiated by this project consume the computing resources in the default quota.
	//
	// example:
	//
	// os_PayAsYouGoQuota
	Quota *string `json:"quota,omitempty" xml:"quota,omitempty"`
}

func (s UpdateProjectDefaultQuotaRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectDefaultQuotaRequest) GoString() string {
	return s.String()
}

func (s *UpdateProjectDefaultQuotaRequest) SetQuota(v string) *UpdateProjectDefaultQuotaRequest {
	s.Quota = &v
	return s
}

type UpdateProjectDefaultQuotaResponseBody struct {
	// The data returned.
	//
	// example:
	//
	// success
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0a06dfe716674588654372173ec0da
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateProjectDefaultQuotaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectDefaultQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProjectDefaultQuotaResponseBody) SetData(v string) *UpdateProjectDefaultQuotaResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateProjectDefaultQuotaResponseBody) SetRequestId(v string) *UpdateProjectDefaultQuotaResponseBody {
	s.RequestId = &v
	return s
}

type UpdateProjectDefaultQuotaResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateProjectDefaultQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateProjectDefaultQuotaResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectDefaultQuotaResponse) GoString() string {
	return s.String()
}

func (s *UpdateProjectDefaultQuotaResponse) SetHeaders(v map[string]*string) *UpdateProjectDefaultQuotaResponse {
	s.Headers = v
	return s
}

func (s *UpdateProjectDefaultQuotaResponse) SetStatusCode(v int32) *UpdateProjectDefaultQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateProjectDefaultQuotaResponse) SetBody(v *UpdateProjectDefaultQuotaResponseBody) *UpdateProjectDefaultQuotaResponse {
	s.Body = v
	return s
}

type UpdateProjectIpWhiteListRequest struct {
	// The request body parameters.
	//
	// example:
	//
	// json {"ipWhiteList":{"ipList": "", // The IP address whitelists are of the STRING data type. Separate multiple IP address whitelists with commas (,). "vpcIpList": "", //} }
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateProjectIpWhiteListRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectIpWhiteListRequest) GoString() string {
	return s.String()
}

func (s *UpdateProjectIpWhiteListRequest) SetBody(v string) *UpdateProjectIpWhiteListRequest {
	s.Body = &v
	return s
}

type UpdateProjectIpWhiteListResponseBody struct {
	// The returned result.
	//
	// example:
	//
	// success
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0bc12e4316675560945192024e1044
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateProjectIpWhiteListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectIpWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProjectIpWhiteListResponseBody) SetData(v string) *UpdateProjectIpWhiteListResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateProjectIpWhiteListResponseBody) SetRequestId(v string) *UpdateProjectIpWhiteListResponseBody {
	s.RequestId = &v
	return s
}

type UpdateProjectIpWhiteListResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateProjectIpWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateProjectIpWhiteListResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectIpWhiteListResponse) GoString() string {
	return s.String()
}

func (s *UpdateProjectIpWhiteListResponse) SetHeaders(v map[string]*string) *UpdateProjectIpWhiteListResponse {
	s.Headers = v
	return s
}

func (s *UpdateProjectIpWhiteListResponse) SetStatusCode(v int32) *UpdateProjectIpWhiteListResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateProjectIpWhiteListResponse) SetBody(v *UpdateProjectIpWhiteListResponseBody) *UpdateProjectIpWhiteListResponse {
	s.Body = v
	return s
}

type UpdateQuotaPlanRequest struct {
	// The request body parameters.
	//
	// example:
	//
	// { "name": "planA", // The quota is a level-1 quota. You can select only the fields that are related to the quota plan. "quota": { "name": "a", "nickName": "aaa_nick", "tenantId": "10001", "regionId": "cn-hangzhou", "parentId": "0", "cluster": "AT-ODPS-TEST3", "parameter": { "minCU": 40, "maxCU": 40, "adhocCU": 0, "elasticMinCU": 40, "elasticMaxCU": 40, "enablePreemptiveScheduling": false, "forceReservedMin":true, "enablePriority":false, "singleJobCULimit":100, "adhocQuotaBeginTimeInSec": 1345, "adhocQuotaEndTimeInSec": 1234, "ignoreAdhocQuota":false }, "subQuotaInfoList": [ { "nickName": "WlmFuxiSecondaryOnlineQuotaTest", "name": "WlmFuxiSecondaryOnlineQuotaTest", "type": "FUXI_ONLINE", "tenantId": "10001", "regionId": "cn-hangzhou", "cluster": "AT-ODPS-TEST3", "parameter": { "minCU": 40, "maxCU": 40, "adhocCU": 0, "elasticMinCU": 40, "elasticMaxCU": 40, "enablePreemptiveScheduling": false, "forceReservedMin":true, "enablePriority":false, "singleJobCULimit":100, "adhocQuotaBeginTimeInSec": 1345, "adhocQuotaEndTimeInSec": 1234, "ignoreAdhocQuota":false } } ] } }
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The ID of the tenant.
	//
	// example:
	//
	// 478403690625249
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s UpdateQuotaPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateQuotaPlanRequest) GoString() string {
	return s.String()
}

func (s *UpdateQuotaPlanRequest) SetBody(v string) *UpdateQuotaPlanRequest {
	s.Body = &v
	return s
}

func (s *UpdateQuotaPlanRequest) SetRegion(v string) *UpdateQuotaPlanRequest {
	s.Region = &v
	return s
}

func (s *UpdateQuotaPlanRequest) SetTenantId(v string) *UpdateQuotaPlanRequest {
	s.TenantId = &v
	return s
}

type UpdateQuotaPlanResponseBody struct {
	// The returned result.
	//
	// example:
	//
	// success
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0a06dfe516688379832875789e2c65
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateQuotaPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateQuotaPlanResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateQuotaPlanResponseBody) SetData(v string) *UpdateQuotaPlanResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateQuotaPlanResponseBody) SetRequestId(v string) *UpdateQuotaPlanResponseBody {
	s.RequestId = &v
	return s
}

type UpdateQuotaPlanResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateQuotaPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateQuotaPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateQuotaPlanResponse) GoString() string {
	return s.String()
}

func (s *UpdateQuotaPlanResponse) SetHeaders(v map[string]*string) *UpdateQuotaPlanResponse {
	s.Headers = v
	return s
}

func (s *UpdateQuotaPlanResponse) SetStatusCode(v int32) *UpdateQuotaPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateQuotaPlanResponse) SetBody(v *UpdateQuotaPlanResponseBody) *UpdateQuotaPlanResponse {
	s.Body = v
	return s
}

type UpdateQuotaScheduleRequest struct {
	// The request body parameters.
	//
	// example:
	//
	// \\# The quota plan immediately takes effect. [ { "type": "once", "plan": "planA", "operator":"userA" } ] # The quota plan is scheduled on a regular basis. [ { "id": "etl_time", "type": "daily", "condition": { "at": "0800", "after": "2022-04-25T04:23:04Z" // optional }, "plan": "planA" }, { "id": "bi", "type": "daily", "condition": { "at": "0900", "after": "2022-04-25T04:23:04Z" // optional }, "plan": "planB" }, ]
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-chengdu
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The ID of the tenant.
	//
	// example:
	//
	// 478403690625249
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s UpdateQuotaScheduleRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateQuotaScheduleRequest) GoString() string {
	return s.String()
}

func (s *UpdateQuotaScheduleRequest) SetBody(v string) *UpdateQuotaScheduleRequest {
	s.Body = &v
	return s
}

func (s *UpdateQuotaScheduleRequest) SetRegion(v string) *UpdateQuotaScheduleRequest {
	s.Region = &v
	return s
}

func (s *UpdateQuotaScheduleRequest) SetTenantId(v string) *UpdateQuotaScheduleRequest {
	s.TenantId = &v
	return s
}

type UpdateQuotaScheduleResponseBody struct {
	// The returned result.
	//
	// example:
	//
	// success
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0a06dfe516691014920015940e1c9d
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateQuotaScheduleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateQuotaScheduleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateQuotaScheduleResponseBody) SetData(v string) *UpdateQuotaScheduleResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateQuotaScheduleResponseBody) SetRequestId(v string) *UpdateQuotaScheduleResponseBody {
	s.RequestId = &v
	return s
}

type UpdateQuotaScheduleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateQuotaScheduleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateQuotaScheduleResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateQuotaScheduleResponse) GoString() string {
	return s.String()
}

func (s *UpdateQuotaScheduleResponse) SetHeaders(v map[string]*string) *UpdateQuotaScheduleResponse {
	s.Headers = v
	return s
}

func (s *UpdateQuotaScheduleResponse) SetStatusCode(v int32) *UpdateQuotaScheduleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateQuotaScheduleResponse) SetBody(v *UpdateQuotaScheduleResponseBody) *UpdateQuotaScheduleResponse {
	s.Body = v
	return s
}

type UpdateTunnelQuotaTimerRequest struct {
	// The request body.
	Body     []*UpdateTunnelQuotaTimerRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	Timezone *string                              `json:"timezone,omitempty" xml:"timezone,omitempty"`
}

func (s UpdateTunnelQuotaTimerRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTunnelQuotaTimerRequest) GoString() string {
	return s.String()
}

func (s *UpdateTunnelQuotaTimerRequest) SetBody(v []*UpdateTunnelQuotaTimerRequestBody) *UpdateTunnelQuotaTimerRequest {
	s.Body = v
	return s
}

func (s *UpdateTunnelQuotaTimerRequest) SetTimezone(v string) *UpdateTunnelQuotaTimerRequest {
	s.Timezone = &v
	return s
}

type UpdateTunnelQuotaTimerRequestBody struct {
	// The start time of the time-specific configuration.
	//
	// example:
	//
	// 00:00
	BeginTime *string `json:"beginTime,omitempty" xml:"beginTime,omitempty"`
	// The end time of the time-specific configuration.
	//
	// example:
	//
	// 08:00
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The parameters for the time-specific configuration.
	TunnelQuotaParameter *UpdateTunnelQuotaTimerRequestBodyTunnelQuotaParameter `json:"tunnelQuotaParameter,omitempty" xml:"tunnelQuotaParameter,omitempty" type:"Struct"`
}

func (s UpdateTunnelQuotaTimerRequestBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTunnelQuotaTimerRequestBody) GoString() string {
	return s.String()
}

func (s *UpdateTunnelQuotaTimerRequestBody) SetBeginTime(v string) *UpdateTunnelQuotaTimerRequestBody {
	s.BeginTime = &v
	return s
}

func (s *UpdateTunnelQuotaTimerRequestBody) SetEndTime(v string) *UpdateTunnelQuotaTimerRequestBody {
	s.EndTime = &v
	return s
}

func (s *UpdateTunnelQuotaTimerRequestBody) SetTunnelQuotaParameter(v *UpdateTunnelQuotaTimerRequestBodyTunnelQuotaParameter) *UpdateTunnelQuotaTimerRequestBody {
	s.TunnelQuotaParameter = v
	return s
}

type UpdateTunnelQuotaTimerRequestBodyTunnelQuotaParameter struct {
	// The number of elastically reserved slots.
	//
	// example:
	//
	// 100
	ElasticReservedSlotNum *int64 `json:"elasticReservedSlotNum,omitempty" xml:"elasticReservedSlotNum,omitempty"`
	// The number of reserved slots.
	//
	// example:
	//
	// 100
	SlotNum *int64 `json:"slotNum,omitempty" xml:"slotNum,omitempty"`
}

func (s UpdateTunnelQuotaTimerRequestBodyTunnelQuotaParameter) String() string {
	return tea.Prettify(s)
}

func (s UpdateTunnelQuotaTimerRequestBodyTunnelQuotaParameter) GoString() string {
	return s.String()
}

func (s *UpdateTunnelQuotaTimerRequestBodyTunnelQuotaParameter) SetElasticReservedSlotNum(v int64) *UpdateTunnelQuotaTimerRequestBodyTunnelQuotaParameter {
	s.ElasticReservedSlotNum = &v
	return s
}

func (s *UpdateTunnelQuotaTimerRequestBodyTunnelQuotaParameter) SetSlotNum(v int64) *UpdateTunnelQuotaTimerRequestBodyTunnelQuotaParameter {
	s.SlotNum = &v
	return s
}

type UpdateTunnelQuotaTimerResponseBody struct {
	// The data returned.
	//
	// example:
	//
	// success
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// The error code.
	//
	// example:
	//
	// OBJECT_NOT_EXIST
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// This object does not exist.
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// The HTTP status code.
	//
	// 	- 1xx: informational response. The request is received and is being processed.
	//
	// 	- 2xx: success. The request is successfully received, understood, and accepted by the server.
	//
	// 	- 3xx: redirection. The request is redirected, and further actions are required to complete the request.
	//
	// 	- 4xx: client error. The request contains invalid request parameters or syntaxes, or specific request conditions cannot be met.
	//
	// 	- 5xx: server error. The server cannot meet requirements due to other reasons.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0bc12e4316675560945192024e1044
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateTunnelQuotaTimerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTunnelQuotaTimerResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTunnelQuotaTimerResponseBody) SetData(v string) *UpdateTunnelQuotaTimerResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateTunnelQuotaTimerResponseBody) SetErrorCode(v string) *UpdateTunnelQuotaTimerResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateTunnelQuotaTimerResponseBody) SetErrorMsg(v string) *UpdateTunnelQuotaTimerResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *UpdateTunnelQuotaTimerResponseBody) SetHttpCode(v int32) *UpdateTunnelQuotaTimerResponseBody {
	s.HttpCode = &v
	return s
}

func (s *UpdateTunnelQuotaTimerResponseBody) SetRequestId(v string) *UpdateTunnelQuotaTimerResponseBody {
	s.RequestId = &v
	return s
}

type UpdateTunnelQuotaTimerResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTunnelQuotaTimerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTunnelQuotaTimerResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTunnelQuotaTimerResponse) GoString() string {
	return s.String()
}

func (s *UpdateTunnelQuotaTimerResponse) SetHeaders(v map[string]*string) *UpdateTunnelQuotaTimerResponse {
	s.Headers = v
	return s
}

func (s *UpdateTunnelQuotaTimerResponse) SetStatusCode(v int32) *UpdateTunnelQuotaTimerResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTunnelQuotaTimerResponse) SetBody(v *UpdateTunnelQuotaTimerResponseBody) *UpdateTunnelQuotaTimerResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"ap-northeast-1":              tea.String("maxcompute.aliyuncs.com"),
		"ap-northeast-2-pop":          tea.String("maxcompute.aliyuncs.com"),
		"ap-south-1":                  tea.String("maxcompute.aliyuncs.com"),
		"ap-southeast-1":              tea.String("maxcompute.aliyuncs.com"),
		"ap-southeast-2":              tea.String("maxcompute.aliyuncs.com"),
		"ap-southeast-3":              tea.String("maxcompute.aliyuncs.com"),
		"ap-southeast-5":              tea.String("maxcompute.aliyuncs.com"),
		"cn-beijing":                  tea.String("maxcompute.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("maxcompute.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("maxcompute.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("maxcompute.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("maxcompute.aliyuncs.com"),
		"cn-chengdu":                  tea.String("maxcompute.aliyuncs.com"),
		"cn-edge-1":                   tea.String("maxcompute.aliyuncs.com"),
		"cn-fujian":                   tea.String("maxcompute.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("maxcompute.aliyuncs.com"),
		"cn-hangzhou":                 tea.String("maxcompute.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("maxcompute.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("maxcompute.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("maxcompute.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("maxcompute.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("maxcompute.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("maxcompute.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("maxcompute.aliyuncs.com"),
		"cn-hongkong":                 tea.String("maxcompute.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("maxcompute.aliyuncs.com"),
		"cn-huhehaote":                tea.String("maxcompute.aliyuncs.com"),
		"cn-north-2-gov-1":            tea.String("maxcompute.aliyuncs.com"),
		"cn-qingdao":                  tea.String("maxcompute.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("maxcompute.aliyuncs.com"),
		"cn-shanghai":                 tea.String("maxcompute.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("maxcompute.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("maxcompute.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("maxcompute.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("maxcompute.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("maxcompute.aliyuncs.com"),
		"cn-shenzhen":                 tea.String("maxcompute.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("maxcompute.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("maxcompute.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("maxcompute.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("maxcompute.aliyuncs.com"),
		"cn-wuhan":                    tea.String("maxcompute.aliyuncs.com"),
		"cn-yushanfang":               tea.String("maxcompute.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("maxcompute.aliyuncs.com"),
		"cn-zhangjiakou":              tea.String("maxcompute.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("maxcompute.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("maxcompute.aliyuncs.com"),
		"eu-central-1":                tea.String("maxcompute.aliyuncs.com"),
		"eu-west-1":                   tea.String("maxcompute.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("maxcompute.aliyuncs.com"),
		"me-east-1":                   tea.String("maxcompute.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("maxcompute.aliyuncs.com"),
		"us-east-1":                   tea.String("maxcompute.aliyuncs.com"),
		"us-west-1":                   tea.String("maxcompute.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("maxcompute"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Activate a Quota Plan Immediately.
//
// Description:
//
// Please ensure that before using this interface, you have fully understood the <props="china">[Pricing and Charges](https://help.aliyun.com/zh/maxcompute/product-overview/computing-pricing-1)
//
// <props="intl">[Pricing and Charges](https://www.alibabacloud.com/help/maxcompute/product-overview/computing-pricing-1) of MaxCompute Elastic Reserved CU.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyComputeQuotaPlanResponse
func (client *Client) ApplyComputeQuotaPlanWithOptions(nickname *string, planName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ApplyComputeQuotaPlanResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ApplyComputeQuotaPlan"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/quotas/" + tea.StringValue(openapiutil.GetEncodeParam(nickname)) + "/computeQuotaPlan/" + tea.StringValue(openapiutil.GetEncodeParam(planName)) + "/apply"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ApplyComputeQuotaPlanResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ApplyComputeQuotaPlanResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Activate a Quota Plan Immediately.
//
// Description:
//
// Please ensure that before using this interface, you have fully understood the <props="china">[Pricing and Charges](https://help.aliyun.com/zh/maxcompute/product-overview/computing-pricing-1)
//
// <props="intl">[Pricing and Charges](https://www.alibabacloud.com/help/maxcompute/product-overview/computing-pricing-1) of MaxCompute Elastic Reserved CU.
//
// @return ApplyComputeQuotaPlanResponse
func (client *Client) ApplyComputeQuotaPlan(nickname *string, planName *string) (_result *ApplyComputeQuotaPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ApplyComputeQuotaPlanResponse{}
	_body, _err := client.ApplyComputeQuotaPlanWithOptions(nickname, planName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Create MaxCompute ComputeQuotaPlan.
//
// Description:
//
// Please ensure that before using this interface, you have fully understood the <props="china">[Pricing and Charges](https://help.aliyun.com/zh/maxcompute/product-overview/computing-pricing-1)
//
// <props="intl">[Pricing and Charges](https://www.alibabacloud.com/help/maxcompute/product-overview/computing-pricing-1) of MaxCompute Elastic Reserved CU.
//
// @param request - CreateComputeQuotaPlanRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateComputeQuotaPlanResponse
func (client *Client) CreateComputeQuotaPlanWithOptions(nickname *string, request *CreateComputeQuotaPlanRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateComputeQuotaPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Quota)) {
		body["quota"] = request.Quota
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateComputeQuotaPlan"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/quotas/" + tea.StringValue(openapiutil.GetEncodeParam(nickname)) + "/computeQuotaPlan"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateComputeQuotaPlanResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateComputeQuotaPlanResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Create MaxCompute ComputeQuotaPlan.
//
// Description:
//
// Please ensure that before using this interface, you have fully understood the <props="china">[Pricing and Charges](https://help.aliyun.com/zh/maxcompute/product-overview/computing-pricing-1)
//
// <props="intl">[Pricing and Charges](https://www.alibabacloud.com/help/maxcompute/product-overview/computing-pricing-1) of MaxCompute Elastic Reserved CU.
//
// @param request - CreateComputeQuotaPlanRequest
//
// @return CreateComputeQuotaPlanResponse
func (client *Client) CreateComputeQuotaPlan(nickname *string, request *CreateComputeQuotaPlanRequest) (_result *CreateComputeQuotaPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateComputeQuotaPlanResponse{}
	_body, _err := client.CreateComputeQuotaPlanWithOptions(nickname, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateMmsDataSourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMmsDataSourceResponse
func (client *Client) CreateMmsDataSourceWithOptions(request *CreateMmsDataSourceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateMmsDataSourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Config)) {
		body["config"] = request.Config
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Networklink)) {
		body["networklink"] = request.Networklink
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateMmsDataSource"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/mms/datasources"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateMmsDataSourceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateMmsDataSourceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - CreateMmsDataSourceRequest
//
// @return CreateMmsDataSourceResponse
func (client *Client) CreateMmsDataSource(request *CreateMmsDataSourceRequest) (_result *CreateMmsDataSourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateMmsDataSourceResponse{}
	_body, _err := client.CreateMmsDataSourceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建数据源的更新元数据操作
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMmsFetchMetadataJobResponse
func (client *Client) CreateMmsFetchMetadataJobWithOptions(sourceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateMmsFetchMetadataJobResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("CreateMmsFetchMetadataJob"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/mms/datasources/" + tea.StringValue(openapiutil.GetEncodeParam(sourceId)) + "/scans"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateMmsFetchMetadataJobResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateMmsFetchMetadataJobResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 创建数据源的更新元数据操作
//
// @return CreateMmsFetchMetadataJobResponse
func (client *Client) CreateMmsFetchMetadataJob(sourceId *string) (_result *CreateMmsFetchMetadataJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateMmsFetchMetadataJobResponse{}
	_body, _err := client.CreateMmsFetchMetadataJobWithOptions(sourceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建迁移任务
//
// @param request - CreateMmsJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMmsJobResponse
func (client *Client) CreateMmsJobWithOptions(sourceId *string, request *CreateMmsJobRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateMmsJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ColumnMapping)) {
		body["columnMapping"] = request.ColumnMapping
	}

	if !tea.BoolValue(util.IsUnset(request.DstDbName)) {
		body["dstDbName"] = request.DstDbName
	}

	if !tea.BoolValue(util.IsUnset(request.DstSchemaName)) {
		body["dstSchemaName"] = request.DstSchemaName
	}

	if !tea.BoolValue(util.IsUnset(request.EnableVerification)) {
		body["enableVerification"] = request.EnableVerification
	}

	if !tea.BoolValue(util.IsUnset(request.Eta)) {
		body["eta"] = request.Eta
	}

	if !tea.BoolValue(util.IsUnset(request.Increment)) {
		body["increment"] = request.Increment
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Others)) {
		body["others"] = request.Others
	}

	if !tea.BoolValue(util.IsUnset(request.PartitionFilters)) {
		body["partitionFilters"] = request.PartitionFilters
	}

	if !tea.BoolValue(util.IsUnset(request.Partitions)) {
		body["partitions"] = request.Partitions
	}

	if !tea.BoolValue(util.IsUnset(request.SchemaOnly)) {
		body["schemaOnly"] = request.SchemaOnly
	}

	if !tea.BoolValue(util.IsUnset(request.SourceId)) {
		body["sourceId"] = request.SourceId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceName)) {
		body["sourceName"] = request.SourceName
	}

	if !tea.BoolValue(util.IsUnset(request.SrcDbName)) {
		body["srcDbName"] = request.SrcDbName
	}

	if !tea.BoolValue(util.IsUnset(request.SrcSchemaName)) {
		body["srcSchemaName"] = request.SrcSchemaName
	}

	if !tea.BoolValue(util.IsUnset(request.TableBlackList)) {
		body["tableBlackList"] = request.TableBlackList
	}

	if !tea.BoolValue(util.IsUnset(request.TableMapping)) {
		body["tableMapping"] = request.TableMapping
	}

	if !tea.BoolValue(util.IsUnset(request.TableWhiteList)) {
		body["tableWhiteList"] = request.TableWhiteList
	}

	if !tea.BoolValue(util.IsUnset(request.Tables)) {
		body["tables"] = request.Tables
	}

	if !tea.BoolValue(util.IsUnset(request.TaskType)) {
		body["taskType"] = request.TaskType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateMmsJob"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/mms/datasources/" + tea.StringValue(openapiutil.GetEncodeParam(sourceId)) + "/jobs"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateMmsJobResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateMmsJobResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 创建迁移任务
//
// @param request - CreateMmsJobRequest
//
// @return CreateMmsJobResponse
func (client *Client) CreateMmsJob(sourceId *string, request *CreateMmsJobRequest) (_result *CreateMmsJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateMmsJobResponse{}
	_body, _err := client.CreateMmsJobWithOptions(sourceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a package.
//
// @param request - CreatePackageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePackageResponse
func (client *Client) CreatePackageWithOptions(projectName *string, request *CreatePackageRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreatePackageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsInstall)) {
		query["isInstall"] = request.IsInstall
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("CreatePackage"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/projects/" + tea.StringValue(openapiutil.GetEncodeParam(projectName)) + "/packages"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreatePackageResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreatePackageResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Creates a package.
//
// @param request - CreatePackageRequest
//
// @return CreatePackageResponse
func (client *Client) CreatePackage(projectName *string, request *CreatePackageRequest) (_result *CreatePackageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreatePackageResponse{}
	_body, _err := client.CreatePackageWithOptions(projectName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a MaxCompute project.
//
// @param request - CreateProjectRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateProjectResponse
func (client *Client) CreateProjectWithOptions(request *CreateProjectRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("CreateProject"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/projects"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateProjectResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateProjectResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Creates a MaxCompute project.
//
// @param request - CreateProjectRequest
//
// @return CreateProjectResponse
func (client *Client) CreateProject(request *CreateProjectRequest) (_result *CreateProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateProjectResponse{}
	_body, _err := client.CreateProjectWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a quota plan.
//
// @param request - CreateQuotaPlanRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateQuotaPlanResponse
func (client *Client) CreateQuotaPlanWithOptions(nickname *string, request *CreateQuotaPlanRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateQuotaPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		query["tenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("CreateQuotaPlan"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/quotas/" + tea.StringValue(openapiutil.GetEncodeParam(nickname)) + "/plans"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateQuotaPlanResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateQuotaPlanResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Creates a quota plan.
//
// @param request - CreateQuotaPlanRequest
//
// @return CreateQuotaPlanResponse
func (client *Client) CreateQuotaPlan(nickname *string, request *CreateQuotaPlanRequest) (_result *CreateQuotaPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateQuotaPlanResponse{}
	_body, _err := client.CreateQuotaPlanWithOptions(nickname, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a role at the MaxCompute project level.
//
// @param request - CreateRoleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRoleResponse
func (client *Client) CreateRoleWithOptions(projectName *string, request *CreateRoleRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRole"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/projects/" + tea.StringValue(openapiutil.GetEncodeParam(projectName)) + "/roles"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateRoleResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateRoleResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Creates a role at the MaxCompute project level.
//
// @param request - CreateRoleRequest
//
// @return CreateRoleResponse
func (client *Client) CreateRole(projectName *string, request *CreateRoleRequest) (_result *CreateRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateRoleResponse{}
	_body, _err := client.CreateRoleWithOptions(projectName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Delete MaxCompute compute quota plan.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteComputeQuotaPlanResponse
func (client *Client) DeleteComputeQuotaPlanWithOptions(nickname *string, planName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteComputeQuotaPlanResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteComputeQuotaPlan"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/quotas/" + tea.StringValue(openapiutil.GetEncodeParam(nickname)) + "/computeQuotaPlan/" + tea.StringValue(openapiutil.GetEncodeParam(planName))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteComputeQuotaPlanResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteComputeQuotaPlanResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Delete MaxCompute compute quota plan.
//
// @return DeleteComputeQuotaPlanResponse
func (client *Client) DeleteComputeQuotaPlan(nickname *string, planName *string) (_result *DeleteComputeQuotaPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteComputeQuotaPlanResponse{}
	_body, _err := client.DeleteComputeQuotaPlanWithOptions(nickname, planName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除数据源
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMmsDataSourceResponse
func (client *Client) DeleteMmsDataSourceWithOptions(sourceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteMmsDataSourceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteMmsDataSource"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/mms/datasources/" + tea.StringValue(openapiutil.GetEncodeParam(sourceId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteMmsDataSourceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteMmsDataSourceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 删除数据源
//
// @return DeleteMmsDataSourceResponse
func (client *Client) DeleteMmsDataSource(sourceId *string) (_result *DeleteMmsDataSourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteMmsDataSourceResponse{}
	_body, _err := client.DeleteMmsDataSourceWithOptions(sourceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMmsJobResponse
func (client *Client) DeleteMmsJobWithOptions(sourceId *string, jobId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteMmsJobResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteMmsJob"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/mms/datasources/" + tea.StringValue(openapiutil.GetEncodeParam(sourceId)) + "/jobs/" + tea.StringValue(openapiutil.GetEncodeParam(jobId))),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteMmsJobResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteMmsJobResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @return DeleteMmsJobResponse
func (client *Client) DeleteMmsJob(sourceId *string, jobId *string) (_result *DeleteMmsJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteMmsJobResponse{}
	_body, _err := client.DeleteMmsJobWithOptions(sourceId, jobId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a quota plan.
//
// @param request - DeleteQuotaPlanRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteQuotaPlanResponse
func (client *Client) DeleteQuotaPlanWithOptions(nickname *string, planName *string, request *DeleteQuotaPlanRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteQuotaPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		query["tenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteQuotaPlan"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/quotas/" + tea.StringValue(openapiutil.GetEncodeParam(nickname)) + "/plans/" + tea.StringValue(openapiutil.GetEncodeParam(planName))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteQuotaPlanResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteQuotaPlanResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Deletes a quota plan.
//
// @param request - DeleteQuotaPlanRequest
//
// @return DeleteQuotaPlanResponse
func (client *Client) DeleteQuotaPlan(nickname *string, planName *string, request *DeleteQuotaPlanRequest) (_result *DeleteQuotaPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteQuotaPlanResponse{}
	_body, _err := client.DeleteQuotaPlanWithOptions(nickname, planName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// GetComputeEffectivePlan.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetComputeEffectivePlanResponse
func (client *Client) GetComputeEffectivePlanWithOptions(nickname *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetComputeEffectivePlanResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetComputeEffectivePlan"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/quotas/" + tea.StringValue(openapiutil.GetEncodeParam(nickname)) + "/computeEffectivePlan/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetComputeEffectivePlanResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetComputeEffectivePlanResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// GetComputeEffectivePlan.
//
// @return GetComputeEffectivePlanResponse
func (client *Client) GetComputeEffectivePlan(nickname *string) (_result *GetComputeEffectivePlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetComputeEffectivePlanResponse{}
	_body, _err := client.GetComputeEffectivePlanWithOptions(nickname, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Get detailed information of a single compute quota plan.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetComputeQuotaPlanResponse
func (client *Client) GetComputeQuotaPlanWithOptions(nickname *string, planName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetComputeQuotaPlanResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetComputeQuotaPlan"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/quotas/" + tea.StringValue(openapiutil.GetEncodeParam(nickname)) + "/computeQuotaPlan/" + tea.StringValue(openapiutil.GetEncodeParam(planName))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetComputeQuotaPlanResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetComputeQuotaPlanResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Get detailed information of a single compute quota plan.
//
// @return GetComputeQuotaPlanResponse
func (client *Client) GetComputeQuotaPlan(nickname *string, planName *string) (_result *GetComputeQuotaPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetComputeQuotaPlanResponse{}
	_body, _err := client.GetComputeQuotaPlanWithOptions(nickname, planName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Displays the time-specific configuration of compute quota.
//
// @param request - GetComputeQuotaScheduleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetComputeQuotaScheduleResponse
func (client *Client) GetComputeQuotaScheduleWithOptions(nickname *string, request *GetComputeQuotaScheduleRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetComputeQuotaScheduleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DisplayTimezone)) {
		query["displayTimezone"] = request.DisplayTimezone
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetComputeQuotaSchedule"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/quotas/" + tea.StringValue(openapiutil.GetEncodeParam(nickname)) + "/computeQuotaSchedule"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetComputeQuotaScheduleResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetComputeQuotaScheduleResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Displays the time-specific configuration of compute quota.
//
// @param request - GetComputeQuotaScheduleRequest
//
// @return GetComputeQuotaScheduleResponse
func (client *Client) GetComputeQuotaSchedule(nickname *string, request *GetComputeQuotaScheduleRequest) (_result *GetComputeQuotaScheduleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetComputeQuotaScheduleResponse{}
	_body, _err := client.GetComputeQuotaScheduleWithOptions(nickname, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Performs statistics on all jobs that are complete on a specified day and obtains the total resource usage of each job executor on a daily basis.
//
// @param tmpReq - GetJobResourceUsageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJobResourceUsageResponse
func (client *Client) GetJobResourceUsageWithOptions(tmpReq *GetJobResourceUsageRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetJobResourceUsageResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetJobResourceUsageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.JobOwnerList)) {
		request.JobOwnerListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.JobOwnerList, tea.String("jobOwnerList"), tea.String("simple"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.QuotaNicknameList)) {
		request.QuotaNicknameListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.QuotaNicknameList, tea.String("quotaNicknameList"), tea.String("simple"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Date)) {
		query["date"] = request.Date
	}

	if !tea.BoolValue(util.IsUnset(request.JobOwnerListShrink)) {
		query["jobOwnerList"] = request.JobOwnerListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.QuotaNicknameListShrink)) {
		query["quotaNicknameList"] = request.QuotaNicknameListShrink
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetJobResourceUsage"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/jobs/resourceUsage"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetJobResourceUsageResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetJobResourceUsageResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Performs statistics on all jobs that are complete on a specified day and obtains the total resource usage of each job executor on a daily basis.
//
// @param request - GetJobResourceUsageRequest
//
// @return GetJobResourceUsageResponse
func (client *Client) GetJobResourceUsage(request *GetJobResourceUsageRequest) (_result *GetJobResourceUsageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetJobResourceUsageResponse{}
	_body, _err := client.GetJobResourceUsageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMmsAsyncTaskResponse
func (client *Client) GetMmsAsyncTaskWithOptions(sourceId *string, asyncTaskId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetMmsAsyncTaskResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetMmsAsyncTask"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/mms/datasources/" + tea.StringValue(openapiutil.GetEncodeParam(sourceId)) + "/asyncTasks/" + tea.StringValue(openapiutil.GetEncodeParam(asyncTaskId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetMmsAsyncTaskResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetMmsAsyncTaskResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @return GetMmsAsyncTaskResponse
func (client *Client) GetMmsAsyncTask(sourceId *string, asyncTaskId *string) (_result *GetMmsAsyncTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetMmsAsyncTaskResponse{}
	_body, _err := client.GetMmsAsyncTaskWithOptions(sourceId, asyncTaskId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetMmsDataSourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMmsDataSourceResponse
func (client *Client) GetMmsDataSourceWithOptions(sourceId *string, request *GetMmsDataSourceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetMmsDataSourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.WithConfig)) {
		query["withConfig"] = request.WithConfig
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetMmsDataSource"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/mms/datasources/" + tea.StringValue(openapiutil.GetEncodeParam(sourceId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetMmsDataSourceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetMmsDataSourceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - GetMmsDataSourceRequest
//
// @return GetMmsDataSourceResponse
func (client *Client) GetMmsDataSource(sourceId *string, request *GetMmsDataSourceRequest) (_result *GetMmsDataSourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetMmsDataSourceResponse{}
	_body, _err := client.GetMmsDataSourceWithOptions(sourceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMmsDbResponse
func (client *Client) GetMmsDbWithOptions(sourceId *string, dbId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetMmsDbResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetMmsDb"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/mms/datasources/" + tea.StringValue(openapiutil.GetEncodeParam(sourceId)) + "/dbs/" + tea.StringValue(openapiutil.GetEncodeParam(dbId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetMmsDbResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetMmsDbResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @return GetMmsDbResponse
func (client *Client) GetMmsDb(sourceId *string, dbId *string) (_result *GetMmsDbResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetMmsDbResponse{}
	_body, _err := client.GetMmsDbWithOptions(sourceId, dbId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMmsFetchMetadataJobResponse
func (client *Client) GetMmsFetchMetadataJobWithOptions(sourceId *string, scanId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetMmsFetchMetadataJobResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetMmsFetchMetadataJob"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/mms/datasources/" + tea.StringValue(openapiutil.GetEncodeParam(sourceId)) + "/scans/" + tea.StringValue(openapiutil.GetEncodeParam(scanId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetMmsFetchMetadataJobResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetMmsFetchMetadataJobResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @return GetMmsFetchMetadataJobResponse
func (client *Client) GetMmsFetchMetadataJob(sourceId *string, scanId *string) (_result *GetMmsFetchMetadataJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetMmsFetchMetadataJobResponse{}
	_body, _err := client.GetMmsFetchMetadataJobWithOptions(sourceId, scanId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMmsJobResponse
func (client *Client) GetMmsJobWithOptions(sourceId *string, jobId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetMmsJobResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetMmsJob"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/mms/datasources/" + tea.StringValue(openapiutil.GetEncodeParam(sourceId)) + "/jobs/" + tea.StringValue(openapiutil.GetEncodeParam(jobId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetMmsJobResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetMmsJobResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @return GetMmsJobResponse
func (client *Client) GetMmsJob(sourceId *string, jobId *string) (_result *GetMmsJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetMmsJobResponse{}
	_body, _err := client.GetMmsJobWithOptions(sourceId, jobId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMmsPartitionResponse
func (client *Client) GetMmsPartitionWithOptions(sourceId *string, partitionId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetMmsPartitionResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetMmsPartition"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/mms/datasources/" + tea.StringValue(openapiutil.GetEncodeParam(sourceId)) + "/partitions/" + tea.StringValue(openapiutil.GetEncodeParam(partitionId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetMmsPartitionResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetMmsPartitionResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @return GetMmsPartitionResponse
func (client *Client) GetMmsPartition(sourceId *string, partitionId *string) (_result *GetMmsPartitionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetMmsPartitionResponse{}
	_body, _err := client.GetMmsPartitionWithOptions(sourceId, partitionId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMmsTableResponse
func (client *Client) GetMmsTableWithOptions(sourceId *string, tableId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetMmsTableResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetMmsTable"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/mms/datasources/" + tea.StringValue(openapiutil.GetEncodeParam(sourceId)) + "/tables/" + tea.StringValue(openapiutil.GetEncodeParam(tableId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetMmsTableResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetMmsTableResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @return GetMmsTableResponse
func (client *Client) GetMmsTable(sourceId *string, tableId *string) (_result *GetMmsTableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetMmsTableResponse{}
	_body, _err := client.GetMmsTableWithOptions(sourceId, tableId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMmsTaskResponse
func (client *Client) GetMmsTaskWithOptions(sourceId *string, taskId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetMmsTaskResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetMmsTask"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/mms/datasources/" + tea.StringValue(openapiutil.GetEncodeParam(sourceId)) + "/tasks/" + tea.StringValue(openapiutil.GetEncodeParam(taskId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetMmsTaskResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetMmsTaskResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @return GetMmsTaskResponse
func (client *Client) GetMmsTask(sourceId *string, taskId *string) (_result *GetMmsTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetMmsTaskResponse{}
	_body, _err := client.GetMmsTaskWithOptions(sourceId, taskId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the information about a package.
//
// @param request - GetPackageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPackageResponse
func (client *Client) GetPackageWithOptions(projectName *string, packageName *string, request *GetPackageRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetPackageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SourceProject)) {
		query["sourceProject"] = request.SourceProject
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPackage"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/projects/" + tea.StringValue(openapiutil.GetEncodeParam(projectName)) + "/packages/" + tea.StringValue(openapiutil.GetEncodeParam(packageName))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetPackageResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetPackageResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Obtains the information about a package.
//
// @param request - GetPackageRequest
//
// @return GetPackageResponse
func (client *Client) GetPackage(projectName *string, packageName *string, request *GetPackageRequest) (_result *GetPackageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetPackageResponse{}
	_body, _err := client.GetPackageWithOptions(projectName, packageName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a MaxCompute project.
//
// @param request - GetProjectRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProjectResponse
func (client *Client) GetProjectWithOptions(projectName *string, request *GetProjectRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Verbose)) {
		query["verbose"] = request.Verbose
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetProject"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/projects/" + tea.StringValue(openapiutil.GetEncodeParam(projectName))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetProjectResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetProjectResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the information about a MaxCompute project.
//
// @param request - GetProjectRequest
//
// @return GetProjectResponse
func (client *Client) GetProject(projectName *string, request *GetProjectRequest) (_result *GetProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetProjectResponse{}
	_body, _err := client.GetProjectWithOptions(projectName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the information about a specified level-1 quota.
//
// @param request - GetQuotaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQuotaResponse
func (client *Client) GetQuotaWithOptions(nickname *string, request *GetQuotaRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetQuotaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AkProven)) {
		query["AkProven"] = request.AkProven
	}

	if !tea.BoolValue(util.IsUnset(request.Mock)) {
		query["mock"] = request.Mock
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		query["tenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetQuota"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/quotas/" + tea.StringValue(openapiutil.GetEncodeParam(nickname))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetQuotaResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetQuotaResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Obtains the information about a specified level-1 quota.
//
// @param request - GetQuotaRequest
//
// @return GetQuotaResponse
func (client *Client) GetQuota(nickname *string, request *GetQuotaRequest) (_result *GetQuotaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetQuotaResponse{}
	_body, _err := client.GetQuotaWithOptions(nickname, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the information of a quota plan.
//
// @param request - GetQuotaPlanRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQuotaPlanResponse
func (client *Client) GetQuotaPlanWithOptions(nickname *string, planName *string, request *GetQuotaPlanRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetQuotaPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		query["tenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetQuotaPlan"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/quotas/" + tea.StringValue(openapiutil.GetEncodeParam(nickname)) + "/plans/" + tea.StringValue(openapiutil.GetEncodeParam(planName))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetQuotaPlanResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetQuotaPlanResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Obtains the information of a quota plan.
//
// @param request - GetQuotaPlanRequest
//
// @return GetQuotaPlanResponse
func (client *Client) GetQuotaPlan(nickname *string, planName *string, request *GetQuotaPlanRequest) (_result *GetQuotaPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetQuotaPlanResponse{}
	_body, _err := client.GetQuotaPlanWithOptions(nickname, planName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the scheduling plan for a quota plan.
//
// @param request - GetQuotaScheduleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQuotaScheduleResponse
func (client *Client) GetQuotaScheduleWithOptions(nickname *string, request *GetQuotaScheduleRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetQuotaScheduleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DisplayTimezone)) {
		query["displayTimezone"] = request.DisplayTimezone
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		query["tenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetQuotaSchedule"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/quotas/" + tea.StringValue(openapiutil.GetEncodeParam(nickname)) + "/schedule"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetQuotaScheduleResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetQuotaScheduleResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Obtains the scheduling plan for a quota plan.
//
// @param request - GetQuotaScheduleRequest
//
// @return GetQuotaScheduleResponse
func (client *Client) GetQuotaSchedule(nickname *string, request *GetQuotaScheduleRequest) (_result *GetQuotaScheduleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetQuotaScheduleResponse{}
	_body, _err := client.GetQuotaScheduleWithOptions(nickname, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries quota resource consumption information.
//
// @param tmpReq - GetQuotaUsageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQuotaUsageResponse
func (client *Client) GetQuotaUsageWithOptions(nickname *string, tmpReq *GetQuotaUsageRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetQuotaUsageResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetQuotaUsageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.PlotTypes)) {
		request.PlotTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PlotTypes, tea.String("plotTypes"), tea.String("simple"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.YAxisTypes)) {
		request.YAxisTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.YAxisTypes, tea.String("yAxisTypes"), tea.String("simple"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AggMethod)) {
		query["aggMethod"] = request.AggMethod
	}

	if !tea.BoolValue(util.IsUnset(request.From)) {
		query["from"] = request.From
	}

	if !tea.BoolValue(util.IsUnset(request.PlotTypesShrink)) {
		query["plotTypes"] = request.PlotTypesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ProductId)) {
		query["productId"] = request.ProductId
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.SubQuotaNickname)) {
		query["subQuotaNickname"] = request.SubQuotaNickname
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		query["tenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.To)) {
		query["to"] = request.To
	}

	if !tea.BoolValue(util.IsUnset(request.YAxisTypesShrink)) {
		query["yAxisTypes"] = request.YAxisTypesShrink
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetQuotaUsage"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/quotas/" + tea.StringValue(openapiutil.GetEncodeParam(nickname)) + "/usage"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetQuotaUsageResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetQuotaUsageResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries quota resource consumption information.
//
// @param request - GetQuotaUsageRequest
//
// @return GetQuotaUsageResponse
func (client *Client) GetQuotaUsage(nickname *string, request *GetQuotaUsageRequest) (_result *GetQuotaUsageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetQuotaUsageResponse{}
	_body, _err := client.GetQuotaUsageWithOptions(nickname, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the ACL-based permissions that is granted to a project-level role.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRoleAclResponse
func (client *Client) GetRoleAclWithOptions(projectName *string, roleName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetRoleAclResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetRoleAcl"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/projects/" + tea.StringValue(openapiutil.GetEncodeParam(projectName)) + "/roles/" + tea.StringValue(openapiutil.GetEncodeParam(roleName)) + "/roleAcl"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetRoleAclResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetRoleAclResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Obtains the ACL-based permissions that is granted to a project-level role.
//
// @return GetRoleAclResponse
func (client *Client) GetRoleAcl(projectName *string, roleName *string) (_result *GetRoleAclResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetRoleAclResponse{}
	_body, _err := client.GetRoleAclWithOptions(projectName, roleName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains ACL-based permissions on an object that are granted to a project-level role.
//
// @param request - GetRoleAclOnObjectRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRoleAclOnObjectResponse
func (client *Client) GetRoleAclOnObjectWithOptions(projectName *string, roleName *string, request *GetRoleAclOnObjectRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetRoleAclOnObjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ObjectName)) {
		query["objectName"] = request.ObjectName
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectType)) {
		query["objectType"] = request.ObjectType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRoleAclOnObject"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/projects/" + tea.StringValue(openapiutil.GetEncodeParam(projectName)) + "/roles/" + tea.StringValue(openapiutil.GetEncodeParam(roleName)) + "/acl"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetRoleAclOnObjectResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetRoleAclOnObjectResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Obtains ACL-based permissions on an object that are granted to a project-level role.
//
// @param request - GetRoleAclOnObjectRequest
//
// @return GetRoleAclOnObjectResponse
func (client *Client) GetRoleAclOnObject(projectName *string, roleName *string, request *GetRoleAclOnObjectRequest) (_result *GetRoleAclOnObjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetRoleAclOnObjectResponse{}
	_body, _err := client.GetRoleAclOnObjectWithOptions(projectName, roleName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the policy that is attached to a project-level role.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRolePolicyResponse
func (client *Client) GetRolePolicyWithOptions(projectName *string, roleName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetRolePolicyResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetRolePolicy"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/projects/" + tea.StringValue(openapiutil.GetEncodeParam(projectName)) + "/roles/" + tea.StringValue(openapiutil.GetEncodeParam(roleName)) + "/policy"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetRolePolicyResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetRolePolicyResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Obtains the policy that is attached to a project-level role.
//
// @return GetRolePolicyResponse
func (client *Client) GetRolePolicy(projectName *string, roleName *string) (_result *GetRolePolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetRolePolicyResponse{}
	_body, _err := client.GetRolePolicyWithOptions(projectName, roleName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the running state data of jobs that are in the running state in a specified period of time.
//
// @param tmpReq - GetRunningJobsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRunningJobsResponse
func (client *Client) GetRunningJobsWithOptions(tmpReq *GetRunningJobsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetRunningJobsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetRunningJobsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.JobOwnerList)) {
		request.JobOwnerListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.JobOwnerList, tea.String("jobOwnerList"), tea.String("simple"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.QuotaNicknameList)) {
		request.QuotaNicknameListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.QuotaNicknameList, tea.String("quotaNicknameList"), tea.String("simple"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.From)) {
		query["from"] = request.From
	}

	if !tea.BoolValue(util.IsUnset(request.JobOwnerListShrink)) {
		query["jobOwnerList"] = request.JobOwnerListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.QuotaNicknameListShrink)) {
		query["quotaNicknameList"] = request.QuotaNicknameListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.To)) {
		query["to"] = request.To
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRunningJobs"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/jobs/runningJobs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetRunningJobsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetRunningJobsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Obtains the running state data of jobs that are in the running state in a specified period of time.
//
// @param request - GetRunningJobsRequest
//
// @return GetRunningJobsResponse
func (client *Client) GetRunningJobs(request *GetRunningJobsRequest) (_result *GetRunningJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetRunningJobsResponse{}
	_body, _err := client.GetRunningJobsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Views the information about MaxCompute internal tables, views, external tables, clustered tables, or transactional tables.
//
// @param request - GetTableInfoRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTableInfoResponse
func (client *Client) GetTableInfoWithOptions(projectName *string, tableName *string, request *GetTableInfoRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTableInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SchemaName)) {
		query["schemaName"] = request.SchemaName
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTableInfo"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/projects/" + tea.StringValue(openapiutil.GetEncodeParam(projectName)) + "/tables/" + tea.StringValue(openapiutil.GetEncodeParam(tableName))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetTableInfoResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetTableInfoResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Views the information about MaxCompute internal tables, views, external tables, clustered tables, or transactional tables.
//
// @param request - GetTableInfoRequest
//
// @return GetTableInfoResponse
func (client *Client) GetTableInfo(projectName *string, tableName *string, request *GetTableInfoRequest) (_result *GetTableInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTableInfoResponse{}
	_body, _err := client.GetTableInfoWithOptions(projectName, tableName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the trusted projects of the current project.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTrustedProjectsResponse
func (client *Client) GetTrustedProjectsWithOptions(projectName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTrustedProjectsResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetTrustedProjects"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/projects/" + tea.StringValue(openapiutil.GetEncodeParam(projectName)) + "/trustedProjects"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetTrustedProjectsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetTrustedProjectsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Obtains the trusted projects of the current project.
//
// @return GetTrustedProjectsResponse
func (client *Client) GetTrustedProjects(projectName *string) (_result *GetTrustedProjectsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTrustedProjectsResponse{}
	_body, _err := client.GetTrustedProjectsWithOptions(projectName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Terminates a running job.
//
// @param request - KillJobsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return KillJobsResponse
func (client *Client) KillJobsWithOptions(request *KillJobsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *KillJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		query["tenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("KillJobs"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/jobs/kill"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &KillJobsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &KillJobsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Terminates a running job.
//
// @param request - KillJobsRequest
//
// @return KillJobsResponse
func (client *Client) KillJobs(request *KillJobsRequest) (_result *KillJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &KillJobsResponse{}
	_body, _err := client.KillJobsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Get compute usage of pay-as-you-go jobs.
//
// @param request - ListComputeMetricsByInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListComputeMetricsByInstanceResponse
func (client *Client) ListComputeMetricsByInstanceWithOptions(request *ListComputeMetricsByInstanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListComputeMetricsByInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		body["endDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["instanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.JobOwner)) {
		body["jobOwner"] = request.JobOwner
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectNames)) {
		body["projectNames"] = request.ProjectNames
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		body["region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.Signature)) {
		body["signature"] = request.Signature
	}

	if !tea.BoolValue(util.IsUnset(request.SpecCodes)) {
		body["specCodes"] = request.SpecCodes
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		body["startDate"] = request.StartDate
	}

	if !tea.BoolValue(util.IsUnset(request.Types)) {
		body["types"] = request.Types
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListComputeMetricsByInstance"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/computeMetrics/listByInstance"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListComputeMetricsByInstanceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListComputeMetricsByInstanceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Get compute usage of pay-as-you-go jobs.
//
// @param request - ListComputeMetricsByInstanceRequest
//
// @return ListComputeMetricsByInstanceResponse
func (client *Client) ListComputeMetricsByInstance(request *ListComputeMetricsByInstanceRequest) (_result *ListComputeMetricsByInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListComputeMetricsByInstanceResponse{}
	_body, _err := client.ListComputeMetricsByInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Get computeQuotaPlan list.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListComputeQuotaPlanResponse
func (client *Client) ListComputeQuotaPlanWithOptions(nickname *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListComputeQuotaPlanResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListComputeQuotaPlan"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/quotas/" + tea.StringValue(openapiutil.GetEncodeParam(nickname)) + "/computeQuotaPlan"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListComputeQuotaPlanResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListComputeQuotaPlanResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Get computeQuotaPlan list.
//
// @return ListComputeQuotaPlanResponse
func (client *Client) ListComputeQuotaPlan(nickname *string) (_result *ListComputeQuotaPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListComputeQuotaPlanResponse{}
	_body, _err := client.ListComputeQuotaPlanWithOptions(nickname, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains functions in a MaxCompute project.
//
// @param request - ListFunctionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFunctionsResponse
func (client *Client) ListFunctionsWithOptions(projectName *string, request *ListFunctionsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListFunctionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Marker)) {
		query["marker"] = request.Marker
	}

	if !tea.BoolValue(util.IsUnset(request.MaxItem)) {
		query["maxItem"] = request.MaxItem
	}

	if !tea.BoolValue(util.IsUnset(request.Prefix)) {
		query["prefix"] = request.Prefix
	}

	if !tea.BoolValue(util.IsUnset(request.SchemaName)) {
		query["schemaName"] = request.SchemaName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFunctions"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/projects/" + tea.StringValue(openapiutil.GetEncodeParam(projectName)) + "/functions"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListFunctionsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListFunctionsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Obtains functions in a MaxCompute project.
//
// @param request - ListFunctionsRequest
//
// @return ListFunctionsResponse
func (client *Client) ListFunctions(projectName *string, request *ListFunctionsRequest) (_result *ListFunctionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListFunctionsResponse{}
	_body, _err := client.ListFunctionsWithOptions(projectName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Views a list of jobs.
//
// @param request - ListJobInfosRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListJobInfosResponse
func (client *Client) ListJobInfosWithOptions(request *ListJobInfosRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListJobInfosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AscOrder)) {
		query["ascOrder"] = request.AscOrder
	}

	if !tea.BoolValue(util.IsUnset(request.OrderColumn)) {
		query["orderColumn"] = request.OrderColumn
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExtNodeIdList)) {
		body["extNodeIdList"] = request.ExtNodeIdList
	}

	if !tea.BoolValue(util.IsUnset(request.From)) {
		body["from"] = request.From
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceIdList)) {
		body["instanceIdList"] = request.InstanceIdList
	}

	if !tea.BoolValue(util.IsUnset(request.JobOwnerList)) {
		body["jobOwnerList"] = request.JobOwnerList
	}

	if !tea.BoolValue(util.IsUnset(request.PriorityList)) {
		body["priorityList"] = request.PriorityList
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectList)) {
		body["projectList"] = request.ProjectList
	}

	if !tea.BoolValue(util.IsUnset(request.QuotaNickname)) {
		body["quotaNickname"] = request.QuotaNickname
	}

	if !tea.BoolValue(util.IsUnset(request.SceneTagList)) {
		body["sceneTagList"] = request.SceneTagList
	}

	if !tea.BoolValue(util.IsUnset(request.SignatureList)) {
		body["signatureList"] = request.SignatureList
	}

	if !tea.BoolValue(util.IsUnset(request.SortByList)) {
		body["sortByList"] = request.SortByList
	}

	if !tea.BoolValue(util.IsUnset(request.SortOrderList)) {
		body["sortOrderList"] = request.SortOrderList
	}

	if !tea.BoolValue(util.IsUnset(request.StatusList)) {
		body["statusList"] = request.StatusList
	}

	if !tea.BoolValue(util.IsUnset(request.To)) {
		body["to"] = request.To
	}

	if !tea.BoolValue(util.IsUnset(request.TypeList)) {
		body["typeList"] = request.TypeList
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListJobInfos"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/jobs"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListJobInfosResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListJobInfosResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Views a list of jobs.
//
// @param request - ListJobInfosRequest
//
// @return ListJobInfosResponse
func (client *Client) ListJobInfos(request *ListJobInfosRequest) (_result *ListJobInfosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListJobInfosResponse{}
	_body, _err := client.ListJobInfosWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieve performance metrics for completed jobs.
//
// @param request - ListJobMetricRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListJobMetricResponse
func (client *Client) ListJobMetricWithOptions(request *ListJobMetricRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListJobMetricResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["startTime"] = request.StartTime
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Group)) {
		body["group"] = request.Group
	}

	if !tea.BoolValue(util.IsUnset(request.Metric)) {
		body["metric"] = request.Metric
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		body["period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.Project)) {
		body["project"] = request.Project
	}

	if !tea.BoolValue(util.IsUnset(request.Quota)) {
		body["quota"] = request.Quota
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListJobMetric"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/jobs/metric"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListJobMetricResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListJobMetricResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Retrieve performance metrics for completed jobs.
//
// @param request - ListJobMetricRequest
//
// @return ListJobMetricResponse
func (client *Client) ListJobMetric(request *ListJobMetricRequest) (_result *ListJobMetricResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListJobMetricResponse{}
	_body, _err := client.ListJobMetricWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Views a list of job snapshot data at a specific point in time.
//
// @param request - ListJobSnapshotInfosRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListJobSnapshotInfosResponse
func (client *Client) ListJobSnapshotInfosWithOptions(request *ListJobSnapshotInfosRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListJobSnapshotInfosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AscOrder)) {
		query["ascOrder"] = request.AscOrder
	}

	if !tea.BoolValue(util.IsUnset(request.OrderColumn)) {
		query["orderColumn"] = request.OrderColumn
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExtNodeIdList)) {
		body["extNodeIdList"] = request.ExtNodeIdList
	}

	if !tea.BoolValue(util.IsUnset(request.From)) {
		body["from"] = request.From
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceIdList)) {
		body["instanceIdList"] = request.InstanceIdList
	}

	if !tea.BoolValue(util.IsUnset(request.JobOwnerList)) {
		body["jobOwnerList"] = request.JobOwnerList
	}

	if !tea.BoolValue(util.IsUnset(request.PriorityList)) {
		body["priorityList"] = request.PriorityList
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectList)) {
		body["projectList"] = request.ProjectList
	}

	if !tea.BoolValue(util.IsUnset(request.QuotaNickname)) {
		body["quotaNickname"] = request.QuotaNickname
	}

	if !tea.BoolValue(util.IsUnset(request.SignatureList)) {
		body["signatureList"] = request.SignatureList
	}

	if !tea.BoolValue(util.IsUnset(request.SortByList)) {
		body["sortByList"] = request.SortByList
	}

	if !tea.BoolValue(util.IsUnset(request.SortOrderList)) {
		body["sortOrderList"] = request.SortOrderList
	}

	if !tea.BoolValue(util.IsUnset(request.StatusList)) {
		body["statusList"] = request.StatusList
	}

	if !tea.BoolValue(util.IsUnset(request.To)) {
		body["to"] = request.To
	}

	if !tea.BoolValue(util.IsUnset(request.TypeList)) {
		body["typeList"] = request.TypeList
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListJobSnapshotInfos"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/jobs/snapshot"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListJobSnapshotInfosResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListJobSnapshotInfosResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Views a list of job snapshot data at a specific point in time.
//
// @param request - ListJobSnapshotInfosRequest
//
// @return ListJobSnapshotInfosResponse
func (client *Client) ListJobSnapshotInfos(request *ListJobSnapshotInfosRequest) (_result *ListJobSnapshotInfosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListJobSnapshotInfosResponse{}
	_body, _err := client.ListJobSnapshotInfosWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListMmsDataSourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMmsDataSourcesResponse
func (client *Client) ListMmsDataSourcesWithOptions(request *ListMmsDataSourcesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListMmsDataSourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["pageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListMmsDataSources"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/mms/datasources"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListMmsDataSourcesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListMmsDataSourcesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - ListMmsDataSourcesRequest
//
// @return ListMmsDataSourcesResponse
func (client *Client) ListMmsDataSources(request *ListMmsDataSourcesRequest) (_result *ListMmsDataSourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListMmsDataSourcesResponse{}
	_body, _err := client.ListMmsDataSourcesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取一个数据源内“库”列表
//
// @param tmpReq - ListMmsDbsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMmsDbsResponse
func (client *Client) ListMmsDbsWithOptions(sourceId *string, tmpReq *ListMmsDbsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListMmsDbsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListMmsDbsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Sorter)) {
		request.SorterShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Sorter, tea.String("sorter"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["pageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SorterShrink)) {
		query["sorter"] = request.SorterShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListMmsDbs"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/mms/datasources/" + tea.StringValue(openapiutil.GetEncodeParam(sourceId)) + "/dbs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListMmsDbsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListMmsDbsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取一个数据源内“库”列表
//
// @param request - ListMmsDbsRequest
//
// @return ListMmsDbsResponse
func (client *Client) ListMmsDbs(sourceId *string, request *ListMmsDbsRequest) (_result *ListMmsDbsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListMmsDbsResponse{}
	_body, _err := client.ListMmsDbsWithOptions(sourceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListMmsJobsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMmsJobsResponse
func (client *Client) ListMmsJobsWithOptions(sourceId *string, request *ListMmsJobsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListMmsJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DstDbName)) {
		query["dstDbName"] = request.DstDbName
	}

	if !tea.BoolValue(util.IsUnset(request.DstTableName)) {
		query["dstTableName"] = request.DstTableName
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["pageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SrcDbName)) {
		query["srcDbName"] = request.SrcDbName
	}

	if !tea.BoolValue(util.IsUnset(request.SrcTableName)) {
		query["srcTableName"] = request.SrcTableName
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.Stopped)) {
		query["stopped"] = request.Stopped
	}

	if !tea.BoolValue(util.IsUnset(request.Sorter)) {
		query["sorter"] = request.Sorter
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListMmsJobs"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/mms/datasources/" + tea.StringValue(openapiutil.GetEncodeParam(sourceId)) + "/jobs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListMmsJobsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListMmsJobsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - ListMmsJobsRequest
//
// @return ListMmsJobsResponse
func (client *Client) ListMmsJobs(sourceId *string, request *ListMmsJobsRequest) (_result *ListMmsJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListMmsJobsResponse{}
	_body, _err := client.ListMmsJobsWithOptions(sourceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param tmpReq - ListMmsPartitionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMmsPartitionsResponse
func (client *Client) ListMmsPartitionsWithOptions(sourceId *string, tmpReq *ListMmsPartitionsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListMmsPartitionsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListMmsPartitionsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Status)) {
		request.StatusShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Status, tea.String("status"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbId)) {
		query["dbId"] = request.DbId
	}

	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["dbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.LastDdlTimeEnd)) {
		query["lastDdlTimeEnd"] = request.LastDdlTimeEnd
	}

	if !tea.BoolValue(util.IsUnset(request.LastDdlTimeStart)) {
		query["lastDdlTimeStart"] = request.LastDdlTimeStart
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["pageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StatusShrink)) {
		query["status"] = request.StatusShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		query["tableName"] = request.TableName
	}

	if !tea.BoolValue(util.IsUnset(request.Updated)) {
		query["updated"] = request.Updated
	}

	if !tea.BoolValue(util.IsUnset(request.Value)) {
		query["value"] = request.Value
	}

	if !tea.BoolValue(util.IsUnset(request.Sorter)) {
		query["sorter"] = request.Sorter
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListMmsPartitions"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/mms/datasources/" + tea.StringValue(openapiutil.GetEncodeParam(sourceId)) + "/partitions"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListMmsPartitionsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListMmsPartitionsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - ListMmsPartitionsRequest
//
// @return ListMmsPartitionsResponse
func (client *Client) ListMmsPartitions(sourceId *string, request *ListMmsPartitionsRequest) (_result *ListMmsPartitionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListMmsPartitionsResponse{}
	_body, _err := client.ListMmsPartitionsWithOptions(sourceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param tmpReq - ListMmsTablesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMmsTablesResponse
func (client *Client) ListMmsTablesWithOptions(sourceId *string, tmpReq *ListMmsTablesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListMmsTablesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListMmsTablesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Status)) {
		request.StatusShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Status, tea.String("status"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbId)) {
		query["dbId"] = request.DbId
	}

	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["dbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.HasPartitions)) {
		query["hasPartitions"] = request.HasPartitions
	}

	if !tea.BoolValue(util.IsUnset(request.LastDdlTimeEnd)) {
		query["lastDdlTimeEnd"] = request.LastDdlTimeEnd
	}

	if !tea.BoolValue(util.IsUnset(request.LastDdlTimeStart)) {
		query["lastDdlTimeStart"] = request.LastDdlTimeStart
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OnlyName)) {
		query["onlyName"] = request.OnlyName
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["pageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StatusShrink)) {
		query["status"] = request.StatusShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.Sorter)) {
		query["sorter"] = request.Sorter
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListMmsTables"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/mms/datasources/" + tea.StringValue(openapiutil.GetEncodeParam(sourceId)) + "/tables"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListMmsTablesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListMmsTablesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - ListMmsTablesRequest
//
// @return ListMmsTablesResponse
func (client *Client) ListMmsTables(sourceId *string, request *ListMmsTablesRequest) (_result *ListMmsTablesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListMmsTablesResponse{}
	_body, _err := client.ListMmsTablesWithOptions(sourceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMmsTaskLogsResponse
func (client *Client) ListMmsTaskLogsWithOptions(sourceId *string, taskId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListMmsTaskLogsResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListMmsTaskLogs"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/mms/datasources/" + tea.StringValue(openapiutil.GetEncodeParam(sourceId)) + "/tasks/" + tea.StringValue(openapiutil.GetEncodeParam(taskId)) + "/logs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListMmsTaskLogsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListMmsTaskLogsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @return ListMmsTaskLogsResponse
func (client *Client) ListMmsTaskLogs(sourceId *string, taskId *string) (_result *ListMmsTaskLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListMmsTaskLogsResponse{}
	_body, _err := client.ListMmsTaskLogsWithOptions(sourceId, taskId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListMmsTasksRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMmsTasksResponse
func (client *Client) ListMmsTasksWithOptions(sourceId *string, request *ListMmsTasksRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListMmsTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DstDbName)) {
		query["dstDbName"] = request.DstDbName
	}

	if !tea.BoolValue(util.IsUnset(request.DstTableName)) {
		query["dstTableName"] = request.DstTableName
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["jobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.JobName)) {
		query["jobName"] = request.JobName
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["pageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Partition)) {
		query["partition"] = request.Partition
	}

	if !tea.BoolValue(util.IsUnset(request.SrcDbName)) {
		query["srcDbName"] = request.SrcDbName
	}

	if !tea.BoolValue(util.IsUnset(request.SrcTableName)) {
		query["srcTableName"] = request.SrcTableName
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.Sorter)) {
		query["sorter"] = request.Sorter
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListMmsTasks"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/mms/datasources/" + tea.StringValue(openapiutil.GetEncodeParam(sourceId)) + "/tasks"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListMmsTasksResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListMmsTasksResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - ListMmsTasksRequest
//
// @return ListMmsTasksResponse
func (client *Client) ListMmsTasks(sourceId *string, request *ListMmsTasksRequest) (_result *ListMmsTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListMmsTasksResponse{}
	_body, _err := client.ListMmsTasksWithOptions(sourceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the packages in a MaxCompute project.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPackagesResponse
func (client *Client) ListPackagesWithOptions(projectName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListPackagesResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListPackages"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/projects/" + tea.StringValue(openapiutil.GetEncodeParam(projectName)) + "/packages"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListPackagesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListPackagesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the packages in a MaxCompute project.
//
// @return ListPackagesResponse
func (client *Client) ListPackages(projectName *string) (_result *ListPackagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListPackagesResponse{}
	_body, _err := client.ListPackagesWithOptions(projectName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of users in a project.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProjectUsersResponse
func (client *Client) ListProjectUsersWithOptions(projectName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListProjectUsersResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListProjectUsers"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/projects/" + tea.StringValue(openapiutil.GetEncodeParam(projectName)) + "/users"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListProjectUsersResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListProjectUsersResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries a list of users in a project.
//
// @return ListProjectUsersResponse
func (client *Client) ListProjectUsers(projectName *string) (_result *ListProjectUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListProjectUsersResponse{}
	_body, _err := client.ListProjectUsersWithOptions(projectName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of MaxCompute projects.
//
// @param request - ListProjectsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProjectsResponse
func (client *Client) ListProjectsWithOptions(request *ListProjectsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListProjectsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ListSystemCatalog)) {
		query["listSystemCatalog"] = request.ListSystemCatalog
	}

	if !tea.BoolValue(util.IsUnset(request.Marker)) {
		query["marker"] = request.Marker
	}

	if !tea.BoolValue(util.IsUnset(request.MaxItem)) {
		query["maxItem"] = request.MaxItem
	}

	if !tea.BoolValue(util.IsUnset(request.Prefix)) {
		query["prefix"] = request.Prefix
	}

	if !tea.BoolValue(util.IsUnset(request.QuotaName)) {
		query["quotaName"] = request.QuotaName
	}

	if !tea.BoolValue(util.IsUnset(request.QuotaNickName)) {
		query["quotaNickName"] = request.QuotaNickName
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.SaleTags)) {
		query["saleTags"] = request.SaleTags
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		query["tenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListProjects"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/projects"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListProjectsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListProjectsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries a list of MaxCompute projects.
//
// @param request - ListProjectsRequest
//
// @return ListProjectsResponse
func (client *Client) ListProjects(request *ListProjectsRequest) (_result *ListProjectsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListProjectsResponse{}
	_body, _err := client.ListProjectsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries quotas.
//
// @param request - ListQuotasRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQuotasResponse
func (client *Client) ListQuotasWithOptions(request *ListQuotasRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListQuotasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BillingType)) {
		query["billingType"] = request.BillingType
	}

	if !tea.BoolValue(util.IsUnset(request.Marker)) {
		query["marker"] = request.Marker
	}

	if !tea.BoolValue(util.IsUnset(request.MaxItem)) {
		query["maxItem"] = request.MaxItem
	}

	if !tea.BoolValue(util.IsUnset(request.ProductId)) {
		query["productId"] = request.ProductId
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.SaleTags)) {
		query["saleTags"] = request.SaleTags
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		query["tenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListQuotas"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/quotas"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListQuotasResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListQuotasResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries quotas.
//
// @param request - ListQuotasRequest
//
// @return ListQuotasResponse
func (client *Client) ListQuotas(request *ListQuotasRequest) (_result *ListQuotasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListQuotasResponse{}
	_body, _err := client.ListQuotasWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains quota plans.
//
// @param request - ListQuotasPlansRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQuotasPlansResponse
func (client *Client) ListQuotasPlansWithOptions(nickname *string, request *ListQuotasPlansRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListQuotasPlansResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		query["tenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListQuotasPlans"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/quotas/" + tea.StringValue(openapiutil.GetEncodeParam(nickname)) + "/plans"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListQuotasPlansResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListQuotasPlansResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Obtains quota plans.
//
// @param request - ListQuotasPlansRequest
//
// @return ListQuotasPlansResponse
func (client *Client) ListQuotasPlans(nickname *string, request *ListQuotasPlansRequest) (_result *ListQuotasPlansResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListQuotasPlansResponse{}
	_body, _err := client.ListQuotasPlansWithOptions(nickname, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains resources in a MaxCompute project.
//
// @param request - ListResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourcesResponse
func (client *Client) ListResourcesWithOptions(projectName *string, request *ListResourcesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Marker)) {
		query["marker"] = request.Marker
	}

	if !tea.BoolValue(util.IsUnset(request.MaxItem)) {
		query["maxItem"] = request.MaxItem
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.SchemaName)) {
		query["schemaName"] = request.SchemaName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListResources"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/projects/" + tea.StringValue(openapiutil.GetEncodeParam(projectName)) + "/resources"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListResourcesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListResourcesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Obtains resources in a MaxCompute project.
//
// @param request - ListResourcesRequest
//
// @return ListResourcesResponse
func (client *Client) ListResources(projectName *string, request *ListResourcesRequest) (_result *ListResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListResourcesResponse{}
	_body, _err := client.ListResourcesWithOptions(projectName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains MaxCompute project-level roles.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRolesResponse
func (client *Client) ListRolesWithOptions(projectName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListRolesResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListRoles"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/projects/" + tea.StringValue(openapiutil.GetEncodeParam(projectName)) + "/roles"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListRolesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListRolesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Obtains MaxCompute project-level roles.
//
// @return ListRolesResponse
func (client *Client) ListRoles(projectName *string) (_result *ListRolesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListRolesResponse{}
	_body, _err := client.ListRolesWithOptions(projectName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the storage details of a specific partition in a partitioned table in a MaxCompute project.
//
// @param tmpReq - ListStoragePartitionsInfoRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListStoragePartitionsInfoResponse
func (client *Client) ListStoragePartitionsInfoWithOptions(project *string, table *string, tmpReq *ListStoragePartitionsInfoRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListStoragePartitionsInfoResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListStoragePartitionsInfoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Types)) {
		request.TypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Types, tea.String("types"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AscOrder)) {
		query["ascOrder"] = request.AscOrder
	}

	if !tea.BoolValue(util.IsUnset(request.Date)) {
		query["date"] = request.Date
	}

	if !tea.BoolValue(util.IsUnset(request.OrderColumn)) {
		query["orderColumn"] = request.OrderColumn
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PartitionPrefix)) {
		query["partitionPrefix"] = request.PartitionPrefix
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.Schema)) {
		query["schema"] = request.Schema
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		query["tenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.TypesShrink)) {
		query["types"] = request.TypesShrink
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListStoragePartitionsInfo"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/observations/analysis/storage/projects/" + tea.StringValue(openapiutil.GetEncodeParam(project)) + "/tables/" + tea.StringValue(openapiutil.GetEncodeParam(table)) + "/partitionsInfo"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListStoragePartitionsInfoResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListStoragePartitionsInfoResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the storage details of a specific partition in a partitioned table in a MaxCompute project.
//
// @param request - ListStoragePartitionsInfoRequest
//
// @return ListStoragePartitionsInfoResponse
func (client *Client) ListStoragePartitionsInfo(project *string, table *string, request *ListStoragePartitionsInfoRequest) (_result *ListStoragePartitionsInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListStoragePartitionsInfoResponse{}
	_body, _err := client.ListStoragePartitionsInfoWithOptions(project, table, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the table storage details of a MaxCompute project.
//
// @param tmpReq - ListStorageTablesInfoRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListStorageTablesInfoResponse
func (client *Client) ListStorageTablesInfoWithOptions(project *string, tmpReq *ListStorageTablesInfoRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListStorageTablesInfoResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListStorageTablesInfoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Types)) {
		request.TypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Types, tea.String("types"), tea.String("simple"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AscOrder)) {
		query["ascOrder"] = request.AscOrder
	}

	if !tea.BoolValue(util.IsUnset(request.Date)) {
		query["date"] = request.Date
	}

	if !tea.BoolValue(util.IsUnset(request.OrderColumn)) {
		query["orderColumn"] = request.OrderColumn
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RecentDays)) {
		query["recentDays"] = request.RecentDays
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.Schema)) {
		query["schema"] = request.Schema
	}

	if !tea.BoolValue(util.IsUnset(request.TablePrefix)) {
		query["tablePrefix"] = request.TablePrefix
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		query["tenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.TypesShrink)) {
		query["types"] = request.TypesShrink
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListStorageTablesInfo"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/observations/analysis/storage/projects/" + tea.StringValue(openapiutil.GetEncodeParam(project)) + "/tablesInfo"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListStorageTablesInfoResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListStorageTablesInfoResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the table storage details of a MaxCompute project.
//
// @param request - ListStorageTablesInfoRequest
//
// @return ListStorageTablesInfoResponse
func (client *Client) ListStorageTablesInfo(project *string, request *ListStorageTablesInfoRequest) (_result *ListStorageTablesInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListStorageTablesInfoResponse{}
	_body, _err := client.ListStorageTablesInfoWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains tables in a MaxCompute project.
//
// @param request - ListTablesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTablesResponse
func (client *Client) ListTablesWithOptions(projectName *string, request *ListTablesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTablesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Marker)) {
		query["marker"] = request.Marker
	}

	if !tea.BoolValue(util.IsUnset(request.MaxItem)) {
		query["maxItem"] = request.MaxItem
	}

	if !tea.BoolValue(util.IsUnset(request.Prefix)) {
		query["prefix"] = request.Prefix
	}

	if !tea.BoolValue(util.IsUnset(request.SchemaName)) {
		query["schemaName"] = request.SchemaName
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTables"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/projects/" + tea.StringValue(openapiutil.GetEncodeParam(projectName)) + "/tables"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListTablesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListTablesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Obtains tables in a MaxCompute project.
//
// @param request - ListTablesRequest
//
// @return ListTablesResponse
func (client *Client) ListTables(projectName *string, request *ListTablesRequest) (_result *ListTablesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTablesResponse{}
	_body, _err := client.ListTablesWithOptions(projectName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Displays the time-specific configuration of an exclusive resource group for Tunnel (referred to as Tunnel quota).
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTunnelQuotaTimerResponse
func (client *Client) ListTunnelQuotaTimerWithOptions(nickname *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTunnelQuotaTimerResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListTunnelQuotaTimer"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/tunnel/" + tea.StringValue(openapiutil.GetEncodeParam(nickname)) + "/timers"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListTunnelQuotaTimerResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListTunnelQuotaTimerResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Displays the time-specific configuration of an exclusive resource group for Tunnel (referred to as Tunnel quota).
//
// @return ListTunnelQuotaTimerResponse
func (client *Client) ListTunnelQuotaTimer(nickname *string) (_result *ListTunnelQuotaTimerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTunnelQuotaTimerResponse{}
	_body, _err := client.ListTunnelQuotaTimerWithOptions(nickname, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of all users under a tenant.
//
// @param request - ListUsersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUsersResponse
func (client *Client) ListUsersWithOptions(request *ListUsersRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUsers"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/users"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListUsersResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListUsersResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries a list of all users under a tenant.
//
// @param request - ListUsersRequest
//
// @return ListUsersResponse
func (client *Client) ListUsers(request *ListUsersRequest) (_result *ListUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListUsersResponse{}
	_body, _err := client.ListUsersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains information about the users who are assigned a project-level role.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUsersByRoleResponse
func (client *Client) ListUsersByRoleWithOptions(projectName *string, roleName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListUsersByRoleResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListUsersByRole"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/projects/" + tea.StringValue(openapiutil.GetEncodeParam(projectName)) + "/roles/" + tea.StringValue(openapiutil.GetEncodeParam(roleName)) + "/users"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListUsersByRoleResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListUsersByRoleResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Obtains information about the users who are assigned a project-level role.
//
// @return ListUsersByRoleResponse
func (client *Client) ListUsersByRole(projectName *string, roleName *string) (_result *ListUsersByRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListUsersByRoleResponse{}
	_body, _err := client.ListUsersByRoleWithOptions(projectName, roleName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryQuotaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryQuotaResponse
func (client *Client) QueryQuotaWithOptions(nickname *string, request *QueryQuotaRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryQuotaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AkProven)) {
		query["AkProven"] = request.AkProven
	}

	if !tea.BoolValue(util.IsUnset(request.Mock)) {
		query["mock"] = request.Mock
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		query["tenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryQuota"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/quotas/" + tea.StringValue(openapiutil.GetEncodeParam(nickname)) + "/query"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &QueryQuotaResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &QueryQuotaResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - QueryQuotaRequest
//
// @return QueryQuotaResponse
func (client *Client) QueryQuota(nickname *string, request *QueryQuotaRequest) (_result *QueryQuotaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryQuotaResponse{}
	_body, _err := client.QueryQuotaWithOptions(nickname, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RetryMmsJobResponse
func (client *Client) RetryMmsJobWithOptions(sourceId *string, jobId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RetryMmsJobResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("RetryMmsJob"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/mms/datasources/" + tea.StringValue(openapiutil.GetEncodeParam(sourceId)) + "/jobs/" + tea.StringValue(openapiutil.GetEncodeParam(jobId)) + "/retry"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &RetryMmsJobResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &RetryMmsJobResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @return RetryMmsJobResponse
func (client *Client) RetryMmsJob(sourceId *string, jobId *string) (_result *RetryMmsJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RetryMmsJobResponse{}
	_body, _err := client.RetryMmsJobWithOptions(sourceId, jobId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartMmsJobResponse
func (client *Client) StartMmsJobWithOptions(sourceId *string, jobId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StartMmsJobResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("StartMmsJob"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/mms/datasources/" + tea.StringValue(openapiutil.GetEncodeParam(sourceId)) + "/jobs/" + tea.StringValue(openapiutil.GetEncodeParam(jobId)) + "/start"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &StartMmsJobResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &StartMmsJobResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @return StartMmsJobResponse
func (client *Client) StartMmsJob(sourceId *string, jobId *string) (_result *StartMmsJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StartMmsJobResponse{}
	_body, _err := client.StartMmsJobWithOptions(sourceId, jobId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopMmsJobResponse
func (client *Client) StopMmsJobWithOptions(sourceId *string, jobId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StopMmsJobResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("StopMmsJob"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/mms/datasources/" + tea.StringValue(openapiutil.GetEncodeParam(sourceId)) + "/jobs/" + tea.StringValue(openapiutil.GetEncodeParam(jobId)) + "/stop"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &StopMmsJobResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &StopMmsJobResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @return StopMmsJobResponse
func (client *Client) StopMmsJob(sourceId *string, jobId *string) (_result *StopMmsJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopMmsJobResponse{}
	_body, _err := client.StopMmsJobWithOptions(sourceId, jobId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Update the ComputeQuotaPlan.
//
// Description:
//
// Please ensure that before using this interface, you have fully understood the <props="china">[Pricing and Charges](https://help.aliyun.com/zh/maxcompute/product-overview/computing-pricing-1)
//
// <props="intl">[Pricing and Charges](https://www.alibabacloud.com/help/maxcompute/product-overview/computing-pricing-1) of MaxCompute Elastic Reserved CU.
//
// @param request - UpdateComputeQuotaPlanRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateComputeQuotaPlanResponse
func (client *Client) UpdateComputeQuotaPlanWithOptions(nickname *string, request *UpdateComputeQuotaPlanRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateComputeQuotaPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Quota)) {
		body["quota"] = request.Quota
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateComputeQuotaPlan"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/quotas/" + tea.StringValue(openapiutil.GetEncodeParam(nickname)) + "/computeQuotaPlan"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateComputeQuotaPlanResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateComputeQuotaPlanResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Update the ComputeQuotaPlan.
//
// Description:
//
// Please ensure that before using this interface, you have fully understood the <props="china">[Pricing and Charges](https://help.aliyun.com/zh/maxcompute/product-overview/computing-pricing-1)
//
// <props="intl">[Pricing and Charges](https://www.alibabacloud.com/help/maxcompute/product-overview/computing-pricing-1) of MaxCompute Elastic Reserved CU.
//
// @param request - UpdateComputeQuotaPlanRequest
//
// @return UpdateComputeQuotaPlanResponse
func (client *Client) UpdateComputeQuotaPlan(nickname *string, request *UpdateComputeQuotaPlanRequest) (_result *UpdateComputeQuotaPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateComputeQuotaPlanResponse{}
	_body, _err := client.UpdateComputeQuotaPlanWithOptions(nickname, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Update the time-based plan for computing quota.
//
// Description:
//
// Please ensure that before using this interface, you have fully understood the<props="china">[Pricing and Billing](https://help.aliyun.com/zh/maxcompute/product-overview/computing-pricing-1)
//
// <props="intl">[Pricing and Billing](https://www.alibabacloud.com/help/maxcompute/product-overview/computing-pricing-1) of MaxCompute Elastic Reserved CU.
//
// @param request - UpdateComputeQuotaScheduleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateComputeQuotaScheduleResponse
func (client *Client) UpdateComputeQuotaScheduleWithOptions(nickname *string, request *UpdateComputeQuotaScheduleRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateComputeQuotaScheduleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ScheduleTimezone)) {
		query["scheduleTimezone"] = request.ScheduleTimezone
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    util.ToArray(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateComputeQuotaSchedule"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/quotas/" + tea.StringValue(openapiutil.GetEncodeParam(nickname)) + "/computeQuotaSchedule"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateComputeQuotaScheduleResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateComputeQuotaScheduleResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Update the time-based plan for computing quota.
//
// Description:
//
// Please ensure that before using this interface, you have fully understood the<props="china">[Pricing and Billing](https://help.aliyun.com/zh/maxcompute/product-overview/computing-pricing-1)
//
// <props="intl">[Pricing and Billing](https://www.alibabacloud.com/help/maxcompute/product-overview/computing-pricing-1) of MaxCompute Elastic Reserved CU.
//
// @param request - UpdateComputeQuotaScheduleRequest
//
// @return UpdateComputeQuotaScheduleResponse
func (client *Client) UpdateComputeQuotaSchedule(nickname *string, request *UpdateComputeQuotaScheduleRequest) (_result *UpdateComputeQuotaScheduleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateComputeQuotaScheduleResponse{}
	_body, _err := client.UpdateComputeQuotaScheduleWithOptions(nickname, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Update the basic configuration of the calculation quota, including adding or deleting the sub quota, modifying the basic properties of the secondary quota, and the CU configuration of the effective quota plan.
//
// @param request - UpdateComputeSubQuotaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateComputeSubQuotaResponse
func (client *Client) UpdateComputeSubQuotaWithOptions(nickname *string, request *UpdateComputeSubQuotaRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateComputeSubQuotaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SubQuotaInfoList)) {
		body["subQuotaInfoList"] = request.SubQuotaInfoList
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateComputeSubQuota"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/quotas/" + tea.StringValue(openapiutil.GetEncodeParam(nickname)) + "/computeSubQuota"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateComputeSubQuotaResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateComputeSubQuotaResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Update the basic configuration of the calculation quota, including adding or deleting the sub quota, modifying the basic properties of the secondary quota, and the CU configuration of the effective quota plan.
//
// @param request - UpdateComputeSubQuotaRequest
//
// @return UpdateComputeSubQuotaResponse
func (client *Client) UpdateComputeSubQuota(nickname *string, request *UpdateComputeSubQuotaRequest) (_result *UpdateComputeSubQuotaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateComputeSubQuotaResponse{}
	_body, _err := client.UpdateComputeSubQuotaWithOptions(nickname, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新数据源配置、名称，启/停数据源实例
//
// @param request - UpdateMmsDataSourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMmsDataSourceResponse
func (client *Client) UpdateMmsDataSourceWithOptions(sourceId *string, request *UpdateMmsDataSourceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateMmsDataSourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Action)) {
		body["action"] = request.Action
	}

	if !tea.BoolValue(util.IsUnset(request.Config)) {
		body["config"] = request.Config
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Test)) {
		body["test"] = request.Test
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateMmsDataSource"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/mms/datasources/" + tea.StringValue(openapiutil.GetEncodeParam(sourceId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateMmsDataSourceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateMmsDataSourceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 更新数据源配置、名称，启/停数据源实例
//
// @param request - UpdateMmsDataSourceRequest
//
// @return UpdateMmsDataSourceResponse
func (client *Client) UpdateMmsDataSource(sourceId *string, request *UpdateMmsDataSourceRequest) (_result *UpdateMmsDataSourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateMmsDataSourceResponse{}
	_body, _err := client.UpdateMmsDataSourceWithOptions(sourceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the objects in a package and projects in which the package can be installed.
//
// @param request - UpdatePackageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePackageResponse
func (client *Client) UpdatePackageWithOptions(projectName *string, packageName *string, request *UpdatePackageRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdatePackageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("UpdatePackage"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/projects/" + tea.StringValue(openapiutil.GetEncodeParam(projectName)) + "/packages/" + tea.StringValue(openapiutil.GetEncodeParam(packageName))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdatePackageResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdatePackageResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Updates the objects in a package and projects in which the package can be installed.
//
// @param request - UpdatePackageRequest
//
// @return UpdatePackageResponse
func (client *Client) UpdatePackage(projectName *string, packageName *string, request *UpdatePackageRequest) (_result *UpdatePackageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdatePackageResponse{}
	_body, _err := client.UpdatePackageWithOptions(projectName, packageName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Update Project Basic Information
//
// @param request - UpdateProjectBasicMetaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateProjectBasicMetaResponse
func (client *Client) UpdateProjectBasicMetaWithOptions(projectName *string, request *UpdateProjectBasicMetaRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateProjectBasicMetaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Comment)) {
		body["comment"] = request.Comment
	}

	if !tea.BoolValue(util.IsUnset(request.Properties)) {
		body["properties"] = request.Properties
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateProjectBasicMeta"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/projects/" + tea.StringValue(openapiutil.GetEncodeParam(projectName)) + "/meta"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateProjectBasicMetaResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateProjectBasicMetaResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Update Project Basic Information
//
// @param request - UpdateProjectBasicMetaRequest
//
// @return UpdateProjectBasicMetaResponse
func (client *Client) UpdateProjectBasicMeta(projectName *string, request *UpdateProjectBasicMetaRequest) (_result *UpdateProjectBasicMetaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateProjectBasicMetaResponse{}
	_body, _err := client.UpdateProjectBasicMetaWithOptions(projectName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modify Default Project Compute Quota
//
// @param request - UpdateProjectDefaultQuotaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateProjectDefaultQuotaResponse
func (client *Client) UpdateProjectDefaultQuotaWithOptions(projectName *string, request *UpdateProjectDefaultQuotaRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateProjectDefaultQuotaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Quota)) {
		body["quota"] = request.Quota
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateProjectDefaultQuota"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/projects/" + tea.StringValue(openapiutil.GetEncodeParam(projectName)) + "/quota"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateProjectDefaultQuotaResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateProjectDefaultQuotaResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Modify Default Project Compute Quota
//
// @param request - UpdateProjectDefaultQuotaRequest
//
// @return UpdateProjectDefaultQuotaResponse
func (client *Client) UpdateProjectDefaultQuota(projectName *string, request *UpdateProjectDefaultQuotaRequest) (_result *UpdateProjectDefaultQuotaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateProjectDefaultQuotaResponse{}
	_body, _err := client.UpdateProjectDefaultQuotaWithOptions(projectName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the IP address whitelist of a MaxCompute project.
//
// @param request - UpdateProjectIpWhiteListRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateProjectIpWhiteListResponse
func (client *Client) UpdateProjectIpWhiteListWithOptions(projectName *string, request *UpdateProjectIpWhiteListRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateProjectIpWhiteListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateProjectIpWhiteList"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/projects/" + tea.StringValue(openapiutil.GetEncodeParam(projectName)) + "/ipWhiteList"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateProjectIpWhiteListResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateProjectIpWhiteListResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Modifies the IP address whitelist of a MaxCompute project.
//
// @param request - UpdateProjectIpWhiteListRequest
//
// @return UpdateProjectIpWhiteListResponse
func (client *Client) UpdateProjectIpWhiteList(projectName *string, request *UpdateProjectIpWhiteListRequest) (_result *UpdateProjectIpWhiteListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateProjectIpWhiteListResponse{}
	_body, _err := client.UpdateProjectIpWhiteListWithOptions(projectName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a quota plan.
//
// @param request - UpdateQuotaPlanRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateQuotaPlanResponse
func (client *Client) UpdateQuotaPlanWithOptions(nickname *string, planName *string, request *UpdateQuotaPlanRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateQuotaPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		query["tenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateQuotaPlan"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/quotas/" + tea.StringValue(openapiutil.GetEncodeParam(nickname)) + "/plans/" + tea.StringValue(openapiutil.GetEncodeParam(planName))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateQuotaPlanResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateQuotaPlanResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Updates a quota plan.
//
// @param request - UpdateQuotaPlanRequest
//
// @return UpdateQuotaPlanResponse
func (client *Client) UpdateQuotaPlan(nickname *string, planName *string, request *UpdateQuotaPlanRequest) (_result *UpdateQuotaPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateQuotaPlanResponse{}
	_body, _err := client.UpdateQuotaPlanWithOptions(nickname, planName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the scheduling plan for a quota plan.
//
// @param request - UpdateQuotaScheduleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateQuotaScheduleResponse
func (client *Client) UpdateQuotaScheduleWithOptions(nickname *string, request *UpdateQuotaScheduleRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateQuotaScheduleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		query["tenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateQuotaSchedule"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/quotas/" + tea.StringValue(openapiutil.GetEncodeParam(nickname)) + "/schedule"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateQuotaScheduleResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateQuotaScheduleResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Updates the scheduling plan for a quota plan.
//
// @param request - UpdateQuotaScheduleRequest
//
// @return UpdateQuotaScheduleResponse
func (client *Client) UpdateQuotaSchedule(nickname *string, request *UpdateQuotaScheduleRequest) (_result *UpdateQuotaScheduleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateQuotaScheduleResponse{}
	_body, _err := client.UpdateQuotaScheduleWithOptions(nickname, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the time-specific configuration of an exclusive resource group for Tunnel (referred to as Tunnel quota).
//
// Description:
//
// Before you call this operation, make sure that you are familiar with the [billing and prices](https://www.alibabacloud.com/help/maxcompute/product-overview/data-transfer-fees-hourly-billing) of Tunnel quotas and elastically reserved computing resources.
//
// @param request - UpdateTunnelQuotaTimerRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTunnelQuotaTimerResponse
func (client *Client) UpdateTunnelQuotaTimerWithOptions(nickname *string, request *UpdateTunnelQuotaTimerRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateTunnelQuotaTimerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Timezone)) {
		query["timezone"] = request.Timezone
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    util.ToArray(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateTunnelQuotaTimer"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/tunnel/" + tea.StringValue(openapiutil.GetEncodeParam(nickname)) + "/timers"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateTunnelQuotaTimerResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateTunnelQuotaTimerResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Updates the time-specific configuration of an exclusive resource group for Tunnel (referred to as Tunnel quota).
//
// Description:
//
// Before you call this operation, make sure that you are familiar with the [billing and prices](https://www.alibabacloud.com/help/maxcompute/product-overview/data-transfer-fees-hourly-billing) of Tunnel quotas and elastically reserved computing resources.
//
// @param request - UpdateTunnelQuotaTimerRequest
//
// @return UpdateTunnelQuotaTimerResponse
func (client *Client) UpdateTunnelQuotaTimer(nickname *string, request *UpdateTunnelQuotaTimerRequest) (_result *UpdateTunnelQuotaTimerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateTunnelQuotaTimerResponse{}
	_body, _err := client.UpdateTunnelQuotaTimerWithOptions(nickname, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
