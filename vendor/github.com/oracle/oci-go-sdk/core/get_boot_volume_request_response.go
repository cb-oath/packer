// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package core

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetBootVolumeRequest wrapper for the GetBootVolume operation
type GetBootVolumeRequest struct {

	// The OCID of the boot volume.
	BootVolumeId *string `mandatory:"true" contributesTo:"path" name:"bootVolumeId"`
}

func (request GetBootVolumeRequest) String() string {
	return common.PointerString(request)
}

// GetBootVolumeResponse wrapper for the GetBootVolume operation
type GetBootVolumeResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The BootVolume instance
	BootVolume `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetBootVolumeResponse) String() string {
	return common.PointerString(response)
}
