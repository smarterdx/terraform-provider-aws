// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lambda

// Exports for use in tests only.
var (
	ResourceAlias                     = resourceAlias
	ResourceCodeSigningConfig         = resourceCodeSigningConfig
	ResourceEventSourceMapping        = resourceEventSourceMapping
	ResourceFunctionEventInvokeConfig = resourceFunctionEventInvokeConfig
	ResourceFunctionURL               = resourceFunctionURL

	FindAliasByTwoPartKey                     = findAliasByTwoPartKey
	FindCodeSigningConfigByARN                = findCodeSigningConfigByARN
	FindEventSourceMappingByID                = findEventSourceMappingByID
	FindFunctionEventInvokeConfigByTwoPartKey = findFunctionEventInvokeConfigByTwoPartKey
	FunctionEventInvokeConfigParseResourceID  = functionEventInvokeConfigParseResourceID
	FindFunctionURLByTwoPartKey               = findFunctionURLByTwoPartKey
	SignerServiceIsAvailable                  = signerServiceIsAvailable
)
