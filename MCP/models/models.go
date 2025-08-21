package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// ProgressUpdateStreamSummary represents the ProgressUpdateStreamSummary schema from the OpenAPI specification
type ProgressUpdateStreamSummary struct {
	Progressupdatestreamname interface{} `json:"ProgressUpdateStreamName,omitempty"`
}

// ImportMigrationTaskRequest represents the ImportMigrationTaskRequest schema from the OpenAPI specification
type ImportMigrationTaskRequest struct {
	Dryrun interface{} `json:"DryRun,omitempty"`
	Migrationtaskname interface{} `json:"MigrationTaskName"`
	Progressupdatestream interface{} `json:"ProgressUpdateStream"`
}

// MigrationTaskSummary represents the MigrationTaskSummary schema from the OpenAPI specification
type MigrationTaskSummary struct {
	Statusdetail interface{} `json:"StatusDetail,omitempty"`
	Updatedatetime interface{} `json:"UpdateDateTime,omitempty"`
	Migrationtaskname interface{} `json:"MigrationTaskName,omitempty"`
	Progresspercent interface{} `json:"ProgressPercent,omitempty"`
	Progressupdatestream interface{} `json:"ProgressUpdateStream,omitempty"`
	Status interface{} `json:"Status,omitempty"`
}

// AssociateCreatedArtifactResult represents the AssociateCreatedArtifactResult schema from the OpenAPI specification
type AssociateCreatedArtifactResult struct {
}

// MigrationTask represents the MigrationTask schema from the OpenAPI specification
type MigrationTask struct {
	Progressupdatestream interface{} `json:"ProgressUpdateStream,omitempty"`
	Resourceattributelist interface{} `json:"ResourceAttributeList,omitempty"`
	Task interface{} `json:"Task,omitempty"`
	Updatedatetime interface{} `json:"UpdateDateTime,omitempty"`
	Migrationtaskname interface{} `json:"MigrationTaskName,omitempty"`
}

// DescribeMigrationTaskRequest represents the DescribeMigrationTaskRequest schema from the OpenAPI specification
type DescribeMigrationTaskRequest struct {
	Progressupdatestream interface{} `json:"ProgressUpdateStream"`
	Migrationtaskname interface{} `json:"MigrationTaskName"`
}

// ListApplicationStatesRequest represents the ListApplicationStatesRequest schema from the OpenAPI specification
type ListApplicationStatesRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Applicationids interface{} `json:"ApplicationIds,omitempty"`
}

// ListDiscoveredResourcesResult represents the ListDiscoveredResourcesResult schema from the OpenAPI specification
type ListDiscoveredResourcesResult struct {
	Discoveredresourcelist interface{} `json:"DiscoveredResourceList,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// NotifyMigrationTaskStateRequest represents the NotifyMigrationTaskStateRequest schema from the OpenAPI specification
type NotifyMigrationTaskStateRequest struct {
	Task interface{} `json:"Task"`
	Updatedatetime interface{} `json:"UpdateDateTime"`
	Dryrun interface{} `json:"DryRun,omitempty"`
	Migrationtaskname interface{} `json:"MigrationTaskName"`
	Nextupdateseconds interface{} `json:"NextUpdateSeconds"`
	Progressupdatestream interface{} `json:"ProgressUpdateStream"`
}

// AssociateCreatedArtifactRequest represents the AssociateCreatedArtifactRequest schema from the OpenAPI specification
type AssociateCreatedArtifactRequest struct {
	Progressupdatestream interface{} `json:"ProgressUpdateStream"`
	Createdartifact interface{} `json:"CreatedArtifact"`
	Dryrun interface{} `json:"DryRun,omitempty"`
	Migrationtaskname interface{} `json:"MigrationTaskName"`
}

// ListDiscoveredResourcesRequest represents the ListDiscoveredResourcesRequest schema from the OpenAPI specification
type ListDiscoveredResourcesRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Migrationtaskname interface{} `json:"MigrationTaskName"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Progressupdatestream interface{} `json:"ProgressUpdateStream"`
}

// ListCreatedArtifactsResult represents the ListCreatedArtifactsResult schema from the OpenAPI specification
type ListCreatedArtifactsResult struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Createdartifactlist interface{} `json:"CreatedArtifactList,omitempty"`
}

// DiscoveredResource represents the DiscoveredResource schema from the OpenAPI specification
type DiscoveredResource struct {
	Description interface{} `json:"Description,omitempty"`
	Configurationid interface{} `json:"ConfigurationId"`
}

// ApplicationState represents the ApplicationState schema from the OpenAPI specification
type ApplicationState struct {
	Applicationid interface{} `json:"ApplicationId,omitempty"`
	Applicationstatus interface{} `json:"ApplicationStatus,omitempty"`
	Lastupdatedtime interface{} `json:"LastUpdatedTime,omitempty"`
}

// Task represents the Task schema from the OpenAPI specification
type Task struct {
	Statusdetail interface{} `json:"StatusDetail,omitempty"`
	Progresspercent interface{} `json:"ProgressPercent,omitempty"`
	Status interface{} `json:"Status"`
}

// ListMigrationTasksRequest represents the ListMigrationTasksRequest schema from the OpenAPI specification
type ListMigrationTasksRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Resourcename interface{} `json:"ResourceName,omitempty"`
}

// DeleteProgressUpdateStreamResult represents the DeleteProgressUpdateStreamResult schema from the OpenAPI specification
type DeleteProgressUpdateStreamResult struct {
}

// ListProgressUpdateStreamsResult represents the ListProgressUpdateStreamsResult schema from the OpenAPI specification
type ListProgressUpdateStreamsResult struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Progressupdatestreamsummarylist interface{} `json:"ProgressUpdateStreamSummaryList,omitempty"`
}

// ListApplicationStatesResult represents the ListApplicationStatesResult schema from the OpenAPI specification
type ListApplicationStatesResult struct {
	Applicationstatelist interface{} `json:"ApplicationStateList,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// ListMigrationTasksResult represents the ListMigrationTasksResult schema from the OpenAPI specification
type ListMigrationTasksResult struct {
	Migrationtasksummarylist interface{} `json:"MigrationTaskSummaryList,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// DescribeMigrationTaskResult represents the DescribeMigrationTaskResult schema from the OpenAPI specification
type DescribeMigrationTaskResult struct {
	Migrationtask interface{} `json:"MigrationTask,omitempty"`
}

// AssociateDiscoveredResourceResult represents the AssociateDiscoveredResourceResult schema from the OpenAPI specification
type AssociateDiscoveredResourceResult struct {
}

// CreateProgressUpdateStreamResult represents the CreateProgressUpdateStreamResult schema from the OpenAPI specification
type CreateProgressUpdateStreamResult struct {
}

// NotifyApplicationStateResult represents the NotifyApplicationStateResult schema from the OpenAPI specification
type NotifyApplicationStateResult struct {
}

// PutResourceAttributesRequest represents the PutResourceAttributesRequest schema from the OpenAPI specification
type PutResourceAttributesRequest struct {
	Resourceattributelist interface{} `json:"ResourceAttributeList"`
	Dryrun interface{} `json:"DryRun,omitempty"`
	Migrationtaskname interface{} `json:"MigrationTaskName"`
	Progressupdatestream interface{} `json:"ProgressUpdateStream"`
}

// CreatedArtifact represents the CreatedArtifact schema from the OpenAPI specification
type CreatedArtifact struct {
	Name interface{} `json:"Name"`
	Description interface{} `json:"Description,omitempty"`
}

// DeleteProgressUpdateStreamRequest represents the DeleteProgressUpdateStreamRequest schema from the OpenAPI specification
type DeleteProgressUpdateStreamRequest struct {
	Progressupdatestreamname interface{} `json:"ProgressUpdateStreamName"`
	Dryrun interface{} `json:"DryRun,omitempty"`
}

// DisassociateDiscoveredResourceResult represents the DisassociateDiscoveredResourceResult schema from the OpenAPI specification
type DisassociateDiscoveredResourceResult struct {
}

// DescribeApplicationStateResult represents the DescribeApplicationStateResult schema from the OpenAPI specification
type DescribeApplicationStateResult struct {
	Applicationstatus interface{} `json:"ApplicationStatus,omitempty"`
	Lastupdatedtime interface{} `json:"LastUpdatedTime,omitempty"`
}

// ImportMigrationTaskResult represents the ImportMigrationTaskResult schema from the OpenAPI specification
type ImportMigrationTaskResult struct {
}

// CreateProgressUpdateStreamRequest represents the CreateProgressUpdateStreamRequest schema from the OpenAPI specification
type CreateProgressUpdateStreamRequest struct {
	Progressupdatestreamname interface{} `json:"ProgressUpdateStreamName"`
	Dryrun interface{} `json:"DryRun,omitempty"`
}

// DisassociateCreatedArtifactRequest represents the DisassociateCreatedArtifactRequest schema from the OpenAPI specification
type DisassociateCreatedArtifactRequest struct {
	Createdartifactname interface{} `json:"CreatedArtifactName"`
	Dryrun interface{} `json:"DryRun,omitempty"`
	Migrationtaskname interface{} `json:"MigrationTaskName"`
	Progressupdatestream interface{} `json:"ProgressUpdateStream"`
}

// NotifyMigrationTaskStateResult represents the NotifyMigrationTaskStateResult schema from the OpenAPI specification
type NotifyMigrationTaskStateResult struct {
}

// DisassociateCreatedArtifactResult represents the DisassociateCreatedArtifactResult schema from the OpenAPI specification
type DisassociateCreatedArtifactResult struct {
}

// DescribeApplicationStateRequest represents the DescribeApplicationStateRequest schema from the OpenAPI specification
type DescribeApplicationStateRequest struct {
	Applicationid interface{} `json:"ApplicationId"`
}

// ResourceAttribute represents the ResourceAttribute schema from the OpenAPI specification
type ResourceAttribute struct {
	TypeField interface{} `json:"Type"`
	Value interface{} `json:"Value"`
}

// AssociateDiscoveredResourceRequest represents the AssociateDiscoveredResourceRequest schema from the OpenAPI specification
type AssociateDiscoveredResourceRequest struct {
	Discoveredresource interface{} `json:"DiscoveredResource"`
	Dryrun interface{} `json:"DryRun,omitempty"`
	Migrationtaskname interface{} `json:"MigrationTaskName"`
	Progressupdatestream interface{} `json:"ProgressUpdateStream"`
}

// ListProgressUpdateStreamsRequest represents the ListProgressUpdateStreamsRequest schema from the OpenAPI specification
type ListProgressUpdateStreamsRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// ListCreatedArtifactsRequest represents the ListCreatedArtifactsRequest schema from the OpenAPI specification
type ListCreatedArtifactsRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Migrationtaskname interface{} `json:"MigrationTaskName"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Progressupdatestream interface{} `json:"ProgressUpdateStream"`
}

// DisassociateDiscoveredResourceRequest represents the DisassociateDiscoveredResourceRequest schema from the OpenAPI specification
type DisassociateDiscoveredResourceRequest struct {
	Progressupdatestream interface{} `json:"ProgressUpdateStream"`
	Configurationid interface{} `json:"ConfigurationId"`
	Dryrun interface{} `json:"DryRun,omitempty"`
	Migrationtaskname interface{} `json:"MigrationTaskName"`
}

// PutResourceAttributesResult represents the PutResourceAttributesResult schema from the OpenAPI specification
type PutResourceAttributesResult struct {
}

// NotifyApplicationStateRequest represents the NotifyApplicationStateRequest schema from the OpenAPI specification
type NotifyApplicationStateRequest struct {
	Dryrun interface{} `json:"DryRun,omitempty"`
	Status interface{} `json:"Status"`
	Updatedatetime interface{} `json:"UpdateDateTime,omitempty"`
	Applicationid interface{} `json:"ApplicationId"`
}
