package operationalinsights

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"net/http"
)

// DataSourceKind enumerates the values for data source kind.
type DataSourceKind string

const (
	// AzureActivityLog specifies the azure activity log state for data source kind.
	AzureActivityLog DataSourceKind = "AzureActivityLog"
	// ChangeTrackingCustomRegistry specifies the change tracking custom registry state for data source kind.
	ChangeTrackingCustomRegistry DataSourceKind = "ChangeTrackingCustomRegistry"
	// ChangeTrackingDefaultPath specifies the change tracking default path state for data source kind.
	ChangeTrackingDefaultPath DataSourceKind = "ChangeTrackingDefaultPath"
	// ChangeTrackingDefaultRegistry specifies the change tracking default registry state for data source kind.
	ChangeTrackingDefaultRegistry DataSourceKind = "ChangeTrackingDefaultRegistry"
	// ChangeTrackingPath specifies the change tracking path state for data source kind.
	ChangeTrackingPath DataSourceKind = "ChangeTrackingPath"
	// CustomLog specifies the custom log state for data source kind.
	CustomLog DataSourceKind = "CustomLog"
	// CustomLogCollection specifies the custom log collection state for data source kind.
	CustomLogCollection DataSourceKind = "CustomLogCollection"
	// GenericDataSource specifies the generic data source state for data source kind.
	GenericDataSource DataSourceKind = "GenericDataSource"
	// IISLogs specifies the iis logs state for data source kind.
	IISLogs DataSourceKind = "IISLogs"
	// LinuxPerformanceCollection specifies the linux performance collection state for data source kind.
	LinuxPerformanceCollection DataSourceKind = "LinuxPerformanceCollection"
	// LinuxPerformanceObject specifies the linux performance object state for data source kind.
	LinuxPerformanceObject DataSourceKind = "LinuxPerformanceObject"
	// LinuxSyslog specifies the linux syslog state for data source kind.
	LinuxSyslog DataSourceKind = "LinuxSyslog"
	// LinuxSyslogCollection specifies the linux syslog collection state for data source kind.
	LinuxSyslogCollection DataSourceKind = "LinuxSyslogCollection"
	// WindowsEvent specifies the windows event state for data source kind.
	WindowsEvent DataSourceKind = "WindowsEvent"
	// WindowsPerformanceCounter specifies the windows performance counter state for data source kind.
	WindowsPerformanceCounter DataSourceKind = "WindowsPerformanceCounter"
)

// EntityStatus enumerates the values for entity status.
type EntityStatus string

const (
	// Canceled specifies the canceled state for entity status.
	Canceled EntityStatus = "Canceled"
	// Creating specifies the creating state for entity status.
	Creating EntityStatus = "Creating"
	// Deleting specifies the deleting state for entity status.
	Deleting EntityStatus = "Deleting"
	// Failed specifies the failed state for entity status.
	Failed EntityStatus = "Failed"
	// ProvisioningAccount specifies the provisioning account state for entity status.
	ProvisioningAccount EntityStatus = "ProvisioningAccount"
	// Succeeded specifies the succeeded state for entity status.
	Succeeded EntityStatus = "Succeeded"
)

// SearchSortEnum enumerates the values for search sort enum.
type SearchSortEnum string

const (
	// Asc specifies the asc state for search sort enum.
	Asc SearchSortEnum = "asc"
	// Desc specifies the desc state for search sort enum.
	Desc SearchSortEnum = "desc"
)

// SkuNameEnum enumerates the values for sku name enum.
type SkuNameEnum string

const (
	// Free specifies the free state for sku name enum.
	Free SkuNameEnum = "Free"
	// PerNode specifies the per node state for sku name enum.
	PerNode SkuNameEnum = "PerNode"
	// Premium specifies the premium state for sku name enum.
	Premium SkuNameEnum = "Premium"
	// Standalone specifies the standalone state for sku name enum.
	Standalone SkuNameEnum = "Standalone"
	// Standard specifies the standard state for sku name enum.
	Standard SkuNameEnum = "Standard"
	// Unlimited specifies the unlimited state for sku name enum.
	Unlimited SkuNameEnum = "Unlimited"
)

// StorageInsightState enumerates the values for storage insight state.
type StorageInsightState string

const (
	// ERROR specifies the error state for storage insight state.
	ERROR StorageInsightState = "ERROR"
	// OK specifies the ok state for storage insight state.
	OK StorageInsightState = "OK"
)

// CoreSummary is the core summary of a search.
type CoreSummary struct {
	Status            *string `json:"Status,omitempty"`
	NumberOfDocuments *int64  `json:"NumberOfDocuments,omitempty"`
}

// DataSource is datasources under OMS Workspace.
type DataSource struct {
	autorest.Response `json:"-"`
	ID                *string                 `json:"id,omitempty"`
	Name              *string                 `json:"name,omitempty"`
	Type              *string                 `json:"type,omitempty"`
	Tags              *map[string]*string     `json:"tags,omitempty"`
	Properties        *map[string]interface{} `json:"properties,omitempty"`
	ETag              *string                 `json:"eTag,omitempty"`
	Kind              DataSourceKind          `json:"kind,omitempty"`
}

// DataSourceFilter is dataSource filter. Right now, only filter by kind is supported.
type DataSourceFilter struct {
	Kind DataSourceKind `json:"kind,omitempty"`
}

// DataSourceListResult is the list data source by workspace operation response.
type DataSourceListResult struct {
	autorest.Response `json:"-"`
	Value             *[]DataSource `json:"value,omitempty"`
	NextLink          *string       `json:"nextLink,omitempty"`
}

// DataSourceListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client DataSourceListResult) DataSourceListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// IntelligencePack is intelligence Pack containing a string name and boolean indicating if it's enabled.
type IntelligencePack struct {
	Name    *string `json:"name,omitempty"`
	Enabled *bool   `json:"enabled,omitempty"`
}

// LinkedService is the top level Linked service resource container.
type LinkedService struct {
	autorest.Response        `json:"-"`
	ID                       *string             `json:"id,omitempty"`
	Name                     *string             `json:"name,omitempty"`
	Type                     *string             `json:"type,omitempty"`
	Tags                     *map[string]*string `json:"tags,omitempty"`
	*LinkedServiceProperties `json:"properties,omitempty"`
}

// LinkedServiceListResult is the list linked service operation response.
type LinkedServiceListResult struct {
	autorest.Response `json:"-"`
	Value             *[]LinkedService `json:"value,omitempty"`
}

// LinkedServiceProperties is linked service properties.
type LinkedServiceProperties struct {
	ResourceID *string `json:"resourceId,omitempty"`
}

// LinkTarget is metadata for a workspace that isn't linked to an Azure subscription.
type LinkTarget struct {
	CustomerID    *string `json:"customerId,omitempty"`
	DisplayName   *string `json:"accountName,omitempty"`
	WorkspaceName *string `json:"workspaceName,omitempty"`
	Location      *string `json:"location,omitempty"`
}

// ListIntelligencePack is
type ListIntelligencePack struct {
	autorest.Response `json:"-"`
	Value             *[]IntelligencePack `json:"value,omitempty"`
}

// ListLinkTarget is
type ListLinkTarget struct {
	autorest.Response `json:"-"`
	Value             *[]LinkTarget `json:"value,omitempty"`
}

// ManagementGroup is a management group that is connected to a workspace
type ManagementGroup struct {
	*ManagementGroupProperties `json:"properties,omitempty"`
}

// ManagementGroupProperties is management group properties.
type ManagementGroupProperties struct {
	ServerCount  *int32     `json:"serverCount,omitempty"`
	IsGateway    *bool      `json:"isGateway,omitempty"`
	Name         *string    `json:"name,omitempty"`
	ID           *string    `json:"id,omitempty"`
	Created      *date.Time `json:"created,omitempty"`
	DataReceived *date.Time `json:"dataReceived,omitempty"`
	Version      *string    `json:"version,omitempty"`
	Sku          *string    `json:"sku,omitempty"`
}

// MetricName is the name of a metric.
type MetricName struct {
	Value          *string `json:"value,omitempty"`
	LocalizedValue *string `json:"localizedValue,omitempty"`
}

// ProxyResource is common properties of proxy resource.
type ProxyResource struct {
	ID   *string             `json:"id,omitempty"`
	Name *string             `json:"name,omitempty"`
	Type *string             `json:"type,omitempty"`
	Tags *map[string]*string `json:"tags,omitempty"`
}

// Resource is the resource definition.
type Resource struct {
	ID       *string             `json:"id,omitempty"`
	Name     *string             `json:"name,omitempty"`
	Type     *string             `json:"type,omitempty"`
	Location *string             `json:"location,omitempty"`
	Tags     *map[string]*string `json:"tags,omitempty"`
}

// SavedSearch is value object for saved search results.
type SavedSearch struct {
	autorest.Response      `json:"-"`
	ID                     *string `json:"id,omitempty"`
	Etag                   *string `json:"etag,omitempty"`
	*SavedSearchProperties `json:"properties,omitempty"`
}

// SavedSearchesListResult is the saved search operation response.
type SavedSearchesListResult struct {
	autorest.Response `json:"-"`
	Metadata          *SearchMetadata `json:"__metadata,omitempty"`
	Value             *[]SavedSearch  `json:"value,omitempty"`
}

// SavedSearchProperties is value object for saved search results.
type SavedSearchProperties struct {
	Category    *string `json:"Category,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty"`
	Query       *string `json:"Query,omitempty"`
	Version     *int64  `json:"Version,omitempty"`
	Tags        *[]Tag  `json:"Tags,omitempty"`
}

// SearchError is details for a search error.
type SearchError struct {
	Type    *string `json:"type,omitempty"`
	Message *string `json:"message,omitempty"`
}

// SearchGetSchemaResponse is the get schema operation response.
type SearchGetSchemaResponse struct {
	autorest.Response `json:"-"`
	Metadata          *SearchMetadata      `json:"__metadata,omitempty"`
	Value             *[]SearchSchemaValue `json:"value,omitempty"`
}

// SearchHighlight is highlight details.
type SearchHighlight struct {
	Pre  *string `json:"pre,omitempty"`
	Post *string `json:"post,omitempty"`
}

// SearchMetadata is metadata for search results.
type SearchMetadata struct {
	SearchID                 *string               `json:"RequestId,omitempty"`
	ResultType               *string               `json:"resultType,omitempty"`
	Total                    *int64                `json:"total,omitempty"`
	Top                      *int64                `json:"top,omitempty"`
	ID                       *string               `json:"id,omitempty"`
	CoreSummaries            *[]CoreSummary        `json:"CoreSummaries,omitempty"`
	Status                   *string               `json:"Status,omitempty"`
	StartTime                *date.Time            `json:"StartTime,omitempty"`
	LastUpdated              *date.Time            `json:"LastUpdated,omitempty"`
	ETag                     *string               `json:"ETag,omitempty"`
	Sort                     *[]SearchSort         `json:"sort,omitempty"`
	RequestTime              *int64                `json:"requestTime,omitempty"`
	AggregatedValueField     *string               `json:"aggregatedValueField,omitempty"`
	AggregatedGroupingFields *string               `json:"aggregatedGroupingFields,omitempty"`
	Sum                      *int64                `json:"sum,omitempty"`
	Max                      *int64                `json:"max,omitempty"`
	Schema                   *SearchMetadataSchema `json:"schema,omitempty"`
}

// SearchMetadataSchema is schema metadata for search.
type SearchMetadataSchema struct {
	Name    *string `json:"name,omitempty"`
	Version *int32  `json:"version,omitempty"`
}

// SearchParameters is parameters specifying the search query and range.
type SearchParameters struct {
	Top       *int64           `json:"top,omitempty"`
	Highlight *SearchHighlight `json:"highlight,omitempty"`
	Query     *string          `json:"query,omitempty"`
	Start     *date.Time       `json:"start,omitempty"`
	End       *date.Time       `json:"end,omitempty"`
}

// SearchResultsResponse is the get search result operation response.
type SearchResultsResponse struct {
	autorest.Response `json:"-"`
	ID                *string                   `json:"id,omitempty"`
	Metadata          *SearchMetadata           `json:"__metadata,omitempty"`
	Value             *[]map[string]interface{} `json:"value,omitempty"`
	Error             *SearchError              `json:"error,omitempty"`
}

// SearchSchemaValue is value object for schema results.
type SearchSchemaValue struct {
	Name        *string   `json:"name,omitempty"`
	DisplayName *string   `json:"displayName,omitempty"`
	Type        *string   `json:"type,omitempty"`
	Indexed     *bool     `json:"indexed,omitempty"`
	Stored      *bool     `json:"stored,omitempty"`
	Facet       *bool     `json:"facet,omitempty"`
	OwnerType   *[]string `json:"ownerType,omitempty"`
}

// SearchSort is the sort parameters for search.
type SearchSort struct {
	Name  *string        `json:"name,omitempty"`
	Order SearchSortEnum `json:"order,omitempty"`
}

// SharedKeys is the shared keys for a workspace.
type SharedKeys struct {
	autorest.Response  `json:"-"`
	PrimarySharedKey   *string `json:"primarySharedKey,omitempty"`
	SecondarySharedKey *string `json:"secondarySharedKey,omitempty"`
}

// Sku is the SKU (tier) of a workspace.
type Sku struct {
	Name SkuNameEnum `json:"name,omitempty"`
}

// StorageAccount is describes a storage account connection.
type StorageAccount struct {
	ID  *string `json:"id,omitempty"`
	Key *string `json:"key,omitempty"`
}

// StorageInsight is the top level storage insight resource container.
type StorageInsight struct {
	autorest.Response         `json:"-"`
	ID                        *string             `json:"id,omitempty"`
	Name                      *string             `json:"name,omitempty"`
	Type                      *string             `json:"type,omitempty"`
	Tags                      *map[string]*string `json:"tags,omitempty"`
	*StorageInsightProperties `json:"properties,omitempty"`
	ETag                      *string `json:"eTag,omitempty"`
}

// StorageInsightListResult is the list storage insights operation response.
type StorageInsightListResult struct {
	autorest.Response `json:"-"`
	Value             *[]StorageInsight `json:"value,omitempty"`
	OdataNextLink     *string           `json:"@odata.nextLink,omitempty"`
}

// StorageInsightListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client StorageInsightListResult) StorageInsightListResultPreparer() (*http.Request, error) {
	if client.OdataNextLink == nil || len(to.String(client.OdataNextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.OdataNextLink)))
}

// StorageInsightProperties is storage insight properties.
type StorageInsightProperties struct {
	Containers     *[]string             `json:"containers,omitempty"`
	Tables         *[]string             `json:"tables,omitempty"`
	StorageAccount *StorageAccount       `json:"storageAccount,omitempty"`
	Status         *StorageInsightStatus `json:"status,omitempty"`
}

// StorageInsightStatus is the status of the storage insight.
type StorageInsightStatus struct {
	State       StorageInsightState `json:"state,omitempty"`
	Description *string             `json:"description,omitempty"`
}

// Tag is a tag of a saved search.
type Tag struct {
	Name  *string `json:"Name,omitempty"`
	Value *string `json:"Value,omitempty"`
}

// UsageMetric is a metric describing the usage of a resource.
type UsageMetric struct {
	Name          *MetricName `json:"name,omitempty"`
	Unit          *string     `json:"unit,omitempty"`
	CurrentValue  *float64    `json:"currentValue,omitempty"`
	Limit         *float64    `json:"limit,omitempty"`
	NextResetTime *date.Time  `json:"nextResetTime,omitempty"`
	QuotaPeriod   *string     `json:"quotaPeriod,omitempty"`
}

// Workspace is the top level Workspace resource container.
type Workspace struct {
	autorest.Response    `json:"-"`
	ID                   *string             `json:"id,omitempty"`
	Name                 *string             `json:"name,omitempty"`
	Type                 *string             `json:"type,omitempty"`
	Location             *string             `json:"location,omitempty"`
	Tags                 *map[string]*string `json:"tags,omitempty"`
	*WorkspaceProperties `json:"properties,omitempty"`
	ETag                 *string `json:"eTag,omitempty"`
}

// WorkspaceListManagementGroupsResult is the list workspace managmement groups operation response.
type WorkspaceListManagementGroupsResult struct {
	autorest.Response `json:"-"`
	Value             *[]ManagementGroup `json:"value,omitempty"`
}

// WorkspaceListResult is the list workspaces operation response.
type WorkspaceListResult struct {
	autorest.Response `json:"-"`
	Value             *[]Workspace `json:"value,omitempty"`
}

// WorkspaceListUsagesResult is the list workspace usages operation response.
type WorkspaceListUsagesResult struct {
	autorest.Response `json:"-"`
	Value             *[]UsageMetric `json:"value,omitempty"`
}

// WorkspaceProperties is workspace properties.
type WorkspaceProperties struct {
	ProvisioningState EntityStatus `json:"provisioningState,omitempty"`
	Source            *string      `json:"source,omitempty"`
	CustomerID        *string      `json:"customerId,omitempty"`
	PortalURL         *string      `json:"portalUrl,omitempty"`
	Sku               *Sku         `json:"sku,omitempty"`
	RetentionInDays   *int32       `json:"retentionInDays,omitempty"`
}
