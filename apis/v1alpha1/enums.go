// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

type AccessDeniedExceptionReason string

const (
	AccessDeniedExceptionReason_DEPENDENCY_ACCESS_DENIED AccessDeniedExceptionReason = "DEPENDENCY_ACCESS_DENIED"
	AccessDeniedExceptionReason_UNAUTHORIZED_ACCOUNT     AccessDeniedExceptionReason = "UNAUTHORIZED_ACCOUNT"
)

type ChecksumAggregationMethod string

const (
	ChecksumAggregationMethod_LINEAR ChecksumAggregationMethod = "LINEAR"
)

type ChecksumAlgorithm string

const (
	ChecksumAlgorithm_SHA256 ChecksumAlgorithm = "SHA256"
)

type RequestThrottledExceptionReason string

const (
	RequestThrottledExceptionReason_ACCOUNT_THROTTLED            RequestThrottledExceptionReason = "ACCOUNT_THROTTLED"
	RequestThrottledExceptionReason_DEPENDENCY_REQUEST_THROTTLED RequestThrottledExceptionReason = "DEPENDENCY_REQUEST_THROTTLED"
	RequestThrottledExceptionReason_RESOURCE_LEVEL_THROTTLE      RequestThrottledExceptionReason = "RESOURCE_LEVEL_THROTTLE"
)

type ResourceNotFoundExceptionReason string

const (
	ResourceNotFoundExceptionReason_DEPENDENCY_RESOURCE_NOT_FOUND ResourceNotFoundExceptionReason = "DEPENDENCY_RESOURCE_NOT_FOUND"
	ResourceNotFoundExceptionReason_GRANT_NOT_FOUND               ResourceNotFoundExceptionReason = "GRANT_NOT_FOUND"
	ResourceNotFoundExceptionReason_IMAGE_NOT_FOUND               ResourceNotFoundExceptionReason = "IMAGE_NOT_FOUND"
	ResourceNotFoundExceptionReason_SNAPSHOT_NOT_FOUND            ResourceNotFoundExceptionReason = "SNAPSHOT_NOT_FOUND"
)

type SSEType string

const (
	SSEType_none    SSEType = "none"
	SSEType_sse_ebs SSEType = "sse-ebs"
	SSEType_sse_kms SSEType = "sse-kms"
)

type ServiceQuotaExceededExceptionReason string

const (
	ServiceQuotaExceededExceptionReason_DEPENDENCY_SERVICE_QUOTA_EXCEEDED ServiceQuotaExceededExceptionReason = "DEPENDENCY_SERVICE_QUOTA_EXCEEDED"
)

type Status string

const (
	Status_completed Status = "completed"
	Status_error     Status = "error"
	Status_pending   Status = "pending"
)

type ValidationExceptionReason string

const (
	ValidationExceptionReason_CONFLICTING_BLOCK_UPDATE   ValidationExceptionReason = "CONFLICTING_BLOCK_UPDATE"
	ValidationExceptionReason_INVALID_BLOCK              ValidationExceptionReason = "INVALID_BLOCK"
	ValidationExceptionReason_INVALID_BLOCK_TOKEN        ValidationExceptionReason = "INVALID_BLOCK_TOKEN"
	ValidationExceptionReason_INVALID_CONTENT_ENCODING   ValidationExceptionReason = "INVALID_CONTENT_ENCODING"
	ValidationExceptionReason_INVALID_CUSTOMER_KEY       ValidationExceptionReason = "INVALID_CUSTOMER_KEY"
	ValidationExceptionReason_INVALID_DEPENDENCY_REQUEST ValidationExceptionReason = "INVALID_DEPENDENCY_REQUEST"
	ValidationExceptionReason_INVALID_GRANT_TOKEN        ValidationExceptionReason = "INVALID_GRANT_TOKEN"
	ValidationExceptionReason_INVALID_IMAGE_ID           ValidationExceptionReason = "INVALID_IMAGE_ID"
	ValidationExceptionReason_INVALID_PAGE_TOKEN         ValidationExceptionReason = "INVALID_PAGE_TOKEN"
	ValidationExceptionReason_INVALID_PARAMETER_VALUE    ValidationExceptionReason = "INVALID_PARAMETER_VALUE"
	ValidationExceptionReason_INVALID_SNAPSHOT_ID        ValidationExceptionReason = "INVALID_SNAPSHOT_ID"
	ValidationExceptionReason_INVALID_TAG                ValidationExceptionReason = "INVALID_TAG"
	ValidationExceptionReason_INVALID_VOLUME_SIZE        ValidationExceptionReason = "INVALID_VOLUME_SIZE"
	ValidationExceptionReason_UNRELATED_SNAPSHOTS        ValidationExceptionReason = "UNRELATED_SNAPSHOTS"
	ValidationExceptionReason_WRITE_REQUEST_TIMEOUT      ValidationExceptionReason = "WRITE_REQUEST_TIMEOUT"
)
