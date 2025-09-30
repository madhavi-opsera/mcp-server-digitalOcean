package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// GeneratedType_Webhook_milestone_closed represents the GeneratedType_Webhook_milestone_closed schema from the OpenAPI specification
type GeneratedType_Webhook_milestone_closed struct {
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Milestone Webhooksmilestone `json:"milestone"` // A collection of related issues and pull requests.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType_Webhook_deploy_key_created represents the GeneratedType_Webhook_deploy_key_created schema from the OpenAPI specification
type GeneratedType_Webhook_deploy_key_created struct {
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Key Webhooksdeploykey `json:"key"` // The [`deploy key`](https://docs.github.com/rest/deploy-keys/deploy-keys#get-a-deploy-key) resource.
}

// Deployment represents the Deployment schema from the OpenAPI specification
type Deployment struct {
	Production_environment bool `json:"production_environment,omitempty"` // Specifies if the given environment is one that end-users directly interact with. Default: false.
	Sha string `json:"sha"`
	Task string `json:"task"` // Parameter to specify a task to execute
	Description string `json:"description"`
	Performed_via_github_app GeneratedType_Nullable_integration `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Transient_environment bool `json:"transient_environment,omitempty"` // Specifies if the given environment is will no longer exist at some point in the future. Default: false.
	Updated_at string `json:"updated_at"`
	Creator GeneratedType_Nullable_simple_user `json:"creator"` // A GitHub user.
	Statuses_url string `json:"statuses_url"`
	Environment string `json:"environment"` // Name for the target deployment environment.
	Node_id string `json:"node_id"`
	Payload interface{} `json:"payload"`
	Repository_url string `json:"repository_url"`
	Id int64 `json:"id"` // Unique identifier of the deployment
	Original_environment string `json:"original_environment,omitempty"`
	Url string `json:"url"`
	Created_at string `json:"created_at"`
	Ref string `json:"ref"` // The ref to deploy. This can be a branch, tag, or sha.
}

// GeneratedType_Webhook_discussion_pinned represents the GeneratedType_Webhook_discussion_pinned schema from the OpenAPI specification
type GeneratedType_Webhook_discussion_pinned struct {
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType_Oidc_custom_sub represents the GeneratedType_Oidc_custom_sub schema from the OpenAPI specification
type GeneratedType_Oidc_custom_sub struct {
	Include_claim_keys []string `json:"include_claim_keys"` // Array of unique strings. Each claim key can only contain alphanumeric characters and underscores.
}

// GeneratedType_Organization_create_issue_type represents the GeneratedType_Organization_create_issue_type schema from the OpenAPI specification
type GeneratedType_Organization_create_issue_type struct {
	Color string `json:"color,omitempty"` // Color for the issue type.
	Description string `json:"description,omitempty"` // Description of the issue type.
	Is_enabled bool `json:"is_enabled"` // Whether or not the issue type is enabled at the organization level.
	Name string `json:"name"` // Name of the issue type.
}

// GeneratedType_Secret_scanning_location_pull_request_body represents the GeneratedType_Secret_scanning_location_pull_request_body schema from the OpenAPI specification
type GeneratedType_Secret_scanning_location_pull_request_body struct {
	Pull_request_body_url string `json:"pull_request_body_url"` // The API URL to get the pull request where the secret was detected.
}

// GeneratedType_Simple_user represents the GeneratedType_Simple_user schema from the OpenAPI specification
type GeneratedType_Simple_user struct {
	Subscriptions_url string `json:"subscriptions_url"`
	Gravatar_id string `json:"gravatar_id"`
	Url string `json:"url"`
	Organizations_url string `json:"organizations_url"`
	Avatar_url string `json:"avatar_url"`
	Events_url string `json:"events_url"`
	Followers_url string `json:"followers_url"`
	Received_events_url string `json:"received_events_url"`
	User_view_type string `json:"user_view_type,omitempty"`
	Email string `json:"email,omitempty"`
	Starred_at string `json:"starred_at,omitempty"`
	Name string `json:"name,omitempty"`
	Repos_url string `json:"repos_url"`
	Starred_url string `json:"starred_url"`
	Gists_url string `json:"gists_url"`
	Node_id string `json:"node_id"`
	Following_url string `json:"following_url"`
	Html_url string `json:"html_url"`
	Login string `json:"login"`
	Id int64 `json:"id"`
	Site_admin bool `json:"site_admin"`
	TypeField string `json:"type"`
}

// GeneratedType_Code_scanning_variant_analysis_skipped_repo_group represents the GeneratedType_Code_scanning_variant_analysis_skipped_repo_group schema from the OpenAPI specification
type GeneratedType_Code_scanning_variant_analysis_skipped_repo_group struct {
	Repository_count int `json:"repository_count"` // The total number of repositories that were skipped for this reason.
	Repositories []GeneratedType_Code_scanning_variant_analysis_repository `json:"repositories"` // A list of repositories that were skipped. This list may not include all repositories that were skipped. This is only available when the repository was found and the user has access to it.
}

// GeneratedType_Webhook_projects_v2_item_restored represents the GeneratedType_Webhook_projects_v2_item_restored schema from the OpenAPI specification
type GeneratedType_Webhook_projects_v2_item_restored struct {
	Projects_v2_item GeneratedType_Projects_v2_item `json:"projects_v2_item"` // An item belonging to a project
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes Webhooksprojectchanges `json:"changes"`
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType_Org_repo_custom_property_values represents the GeneratedType_Org_repo_custom_property_values schema from the OpenAPI specification
type GeneratedType_Org_repo_custom_property_values struct {
	Properties []GeneratedType_Custom_property_value `json:"properties"` // List of custom property names and associated values
	Repository_full_name string `json:"repository_full_name"`
	Repository_id int `json:"repository_id"`
	Repository_name string `json:"repository_name"`
}

// GeneratedType_Short_branch represents the GeneratedType_Short_branch schema from the OpenAPI specification
type GeneratedType_Short_branch struct {
	Commit map[string]interface{} `json:"commit"`
	Name string `json:"name"`
	Protected bool `json:"protected"`
	Protection GeneratedType_Branch_protection `json:"protection,omitempty"` // Branch Protection
	Protection_url string `json:"protection_url,omitempty"`
}

// GeneratedType_Deployment_protection_rule represents the GeneratedType_Deployment_protection_rule schema from the OpenAPI specification
type GeneratedType_Deployment_protection_rule struct {
	Node_id string `json:"node_id"` // The node ID for the deployment protection rule.
	App GeneratedType_Custom_deployment_rule_app `json:"app"` // A GitHub App that is providing a custom deployment protection rule.
	Enabled bool `json:"enabled"` // Whether the deployment protection rule is enabled for the environment.
	Id int `json:"id"` // The unique identifier for the deployment protection rule.
}

// GeneratedType_Repository_rule_params_workflow_file_reference represents the GeneratedType_Repository_rule_params_workflow_file_reference schema from the OpenAPI specification
type GeneratedType_Repository_rule_params_workflow_file_reference struct {
	Path string `json:"path"` // The path to the workflow file
	Ref string `json:"ref,omitempty"` // The ref (branch or tag) of the workflow file to use
	Repository_id int `json:"repository_id"` // The ID of the repository where the workflow is defined
	Sha string `json:"sha,omitempty"` // The commit SHA of the workflow file to use
}

// GeneratedType_Repository_rule_required_status_checks represents the GeneratedType_Repository_rule_required_status_checks schema from the OpenAPI specification
type GeneratedType_Repository_rule_required_status_checks struct {
	TypeField string `json:"type"`
	Parameters map[string]interface{} `json:"parameters,omitempty"`
}

// GeneratedType_Webhook_fork represents the GeneratedType_Webhook_fork schema from the OpenAPI specification
type GeneratedType_Webhook_fork struct {
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Forkee interface{} `json:"forkee"` // The created [`repository`](https://docs.github.com/rest/repos/repos#get-a-repository) resource.
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// Project represents the Project schema from the OpenAPI specification
type Project struct {
	Html_url string `json:"html_url"`
	Node_id string `json:"node_id"`
	Organization_permission string `json:"organization_permission,omitempty"` // The baseline permission that all organization members have on this project. Only present if owner is an organization.
	Body string `json:"body"` // Body of the project
	Id int `json:"id"`
	Url string `json:"url"`
	Updated_at string `json:"updated_at"`
	Created_at string `json:"created_at"`
	Private bool `json:"private,omitempty"` // Whether or not this project can be seen by everyone. Only present if owner is an organization.
	Number int `json:"number"`
	Owner_url string `json:"owner_url"`
	Creator GeneratedType_Nullable_simple_user `json:"creator"` // A GitHub user.
	State string `json:"state"` // State of the project; either 'open' or 'closed'
	Name string `json:"name"` // Name of the project
	Columns_url string `json:"columns_url"`
}

// GeneratedType_Webhook_discussion_category_changed represents the GeneratedType_Webhook_discussion_category_changed schema from the OpenAPI specification
type GeneratedType_Webhook_discussion_category_changed struct {
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
}

// GeneratedType_Webhook_check_run_rerequested_form_encoded represents the GeneratedType_Webhook_check_run_rerequested_form_encoded schema from the OpenAPI specification
type GeneratedType_Webhook_check_run_rerequested_form_encoded struct {
	Payload string `json:"payload"` // A URL-encoded string of the check_run.rerequested JSON payload. The decoded payload is a JSON object.
}

// Webhooksissue2 represents the Webhooksissue2 schema from the OpenAPI specification
type Webhooksissue2 struct {
	Labels_url string `json:"labels_url"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	User map[string]interface{} `json:"user"`
	Locked bool `json:"locked,omitempty"`
	Number int `json:"number"`
	Performed_via_github_app map[string]interface{} `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Draft bool `json:"draft,omitempty"`
	Sub_issues_summary map[string]interface{} `json:"sub_issues_summary,omitempty"`
	Repository_url string `json:"repository_url"`
	Labels []map[string]interface{} `json:"labels,omitempty"`
	Milestone map[string]interface{} `json:"milestone"` // A collection of related issues and pull requests.
	State string `json:"state,omitempty"` // State of the issue; either 'open' or 'closed'
	State_reason string `json:"state_reason,omitempty"`
	Assignees []map[string]interface{} `json:"assignees"`
	Node_id string `json:"node_id"`
	Timeline_url string `json:"timeline_url,omitempty"`
	TypeField GeneratedType_Issue_type `json:"type,omitempty"` // The type of issue.
	Comments_url string `json:"comments_url"`
	Url string `json:"url"` // URL for the issue
	Comments int `json:"comments"`
	Updated_at string `json:"updated_at"`
	Reactions map[string]interface{} `json:"reactions"`
	Closed_at string `json:"closed_at"`
	Title string `json:"title"` // Title of the issue
	Created_at string `json:"created_at"`
	Assignee map[string]interface{} `json:"assignee,omitempty"`
	Pull_request map[string]interface{} `json:"pull_request,omitempty"`
	Body string `json:"body"` // Contents of the issue
	Id int64 `json:"id"`
	Events_url string `json:"events_url"`
	Html_url string `json:"html_url"`
	Active_lock_reason string `json:"active_lock_reason"`
}

// GeneratedType_Webhook_ping_form_encoded represents the GeneratedType_Webhook_ping_form_encoded schema from the OpenAPI specification
type GeneratedType_Webhook_ping_form_encoded struct {
	Payload string `json:"payload"` // A URL-encoded string of the ping JSON payload. The decoded payload is a JSON object.
}

// GeneratedType_Webhook_issues_typed represents the GeneratedType_Webhook_issues_typed schema from the OpenAPI specification
type GeneratedType_Webhook_issues_typed struct {
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue Webhooksissue `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	TypeField GeneratedType_Issue_type `json:"type"` // The type of issue.
	Action string `json:"action"`
}

// Status represents the Status schema from the OpenAPI specification
type Status struct {
	State string `json:"state"`
	Url string `json:"url"`
	Target_url string `json:"target_url"`
	Updated_at string `json:"updated_at"`
	Avatar_url string `json:"avatar_url"`
	Description string `json:"description"`
	Node_id string `json:"node_id"`
	Context string `json:"context"`
	Creator GeneratedType_Nullable_simple_user `json:"creator"` // A GitHub user.
	Created_at string `json:"created_at"`
	Id int `json:"id"`
}

// GeneratedType_Codespaces_user_public_key represents the GeneratedType_Codespaces_user_public_key schema from the OpenAPI specification
type GeneratedType_Codespaces_user_public_key struct {
	Key string `json:"key"` // The Base64 encoded public key.
	Key_id string `json:"key_id"` // The identifier for the key.
}

// GeneratedType_Nullable_codespace_machine represents the GeneratedType_Nullable_codespace_machine schema from the OpenAPI specification
type GeneratedType_Nullable_codespace_machine struct {
	Storage_in_bytes int `json:"storage_in_bytes"` // How much storage is available to the codespace.
	Cpus int `json:"cpus"` // How many cores are available to the codespace.
	Display_name string `json:"display_name"` // The display name of the machine includes cores, memory, and storage.
	Memory_in_bytes int `json:"memory_in_bytes"` // How much memory is available to the codespace.
	Name string `json:"name"` // The name of the machine.
	Operating_system string `json:"operating_system"` // The operating system of the machine.
	Prebuild_availability string `json:"prebuild_availability"` // Whether a prebuild is currently available when creating a codespace for this machine and repository. If a branch was not specified as a ref, the default branch will be assumed. Value will be "null" if prebuilds are not supported or prebuild availability could not be determined. Value will be "none" if no prebuild is available. Latest values "ready" and "in_progress" indicate the prebuild availability status.
}

// Workflow represents the Workflow schema from the OpenAPI specification
type Workflow struct {
	State string `json:"state"`
	Badge_url string `json:"badge_url"`
	Deleted_at string `json:"deleted_at,omitempty"`
	Path string `json:"path"`
	Url string `json:"url"`
	Html_url string `json:"html_url"`
	Node_id string `json:"node_id"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
	Id int `json:"id"`
	Name string `json:"name"`
}

// GeneratedType_Organization_secret_scanning_alert represents the GeneratedType_Organization_secret_scanning_alert schema from the OpenAPI specification
type GeneratedType_Organization_secret_scanning_alert struct {
	Secret_type string `json:"secret_type,omitempty"` // The type of secret that secret scanning detected.
	Updated_at string `json:"updated_at,omitempty"` // The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Validity string `json:"validity,omitempty"` // The token status as of the latest validity check.
	Resolved_at string `json:"resolved_at,omitempty"` // The time that the alert was resolved in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Locations_url string `json:"locations_url,omitempty"` // The REST API URL of the code locations for this alert.
	Push_protection_bypass_request_comment string `json:"push_protection_bypass_request_comment,omitempty"` // An optional comment when requesting a push protection bypass.
	Html_url string `json:"html_url,omitempty"` // The GitHub URL of the alert resource.
	Push_protection_bypass_request_reviewer GeneratedType_Nullable_simple_user `json:"push_protection_bypass_request_reviewer,omitempty"` // A GitHub user.
	Resolution_comment string `json:"resolution_comment,omitempty"` // The comment that was optionally added when this alert was closed
	Push_protection_bypassed_at string `json:"push_protection_bypassed_at,omitempty"` // The time that push protection was bypassed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Push_protection_bypassed_by GeneratedType_Nullable_simple_user `json:"push_protection_bypassed_by,omitempty"` // A GitHub user.
	Has_more_locations bool `json:"has_more_locations,omitempty"` // A boolean value representing whether or not the token in the alert was detected in more than one location.
	Multi_repo bool `json:"multi_repo,omitempty"` // Whether the detected secret was found in multiple repositories in the same organization or enterprise.
	Resolution string `json:"resolution,omitempty"` // **Required when the `state` is `resolved`.** The reason for resolving the alert.
	Push_protection_bypassed bool `json:"push_protection_bypassed,omitempty"` // Whether push protection was bypassed for the detected secret.
	Resolved_by GeneratedType_Nullable_simple_user `json:"resolved_by,omitempty"` // A GitHub user.
	Created_at string `json:"created_at,omitempty"` // The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Number int `json:"number,omitempty"` // The security alert number.
	Secret_type_display_name string `json:"secret_type_display_name,omitempty"` // User-friendly name for the detected secret, matching the `secret_type`. For a list of built-in patterns, see "[Supported secret scanning patterns](https://docs.github.com/code-security/secret-scanning/introduction/supported-secret-scanning-patterns#supported-secrets)."
	First_location_detected interface{} `json:"first_location_detected,omitempty"` // Details on the location where the token was initially detected. This can be a commit, wiki commit, issue, discussion, pull request.
	Is_base64_encoded bool `json:"is_base64_encoded,omitempty"` // A boolean value representing whether or not alert is base64 encoded
	Publicly_leaked bool `json:"publicly_leaked,omitempty"` // Whether the secret was publicly leaked.
	Push_protection_bypass_request_reviewer_comment string `json:"push_protection_bypass_request_reviewer_comment,omitempty"` // An optional comment when reviewing a push protection bypass.
	State string `json:"state,omitempty"` // Sets the state of the secret scanning alert. You must provide `resolution` when you set the state to `resolved`.
	Repository GeneratedType_Simple_repository `json:"repository,omitempty"` // A GitHub repository.
	Secret string `json:"secret,omitempty"` // The secret that was detected.
	Push_protection_bypass_request_html_url string `json:"push_protection_bypass_request_html_url,omitempty"` // The URL to a push protection bypass request.
	Url string `json:"url,omitempty"` // The REST API URL of the alert resource.
}

// GeneratedType_Webhook_discussion_comment_edited represents the GeneratedType_Webhook_discussion_comment_edited schema from the OpenAPI specification
type GeneratedType_Webhook_discussion_comment_edited struct {
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Changes map[string]interface{} `json:"changes"`
	Comment Webhookscomment `json:"comment"`
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Action string `json:"action"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
}

// GeneratedType_Added_to_project_issue_event represents the GeneratedType_Added_to_project_issue_event schema from the OpenAPI specification
type GeneratedType_Added_to_project_issue_event struct {
	Created_at string `json:"created_at"`
	Event string `json:"event"`
	Url string `json:"url"`
	Commit_url string `json:"commit_url"`
	Id int `json:"id"`
	Project_card map[string]interface{} `json:"project_card,omitempty"`
	Commit_id string `json:"commit_id"`
	Node_id string `json:"node_id"`
	Performed_via_github_app GeneratedType_Nullable_integration `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Actor GeneratedType_Simple_user `json:"actor"` // A GitHub user.
}

// Event represents the Event schema from the OpenAPI specification
type Event struct {
	Created_at string `json:"created_at"`
	Id string `json:"id"`
	Org Actor `json:"org,omitempty"` // Actor
	Payload map[string]interface{} `json:"payload"`
	Public bool `json:"public"`
	Repo map[string]interface{} `json:"repo"`
	TypeField string `json:"type"`
	Actor Actor `json:"actor"` // Actor
}

// GeneratedType_Secret_scanning_location_issue_title represents the GeneratedType_Secret_scanning_location_issue_title schema from the OpenAPI specification
type GeneratedType_Secret_scanning_location_issue_title struct {
	Issue_title_url string `json:"issue_title_url"` // The API URL to get the issue where the secret was detected.
}

// GeneratedType_License_content represents the GeneratedType_License_content schema from the OpenAPI specification
type GeneratedType_License_content struct {
	Html_url string `json:"html_url"`
	Size int `json:"size"`
	TypeField string `json:"type"`
	Download_url string `json:"download_url"`
	Encoding string `json:"encoding"`
	Name string `json:"name"`
	Sha string `json:"sha"`
	Path string `json:"path"`
	Url string `json:"url"`
	Links map[string]interface{} `json:"_links"`
	Content string `json:"content"`
	Git_url string `json:"git_url"`
	License GeneratedType_Nullable_license_simple `json:"license"` // License Simple
}

// Verification represents the Verification schema from the OpenAPI specification
type Verification struct {
	Reason string `json:"reason"`
	Signature string `json:"signature"`
	Verified bool `json:"verified"`
	Verified_at string `json:"verified_at"`
	Payload string `json:"payload"`
}

// GeneratedType_Secret_scanning_alert represents the GeneratedType_Secret_scanning_alert schema from the OpenAPI specification
type GeneratedType_Secret_scanning_alert struct {
	Html_url string `json:"html_url,omitempty"` // The GitHub URL of the alert resource.
	Is_base64_encoded bool `json:"is_base64_encoded,omitempty"` // A boolean value representing whether or not alert is base64 encoded
	Push_protection_bypassed_at string `json:"push_protection_bypassed_at,omitempty"` // The time that push protection was bypassed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Secret_type_display_name string `json:"secret_type_display_name,omitempty"` // User-friendly name for the detected secret, matching the `secret_type`. For a list of built-in patterns, see "[Supported secret scanning patterns](https://docs.github.com/code-security/secret-scanning/introduction/supported-secret-scanning-patterns#supported-secrets)."
	Validity string `json:"validity,omitempty"` // The token status as of the latest validity check.
	Resolved_at string `json:"resolved_at,omitempty"` // The time that the alert was resolved in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	State string `json:"state,omitempty"` // Sets the state of the secret scanning alert. You must provide `resolution` when you set the state to `resolved`.
	Push_protection_bypass_request_html_url string `json:"push_protection_bypass_request_html_url,omitempty"` // The URL to a push protection bypass request.
	Secret_type string `json:"secret_type,omitempty"` // The type of secret that secret scanning detected.
	First_location_detected interface{} `json:"first_location_detected,omitempty"` // Details on the location where the token was initially detected. This can be a commit, wiki commit, issue, discussion, pull request.
	Push_protection_bypass_request_reviewer GeneratedType_Nullable_simple_user `json:"push_protection_bypass_request_reviewer,omitempty"` // A GitHub user.
	Multi_repo bool `json:"multi_repo,omitempty"` // Whether the detected secret was found in multiple repositories under the same organization or enterprise.
	Created_at string `json:"created_at,omitempty"` // The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Push_protection_bypassed_by GeneratedType_Nullable_simple_user `json:"push_protection_bypassed_by,omitempty"` // A GitHub user.
	Publicly_leaked bool `json:"publicly_leaked,omitempty"` // Whether the detected secret was publicly leaked.
	Resolution_comment string `json:"resolution_comment,omitempty"` // An optional comment to resolve an alert.
	Locations_url string `json:"locations_url,omitempty"` // The REST API URL of the code locations for this alert.
	Push_protection_bypass_request_comment string `json:"push_protection_bypass_request_comment,omitempty"` // An optional comment when requesting a push protection bypass.
	Secret string `json:"secret,omitempty"` // The secret that was detected.
	Updated_at string `json:"updated_at,omitempty"` // The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Has_more_locations bool `json:"has_more_locations,omitempty"` // A boolean value representing whether or not the token in the alert was detected in more than one location.
	Number int `json:"number,omitempty"` // The security alert number.
	Push_protection_bypass_request_reviewer_comment string `json:"push_protection_bypass_request_reviewer_comment,omitempty"` // An optional comment when reviewing a push protection bypass.
	Push_protection_bypassed bool `json:"push_protection_bypassed,omitempty"` // Whether push protection was bypassed for the detected secret.
	Resolution string `json:"resolution,omitempty"` // **Required when the `state` is `resolved`.** The reason for resolving the alert.
	Resolved_by GeneratedType_Nullable_simple_user `json:"resolved_by,omitempty"` // A GitHub user.
	Url string `json:"url,omitempty"` // The REST API URL of the alert resource.
}

// GeneratedType_Actions_organization_permissions represents the GeneratedType_Actions_organization_permissions schema from the OpenAPI specification
type GeneratedType_Actions_organization_permissions struct {
	Enabled_repositories string `json:"enabled_repositories"` // The policy that controls the repositories in the organization that are allowed to run GitHub Actions.
	Selected_actions_url string `json:"selected_actions_url,omitempty"` // The API URL to use to get or set the actions and reusable workflows that are allowed to run, when `allowed_actions` is set to `selected`.
	Selected_repositories_url string `json:"selected_repositories_url,omitempty"` // The API URL to use to get or set the selected repositories that are allowed to run GitHub Actions, when `enabled_repositories` is set to `selected`.
	Allowed_actions string `json:"allowed_actions,omitempty"` // The permissions policy that controls the actions and reusable workflows that are allowed to run.
}

// GeneratedType_Repository_rule_params_code_scanning_tool represents the GeneratedType_Repository_rule_params_code_scanning_tool schema from the OpenAPI specification
type GeneratedType_Repository_rule_params_code_scanning_tool struct {
	Security_alerts_threshold string `json:"security_alerts_threshold"` // The severity level at which code scanning results that raise security alerts block a reference update. For more information on security severity levels, see "[About code scanning alerts](https://docs.github.com/code-security/code-scanning/managing-code-scanning-alerts/about-code-scanning-alerts#about-alert-severity-and-security-severity-levels)."
	Tool string `json:"tool"` // The name of a code scanning tool
	Alerts_threshold string `json:"alerts_threshold"` // The severity level at which code scanning results that raise alerts block a reference update. For more information on alert severity levels, see "[About code scanning alerts](https://docs.github.com/code-security/code-scanning/managing-code-scanning-alerts/about-code-scanning-alerts#about-alert-severity-and-security-severity-levels)."
}

// GeneratedType_Timeline_commit_commented_event represents the GeneratedType_Timeline_commit_commented_event schema from the OpenAPI specification
type GeneratedType_Timeline_commit_commented_event struct {
	Commit_id string `json:"commit_id,omitempty"`
	Event string `json:"event,omitempty"`
	Node_id string `json:"node_id,omitempty"`
	Comments []GeneratedType_Commit_comment `json:"comments,omitempty"`
}

// GeneratedType_Sub_issues_summary represents the GeneratedType_Sub_issues_summary schema from the OpenAPI specification
type GeneratedType_Sub_issues_summary struct {
	Completed int `json:"completed"`
	Percent_completed int `json:"percent_completed"`
	Total int `json:"total"`
}

// GeneratedType_Code_of_conduct_simple represents the GeneratedType_Code_of_conduct_simple schema from the OpenAPI specification
type GeneratedType_Code_of_conduct_simple struct {
	Key string `json:"key"`
	Name string `json:"name"`
	Url string `json:"url"`
	Html_url string `json:"html_url"`
}

// GeneratedType_Webhook_project_card_deleted represents the GeneratedType_Webhook_project_card_deleted schema from the OpenAPI specification
type GeneratedType_Webhook_project_card_deleted struct {
	Repository GeneratedType_Nullable_repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Project_card map[string]interface{} `json:"project_card"`
}

// GeneratedType_Webhook_issues_unassigned represents the GeneratedType_Webhook_issues_unassigned schema from the OpenAPI specification
type GeneratedType_Webhook_issues_unassigned struct {
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"` // The action that was performed.
	Assignee Webhooksusermannequin `json:"assignee,omitempty"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue Webhooksissue `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType_Simple_classroom represents the GeneratedType_Simple_classroom schema from the OpenAPI specification
type GeneratedType_Simple_classroom struct {
	Name string `json:"name"` // The name of the classroom.
	Url string `json:"url"` // The url of the classroom on GitHub Classroom.
	Archived bool `json:"archived"` // Returns whether classroom is archived or not.
	Id int `json:"id"` // Unique identifier of the classroom.
}

// GeneratedType_Webhook_issues_reopened represents the GeneratedType_Webhook_issues_reopened schema from the OpenAPI specification
type GeneratedType_Webhook_issues_reopened struct {
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType_Content_file represents the GeneratedType_Content_file schema from the OpenAPI specification
type GeneratedType_Content_file struct {
	Html_url string `json:"html_url"`
	Sha string `json:"sha"`
	Git_url string `json:"git_url"`
	Path string `json:"path"`
	TypeField string `json:"type"`
	Download_url string `json:"download_url"`
	Submodule_git_url string `json:"submodule_git_url,omitempty"`
	Url string `json:"url"`
	Links map[string]interface{} `json:"_links"`
	Encoding string `json:"encoding"`
	Name string `json:"name"`
	Size int `json:"size"`
	Target string `json:"target,omitempty"`
	Content string `json:"content"`
}

// GeneratedType_Webhook_discussion_reopened represents the GeneratedType_Webhook_discussion_reopened schema from the OpenAPI specification
type GeneratedType_Webhook_discussion_reopened struct {
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType_Webhook_organization_member_invited represents the GeneratedType_Webhook_organization_member_invited schema from the OpenAPI specification
type GeneratedType_Webhook_organization_member_invited struct {
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	User Webhooksuser `json:"user,omitempty"`
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Invitation map[string]interface{} `json:"invitation"` // The invitation for the user or email if the action is `member_invited`.
}

// GeneratedType_Clone_traffic represents the GeneratedType_Clone_traffic schema from the OpenAPI specification
type GeneratedType_Clone_traffic struct {
	Clones []Traffic `json:"clones"`
	Count int `json:"count"`
	Uniques int `json:"uniques"`
}

// GeneratedType_Nullable_integration represents the GeneratedType_Nullable_integration schema from the OpenAPI specification
type GeneratedType_Nullable_integration struct {
	Slug string `json:"slug,omitempty"` // The slug name of the GitHub app
	Created_at string `json:"created_at"`
	Node_id string `json:"node_id"`
	Client_id string `json:"client_id,omitempty"`
	Description string `json:"description"`
	External_url string `json:"external_url"`
	Owner interface{} `json:"owner"`
	Events []string `json:"events"` // The list of events for the GitHub app. Note that the `installation_target`, `security_advisory`, and `meta` events are not included because they are global events and not specific to an installation.
	Html_url string `json:"html_url"`
	Id int `json:"id"` // Unique identifier of the GitHub app
	Installations_count int `json:"installations_count,omitempty"` // The number of installations associated with the GitHub app. Only returned when the integration is requesting details about itself.
	Name string `json:"name"` // The name of the GitHub app
	Updated_at string `json:"updated_at"`
	Permissions map[string]interface{} `json:"permissions"` // The set of permissions for the GitHub app
}

// GeneratedType_Webhook_personal_access_token_request_created represents the GeneratedType_Webhook_personal_access_token_request_created schema from the OpenAPI specification
type GeneratedType_Webhook_personal_access_token_request_created struct {
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Personal_access_token_request GeneratedType_Personal_access_token_request `json:"personal_access_token_request"` // Details of a Personal Access Token Request.
}

// GeneratedType_Webhook_star_created represents the GeneratedType_Webhook_star_created schema from the OpenAPI specification
type GeneratedType_Webhook_star_created struct {
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Starred_at string `json:"starred_at"` // The time the star was created. This is a timestamp in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`. Will be `null` for the `deleted` action.
}

// GeneratedType_Secret_scanning_scan represents the GeneratedType_Secret_scanning_scan schema from the OpenAPI specification
type GeneratedType_Secret_scanning_scan struct {
	Completed_at string `json:"completed_at,omitempty"` // The time that the scan was completed. Empty if the scan is running
	Started_at string `json:"started_at,omitempty"` // The time that the scan was started. Empty if the scan is pending
	Status string `json:"status,omitempty"` // The state of the scan. Either "completed", "running", or "pending"
	TypeField string `json:"type,omitempty"` // The type of scan
}

// GeneratedType_Webhook_team_added_to_repository represents the GeneratedType_Webhook_team_added_to_repository schema from the OpenAPI specification
type GeneratedType_Webhook_team_added_to_repository struct {
	Team Webhooksteam1 `json:"team"` // Groups of organization members that gives permissions on specified repositories.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository map[string]interface{} `json:"repository,omitempty"` // A git repository
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
}

// Stargazer represents the Stargazer schema from the OpenAPI specification
type Stargazer struct {
	User GeneratedType_Nullable_simple_user `json:"user"` // A GitHub user.
	Starred_at string `json:"starred_at"`
}

// GeneratedType_Code_scanning_default_setup_options represents the GeneratedType_Code_scanning_default_setup_options schema from the OpenAPI specification
type GeneratedType_Code_scanning_default_setup_options struct {
	Runner_type string `json:"runner_type,omitempty"` // Whether to use labeled runners or standard GitHub runners.
	Runner_label string `json:"runner_label,omitempty"` // The label of the runner to use for code scanning default setup when runner_type is 'labeled'.
}

// GeneratedType_Code_scanning_default_setup_update represents the GeneratedType_Code_scanning_default_setup_update schema from the OpenAPI specification
type GeneratedType_Code_scanning_default_setup_update struct {
	Query_suite string `json:"query_suite,omitempty"` // CodeQL query suite to be used.
	Runner_label string `json:"runner_label,omitempty"` // Runner label to be used if the runner type is labeled.
	Runner_type string `json:"runner_type,omitempty"` // Runner type to be used.
	State string `json:"state,omitempty"` // The desired state of code scanning default setup.
	Threat_model string `json:"threat_model,omitempty"` // Threat model to be used for code scanning analysis. Use `remote` to analyze only network sources and `remote_and_local` to include local sources like filesystem access, command-line arguments, database reads, environment variable and standard input.
	Languages []string `json:"languages,omitempty"` // CodeQL languages to be analyzed.
}

// GeneratedType_Branch_short represents the GeneratedType_Branch_short schema from the OpenAPI specification
type GeneratedType_Branch_short struct {
	Name string `json:"name"`
	Protected bool `json:"protected"`
	Commit map[string]interface{} `json:"commit"`
}

// GeneratedType_Webhook_repository_ruleset_deleted represents the GeneratedType_Webhook_repository_ruleset_deleted schema from the OpenAPI specification
type GeneratedType_Webhook_repository_ruleset_deleted struct {
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Repository_ruleset GeneratedType_Repository_ruleset `json:"repository_ruleset"` // A set of rules to apply when specified conditions are met.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType_Repository_rule_violation_error represents the GeneratedType_Repository_rule_violation_error schema from the OpenAPI specification
type GeneratedType_Repository_rule_violation_error struct {
	Status string `json:"status,omitempty"`
	Documentation_url string `json:"documentation_url,omitempty"`
	Message string `json:"message,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// GeneratedType_Timeline_reviewed_event represents the GeneratedType_Timeline_reviewed_event schema from the OpenAPI specification
type GeneratedType_Timeline_reviewed_event struct {
	Body_text string `json:"body_text,omitempty"`
	Id int `json:"id"` // Unique identifier of the review
	State string `json:"state"`
	Links map[string]interface{} `json:"_links"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Body string `json:"body"` // The text of the review.
	Event string `json:"event"`
	Html_url string `json:"html_url"`
	Pull_request_url string `json:"pull_request_url"`
	Submitted_at string `json:"submitted_at,omitempty"`
	User GeneratedType_Simple_user `json:"user"` // A GitHub user.
	Body_html string `json:"body_html,omitempty"`
	Commit_id string `json:"commit_id"` // A commit SHA for the review.
	Node_id string `json:"node_id"`
}

// GeneratedType_Nullable_community_health_file represents the GeneratedType_Nullable_community_health_file schema from the OpenAPI specification
type GeneratedType_Nullable_community_health_file struct {
	Html_url string `json:"html_url"`
	Url string `json:"url"`
}

// GeneratedType_Simple_classroom_assignment represents the GeneratedType_Simple_classroom_assignment schema from the OpenAPI specification
type GeneratedType_Simple_classroom_assignment struct {
	Invite_link string `json:"invite_link"` // The link that a student can use to accept the assignment.
	Classroom GeneratedType_Simple_classroom `json:"classroom"` // A GitHub Classroom classroom
	Editor string `json:"editor"` // The selected editor for the assignment.
	Language string `json:"language"` // The programming language used in the assignment.
	Invitations_enabled bool `json:"invitations_enabled"` // Whether the invitation link is enabled. Visiting an enabled invitation link will accept the assignment.
	Slug string `json:"slug"` // Sluggified name of the assignment.
	Max_members int `json:"max_members,omitempty"` // The maximum allowable members per team.
	Students_are_repo_admins bool `json:"students_are_repo_admins"` // Whether students are admins on created repository on accepted assignment.
	Deadline string `json:"deadline"` // The time at which the assignment is due.
	Feedback_pull_requests_enabled bool `json:"feedback_pull_requests_enabled"` // Whether feedback pull request will be created on assignment acceptance.
	Max_teams int `json:"max_teams,omitempty"` // The maximum allowable teams for the assignment.
	Accepted int `json:"accepted"` // The number of students that have accepted the assignment.
	Passing int `json:"passing"` // The number of students that have passed the assignment.
	Public_repo bool `json:"public_repo"` // Whether an accepted assignment creates a public repository.
	TypeField string `json:"type"` // Whether it's a Group Assignment or Individual Assignment.
	Id int `json:"id"` // Unique identifier of the repository.
	Submitted int `json:"submitted"` // The number of students that have submitted the assignment.
	Title string `json:"title"` // Assignment title.
}

// GeneratedType_Webhook_code_scanning_alert_reopened_by_user represents the GeneratedType_Webhook_code_scanning_alert_reopened_by_user schema from the OpenAPI specification
type GeneratedType_Webhook_code_scanning_alert_reopened_by_user struct {
	Alert map[string]interface{} `json:"alert"` // The code scanning alert involved in the event.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Ref string `json:"ref"` // The Git reference of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Commit_oid string `json:"commit_oid"` // The commit SHA of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// Root represents the Root schema from the OpenAPI specification
type Root struct {
	Issue_search_url string `json:"issue_search_url"`
	Following_url string `json:"following_url"`
	User_url string `json:"user_url"`
	Organization_teams_url string `json:"organization_teams_url"`
	Repository_url string `json:"repository_url"`
	Public_gists_url string `json:"public_gists_url"`
	Emojis_url string `json:"emojis_url"`
	Keys_url string `json:"keys_url"`
	User_repositories_url string `json:"user_repositories_url"`
	Current_user_repositories_url string `json:"current_user_repositories_url"`
	Hub_url string `json:"hub_url,omitempty"`
	Rate_limit_url string `json:"rate_limit_url"`
	Current_user_authorizations_html_url string `json:"current_user_authorizations_html_url"`
	Authorizations_url string `json:"authorizations_url"`
	Feeds_url string `json:"feeds_url"`
	Starred_url string `json:"starred_url"`
	Gists_url string `json:"gists_url"`
	Commit_search_url string `json:"commit_search_url"`
	Events_url string `json:"events_url"`
	Label_search_url string `json:"label_search_url"`
	Starred_gists_url string `json:"starred_gists_url"`
	Topic_search_url string `json:"topic_search_url,omitempty"`
	User_search_url string `json:"user_search_url"`
	Issues_url string `json:"issues_url"`
	Emails_url string `json:"emails_url"`
	Organization_repositories_url string `json:"organization_repositories_url"`
	Followers_url string `json:"followers_url"`
	Organization_url string `json:"organization_url"`
	Notifications_url string `json:"notifications_url"`
	Repository_search_url string `json:"repository_search_url"`
	Current_user_url string `json:"current_user_url"`
	User_organizations_url string `json:"user_organizations_url"`
	Code_search_url string `json:"code_search_url"`
}

// GeneratedType_Webhook_repository_advisory_published represents the GeneratedType_Webhook_repository_advisory_published schema from the OpenAPI specification
type GeneratedType_Webhook_repository_advisory_published struct {
	Repository_advisory GeneratedType_Repository_advisory `json:"repository_advisory"` // A repository security advisory.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType_Webhook_code_scanning_alert_appeared_in_branch represents the GeneratedType_Webhook_code_scanning_alert_appeared_in_branch schema from the OpenAPI specification
type GeneratedType_Webhook_code_scanning_alert_appeared_in_branch struct {
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Ref string `json:"ref"` // The Git reference of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Alert map[string]interface{} `json:"alert"` // The code scanning alert involved in the event.
	Commit_oid string `json:"commit_oid"` // The commit SHA of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Action string `json:"action"`
}

// GeneratedType_Repository_rule_file_path_restriction represents the GeneratedType_Repository_rule_file_path_restriction schema from the OpenAPI specification
type GeneratedType_Repository_rule_file_path_restriction struct {
	TypeField string `json:"type"`
	Parameters map[string]interface{} `json:"parameters,omitempty"`
}

// GeneratedType_Timeline_comment_event represents the GeneratedType_Timeline_comment_event schema from the OpenAPI specification
type GeneratedType_Timeline_comment_event struct {
	Event string `json:"event"`
	Reactions GeneratedType_Reaction_rollup `json:"reactions,omitempty"`
	Created_at string `json:"created_at"`
	Node_id string `json:"node_id"`
	Url string `json:"url"` // URL for the issue comment
	Body string `json:"body,omitempty"` // Contents of the issue comment
	Body_html string `json:"body_html,omitempty"`
	Updated_at string `json:"updated_at"`
	Issue_url string `json:"issue_url"`
	Body_text string `json:"body_text,omitempty"`
	Html_url string `json:"html_url"`
	Actor GeneratedType_Simple_user `json:"actor"` // A GitHub user.
	User GeneratedType_Simple_user `json:"user"` // A GitHub user.
	Id int `json:"id"` // Unique identifier of the issue comment
	Performed_via_github_app GeneratedType_Nullable_integration `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Author_association string `json:"author_association"` // How the author is associated with the repository.
}

// GeneratedType_Check_annotation represents the GeneratedType_Check_annotation schema from the OpenAPI specification
type GeneratedType_Check_annotation struct {
	Message string `json:"message"`
	Start_line int `json:"start_line"`
	Title string `json:"title"`
	End_line int `json:"end_line"`
	Annotation_level string `json:"annotation_level"`
	Blob_href string `json:"blob_href"`
	End_column int `json:"end_column"`
	Path string `json:"path"`
	Raw_details string `json:"raw_details"`
	Start_column int `json:"start_column"`
}

// GeneratedType_Webhook_branch_protection_configuration_disabled represents the GeneratedType_Webhook_branch_protection_configuration_disabled schema from the OpenAPI specification
type GeneratedType_Webhook_branch_protection_configuration_disabled struct {
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
}

// Commit represents the Commit schema from the OpenAPI specification
type Commit struct {
	Url string `json:"url"`
	Author interface{} `json:"author"`
	Parents []map[string]interface{} `json:"parents"`
	Stats map[string]interface{} `json:"stats,omitempty"`
	Html_url string `json:"html_url"`
	Files []GeneratedType_Diff_entry `json:"files,omitempty"`
	Node_id string `json:"node_id"`
	Sha string `json:"sha"`
	Comments_url string `json:"comments_url"`
	Commit map[string]interface{} `json:"commit"`
	Committer interface{} `json:"committer"`
}

// GeneratedType_Actions_repository_permissions represents the GeneratedType_Actions_repository_permissions schema from the OpenAPI specification
type GeneratedType_Actions_repository_permissions struct {
	Selected_actions_url string `json:"selected_actions_url,omitempty"` // The API URL to use to get or set the actions and reusable workflows that are allowed to run, when `allowed_actions` is set to `selected`.
	Allowed_actions string `json:"allowed_actions,omitempty"` // The permissions policy that controls the actions and reusable workflows that are allowed to run.
	Enabled bool `json:"enabled"` // Whether GitHub Actions is enabled on the repository.
}

// Hovercard represents the Hovercard schema from the OpenAPI specification
type Hovercard struct {
	Contexts []map[string]interface{} `json:"contexts"`
}

// Manifest represents the Manifest schema from the OpenAPI specification
type Manifest struct {
	File map[string]interface{} `json:"file,omitempty"`
	Metadata Metadata `json:"metadata,omitempty"` // User-defined metadata to store domain-specific information limited to 8 keys with scalar values.
	Name string `json:"name"` // The name of the manifest.
	Resolved map[string]interface{} `json:"resolved,omitempty"` // A collection of resolved package dependencies.
}

// GeneratedType_Gist_simple represents the GeneratedType_Gist_simple schema from the OpenAPI specification
type GeneratedType_Gist_simple struct {
	Comments_enabled bool `json:"comments_enabled,omitempty"`
	User string `json:"user,omitempty"`
	Commits_url string `json:"commits_url,omitempty"`
	Description string `json:"description,omitempty"`
	History []GeneratedType_Gist_history `json:"history,omitempty"`
	Comments_url string `json:"comments_url,omitempty"`
	Forks []map[string]interface{} `json:"forks,omitempty"`
	Truncated bool `json:"truncated,omitempty"`
	Git_push_url string `json:"git_push_url,omitempty"`
	Comments int `json:"comments,omitempty"`
	Created_at string `json:"created_at,omitempty"`
	Html_url string `json:"html_url,omitempty"`
	Url string `json:"url,omitempty"`
	Fork_of map[string]interface{} `json:"fork_of,omitempty"` // Gist
	Forks_url string `json:"forks_url,omitempty"`
	Node_id string `json:"node_id,omitempty"`
	Owner GeneratedType_Simple_user `json:"owner,omitempty"` // A GitHub user.
	Git_pull_url string `json:"git_pull_url,omitempty"`
	Public bool `json:"public,omitempty"`
	Files map[string]interface{} `json:"files,omitempty"`
	Id string `json:"id,omitempty"`
	Updated_at string `json:"updated_at,omitempty"`
}

// Webhooksteam1 represents the Webhooksteam1 schema from the OpenAPI specification
type Webhooksteam1 struct {
	Html_url string `json:"html_url,omitempty"`
	Name string `json:"name"` // Name of the team
	Parent map[string]interface{} `json:"parent,omitempty"`
	Privacy string `json:"privacy,omitempty"`
	Permission string `json:"permission,omitempty"` // Permission that the team will have for its repositories
	Repositories_url string `json:"repositories_url,omitempty"`
	Slug string `json:"slug,omitempty"`
	Deleted bool `json:"deleted,omitempty"`
	Id int `json:"id"` // Unique identifier of the team
	Members_url string `json:"members_url,omitempty"`
	Node_id string `json:"node_id,omitempty"`
	Description string `json:"description,omitempty"` // Description of the team
	Notification_setting string `json:"notification_setting,omitempty"` // Whether team members will receive notifications when their team is @mentioned
	Url string `json:"url,omitempty"` // URL for the team
}

// GeneratedType_Deployment_branch_policy_settings represents the GeneratedType_Deployment_branch_policy_settings schema from the OpenAPI specification
type GeneratedType_Deployment_branch_policy_settings struct {
	Protected_branches bool `json:"protected_branches"` // Whether only branches with branch protection rules can deploy to this environment. If `protected_branches` is `true`, `custom_branch_policies` must be `false`; if `protected_branches` is `false`, `custom_branch_policies` must be `true`.
	Custom_branch_policies bool `json:"custom_branch_policies"` // Whether only branches that match the specified name patterns can deploy to this environment. If `custom_branch_policies` is `true`, `protected_branches` must be `false`; if `custom_branch_policies` is `false`, `protected_branches` must be `true`.
}

// GeneratedType_Webhook_organization_member_added represents the GeneratedType_Webhook_organization_member_added schema from the OpenAPI specification
type GeneratedType_Webhook_organization_member_added struct {
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Membership Webhooksmembership `json:"membership"` // The membership between the user and the organization. Not present when the action is `member_invited`.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType_Webhook_repository_transferred represents the GeneratedType_Webhook_repository_transferred schema from the OpenAPI specification
type GeneratedType_Webhook_repository_transferred struct {
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType_Link_with_type represents the GeneratedType_Link_with_type schema from the OpenAPI specification
type GeneratedType_Link_with_type struct {
	Href string `json:"href"`
	TypeField string `json:"type"`
}

// GeneratedType_Code_of_conduct represents the GeneratedType_Code_of_conduct schema from the OpenAPI specification
type GeneratedType_Code_of_conduct struct {
	Name string `json:"name"`
	Url string `json:"url"`
	Body string `json:"body,omitempty"`
	Html_url string `json:"html_url"`
	Key string `json:"key"`
}

// GeneratedType_Codespaces_permissions_check_for_devcontainer represents the GeneratedType_Codespaces_permissions_check_for_devcontainer schema from the OpenAPI specification
type GeneratedType_Codespaces_permissions_check_for_devcontainer struct {
	Accepted bool `json:"accepted"` // Whether the user has accepted the permissions defined by the devcontainer config
}

// GeneratedType_Team_organization represents the GeneratedType_Team_organization schema from the OpenAPI specification
type GeneratedType_Team_organization struct {
	Hooks_url string `json:"hooks_url"`
	Disk_usage int `json:"disk_usage,omitempty"`
	Location string `json:"location,omitempty"`
	Public_repos int `json:"public_repos"`
	Members_allowed_repository_creation_type string `json:"members_allowed_repository_creation_type,omitempty"`
	Updated_at string `json:"updated_at"`
	Members_can_create_pages bool `json:"members_can_create_pages,omitempty"`
	Default_repository_permission string `json:"default_repository_permission,omitempty"`
	Followers int `json:"followers"`
	Name string `json:"name,omitempty"`
	Collaborators int `json:"collaborators,omitempty"`
	Members_can_create_private_repositories bool `json:"members_can_create_private_repositories,omitempty"`
	Avatar_url string `json:"avatar_url"`
	Issues_url string `json:"issues_url"`
	Members_can_create_public_pages bool `json:"members_can_create_public_pages,omitempty"`
	Public_gists int `json:"public_gists"`
	Owned_private_repos int `json:"owned_private_repos,omitempty"`
	Following int `json:"following"`
	Billing_email string `json:"billing_email,omitempty"`
	Twitter_username string `json:"twitter_username,omitempty"`
	TypeField string `json:"type"`
	Events_url string `json:"events_url"`
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"`
	Members_can_create_repositories bool `json:"members_can_create_repositories,omitempty"`
	Repos_url string `json:"repos_url"`
	Has_repository_projects bool `json:"has_repository_projects"`
	Url string `json:"url"`
	Node_id string `json:"node_id"`
	Members_can_create_private_pages bool `json:"members_can_create_private_pages,omitempty"`
	Plan map[string]interface{} `json:"plan,omitempty"`
	Archived_at string `json:"archived_at"`
	Members_can_create_public_repositories bool `json:"members_can_create_public_repositories,omitempty"`
	Email string `json:"email,omitempty"`
	Two_factor_requirement_enabled bool `json:"two_factor_requirement_enabled,omitempty"`
	Is_verified bool `json:"is_verified,omitempty"`
	Private_gists int `json:"private_gists,omitempty"`
	Has_organization_projects bool `json:"has_organization_projects"`
	Html_url string `json:"html_url"`
	Blog string `json:"blog,omitempty"`
	Public_members_url string `json:"public_members_url"`
	Company string `json:"company,omitempty"`
	Members_can_fork_private_repositories bool `json:"members_can_fork_private_repositories,omitempty"`
	Total_private_repos int `json:"total_private_repos,omitempty"`
	Id int `json:"id"`
	Members_url string `json:"members_url"`
	Created_at string `json:"created_at"`
	Description string `json:"description"`
	Login string `json:"login"`
	Members_can_create_internal_repositories bool `json:"members_can_create_internal_repositories,omitempty"`
}

// GeneratedType_Code_security_configuration_for_repository represents the GeneratedType_Code_security_configuration_for_repository schema from the OpenAPI specification
type GeneratedType_Code_security_configuration_for_repository struct {
	Status string `json:"status,omitempty"` // The attachment status of the code security configuration on the repository.
	Configuration GeneratedType_Code_security_configuration `json:"configuration,omitempty"` // A code security configuration
}

// GeneratedType_Webhook_project_card_created represents the GeneratedType_Webhook_project_card_created schema from the OpenAPI specification
type GeneratedType_Webhook_project_card_created struct {
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Project_card Webhooksprojectcard `json:"project_card"`
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType_Webhook_status represents the GeneratedType_Webhook_status schema from the OpenAPI specification
type GeneratedType_Webhook_status struct {
	Branches []map[string]interface{} `json:"branches"` // An array of branch objects containing the status' SHA. Each branch contains the given SHA, but the SHA may or may not be the head of the branch. The array includes a maximum of 10 branches.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Name string `json:"name"`
	Commit map[string]interface{} `json:"commit"`
	Description string `json:"description"` // The optional human-readable description added to the status.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Context string `json:"context"`
	Avatar_url string `json:"avatar_url,omitempty"`
	Id int `json:"id"` // The unique identifier of the status.
	Target_url string `json:"target_url"` // The optional link added to the status.
	Updated_at string `json:"updated_at"`
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Created_at string `json:"created_at"`
	Sha string `json:"sha"` // The Commit SHA.
	State string `json:"state"` // The new state. Can be `pending`, `success`, `failure`, or `error`.
}

// GeneratedType_Timeline_issue_events represents the GeneratedType_Timeline_issue_events schema from the OpenAPI specification
type GeneratedType_Timeline_issue_events struct {
}

// Webhookschanges represents the Webhookschanges schema from the OpenAPI specification
type Webhookschanges struct {
	Body map[string]interface{} `json:"body,omitempty"`
}

// GeneratedType_Runner_label represents the GeneratedType_Runner_label schema from the OpenAPI specification
type GeneratedType_Runner_label struct {
	Name string `json:"name"` // Name of the label.
	TypeField string `json:"type,omitempty"` // The type of label. Read-only labels are applied automatically when the runner is configured.
	Id int `json:"id,omitempty"` // Unique identifier of the label.
}

// GeneratedType_Repository_advisory_create represents the GeneratedType_Repository_advisory_create schema from the OpenAPI specification
type GeneratedType_Repository_advisory_create struct {
	Cwe_ids []string `json:"cwe_ids,omitempty"` // A list of Common Weakness Enumeration (CWE) IDs.
	Description string `json:"description"` // A detailed description of what the advisory impacts.
	Cve_id string `json:"cve_id,omitempty"` // The Common Vulnerabilities and Exposures (CVE) ID.
	Start_private_fork bool `json:"start_private_fork,omitempty"` // Whether to create a temporary private fork of the repository to collaborate on a fix.
	Vulnerabilities []map[string]interface{} `json:"vulnerabilities"` // A product affected by the vulnerability detailed in a repository security advisory.
	Cvss_vector_string string `json:"cvss_vector_string,omitempty"` // The CVSS vector that calculates the severity of the advisory. You must choose between setting this field or `severity`.
	Severity string `json:"severity,omitempty"` // The severity of the advisory. You must choose between setting this field or `cvss_vector_string`.
	Summary string `json:"summary"` // A short summary of the advisory.
	Credits []map[string]interface{} `json:"credits,omitempty"` // A list of users receiving credit for their participation in the security advisory.
}

// GeneratedType_Webhook_discussion_locked represents the GeneratedType_Webhook_discussion_locked schema from the OpenAPI specification
type GeneratedType_Webhook_discussion_locked struct {
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType_Webhook_push represents the GeneratedType_Webhook_push schema from the OpenAPI specification
type GeneratedType_Webhook_push struct {
	After string `json:"after"` // The SHA of the most recent commit on `ref` after the push.
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Ref string `json:"ref"` // The full git ref that was pushed. Example: `refs/heads/main` or `refs/tags/v3.14.1`.
	Pusher map[string]interface{} `json:"pusher"` // Metaproperties for Git author/committer information.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Compare string `json:"compare"` // URL that shows the changes in this `ref` update, from the `before` commit to the `after` commit. For a newly created `ref` that is directly based on the default branch, this is the comparison between the head of the default branch and the `after` commit. Otherwise, this shows all commits until the `after` commit.
	Head_commit map[string]interface{} `json:"head_commit"`
	Repository map[string]interface{} `json:"repository"` // A git repository
	Forced bool `json:"forced"` // Whether this push was a force push of the `ref`.
	Before string `json:"before"` // The SHA of the most recent commit on `ref` before the push.
	Base_ref string `json:"base_ref"`
	Commits []map[string]interface{} `json:"commits"` // An array of commit objects describing the pushed commits. (Pushed commits are all commits that are included in the `compare` between the `before` commit and the `after` commit.) The array includes a maximum of 2048 commits. If necessary, you can use the [Commits API](https://docs.github.com/rest/commits) to fetch additional commits.
	Deleted bool `json:"deleted"` // Whether this push deleted the `ref`.
	Created bool `json:"created"` // Whether this push created the `ref`.
}

// Webhooksprojectchanges represents the Webhooksprojectchanges schema from the OpenAPI specification
type Webhooksprojectchanges struct {
	Archived_at map[string]interface{} `json:"archived_at,omitempty"`
}

// GeneratedType_Webhook_repository_privatized represents the GeneratedType_Webhook_repository_privatized schema from the OpenAPI specification
type GeneratedType_Webhook_repository_privatized struct {
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType_Webhook_meta_deleted represents the GeneratedType_Webhook_meta_deleted schema from the OpenAPI specification
type GeneratedType_Webhook_meta_deleted struct {
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Nullable_repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Hook map[string]interface{} `json:"hook"` // The deleted webhook. This will contain different keys based on the type of webhook it is: repository, organization, business, app, or GitHub Marketplace.
	Hook_id int `json:"hook_id"` // The id of the modified webhook.
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType_Codespace_machine represents the GeneratedType_Codespace_machine schema from the OpenAPI specification
type GeneratedType_Codespace_machine struct {
	Operating_system string `json:"operating_system"` // The operating system of the machine.
	Prebuild_availability string `json:"prebuild_availability"` // Whether a prebuild is currently available when creating a codespace for this machine and repository. If a branch was not specified as a ref, the default branch will be assumed. Value will be "null" if prebuilds are not supported or prebuild availability could not be determined. Value will be "none" if no prebuild is available. Latest values "ready" and "in_progress" indicate the prebuild availability status.
	Storage_in_bytes int `json:"storage_in_bytes"` // How much storage is available to the codespace.
	Cpus int `json:"cpus"` // How many cores are available to the codespace.
	Display_name string `json:"display_name"` // The display name of the machine includes cores, memory, and storage.
	Memory_in_bytes int `json:"memory_in_bytes"` // How much memory is available to the codespace.
	Name string `json:"name"` // The name of the machine.
}

// GeneratedType_Webhook_milestone_deleted represents the GeneratedType_Webhook_milestone_deleted schema from the OpenAPI specification
type GeneratedType_Webhook_milestone_deleted struct {
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Milestone Webhooksmilestone `json:"milestone"` // A collection of related issues and pull requests.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType_Copilot_dotcom_chat represents the GeneratedType_Copilot_dotcom_chat schema from the OpenAPI specification
type GeneratedType_Copilot_dotcom_chat struct {
	Models []map[string]interface{} `json:"models,omitempty"` // List of model metrics for a custom models and the default model.
	Total_engaged_users int `json:"total_engaged_users,omitempty"` // Total number of users who prompted Copilot Chat on github.com at least once.
}

// GeneratedType_Milestoned_issue_event represents the GeneratedType_Milestoned_issue_event schema from the OpenAPI specification
type GeneratedType_Milestoned_issue_event struct {
	Commit_url string `json:"commit_url"`
	Event string `json:"event"`
	Milestone map[string]interface{} `json:"milestone"`
	Commit_id string `json:"commit_id"`
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	Created_at string `json:"created_at"`
	Performed_via_github_app GeneratedType_Nullable_integration `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Url string `json:"url"`
	Actor GeneratedType_Simple_user `json:"actor"` // A GitHub user.
}

// Traffic represents the Traffic schema from the OpenAPI specification
type Traffic struct {
	Uniques int `json:"uniques"`
	Count int `json:"count"`
	Timestamp string `json:"timestamp"`
}

// GeneratedType_Webhook_security_advisory_published represents the GeneratedType_Webhook_security_advisory_published schema from the OpenAPI specification
type GeneratedType_Webhook_security_advisory_published struct {
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Security_advisory Webhookssecurityadvisory `json:"security_advisory"` // The details of the security advisory, including summary, description, and severity.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType_Repository_rule_code_scanning represents the GeneratedType_Repository_rule_code_scanning schema from the OpenAPI specification
type GeneratedType_Repository_rule_code_scanning struct {
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	TypeField string `json:"type"`
}

// Dependency represents the Dependency schema from the OpenAPI specification
type Dependency struct {
	Package_url string `json:"package_url,omitempty"` // Package-url (PURL) of dependency. See https://github.com/package-url/purl-spec for more details.
	Relationship string `json:"relationship,omitempty"` // A notation of whether a dependency is requested directly by this manifest or is a dependency of another dependency.
	Scope string `json:"scope,omitempty"` // A notation of whether the dependency is required for the primary build artifact (runtime) or is only used for development. Future versions of this specification may allow for more granular scopes.
	Dependencies []string `json:"dependencies,omitempty"` // Array of package-url (PURLs) of direct child dependencies.
	Metadata Metadata `json:"metadata,omitempty"` // User-defined metadata to store domain-specific information limited to 8 keys with scalar values.
}

// GeneratedType_Webhook_discussion_deleted represents the GeneratedType_Webhook_discussion_deleted schema from the OpenAPI specification
type GeneratedType_Webhook_discussion_deleted struct {
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType_Webhook_discussion_unanswered represents the GeneratedType_Webhook_discussion_unanswered schema from the OpenAPI specification
type GeneratedType_Webhook_discussion_unanswered struct {
	Old_answer Webhooksanswer `json:"old_answer"`
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
}

// GeneratedType_Ruleset_version_with_state represents the GeneratedType_Ruleset_version_with_state schema from the OpenAPI specification
type GeneratedType_Ruleset_version_with_state struct {
	Updated_at string `json:"updated_at"`
	Version_id int `json:"version_id"` // The ID of the previous version of the ruleset
	Actor map[string]interface{} `json:"actor"` // The actor who updated the ruleset
	State map[string]interface{} `json:"state"` // The state of the ruleset version
}

// GeneratedType_Webhook_workflow_run_completed represents the GeneratedType_Webhook_workflow_run_completed schema from the OpenAPI specification
type GeneratedType_Webhook_workflow_run_completed struct {
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Workflow Webhooksworkflow `json:"workflow"`
	Workflow_run map[string]interface{} `json:"workflow_run"`
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType_Repository_rule_params_status_check_configuration represents the GeneratedType_Repository_rule_params_status_check_configuration schema from the OpenAPI specification
type GeneratedType_Repository_rule_params_status_check_configuration struct {
	Context string `json:"context"` // The status check context name that must be present on the commit.
	Integration_id int `json:"integration_id,omitempty"` // The optional integration ID that this status check must originate from.
}

// GeneratedType_Webhook_repository_renamed represents the GeneratedType_Webhook_repository_renamed schema from the OpenAPI specification
type GeneratedType_Webhook_repository_renamed struct {
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
}

// GeneratedType_Workflow_usage represents the GeneratedType_Workflow_usage schema from the OpenAPI specification
type GeneratedType_Workflow_usage struct {
	Billable map[string]interface{} `json:"billable"`
}

// GeneratedType_Webhook_pull_request_converted_to_draft represents the GeneratedType_Webhook_pull_request_converted_to_draft schema from the OpenAPI specification
type GeneratedType_Webhook_pull_request_converted_to_draft struct {
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request GeneratedType_Pull_request_webhook `json:"pull_request"`
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// Email represents the Email schema from the OpenAPI specification
type Email struct {
	Email string `json:"email"`
	Primary bool `json:"primary"`
	Verified bool `json:"verified"`
	Visibility string `json:"visibility"`
}

// GeneratedType_Webhook_label_deleted represents the GeneratedType_Webhook_label_deleted schema from the OpenAPI specification
type GeneratedType_Webhook_label_deleted struct {
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Label Webhookslabel `json:"label"`
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType_Webhook_member_removed represents the GeneratedType_Webhook_member_removed schema from the OpenAPI specification
type GeneratedType_Webhook_member_removed struct {
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Member Webhooksuser `json:"member"`
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType_Workflow_run_usage represents the GeneratedType_Workflow_run_usage schema from the OpenAPI specification
type GeneratedType_Workflow_run_usage struct {
	Run_duration_ms int `json:"run_duration_ms,omitempty"`
	Billable map[string]interface{} `json:"billable"`
}

// GeneratedType_Webhook_github_app_authorization_revoked represents the GeneratedType_Webhook_github_app_authorization_revoked schema from the OpenAPI specification
type GeneratedType_Webhook_github_app_authorization_revoked struct {
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType_Pages_health_check represents the GeneratedType_Pages_health_check schema from the OpenAPI specification
type GeneratedType_Pages_health_check struct {
	Alt_domain map[string]interface{} `json:"alt_domain,omitempty"`
	Domain map[string]interface{} `json:"domain,omitempty"`
}

// GeneratedType_Webhook_team_removed_from_repository represents the GeneratedType_Webhook_team_removed_from_repository schema from the OpenAPI specification
type GeneratedType_Webhook_team_removed_from_repository struct {
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Team Webhooksteam1 `json:"team"` // Groups of organization members that gives permissions on specified repositories.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository map[string]interface{} `json:"repository,omitempty"` // A git repository
}

// GeneratedType_Hook_delivery_item represents the GeneratedType_Hook_delivery_item schema from the OpenAPI specification
type GeneratedType_Hook_delivery_item struct {
	Delivered_at string `json:"delivered_at"` // Time when the webhook delivery occurred.
	Duration float64 `json:"duration"` // Time spent delivering.
	Guid string `json:"guid"` // Unique identifier for the event (shared with all deliveries for all webhooks that subscribe to this event).
	Id int64 `json:"id"` // Unique identifier of the webhook delivery.
	Event string `json:"event"` // The event that triggered the delivery.
	Redelivery bool `json:"redelivery"` // Whether the webhook delivery is a redelivery.
	Repository_id int64 `json:"repository_id"` // The id of the repository associated with this event.
	Action string `json:"action"` // The type of activity for the event that triggered the delivery.
	Status_code int `json:"status_code"` // Status code received when delivery was made.
	Throttled_at string `json:"throttled_at,omitempty"` // Time when the webhook delivery was throttled.
	Installation_id int64 `json:"installation_id"` // The id of the GitHub App installation associated with this event.
	Status string `json:"status"` // Describes the response returned after attempting the delivery.
}

// GeneratedType_Webhook_registry_package_updated represents the GeneratedType_Webhook_registry_package_updated schema from the OpenAPI specification
type GeneratedType_Webhook_registry_package_updated struct {
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Registry_package map[string]interface{} `json:"registry_package"`
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// Discussion represents the Discussion schema from the OpenAPI specification
type Discussion struct {
	Body string `json:"body"`
	Labels []Label `json:"labels,omitempty"`
	Locked bool `json:"locked"`
	Updated_at string `json:"updated_at"`
	Created_at string `json:"created_at"`
	Answer_html_url string `json:"answer_html_url"`
	Node_id string `json:"node_id"`
	Comments int `json:"comments"`
	User map[string]interface{} `json:"user"`
	Category map[string]interface{} `json:"category"`
	Id int `json:"id"`
	Title string `json:"title"`
	Active_lock_reason string `json:"active_lock_reason"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Html_url string `json:"html_url"`
	State_reason string `json:"state_reason"` // The reason for the current state
	Answer_chosen_by map[string]interface{} `json:"answer_chosen_by"`
	Number int `json:"number"`
	State string `json:"state"` // The current state of the discussion. `converting` means that the discussion is being converted from an issue. `transferring` means that the discussion is being transferred from another repository.
	Timeline_url string `json:"timeline_url,omitempty"`
	Reactions map[string]interface{} `json:"reactions,omitempty"`
	Repository_url string `json:"repository_url"`
	Answer_chosen_at string `json:"answer_chosen_at"`
}

// GeneratedType_Webhook_issues_opened represents the GeneratedType_Webhook_issues_opened schema from the OpenAPI specification
type GeneratedType_Webhook_issues_opened struct {
	Changes map[string]interface{} `json:"changes,omitempty"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType_Classroom_accepted_assignment represents the GeneratedType_Classroom_accepted_assignment schema from the OpenAPI specification
type GeneratedType_Classroom_accepted_assignment struct {
	Commit_count int `json:"commit_count"` // Count of student commits.
	Grade string `json:"grade"` // Most recent grade.
	Id int `json:"id"` // Unique identifier of the repository.
	Passing bool `json:"passing"` // Whether a submission passed.
	Repository GeneratedType_Simple_classroom_repository `json:"repository"` // A GitHub repository view for Classroom
	Students []GeneratedType_Simple_classroom_user `json:"students"`
	Submitted bool `json:"submitted"` // Whether an accepted assignment has been submitted.
	Assignment GeneratedType_Simple_classroom_assignment `json:"assignment"` // A GitHub Classroom assignment
}

// GeneratedType_Webhook_security_and_analysis represents the GeneratedType_Webhook_security_and_analysis schema from the OpenAPI specification
type GeneratedType_Webhook_security_and_analysis struct {
	Repository GeneratedType_Full_repository `json:"repository"` // Full Repository
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Changes map[string]interface{} `json:"changes"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType_Porter_author represents the GeneratedType_Porter_author schema from the OpenAPI specification
type GeneratedType_Porter_author struct {
	Id int `json:"id"`
	Import_url string `json:"import_url"`
	Name string `json:"name"`
	Remote_id string `json:"remote_id"`
	Remote_name string `json:"remote_name"`
	Url string `json:"url"`
	Email string `json:"email"`
}

// GeneratedType_Actions_workflow_access_to_repository represents the GeneratedType_Actions_workflow_access_to_repository schema from the OpenAPI specification
type GeneratedType_Actions_workflow_access_to_repository struct {
	Access_level string `json:"access_level"` // Defines the level of access that workflows outside of the repository have to actions and reusable workflows within the repository. `none` means the access is only possible from workflows in this repository. `user` level access allows sharing across user owned private repositories only. `organization` level access allows sharing across the organization.
}

// GeneratedType_Secret_scanning_location_discussion_comment represents the GeneratedType_Secret_scanning_location_discussion_comment schema from the OpenAPI specification
type GeneratedType_Secret_scanning_location_discussion_comment struct {
	Discussion_comment_url string `json:"discussion_comment_url"` // The API URL to get the discussion comment where the secret was detected.
}

// GeneratedType_Validation_error_simple represents the GeneratedType_Validation_error_simple schema from the OpenAPI specification
type GeneratedType_Validation_error_simple struct {
	Documentation_url string `json:"documentation_url"`
	Errors []string `json:"errors,omitempty"`
	Message string `json:"message"`
}

// GeneratedType_Webhook_deployment_review_requested represents the GeneratedType_Webhook_deployment_review_requested schema from the OpenAPI specification
type GeneratedType_Webhook_deployment_review_requested struct {
	Workflow_run map[string]interface{} `json:"workflow_run"`
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Reviewers []map[string]interface{} `json:"reviewers"`
	Since string `json:"since"`
	Action string `json:"action"`
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Environment string `json:"environment"`
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Requestor Webhooksuser `json:"requestor"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Workflow_job_run map[string]interface{} `json:"workflow_job_run"`
}

// GeneratedType_Issue_event_project_card represents the GeneratedType_Issue_event_project_card schema from the OpenAPI specification
type GeneratedType_Issue_event_project_card struct {
	Url string `json:"url"`
	Column_name string `json:"column_name"`
	Id int `json:"id"`
	Previous_column_name string `json:"previous_column_name,omitempty"`
	Project_id int `json:"project_id"`
	Project_url string `json:"project_url"`
}

// Webhooksdeploykey represents the Webhooksdeploykey schema from the OpenAPI specification
type Webhooksdeploykey struct {
	Last_used string `json:"last_used,omitempty"`
	Url string `json:"url"`
	Read_only bool `json:"read_only"`
	Title string `json:"title"`
	Id int `json:"id"`
	Key string `json:"key"`
	Verified bool `json:"verified"`
	Added_by string `json:"added_by,omitempty"`
	Created_at string `json:"created_at"`
	Enabled bool `json:"enabled,omitempty"`
}

// Hook represents the Hook schema from the OpenAPI specification
type Hook struct {
	Last_response GeneratedType_Hook_response `json:"last_response"`
	Ping_url string `json:"ping_url"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	Active bool `json:"active"` // Determines whether the hook is actually triggered on pushes.
	Config GeneratedType_Webhook_config `json:"config"` // Configuration object of the webhook
	Id int `json:"id"` // Unique identifier of the webhook.
	Name string `json:"name"` // The name of a valid service, use 'web' for a webhook.
	TypeField string `json:"type"`
	Created_at string `json:"created_at"`
	Deliveries_url string `json:"deliveries_url,omitempty"`
	Events []string `json:"events"` // Determines what events the hook is triggered for. Default: ['push'].
	Test_url string `json:"test_url"`
}

// Webhooksteam represents the Webhooksteam schema from the OpenAPI specification
type Webhooksteam struct {
	Notification_setting string `json:"notification_setting,omitempty"`
	Url string `json:"url,omitempty"` // URL for the team
	Deleted bool `json:"deleted,omitempty"`
	Html_url string `json:"html_url,omitempty"`
	Node_id string `json:"node_id,omitempty"`
	Privacy string `json:"privacy,omitempty"`
	Repositories_url string `json:"repositories_url,omitempty"`
	Members_url string `json:"members_url,omitempty"`
	Parent map[string]interface{} `json:"parent,omitempty"`
	Slug string `json:"slug,omitempty"`
	Id int `json:"id"` // Unique identifier of the team
	Name string `json:"name"` // Name of the team
	Permission string `json:"permission,omitempty"` // Permission that the team will have for its repositories
	Description string `json:"description,omitempty"` // Description of the team
}

// Contributor represents the Contributor schema from the OpenAPI specification
type Contributor struct {
	Node_id string `json:"node_id,omitempty"`
	Gists_url string `json:"gists_url,omitempty"`
	Subscriptions_url string `json:"subscriptions_url,omitempty"`
	Following_url string `json:"following_url,omitempty"`
	Organizations_url string `json:"organizations_url,omitempty"`
	Url string `json:"url,omitempty"`
	Events_url string `json:"events_url,omitempty"`
	Site_admin bool `json:"site_admin,omitempty"`
	User_view_type string `json:"user_view_type,omitempty"`
	Gravatar_id string `json:"gravatar_id,omitempty"`
	Login string `json:"login,omitempty"`
	Received_events_url string `json:"received_events_url,omitempty"`
	Html_url string `json:"html_url,omitempty"`
	Name string `json:"name,omitempty"`
	Avatar_url string `json:"avatar_url,omitempty"`
	Followers_url string `json:"followers_url,omitempty"`
	TypeField string `json:"type"`
	Id int `json:"id,omitempty"`
	Starred_url string `json:"starred_url,omitempty"`
	Repos_url string `json:"repos_url,omitempty"`
	Contributions int `json:"contributions"`
	Email string `json:"email,omitempty"`
}

// GeneratedType_Private_vulnerability_report_create represents the GeneratedType_Private_vulnerability_report_create schema from the OpenAPI specification
type GeneratedType_Private_vulnerability_report_create struct {
	Summary string `json:"summary"` // A short summary of the advisory.
	Vulnerabilities []map[string]interface{} `json:"vulnerabilities,omitempty"` // An array of products affected by the vulnerability detailed in a repository security advisory.
	Cvss_vector_string string `json:"cvss_vector_string,omitempty"` // The CVSS vector that calculates the severity of the advisory. You must choose between setting this field or `severity`.
	Cwe_ids []string `json:"cwe_ids,omitempty"` // A list of Common Weakness Enumeration (CWE) IDs.
	Description string `json:"description"` // A detailed description of what the advisory impacts.
	Severity string `json:"severity,omitempty"` // The severity of the advisory. You must choose between setting this field or `cvss_vector_string`.
	Start_private_fork bool `json:"start_private_fork,omitempty"` // Whether to create a temporary private fork of the repository to collaborate on a fix.
}

// Webhooksrelease1 represents the Webhooksrelease1 schema from the OpenAPI specification
type Webhooksrelease1 struct {
	Discussion_url string `json:"discussion_url,omitempty"`
	Node_id string `json:"node_id"`
	Created_at string `json:"created_at"`
	Prerelease bool `json:"prerelease"` // Whether the release is identified as a prerelease or a full release.
	Body string `json:"body"`
	Name string `json:"name"`
	Tarball_url string `json:"tarball_url"`
	Target_commitish string `json:"target_commitish"` // Specifies the commitish value that determines where the Git tag is created from.
	Author map[string]interface{} `json:"author"`
	Zipball_url string `json:"zipball_url"`
	Reactions map[string]interface{} `json:"reactions,omitempty"`
	Upload_url string `json:"upload_url"`
	Draft bool `json:"draft"` // Whether the release is a draft or published
	Html_url string `json:"html_url"`
	Tag_name string `json:"tag_name"` // The name of the tag.
	Assets []map[string]interface{} `json:"assets"`
	Assets_url string `json:"assets_url"`
	Url string `json:"url"`
	Id int `json:"id"`
	Published_at string `json:"published_at"`
}

// GeneratedType_Issue_event_milestone represents the GeneratedType_Issue_event_milestone schema from the OpenAPI specification
type GeneratedType_Issue_event_milestone struct {
	Title string `json:"title"`
}

// Webhooksprojectcolumn represents the Webhooksprojectcolumn schema from the OpenAPI specification
type Webhooksprojectcolumn struct {
	After_id int `json:"after_id,omitempty"`
	Created_at string `json:"created_at"`
	Id int `json:"id"` // The unique identifier of the project column
	Url string `json:"url"`
	Name string `json:"name"` // Name of the project column
	Node_id string `json:"node_id"`
	Project_url string `json:"project_url"`
	Cards_url string `json:"cards_url"`
	Updated_at string `json:"updated_at"`
}

// GeneratedType_Webhook_delete represents the GeneratedType_Webhook_delete schema from the OpenAPI specification
type GeneratedType_Webhook_delete struct {
	Ref_type string `json:"ref_type"` // The type of Git ref object deleted in the repository.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pusher_type string `json:"pusher_type"` // The pusher type for the event. Can be either `user` or a deploy key.
	Ref string `json:"ref"` // The [`git ref`](https://docs.github.com/rest/git/refs#get-a-reference) resource.
}

// GeneratedType_Webhook_sponsorship_created represents the GeneratedType_Webhook_sponsorship_created schema from the OpenAPI specification
type GeneratedType_Webhook_sponsorship_created struct {
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Sponsorship Webhookssponsorship `json:"sponsorship"`
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// Package represents the Package schema from the OpenAPI specification
type Package struct {
	Version_count int `json:"version_count"` // The number of versions of the package.
	Name string `json:"name"` // The name of the package.
	Created_at string `json:"created_at"`
	Package_type string `json:"package_type"`
	Repository GeneratedType_Nullable_minimal_repository `json:"repository,omitempty"` // Minimal Repository
	Updated_at string `json:"updated_at"`
	Html_url string `json:"html_url"`
	Id int `json:"id"` // Unique identifier of the package.
	Visibility string `json:"visibility"`
	Owner GeneratedType_Nullable_simple_user `json:"owner,omitempty"` // A GitHub user.
	Url string `json:"url"`
}

// GeneratedType_Simple_installation represents the GeneratedType_Simple_installation schema from the OpenAPI specification
type GeneratedType_Simple_installation struct {
	Id int `json:"id"` // The ID of the installation.
	Node_id string `json:"node_id"` // The global node ID of the installation.
}

// GeneratedType_Repository_collaborator_permission represents the GeneratedType_Repository_collaborator_permission schema from the OpenAPI specification
type GeneratedType_Repository_collaborator_permission struct {
	User GeneratedType_Nullable_collaborator `json:"user"` // Collaborator
	Permission string `json:"permission"`
	Role_name string `json:"role_name"`
}

// GeneratedType_Content_traffic represents the GeneratedType_Content_traffic schema from the OpenAPI specification
type GeneratedType_Content_traffic struct {
	Uniques int `json:"uniques"`
	Count int `json:"count"`
	Path string `json:"path"`
	Title string `json:"title"`
}

// GeneratedType_Full_repository represents the GeneratedType_Full_repository schema from the OpenAPI specification
type GeneratedType_Full_repository struct {
	Size int `json:"size"` // The size of the repository, in kilobytes. Size is calculated hourly. When a repository is initially created, the size is 0.
	Open_issues int `json:"open_issues"`
	Name string `json:"name"`
	Tags_url string `json:"tags_url"`
	Temp_clone_token string `json:"temp_clone_token,omitempty"`
	Deployments_url string `json:"deployments_url"`
	Organization GeneratedType_Nullable_simple_user `json:"organization,omitempty"` // A GitHub user.
	Forks int `json:"forks"`
	Subscribers_count int `json:"subscribers_count"`
	Events_url string `json:"events_url"`
	Has_wiki bool `json:"has_wiki"`
	Disabled bool `json:"disabled"` // Returns whether or not this repository disabled.
	Has_pages bool `json:"has_pages"`
	Mirror_url string `json:"mirror_url"`
	Has_downloads bool `json:"has_downloads,omitempty"`
	Clone_url string `json:"clone_url"`
	Archived bool `json:"archived"`
	Issues_url string `json:"issues_url"`
	Source Repository `json:"source,omitempty"` // A repository on GitHub.
	Archive_url string `json:"archive_url"`
	Description string `json:"description"`
	Keys_url string `json:"keys_url"`
	Homepage string `json:"homepage"`
	Contents_url string `json:"contents_url"`
	Ssh_url string `json:"ssh_url"`
	Issue_events_url string `json:"issue_events_url"`
	Notifications_url string `json:"notifications_url"`
	Has_discussions bool `json:"has_discussions"`
	Parent Repository `json:"parent,omitempty"` // A repository on GitHub.
	Owner GeneratedType_Simple_user `json:"owner"` // A GitHub user.
	Forks_count int `json:"forks_count"`
	Labels_url string `json:"labels_url"`
	Is_template bool `json:"is_template,omitempty"`
	Allow_merge_commit bool `json:"allow_merge_commit,omitempty"`
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"`
	Language string `json:"language"`
	Open_issues_count int `json:"open_issues_count"`
	Blobs_url string `json:"blobs_url"`
	Code_of_conduct GeneratedType_Code_of_conduct_simple `json:"code_of_conduct,omitempty"` // Code of Conduct Simple
	Comments_url string `json:"comments_url"`
	Branches_url string `json:"branches_url"`
	Template_repository GeneratedType_Nullable_repository `json:"template_repository,omitempty"` // A repository on GitHub.
	Allow_squash_merge bool `json:"allow_squash_merge,omitempty"`
	Issue_comment_url string `json:"issue_comment_url"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Html_url string `json:"html_url"`
	Allow_forking bool `json:"allow_forking,omitempty"`
	Trees_url string `json:"trees_url"`
	Subscription_url string `json:"subscription_url"`
	Merge_commit_title string `json:"merge_commit_title,omitempty"` // The default value for a merge commit title. - `PR_TITLE` - default to the pull request's title. - `MERGE_MESSAGE` - default to the classic title for a merge message (e.g., Merge pull request #123 from branch-name).
	License GeneratedType_Nullable_license_simple `json:"license"` // License Simple
	Teams_url string `json:"teams_url"`
	Commits_url string `json:"commits_url"`
	Security_and_analysis GeneratedType_Security_and_analysis `json:"security_and_analysis,omitempty"`
	Stargazers_url string `json:"stargazers_url"`
	Assignees_url string `json:"assignees_url"`
	Has_issues bool `json:"has_issues"`
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"`
	Forks_url string `json:"forks_url"`
	Watchers_count int `json:"watchers_count"`
	Languages_url string `json:"languages_url"`
	Network_count int `json:"network_count"`
	Id int64 `json:"id"`
	Anonymous_access_enabled bool `json:"anonymous_access_enabled,omitempty"` // Whether anonymous git access is allowed.
	Contributors_url string `json:"contributors_url"`
	Full_name string `json:"full_name"`
	Default_branch string `json:"default_branch"`
	Use_squash_pr_title_as_default bool `json:"use_squash_pr_title_as_default,omitempty"`
	Pulls_url string `json:"pulls_url"`
	Fork bool `json:"fork"`
	Subscribers_url string `json:"subscribers_url"`
	Allow_update_branch bool `json:"allow_update_branch,omitempty"`
	Git_commits_url string `json:"git_commits_url"`
	Git_refs_url string `json:"git_refs_url"`
	Merges_url string `json:"merges_url"`
	Url string `json:"url"`
	Merge_commit_message string `json:"merge_commit_message,omitempty"` // The default value for a merge commit message. - `PR_TITLE` - default to the pull request's title. - `PR_BODY` - default to the pull request's body. - `BLANK` - default to a blank commit message.
	Squash_merge_commit_message string `json:"squash_merge_commit_message,omitempty"` // The default value for a squash merge commit message: - `PR_BODY` - default to the pull request's body. - `COMMIT_MESSAGES` - default to the branch's commit messages. - `BLANK` - default to a blank commit message.
	Downloads_url string `json:"downloads_url"`
	Master_branch string `json:"master_branch,omitempty"`
	Allow_auto_merge bool `json:"allow_auto_merge,omitempty"`
	Has_projects bool `json:"has_projects"`
	Releases_url string `json:"releases_url"`
	Private bool `json:"private"`
	Custom_properties map[string]interface{} `json:"custom_properties,omitempty"` // The custom properties that were defined for the repository. The keys are the custom property names, and the values are the corresponding custom property values.
	Stargazers_count int `json:"stargazers_count"`
	Visibility string `json:"visibility,omitempty"` // The repository visibility: public, private, or internal.
	Pushed_at string `json:"pushed_at"`
	Svn_url string `json:"svn_url"`
	Node_id string `json:"node_id"`
	Compare_url string `json:"compare_url"`
	Topics []string `json:"topics,omitempty"`
	Git_url string `json:"git_url"`
	Allow_rebase_merge bool `json:"allow_rebase_merge,omitempty"`
	Collaborators_url string `json:"collaborators_url"`
	Git_tags_url string `json:"git_tags_url"`
	Updated_at string `json:"updated_at"`
	Hooks_url string `json:"hooks_url"`
	Watchers int `json:"watchers"`
	Squash_merge_commit_title string `json:"squash_merge_commit_title,omitempty"` // The default value for a squash merge commit title: - `PR_TITLE` - default to the pull request's title. - `COMMIT_OR_PR_TITLE` - default to the commit's title (if only one commit) or the pull request's title (when more than one commit).
	Milestones_url string `json:"milestones_url"`
	Statuses_url string `json:"statuses_url"`
	Created_at string `json:"created_at"`
}

// GeneratedType_Webhook_discussion_unlocked represents the GeneratedType_Webhook_discussion_unlocked schema from the OpenAPI specification
type GeneratedType_Webhook_discussion_unlocked struct {
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType_Simple_repository represents the GeneratedType_Simple_repository schema from the OpenAPI specification
type GeneratedType_Simple_repository struct {
	Owner GeneratedType_Simple_user `json:"owner"` // A GitHub user.
	Assignees_url string `json:"assignees_url"` // A template for the API URL to list the available assignees for issues in the repository.
	Issues_url string `json:"issues_url"` // A template for the API URL to get information about issues on the repository.
	Git_commits_url string `json:"git_commits_url"` // A template for the API URL to get information about Git commits of the repository.
	Compare_url string `json:"compare_url"` // A template for the API URL to compare two commits or refs.
	Tags_url string `json:"tags_url"` // The API URL to get information about tags on the repository.
	Url string `json:"url"` // The URL to get more information about the repository from the GitHub API.
	Full_name string `json:"full_name"` // The full, globally unique, name of the repository.
	Collaborators_url string `json:"collaborators_url"` // A template for the API URL to get information about collaborators of the repository.
	Git_refs_url string `json:"git_refs_url"` // A template for the API URL to get information about Git refs of the repository.
	Commits_url string `json:"commits_url"` // A template for the API URL to get information about commits on the repository.
	Contents_url string `json:"contents_url"` // A template for the API URL to get the contents of the repository.
	Downloads_url string `json:"downloads_url"` // The API URL to list the downloads on the repository.
	Events_url string `json:"events_url"` // The API URL to list the events of the repository.
	Trees_url string `json:"trees_url"` // A template for the API URL to create or retrieve a raw Git tree of the repository.
	Description string `json:"description"` // The repository description.
	Deployments_url string `json:"deployments_url"` // The API URL to list the deployments of the repository.
	Merges_url string `json:"merges_url"` // The API URL to merge branches in the repository.
	Private bool `json:"private"` // Whether the repository is private.
	Blobs_url string `json:"blobs_url"` // A template for the API URL to create or retrieve a raw Git blob in the repository.
	Name string `json:"name"` // The name of the repository.
	Releases_url string `json:"releases_url"` // A template for the API URL to get information about releases on the repository.
	Id int64 `json:"id"` // A unique identifier of the repository.
	Notifications_url string `json:"notifications_url"` // A template for the API URL to get information about notifications on the repository.
	Git_tags_url string `json:"git_tags_url"` // A template for the API URL to get information about Git tags of the repository.
	Keys_url string `json:"keys_url"` // A template for the API URL to get information about deploy keys on the repository.
	Forks_url string `json:"forks_url"` // The API URL to list the forks of the repository.
	Comments_url string `json:"comments_url"` // A template for the API URL to get information about comments on the repository.
	Subscribers_url string `json:"subscribers_url"` // The API URL to list the subscribers on the repository.
	Hooks_url string `json:"hooks_url"` // The API URL to list the hooks on the repository.
	Fork bool `json:"fork"` // Whether the repository is a fork.
	Node_id string `json:"node_id"` // The GraphQL identifier of the repository.
	Issue_comment_url string `json:"issue_comment_url"` // A template for the API URL to get information about issue comments on the repository.
	Pulls_url string `json:"pulls_url"` // A template for the API URL to get information about pull requests on the repository.
	Languages_url string `json:"languages_url"` // The API URL to get information about the languages of the repository.
	Archive_url string `json:"archive_url"` // A template for the API URL to download the repository as an archive.
	Subscription_url string `json:"subscription_url"` // The API URL to subscribe to notifications for this repository.
	Contributors_url string `json:"contributors_url"` // A template for the API URL to list the contributors to the repository.
	Html_url string `json:"html_url"` // The URL to view the repository on GitHub.com.
	Stargazers_url string `json:"stargazers_url"` // The API URL to list the stargazers on the repository.
	Statuses_url string `json:"statuses_url"` // A template for the API URL to get information about statuses of a commit.
	Issue_events_url string `json:"issue_events_url"` // A template for the API URL to get information about issue events on the repository.
	Milestones_url string `json:"milestones_url"` // A template for the API URL to get information about milestones of the repository.
	Labels_url string `json:"labels_url"` // A template for the API URL to get information about labels of the repository.
	Branches_url string `json:"branches_url"` // A template for the API URL to get information about branches in the repository.
	Teams_url string `json:"teams_url"` // The API URL to list the teams on the repository.
}

// GeneratedType_Webhook_discussion_edited represents the GeneratedType_Webhook_discussion_edited schema from the OpenAPI specification
type GeneratedType_Webhook_discussion_edited struct {
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes,omitempty"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
}

// GeneratedType_Issue_event_rename represents the GeneratedType_Issue_event_rename schema from the OpenAPI specification
type GeneratedType_Issue_event_rename struct {
	From string `json:"from"`
	To string `json:"to"`
}

// GeneratedType_Repository_rule_creation represents the GeneratedType_Repository_rule_creation schema from the OpenAPI specification
type GeneratedType_Repository_rule_creation struct {
	TypeField string `json:"type"`
}

// GeneratedType_Actions_get_default_workflow_permissions represents the GeneratedType_Actions_get_default_workflow_permissions schema from the OpenAPI specification
type GeneratedType_Actions_get_default_workflow_permissions struct {
	Can_approve_pull_request_reviews bool `json:"can_approve_pull_request_reviews"` // Whether GitHub Actions can approve pull requests. Enabling this can be a security risk.
	Default_workflow_permissions string `json:"default_workflow_permissions"` // The default workflow permissions granted to the GITHUB_TOKEN when running workflows.
}

// GeneratedType_Custom_deployment_rule_app represents the GeneratedType_Custom_deployment_rule_app schema from the OpenAPI specification
type GeneratedType_Custom_deployment_rule_app struct {
	Id int `json:"id"` // The unique identifier of the deployment protection rule integration.
	Integration_url string `json:"integration_url"` // The URL for the endpoint to get details about the app.
	Node_id string `json:"node_id"` // The node ID for the deployment protection rule integration.
	Slug string `json:"slug"` // The slugified name of the deployment protection rule integration.
}

// GeneratedType_Webhook_project_card_converted represents the GeneratedType_Webhook_project_card_converted schema from the OpenAPI specification
type GeneratedType_Webhook_project_card_converted struct {
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Project_card Webhooksprojectcard `json:"project_card"`
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
}

// GeneratedType_Code_security_configuration_repositories represents the GeneratedType_Code_security_configuration_repositories schema from the OpenAPI specification
type GeneratedType_Code_security_configuration_repositories struct {
	Repository GeneratedType_Simple_repository `json:"repository,omitempty"` // A GitHub repository.
	Status string `json:"status,omitempty"` // The attachment status of the code security configuration on the repository.
}

// GeneratedType_Webhook_issue_comment_created represents the GeneratedType_Webhook_issue_comment_created schema from the OpenAPI specification
type GeneratedType_Webhook_issue_comment_created struct {
	Comment map[string]interface{} `json:"comment"` // The [comment](https://docs.github.com/rest/issues/comments#get-an-issue-comment) itself.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) the comment belongs to.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType_Webhook_sponsorship_pending_cancellation represents the GeneratedType_Webhook_sponsorship_pending_cancellation schema from the OpenAPI specification
type GeneratedType_Webhook_sponsorship_pending_cancellation struct {
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Sponsorship Webhookssponsorship `json:"sponsorship"`
	Action string `json:"action"`
	Effective_date string `json:"effective_date,omitempty"` // The `pending_cancellation` and `pending_tier_change` event types will include the date the cancellation or tier change will take effect.
}

// GeneratedType_Webhook_sub_issues_sub_issue_removed represents the GeneratedType_Webhook_sub_issues_sub_issue_removed schema from the OpenAPI specification
type GeneratedType_Webhook_sub_issues_sub_issue_removed struct {
	Sub_issue_id float64 `json:"sub_issue_id"` // The ID of the sub-issue.
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Parent_issue_id float64 `json:"parent_issue_id"` // The ID of the parent issue.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sub_issue_repo Repository `json:"sub_issue_repo"` // A repository on GitHub.
	Parent_issue Issue `json:"parent_issue"` // Issues are a great way to keep track of tasks, enhancements, and bugs for your projects.
	Sub_issue Issue `json:"sub_issue"` // Issues are a great way to keep track of tasks, enhancements, and bugs for your projects.
}

// GeneratedType_Dependabot_alert_security_vulnerability represents the GeneratedType_Dependabot_alert_security_vulnerability schema from the OpenAPI specification
type GeneratedType_Dependabot_alert_security_vulnerability struct {
	First_patched_version map[string]interface{} `json:"first_patched_version"` // Details pertaining to the package version that patches this vulnerability.
	PackageField GeneratedType_Dependabot_alert_package `json:"package"` // Details for the vulnerable package.
	Severity string `json:"severity"` // The severity of the vulnerability.
	Vulnerable_version_range string `json:"vulnerable_version_range"` // Conditions that identify vulnerable versions of this vulnerability's package.
}

// GeneratedType_Webhook_team_add represents the GeneratedType_Webhook_team_add schema from the OpenAPI specification
type GeneratedType_Webhook_team_add struct {
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Team Webhooksteam1 `json:"team"` // Groups of organization members that gives permissions on specified repositories.
}

// GeneratedType_Ruleset_version represents the GeneratedType_Ruleset_version schema from the OpenAPI specification
type GeneratedType_Ruleset_version struct {
	Updated_at string `json:"updated_at"`
	Version_id int `json:"version_id"` // The ID of the previous version of the ruleset
	Actor map[string]interface{} `json:"actor"` // The actor who updated the ruleset
}

// GeneratedType_Webhook_watch_started represents the GeneratedType_Webhook_watch_started schema from the OpenAPI specification
type GeneratedType_Webhook_watch_started struct {
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType_Check_run represents the GeneratedType_Check_run schema from the OpenAPI specification
type GeneratedType_Check_run struct {
	App GeneratedType_Nullable_integration `json:"app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Pull_requests []GeneratedType_Pull_request_minimal `json:"pull_requests"` // Pull requests that are open with a `head_sha` or `head_branch` that matches the check. The returned pull requests do not necessarily indicate pull requests that triggered the check.
	Head_sha string `json:"head_sha"` // The SHA of the commit that is being checked.
	Conclusion string `json:"conclusion"`
	External_id string `json:"external_id"`
	Id int64 `json:"id"` // The id of the check.
	Node_id string `json:"node_id"`
	Output map[string]interface{} `json:"output"`
	Url string `json:"url"`
	Started_at string `json:"started_at"`
	Completed_at string `json:"completed_at"`
	Details_url string `json:"details_url"`
	Status string `json:"status"` // The phase of the lifecycle that the check is currently in. Statuses of waiting, requested, and pending are reserved for GitHub Actions check runs.
	Name string `json:"name"` // The name of the check.
	Deployment GeneratedType_Deployment_simple `json:"deployment,omitempty"` // A deployment created as the result of an Actions check run from a workflow that references an environment
	Html_url string `json:"html_url"`
	Check_suite map[string]interface{} `json:"check_suite"`
}

// GeneratedType_Dependabot_alert_security_advisory represents the GeneratedType_Dependabot_alert_security_advisory schema from the OpenAPI specification
type GeneratedType_Dependabot_alert_security_advisory struct {
	Description string `json:"description"` // A long-form Markdown-supported description of the advisory.
	Vulnerabilities []GeneratedType_Dependabot_alert_security_vulnerability `json:"vulnerabilities"` // Vulnerable version range information for the advisory.
	Severity string `json:"severity"` // The severity of the advisory.
	Identifiers []map[string]interface{} `json:"identifiers"` // Values that identify this advisory among security information sources.
	Cve_id string `json:"cve_id"` // The unique CVE ID assigned to the advisory.
	Cvss map[string]interface{} `json:"cvss"` // Details for the advisory pertaining to the Common Vulnerability Scoring System.
	References []map[string]interface{} `json:"references"` // Links to additional advisory information.
	Updated_at string `json:"updated_at"` // The time that the advisory was last modified in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Withdrawn_at string `json:"withdrawn_at"` // The time that the advisory was withdrawn in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Cvss_severities GeneratedType_Cvss_severities `json:"cvss_severities,omitempty"`
	Published_at string `json:"published_at"` // The time that the advisory was published in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Cwes []map[string]interface{} `json:"cwes"` // Details for the advisory pertaining to Common Weakness Enumeration.
	Epss GeneratedType_Security_advisory_epss `json:"epss,omitempty"` // The EPSS scores as calculated by the [Exploit Prediction Scoring System](https://www.first.org/epss).
	Ghsa_id string `json:"ghsa_id"` // The unique GitHub Security Advisory ID assigned to the advisory.
	Summary string `json:"summary"` // A short, plain text summary of the advisory.
}

// GeneratedType_Webhook_discussion_labeled represents the GeneratedType_Webhook_discussion_labeled schema from the OpenAPI specification
type GeneratedType_Webhook_discussion_labeled struct {
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Label Webhookslabel `json:"label"`
}

// Migration represents the Migration schema from the OpenAPI specification
type Migration struct {
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	Id int64 `json:"id"`
	Exclude_metadata bool `json:"exclude_metadata"`
	Exclude_owner_projects bool `json:"exclude_owner_projects"`
	Exclude []string `json:"exclude,omitempty"` // Exclude related items from being returned in the response in order to improve performance of the request. The array can include any of: `"repositories"`.
	Node_id string `json:"node_id"`
	Owner GeneratedType_Nullable_simple_user `json:"owner"` // A GitHub user.
	Lock_repositories bool `json:"lock_repositories"`
	Org_metadata_only bool `json:"org_metadata_only"`
	Exclude_attachments bool `json:"exclude_attachments"`
	Exclude_git_data bool `json:"exclude_git_data"`
	Repositories []Repository `json:"repositories"` // The repositories included in the migration. Only returned for export migrations.
	Archive_url string `json:"archive_url,omitempty"`
	Guid string `json:"guid"`
	Exclude_releases bool `json:"exclude_releases"`
	State string `json:"state"`
	Created_at string `json:"created_at"`
}

// GeneratedType_Webhook_project_card_moved represents the GeneratedType_Webhook_project_card_moved schema from the OpenAPI specification
type GeneratedType_Webhook_project_card_moved struct {
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Project_card interface{} `json:"project_card"`
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes,omitempty"`
}

// GeneratedType_Webhook_team_edited represents the GeneratedType_Webhook_team_edited schema from the OpenAPI specification
type GeneratedType_Webhook_team_edited struct {
	Repository map[string]interface{} `json:"repository,omitempty"` // A git repository
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Team Webhooksteam1 `json:"team"` // Groups of organization members that gives permissions on specified repositories.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"` // The changes to the team if the action was `edited`.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType_Webhook_pull_request_review_comment_edited represents the GeneratedType_Webhook_pull_request_review_comment_edited schema from the OpenAPI specification
type GeneratedType_Webhook_pull_request_review_comment_edited struct {
	Comment Webhooksreviewcomment `json:"comment"` // The [comment](https://docs.github.com/rest/pulls/comments#get-a-review-comment-for-a-pull-request) itself.
	Pull_request map[string]interface{} `json:"pull_request"`
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Action string `json:"action"`
	Changes Webhookschanges `json:"changes"` // The changes to the comment.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType_Webhook_check_suite_rerequested represents the GeneratedType_Webhook_check_suite_rerequested schema from the OpenAPI specification
type GeneratedType_Webhook_check_suite_rerequested struct {
	Action string `json:"action"`
	Check_suite map[string]interface{} `json:"check_suite"` // The [check_suite](https://docs.github.com/rest/checks/suites#get-a-check-suite).
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
}

// Autolink represents the Autolink schema from the OpenAPI specification
type Autolink struct {
	Id int `json:"id"`
	Is_alphanumeric bool `json:"is_alphanumeric"` // Whether this autolink reference matches alphanumeric characters. If false, this autolink reference only matches numeric characters.
	Key_prefix string `json:"key_prefix"` // The prefix of a key that is linkified.
	Url_template string `json:"url_template"` // A template for the target URL that is generated if a key was found.
}

// GeneratedType_Secret_scanning_location_pull_request_title represents the GeneratedType_Secret_scanning_location_pull_request_title schema from the OpenAPI specification
type GeneratedType_Secret_scanning_location_pull_request_title struct {
	Pull_request_title_url string `json:"pull_request_title_url"` // The API URL to get the pull request where the secret was detected.
}

// GeneratedType_Webhook_projects_v2_item_deleted represents the GeneratedType_Webhook_projects_v2_item_deleted schema from the OpenAPI specification
type GeneratedType_Webhook_projects_v2_item_deleted struct {
	Action string `json:"action"`
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Projects_v2_item GeneratedType_Projects_v2_item `json:"projects_v2_item"` // An item belonging to a project
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
}

// GeneratedType_Webhook_check_suite_completed represents the GeneratedType_Webhook_check_suite_completed schema from the OpenAPI specification
type GeneratedType_Webhook_check_suite_completed struct {
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Check_suite map[string]interface{} `json:"check_suite"` // The [check_suite](https://docs.github.com/rest/checks/suites#get-a-check-suite).
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType_Deploy_key represents the GeneratedType_Deploy_key schema from the OpenAPI specification
type GeneratedType_Deploy_key struct {
	Title string `json:"title"`
	Added_by string `json:"added_by,omitempty"`
	Id int `json:"id"`
	Key string `json:"key"`
	Read_only bool `json:"read_only"`
	Url string `json:"url"`
	Enabled bool `json:"enabled,omitempty"`
	Last_used string `json:"last_used,omitempty"`
	Verified bool `json:"verified"`
	Created_at string `json:"created_at"`
}

// GeneratedType_Deployment_status represents the GeneratedType_Deployment_status schema from the OpenAPI specification
type GeneratedType_Deployment_status struct {
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	Description string `json:"description"` // A short description of the status.
	State string `json:"state"` // The state of the status.
	Deployment_url string `json:"deployment_url"`
	Environment_url string `json:"environment_url,omitempty"` // The URL for accessing your environment.
	Log_url string `json:"log_url,omitempty"` // The URL to associate with this status.
	Repository_url string `json:"repository_url"`
	Node_id string `json:"node_id"`
	Performed_via_github_app GeneratedType_Nullable_integration `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Creator GeneratedType_Nullable_simple_user `json:"creator"` // A GitHub user.
	Target_url string `json:"target_url"` // Closing down notice: the URL to associate with this status.
	Created_at string `json:"created_at"`
	Id int64 `json:"id"`
	Environment string `json:"environment,omitempty"` // The environment of the deployment that the status is for.
}

// GeneratedType_Webhook_release_created represents the GeneratedType_Webhook_release_created schema from the OpenAPI specification
type GeneratedType_Webhook_release_created struct {
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Release Webhooksrelease `json:"release"` // The [release](https://docs.github.com/rest/releases/releases/#get-a-release) object.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
}

// GeneratedType_Moved_column_in_project_issue_event represents the GeneratedType_Moved_column_in_project_issue_event schema from the OpenAPI specification
type GeneratedType_Moved_column_in_project_issue_event struct {
	Event string `json:"event"`
	Created_at string `json:"created_at"`
	Node_id string `json:"node_id"`
	Performed_via_github_app GeneratedType_Nullable_integration `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Project_card map[string]interface{} `json:"project_card,omitempty"`
	Commit_id string `json:"commit_id"`
	Commit_url string `json:"commit_url"`
	Id int `json:"id"`
	Url string `json:"url"`
	Actor GeneratedType_Simple_user `json:"actor"` // A GitHub user.
}

// GeneratedType_Code_scanning_analysis represents the GeneratedType_Code_scanning_analysis schema from the OpenAPI specification
type GeneratedType_Code_scanning_analysis struct {
	Analysis_key string `json:"analysis_key"` // Identifies the configuration under which the analysis was executed. For example, in GitHub Actions this includes the workflow filename and job name.
	Id int `json:"id"` // Unique identifier for this analysis.
	Category string `json:"category,omitempty"` // Identifies the configuration under which the analysis was executed. Used to distinguish between multiple analyses for the same tool and commit, but performed on different languages or different parts of the code.
	ErrorField string `json:"error"`
	Commit_sha string `json:"commit_sha"` // The SHA of the commit to which the analysis you are uploading relates.
	Deletable bool `json:"deletable"`
	Environment string `json:"environment"` // Identifies the variable values associated with the environment in which this analysis was performed.
	Url string `json:"url"` // The REST API URL of the analysis resource.
	Tool GeneratedType_Code_scanning_analysis_tool `json:"tool"`
	Warning string `json:"warning"` // Warning generated when processing the analysis
	Results_count int `json:"results_count"` // The total number of results in the analysis.
	Created_at string `json:"created_at"` // The time that the analysis was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Ref string `json:"ref"` // The Git reference, formatted as `refs/pull/<number>/merge`, `refs/pull/<number>/head`, `refs/heads/<branch name>` or simply `<branch name>`.
	Rules_count int `json:"rules_count"` // The total number of rules used in the analysis.
	Sarif_id string `json:"sarif_id"` // An identifier for the upload.
}

// GeneratedType_Webhook_branch_protection_rule_created represents the GeneratedType_Webhook_branch_protection_rule_created schema from the OpenAPI specification
type GeneratedType_Webhook_branch_protection_rule_created struct {
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Rule Webhooksrule `json:"rule"` // The branch protection rule. Includes a `name` and all the [branch protection settings](https://docs.github.com/github/administering-a-repository/defining-the-mergeability-of-pull-requests/about-protected-branches#about-branch-protection-settings) applied to branches that match the name. Binary settings are boolean. Multi-level configurations are one of `off`, `non_admins`, or `everyone`. Actor and build lists are arrays of strings.
}

// Codespace represents the Codespace schema from the OpenAPI specification
type Codespace struct {
	Display_name string `json:"display_name,omitempty"` // Display name for this codespace.
	Git_status map[string]interface{} `json:"git_status"` // Details about the codespace's git repository.
	Url string `json:"url"` // API URL for this codespace.
	Last_known_stop_notice string `json:"last_known_stop_notice,omitempty"` // The text to display to a user when a codespace has been stopped for a potentially actionable reason.
	Pending_operation_disabled_reason string `json:"pending_operation_disabled_reason,omitempty"` // Text to show user when codespace is disabled by a pending operation
	Prebuild bool `json:"prebuild"` // Whether the codespace was created from a prebuild.
	Publish_url string `json:"publish_url,omitempty"` // API URL to publish this codespace to a new repository.
	Location string `json:"location"` // The initally assigned location of a new codespace.
	Stop_url string `json:"stop_url"` // API URL to stop this codespace.
	Owner GeneratedType_Simple_user `json:"owner"` // A GitHub user.
	Idle_timeout_minutes int `json:"idle_timeout_minutes"` // The number of minutes of inactivity after which this codespace will be automatically stopped.
	Last_used_at string `json:"last_used_at"` // Last known time this codespace was started.
	Pending_operation bool `json:"pending_operation,omitempty"` // Whether or not a codespace has a pending async operation. This would mean that the codespace is temporarily unavailable. The only thing that you can do with a codespace in this state is delete it.
	Repository GeneratedType_Minimal_repository `json:"repository"` // Minimal Repository
	Recent_folders []string `json:"recent_folders"`
	Billable_owner GeneratedType_Simple_user `json:"billable_owner"` // A GitHub user.
	Start_url string `json:"start_url"` // API URL to start this codespace.
	Web_url string `json:"web_url"` // URL to access this codespace on the web.
	Machine GeneratedType_Nullable_codespace_machine `json:"machine"` // A description of the machine powering a codespace.
	Pulls_url string `json:"pulls_url"` // API URL for the Pull Request associated with this codespace, if any.
	Retention_expires_at string `json:"retention_expires_at,omitempty"` // When a codespace will be auto-deleted based on the "retention_period_minutes" and "last_used_at"
	Runtime_constraints map[string]interface{} `json:"runtime_constraints,omitempty"`
	Name string `json:"name"` // Automatically generated name of this codespace.
	State string `json:"state"` // State of this codespace.
	Devcontainer_path string `json:"devcontainer_path,omitempty"` // Path to devcontainer.json from repo root used to create Codespace.
	Created_at string `json:"created_at"`
	Environment_id string `json:"environment_id"` // UUID identifying this codespace's environment.
	Updated_at string `json:"updated_at"`
	Id int64 `json:"id"`
	Retention_period_minutes int `json:"retention_period_minutes,omitempty"` // Duration in minutes after codespace has gone idle in which it will be deleted. Must be integer minutes between 0 and 43200 (30 days).
	Machines_url string `json:"machines_url"` // API URL to access available alternate machine types for this codespace.
	Idle_timeout_notice string `json:"idle_timeout_notice,omitempty"` // Text to show user when codespace idle timeout minutes has been overriden by an organization policy
}

// GeneratedType_Webhook_secret_scanning_alert_location_created represents the GeneratedType_Webhook_secret_scanning_alert_location_created schema from the OpenAPI specification
type GeneratedType_Webhook_secret_scanning_alert_location_created struct {
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action,omitempty"`
	Alert GeneratedType_Secret_scanning_alert_webhook `json:"alert"`
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Location GeneratedType_Secret_scanning_location `json:"location"`
}

// GeneratedType_Secret_scanning_scan_history represents the GeneratedType_Secret_scanning_scan_history schema from the OpenAPI specification
type GeneratedType_Secret_scanning_scan_history struct {
	Custom_pattern_backfill_scans []interface{} `json:"custom_pattern_backfill_scans,omitempty"`
	Incremental_scans []GeneratedType_Secret_scanning_scan `json:"incremental_scans,omitempty"`
	Pattern_update_scans []GeneratedType_Secret_scanning_scan `json:"pattern_update_scans,omitempty"`
	Backfill_scans []GeneratedType_Secret_scanning_scan `json:"backfill_scans,omitempty"`
}

// GeneratedType_Repository_ruleset_conditions_repository_property_target represents the GeneratedType_Repository_ruleset_conditions_repository_property_target schema from the OpenAPI specification
type GeneratedType_Repository_ruleset_conditions_repository_property_target struct {
	Repository_property map[string]interface{} `json:"repository_property"`
}

// GeneratedType_Installation_token represents the GeneratedType_Installation_token schema from the OpenAPI specification
type GeneratedType_Installation_token struct {
	Repository_selection string `json:"repository_selection,omitempty"`
	Single_file string `json:"single_file,omitempty"`
	Single_file_paths []string `json:"single_file_paths,omitempty"`
	Token string `json:"token"`
	Expires_at string `json:"expires_at"`
	Has_multiple_single_files bool `json:"has_multiple_single_files,omitempty"`
	Permissions GeneratedType_App_permissions `json:"permissions,omitempty"` // The permissions granted to the user access token.
	Repositories []Repository `json:"repositories,omitempty"`
}

// GeneratedType_Projects_v2_item represents the GeneratedType_Projects_v2_item schema from the OpenAPI specification
type GeneratedType_Projects_v2_item struct {
	Content_node_id string `json:"content_node_id"`
	Content_type string `json:"content_type"` // The type of content tracked in a project item
	Project_node_id string `json:"project_node_id,omitempty"`
	Archived_at string `json:"archived_at"`
	Created_at string `json:"created_at"`
	Id float64 `json:"id"`
	Node_id string `json:"node_id,omitempty"`
	Updated_at string `json:"updated_at"`
	Creator GeneratedType_Simple_user `json:"creator,omitempty"` // A GitHub user.
}

// Runner represents the Runner schema from the OpenAPI specification
type Runner struct {
	Status string `json:"status"` // The status of the runner.
	Busy bool `json:"busy"`
	Ephemeral bool `json:"ephemeral,omitempty"`
	Id int `json:"id"` // The ID of the runner.
	Labels []GeneratedType_Runner_label `json:"labels"`
	Name string `json:"name"` // The name of the runner.
	Os string `json:"os"` // The Operating System of the runner.
	Runner_group_id int `json:"runner_group_id,omitempty"` // The ID of the runner group.
}

// Webhooksworkflow represents the Webhooksworkflow schema from the OpenAPI specification
type Webhooksworkflow struct {
	Node_id string `json:"node_id"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	Path string `json:"path"`
	Html_url string `json:"html_url"`
	State string `json:"state"`
	Badge_url string `json:"badge_url"`
	Created_at string `json:"created_at"`
	Id int `json:"id"`
	Name string `json:"name"`
}

// GeneratedType_Webhook_code_scanning_alert_fixed represents the GeneratedType_Webhook_code_scanning_alert_fixed schema from the OpenAPI specification
type GeneratedType_Webhook_code_scanning_alert_fixed struct {
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Ref string `json:"ref"` // The Git reference of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Alert map[string]interface{} `json:"alert"` // The code scanning alert involved in the event.
	Commit_oid string `json:"commit_oid"` // The commit SHA of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Action string `json:"action"`
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// Vulnerability represents the Vulnerability schema from the OpenAPI specification
type Vulnerability struct {
	First_patched_version string `json:"first_patched_version"` // The package version that resolves the vulnerability.
	PackageField map[string]interface{} `json:"package"` // The name of the package affected by the vulnerability.
	Vulnerable_functions []string `json:"vulnerable_functions"` // The functions in the package that are affected by the vulnerability.
	Vulnerable_version_range string `json:"vulnerable_version_range"` // The range of the package versions affected by the vulnerability.
}

// GeneratedType_Git_tree represents the GeneratedType_Git_tree schema from the OpenAPI specification
type GeneratedType_Git_tree struct {
	Truncated bool `json:"truncated"`
	Url string `json:"url,omitempty"`
	Sha string `json:"sha"`
	Tree []map[string]interface{} `json:"tree"` // Objects specifying a tree structure
}

// Webhookspreviousmarketplacepurchase represents the Webhookspreviousmarketplacepurchase schema from the OpenAPI specification
type Webhookspreviousmarketplacepurchase struct {
	Next_billing_date string `json:"next_billing_date,omitempty"`
	On_free_trial bool `json:"on_free_trial"`
	Plan map[string]interface{} `json:"plan"`
	Unit_count int `json:"unit_count"`
	Account map[string]interface{} `json:"account"`
	Billing_cycle string `json:"billing_cycle"`
	Free_trial_ends_on interface{} `json:"free_trial_ends_on"`
}

// GeneratedType_Webhook_discussion_comment_created represents the GeneratedType_Webhook_discussion_comment_created schema from the OpenAPI specification
type GeneratedType_Webhook_discussion_comment_created struct {
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Comment Webhookscomment `json:"comment"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType_Dependabot_alert_package represents the GeneratedType_Dependabot_alert_package schema from the OpenAPI specification
type GeneratedType_Dependabot_alert_package struct {
	Ecosystem string `json:"ecosystem"` // The package's language or package management ecosystem.
	Name string `json:"name"` // The unique package name within its ecosystem.
}

// Tag represents the Tag schema from the OpenAPI specification
type Tag struct {
	Tarball_url string `json:"tarball_url"`
	Zipball_url string `json:"zipball_url"`
	Commit map[string]interface{} `json:"commit"`
	Name string `json:"name"`
	Node_id string `json:"node_id"`
}

// GeneratedType_Demilestoned_issue_event represents the GeneratedType_Demilestoned_issue_event schema from the OpenAPI specification
type GeneratedType_Demilestoned_issue_event struct {
	Commit_id string `json:"commit_id"`
	Created_at string `json:"created_at"`
	Node_id string `json:"node_id"`
	Url string `json:"url"`
	Commit_url string `json:"commit_url"`
	Event string `json:"event"`
	Milestone map[string]interface{} `json:"milestone"`
	Actor GeneratedType_Simple_user `json:"actor"` // A GitHub user.
	Performed_via_github_app GeneratedType_Nullable_integration `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Id int `json:"id"`
}

// GeneratedType_Check_suite represents the GeneratedType_Check_suite schema from the OpenAPI specification
type GeneratedType_Check_suite struct {
	Check_runs_url string `json:"check_runs_url"`
	Updated_at string `json:"updated_at"`
	After string `json:"after"`
	Conclusion string `json:"conclusion"`
	Created_at string `json:"created_at"`
	Repository GeneratedType_Minimal_repository `json:"repository"` // Minimal Repository
	Runs_rerequestable bool `json:"runs_rerequestable,omitempty"`
	Before string `json:"before"`
	Head_sha string `json:"head_sha"` // The SHA of the head commit that is being checked.
	Id int64 `json:"id"`
	Pull_requests []GeneratedType_Pull_request_minimal `json:"pull_requests"`
	App GeneratedType_Nullable_integration `json:"app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Status string `json:"status"` // The phase of the lifecycle that the check suite is currently in. Statuses of waiting, requested, and pending are reserved for GitHub Actions check suites.
	Url string `json:"url"`
	Latest_check_runs_count int `json:"latest_check_runs_count"`
	Node_id string `json:"node_id"`
	Head_branch string `json:"head_branch"`
	Head_commit GeneratedType_Simple_commit `json:"head_commit"` // A commit.
	Rerequestable bool `json:"rerequestable,omitempty"`
}

// GeneratedType_Webhook_public represents the GeneratedType_Webhook_public schema from the OpenAPI specification
type GeneratedType_Webhook_public struct {
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
}

// GeneratedType_Webhook_rubygems_metadata represents the GeneratedType_Webhook_rubygems_metadata schema from the OpenAPI specification
type GeneratedType_Webhook_rubygems_metadata struct {
	Dependencies []map[string]interface{} `json:"dependencies,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Repo string `json:"repo,omitempty"`
	Description string `json:"description,omitempty"`
	Name string `json:"name,omitempty"`
	Version_info map[string]interface{} `json:"version_info,omitempty"`
	Commit_oid string `json:"commit_oid,omitempty"`
	Homepage string `json:"homepage,omitempty"`
	Platform string `json:"platform,omitempty"`
	Readme string `json:"readme,omitempty"`
}

// GeneratedType_Api_overview represents the GeneratedType_Api_overview schema from the OpenAPI specification
type GeneratedType_Api_overview struct {
	Verifiable_password_authentication bool `json:"verifiable_password_authentication"`
	Pages []string `json:"pages,omitempty"`
	Copilot []string `json:"copilot,omitempty"`
	Actions []string `json:"actions,omitempty"`
	Hooks []string `json:"hooks,omitempty"`
	Api []string `json:"api,omitempty"`
	Web []string `json:"web,omitempty"`
	Importer []string `json:"importer,omitempty"`
	Github_enterprise_importer []string `json:"github_enterprise_importer,omitempty"`
	Packages []string `json:"packages,omitempty"`
	Codespaces []string `json:"codespaces,omitempty"`
	Dependabot []string `json:"dependabot,omitempty"`
	Ssh_keys []string `json:"ssh_keys,omitempty"`
	Ssh_key_fingerprints map[string]interface{} `json:"ssh_key_fingerprints,omitempty"`
	Actions_macos []string `json:"actions_macos,omitempty"`
	Git []string `json:"git,omitempty"`
	Domains map[string]interface{} `json:"domains,omitempty"`
}

// GeneratedType_Pull_request_merge_result represents the GeneratedType_Pull_request_merge_result schema from the OpenAPI specification
type GeneratedType_Pull_request_merge_result struct {
	Merged bool `json:"merged"`
	Message string `json:"message"`
	Sha string `json:"sha"`
}

// GeneratedType_Webhook_sponsorship_pending_tier_change represents the GeneratedType_Webhook_sponsorship_pending_tier_change schema from the OpenAPI specification
type GeneratedType_Webhook_sponsorship_pending_tier_change struct {
	Changes Webhookschanges8 `json:"changes"`
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Sponsorship Webhookssponsorship `json:"sponsorship"`
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Effective_date string `json:"effective_date,omitempty"` // The `pending_cancellation` and `pending_tier_change` event types will include the date the cancellation or tier change will take effect.
}

// GeneratedType_Check_suite_preference represents the GeneratedType_Check_suite_preference schema from the OpenAPI specification
type GeneratedType_Check_suite_preference struct {
	Preferences map[string]interface{} `json:"preferences"`
	Repository GeneratedType_Minimal_repository `json:"repository"` // Minimal Repository
}

// GeneratedType_Copilot_organization_seat_breakdown represents the GeneratedType_Copilot_organization_seat_breakdown schema from the OpenAPI specification
type GeneratedType_Copilot_organization_seat_breakdown struct {
	Active_this_cycle int `json:"active_this_cycle,omitempty"` // The number of seats that have used Copilot during the current billing cycle.
	Added_this_cycle int `json:"added_this_cycle,omitempty"` // Seats added during the current billing cycle.
	Inactive_this_cycle int `json:"inactive_this_cycle,omitempty"` // The number of seats that have not used Copilot during the current billing cycle.
	Pending_cancellation int `json:"pending_cancellation,omitempty"` // The number of seats that are pending cancellation at the end of the current billing cycle.
	Pending_invitation int `json:"pending_invitation,omitempty"` // The number of users who have been invited to receive a Copilot seat through this organization.
	Total int `json:"total,omitempty"` // The total number of seats being billed for the organization as of the current billing cycle.
}

// GeneratedType_Simple_commit represents the GeneratedType_Simple_commit schema from the OpenAPI specification
type GeneratedType_Simple_commit struct {
	Timestamp string `json:"timestamp"` // Timestamp of the commit
	Tree_id string `json:"tree_id"` // SHA for the commit's tree
	Author map[string]interface{} `json:"author"` // Information about the Git author
	Committer map[string]interface{} `json:"committer"` // Information about the Git committer
	Id string `json:"id"` // SHA for the commit
	Message string `json:"message"` // Message describing the purpose of the commit
}

// GeneratedType_Webhook_pull_request_auto_merge_disabled represents the GeneratedType_Webhook_pull_request_auto_merge_disabled schema from the OpenAPI specification
type GeneratedType_Webhook_pull_request_auto_merge_disabled struct {
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Reason string `json:"reason"`
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Number int `json:"number"`
	Pull_request map[string]interface{} `json:"pull_request"`
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// Webhooksreviewcomment represents the Webhooksreviewcomment schema from the OpenAPI specification
type Webhooksreviewcomment struct {
	Pull_request_review_id int `json:"pull_request_review_id"` // The ID of the pull request review to which the comment belongs.
	Original_position int `json:"original_position"` // The index of the original line in the diff to which the comment applies.
	Created_at string `json:"created_at"`
	Original_start_line int `json:"original_start_line"` // The first line of the range for a multi-line comment.
	Pull_request_url string `json:"pull_request_url"` // URL for the pull request that the review comment belongs to.
	Updated_at string `json:"updated_at"`
	Url string `json:"url"` // URL for the pull request review comment
	Position int `json:"position"` // The line index in the diff to which the comment applies.
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Diff_hunk string `json:"diff_hunk"` // The diff of the line that the comment refers to.
	In_reply_to_id int `json:"in_reply_to_id,omitempty"` // The comment ID to reply to.
	Line int `json:"line"` // The line of the blob to which the comment applies. The last line of the range for a multi-line comment
	Reactions map[string]interface{} `json:"reactions"`
	Subject_type string `json:"subject_type,omitempty"` // The level at which the comment is targeted, can be a diff line or a file.
	User map[string]interface{} `json:"user"`
	Commit_id string `json:"commit_id"` // The SHA of the commit to which the comment applies.
	Html_url string `json:"html_url"` // HTML URL for the pull request review comment.
	Node_id string `json:"node_id"` // The node ID of the pull request review comment.
	Side string `json:"side"` // The side of the first line of the range for a multi-line comment.
	Start_line int `json:"start_line"` // The first line of the range for a multi-line comment.
	Links map[string]interface{} `json:"_links"`
	Body string `json:"body"` // The text of the comment.
	Original_commit_id string `json:"original_commit_id"` // The SHA of the original commit to which the comment applies.
	Id int `json:"id"` // The ID of the pull request review comment.
	Path string `json:"path"` // The relative path of the file to which the comment applies.
	Start_side string `json:"start_side"` // The side of the first line of the range for a multi-line comment.
	Original_line int `json:"original_line"` // The line of the blob to which the comment applies. The last line of the range for a multi-line comment
}

// GeneratedType_Webhook_projects_v2_item_edited represents the GeneratedType_Webhook_projects_v2_item_edited schema from the OpenAPI specification
type GeneratedType_Webhook_projects_v2_item_edited struct {
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Projects_v2_item GeneratedType_Projects_v2_item `json:"projects_v2_item"` // An item belonging to a project
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes interface{} `json:"changes,omitempty"` // The changes made to the item may involve modifications in the item's fields and draft issue body. It includes altered values for text, number, date, single select, and iteration fields, along with the GraphQL node ID of the changed field.
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType_Nullable_simple_repository represents the GeneratedType_Nullable_simple_repository schema from the OpenAPI specification
type GeneratedType_Nullable_simple_repository struct {
	Trees_url string `json:"trees_url"` // A template for the API URL to create or retrieve a raw Git tree of the repository.
	Git_tags_url string `json:"git_tags_url"` // A template for the API URL to get information about Git tags of the repository.
	Collaborators_url string `json:"collaborators_url"` // A template for the API URL to get information about collaborators of the repository.
	Contributors_url string `json:"contributors_url"` // A template for the API URL to list the contributors to the repository.
	Pulls_url string `json:"pulls_url"` // A template for the API URL to get information about pull requests on the repository.
	Archive_url string `json:"archive_url"` // A template for the API URL to download the repository as an archive.
	Blobs_url string `json:"blobs_url"` // A template for the API URL to create or retrieve a raw Git blob in the repository.
	Hooks_url string `json:"hooks_url"` // The API URL to list the hooks on the repository.
	Private bool `json:"private"` // Whether the repository is private.
	Description string `json:"description"` // The repository description.
	Git_refs_url string `json:"git_refs_url"` // A template for the API URL to get information about Git refs of the repository.
	Url string `json:"url"` // The URL to get more information about the repository from the GitHub API.
	Deployments_url string `json:"deployments_url"` // The API URL to list the deployments of the repository.
	Git_commits_url string `json:"git_commits_url"` // A template for the API URL to get information about Git commits of the repository.
	Labels_url string `json:"labels_url"` // A template for the API URL to get information about labels of the repository.
	Issues_url string `json:"issues_url"` // A template for the API URL to get information about issues on the repository.
	Milestones_url string `json:"milestones_url"` // A template for the API URL to get information about milestones of the repository.
	Html_url string `json:"html_url"` // The URL to view the repository on GitHub.com.
	Tags_url string `json:"tags_url"` // The API URL to get information about tags on the repository.
	Name string `json:"name"` // The name of the repository.
	Statuses_url string `json:"statuses_url"` // A template for the API URL to get information about statuses of a commit.
	Downloads_url string `json:"downloads_url"` // The API URL to list the downloads on the repository.
	Notifications_url string `json:"notifications_url"` // A template for the API URL to get information about notifications on the repository.
	Issue_comment_url string `json:"issue_comment_url"` // A template for the API URL to get information about issue comments on the repository.
	Owner GeneratedType_Simple_user `json:"owner"` // A GitHub user.
	Contents_url string `json:"contents_url"` // A template for the API URL to get the contents of the repository.
	Node_id string `json:"node_id"` // The GraphQL identifier of the repository.
	Subscribers_url string `json:"subscribers_url"` // The API URL to list the subscribers on the repository.
	Fork bool `json:"fork"` // Whether the repository is a fork.
	Full_name string `json:"full_name"` // The full, globally unique, name of the repository.
	Assignees_url string `json:"assignees_url"` // A template for the API URL to list the available assignees for issues in the repository.
	Forks_url string `json:"forks_url"` // The API URL to list the forks of the repository.
	Languages_url string `json:"languages_url"` // The API URL to get information about the languages of the repository.
	Branches_url string `json:"branches_url"` // A template for the API URL to get information about branches in the repository.
	Comments_url string `json:"comments_url"` // A template for the API URL to get information about comments on the repository.
	Events_url string `json:"events_url"` // The API URL to list the events of the repository.
	Teams_url string `json:"teams_url"` // The API URL to list the teams on the repository.
	Keys_url string `json:"keys_url"` // A template for the API URL to get information about deploy keys on the repository.
	Stargazers_url string `json:"stargazers_url"` // The API URL to list the stargazers on the repository.
	Subscription_url string `json:"subscription_url"` // The API URL to subscribe to notifications for this repository.
	Compare_url string `json:"compare_url"` // A template for the API URL to compare two commits or refs.
	Issue_events_url string `json:"issue_events_url"` // A template for the API URL to get information about issue events on the repository.
	Releases_url string `json:"releases_url"` // A template for the API URL to get information about releases on the repository.
	Commits_url string `json:"commits_url"` // A template for the API URL to get information about commits on the repository.
	Id int64 `json:"id"` // A unique identifier of the repository.
	Merges_url string `json:"merges_url"` // The API URL to merge branches in the repository.
}

// GeneratedType_Package_version represents the GeneratedType_Package_version schema from the OpenAPI specification
type GeneratedType_Package_version struct {
	Created_at string `json:"created_at"`
	Id int `json:"id"` // Unique identifier of the package version.
	Package_html_url string `json:"package_html_url"`
	Deleted_at string `json:"deleted_at,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Name string `json:"name"` // The name of the package version.
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	Description string `json:"description,omitempty"`
	Html_url string `json:"html_url,omitempty"`
	License string `json:"license,omitempty"`
}

// GeneratedType_Repository_rule_commit_author_email_pattern represents the GeneratedType_Repository_rule_commit_author_email_pattern schema from the OpenAPI specification
type GeneratedType_Repository_rule_commit_author_email_pattern struct {
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	TypeField string `json:"type"`
}

// GeneratedType_Code_scanning_alert_items represents the GeneratedType_Code_scanning_alert_items schema from the OpenAPI specification
type GeneratedType_Code_scanning_alert_items struct {
	Updated_at string `json:"updated_at,omitempty"` // The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	State string `json:"state"` // State of a code scanning alert.
	Most_recent_instance GeneratedType_Code_scanning_alert_instance `json:"most_recent_instance"`
	Dismissed_comment string `json:"dismissed_comment,omitempty"` // The dismissal comment associated with the dismissal of the alert.
	Tool GeneratedType_Code_scanning_analysis_tool `json:"tool"`
	Html_url string `json:"html_url"` // The GitHub URL of the alert resource.
	Dismissed_by GeneratedType_Nullable_simple_user `json:"dismissed_by"` // A GitHub user.
	Dismissed_at string `json:"dismissed_at"` // The time that the alert was dismissed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Dismissed_reason string `json:"dismissed_reason"` // **Required when the state is dismissed.** The reason for dismissing or closing the alert.
	Url string `json:"url"` // The REST API URL of the alert resource.
	Instances_url string `json:"instances_url"` // The REST API URL for fetching the list of instances for an alert.
	Fixed_at string `json:"fixed_at,omitempty"` // The time that the alert was no longer detected and was considered fixed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Rule GeneratedType_Code_scanning_alert_rule_summary `json:"rule"`
	Dismissal_approved_by GeneratedType_Nullable_simple_user `json:"dismissal_approved_by,omitempty"` // A GitHub user.
	Number int `json:"number"` // The security alert number.
	Created_at string `json:"created_at"` // The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
}

// GeneratedType_Webhook_security_advisory_withdrawn represents the GeneratedType_Webhook_security_advisory_withdrawn schema from the OpenAPI specification
type GeneratedType_Webhook_security_advisory_withdrawn struct {
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Security_advisory map[string]interface{} `json:"security_advisory"` // The details of the security advisory, including summary, description, and severity.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
}

// GeneratedType_Participation_stats represents the GeneratedType_Participation_stats schema from the OpenAPI specification
type GeneratedType_Participation_stats struct {
	Owner []int `json:"owner"`
	All []int `json:"all"`
}

// GeneratedType_Webhook_secret_scanning_alert_validated represents the GeneratedType_Webhook_secret_scanning_alert_validated schema from the OpenAPI specification
type GeneratedType_Webhook_secret_scanning_alert_validated struct {
	Action string `json:"action"`
	Alert GeneratedType_Secret_scanning_alert_webhook `json:"alert"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
}

// GeneratedType_Webhook_issues_unpinned represents the GeneratedType_Webhook_issues_unpinned schema from the OpenAPI specification
type GeneratedType_Webhook_issues_unpinned struct {
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue Webhooksissue2 `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType_Webhook_projects_v2_project_reopened represents the GeneratedType_Webhook_projects_v2_project_reopened schema from the OpenAPI specification
type GeneratedType_Webhook_projects_v2_project_reopened struct {
	Projects_v2 GeneratedType_Projects_v2 `json:"projects_v2"` // A projects v2 project
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType_Webhook_dependabot_alert_auto_dismissed represents the GeneratedType_Webhook_dependabot_alert_auto_dismissed schema from the OpenAPI specification
type GeneratedType_Webhook_dependabot_alert_auto_dismissed struct {
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Alert GeneratedType_Dependabot_alert `json:"alert"` // A Dependabot alert.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType_Webhook_issues_untyped represents the GeneratedType_Webhook_issues_untyped schema from the OpenAPI specification
type GeneratedType_Webhook_issues_untyped struct {
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue Webhooksissue `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	TypeField GeneratedType_Issue_type `json:"type"` // The type of issue.
}

// GeneratedType_Release_notes_content represents the GeneratedType_Release_notes_content schema from the OpenAPI specification
type GeneratedType_Release_notes_content struct {
	Body string `json:"body"` // The generated body describing the contents of the release supporting markdown formatting
	Name string `json:"name"` // The generated name of the release
}

// GeneratedType_Webhook_check_run_created represents the GeneratedType_Webhook_check_run_created schema from the OpenAPI specification
type GeneratedType_Webhook_check_run_created struct {
	Check_run GeneratedType_Check_run_with_simple_check_suite `json:"check_run"` // A check performed on the code of a given code change
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action,omitempty"`
}

// GeneratedType_Webhook_repository_dispatch_sample represents the GeneratedType_Webhook_repository_dispatch_sample schema from the OpenAPI specification
type GeneratedType_Webhook_repository_dispatch_sample struct {
	Installation GeneratedType_Simple_installation `json:"installation"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"` // The `event_type` that was specified in the `POST /repos/{owner}/{repo}/dispatches` request body.
	Branch string `json:"branch"`
	Client_payload map[string]interface{} `json:"client_payload"` // The `client_payload` that was specified in the `POST /repos/{owner}/{repo}/dispatches` request body.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType_Webhook_project_edited represents the GeneratedType_Webhook_project_edited schema from the OpenAPI specification
type GeneratedType_Webhook_project_edited struct {
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Project Webhooksproject `json:"project"`
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes,omitempty"` // The changes to the project if the action was `edited`.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType_Nullable_minimal_repository represents the GeneratedType_Nullable_minimal_repository schema from the OpenAPI specification
type GeneratedType_Nullable_minimal_repository struct {
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Archived bool `json:"archived,omitempty"`
	Stargazers_count int `json:"stargazers_count,omitempty"`
	Updated_at string `json:"updated_at,omitempty"`
	Node_id string `json:"node_id"`
	Default_branch string `json:"default_branch,omitempty"`
	Name string `json:"name"`
	Forks int `json:"forks,omitempty"`
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"`
	Has_downloads bool `json:"has_downloads,omitempty"`
	Downloads_url string `json:"downloads_url"`
	Events_url string `json:"events_url"`
	Pushed_at string `json:"pushed_at,omitempty"`
	Open_issues_count int `json:"open_issues_count,omitempty"`
	Subscription_url string `json:"subscription_url"`
	Created_at string `json:"created_at,omitempty"`
	Pulls_url string `json:"pulls_url"`
	Topics []string `json:"topics,omitempty"`
	Description string `json:"description"`
	Git_url string `json:"git_url,omitempty"`
	Branches_url string `json:"branches_url"`
	Tags_url string `json:"tags_url"`
	Network_count int `json:"network_count,omitempty"`
	Security_and_analysis GeneratedType_Security_and_analysis `json:"security_and_analysis,omitempty"`
	License map[string]interface{} `json:"license,omitempty"`
	Merges_url string `json:"merges_url"`
	Issue_events_url string `json:"issue_events_url"`
	Url string `json:"url"`
	Assignees_url string `json:"assignees_url"`
	Allow_forking bool `json:"allow_forking,omitempty"`
	Code_of_conduct GeneratedType_Code_of_conduct `json:"code_of_conduct,omitempty"` // Code Of Conduct
	Fork bool `json:"fork"`
	Temp_clone_token string `json:"temp_clone_token,omitempty"`
	Git_tags_url string `json:"git_tags_url"`
	Owner GeneratedType_Simple_user `json:"owner"` // A GitHub user.
	Hooks_url string `json:"hooks_url"`
	Compare_url string `json:"compare_url"`
	Has_discussions bool `json:"has_discussions,omitempty"`
	Git_refs_url string `json:"git_refs_url"`
	Subscribers_count int `json:"subscribers_count,omitempty"`
	Releases_url string `json:"releases_url"`
	Language string `json:"language,omitempty"`
	Has_projects bool `json:"has_projects,omitempty"`
	Full_name string `json:"full_name"`
	Private bool `json:"private"`
	Forks_url string `json:"forks_url"`
	Stargazers_url string `json:"stargazers_url"`
	Deployments_url string `json:"deployments_url"`
	Collaborators_url string `json:"collaborators_url"`
	Has_issues bool `json:"has_issues,omitempty"`
	Open_issues int `json:"open_issues,omitempty"`
	Blobs_url string `json:"blobs_url"`
	Watchers int `json:"watchers,omitempty"`
	Archive_url string `json:"archive_url"`
	Keys_url string `json:"keys_url"`
	Languages_url string `json:"languages_url"`
	Svn_url string `json:"svn_url,omitempty"`
	Issue_comment_url string `json:"issue_comment_url"`
	Teams_url string `json:"teams_url"`
	Mirror_url string `json:"mirror_url,omitempty"`
	Homepage string `json:"homepage,omitempty"`
	Visibility string `json:"visibility,omitempty"`
	Id int64 `json:"id"`
	Size int `json:"size,omitempty"` // The size of the repository, in kilobytes. Size is calculated hourly. When a repository is initially created, the size is 0.
	Ssh_url string `json:"ssh_url,omitempty"`
	Disabled bool `json:"disabled,omitempty"`
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"`
	Watchers_count int `json:"watchers_count,omitempty"`
	Comments_url string `json:"comments_url"`
	Forks_count int `json:"forks_count,omitempty"`
	Milestones_url string `json:"milestones_url"`
	Commits_url string `json:"commits_url"`
	Has_pages bool `json:"has_pages,omitempty"`
	Labels_url string `json:"labels_url"`
	Clone_url string `json:"clone_url,omitempty"`
	Git_commits_url string `json:"git_commits_url"`
	Html_url string `json:"html_url"`
	Contents_url string `json:"contents_url"`
	Custom_properties map[string]interface{} `json:"custom_properties,omitempty"` // The custom properties that were defined for the repository. The keys are the custom property names, and the values are the corresponding custom property values.
	Is_template bool `json:"is_template,omitempty"`
	Issues_url string `json:"issues_url"`
	Trees_url string `json:"trees_url"`
	Role_name string `json:"role_name,omitempty"`
	Notifications_url string `json:"notifications_url"`
	Subscribers_url string `json:"subscribers_url"`
	Contributors_url string `json:"contributors_url"`
	Statuses_url string `json:"statuses_url"`
	Has_wiki bool `json:"has_wiki,omitempty"`
}

// GeneratedType_Webhook_issues_milestoned represents the GeneratedType_Webhook_issues_milestoned schema from the OpenAPI specification
type GeneratedType_Webhook_issues_milestoned struct {
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Milestone Webhooksmilestone `json:"milestone"` // A collection of related issues and pull requests.
}

// GeneratedType_Webhook_discussion_unpinned represents the GeneratedType_Webhook_discussion_unpinned schema from the OpenAPI specification
type GeneratedType_Webhook_discussion_unpinned struct {
	Action string `json:"action"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
}

// GeneratedType_Webhook_discussion_answered represents the GeneratedType_Webhook_discussion_answered schema from the OpenAPI specification
type GeneratedType_Webhook_discussion_answered struct {
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Answer Webhooksanswer `json:"answer"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// Webhooksmarketplacepurchase represents the Webhooksmarketplacepurchase schema from the OpenAPI specification
type Webhooksmarketplacepurchase struct {
	Account map[string]interface{} `json:"account"`
	Billing_cycle string `json:"billing_cycle"`
	Free_trial_ends_on string `json:"free_trial_ends_on"`
	Next_billing_date string `json:"next_billing_date"`
	On_free_trial bool `json:"on_free_trial"`
	Plan map[string]interface{} `json:"plan"`
	Unit_count int `json:"unit_count"`
}

// GeneratedType_Webhook_sub_issues_sub_issue_added represents the GeneratedType_Webhook_sub_issues_sub_issue_added schema from the OpenAPI specification
type GeneratedType_Webhook_sub_issues_sub_issue_added struct {
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Sub_issue_repo Repository `json:"sub_issue_repo"` // A repository on GitHub.
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Sub_issue Issue `json:"sub_issue"` // Issues are a great way to keep track of tasks, enhancements, and bugs for your projects.
	Action string `json:"action"`
	Parent_issue_id float64 `json:"parent_issue_id"` // The ID of the parent issue.
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Parent_issue Issue `json:"parent_issue"` // Issues are a great way to keep track of tasks, enhancements, and bugs for your projects.
	Sub_issue_id float64 `json:"sub_issue_id"` // The ID of the sub-issue.
}

// GeneratedType_Webhook_deployment_status_created represents the GeneratedType_Webhook_deployment_status_created schema from the OpenAPI specification
type GeneratedType_Webhook_deployment_status_created struct {
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Deployment_status map[string]interface{} `json:"deployment_status"` // The [deployment status](https://docs.github.com/rest/deployments/statuses#list-deployment-statuses).
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Workflow Webhooksworkflow `json:"workflow,omitempty"`
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Workflow_run map[string]interface{} `json:"workflow_run,omitempty"`
	Action string `json:"action"`
	Deployment map[string]interface{} `json:"deployment"` // The [deployment](https://docs.github.com/rest/deployments/deployments#list-deployments).
	Check_run map[string]interface{} `json:"check_run,omitempty"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType_Webhook_repository_vulnerability_alert_reopen represents the GeneratedType_Webhook_repository_vulnerability_alert_reopen schema from the OpenAPI specification
type GeneratedType_Webhook_repository_vulnerability_alert_reopen struct {
	Action string `json:"action"`
	Alert Webhooksalert `json:"alert"` // The security alert of the vulnerable dependency.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
}

// Snapshot represents the Snapshot schema from the OpenAPI specification
type Snapshot struct {
	Detector map[string]interface{} `json:"detector"` // A description of the detector used.
	Job map[string]interface{} `json:"job"`
	Manifests map[string]interface{} `json:"manifests,omitempty"` // A collection of package manifests, which are a collection of related dependencies declared in a file or representing a logical group of dependencies.
	Metadata Metadata `json:"metadata,omitempty"` // User-defined metadata to store domain-specific information limited to 8 keys with scalar values.
	Ref string `json:"ref"` // The repository branch that triggered this snapshot.
	Scanned string `json:"scanned"` // The time at which the snapshot was scanned.
	Sha string `json:"sha"` // The commit SHA associated with this dependency snapshot. Maximum length: 40 characters.
	Version int `json:"version"` // The version of the repository snapshot submission.
}

// GeneratedType_Personal_access_token_request represents the GeneratedType_Personal_access_token_request schema from the OpenAPI specification
type GeneratedType_Personal_access_token_request struct {
	Created_at string `json:"created_at"` // Date and time when the request for access was created.
	Owner GeneratedType_Simple_user `json:"owner"` // A GitHub user.
	Repository_selection string `json:"repository_selection"` // Type of repository selection requested.
	Token_expires_at string `json:"token_expires_at"` // Date and time when the associated fine-grained personal access token expires.
	Permissions_added map[string]interface{} `json:"permissions_added"` // New requested permissions, categorized by type of permission.
	Permissions_result map[string]interface{} `json:"permissions_result"` // Permissions requested, categorized by type of permission. This field incorporates `permissions_added` and `permissions_upgraded`.
	Token_name string `json:"token_name"` // The name given to the user's token. This field can also be found in an organization's settings page for Active Tokens.
	Id int `json:"id"` // Unique identifier of the request for access via fine-grained personal access token. Used as the `pat_request_id` parameter in the list and review API calls.
	Repositories []map[string]interface{} `json:"repositories"` // An array of repository objects the token is requesting access to. This field is only populated when `repository_selection` is `subset`.
	Repository_count int `json:"repository_count"` // The number of repositories the token is requesting access to. This field is only populated when `repository_selection` is `subset`.
	Token_last_used_at string `json:"token_last_used_at"` // Date and time when the associated fine-grained personal access token was last used for authentication.
	Permissions_upgraded map[string]interface{} `json:"permissions_upgraded"` // Requested permissions that elevate access for a previously approved request for access, categorized by type of permission.
	Token_expired bool `json:"token_expired"` // Whether the associated fine-grained personal access token has expired.
	Token_id int `json:"token_id"` // Unique identifier of the user's token. This field can also be found in audit log events and the organization's settings for their PAT grants.
}

// GeneratedType_Webhook_issues_closed represents the GeneratedType_Webhook_issues_closed schema from the OpenAPI specification
type GeneratedType_Webhook_issues_closed struct {
	Action string `json:"action"` // The action that was performed.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
}

// GeneratedType_Webhook_secret_scanning_scan_completed represents the GeneratedType_Webhook_secret_scanning_scan_completed schema from the OpenAPI specification
type GeneratedType_Webhook_secret_scanning_scan_completed struct {
	Custom_pattern_scope string `json:"custom_pattern_scope,omitempty"` // If the scan was triggered by a custom pattern update, this will be the scope of the pattern that was updated
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Started_at string `json:"started_at"` // The time that the alert was resolved in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Completed_at string `json:"completed_at"` // The time that the alert was resolved in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Custom_pattern_name string `json:"custom_pattern_name,omitempty"` // If the scan was triggered by a custom pattern update, this will be the name of the pattern that was updated
	Secret_types []string `json:"secret_types,omitempty"` // List of patterns that were updated. This will be empty for normal backfill scans or custom pattern updates
	Source string `json:"source"` // What type of content was scanned
	TypeField string `json:"type"` // What type of scan was completed
	Action string `json:"action"`
}

// GeneratedType_Webhook_release_prereleased represents the GeneratedType_Webhook_release_prereleased schema from the OpenAPI specification
type GeneratedType_Webhook_release_prereleased struct {
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Release map[string]interface{} `json:"release"` // The [release](https://docs.github.com/rest/releases/releases/#get-a-release) object.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType_Nullable_repository_webhooks represents the GeneratedType_Nullable_repository_webhooks schema from the OpenAPI specification
type GeneratedType_Nullable_repository_webhooks struct {
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Blobs_url string `json:"blobs_url"`
	Teams_url string `json:"teams_url"`
	Mirror_url string `json:"mirror_url"`
	Comments_url string `json:"comments_url"`
	Branches_url string `json:"branches_url"`
	Issue_events_url string `json:"issue_events_url"`
	Open_issues int `json:"open_issues"`
	Stargazers_count int `json:"stargazers_count"`
	Has_discussions bool `json:"has_discussions,omitempty"` // Whether discussions are enabled.
	Fork bool `json:"fork"`
	Organization GeneratedType_Nullable_simple_user `json:"organization,omitempty"` // A GitHub user.
	Compare_url string `json:"compare_url"`
	Releases_url string `json:"releases_url"`
	Temp_clone_token string `json:"temp_clone_token,omitempty"`
	Git_refs_url string `json:"git_refs_url"`
	Contributors_url string `json:"contributors_url"`
	Deployments_url string `json:"deployments_url"`
	Merge_commit_title string `json:"merge_commit_title,omitempty"` // The default value for a merge commit title. - `PR_TITLE` - default to the pull request's title. - `MERGE_MESSAGE` - default to the classic title for a merge message (e.g., Merge pull request #123 from branch-name).
	Issue_comment_url string `json:"issue_comment_url"`
	Visibility string `json:"visibility,omitempty"` // The repository visibility: public, private, or internal.
	Archived bool `json:"archived"` // Whether the repository is archived.
	Has_downloads bool `json:"has_downloads"` // Whether downloads are enabled.
	Owner GeneratedType_Simple_user `json:"owner"` // A GitHub user.
	Collaborators_url string `json:"collaborators_url"`
	Issues_url string `json:"issues_url"`
	Watchers_count int `json:"watchers_count"`
	Name string `json:"name"` // The name of the repository.
	Subscribers_url string `json:"subscribers_url"`
	Tags_url string `json:"tags_url"`
	Clone_url string `json:"clone_url"`
	Notifications_url string `json:"notifications_url"`
	Has_issues bool `json:"has_issues"` // Whether issues are enabled.
	Template_repository map[string]interface{} `json:"template_repository,omitempty"`
	Subscribers_count int `json:"subscribers_count,omitempty"`
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"` // Whether to require contributors to sign off on web-based commits
	Commits_url string `json:"commits_url"`
	Keys_url string `json:"keys_url"`
	Language string `json:"language"`
	Id int64 `json:"id"` // Unique identifier of the repository
	Full_name string `json:"full_name"`
	Git_commits_url string `json:"git_commits_url"`
	Anonymous_access_enabled bool `json:"anonymous_access_enabled,omitempty"` // Whether anonymous git access is enabled for this repository
	Milestones_url string `json:"milestones_url"`
	Forks int `json:"forks"`
	Custom_properties map[string]interface{} `json:"custom_properties,omitempty"` // The custom properties that were defined for the repository. The keys are the custom property names, and the values are the corresponding custom property values.
	Languages_url string `json:"languages_url"`
	Forks_url string `json:"forks_url"`
	Git_url string `json:"git_url"`
	Use_squash_pr_title_as_default bool `json:"use_squash_pr_title_as_default,omitempty"` // Whether a squash merge commit can use the pull request title as default. **This property is closing down. Please use `squash_merge_commit_title` instead.
	Allow_auto_merge bool `json:"allow_auto_merge,omitempty"` // Whether to allow Auto-merge to be used on pull requests.
	Pushed_at string `json:"pushed_at"`
	Trees_url string `json:"trees_url"`
	Open_issues_count int `json:"open_issues_count"`
	Created_at string `json:"created_at"`
	Squash_merge_commit_title string `json:"squash_merge_commit_title,omitempty"` // The default value for a squash merge commit title: - `PR_TITLE` - default to the pull request's title. - `COMMIT_OR_PR_TITLE` - default to the commit's title (if only one commit) or the pull request's title (when more than one commit).
	Hooks_url string `json:"hooks_url"`
	Has_wiki bool `json:"has_wiki"` // Whether the wiki is enabled.
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	Node_id string `json:"node_id"`
	Statuses_url string `json:"statuses_url"`
	Labels_url string `json:"labels_url"`
	Allow_merge_commit bool `json:"allow_merge_commit,omitempty"` // Whether to allow merge commits for pull requests.
	Topics []string `json:"topics,omitempty"`
	Watchers int `json:"watchers"`
	Size int `json:"size"` // The size of the repository, in kilobytes. Size is calculated hourly. When a repository is initially created, the size is 0.
	Subscription_url string `json:"subscription_url"`
	Allow_forking bool `json:"allow_forking,omitempty"` // Whether to allow forking this repo
	Allow_update_branch bool `json:"allow_update_branch,omitempty"` // Whether or not a pull request head branch that is behind its base branch can always be updated even if it is not required to be up to date before merging.
	Archive_url string `json:"archive_url"`
	Disabled bool `json:"disabled"` // Returns whether or not this repository disabled.
	Starred_at string `json:"starred_at,omitempty"`
	Has_projects bool `json:"has_projects"` // Whether projects are enabled.
	Network_count int `json:"network_count,omitempty"`
	Squash_merge_commit_message string `json:"squash_merge_commit_message,omitempty"` // The default value for a squash merge commit message: - `PR_BODY` - default to the pull request's body. - `COMMIT_MESSAGES` - default to the branch's commit messages. - `BLANK` - default to a blank commit message.
	Homepage string `json:"homepage"`
	Is_template bool `json:"is_template,omitempty"` // Whether this repository acts as a template that can be used to generate new repositories.
	Stargazers_url string `json:"stargazers_url"`
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"` // Whether to delete head branches when pull requests are merged
	Git_tags_url string `json:"git_tags_url"`
	Default_branch string `json:"default_branch"` // The default branch of the repository.
	Has_pages bool `json:"has_pages"`
	License GeneratedType_Nullable_license_simple `json:"license"` // License Simple
	Merges_url string `json:"merges_url"`
	Master_branch string `json:"master_branch,omitempty"`
	Html_url string `json:"html_url"`
	Downloads_url string `json:"downloads_url"`
	Allow_squash_merge bool `json:"allow_squash_merge,omitempty"` // Whether to allow squash merges for pull requests.
	Contents_url string `json:"contents_url"`
	Description string `json:"description"`
	Ssh_url string `json:"ssh_url"`
	Assignees_url string `json:"assignees_url"`
	Merge_commit_message string `json:"merge_commit_message,omitempty"` // The default value for a merge commit message. - `PR_TITLE` - default to the pull request's title. - `PR_BODY` - default to the pull request's body. - `BLANK` - default to a blank commit message.
	Svn_url string `json:"svn_url"`
	Pulls_url string `json:"pulls_url"`
	Allow_rebase_merge bool `json:"allow_rebase_merge,omitempty"` // Whether to allow rebase merges for pull requests.
	Private bool `json:"private"` // Whether the repository is private or public.
	Events_url string `json:"events_url"`
	Forks_count int `json:"forks_count"`
}

// GeneratedType_State_change_issue_event represents the GeneratedType_State_change_issue_event schema from the OpenAPI specification
type GeneratedType_State_change_issue_event struct {
	Commit_id string `json:"commit_id"`
	Created_at string `json:"created_at"`
	Event string `json:"event"`
	Id int `json:"id"`
	Url string `json:"url"`
	Performed_via_github_app GeneratedType_Nullable_integration `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Actor GeneratedType_Simple_user `json:"actor"` // A GitHub user.
	Commit_url string `json:"commit_url"`
	Node_id string `json:"node_id"`
	State_reason string `json:"state_reason,omitempty"`
}

// GeneratedType_Webhook_deployment_protection_rule_requested represents the GeneratedType_Webhook_deployment_protection_rule_requested schema from the OpenAPI specification
type GeneratedType_Webhook_deployment_protection_rule_requested struct {
	Environment string `json:"environment,omitempty"` // The name of the environment that has the deployment protection rule.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Deployment Deployment `json:"deployment,omitempty"` // A request for a specific ref(branch,sha,tag) to be deployed
	Deployment_callback_url string `json:"deployment_callback_url,omitempty"` // The URL to review the deployment protection rule.
	Event string `json:"event,omitempty"` // The event that triggered the deployment protection rule.
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action,omitempty"`
	Pull_requests []GeneratedType_Pull_request `json:"pull_requests,omitempty"`
}

// GeneratedType_Webhook_project_column_created represents the GeneratedType_Webhook_project_column_created schema from the OpenAPI specification
type GeneratedType_Webhook_project_column_created struct {
	Project_column Webhooksprojectcolumn `json:"project_column"`
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// Webhooksprojectcard represents the Webhooksprojectcard schema from the OpenAPI specification
type Webhooksprojectcard struct {
	Updated_at string `json:"updated_at"`
	Node_id string `json:"node_id"`
	Url string `json:"url"`
	Column_url string `json:"column_url"`
	Creator map[string]interface{} `json:"creator"`
	Id int `json:"id"` // The project card's ID
	Note string `json:"note"`
	Project_url string `json:"project_url"`
	Archived bool `json:"archived"` // Whether or not the card is archived
	Content_url string `json:"content_url,omitempty"`
	Created_at string `json:"created_at"`
	After_id int `json:"after_id,omitempty"`
	Column_id int `json:"column_id"`
}

// GeneratedType_Code_scanning_variant_analysis_repository represents the GeneratedType_Code_scanning_variant_analysis_repository schema from the OpenAPI specification
type GeneratedType_Code_scanning_variant_analysis_repository struct {
	Stargazers_count int `json:"stargazers_count"`
	Updated_at string `json:"updated_at"`
	Full_name string `json:"full_name"` // The full, globally unique, name of the repository.
	Id int `json:"id"` // A unique identifier of the repository.
	Name string `json:"name"` // The name of the repository.
	Private bool `json:"private"` // Whether the repository is private.
}

// GeneratedType_Code_scanning_default_setup_update_response represents the GeneratedType_Code_scanning_default_setup_update_response schema from the OpenAPI specification
type GeneratedType_Code_scanning_default_setup_update_response struct {
	Run_id int `json:"run_id,omitempty"` // ID of the corresponding run.
	Run_url string `json:"run_url,omitempty"` // URL of the corresponding run.
}

// GeneratedType_Webhook_project_column_deleted represents the GeneratedType_Webhook_project_column_deleted schema from the OpenAPI specification
type GeneratedType_Webhook_project_column_deleted struct {
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Project_column Webhooksprojectcolumn `json:"project_column"`
	Repository GeneratedType_Nullable_repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
}

// GeneratedType_Actions_hosted_runner represents the GeneratedType_Actions_hosted_runner schema from the OpenAPI specification
type GeneratedType_Actions_hosted_runner struct {
	Last_active_on string `json:"last_active_on,omitempty"` // The time at which the runner was last used, in ISO 8601 format.
	Public_ip_enabled bool `json:"public_ip_enabled"` // Whether public IP is enabled for the hosted runners.
	Platform string `json:"platform"` // The operating system of the image.
	Runner_group_id int `json:"runner_group_id,omitempty"` // The unique identifier of the group that the hosted runner belongs to.
	Id int `json:"id"` // The unique identifier of the hosted runner.
	Maximum_runners int `json:"maximum_runners,omitempty"` // The maximum amount of hosted runners. Runners will not scale automatically above this number. Use this setting to limit your cost.
	Name string `json:"name"` // The name of the hosted runner.
	Public_ips []GeneratedType_Public_ip `json:"public_ips,omitempty"` // The public IP ranges when public IP is enabled for the hosted runners.
	Machine_size_details GeneratedType_Actions_hosted_runner_machine_spec `json:"machine_size_details"` // Provides details of a particular machine spec.
	Status string `json:"status"` // The status of the runner.
	Image_details GeneratedType_Nullable_actions_hosted_runner_pool_image `json:"image_details"` // Provides details of a hosted runner image
}

// Link represents the Link schema from the OpenAPI specification
type Link struct {
	Href string `json:"href"`
}

// GeneratedType_Webhook_secret_scanning_alert_location_created_form_encoded represents the GeneratedType_Webhook_secret_scanning_alert_location_created_form_encoded schema from the OpenAPI specification
type GeneratedType_Webhook_secret_scanning_alert_location_created_form_encoded struct {
	Payload string `json:"payload"` // A URL-encoded string of the secret_scanning_alert_location.created JSON payload. The decoded payload is a JSON object.
}

// GeneratedType_Referenced_workflow represents the GeneratedType_Referenced_workflow schema from the OpenAPI specification
type GeneratedType_Referenced_workflow struct {
	Sha string `json:"sha"`
	Path string `json:"path"`
	Ref string `json:"ref,omitempty"`
}

// GeneratedType_Webhook_check_run_completed represents the GeneratedType_Webhook_check_run_completed schema from the OpenAPI specification
type GeneratedType_Webhook_check_run_completed struct {
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action,omitempty"`
	Check_run GeneratedType_Check_run_with_simple_check_suite `json:"check_run"` // A check performed on the code of a given code change
}

// GeneratedType_Code_scanning_sarifs_status represents the GeneratedType_Code_scanning_sarifs_status schema from the OpenAPI specification
type GeneratedType_Code_scanning_sarifs_status struct {
	Analyses_url string `json:"analyses_url,omitempty"` // The REST API URL for getting the analyses associated with the upload.
	Errors []string `json:"errors,omitempty"` // Any errors that ocurred during processing of the delivery.
	Processing_status string `json:"processing_status,omitempty"` // `pending` files have not yet been processed, while `complete` means results from the SARIF have been stored. `failed` files have either not been processed at all, or could only be partially processed.
}

// GeneratedType_Webhook_projects_v2_item_reordered represents the GeneratedType_Webhook_projects_v2_item_reordered schema from the OpenAPI specification
type GeneratedType_Webhook_projects_v2_item_reordered struct {
	Changes map[string]interface{} `json:"changes"`
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Projects_v2_item GeneratedType_Projects_v2_item `json:"projects_v2_item"` // An item belonging to a project
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType_Pull_request_minimal represents the GeneratedType_Pull_request_minimal schema from the OpenAPI specification
type GeneratedType_Pull_request_minimal struct {
	Id int64 `json:"id"`
	Number int `json:"number"`
	Url string `json:"url"`
	Base map[string]interface{} `json:"base"`
	Head map[string]interface{} `json:"head"`
}

// GeneratedType_Check_run_with_simple_check_suite represents the GeneratedType_Check_run_with_simple_check_suite schema from the OpenAPI specification
type GeneratedType_Check_run_with_simple_check_suite struct {
	Completed_at string `json:"completed_at"`
	Deployment GeneratedType_Deployment_simple `json:"deployment,omitempty"` // A deployment created as the result of an Actions check run from a workflow that references an environment
	Head_sha string `json:"head_sha"` // The SHA of the commit that is being checked.
	Started_at string `json:"started_at"`
	Node_id string `json:"node_id"`
	Output map[string]interface{} `json:"output"`
	Conclusion string `json:"conclusion"`
	Html_url string `json:"html_url"`
	Status string `json:"status"` // The phase of the lifecycle that the check is currently in.
	Url string `json:"url"`
	Id int `json:"id"` // The id of the check.
	Name string `json:"name"` // The name of the check.
	Details_url string `json:"details_url"`
	App Integration `json:"app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	External_id string `json:"external_id"`
	Check_suite GeneratedType_Simple_check_suite `json:"check_suite"` // A suite of checks performed on the code of a given code change
	Pull_requests []GeneratedType_Pull_request_minimal `json:"pull_requests"`
}

// GeneratedType_Webhook_pull_request_edited represents the GeneratedType_Webhook_pull_request_edited schema from the OpenAPI specification
type GeneratedType_Webhook_pull_request_edited struct {
	Action string `json:"action"`
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Number int `json:"number"` // The pull request number.
	Pull_request GeneratedType_Pull_request_webhook `json:"pull_request"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Changes map[string]interface{} `json:"changes"` // The changes to the comment if the action was `edited`.
}

// GeneratedType_Git_ref represents the GeneratedType_Git_ref schema from the OpenAPI specification
type GeneratedType_Git_ref struct {
	Node_id string `json:"node_id"`
	Object map[string]interface{} `json:"object"`
	Ref string `json:"ref"`
	Url string `json:"url"`
}

// GeneratedType_Codespaces_public_key represents the GeneratedType_Codespaces_public_key schema from the OpenAPI specification
type GeneratedType_Codespaces_public_key struct {
	Title string `json:"title,omitempty"`
	Url string `json:"url,omitempty"`
	Created_at string `json:"created_at,omitempty"`
	Id int `json:"id,omitempty"`
	Key string `json:"key"` // The Base64 encoded public key.
	Key_id string `json:"key_id"` // The identifier for the key.
}

// Webhooksrule represents the Webhooksrule schema from the OpenAPI specification
type Webhooksrule struct {
	Authorized_actors_only bool `json:"authorized_actors_only"`
	Lock_allows_fork_sync bool `json:"lock_allows_fork_sync,omitempty"` // Whether users can pull changes from upstream when the branch is locked. Set to `true` to allow users to pull changes from upstream when the branch is locked. This setting is only applicable for forks.
	Required_status_checks_enforcement_level string `json:"required_status_checks_enforcement_level"`
	Signature_requirement_enforcement_level string `json:"signature_requirement_enforcement_level"`
	Allow_deletions_enforcement_level string `json:"allow_deletions_enforcement_level"`
	Name string `json:"name"`
	Admin_enforced bool `json:"admin_enforced"`
	Authorized_dismissal_actors_only bool `json:"authorized_dismissal_actors_only"`
	Merge_queue_enforcement_level string `json:"merge_queue_enforcement_level"`
	Lock_branch_enforcement_level string `json:"lock_branch_enforcement_level"` // The enforcement level of the branch lock setting. `off` means the branch is not locked, `non_admins` means the branch is read-only for non_admins, and `everyone` means the branch is read-only for everyone.
	Ignore_approvals_from_contributors bool `json:"ignore_approvals_from_contributors"`
	Strict_required_status_checks_policy bool `json:"strict_required_status_checks_policy"`
	Updated_at string `json:"updated_at"`
	Required_approving_review_count int `json:"required_approving_review_count"`
	Allow_force_pushes_enforcement_level string `json:"allow_force_pushes_enforcement_level"`
	Id int `json:"id"`
	Require_code_owner_review bool `json:"require_code_owner_review"`
	Required_status_checks []string `json:"required_status_checks"`
	Require_last_push_approval bool `json:"require_last_push_approval,omitempty"` // Whether the most recent push must be approved by someone other than the person who pushed it
	Required_deployments_enforcement_level string `json:"required_deployments_enforcement_level"`
	Create_protected bool `json:"create_protected,omitempty"`
	Created_at string `json:"created_at"`
	Linear_history_requirement_enforcement_level string `json:"linear_history_requirement_enforcement_level"`
	Pull_request_reviews_enforcement_level string `json:"pull_request_reviews_enforcement_level"`
	Required_conversation_resolution_level string `json:"required_conversation_resolution_level"`
	Authorized_actor_names []string `json:"authorized_actor_names"`
	Dismiss_stale_reviews_on_push bool `json:"dismiss_stale_reviews_on_push"`
	Repository_id int `json:"repository_id"`
}

// GeneratedType_Public_ip represents the GeneratedType_Public_ip schema from the OpenAPI specification
type GeneratedType_Public_ip struct {
	Enabled bool `json:"enabled,omitempty"` // Whether public IP is enabled.
	Length int `json:"length,omitempty"` // The length of the IP prefix.
	Prefix string `json:"prefix,omitempty"` // The prefix for the public IP.
}

// GeneratedType_Webhook_projects_v2_item_created represents the GeneratedType_Webhook_projects_v2_item_created schema from the OpenAPI specification
type GeneratedType_Webhook_projects_v2_item_created struct {
	Projects_v2_item GeneratedType_Projects_v2_item `json:"projects_v2_item"` // An item belonging to a project
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// Team represents the Team schema from the OpenAPI specification
type Team struct {
	Node_id string `json:"node_id"`
	Slug string `json:"slug"`
	Url string `json:"url"`
	Description string `json:"description"`
	Html_url string `json:"html_url"`
	Members_url string `json:"members_url"`
	Name string `json:"name"`
	Privacy string `json:"privacy,omitempty"`
	Parent GeneratedType_Nullable_team_simple `json:"parent"` // Groups of organization members that gives permissions on specified repositories.
	Permission string `json:"permission"`
	Repositories_url string `json:"repositories_url"`
	Notification_setting string `json:"notification_setting,omitempty"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Id int `json:"id"`
}

// GeneratedType_Webhook_projects_v2_project_deleted represents the GeneratedType_Webhook_projects_v2_project_deleted schema from the OpenAPI specification
type GeneratedType_Webhook_projects_v2_project_deleted struct {
	Projects_v2 GeneratedType_Projects_v2 `json:"projects_v2"` // A projects v2 project
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType_Webhook_workflow_dispatch represents the GeneratedType_Webhook_workflow_dispatch schema from the OpenAPI specification
type GeneratedType_Webhook_workflow_dispatch struct {
	Inputs map[string]interface{} `json:"inputs"`
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Ref string `json:"ref"`
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Workflow string `json:"workflow"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType_Webhook_issues_demilestoned represents the GeneratedType_Webhook_issues_demilestoned schema from the OpenAPI specification
type GeneratedType_Webhook_issues_demilestoned struct {
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Milestone Webhooksmilestone `json:"milestone,omitempty"` // A collection of related issues and pull requests.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
}

// GeneratedType_Integration_installation_request represents the GeneratedType_Integration_installation_request schema from the OpenAPI specification
type GeneratedType_Integration_installation_request struct {
	Created_at string `json:"created_at"`
	Id int `json:"id"` // Unique identifier of the request installation.
	Node_id string `json:"node_id,omitempty"`
	Requester GeneratedType_Simple_user `json:"requester"` // A GitHub user.
	Account interface{} `json:"account"`
}

// GeneratedType_Rate_limit_overview represents the GeneratedType_Rate_limit_overview schema from the OpenAPI specification
type GeneratedType_Rate_limit_overview struct {
	Rate GeneratedType_Rate_limit `json:"rate"`
	Resources map[string]interface{} `json:"resources"`
}

// GeneratedType_Webhook_dependabot_alert_reopened represents the GeneratedType_Webhook_dependabot_alert_reopened schema from the OpenAPI specification
type GeneratedType_Webhook_dependabot_alert_reopened struct {
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Alert GeneratedType_Dependabot_alert `json:"alert"` // A Dependabot alert.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType_Status_check_policy represents the GeneratedType_Status_check_policy schema from the OpenAPI specification
type GeneratedType_Status_check_policy struct {
	Strict bool `json:"strict"`
	Url string `json:"url"`
	Checks []map[string]interface{} `json:"checks"`
	Contexts []string `json:"contexts"`
	Contexts_url string `json:"contexts_url"`
}

// GeneratedType_Validation_error represents the GeneratedType_Validation_error schema from the OpenAPI specification
type GeneratedType_Validation_error struct {
	Documentation_url string `json:"documentation_url"`
	Errors []map[string]interface{} `json:"errors,omitempty"`
	Message string `json:"message"`
}

// GeneratedType_Webhook_deployment_review_approved represents the GeneratedType_Webhook_deployment_review_approved schema from the OpenAPI specification
type GeneratedType_Webhook_deployment_review_approved struct {
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Reviewers []map[string]interface{} `json:"reviewers,omitempty"`
	Since string `json:"since"`
	Workflow_run map[string]interface{} `json:"workflow_run"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Workflow_job_run Webhooksworkflowjobrun `json:"workflow_job_run,omitempty"`
	Action string `json:"action"`
	Approver Webhooksapprover `json:"approver,omitempty"`
	Comment string `json:"comment,omitempty"`
	Workflow_job_runs []map[string]interface{} `json:"workflow_job_runs,omitempty"`
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
}

// GeneratedType_Empty_object represents the GeneratedType_Empty_object schema from the OpenAPI specification
type GeneratedType_Empty_object struct {
}

// GeneratedType_Git_tag represents the GeneratedType_Git_tag schema from the OpenAPI specification
type GeneratedType_Git_tag struct {
	Tagger map[string]interface{} `json:"tagger"`
	Url string `json:"url"` // URL for the tag
	Verification Verification `json:"verification,omitempty"`
	Message string `json:"message"` // Message describing the purpose of the tag
	Node_id string `json:"node_id"`
	Object map[string]interface{} `json:"object"`
	Sha string `json:"sha"`
	Tag string `json:"tag"` // Name of the tag
}

// GeneratedType_Webhook_personal_access_token_request_cancelled represents the GeneratedType_Webhook_personal_access_token_request_cancelled schema from the OpenAPI specification
type GeneratedType_Webhook_personal_access_token_request_cancelled struct {
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Personal_access_token_request GeneratedType_Personal_access_token_request `json:"personal_access_token_request"` // Details of a Personal Access Token Request.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType_Webhook_secret_scanning_alert_resolved represents the GeneratedType_Webhook_secret_scanning_alert_resolved schema from the OpenAPI specification
type GeneratedType_Webhook_secret_scanning_alert_resolved struct {
	Action string `json:"action"`
	Alert GeneratedType_Secret_scanning_alert_webhook `json:"alert"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
}

// GeneratedType_Repository_rule_update represents the GeneratedType_Repository_rule_update schema from the OpenAPI specification
type GeneratedType_Repository_rule_update struct {
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	TypeField string `json:"type"`
}

// Webhookschanges8 represents the Webhookschanges8 schema from the OpenAPI specification
type Webhookschanges8 struct {
	Tier map[string]interface{} `json:"tier"`
}

// GeneratedType_Code_scanning_sarifs_receipt represents the GeneratedType_Code_scanning_sarifs_receipt schema from the OpenAPI specification
type GeneratedType_Code_scanning_sarifs_receipt struct {
	Id string `json:"id,omitempty"` // An identifier for the upload.
	Url string `json:"url,omitempty"` // The REST API URL for checking the status of the upload.
}

// GeneratedType_Dependency_graph_spdx_sbom represents the GeneratedType_Dependency_graph_spdx_sbom schema from the OpenAPI specification
type GeneratedType_Dependency_graph_spdx_sbom struct {
	Sbom map[string]interface{} `json:"sbom"`
}

// GeneratedType_Webhook_pull_request_enqueued represents the GeneratedType_Webhook_pull_request_enqueued schema from the OpenAPI specification
type GeneratedType_Webhook_pull_request_enqueued struct {
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Number int `json:"number"`
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType_Webhook_projects_v2_project_edited represents the GeneratedType_Webhook_projects_v2_project_edited schema from the OpenAPI specification
type GeneratedType_Webhook_projects_v2_project_edited struct {
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Projects_v2 GeneratedType_Projects_v2 `json:"projects_v2"` // A projects v2 project
}

// GeneratedType_Webhook_workflow_run_in_progress represents the GeneratedType_Webhook_workflow_run_in_progress schema from the OpenAPI specification
type GeneratedType_Webhook_workflow_run_in_progress struct {
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Workflow Webhooksworkflow `json:"workflow"`
	Workflow_run map[string]interface{} `json:"workflow_run"`
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType_Org_private_registry_configuration represents the GeneratedType_Org_private_registry_configuration schema from the OpenAPI specification
type GeneratedType_Org_private_registry_configuration struct {
	Updated_at string `json:"updated_at"`
	Username string `json:"username,omitempty"` // The username to use when authenticating with the private registry.
	Visibility string `json:"visibility"` // Which type of organization repositories have access to the private registry.
	Created_at string `json:"created_at"`
	Name string `json:"name"` // The name of the private registry configuration.
	Registry_type string `json:"registry_type"` // The registry type.
}

// GeneratedType_Code_scanning_alert represents the GeneratedType_Code_scanning_alert schema from the OpenAPI specification
type GeneratedType_Code_scanning_alert struct {
	Url string `json:"url"` // The REST API URL of the alert resource.
	Number int `json:"number"` // The security alert number.
	Dismissed_comment string `json:"dismissed_comment,omitempty"` // The dismissal comment associated with the dismissal of the alert.
	Fixed_at string `json:"fixed_at,omitempty"` // The time that the alert was no longer detected and was considered fixed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Created_at string `json:"created_at"` // The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Most_recent_instance GeneratedType_Code_scanning_alert_instance `json:"most_recent_instance"`
	Rule GeneratedType_Code_scanning_alert_rule `json:"rule"`
	Instances_url string `json:"instances_url"` // The REST API URL for fetching the list of instances for an alert.
	Tool GeneratedType_Code_scanning_analysis_tool `json:"tool"`
	Updated_at string `json:"updated_at,omitempty"` // The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Dismissed_at string `json:"dismissed_at"` // The time that the alert was dismissed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Dismissed_by GeneratedType_Nullable_simple_user `json:"dismissed_by"` // A GitHub user.
	State string `json:"state"` // State of a code scanning alert.
	Dismissal_approved_by GeneratedType_Nullable_simple_user `json:"dismissal_approved_by,omitempty"` // A GitHub user.
	Html_url string `json:"html_url"` // The GitHub URL of the alert resource.
	Dismissed_reason string `json:"dismissed_reason"` // **Required when the state is dismissed.** The reason for dismissing or closing the alert.
}

// GeneratedType_Secret_scanning_location_issue_body represents the GeneratedType_Secret_scanning_location_issue_body schema from the OpenAPI specification
type GeneratedType_Secret_scanning_location_issue_body struct {
	Issue_body_url string `json:"issue_body_url"` // The API URL to get the issue where the secret was detected.
}

// GeneratedType_Issue_type represents the GeneratedType_Issue_type schema from the OpenAPI specification
type GeneratedType_Issue_type struct {
	Is_enabled bool `json:"is_enabled,omitempty"` // The enabled state of the issue type.
	Name string `json:"name"` // The name of the issue type.
	Node_id string `json:"node_id"` // The node identifier of the issue type.
	Updated_at string `json:"updated_at,omitempty"` // The time the issue type last updated.
	Color string `json:"color,omitempty"` // The color of the issue type.
	Created_at string `json:"created_at,omitempty"` // The time the issue type created.
	Description string `json:"description"` // The description of the issue type.
	Id int `json:"id"` // The unique identifier of the issue type.
}

// GeneratedType_Public_user represents the GeneratedType_Public_user schema from the OpenAPI specification
type GeneratedType_Public_user struct {
	Public_repos int `json:"public_repos"`
	Following_url string `json:"following_url"`
	Bio string `json:"bio"`
	Created_at string `json:"created_at"`
	Owned_private_repos int `json:"owned_private_repos,omitempty"`
	Followers_url string `json:"followers_url"`
	Id int64 `json:"id"`
	Public_gists int `json:"public_gists"`
	Site_admin bool `json:"site_admin"`
	Events_url string `json:"events_url"`
	Plan map[string]interface{} `json:"plan,omitempty"`
	Avatar_url string `json:"avatar_url"`
	Disk_usage int `json:"disk_usage,omitempty"`
	Repos_url string `json:"repos_url"`
	TypeField string `json:"type"`
	Email string `json:"email"`
	Company string `json:"company"`
	Twitter_username string `json:"twitter_username,omitempty"`
	Html_url string `json:"html_url"`
	Name string `json:"name"`
	Blog string `json:"blog"`
	Hireable bool `json:"hireable"`
	Notification_email string `json:"notification_email,omitempty"`
	Url string `json:"url"`
	User_view_type string `json:"user_view_type,omitempty"`
	Gists_url string `json:"gists_url"`
	Organizations_url string `json:"organizations_url"`
	Following int `json:"following"`
	Location string `json:"location"`
	Total_private_repos int `json:"total_private_repos,omitempty"`
	Login string `json:"login"`
	Updated_at string `json:"updated_at"`
	Starred_url string `json:"starred_url"`
	Followers int `json:"followers"`
	Collaborators int `json:"collaborators,omitempty"`
	Node_id string `json:"node_id"`
	Private_gists int `json:"private_gists,omitempty"`
	Gravatar_id string `json:"gravatar_id"`
	Received_events_url string `json:"received_events_url"`
	Subscriptions_url string `json:"subscriptions_url"`
}

// GeneratedType_Webhook_discussion_unlabeled represents the GeneratedType_Webhook_discussion_unlabeled schema from the OpenAPI specification
type GeneratedType_Webhook_discussion_unlabeled struct {
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Label Webhookslabel `json:"label"`
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType_User_search_result_item represents the GeneratedType_User_search_result_item schema from the OpenAPI specification
type GeneratedType_User_search_result_item struct {
	Url string `json:"url"`
	Starred_url string `json:"starred_url"`
	Bio string `json:"bio,omitempty"`
	Following_url string `json:"following_url"`
	Received_events_url string `json:"received_events_url"`
	Suspended_at string `json:"suspended_at,omitempty"`
	Organizations_url string `json:"organizations_url"`
	Score float64 `json:"score"`
	Blog string `json:"blog,omitempty"`
	Text_matches []map[string]interface{} `json:"text_matches,omitempty"`
	Following int `json:"following,omitempty"`
	Followers_url string `json:"followers_url"`
	Repos_url string `json:"repos_url"`
	Gravatar_id string `json:"gravatar_id"`
	Company string `json:"company,omitempty"`
	Hireable bool `json:"hireable,omitempty"`
	Location string `json:"location,omitempty"`
	Events_url string `json:"events_url"`
	User_view_type string `json:"user_view_type,omitempty"`
	Email string `json:"email,omitempty"`
	Site_admin bool `json:"site_admin"`
	Subscriptions_url string `json:"subscriptions_url"`
	Public_gists int `json:"public_gists,omitempty"`
	Public_repos int `json:"public_repos,omitempty"`
	Followers int `json:"followers,omitempty"`
	Updated_at string `json:"updated_at,omitempty"`
	TypeField string `json:"type"`
	Node_id string `json:"node_id"`
	Html_url string `json:"html_url"`
	Login string `json:"login"`
	Avatar_url string `json:"avatar_url"`
	Id int64 `json:"id"`
	Name string `json:"name,omitempty"`
	Created_at string `json:"created_at,omitempty"`
	Gists_url string `json:"gists_url"`
}

// GeneratedType_Review_requested_issue_event represents the GeneratedType_Review_requested_issue_event schema from the OpenAPI specification
type GeneratedType_Review_requested_issue_event struct {
	Event string `json:"event"`
	Review_requester GeneratedType_Simple_user `json:"review_requester"` // A GitHub user.
	Performed_via_github_app GeneratedType_Nullable_integration `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Requested_team Team `json:"requested_team,omitempty"` // Groups of organization members that gives permissions on specified repositories.
	Id int `json:"id"`
	Requested_reviewer GeneratedType_Simple_user `json:"requested_reviewer,omitempty"` // A GitHub user.
	Url string `json:"url"`
	Actor GeneratedType_Simple_user `json:"actor"` // A GitHub user.
	Commit_url string `json:"commit_url"`
	Node_id string `json:"node_id"`
	Commit_id string `json:"commit_id"`
	Created_at string `json:"created_at"`
}

// GeneratedType_Issue_event represents the GeneratedType_Issue_event schema from the OpenAPI specification
type GeneratedType_Issue_event struct {
	Id int64 `json:"id"`
	Issue GeneratedType_Nullable_issue `json:"issue,omitempty"` // Issues are a great way to keep track of tasks, enhancements, and bugs for your projects.
	Lock_reason string `json:"lock_reason,omitempty"`
	Review_requester GeneratedType_Nullable_simple_user `json:"review_requester,omitempty"` // A GitHub user.
	Requested_team Team `json:"requested_team,omitempty"` // Groups of organization members that gives permissions on specified repositories.
	Author_association string `json:"author_association,omitempty"` // How the author is associated with the repository.
	Commit_url string `json:"commit_url"`
	Actor GeneratedType_Nullable_simple_user `json:"actor"` // A GitHub user.
	Node_id string `json:"node_id"`
	Performed_via_github_app GeneratedType_Nullable_integration `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Requested_reviewer GeneratedType_Nullable_simple_user `json:"requested_reviewer,omitempty"` // A GitHub user.
	Assigner GeneratedType_Nullable_simple_user `json:"assigner,omitempty"` // A GitHub user.
	Event string `json:"event"`
	Url string `json:"url"`
	Assignee GeneratedType_Nullable_simple_user `json:"assignee,omitempty"` // A GitHub user.
	Commit_id string `json:"commit_id"`
	Dismissed_review GeneratedType_Issue_event_dismissed_review `json:"dismissed_review,omitempty"`
	Created_at string `json:"created_at"`
	Project_card GeneratedType_Issue_event_project_card `json:"project_card,omitempty"` // Issue Event Project Card
	Milestone GeneratedType_Issue_event_milestone `json:"milestone,omitempty"` // Issue Event Milestone
	Label GeneratedType_Issue_event_label `json:"label,omitempty"` // Issue Event Label
	Rename GeneratedType_Issue_event_rename `json:"rename,omitempty"` // Issue Event Rename
}

// GeneratedType_Webhook_release_edited represents the GeneratedType_Webhook_release_edited schema from the OpenAPI specification
type GeneratedType_Webhook_release_edited struct {
	Changes map[string]interface{} `json:"changes"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Release Webhooksrelease `json:"release"` // The [release](https://docs.github.com/rest/releases/releases/#get-a-release) object.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType_Repository_rule_max_file_path_length represents the GeneratedType_Repository_rule_max_file_path_length schema from the OpenAPI specification
type GeneratedType_Repository_rule_max_file_path_length struct {
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	TypeField string `json:"type"`
}

// GeneratedType_Repository_rule_detailed represents the GeneratedType_Repository_rule_detailed schema from the OpenAPI specification
type GeneratedType_Repository_rule_detailed struct {
}

// GeneratedType_Webhook_check_run_completed_form_encoded represents the GeneratedType_Webhook_check_run_completed_form_encoded schema from the OpenAPI specification
type GeneratedType_Webhook_check_run_completed_form_encoded struct {
	Payload string `json:"payload"` // A URL-encoded string of the check_run.completed JSON payload. The decoded payload is a JSON object.
}

// GeneratedType_Rule_suite represents the GeneratedType_Rule_suite schema from the OpenAPI specification
type GeneratedType_Rule_suite struct {
	Pushed_at string `json:"pushed_at,omitempty"`
	Repository_id int `json:"repository_id,omitempty"` // The ID of the repository associated with the rule evaluation.
	Result string `json:"result,omitempty"` // The result of the rule evaluations for rules with the `active` enforcement status.
	Rule_evaluations []map[string]interface{} `json:"rule_evaluations,omitempty"` // Details on the evaluated rules.
	Actor_id int `json:"actor_id,omitempty"` // The number that identifies the user.
	Id int `json:"id,omitempty"` // The unique identifier of the rule insight.
	Ref string `json:"ref,omitempty"` // The ref name that the evaluation ran on.
	Repository_name string `json:"repository_name,omitempty"` // The name of the repository without the `.git` extension.
	Actor_name string `json:"actor_name,omitempty"` // The handle for the GitHub user account.
	Before_sha string `json:"before_sha,omitempty"` // The first commit sha before the push evaluation.
	After_sha string `json:"after_sha,omitempty"` // The last commit sha in the push evaluation.
	Evaluation_result string `json:"evaluation_result,omitempty"` // The result of the rule evaluations for rules with the `active` and `evaluate` enforcement statuses, demonstrating whether rules would pass or fail if all rules in the rule suite were `active`. Null if no rules with `evaluate` enforcement status were run.
}

// Enterprise represents the Enterprise schema from the OpenAPI specification
type Enterprise struct {
	Node_id string `json:"node_id"`
	Updated_at string `json:"updated_at"`
	Website_url string `json:"website_url,omitempty"` // The enterprise's website URL.
	Description string `json:"description,omitempty"` // A short description of the enterprise.
	Slug string `json:"slug"` // The slug url identifier for the enterprise.
	Avatar_url string `json:"avatar_url"`
	Html_url string `json:"html_url"`
	Name string `json:"name"` // The name of the enterprise.
	Id int `json:"id"` // Unique identifier of the enterprise
	Created_at string `json:"created_at"`
}

// GeneratedType_Deployment_branch_policy_name_pattern_with_type represents the GeneratedType_Deployment_branch_policy_name_pattern_with_type schema from the OpenAPI specification
type GeneratedType_Deployment_branch_policy_name_pattern_with_type struct {
	TypeField string `json:"type,omitempty"` // Whether this rule targets a branch or tag
	Name string `json:"name"` // The name pattern that branches or tags must match in order to deploy to the environment. Wildcard characters will not match `/`. For example, to match branches that begin with `release/` and contain an additional single slash, use `release/*/*`. For more information about pattern matching syntax, see the [Ruby File.fnmatch documentation](https://ruby-doc.org/core-2.5.1/File.html#method-c-fnmatch).
}

// GeneratedType_Webhook_membership_added represents the GeneratedType_Webhook_membership_added schema from the OpenAPI specification
type GeneratedType_Webhook_membership_added struct {
	Team Webhooksteam `json:"team"` // Groups of organization members that gives permissions on specified repositories.
	Scope string `json:"scope"` // The scope of the membership. Currently, can only be `team`.
	Sender map[string]interface{} `json:"sender"`
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Member Webhooksuser `json:"member"`
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType_Interaction_limit represents the GeneratedType_Interaction_limit schema from the OpenAPI specification
type GeneratedType_Interaction_limit struct {
	Limit string `json:"limit"` // The type of GitHub user that can comment, open issues, or create pull requests while the interaction limit is in effect.
	Expiry string `json:"expiry,omitempty"` // The duration of the interaction restriction. Default: `one_day`.
}

// GeneratedType_Webhook_issues_deleted represents the GeneratedType_Webhook_issues_deleted schema from the OpenAPI specification
type GeneratedType_Webhook_issues_deleted struct {
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType_Team_discussion represents the GeneratedType_Team_discussion schema from the OpenAPI specification
type GeneratedType_Team_discussion struct {
	Comments_url string `json:"comments_url"`
	Url string `json:"url"`
	Body string `json:"body"` // The main text of the discussion.
	Last_edited_at string `json:"last_edited_at"`
	Private bool `json:"private"` // Whether or not this discussion should be restricted to team members and organization owners.
	Node_id string `json:"node_id"`
	Pinned bool `json:"pinned"` // Whether or not this discussion should be pinned for easy retrieval.
	Updated_at string `json:"updated_at"`
	Created_at string `json:"created_at"`
	Team_url string `json:"team_url"`
	Reactions GeneratedType_Reaction_rollup `json:"reactions,omitempty"`
	Author GeneratedType_Nullable_simple_user `json:"author"` // A GitHub user.
	Comments_count int `json:"comments_count"`
	Html_url string `json:"html_url"`
	Title string `json:"title"` // The title of the discussion.
	Body_version string `json:"body_version"` // The current version of the body content. If provided, this update operation will be rejected if the given version does not match the latest version on the server.
	Number int `json:"number"` // The unique sequence number of a team discussion.
	Body_html string `json:"body_html"`
}

// GeneratedType_Timeline_committed_event represents the GeneratedType_Timeline_committed_event schema from the OpenAPI specification
type GeneratedType_Timeline_committed_event struct {
	Parents []map[string]interface{} `json:"parents"`
	Url string `json:"url"`
	Author map[string]interface{} `json:"author"` // Identifying information for the git-user
	Html_url string `json:"html_url"`
	Message string `json:"message"` // Message describing the purpose of the commit
	Verification map[string]interface{} `json:"verification"`
	Sha string `json:"sha"` // SHA for the commit
	Tree map[string]interface{} `json:"tree"`
	Committer map[string]interface{} `json:"committer"` // Identifying information for the git-user
	Event string `json:"event,omitempty"`
	Node_id string `json:"node_id"`
}

// GeneratedType_Organization_simple_webhooks represents the GeneratedType_Organization_simple_webhooks schema from the OpenAPI specification
type GeneratedType_Organization_simple_webhooks struct {
	Hooks_url string `json:"hooks_url"`
	Repos_url string `json:"repos_url"`
	Url string `json:"url"`
	Members_url string `json:"members_url"`
	Node_id string `json:"node_id"`
	Avatar_url string `json:"avatar_url"`
	Id int `json:"id"`
	Issues_url string `json:"issues_url"`
	Login string `json:"login"`
	Public_members_url string `json:"public_members_url"`
	Description string `json:"description"`
	Events_url string `json:"events_url"`
}

// GeneratedType_Content_tree represents the GeneratedType_Content_tree schema from the OpenAPI specification
type GeneratedType_Content_tree struct {
	TypeField string `json:"type"`
	Git_url string `json:"git_url"`
	Name string `json:"name"`
	Sha string `json:"sha"`
	Content string `json:"content,omitempty"`
	Url string `json:"url"`
	Download_url string `json:"download_url"`
	Entries []map[string]interface{} `json:"entries,omitempty"`
	Html_url string `json:"html_url"`
	Path string `json:"path"`
	Size int `json:"size"`
	Links map[string]interface{} `json:"_links"`
	Encoding string `json:"encoding,omitempty"`
}

// GeneratedType_Team_simple represents the GeneratedType_Team_simple schema from the OpenAPI specification
type GeneratedType_Team_simple struct {
	Ldap_dn string `json:"ldap_dn,omitempty"` // Distinguished Name (DN) that team maps to within LDAP environment
	Name string `json:"name"` // Name of the team
	Notification_setting string `json:"notification_setting,omitempty"` // The notification setting the team has set
	Privacy string `json:"privacy,omitempty"` // The level of privacy this team should have
	Id int `json:"id"` // Unique identifier of the team
	Repositories_url string `json:"repositories_url"`
	Members_url string `json:"members_url"`
	Permission string `json:"permission"` // Permission that the team will have for its repositories
	Url string `json:"url"` // URL for the team
	Description string `json:"description"` // Description of the team
	Html_url string `json:"html_url"`
	Node_id string `json:"node_id"`
	Slug string `json:"slug"`
}

// GeneratedType_Actions_billing_usage represents the GeneratedType_Actions_billing_usage schema from the OpenAPI specification
type GeneratedType_Actions_billing_usage struct {
	Total_paid_minutes_used int `json:"total_paid_minutes_used"` // The total paid GitHub Actions minutes used.
	Included_minutes int `json:"included_minutes"` // The amount of free GitHub Actions minutes available.
	Minutes_used_breakdown map[string]interface{} `json:"minutes_used_breakdown"`
	Total_minutes_used int `json:"total_minutes_used"` // The sum of the free and paid GitHub Actions minutes used.
}

// GeneratedType_View_traffic represents the GeneratedType_View_traffic schema from the OpenAPI specification
type GeneratedType_View_traffic struct {
	Count int `json:"count"`
	Uniques int `json:"uniques"`
	Views []Traffic `json:"views"`
}

// GeneratedType_Repository_rule_committer_email_pattern represents the GeneratedType_Repository_rule_committer_email_pattern schema from the OpenAPI specification
type GeneratedType_Repository_rule_committer_email_pattern struct {
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	TypeField string `json:"type"`
}

// GeneratedType_Webhook_security_advisory_updated represents the GeneratedType_Webhook_security_advisory_updated schema from the OpenAPI specification
type GeneratedType_Webhook_security_advisory_updated struct {
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Security_advisory Webhookssecurityadvisory `json:"security_advisory"` // The details of the security advisory, including summary, description, and severity.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
}

// GeneratedType_Codespace_with_full_repository represents the GeneratedType_Codespace_with_full_repository schema from the OpenAPI specification
type GeneratedType_Codespace_with_full_repository struct {
	Pending_operation_disabled_reason string `json:"pending_operation_disabled_reason,omitempty"` // Text to show user when codespace is disabled by a pending operation
	Git_status map[string]interface{} `json:"git_status"` // Details about the codespace's git repository.
	Runtime_constraints map[string]interface{} `json:"runtime_constraints,omitempty"`
	State string `json:"state"` // State of this codespace.
	Pending_operation bool `json:"pending_operation,omitempty"` // Whether or not a codespace has a pending async operation. This would mean that the codespace is temporarily unavailable. The only thing that you can do with a codespace in this state is delete it.
	Start_url string `json:"start_url"` // API URL to start this codespace.
	Devcontainer_path string `json:"devcontainer_path,omitempty"` // Path to devcontainer.json from repo root used to create Codespace.
	Pulls_url string `json:"pulls_url"` // API URL for the Pull Request associated with this codespace, if any.
	Repository GeneratedType_Full_repository `json:"repository"` // Full Repository
	Id int64 `json:"id"`
	Prebuild bool `json:"prebuild"` // Whether the codespace was created from a prebuild.
	Stop_url string `json:"stop_url"` // API URL to stop this codespace.
	Machines_url string `json:"machines_url"` // API URL to access available alternate machine types for this codespace.
	Name string `json:"name"` // Automatically generated name of this codespace.
	Billable_owner GeneratedType_Simple_user `json:"billable_owner"` // A GitHub user.
	Idle_timeout_notice string `json:"idle_timeout_notice,omitempty"` // Text to show user when codespace idle timeout minutes has been overriden by an organization policy
	Location string `json:"location"` // The initally assigned location of a new codespace.
	Environment_id string `json:"environment_id"` // UUID identifying this codespace's environment.
	Display_name string `json:"display_name,omitempty"` // Display name for this codespace.
	Machine GeneratedType_Nullable_codespace_machine `json:"machine"` // A description of the machine powering a codespace.
	Updated_at string `json:"updated_at"`
	Idle_timeout_minutes int `json:"idle_timeout_minutes"` // The number of minutes of inactivity after which this codespace will be automatically stopped.
	Retention_period_minutes int `json:"retention_period_minutes,omitempty"` // Duration in minutes after codespace has gone idle in which it will be deleted. Must be integer minutes between 0 and 43200 (30 days).
	Last_used_at string `json:"last_used_at"` // Last known time this codespace was started.
	Owner GeneratedType_Simple_user `json:"owner"` // A GitHub user.
	Retention_expires_at string `json:"retention_expires_at,omitempty"` // When a codespace will be auto-deleted based on the "retention_period_minutes" and "last_used_at"
	Publish_url string `json:"publish_url,omitempty"` // API URL to publish this codespace to a new repository.
	Created_at string `json:"created_at"`
	Recent_folders []string `json:"recent_folders"`
	Url string `json:"url"` // API URL for this codespace.
	Web_url string `json:"web_url"` // URL to access this codespace on the web.
}

// GeneratedType_Tag_protection represents the GeneratedType_Tag_protection schema from the OpenAPI specification
type GeneratedType_Tag_protection struct {
	Updated_at string `json:"updated_at,omitempty"`
	Created_at string `json:"created_at,omitempty"`
	Enabled bool `json:"enabled,omitempty"`
	Id int `json:"id,omitempty"`
	Pattern string `json:"pattern"`
}

// GeneratedType_Webhook_installation_repositories_added represents the GeneratedType_Webhook_installation_repositories_added schema from the OpenAPI specification
type GeneratedType_Webhook_installation_repositories_added struct {
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Repositories_added []map[string]interface{} `json:"repositories_added"` // An array of repository objects, which were added to the installation.
	Repository_selection string `json:"repository_selection"` // Describe whether all repositories have been selected or there's a selection involved
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repositories_removed []map[string]interface{} `json:"repositories_removed"` // An array of repository objects, which were removed from the installation.
	Requester Webhooksuser `json:"requester"`
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Installation Installation `json:"installation"` // Installation
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// Topic represents the Topic schema from the OpenAPI specification
type Topic struct {
	Names []string `json:"names"`
}

// GeneratedType_Webhook_repository_created represents the GeneratedType_Webhook_repository_created schema from the OpenAPI specification
type GeneratedType_Webhook_repository_created struct {
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType_Webhook_merge_group_destroyed represents the GeneratedType_Webhook_merge_group_destroyed schema from the OpenAPI specification
type GeneratedType_Webhook_merge_group_destroyed struct {
	Merge_group GeneratedType_Merge_group `json:"merge_group"` // A group of pull requests that the merge queue has grouped together to be merged.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Reason string `json:"reason,omitempty"` // Explains why the merge group is being destroyed. The group could have been merged, removed from the queue (dequeued), or invalidated by an earlier queue entry being dequeued (invalidated).
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType_Webhook_pull_request_review_thread_resolved represents the GeneratedType_Webhook_pull_request_review_thread_resolved schema from the OpenAPI specification
type GeneratedType_Webhook_pull_request_review_thread_resolved struct {
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Thread map[string]interface{} `json:"thread"`
	Action string `json:"action"`
}

// GeneratedType_Webhook_installation_deleted represents the GeneratedType_Webhook_installation_deleted schema from the OpenAPI specification
type GeneratedType_Webhook_installation_deleted struct {
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation Installation `json:"installation"` // Installation
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repositories []map[string]interface{} `json:"repositories,omitempty"` // An array of repository objects that the installation can access.
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Requester interface{} `json:"requester,omitempty"`
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
}

// GeneratedType_Webhook_repository_deleted represents the GeneratedType_Webhook_repository_deleted schema from the OpenAPI specification
type GeneratedType_Webhook_repository_deleted struct {
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType_Webhook_repository_ruleset_edited represents the GeneratedType_Webhook_repository_ruleset_edited schema from the OpenAPI specification
type GeneratedType_Webhook_repository_ruleset_edited struct {
	Changes map[string]interface{} `json:"changes,omitempty"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Repository_ruleset GeneratedType_Repository_ruleset `json:"repository_ruleset"` // A set of rules to apply when specified conditions are met.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType_Billing_usage_report represents the GeneratedType_Billing_usage_report schema from the OpenAPI specification
type GeneratedType_Billing_usage_report struct {
	Usageitems []map[string]interface{} `json:"usageItems,omitempty"`
}

// GeneratedType_Nullable_code_of_conduct_simple represents the GeneratedType_Nullable_code_of_conduct_simple schema from the OpenAPI specification
type GeneratedType_Nullable_code_of_conduct_simple struct {
	Html_url string `json:"html_url"`
	Key string `json:"key"`
	Name string `json:"name"`
	Url string `json:"url"`
}

// GeneratedType_Webhook_discussion_transferred represents the GeneratedType_Webhook_discussion_transferred schema from the OpenAPI specification
type GeneratedType_Webhook_discussion_transferred struct {
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
}

// GeneratedType_Webhook_workflow_run_requested represents the GeneratedType_Webhook_workflow_run_requested schema from the OpenAPI specification
type GeneratedType_Webhook_workflow_run_requested struct {
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Workflow Webhooksworkflow `json:"workflow"`
	Workflow_run map[string]interface{} `json:"workflow_run"`
	Action string `json:"action"`
}

// GeneratedType_Webhook_issues_transferred represents the GeneratedType_Webhook_issues_transferred schema from the OpenAPI specification
type GeneratedType_Webhook_issues_transferred struct {
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue Webhooksissue2 `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
}

// GeneratedType_Repository_subscription represents the GeneratedType_Repository_subscription schema from the OpenAPI specification
type GeneratedType_Repository_subscription struct {
	Repository_url string `json:"repository_url"`
	Subscribed bool `json:"subscribed"` // Determines if notifications should be received from this repository.
	Url string `json:"url"`
	Created_at string `json:"created_at"`
	Ignored bool `json:"ignored"` // Determines if all notifications should be blocked from this repository.
	Reason string `json:"reason"`
}

// GeneratedType_Repository_webhooks represents the GeneratedType_Repository_webhooks schema from the OpenAPI specification
type GeneratedType_Repository_webhooks struct {
	Private bool `json:"private"` // Whether the repository is private or public.
	Allow_merge_commit bool `json:"allow_merge_commit,omitempty"` // Whether to allow merge commits for pull requests.
	Has_downloads bool `json:"has_downloads"` // Whether downloads are enabled.
	Owner GeneratedType_Simple_user `json:"owner"` // A GitHub user.
	Merge_commit_title string `json:"merge_commit_title,omitempty"` // The default value for a merge commit title. - `PR_TITLE` - default to the pull request's title. - `MERGE_MESSAGE` - default to the classic title for a merge message (e.g., Merge pull request #123 from branch-name).
	Pulls_url string `json:"pulls_url"`
	Contents_url string `json:"contents_url"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Languages_url string `json:"languages_url"`
	Mirror_url string `json:"mirror_url"`
	Node_id string `json:"node_id"`
	Ssh_url string `json:"ssh_url"`
	Id int64 `json:"id"` // Unique identifier of the repository
	Comments_url string `json:"comments_url"`
	Use_squash_pr_title_as_default bool `json:"use_squash_pr_title_as_default,omitempty"` // Whether a squash merge commit can use the pull request title as default. **This property is closing down. Please use `squash_merge_commit_title` instead.
	Visibility string `json:"visibility,omitempty"` // The repository visibility: public, private, or internal.
	Open_issues int `json:"open_issues"`
	Organization GeneratedType_Nullable_simple_user `json:"organization,omitempty"` // A GitHub user.
	Has_issues bool `json:"has_issues"` // Whether issues are enabled.
	Anonymous_access_enabled bool `json:"anonymous_access_enabled,omitempty"` // Whether anonymous git access is enabled for this repository
	Watchers_count int `json:"watchers_count"`
	Language string `json:"language"`
	Assignees_url string `json:"assignees_url"`
	Releases_url string `json:"releases_url"`
	Labels_url string `json:"labels_url"`
	Issue_comment_url string `json:"issue_comment_url"`
	Teams_url string `json:"teams_url"`
	Disabled bool `json:"disabled"` // Returns whether or not this repository disabled.
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"` // Whether to require contributors to sign off on web-based commits
	Homepage string `json:"homepage"`
	Created_at string `json:"created_at"`
	Git_url string `json:"git_url"`
	Allow_forking bool `json:"allow_forking,omitempty"` // Whether to allow forking this repo
	Git_tags_url string `json:"git_tags_url"`
	Hooks_url string `json:"hooks_url"`
	Pushed_at string `json:"pushed_at"`
	Blobs_url string `json:"blobs_url"`
	Keys_url string `json:"keys_url"`
	Forks int `json:"forks"`
	Has_wiki bool `json:"has_wiki"` // Whether the wiki is enabled.
	Html_url string `json:"html_url"`
	Contributors_url string `json:"contributors_url"`
	Svn_url string `json:"svn_url"`
	Issue_events_url string `json:"issue_events_url"`
	Compare_url string `json:"compare_url"`
	Description string `json:"description"`
	Master_branch string `json:"master_branch,omitempty"`
	Stargazers_url string `json:"stargazers_url"`
	Template_repository map[string]interface{} `json:"template_repository,omitempty"`
	Branches_url string `json:"branches_url"`
	Topics []string `json:"topics,omitempty"`
	Custom_properties map[string]interface{} `json:"custom_properties,omitempty"` // The custom properties that were defined for the repository. The keys are the custom property names, and the values are the corresponding custom property values.
	Subscribers_url string `json:"subscribers_url"`
	Full_name string `json:"full_name"`
	Squash_merge_commit_message string `json:"squash_merge_commit_message,omitempty"` // The default value for a squash merge commit message: - `PR_BODY` - default to the pull request's body. - `COMMIT_MESSAGES` - default to the branch's commit messages. - `BLANK` - default to a blank commit message.
	Commits_url string `json:"commits_url"`
	Has_discussions bool `json:"has_discussions,omitempty"` // Whether discussions are enabled.
	Forks_count int `json:"forks_count"`
	Merge_commit_message string `json:"merge_commit_message,omitempty"` // The default value for a merge commit message. - `PR_TITLE` - default to the pull request's title. - `PR_BODY` - default to the pull request's body. - `BLANK` - default to a blank commit message.
	Events_url string `json:"events_url"`
	Downloads_url string `json:"downloads_url"`
	Deployments_url string `json:"deployments_url"`
	Git_refs_url string `json:"git_refs_url"`
	Archived bool `json:"archived"` // Whether the repository is archived.
	Forks_url string `json:"forks_url"`
	Allow_squash_merge bool `json:"allow_squash_merge,omitempty"` // Whether to allow squash merges for pull requests.
	Git_commits_url string `json:"git_commits_url"`
	Merges_url string `json:"merges_url"`
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"` // Whether to delete head branches when pull requests are merged
	Statuses_url string `json:"statuses_url"`
	Notifications_url string `json:"notifications_url"`
	Subscribers_count int `json:"subscribers_count,omitempty"`
	Updated_at string `json:"updated_at"`
	Is_template bool `json:"is_template,omitempty"` // Whether this repository acts as a template that can be used to generate new repositories.
	Subscription_url string `json:"subscription_url"`
	Has_pages bool `json:"has_pages"`
	Collaborators_url string `json:"collaborators_url"`
	Stargazers_count int `json:"stargazers_count"`
	Issues_url string `json:"issues_url"`
	Allow_update_branch bool `json:"allow_update_branch,omitempty"` // Whether or not a pull request head branch that is behind its base branch can always be updated even if it is not required to be up to date before merging.
	Allow_rebase_merge bool `json:"allow_rebase_merge,omitempty"` // Whether to allow rebase merges for pull requests.
	Default_branch string `json:"default_branch"` // The default branch of the repository.
	Size int `json:"size"` // The size of the repository, in kilobytes. Size is calculated hourly. When a repository is initially created, the size is 0.
	Allow_auto_merge bool `json:"allow_auto_merge,omitempty"` // Whether to allow Auto-merge to be used on pull requests.
	Open_issues_count int `json:"open_issues_count"`
	Clone_url string `json:"clone_url"`
	Temp_clone_token string `json:"temp_clone_token,omitempty"`
	Network_count int `json:"network_count,omitempty"`
	Has_projects bool `json:"has_projects"` // Whether projects are enabled.
	License GeneratedType_Nullable_license_simple `json:"license"` // License Simple
	Name string `json:"name"` // The name of the repository.
	Starred_at string `json:"starred_at,omitempty"`
	Tags_url string `json:"tags_url"`
	Trees_url string `json:"trees_url"`
	Fork bool `json:"fork"`
	Url string `json:"url"`
	Archive_url string `json:"archive_url"`
	Squash_merge_commit_title string `json:"squash_merge_commit_title,omitempty"` // The default value for a squash merge commit title: - `PR_TITLE` - default to the pull request's title. - `COMMIT_OR_PR_TITLE` - default to the commit's title (if only one commit) or the pull request's title (when more than one commit).
	Watchers int `json:"watchers"`
	Milestones_url string `json:"milestones_url"`
}

// GeneratedType_Webhook_repository_advisory_reported represents the GeneratedType_Webhook_repository_advisory_reported schema from the OpenAPI specification
type GeneratedType_Webhook_repository_advisory_reported struct {
	Repository_advisory GeneratedType_Repository_advisory `json:"repository_advisory"` // A repository security advisory.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType_Webhook_installation_repositories_removed represents the GeneratedType_Webhook_installation_repositories_removed schema from the OpenAPI specification
type GeneratedType_Webhook_installation_repositories_removed struct {
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Installation Installation `json:"installation"` // Installation
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repositories_added []map[string]interface{} `json:"repositories_added"` // An array of repository objects, which were added to the installation.
	Repositories_removed []map[string]interface{} `json:"repositories_removed"` // An array of repository objects, which were removed from the installation.
	Repository_selection string `json:"repository_selection"` // Describe whether all repositories have been selected or there's a selection involved
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Requester Webhooksuser `json:"requester"`
}

// GeneratedType_Enterprise_webhooks represents the GeneratedType_Enterprise_webhooks schema from the OpenAPI specification
type GeneratedType_Enterprise_webhooks struct {
	Id int `json:"id"` // Unique identifier of the enterprise
	Updated_at string `json:"updated_at"`
	Description string `json:"description,omitempty"` // A short description of the enterprise.
	Slug string `json:"slug"` // The slug url identifier for the enterprise.
	Website_url string `json:"website_url,omitempty"` // The enterprise's website URL.
	Created_at string `json:"created_at"`
	Html_url string `json:"html_url"`
	Name string `json:"name"` // The name of the enterprise.
	Avatar_url string `json:"avatar_url"`
	Node_id string `json:"node_id"`
}

// GeneratedType_Webhook_workflow_job_completed represents the GeneratedType_Webhook_workflow_job_completed schema from the OpenAPI specification
type GeneratedType_Webhook_workflow_job_completed struct {
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Workflow_job interface{} `json:"workflow_job"`
	Action string `json:"action"`
	Deployment Deployment `json:"deployment,omitempty"` // A request for a specific ref(branch,sha,tag) to be deployed
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType_Runner_groups_org represents the GeneratedType_Runner_groups_org schema from the OpenAPI specification
type GeneratedType_Runner_groups_org struct {
	Selected_workflows []string `json:"selected_workflows,omitempty"` // List of workflows the runner group should be allowed to run. This setting will be ignored unless `restricted_to_workflows` is set to `true`.
	Workflow_restrictions_read_only bool `json:"workflow_restrictions_read_only,omitempty"` // If `true`, the `restricted_to_workflows` and `selected_workflows` fields cannot be modified.
	Allows_public_repositories bool `json:"allows_public_repositories"`
	Inherited bool `json:"inherited"`
	Restricted_to_workflows bool `json:"restricted_to_workflows,omitempty"` // If `true`, the runner group will be restricted to running only the workflows specified in the `selected_workflows` array.
	Runners_url string `json:"runners_url"`
	Selected_repositories_url string `json:"selected_repositories_url,omitempty"` // Link to the selected repositories resource for this runner group. Not present unless visibility was set to `selected`
	DefaultField bool `json:"default"`
	Id float64 `json:"id"`
	Name string `json:"name"`
	Visibility string `json:"visibility"`
	Hosted_runners_url string `json:"hosted_runners_url,omitempty"`
	Inherited_allows_public_repositories bool `json:"inherited_allows_public_repositories,omitempty"`
	Network_configuration_id string `json:"network_configuration_id,omitempty"` // The identifier of a hosted compute network configuration.
}

// GeneratedType_Webhook_custom_property_updated represents the GeneratedType_Webhook_custom_property_updated schema from the OpenAPI specification
type GeneratedType_Webhook_custom_property_updated struct {
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Definition GeneratedType_Custom_property `json:"definition"` // Custom property defined on an organization
}

// Page represents the Page schema from the OpenAPI specification
type Page struct {
	Html_url string `json:"html_url,omitempty"` // The web address the Page can be accessed from.
	Https_certificate GeneratedType_Pages_https_certificate `json:"https_certificate,omitempty"`
	Status string `json:"status"` // The status of the most recent build of the Page.
	Custom_404 bool `json:"custom_404"` // Whether the Page has a custom 404 page.
	Pending_domain_unverified_at string `json:"pending_domain_unverified_at,omitempty"` // The timestamp when a pending domain becomes unverified.
	Protected_domain_state string `json:"protected_domain_state,omitempty"` // The state if the domain is verified
	Public bool `json:"public"` // Whether the GitHub Pages site is publicly visible. If set to `true`, the site is accessible to anyone on the internet. If set to `false`, the site will only be accessible to users who have at least `read` access to the repository that published the site.
	Source GeneratedType_Pages_source_hash `json:"source,omitempty"`
	Https_enforced bool `json:"https_enforced,omitempty"` // Whether https is enabled on the domain
	Url string `json:"url"` // The API address for accessing this Page resource.
	Build_type string `json:"build_type,omitempty"` // The process in which the Page will be built.
	Cname string `json:"cname"` // The Pages site's custom domain
}

// GeneratedType_Webhook_repository_vulnerability_alert_create represents the GeneratedType_Webhook_repository_vulnerability_alert_create schema from the OpenAPI specification
type GeneratedType_Webhook_repository_vulnerability_alert_create struct {
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Alert Webhooksalert `json:"alert"` // The security alert of the vulnerable dependency.
}

// GeneratedType_Pull_request_simple represents the GeneratedType_Pull_request_simple schema from the OpenAPI specification
type GeneratedType_Pull_request_simple struct {
	Merge_commit_sha string `json:"merge_commit_sha"`
	Links map[string]interface{} `json:"_links"`
	User GeneratedType_Nullable_simple_user `json:"user"` // A GitHub user.
	Body string `json:"body"`
	Review_comment_url string `json:"review_comment_url"`
	Head map[string]interface{} `json:"head"`
	Merged_at string `json:"merged_at"`
	Locked bool `json:"locked"`
	Milestone GeneratedType_Nullable_milestone `json:"milestone"` // A collection of related issues and pull requests.
	Base map[string]interface{} `json:"base"`
	Review_comments_url string `json:"review_comments_url"`
	Assignees []GeneratedType_Simple_user `json:"assignees,omitempty"`
	Commits_url string `json:"commits_url"`
	Id int64 `json:"id"`
	State string `json:"state"`
	Node_id string `json:"node_id"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	Assignee GeneratedType_Nullable_simple_user `json:"assignee"` // A GitHub user.
	Active_lock_reason string `json:"active_lock_reason,omitempty"`
	Auto_merge GeneratedType_Auto_merge `json:"auto_merge"` // The status of auto merging a pull request.
	Created_at string `json:"created_at"`
	Diff_url string `json:"diff_url"`
	Number int `json:"number"`
	Closed_at string `json:"closed_at"`
	Issue_url string `json:"issue_url"`
	Labels []map[string]interface{} `json:"labels"`
	Title string `json:"title"`
	Draft bool `json:"draft,omitempty"` // Indicates whether or not the pull request is a draft.
	Statuses_url string `json:"statuses_url"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Requested_teams []Team `json:"requested_teams,omitempty"`
	Comments_url string `json:"comments_url"`
	Patch_url string `json:"patch_url"`
	Requested_reviewers []GeneratedType_Simple_user `json:"requested_reviewers,omitempty"`
	Html_url string `json:"html_url"`
}

// GeneratedType_Marketplace_purchase represents the GeneratedType_Marketplace_purchase schema from the OpenAPI specification
type GeneratedType_Marketplace_purchase struct {
	Url string `json:"url"`
	Email string `json:"email,omitempty"`
	Id int `json:"id"`
	Login string `json:"login"`
	Marketplace_pending_change map[string]interface{} `json:"marketplace_pending_change,omitempty"`
	Marketplace_purchase map[string]interface{} `json:"marketplace_purchase"`
	Organization_billing_email string `json:"organization_billing_email,omitempty"`
	TypeField string `json:"type"`
}

// GeneratedType_Webhook_code_scanning_alert_closed_by_user represents the GeneratedType_Webhook_code_scanning_alert_closed_by_user schema from the OpenAPI specification
type GeneratedType_Webhook_code_scanning_alert_closed_by_user struct {
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Ref string `json:"ref"` // The Git reference of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Commit_oid string `json:"commit_oid"` // The commit SHA of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Alert map[string]interface{} `json:"alert"` // The code scanning alert involved in the event.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType_Topic_search_result_item represents the GeneratedType_Topic_search_result_item schema from the OpenAPI specification
type GeneratedType_Topic_search_result_item struct {
	Curated bool `json:"curated"`
	Featured bool `json:"featured"`
	Released string `json:"released"`
	Aliases []map[string]interface{} `json:"aliases,omitempty"`
	Text_matches []map[string]interface{} `json:"text_matches,omitempty"`
	Repository_count int `json:"repository_count,omitempty"`
	Short_description string `json:"short_description"`
	Created_by string `json:"created_by"`
	Display_name string `json:"display_name"`
	Description string `json:"description"`
	Updated_at string `json:"updated_at"`
	Score float64 `json:"score"`
	Logo_url string `json:"logo_url,omitempty"`
	Related []map[string]interface{} `json:"related,omitempty"`
	Created_at string `json:"created_at"`
	Name string `json:"name"`
}

// GeneratedType_Webhook_projects_v2_item_converted represents the GeneratedType_Webhook_projects_v2_item_converted schema from the OpenAPI specification
type GeneratedType_Webhook_projects_v2_item_converted struct {
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Projects_v2_item GeneratedType_Projects_v2_item `json:"projects_v2_item"` // An item belonging to a project
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
}

// GeneratedType_Webhook_pull_request_auto_merge_enabled represents the GeneratedType_Webhook_pull_request_auto_merge_enabled schema from the OpenAPI specification
type GeneratedType_Webhook_pull_request_auto_merge_enabled struct {
	Action string `json:"action"`
	Number int `json:"number"`
	Pull_request map[string]interface{} `json:"pull_request"`
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Reason string `json:"reason,omitempty"`
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType_Team_full represents the GeneratedType_Team_full schema from the OpenAPI specification
type GeneratedType_Team_full struct {
	Html_url string `json:"html_url"`
	Ldap_dn string `json:"ldap_dn,omitempty"` // Distinguished Name (DN) that team maps to within LDAP environment
	Organization GeneratedType_Team_organization `json:"organization"` // Team Organization
	Updated_at string `json:"updated_at"`
	Node_id string `json:"node_id"`
	Repos_count int `json:"repos_count"`
	Notification_setting string `json:"notification_setting,omitempty"` // The notification setting the team has set
	Name string `json:"name"` // Name of the team
	Members_url string `json:"members_url"`
	Repositories_url string `json:"repositories_url"`
	Members_count int `json:"members_count"`
	Permission string `json:"permission"` // Permission that the team will have for its repositories
	Privacy string `json:"privacy,omitempty"` // The level of privacy this team should have
	Slug string `json:"slug"`
	Url string `json:"url"` // URL for the team
	Id int `json:"id"` // Unique identifier of the team
	Parent GeneratedType_Nullable_team_simple `json:"parent,omitempty"` // Groups of organization members that gives permissions on specified repositories.
	Created_at string `json:"created_at"`
	Description string `json:"description"`
}

// GeneratedType_Referrer_traffic represents the GeneratedType_Referrer_traffic schema from the OpenAPI specification
type GeneratedType_Referrer_traffic struct {
	Count int `json:"count"`
	Referrer string `json:"referrer"`
	Uniques int `json:"uniques"`
}

// GeneratedType_Webhook_project_closed represents the GeneratedType_Webhook_project_closed schema from the OpenAPI specification
type GeneratedType_Webhook_project_closed struct {
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Project Webhooksproject `json:"project"`
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// Webhookssecurityadvisory represents the Webhookssecurityadvisory schema from the OpenAPI specification
type Webhookssecurityadvisory struct {
	Description string `json:"description"`
	Identifiers []map[string]interface{} `json:"identifiers"`
	Cvss map[string]interface{} `json:"cvss"`
	Cvss_severities GeneratedType_Cvss_severities `json:"cvss_severities,omitempty"`
	Cwes []map[string]interface{} `json:"cwes"`
	Published_at string `json:"published_at"`
	References []map[string]interface{} `json:"references"`
	Ghsa_id string `json:"ghsa_id"`
	Vulnerabilities []map[string]interface{} `json:"vulnerabilities"`
	Withdrawn_at string `json:"withdrawn_at"`
	Severity string `json:"severity"`
	Summary string `json:"summary"`
	Updated_at string `json:"updated_at"`
}

// GeneratedType_Webhook_commit_comment_created represents the GeneratedType_Webhook_commit_comment_created schema from the OpenAPI specification
type GeneratedType_Webhook_commit_comment_created struct {
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"` // The action performed. Can be `created`.
	Comment map[string]interface{} `json:"comment"` // The [commit comment](${externalDocsUpapp/api/description/components/schemas/webhooks/issue-comment-created.yamlrl}/rest/commits/comments#get-a-commit-comment) resource.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType_Repository_rule_file_extension_restriction represents the GeneratedType_Repository_rule_file_extension_restriction schema from the OpenAPI specification
type GeneratedType_Repository_rule_file_extension_restriction struct {
	TypeField string `json:"type"`
	Parameters map[string]interface{} `json:"parameters,omitempty"`
}

// GeneratedType_Webhook_marketplace_purchase_pending_change represents the GeneratedType_Webhook_marketplace_purchase_pending_change schema from the OpenAPI specification
type GeneratedType_Webhook_marketplace_purchase_pending_change struct {
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Effective_date string `json:"effective_date"`
	Marketplace_purchase Webhooksmarketplacepurchase `json:"marketplace_purchase"`
	Previous_marketplace_purchase map[string]interface{} `json:"previous_marketplace_purchase,omitempty"`
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType_Runner_application represents the GeneratedType_Runner_application schema from the OpenAPI specification
type GeneratedType_Runner_application struct {
	Os string `json:"os"`
	Sha256_checksum string `json:"sha256_checksum,omitempty"`
	Temp_download_token string `json:"temp_download_token,omitempty"` // A short lived bearer token used to download the runner, if needed.
	Architecture string `json:"architecture"`
	Download_url string `json:"download_url"`
	Filename string `json:"filename"`
}

// GeneratedType_Repository_rule_pull_request represents the GeneratedType_Repository_rule_pull_request schema from the OpenAPI specification
type GeneratedType_Repository_rule_pull_request struct {
	TypeField string `json:"type"`
	Parameters map[string]interface{} `json:"parameters,omitempty"`
}

// Classroom represents the Classroom schema from the OpenAPI specification
type Classroom struct {
	Id int `json:"id"` // Unique identifier of the classroom.
	Name string `json:"name"` // The name of the classroom.
	Organization GeneratedType_Simple_classroom_organization `json:"organization"` // A GitHub organization.
	Url string `json:"url"` // The URL of the classroom on GitHub Classroom.
	Archived bool `json:"archived"` // Whether classroom is archived.
}

// GeneratedType_Starred_repository represents the GeneratedType_Starred_repository schema from the OpenAPI specification
type GeneratedType_Starred_repository struct {
	Repo Repository `json:"repo"` // A repository on GitHub.
	Starred_at string `json:"starred_at"`
}

// Language represents the Language schema from the OpenAPI specification
type Language struct {
}

// Activity represents the Activity schema from the OpenAPI specification
type Activity struct {
	Before string `json:"before"` // The SHA of the commit before the activity.
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	Ref string `json:"ref"` // The full Git reference, formatted as `refs/heads/<branch name>`.
	Timestamp string `json:"timestamp"` // The time when the activity occurred.
	Activity_type string `json:"activity_type"` // The type of the activity that was performed.
	Actor GeneratedType_Nullable_simple_user `json:"actor"` // A GitHub user.
	After string `json:"after"` // The SHA of the commit after the activity.
}

// GeneratedType_Webhook_pull_request_unassigned represents the GeneratedType_Webhook_pull_request_unassigned schema from the OpenAPI specification
type GeneratedType_Webhook_pull_request_unassigned struct {
	Pull_request map[string]interface{} `json:"pull_request"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Number int `json:"number"` // The pull request number.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Assignee Webhooksusermannequin `json:"assignee,omitempty"`
}

// GeneratedType_Actions_hosted_runner_image represents the GeneratedType_Actions_hosted_runner_image schema from the OpenAPI specification
type GeneratedType_Actions_hosted_runner_image struct {
	Source string `json:"source"` // The image provider.
	Display_name string `json:"display_name"` // Display name for this image.
	Id string `json:"id"` // The ID of the image. Use this ID for the `image` parameter when creating a new larger runner.
	Platform string `json:"platform"` // The operating system of the image.
	Size_gb int `json:"size_gb"` // Image size in GB.
}

// GeneratedType_Secret_scanning_alert_webhook represents the GeneratedType_Secret_scanning_alert_webhook schema from the OpenAPI specification
type GeneratedType_Secret_scanning_alert_webhook struct {
	Push_protection_bypassed_at string `json:"push_protection_bypassed_at,omitempty"` // The time that push protection was bypassed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Push_protection_bypass_request_html_url string `json:"push_protection_bypass_request_html_url,omitempty"` // The URL to a push protection bypass request.
	Resolution_comment string `json:"resolution_comment,omitempty"` // An optional comment to resolve an alert.
	Push_protection_bypass_request_comment string `json:"push_protection_bypass_request_comment,omitempty"` // An optional comment when requesting a push protection bypass.
	Secret_type_display_name string `json:"secret_type_display_name,omitempty"` // User-friendly name for the detected secret, matching the `secret_type`. For a list of built-in patterns, see "[Supported secret scanning patterns](https://docs.github.com/code-security/secret-scanning/introduction/supported-secret-scanning-patterns#supported-secrets)."
	Secret_type string `json:"secret_type,omitempty"` // The type of secret that secret scanning detected.
	Number int `json:"number,omitempty"` // The security alert number.
	Resolved_at string `json:"resolved_at,omitempty"` // The time that the alert was resolved in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Publicly_leaked bool `json:"publicly_leaked,omitempty"` // Whether the detected secret was publicly leaked.
	Url string `json:"url,omitempty"` // The REST API URL of the alert resource.
	Validity string `json:"validity,omitempty"` // The token status as of the latest validity check.
	Multi_repo bool `json:"multi_repo,omitempty"` // Whether the detected secret was found in multiple repositories in the same organization or business.
	Push_protection_bypass_request_reviewer GeneratedType_Nullable_simple_user `json:"push_protection_bypass_request_reviewer,omitempty"` // A GitHub user.
	Push_protection_bypassed bool `json:"push_protection_bypassed,omitempty"` // Whether push protection was bypassed for the detected secret.
	Resolution string `json:"resolution,omitempty"` // The reason for resolving the alert.
	Push_protection_bypassed_by GeneratedType_Nullable_simple_user `json:"push_protection_bypassed_by,omitempty"` // A GitHub user.
	Created_at string `json:"created_at,omitempty"` // The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Resolved_by GeneratedType_Nullable_simple_user `json:"resolved_by,omitempty"` // A GitHub user.
	Updated_at string `json:"updated_at,omitempty"` // The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Html_url string `json:"html_url,omitempty"` // The GitHub URL of the alert resource.
	Locations_url string `json:"locations_url,omitempty"` // The REST API URL of the code locations for this alert.
	Push_protection_bypass_request_reviewer_comment string `json:"push_protection_bypass_request_reviewer_comment,omitempty"` // An optional comment when reviewing a push protection bypass.
}

// GeneratedType_Actions_hosted_runner_limits represents the GeneratedType_Actions_hosted_runner_limits schema from the OpenAPI specification
type GeneratedType_Actions_hosted_runner_limits struct {
	Public_ips map[string]interface{} `json:"public_ips"` // Provides details of static public IP limits for GitHub-hosted Hosted Runners
}

// GeneratedType_Repository_rule_params_required_reviewer_configuration represents the GeneratedType_Repository_rule_params_required_reviewer_configuration schema from the OpenAPI specification
type GeneratedType_Repository_rule_params_required_reviewer_configuration struct {
	File_patterns []string `json:"file_patterns"` // Array of file patterns. Pull requests which change matching files must be approved by the specified team. File patterns use the same syntax as `.gitignore` files.
	Minimum_approvals int `json:"minimum_approvals"` // Minimum number of approvals required from the specified team. If set to zero, the team will be added to the pull request but approval is optional.
	Reviewer GeneratedType_Repository_rule_params_reviewer `json:"reviewer"` // A required reviewing team
}

// GeneratedType_Dependabot_alert represents the GeneratedType_Dependabot_alert schema from the OpenAPI specification
type GeneratedType_Dependabot_alert struct {
	Security_advisory GeneratedType_Dependabot_alert_security_advisory `json:"security_advisory"` // Details for the GitHub Security Advisory.
	Auto_dismissed_at string `json:"auto_dismissed_at,omitempty"` // The time that the alert was auto-dismissed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Number int `json:"number"` // The security alert number.
	State string `json:"state"` // The state of the Dependabot alert.
	Fixed_at string `json:"fixed_at"` // The time that the alert was no longer detected and was considered fixed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Html_url string `json:"html_url"` // The GitHub URL of the alert resource.
	Updated_at string `json:"updated_at"` // The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Security_vulnerability GeneratedType_Dependabot_alert_security_vulnerability `json:"security_vulnerability"` // Details pertaining to one vulnerable version range for the advisory.
	Dependency map[string]interface{} `json:"dependency"` // Details for the vulnerable dependency.
	Dismissed_by GeneratedType_Nullable_simple_user `json:"dismissed_by"` // A GitHub user.
	Dismissed_comment string `json:"dismissed_comment"` // An optional comment associated with the alert's dismissal.
	Url string `json:"url"` // The REST API URL of the alert resource.
	Dismissed_at string `json:"dismissed_at"` // The time that the alert was dismissed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Dismissed_reason string `json:"dismissed_reason"` // The reason that the alert was dismissed.
	Created_at string `json:"created_at"` // The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
}

// Webhooksrelease represents the Webhooksrelease schema from the OpenAPI specification
type Webhooksrelease struct {
	Prerelease bool `json:"prerelease"` // Whether the release is identified as a prerelease or a full release.
	Body string `json:"body"`
	Node_id string `json:"node_id"`
	Zipball_url string `json:"zipball_url"`
	Assets_url string `json:"assets_url"`
	Html_url string `json:"html_url"`
	Discussion_url string `json:"discussion_url,omitempty"`
	Url string `json:"url"`
	Target_commitish string `json:"target_commitish"` // Specifies the commitish value that determines where the Git tag is created from.
	Id int `json:"id"`
	Upload_url string `json:"upload_url"`
	Assets []map[string]interface{} `json:"assets"`
	Name string `json:"name"`
	Tarball_url string `json:"tarball_url"`
	Created_at string `json:"created_at"`
	Published_at string `json:"published_at"`
	Reactions map[string]interface{} `json:"reactions,omitempty"`
	Tag_name string `json:"tag_name"` // The name of the tag.
	Draft bool `json:"draft"` // Whether the release is a draft or published
	Author map[string]interface{} `json:"author"`
}

// GeneratedType_Repository_ruleset_conditions represents the GeneratedType_Repository_ruleset_conditions schema from the OpenAPI specification
type GeneratedType_Repository_ruleset_conditions struct {
	Ref_name map[string]interface{} `json:"ref_name,omitempty"`
}

// GeneratedType_Secret_scanning_push_protection_bypass represents the GeneratedType_Secret_scanning_push_protection_bypass schema from the OpenAPI specification
type GeneratedType_Secret_scanning_push_protection_bypass struct {
	Reason string `json:"reason,omitempty"` // The reason for bypassing push protection.
	Token_type string `json:"token_type,omitempty"` // The token type this bypass is for.
	Expire_at string `json:"expire_at,omitempty"` // The time that the bypass will expire in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
}

// GeneratedType_Custom_property_value represents the GeneratedType_Custom_property_value schema from the OpenAPI specification
type GeneratedType_Custom_property_value struct {
	Property_name string `json:"property_name"` // The name of the property
	Value string `json:"value"` // The value assigned to the property
}

// GeneratedType_Repository_rule_max_file_size represents the GeneratedType_Repository_rule_max_file_size schema from the OpenAPI specification
type GeneratedType_Repository_rule_max_file_size struct {
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	TypeField string `json:"type"`
}

// GeneratedType_Interaction_limit_response represents the GeneratedType_Interaction_limit_response schema from the OpenAPI specification
type GeneratedType_Interaction_limit_response struct {
	Expires_at string `json:"expires_at"`
	Limit string `json:"limit"` // The type of GitHub user that can comment, open issues, or create pull requests while the interaction limit is in effect.
	Origin string `json:"origin"`
}

// GeneratedType_Timeline_cross_referenced_event represents the GeneratedType_Timeline_cross_referenced_event schema from the OpenAPI specification
type GeneratedType_Timeline_cross_referenced_event struct {
	Actor GeneratedType_Simple_user `json:"actor,omitempty"` // A GitHub user.
	Created_at string `json:"created_at"`
	Event string `json:"event"`
	Source map[string]interface{} `json:"source"`
	Updated_at string `json:"updated_at"`
}

// GeneratedType_Webhook_registry_package_published represents the GeneratedType_Webhook_registry_package_published schema from the OpenAPI specification
type GeneratedType_Webhook_registry_package_published struct {
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Registry_package map[string]interface{} `json:"registry_package"`
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType_Nullable_organization_simple represents the GeneratedType_Nullable_organization_simple schema from the OpenAPI specification
type GeneratedType_Nullable_organization_simple struct {
	Members_url string `json:"members_url"`
	Repos_url string `json:"repos_url"`
	Url string `json:"url"`
	Avatar_url string `json:"avatar_url"`
	Issues_url string `json:"issues_url"`
	Public_members_url string `json:"public_members_url"`
	Description string `json:"description"`
	Events_url string `json:"events_url"`
	Hooks_url string `json:"hooks_url"`
	Id int `json:"id"`
	Login string `json:"login"`
	Node_id string `json:"node_id"`
}

// GeneratedType_Repository_rule_params_reviewer represents the GeneratedType_Repository_rule_params_reviewer schema from the OpenAPI specification
type GeneratedType_Repository_rule_params_reviewer struct {
	Id int `json:"id"` // ID of the reviewer which must review changes to matching files.
	TypeField string `json:"type"` // The type of the reviewer
}

// GeneratedType_Webhook_branch_protection_rule_edited represents the GeneratedType_Webhook_branch_protection_rule_edited schema from the OpenAPI specification
type GeneratedType_Webhook_branch_protection_rule_edited struct {
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Rule Webhooksrule `json:"rule"` // The branch protection rule. Includes a `name` and all the [branch protection settings](https://docs.github.com/github/administering-a-repository/defining-the-mergeability-of-pull-requests/about-protected-branches#about-branch-protection-settings) applied to branches that match the name. Binary settings are boolean. Multi-level configurations are one of `off`, `non_admins`, or `everyone`. Actor and build lists are arrays of strings.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes,omitempty"` // If the action was `edited`, the changes to the rule.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType_Repository_ruleset_conditions_repository_property_spec represents the GeneratedType_Repository_ruleset_conditions_repository_property_spec schema from the OpenAPI specification
type GeneratedType_Repository_ruleset_conditions_repository_property_spec struct {
	Name string `json:"name"` // The name of the repository property to target
	Property_values []string `json:"property_values"` // The values to match for the repository property
	Source string `json:"source,omitempty"` // The source of the repository property. Defaults to 'custom' if not specified.
}

// GeneratedType_Repository_rule_required_signatures represents the GeneratedType_Repository_rule_required_signatures schema from the OpenAPI specification
type GeneratedType_Repository_rule_required_signatures struct {
	TypeField string `json:"type"`
}

// GeneratedType_Ssh_signing_key represents the GeneratedType_Ssh_signing_key schema from the OpenAPI specification
type GeneratedType_Ssh_signing_key struct {
	Key string `json:"key"`
	Title string `json:"title"`
	Created_at string `json:"created_at"`
	Id int `json:"id"`
}

// GeneratedType_Team_role_assignment represents the GeneratedType_Team_role_assignment schema from the OpenAPI specification
type GeneratedType_Team_role_assignment struct {
	Id int `json:"id"`
	Permission string `json:"permission"`
	Repositories_url string `json:"repositories_url"`
	Url string `json:"url"`
	Name string `json:"name"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Node_id string `json:"node_id"`
	Parent GeneratedType_Nullable_team_simple `json:"parent"` // Groups of organization members that gives permissions on specified repositories.
	Html_url string `json:"html_url"`
	Description string `json:"description"`
	Assignment string `json:"assignment,omitempty"` // Determines if the team has a direct, indirect, or mixed relationship to a role
	Members_url string `json:"members_url"`
	Notification_setting string `json:"notification_setting,omitempty"`
	Privacy string `json:"privacy,omitempty"`
	Slug string `json:"slug"`
}

// GeneratedType_Secret_scanning_location_discussion_title represents the GeneratedType_Secret_scanning_location_discussion_title schema from the OpenAPI specification
type GeneratedType_Secret_scanning_location_discussion_title struct {
	Discussion_title_url string `json:"discussion_title_url"` // The URL to the discussion where the secret was detected.
}

// GeneratedType_Gitignore_template represents the GeneratedType_Gitignore_template schema from the OpenAPI specification
type GeneratedType_Gitignore_template struct {
	Name string `json:"name"`
	Source string `json:"source"`
}

// Milestone represents the Milestone schema from the OpenAPI specification
type Milestone struct {
	Id int `json:"id"`
	Labels_url string `json:"labels_url"`
	Url string `json:"url"`
	Closed_at string `json:"closed_at"`
	Description string `json:"description"`
	Creator GeneratedType_Nullable_simple_user `json:"creator"` // A GitHub user.
	Node_id string `json:"node_id"`
	State string `json:"state"` // The state of the milestone.
	Html_url string `json:"html_url"`
	Title string `json:"title"` // The title of the milestone.
	Due_on string `json:"due_on"`
	Closed_issues int `json:"closed_issues"`
	Updated_at string `json:"updated_at"`
	Number int `json:"number"` // The number of the milestone.
	Open_issues int `json:"open_issues"`
	Created_at string `json:"created_at"`
}

// GeneratedType_Pull_request_webhook represents the GeneratedType_Pull_request_webhook schema from the OpenAPI specification
type GeneratedType_Pull_request_webhook struct {
	Number int `json:"number"` // Number uniquely identifying the pull request within its repository.
	Node_id string `json:"node_id"`
	Rebaseable bool `json:"rebaseable,omitempty"`
	Updated_at string `json:"updated_at"`
	Merged_at string `json:"merged_at"`
	Milestone GeneratedType_Nullable_milestone `json:"milestone"` // A collection of related issues and pull requests.
	Title string `json:"title"` // The title of the pull request.
	Mergeable_state string `json:"mergeable_state"`
	Body string `json:"body"`
	Issue_url string `json:"issue_url"`
	Patch_url string `json:"patch_url"`
	Commits_url string `json:"commits_url"`
	Active_lock_reason string `json:"active_lock_reason,omitempty"`
	Base map[string]interface{} `json:"base"`
	Locked bool `json:"locked"`
	Review_comments_url string `json:"review_comments_url"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Commits int `json:"commits"`
	Comments int `json:"comments"`
	Requested_reviewers []GeneratedType_Simple_user `json:"requested_reviewers,omitempty"`
	Html_url string `json:"html_url"`
	Created_at string `json:"created_at"`
	Assignee GeneratedType_Nullable_simple_user `json:"assignee"` // A GitHub user.
	Deletions int `json:"deletions"`
	Additions int `json:"additions"`
	Closed_at string `json:"closed_at"`
	Labels []map[string]interface{} `json:"labels"`
	Head map[string]interface{} `json:"head"`
	Merge_commit_sha string `json:"merge_commit_sha"`
	Url string `json:"url"`
	Review_comment_url string `json:"review_comment_url"`
	Review_comments int `json:"review_comments"`
	User GeneratedType_Simple_user `json:"user"` // A GitHub user.
	Requested_teams []GeneratedType_Team_simple `json:"requested_teams,omitempty"`
	Auto_merge GeneratedType_Auto_merge `json:"auto_merge"` // The status of auto merging a pull request.
	Mergeable bool `json:"mergeable"`
	Changed_files int `json:"changed_files"`
	Maintainer_can_modify bool `json:"maintainer_can_modify"` // Indicates whether maintainers can modify the pull request.
	Assignees []GeneratedType_Simple_user `json:"assignees,omitempty"`
	Comments_url string `json:"comments_url"`
	Merged_by GeneratedType_Nullable_simple_user `json:"merged_by"` // A GitHub user.
	Draft bool `json:"draft,omitempty"` // Indicates whether or not the pull request is a draft.
	State string `json:"state"` // State of this Pull Request. Either `open` or `closed`.
	Links map[string]interface{} `json:"_links"`
	Merged bool `json:"merged"`
	Diff_url string `json:"diff_url"`
	Id int64 `json:"id"`
	Statuses_url string `json:"statuses_url"`
	Allow_auto_merge bool `json:"allow_auto_merge,omitempty"` // Whether to allow auto-merge for pull requests.
	Allow_update_branch bool `json:"allow_update_branch,omitempty"` // Whether to allow updating the pull request's branch.
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"` // Whether to delete head branches when pull requests are merged.
	Merge_commit_message string `json:"merge_commit_message,omitempty"` // The default value for a merge commit message. - `PR_TITLE` - default to the pull request's title. - `PR_BODY` - default to the pull request's body. - `BLANK` - default to a blank commit message.
	Merge_commit_title string `json:"merge_commit_title,omitempty"` // The default value for a merge commit title. - `PR_TITLE` - default to the pull request's title. - `MERGE_MESSAGE` - default to the classic title for a merge message (e.g., "Merge pull request #123 from branch-name").
	Squash_merge_commit_message string `json:"squash_merge_commit_message,omitempty"` // The default value for a squash merge commit message: - `PR_BODY` - default to the pull request's body. - `COMMIT_MESSAGES` - default to the branch's commit messages. - `BLANK` - default to a blank commit message.
	Squash_merge_commit_title string `json:"squash_merge_commit_title,omitempty"` // The default value for a squash merge commit title: - `PR_TITLE` - default to the pull request's title. - `COMMIT_OR_PR_TITLE` - default to the commit's title (if only one commit) or the pull request's title (when more than one commit).
	Use_squash_pr_title_as_default bool `json:"use_squash_pr_title_as_default,omitempty"` // Whether a squash merge commit can use the pull request title as default. **This property is closing down. Please use `squash_merge_commit_title` instead.**
}

// GeneratedType_License_simple represents the GeneratedType_License_simple schema from the OpenAPI specification
type GeneratedType_License_simple struct {
	Url string `json:"url"`
	Html_url string `json:"html_url,omitempty"`
	Key string `json:"key"`
	Name string `json:"name"`
	Node_id string `json:"node_id"`
	Spdx_id string `json:"spdx_id"`
}

// GeneratedType_Webhook_issue_comment_deleted represents the GeneratedType_Webhook_issue_comment_deleted schema from the OpenAPI specification
type GeneratedType_Webhook_issue_comment_deleted struct {
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) the comment belongs to.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Comment Webhooksissuecomment `json:"comment"` // The [comment](https://docs.github.com/rest/issues/comments#get-an-issue-comment) itself.
}

// GeneratedType_Webhook_repository_ruleset_created represents the GeneratedType_Webhook_repository_ruleset_created schema from the OpenAPI specification
type GeneratedType_Webhook_repository_ruleset_created struct {
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Repository_ruleset GeneratedType_Repository_ruleset `json:"repository_ruleset"` // A set of rules to apply when specified conditions are met.
}

// Metadata represents the Metadata schema from the OpenAPI specification
type Metadata struct {
}

// GeneratedType_Repository_ruleset represents the GeneratedType_Repository_ruleset schema from the OpenAPI specification
type GeneratedType_Repository_ruleset struct {
	Updated_at string `json:"updated_at,omitempty"`
	Enforcement string `json:"enforcement"` // The enforcement level of the ruleset. `evaluate` allows admins to test rules before enforcing them. Admins can view insights on the Rule Insights page (`evaluate` is only available with GitHub Enterprise).
	Rules []GeneratedType_Repository_rule `json:"rules,omitempty"`
	Source string `json:"source"` // The name of the source
	Conditions interface{} `json:"conditions,omitempty"`
	Node_id string `json:"node_id,omitempty"`
	Links map[string]interface{} `json:"_links,omitempty"`
	Bypass_actors []GeneratedType_Repository_ruleset_bypass_actor `json:"bypass_actors,omitempty"` // The actors that can bypass the rules in this ruleset
	Id int `json:"id"` // The ID of the ruleset
	Source_type string `json:"source_type,omitempty"` // The type of the source of the ruleset
	Target string `json:"target,omitempty"` // The target of the ruleset
	Created_at string `json:"created_at,omitempty"`
	Current_user_can_bypass string `json:"current_user_can_bypass,omitempty"` // The bypass type of the user making the API request for this ruleset. This field is only returned when querying the repository-level endpoint.
	Name string `json:"name"` // The name of the ruleset
}

// GeneratedType_Codespace_export_details represents the GeneratedType_Codespace_export_details schema from the OpenAPI specification
type GeneratedType_Codespace_export_details struct {
	Sha string `json:"sha,omitempty"` // Git commit SHA of the exported branch
	State string `json:"state,omitempty"` // State of the latest export
	Branch string `json:"branch,omitempty"` // Name of the exported branch
	Completed_at string `json:"completed_at,omitempty"` // Completion time of the last export operation
	Export_url string `json:"export_url,omitempty"` // Url for fetching export details
	Html_url string `json:"html_url,omitempty"` // Web url for the exported branch
	Id string `json:"id,omitempty"` // Id for the export details
}

// GeneratedType_Webhook_personal_access_token_request_denied represents the GeneratedType_Webhook_personal_access_token_request_denied schema from the OpenAPI specification
type GeneratedType_Webhook_personal_access_token_request_denied struct {
	Installation GeneratedType_Simple_installation `json:"installation"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Personal_access_token_request GeneratedType_Personal_access_token_request `json:"personal_access_token_request"` // Details of a Personal Access Token Request.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType_Webhook_secret_scanning_alert_created represents the GeneratedType_Webhook_secret_scanning_alert_created schema from the OpenAPI specification
type GeneratedType_Webhook_secret_scanning_alert_created struct {
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Alert GeneratedType_Secret_scanning_alert_webhook `json:"alert"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType_Webhook_membership_removed represents the GeneratedType_Webhook_membership_removed schema from the OpenAPI specification
type GeneratedType_Webhook_membership_removed struct {
	Member Webhooksuser `json:"member"`
	Action string `json:"action"`
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Scope string `json:"scope"` // The scope of the membership. Currently, can only be `team`.
	Sender map[string]interface{} `json:"sender"`
	Team Webhooksteam `json:"team"` // Groups of organization members that gives permissions on specified repositories.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType_Webhook_sponsorship_cancelled represents the GeneratedType_Webhook_sponsorship_cancelled schema from the OpenAPI specification
type GeneratedType_Webhook_sponsorship_cancelled struct {
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Sponsorship Webhookssponsorship `json:"sponsorship"`
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType_Simple_classroom_user represents the GeneratedType_Simple_classroom_user schema from the OpenAPI specification
type GeneratedType_Simple_classroom_user struct {
	Avatar_url string `json:"avatar_url"`
	Html_url string `json:"html_url"`
	Id int `json:"id"`
	Login string `json:"login"`
}

// GeneratedType_Gist_comment represents the GeneratedType_Gist_comment schema from the OpenAPI specification
type GeneratedType_Gist_comment struct {
	Created_at string `json:"created_at"`
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	User GeneratedType_Nullable_simple_user `json:"user"` // A GitHub user.
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Body string `json:"body"` // The comment text.
}

// GeneratedType_Simple_classroom_organization represents the GeneratedType_Simple_classroom_organization schema from the OpenAPI specification
type GeneratedType_Simple_classroom_organization struct {
	Avatar_url string `json:"avatar_url"`
	Html_url string `json:"html_url"`
	Id int `json:"id"`
	Login string `json:"login"`
	Name string `json:"name"`
	Node_id string `json:"node_id"`
}

// GeneratedType_Combined_commit_status represents the GeneratedType_Combined_commit_status schema from the OpenAPI specification
type GeneratedType_Combined_commit_status struct {
	Total_count int `json:"total_count"`
	Url string `json:"url"`
	Commit_url string `json:"commit_url"`
	Repository GeneratedType_Minimal_repository `json:"repository"` // Minimal Repository
	Sha string `json:"sha"`
	State string `json:"state"`
	Statuses []GeneratedType_Simple_commit_status `json:"statuses"`
}

// Webhooksworkflowjobrun represents the Webhooksworkflowjobrun schema from the OpenAPI specification
type Webhooksworkflowjobrun struct {
	Updated_at string `json:"updated_at"`
	Conclusion interface{} `json:"conclusion"`
	Created_at string `json:"created_at"`
	Environment string `json:"environment"`
	Html_url string `json:"html_url"`
	Id int `json:"id"`
	Name interface{} `json:"name"`
	Status string `json:"status"`
}

// GeneratedType_Webhook_secret_scanning_alert_reopened represents the GeneratedType_Webhook_secret_scanning_alert_reopened schema from the OpenAPI specification
type GeneratedType_Webhook_secret_scanning_alert_reopened struct {
	Alert GeneratedType_Secret_scanning_alert_webhook `json:"alert"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType_Webhook_milestone_opened represents the GeneratedType_Webhook_milestone_opened schema from the OpenAPI specification
type GeneratedType_Webhook_milestone_opened struct {
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Milestone Webhooksmilestone3 `json:"milestone"` // A collection of related issues and pull requests.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType_Webhook_team_created represents the GeneratedType_Webhook_team_created schema from the OpenAPI specification
type GeneratedType_Webhook_team_created struct {
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository map[string]interface{} `json:"repository,omitempty"` // A git repository
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Team Webhooksteam1 `json:"team"` // Groups of organization members that gives permissions on specified repositories.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// Webhooksissuecomment represents the Webhooksissuecomment schema from the OpenAPI specification
type Webhooksissuecomment struct {
	Performed_via_github_app Integration `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Updated_at string `json:"updated_at"`
	Html_url string `json:"html_url"`
	Issue_url string `json:"issue_url"`
	Created_at string `json:"created_at"`
	Reactions map[string]interface{} `json:"reactions"`
	Url string `json:"url"` // URL for the issue comment
	User map[string]interface{} `json:"user"`
	Body string `json:"body"` // Contents of the issue comment
	Id int64 `json:"id"` // Unique identifier of the issue comment
	Node_id string `json:"node_id"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
}

// GeneratedType_Webhook_issues_labeled represents the GeneratedType_Webhook_issues_labeled schema from the OpenAPI specification
type GeneratedType_Webhook_issues_labeled struct {
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Label Webhookslabel `json:"label,omitempty"`
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
}

// GeneratedType_Marketplace_listing_plan represents the GeneratedType_Marketplace_listing_plan schema from the OpenAPI specification
type GeneratedType_Marketplace_listing_plan struct {
	Price_model string `json:"price_model"`
	State string `json:"state"`
	Unit_name string `json:"unit_name"`
	Accounts_url string `json:"accounts_url"`
	Description string `json:"description"`
	Url string `json:"url"`
	Has_free_trial bool `json:"has_free_trial"`
	Id int `json:"id"`
	Monthly_price_in_cents int `json:"monthly_price_in_cents"`
	Name string `json:"name"`
	Yearly_price_in_cents int `json:"yearly_price_in_cents"`
	Bullets []string `json:"bullets"`
	Number int `json:"number"`
}

// GeneratedType_Webhook_personal_access_token_request_approved represents the GeneratedType_Webhook_personal_access_token_request_approved schema from the OpenAPI specification
type GeneratedType_Webhook_personal_access_token_request_approved struct {
	Installation GeneratedType_Simple_installation `json:"installation"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Personal_access_token_request GeneratedType_Personal_access_token_request `json:"personal_access_token_request"` // Details of a Personal Access Token Request.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType_Webhook_sub_issues_parent_issue_removed represents the GeneratedType_Webhook_sub_issues_parent_issue_removed schema from the OpenAPI specification
type GeneratedType_Webhook_sub_issues_parent_issue_removed struct {
	Parent_issue_id float64 `json:"parent_issue_id"` // The ID of the parent issue.
	Parent_issue_repo Repository `json:"parent_issue_repo"` // A repository on GitHub.
	Sub_issue Issue `json:"sub_issue"` // Issues are a great way to keep track of tasks, enhancements, and bugs for your projects.
	Sub_issue_id float64 `json:"sub_issue_id"` // The ID of the sub-issue.
	Parent_issue Issue `json:"parent_issue"` // Issues are a great way to keep track of tasks, enhancements, and bugs for your projects.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// Webhooksusermannequin represents the Webhooksusermannequin schema from the OpenAPI specification
type Webhooksusermannequin struct {
	Node_id string `json:"node_id,omitempty"`
	Starred_url string `json:"starred_url,omitempty"`
	TypeField string `json:"type,omitempty"`
	Avatar_url string `json:"avatar_url,omitempty"`
	Gists_url string `json:"gists_url,omitempty"`
	Site_admin bool `json:"site_admin,omitempty"`
	Events_url string `json:"events_url,omitempty"`
	Received_events_url string `json:"received_events_url,omitempty"`
	Html_url string `json:"html_url,omitempty"`
	Login string `json:"login"`
	Subscriptions_url string `json:"subscriptions_url,omitempty"`
	Followers_url string `json:"followers_url,omitempty"`
	Deleted bool `json:"deleted,omitempty"`
	Following_url string `json:"following_url,omitempty"`
	Repos_url string `json:"repos_url,omitempty"`
	Email string `json:"email,omitempty"`
	Organizations_url string `json:"organizations_url,omitempty"`
	User_view_type string `json:"user_view_type,omitempty"`
	Gravatar_id string `json:"gravatar_id,omitempty"`
	Id int `json:"id"`
	Url string `json:"url,omitempty"`
	Name string `json:"name,omitempty"`
}

// GeneratedType_Secret_scanning_location_pull_request_comment represents the GeneratedType_Secret_scanning_location_pull_request_comment schema from the OpenAPI specification
type GeneratedType_Secret_scanning_location_pull_request_comment struct {
	Pull_request_comment_url string `json:"pull_request_comment_url"` // The API URL to get the pull request comment where the secret was detected.
}

// GeneratedType_Code_scanning_autofix_commits_response represents the GeneratedType_Code_scanning_autofix_commits_response schema from the OpenAPI specification
type GeneratedType_Code_scanning_autofix_commits_response struct {
	Sha string `json:"sha,omitempty"` // SHA of commit with autofix.
	Target_ref string `json:"target_ref,omitempty"` // The Git reference of target branch for the commit. For more information, see "[Git References](https://git-scm.com/book/en/v2/Git-Internals-Git-References)" in the Git documentation.
}

// GeneratedType_Webhook_repository_vulnerability_alert_resolve represents the GeneratedType_Webhook_repository_vulnerability_alert_resolve schema from the OpenAPI specification
type GeneratedType_Webhook_repository_vulnerability_alert_resolve struct {
	Alert map[string]interface{} `json:"alert"` // The security alert of the vulnerable dependency.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType_Webhook_pull_request_review_edited represents the GeneratedType_Webhook_pull_request_review_edited schema from the OpenAPI specification
type GeneratedType_Webhook_pull_request_review_edited struct {
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Review Webhooksreview `json:"review"` // The review that was affected.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Pull_request map[string]interface{} `json:"pull_request"`
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType_Marketplace_account represents the GeneratedType_Marketplace_account schema from the OpenAPI specification
type GeneratedType_Marketplace_account struct {
	Login string `json:"login"`
	Node_id string `json:"node_id,omitempty"`
	Organization_billing_email string `json:"organization_billing_email,omitempty"`
	TypeField string `json:"type"`
	Url string `json:"url"`
	Email string `json:"email,omitempty"`
	Id int `json:"id"`
}

// GeneratedType_User_marketplace_purchase represents the GeneratedType_User_marketplace_purchase schema from the OpenAPI specification
type GeneratedType_User_marketplace_purchase struct {
	Next_billing_date string `json:"next_billing_date"`
	On_free_trial bool `json:"on_free_trial"`
	Plan GeneratedType_Marketplace_listing_plan `json:"plan"` // Marketplace Listing Plan
	Unit_count int `json:"unit_count"`
	Updated_at string `json:"updated_at"`
	Account GeneratedType_Marketplace_account `json:"account"`
	Billing_cycle string `json:"billing_cycle"`
	Free_trial_ends_on string `json:"free_trial_ends_on"`
}

// GeneratedType_Nullable_scoped_installation represents the GeneratedType_Nullable_scoped_installation schema from the OpenAPI specification
type GeneratedType_Nullable_scoped_installation struct {
	Single_file_paths []string `json:"single_file_paths,omitempty"`
	Account GeneratedType_Simple_user `json:"account"` // A GitHub user.
	Has_multiple_single_files bool `json:"has_multiple_single_files,omitempty"`
	Permissions GeneratedType_App_permissions `json:"permissions"` // The permissions granted to the user access token.
	Repositories_url string `json:"repositories_url"`
	Repository_selection string `json:"repository_selection"` // Describe whether all repositories have been selected or there's a selection involved
	Single_file_name string `json:"single_file_name"`
}

// GeneratedType_Repo_codespaces_secret represents the GeneratedType_Repo_codespaces_secret schema from the OpenAPI specification
type GeneratedType_Repo_codespaces_secret struct {
	Updated_at string `json:"updated_at"`
	Created_at string `json:"created_at"`
	Name string `json:"name"` // The name of the secret.
}

// GeneratedType_Webhook_project_column_edited represents the GeneratedType_Webhook_project_column_edited schema from the OpenAPI specification
type GeneratedType_Webhook_project_column_edited struct {
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Project_column Webhooksprojectcolumn `json:"project_column"`
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType_Nullable_simple_commit represents the GeneratedType_Nullable_simple_commit schema from the OpenAPI specification
type GeneratedType_Nullable_simple_commit struct {
	Message string `json:"message"` // Message describing the purpose of the commit
	Timestamp string `json:"timestamp"` // Timestamp of the commit
	Tree_id string `json:"tree_id"` // SHA for the commit's tree
	Author map[string]interface{} `json:"author"` // Information about the Git author
	Committer map[string]interface{} `json:"committer"` // Information about the Git committer
	Id string `json:"id"` // SHA for the commit
}

// GeneratedType_Repository_rule_merge_queue represents the GeneratedType_Repository_rule_merge_queue schema from the OpenAPI specification
type GeneratedType_Repository_rule_merge_queue struct {
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	TypeField string `json:"type"`
}

// Thread represents the Thread schema from the OpenAPI specification
type Thread struct {
	Subject map[string]interface{} `json:"subject"`
	Unread bool `json:"unread"`
	Id string `json:"id"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	Repository GeneratedType_Minimal_repository `json:"repository"` // Minimal Repository
	Reason string `json:"reason"`
	Subscription_url string `json:"subscription_url"`
	Last_read_at string `json:"last_read_at"`
}

// GeneratedType_Organization_actions_secret represents the GeneratedType_Organization_actions_secret schema from the OpenAPI specification
type GeneratedType_Organization_actions_secret struct {
	Visibility string `json:"visibility"` // Visibility of a secret
	Created_at string `json:"created_at"`
	Name string `json:"name"` // The name of the secret.
	Selected_repositories_url string `json:"selected_repositories_url,omitempty"`
	Updated_at string `json:"updated_at"`
}

// GeneratedType_Secret_scanning_location represents the GeneratedType_Secret_scanning_location schema from the OpenAPI specification
type GeneratedType_Secret_scanning_location struct {
	Details interface{} `json:"details,omitempty"`
	TypeField string `json:"type,omitempty"` // The location type. Because secrets may be found in different types of resources (ie. code, comments, issues, pull requests, discussions), this field identifies the type of resource where the secret was found.
}

// GeneratedType_Webhook_dependabot_alert_reintroduced represents the GeneratedType_Webhook_dependabot_alert_reintroduced schema from the OpenAPI specification
type GeneratedType_Webhook_dependabot_alert_reintroduced struct {
	Action string `json:"action"`
	Alert GeneratedType_Dependabot_alert `json:"alert"` // A Dependabot alert.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
}

// GeneratedType_Webhook_pull_request_demilestoned represents the GeneratedType_Webhook_pull_request_demilestoned schema from the OpenAPI specification
type GeneratedType_Webhook_pull_request_demilestoned struct {
	Milestone Milestone `json:"milestone,omitempty"` // A collection of related issues and pull requests.
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request Webhookspullrequest5 `json:"pull_request"`
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType_Commit_comment represents the GeneratedType_Commit_comment schema from the OpenAPI specification
type GeneratedType_Commit_comment struct {
	Path string `json:"path"`
	Created_at string `json:"created_at"`
	Html_url string `json:"html_url"`
	Id int `json:"id"`
	Position int `json:"position"`
	Commit_id string `json:"commit_id"`
	Line int `json:"line"`
	Reactions GeneratedType_Reaction_rollup `json:"reactions,omitempty"`
	Url string `json:"url"`
	Updated_at string `json:"updated_at"`
	User GeneratedType_Nullable_simple_user `json:"user"` // A GitHub user.
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Body string `json:"body"`
	Node_id string `json:"node_id"`
}

// GeneratedType_Reaction_rollup represents the GeneratedType_Reaction_rollup schema from the OpenAPI specification
type GeneratedType_Reaction_rollup struct {
	Field1 int `json:"-1"`
	Confused int `json:"confused"`
	Laugh int `json:"laugh"`
	Total_count int `json:"total_count"`
	Field1_1 int `json:"+1"`
	Hooray int `json:"hooray"`
	Rocket int `json:"rocket"`
	Url string `json:"url"`
	Eyes int `json:"eyes"`
	Heart int `json:"heart"`
}

// GeneratedType_Nullable_milestone represents the GeneratedType_Nullable_milestone schema from the OpenAPI specification
type GeneratedType_Nullable_milestone struct {
	Open_issues int `json:"open_issues"`
	Created_at string `json:"created_at"`
	Number int `json:"number"` // The number of the milestone.
	Updated_at string `json:"updated_at"`
	Creator GeneratedType_Nullable_simple_user `json:"creator"` // A GitHub user.
	Description string `json:"description"`
	Url string `json:"url"`
	Node_id string `json:"node_id"`
	Labels_url string `json:"labels_url"`
	Id int `json:"id"`
	Html_url string `json:"html_url"`
	Closed_at string `json:"closed_at"`
	Title string `json:"title"` // The title of the milestone.
	Due_on string `json:"due_on"`
	Closed_issues int `json:"closed_issues"`
	State string `json:"state"` // The state of the milestone.
}

// GeneratedType_Webhook_custom_property_deleted represents the GeneratedType_Webhook_custom_property_deleted schema from the OpenAPI specification
type GeneratedType_Webhook_custom_property_deleted struct {
	Definition map[string]interface{} `json:"definition"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType_Webhook_issues_unlabeled represents the GeneratedType_Webhook_issues_unlabeled schema from the OpenAPI specification
type GeneratedType_Webhook_issues_unlabeled struct {
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue Webhooksissue `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Label Webhookslabel `json:"label,omitempty"`
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType_Repository_advisory_update represents the GeneratedType_Repository_advisory_update schema from the OpenAPI specification
type GeneratedType_Repository_advisory_update struct {
	Severity string `json:"severity,omitempty"` // The severity of the advisory. You must choose between setting this field or `cvss_vector_string`.
	Vulnerabilities []map[string]interface{} `json:"vulnerabilities,omitempty"` // A product affected by the vulnerability detailed in a repository security advisory.
	Credits []map[string]interface{} `json:"credits,omitempty"` // A list of users receiving credit for their participation in the security advisory.
	Cve_id string `json:"cve_id,omitempty"` // The Common Vulnerabilities and Exposures (CVE) ID.
	Cvss_vector_string string `json:"cvss_vector_string,omitempty"` // The CVSS vector that calculates the severity of the advisory. You must choose between setting this field or `severity`.
	Cwe_ids []string `json:"cwe_ids,omitempty"` // A list of Common Weakness Enumeration (CWE) IDs.
	Description string `json:"description,omitempty"` // A detailed description of what the advisory impacts.
	State string `json:"state,omitempty"` // The state of the advisory.
	Collaborating_users []string `json:"collaborating_users,omitempty"` // A list of usernames who have been granted write access to the advisory.
	Summary string `json:"summary,omitempty"` // A short summary of the advisory.
	Collaborating_teams []string `json:"collaborating_teams,omitempty"` // A list of team slugs which have been granted write access to the advisory.
}

// GeneratedType_Webhook_projects_v2_item_archived represents the GeneratedType_Webhook_projects_v2_item_archived schema from the OpenAPI specification
type GeneratedType_Webhook_projects_v2_item_archived struct {
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Projects_v2_item GeneratedType_Projects_v2_item `json:"projects_v2_item"` // An item belonging to a project
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes Webhooksprojectchanges `json:"changes"`
}

// GeneratedType_Organization_role represents the GeneratedType_Organization_role schema from the OpenAPI specification
type GeneratedType_Organization_role struct {
	Created_at string `json:"created_at"` // The date and time the role was created.
	Description string `json:"description,omitempty"` // A short description about who this role is for or what permissions it grants.
	Permissions []string `json:"permissions"` // A list of permissions included in this role.
	Base_role string `json:"base_role,omitempty"` // The system role from which this role inherits permissions.
	Source string `json:"source,omitempty"` // Source answers the question, "where did this role come from?"
	Updated_at string `json:"updated_at"` // The date and time the role was last updated.
	Id int64 `json:"id"` // The unique identifier of the role.
	Organization GeneratedType_Nullable_simple_user `json:"organization"` // A GitHub user.
	Name string `json:"name"` // The name of the role.
}

// GeneratedType_Webhook_installation_suspend represents the GeneratedType_Webhook_installation_suspend schema from the OpenAPI specification
type GeneratedType_Webhook_installation_suspend struct {
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation Installation `json:"installation"` // Installation
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repositories []map[string]interface{} `json:"repositories,omitempty"` // An array of repository objects that the installation can access.
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Requester interface{} `json:"requester,omitempty"`
}

// GeneratedType_Actions_set_default_workflow_permissions represents the GeneratedType_Actions_set_default_workflow_permissions schema from the OpenAPI specification
type GeneratedType_Actions_set_default_workflow_permissions struct {
	Can_approve_pull_request_reviews bool `json:"can_approve_pull_request_reviews,omitempty"` // Whether GitHub Actions can approve pull requests. Enabling this can be a security risk.
	Default_workflow_permissions string `json:"default_workflow_permissions,omitempty"` // The default workflow permissions granted to the GITHUB_TOKEN when running workflows.
}

// GeneratedType_Copilot_seat_details represents the GeneratedType_Copilot_seat_details schema from the OpenAPI specification
type GeneratedType_Copilot_seat_details struct {
	Last_activity_at string `json:"last_activity_at,omitempty"` // Timestamp of user's last GitHub Copilot activity, in ISO 8601 format.
	Organization GeneratedType_Nullable_organization_simple `json:"organization,omitempty"` // A GitHub organization.
	Last_activity_editor string `json:"last_activity_editor,omitempty"` // Last editor that was used by the user for a GitHub Copilot completion.
	Pending_cancellation_date string `json:"pending_cancellation_date,omitempty"` // The pending cancellation date for the seat, in `YYYY-MM-DD` format. This will be null unless the assignee's Copilot access has been canceled during the current billing cycle. If the seat has been cancelled, this corresponds to the start of the organization's next billing cycle.
	Plan_type string `json:"plan_type,omitempty"` // The Copilot plan of the organization, or the parent enterprise, when applicable.
	Assigning_team interface{} `json:"assigning_team,omitempty"` // The team through which the assignee is granted access to GitHub Copilot, if applicable.
	Created_at string `json:"created_at"` // Timestamp of when the assignee was last granted access to GitHub Copilot, in ISO 8601 format.
	Updated_at string `json:"updated_at,omitempty"` // **Closing down notice:** This field is no longer relevant and is closing down. Use the `created_at` field to determine when the assignee was last granted access to GitHub Copilot. Timestamp of when the assignee's GitHub Copilot access was last updated, in ISO 8601 format.
	Assignee GeneratedType_Nullable_simple_user `json:"assignee,omitempty"` // A GitHub user.
}

// GeneratedType_Selected_actions represents the GeneratedType_Selected_actions schema from the OpenAPI specification
type GeneratedType_Selected_actions struct {
	Verified_allowed bool `json:"verified_allowed,omitempty"` // Whether actions from GitHub Marketplace verified creators are allowed. Set to `true` to allow all actions by GitHub Marketplace verified creators.
	Github_owned_allowed bool `json:"github_owned_allowed,omitempty"` // Whether GitHub-owned actions are allowed. For example, this includes the actions in the `actions` organization.
	Patterns_allowed []string `json:"patterns_allowed,omitempty"` // Specifies a list of string-matching patterns to allow specific action(s) and reusable workflow(s). Wildcards, tags, and SHAs are allowed. For example, `monalisa/octocat@*`, `monalisa/octocat@v2`, `monalisa/*`. > [!NOTE] > The `patterns_allowed` setting only applies to public repositories.
}

// GeneratedType_Actions_hosted_runner_machine_spec represents the GeneratedType_Actions_hosted_runner_machine_spec schema from the OpenAPI specification
type GeneratedType_Actions_hosted_runner_machine_spec struct {
	Storage_gb int `json:"storage_gb"` // The available SSD storage for the machine spec.
	Cpu_cores int `json:"cpu_cores"` // The number of cores.
	Id string `json:"id"` // The ID used for the `size` parameter when creating a new runner.
	Memory_gb int `json:"memory_gb"` // The available RAM for the machine spec.
}

// GeneratedType_Actions_cache_usage_org_enterprise represents the GeneratedType_Actions_cache_usage_org_enterprise schema from the OpenAPI specification
type GeneratedType_Actions_cache_usage_org_enterprise struct {
	Total_active_caches_count int `json:"total_active_caches_count"` // The count of active caches across all repositories of an enterprise or an organization.
	Total_active_caches_size_in_bytes int `json:"total_active_caches_size_in_bytes"` // The total size in bytes of all active cache items across all repositories of an enterprise or an organization.
}

// GeneratedType_Repository_advisory_credit represents the GeneratedType_Repository_advisory_credit schema from the OpenAPI specification
type GeneratedType_Repository_advisory_credit struct {
	State string `json:"state"` // The state of the user's acceptance of the credit.
	TypeField string `json:"type"` // The type of credit the user is receiving.
	User GeneratedType_Simple_user `json:"user"` // A GitHub user.
}

// GeneratedType_Secret_scanning_location_pull_request_review_comment represents the GeneratedType_Secret_scanning_location_pull_request_review_comment schema from the OpenAPI specification
type GeneratedType_Secret_scanning_location_pull_request_review_comment struct {
	Pull_request_review_comment_url string `json:"pull_request_review_comment_url"` // The API URL to get the pull request review comment where the secret was detected.
}

// GeneratedType_Labeled_issue_event represents the GeneratedType_Labeled_issue_event schema from the OpenAPI specification
type GeneratedType_Labeled_issue_event struct {
	Commit_id string `json:"commit_id"`
	Created_at string `json:"created_at"`
	Url string `json:"url"`
	Event string `json:"event"`
	Id int `json:"id"`
	Label map[string]interface{} `json:"label"`
	Commit_url string `json:"commit_url"`
	Node_id string `json:"node_id"`
	Performed_via_github_app GeneratedType_Nullable_integration `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Actor GeneratedType_Simple_user `json:"actor"` // A GitHub user.
}

// GeneratedType_Projects_v2_single_select_option represents the GeneratedType_Projects_v2_single_select_option schema from the OpenAPI specification
type GeneratedType_Projects_v2_single_select_option struct {
	Description string `json:"description,omitempty"`
	Id string `json:"id"`
	Name string `json:"name"`
	Color string `json:"color,omitempty"`
}

// GeneratedType_Webhook_organization_renamed represents the GeneratedType_Webhook_organization_renamed schema from the OpenAPI specification
type GeneratedType_Webhook_organization_renamed struct {
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes,omitempty"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Membership Webhooksmembership `json:"membership,omitempty"` // The membership between the user and the organization. Not present when the action is `member_invited`.
}

// GeneratedType_Webhook_org_block_unblocked represents the GeneratedType_Webhook_org_block_unblocked schema from the OpenAPI specification
type GeneratedType_Webhook_org_block_unblocked struct {
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Blocked_user Webhooksuser `json:"blocked_user"`
}

// GeneratedType_Webhook_pull_request_locked represents the GeneratedType_Webhook_pull_request_locked schema from the OpenAPI specification
type GeneratedType_Webhook_pull_request_locked struct {
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// Feed represents the Feed schema from the OpenAPI specification
type Feed struct {
	User_url string `json:"user_url"`
	Current_user_organization_urls []string `json:"current_user_organization_urls,omitempty"`
	Current_user_public_url string `json:"current_user_public_url,omitempty"`
	Current_user_url string `json:"current_user_url,omitempty"`
	Repository_discussions_url string `json:"repository_discussions_url,omitempty"` // A feed of discussions for a given repository.
	Links map[string]interface{} `json:"_links"`
	Current_user_actor_url string `json:"current_user_actor_url,omitempty"`
	Current_user_organization_url string `json:"current_user_organization_url,omitempty"`
	Repository_discussions_category_url string `json:"repository_discussions_category_url,omitempty"` // A feed of discussions for a given repository and category.
	Security_advisories_url string `json:"security_advisories_url,omitempty"`
	Timeline_url string `json:"timeline_url"`
}

// GeneratedType_Classroom_assignment_grade represents the GeneratedType_Classroom_assignment_grade schema from the OpenAPI specification
type GeneratedType_Classroom_assignment_grade struct {
	Github_username string `json:"github_username"` // GitHub username of the student
	Group_name string `json:"group_name,omitempty"` // If a group assignment, name of the group the student is in
	Starter_code_url string `json:"starter_code_url"` // URL of the starter code for the assignment
	Student_repository_url string `json:"student_repository_url"` // URL of the student's assignment repository
	Points_awarded int `json:"points_awarded"` // Number of points awarded to the student
	Student_repository_name string `json:"student_repository_name"` // Name of the student's assignment repository
	Assignment_name string `json:"assignment_name"` // Name of the assignment
	Assignment_url string `json:"assignment_url"` // URL of the assignment
	Points_available int `json:"points_available"` // Number of points available for the assignment
	Roster_identifier string `json:"roster_identifier"` // Roster identifier of the student
	Submission_timestamp string `json:"submission_timestamp"` // Timestamp of the student's assignment submission
}

// GeneratedType_Actions_variable represents the GeneratedType_Actions_variable schema from the OpenAPI specification
type GeneratedType_Actions_variable struct {
	Created_at string `json:"created_at"` // The date and time at which the variable was created, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Name string `json:"name"` // The name of the variable.
	Updated_at string `json:"updated_at"` // The date and time at which the variable was last updated, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Value string `json:"value"` // The value of the variable.
}

// GeneratedType_Content_symlink represents the GeneratedType_Content_symlink schema from the OpenAPI specification
type GeneratedType_Content_symlink struct {
	Git_url string `json:"git_url"`
	Sha string `json:"sha"`
	TypeField string `json:"type"`
	Url string `json:"url"`
	Links map[string]interface{} `json:"_links"`
	Size int `json:"size"`
	Html_url string `json:"html_url"`
	Path string `json:"path"`
	Name string `json:"name"`
	Target string `json:"target"`
	Download_url string `json:"download_url"`
}

// GeneratedType_Scim_error represents the GeneratedType_Scim_error schema from the OpenAPI specification
type GeneratedType_Scim_error struct {
	Message string `json:"message,omitempty"`
	Schemas []string `json:"schemas,omitempty"`
	Scimtype string `json:"scimType,omitempty"`
	Status int `json:"status,omitempty"`
	Detail string `json:"detail,omitempty"`
	Documentation_url string `json:"documentation_url,omitempty"`
}

// GeneratedType_Codespaces_org_secret represents the GeneratedType_Codespaces_org_secret schema from the OpenAPI specification
type GeneratedType_Codespaces_org_secret struct {
	Visibility string `json:"visibility"` // The type of repositories in the organization that the secret is visible to
	Created_at string `json:"created_at"` // The date and time at which the secret was created, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Name string `json:"name"` // The name of the secret
	Selected_repositories_url string `json:"selected_repositories_url,omitempty"` // The API URL at which the list of repositories this secret is visible to can be retrieved
	Updated_at string `json:"updated_at"` // The date and time at which the secret was created, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
}

// GeneratedType_Org_private_registry_configuration_with_selected_repositories represents the GeneratedType_Org_private_registry_configuration_with_selected_repositories schema from the OpenAPI specification
type GeneratedType_Org_private_registry_configuration_with_selected_repositories struct {
	Username string `json:"username,omitempty"` // The username to use when authenticating with the private registry.
	Visibility string `json:"visibility"` // Which type of organization repositories have access to the private registry. `selected` means only the repositories specified by `selected_repository_ids` can access the private registry.
	Created_at string `json:"created_at"`
	Name string `json:"name"` // The name of the private registry configuration.
	Registry_type string `json:"registry_type"` // The registry type.
	Selected_repository_ids []int `json:"selected_repository_ids,omitempty"` // An array of repository IDs that can access the organization private registry when `visibility` is set to `selected`.
	Updated_at string `json:"updated_at"`
}

// GeneratedType_Webhook_marketplace_purchase_changed represents the GeneratedType_Webhook_marketplace_purchase_changed schema from the OpenAPI specification
type GeneratedType_Webhook_marketplace_purchase_changed struct {
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Previous_marketplace_purchase map[string]interface{} `json:"previous_marketplace_purchase,omitempty"`
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Action string `json:"action"`
	Effective_date string `json:"effective_date"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Marketplace_purchase Webhooksmarketplacepurchase `json:"marketplace_purchase"`
}

// GeneratedType_Webhook_pull_request_review_thread_unresolved represents the GeneratedType_Webhook_pull_request_review_thread_unresolved schema from the OpenAPI specification
type GeneratedType_Webhook_pull_request_review_thread_unresolved struct {
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Thread map[string]interface{} `json:"thread"`
	Action string `json:"action"`
}

// GeneratedType_Short_blob represents the GeneratedType_Short_blob schema from the OpenAPI specification
type GeneratedType_Short_blob struct {
	Sha string `json:"sha"`
	Url string `json:"url"`
}

// GeneratedType_Billing_usage_report_user represents the GeneratedType_Billing_usage_report_user schema from the OpenAPI specification
type GeneratedType_Billing_usage_report_user struct {
	Usageitems []map[string]interface{} `json:"usageItems,omitempty"`
}

// GeneratedType_Hook_response represents the GeneratedType_Hook_response schema from the OpenAPI specification
type GeneratedType_Hook_response struct {
	Message string `json:"message"`
	Status string `json:"status"`
	Code int `json:"code"`
}

// GeneratedType_Actions_public_key represents the GeneratedType_Actions_public_key schema from the OpenAPI specification
type GeneratedType_Actions_public_key struct {
	Key string `json:"key"` // The Base64 encoded public key.
	Key_id string `json:"key_id"` // The identifier for the key.
	Title string `json:"title,omitempty"`
	Url string `json:"url,omitempty"`
	Created_at string `json:"created_at,omitempty"`
	Id int `json:"id,omitempty"`
}

// GeneratedType_Enterprise_team represents the GeneratedType_Enterprise_team schema from the OpenAPI specification
type GeneratedType_Enterprise_team struct {
	Group_id string `json:"group_id,omitempty"`
	Updated_at string `json:"updated_at"`
	Html_url string `json:"html_url"`
	Id int64 `json:"id"`
	Organization_selection_type string `json:"organization_selection_type,omitempty"`
	Url string `json:"url"`
	Created_at string `json:"created_at"`
	Group_name string `json:"group_name,omitempty"`
	Name string `json:"name"`
	Description string `json:"description,omitempty"`
	Members_url string `json:"members_url"`
	Slug string `json:"slug"`
	Sync_to_organizations string `json:"sync_to_organizations,omitempty"`
}

// Webhookspullrequest5 represents the Webhookspullrequest5 schema from the OpenAPI specification
type Webhookspullrequest5 struct {
	Mergeable bool `json:"mergeable,omitempty"`
	Closed_at string `json:"closed_at"`
	Title string `json:"title"` // The title of the pull request.
	Links map[string]interface{} `json:"_links"`
	Merge_commit_sha string `json:"merge_commit_sha"`
	Locked bool `json:"locked"`
	Review_comment_url string `json:"review_comment_url"`
	Patch_url string `json:"patch_url"`
	Statuses_url string `json:"statuses_url"`
	Base map[string]interface{} `json:"base"`
	Mergeable_state string `json:"mergeable_state,omitempty"`
	Created_at string `json:"created_at"`
	Rebaseable bool `json:"rebaseable,omitempty"`
	Active_lock_reason string `json:"active_lock_reason"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Body string `json:"body"`
	Maintainer_can_modify bool `json:"maintainer_can_modify,omitempty"` // Indicates whether maintainers can modify the pull request.
	Commits_url string `json:"commits_url"`
	Requested_reviewers []interface{} `json:"requested_reviewers"`
	Html_url string `json:"html_url"`
	Review_comments int `json:"review_comments,omitempty"`
	Id int `json:"id"`
	User map[string]interface{} `json:"user"`
	Additions int `json:"additions,omitempty"`
	Url string `json:"url"`
	Draft bool `json:"draft"` // Indicates whether or not the pull request is a draft.
	Merged_by map[string]interface{} `json:"merged_by,omitempty"`
	Comments int `json:"comments,omitempty"`
	Node_id string `json:"node_id"`
	Review_comments_url string `json:"review_comments_url"`
	Number int `json:"number"` // Number uniquely identifying the pull request within its repository.
	Diff_url string `json:"diff_url"`
	Head map[string]interface{} `json:"head"`
	Labels []map[string]interface{} `json:"labels"`
	Updated_at string `json:"updated_at"`
	Milestone map[string]interface{} `json:"milestone"` // A collection of related issues and pull requests.
	Auto_merge map[string]interface{} `json:"auto_merge"` // The status of auto merging a pull request.
	Deletions int `json:"deletions,omitempty"`
	Requested_teams []map[string]interface{} `json:"requested_teams"`
	Comments_url string `json:"comments_url"`
	Changed_files int `json:"changed_files,omitempty"`
	Issue_url string `json:"issue_url"`
	Merged_at string `json:"merged_at"`
	Commits int `json:"commits,omitempty"`
	Merged bool `json:"merged,omitempty"`
	Assignee map[string]interface{} `json:"assignee"`
	State string `json:"state"` // State of this Pull Request. Either `open` or `closed`.
	Assignees []map[string]interface{} `json:"assignees"`
}

// GeneratedType_Repository_ruleset_conditions_repository_id_target represents the GeneratedType_Repository_ruleset_conditions_repository_id_target schema from the OpenAPI specification
type GeneratedType_Repository_ruleset_conditions_repository_id_target struct {
	Repository_id map[string]interface{} `json:"repository_id"`
}

// GeneratedType_Webhook_check_run_created_form_encoded represents the GeneratedType_Webhook_check_run_created_form_encoded schema from the OpenAPI specification
type GeneratedType_Webhook_check_run_created_form_encoded struct {
	Payload string `json:"payload"` // A URL-encoded string of the check_run.created JSON payload. The decoded payload is a JSON object.
}

// GeneratedType_Project_collaborator_permission represents the GeneratedType_Project_collaborator_permission schema from the OpenAPI specification
type GeneratedType_Project_collaborator_permission struct {
	User GeneratedType_Nullable_simple_user `json:"user"` // A GitHub user.
	Permission string `json:"permission"`
}

// GeneratedType_Removed_from_project_issue_event represents the GeneratedType_Removed_from_project_issue_event schema from the OpenAPI specification
type GeneratedType_Removed_from_project_issue_event struct {
	Created_at string `json:"created_at"`
	Event string `json:"event"`
	Commit_url string `json:"commit_url"`
	Node_id string `json:"node_id"`
	Performed_via_github_app GeneratedType_Nullable_integration `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Url string `json:"url"`
	Actor GeneratedType_Simple_user `json:"actor"` // A GitHub user.
	Commit_id string `json:"commit_id"`
	Id int `json:"id"`
	Project_card map[string]interface{} `json:"project_card,omitempty"`
}

// GeneratedType_Organization_simple represents the GeneratedType_Organization_simple schema from the OpenAPI specification
type GeneratedType_Organization_simple struct {
	Url string `json:"url"`
	Avatar_url string `json:"avatar_url"`
	Id int `json:"id"`
	Login string `json:"login"`
	Description string `json:"description"`
	Issues_url string `json:"issues_url"`
	Members_url string `json:"members_url"`
	Node_id string `json:"node_id"`
	Public_members_url string `json:"public_members_url"`
	Repos_url string `json:"repos_url"`
	Events_url string `json:"events_url"`
	Hooks_url string `json:"hooks_url"`
}

// Integration represents the Integration schema from the OpenAPI specification
type Integration struct {
	Description string `json:"description"`
	External_url string `json:"external_url"`
	Name string `json:"name"` // The name of the GitHub app
	Owner interface{} `json:"owner"`
	Slug string `json:"slug,omitempty"` // The slug name of the GitHub app
	Client_id string `json:"client_id,omitempty"`
	Events []string `json:"events"` // The list of events for the GitHub app. Note that the `installation_target`, `security_advisory`, and `meta` events are not included because they are global events and not specific to an installation.
	Node_id string `json:"node_id"`
	Permissions map[string]interface{} `json:"permissions"` // The set of permissions for the GitHub app
	Updated_at string `json:"updated_at"`
	Created_at string `json:"created_at"`
	Html_url string `json:"html_url"`
	Id int `json:"id"` // Unique identifier of the GitHub app
	Installations_count int `json:"installations_count,omitempty"` // The number of installations associated with the GitHub app. Only returned when the integration is requesting details about itself.
}

// GeneratedType_Nullable_git_user represents the GeneratedType_Nullable_git_user schema from the OpenAPI specification
type GeneratedType_Nullable_git_user struct {
	Date string `json:"date,omitempty"`
	Email string `json:"email,omitempty"`
	Name string `json:"name,omitempty"`
}

// GeneratedType_Webhook_projects_v2_project_closed represents the GeneratedType_Webhook_projects_v2_project_closed schema from the OpenAPI specification
type GeneratedType_Webhook_projects_v2_project_closed struct {
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Projects_v2 GeneratedType_Projects_v2 `json:"projects_v2"` // A projects v2 project
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// Import represents the Import schema from the OpenAPI specification
type Import struct {
	Message string `json:"message,omitempty"`
	Tfvc_project string `json:"tfvc_project,omitempty"`
	Error_message string `json:"error_message,omitempty"`
	Use_lfs bool `json:"use_lfs,omitempty"`
	Push_percent int `json:"push_percent,omitempty"`
	Authors_url string `json:"authors_url"`
	Project_choices []map[string]interface{} `json:"project_choices,omitempty"`
	Import_percent int `json:"import_percent,omitempty"`
	Url string `json:"url"`
	Status string `json:"status"`
	Status_text string `json:"status_text,omitempty"`
	Svn_root string `json:"svn_root,omitempty"`
	Vcs_url string `json:"vcs_url"` // The URL of the originating repository.
	Commit_count int `json:"commit_count,omitempty"`
	Repository_url string `json:"repository_url"`
	Vcs string `json:"vcs"`
	Failed_step string `json:"failed_step,omitempty"`
	Large_files_count int `json:"large_files_count,omitempty"`
	Html_url string `json:"html_url"`
	Large_files_size int `json:"large_files_size,omitempty"`
	Svc_root string `json:"svc_root,omitempty"`
	Authors_count int `json:"authors_count,omitempty"`
	Has_large_files bool `json:"has_large_files,omitempty"`
}

// Webhooksreview represents the Webhooksreview schema from the OpenAPI specification
type Webhooksreview struct {
	Links map[string]interface{} `json:"_links"`
	Commit_id string `json:"commit_id"` // A commit SHA for the review.
	State string `json:"state"`
	Body string `json:"body"` // The text of the review.
	Pull_request_url string `json:"pull_request_url"`
	Submitted_at string `json:"submitted_at"`
	Node_id string `json:"node_id"`
	User map[string]interface{} `json:"user"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Html_url string `json:"html_url"`
	Id int `json:"id"` // Unique identifier of the review
}

// GeneratedType_Code_scanning_alert_rule represents the GeneratedType_Code_scanning_alert_rule schema from the OpenAPI specification
type GeneratedType_Code_scanning_alert_rule struct {
	Description string `json:"description,omitempty"` // A short description of the rule used to detect the alert.
	Help_uri string `json:"help_uri,omitempty"` // A link to the documentation for the rule used to detect the alert.
	Name string `json:"name,omitempty"` // The name of the rule used to detect the alert.
	Security_severity_level string `json:"security_severity_level,omitempty"` // The security severity of the alert.
	Full_description string `json:"full_description,omitempty"` // A description of the rule used to detect the alert.
	Help string `json:"help,omitempty"` // Detailed documentation for the rule as GitHub Flavored Markdown.
	Id string `json:"id,omitempty"` // A unique identifier for the rule used to detect the alert.
	Severity string `json:"severity,omitempty"` // The severity of the alert.
	Tags []string `json:"tags,omitempty"` // A set of tags applicable for the rule.
}

// GeneratedType_Webhook_custom_property_created represents the GeneratedType_Webhook_custom_property_created schema from the OpenAPI specification
type GeneratedType_Webhook_custom_property_created struct {
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Definition GeneratedType_Custom_property `json:"definition"` // Custom property defined on an organization
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType_Branch_with_protection represents the GeneratedType_Branch_with_protection schema from the OpenAPI specification
type GeneratedType_Branch_with_protection struct {
	Links map[string]interface{} `json:"_links"`
	Commit Commit `json:"commit"` // Commit
	Name string `json:"name"`
	Pattern string `json:"pattern,omitempty"`
	Protected bool `json:"protected"`
	Protection GeneratedType_Branch_protection `json:"protection"` // Branch Protection
	Protection_url string `json:"protection_url"`
	Required_approving_review_count int `json:"required_approving_review_count,omitempty"`
}

// GeneratedType_Webhook_installation_new_permissions_accepted represents the GeneratedType_Webhook_installation_new_permissions_accepted schema from the OpenAPI specification
type GeneratedType_Webhook_installation_new_permissions_accepted struct {
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation Installation `json:"installation"` // Installation
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repositories []map[string]interface{} `json:"repositories,omitempty"` // An array of repository objects that the installation can access.
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Requester interface{} `json:"requester,omitempty"`
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType_Packages_billing_usage represents the GeneratedType_Packages_billing_usage schema from the OpenAPI specification
type GeneratedType_Packages_billing_usage struct {
	Included_gigabytes_bandwidth int `json:"included_gigabytes_bandwidth"` // Free storage space (GB) for GitHub Packages.
	Total_gigabytes_bandwidth_used int `json:"total_gigabytes_bandwidth_used"` // Sum of the free and paid storage space (GB) for GitHuub Packages.
	Total_paid_gigabytes_bandwidth_used int `json:"total_paid_gigabytes_bandwidth_used"` // Total paid storage space (GB) for GitHuub Packages.
}

// GeneratedType_Simple_commit_status represents the GeneratedType_Simple_commit_status schema from the OpenAPI specification
type GeneratedType_Simple_commit_status struct {
	Url string `json:"url"`
	Avatar_url string `json:"avatar_url"`
	Node_id string `json:"node_id"`
	Context string `json:"context"`
	Created_at string `json:"created_at"`
	Required bool `json:"required,omitempty"`
	State string `json:"state"`
	Description string `json:"description"`
	Id int `json:"id"`
	Target_url string `json:"target_url"`
	Updated_at string `json:"updated_at"`
}

// GeneratedType_Secret_scanning_location_wiki_commit represents the GeneratedType_Secret_scanning_location_wiki_commit schema from the OpenAPI specification
type GeneratedType_Secret_scanning_location_wiki_commit struct {
	Start_line float64 `json:"start_line"` // Line number at which the secret starts in the file
	Commit_url string `json:"commit_url"` // The GitHub URL to get the associated wiki commit
	End_column float64 `json:"end_column"` // The column at which the secret ends within the end line when the file is interpreted as 8-bit ASCII.
	End_line float64 `json:"end_line"` // Line number at which the secret ends in the file
	Path string `json:"path"` // The file path of the wiki page
	Blob_sha string `json:"blob_sha"` // SHA-1 hash ID of the associated blob
	Commit_sha string `json:"commit_sha"` // SHA-1 hash ID of the associated commit
	Page_url string `json:"page_url"` // The GitHub URL to get the associated wiki page
	Start_column float64 `json:"start_column"` // The column at which the secret starts within the start line when the file is interpreted as 8-bit ASCII.
}

// GeneratedType_Webhook_pull_request_unlabeled represents the GeneratedType_Webhook_pull_request_unlabeled schema from the OpenAPI specification
type GeneratedType_Webhook_pull_request_unlabeled struct {
	Action string `json:"action"`
	Label Webhookslabel `json:"label,omitempty"`
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Number int `json:"number"` // The pull request number.
}

// GeneratedType_Workflow_run represents the GeneratedType_Workflow_run schema from the OpenAPI specification
type GeneratedType_Workflow_run struct {
	Event string `json:"event"`
	Jobs_url string `json:"jobs_url"` // The URL to the jobs for the workflow run.
	Workflow_url string `json:"workflow_url"` // The URL to the workflow.
	Cancel_url string `json:"cancel_url"` // The URL to cancel the workflow run.
	Workflow_id int `json:"workflow_id"` // The ID of the parent workflow.
	Head_commit GeneratedType_Nullable_simple_commit `json:"head_commit"` // A commit.
	Url string `json:"url"` // The URL to the workflow run.
	Check_suite_node_id string `json:"check_suite_node_id,omitempty"` // The node ID of the associated check suite.
	Head_repository GeneratedType_Minimal_repository `json:"head_repository"` // Minimal Repository
	Head_repository_id int `json:"head_repository_id,omitempty"`
	Previous_attempt_url string `json:"previous_attempt_url,omitempty"` // The URL to the previous attempted run of this workflow, if one exists.
	Node_id string `json:"node_id"`
	Referenced_workflows []GeneratedType_Referenced_workflow `json:"referenced_workflows,omitempty"`
	Rerun_url string `json:"rerun_url"` // The URL to rerun the workflow run.
	Conclusion string `json:"conclusion"`
	Display_title string `json:"display_title"` // The event-specific title associated with the run or the run-name if set, or the value of `run-name` if it is set in the workflow.
	Run_attempt int `json:"run_attempt,omitempty"` // Attempt number of the run, 1 for first attempt and higher if the workflow was re-run.
	Artifacts_url string `json:"artifacts_url"` // The URL to the artifacts for the workflow run.
	Actor GeneratedType_Simple_user `json:"actor,omitempty"` // A GitHub user.
	Name string `json:"name,omitempty"` // The name of the workflow run.
	Updated_at string `json:"updated_at"`
	Run_number int `json:"run_number"` // The auto incrementing run number for the workflow run.
	Check_suite_id int `json:"check_suite_id,omitempty"` // The ID of the associated check suite.
	Repository GeneratedType_Minimal_repository `json:"repository"` // Minimal Repository
	Triggering_actor GeneratedType_Simple_user `json:"triggering_actor,omitempty"` // A GitHub user.
	Pull_requests []GeneratedType_Pull_request_minimal `json:"pull_requests"` // Pull requests that are open with a `head_sha` or `head_branch` that matches the workflow run. The returned pull requests do not necessarily indicate pull requests that triggered the run.
	Head_sha string `json:"head_sha"` // The SHA of the head commit that points to the version of the workflow being run.
	Logs_url string `json:"logs_url"` // The URL to download the logs for the workflow run.
	Id int `json:"id"` // The ID of the workflow run.
	Created_at string `json:"created_at"`
	Check_suite_url string `json:"check_suite_url"` // The URL to the associated check suite.
	Path string `json:"path"` // The full path of the workflow
	Status string `json:"status"`
	Html_url string `json:"html_url"`
	Run_started_at string `json:"run_started_at,omitempty"` // The start time of the latest run. Resets on re-run.
	Head_branch string `json:"head_branch"`
}

// GeneratedType_Cvss_severities represents the GeneratedType_Cvss_severities schema from the OpenAPI specification
type GeneratedType_Cvss_severities struct {
	Cvss_v4 map[string]interface{} `json:"cvss_v4,omitempty"`
	Cvss_v3 map[string]interface{} `json:"cvss_v3,omitempty"`
}

// GeneratedType_Secret_scanning_location_commit represents the GeneratedType_Secret_scanning_location_commit schema from the OpenAPI specification
type GeneratedType_Secret_scanning_location_commit struct {
	Commit_sha string `json:"commit_sha"` // SHA-1 hash ID of the associated commit
	End_column float64 `json:"end_column"` // The column at which the secret ends within the end line when the file is interpreted as 8BIT ASCII
	End_line float64 `json:"end_line"` // Line number at which the secret ends in the file
	Start_line float64 `json:"start_line"` // Line number at which the secret starts in the file
	Start_column float64 `json:"start_column"` // The column at which the secret starts within the start line when the file is interpreted as 8BIT ASCII
	Blob_sha string `json:"blob_sha"` // SHA-1 hash ID of the associated blob
	Commit_url string `json:"commit_url"` // The API URL to get the associated commit resource
	Path string `json:"path"` // The file path in the repository
	Blob_url string `json:"blob_url"` // The API URL to get the associated blob resource
}

// GeneratedType_Webhook_sub_issues_parent_issue_added represents the GeneratedType_Webhook_sub_issues_parent_issue_added schema from the OpenAPI specification
type GeneratedType_Webhook_sub_issues_parent_issue_added struct {
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Parent_issue_id float64 `json:"parent_issue_id"` // The ID of the parent issue.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Sub_issue Issue `json:"sub_issue"` // Issues are a great way to keep track of tasks, enhancements, and bugs for your projects.
	Sub_issue_id float64 `json:"sub_issue_id"` // The ID of the sub-issue.
	Parent_issue_repo Repository `json:"parent_issue_repo"` // A repository on GitHub.
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Parent_issue Issue `json:"parent_issue"` // Issues are a great way to keep track of tasks, enhancements, and bugs for your projects.
}

// GeneratedType_Organization_full represents the GeneratedType_Organization_full schema from the OpenAPI specification
type GeneratedType_Organization_full struct {
	Members_can_create_public_pages bool `json:"members_can_create_public_pages,omitempty"`
	Members_allowed_repository_creation_type string `json:"members_allowed_repository_creation_type,omitempty"`
	Members_can_delete_issues bool `json:"members_can_delete_issues,omitempty"`
	Public_repos int `json:"public_repos"`
	Hooks_url string `json:"hooks_url"`
	Private_gists int `json:"private_gists,omitempty"`
	Members_can_create_repositories bool `json:"members_can_create_repositories,omitempty"`
	Owned_private_repos int `json:"owned_private_repos,omitempty"`
	Default_repository_branch string `json:"default_repository_branch,omitempty"` // The default branch for repositories created in this organization.
	Disk_usage int `json:"disk_usage,omitempty"`
	Members_can_delete_repositories bool `json:"members_can_delete_repositories,omitempty"`
	Advanced_security_enabled_for_new_repositories bool `json:"advanced_security_enabled_for_new_repositories,omitempty"` // **Endpoint closing down notice.** Please use [code security configurations](https://docs.github.com/rest/code-security/configurations) instead. Whether GitHub Advanced Security is enabled for new repositories and repositories transferred to this organization. This field is only visible to organization owners or members of a team with the security manager role.
	Members_can_fork_private_repositories bool `json:"members_can_fork_private_repositories,omitempty"`
	Email string `json:"email,omitempty"`
	Is_verified bool `json:"is_verified,omitempty"`
	Created_at string `json:"created_at"`
	Issues_url string `json:"issues_url"`
	Members_can_create_private_repositories bool `json:"members_can_create_private_repositories,omitempty"`
	Avatar_url string `json:"avatar_url"`
	Twitter_username string `json:"twitter_username,omitempty"`
	Public_members_url string `json:"public_members_url"`
	Members_can_create_public_repositories bool `json:"members_can_create_public_repositories,omitempty"`
	Secret_scanning_enabled_for_new_repositories bool `json:"secret_scanning_enabled_for_new_repositories,omitempty"` // **Endpoint closing down notice.** Please use [code security configurations](https://docs.github.com/rest/code-security/configurations) instead. Whether secret scanning is automatically enabled for new repositories and repositories transferred to this organization. This field is only visible to organization owners or members of a team with the security manager role.
	Secret_scanning_push_protection_custom_link string `json:"secret_scanning_push_protection_custom_link,omitempty"` // An optional URL string to display to contributors who are blocked from pushing a secret.
	Two_factor_requirement_enabled bool `json:"two_factor_requirement_enabled,omitempty"`
	Following int `json:"following"`
	Has_organization_projects bool `json:"has_organization_projects"`
	Has_repository_projects bool `json:"has_repository_projects"`
	Members_can_create_private_pages bool `json:"members_can_create_private_pages,omitempty"`
	TypeField string `json:"type"`
	Blog string `json:"blog,omitempty"`
	Dependency_graph_enabled_for_new_repositories bool `json:"dependency_graph_enabled_for_new_repositories,omitempty"` // **Endpoint closing down notice.** Please use [code security configurations](https://docs.github.com/rest/code-security/configurations) instead. Whether dependency graph is automatically enabled for new repositories and repositories transferred to this organization. This field is only visible to organization owners or members of a team with the security manager role.
	Total_private_repos int `json:"total_private_repos,omitempty"`
	Description string `json:"description"`
	Deploy_keys_enabled_for_repositories bool `json:"deploy_keys_enabled_for_repositories,omitempty"` // Controls whether or not deploy keys may be added and used for repositories in the organization.
	Members_can_create_pages bool `json:"members_can_create_pages,omitempty"`
	Members_can_invite_outside_collaborators bool `json:"members_can_invite_outside_collaborators,omitempty"`
	Followers int `json:"followers"`
	Plan map[string]interface{} `json:"plan,omitempty"`
	Secret_scanning_push_protection_custom_link_enabled bool `json:"secret_scanning_push_protection_custom_link_enabled,omitempty"` // Whether a custom link is shown to contributors who are blocked from pushing a secret by push protection.
	Members_url string `json:"members_url"`
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"`
	Readers_can_create_discussions bool `json:"readers_can_create_discussions,omitempty"`
	Id int `json:"id"`
	Members_can_create_teams bool `json:"members_can_create_teams,omitempty"`
	Events_url string `json:"events_url"`
	Public_gists int `json:"public_gists"`
	Members_can_create_internal_repositories bool `json:"members_can_create_internal_repositories,omitempty"`
	Node_id string `json:"node_id"`
	Dependabot_alerts_enabled_for_new_repositories bool `json:"dependabot_alerts_enabled_for_new_repositories,omitempty"` // **Endpoint closing down notice.** Please use [code security configurations](https://docs.github.com/rest/code-security/configurations) instead. Whether Dependabot alerts are automatically enabled for new repositories and repositories transferred to this organization. This field is only visible to organization owners or members of a team with the security manager role.
	Archived_at string `json:"archived_at"`
	Collaborators int `json:"collaborators,omitempty"` // The number of collaborators on private repositories. This field may be null if the number of private repositories is over 50,000.
	Company string `json:"company,omitempty"`
	Repos_url string `json:"repos_url"`
	Members_can_change_repo_visibility bool `json:"members_can_change_repo_visibility,omitempty"`
	Updated_at string `json:"updated_at"`
	Display_commenter_full_name_setting_enabled bool `json:"display_commenter_full_name_setting_enabled,omitempty"`
	Html_url string `json:"html_url"`
	Location string `json:"location,omitempty"`
	Login string `json:"login"`
	Secret_scanning_push_protection_enabled_for_new_repositories bool `json:"secret_scanning_push_protection_enabled_for_new_repositories,omitempty"` // **Endpoint closing down notice.** Please use [code security configurations](https://docs.github.com/rest/code-security/configurations) instead. Whether secret scanning push protection is automatically enabled for new repositories and repositories transferred to this organization. This field is only visible to organization owners or members of a team with the security manager role.
	Dependabot_security_updates_enabled_for_new_repositories bool `json:"dependabot_security_updates_enabled_for_new_repositories,omitempty"` // **Endpoint closing down notice.** Please use [code security configurations](https://docs.github.com/rest/code-security/configurations) instead. Whether Dependabot security updates are automatically enabled for new repositories and repositories transferred to this organization. This field is only visible to organization owners or members of a team with the security manager role.
	Billing_email string `json:"billing_email,omitempty"`
	Name string `json:"name,omitempty"`
	Members_can_view_dependency_insights bool `json:"members_can_view_dependency_insights,omitempty"`
	Default_repository_permission string `json:"default_repository_permission,omitempty"`
	Url string `json:"url"`
}

// GeneratedType_Hook_delivery represents the GeneratedType_Hook_delivery schema from the OpenAPI specification
type GeneratedType_Hook_delivery struct {
	Throttled_at string `json:"throttled_at,omitempty"` // Time when the webhook delivery was throttled.
	Status string `json:"status"` // Description of the status of the attempted delivery
	Response map[string]interface{} `json:"response"`
	Status_code int `json:"status_code"` // Status code received when delivery was made.
	Installation_id int `json:"installation_id"` // The id of the GitHub App installation associated with this event.
	Request map[string]interface{} `json:"request"`
	Url string `json:"url,omitempty"` // The URL target of the delivery.
	Duration float64 `json:"duration"` // Time spent delivering.
	Delivered_at string `json:"delivered_at"` // Time when the delivery was delivered.
	Action string `json:"action"` // The type of activity for the event that triggered the delivery.
	Event string `json:"event"` // The event that triggered the delivery.
	Guid string `json:"guid"` // Unique identifier for the event (shared with all deliveries for all webhooks that subscribe to this event).
	Redelivery bool `json:"redelivery"` // Whether the delivery is a redelivery.
	Id int `json:"id"` // Unique identifier of the delivery.
	Repository_id int `json:"repository_id"` // The id of the repository associated with this event.
}

// GeneratedType_Copilot_ide_chat represents the GeneratedType_Copilot_ide_chat schema from the OpenAPI specification
type GeneratedType_Copilot_ide_chat struct {
	Editors []map[string]interface{} `json:"editors,omitempty"`
	Total_engaged_users int `json:"total_engaged_users,omitempty"` // Total number of users who prompted Copilot Chat in the IDE.
}

// GeneratedType_Repository_rule_deletion represents the GeneratedType_Repository_rule_deletion schema from the OpenAPI specification
type GeneratedType_Repository_rule_deletion struct {
	TypeField string `json:"type"`
}

// GeneratedType_Webhook_projects_v2_project_created represents the GeneratedType_Webhook_projects_v2_project_created schema from the OpenAPI specification
type GeneratedType_Webhook_projects_v2_project_created struct {
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Projects_v2 GeneratedType_Projects_v2 `json:"projects_v2"` // A projects v2 project
}

// GeneratedType_Copilot_usage_metrics_day represents the GeneratedType_Copilot_usage_metrics_day schema from the OpenAPI specification
type GeneratedType_Copilot_usage_metrics_day struct {
	Copilot_dotcom_pull_requests GeneratedType_Copilot_dotcom_pull_requests `json:"copilot_dotcom_pull_requests,omitempty"` // Usage metrics for Copilot for pull requests.
	Copilot_ide_chat GeneratedType_Copilot_ide_chat `json:"copilot_ide_chat,omitempty"` // Usage metrics for Copilot Chat in the IDE.
	Copilot_ide_code_completions GeneratedType_Copilot_ide_code_completions `json:"copilot_ide_code_completions,omitempty"` // Usage metrics for Copilot editor code completions in the IDE.
	Date string `json:"date"` // The date for which the usage metrics are aggregated, in `YYYY-MM-DD` format.
	Total_active_users int `json:"total_active_users,omitempty"` // The total number of Copilot users with activity belonging to any Copilot feature, globally, for the given day. Includes passive activity such as receiving a code suggestion, as well as engagement activity such as accepting a code suggestion or prompting chat. Does not include authentication events. Is not limited to the individual features detailed on the endpoint.
	Total_engaged_users int `json:"total_engaged_users,omitempty"` // The total number of Copilot users who engaged with any Copilot feature, for the given day. Examples include but are not limited to accepting a code suggestion, prompting Copilot chat, or triggering a PR Summary. Does not include authentication events. Is not limited to the individual features detailed on the endpoint.
	Copilot_dotcom_chat GeneratedType_Copilot_dotcom_chat `json:"copilot_dotcom_chat,omitempty"` // Usage metrics for Copilot Chat in GitHub.com
}

// GeneratedType_Code_scanning_analysis_tool represents the GeneratedType_Code_scanning_analysis_tool schema from the OpenAPI specification
type GeneratedType_Code_scanning_analysis_tool struct {
	Version string `json:"version,omitempty"` // The version of the tool used to generate the code scanning analysis.
	Guid string `json:"guid,omitempty"` // The GUID of the tool used to generate the code scanning analysis, if provided in the uploaded SARIF data.
	Name string `json:"name,omitempty"` // The name of the tool used to generate the code scanning analysis.
}

// GeneratedType_Webhook_repository_publicized represents the GeneratedType_Webhook_repository_publicized schema from the OpenAPI specification
type GeneratedType_Webhook_repository_publicized struct {
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType_Webhook_workflow_job_in_progress represents the GeneratedType_Webhook_workflow_job_in_progress schema from the OpenAPI specification
type GeneratedType_Webhook_workflow_job_in_progress struct {
	Deployment Deployment `json:"deployment,omitempty"` // A request for a specific ref(branch,sha,tag) to be deployed
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Workflow_job interface{} `json:"workflow_job"`
	Action string `json:"action"`
}

// GeneratedType_Gpg_key represents the GeneratedType_Gpg_key schema from the OpenAPI specification
type GeneratedType_Gpg_key struct {
	Primary_key_id int `json:"primary_key_id"`
	Emails []map[string]interface{} `json:"emails"`
	Subkeys []map[string]interface{} `json:"subkeys"`
	Can_certify bool `json:"can_certify"`
	Can_encrypt_comms bool `json:"can_encrypt_comms"`
	Public_key string `json:"public_key"`
	Name string `json:"name,omitempty"`
	Revoked bool `json:"revoked"`
	Expires_at string `json:"expires_at"`
	Can_encrypt_storage bool `json:"can_encrypt_storage"`
	Created_at string `json:"created_at"`
	Key_id string `json:"key_id"`
	Raw_key string `json:"raw_key"`
	Can_sign bool `json:"can_sign"`
	Id int64 `json:"id"`
}

// GeneratedType_Commit_search_result_item represents the GeneratedType_Commit_search_result_item schema from the OpenAPI specification
type GeneratedType_Commit_search_result_item struct {
	Author GeneratedType_Nullable_simple_user `json:"author"` // A GitHub user.
	Committer GeneratedType_Nullable_git_user `json:"committer"` // Metaproperties for Git author/committer information.
	Text_matches []map[string]interface{} `json:"text_matches,omitempty"`
	Parents []map[string]interface{} `json:"parents"`
	Url string `json:"url"`
	Commit map[string]interface{} `json:"commit"`
	Html_url string `json:"html_url"`
	Repository GeneratedType_Minimal_repository `json:"repository"` // Minimal Repository
	Score float64 `json:"score"`
	Sha string `json:"sha"`
	Comments_url string `json:"comments_url"`
	Node_id string `json:"node_id"`
}

// GeneratedType_Organization_programmatic_access_grant represents the GeneratedType_Organization_programmatic_access_grant schema from the OpenAPI specification
type GeneratedType_Organization_programmatic_access_grant struct {
	Repository_selection string `json:"repository_selection"` // Type of repository selection requested.
	Token_expires_at string `json:"token_expires_at"` // Date and time when the associated fine-grained personal access token expires.
	Id int `json:"id"` // Unique identifier of the fine-grained personal access token grant. The `pat_id` used to get details about an approved fine-grained personal access token.
	Token_expired bool `json:"token_expired"` // Whether the associated fine-grained personal access token has expired.
	Token_last_used_at string `json:"token_last_used_at"` // Date and time when the associated fine-grained personal access token was last used for authentication.
	Owner GeneratedType_Simple_user `json:"owner"` // A GitHub user.
	Token_id int `json:"token_id"` // Unique identifier of the user's token. This field can also be found in audit log events and the organization's settings for their PAT grants.
	Access_granted_at string `json:"access_granted_at"` // Date and time when the fine-grained personal access token was approved to access the organization.
	Permissions map[string]interface{} `json:"permissions"` // Permissions requested, categorized by type of permission.
	Repositories_url string `json:"repositories_url"` // URL to the list of repositories the fine-grained personal access token can access. Only follow when `repository_selection` is `subset`.
	Token_name string `json:"token_name"` // The name given to the user's token. This field can also be found in an organization's settings page for Active Tokens.
}

// GeneratedType_Webhook_repository_vulnerability_alert_dismiss represents the GeneratedType_Webhook_repository_vulnerability_alert_dismiss schema from the OpenAPI specification
type GeneratedType_Webhook_repository_vulnerability_alert_dismiss struct {
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Alert map[string]interface{} `json:"alert"` // The security alert of the vulnerable dependency.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType_Merged_upstream represents the GeneratedType_Merged_upstream schema from the OpenAPI specification
type GeneratedType_Merged_upstream struct {
	Base_branch string `json:"base_branch,omitempty"`
	Merge_type string `json:"merge_type,omitempty"`
	Message string `json:"message,omitempty"`
}

// GeneratedType_Nullable_team_simple represents the GeneratedType_Nullable_team_simple schema from the OpenAPI specification
type GeneratedType_Nullable_team_simple struct {
	Ldap_dn string `json:"ldap_dn,omitempty"` // Distinguished Name (DN) that team maps to within LDAP environment
	Permission string `json:"permission"` // Permission that the team will have for its repositories
	Slug string `json:"slug"`
	Description string `json:"description"` // Description of the team
	Html_url string `json:"html_url"`
	Name string `json:"name"` // Name of the team
	Privacy string `json:"privacy,omitempty"` // The level of privacy this team should have
	Id int `json:"id"` // Unique identifier of the team
	Node_id string `json:"node_id"`
	Notification_setting string `json:"notification_setting,omitempty"` // The notification setting the team has set
	Url string `json:"url"` // URL for the team
	Members_url string `json:"members_url"`
	Repositories_url string `json:"repositories_url"`
}

// GeneratedType_Diff_entry represents the GeneratedType_Diff_entry schema from the OpenAPI specification
type GeneratedType_Diff_entry struct {
	Contents_url string `json:"contents_url"`
	Filename string `json:"filename"`
	Patch string `json:"patch,omitempty"`
	Raw_url string `json:"raw_url"`
	Previous_filename string `json:"previous_filename,omitempty"`
	Sha string `json:"sha"`
	Additions int `json:"additions"`
	Deletions int `json:"deletions"`
	Status string `json:"status"`
	Blob_url string `json:"blob_url"`
	Changes int `json:"changes"`
}

// GeneratedType_Pages_https_certificate represents the GeneratedType_Pages_https_certificate schema from the OpenAPI specification
type GeneratedType_Pages_https_certificate struct {
	Expires_at string `json:"expires_at,omitempty"`
	State string `json:"state"`
	Description string `json:"description"`
	Domains []string `json:"domains"` // Array of the domain set and its alternate name (if it is configured)
}

// GeneratedType_Webhook_repository_archived represents the GeneratedType_Webhook_repository_archived schema from the OpenAPI specification
type GeneratedType_Webhook_repository_archived struct {
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType_Webhook_issues_locked represents the GeneratedType_Webhook_issues_locked schema from the OpenAPI specification
type GeneratedType_Webhook_issues_locked struct {
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType_Team_project represents the GeneratedType_Team_project schema from the OpenAPI specification
type GeneratedType_Team_project struct {
	Private bool `json:"private,omitempty"` // Whether the project is private or not. Only present when owner is an organization.
	Permissions map[string]interface{} `json:"permissions"`
	Name string `json:"name"`
	Organization_permission string `json:"organization_permission,omitempty"` // The organization permission for this project. Only present when owner is an organization.
	State string `json:"state"`
	Url string `json:"url"`
	Html_url string `json:"html_url"`
	Body string `json:"body"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
	Columns_url string `json:"columns_url"`
	Creator GeneratedType_Simple_user `json:"creator"` // A GitHub user.
	Owner_url string `json:"owner_url"`
	Number int `json:"number"`
	Id int `json:"id"`
	Node_id string `json:"node_id"`
}

// GeneratedType_Commit_comparison represents the GeneratedType_Commit_comparison schema from the OpenAPI specification
type GeneratedType_Commit_comparison struct {
	Diff_url string `json:"diff_url"`
	Files []GeneratedType_Diff_entry `json:"files,omitempty"`
	Merge_base_commit Commit `json:"merge_base_commit"` // Commit
	Permalink_url string `json:"permalink_url"`
	Status string `json:"status"`
	Base_commit Commit `json:"base_commit"` // Commit
	Behind_by int `json:"behind_by"`
	Total_commits int `json:"total_commits"`
	Commits []Commit `json:"commits"`
	Html_url string `json:"html_url"`
	Patch_url string `json:"patch_url"`
	Ahead_by int `json:"ahead_by"`
	Url string `json:"url"`
}

// GeneratedType_Code_scanning_variant_analysis represents the GeneratedType_Code_scanning_variant_analysis schema from the OpenAPI specification
type GeneratedType_Code_scanning_variant_analysis struct {
	Completed_at string `json:"completed_at,omitempty"` // The date and time at which the variant analysis was completed, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ. Will be null if the variant analysis has not yet completed or this information is not available.
	Query_pack_url string `json:"query_pack_url"` // The download url for the query pack.
	Actions_workflow_run_id int `json:"actions_workflow_run_id,omitempty"` // The GitHub Actions workflow run used to execute this variant analysis. This is only available if the workflow run has started.
	Scanned_repositories []map[string]interface{} `json:"scanned_repositories,omitempty"`
	Created_at string `json:"created_at,omitempty"` // The date and time at which the variant analysis was created, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Status string `json:"status"`
	Updated_at string `json:"updated_at,omitempty"` // The date and time at which the variant analysis was last updated, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Actor GeneratedType_Simple_user `json:"actor"` // A GitHub user.
	Controller_repo GeneratedType_Simple_repository `json:"controller_repo"` // A GitHub repository.
	Failure_reason string `json:"failure_reason,omitempty"` // The reason for a failure of the variant analysis. This is only available if the variant analysis has failed.
	Id int `json:"id"` // The ID of the variant analysis.
	Query_language string `json:"query_language"` // The language targeted by the CodeQL query
	Skipped_repositories map[string]interface{} `json:"skipped_repositories,omitempty"` // Information about repositories that were skipped from processing. This information is only available to the user that initiated the variant analysis.
}

// GeneratedType_Release_asset represents the GeneratedType_Release_asset schema from the OpenAPI specification
type GeneratedType_Release_asset struct {
	State string `json:"state"` // State of the release asset.
	Content_type string `json:"content_type"`
	Name string `json:"name"` // The file name of the asset.
	Node_id string `json:"node_id"`
	Updated_at string `json:"updated_at"`
	Uploader GeneratedType_Nullable_simple_user `json:"uploader"` // A GitHub user.
	Url string `json:"url"`
	Browser_download_url string `json:"browser_download_url"`
	Download_count int `json:"download_count"`
	Label string `json:"label"`
	Size int `json:"size"`
	Created_at string `json:"created_at"`
	Digest string `json:"digest"`
	Id int `json:"id"`
}

// GeneratedType_Webhook_pull_request_review_comment_created represents the GeneratedType_Webhook_pull_request_review_comment_created schema from the OpenAPI specification
type GeneratedType_Webhook_pull_request_review_comment_created struct {
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Comment map[string]interface{} `json:"comment"` // The [comment](https://docs.github.com/rest/pulls/comments#get-a-review-comment-for-a-pull-request) itself.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType_Gist_commit represents the GeneratedType_Gist_commit schema from the OpenAPI specification
type GeneratedType_Gist_commit struct {
	User GeneratedType_Nullable_simple_user `json:"user"` // A GitHub user.
	Version string `json:"version"`
	Change_status map[string]interface{} `json:"change_status"`
	Committed_at string `json:"committed_at"`
	Url string `json:"url"`
}

// Webhooksmembership represents the Webhooksmembership schema from the OpenAPI specification
type Webhooksmembership struct {
	Url string `json:"url"`
	User map[string]interface{} `json:"user"`
	Organization_url string `json:"organization_url"`
	Role string `json:"role"`
	State string `json:"state"`
}

// GeneratedType_Commit_activity represents the GeneratedType_Commit_activity schema from the OpenAPI specification
type GeneratedType_Commit_activity struct {
	Days []int `json:"days"`
	Total int `json:"total"`
	Week int `json:"week"`
}

// GeneratedType_Pull_request represents the GeneratedType_Pull_request schema from the OpenAPI specification
type GeneratedType_Pull_request struct {
	Changed_files int `json:"changed_files"`
	Maintainer_can_modify bool `json:"maintainer_can_modify"` // Indicates whether maintainers can modify the pull request.
	Assignees []GeneratedType_Simple_user `json:"assignees,omitempty"`
	Comments_url string `json:"comments_url"`
	Merged_by GeneratedType_Nullable_simple_user `json:"merged_by"` // A GitHub user.
	Draft bool `json:"draft,omitempty"` // Indicates whether or not the pull request is a draft.
	State string `json:"state"` // State of this Pull Request. Either `open` or `closed`.
	Links map[string]interface{} `json:"_links"`
	Merged bool `json:"merged"`
	Diff_url string `json:"diff_url"`
	Id int64 `json:"id"`
	Statuses_url string `json:"statuses_url"`
	Number int `json:"number"` // Number uniquely identifying the pull request within its repository.
	Node_id string `json:"node_id"`
	Rebaseable bool `json:"rebaseable,omitempty"`
	Updated_at string `json:"updated_at"`
	Merged_at string `json:"merged_at"`
	Milestone GeneratedType_Nullable_milestone `json:"milestone"` // A collection of related issues and pull requests.
	Title string `json:"title"` // The title of the pull request.
	Mergeable_state string `json:"mergeable_state"`
	Body string `json:"body"`
	Issue_url string `json:"issue_url"`
	Patch_url string `json:"patch_url"`
	Commits_url string `json:"commits_url"`
	Active_lock_reason string `json:"active_lock_reason,omitempty"`
	Base map[string]interface{} `json:"base"`
	Locked bool `json:"locked"`
	Review_comments_url string `json:"review_comments_url"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Commits int `json:"commits"`
	Comments int `json:"comments"`
	Requested_reviewers []GeneratedType_Simple_user `json:"requested_reviewers,omitempty"`
	Html_url string `json:"html_url"`
	Created_at string `json:"created_at"`
	Assignee GeneratedType_Nullable_simple_user `json:"assignee"` // A GitHub user.
	Deletions int `json:"deletions"`
	Additions int `json:"additions"`
	Closed_at string `json:"closed_at"`
	Labels []map[string]interface{} `json:"labels"`
	Head map[string]interface{} `json:"head"`
	Merge_commit_sha string `json:"merge_commit_sha"`
	Url string `json:"url"`
	Review_comment_url string `json:"review_comment_url"`
	Review_comments int `json:"review_comments"`
	User GeneratedType_Simple_user `json:"user"` // A GitHub user.
	Requested_teams []GeneratedType_Team_simple `json:"requested_teams,omitempty"`
	Auto_merge GeneratedType_Auto_merge `json:"auto_merge"` // The status of auto merging a pull request.
	Mergeable bool `json:"mergeable"`
}

// GeneratedType_Protected_branch represents the GeneratedType_Protected_branch schema from the OpenAPI specification
type GeneratedType_Protected_branch struct {
	Allow_deletions map[string]interface{} `json:"allow_deletions,omitempty"`
	Allow_fork_syncing map[string]interface{} `json:"allow_fork_syncing,omitempty"` // Whether users can pull changes from upstream when the branch is locked. Set to `true` to allow fork syncing. Set to `false` to prevent fork syncing.
	Block_creations map[string]interface{} `json:"block_creations,omitempty"`
	Lock_branch map[string]interface{} `json:"lock_branch,omitempty"` // Whether to set the branch as read-only. If this is true, users will not be able to push to the branch.
	Restrictions GeneratedType_Branch_restriction_policy `json:"restrictions,omitempty"` // Branch Restriction Policy
	Enforce_admins map[string]interface{} `json:"enforce_admins,omitempty"`
	Required_conversation_resolution map[string]interface{} `json:"required_conversation_resolution,omitempty"`
	Url string `json:"url"`
	Allow_force_pushes map[string]interface{} `json:"allow_force_pushes,omitempty"`
	Required_linear_history map[string]interface{} `json:"required_linear_history,omitempty"`
	Required_pull_request_reviews map[string]interface{} `json:"required_pull_request_reviews,omitempty"`
	Required_signatures map[string]interface{} `json:"required_signatures,omitempty"`
	Required_status_checks GeneratedType_Status_check_policy `json:"required_status_checks,omitempty"` // Status Check Policy
}

// GeneratedType_Team_discussion_comment represents the GeneratedType_Team_discussion_comment schema from the OpenAPI specification
type GeneratedType_Team_discussion_comment struct {
	Reactions GeneratedType_Reaction_rollup `json:"reactions,omitempty"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	Author GeneratedType_Nullable_simple_user `json:"author"` // A GitHub user.
	Body string `json:"body"` // The main text of the comment.
	Discussion_url string `json:"discussion_url"`
	Number int `json:"number"` // The unique sequence number of a team discussion comment.
	Body_version string `json:"body_version"` // The current version of the body content. If provided, this update operation will be rejected if the given version does not match the latest version on the server.
	Created_at string `json:"created_at"`
	Html_url string `json:"html_url"`
	Node_id string `json:"node_id"`
	Body_html string `json:"body_html"`
	Last_edited_at string `json:"last_edited_at"`
}

// Webhooksmilestone3 represents the Webhooksmilestone3 schema from the OpenAPI specification
type Webhooksmilestone3 struct {
	Html_url string `json:"html_url"`
	Updated_at string `json:"updated_at"`
	Labels_url string `json:"labels_url"`
	Node_id string `json:"node_id"`
	Open_issues int `json:"open_issues"`
	Closed_at string `json:"closed_at"`
	Closed_issues int `json:"closed_issues"`
	Id int `json:"id"`
	Number int `json:"number"` // The number of the milestone.
	Title string `json:"title"` // The title of the milestone.
	Created_at string `json:"created_at"`
	Due_on string `json:"due_on"`
	Creator map[string]interface{} `json:"creator"`
	Description string `json:"description"`
	State string `json:"state"` // The state of the milestone.
	Url string `json:"url"`
}

// GeneratedType_Dependabot_repository_access_details represents the GeneratedType_Dependabot_repository_access_details schema from the OpenAPI specification
type GeneratedType_Dependabot_repository_access_details struct {
	Accessible_repositories []GeneratedType_Nullable_simple_repository `json:"accessible_repositories,omitempty"`
	Default_level string `json:"default_level,omitempty"` // The default repository access level for Dependabot updates.
}

// GeneratedType_Webhook_workflow_job_waiting represents the GeneratedType_Webhook_workflow_job_waiting schema from the OpenAPI specification
type GeneratedType_Webhook_workflow_job_waiting struct {
	Deployment Deployment `json:"deployment,omitempty"` // A request for a specific ref(branch,sha,tag) to be deployed
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Workflow_job map[string]interface{} `json:"workflow_job"`
	Action string `json:"action"`
}

// GeneratedType_Community_profile represents the GeneratedType_Community_profile schema from the OpenAPI specification
type GeneratedType_Community_profile struct {
	Files map[string]interface{} `json:"files"`
	Health_percentage int `json:"health_percentage"`
	Updated_at string `json:"updated_at"`
	Content_reports_enabled bool `json:"content_reports_enabled,omitempty"`
	Description string `json:"description"`
	Documentation string `json:"documentation"`
}

// GeneratedType_Environment_approvals represents the GeneratedType_Environment_approvals schema from the OpenAPI specification
type GeneratedType_Environment_approvals struct {
	Comment string `json:"comment"` // The comment submitted with the deployment review
	Environments []map[string]interface{} `json:"environments"` // The list of environments that were approved or rejected
	State string `json:"state"` // Whether deployment to the environment(s) was approved or rejected or pending (with comments)
	User GeneratedType_Simple_user `json:"user"` // A GitHub user.
}

// Issue represents the Issue schema from the OpenAPI specification
type Issue struct {
	Assignee GeneratedType_Nullable_simple_user `json:"assignee"` // A GitHub user.
	Pull_request map[string]interface{} `json:"pull_request,omitempty"`
	User GeneratedType_Nullable_simple_user `json:"user"` // A GitHub user.
	Comments_url string `json:"comments_url"`
	Assignees []GeneratedType_Simple_user `json:"assignees,omitempty"`
	Body_text string `json:"body_text,omitempty"`
	Locked bool `json:"locked"`
	Milestone GeneratedType_Nullable_milestone `json:"milestone"` // A collection of related issues and pull requests.
	Sub_issues_summary GeneratedType_Sub_issues_summary `json:"sub_issues_summary,omitempty"`
	Closed_at string `json:"closed_at"`
	Url string `json:"url"` // URL for the issue
	Closed_by GeneratedType_Nullable_simple_user `json:"closed_by,omitempty"` // A GitHub user.
	Id int64 `json:"id"`
	Reactions GeneratedType_Reaction_rollup `json:"reactions,omitempty"`
	Timeline_url string `json:"timeline_url,omitempty"`
	Node_id string `json:"node_id"`
	TypeField GeneratedType_Issue_type `json:"type,omitempty"` // The type of issue.
	Updated_at string `json:"updated_at"`
	Comments int `json:"comments"`
	Events_url string `json:"events_url"`
	Draft bool `json:"draft,omitempty"`
	Labels []interface{} `json:"labels"` // Labels to associate with this issue; pass one or more label names to replace the set of labels on this issue; send an empty array to clear all labels from the issue; note that the labels are silently dropped for users without push access to the repository
	Body_html string `json:"body_html,omitempty"`
	Created_at string `json:"created_at"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Title string `json:"title"` // Title of the issue
	Body string `json:"body,omitempty"` // Contents of the issue
	State_reason string `json:"state_reason,omitempty"` // The reason for the current state
	Html_url string `json:"html_url"`
	Number int `json:"number"` // Number uniquely identifying the issue within its repository
	State string `json:"state"` // State of the issue; either 'open' or 'closed'
	Performed_via_github_app GeneratedType_Nullable_integration `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Active_lock_reason string `json:"active_lock_reason,omitempty"`
	Repository_url string `json:"repository_url"`
	Labels_url string `json:"labels_url"`
}

// GeneratedType_Webhook_marketplace_purchase_pending_change_cancelled represents the GeneratedType_Webhook_marketplace_purchase_pending_change_cancelled schema from the OpenAPI specification
type GeneratedType_Webhook_marketplace_purchase_pending_change_cancelled struct {
	Effective_date string `json:"effective_date"`
	Marketplace_purchase map[string]interface{} `json:"marketplace_purchase"`
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Action string `json:"action"`
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Previous_marketplace_purchase Webhookspreviousmarketplacepurchase `json:"previous_marketplace_purchase,omitempty"`
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType_Repository_rule_ruleset_info represents the GeneratedType_Repository_rule_ruleset_info schema from the OpenAPI specification
type GeneratedType_Repository_rule_ruleset_info struct {
	Ruleset_source_type string `json:"ruleset_source_type,omitempty"` // The type of source for the ruleset that includes this rule.
	Ruleset_id int `json:"ruleset_id,omitempty"` // The ID of the ruleset that includes this rule.
	Ruleset_source string `json:"ruleset_source,omitempty"` // The name of the source of the ruleset that includes this rule.
}

// GeneratedType_Webhook_member_edited represents the GeneratedType_Webhook_member_edited schema from the OpenAPI specification
type GeneratedType_Webhook_member_edited struct {
	Changes map[string]interface{} `json:"changes"` // The changes to the collaborator permissions
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Member Webhooksuser `json:"member"`
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// Webhookslabel represents the Webhookslabel schema from the OpenAPI specification
type Webhookslabel struct {
	Description string `json:"description"`
	Id int `json:"id"`
	Name string `json:"name"` // The name of the label.
	Node_id string `json:"node_id"`
	Url string `json:"url"` // URL for the label
	Color string `json:"color"` // 6-character hex code, without the leading #, identifying the color
	DefaultField bool `json:"default"`
}

// GeneratedType_Webhook_installation_unsuspend represents the GeneratedType_Webhook_installation_unsuspend schema from the OpenAPI specification
type GeneratedType_Webhook_installation_unsuspend struct {
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation Installation `json:"installation"` // Installation
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repositories []map[string]interface{} `json:"repositories,omitempty"` // An array of repository objects that the installation can access.
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Requester interface{} `json:"requester,omitempty"`
}

// GeneratedType_Branch_protection represents the GeneratedType_Branch_protection schema from the OpenAPI specification
type GeneratedType_Branch_protection struct {
	Allow_fork_syncing map[string]interface{} `json:"allow_fork_syncing,omitempty"` // Whether users can pull changes from upstream when the branch is locked. Set to `true` to allow fork syncing. Set to `false` to prevent fork syncing.
	Required_signatures map[string]interface{} `json:"required_signatures,omitempty"`
	Required_conversation_resolution map[string]interface{} `json:"required_conversation_resolution,omitempty"`
	Allow_deletions map[string]interface{} `json:"allow_deletions,omitempty"`
	Allow_force_pushes map[string]interface{} `json:"allow_force_pushes,omitempty"`
	Required_linear_history map[string]interface{} `json:"required_linear_history,omitempty"`
	Block_creations map[string]interface{} `json:"block_creations,omitempty"`
	Name string `json:"name,omitempty"`
	Lock_branch map[string]interface{} `json:"lock_branch,omitempty"` // Whether to set the branch as read-only. If this is true, users will not be able to push to the branch.
	Enforce_admins GeneratedType_Protected_branch_admin_enforced `json:"enforce_admins,omitempty"` // Protected Branch Admin Enforced
	Restrictions GeneratedType_Branch_restriction_policy `json:"restrictions,omitempty"` // Branch Restriction Policy
	Enabled bool `json:"enabled,omitempty"`
	Required_pull_request_reviews GeneratedType_Protected_branch_pull_request_review `json:"required_pull_request_reviews,omitempty"` // Protected Branch Pull Request Review
	Required_status_checks GeneratedType_Protected_branch_required_status_check `json:"required_status_checks,omitempty"` // Protected Branch Required Status Check
	Url string `json:"url,omitempty"`
	Protection_url string `json:"protection_url,omitempty"`
}

// GeneratedType_Repository_rule_required_deployments represents the GeneratedType_Repository_rule_required_deployments schema from the OpenAPI specification
type GeneratedType_Repository_rule_required_deployments struct {
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	TypeField string `json:"type"`
}

// GeneratedType_Review_comment represents the GeneratedType_Review_comment schema from the OpenAPI specification
type GeneratedType_Review_comment struct {
	Links map[string]interface{} `json:"_links"`
	Line int `json:"line,omitempty"` // The line of the blob to which the comment applies. The last line of the range for a multi-line comment
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Original_line int `json:"original_line,omitempty"` // The original line of the blob to which the comment applies. The last line of the range for a multi-line comment
	Side string `json:"side,omitempty"` // The side of the first line of the range for a multi-line comment.
	User GeneratedType_Nullable_simple_user `json:"user"` // A GitHub user.
	Diff_hunk string `json:"diff_hunk"`
	Html_url string `json:"html_url"`
	Body string `json:"body"`
	Subject_type string `json:"subject_type,omitempty"` // The level at which the comment is targeted, can be a diff line or a file.
	Created_at string `json:"created_at"`
	Pull_request_url string `json:"pull_request_url"`
	Original_commit_id string `json:"original_commit_id"`
	Path string `json:"path"`
	Original_position int `json:"original_position"`
	Original_start_line int `json:"original_start_line,omitempty"` // The original first line of the range for a multi-line comment.
	In_reply_to_id int `json:"in_reply_to_id,omitempty"`
	Updated_at string `json:"updated_at"`
	Pull_request_review_id int64 `json:"pull_request_review_id"`
	Start_side string `json:"start_side,omitempty"` // The side of the first line of the range for a multi-line comment.
	Reactions GeneratedType_Reaction_rollup `json:"reactions,omitempty"`
	Body_text string `json:"body_text,omitempty"`
	Commit_id string `json:"commit_id"`
	Url string `json:"url"`
	Id int64 `json:"id"`
	Start_line int `json:"start_line,omitempty"` // The first line of the range for a multi-line comment.
	Body_html string `json:"body_html,omitempty"`
	Position int `json:"position"`
	Node_id string `json:"node_id"`
}

// GeneratedType_Webhook_package_updated represents the GeneratedType_Webhook_package_updated schema from the OpenAPI specification
type GeneratedType_Webhook_package_updated struct {
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	PackageField map[string]interface{} `json:"package"` // Information about the package.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
}

// GeneratedType_Code_scanning_organization_alert_items represents the GeneratedType_Code_scanning_organization_alert_items schema from the OpenAPI specification
type GeneratedType_Code_scanning_organization_alert_items struct {
	Tool GeneratedType_Code_scanning_analysis_tool `json:"tool"`
	Created_at string `json:"created_at"` // The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Dismissed_at string `json:"dismissed_at"` // The time that the alert was dismissed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Url string `json:"url"` // The REST API URL of the alert resource.
	State string `json:"state"` // State of a code scanning alert.
	Html_url string `json:"html_url"` // The GitHub URL of the alert resource.
	Dismissed_comment string `json:"dismissed_comment,omitempty"` // The dismissal comment associated with the dismissal of the alert.
	Rule GeneratedType_Code_scanning_alert_rule_summary `json:"rule"`
	Dismissed_by GeneratedType_Nullable_simple_user `json:"dismissed_by"` // A GitHub user.
	Fixed_at string `json:"fixed_at,omitempty"` // The time that the alert was no longer detected and was considered fixed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Dismissed_reason string `json:"dismissed_reason"` // **Required when the state is dismissed.** The reason for dismissing or closing the alert.
	Repository GeneratedType_Simple_repository `json:"repository"` // A GitHub repository.
	Dismissal_approved_by GeneratedType_Nullable_simple_user `json:"dismissal_approved_by,omitempty"` // A GitHub user.
	Instances_url string `json:"instances_url"` // The REST API URL for fetching the list of instances for an alert.
	Number int `json:"number"` // The security alert number.
	Updated_at string `json:"updated_at,omitempty"` // The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Most_recent_instance GeneratedType_Code_scanning_alert_instance `json:"most_recent_instance"`
}

// GeneratedType_Base_gist represents the GeneratedType_Base_gist schema from the OpenAPI specification
type GeneratedType_Base_gist struct {
	Git_push_url string `json:"git_push_url"`
	Truncated bool `json:"truncated,omitempty"`
	History []interface{} `json:"history,omitempty"`
	Html_url string `json:"html_url"`
	Owner GeneratedType_Simple_user `json:"owner,omitempty"` // A GitHub user.
	Comments int `json:"comments"`
	Comments_url string `json:"comments_url"`
	Git_pull_url string `json:"git_pull_url"`
	Comments_enabled bool `json:"comments_enabled,omitempty"`
	Id string `json:"id"`
	Public bool `json:"public"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
	Files map[string]interface{} `json:"files"`
	Url string `json:"url"`
	User GeneratedType_Nullable_simple_user `json:"user"` // A GitHub user.
	Node_id string `json:"node_id"`
	Forks []interface{} `json:"forks,omitempty"`
	Forks_url string `json:"forks_url"`
	Commits_url string `json:"commits_url"`
	Description string `json:"description"`
}

// GeneratedType_Combined_billing_usage represents the GeneratedType_Combined_billing_usage schema from the OpenAPI specification
type GeneratedType_Combined_billing_usage struct {
	Days_left_in_billing_cycle int `json:"days_left_in_billing_cycle"` // Numbers of days left in billing cycle.
	Estimated_paid_storage_for_month int `json:"estimated_paid_storage_for_month"` // Estimated storage space (GB) used in billing cycle.
	Estimated_storage_for_month int `json:"estimated_storage_for_month"` // Estimated sum of free and paid storage space (GB) used in billing cycle.
}

// GeneratedType_Webhook_projects_v2_status_update_created represents the GeneratedType_Webhook_projects_v2_status_update_created schema from the OpenAPI specification
type GeneratedType_Webhook_projects_v2_status_update_created struct {
	Action string `json:"action"`
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Projects_v2_status_update GeneratedType_Projects_v2_status_update `json:"projects_v2_status_update"` // An status update belonging to a project
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
}

// GeneratedType_Organization_update_issue_type represents the GeneratedType_Organization_update_issue_type schema from the OpenAPI specification
type GeneratedType_Organization_update_issue_type struct {
	Color string `json:"color,omitempty"` // Color for the issue type.
	Description string `json:"description,omitempty"` // Description of the issue type.
	Is_enabled bool `json:"is_enabled"` // Whether or not the issue type is enabled at the organization level.
	Name string `json:"name"` // Name of the issue type.
}

// GeneratedType_Nullable_actions_hosted_runner_pool_image represents the GeneratedType_Nullable_actions_hosted_runner_pool_image schema from the OpenAPI specification
type GeneratedType_Nullable_actions_hosted_runner_pool_image struct {
	Display_name string `json:"display_name"` // Display name for this image.
	Id string `json:"id"` // The ID of the image. Use this ID for the `image` parameter when creating a new larger runner.
	Size_gb int `json:"size_gb"` // Image size in GB.
	Source string `json:"source"` // The image provider.
}

// GeneratedType_Social_account represents the GeneratedType_Social_account schema from the OpenAPI specification
type GeneratedType_Social_account struct {
	Provider string `json:"provider"`
	Url string `json:"url"`
}

// GeneratedType_Copilot_ide_code_completions represents the GeneratedType_Copilot_ide_code_completions schema from the OpenAPI specification
type GeneratedType_Copilot_ide_code_completions struct {
	Editors []map[string]interface{} `json:"editors,omitempty"`
	Languages []map[string]interface{} `json:"languages,omitempty"` // Code completion metrics for active languages.
	Total_engaged_users int `json:"total_engaged_users,omitempty"` // Number of users who accepted at least one Copilot code suggestion, across all active editors. Includes both full and partial acceptances.
}

// GeneratedType_Projects_v2_iteration_setting represents the GeneratedType_Projects_v2_iteration_setting schema from the OpenAPI specification
type GeneratedType_Projects_v2_iteration_setting struct {
	Duration float64 `json:"duration,omitempty"`
	Id string `json:"id"`
	Start_date string `json:"start_date,omitempty"`
	Title string `json:"title"`
}

// GeneratedType_Deployment_branch_policy represents the GeneratedType_Deployment_branch_policy schema from the OpenAPI specification
type GeneratedType_Deployment_branch_policy struct {
	Id int `json:"id,omitempty"` // The unique identifier of the branch or tag policy.
	Name string `json:"name,omitempty"` // The name pattern that branches or tags must match in order to deploy to the environment.
	Node_id string `json:"node_id,omitempty"`
	TypeField string `json:"type,omitempty"` // Whether this rule targets a branch or tag.
}

// GeneratedType_Pull_request_review_comment represents the GeneratedType_Pull_request_review_comment schema from the OpenAPI specification
type GeneratedType_Pull_request_review_comment struct {
	Original_commit_id string `json:"original_commit_id"` // The SHA of the original commit to which the comment applies.
	Path string `json:"path"` // The relative path of the file to which the comment applies.
	In_reply_to_id int `json:"in_reply_to_id,omitempty"` // The comment ID to reply to.
	Position int `json:"position,omitempty"` // The line index in the diff to which the comment applies. This field is closing down; use `line` instead.
	Commit_id string `json:"commit_id"` // The SHA of the commit to which the comment applies.
	Pull_request_review_id int64 `json:"pull_request_review_id"` // The ID of the pull request review to which the comment belongs.
	Start_side string `json:"start_side,omitempty"` // The side of the first line of the range for a multi-line comment.
	Url string `json:"url"` // URL for the pull request review comment
	Diff_hunk string `json:"diff_hunk"` // The diff of the line that the comment refers to.
	Id int64 `json:"id"` // The ID of the pull request review comment.
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Node_id string `json:"node_id"` // The node ID of the pull request review comment.
	Start_line int `json:"start_line,omitempty"` // The first line of the range for a multi-line comment.
	Updated_at string `json:"updated_at"`
	Created_at string `json:"created_at"`
	Links map[string]interface{} `json:"_links"`
	Body_text string `json:"body_text,omitempty"`
	Reactions GeneratedType_Reaction_rollup `json:"reactions,omitempty"`
	Original_line int `json:"original_line,omitempty"` // The line of the blob to which the comment applies. The last line of the range for a multi-line comment
	Side string `json:"side,omitempty"` // The side of the diff to which the comment applies. The side of the last line of the range for a multi-line comment
	User GeneratedType_Simple_user `json:"user"` // A GitHub user.
	Line int `json:"line,omitempty"` // The line of the blob to which the comment applies. The last line of the range for a multi-line comment
	Body_html string `json:"body_html,omitempty"`
	Html_url string `json:"html_url"` // HTML URL for the pull request review comment.
	Subject_type string `json:"subject_type,omitempty"` // The level at which the comment is targeted, can be a diff line or a file.
	Original_start_line int `json:"original_start_line,omitempty"` // The first line of the range for a multi-line comment.
	Body string `json:"body"` // The text of the comment.
	Original_position int `json:"original_position,omitempty"` // The index of the original line in the diff to which the comment applies. This field is closing down; use `original_line` instead.
	Pull_request_url string `json:"pull_request_url"` // URL for the pull request that the review comment belongs to.
}

// GeneratedType_Repository_advisory represents the GeneratedType_Repository_advisory schema from the OpenAPI specification
type GeneratedType_Repository_advisory struct {
	Cvss_severities GeneratedType_Cvss_severities `json:"cvss_severities,omitempty"`
	Cvss map[string]interface{} `json:"cvss"`
	Cwe_ids []string `json:"cwe_ids"` // A list of only the CWE IDs.
	Identifiers []map[string]interface{} `json:"identifiers"`
	Created_at string `json:"created_at"` // The date and time of when the advisory was created, in ISO 8601 format.
	Published_at string `json:"published_at"` // The date and time of when the advisory was published, in ISO 8601 format.
	State string `json:"state"` // The state of the advisory.
	Ghsa_id string `json:"ghsa_id"` // The GitHub Security Advisory ID.
	Severity string `json:"severity"` // The severity of the advisory.
	Withdrawn_at string `json:"withdrawn_at"` // The date and time of when the advisory was withdrawn, in ISO 8601 format.
	Closed_at string `json:"closed_at"` // The date and time of when the advisory was closed, in ISO 8601 format.
	Html_url string `json:"html_url"` // The URL for the advisory.
	Collaborating_teams []Team `json:"collaborating_teams"` // A list of teams that collaborate on the advisory.
	Credits []map[string]interface{} `json:"credits"`
	Cwes []map[string]interface{} `json:"cwes"`
	Url string `json:"url"` // The API URL for the advisory.
	Credits_detailed []GeneratedType_Repository_advisory_credit `json:"credits_detailed"`
	Description string `json:"description"` // A detailed description of what the advisory entails.
	Private_fork interface{} `json:"private_fork"` // A temporary private fork of the advisory's repository for collaborating on a fix.
	Publisher interface{} `json:"publisher"` // The publisher of the advisory.
	Submission map[string]interface{} `json:"submission"`
	Summary string `json:"summary"` // A short summary of the advisory.
	Updated_at string `json:"updated_at"` // The date and time of when the advisory was last updated, in ISO 8601 format.
	Vulnerabilities []GeneratedType_Repository_advisory_vulnerability `json:"vulnerabilities"`
	Author interface{} `json:"author"` // The author of the advisory.
	Collaborating_users []GeneratedType_Simple_user `json:"collaborating_users"` // A list of users that collaborate on the advisory.
	Cve_id string `json:"cve_id"` // The Common Vulnerabilities and Exposures (CVE) ID.
}

// GeneratedType_Webhook_check_run_rerequested represents the GeneratedType_Webhook_check_run_rerequested schema from the OpenAPI specification
type GeneratedType_Webhook_check_run_rerequested struct {
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action,omitempty"`
	Check_run GeneratedType_Check_run_with_simple_check_suite `json:"check_run"` // A check performed on the code of a given code change
}

// GeneratedType_Webhook_star_deleted represents the GeneratedType_Webhook_star_deleted schema from the OpenAPI specification
type GeneratedType_Webhook_star_deleted struct {
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Starred_at interface{} `json:"starred_at"` // The time the star was created. This is a timestamp in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`. Will be `null` for the `deleted` action.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType_Webhook_release_unpublished represents the GeneratedType_Webhook_release_unpublished schema from the OpenAPI specification
type GeneratedType_Webhook_release_unpublished struct {
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Release Webhooksrelease1 `json:"release"` // The [release](https://docs.github.com/rest/releases/releases/#get-a-release) object.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType_Repository_rule represents the GeneratedType_Repository_rule schema from the OpenAPI specification
type GeneratedType_Repository_rule struct {
}

// GeneratedType_Contributor_activity represents the GeneratedType_Contributor_activity schema from the OpenAPI specification
type GeneratedType_Contributor_activity struct {
	Author GeneratedType_Nullable_simple_user `json:"author"` // A GitHub user.
	Total int `json:"total"`
	Weeks []map[string]interface{} `json:"weeks"`
}

// GeneratedType_Webhook_organization_member_removed represents the GeneratedType_Webhook_organization_member_removed schema from the OpenAPI specification
type GeneratedType_Webhook_organization_member_removed struct {
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Membership Webhooksmembership `json:"membership"` // The membership between the user and the organization. Not present when the action is `member_invited`.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType_Webhook_pull_request_closed represents the GeneratedType_Webhook_pull_request_closed schema from the OpenAPI specification
type GeneratedType_Webhook_pull_request_closed struct {
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request GeneratedType_Pull_request_webhook `json:"pull_request"`
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType_Renamed_issue_event represents the GeneratedType_Renamed_issue_event schema from the OpenAPI specification
type GeneratedType_Renamed_issue_event struct {
	Commit_url string `json:"commit_url"`
	Created_at string `json:"created_at"`
	Performed_via_github_app GeneratedType_Nullable_integration `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Rename map[string]interface{} `json:"rename"`
	Url string `json:"url"`
	Actor GeneratedType_Simple_user `json:"actor"` // A GitHub user.
	Event string `json:"event"`
	Commit_id string `json:"commit_id"`
	Id int `json:"id"`
	Node_id string `json:"node_id"`
}

// GeneratedType_Webhook_check_suite_requested represents the GeneratedType_Webhook_check_suite_requested schema from the OpenAPI specification
type GeneratedType_Webhook_check_suite_requested struct {
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Check_suite map[string]interface{} `json:"check_suite"` // The [check_suite](https://docs.github.com/rest/checks/suites#get-a-check-suite).
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType_Webhook_org_block_blocked represents the GeneratedType_Webhook_org_block_blocked schema from the OpenAPI specification
type GeneratedType_Webhook_org_block_blocked struct {
	Action string `json:"action"`
	Blocked_user Webhooksuser `json:"blocked_user"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
}

// GeneratedType_Nullable_issue represents the GeneratedType_Nullable_issue schema from the OpenAPI specification
type GeneratedType_Nullable_issue struct {
	Comments_url string `json:"comments_url"`
	Performed_via_github_app GeneratedType_Nullable_integration `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Active_lock_reason string `json:"active_lock_reason,omitempty"`
	Reactions GeneratedType_Reaction_rollup `json:"reactions,omitempty"`
	Comments int `json:"comments"`
	Labels []interface{} `json:"labels"` // Labels to associate with this issue; pass one or more label names to replace the set of labels on this issue; send an empty array to clear all labels from the issue; note that the labels are silently dropped for users without push access to the repository
	Number int `json:"number"` // Number uniquely identifying the issue within its repository
	Repository_url string `json:"repository_url"`
	Sub_issues_summary GeneratedType_Sub_issues_summary `json:"sub_issues_summary,omitempty"`
	Draft bool `json:"draft,omitempty"`
	Labels_url string `json:"labels_url"`
	Updated_at string `json:"updated_at"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Body_html string `json:"body_html,omitempty"`
	Closed_at string `json:"closed_at"`
	Pull_request map[string]interface{} `json:"pull_request,omitempty"`
	User GeneratedType_Nullable_simple_user `json:"user"` // A GitHub user.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Timeline_url string `json:"timeline_url,omitempty"`
	Body_text string `json:"body_text,omitempty"`
	Closed_by GeneratedType_Nullable_simple_user `json:"closed_by,omitempty"` // A GitHub user.
	Events_url string `json:"events_url"`
	Node_id string `json:"node_id"`
	Html_url string `json:"html_url"`
	State string `json:"state"` // State of the issue; either 'open' or 'closed'
	Assignees []GeneratedType_Simple_user `json:"assignees,omitempty"`
	Locked bool `json:"locked"`
	Created_at string `json:"created_at"`
	Url string `json:"url"` // URL for the issue
	Assignee GeneratedType_Nullable_simple_user `json:"assignee"` // A GitHub user.
	Milestone GeneratedType_Nullable_milestone `json:"milestone"` // A collection of related issues and pull requests.
	State_reason string `json:"state_reason,omitempty"` // The reason for the current state
	Title string `json:"title"` // Title of the issue
	TypeField GeneratedType_Issue_type `json:"type,omitempty"` // The type of issue.
	Body string `json:"body,omitempty"` // Contents of the issue
	Id int64 `json:"id"`
}

// GeneratedType_Code_scanning_analysis_deletion represents the GeneratedType_Code_scanning_analysis_deletion schema from the OpenAPI specification
type GeneratedType_Code_scanning_analysis_deletion struct {
	Next_analysis_url string `json:"next_analysis_url"` // Next deletable analysis in chain, without last analysis deletion confirmation
	Confirm_delete_url string `json:"confirm_delete_url"` // Next deletable analysis in chain, with last analysis deletion confirmation
}

// GeneratedType_Webhook_project_card_edited represents the GeneratedType_Webhook_project_card_edited schema from the OpenAPI specification
type GeneratedType_Webhook_project_card_edited struct {
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Project_card Webhooksprojectcard `json:"project_card"`
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// Webhooksalert represents the Webhooksalert schema from the OpenAPI specification
type Webhooksalert struct {
	Affected_range string `json:"affected_range"`
	Number int `json:"number"`
	Dismiss_reason string `json:"dismiss_reason,omitempty"`
	Ghsa_id string `json:"ghsa_id"`
	Dismissed_at string `json:"dismissed_at,omitempty"`
	Dismisser map[string]interface{} `json:"dismisser,omitempty"`
	External_reference string `json:"external_reference"`
	Fix_reason string `json:"fix_reason,omitempty"`
	Node_id string `json:"node_id"`
	Fixed_in string `json:"fixed_in,omitempty"`
	Created_at string `json:"created_at"`
	External_identifier string `json:"external_identifier"`
	Fixed_at string `json:"fixed_at,omitempty"`
	State string `json:"state"`
	Affected_package_name string `json:"affected_package_name"`
	Id int `json:"id"`
	Severity string `json:"severity"`
}

// Job represents the Job schema from the OpenAPI specification
type Job struct {
	Runner_name string `json:"runner_name"` // The name of the runner to which this job has been assigned. (If a runner hasn't yet been assigned, this will be null.)
	Head_branch string `json:"head_branch"` // The name of the current branch.
	Labels []string `json:"labels"` // Labels for the workflow job. Specified by the "runs_on" attribute in the action's workflow file.
	Node_id string `json:"node_id"`
	Html_url string `json:"html_url"`
	Url string `json:"url"`
	Runner_group_id int `json:"runner_group_id"` // The ID of the runner group to which this job has been assigned. (If a runner hasn't yet been assigned, this will be null.)
	Steps []map[string]interface{} `json:"steps,omitempty"` // Steps in this job.
	Conclusion string `json:"conclusion"` // The outcome of the job.
	Run_attempt int `json:"run_attempt,omitempty"` // Attempt number of the associated workflow run, 1 for first attempt and higher if the workflow was re-run.
	Check_run_url string `json:"check_run_url"`
	Created_at string `json:"created_at"` // The time that the job created, in ISO 8601 format.
	Completed_at string `json:"completed_at"` // The time that the job finished, in ISO 8601 format.
	Head_sha string `json:"head_sha"` // The SHA of the commit that is being run.
	Status string `json:"status"` // The phase of the lifecycle that the job is currently in.
	Id int `json:"id"` // The id of the job.
	Run_id int `json:"run_id"` // The id of the associated workflow run.
	Workflow_name string `json:"workflow_name"` // The name of the workflow.
	Name string `json:"name"` // The name of the job.
	Runner_group_name string `json:"runner_group_name"` // The name of the runner group to which this job has been assigned. (If a runner hasn't yet been assigned, this will be null.)
	Runner_id int `json:"runner_id"` // The ID of the runner to which this job has been assigned. (If a runner hasn't yet been assigned, this will be null.)
	Run_url string `json:"run_url"`
	Started_at string `json:"started_at"` // The time that the job started, in ISO 8601 format.
}

// GeneratedType_Webhook_discussion_created represents the GeneratedType_Webhook_discussion_created schema from the OpenAPI specification
type GeneratedType_Webhook_discussion_created struct {
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType_Repository_rule_branch_name_pattern represents the GeneratedType_Repository_rule_branch_name_pattern schema from the OpenAPI specification
type GeneratedType_Repository_rule_branch_name_pattern struct {
	TypeField string `json:"type"`
	Parameters map[string]interface{} `json:"parameters,omitempty"`
}

// GeneratedType_Organization_invitation represents the GeneratedType_Organization_invitation schema from the OpenAPI specification
type GeneratedType_Organization_invitation struct {
	Failed_at string `json:"failed_at,omitempty"`
	Failed_reason string `json:"failed_reason,omitempty"`
	Invitation_teams_url string `json:"invitation_teams_url"`
	Login string `json:"login"`
	Invitation_source string `json:"invitation_source,omitempty"`
	Inviter GeneratedType_Simple_user `json:"inviter"` // A GitHub user.
	Node_id string `json:"node_id"`
	Role string `json:"role"`
	Team_count int `json:"team_count"`
	Created_at string `json:"created_at"`
	Id int64 `json:"id"`
	Email string `json:"email"`
}

// GeneratedType_Code_scanning_variant_analysis_repo_task represents the GeneratedType_Code_scanning_variant_analysis_repo_task schema from the OpenAPI specification
type GeneratedType_Code_scanning_variant_analysis_repo_task struct {
	Repository GeneratedType_Simple_repository `json:"repository"` // A GitHub repository.
	Result_count int `json:"result_count,omitempty"` // The number of results in the case of a successful analysis. This is only available for successful analyses.
	Source_location_prefix string `json:"source_location_prefix,omitempty"` // The source location prefix to use. This is only available for successful analyses.
	Analysis_status string `json:"analysis_status"` // The new status of the CodeQL variant analysis repository task.
	Artifact_size_in_bytes int `json:"artifact_size_in_bytes,omitempty"` // The size of the artifact. This is only available for successful analyses.
	Artifact_url string `json:"artifact_url,omitempty"` // The URL of the artifact. This is only available for successful analyses.
	Database_commit_sha string `json:"database_commit_sha,omitempty"` // The SHA of the commit the CodeQL database was built against. This is only available for successful analyses.
	Failure_message string `json:"failure_message,omitempty"` // The reason of the failure of this repo task. This is only available if the repository task has failed.
}

// GeneratedType_Locked_issue_event represents the GeneratedType_Locked_issue_event schema from the OpenAPI specification
type GeneratedType_Locked_issue_event struct {
	Created_at string `json:"created_at"`
	Node_id string `json:"node_id"`
	Commit_url string `json:"commit_url"`
	Event string `json:"event"`
	Id int `json:"id"`
	Lock_reason string `json:"lock_reason"`
	Performed_via_github_app GeneratedType_Nullable_integration `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Url string `json:"url"`
	Actor GeneratedType_Simple_user `json:"actor"` // A GitHub user.
	Commit_id string `json:"commit_id"`
}

// GeneratedType_Security_and_analysis represents the GeneratedType_Security_and_analysis schema from the OpenAPI specification
type GeneratedType_Security_and_analysis struct {
	Secret_scanning_push_protection map[string]interface{} `json:"secret_scanning_push_protection,omitempty"`
	Advanced_security map[string]interface{} `json:"advanced_security,omitempty"`
	Code_security map[string]interface{} `json:"code_security,omitempty"`
	Dependabot_security_updates map[string]interface{} `json:"dependabot_security_updates,omitempty"` // Enable or disable Dependabot security updates for the repository.
	Secret_scanning map[string]interface{} `json:"secret_scanning,omitempty"`
	Secret_scanning_ai_detection map[string]interface{} `json:"secret_scanning_ai_detection,omitempty"`
	Secret_scanning_non_provider_patterns map[string]interface{} `json:"secret_scanning_non_provider_patterns,omitempty"`
}

// GeneratedType_Deployment_simple represents the GeneratedType_Deployment_simple schema from the OpenAPI specification
type GeneratedType_Deployment_simple struct {
	Created_at string `json:"created_at"`
	Description string `json:"description"`
	Original_environment string `json:"original_environment,omitempty"`
	Production_environment bool `json:"production_environment,omitempty"` // Specifies if the given environment is one that end-users directly interact with. Default: false.
	Repository_url string `json:"repository_url"`
	Transient_environment bool `json:"transient_environment,omitempty"` // Specifies if the given environment is will no longer exist at some point in the future. Default: false.
	Environment string `json:"environment"` // Name for the target deployment environment.
	Id int `json:"id"` // Unique identifier of the deployment
	Task string `json:"task"` // Parameter to specify a task to execute
	Url string `json:"url"`
	Node_id string `json:"node_id"`
	Performed_via_github_app GeneratedType_Nullable_integration `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Statuses_url string `json:"statuses_url"`
	Updated_at string `json:"updated_at"`
}

// GeneratedType_Webhook_pull_request_labeled represents the GeneratedType_Webhook_pull_request_labeled schema from the OpenAPI specification
type GeneratedType_Webhook_pull_request_labeled struct {
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Action string `json:"action"`
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Pull_request map[string]interface{} `json:"pull_request"`
	Label Webhookslabel `json:"label,omitempty"`
}

// GeneratedType_Page_build_status represents the GeneratedType_Page_build_status schema from the OpenAPI specification
type GeneratedType_Page_build_status struct {
	Status string `json:"status"`
	Url string `json:"url"`
}

// GeneratedType_Code_scanning_alert_instance represents the GeneratedType_Code_scanning_alert_instance schema from the OpenAPI specification
type GeneratedType_Code_scanning_alert_instance struct {
	Analysis_key string `json:"analysis_key,omitempty"` // Identifies the configuration under which the analysis was executed. For example, in GitHub Actions this includes the workflow filename and job name.
	Classifications []string `json:"classifications,omitempty"` // Classifications that have been applied to the file that triggered the alert. For example identifying it as documentation, or a generated file.
	Location GeneratedType_Code_scanning_alert_location `json:"location,omitempty"` // Describe a region within a file for the alert.
	Category string `json:"category,omitempty"` // Identifies the configuration under which the analysis was executed. Used to distinguish between multiple analyses for the same tool and commit, but performed on different languages or different parts of the code.
	Commit_sha string `json:"commit_sha,omitempty"`
	Ref string `json:"ref,omitempty"` // The Git reference, formatted as `refs/pull/<number>/merge`, `refs/pull/<number>/head`, `refs/heads/<branch name>` or simply `<branch name>`.
	Environment string `json:"environment,omitempty"` // Identifies the variable values associated with the environment in which the analysis that generated this alert instance was performed, such as the language that was analyzed.
	Html_url string `json:"html_url,omitempty"`
	Message map[string]interface{} `json:"message,omitempty"`
	State string `json:"state,omitempty"` // State of a code scanning alert.
}

// GeneratedType_Webhook_dependabot_alert_dismissed represents the GeneratedType_Webhook_dependabot_alert_dismissed schema from the OpenAPI specification
type GeneratedType_Webhook_dependabot_alert_dismissed struct {
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Alert GeneratedType_Dependabot_alert `json:"alert"` // A Dependabot alert.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType_Private_user represents the GeneratedType_Private_user schema from the OpenAPI specification
type GeneratedType_Private_user struct {
	Email string `json:"email"`
	Gists_url string `json:"gists_url"`
	Gravatar_id string `json:"gravatar_id"`
	Id int64 `json:"id"`
	Updated_at string `json:"updated_at"`
	Html_url string `json:"html_url"`
	Total_private_repos int `json:"total_private_repos"`
	User_view_type string `json:"user_view_type,omitempty"`
	Public_gists int `json:"public_gists"`
	Plan map[string]interface{} `json:"plan,omitempty"`
	Events_url string `json:"events_url"`
	Notification_email string `json:"notification_email,omitempty"`
	Received_events_url string `json:"received_events_url"`
	Starred_url string `json:"starred_url"`
	Owned_private_repos int `json:"owned_private_repos"`
	Public_repos int `json:"public_repos"`
	Hireable bool `json:"hireable"`
	Login string `json:"login"`
	Organizations_url string `json:"organizations_url"`
	Avatar_url string `json:"avatar_url"`
	Twitter_username string `json:"twitter_username,omitempty"`
	Bio string `json:"bio"`
	Url string `json:"url"`
	Business_plus bool `json:"business_plus,omitempty"`
	Following int `json:"following"`
	Ldap_dn string `json:"ldap_dn,omitempty"`
	Site_admin bool `json:"site_admin"`
	Followers int `json:"followers"`
	Location string `json:"location"`
	Created_at string `json:"created_at"`
	Followers_url string `json:"followers_url"`
	Company string `json:"company"`
	Subscriptions_url string `json:"subscriptions_url"`
	Repos_url string `json:"repos_url"`
	Two_factor_authentication bool `json:"two_factor_authentication"`
	Name string `json:"name"`
	Blog string `json:"blog"`
	Disk_usage int `json:"disk_usage"`
	TypeField string `json:"type"`
	Collaborators int `json:"collaborators"`
	Following_url string `json:"following_url"`
	Private_gists int `json:"private_gists"`
	Node_id string `json:"node_id"`
}

// GeneratedType_Check_automated_security_fixes represents the GeneratedType_Check_automated_security_fixes schema from the OpenAPI specification
type GeneratedType_Check_automated_security_fixes struct {
	Enabled bool `json:"enabled"` // Whether Dependabot security updates are enabled for the repository.
	Paused bool `json:"paused"` // Whether Dependabot security updates are paused for the repository.
}

// GeneratedType_Webhook_milestone_created represents the GeneratedType_Webhook_milestone_created schema from the OpenAPI specification
type GeneratedType_Webhook_milestone_created struct {
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Milestone Webhooksmilestone3 `json:"milestone"` // A collection of related issues and pull requests.
}

// GeneratedType_Content_submodule represents the GeneratedType_Content_submodule schema from the OpenAPI specification
type GeneratedType_Content_submodule struct {
	Submodule_git_url string `json:"submodule_git_url"`
	Url string `json:"url"`
	Git_url string `json:"git_url"`
	Name string `json:"name"`
	Path string `json:"path"`
	Size int `json:"size"`
	TypeField string `json:"type"`
	Links map[string]interface{} `json:"_links"`
	Download_url string `json:"download_url"`
	Sha string `json:"sha"`
	Html_url string `json:"html_url"`
}

// GeneratedType_Review_dismissed_issue_event represents the GeneratedType_Review_dismissed_issue_event schema from the OpenAPI specification
type GeneratedType_Review_dismissed_issue_event struct {
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	Commit_url string `json:"commit_url"`
	Event string `json:"event"`
	Actor GeneratedType_Simple_user `json:"actor"` // A GitHub user.
	Dismissed_review map[string]interface{} `json:"dismissed_review"`
	Url string `json:"url"`
	Commit_id string `json:"commit_id"`
	Performed_via_github_app GeneratedType_Nullable_integration `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Created_at string `json:"created_at"`
}

// GeneratedType_Repository_advisory_vulnerability represents the GeneratedType_Repository_advisory_vulnerability schema from the OpenAPI specification
type GeneratedType_Repository_advisory_vulnerability struct {
	Vulnerable_functions []string `json:"vulnerable_functions"` // The functions in the package that are affected.
	Vulnerable_version_range string `json:"vulnerable_version_range"` // The range of the package versions affected by the vulnerability.
	PackageField map[string]interface{} `json:"package"` // The name of the package affected by the vulnerability.
	Patched_versions string `json:"patched_versions"` // The package version(s) that resolve the vulnerability.
}

// GeneratedType_Webhook_deploy_key_deleted represents the GeneratedType_Webhook_deploy_key_deleted schema from the OpenAPI specification
type GeneratedType_Webhook_deploy_key_deleted struct {
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Key Webhooksdeploykey `json:"key"` // The [`deploy key`](https://docs.github.com/rest/deploy-keys/deploy-keys#get-a-deploy-key) resource.
}

// GeneratedType_Converted_note_to_issue_issue_event represents the GeneratedType_Converted_note_to_issue_issue_event schema from the OpenAPI specification
type GeneratedType_Converted_note_to_issue_issue_event struct {
	Commit_id string `json:"commit_id"`
	Commit_url string `json:"commit_url"`
	Created_at string `json:"created_at"`
	Performed_via_github_app Integration `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Node_id string `json:"node_id"`
	Project_card map[string]interface{} `json:"project_card,omitempty"`
	Url string `json:"url"`
	Event string `json:"event"`
	Actor GeneratedType_Simple_user `json:"actor"` // A GitHub user.
	Id int `json:"id"`
}

// GeneratedType_Codeowners_errors represents the GeneratedType_Codeowners_errors schema from the OpenAPI specification
type GeneratedType_Codeowners_errors struct {
	Errors []map[string]interface{} `json:"errors"`
}

// GeneratedType_Webhook_projects_v2_status_update_deleted represents the GeneratedType_Webhook_projects_v2_status_update_deleted schema from the OpenAPI specification
type GeneratedType_Webhook_projects_v2_status_update_deleted struct {
	Action string `json:"action"`
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Projects_v2_status_update GeneratedType_Projects_v2_status_update `json:"projects_v2_status_update"` // An status update belonging to a project
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
}

// GeneratedType_Timeline_unassigned_issue_event represents the GeneratedType_Timeline_unassigned_issue_event schema from the OpenAPI specification
type GeneratedType_Timeline_unassigned_issue_event struct {
	Created_at string `json:"created_at"`
	Actor GeneratedType_Simple_user `json:"actor"` // A GitHub user.
	Commit_url string `json:"commit_url"`
	Event string `json:"event"`
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	Performed_via_github_app GeneratedType_Nullable_integration `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Commit_id string `json:"commit_id"`
	Url string `json:"url"`
	Assignee GeneratedType_Simple_user `json:"assignee"` // A GitHub user.
}

// GeneratedType_Dependabot_alert_with_repository represents the GeneratedType_Dependabot_alert_with_repository schema from the OpenAPI specification
type GeneratedType_Dependabot_alert_with_repository struct {
	Updated_at string `json:"updated_at"` // The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Dismissed_by GeneratedType_Nullable_simple_user `json:"dismissed_by"` // A GitHub user.
	Dismissed_at string `json:"dismissed_at"` // The time that the alert was dismissed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Security_advisory GeneratedType_Dependabot_alert_security_advisory `json:"security_advisory"` // Details for the GitHub Security Advisory.
	Created_at string `json:"created_at"` // The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Dependency map[string]interface{} `json:"dependency"` // Details for the vulnerable dependency.
	Url string `json:"url"` // The REST API URL of the alert resource.
	Auto_dismissed_at string `json:"auto_dismissed_at,omitempty"` // The time that the alert was auto-dismissed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	State string `json:"state"` // The state of the Dependabot alert.
	Fixed_at string `json:"fixed_at"` // The time that the alert was no longer detected and was considered fixed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Repository GeneratedType_Simple_repository `json:"repository"` // A GitHub repository.
	Dismissed_comment string `json:"dismissed_comment"` // An optional comment associated with the alert's dismissal.
	Dismissed_reason string `json:"dismissed_reason"` // The reason that the alert was dismissed.
	Html_url string `json:"html_url"` // The GitHub URL of the alert resource.
	Number int `json:"number"` // The security alert number.
	Security_vulnerability GeneratedType_Dependabot_alert_security_vulnerability `json:"security_vulnerability"` // Details pertaining to one vulnerable version range for the advisory.
}

// GeneratedType_Page_build represents the GeneratedType_Page_build schema from the OpenAPI specification
type GeneratedType_Page_build struct {
	Status string `json:"status"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	Commit string `json:"commit"`
	Created_at string `json:"created_at"`
	Duration int `json:"duration"`
	ErrorField map[string]interface{} `json:"error"`
	Pusher GeneratedType_Nullable_simple_user `json:"pusher"` // A GitHub user.
}

// GeneratedType_Code_scanning_alert_location represents the GeneratedType_Code_scanning_alert_location schema from the OpenAPI specification
type GeneratedType_Code_scanning_alert_location struct {
	End_line int `json:"end_line,omitempty"`
	Path string `json:"path,omitempty"`
	Start_column int `json:"start_column,omitempty"`
	Start_line int `json:"start_line,omitempty"`
	End_column int `json:"end_column,omitempty"`
}

// Environment represents the Environment schema from the OpenAPI specification
type Environment struct {
	Url string `json:"url"`
	Created_at string `json:"created_at"` // The time that the environment was created, in ISO 8601 format.
	Deployment_branch_policy GeneratedType_Deployment_branch_policy_settings `json:"deployment_branch_policy,omitempty"` // The type of deployment branch policy for this environment. To allow all branches to deploy, set to `null`.
	Node_id string `json:"node_id"`
	Html_url string `json:"html_url"`
	Id int64 `json:"id"` // The id of the environment.
	Name string `json:"name"` // The name of the environment.
	Protection_rules []interface{} `json:"protection_rules,omitempty"` // Built-in deployment protection rules for the environment.
	Updated_at string `json:"updated_at"` // The time that the environment was last updated, in ISO 8601 format.
}

// GeneratedType_Webhook_repository_import represents the GeneratedType_Webhook_repository_import schema from the OpenAPI specification
type GeneratedType_Webhook_repository_import struct {
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Status string `json:"status"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType_Pages_source_hash represents the GeneratedType_Pages_source_hash schema from the OpenAPI specification
type GeneratedType_Pages_source_hash struct {
	Branch string `json:"branch"`
	Path string `json:"path"`
}

// GeneratedType_Webhook_pull_request_unlocked represents the GeneratedType_Webhook_pull_request_unlocked schema from the OpenAPI specification
type GeneratedType_Webhook_pull_request_unlocked struct {
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType_Org_hook represents the GeneratedType_Org_hook schema from the OpenAPI specification
type GeneratedType_Org_hook struct {
	Deliveries_url string `json:"deliveries_url,omitempty"`
	Id int `json:"id"`
	Config map[string]interface{} `json:"config"`
	Events []string `json:"events"`
	Name string `json:"name"`
	Updated_at string `json:"updated_at"`
	Active bool `json:"active"`
	Ping_url string `json:"ping_url"`
	TypeField string `json:"type"`
	Url string `json:"url"`
	Created_at string `json:"created_at"`
}

// Webhookscomment represents the Webhookscomment schema from the OpenAPI specification
type Webhookscomment struct {
	Repository_url string `json:"repository_url"`
	Created_at string `json:"created_at"`
	Discussion_id int `json:"discussion_id"`
	Child_comment_count int `json:"child_comment_count"`
	Html_url string `json:"html_url"`
	Parent_id int `json:"parent_id"`
	Reactions map[string]interface{} `json:"reactions"`
	Updated_at string `json:"updated_at"`
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	User map[string]interface{} `json:"user"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Body string `json:"body"`
}

// GeneratedType_Webhook_project_column_moved represents the GeneratedType_Webhook_project_column_moved schema from the OpenAPI specification
type GeneratedType_Webhook_project_column_moved struct {
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Project_column Webhooksprojectcolumn `json:"project_column"`
}

// GeneratedType_Repo_search_result_item represents the GeneratedType_Repo_search_result_item schema from the OpenAPI specification
type GeneratedType_Repo_search_result_item struct {
	Score float64 `json:"score"`
	Contributors_url string `json:"contributors_url"`
	Description string `json:"description"`
	Pulls_url string `json:"pulls_url"`
	Deployments_url string `json:"deployments_url"`
	Topics []string `json:"topics,omitempty"`
	Archive_url string `json:"archive_url"`
	Merges_url string `json:"merges_url"`
	Allow_squash_merge bool `json:"allow_squash_merge,omitempty"`
	Has_pages bool `json:"has_pages"`
	Allow_rebase_merge bool `json:"allow_rebase_merge,omitempty"`
	Issue_comment_url string `json:"issue_comment_url"`
	Comments_url string `json:"comments_url"`
	Mirror_url string `json:"mirror_url"`
	Updated_at string `json:"updated_at"`
	Notifications_url string `json:"notifications_url"`
	Branches_url string `json:"branches_url"`
	Releases_url string `json:"releases_url"`
	Assignees_url string `json:"assignees_url"`
	Node_id string `json:"node_id"`
	Git_refs_url string `json:"git_refs_url"`
	Fork bool `json:"fork"`
	Allow_auto_merge bool `json:"allow_auto_merge,omitempty"`
	Labels_url string `json:"labels_url"`
	Size int `json:"size"`
	Collaborators_url string `json:"collaborators_url"`
	Events_url string `json:"events_url"`
	Allow_merge_commit bool `json:"allow_merge_commit,omitempty"`
	Issue_events_url string `json:"issue_events_url"`
	Keys_url string `json:"keys_url"`
	Open_issues_count int `json:"open_issues_count"`
	Git_commits_url string `json:"git_commits_url"`
	Archived bool `json:"archived"`
	Subscribers_url string `json:"subscribers_url"`
	Has_discussions bool `json:"has_discussions,omitempty"`
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"`
	Visibility string `json:"visibility,omitempty"` // The repository visibility: public, private, or internal.
	Issues_url string `json:"issues_url"`
	Has_wiki bool `json:"has_wiki"`
	Has_issues bool `json:"has_issues"`
	Private bool `json:"private"`
	Ssh_url string `json:"ssh_url"`
	Is_template bool `json:"is_template,omitempty"`
	License GeneratedType_Nullable_license_simple `json:"license"` // License Simple
	Subscription_url string `json:"subscription_url"`
	Compare_url string `json:"compare_url"`
	Git_url string `json:"git_url"`
	Downloads_url string `json:"downloads_url"`
	Contents_url string `json:"contents_url"`
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"`
	Svn_url string `json:"svn_url"`
	Language string `json:"language"`
	Milestones_url string `json:"milestones_url"`
	Hooks_url string `json:"hooks_url"`
	Pushed_at string `json:"pushed_at"`
	Has_downloads bool `json:"has_downloads"`
	Languages_url string `json:"languages_url"`
	Commits_url string `json:"commits_url"`
	Clone_url string `json:"clone_url"`
	Watchers_count int `json:"watchers_count"`
	Forks_count int `json:"forks_count"`
	Full_name string `json:"full_name"`
	Statuses_url string `json:"statuses_url"`
	Watchers int `json:"watchers"`
	Open_issues int `json:"open_issues"`
	Master_branch string `json:"master_branch,omitempty"`
	Forks_url string `json:"forks_url"`
	Stargazers_url string `json:"stargazers_url"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Html_url string `json:"html_url"`
	Homepage string `json:"homepage"`
	Default_branch string `json:"default_branch"`
	Url string `json:"url"`
	Has_projects bool `json:"has_projects"`
	Owner GeneratedType_Nullable_simple_user `json:"owner"` // A GitHub user.
	Temp_clone_token string `json:"temp_clone_token,omitempty"`
	Name string `json:"name"`
	Created_at string `json:"created_at"`
	Disabled bool `json:"disabled"` // Returns whether or not this repository disabled.
	Id int `json:"id"`
	Teams_url string `json:"teams_url"`
	Forks int `json:"forks"`
	Text_matches []map[string]interface{} `json:"text_matches,omitempty"`
	Allow_forking bool `json:"allow_forking,omitempty"`
	Stargazers_count int `json:"stargazers_count"`
	Tags_url string `json:"tags_url"`
	Git_tags_url string `json:"git_tags_url"`
	Trees_url string `json:"trees_url"`
	Blobs_url string `json:"blobs_url"`
}

// GeneratedType_Webhook_deployment_created represents the GeneratedType_Webhook_deployment_created schema from the OpenAPI specification
type GeneratedType_Webhook_deployment_created struct {
	Workflow Webhooksworkflow `json:"workflow"`
	Action string `json:"action"`
	Deployment map[string]interface{} `json:"deployment"` // The [deployment](https://docs.github.com/rest/deployments/deployments#list-deployments).
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Workflow_run map[string]interface{} `json:"workflow_run"`
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType_Webhook_project_deleted represents the GeneratedType_Webhook_project_deleted schema from the OpenAPI specification
type GeneratedType_Webhook_project_deleted struct {
	Repository GeneratedType_Nullable_repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Project Webhooksproject `json:"project"`
}

// GeneratedType_Webhook_ping represents the GeneratedType_Webhook_ping schema from the OpenAPI specification
type GeneratedType_Webhook_ping struct {
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Zen string `json:"zen,omitempty"` // Random string of GitHub zen.
	Hook map[string]interface{} `json:"hook,omitempty"` // The webhook that is being pinged
	Hook_id int `json:"hook_id,omitempty"` // The ID of the webhook that triggered the ping.
}

// GeneratedType_Assigned_issue_event represents the GeneratedType_Assigned_issue_event schema from the OpenAPI specification
type GeneratedType_Assigned_issue_event struct {
	Url string `json:"url"`
	Actor GeneratedType_Simple_user `json:"actor"` // A GitHub user.
	Node_id string `json:"node_id"`
	Commit_id string `json:"commit_id"`
	Commit_url string `json:"commit_url"`
	Assignee GeneratedType_Simple_user `json:"assignee"` // A GitHub user.
	Event string `json:"event"`
	Id int `json:"id"`
	Assigner GeneratedType_Simple_user `json:"assigner"` // A GitHub user.
	Created_at string `json:"created_at"`
	Performed_via_github_app Integration `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
}

// GeneratedType_Org_ruleset_conditions represents the GeneratedType_Org_ruleset_conditions schema from the OpenAPI specification
type GeneratedType_Org_ruleset_conditions struct {
}

// GeneratedType_Webhook_marketplace_purchase_purchased represents the GeneratedType_Webhook_marketplace_purchase_purchased schema from the OpenAPI specification
type GeneratedType_Webhook_marketplace_purchase_purchased struct {
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Marketplace_purchase Webhooksmarketplacepurchase `json:"marketplace_purchase"`
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Previous_marketplace_purchase Webhookspreviousmarketplacepurchase `json:"previous_marketplace_purchase,omitempty"`
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Effective_date string `json:"effective_date"`
}

// GeneratedType_Webhook_code_scanning_alert_created represents the GeneratedType_Webhook_code_scanning_alert_created schema from the OpenAPI specification
type GeneratedType_Webhook_code_scanning_alert_created struct {
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Alert map[string]interface{} `json:"alert"` // The code scanning alert involved in the event.
	Commit_oid string `json:"commit_oid"` // The commit SHA of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Action string `json:"action"`
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Ref string `json:"ref"` // The Git reference of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType_Webhook_pull_request_milestoned represents the GeneratedType_Webhook_pull_request_milestoned schema from the OpenAPI specification
type GeneratedType_Webhook_pull_request_milestoned struct {
	Pull_request Webhookspullrequest5 `json:"pull_request"`
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Milestone Milestone `json:"milestone,omitempty"` // A collection of related issues and pull requests.
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType_Issue_search_result_item represents the GeneratedType_Issue_search_result_item schema from the OpenAPI specification
type GeneratedType_Issue_search_result_item struct {
	Number int `json:"number"`
	User GeneratedType_Nullable_simple_user `json:"user"` // A GitHub user.
	Body string `json:"body,omitempty"`
	Labels_url string `json:"labels_url"`
	Timeline_url string `json:"timeline_url,omitempty"`
	Events_url string `json:"events_url"`
	Id int64 `json:"id"`
	Url string `json:"url"`
	Labels []map[string]interface{} `json:"labels"`
	Pull_request map[string]interface{} `json:"pull_request,omitempty"`
	Locked bool `json:"locked"`
	Assignee GeneratedType_Nullable_simple_user `json:"assignee"` // A GitHub user.
	Created_at string `json:"created_at"`
	Comments int `json:"comments"`
	Html_url string `json:"html_url"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Node_id string `json:"node_id"`
	Body_text string `json:"body_text,omitempty"`
	Active_lock_reason string `json:"active_lock_reason,omitempty"`
	State_reason string `json:"state_reason,omitempty"`
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Score float64 `json:"score"`
	Text_matches []map[string]interface{} `json:"text_matches,omitempty"`
	TypeField GeneratedType_Issue_type `json:"type,omitempty"` // The type of issue.
	Closed_at string `json:"closed_at"`
	Updated_at string `json:"updated_at"`
	Body_html string `json:"body_html,omitempty"`
	Performed_via_github_app GeneratedType_Nullable_integration `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Comments_url string `json:"comments_url"`
	Sub_issues_summary map[string]interface{} `json:"sub_issues_summary,omitempty"`
	Assignees []GeneratedType_Simple_user `json:"assignees,omitempty"`
	Reactions GeneratedType_Reaction_rollup `json:"reactions,omitempty"`
	Repository_url string `json:"repository_url"`
	Draft bool `json:"draft,omitempty"`
	State string `json:"state"`
	Title string `json:"title"`
	Milestone GeneratedType_Nullable_milestone `json:"milestone"` // A collection of related issues and pull requests.
}

// GeneratedType_Timeline_assigned_issue_event represents the GeneratedType_Timeline_assigned_issue_event schema from the OpenAPI specification
type GeneratedType_Timeline_assigned_issue_event struct {
	Commit_id string `json:"commit_id"`
	Event string `json:"event"`
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	Url string `json:"url"`
	Assignee GeneratedType_Simple_user `json:"assignee"` // A GitHub user.
	Commit_url string `json:"commit_url"`
	Created_at string `json:"created_at"`
	Performed_via_github_app GeneratedType_Nullable_integration `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Actor GeneratedType_Simple_user `json:"actor"` // A GitHub user.
}

// GeneratedType_Webhook_deployment_review_rejected represents the GeneratedType_Webhook_deployment_review_rejected schema from the OpenAPI specification
type GeneratedType_Webhook_deployment_review_rejected struct {
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Reviewers []map[string]interface{} `json:"reviewers,omitempty"`
	Workflow_job_run Webhooksworkflowjobrun `json:"workflow_job_run,omitempty"`
	Action string `json:"action"`
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Comment string `json:"comment,omitempty"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Since string `json:"since"`
	Workflow_job_runs []map[string]interface{} `json:"workflow_job_runs,omitempty"`
	Workflow_run map[string]interface{} `json:"workflow_run"`
	Approver Webhooksapprover `json:"approver,omitempty"`
}

// GeneratedType_Webhook_repository_edited represents the GeneratedType_Webhook_repository_edited schema from the OpenAPI specification
type GeneratedType_Webhook_repository_edited struct {
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType_Security_advisory_epss represents the GeneratedType_Security_advisory_epss schema from the OpenAPI specification
type GeneratedType_Security_advisory_epss struct {
	Percentage float64 `json:"percentage,omitempty"`
	Percentile float64 `json:"percentile,omitempty"`
}

// GeneratedType_Copilot_organization_details represents the GeneratedType_Copilot_organization_details schema from the OpenAPI specification
type GeneratedType_Copilot_organization_details struct {
	Seat_breakdown GeneratedType_Copilot_organization_seat_breakdown `json:"seat_breakdown"` // The breakdown of Copilot Business seats for the organization.
	Seat_management_setting string `json:"seat_management_setting"` // The mode of assigning new seats.
	Cli string `json:"cli,omitempty"` // The organization policy for allowing or disallowing Copilot in the CLI.
	Ide_chat string `json:"ide_chat,omitempty"` // The organization policy for allowing or disallowing Copilot Chat in the IDE.
	Plan_type string `json:"plan_type,omitempty"` // The Copilot plan of the organization, or the parent enterprise, when applicable.
	Platform_chat string `json:"platform_chat,omitempty"` // The organization policy for allowing or disallowing Copilot features on GitHub.com.
	Public_code_suggestions string `json:"public_code_suggestions"` // The organization policy for allowing or blocking suggestions matching public code (duplication detection filter).
}

// GeneratedType_Webhook_discussion_closed represents the GeneratedType_Webhook_discussion_closed schema from the OpenAPI specification
type GeneratedType_Webhook_discussion_closed struct {
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType_Timeline_line_commented_event represents the GeneratedType_Timeline_line_commented_event schema from the OpenAPI specification
type GeneratedType_Timeline_line_commented_event struct {
	Event string `json:"event,omitempty"`
	Node_id string `json:"node_id,omitempty"`
	Comments []GeneratedType_Pull_request_review_comment `json:"comments,omitempty"`
}

// GeneratedType_Webhook_pull_request_reopened represents the GeneratedType_Webhook_pull_request_reopened schema from the OpenAPI specification
type GeneratedType_Webhook_pull_request_reopened struct {
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request GeneratedType_Pull_request_webhook `json:"pull_request"`
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType_Webhook_member_added represents the GeneratedType_Webhook_member_added schema from the OpenAPI specification
type GeneratedType_Webhook_member_added struct {
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Member Webhooksuser `json:"member"`
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes,omitempty"`
}

// GeneratedType_Protected_branch_pull_request_review represents the GeneratedType_Protected_branch_pull_request_review schema from the OpenAPI specification
type GeneratedType_Protected_branch_pull_request_review struct {
	Dismissal_restrictions map[string]interface{} `json:"dismissal_restrictions,omitempty"`
	Require_code_owner_reviews bool `json:"require_code_owner_reviews"`
	Require_last_push_approval bool `json:"require_last_push_approval,omitempty"` // Whether the most recent push must be approved by someone other than the person who pushed it.
	Required_approving_review_count int `json:"required_approving_review_count,omitempty"`
	Url string `json:"url,omitempty"`
	Bypass_pull_request_allowances map[string]interface{} `json:"bypass_pull_request_allowances,omitempty"` // Allow specific users, teams, or apps to bypass pull request requirements.
	Dismiss_stale_reviews bool `json:"dismiss_stale_reviews"`
}

// GeneratedType_Webhook_label_created represents the GeneratedType_Webhook_label_created schema from the OpenAPI specification
type GeneratedType_Webhook_label_created struct {
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Label Webhookslabel `json:"label"`
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType_Webhook_organization_deleted represents the GeneratedType_Webhook_organization_deleted schema from the OpenAPI specification
type GeneratedType_Webhook_organization_deleted struct {
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Membership Webhooksmembership `json:"membership,omitempty"` // The membership between the user and the organization. Not present when the action is `member_invited`.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType_Organization_programmatic_access_grant_request represents the GeneratedType_Organization_programmatic_access_grant_request schema from the OpenAPI specification
type GeneratedType_Organization_programmatic_access_grant_request struct {
	Reason string `json:"reason"` // Reason for requesting access.
	Token_last_used_at string `json:"token_last_used_at"` // Date and time when the associated fine-grained personal access token was last used for authentication.
	Permissions map[string]interface{} `json:"permissions"` // Permissions requested, categorized by type of permission.
	Token_expired bool `json:"token_expired"` // Whether the associated fine-grained personal access token has expired.
	Token_expires_at string `json:"token_expires_at"` // Date and time when the associated fine-grained personal access token expires.
	Token_id int `json:"token_id"` // Unique identifier of the user's token. This field can also be found in audit log events and the organization's settings for their PAT grants.
	Token_name string `json:"token_name"` // The name given to the user's token. This field can also be found in an organization's settings page for Active Tokens.
	Repository_selection string `json:"repository_selection"` // Type of repository selection requested.
	Created_at string `json:"created_at"` // Date and time when the request for access was created.
	Id int `json:"id"` // Unique identifier of the request for access via fine-grained personal access token. The `pat_request_id` used to review PAT requests.
	Repositories_url string `json:"repositories_url"` // URL to the list of repositories requested to be accessed via fine-grained personal access token. Should only be followed when `repository_selection` is `subset`.
	Owner GeneratedType_Simple_user `json:"owner"` // A GitHub user.
}

// Actor represents the Actor schema from the OpenAPI specification
type Actor struct {
	Display_login string `json:"display_login,omitempty"`
	Gravatar_id string `json:"gravatar_id"`
	Id int `json:"id"`
	Login string `json:"login"`
	Url string `json:"url"`
	Avatar_url string `json:"avatar_url"`
}

// GeneratedType_Code_search_result_item represents the GeneratedType_Code_search_result_item schema from the OpenAPI specification
type GeneratedType_Code_search_result_item struct {
	Sha string `json:"sha"`
	Language string `json:"language,omitempty"`
	Line_numbers []string `json:"line_numbers,omitempty"`
	Score float64 `json:"score"`
	Url string `json:"url"`
	File_size int `json:"file_size,omitempty"`
	Git_url string `json:"git_url"`
	Path string `json:"path"`
	Repository GeneratedType_Minimal_repository `json:"repository"` // Minimal Repository
	Text_matches []map[string]interface{} `json:"text_matches,omitempty"`
	Html_url string `json:"html_url"`
	Name string `json:"name"`
	Last_modified_at string `json:"last_modified_at,omitempty"`
}

// GeneratedType_Code_scanning_alert_rule_summary represents the GeneratedType_Code_scanning_alert_rule_summary schema from the OpenAPI specification
type GeneratedType_Code_scanning_alert_rule_summary struct {
	Description string `json:"description,omitempty"` // A short description of the rule used to detect the alert.
	Full_description string `json:"full_description,omitempty"` // A description of the rule used to detect the alert.
	Id string `json:"id,omitempty"` // A unique identifier for the rule used to detect the alert.
	Name string `json:"name,omitempty"` // The name of the rule used to detect the alert.
	Security_severity_level string `json:"security_severity_level,omitempty"` // The security severity of the alert.
	Severity string `json:"severity,omitempty"` // The severity of the alert.
	Help string `json:"help,omitempty"` // Detailed documentation for the rule as GitHub Flavored Markdown.
	Help_uri string `json:"help_uri,omitempty"` // A link to the documentation for the rule used to detect the alert.
	Tags []string `json:"tags,omitempty"` // A set of tags applicable for the rule.
}

// Blob represents the Blob schema from the OpenAPI specification
type Blob struct {
	Node_id string `json:"node_id"`
	Sha string `json:"sha"`
	Size int `json:"size"`
	Url string `json:"url"`
	Content string `json:"content"`
	Encoding string `json:"encoding"`
	Highlighted_content string `json:"highlighted_content,omitempty"`
}

// Authorization represents the Authorization schema from the OpenAPI specification
type Authorization struct {
	Expires_at string `json:"expires_at"`
	Scopes []string `json:"scopes"` // A list of scopes that this authorization is in.
	Token_last_eight string `json:"token_last_eight"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	Installation GeneratedType_Nullable_scoped_installation `json:"installation,omitempty"`
	Fingerprint string `json:"fingerprint"`
	Hashed_token string `json:"hashed_token"`
	Created_at string `json:"created_at"`
	User GeneratedType_Nullable_simple_user `json:"user,omitempty"` // A GitHub user.
	Token string `json:"token"`
	Note_url string `json:"note_url"`
	Note string `json:"note"`
	Id int64 `json:"id"`
	App map[string]interface{} `json:"app"`
}

// GeneratedType_Oidc_custom_sub_repo represents the GeneratedType_Oidc_custom_sub_repo schema from the OpenAPI specification
type GeneratedType_Oidc_custom_sub_repo struct {
	Include_claim_keys []string `json:"include_claim_keys,omitempty"` // Array of unique strings. Each claim key can only contain alphanumeric characters and underscores.
	Use_default bool `json:"use_default"` // Whether to use the default template or not. If `true`, the `include_claim_keys` field is ignored.
}

// GeneratedType_Porter_large_file represents the GeneratedType_Porter_large_file schema from the OpenAPI specification
type GeneratedType_Porter_large_file struct {
	Oid string `json:"oid"`
	Path string `json:"path"`
	Ref_name string `json:"ref_name"`
	Size int `json:"size"`
}

// GeneratedType_Team_repository represents the GeneratedType_Team_repository schema from the OpenAPI specification
type GeneratedType_Team_repository struct {
	Node_id string `json:"node_id"`
	Keys_url string `json:"keys_url"`
	Allow_auto_merge bool `json:"allow_auto_merge,omitempty"` // Whether to allow Auto-merge to be used on pull requests.
	Ssh_url string `json:"ssh_url"`
	Role_name string `json:"role_name,omitempty"`
	Branches_url string `json:"branches_url"`
	Has_downloads bool `json:"has_downloads"` // Whether downloads are enabled.
	Comments_url string `json:"comments_url"`
	Pushed_at string `json:"pushed_at"`
	Milestones_url string `json:"milestones_url"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Temp_clone_token string `json:"temp_clone_token,omitempty"`
	Id int `json:"id"` // Unique identifier of the repository
	Subscribers_url string `json:"subscribers_url"`
	Git_tags_url string `json:"git_tags_url"`
	Url string `json:"url"`
	Blobs_url string `json:"blobs_url"`
	Languages_url string `json:"languages_url"`
	Open_issues int `json:"open_issues"`
	Events_url string `json:"events_url"`
	Is_template bool `json:"is_template,omitempty"` // Whether this repository acts as a template that can be used to generate new repositories.
	Issue_comment_url string `json:"issue_comment_url"`
	Archive_url string `json:"archive_url"`
	Has_issues bool `json:"has_issues"` // Whether issues are enabled.
	Open_issues_count int `json:"open_issues_count"`
	Contributors_url string `json:"contributors_url"`
	Has_projects bool `json:"has_projects"` // Whether projects are enabled.
	Private bool `json:"private"` // Whether the repository is private or public.
	Name string `json:"name"` // The name of the repository.
	Forks_url string `json:"forks_url"`
	Assignees_url string `json:"assignees_url"`
	Updated_at string `json:"updated_at"`
	Created_at string `json:"created_at"`
	Language string `json:"language"`
	Subscription_url string `json:"subscription_url"`
	Forks int `json:"forks"`
	Mirror_url string `json:"mirror_url"`
	Teams_url string `json:"teams_url"`
	Svn_url string `json:"svn_url"`
	Has_pages bool `json:"has_pages"`
	Default_branch string `json:"default_branch"` // The default branch of the repository.
	Watchers int `json:"watchers"`
	Size int `json:"size"`
	Visibility string `json:"visibility,omitempty"` // The repository visibility: public, private, or internal.
	Html_url string `json:"html_url"`
	Issue_events_url string `json:"issue_events_url"`
	Downloads_url string `json:"downloads_url"`
	Git_refs_url string `json:"git_refs_url"`
	Commits_url string `json:"commits_url"`
	Git_commits_url string `json:"git_commits_url"`
	Allow_rebase_merge bool `json:"allow_rebase_merge,omitempty"` // Whether to allow rebase merges for pull requests.
	Deployments_url string `json:"deployments_url"`
	Owner GeneratedType_Nullable_simple_user `json:"owner"` // A GitHub user.
	Forks_count int `json:"forks_count"`
	Homepage string `json:"homepage"`
	Clone_url string `json:"clone_url"`
	Compare_url string `json:"compare_url"`
	Topics []string `json:"topics,omitempty"`
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"` // Whether to delete head branches when pull requests are merged
	Disabled bool `json:"disabled"` // Returns whether or not this repository disabled.
	Has_wiki bool `json:"has_wiki"` // Whether the wiki is enabled.
	Allow_merge_commit bool `json:"allow_merge_commit,omitempty"` // Whether to allow merge commits for pull requests.
	Statuses_url string `json:"statuses_url"`
	License GeneratedType_Nullable_license_simple `json:"license"` // License Simple
	Collaborators_url string `json:"collaborators_url"`
	Description string `json:"description"`
	Allow_forking bool `json:"allow_forking,omitempty"` // Whether to allow forking this repo
	Fork bool `json:"fork"`
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"` // Whether to require contributors to sign off on web-based commits
	Releases_url string `json:"releases_url"`
	Git_url string `json:"git_url"`
	Labels_url string `json:"labels_url"`
	Contents_url string `json:"contents_url"`
	Notifications_url string `json:"notifications_url"`
	Pulls_url string `json:"pulls_url"`
	Allow_squash_merge bool `json:"allow_squash_merge,omitempty"` // Whether to allow squash merges for pull requests.
	Stargazers_url string `json:"stargazers_url"`
	Tags_url string `json:"tags_url"`
	Network_count int `json:"network_count,omitempty"`
	Subscribers_count int `json:"subscribers_count,omitempty"`
	Watchers_count int `json:"watchers_count"`
	Full_name string `json:"full_name"`
	Issues_url string `json:"issues_url"`
	Merges_url string `json:"merges_url"`
	Trees_url string `json:"trees_url"`
	Hooks_url string `json:"hooks_url"`
	Stargazers_count int `json:"stargazers_count"`
	Archived bool `json:"archived"` // Whether the repository is archived.
	Master_branch string `json:"master_branch,omitempty"`
}

// GeneratedType_Webhook_sponsorship_edited represents the GeneratedType_Webhook_sponsorship_edited schema from the OpenAPI specification
type GeneratedType_Webhook_sponsorship_edited struct {
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Sponsorship Webhookssponsorship `json:"sponsorship"`
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
}

// GeneratedType_Webhook_check_run_requested_action_form_encoded represents the GeneratedType_Webhook_check_run_requested_action_form_encoded schema from the OpenAPI specification
type GeneratedType_Webhook_check_run_requested_action_form_encoded struct {
	Payload string `json:"payload"` // A URL-encoded string of the check_run.requested_action JSON payload. The decoded payload is a JSON object.
}

// GeneratedType_Webhook_pull_request_opened represents the GeneratedType_Webhook_pull_request_opened schema from the OpenAPI specification
type GeneratedType_Webhook_pull_request_opened struct {
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request GeneratedType_Pull_request_webhook `json:"pull_request"`
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType_Webhook_secret_scanning_alert_publicly_leaked represents the GeneratedType_Webhook_secret_scanning_alert_publicly_leaked schema from the OpenAPI specification
type GeneratedType_Webhook_secret_scanning_alert_publicly_leaked struct {
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Alert GeneratedType_Secret_scanning_alert_webhook `json:"alert"`
}

// GeneratedType_Webhook_pull_request_review_submitted represents the GeneratedType_Webhook_pull_request_review_submitted schema from the OpenAPI specification
type GeneratedType_Webhook_pull_request_review_submitted struct {
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Review Webhooksreview `json:"review"` // The review that was affected.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request map[string]interface{} `json:"pull_request"`
}

// GeneratedType_Actions_secret represents the GeneratedType_Actions_secret schema from the OpenAPI specification
type GeneratedType_Actions_secret struct {
	Created_at string `json:"created_at"`
	Name string `json:"name"` // The name of the secret.
	Updated_at string `json:"updated_at"`
}

// GeneratedType_Webhook_dependabot_alert_auto_reopened represents the GeneratedType_Webhook_dependabot_alert_auto_reopened schema from the OpenAPI specification
type GeneratedType_Webhook_dependabot_alert_auto_reopened struct {
	Action string `json:"action"`
	Alert GeneratedType_Dependabot_alert `json:"alert"` // A Dependabot alert.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
}

// Webhooksmilestone represents the Webhooksmilestone schema from the OpenAPI specification
type Webhooksmilestone struct {
	State string `json:"state"` // The state of the milestone.
	Closed_issues int `json:"closed_issues"`
	Created_at string `json:"created_at"`
	Id int `json:"id"`
	Html_url string `json:"html_url"`
	Url string `json:"url"`
	Labels_url string `json:"labels_url"`
	Due_on string `json:"due_on"`
	Updated_at string `json:"updated_at"`
	Node_id string `json:"node_id"`
	Open_issues int `json:"open_issues"`
	Title string `json:"title"` // The title of the milestone.
	Closed_at string `json:"closed_at"`
	Creator map[string]interface{} `json:"creator"`
	Description string `json:"description"`
	Number int `json:"number"` // The number of the milestone.
}

// GeneratedType_Actions_cache_list represents the GeneratedType_Actions_cache_list schema from the OpenAPI specification
type GeneratedType_Actions_cache_list struct {
	Actions_caches []map[string]interface{} `json:"actions_caches"` // Array of caches
	Total_count int `json:"total_count"` // Total number of caches
}

// GeneratedType_Webhook_project_created represents the GeneratedType_Webhook_project_created schema from the OpenAPI specification
type GeneratedType_Webhook_project_created struct {
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Project Webhooksproject `json:"project"`
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType_Dependabot_public_key represents the GeneratedType_Dependabot_public_key schema from the OpenAPI specification
type GeneratedType_Dependabot_public_key struct {
	Key_id string `json:"key_id"` // The identifier for the key.
	Key string `json:"key"` // The Base64 encoded public key.
}

// GeneratedType_Webhook_pull_request_review_dismissed represents the GeneratedType_Webhook_pull_request_review_dismissed schema from the OpenAPI specification
type GeneratedType_Webhook_pull_request_review_dismissed struct {
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Review map[string]interface{} `json:"review"` // The review that was affected.
}

// GeneratedType_Webhook_pull_request_review_comment_deleted represents the GeneratedType_Webhook_pull_request_review_comment_deleted schema from the OpenAPI specification
type GeneratedType_Webhook_pull_request_review_comment_deleted struct {
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Comment Webhooksreviewcomment `json:"comment"` // The [comment](https://docs.github.com/rest/pulls/comments#get-a-review-comment-for-a-pull-request) itself.
}

// GeneratedType_Simple_check_suite represents the GeneratedType_Simple_check_suite schema from the OpenAPI specification
type GeneratedType_Simple_check_suite struct {
	Created_at string `json:"created_at,omitempty"`
	Head_sha string `json:"head_sha,omitempty"` // The SHA of the head commit that is being checked.
	Node_id string `json:"node_id,omitempty"`
	Conclusion string `json:"conclusion,omitempty"`
	Id int `json:"id,omitempty"`
	Updated_at string `json:"updated_at,omitempty"`
	Before string `json:"before,omitempty"`
	Head_branch string `json:"head_branch,omitempty"`
	App Integration `json:"app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Pull_requests []GeneratedType_Pull_request_minimal `json:"pull_requests,omitempty"`
	Repository GeneratedType_Minimal_repository `json:"repository,omitempty"` // Minimal Repository
	Status string `json:"status,omitempty"`
	Url string `json:"url,omitempty"`
	After string `json:"after,omitempty"`
}

// GeneratedType_Branch_restriction_policy represents the GeneratedType_Branch_restriction_policy schema from the OpenAPI specification
type GeneratedType_Branch_restriction_policy struct {
	Users_url string `json:"users_url"`
	Apps []map[string]interface{} `json:"apps"`
	Apps_url string `json:"apps_url"`
	Teams []map[string]interface{} `json:"teams"`
	Teams_url string `json:"teams_url"`
	Url string `json:"url"`
	Users []map[string]interface{} `json:"users"`
}

// GeneratedType_Actions_cache_usage_by_repository represents the GeneratedType_Actions_cache_usage_by_repository schema from the OpenAPI specification
type GeneratedType_Actions_cache_usage_by_repository struct {
	Active_caches_size_in_bytes int `json:"active_caches_size_in_bytes"` // The sum of the size in bytes of all the active cache items in the repository.
	Full_name string `json:"full_name"` // The repository owner and name for the cache usage being shown.
	Active_caches_count int `json:"active_caches_count"` // The number of active caches in the repository.
}

// Repository represents the Repository schema from the OpenAPI specification
type Repository struct {
	Pulls_url string `json:"pulls_url"`
	Git_url string `json:"git_url"`
	Teams_url string `json:"teams_url"`
	Forks int `json:"forks"`
	Commits_url string `json:"commits_url"`
	Has_downloads bool `json:"has_downloads"` // Whether downloads are enabled.
	Git_commits_url string `json:"git_commits_url"`
	Contents_url string `json:"contents_url"`
	Branches_url string `json:"branches_url"`
	Stargazers_url string `json:"stargazers_url"`
	Visibility string `json:"visibility,omitempty"` // The repository visibility: public, private, or internal.
	Allow_squash_merge bool `json:"allow_squash_merge,omitempty"` // Whether to allow squash merges for pull requests.
	Forks_count int `json:"forks_count"`
	Name string `json:"name"` // The name of the repository.
	Merge_commit_title string `json:"merge_commit_title,omitempty"` // The default value for a merge commit title. - `PR_TITLE` - default to the pull request's title. - `MERGE_MESSAGE` - default to the classic title for a merge message (e.g., Merge pull request #123 from branch-name).
	Ssh_url string `json:"ssh_url"`
	Open_issues int `json:"open_issues"`
	Subscribers_url string `json:"subscribers_url"`
	Collaborators_url string `json:"collaborators_url"`
	Issues_url string `json:"issues_url"`
	Has_wiki bool `json:"has_wiki"` // Whether the wiki is enabled.
	Updated_at string `json:"updated_at"`
	Watchers_count int `json:"watchers_count"`
	Disabled bool `json:"disabled"` // Returns whether or not this repository disabled.
	Open_issues_count int `json:"open_issues_count"`
	Keys_url string `json:"keys_url"`
	Default_branch string `json:"default_branch"` // The default branch of the repository.
	Languages_url string `json:"languages_url"`
	Notifications_url string `json:"notifications_url"`
	Squash_merge_commit_title string `json:"squash_merge_commit_title,omitempty"` // The default value for a squash merge commit title: - `PR_TITLE` - default to the pull request's title. - `COMMIT_OR_PR_TITLE` - default to the commit's title (if only one commit) or the pull request's title (when more than one commit).
	Blobs_url string `json:"blobs_url"`
	Created_at string `json:"created_at"`
	Starred_at string `json:"starred_at,omitempty"`
	Git_refs_url string `json:"git_refs_url"`
	Allow_forking bool `json:"allow_forking,omitempty"` // Whether to allow forking this repo
	Anonymous_access_enabled bool `json:"anonymous_access_enabled,omitempty"` // Whether anonymous git access is enabled for this repository
	Issue_events_url string `json:"issue_events_url"`
	Deployments_url string `json:"deployments_url"`
	Compare_url string `json:"compare_url"`
	Has_discussions bool `json:"has_discussions,omitempty"` // Whether discussions are enabled.
	Mirror_url string `json:"mirror_url"`
	Allow_rebase_merge bool `json:"allow_rebase_merge,omitempty"` // Whether to allow rebase merges for pull requests.
	Forks_url string `json:"forks_url"`
	Has_pages bool `json:"has_pages"`
	Has_issues bool `json:"has_issues"` // Whether issues are enabled.
	Git_tags_url string `json:"git_tags_url"`
	Use_squash_pr_title_as_default bool `json:"use_squash_pr_title_as_default,omitempty"` // Whether a squash merge commit can use the pull request title as default. **This property is closing down. Please use `squash_merge_commit_title` instead.
	Hooks_url string `json:"hooks_url"`
	Has_projects bool `json:"has_projects"` // Whether projects are enabled.
	Is_template bool `json:"is_template,omitempty"` // Whether this repository acts as a template that can be used to generate new repositories.
	License GeneratedType_Nullable_license_simple `json:"license"` // License Simple
	Id int64 `json:"id"` // Unique identifier of the repository
	Assignees_url string `json:"assignees_url"`
	Labels_url string `json:"labels_url"`
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"` // Whether to delete head branches when pull requests are merged
	Clone_url string `json:"clone_url"`
	Full_name string `json:"full_name"`
	Squash_merge_commit_message string `json:"squash_merge_commit_message,omitempty"` // The default value for a squash merge commit message: - `PR_BODY` - default to the pull request's body. - `COMMIT_MESSAGES` - default to the branch's commit messages. - `BLANK` - default to a blank commit message.
	Stargazers_count int `json:"stargazers_count"`
	Archived bool `json:"archived"` // Whether the repository is archived.
	Comments_url string `json:"comments_url"`
	Statuses_url string `json:"statuses_url"`
	Trees_url string `json:"trees_url"`
	Allow_update_branch bool `json:"allow_update_branch,omitempty"` // Whether or not a pull request head branch that is behind its base branch can always be updated even if it is not required to be up to date before merging.
	Pushed_at string `json:"pushed_at"`
	Master_branch string `json:"master_branch,omitempty"`
	Milestones_url string `json:"milestones_url"`
	Homepage string `json:"homepage"`
	Allow_merge_commit bool `json:"allow_merge_commit,omitempty"` // Whether to allow merge commits for pull requests.
	Watchers int `json:"watchers"`
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"` // Whether to require contributors to sign off on web-based commits
	Code_search_index_status map[string]interface{} `json:"code_search_index_status,omitempty"` // The status of the code search index for this repository
	Private bool `json:"private"` // Whether the repository is private or public.
	Merge_commit_message string `json:"merge_commit_message,omitempty"` // The default value for a merge commit message. - `PR_TITLE` - default to the pull request's title. - `PR_BODY` - default to the pull request's body. - `BLANK` - default to a blank commit message.
	Archive_url string `json:"archive_url"`
	Downloads_url string `json:"downloads_url"`
	Svn_url string `json:"svn_url"`
	Events_url string `json:"events_url"`
	Owner GeneratedType_Simple_user `json:"owner"` // A GitHub user.
	Temp_clone_token string `json:"temp_clone_token,omitempty"`
	Fork bool `json:"fork"`
	Allow_auto_merge bool `json:"allow_auto_merge,omitempty"` // Whether to allow Auto-merge to be used on pull requests.
	Html_url string `json:"html_url"`
	Tags_url string `json:"tags_url"`
	Language string `json:"language"`
	Description string `json:"description"`
	Node_id string `json:"node_id"`
	Releases_url string `json:"releases_url"`
	Url string `json:"url"`
	Issue_comment_url string `json:"issue_comment_url"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Merges_url string `json:"merges_url"`
	Size int `json:"size"` // The size of the repository, in kilobytes. Size is calculated hourly. When a repository is initially created, the size is 0.
	Contributors_url string `json:"contributors_url"`
	Topics []string `json:"topics,omitempty"`
	Subscription_url string `json:"subscription_url"`
}

// GeneratedType_Webhook_project_reopened represents the GeneratedType_Webhook_project_reopened schema from the OpenAPI specification
type GeneratedType_Webhook_project_reopened struct {
	Project Webhooksproject `json:"project"`
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType_Repository_invitation represents the GeneratedType_Repository_invitation schema from the OpenAPI specification
type GeneratedType_Repository_invitation struct {
	Created_at string `json:"created_at"`
	Permissions string `json:"permissions"` // The permission associated with the invitation.
	Repository GeneratedType_Minimal_repository `json:"repository"` // Minimal Repository
	Expired bool `json:"expired,omitempty"` // Whether or not the invitation has expired
	Invitee GeneratedType_Nullable_simple_user `json:"invitee"` // A GitHub user.
	Inviter GeneratedType_Nullable_simple_user `json:"inviter"` // A GitHub user.
	Html_url string `json:"html_url"`
	Id int64 `json:"id"` // Unique identifier of the repository invitation.
	Node_id string `json:"node_id"`
	Url string `json:"url"` // URL for the repository invitation
}

// GeneratedType_Repository_rule_non_fast_forward represents the GeneratedType_Repository_rule_non_fast_forward schema from the OpenAPI specification
type GeneratedType_Repository_rule_non_fast_forward struct {
	TypeField string `json:"type"`
}

// GeneratedType_Deployment_branch_policy_name_pattern represents the GeneratedType_Deployment_branch_policy_name_pattern schema from the OpenAPI specification
type GeneratedType_Deployment_branch_policy_name_pattern struct {
	Name string `json:"name"` // The name pattern that branches must match in order to deploy to the environment. Wildcard characters will not match `/`. For example, to match branches that begin with `release/` and contain an additional single slash, use `release/*/*`. For more information about pattern matching syntax, see the [Ruby File.fnmatch documentation](https://ruby-doc.org/core-2.5.1/File.html#method-c-fnmatch).
}

// Release represents the Release schema from the OpenAPI specification
type Release struct {
	Node_id string `json:"node_id"`
	Url string `json:"url"`
	Html_url string `json:"html_url"`
	Target_commitish string `json:"target_commitish"` // Specifies the commitish value that determines where the Git tag is created from.
	Tag_name string `json:"tag_name"` // The name of the tag.
	Body_text string `json:"body_text,omitempty"`
	Created_at string `json:"created_at"`
	Draft bool `json:"draft"` // true to create a draft (unpublished) release, false to create a published one.
	Upload_url string `json:"upload_url"`
	Author GeneratedType_Simple_user `json:"author"` // A GitHub user.
	Mentions_count int `json:"mentions_count,omitempty"`
	Published_at string `json:"published_at"`
	Tarball_url string `json:"tarball_url"`
	Body_html string `json:"body_html,omitempty"`
	Discussion_url string `json:"discussion_url,omitempty"` // The URL of the release discussion.
	Id int `json:"id"`
	Name string `json:"name"`
	Assets []GeneratedType_Release_asset `json:"assets"`
	Zipball_url string `json:"zipball_url"`
	Body string `json:"body,omitempty"`
	Prerelease bool `json:"prerelease"` // Whether to identify the release as a prerelease or a full release.
	Assets_url string `json:"assets_url"`
	Reactions GeneratedType_Reaction_rollup `json:"reactions,omitempty"`
}

// GeneratedType_Codespaces_secret represents the GeneratedType_Codespaces_secret schema from the OpenAPI specification
type GeneratedType_Codespaces_secret struct {
	Created_at string `json:"created_at"` // The date and time at which the secret was created, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Name string `json:"name"` // The name of the secret
	Selected_repositories_url string `json:"selected_repositories_url"` // The API URL at which the list of repositories this secret is visible to can be retrieved
	Updated_at string `json:"updated_at"` // The date and time at which the secret was last updated, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Visibility string `json:"visibility"` // The type of repositories in the organization that the secret is visible to
}

// GeneratedType_Organization_actions_variable represents the GeneratedType_Organization_actions_variable schema from the OpenAPI specification
type GeneratedType_Organization_actions_variable struct {
	Updated_at string `json:"updated_at"` // The date and time at which the variable was last updated, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Value string `json:"value"` // The value of the variable.
	Visibility string `json:"visibility"` // Visibility of a variable
	Created_at string `json:"created_at"` // The date and time at which the variable was created, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Name string `json:"name"` // The name of the variable.
	Selected_repositories_url string `json:"selected_repositories_url,omitempty"`
}

// GeneratedType_Webhook_issues_pinned represents the GeneratedType_Webhook_issues_pinned schema from the OpenAPI specification
type GeneratedType_Webhook_issues_pinned struct {
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue Webhooksissue2 `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
}

// GeneratedType_Webhook_pull_request_synchronize represents the GeneratedType_Webhook_pull_request_synchronize schema from the OpenAPI specification
type GeneratedType_Webhook_pull_request_synchronize struct {
	Before string `json:"before"`
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Number int `json:"number"` // The pull request number.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	After string `json:"after"`
}

// GeneratedType_Webhook_merge_group_checks_requested represents the GeneratedType_Webhook_merge_group_checks_requested schema from the OpenAPI specification
type GeneratedType_Webhook_merge_group_checks_requested struct {
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Merge_group GeneratedType_Merge_group `json:"merge_group"` // A group of pull requests that the merge queue has grouped together to be merged.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType_Webhook_discussion_comment_deleted represents the GeneratedType_Webhook_discussion_comment_deleted schema from the OpenAPI specification
type GeneratedType_Webhook_discussion_comment_deleted struct {
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Comment Webhookscomment `json:"comment"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// Webhooksproject represents the Webhooksproject schema from the OpenAPI specification
type Webhooksproject struct {
	Updated_at string `json:"updated_at"`
	Columns_url string `json:"columns_url"`
	Created_at string `json:"created_at"`
	Creator map[string]interface{} `json:"creator"`
	Html_url string `json:"html_url"`
	Id int `json:"id"`
	State string `json:"state"` // State of the project; either 'open' or 'closed'
	Body string `json:"body"` // Body of the project
	Owner_url string `json:"owner_url"`
	Url string `json:"url"`
	Name string `json:"name"` // Name of the project
	Node_id string `json:"node_id"`
	Number int `json:"number"`
}

// GeneratedType_Repository_rule_workflows represents the GeneratedType_Repository_rule_workflows schema from the OpenAPI specification
type GeneratedType_Repository_rule_workflows struct {
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	TypeField string `json:"type"`
}

// GeneratedType_Review_custom_gates_state_required represents the GeneratedType_Review_custom_gates_state_required schema from the OpenAPI specification
type GeneratedType_Review_custom_gates_state_required struct {
	State string `json:"state"` // Whether to approve or reject deployment to the specified environments.
	Comment string `json:"comment,omitempty"` // Optional comment to include with the review.
	Environment_name string `json:"environment_name"` // The name of the environment to approve or reject.
}

// GeneratedType_Campaign_summary represents the GeneratedType_Campaign_summary schema from the OpenAPI specification
type GeneratedType_Campaign_summary struct {
	Number int `json:"number"` // The number of the newly created campaign
	Published_at string `json:"published_at,omitempty"` // The date and time the campaign was published, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Updated_at string `json:"updated_at"` // The date and time the campaign was last updated, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Contact_link string `json:"contact_link"` // The contact link of the campaign.
	Ends_at string `json:"ends_at"` // The date and time the campaign has ended, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Managers []GeneratedType_Simple_user `json:"managers"` // The campaign managers
	State string `json:"state"` // Indicates whether a campaign is open or closed
	Team_managers []Team `json:"team_managers,omitempty"` // The campaign team managers
	Alert_stats map[string]interface{} `json:"alert_stats,omitempty"`
	Closed_at string `json:"closed_at,omitempty"` // The date and time the campaign was closed, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ. Will be null if the campaign is still open.
	Created_at string `json:"created_at"` // The date and time the campaign was created, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Description string `json:"description"` // The campaign description
	Name string `json:"name,omitempty"` // The campaign name
}

// GeneratedType_Webhook_code_scanning_alert_reopened represents the GeneratedType_Webhook_code_scanning_alert_reopened schema from the OpenAPI specification
type GeneratedType_Webhook_code_scanning_alert_reopened struct {
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Ref string `json:"ref"` // The Git reference of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Alert map[string]interface{} `json:"alert"` // The code scanning alert involved in the event.
	Commit_oid string `json:"commit_oid"` // The commit SHA of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType_Repository_rule_required_linear_history represents the GeneratedType_Repository_rule_required_linear_history schema from the OpenAPI specification
type GeneratedType_Repository_rule_required_linear_history struct {
	TypeField string `json:"type"`
}

// GeneratedType_Webhook_sponsorship_tier_changed represents the GeneratedType_Webhook_sponsorship_tier_changed schema from the OpenAPI specification
type GeneratedType_Webhook_sponsorship_tier_changed struct {
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Sponsorship Webhookssponsorship `json:"sponsorship"`
	Action string `json:"action"`
	Changes Webhookschanges8 `json:"changes"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType_Webhook_issue_comment_edited represents the GeneratedType_Webhook_issue_comment_edited schema from the OpenAPI specification
type GeneratedType_Webhook_issue_comment_edited struct {
	Comment Webhooksissuecomment `json:"comment"` // The [comment](https://docs.github.com/rest/issues/comments#get-an-issue-comment) itself.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Issue interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) the comment belongs to.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Action string `json:"action"`
	Changes Webhookschanges `json:"changes"` // The changes to the comment.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType_Projects_v2 represents the GeneratedType_Projects_v2 schema from the OpenAPI specification
type GeneratedType_Projects_v2 struct {
	Title string `json:"title"`
	Updated_at string `json:"updated_at"`
	Closed_at string `json:"closed_at"`
	Creator GeneratedType_Simple_user `json:"creator"` // A GitHub user.
	Deleted_at string `json:"deleted_at"`
	Deleted_by GeneratedType_Nullable_simple_user `json:"deleted_by"` // A GitHub user.
	Node_id string `json:"node_id"`
	Number int `json:"number"`
	Short_description string `json:"short_description"`
	Created_at string `json:"created_at"`
	Id float64 `json:"id"`
	Description string `json:"description"`
	Owner GeneratedType_Simple_user `json:"owner"` // A GitHub user.
	Public bool `json:"public"`
}

// GeneratedType_Webhook_pull_request_assigned represents the GeneratedType_Webhook_pull_request_assigned schema from the OpenAPI specification
type GeneratedType_Webhook_pull_request_assigned struct {
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Assignee Webhooksuser `json:"assignee"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Action string `json:"action"`
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType_Review_request_removed_issue_event represents the GeneratedType_Review_request_removed_issue_event schema from the OpenAPI specification
type GeneratedType_Review_request_removed_issue_event struct {
	Created_at string `json:"created_at"`
	Actor GeneratedType_Simple_user `json:"actor"` // A GitHub user.
	Commit_id string `json:"commit_id"`
	Requested_reviewer GeneratedType_Simple_user `json:"requested_reviewer,omitempty"` // A GitHub user.
	Requested_team Team `json:"requested_team,omitempty"` // Groups of organization members that gives permissions on specified repositories.
	Review_requester GeneratedType_Simple_user `json:"review_requester"` // A GitHub user.
	Event string `json:"event"`
	Node_id string `json:"node_id"`
	Performed_via_github_app GeneratedType_Nullable_integration `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Url string `json:"url"`
	Id int `json:"id"`
	Commit_url string `json:"commit_url"`
}

// GeneratedType_Webhook_custom_property_values_updated represents the GeneratedType_Webhook_custom_property_values_updated schema from the OpenAPI specification
type GeneratedType_Webhook_custom_property_values_updated struct {
	New_property_values []GeneratedType_Custom_property_value `json:"new_property_values"` // The new custom property values for the repository.
	Old_property_values []GeneratedType_Custom_property_value `json:"old_property_values"` // The old custom property values for the repository.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType_Webhook_team_deleted represents the GeneratedType_Webhook_team_deleted schema from the OpenAPI specification
type GeneratedType_Webhook_team_deleted struct {
	Repository map[string]interface{} `json:"repository,omitempty"` // A git repository
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Team Webhooksteam1 `json:"team"` // Groups of organization members that gives permissions on specified repositories.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// Artifact represents the Artifact schema from the OpenAPI specification
type Artifact struct {
	Size_in_bytes int `json:"size_in_bytes"` // The size in bytes of the artifact.
	Updated_at string `json:"updated_at"`
	Created_at string `json:"created_at"`
	Workflow_run map[string]interface{} `json:"workflow_run,omitempty"`
	Archive_download_url string `json:"archive_download_url"`
	Node_id string `json:"node_id"`
	Url string `json:"url"`
	Digest string `json:"digest,omitempty"` // The SHA256 digest of the artifact. This field will only be populated on artifacts uploaded with upload-artifact v4 or newer. For older versions, this field will be null.
	Expired bool `json:"expired"` // Whether or not the artifact has expired.
	Expires_at string `json:"expires_at"`
	Id int `json:"id"`
	Name string `json:"name"` // The name of the artifact.
}

// GeneratedType_Repository_ruleset_bypass_actor represents the GeneratedType_Repository_ruleset_bypass_actor schema from the OpenAPI specification
type GeneratedType_Repository_ruleset_bypass_actor struct {
	Actor_id int `json:"actor_id,omitempty"` // The ID of the actor that can bypass a ruleset. If `actor_type` is `OrganizationAdmin`, this should be `1`. If `actor_type` is `DeployKey`, this should be null. `OrganizationAdmin` is not applicable for personal repositories.
	Actor_type string `json:"actor_type"` // The type of actor that can bypass a ruleset.
	Bypass_mode string `json:"bypass_mode,omitempty"` // When the specified actor can bypass the ruleset. `pull_request` means that an actor can only bypass rules on pull requests. `pull_request` is not applicable for the `DeployKey` actor type. Also, `pull_request` is only applicable to branch rulesets.
}

// Webhooksapprover represents the Webhooksapprover schema from the OpenAPI specification
type Webhooksapprover struct {
	Gravatar_id string `json:"gravatar_id,omitempty"`
	Repos_url string `json:"repos_url,omitempty"`
	Following_url string `json:"following_url,omitempty"`
	Node_id string `json:"node_id,omitempty"`
	Login string `json:"login,omitempty"`
	Site_admin bool `json:"site_admin,omitempty"`
	Followers_url string `json:"followers_url,omitempty"`
	Subscriptions_url string `json:"subscriptions_url,omitempty"`
	Html_url string `json:"html_url,omitempty"`
	Starred_url string `json:"starred_url,omitempty"`
	Received_events_url string `json:"received_events_url,omitempty"`
	Gists_url string `json:"gists_url,omitempty"`
	Url string `json:"url,omitempty"`
	User_view_type string `json:"user_view_type,omitempty"`
	Organizations_url string `json:"organizations_url,omitempty"`
	Avatar_url string `json:"avatar_url,omitempty"`
	Events_url string `json:"events_url,omitempty"`
	Id int `json:"id,omitempty"`
	TypeField string `json:"type,omitempty"`
}

// GeneratedType_Nullable_repository represents the GeneratedType_Nullable_repository schema from the OpenAPI specification
type GeneratedType_Nullable_repository struct {
	Milestones_url string `json:"milestones_url"`
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"` // Whether to delete head branches when pull requests are merged
	Keys_url string `json:"keys_url"`
	Merge_commit_message string `json:"merge_commit_message,omitempty"` // The default value for a merge commit message. - `PR_TITLE` - default to the pull request's title. - `PR_BODY` - default to the pull request's body. - `BLANK` - default to a blank commit message.
	Has_discussions bool `json:"has_discussions,omitempty"` // Whether discussions are enabled.
	Svn_url string `json:"svn_url"`
	Archived bool `json:"archived"` // Whether the repository is archived.
	Allow_squash_merge bool `json:"allow_squash_merge,omitempty"` // Whether to allow squash merges for pull requests.
	Forks int `json:"forks"`
	Commits_url string `json:"commits_url"`
	Language string `json:"language"`
	Topics []string `json:"topics,omitempty"`
	Owner GeneratedType_Simple_user `json:"owner"` // A GitHub user.
	Visibility string `json:"visibility,omitempty"` // The repository visibility: public, private, or internal.
	Master_branch string `json:"master_branch,omitempty"`
	Code_search_index_status map[string]interface{} `json:"code_search_index_status,omitempty"` // The status of the code search index for this repository
	Contributors_url string `json:"contributors_url"`
	Is_template bool `json:"is_template,omitempty"` // Whether this repository acts as a template that can be used to generate new repositories.
	Hooks_url string `json:"hooks_url"`
	Anonymous_access_enabled bool `json:"anonymous_access_enabled,omitempty"` // Whether anonymous git access is enabled for this repository
	Description string `json:"description"`
	Allow_rebase_merge bool `json:"allow_rebase_merge,omitempty"` // Whether to allow rebase merges for pull requests.
	Comments_url string `json:"comments_url"`
	Merges_url string `json:"merges_url"`
	Squash_merge_commit_title string `json:"squash_merge_commit_title,omitempty"` // The default value for a squash merge commit title: - `PR_TITLE` - default to the pull request's title. - `COMMIT_OR_PR_TITLE` - default to the commit's title (if only one commit) or the pull request's title (when more than one commit).
	Watchers int `json:"watchers"`
	Open_issues_count int `json:"open_issues_count"`
	Issue_events_url string `json:"issue_events_url"`
	Events_url string `json:"events_url"`
	Collaborators_url string `json:"collaborators_url"`
	Has_pages bool `json:"has_pages"`
	Clone_url string `json:"clone_url"`
	Created_at string `json:"created_at"`
	Pulls_url string `json:"pulls_url"`
	Default_branch string `json:"default_branch"` // The default branch of the repository.
	Open_issues int `json:"open_issues"`
	Full_name string `json:"full_name"`
	Subscribers_url string `json:"subscribers_url"`
	Git_url string `json:"git_url"`
	Languages_url string `json:"languages_url"`
	Stargazers_url string `json:"stargazers_url"`
	Url string `json:"url"`
	Has_projects bool `json:"has_projects"` // Whether projects are enabled.
	Deployments_url string `json:"deployments_url"`
	Downloads_url string `json:"downloads_url"`
	Stargazers_count int `json:"stargazers_count"`
	Starred_at string `json:"starred_at,omitempty"`
	Allow_auto_merge bool `json:"allow_auto_merge,omitempty"` // Whether to allow Auto-merge to be used on pull requests.
	Allow_update_branch bool `json:"allow_update_branch,omitempty"` // Whether or not a pull request head branch that is behind its base branch can always be updated even if it is not required to be up to date before merging.
	Trees_url string `json:"trees_url"`
	Git_commits_url string `json:"git_commits_url"`
	Branches_url string `json:"branches_url"`
	Private bool `json:"private"` // Whether the repository is private or public.
	Blobs_url string `json:"blobs_url"`
	Teams_url string `json:"teams_url"`
	Forks_count int `json:"forks_count"`
	Statuses_url string `json:"statuses_url"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Tags_url string `json:"tags_url"`
	Compare_url string `json:"compare_url"`
	Has_wiki bool `json:"has_wiki"` // Whether the wiki is enabled.
	Temp_clone_token string `json:"temp_clone_token,omitempty"`
	Id int64 `json:"id"` // Unique identifier of the repository
	Use_squash_pr_title_as_default bool `json:"use_squash_pr_title_as_default,omitempty"` // Whether a squash merge commit can use the pull request title as default. **This property is closing down. Please use `squash_merge_commit_title` instead.
	Archive_url string `json:"archive_url"`
	Issues_url string `json:"issues_url"`
	Allow_forking bool `json:"allow_forking,omitempty"` // Whether to allow forking this repo
	Html_url string `json:"html_url"`
	Disabled bool `json:"disabled"` // Returns whether or not this repository disabled.
	Squash_merge_commit_message string `json:"squash_merge_commit_message,omitempty"` // The default value for a squash merge commit message: - `PR_BODY` - default to the pull request's body. - `COMMIT_MESSAGES` - default to the branch's commit messages. - `BLANK` - default to a blank commit message.
	Size int `json:"size"` // The size of the repository, in kilobytes. Size is calculated hourly. When a repository is initially created, the size is 0.
	License GeneratedType_Nullable_license_simple `json:"license"` // License Simple
	Name string `json:"name"` // The name of the repository.
	Git_refs_url string `json:"git_refs_url"`
	Releases_url string `json:"releases_url"`
	Labels_url string `json:"labels_url"`
	Fork bool `json:"fork"`
	Has_issues bool `json:"has_issues"` // Whether issues are enabled.
	Watchers_count int `json:"watchers_count"`
	Git_tags_url string `json:"git_tags_url"`
	Assignees_url string `json:"assignees_url"`
	Forks_url string `json:"forks_url"`
	Has_downloads bool `json:"has_downloads"` // Whether downloads are enabled.
	Ssh_url string `json:"ssh_url"`
	Merge_commit_title string `json:"merge_commit_title,omitempty"` // The default value for a merge commit title. - `PR_TITLE` - default to the pull request's title. - `MERGE_MESSAGE` - default to the classic title for a merge message (e.g., Merge pull request #123 from branch-name).
	Homepage string `json:"homepage"`
	Allow_merge_commit bool `json:"allow_merge_commit,omitempty"` // Whether to allow merge commits for pull requests.
	Contents_url string `json:"contents_url"`
	Subscription_url string `json:"subscription_url"`
	Issue_comment_url string `json:"issue_comment_url"`
	Updated_at string `json:"updated_at"`
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"` // Whether to require contributors to sign off on web-based commits
	Notifications_url string `json:"notifications_url"`
	Pushed_at string `json:"pushed_at"`
	Mirror_url string `json:"mirror_url"`
	Node_id string `json:"node_id"`
}

// GeneratedType_Webhook_installation_target_renamed represents the GeneratedType_Webhook_installation_target_renamed schema from the OpenAPI specification
type GeneratedType_Webhook_installation_target_renamed struct {
	Account map[string]interface{} `json:"account"`
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Changes map[string]interface{} `json:"changes"`
	Installation GeneratedType_Simple_installation `json:"installation"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Target_type string `json:"target_type"`
	Action string `json:"action"`
}

// GeneratedType_Team_membership represents the GeneratedType_Team_membership schema from the OpenAPI specification
type GeneratedType_Team_membership struct {
	State string `json:"state"` // The state of the user's membership in the team.
	Url string `json:"url"`
	Role string `json:"role"` // The role of the user in the team.
}

// GeneratedType_Code_scanning_autofix represents the GeneratedType_Code_scanning_autofix schema from the OpenAPI specification
type GeneratedType_Code_scanning_autofix struct {
	Description string `json:"description"` // The description of an autofix.
	Started_at string `json:"started_at"` // The start time of an autofix in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Status string `json:"status"` // The status of an autofix.
}

// GeneratedType_Simple_classroom_repository represents the GeneratedType_Simple_classroom_repository schema from the OpenAPI specification
type GeneratedType_Simple_classroom_repository struct {
	Id int `json:"id"` // A unique identifier of the repository.
	Node_id string `json:"node_id"` // The GraphQL identifier of the repository.
	Private bool `json:"private"` // Whether the repository is private.
	Default_branch string `json:"default_branch"` // The default branch for the repository.
	Full_name string `json:"full_name"` // The full, globally unique name of the repository.
	Html_url string `json:"html_url"` // The URL to view the repository on GitHub.com.
}

// GeneratedType_Rate_limit represents the GeneratedType_Rate_limit schema from the OpenAPI specification
type GeneratedType_Rate_limit struct {
	Limit int `json:"limit"`
	Remaining int `json:"remaining"`
	Reset int `json:"reset"`
	Used int `json:"used"`
}

// GeneratedType_Webhook_dependabot_alert_fixed represents the GeneratedType_Webhook_dependabot_alert_fixed schema from the OpenAPI specification
type GeneratedType_Webhook_dependabot_alert_fixed struct {
	Alert GeneratedType_Dependabot_alert `json:"alert"` // A Dependabot alert.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType_Key_simple represents the GeneratedType_Key_simple schema from the OpenAPI specification
type GeneratedType_Key_simple struct {
	Id int `json:"id"`
	Key string `json:"key"`
	Created_at string `json:"created_at,omitempty"`
}

// GeneratedType_Webhook_release_deleted represents the GeneratedType_Webhook_release_deleted schema from the OpenAPI specification
type GeneratedType_Webhook_release_deleted struct {
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Release Webhooksrelease `json:"release"` // The [release](https://docs.github.com/rest/releases/releases/#get-a-release) object.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType_Webhook_pull_request_dequeued represents the GeneratedType_Webhook_pull_request_dequeued schema from the OpenAPI specification
type GeneratedType_Webhook_pull_request_dequeued struct {
	Number int `json:"number"`
	Action string `json:"action"`
	Pull_request map[string]interface{} `json:"pull_request"`
	Reason string `json:"reason"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// Webhooksissue represents the Webhooksissue schema from the OpenAPI specification
type Webhooksissue struct {
	Created_at string `json:"created_at"`
	Performed_via_github_app map[string]interface{} `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Html_url string `json:"html_url"`
	Locked bool `json:"locked,omitempty"`
	Events_url string `json:"events_url"`
	Labels_url string `json:"labels_url"`
	Number int `json:"number"`
	TypeField GeneratedType_Issue_type `json:"type,omitempty"` // The type of issue.
	User map[string]interface{} `json:"user"`
	Title string `json:"title"` // Title of the issue
	Url string `json:"url"` // URL for the issue
	Comments int `json:"comments"`
	Labels []map[string]interface{} `json:"labels,omitempty"`
	State string `json:"state,omitempty"` // State of the issue; either 'open' or 'closed'
	Updated_at string `json:"updated_at"`
	Active_lock_reason string `json:"active_lock_reason"`
	Repository_url string `json:"repository_url"`
	Draft bool `json:"draft,omitempty"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Reactions map[string]interface{} `json:"reactions"`
	Milestone map[string]interface{} `json:"milestone"` // A collection of related issues and pull requests.
	State_reason string `json:"state_reason,omitempty"`
	Node_id string `json:"node_id"`
	Assignee map[string]interface{} `json:"assignee,omitempty"`
	Timeline_url string `json:"timeline_url,omitempty"`
	Pull_request map[string]interface{} `json:"pull_request,omitempty"`
	Closed_at string `json:"closed_at"`
	Body string `json:"body"` // Contents of the issue
	Comments_url string `json:"comments_url"`
	Id int64 `json:"id"`
	Sub_issues_summary map[string]interface{} `json:"sub_issues_summary,omitempty"`
	Assignees []map[string]interface{} `json:"assignees"`
}

// GeneratedType_Protected_branch_admin_enforced represents the GeneratedType_Protected_branch_admin_enforced schema from the OpenAPI specification
type GeneratedType_Protected_branch_admin_enforced struct {
	Enabled bool `json:"enabled"`
	Url string `json:"url"`
}

// GeneratedType_Webhook_issues_assigned represents the GeneratedType_Webhook_issues_assigned schema from the OpenAPI specification
type GeneratedType_Webhook_issues_assigned struct {
	Action string `json:"action"` // The action that was performed.
	Assignee Webhooksuser `json:"assignee,omitempty"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue Webhooksissue `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
}

// GeneratedType_Webhook_release_published represents the GeneratedType_Webhook_release_published schema from the OpenAPI specification
type GeneratedType_Webhook_release_published struct {
	Release Webhooksrelease1 `json:"release"` // The [release](https://docs.github.com/rest/releases/releases/#get-a-release) object.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType_Projects_v2_status_update represents the GeneratedType_Projects_v2_status_update schema from the OpenAPI specification
type GeneratedType_Projects_v2_status_update struct {
	Id float64 `json:"id"`
	Project_node_id string `json:"project_node_id,omitempty"`
	Status string `json:"status,omitempty"`
	Updated_at string `json:"updated_at"`
	Created_at string `json:"created_at"`
	Start_date string `json:"start_date,omitempty"`
	Target_date string `json:"target_date,omitempty"`
	Body string `json:"body,omitempty"` // Body of the status update
	Creator GeneratedType_Simple_user `json:"creator,omitempty"` // A GitHub user.
	Node_id string `json:"node_id"`
}

// GeneratedType_Global_advisory represents the GeneratedType_Global_advisory schema from the OpenAPI specification
type GeneratedType_Global_advisory struct {
	Cvss map[string]interface{} `json:"cvss"`
	Cvss_severities GeneratedType_Cvss_severities `json:"cvss_severities,omitempty"`
	Ghsa_id string `json:"ghsa_id"` // The GitHub Security Advisory ID.
	Url string `json:"url"` // The API URL for the advisory.
	Epss GeneratedType_Security_advisory_epss `json:"epss,omitempty"` // The EPSS scores as calculated by the [Exploit Prediction Scoring System](https://www.first.org/epss).
	Source_code_location string `json:"source_code_location"` // The URL of the advisory's source code.
	Summary string `json:"summary"` // A short summary of the advisory.
	Vulnerabilities []Vulnerability `json:"vulnerabilities"` // The products and respective version ranges affected by the advisory.
	TypeField string `json:"type"` // The type of advisory.
	Html_url string `json:"html_url"` // The URL for the advisory.
	Repository_advisory_url string `json:"repository_advisory_url"` // The API URL for the repository advisory.
	Withdrawn_at string `json:"withdrawn_at"` // The date and time of when the advisory was withdrawn, in ISO 8601 format.
	Github_reviewed_at string `json:"github_reviewed_at"` // The date and time of when the advisory was reviewed by GitHub, in ISO 8601 format.
	Description string `json:"description"` // A detailed description of what the advisory entails.
	Identifiers []map[string]interface{} `json:"identifiers"`
	References []string `json:"references"`
	Cwes []map[string]interface{} `json:"cwes"`
	Credits []map[string]interface{} `json:"credits"` // The users who contributed to the advisory.
	Nvd_published_at string `json:"nvd_published_at"` // The date and time when the advisory was published in the National Vulnerability Database, in ISO 8601 format. This field is only populated when the advisory is imported from the National Vulnerability Database.
	Severity string `json:"severity"` // The severity of the advisory.
	Updated_at string `json:"updated_at"` // The date and time of when the advisory was last updated, in ISO 8601 format.
	Cve_id string `json:"cve_id"` // The Common Vulnerabilities and Exposures (CVE) ID.
	Published_at string `json:"published_at"` // The date and time of when the advisory was published, in ISO 8601 format.
}

// Webhooksanswer represents the Webhooksanswer schema from the OpenAPI specification
type Webhooksanswer struct {
	Node_id string `json:"node_id"`
	Child_comment_count int `json:"child_comment_count"`
	Discussion_id int `json:"discussion_id"`
	Parent_id interface{} `json:"parent_id"`
	Repository_url string `json:"repository_url"`
	Body string `json:"body"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
	User map[string]interface{} `json:"user"`
	Id int `json:"id"`
	Reactions map[string]interface{} `json:"reactions,omitempty"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Html_url string `json:"html_url"`
}

// GeneratedType_Page_deployment represents the GeneratedType_Page_deployment schema from the OpenAPI specification
type GeneratedType_Page_deployment struct {
	Preview_url string `json:"preview_url,omitempty"` // The URI to the deployed GitHub Pages preview.
	Status_url string `json:"status_url"` // The URI to monitor GitHub Pages deployment status.
	Id string `json:"id"` // The ID of the GitHub Pages deployment. This is the Git SHA of the deployed commit.
	Page_url string `json:"page_url"` // The URI to the deployed GitHub Pages.
}

// GeneratedType_Git_commit represents the GeneratedType_Git_commit schema from the OpenAPI specification
type GeneratedType_Git_commit struct {
	Parents []map[string]interface{} `json:"parents"`
	Sha string `json:"sha"` // SHA for the commit
	Author map[string]interface{} `json:"author"` // Identifying information for the git-user
	Html_url string `json:"html_url"`
	Committer map[string]interface{} `json:"committer"` // Identifying information for the git-user
	Message string `json:"message"` // Message describing the purpose of the commit
	Node_id string `json:"node_id"`
	Tree map[string]interface{} `json:"tree"`
	Url string `json:"url"`
	Verification map[string]interface{} `json:"verification"`
}

// GeneratedType_Unassigned_issue_event represents the GeneratedType_Unassigned_issue_event schema from the OpenAPI specification
type GeneratedType_Unassigned_issue_event struct {
	Commit_url string `json:"commit_url"`
	Assignee GeneratedType_Simple_user `json:"assignee"` // A GitHub user.
	Event string `json:"event"`
	Performed_via_github_app GeneratedType_Nullable_integration `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Actor GeneratedType_Simple_user `json:"actor"` // A GitHub user.
	Assigner GeneratedType_Simple_user `json:"assigner"` // A GitHub user.
	Commit_id string `json:"commit_id"`
	Created_at string `json:"created_at"`
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	Url string `json:"url"`
}

// GeneratedType_Code_scanning_autofix_commits represents the GeneratedType_Code_scanning_autofix_commits schema from the OpenAPI specification
type GeneratedType_Code_scanning_autofix_commits struct {
	Message string `json:"message,omitempty"` // Commit message to be used.
	Target_ref string `json:"target_ref,omitempty"` // The Git reference of target branch for the commit. Branch needs to already exist. For more information, see "[Git References](https://git-scm.com/book/en/v2/Git-Internals-Git-References)" in the Git documentation.
}

// GeneratedType_Webhook_config represents the GeneratedType_Webhook_config schema from the OpenAPI specification
type GeneratedType_Webhook_config struct {
	Secret string `json:"secret,omitempty"` // If provided, the `secret` will be used as the `key` to generate the HMAC hex digest value for [delivery signature headers](https://docs.github.com/webhooks/event-payloads/#delivery-headers).
	Url string `json:"url,omitempty"` // The URL to which the payloads will be delivered.
	Content_type string `json:"content_type,omitempty"` // The media type used to serialize the payloads. Supported values include `json` and `form`. The default is `form`.
	Insecure_ssl string `json:"insecure_ssl,omitempty"`
}

// GeneratedType_Issue_comment represents the GeneratedType_Issue_comment schema from the OpenAPI specification
type GeneratedType_Issue_comment struct {
	Body_text string `json:"body_text,omitempty"`
	Performed_via_github_app GeneratedType_Nullable_integration `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Url string `json:"url"` // URL for the issue comment
	Body string `json:"body,omitempty"` // Contents of the issue comment
	Body_html string `json:"body_html,omitempty"`
	Node_id string `json:"node_id"`
	Updated_at string `json:"updated_at"`
	Created_at string `json:"created_at"`
	Id int64 `json:"id"` // Unique identifier of the issue comment
	Reactions GeneratedType_Reaction_rollup `json:"reactions,omitempty"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Html_url string `json:"html_url"`
	Issue_url string `json:"issue_url"`
	User GeneratedType_Nullable_simple_user `json:"user"` // A GitHub user.
}

// GeneratedType_Secret_scanning_location_issue_comment represents the GeneratedType_Secret_scanning_location_issue_comment schema from the OpenAPI specification
type GeneratedType_Secret_scanning_location_issue_comment struct {
	Issue_comment_url string `json:"issue_comment_url"` // The API URL to get the issue comment where the secret was detected.
}

// Key represents the Key schema from the OpenAPI specification
type Key struct {
	Id int64 `json:"id"`
	Key string `json:"key"`
	Read_only bool `json:"read_only"`
	Title string `json:"title"`
	Url string `json:"url"`
	Verified bool `json:"verified"`
	Created_at string `json:"created_at"`
}

// Webhooksuser represents the Webhooksuser schema from the OpenAPI specification
type Webhooksuser struct {
	Site_admin bool `json:"site_admin,omitempty"`
	Gists_url string `json:"gists_url,omitempty"`
	Name string `json:"name,omitempty"`
	Node_id string `json:"node_id,omitempty"`
	Repos_url string `json:"repos_url,omitempty"`
	Followers_url string `json:"followers_url,omitempty"`
	Avatar_url string `json:"avatar_url,omitempty"`
	Html_url string `json:"html_url,omitempty"`
	Starred_url string `json:"starred_url,omitempty"`
	Url string `json:"url,omitempty"`
	Following_url string `json:"following_url,omitempty"`
	Id int64 `json:"id"`
	Organizations_url string `json:"organizations_url,omitempty"`
	Gravatar_id string `json:"gravatar_id,omitempty"`
	Subscriptions_url string `json:"subscriptions_url,omitempty"`
	Email string `json:"email,omitempty"`
	Events_url string `json:"events_url,omitempty"`
	TypeField string `json:"type,omitempty"`
	Deleted bool `json:"deleted,omitempty"`
	Login string `json:"login"`
	Received_events_url string `json:"received_events_url,omitempty"`
	User_view_type string `json:"user_view_type,omitempty"`
}

// GeneratedType_Basic_error represents the GeneratedType_Basic_error schema from the OpenAPI specification
type GeneratedType_Basic_error struct {
	Message string `json:"message,omitempty"`
	Status string `json:"status,omitempty"`
	Url string `json:"url,omitempty"`
	Documentation_url string `json:"documentation_url,omitempty"`
}

// GeneratedType_Label_search_result_item represents the GeneratedType_Label_search_result_item schema from the OpenAPI specification
type GeneratedType_Label_search_result_item struct {
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	Score float64 `json:"score"`
	Url string `json:"url"`
	Description string `json:"description"`
	DefaultField bool `json:"default"`
	Color string `json:"color"`
	Name string `json:"name"`
	Text_matches []map[string]interface{} `json:"text_matches,omitempty"`
}

// GeneratedType_Organization_dependabot_secret represents the GeneratedType_Organization_dependabot_secret schema from the OpenAPI specification
type GeneratedType_Organization_dependabot_secret struct {
	Created_at string `json:"created_at"`
	Name string `json:"name"` // The name of the secret.
	Selected_repositories_url string `json:"selected_repositories_url,omitempty"`
	Updated_at string `json:"updated_at"`
	Visibility string `json:"visibility"` // Visibility of a secret
}

// GeneratedType_Network_configuration represents the GeneratedType_Network_configuration schema from the OpenAPI specification
type GeneratedType_Network_configuration struct {
	Name string `json:"name"` // The name of the network configuration.
	Network_settings_ids []string `json:"network_settings_ids,omitempty"` // The unique identifier of each network settings in the configuration.
	Compute_service string `json:"compute_service,omitempty"` // The hosted compute service the network configuration supports.
	Created_on string `json:"created_on"` // The time at which the network configuration was created, in ISO 8601 format.
	Id string `json:"id"` // The unique identifier of the network configuration.
}

// GeneratedType_Webhook_projects_v2_status_update_edited represents the GeneratedType_Webhook_projects_v2_status_update_edited schema from the OpenAPI specification
type GeneratedType_Webhook_projects_v2_status_update_edited struct {
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Projects_v2_status_update GeneratedType_Projects_v2_status_update `json:"projects_v2_status_update"` // An status update belonging to a project
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes,omitempty"`
}

// GeneratedType_Issue_event_dismissed_review represents the GeneratedType_Issue_event_dismissed_review schema from the OpenAPI specification
type GeneratedType_Issue_event_dismissed_review struct {
	Dismissal_commit_id string `json:"dismissal_commit_id,omitempty"`
	Dismissal_message string `json:"dismissal_message"`
	Review_id int `json:"review_id"`
	State string `json:"state"`
}

// GeneratedType_Webhook_issues_edited represents the GeneratedType_Webhook_issues_edited schema from the OpenAPI specification
type GeneratedType_Webhook_issues_edited struct {
	Changes map[string]interface{} `json:"changes"` // The changes to the issue.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Action string `json:"action"`
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Label Webhookslabel `json:"label,omitempty"`
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType_Pull_request_review represents the GeneratedType_Pull_request_review schema from the OpenAPI specification
type GeneratedType_Pull_request_review struct {
	Commit_id string `json:"commit_id"` // A commit SHA for the review. If the commit object was garbage collected or forcibly deleted, then it no longer exists in Git and this value will be `null`.
	Html_url string `json:"html_url"`
	State string `json:"state"`
	Links map[string]interface{} `json:"_links"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Body string `json:"body"` // The text of the review.
	Body_text string `json:"body_text,omitempty"`
	Id int64 `json:"id"` // Unique identifier of the review
	Body_html string `json:"body_html,omitempty"`
	Node_id string `json:"node_id"`
	Pull_request_url string `json:"pull_request_url"`
	Submitted_at string `json:"submitted_at,omitempty"`
	User GeneratedType_Nullable_simple_user `json:"user"` // A GitHub user.
}

// GeneratedType_Dependabot_secret represents the GeneratedType_Dependabot_secret schema from the OpenAPI specification
type GeneratedType_Dependabot_secret struct {
	Created_at string `json:"created_at"`
	Name string `json:"name"` // The name of the secret.
	Updated_at string `json:"updated_at"`
}

// Installation represents the Installation schema from the OpenAPI specification
type Installation struct {
	Created_at string `json:"created_at"`
	Single_file_name string `json:"single_file_name"`
	Events []string `json:"events"`
	Access_tokens_url string `json:"access_tokens_url"`
	Id int `json:"id"` // The ID of the installation.
	Suspended_at string `json:"suspended_at"`
	Repositories_url string `json:"repositories_url"`
	Contact_email string `json:"contact_email,omitempty"`
	App_slug string `json:"app_slug"`
	Repository_selection string `json:"repository_selection"` // Describe whether all repositories have been selected or there's a selection involved
	Account interface{} `json:"account"`
	Target_type string `json:"target_type"`
	Html_url string `json:"html_url"`
	Has_multiple_single_files bool `json:"has_multiple_single_files,omitempty"`
	Permissions GeneratedType_App_permissions `json:"permissions"` // The permissions granted to the user access token.
	Suspended_by GeneratedType_Nullable_simple_user `json:"suspended_by"` // A GitHub user.
	Single_file_paths []string `json:"single_file_paths,omitempty"`
	Target_id int `json:"target_id"` // The ID of the user or organization this token is being scoped to.
	Updated_at string `json:"updated_at"`
	App_id int `json:"app_id"`
}

// GeneratedType_Authentication_token represents the GeneratedType_Authentication_token schema from the OpenAPI specification
type GeneratedType_Authentication_token struct {
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Repositories []Repository `json:"repositories,omitempty"` // The repositories this token has access to
	Repository_selection string `json:"repository_selection,omitempty"` // Describe whether all repositories have been selected or there's a selection involved
	Single_file string `json:"single_file,omitempty"`
	Token string `json:"token"` // The token used for authentication
	Expires_at string `json:"expires_at"` // The time this token expires
}

// GeneratedType_Secret_scanning_location_discussion_body represents the GeneratedType_Secret_scanning_location_discussion_body schema from the OpenAPI specification
type GeneratedType_Secret_scanning_location_discussion_body struct {
	Discussion_body_url string `json:"discussion_body_url"` // The URL to the discussion where the secret was detected.
}

// GeneratedType_Webhook_dependabot_alert_created represents the GeneratedType_Webhook_dependabot_alert_created schema from the OpenAPI specification
type GeneratedType_Webhook_dependabot_alert_created struct {
	Action string `json:"action"`
	Alert GeneratedType_Dependabot_alert `json:"alert"` // A Dependabot alert.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
}

// GeneratedType_Merge_group represents the GeneratedType_Merge_group schema from the OpenAPI specification
type GeneratedType_Merge_group struct {
	Head_commit GeneratedType_Simple_commit `json:"head_commit"` // A commit.
	Head_ref string `json:"head_ref"` // The full ref of the merge group.
	Head_sha string `json:"head_sha"` // The SHA of the merge group.
	Base_ref string `json:"base_ref"` // The full ref of the branch the merge group will be merged into.
	Base_sha string `json:"base_sha"` // The SHA of the merge group's parent commit.
}

// GeneratedType_Repository_ruleset_conditions_repository_name_target represents the GeneratedType_Repository_ruleset_conditions_repository_name_target schema from the OpenAPI specification
type GeneratedType_Repository_ruleset_conditions_repository_name_target struct {
	Repository_name map[string]interface{} `json:"repository_name"`
}

// GeneratedType_Webhook_issues_unlocked represents the GeneratedType_Webhook_issues_unlocked schema from the OpenAPI specification
type GeneratedType_Webhook_issues_unlocked struct {
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType_Auto_merge represents the GeneratedType_Auto_merge schema from the OpenAPI specification
type GeneratedType_Auto_merge struct {
	Merge_method string `json:"merge_method"` // The merge method to use.
	Commit_message string `json:"commit_message"` // Commit message for the merge commit.
	Commit_title string `json:"commit_title"` // Title for the merge commit message.
	Enabled_by GeneratedType_Simple_user `json:"enabled_by"` // A GitHub user.
}

// GeneratedType_Code_security_configuration represents the GeneratedType_Code_security_configuration schema from the OpenAPI specification
type GeneratedType_Code_security_configuration struct {
	Code_scanning_default_setup string `json:"code_scanning_default_setup,omitempty"` // The enablement status of code scanning default setup
	Html_url string `json:"html_url,omitempty"` // The URL of the configuration
	Description string `json:"description,omitempty"` // A description of the code security configuration
	Updated_at string `json:"updated_at,omitempty"`
	Secret_scanning_delegated_bypass_options map[string]interface{} `json:"secret_scanning_delegated_bypass_options,omitempty"` // Feature options for secret scanning delegated bypass
	Created_at string `json:"created_at,omitempty"`
	Secret_scanning_non_provider_patterns string `json:"secret_scanning_non_provider_patterns,omitempty"` // The enablement status of secret scanning non-provider patterns
	Secret_scanning string `json:"secret_scanning,omitempty"` // The enablement status of secret scanning
	Secret_scanning_delegated_alert_dismissal string `json:"secret_scanning_delegated_alert_dismissal,omitempty"` // The enablement status of secret scanning delegated alert dismissal
	Target_type string `json:"target_type,omitempty"` // The type of the code security configuration.
	Url string `json:"url,omitempty"` // The URL of the configuration
	Code_scanning_default_setup_options map[string]interface{} `json:"code_scanning_default_setup_options,omitempty"` // Feature options for code scanning default setup
	Code_scanning_delegated_alert_dismissal string `json:"code_scanning_delegated_alert_dismissal,omitempty"` // The enablement status of code scanning delegated alert dismissal
	Code_scanning_options map[string]interface{} `json:"code_scanning_options,omitempty"` // Feature options for code scanning
	Dependabot_alerts string `json:"dependabot_alerts,omitempty"` // The enablement status of Dependabot alerts
	Secret_scanning_validity_checks string `json:"secret_scanning_validity_checks,omitempty"` // The enablement status of secret scanning validity checks
	Dependabot_security_updates string `json:"dependabot_security_updates,omitempty"` // The enablement status of Dependabot security updates
	Advanced_security string `json:"advanced_security,omitempty"` // The enablement status of GitHub Advanced Security
	Dependency_graph string `json:"dependency_graph,omitempty"` // The enablement status of Dependency Graph
	Dependency_graph_autosubmit_action_options map[string]interface{} `json:"dependency_graph_autosubmit_action_options,omitempty"` // Feature options for Automatic dependency submission
	Secret_scanning_delegated_bypass string `json:"secret_scanning_delegated_bypass,omitempty"` // The enablement status of secret scanning delegated bypass
	Secret_scanning_generic_secrets string `json:"secret_scanning_generic_secrets,omitempty"` // The enablement status of Copilot secret scanning
	Secret_scanning_push_protection string `json:"secret_scanning_push_protection,omitempty"` // The enablement status of secret scanning push protection
	Dependency_graph_autosubmit_action string `json:"dependency_graph_autosubmit_action,omitempty"` // The enablement status of Automatic dependency submission
	Name string `json:"name,omitempty"` // The name of the code security configuration. Must be unique within the organization.
	Enforcement string `json:"enforcement,omitempty"` // The enforcement status for a security configuration
	Id int `json:"id,omitempty"` // The ID of the code security configuration
	Private_vulnerability_reporting string `json:"private_vulnerability_reporting,omitempty"` // The enablement status of private vulnerability reporting
}

// GeneratedType_Thread_subscription represents the GeneratedType_Thread_subscription schema from the OpenAPI specification
type GeneratedType_Thread_subscription struct {
	Created_at string `json:"created_at"`
	Ignored bool `json:"ignored"`
	Reason string `json:"reason"`
	Repository_url string `json:"repository_url,omitempty"`
	Subscribed bool `json:"subscribed"`
	Thread_url string `json:"thread_url,omitempty"`
	Url string `json:"url"`
}

// GeneratedType_Classroom_assignment represents the GeneratedType_Classroom_assignment schema from the OpenAPI specification
type GeneratedType_Classroom_assignment struct {
	Feedback_pull_requests_enabled bool `json:"feedback_pull_requests_enabled"` // Whether feedback pull request will be created when a student accepts the assignment.
	Language string `json:"language"` // The programming language used in the assignment.
	Classroom Classroom `json:"classroom"` // A GitHub Classroom classroom
	Max_members int `json:"max_members"` // The maximum allowable members per team.
	Accepted int `json:"accepted"` // The number of students that have accepted the assignment.
	Slug string `json:"slug"` // Sluggified name of the assignment.
	Max_teams int `json:"max_teams"` // The maximum allowable teams for the assignment.
	Id int `json:"id"` // Unique identifier of the repository.
	Title string `json:"title"` // Assignment title.
	Deadline string `json:"deadline"` // The time at which the assignment is due.
	Editor string `json:"editor"` // The selected editor for the assignment.
	Invite_link string `json:"invite_link"` // The link that a student can use to accept the assignment.
	Students_are_repo_admins bool `json:"students_are_repo_admins"` // Whether students are admins on created repository when a student accepts the assignment.
	Invitations_enabled bool `json:"invitations_enabled"` // Whether the invitation link is enabled. Visiting an enabled invitation link will accept the assignment.
	Starter_code_repository GeneratedType_Simple_classroom_repository `json:"starter_code_repository"` // A GitHub repository view for Classroom
	Submitted int `json:"submitted"` // The number of students that have submitted the assignment.
	Passing int `json:"passing"` // The number of students that have passed the assignment.
	Public_repo bool `json:"public_repo"` // Whether an accepted assignment creates a public repository.
	TypeField string `json:"type"` // Whether it's a group assignment or individual assignment.
}

// GeneratedType_Webhook_check_run_requested_action represents the GeneratedType_Webhook_check_run_requested_action schema from the OpenAPI specification
type GeneratedType_Webhook_check_run_requested_action struct {
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Check_run GeneratedType_Check_run_with_simple_check_suite `json:"check_run"` // A check performed on the code of a given code change
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Requested_action map[string]interface{} `json:"requested_action,omitempty"` // The action requested by the user.
}

// GeneratedType_Webhook_pull_request_ready_for_review represents the GeneratedType_Webhook_pull_request_ready_for_review schema from the OpenAPI specification
type GeneratedType_Webhook_pull_request_ready_for_review struct {
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request GeneratedType_Pull_request_webhook `json:"pull_request"`
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Number int `json:"number"` // The pull request number.
}

// GeneratedType_Repository_rule_params_restricted_commits represents the GeneratedType_Repository_rule_params_restricted_commits schema from the OpenAPI specification
type GeneratedType_Repository_rule_params_restricted_commits struct {
	Reason string `json:"reason,omitempty"` // Reason for restriction
	Oid string `json:"oid"` // Full or abbreviated commit hash to reject
}

// License represents the License schema from the OpenAPI specification
type License struct {
	Featured bool `json:"featured"`
	Html_url string `json:"html_url"`
	Limitations []string `json:"limitations"`
	Description string `json:"description"`
	Implementation string `json:"implementation"`
	Key string `json:"key"`
	Node_id string `json:"node_id"`
	Spdx_id string `json:"spdx_id"`
	Url string `json:"url"`
	Body string `json:"body"`
	Permissions []string `json:"permissions"`
	Conditions []string `json:"conditions"`
	Name string `json:"name"`
}

// GeneratedType_Webhook_package_published represents the GeneratedType_Webhook_package_published schema from the OpenAPI specification
type GeneratedType_Webhook_package_published struct {
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	PackageField map[string]interface{} `json:"package"` // Information about the package.
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType_Code_scanning_default_setup represents the GeneratedType_Code_scanning_default_setup schema from the OpenAPI specification
type GeneratedType_Code_scanning_default_setup struct {
	Runner_type string `json:"runner_type,omitempty"` // Runner type to be used.
	Schedule string `json:"schedule,omitempty"` // The frequency of the periodic analysis.
	State string `json:"state,omitempty"` // Code scanning default setup has been configured or not.
	Threat_model string `json:"threat_model,omitempty"` // Threat model to be used for code scanning analysis. Use `remote` to analyze only network sources and `remote_and_local` to include local sources like filesystem access, command-line arguments, database reads, environment variable and standard input.
	Updated_at string `json:"updated_at,omitempty"` // Timestamp of latest configuration update.
	Languages []string `json:"languages,omitempty"` // Languages to be analyzed.
	Query_suite string `json:"query_suite,omitempty"` // CodeQL query suite to be used.
	Runner_label string `json:"runner_label,omitempty"` // Runner label to be used if the runner type is labeled.
}

// GeneratedType_Webhook_branch_protection_configuration_enabled represents the GeneratedType_Webhook_branch_protection_configuration_enabled schema from the OpenAPI specification
type GeneratedType_Webhook_branch_protection_configuration_enabled struct {
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType_Webhook_workflow_job_queued represents the GeneratedType_Webhook_workflow_job_queued schema from the OpenAPI specification
type GeneratedType_Webhook_workflow_job_queued struct {
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Workflow_job map[string]interface{} `json:"workflow_job"`
	Action string `json:"action"`
	Deployment Deployment `json:"deployment,omitempty"` // A request for a specific ref(branch,sha,tag) to be deployed
}

// GeneratedType_User_role_assignment represents the GeneratedType_User_role_assignment schema from the OpenAPI specification
type GeneratedType_User_role_assignment struct {
	Starred_url string `json:"starred_url"`
	TypeField string `json:"type"`
	Id int `json:"id"`
	Organizations_url string `json:"organizations_url"`
	Received_events_url string `json:"received_events_url"`
	Email string `json:"email,omitempty"`
	Following_url string `json:"following_url"`
	Url string `json:"url"`
	Site_admin bool `json:"site_admin"`
	Inherited_from []GeneratedType_Team_simple `json:"inherited_from,omitempty"` // Team the user has gotten the role through
	Assignment string `json:"assignment,omitempty"` // Determines if the user has a direct, indirect, or mixed relationship to a role
	User_view_type string `json:"user_view_type,omitempty"`
	Login string `json:"login"`
	Repos_url string `json:"repos_url"`
	Subscriptions_url string `json:"subscriptions_url"`
	Gravatar_id string `json:"gravatar_id"`
	Html_url string `json:"html_url"`
	Starred_at string `json:"starred_at,omitempty"`
	Name string `json:"name,omitempty"`
	Avatar_url string `json:"avatar_url"`
	Followers_url string `json:"followers_url"`
	Gists_url string `json:"gists_url"`
	Node_id string `json:"node_id"`
	Events_url string `json:"events_url"`
}

// GeneratedType_Project_card represents the GeneratedType_Project_card schema from the OpenAPI specification
type GeneratedType_Project_card struct {
	Url string `json:"url"`
	Column_name string `json:"column_name,omitempty"`
	Column_url string `json:"column_url"`
	Id int64 `json:"id"` // The project card's ID
	Node_id string `json:"node_id"`
	Project_id string `json:"project_id,omitempty"`
	Project_url string `json:"project_url"`
	Updated_at string `json:"updated_at"`
	Created_at string `json:"created_at"`
	Creator GeneratedType_Nullable_simple_user `json:"creator"` // A GitHub user.
	Note string `json:"note"`
	Content_url string `json:"content_url,omitempty"`
	Archived bool `json:"archived,omitempty"` // Whether or not the card is archived
}

// GeneratedType_Pull_request_review_request represents the GeneratedType_Pull_request_review_request schema from the OpenAPI specification
type GeneratedType_Pull_request_review_request struct {
	Teams []Team `json:"teams"`
	Users []GeneratedType_Simple_user `json:"users"`
}

// GeneratedType_Nullable_license_simple represents the GeneratedType_Nullable_license_simple schema from the OpenAPI specification
type GeneratedType_Nullable_license_simple struct {
	Url string `json:"url"`
	Html_url string `json:"html_url,omitempty"`
	Key string `json:"key"`
	Name string `json:"name"`
	Node_id string `json:"node_id"`
	Spdx_id string `json:"spdx_id"`
}

// GeneratedType_Project_column represents the GeneratedType_Project_column schema from the OpenAPI specification
type GeneratedType_Project_column struct {
	Name string `json:"name"` // Name of the project column
	Node_id string `json:"node_id"`
	Project_url string `json:"project_url"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	Cards_url string `json:"cards_url"`
	Created_at string `json:"created_at"`
	Id int `json:"id"` // The unique identifier of the project column
}

// GeneratedType_Unlabeled_issue_event represents the GeneratedType_Unlabeled_issue_event schema from the OpenAPI specification
type GeneratedType_Unlabeled_issue_event struct {
	Commit_url string `json:"commit_url"`
	Label map[string]interface{} `json:"label"`
	Node_id string `json:"node_id"`
	Performed_via_github_app GeneratedType_Nullable_integration `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Created_at string `json:"created_at"`
	Event string `json:"event"`
	Actor GeneratedType_Simple_user `json:"actor"` // A GitHub user.
	Commit_id string `json:"commit_id"`
	Id int `json:"id"`
	Url string `json:"url"`
}

// GeneratedType_Pages_deployment_status represents the GeneratedType_Pages_deployment_status schema from the OpenAPI specification
type GeneratedType_Pages_deployment_status struct {
	Status string `json:"status,omitempty"` // The current status of the deployment.
}

// GeneratedType_Nullable_simple_user represents the GeneratedType_Nullable_simple_user schema from the OpenAPI specification
type GeneratedType_Nullable_simple_user struct {
	Followers_url string `json:"followers_url"`
	Starred_at string `json:"starred_at,omitempty"`
	Node_id string `json:"node_id"`
	Following_url string `json:"following_url"`
	Organizations_url string `json:"organizations_url"`
	Avatar_url string `json:"avatar_url"`
	Email string `json:"email,omitempty"`
	User_view_type string `json:"user_view_type,omitempty"`
	Id int64 `json:"id"`
	Subscriptions_url string `json:"subscriptions_url"`
	Received_events_url string `json:"received_events_url"`
	Gists_url string `json:"gists_url"`
	Login string `json:"login"`
	Site_admin bool `json:"site_admin"`
	Html_url string `json:"html_url"`
	TypeField string `json:"type"`
	Name string `json:"name,omitempty"`
	Starred_url string `json:"starred_url"`
	Repos_url string `json:"repos_url"`
	Gravatar_id string `json:"gravatar_id"`
	Url string `json:"url"`
	Events_url string `json:"events_url"`
}

// GeneratedType_Pending_deployment represents the GeneratedType_Pending_deployment schema from the OpenAPI specification
type GeneratedType_Pending_deployment struct {
	Current_user_can_approve bool `json:"current_user_can_approve"` // Whether the currently authenticated user can approve the deployment
	Environment map[string]interface{} `json:"environment"`
	Reviewers []map[string]interface{} `json:"reviewers"` // The people or teams that may approve jobs that reference the environment. You can list up to six users or teams as reviewers. The reviewers must have at least read access to the repository. Only one of the required reviewers needs to approve the job for it to proceed.
	Wait_timer int `json:"wait_timer"` // The set duration of the wait timer
	Wait_timer_started_at string `json:"wait_timer_started_at"` // The time that the wait timer began.
}

// GeneratedType_Org_membership represents the GeneratedType_Org_membership schema from the OpenAPI specification
type GeneratedType_Org_membership struct {
	Organization GeneratedType_Organization_simple `json:"organization"` // A GitHub organization.
	Organization_url string `json:"organization_url"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Role string `json:"role"` // The user's membership type in the organization.
	State string `json:"state"` // The state of the member in the organization. The `pending` state indicates the user has not yet accepted an invitation.
	Url string `json:"url"`
	User GeneratedType_Nullable_simple_user `json:"user"` // A GitHub user.
}

// GeneratedType_Webhook_release_released represents the GeneratedType_Webhook_release_released schema from the OpenAPI specification
type GeneratedType_Webhook_release_released struct {
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Release Webhooksrelease `json:"release"` // The [release](https://docs.github.com/rest/releases/releases/#get-a-release) object.
}

// Label represents the Label schema from the OpenAPI specification
type Label struct {
	Node_id string `json:"node_id"`
	Url string `json:"url"` // URL for the label
	Color string `json:"color"` // 6-character hex code, without the leading #, identifying the color
	DefaultField bool `json:"default"` // Whether this label comes by default in a new repository.
	Description string `json:"description"` // Optional description of the label, such as its purpose.
	Id int64 `json:"id"` // Unique identifier for the label.
	Name string `json:"name"` // The name of the label.
}

// GeneratedType_Webhook_marketplace_purchase_cancelled represents the GeneratedType_Webhook_marketplace_purchase_cancelled schema from the OpenAPI specification
type GeneratedType_Webhook_marketplace_purchase_cancelled struct {
	Effective_date string `json:"effective_date"`
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Marketplace_purchase Webhooksmarketplacepurchase `json:"marketplace_purchase"`
	Previous_marketplace_purchase Webhookspreviousmarketplacepurchase `json:"previous_marketplace_purchase,omitempty"`
}

// GeneratedType_Custom_property represents the GeneratedType_Custom_property schema from the OpenAPI specification
type GeneratedType_Custom_property struct {
	Required bool `json:"required,omitempty"` // Whether the property is required.
	Url string `json:"url,omitempty"` // The URL that can be used to fetch, update, or delete info about this property via the API.
	Default_value string `json:"default_value,omitempty"` // Default value of the property
	Values_editable_by string `json:"values_editable_by,omitempty"` // Who can edit the values of the property
	Property_name string `json:"property_name"` // The name of the property
	Source_type string `json:"source_type,omitempty"` // The source type of the property
	Value_type string `json:"value_type"` // The type of the value for the property
	Description string `json:"description,omitempty"` // Short description of the property
	Allowed_values []string `json:"allowed_values,omitempty"` // An ordered list of the allowed values of the property. The property can have up to 200 allowed values.
}

// GeneratedType_Webhook_installation_created represents the GeneratedType_Webhook_installation_created schema from the OpenAPI specification
type GeneratedType_Webhook_installation_created struct {
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repositories []map[string]interface{} `json:"repositories,omitempty"` // An array of repository objects that the installation can access.
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Requester Webhooksuser `json:"requester,omitempty"`
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation Installation `json:"installation"` // Installation
}

// Webhookssponsorship represents the Webhookssponsorship schema from the OpenAPI specification
type Webhookssponsorship struct {
	Maintainer map[string]interface{} `json:"maintainer,omitempty"`
	Node_id string `json:"node_id"`
	Privacy_level string `json:"privacy_level"`
	Sponsor map[string]interface{} `json:"sponsor"`
	Sponsorable map[string]interface{} `json:"sponsorable"`
	Tier map[string]interface{} `json:"tier"` // The `tier_changed` and `pending_tier_change` will include the original tier before the change or pending change. For more information, see the pending tier change payload.
	Created_at string `json:"created_at"`
}

// GeneratedType_Repository_rule_tag_name_pattern represents the GeneratedType_Repository_rule_tag_name_pattern schema from the OpenAPI specification
type GeneratedType_Repository_rule_tag_name_pattern struct {
	TypeField string `json:"type"`
	Parameters map[string]interface{} `json:"parameters,omitempty"`
}

// GeneratedType_Review_custom_gates_comment_required represents the GeneratedType_Review_custom_gates_comment_required schema from the OpenAPI specification
type GeneratedType_Review_custom_gates_comment_required struct {
	Comment string `json:"comment"` // Comment associated with the pending deployment protection rule. **Required when state is not provided.**
	Environment_name string `json:"environment_name"` // The name of the environment to approve or reject.
}

// GeneratedType_Webhook_custom_property_promoted_to_enterprise represents the GeneratedType_Webhook_custom_property_promoted_to_enterprise schema from the OpenAPI specification
type GeneratedType_Webhook_custom_property_promoted_to_enterprise struct {
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Definition GeneratedType_Custom_property `json:"definition"` // Custom property defined on an organization
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType_Api_insights_summary_stats represents the GeneratedType_Api_insights_summary_stats schema from the OpenAPI specification
type GeneratedType_Api_insights_summary_stats struct {
	Rate_limited_request_count int64 `json:"rate_limited_request_count,omitempty"` // The total number of requests that were rate limited within the queried time period
	Total_request_count int64 `json:"total_request_count,omitempty"` // The total number of requests within the queried time period
}

// GeneratedType_Custom_property_set_payload represents the GeneratedType_Custom_property_set_payload schema from the OpenAPI specification
type GeneratedType_Custom_property_set_payload struct {
	Required bool `json:"required,omitempty"` // Whether the property is required.
	Value_type string `json:"value_type"` // The type of the value for the property
	Values_editable_by string `json:"values_editable_by,omitempty"` // Who can edit the values of the property
	Allowed_values []string `json:"allowed_values,omitempty"` // An ordered list of the allowed values of the property. The property can have up to 200 allowed values.
	Default_value string `json:"default_value,omitempty"` // Default value of the property
	Description string `json:"description,omitempty"` // Short description of the property
}

// GeneratedType_Webhook_page_build represents the GeneratedType_Webhook_page_build schema from the OpenAPI specification
type GeneratedType_Webhook_page_build struct {
	Id int `json:"id"`
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Build map[string]interface{} `json:"build"` // The [List GitHub Pages builds](https://docs.github.com/rest/pages/pages#list-github-pages-builds) itself.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType_Webhook_repository_unarchived represents the GeneratedType_Webhook_repository_unarchived schema from the OpenAPI specification
type GeneratedType_Webhook_repository_unarchived struct {
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType_File_commit represents the GeneratedType_File_commit schema from the OpenAPI specification
type GeneratedType_File_commit struct {
	Commit map[string]interface{} `json:"commit"`
	Content map[string]interface{} `json:"content"`
}

// GeneratedType_App_permissions represents the GeneratedType_App_permissions schema from the OpenAPI specification
type GeneratedType_App_permissions struct {
	Organization_projects string `json:"organization_projects,omitempty"` // The level of permission to grant the access token to manage organization projects and projects public preview (where available).
	Organization_user_blocking string `json:"organization_user_blocking,omitempty"` // The level of permission to grant the access token to view and manage users blocked by the organization.
	Organization_custom_org_roles string `json:"organization_custom_org_roles,omitempty"` // The level of permission to grant the access token for custom organization roles management.
	Interaction_limits string `json:"interaction_limits,omitempty"` // The level of permission to grant the access token to view and manage interaction limits on a repository.
	Organization_events string `json:"organization_events,omitempty"` // The level of permission to grant the access token to view events triggered by an activity in an organization.
	Administration string `json:"administration,omitempty"` // The level of permission to grant the access token for repository creation, deletion, settings, teams, and collaborators creation.
	Organization_copilot_seat_management string `json:"organization_copilot_seat_management,omitempty"` // The level of permission to grant the access token for managing access to GitHub Copilot for members of an organization with a Copilot Business subscription. This property is in public preview and is subject to change.
	Pull_requests string `json:"pull_requests,omitempty"` // The level of permission to grant the access token for pull requests and related comments, assignees, labels, milestones, and merges.
	Organization_packages string `json:"organization_packages,omitempty"` // The level of permission to grant the access token for organization packages published to GitHub Packages.
	Followers string `json:"followers,omitempty"` // The level of permission to grant the access token to manage the followers belonging to a user.
	Organization_personal_access_tokens string `json:"organization_personal_access_tokens,omitempty"` // The level of permission to grant the access token for viewing and managing fine-grained personal access token requests to an organization.
	Secret_scanning_alerts string `json:"secret_scanning_alerts,omitempty"` // The level of permission to grant the access token to view and manage secret scanning alerts.
	Repository_hooks string `json:"repository_hooks,omitempty"` // The level of permission to grant the access token to manage the post-receive hooks for a repository.
	Checks string `json:"checks,omitempty"` // The level of permission to grant the access token for checks on code.
	Organization_secrets string `json:"organization_secrets,omitempty"` // The level of permission to grant the access token to manage organization secrets.
	Organization_plan string `json:"organization_plan,omitempty"` // The level of permission to grant the access token for viewing an organization's plan.
	Members string `json:"members,omitempty"` // The level of permission to grant the access token for organization teams and members.
	Statuses string `json:"statuses,omitempty"` // The level of permission to grant the access token for commit statuses.
	Profile string `json:"profile,omitempty"` // The level of permission to grant the access token to manage the profile settings belonging to a user.
	Contents string `json:"contents,omitempty"` // The level of permission to grant the access token for repository contents, commits, branches, downloads, releases, and merges.
	Codespaces string `json:"codespaces,omitempty"` // The level of permission to grant the access token to create, edit, delete, and list Codespaces.
	Organization_custom_roles string `json:"organization_custom_roles,omitempty"` // The level of permission to grant the access token for custom repository roles management.
	Organization_personal_access_token_requests string `json:"organization_personal_access_token_requests,omitempty"` // The level of permission to grant the access token for viewing and managing fine-grained personal access tokens that have been approved by an organization.
	Packages string `json:"packages,omitempty"` // The level of permission to grant the access token for packages published to GitHub Packages.
	Organization_announcement_banners string `json:"organization_announcement_banners,omitempty"` // The level of permission to grant the access token to view and manage announcement banners for an organization.
	Pages string `json:"pages,omitempty"` // The level of permission to grant the access token to retrieve Pages statuses, configuration, and builds, as well as create new builds.
	Organization_self_hosted_runners string `json:"organization_self_hosted_runners,omitempty"` // The level of permission to grant the access token to view and manage GitHub Actions self-hosted runners available to an organization.
	Email_addresses string `json:"email_addresses,omitempty"` // The level of permission to grant the access token to manage the email addresses belonging to a user.
	Metadata string `json:"metadata,omitempty"` // The level of permission to grant the access token to search repositories, list collaborators, and access repository metadata.
	Starring string `json:"starring,omitempty"` // The level of permission to grant the access token to list and manage repositories a user is starring.
	Actions string `json:"actions,omitempty"` // The level of permission to grant the access token for GitHub Actions workflows, workflow runs, and artifacts.
	Team_discussions string `json:"team_discussions,omitempty"` // The level of permission to grant the access token to manage team discussions and related comments.
	Repository_custom_properties string `json:"repository_custom_properties,omitempty"` // The level of permission to grant the access token to view and edit custom properties for a repository, when allowed by the property.
	Gpg_keys string `json:"gpg_keys,omitempty"` // The level of permission to grant the access token to view and manage GPG keys belonging to a user.
	Repository_projects string `json:"repository_projects,omitempty"` // The level of permission to grant the access token to manage repository projects, columns, and cards.
	Environments string `json:"environments,omitempty"` // The level of permission to grant the access token for managing repository environments.
	Organization_custom_properties string `json:"organization_custom_properties,omitempty"` // The level of permission to grant the access token for custom property management.
	Single_file string `json:"single_file,omitempty"` // The level of permission to grant the access token to manage just a single file.
	Security_events string `json:"security_events,omitempty"` // The level of permission to grant the access token to view and manage security events like code scanning alerts.
	Workflows string `json:"workflows,omitempty"` // The level of permission to grant the access token to update GitHub Actions workflow files.
	Dependabot_secrets string `json:"dependabot_secrets,omitempty"` // The level of permission to grant the access token to manage Dependabot secrets.
	Issues string `json:"issues,omitempty"` // The level of permission to grant the access token for issues and related comments, assignees, labels, and milestones.
	Secrets string `json:"secrets,omitempty"` // The level of permission to grant the access token to manage repository secrets.
	Git_ssh_keys string `json:"git_ssh_keys,omitempty"` // The level of permission to grant the access token to manage git SSH keys.
	Organization_administration string `json:"organization_administration,omitempty"` // The level of permission to grant the access token to manage access to an organization.
	Organization_hooks string `json:"organization_hooks,omitempty"` // The level of permission to grant the access token to manage the post-receive hooks for an organization.
	Vulnerability_alerts string `json:"vulnerability_alerts,omitempty"` // The level of permission to grant the access token to manage Dependabot alerts.
	Deployments string `json:"deployments,omitempty"` // The level of permission to grant the access token for deployments and deployment statuses.
}

// GeneratedType_Issue_event_label represents the GeneratedType_Issue_event_label schema from the OpenAPI specification
type GeneratedType_Issue_event_label struct {
	Name string `json:"name"`
	Color string `json:"color"`
}

// GeneratedType_Webhook_milestone_edited represents the GeneratedType_Webhook_milestone_edited schema from the OpenAPI specification
type GeneratedType_Webhook_milestone_edited struct {
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"` // The changes to the milestone if the action was `edited`.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Milestone Webhooksmilestone `json:"milestone"` // A collection of related issues and pull requests.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType_Webhook_create represents the GeneratedType_Webhook_create schema from the OpenAPI specification
type GeneratedType_Webhook_create struct {
	Ref string `json:"ref"` // The [`git ref`](https://docs.github.com/rest/git/refs#get-a-reference) resource.
	Ref_type string `json:"ref_type"` // The type of Git ref object created in the repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Description string `json:"description"` // The repository's current description.
	Master_branch string `json:"master_branch"` // The name of the repository's default branch (usually `main`).
	Pusher_type string `json:"pusher_type"` // The pusher type for the event. Can be either `user` or a deploy key.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType_Copilot_dotcom_pull_requests represents the GeneratedType_Copilot_dotcom_pull_requests schema from the OpenAPI specification
type GeneratedType_Copilot_dotcom_pull_requests struct {
	Repositories []map[string]interface{} `json:"repositories,omitempty"` // Repositories in which users used Copilot for Pull Requests to generate pull request summaries
	Total_engaged_users int `json:"total_engaged_users,omitempty"` // The number of users who used Copilot for Pull Requests on github.com to generate a pull request summary at least once.
}

// GeneratedType_Network_settings represents the GeneratedType_Network_settings schema from the OpenAPI specification
type GeneratedType_Network_settings struct {
	Region string `json:"region"` // The location of the subnet this network settings resource is configured for.
	Subnet_id string `json:"subnet_id"` // The subnet this network settings resource is configured for.
	Id string `json:"id"` // The unique identifier of the network settings resource.
	Name string `json:"name"` // The name of the network settings resource.
	Network_configuration_id string `json:"network_configuration_id,omitempty"` // The identifier of the network configuration that is using this settings resource.
}

// GeneratedType_Webhook_gollum represents the GeneratedType_Webhook_gollum schema from the OpenAPI specification
type GeneratedType_Webhook_gollum struct {
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pages []map[string]interface{} `json:"pages"` // The pages that were updated.
}

// GeneratedType_Minimal_repository represents the GeneratedType_Minimal_repository schema from the OpenAPI specification
type GeneratedType_Minimal_repository struct {
	Open_issues int `json:"open_issues,omitempty"`
	Open_issues_count int `json:"open_issues_count,omitempty"`
	Is_template bool `json:"is_template,omitempty"`
	Id int64 `json:"id"`
	Url string `json:"url"`
	Contents_url string `json:"contents_url"`
	Trees_url string `json:"trees_url"`
	Default_branch string `json:"default_branch,omitempty"`
	Pushed_at string `json:"pushed_at,omitempty"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Allow_forking bool `json:"allow_forking,omitempty"`
	Forks_url string `json:"forks_url"`
	Watchers int `json:"watchers,omitempty"`
	Code_of_conduct GeneratedType_Code_of_conduct `json:"code_of_conduct,omitempty"` // Code Of Conduct
	Pulls_url string `json:"pulls_url"`
	Compare_url string `json:"compare_url"`
	Forks_count int `json:"forks_count,omitempty"`
	Has_projects bool `json:"has_projects,omitempty"`
	Mirror_url string `json:"mirror_url,omitempty"`
	Ssh_url string `json:"ssh_url,omitempty"`
	Node_id string `json:"node_id"`
	Subscription_url string `json:"subscription_url"`
	Statuses_url string `json:"statuses_url"`
	Keys_url string `json:"keys_url"`
	Subscribers_count int `json:"subscribers_count,omitempty"`
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"`
	Git_url string `json:"git_url,omitempty"`
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"`
	Name string `json:"name"`
	Contributors_url string `json:"contributors_url"`
	Has_wiki bool `json:"has_wiki,omitempty"`
	Fork bool `json:"fork"`
	Assignees_url string `json:"assignees_url"`
	Network_count int `json:"network_count,omitempty"`
	Issue_comment_url string `json:"issue_comment_url"`
	Stargazers_count int `json:"stargazers_count,omitempty"`
	Topics []string `json:"topics,omitempty"`
	Custom_properties map[string]interface{} `json:"custom_properties,omitempty"` // The custom properties that were defined for the repository. The keys are the custom property names, and the values are the corresponding custom property values.
	Teams_url string `json:"teams_url"`
	Git_commits_url string `json:"git_commits_url"`
	Has_discussions bool `json:"has_discussions,omitempty"`
	Labels_url string `json:"labels_url"`
	Visibility string `json:"visibility,omitempty"`
	Stargazers_url string `json:"stargazers_url"`
	Svn_url string `json:"svn_url,omitempty"`
	Created_at string `json:"created_at,omitempty"`
	Collaborators_url string `json:"collaborators_url"`
	Security_and_analysis GeneratedType_Security_and_analysis `json:"security_and_analysis,omitempty"`
	Downloads_url string `json:"downloads_url"`
	Notifications_url string `json:"notifications_url"`
	Homepage string `json:"homepage,omitempty"`
	Language string `json:"language,omitempty"`
	Tags_url string `json:"tags_url"`
	Forks int `json:"forks,omitempty"`
	Events_url string `json:"events_url"`
	Merges_url string `json:"merges_url"`
	Disabled bool `json:"disabled,omitempty"`
	Subscribers_url string `json:"subscribers_url"`
	Html_url string `json:"html_url"`
	Commits_url string `json:"commits_url"`
	Clone_url string `json:"clone_url,omitempty"`
	Description string `json:"description"`
	Private bool `json:"private"`
	Branches_url string `json:"branches_url"`
	Issue_events_url string `json:"issue_events_url"`
	Archive_url string `json:"archive_url"`
	Languages_url string `json:"languages_url"`
	Git_refs_url string `json:"git_refs_url"`
	Owner GeneratedType_Simple_user `json:"owner"` // A GitHub user.
	Size int `json:"size,omitempty"` // The size of the repository, in kilobytes. Size is calculated hourly. When a repository is initially created, the size is 0.
	Issues_url string `json:"issues_url"`
	Full_name string `json:"full_name"`
	License map[string]interface{} `json:"license,omitempty"`
	Hooks_url string `json:"hooks_url"`
	Temp_clone_token string `json:"temp_clone_token,omitempty"`
	Has_pages bool `json:"has_pages,omitempty"`
	Blobs_url string `json:"blobs_url"`
	Archived bool `json:"archived,omitempty"`
	Comments_url string `json:"comments_url"`
	Role_name string `json:"role_name,omitempty"`
	Releases_url string `json:"releases_url"`
	Has_downloads bool `json:"has_downloads,omitempty"`
	Milestones_url string `json:"milestones_url"`
	Git_tags_url string `json:"git_tags_url"`
	Watchers_count int `json:"watchers_count,omitempty"`
	Has_issues bool `json:"has_issues,omitempty"`
	Deployments_url string `json:"deployments_url"`
	Updated_at string `json:"updated_at,omitempty"`
}

// Reaction represents the Reaction schema from the OpenAPI specification
type Reaction struct {
	Created_at string `json:"created_at"`
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	User GeneratedType_Nullable_simple_user `json:"user"` // A GitHub user.
	Content string `json:"content"` // The reaction to use
}

// GeneratedType_Webhook_branch_protection_rule_deleted represents the GeneratedType_Webhook_branch_protection_rule_deleted schema from the OpenAPI specification
type GeneratedType_Webhook_branch_protection_rule_deleted struct {
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Rule Webhooksrule `json:"rule"` // The branch protection rule. Includes a `name` and all the [branch protection settings](https://docs.github.com/github/administering-a-repository/defining-the-mergeability-of-pull-requests/about-protected-branches#about-branch-protection-settings) applied to branches that match the name. Binary settings are boolean. Multi-level configurations are one of `off`, `non_admins`, or `everyone`. Actor and build lists are arrays of strings.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType_Webhook_label_edited represents the GeneratedType_Webhook_label_edited schema from the OpenAPI specification
type GeneratedType_Webhook_label_edited struct {
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes,omitempty"` // The changes to the label if the action was `edited`.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Label Webhookslabel `json:"label"`
}

// GeneratedType_Protected_branch_required_status_check represents the GeneratedType_Protected_branch_required_status_check schema from the OpenAPI specification
type GeneratedType_Protected_branch_required_status_check struct {
	Enforcement_level string `json:"enforcement_level,omitempty"`
	Strict bool `json:"strict,omitempty"`
	Url string `json:"url,omitempty"`
	Checks []map[string]interface{} `json:"checks"`
	Contexts []string `json:"contexts"`
	Contexts_url string `json:"contexts_url,omitempty"`
}

// GeneratedType_Secret_scanning_location_pull_request_review represents the GeneratedType_Secret_scanning_location_pull_request_review schema from the OpenAPI specification
type GeneratedType_Secret_scanning_location_pull_request_review struct {
	Pull_request_review_url string `json:"pull_request_review_url"` // The API URL to get the pull request review where the secret was detected.
}

// GeneratedType_Nullable_collaborator represents the GeneratedType_Nullable_collaborator schema from the OpenAPI specification
type GeneratedType_Nullable_collaborator struct {
	TypeField string `json:"type"`
	Email string `json:"email,omitempty"`
	Gists_url string `json:"gists_url"`
	User_view_type string `json:"user_view_type,omitempty"`
	Following_url string `json:"following_url"`
	Gravatar_id string `json:"gravatar_id"`
	Organizations_url string `json:"organizations_url"`
	Starred_url string `json:"starred_url"`
	Node_id string `json:"node_id"`
	Role_name string `json:"role_name"`
	Events_url string `json:"events_url"`
	Followers_url string `json:"followers_url"`
	Login string `json:"login"`
	Id int64 `json:"id"`
	Received_events_url string `json:"received_events_url"`
	Repos_url string `json:"repos_url"`
	Site_admin bool `json:"site_admin"`
	Name string `json:"name,omitempty"`
	Url string `json:"url"`
	Html_url string `json:"html_url"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Avatar_url string `json:"avatar_url"`
	Subscriptions_url string `json:"subscriptions_url"`
}

// Collaborator represents the Collaborator schema from the OpenAPI specification
type Collaborator struct {
	Starred_url string `json:"starred_url"`
	Subscriptions_url string `json:"subscriptions_url"`
	Node_id string `json:"node_id"`
	Received_events_url string `json:"received_events_url"`
	TypeField string `json:"type"`
	Events_url string `json:"events_url"`
	User_view_type string `json:"user_view_type,omitempty"`
	Site_admin bool `json:"site_admin"`
	Role_name string `json:"role_name"`
	Followers_url string `json:"followers_url"`
	Gists_url string `json:"gists_url"`
	Html_url string `json:"html_url"`
	Email string `json:"email,omitempty"`
	Gravatar_id string `json:"gravatar_id"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Avatar_url string `json:"avatar_url"`
	Repos_url string `json:"repos_url"`
	Following_url string `json:"following_url"`
	Organizations_url string `json:"organizations_url"`
	Url string `json:"url"`
	Id int64 `json:"id"`
	Login string `json:"login"`
	Name string `json:"name,omitempty"`
}

// GeneratedType_Repository_rule_commit_message_pattern represents the GeneratedType_Repository_rule_commit_message_pattern schema from the OpenAPI specification
type GeneratedType_Repository_rule_commit_message_pattern struct {
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	TypeField string `json:"type"`
}

// GeneratedType_Code_scanning_codeql_database represents the GeneratedType_Code_scanning_codeql_database schema from the OpenAPI specification
type GeneratedType_Code_scanning_codeql_database struct {
	Created_at string `json:"created_at"` // The date and time at which the CodeQL database was created, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Uploader GeneratedType_Simple_user `json:"uploader"` // A GitHub user.
	Id int `json:"id"` // The ID of the CodeQL database.
	Name string `json:"name"` // The name of the CodeQL database.
	Content_type string `json:"content_type"` // The MIME type of the CodeQL database file.
	Language string `json:"language"` // The language of the CodeQL database.
	Url string `json:"url"` // The URL at which to download the CodeQL database. The `Accept` header must be set to the value of the `content_type` property.
	Commit_oid string `json:"commit_oid,omitempty"` // The commit SHA of the repository at the time the CodeQL database was created.
	Size int `json:"size"` // The size of the CodeQL database file in bytes.
	Updated_at string `json:"updated_at"` // The date and time at which the CodeQL database was last updated, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
}

// GeneratedType_Gist_history represents the GeneratedType_Gist_history schema from the OpenAPI specification
type GeneratedType_Gist_history struct {
	Url string `json:"url,omitempty"`
	User GeneratedType_Nullable_simple_user `json:"user,omitempty"` // A GitHub user.
	Version string `json:"version,omitempty"`
	Change_status map[string]interface{} `json:"change_status,omitempty"`
	Committed_at string `json:"committed_at,omitempty"`
}
