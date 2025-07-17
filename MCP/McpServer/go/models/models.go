package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Login string `json:"login"`
	Url string `json:"url"`
	Avatar_url string `json:"avatar_url"`
	Description string `json:"description"`
	Hooks_url string `json:"hooks_url"`
	Issues_url string `json:"issues_url"`
	Members_url string `json:"members_url"`
	Events_url string `json:"events_url"`
	Id int `json:"id"`
	Repos_url string `json:"repos_url"`
	Node_id string `json:"node_id"`
	Public_members_url string `json:"public_members_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Copilot_dotcom_pull_requests GeneratedType `json:"copilot_dotcom_pull_requests,omitempty"` // Usage metrics for Copilot for pull requests.
	Copilot_ide_chat GeneratedType `json:"copilot_ide_chat,omitempty"` // Usage metrics for Copilot Chat in the IDE.
	Copilot_ide_code_completions GeneratedType `json:"copilot_ide_code_completions,omitempty"` // Usage metrics for Copilot editor code completions in the IDE.
	Date string `json:"date"` // The date for which the usage metrics are aggregated, in `YYYY-MM-DD` format.
	Total_active_users int `json:"total_active_users,omitempty"` // The total number of Copilot users with activity belonging to any Copilot feature, globally, for the given day. Includes passive activity such as receiving a code suggestion, as well as engagement activity such as accepting a code suggestion or prompting chat. Does not include authentication events. Is not limited to the individual features detailed on the endpoint.
	Total_engaged_users int `json:"total_engaged_users,omitempty"` // The total number of Copilot users who engaged with any Copilot feature, for the given day. Examples include but are not limited to accepting a code suggestion, prompting Copilot chat, or triggering a PR Summary. Does not include authentication events. Is not limited to the individual features detailed on the endpoint.
	Copilot_dotcom_chat GeneratedType `json:"copilot_dotcom_chat,omitempty"` // Usage metrics for Copilot Chat in GitHub.com
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Number int `json:"number"` // The pull request number.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Assignee Webhooksuser `json:"assignee"`
	Pull_request map[string]interface{} `json:"pull_request"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Created_at string `json:"created_at"`
	Name string `json:"name"` // The name of the secret.
	Updated_at string `json:"updated_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Milestone Webhooksmilestone `json:"milestone"` // A collection of related issues and pull requests.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Release Webhooksrelease1 `json:"release"` // The [release](https://docs.github.com/rest/releases/releases/#get-a-release) object.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Cve_id string `json:"cve_id,omitempty"` // The Common Vulnerabilities and Exposures (CVE) ID.
	Cvss_vector_string string `json:"cvss_vector_string,omitempty"` // The CVSS vector that calculates the severity of the advisory. You must choose between setting this field or `severity`.
	Severity string `json:"severity,omitempty"` // The severity of the advisory. You must choose between setting this field or `cvss_vector_string`.
	Credits []map[string]interface{} `json:"credits,omitempty"` // A list of users receiving credit for their participation in the security advisory.
	Vulnerabilities []map[string]interface{} `json:"vulnerabilities"` // A product affected by the vulnerability detailed in a repository security advisory.
	Cwe_ids []string `json:"cwe_ids,omitempty"` // A list of Common Weakness Enumeration (CWE) IDs.
	Description string `json:"description"` // A detailed description of what the advisory impacts.
	Start_private_fork bool `json:"start_private_fork,omitempty"` // Whether to create a temporary private fork of the repository to collaborate on a fix.
	Summary string `json:"summary"` // A short summary of the advisory.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Milestone Webhooksmilestone3 `json:"milestone"` // A collection of related issues and pull requests.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Commit_count int `json:"commit_count"` // Count of student commits.
	Grade string `json:"grade"` // Most recent grade.
	Id int `json:"id"` // Unique identifier of the repository.
	Passing bool `json:"passing"` // Whether a submission passed.
	Repository GeneratedType `json:"repository"` // A GitHub repository view for Classroom
	Students []GeneratedType `json:"students"`
	Submitted bool `json:"submitted"` // Whether an accepted assignment has been submitted.
	Assignment GeneratedType `json:"assignment"` // A GitHub Classroom assignment
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Html_url string `json:"html_url"`
	Key string `json:"key"`
	Name string `json:"name"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Database_commit_sha string `json:"database_commit_sha,omitempty"` // The SHA of the commit the CodeQL database was built against. This is only available for successful analyses.
	Failure_message string `json:"failure_message,omitempty"` // The reason of the failure of this repo task. This is only available if the repository task has failed.
	Repository GeneratedType `json:"repository"` // A GitHub repository.
	Result_count int `json:"result_count,omitempty"` // The number of results in the case of a successful analysis. This is only available for successful analyses.
	Source_location_prefix string `json:"source_location_prefix,omitempty"` // The source location prefix to use. This is only available for successful analyses.
	Analysis_status string `json:"analysis_status"` // The new status of the CodeQL variant analysis repository task.
	Artifact_size_in_bytes int `json:"artifact_size_in_bytes,omitempty"` // The size of the artifact. This is only available for successful analyses.
	Artifact_url string `json:"artifact_url,omitempty"` // The URL of the artifact. This is only available for successful analyses.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Commit_oid string `json:"commit_oid"` // The commit SHA of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Alert map[string]interface{} `json:"alert"` // The code scanning alert involved in the event.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Ref string `json:"ref"` // The Git reference of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Blocked_user Webhooksuser `json:"blocked_user"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Release Webhooksrelease `json:"release"` // The [release](https://docs.github.com/rest/releases/releases/#get-a-release) object.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Release Webhooksrelease `json:"release"` // The [release](https://docs.github.com/rest/releases/releases/#get-a-release) object.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Actor map[string]interface{} `json:"actor"` // The actor who updated the ruleset
	Updated_at string `json:"updated_at"`
	Version_id int `json:"version_id"` // The ID of the previous version of the ruleset
	State map[string]interface{} `json:"state"` // The state of the ruleset version
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Hook map[string]interface{} `json:"hook"` // The deleted webhook. This will contain different keys based on the type of webhook it is: repository, organization, business, app, or GitHub Marketplace.
	Hook_id int `json:"hook_id"` // The id of the modified webhook.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Token string `json:"token"` // The token used for authentication
	Expires_at string `json:"expires_at"` // The time this token expires
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Repositories []Repository `json:"repositories,omitempty"` // The repositories this token has access to
	Repository_selection string `json:"repository_selection,omitempty"` // Describe whether all repositories have been selected or there's a selection involved
	Single_file string `json:"single_file,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Target string `json:"target"`
	_links map[string]interface{} `json:"_links"`
	Git_url string `json:"git_url"`
	Html_url string `json:"html_url"`
	Size int `json:"size"`
	Url string `json:"url"`
	Download_url string `json:"download_url"`
	Sha string `json:"sha"`
	Path string `json:"path"`
	Type string `json:"type"`
	Name string `json:"name"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization_url string `json:"organization_url"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Role string `json:"role"` // The user's membership type in the organization.
	State string `json:"state"` // The state of the member in the organization. The `pending` state indicates the user has not yet accepted an invitation.
	Url string `json:"url"`
	User GeneratedType `json:"user"` // A GitHub user.
	Organization GeneratedType `json:"organization"` // A GitHub organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sha string `json:"sha"`
	Type string `json:"type"`
	Url string `json:"url"`
	_links map[string]interface{} `json:"_links"`
	Download_url string `json:"download_url"`
	Html_url string `json:"html_url"`
	Name string `json:"name"`
	Encoding string `json:"encoding"`
	Size int `json:"size"`
	Git_url string `json:"git_url"`
	License GeneratedType `json:"license"` // License Simple
	Content string `json:"content"`
	Path string `json:"path"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Milestones_url string `json:"milestones_url"` // A template for the API URL to get information about milestones of the repository.
	Contributors_url string `json:"contributors_url"` // A template for the API URL to list the contributors to the repository.
	Git_commits_url string `json:"git_commits_url"` // A template for the API URL to get information about Git commits of the repository.
	Issue_events_url string `json:"issue_events_url"` // A template for the API URL to get information about issue events on the repository.
	Contents_url string `json:"contents_url"` // A template for the API URL to get the contents of the repository.
	Statuses_url string `json:"statuses_url"` // A template for the API URL to get information about statuses of a commit.
	Assignees_url string `json:"assignees_url"` // A template for the API URL to list the available assignees for issues in the repository.
	Subscription_url string `json:"subscription_url"` // The API URL to subscribe to notifications for this repository.
	Keys_url string `json:"keys_url"` // A template for the API URL to get information about deploy keys on the repository.
	Compare_url string `json:"compare_url"` // A template for the API URL to compare two commits or refs.
	Id int64 `json:"id"` // A unique identifier of the repository.
	Tags_url string `json:"tags_url"` // The API URL to get information about tags on the repository.
	Description string `json:"description"` // The repository description.
	Downloads_url string `json:"downloads_url"` // The API URL to list the downloads on the repository.
	Forks_url string `json:"forks_url"` // The API URL to list the forks of the repository.
	Pulls_url string `json:"pulls_url"` // A template for the API URL to get information about pull requests on the repository.
	Stargazers_url string `json:"stargazers_url"` // The API URL to list the stargazers on the repository.
	Comments_url string `json:"comments_url"` // A template for the API URL to get information about comments on the repository.
	Blobs_url string `json:"blobs_url"` // A template for the API URL to create or retrieve a raw Git blob in the repository.
	Hooks_url string `json:"hooks_url"` // The API URL to list the hooks on the repository.
	Subscribers_url string `json:"subscribers_url"` // The API URL to list the subscribers on the repository.
	Merges_url string `json:"merges_url"` // The API URL to merge branches in the repository.
	Private bool `json:"private"` // Whether the repository is private.
	Branches_url string `json:"branches_url"` // A template for the API URL to get information about branches in the repository.
	Archive_url string `json:"archive_url"` // A template for the API URL to download the repository as an archive.
	Releases_url string `json:"releases_url"` // A template for the API URL to get information about releases on the repository.
	Full_name string `json:"full_name"` // The full, globally unique, name of the repository.
	Labels_url string `json:"labels_url"` // A template for the API URL to get information about labels of the repository.
	Name string `json:"name"` // The name of the repository.
	Node_id string `json:"node_id"` // The GraphQL identifier of the repository.
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Teams_url string `json:"teams_url"` // The API URL to list the teams on the repository.
	Url string `json:"url"` // The URL to get more information about the repository from the GitHub API.
	Html_url string `json:"html_url"` // The URL to view the repository on GitHub.com.
	Fork bool `json:"fork"` // Whether the repository is a fork.
	Commits_url string `json:"commits_url"` // A template for the API URL to get information about commits on the repository.
	Issue_comment_url string `json:"issue_comment_url"` // A template for the API URL to get information about issue comments on the repository.
	Trees_url string `json:"trees_url"` // A template for the API URL to create or retrieve a raw Git tree of the repository.
	Notifications_url string `json:"notifications_url"` // A template for the API URL to get information about notifications on the repository.
	Collaborators_url string `json:"collaborators_url"` // A template for the API URL to get information about collaborators of the repository.
	Git_refs_url string `json:"git_refs_url"` // A template for the API URL to get information about Git refs of the repository.
	Deployments_url string `json:"deployments_url"` // The API URL to list the deployments of the repository.
	Events_url string `json:"events_url"` // The API URL to list the events of the repository.
	Languages_url string `json:"languages_url"` // The API URL to get information about the languages of the repository.
	Git_tags_url string `json:"git_tags_url"` // A template for the API URL to get information about Git tags of the repository.
	Issues_url string `json:"issues_url"` // A template for the API URL to get information about issues on the repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Oid string `json:"oid"` // Full or abbreviated commit hash to reject
	Reason string `json:"reason,omitempty"` // Reason for restriction
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Source_type string `json:"source_type,omitempty"` // The type of the source of the ruleset
	Target string `json:"target,omitempty"` // The target of the ruleset
	_links map[string]interface{} `json:"_links,omitempty"`
	Bypass_actors []GeneratedType `json:"bypass_actors,omitempty"` // The actors that can bypass the rules in this ruleset
	Conditions interface{} `json:"conditions,omitempty"`
	Current_user_can_bypass string `json:"current_user_can_bypass,omitempty"` // The bypass type of the user making the API request for this ruleset. This field is only returned when querying the repository-level endpoint.
	Source string `json:"source"` // The name of the source
	Updated_at string `json:"updated_at,omitempty"`
	Created_at string `json:"created_at,omitempty"`
	Enforcement string `json:"enforcement"` // The enforcement level of the ruleset. `evaluate` allows admins to test rules before enforcing them. Admins can view insights on the Rule Insights page (`evaluate` is only available with GitHub Enterprise).
	Id int `json:"id"` // The ID of the ruleset
	Name string `json:"name"` // The name of the ruleset
	Node_id string `json:"node_id,omitempty"`
	Rules []GeneratedType `json:"rules,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Repository_advisory GeneratedType `json:"repository_advisory"` // A repository security advisory.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	Type string `json:"type"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Label Webhookslabel `json:"label,omitempty"`
	Changes map[string]interface{} `json:"changes"` // The changes to the issue.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// Webhooksrelease1 represents the Webhooksrelease1 schema from the OpenAPI specification
type Webhooksrelease1 struct {
	Assets_url string `json:"assets_url"`
	Tarball_url string `json:"tarball_url"`
	Author map[string]interface{} `json:"author"`
	Published_at string `json:"published_at"`
	Zipball_url string `json:"zipball_url"`
	Node_id string `json:"node_id"`
	Target_commitish string `json:"target_commitish"` // Specifies the commitish value that determines where the Git tag is created from.
	Id int `json:"id"`
	Reactions map[string]interface{} `json:"reactions,omitempty"`
	Tag_name string `json:"tag_name"` // The name of the tag.
	Prerelease bool `json:"prerelease"` // Whether the release is identified as a prerelease or a full release.
	Discussion_url string `json:"discussion_url,omitempty"`
	Draft bool `json:"draft"` // Whether the release is a draft or published
	Name string `json:"name"`
	Body string `json:"body"`
	Assets []map[string]interface{} `json:"assets"`
	Upload_url string `json:"upload_url"`
	Url string `json:"url"`
	Created_at string `json:"created_at"`
	Html_url string `json:"html_url"`
}

// Webhooksprojectcard represents the Webhooksprojectcard schema from the OpenAPI specification
type Webhooksprojectcard struct {
	Column_id int `json:"column_id"`
	Created_at string `json:"created_at"`
	Creator map[string]interface{} `json:"creator"`
	Id int `json:"id"` // The project card's ID
	Node_id string `json:"node_id"`
	Note string `json:"note"`
	After_id int `json:"after_id,omitempty"`
	Archived bool `json:"archived"` // Whether or not the card is archived
	Url string `json:"url"`
	Column_url string `json:"column_url"`
	Content_url string `json:"content_url,omitempty"`
	Project_url string `json:"project_url"`
	Updated_at string `json:"updated_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Alert GeneratedType `json:"alert"` // A Dependabot alert.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Created_at string `json:"created_at"`
	Event string `json:"event"`
	Source map[string]interface{} `json:"source"`
	Updated_at string `json:"updated_at"`
	Actor GeneratedType `json:"actor,omitempty"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	_links map[string]interface{} `json:"_links"`
	Download_url string `json:"download_url"`
	Html_url string `json:"html_url"`
	Path string `json:"path"`
	Type string `json:"type"`
	Git_url string `json:"git_url"`
	Name string `json:"name"`
	Size int `json:"size"`
	Submodule_git_url string `json:"submodule_git_url"`
	Sha string `json:"sha"`
	Url string `json:"url"`
}

// Runner represents the Runner schema from the OpenAPI specification
type Runner struct {
	Labels []GeneratedType `json:"labels"`
	Name string `json:"name"` // The name of the runner.
	Os string `json:"os"` // The Operating System of the runner.
	Runner_group_id int `json:"runner_group_id,omitempty"` // The ID of the runner group.
	Status string `json:"status"` // The status of the runner.
	Busy bool `json:"busy"`
	Ephemeral bool `json:"ephemeral,omitempty"`
	Id int `json:"id"` // The ID of the runner.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Start_line int `json:"start_line,omitempty"`
	End_column int `json:"end_column,omitempty"`
	End_line int `json:"end_line,omitempty"`
	Path string `json:"path,omitempty"`
	Start_column int `json:"start_column,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sha string `json:"sha"`
	Merged bool `json:"merged"`
	Message string `json:"message"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Context string `json:"context"` // The status check context name that must be present on the commit.
	Integration_id int `json:"integration_id,omitempty"` // The optional integration ID that this status check must originate from.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Milestone Webhooksmilestone `json:"milestone"` // A collection of related issues and pull requests.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Reason string `json:"reason"` // Reason for requesting access.
	Id int `json:"id"` // Unique identifier of the request for access via fine-grained personal access token. The `pat_request_id` used to review PAT requests.
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Token_expired bool `json:"token_expired"` // Whether the associated fine-grained personal access token has expired.
	Token_last_used_at string `json:"token_last_used_at"` // Date and time when the associated fine-grained personal access token was last used for authentication.
	Token_name string `json:"token_name"` // The name given to the user's token. This field can also be found in an organization's settings page for Active Tokens.
	Repositories_url string `json:"repositories_url"` // URL to the list of repositories requested to be accessed via fine-grained personal access token. Should only be followed when `repository_selection` is `subset`.
	Repository_selection string `json:"repository_selection"` // Type of repository selection requested.
	Token_expires_at string `json:"token_expires_at"` // Date and time when the associated fine-grained personal access token expires.
	Token_id int `json:"token_id"` // Unique identifier of the user's token. This field can also be found in audit log events and the organization's settings for their PAT grants.
	Created_at string `json:"created_at"` // Date and time when the request for access was created.
	Permissions map[string]interface{} `json:"permissions"` // Permissions requested, categorized by type of permission.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Commit_id string `json:"commit_id"`
	Commit_url string `json:"commit_url"`
	Milestone map[string]interface{} `json:"milestone"`
	Url string `json:"url"`
	Event string `json:"event"`
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Created_at string `json:"created_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Updated_at string `json:"updated_at"`
	Browser_download_url string `json:"browser_download_url"`
	Digest string `json:"digest"`
	State string `json:"state"` // State of the release asset.
	Uploader GeneratedType `json:"uploader"` // A GitHub user.
	Content_type string `json:"content_type"`
	Name string `json:"name"` // The file name of the asset.
	Node_id string `json:"node_id"`
	Id int `json:"id"`
	Label string `json:"label"`
	Url string `json:"url"`
	Created_at string `json:"created_at"`
	Download_count int `json:"download_count"`
	Size int `json:"size"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Package map[string]interface{} `json:"package"` // Information about the package.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Editors []map[string]interface{} `json:"editors,omitempty"`
	Total_engaged_users int `json:"total_engaged_users,omitempty"` // Total number of users who prompted Copilot Chat in the IDE.
}

// Artifact represents the Artifact schema from the OpenAPI specification
type Artifact struct {
	Digest string `json:"digest,omitempty"` // The SHA256 digest of the artifact. This field will only be populated on artifacts uploaded with upload-artifact v4 or newer. For older versions, this field will be null.
	Expires_at string `json:"expires_at"`
	Archive_download_url string `json:"archive_download_url"`
	Id int `json:"id"`
	Created_at string `json:"created_at"`
	Expired bool `json:"expired"` // Whether or not the artifact has expired.
	Size_in_bytes int `json:"size_in_bytes"` // The size in bytes of the artifact.
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	Name string `json:"name"` // The name of the artifact.
	Node_id string `json:"node_id"`
	Workflow_run map[string]interface{} `json:"workflow_run,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	State_reason string `json:"state_reason,omitempty"` // The reason for the current state
	Created_at string `json:"created_at"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Active_lock_reason string `json:"active_lock_reason,omitempty"`
	Node_id string `json:"node_id"`
	Html_url string `json:"html_url"`
	Number int `json:"number"` // Number uniquely identifying the issue within its repository
	Assignee GeneratedType `json:"assignee"` // A GitHub user.
	State string `json:"state"` // State of the issue; either 'open' or 'closed'
	Timeline_url string `json:"timeline_url,omitempty"`
	Url string `json:"url"` // URL for the issue
	Updated_at string `json:"updated_at"`
	Sub_issues_summary GeneratedType `json:"sub_issues_summary,omitempty"`
	User GeneratedType `json:"user"` // A GitHub user.
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Comments int `json:"comments"`
	Closed_by GeneratedType `json:"closed_by,omitempty"` // A GitHub user.
	Title string `json:"title"` // Title of the issue
	Assignees []GeneratedType `json:"assignees,omitempty"`
	Body string `json:"body,omitempty"` // Contents of the issue
	Closed_at string `json:"closed_at"`
	Comments_url string `json:"comments_url"`
	Draft bool `json:"draft,omitempty"`
	Labels_url string `json:"labels_url"`
	Reactions GeneratedType `json:"reactions,omitempty"`
	Body_text string `json:"body_text,omitempty"`
	Milestone GeneratedType `json:"milestone"` // A collection of related issues and pull requests.
	Pull_request map[string]interface{} `json:"pull_request,omitempty"`
	Labels []interface{} `json:"labels"` // Labels to associate with this issue; pass one or more label names to replace the set of labels on this issue; send an empty array to clear all labels from the issue; note that the labels are silently dropped for users without push access to the repository
	Locked bool `json:"locked"`
	Type GeneratedType `json:"type,omitempty"` // The type of issue.
	Id int64 `json:"id"`
	Repository_url string `json:"repository_url"`
	Body_html string `json:"body_html,omitempty"`
	Events_url string `json:"events_url"`
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Percentage float64 `json:"percentage,omitempty"`
	Percentile float64 `json:"percentile,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Comment map[string]interface{} `json:"comment"` // The [comment](https://docs.github.com/rest/pulls/comments#get-a-review-comment-for-a-pull-request) itself.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Type string `json:"type"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	State string `json:"state"` // The state of the user's acceptance of the credit.
	Type string `json:"type"` // The type of credit the user is receiving.
	User GeneratedType `json:"user"` // A GitHub user.
}

// Autolink represents the Autolink schema from the OpenAPI specification
type Autolink struct {
	Id int `json:"id"`
	Is_alphanumeric bool `json:"is_alphanumeric"` // Whether this autolink reference matches alphanumeric characters. If false, this autolink reference only matches numeric characters.
	Key_prefix string `json:"key_prefix"` // The prefix of a key that is linkified.
	Url_template string `json:"url_template"` // A template for the target URL that is generated if a key was found.
}

// Webhookslabel represents the Webhookslabel schema from the OpenAPI specification
type Webhookslabel struct {
	Description string `json:"description"`
	Id int `json:"id"`
	Name string `json:"name"` // The name of the label.
	Node_id string `json:"node_id"`
	Url string `json:"url"` // URL for the label
	Color string `json:"color"` // 6-character hex code, without the leading #, identifying the color
	Default bool `json:"default"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Multi_repo bool `json:"multi_repo,omitempty"` // Whether the detected secret was found in multiple repositories under the same organization or enterprise.
	Push_protection_bypass_request_html_url string `json:"push_protection_bypass_request_html_url,omitempty"` // The URL to a push protection bypass request.
	State string `json:"state,omitempty"` // Sets the state of the secret scanning alert. You must provide `resolution` when you set the state to `resolved`.
	Resolution string `json:"resolution,omitempty"` // **Required when the `state` is `resolved`.** The reason for resolving the alert.
	Resolution_comment string `json:"resolution_comment,omitempty"` // An optional comment to resolve an alert.
	Locations_url string `json:"locations_url,omitempty"` // The REST API URL of the code locations for this alert.
	Push_protection_bypass_request_comment string `json:"push_protection_bypass_request_comment,omitempty"` // An optional comment when requesting a push protection bypass.
	Resolved_at string `json:"resolved_at,omitempty"` // The time that the alert was resolved in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Validity string `json:"validity,omitempty"` // The token status as of the latest validity check.
	Is_base64_encoded bool `json:"is_base64_encoded,omitempty"` // A boolean value representing whether or not alert is base64 encoded
	Push_protection_bypassed bool `json:"push_protection_bypassed,omitempty"` // Whether push protection was bypassed for the detected secret.
	Updated_at string `json:"updated_at,omitempty"` // The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Push_protection_bypass_request_reviewer_comment string `json:"push_protection_bypass_request_reviewer_comment,omitempty"` // An optional comment when reviewing a push protection bypass.
	Url string `json:"url,omitempty"` // The REST API URL of the alert resource.
	Secret string `json:"secret,omitempty"` // The secret that was detected.
	Publicly_leaked bool `json:"publicly_leaked,omitempty"` // Whether the detected secret was publicly leaked.
	Created_at string `json:"created_at,omitempty"` // The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Html_url string `json:"html_url,omitempty"` // The GitHub URL of the alert resource.
	Resolved_by GeneratedType `json:"resolved_by,omitempty"` // A GitHub user.
	Number int `json:"number,omitempty"` // The security alert number.
	Secret_type_display_name string `json:"secret_type_display_name,omitempty"` // User-friendly name for the detected secret, matching the `secret_type`. For a list of built-in patterns, see "[Supported secret scanning patterns](https://docs.github.com/code-security/secret-scanning/introduction/supported-secret-scanning-patterns#supported-secrets)."
	Push_protection_bypass_request_reviewer GeneratedType `json:"push_protection_bypass_request_reviewer,omitempty"` // A GitHub user.
	Secret_type string `json:"secret_type,omitempty"` // The type of secret that secret scanning detected.
	Has_more_locations bool `json:"has_more_locations,omitempty"` // A boolean value representing whether or not the token in the alert was detected in more than one location.
	Push_protection_bypassed_by GeneratedType `json:"push_protection_bypassed_by,omitempty"` // A GitHub user.
	Push_protection_bypassed_at string `json:"push_protection_bypassed_at,omitempty"` // The time that push protection was bypassed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	First_location_detected GeneratedType `json:"first_location_detected,omitempty"` // Details on the location where the token was initially detected. This can be a commit, wiki commit, issue, discussion, pull request.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Project_column Webhooksprojectcolumn `json:"project_column"`
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Node_id string `json:"node_id"`
	Event string `json:"event"`
	Updated_at string `json:"updated_at"`
	Created_at string `json:"created_at"`
	Body_html string `json:"body_html,omitempty"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Id int `json:"id"` // Unique identifier of the issue comment
	Performed_via_github_app GeneratedType `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Reactions GeneratedType `json:"reactions,omitempty"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Issue_url string `json:"issue_url"`
	Html_url string `json:"html_url"`
	User GeneratedType `json:"user"` // A GitHub user.
	Body_text string `json:"body_text,omitempty"`
	Url string `json:"url"` // URL for the issue comment
	Body string `json:"body,omitempty"` // Contents of the issue comment
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Workflow_job interface{} `json:"workflow_job"`
	Action string `json:"action"`
	Deployment Deployment `json:"deployment,omitempty"` // A request for a specific ref(branch,sha,tag) to be deployed
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Payload string `json:"payload"` // A URL-encoded string of the check_run.requested_action JSON payload. The decoded payload is a JSON object.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Email string `json:"email,omitempty"`
	Name string `json:"name,omitempty"`
	Date string `json:"date,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Starred_url string `json:"starred_url"`
	Assignment string `json:"assignment,omitempty"` // Determines if the user has a direct, indirect, or mixed relationship to a role
	Login string `json:"login"`
	Type string `json:"type"`
	Received_events_url string `json:"received_events_url"`
	Followers_url string `json:"followers_url"`
	Gists_url string `json:"gists_url"`
	Html_url string `json:"html_url"`
	User_view_type string `json:"user_view_type,omitempty"`
	Organizations_url string `json:"organizations_url"`
	Name string `json:"name,omitempty"`
	Avatar_url string `json:"avatar_url"`
	Events_url string `json:"events_url"`
	Inherited_from []GeneratedType `json:"inherited_from,omitempty"` // Team the user has gotten the role through
	Node_id string `json:"node_id"`
	Repos_url string `json:"repos_url"`
	Url string `json:"url"`
	Email string `json:"email,omitempty"`
	Gravatar_id string `json:"gravatar_id"`
	Id int `json:"id"`
	Following_url string `json:"following_url"`
	Starred_at string `json:"starred_at,omitempty"`
	Subscriptions_url string `json:"subscriptions_url"`
	Site_admin bool `json:"site_admin"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id int `json:"id"` // Unique identifier of the classroom.
	Name string `json:"name"` // The name of the classroom.
	Url string `json:"url"` // The url of the classroom on GitHub Classroom.
	Archived bool `json:"archived"` // Returns whether classroom is archived or not.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Oid string `json:"oid"`
	Path string `json:"path"`
	Ref_name string `json:"ref_name"`
	Size int `json:"size"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Type string `json:"type"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Display_name string `json:"display_name"` // Display name for this image.
	Id string `json:"id"` // The ID of the image. Use this ID for the `image` parameter when creating a new larger runner.
	Platform string `json:"platform"` // The operating system of the image.
	Size_gb int `json:"size_gb"` // Image size in GB.
	Source string `json:"source"` // The image provider.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Ruleset_source_type string `json:"ruleset_source_type,omitempty"` // The type of source for the ruleset that includes this rule.
	Ruleset_id int `json:"ruleset_id,omitempty"` // The ID of the ruleset that includes this rule.
	Ruleset_source string `json:"ruleset_source,omitempty"` // The name of the source of the ruleset that includes this rule.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Can_approve_pull_request_reviews bool `json:"can_approve_pull_request_reviews"` // Whether GitHub Actions can approve pull requests. Enabling this can be a security risk.
	Default_workflow_permissions string `json:"default_workflow_permissions"` // The default workflow permissions granted to the GITHUB_TOKEN when running workflows.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Issue_comment_url string `json:"issue_comment_url"` // The API URL to get the issue comment where the secret was detected.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Location string `json:"location"` // The initally assigned location of a new codespace.
	Display_name string `json:"display_name,omitempty"` // Display name for this codespace.
	Prebuild bool `json:"prebuild"` // Whether the codespace was created from a prebuild.
	Idle_timeout_minutes int `json:"idle_timeout_minutes"` // The number of minutes of inactivity after which this codespace will be automatically stopped.
	Billable_owner GeneratedType `json:"billable_owner"` // A GitHub user.
	Stop_url string `json:"stop_url"` // API URL to stop this codespace.
	Machines_url string `json:"machines_url"` // API URL to access available alternate machine types for this codespace.
	Id int64 `json:"id"`
	Runtime_constraints map[string]interface{} `json:"runtime_constraints,omitempty"`
	Pulls_url string `json:"pulls_url"` // API URL for the Pull Request associated with this codespace, if any.
	Web_url string `json:"web_url"` // URL to access this codespace on the web.
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Updated_at string `json:"updated_at"`
	Pending_operation bool `json:"pending_operation,omitempty"` // Whether or not a codespace has a pending async operation. This would mean that the codespace is temporarily unavailable. The only thing that you can do with a codespace in this state is delete it.
	Publish_url string `json:"publish_url,omitempty"` // API URL to publish this codespace to a new repository.
	Start_url string `json:"start_url"` // API URL to start this codespace.
	Pending_operation_disabled_reason string `json:"pending_operation_disabled_reason,omitempty"` // Text to show user when codespace is disabled by a pending operation
	Retention_expires_at string `json:"retention_expires_at,omitempty"` // When a codespace will be auto-deleted based on the "retention_period_minutes" and "last_used_at"
	Retention_period_minutes int `json:"retention_period_minutes,omitempty"` // Duration in minutes after codespace has gone idle in which it will be deleted. Must be integer minutes between 0 and 43200 (30 days).
	Name string `json:"name"` // Automatically generated name of this codespace.
	Last_used_at string `json:"last_used_at"` // Last known time this codespace was started.
	Environment_id string `json:"environment_id"` // UUID identifying this codespace's environment.
	Machine GeneratedType `json:"machine"` // A description of the machine powering a codespace.
	Recent_folders []string `json:"recent_folders"`
	Repository GeneratedType `json:"repository"` // Full Repository
	Devcontainer_path string `json:"devcontainer_path,omitempty"` // Path to devcontainer.json from repo root used to create Codespace.
	Git_status map[string]interface{} `json:"git_status"` // Details about the codespace's git repository.
	Idle_timeout_notice string `json:"idle_timeout_notice,omitempty"` // Text to show user when codespace idle timeout minutes has been overriden by an organization policy
	Created_at string `json:"created_at"`
	State string `json:"state"` // State of this codespace.
	Url string `json:"url"` // API URL for this codespace.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Commits []Commit `json:"commits"`
	Diff_url string `json:"diff_url"`
	Total_commits int `json:"total_commits"`
	Url string `json:"url"`
	Permalink_url string `json:"permalink_url"`
	Status string `json:"status"`
	Base_commit Commit `json:"base_commit"` // Commit
	Files []GeneratedType `json:"files,omitempty"`
	Html_url string `json:"html_url"`
	Merge_base_commit Commit `json:"merge_base_commit"` // Commit
	Ahead_by int `json:"ahead_by"`
	Behind_by int `json:"behind_by"`
	Patch_url string `json:"patch_url"`
}

// Thread represents the Thread schema from the OpenAPI specification
type Thread struct {
	Unread bool `json:"unread"`
	Repository GeneratedType `json:"repository"` // Minimal Repository
	Id string `json:"id"`
	Subscription_url string `json:"subscription_url"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	Last_read_at string `json:"last_read_at"`
	Reason string `json:"reason"`
	Subject map[string]interface{} `json:"subject"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Team Webhooksteam1 `json:"team"` // Groups of organization members that gives permissions on specified repositories.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository map[string]interface{} `json:"repository,omitempty"` // A git repository
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Name string `json:"name"` // The name of the package version.
	Description string `json:"description,omitempty"`
	Package_html_url string `json:"package_html_url"`
	Updated_at string `json:"updated_at"`
	Created_at string `json:"created_at"`
	License string `json:"license,omitempty"`
	Id int `json:"id"` // Unique identifier of the package version.
	Url string `json:"url"`
	Deleted_at string `json:"deleted_at,omitempty"`
	Html_url string `json:"html_url,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Dependency_graph string `json:"dependency_graph,omitempty"` // The enablement status of Dependency Graph
	Secret_scanning string `json:"secret_scanning,omitempty"` // The enablement status of secret scanning
	Target_type string `json:"target_type,omitempty"` // The type of the code security configuration.
	Updated_at string `json:"updated_at,omitempty"`
	Url string `json:"url,omitempty"` // The URL of the configuration
	Advanced_security string `json:"advanced_security,omitempty"` // The enablement status of GitHub Advanced Security
	Dependabot_security_updates string `json:"dependabot_security_updates,omitempty"` // The enablement status of Dependabot security updates
	Secret_scanning_delegated_bypass string `json:"secret_scanning_delegated_bypass,omitempty"` // The enablement status of secret scanning delegated bypass
	Private_vulnerability_reporting string `json:"private_vulnerability_reporting,omitempty"` // The enablement status of private vulnerability reporting
	Enforcement string `json:"enforcement,omitempty"` // The enforcement status for a security configuration
	Html_url string `json:"html_url,omitempty"` // The URL of the configuration
	Secret_scanning_delegated_alert_dismissal string `json:"secret_scanning_delegated_alert_dismissal,omitempty"` // The enablement status of secret scanning delegated alert dismissal
	Secret_scanning_delegated_bypass_options map[string]interface{} `json:"secret_scanning_delegated_bypass_options,omitempty"` // Feature options for secret scanning delegated bypass
	Secret_scanning_non_provider_patterns string `json:"secret_scanning_non_provider_patterns,omitempty"` // The enablement status of secret scanning non-provider patterns
	Secret_scanning_push_protection string `json:"secret_scanning_push_protection,omitempty"` // The enablement status of secret scanning push protection
	Code_scanning_delegated_alert_dismissal string `json:"code_scanning_delegated_alert_dismissal,omitempty"` // The enablement status of code scanning delegated alert dismissal
	Dependency_graph_autosubmit_action_options map[string]interface{} `json:"dependency_graph_autosubmit_action_options,omitempty"` // Feature options for Automatic dependency submission
	Id int `json:"id,omitempty"` // The ID of the code security configuration
	Name string `json:"name,omitempty"` // The name of the code security configuration. Must be unique within the organization.
	Dependency_graph_autosubmit_action string `json:"dependency_graph_autosubmit_action,omitempty"` // The enablement status of Automatic dependency submission
	Description string `json:"description,omitempty"` // A description of the code security configuration
	Created_at string `json:"created_at,omitempty"`
	Code_scanning_default_setup_options map[string]interface{} `json:"code_scanning_default_setup_options,omitempty"` // Feature options for code scanning default setup
	Code_scanning_default_setup string `json:"code_scanning_default_setup,omitempty"` // The enablement status of code scanning default setup
	Secret_scanning_generic_secrets string `json:"secret_scanning_generic_secrets,omitempty"` // The enablement status of Copilot secret scanning
	Secret_scanning_validity_checks string `json:"secret_scanning_validity_checks,omitempty"` // The enablement status of secret scanning validity checks
	Code_scanning_options map[string]interface{} `json:"code_scanning_options,omitempty"` // Feature options for code scanning
	Dependabot_alerts string `json:"dependabot_alerts,omitempty"` // The enablement status of Dependabot alerts
}

// Vulnerability represents the Vulnerability schema from the OpenAPI specification
type Vulnerability struct {
	Package map[string]interface{} `json:"package"` // The name of the package affected by the vulnerability.
	Vulnerable_functions []string `json:"vulnerable_functions"` // The functions in the package that are affected by the vulnerability.
	Vulnerable_version_range string `json:"vulnerable_version_range"` // The range of the package versions affected by the vulnerability.
	First_patched_version string `json:"first_patched_version"` // The package version that resolves the vulnerability.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Projects_v2_status_update GeneratedType `json:"projects_v2_status_update"` // An status update belonging to a project
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Access_level string `json:"access_level"` // Defines the level of access that workflows outside of the repository have to actions and reusable workflows within the repository. `none` means the access is only possible from workflows in this repository. `user` level access allows sharing across user owned private repositories only. `organization` level access allows sharing across the organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Created_at string `json:"created_at"`
	Event string `json:"event"`
	Id int `json:"id"`
	Url string `json:"url"`
	Assignee GeneratedType `json:"assignee"` // A GitHub user.
	Node_id string `json:"node_id"`
	Commit_id string `json:"commit_id"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Commit_url string `json:"commit_url"`
}

// Release represents the Release schema from the OpenAPI specification
type Release struct {
	Reactions GeneratedType `json:"reactions,omitempty"`
	Body string `json:"body,omitempty"`
	Discussion_url string `json:"discussion_url,omitempty"` // The URL of the release discussion.
	Name string `json:"name"`
	Created_at string `json:"created_at"`
	Prerelease bool `json:"prerelease"` // Whether to identify the release as a prerelease or a full release.
	Zipball_url string `json:"zipball_url"`
	Author GeneratedType `json:"author"` // A GitHub user.
	Upload_url string `json:"upload_url"`
	Tag_name string `json:"tag_name"` // The name of the tag.
	Assets_url string `json:"assets_url"`
	Body_text string `json:"body_text,omitempty"`
	Mentions_count int `json:"mentions_count,omitempty"`
	Draft bool `json:"draft"` // true to create a draft (unpublished) release, false to create a published one.
	Body_html string `json:"body_html,omitempty"`
	Published_at string `json:"published_at"`
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	Target_commitish string `json:"target_commitish"` // Specifies the commitish value that determines where the Git tag is created from.
	Url string `json:"url"`
	Assets []GeneratedType `json:"assets"`
	Html_url string `json:"html_url"`
	Tarball_url string `json:"tarball_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Parent_issue_id float64 `json:"parent_issue_id"` // The ID of the parent issue.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Action string `json:"action"`
	Parent_issue Issue `json:"parent_issue"` // Issues are a great way to keep track of tasks, enhancements, and bugs for your projects.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Sub_issue Issue `json:"sub_issue"` // Issues are a great way to keep track of tasks, enhancements, and bugs for your projects.
	Sub_issue_id float64 `json:"sub_issue_id"` // The ID of the sub-issue.
	Sub_issue_repo Repository `json:"sub_issue_repo"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Property_name string `json:"property_name"` // The name of the property
	Source_type string `json:"source_type,omitempty"` // The source type of the property
	Url string `json:"url,omitempty"` // The URL that can be used to fetch, update, or delete info about this property via the API.
	Required bool `json:"required,omitempty"` // Whether the property is required.
	Values_editable_by string `json:"values_editable_by,omitempty"` // Who can edit the values of the property
	Value_type string `json:"value_type"` // The type of the value for the property
	Allowed_values []string `json:"allowed_values,omitempty"` // An ordered list of the allowed values of the property. The property can have up to 200 allowed values.
	Default_value string `json:"default_value,omitempty"` // Default value of the property
	Description string `json:"description,omitempty"` // Short description of the property
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Project_card map[string]interface{} `json:"project_card"`
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Team Webhooksteam1 `json:"team"` // Groups of organization members that gives permissions on specified repositories.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository map[string]interface{} `json:"repository,omitempty"` // A git repository
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Requested_reviewer GeneratedType `json:"requested_reviewer,omitempty"` // A GitHub user.
	Requested_team Team `json:"requested_team,omitempty"` // Groups of organization members that gives permissions on specified repositories.
	Url string `json:"url"`
	Commit_id string `json:"commit_id"`
	Commit_url string `json:"commit_url"`
	Created_at string `json:"created_at"`
	Node_id string `json:"node_id"`
	Review_requester GeneratedType `json:"review_requester"` // A GitHub user.
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Event string `json:"event"`
	Id int `json:"id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Selected_repositories_url string `json:"selected_repositories_url"` // The API URL at which the list of repositories this secret is visible to can be retrieved
	Updated_at string `json:"updated_at"` // The date and time at which the secret was last updated, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Visibility string `json:"visibility"` // The type of repositories in the organization that the secret is visible to
	Created_at string `json:"created_at"` // The date and time at which the secret was created, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Name string `json:"name"` // The name of the secret
}

// Contributor represents the Contributor schema from the OpenAPI specification
type Contributor struct {
	Starred_url string `json:"starred_url,omitempty"`
	Repos_url string `json:"repos_url,omitempty"`
	Received_events_url string `json:"received_events_url,omitempty"`
	Following_url string `json:"following_url,omitempty"`
	Gravatar_id string `json:"gravatar_id,omitempty"`
	Type string `json:"type"`
	Contributions int `json:"contributions"`
	Site_admin bool `json:"site_admin,omitempty"`
	Followers_url string `json:"followers_url,omitempty"`
	Gists_url string `json:"gists_url,omitempty"`
	Login string `json:"login,omitempty"`
	Name string `json:"name,omitempty"`
	Node_id string `json:"node_id,omitempty"`
	Organizations_url string `json:"organizations_url,omitempty"`
	Subscriptions_url string `json:"subscriptions_url,omitempty"`
	Url string `json:"url,omitempty"`
	Id int `json:"id,omitempty"`
	User_view_type string `json:"user_view_type,omitempty"`
	Avatar_url string `json:"avatar_url,omitempty"`
	Email string `json:"email,omitempty"`
	Events_url string `json:"events_url,omitempty"`
	Html_url string `json:"html_url,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Total_paid_minutes_used int `json:"total_paid_minutes_used"` // The total paid GitHub Actions minutes used.
	Included_minutes int `json:"included_minutes"` // The amount of free GitHub Actions minutes available.
	Minutes_used_breakdown map[string]interface{} `json:"minutes_used_breakdown"`
	Total_minutes_used int `json:"total_minutes_used"` // The sum of the free and paid GitHub Actions minutes used.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id int `json:"id"`
	Key string `json:"key"`
	Created_at string `json:"created_at,omitempty"`
}

// Snapshot represents the Snapshot schema from the OpenAPI specification
type Snapshot struct {
	Scanned string `json:"scanned"` // The time at which the snapshot was scanned.
	Sha string `json:"sha"` // The commit SHA associated with this dependency snapshot. Maximum length: 40 characters.
	Version int `json:"version"` // The version of the repository snapshot submission.
	Detector map[string]interface{} `json:"detector"` // A description of the detector used.
	Job map[string]interface{} `json:"job"`
	Manifests map[string]interface{} `json:"manifests,omitempty"` // A collection of package manifests, which are a collection of related dependencies declared in a file or representing a logical group of dependencies.
	Metadata Metadata `json:"metadata,omitempty"` // User-defined metadata to store domain-specific information limited to 8 keys with scalar values.
	Ref string `json:"ref"` // The repository branch that triggered this snapshot.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Label Webhookslabel `json:"label"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes,omitempty"` // The changes to the label if the action was `edited`.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Run_id int `json:"run_id,omitempty"` // ID of the corresponding run.
	Run_url string `json:"run_url,omitempty"` // URL of the corresponding run.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Project_card Webhooksprojectcard `json:"project_card"`
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Color string `json:"color,omitempty"` // Color for the issue type.
	Description string `json:"description,omitempty"` // Description of the issue type.
	Is_enabled bool `json:"is_enabled"` // Whether or not the issue type is enabled at the organization level.
	Name string `json:"name"` // Name of the issue type.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Node_id string `json:"node_id"`
	Updated_at string `json:"updated_at"`
	Git_pull_url string `json:"git_pull_url"`
	Truncated bool `json:"truncated,omitempty"`
	Comments_enabled bool `json:"comments_enabled,omitempty"`
	Created_at string `json:"created_at"`
	History []interface{} `json:"history,omitempty"`
	Html_url string `json:"html_url"`
	Comments_url string `json:"comments_url"`
	User GeneratedType `json:"user"` // A GitHub user.
	Commits_url string `json:"commits_url"`
	Description string `json:"description"`
	Files map[string]interface{} `json:"files"`
	Owner GeneratedType `json:"owner,omitempty"` // A GitHub user.
	Public bool `json:"public"`
	Url string `json:"url"`
	Git_push_url string `json:"git_push_url"`
	Comments int `json:"comments"`
	Forks_url string `json:"forks_url"`
	Forks []interface{} `json:"forks,omitempty"`
	Id string `json:"id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Created_at string `json:"created_at"`
	Created_by string `json:"created_by"`
	Featured bool `json:"featured"`
	Released string `json:"released"`
	Display_name string `json:"display_name"`
	Aliases []map[string]interface{} `json:"aliases,omitempty"`
	Repository_count int `json:"repository_count,omitempty"`
	Short_description string `json:"short_description"`
	Name string `json:"name"`
	Text_matches []map[string]interface{} `json:"text_matches,omitempty"`
	Related []map[string]interface{} `json:"related,omitempty"`
	Updated_at string `json:"updated_at"`
	Curated bool `json:"curated"`
	Description string `json:"description"`
	Logo_url string `json:"logo_url,omitempty"`
	Score float64 `json:"score"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id string `json:"id"` // The ID of the GitHub Pages deployment. This is the Git SHA of the deployed commit.
	Page_url string `json:"page_url"` // The URI to the deployed GitHub Pages.
	Preview_url string `json:"preview_url,omitempty"` // The URI to the deployed GitHub Pages preview.
	Status_url string `json:"status_url"` // The URI to monitor GitHub Pages deployment status.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Guid string `json:"guid,omitempty"` // The GUID of the tool used to generate the code scanning analysis, if provided in the uploaded SARIF data.
	Name string `json:"name,omitempty"` // The name of the tool used to generate the code scanning analysis.
	Version string `json:"version,omitempty"` // The version of the tool used to generate the code scanning analysis.
}

// Team represents the Team schema from the OpenAPI specification
type Team struct {
	Name string `json:"name"`
	Notification_setting string `json:"notification_setting,omitempty"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Repositories_url string `json:"repositories_url"`
	Url string `json:"url"`
	Id int `json:"id"`
	Permission string `json:"permission"`
	Privacy string `json:"privacy,omitempty"`
	Description string `json:"description"`
	Node_id string `json:"node_id"`
	Parent GeneratedType `json:"parent"` // Groups of organization members that gives permissions on specified repositories.
	Slug string `json:"slug"`
	Html_url string `json:"html_url"`
	Members_url string `json:"members_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	To string `json:"to"`
	From string `json:"from"`
}

// Activity represents the Activity schema from the OpenAPI specification
type Activity struct {
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	Ref string `json:"ref"` // The full Git reference, formatted as `refs/heads/<branch name>`.
	Timestamp string `json:"timestamp"` // The time when the activity occurred.
	Activity_type string `json:"activity_type"` // The type of the activity that was performed.
	Actor GeneratedType `json:"actor"` // A GitHub user.
	After string `json:"after"` // The SHA of the commit after the activity.
	Before string `json:"before"` // The SHA of the commit before the activity.
}

// Language represents the Language schema from the OpenAPI specification
type Language struct {
}

// Authorization represents the Authorization schema from the OpenAPI specification
type Authorization struct {
	Created_at string `json:"created_at"`
	Expires_at string `json:"expires_at"`
	Id int64 `json:"id"`
	Installation GeneratedType `json:"installation,omitempty"`
	Scopes []string `json:"scopes"` // A list of scopes that this authorization is in.
	Token_last_eight string `json:"token_last_eight"`
	User GeneratedType `json:"user,omitempty"` // A GitHub user.
	App map[string]interface{} `json:"app"`
	Note string `json:"note"`
	Url string `json:"url"`
	Hashed_token string `json:"hashed_token"`
	Updated_at string `json:"updated_at"`
	Note_url string `json:"note_url"`
	Token string `json:"token"`
	Fingerprint string `json:"fingerprint"`
}

// Webhooksmarketplacepurchase represents the Webhooksmarketplacepurchase schema from the OpenAPI specification
type Webhooksmarketplacepurchase struct {
	On_free_trial bool `json:"on_free_trial"`
	Plan map[string]interface{} `json:"plan"`
	Unit_count int `json:"unit_count"`
	Account map[string]interface{} `json:"account"`
	Billing_cycle string `json:"billing_cycle"`
	Free_trial_ends_on string `json:"free_trial_ends_on"`
	Next_billing_date string `json:"next_billing_date"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Errors []string `json:"errors,omitempty"`
	Message string `json:"message"`
	Documentation_url string `json:"documentation_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Description string `json:"description"`
	Domains []string `json:"domains"` // Array of the domain set and its alternate name (if it is configured)
	Expires_at string `json:"expires_at,omitempty"`
	State string `json:"state"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Name string `json:"name"`
	Color string `json:"color"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Pull_request_review_comment_url string `json:"pull_request_review_comment_url"` // The API URL to get the pull request review comment where the secret was detected.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url"`
	Role string `json:"role"` // The role of the user in the team.
	State string `json:"state"` // The state of the user's membership in the team.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Type GeneratedType `json:"type,omitempty"` // The type of issue.
	Assignee GeneratedType `json:"assignee"` // A GitHub user.
	Pull_request map[string]interface{} `json:"pull_request,omitempty"`
	Events_url string `json:"events_url"`
	Assignees []GeneratedType `json:"assignees,omitempty"`
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Score float64 `json:"score"`
	Body_text string `json:"body_text,omitempty"`
	Labels_url string `json:"labels_url"`
	Text_matches []map[string]interface{} `json:"text_matches,omitempty"`
	Closed_at string `json:"closed_at"`
	Locked bool `json:"locked"`
	Repository_url string `json:"repository_url"`
	User GeneratedType `json:"user"` // A GitHub user.
	Created_at string `json:"created_at"`
	Html_url string `json:"html_url"`
	Url string `json:"url"`
	Draft bool `json:"draft,omitempty"`
	Timeline_url string `json:"timeline_url,omitempty"`
	Number int `json:"number"`
	Sub_issues_summary map[string]interface{} `json:"sub_issues_summary,omitempty"`
	Comments_url string `json:"comments_url"`
	State_reason string `json:"state_reason,omitempty"`
	Updated_at string `json:"updated_at"`
	Active_lock_reason string `json:"active_lock_reason,omitempty"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Reactions GeneratedType `json:"reactions,omitempty"`
	Id int64 `json:"id"`
	Body_html string `json:"body_html,omitempty"`
	Labels []map[string]interface{} `json:"labels"`
	Title string `json:"title"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Body string `json:"body,omitempty"`
	Milestone GeneratedType `json:"milestone"` // A collection of related issues and pull requests.
	Node_id string `json:"node_id"`
	State string `json:"state"`
	Comments int `json:"comments"`
}

// Webhookscomment represents the Webhookscomment schema from the OpenAPI specification
type Webhookscomment struct {
	Node_id string `json:"node_id"`
	Reactions map[string]interface{} `json:"reactions"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Html_url string `json:"html_url"`
	Body string `json:"body"`
	Child_comment_count int `json:"child_comment_count"`
	Discussion_id int `json:"discussion_id"`
	Id int `json:"id"`
	Parent_id int `json:"parent_id"`
	Repository_url string `json:"repository_url"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
	User map[string]interface{} `json:"user"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Seat_management_setting string `json:"seat_management_setting"` // The mode of assigning new seats.
	Cli string `json:"cli,omitempty"` // The organization policy for allowing or disallowing Copilot in the CLI.
	Ide_chat string `json:"ide_chat,omitempty"` // The organization policy for allowing or disallowing Copilot Chat in the IDE.
	Plan_type string `json:"plan_type,omitempty"` // The Copilot plan of the organization, or the parent enterprise, when applicable.
	Platform_chat string `json:"platform_chat,omitempty"` // The organization policy for allowing or disallowing Copilot features on GitHub.com.
	Public_code_suggestions string `json:"public_code_suggestions"` // The organization policy for allowing or blocking suggestions matching public code (duplication detection filter).
	Seat_breakdown GeneratedType `json:"seat_breakdown"` // The breakdown of Copilot Business seats for the organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue Webhooksissue2 `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	User_view_type string `json:"user_view_type,omitempty"`
	Login string `json:"login"`
	Site_admin bool `json:"site_admin"`
	Url string `json:"url"`
	Avatar_url string `json:"avatar_url"`
	Followers_url string `json:"followers_url"`
	Type string `json:"type"`
	Repos_url string `json:"repos_url"`
	Email string `json:"email,omitempty"`
	Subscriptions_url string `json:"subscriptions_url"`
	Following_url string `json:"following_url"`
	Name string `json:"name,omitempty"`
	Starred_url string `json:"starred_url"`
	Gists_url string `json:"gists_url"`
	Id int64 `json:"id"`
	Node_id string `json:"node_id"`
	Gravatar_id string `json:"gravatar_id"`
	Organizations_url string `json:"organizations_url"`
	Received_events_url string `json:"received_events_url"`
	Html_url string `json:"html_url"`
	Starred_at string `json:"starred_at,omitempty"`
	Events_url string `json:"events_url"`
}

// Webhooksissuecomment represents the Webhooksissuecomment schema from the OpenAPI specification
type Webhooksissuecomment struct {
	Url string `json:"url"` // URL for the issue comment
	Html_url string `json:"html_url"`
	Id int64 `json:"id"` // Unique identifier of the issue comment
	Issue_url string `json:"issue_url"`
	User map[string]interface{} `json:"user"`
	Created_at string `json:"created_at"`
	Reactions map[string]interface{} `json:"reactions"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Body string `json:"body"` // Contents of the issue comment
	Node_id string `json:"node_id"`
	Performed_via_github_app Integration `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Updated_at string `json:"updated_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Sponsorship Webhookssponsorship `json:"sponsorship"`
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Message string `json:"message"` // Message describing the purpose of the tag
	Node_id string `json:"node_id"`
	Object map[string]interface{} `json:"object"`
	Sha string `json:"sha"`
	Tag string `json:"tag"` // Name of the tag
	Tagger map[string]interface{} `json:"tagger"`
	Url string `json:"url"` // URL for the tag
	Verification Verification `json:"verification,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Cpus int `json:"cpus"` // How many cores are available to the codespace.
	Display_name string `json:"display_name"` // The display name of the machine includes cores, memory, and storage.
	Memory_in_bytes int `json:"memory_in_bytes"` // How much memory is available to the codespace.
	Name string `json:"name"` // The name of the machine.
	Operating_system string `json:"operating_system"` // The operating system of the machine.
	Prebuild_availability string `json:"prebuild_availability"` // Whether a prebuild is currently available when creating a codespace for this machine and repository. If a branch was not specified as a ref, the default branch will be assumed. Value will be "null" if prebuilds are not supported or prebuild availability could not be determined. Value will be "none" if no prebuild is available. Latest values "ready" and "in_progress" indicate the prebuild availability status.
	Storage_in_bytes int `json:"storage_in_bytes"` // How much storage is available to the codespace.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sha string `json:"sha"`
	Tree []map[string]interface{} `json:"tree"` // Objects specifying a tree structure
	Truncated bool `json:"truncated"`
	Url string `json:"url,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Merge_method string `json:"merge_method"` // The merge method to use.
	Commit_message string `json:"commit_message"` // Commit message for the merge commit.
	Commit_title string `json:"commit_title"` // Title for the merge commit message.
	Enabled_by GeneratedType `json:"enabled_by"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Thread map[string]interface{} `json:"thread"`
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Team Webhooksteam1 `json:"team"` // Groups of organization members that gives permissions on specified repositories.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository map[string]interface{} `json:"repository,omitempty"` // A git repository
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Repository_advisory GeneratedType `json:"repository_advisory"` // A repository security advisory.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Commit_oid string `json:"commit_oid"` // The commit SHA of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Ref string `json:"ref"` // The Git reference of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Alert map[string]interface{} `json:"alert"` // The code scanning alert involved in the event.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Type string `json:"type"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Start_line float64 `json:"start_line"` // Line number at which the secret starts in the file
	End_line float64 `json:"end_line"` // Line number at which the secret ends in the file
	End_column float64 `json:"end_column"` // The column at which the secret ends within the end line when the file is interpreted as 8-bit ASCII.
	Page_url string `json:"page_url"` // The GitHub URL to get the associated wiki page
	Blob_sha string `json:"blob_sha"` // SHA-1 hash ID of the associated blob
	Commit_url string `json:"commit_url"` // The GitHub URL to get the associated wiki commit
	Path string `json:"path"` // The file path of the wiki page
	Start_column float64 `json:"start_column"` // The column at which the secret starts within the start line when the file is interpreted as 8-bit ASCII.
	Commit_sha string `json:"commit_sha"` // SHA-1 hash ID of the associated commit
}

// Job represents the Job schema from the OpenAPI specification
type Job struct {
	Run_id int `json:"run_id"` // The id of the associated workflow run.
	Steps []map[string]interface{} `json:"steps,omitempty"` // Steps in this job.
	Head_branch string `json:"head_branch"` // The name of the current branch.
	Node_id string `json:"node_id"`
	Url string `json:"url"`
	Head_sha string `json:"head_sha"` // The SHA of the commit that is being run.
	Workflow_name string `json:"workflow_name"` // The name of the workflow.
	Check_run_url string `json:"check_run_url"`
	Started_at string `json:"started_at"` // The time that the job started, in ISO 8601 format.
	Completed_at string `json:"completed_at"` // The time that the job finished, in ISO 8601 format.
	Conclusion string `json:"conclusion"` // The outcome of the job.
	Created_at string `json:"created_at"` // The time that the job created, in ISO 8601 format.
	Html_url string `json:"html_url"`
	Runner_group_id int `json:"runner_group_id"` // The ID of the runner group to which this job has been assigned. (If a runner hasn't yet been assigned, this will be null.)
	Id int `json:"id"` // The id of the job.
	Labels []string `json:"labels"` // Labels for the workflow job. Specified by the "runs_on" attribute in the action's workflow file.
	Name string `json:"name"` // The name of the job.
	Run_attempt int `json:"run_attempt,omitempty"` // Attempt number of the associated workflow run, 1 for first attempt and higher if the workflow was re-run.
	Runner_id int `json:"runner_id"` // The ID of the runner to which this job has been assigned. (If a runner hasn't yet been assigned, this will be null.)
	Runner_name string `json:"runner_name"` // The name of the runner to which this job has been assigned. (If a runner hasn't yet been assigned, this will be null.)
	Run_url string `json:"run_url"`
	Runner_group_name string `json:"runner_group_name"` // The name of the runner group to which this job has been assigned. (If a runner hasn't yet been assigned, this will be null.)
	Status string `json:"status"` // The phase of the lifecycle that the job is currently in.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	New_property_values []GeneratedType `json:"new_property_values"` // The new custom property values for the repository.
	Old_property_values []GeneratedType `json:"old_property_values"` // The old custom property values for the repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Forks int `json:"forks"`
	Fork bool `json:"fork"`
	Name string `json:"name"`
	Disabled bool `json:"disabled"` // Returns whether or not this repository disabled.
	Assignees_url string `json:"assignees_url"`
	Default_branch string `json:"default_branch"`
	Html_url string `json:"html_url"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub user.
	Merges_url string `json:"merges_url"`
	Subscribers_url string `json:"subscribers_url"`
	Deployments_url string `json:"deployments_url"`
	Subscription_url string `json:"subscription_url"`
	Commits_url string `json:"commits_url"`
	Open_issues_count int `json:"open_issues_count"`
	Private bool `json:"private"`
	Allow_update_branch bool `json:"allow_update_branch,omitempty"`
	Branches_url string `json:"branches_url"`
	Watchers_count int `json:"watchers_count"`
	Created_at string `json:"created_at"`
	Pulls_url string `json:"pulls_url"`
	Allow_squash_merge bool `json:"allow_squash_merge,omitempty"`
	Releases_url string `json:"releases_url"`
	Archive_url string `json:"archive_url"`
	Url string `json:"url"`
	License GeneratedType `json:"license"` // License Simple
	Milestones_url string `json:"milestones_url"`
	Updated_at string `json:"updated_at"`
	Visibility string `json:"visibility,omitempty"` // The repository visibility: public, private, or internal.
	Merge_commit_title string `json:"merge_commit_title,omitempty"` // The default value for a merge commit title. - `PR_TITLE` - default to the pull request's title. - `MERGE_MESSAGE` - default to the classic title for a merge message (e.g., Merge pull request #123 from branch-name).
	Comments_url string `json:"comments_url"`
	Trees_url string `json:"trees_url"`
	Watchers int `json:"watchers"`
	Description string `json:"description"`
	Clone_url string `json:"clone_url"`
	Squash_merge_commit_message string `json:"squash_merge_commit_message,omitempty"` // The default value for a squash merge commit message: - `PR_BODY` - default to the pull request's body. - `COMMIT_MESSAGES` - default to the branch's commit messages. - `BLANK` - default to a blank commit message.
	Compare_url string `json:"compare_url"`
	Has_issues bool `json:"has_issues"`
	Forks_url string `json:"forks_url"`
	Collaborators_url string `json:"collaborators_url"`
	Master_branch string `json:"master_branch,omitempty"`
	Git_refs_url string `json:"git_refs_url"`
	Has_downloads bool `json:"has_downloads,omitempty"`
	Contributors_url string `json:"contributors_url"`
	Size int `json:"size"` // The size of the repository, in kilobytes. Size is calculated hourly. When a repository is initially created, the size is 0.
	Blobs_url string `json:"blobs_url"`
	Allow_forking bool `json:"allow_forking,omitempty"`
	Keys_url string `json:"keys_url"`
	Issues_url string `json:"issues_url"`
	Language string `json:"language"`
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"`
	Tags_url string `json:"tags_url"`
	Squash_merge_commit_title string `json:"squash_merge_commit_title,omitempty"` // The default value for a squash merge commit title: - `PR_TITLE` - default to the pull request's title. - `COMMIT_OR_PR_TITLE` - default to the commit's title (if only one commit) or the pull request's title (when more than one commit).
	Issue_comment_url string `json:"issue_comment_url"`
	Homepage string `json:"homepage"`
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Open_issues int `json:"open_issues"`
	Source Repository `json:"source,omitempty"` // A repository on GitHub.
	Has_projects bool `json:"has_projects"`
	Subscribers_count int `json:"subscribers_count"`
	Allow_merge_commit bool `json:"allow_merge_commit,omitempty"`
	Node_id string `json:"node_id"`
	Allow_auto_merge bool `json:"allow_auto_merge,omitempty"`
	Is_template bool `json:"is_template,omitempty"`
	Git_commits_url string `json:"git_commits_url"`
	Stargazers_url string `json:"stargazers_url"`
	Forks_count int `json:"forks_count"`
	Network_count int `json:"network_count"`
	Archived bool `json:"archived"`
	Security_and_analysis GeneratedType `json:"security_and_analysis,omitempty"`
	Custom_properties map[string]interface{} `json:"custom_properties,omitempty"` // The custom properties that were defined for the repository. The keys are the custom property names, and the values are the corresponding custom property values.
	Parent Repository `json:"parent,omitempty"` // A repository on GitHub.
	Notifications_url string `json:"notifications_url"`
	Git_url string `json:"git_url"`
	Statuses_url string `json:"statuses_url"`
	Teams_url string `json:"teams_url"`
	Allow_rebase_merge bool `json:"allow_rebase_merge,omitempty"`
	Merge_commit_message string `json:"merge_commit_message,omitempty"` // The default value for a merge commit message. - `PR_TITLE` - default to the pull request's title. - `PR_BODY` - default to the pull request's body. - `BLANK` - default to a blank commit message.
	Code_of_conduct GeneratedType `json:"code_of_conduct,omitempty"` // Code of Conduct Simple
	Has_wiki bool `json:"has_wiki"`
	Template_repository GeneratedType `json:"template_repository,omitempty"` // A repository on GitHub.
	Id int64 `json:"id"`
	Issue_events_url string `json:"issue_events_url"`
	Git_tags_url string `json:"git_tags_url"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Temp_clone_token string `json:"temp_clone_token,omitempty"`
	Downloads_url string `json:"downloads_url"`
	Ssh_url string `json:"ssh_url"`
	Labels_url string `json:"labels_url"`
	Mirror_url string `json:"mirror_url"`
	Anonymous_access_enabled bool `json:"anonymous_access_enabled,omitempty"` // Whether anonymous git access is allowed.
	Events_url string `json:"events_url"`
	Languages_url string `json:"languages_url"`
	Contents_url string `json:"contents_url"`
	Pushed_at string `json:"pushed_at"`
	Full_name string `json:"full_name"`
	Use_squash_pr_title_as_default bool `json:"use_squash_pr_title_as_default,omitempty"`
	Svn_url string `json:"svn_url"`
	Has_pages bool `json:"has_pages"`
	Stargazers_count int `json:"stargazers_count"`
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"`
	Hooks_url string `json:"hooks_url"`
	Has_discussions bool `json:"has_discussions"`
	Topics []string `json:"topics,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Error map[string]interface{} `json:"error"`
	Pusher GeneratedType `json:"pusher"` // A GitHub user.
	Status string `json:"status"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	Commit string `json:"commit"`
	Created_at string `json:"created_at"`
	Duration int `json:"duration"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Sponsorship Webhookssponsorship `json:"sponsorship"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Inputs map[string]interface{} `json:"inputs"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Ref string `json:"ref"`
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Workflow string `json:"workflow"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Issues string `json:"issues,omitempty"` // The level of permission to grant the access token for issues and related comments, assignees, labels, and milestones.
	Organization_hooks string `json:"organization_hooks,omitempty"` // The level of permission to grant the access token to manage the post-receive hooks for an organization.
	Repository_custom_properties string `json:"repository_custom_properties,omitempty"` // The level of permission to grant the access token to view and edit custom properties for a repository, when allowed by the property.
	Pages string `json:"pages,omitempty"` // The level of permission to grant the access token to retrieve Pages statuses, configuration, and builds, as well as create new builds.
	Organization_user_blocking string `json:"organization_user_blocking,omitempty"` // The level of permission to grant the access token to view and manage users blocked by the organization.
	Organization_personal_access_tokens string `json:"organization_personal_access_tokens,omitempty"` // The level of permission to grant the access token for viewing and managing fine-grained personal access token requests to an organization.
	Security_events string `json:"security_events,omitempty"` // The level of permission to grant the access token to view and manage security events like code scanning alerts.
	Repository_projects string `json:"repository_projects,omitempty"` // The level of permission to grant the access token to manage repository projects, columns, and cards.
	Organization_events string `json:"organization_events,omitempty"` // The level of permission to grant the access token to view events triggered by an activity in an organization.
	Organization_self_hosted_runners string `json:"organization_self_hosted_runners,omitempty"` // The level of permission to grant the access token to view and manage GitHub Actions self-hosted runners available to an organization.
	Single_file string `json:"single_file,omitempty"` // The level of permission to grant the access token to manage just a single file.
	Secret_scanning_alerts string `json:"secret_scanning_alerts,omitempty"` // The level of permission to grant the access token to view and manage secret scanning alerts.
	Deployments string `json:"deployments,omitempty"` // The level of permission to grant the access token for deployments and deployment statuses.
	Secrets string `json:"secrets,omitempty"` // The level of permission to grant the access token to manage repository secrets.
	Organization_packages string `json:"organization_packages,omitempty"` // The level of permission to grant the access token for organization packages published to GitHub Packages.
	Git_ssh_keys string `json:"git_ssh_keys,omitempty"` // The level of permission to grant the access token to manage git SSH keys.
	Statuses string `json:"statuses,omitempty"` // The level of permission to grant the access token for commit statuses.
	Organization_secrets string `json:"organization_secrets,omitempty"` // The level of permission to grant the access token to manage organization secrets.
	Organization_announcement_banners string `json:"organization_announcement_banners,omitempty"` // The level of permission to grant the access token to view and manage announcement banners for an organization.
	Organization_custom_roles string `json:"organization_custom_roles,omitempty"` // The level of permission to grant the access token for custom repository roles management.
	Packages string `json:"packages,omitempty"` // The level of permission to grant the access token for packages published to GitHub Packages.
	Pull_requests string `json:"pull_requests,omitempty"` // The level of permission to grant the access token for pull requests and related comments, assignees, labels, milestones, and merges.
	Organization_administration string `json:"organization_administration,omitempty"` // The level of permission to grant the access token to manage access to an organization.
	Organization_projects string `json:"organization_projects,omitempty"` // The level of permission to grant the access token to manage organization projects and projects public preview (where available).
	Codespaces string `json:"codespaces,omitempty"` // The level of permission to grant the access token to create, edit, delete, and list Codespaces.
	Repository_hooks string `json:"repository_hooks,omitempty"` // The level of permission to grant the access token to manage the post-receive hooks for a repository.
	Dependabot_secrets string `json:"dependabot_secrets,omitempty"` // The level of permission to grant the access token to manage Dependabot secrets.
	Metadata string `json:"metadata,omitempty"` // The level of permission to grant the access token to search repositories, list collaborators, and access repository metadata.
	Team_discussions string `json:"team_discussions,omitempty"` // The level of permission to grant the access token to manage team discussions and related comments.
	Actions string `json:"actions,omitempty"` // The level of permission to grant the access token for GitHub Actions workflows, workflow runs, and artifacts.
	Profile string `json:"profile,omitempty"` // The level of permission to grant the access token to manage the profile settings belonging to a user.
	Starring string `json:"starring,omitempty"` // The level of permission to grant the access token to list and manage repositories a user is starring.
	Interaction_limits string `json:"interaction_limits,omitempty"` // The level of permission to grant the access token to view and manage interaction limits on a repository.
	Organization_plan string `json:"organization_plan,omitempty"` // The level of permission to grant the access token for viewing an organization's plan.
	Members string `json:"members,omitempty"` // The level of permission to grant the access token for organization teams and members.
	Vulnerability_alerts string `json:"vulnerability_alerts,omitempty"` // The level of permission to grant the access token to manage Dependabot alerts.
	Checks string `json:"checks,omitempty"` // The level of permission to grant the access token for checks on code.
	Organization_custom_org_roles string `json:"organization_custom_org_roles,omitempty"` // The level of permission to grant the access token for custom organization roles management.
	Followers string `json:"followers,omitempty"` // The level of permission to grant the access token to manage the followers belonging to a user.
	Gpg_keys string `json:"gpg_keys,omitempty"` // The level of permission to grant the access token to view and manage GPG keys belonging to a user.
	Organization_personal_access_token_requests string `json:"organization_personal_access_token_requests,omitempty"` // The level of permission to grant the access token for viewing and managing fine-grained personal access tokens that have been approved by an organization.
	Organization_copilot_seat_management string `json:"organization_copilot_seat_management,omitempty"` // The level of permission to grant the access token for managing access to GitHub Copilot for members of an organization with a Copilot Business subscription. This property is in public preview and is subject to change.
	Workflows string `json:"workflows,omitempty"` // The level of permission to grant the access token to update GitHub Actions workflow files.
	Email_addresses string `json:"email_addresses,omitempty"` // The level of permission to grant the access token to manage the email addresses belonging to a user.
	Contents string `json:"contents,omitempty"` // The level of permission to grant the access token for repository contents, commits, branches, downloads, releases, and merges.
	Organization_custom_properties string `json:"organization_custom_properties,omitempty"` // The level of permission to grant the access token for custom property management.
	Environments string `json:"environments,omitempty"` // The level of permission to grant the access token for managing repository environments.
	Administration string `json:"administration,omitempty"` // The level of permission to grant the access token for repository creation, deletion, settings, teams, and collaborators creation.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Review Webhooksreview `json:"review"` // The review that was affected.
	Changes map[string]interface{} `json:"changes"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request map[string]interface{} `json:"pull_request"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Repository_ruleset GeneratedType `json:"repository_ruleset"` // A set of rules to apply when specified conditions are met.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// Issue represents the Issue schema from the OpenAPI specification
type Issue struct {
	Closed_at string `json:"closed_at"`
	Timeline_url string `json:"timeline_url,omitempty"`
	Updated_at string `json:"updated_at"`
	Comments int `json:"comments"`
	Draft bool `json:"draft,omitempty"`
	Body_text string `json:"body_text,omitempty"`
	Node_id string `json:"node_id"`
	User GeneratedType `json:"user"` // A GitHub user.
	Repository_url string `json:"repository_url"`
	Labels_url string `json:"labels_url"`
	Milestone GeneratedType `json:"milestone"` // A collection of related issues and pull requests.
	Labels []interface{} `json:"labels"` // Labels to associate with this issue; pass one or more label names to replace the set of labels on this issue; send an empty array to clear all labels from the issue; note that the labels are silently dropped for users without push access to the repository
	Pull_request map[string]interface{} `json:"pull_request,omitempty"`
	State string `json:"state"` // State of the issue; either 'open' or 'closed'
	Performed_via_github_app GeneratedType `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Reactions GeneratedType `json:"reactions,omitempty"`
	Number int `json:"number"` // Number uniquely identifying the issue within its repository
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Assignee GeneratedType `json:"assignee"` // A GitHub user.
	State_reason string `json:"state_reason,omitempty"` // The reason for the current state
	Events_url string `json:"events_url"`
	Title string `json:"title"` // Title of the issue
	Sub_issues_summary GeneratedType `json:"sub_issues_summary,omitempty"`
	Id int64 `json:"id"`
	Url string `json:"url"` // URL for the issue
	Body_html string `json:"body_html,omitempty"`
	Active_lock_reason string `json:"active_lock_reason,omitempty"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Html_url string `json:"html_url"`
	Body string `json:"body,omitempty"` // Contents of the issue
	Closed_by GeneratedType `json:"closed_by,omitempty"` // A GitHub user.
	Comments_url string `json:"comments_url"`
	Created_at string `json:"created_at"`
	Type GeneratedType `json:"type,omitempty"` // The type of issue.
	Assignees []GeneratedType `json:"assignees,omitempty"`
	Locked bool `json:"locked"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Status string `json:"status,omitempty"` // The current status of the deployment.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Projects_v2 GeneratedType `json:"projects_v2"` // A projects v2 project
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
}

// Webhooksdeploykey represents the Webhooksdeploykey schema from the OpenAPI specification
type Webhooksdeploykey struct {
	Id int `json:"id"`
	Url string `json:"url"`
	Created_at string `json:"created_at"`
	Enabled bool `json:"enabled,omitempty"`
	Last_used string `json:"last_used,omitempty"`
	Read_only bool `json:"read_only"`
	Title string `json:"title"`
	Verified bool `json:"verified"`
	Key string `json:"key"`
	Added_by string `json:"added_by,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Html_url string `json:"html_url"`
	Id int `json:"id"`
	Login string `json:"login"`
	Name string `json:"name"`
	Node_id string `json:"node_id"`
	Avatar_url string `json:"avatar_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"` // The changes to the milestone if the action was `edited`.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Milestone Webhooksmilestone `json:"milestone"` // A collection of related issues and pull requests.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Repositories_added []map[string]interface{} `json:"repositories_added"` // An array of repository objects, which were added to the installation.
	Repository_selection string `json:"repository_selection"` // Describe whether all repositories have been selected or there's a selection involved
	Repositories_removed []map[string]interface{} `json:"repositories_removed"` // An array of repository objects, which were removed from the installation.
	Requester Webhooksuser `json:"requester"`
	Action string `json:"action"`
	Installation Installation `json:"installation"` // Installation
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Type string `json:"type"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Member Webhooksuser `json:"member"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Forks_url string `json:"forks_url,omitempty"`
	Node_id string `json:"node_id,omitempty"`
	User string `json:"user,omitempty"`
	Created_at string `json:"created_at,omitempty"`
	Truncated bool `json:"truncated,omitempty"`
	History []GeneratedType `json:"history,omitempty"`
	Id string `json:"id,omitempty"`
	Comments int `json:"comments,omitempty"`
	Public bool `json:"public,omitempty"`
	Url string `json:"url,omitempty"`
	Comments_enabled bool `json:"comments_enabled,omitempty"`
	Git_pull_url string `json:"git_pull_url,omitempty"`
	Owner GeneratedType `json:"owner,omitempty"` // A GitHub user.
	Files map[string]interface{} `json:"files,omitempty"`
	Fork_of map[string]interface{} `json:"fork_of,omitempty"` // Gist
	Comments_url string `json:"comments_url,omitempty"`
	Html_url string `json:"html_url,omitempty"`
	Git_push_url string `json:"git_push_url,omitempty"`
	Updated_at string `json:"updated_at,omitempty"`
	Commits_url string `json:"commits_url,omitempty"`
	Description string `json:"description,omitempty"`
	Forks []map[string]interface{} `json:"forks,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Type string `json:"type"`
	Parameters map[string]interface{} `json:"parameters,omitempty"`
}

// Webhooksrule represents the Webhooksrule schema from the OpenAPI specification
type Webhooksrule struct {
	Id int `json:"id"`
	Lock_allows_fork_sync bool `json:"lock_allows_fork_sync,omitempty"` // Whether users can pull changes from upstream when the branch is locked. Set to `true` to allow users to pull changes from upstream when the branch is locked. This setting is only applicable for forks.
	Required_deployments_enforcement_level string `json:"required_deployments_enforcement_level"`
	Authorized_actor_names []string `json:"authorized_actor_names"`
	Allow_deletions_enforcement_level string `json:"allow_deletions_enforcement_level"`
	Pull_request_reviews_enforcement_level string `json:"pull_request_reviews_enforcement_level"`
	Required_approving_review_count int `json:"required_approving_review_count"`
	Dismiss_stale_reviews_on_push bool `json:"dismiss_stale_reviews_on_push"`
	Admin_enforced bool `json:"admin_enforced"`
	Name string `json:"name"`
	Require_code_owner_review bool `json:"require_code_owner_review"`
	Required_conversation_resolution_level string `json:"required_conversation_resolution_level"`
	Authorized_dismissal_actors_only bool `json:"authorized_dismissal_actors_only"`
	Required_status_checks_enforcement_level string `json:"required_status_checks_enforcement_level"`
	Strict_required_status_checks_policy bool `json:"strict_required_status_checks_policy"`
	Linear_history_requirement_enforcement_level string `json:"linear_history_requirement_enforcement_level"`
	Signature_requirement_enforcement_level string `json:"signature_requirement_enforcement_level"`
	Lock_branch_enforcement_level string `json:"lock_branch_enforcement_level"` // The enforcement level of the branch lock setting. `off` means the branch is not locked, `non_admins` means the branch is read-only for non_admins, and `everyone` means the branch is read-only for everyone.
	Repository_id int `json:"repository_id"`
	Updated_at string `json:"updated_at"`
	Required_status_checks []string `json:"required_status_checks"`
	Require_last_push_approval bool `json:"require_last_push_approval,omitempty"` // Whether the most recent push must be approved by someone other than the person who pushed it
	Created_at string `json:"created_at"`
	Ignore_approvals_from_contributors bool `json:"ignore_approvals_from_contributors"`
	Allow_force_pushes_enforcement_level string `json:"allow_force_pushes_enforcement_level"`
	Create_protected bool `json:"create_protected,omitempty"`
	Merge_queue_enforcement_level string `json:"merge_queue_enforcement_level"`
	Authorized_actors_only bool `json:"authorized_actors_only"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Number int `json:"number"`
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Action string `json:"action"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Pull_request map[string]interface{} `json:"pull_request"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Reason string `json:"reason"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repos_url string `json:"repos_url"`
	Avatar_url string `json:"avatar_url"`
	Description string `json:"description"`
	Events_url string `json:"events_url"`
	Id int `json:"id"`
	Members_url string `json:"members_url"`
	Issues_url string `json:"issues_url"`
	Url string `json:"url"`
	Login string `json:"login"`
	Node_id string `json:"node_id"`
	Hooks_url string `json:"hooks_url"`
	Public_members_url string `json:"public_members_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Accepted bool `json:"accepted"` // Whether the user has accepted the permissions defined by the devcontainer config
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Version string `json:"version"`
	Change_status map[string]interface{} `json:"change_status"`
	Committed_at string `json:"committed_at"`
	Url string `json:"url"`
	User GeneratedType `json:"user"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Merge_group GeneratedType `json:"merge_group"` // A group of pull requests that the merge queue has grouped together to be merged.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Reason string `json:"reason,omitempty"` // Explains why the merge group is being destroyed. The group could have been merged, removed from the queue (dequeued), or invalidated by an earlier queue entry being dequeued (invalidated).
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
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

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Pull_request map[string]interface{} `json:"pull_request"`
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Number int `json:"number"`
	Reason string `json:"reason"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Action string `json:"action"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// Metadata represents the Metadata schema from the OpenAPI specification
type Metadata struct {
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Discussion_url string `json:"discussion_url"`
	Number int `json:"number"` // The unique sequence number of a team discussion comment.
	Url string `json:"url"`
	Author GeneratedType `json:"author"` // A GitHub user.
	Body string `json:"body"` // The main text of the comment.
	Last_edited_at string `json:"last_edited_at"`
	Updated_at string `json:"updated_at"`
	Html_url string `json:"html_url"`
	Reactions GeneratedType `json:"reactions,omitempty"`
	Body_html string `json:"body_html"`
	Body_version string `json:"body_version"` // The current version of the body content. If provided, this update operation will be rejected if the given version does not match the latest version on the server.
	Created_at string `json:"created_at"`
	Node_id string `json:"node_id"`
}

// Hook represents the Hook schema from the OpenAPI specification
type Hook struct {
	Type string `json:"type"`
	Deliveries_url string `json:"deliveries_url,omitempty"`
	Ping_url string `json:"ping_url"`
	Test_url string `json:"test_url"`
	Updated_at string `json:"updated_at"`
	Config GeneratedType `json:"config"` // Configuration object of the webhook
	Events []string `json:"events"` // Determines what events the hook is triggered for. Default: ['push'].
	Url string `json:"url"`
	Active bool `json:"active"` // Determines whether the hook is actually triggered on pushes.
	Last_response GeneratedType `json:"last_response"`
	Name string `json:"name"` // The name of a valid service, use 'web' for a webhook.
	Created_at string `json:"created_at"`
	Id int `json:"id"` // Unique identifier of the webhook.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request GeneratedType `json:"pull_request"`
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Type string `json:"type"`
	Parameters map[string]interface{} `json:"parameters,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Changes map[string]interface{} `json:"changes"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// Codespace represents the Codespace schema from the OpenAPI specification
type Codespace struct {
	Prebuild bool `json:"prebuild"` // Whether the codespace was created from a prebuild.
	Updated_at string `json:"updated_at"`
	Web_url string `json:"web_url"` // URL to access this codespace on the web.
	State string `json:"state"` // State of this codespace.
	Environment_id string `json:"environment_id"` // UUID identifying this codespace's environment.
	Pulls_url string `json:"pulls_url"` // API URL for the Pull Request associated with this codespace, if any.
	Idle_timeout_notice string `json:"idle_timeout_notice,omitempty"` // Text to show user when codespace idle timeout minutes has been overriden by an organization policy
	Recent_folders []string `json:"recent_folders"`
	Url string `json:"url"` // API URL for this codespace.
	Location string `json:"location"` // The initally assigned location of a new codespace.
	Start_url string `json:"start_url"` // API URL to start this codespace.
	Pending_operation bool `json:"pending_operation,omitempty"` // Whether or not a codespace has a pending async operation. This would mean that the codespace is temporarily unavailable. The only thing that you can do with a codespace in this state is delete it.
	Retention_period_minutes int `json:"retention_period_minutes,omitempty"` // Duration in minutes after codespace has gone idle in which it will be deleted. Must be integer minutes between 0 and 43200 (30 days).
	Pending_operation_disabled_reason string `json:"pending_operation_disabled_reason,omitempty"` // Text to show user when codespace is disabled by a pending operation
	Repository GeneratedType `json:"repository"` // Minimal Repository
	Devcontainer_path string `json:"devcontainer_path,omitempty"` // Path to devcontainer.json from repo root used to create Codespace.
	Git_status map[string]interface{} `json:"git_status"` // Details about the codespace's git repository.
	Name string `json:"name"` // Automatically generated name of this codespace.
	Idle_timeout_minutes int `json:"idle_timeout_minutes"` // The number of minutes of inactivity after which this codespace will be automatically stopped.
	Created_at string `json:"created_at"`
	Billable_owner GeneratedType `json:"billable_owner"` // A GitHub user.
	Machine GeneratedType `json:"machine"` // A description of the machine powering a codespace.
	Display_name string `json:"display_name,omitempty"` // Display name for this codespace.
	Machines_url string `json:"machines_url"` // API URL to access available alternate machine types for this codespace.
	Last_used_at string `json:"last_used_at"` // Last known time this codespace was started.
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Stop_url string `json:"stop_url"` // API URL to stop this codespace.
	Id int64 `json:"id"`
	Retention_expires_at string `json:"retention_expires_at,omitempty"` // When a codespace will be auto-deleted based on the "retention_period_minutes" and "last_used_at"
	Publish_url string `json:"publish_url,omitempty"` // API URL to publish this codespace to a new repository.
	Runtime_constraints map[string]interface{} `json:"runtime_constraints,omitempty"`
	Last_known_stop_notice string `json:"last_known_stop_notice,omitempty"` // The text to display to a user when a codespace has been stopped for a potentially actionable reason.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Github_owned_allowed bool `json:"github_owned_allowed,omitempty"` // Whether GitHub-owned actions are allowed. For example, this includes the actions in the `actions` organization.
	Patterns_allowed []string `json:"patterns_allowed,omitempty"` // Specifies a list of string-matching patterns to allow specific action(s) and reusable workflow(s). Wildcards, tags, and SHAs are allowed. For example, `monalisa/octocat@*`, `monalisa/octocat@v2`, `monalisa/*`. > [!NOTE] > The `patterns_allowed` setting only applies to public repositories.
	Verified_allowed bool `json:"verified_allowed,omitempty"` // Whether actions from GitHub Marketplace verified creators are allowed. Set to `true` to allow all actions by GitHub Marketplace verified creators.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url"`
	Html_url string `json:"html_url"`
	Id int64 `json:"id"`
	Name string `json:"name"`
	Created_at string `json:"created_at"`
	Group_name string `json:"group_name,omitempty"`
	Updated_at string `json:"updated_at"`
	Description string `json:"description,omitempty"`
	Members_url string `json:"members_url"`
	Organization_selection_type string `json:"organization_selection_type,omitempty"`
	Slug string `json:"slug"`
	Sync_to_organizations string `json:"sync_to_organizations,omitempty"`
	Group_id string `json:"group_id,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Workflow Webhooksworkflow `json:"workflow"`
	Workflow_run map[string]interface{} `json:"workflow_run"`
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Release Webhooksrelease `json:"release"` // The [release](https://docs.github.com/rest/releases/releases/#get-a-release) object.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id string `json:"id,omitempty"` // An identifier for the upload.
	Url string `json:"url,omitempty"` // The REST API URL for checking the status of the upload.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Documentation_url string `json:"documentation_url"`
	Errors []map[string]interface{} `json:"errors,omitempty"`
	Message string `json:"message"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Created_at string `json:"created_at"`
	Description string `json:"description"`
	Node_id string `json:"node_id"`
	State string `json:"state"`
	Id int `json:"id"`
	Updated_at string `json:"updated_at"`
	Avatar_url string `json:"avatar_url"`
	Context string `json:"context"`
	Required bool `json:"required,omitempty"`
	Target_url string `json:"target_url"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Node_id string `json:"node_id,omitempty"`
	Repository GeneratedType `json:"repository,omitempty"` // Minimal Repository
	Pull_requests []GeneratedType `json:"pull_requests,omitempty"`
	Before string `json:"before,omitempty"`
	Url string `json:"url,omitempty"`
	After string `json:"after,omitempty"`
	Conclusion string `json:"conclusion,omitempty"`
	Created_at string `json:"created_at,omitempty"`
	Head_sha string `json:"head_sha,omitempty"` // The SHA of the head commit that is being checked.
	Id int `json:"id,omitempty"`
	Status string `json:"status,omitempty"`
	Updated_at string `json:"updated_at,omitempty"`
	App Integration `json:"app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Head_branch string `json:"head_branch,omitempty"`
}

// Stargazer represents the Stargazer schema from the OpenAPI specification
type Stargazer struct {
	Starred_at string `json:"starred_at"`
	User GeneratedType `json:"user"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Alert GeneratedType `json:"alert"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Label Webhookslabel `json:"label,omitempty"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Number int `json:"number"` // The pull request number.
	Pull_request map[string]interface{} `json:"pull_request"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// Page represents the Page schema from the OpenAPI specification
type Page struct {
	Build_type string `json:"build_type,omitempty"` // The process in which the Page will be built.
	Protected_domain_state string `json:"protected_domain_state,omitempty"` // The state if the domain is verified
	Status string `json:"status"` // The status of the most recent build of the Page.
	Cname string `json:"cname"` // The Pages site's custom domain
	Public bool `json:"public"` // Whether the GitHub Pages site is publicly visible. If set to `true`, the site is accessible to anyone on the internet. If set to `false`, the site will only be accessible to users who have at least `read` access to the repository that published the site.
	Source GeneratedType `json:"source,omitempty"`
	Url string `json:"url"` // The API address for accessing this Page resource.
	Custom_404 bool `json:"custom_404"` // Whether the Page has a custom 404 page.
	Html_url string `json:"html_url,omitempty"` // The web address the Page can be accessed from.
	Https_certificate GeneratedType `json:"https_certificate,omitempty"`
	Https_enforced bool `json:"https_enforced,omitempty"` // Whether https is enabled on the domain
	Pending_domain_unverified_at string `json:"pending_domain_unverified_at,omitempty"` // The timestamp when a pending domain becomes unverified.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Selected_actions_url string `json:"selected_actions_url,omitempty"` // The API URL to use to get or set the actions and reusable workflows that are allowed to run, when `allowed_actions` is set to `selected`.
	Selected_repositories_url string `json:"selected_repositories_url,omitempty"` // The API URL to use to get or set the selected repositories that are allowed to run GitHub Actions, when `enabled_repositories` is set to `selected`.
	Allowed_actions string `json:"allowed_actions,omitempty"` // The permissions policy that controls the actions and reusable workflows that are allowed to run.
	Enabled_repositories string `json:"enabled_repositories"` // The policy that controls the repositories in the organization that are allowed to run GitHub Actions.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Changes map[string]interface{} `json:"changes"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Comment Webhookscomment `json:"comment"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Action string `json:"action"`
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Number int `json:"number"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Commit_url string `json:"commit_url"`
	Project_card map[string]interface{} `json:"project_card,omitempty"`
	Created_at string `json:"created_at"`
	Event string `json:"event"`
	Node_id string `json:"node_id"`
	Commit_id string `json:"commit_id"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Id int `json:"id"`
	Url string `json:"url"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Network_count int `json:"network_count,omitempty"`
	Contributors_url string `json:"contributors_url"`
	Has_discussions bool `json:"has_discussions,omitempty"`
	Role_name string `json:"role_name,omitempty"`
	Html_url string `json:"html_url"`
	Has_downloads bool `json:"has_downloads,omitempty"`
	Watchers_count int `json:"watchers_count,omitempty"`
	Size int `json:"size,omitempty"` // The size of the repository, in kilobytes. Size is calculated hourly. When a repository is initially created, the size is 0.
	Description string `json:"description"`
	Git_tags_url string `json:"git_tags_url"`
	Updated_at string `json:"updated_at,omitempty"`
	Issues_url string `json:"issues_url"`
	Full_name string `json:"full_name"`
	License map[string]interface{} `json:"license,omitempty"`
	Clone_url string `json:"clone_url,omitempty"`
	Has_projects bool `json:"has_projects,omitempty"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Collaborators_url string `json:"collaborators_url"`
	Compare_url string `json:"compare_url"`
	Custom_properties map[string]interface{} `json:"custom_properties,omitempty"` // The custom properties that were defined for the repository. The keys are the custom property names, and the values are the corresponding custom property values.
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"`
	Notifications_url string `json:"notifications_url"`
	Stargazers_count int `json:"stargazers_count,omitempty"`
	Is_template bool `json:"is_template,omitempty"`
	Fork bool `json:"fork"`
	Statuses_url string `json:"statuses_url"`
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"`
	Trees_url string `json:"trees_url"`
	Stargazers_url string `json:"stargazers_url"`
	Subscribers_url string `json:"subscribers_url"`
	Archive_url string `json:"archive_url"`
	Forks_url string `json:"forks_url"`
	Issue_comment_url string `json:"issue_comment_url"`
	Ssh_url string `json:"ssh_url,omitempty"`
	Git_commits_url string `json:"git_commits_url"`
	Watchers int `json:"watchers,omitempty"`
	Merges_url string `json:"merges_url"`
	Svn_url string `json:"svn_url,omitempty"`
	Assignees_url string `json:"assignees_url"`
	Deployments_url string `json:"deployments_url"`
	Default_branch string `json:"default_branch,omitempty"`
	Forks_count int `json:"forks_count,omitempty"`
	Labels_url string `json:"labels_url"`
	Milestones_url string `json:"milestones_url"`
	Releases_url string `json:"releases_url"`
	Events_url string `json:"events_url"`
	Id int64 `json:"id"`
	Created_at string `json:"created_at,omitempty"`
	Languages_url string `json:"languages_url"`
	Pushed_at string `json:"pushed_at,omitempty"`
	Archived bool `json:"archived,omitempty"`
	Blobs_url string `json:"blobs_url"`
	Has_issues bool `json:"has_issues,omitempty"`
	Mirror_url string `json:"mirror_url,omitempty"`
	Open_issues int `json:"open_issues,omitempty"`
	Open_issues_count int `json:"open_issues_count,omitempty"`
	Teams_url string `json:"teams_url"`
	Node_id string `json:"node_id"`
	Security_and_analysis GeneratedType `json:"security_and_analysis,omitempty"`
	Commits_url string `json:"commits_url"`
	Code_of_conduct GeneratedType `json:"code_of_conduct,omitempty"` // Code Of Conduct
	Has_wiki bool `json:"has_wiki,omitempty"`
	Language string `json:"language,omitempty"`
	Private bool `json:"private"`
	Subscribers_count int `json:"subscribers_count,omitempty"`
	Homepage string `json:"homepage,omitempty"`
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Git_url string `json:"git_url,omitempty"`
	Url string `json:"url"`
	Subscription_url string `json:"subscription_url"`
	Hooks_url string `json:"hooks_url"`
	Tags_url string `json:"tags_url"`
	Comments_url string `json:"comments_url"`
	Temp_clone_token string `json:"temp_clone_token,omitempty"`
	Allow_forking bool `json:"allow_forking,omitempty"`
	Name string `json:"name"`
	Visibility string `json:"visibility,omitempty"`
	Keys_url string `json:"keys_url"`
	Git_refs_url string `json:"git_refs_url"`
	Disabled bool `json:"disabled,omitempty"`
	Contents_url string `json:"contents_url"`
	Branches_url string `json:"branches_url"`
	Has_pages bool `json:"has_pages,omitempty"`
	Topics []string `json:"topics,omitempty"`
	Forks int `json:"forks,omitempty"`
	Downloads_url string `json:"downloads_url"`
	Issue_events_url string `json:"issue_events_url"`
	Pulls_url string `json:"pulls_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Comments []GeneratedType `json:"comments,omitempty"`
	Event string `json:"event,omitempty"`
	Node_id string `json:"node_id,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Reactions GeneratedType `json:"reactions,omitempty"`
	Body_html string `json:"body_html,omitempty"`
	Created_at string `json:"created_at"`
	Position int `json:"position,omitempty"` // The line index in the diff to which the comment applies. This field is closing down; use `line` instead.
	Original_line int `json:"original_line,omitempty"` // The line of the blob to which the comment applies. The last line of the range for a multi-line comment
	Updated_at string `json:"updated_at"`
	Start_line int `json:"start_line,omitempty"` // The first line of the range for a multi-line comment.
	_links map[string]interface{} `json:"_links"`
	Pull_request_review_id int64 `json:"pull_request_review_id"` // The ID of the pull request review to which the comment belongs.
	Start_side string `json:"start_side,omitempty"` // The side of the first line of the range for a multi-line comment.
	Subject_type string `json:"subject_type,omitempty"` // The level at which the comment is targeted, can be a diff line or a file.
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Url string `json:"url"` // URL for the pull request review comment
	Line int `json:"line,omitempty"` // The line of the blob to which the comment applies. The last line of the range for a multi-line comment
	Original_start_line int `json:"original_start_line,omitempty"` // The first line of the range for a multi-line comment.
	User GeneratedType `json:"user"` // A GitHub user.
	In_reply_to_id int `json:"in_reply_to_id,omitempty"` // The comment ID to reply to.
	Id int64 `json:"id"` // The ID of the pull request review comment.
	Path string `json:"path"` // The relative path of the file to which the comment applies.
	Diff_hunk string `json:"diff_hunk"` // The diff of the line that the comment refers to.
	Side string `json:"side,omitempty"` // The side of the diff to which the comment applies. The side of the last line of the range for a multi-line comment
	Body string `json:"body"` // The text of the comment.
	Commit_id string `json:"commit_id"` // The SHA of the commit to which the comment applies.
	Original_position int `json:"original_position,omitempty"` // The index of the original line in the diff to which the comment applies. This field is closing down; use `original_line` instead.
	Original_commit_id string `json:"original_commit_id"` // The SHA of the original commit to which the comment applies.
	Node_id string `json:"node_id"` // The node ID of the pull request review comment.
	Body_text string `json:"body_text,omitempty"`
	Pull_request_url string `json:"pull_request_url"` // URL for the pull request that the review comment belongs to.
	Html_url string `json:"html_url"` // HTML URL for the pull request review comment.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Protected bool `json:"protected"`
	Protection GeneratedType `json:"protection,omitempty"` // Branch Protection
	Protection_url string `json:"protection_url,omitempty"`
	Commit map[string]interface{} `json:"commit"`
	Name string `json:"name"`
}

// Manifest represents the Manifest schema from the OpenAPI specification
type Manifest struct {
	Resolved map[string]interface{} `json:"resolved,omitempty"` // A collection of resolved package dependencies.
	File map[string]interface{} `json:"file,omitempty"`
	Metadata Metadata `json:"metadata,omitempty"` // User-defined metadata to store domain-specific information limited to 8 keys with scalar values.
	Name string `json:"name"` // The name of the manifest.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Subscribers_url string `json:"subscribers_url"`
	Clone_url string `json:"clone_url,omitempty"`
	Compare_url string `json:"compare_url"`
	Is_template bool `json:"is_template,omitempty"`
	Git_commits_url string `json:"git_commits_url"`
	Open_issues_count int `json:"open_issues_count,omitempty"`
	Created_at string `json:"created_at,omitempty"`
	Visibility string `json:"visibility,omitempty"`
	Git_tags_url string `json:"git_tags_url"`
	Events_url string `json:"events_url"`
	Custom_properties map[string]interface{} `json:"custom_properties,omitempty"` // The custom properties that were defined for the repository. The keys are the custom property names, and the values are the corresponding custom property values.
	Description string `json:"description"`
	Trees_url string `json:"trees_url"`
	Languages_url string `json:"languages_url"`
	Open_issues int `json:"open_issues,omitempty"`
	Homepage string `json:"homepage,omitempty"`
	Milestones_url string `json:"milestones_url"`
	Fork bool `json:"fork"`
	Pushed_at string `json:"pushed_at,omitempty"`
	Security_and_analysis GeneratedType `json:"security_and_analysis,omitempty"`
	Name string `json:"name"`
	Branches_url string `json:"branches_url"`
	Archive_url string `json:"archive_url"`
	Url string `json:"url"`
	Merges_url string `json:"merges_url"`
	Labels_url string `json:"labels_url"`
	Has_pages bool `json:"has_pages,omitempty"`
	Node_id string `json:"node_id"`
	Downloads_url string `json:"downloads_url"`
	Hooks_url string `json:"hooks_url"`
	Git_url string `json:"git_url,omitempty"`
	Updated_at string `json:"updated_at,omitempty"`
	Has_projects bool `json:"has_projects,omitempty"`
	Subscription_url string `json:"subscription_url"`
	Mirror_url string `json:"mirror_url,omitempty"`
	Has_issues bool `json:"has_issues,omitempty"`
	Temp_clone_token string `json:"temp_clone_token,omitempty"`
	Notifications_url string `json:"notifications_url"`
	Svn_url string `json:"svn_url,omitempty"`
	Comments_url string `json:"comments_url"`
	License map[string]interface{} `json:"license,omitempty"`
	Allow_forking bool `json:"allow_forking,omitempty"`
	Keys_url string `json:"keys_url"`
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"`
	Default_branch string `json:"default_branch,omitempty"`
	Stargazers_url string `json:"stargazers_url"`
	Disabled bool `json:"disabled,omitempty"`
	Tags_url string `json:"tags_url"`
	Forks_count int `json:"forks_count,omitempty"`
	Has_discussions bool `json:"has_discussions,omitempty"`
	Topics []string `json:"topics,omitempty"`
	Html_url string `json:"html_url"`
	Issues_url string `json:"issues_url"`
	Collaborators_url string `json:"collaborators_url"`
	Git_refs_url string `json:"git_refs_url"`
	Subscribers_count int `json:"subscribers_count,omitempty"`
	Issue_events_url string `json:"issue_events_url"`
	Forks_url string `json:"forks_url"`
	Full_name string `json:"full_name"`
	Id int64 `json:"id"`
	Network_count int `json:"network_count,omitempty"`
	Contributors_url string `json:"contributors_url"`
	Has_downloads bool `json:"has_downloads,omitempty"`
	Teams_url string `json:"teams_url"`
	Has_wiki bool `json:"has_wiki,omitempty"`
	Blobs_url string `json:"blobs_url"`
	Commits_url string `json:"commits_url"`
	Watchers_count int `json:"watchers_count,omitempty"`
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"`
	Watchers int `json:"watchers,omitempty"`
	Code_of_conduct GeneratedType `json:"code_of_conduct,omitempty"` // Code Of Conduct
	Issue_comment_url string `json:"issue_comment_url"`
	Releases_url string `json:"releases_url"`
	Size int `json:"size,omitempty"` // The size of the repository, in kilobytes. Size is calculated hourly. When a repository is initially created, the size is 0.
	Ssh_url string `json:"ssh_url,omitempty"`
	Archived bool `json:"archived,omitempty"`
	Language string `json:"language,omitempty"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Assignees_url string `json:"assignees_url"`
	Role_name string `json:"role_name,omitempty"`
	Deployments_url string `json:"deployments_url"`
	Forks int `json:"forks,omitempty"`
	Pulls_url string `json:"pulls_url"`
	Contents_url string `json:"contents_url"`
	Private bool `json:"private"`
	Stargazers_count int `json:"stargazers_count,omitempty"`
	Statuses_url string `json:"statuses_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Membership Webhooksmembership `json:"membership"` // The membership between the user and the organization. Not present when the action is `member_invited`.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Discussion_title_url string `json:"discussion_title_url"` // The URL to the discussion where the secret was detected.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Color string `json:"color,omitempty"` // Color for the issue type.
	Description string `json:"description,omitempty"` // Description of the issue type.
	Is_enabled bool `json:"is_enabled"` // Whether or not the issue type is enabled at the organization level.
	Name string `json:"name"` // Name of the issue type.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository map[string]interface{} `json:"repository,omitempty"` // A git repository
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Team Webhooksteam1 `json:"team"` // Groups of organization members that gives permissions on specified repositories.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"` // The changes to the team if the action was `edited`.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Previous_marketplace_purchase Webhookspreviousmarketplacepurchase `json:"previous_marketplace_purchase,omitempty"`
	Effective_date string `json:"effective_date"`
	Marketplace_purchase Webhooksmarketplacepurchase `json:"marketplace_purchase"`
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// Webhookssponsorship represents the Webhookssponsorship schema from the OpenAPI specification
type Webhookssponsorship struct {
	Created_at string `json:"created_at"`
	Maintainer map[string]interface{} `json:"maintainer,omitempty"`
	Node_id string `json:"node_id"`
	Privacy_level string `json:"privacy_level"`
	Sponsor map[string]interface{} `json:"sponsor"`
	Sponsorable map[string]interface{} `json:"sponsorable"`
	Tier map[string]interface{} `json:"tier"` // The `tier_changed` and `pending_tier_change` will include the original tier before the change or pending change. For more information, see the pending tier change payload.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Runner_type string `json:"runner_type,omitempty"` // Runner type to be used.
	State string `json:"state,omitempty"` // The desired state of code scanning default setup.
	Threat_model string `json:"threat_model,omitempty"` // Threat model to be used for code scanning analysis. Use `remote` to analyze only network sources and `remote_and_local` to include local sources like filesystem access, command-line arguments, database reads, environment variable and standard input.
	Languages []string `json:"languages,omitempty"` // CodeQL languages to be analyzed.
	Query_suite string `json:"query_suite,omitempty"` // CodeQL query suite to be used.
	Runner_label string `json:"runner_label,omitempty"` // Runner label to be used if the runner type is labeled.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id int `json:"id"`
	Login string `json:"login"`
	Node_id string `json:"node_id,omitempty"`
	Organization_billing_email string `json:"organization_billing_email,omitempty"`
	Type string `json:"type"`
	Url string `json:"url"`
	Email string `json:"email,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Label Webhookslabel `json:"label"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// Webhooksrelease represents the Webhooksrelease schema from the OpenAPI specification
type Webhooksrelease struct {
	Draft bool `json:"draft"` // Whether the release is a draft or published
	Assets_url string `json:"assets_url"`
	Author map[string]interface{} `json:"author"`
	Body string `json:"body"`
	Discussion_url string `json:"discussion_url,omitempty"`
	Upload_url string `json:"upload_url"`
	Tarball_url string `json:"tarball_url"`
	Node_id string `json:"node_id"`
	Published_at string `json:"published_at"`
	Target_commitish string `json:"target_commitish"` // Specifies the commitish value that determines where the Git tag is created from.
	Tag_name string `json:"tag_name"` // The name of the tag.
	Html_url string `json:"html_url"`
	Id int `json:"id"`
	Created_at string `json:"created_at"`
	Prerelease bool `json:"prerelease"` // Whether the release is identified as a prerelease or a full release.
	Assets []map[string]interface{} `json:"assets"`
	Name string `json:"name"`
	Url string `json:"url"`
	Reactions map[string]interface{} `json:"reactions,omitempty"`
	Zipball_url string `json:"zipball_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action,omitempty"`
	Check_run GeneratedType `json:"check_run"` // A check performed on the code of a given code change
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Change_status map[string]interface{} `json:"change_status,omitempty"`
	Committed_at string `json:"committed_at,omitempty"`
	Url string `json:"url,omitempty"`
	User GeneratedType `json:"user,omitempty"` // A GitHub user.
	Version string `json:"version,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	_links map[string]interface{} `json:"_links"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Commit_id string `json:"commit_id"` // A commit SHA for the review. If the commit object was garbage collected or forcibly deleted, then it no longer exists in Git and this value will be `null`.
	Pull_request_url string `json:"pull_request_url"`
	Body_text string `json:"body_text,omitempty"`
	Html_url string `json:"html_url"`
	Id int64 `json:"id"` // Unique identifier of the review
	Node_id string `json:"node_id"`
	State string `json:"state"`
	Submitted_at string `json:"submitted_at,omitempty"`
	Body_html string `json:"body_html,omitempty"`
	User GeneratedType `json:"user"` // A GitHub user.
	Body string `json:"body"` // The text of the review.
}

// Package represents the Package schema from the OpenAPI specification
type Package struct {
	Visibility string `json:"visibility"`
	Created_at string `json:"created_at"`
	Name string `json:"name"` // The name of the package.
	Version_count int `json:"version_count"` // The number of versions of the package.
	Url string `json:"url"`
	Html_url string `json:"html_url"`
	Id int `json:"id"` // Unique identifier of the package.
	Owner GeneratedType `json:"owner,omitempty"` // A GitHub user.
	Package_type string `json:"package_type"`
	Updated_at string `json:"updated_at"`
	Repository GeneratedType `json:"repository,omitempty"` // Minimal Repository
}

// Deployment represents the Deployment schema from the OpenAPI specification
type Deployment struct {
	Transient_environment bool `json:"transient_environment,omitempty"` // Specifies if the given environment is will no longer exist at some point in the future. Default: false.
	Url string `json:"url"`
	Created_at string `json:"created_at"`
	Environment string `json:"environment"` // Name for the target deployment environment.
	Creator GeneratedType `json:"creator"` // A GitHub user.
	Node_id string `json:"node_id"`
	Repository_url string `json:"repository_url"`
	Statuses_url string `json:"statuses_url"`
	Description string `json:"description"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Ref string `json:"ref"` // The ref to deploy. This can be a branch, tag, or sha.
	Sha string `json:"sha"`
	Production_environment bool `json:"production_environment,omitempty"` // Specifies if the given environment is one that end-users directly interact with. Default: false.
	Original_environment string `json:"original_environment,omitempty"`
	Task string `json:"task"` // Parameter to specify a task to execute
	Updated_at string `json:"updated_at"`
	Id int64 `json:"id"` // Unique identifier of the deployment
	Payload interface{} `json:"payload"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Duration float64 `json:"duration,omitempty"`
	Id string `json:"id"`
	Start_date string `json:"start_date,omitempty"`
	Title string `json:"title"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Project_url string `json:"project_url"`
	Updated_at string `json:"updated_at"`
	Created_at string `json:"created_at"`
	Project_id string `json:"project_id,omitempty"`
	Archived bool `json:"archived,omitempty"` // Whether or not the card is archived
	Column_url string `json:"column_url"`
	Content_url string `json:"content_url,omitempty"`
	Node_id string `json:"node_id"`
	Column_name string `json:"column_name,omitempty"`
	Creator GeneratedType `json:"creator"` // A GitHub user.
	Id int64 `json:"id"` // The project card's ID
	Note string `json:"note"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Project Webhooksproject `json:"project"`
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// Verification represents the Verification schema from the OpenAPI specification
type Verification struct {
	Signature string `json:"signature"`
	Verified bool `json:"verified"`
	Verified_at string `json:"verified_at"`
	Payload string `json:"payload"`
	Reason string `json:"reason"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url"`
	Status string `json:"status"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Effective_date string `json:"effective_date"`
	Previous_marketplace_purchase map[string]interface{} `json:"previous_marketplace_purchase,omitempty"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Marketplace_purchase Webhooksmarketplacepurchase `json:"marketplace_purchase"`
	Action string `json:"action"`
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	Type string `json:"type"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Personal_access_token_request GeneratedType `json:"personal_access_token_request"` // Details of a Personal Access Token Request.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Project Webhooksproject `json:"project"`
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Registry_package map[string]interface{} `json:"registry_package"`
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url"`
	Details_url string `json:"details_url"`
	App GeneratedType `json:"app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Status string `json:"status"` // The phase of the lifecycle that the check is currently in. Statuses of waiting, requested, and pending are reserved for GitHub Actions check runs.
	Output map[string]interface{} `json:"output"`
	Deployment GeneratedType `json:"deployment,omitempty"` // A deployment created as the result of an Actions check run from a workflow that references an environment
	Name string `json:"name"` // The name of the check.
	Node_id string `json:"node_id"`
	External_id string `json:"external_id"`
	Conclusion string `json:"conclusion"`
	Check_suite map[string]interface{} `json:"check_suite"`
	Head_sha string `json:"head_sha"` // The SHA of the commit that is being checked.
	Started_at string `json:"started_at"`
	Completed_at string `json:"completed_at"`
	Pull_requests []GeneratedType `json:"pull_requests"` // Pull requests that are open with a `head_sha` or `head_branch` that matches the check. The returned pull requests do not necessarily indicate pull requests that triggered the check.
	Html_url string `json:"html_url"`
	Id int64 `json:"id"` // The id of the check.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Squash_merge_commit_message string `json:"squash_merge_commit_message,omitempty"` // The default value for a squash merge commit message: - `PR_BODY` - default to the pull request's body. - `COMMIT_MESSAGES` - default to the branch's commit messages. - `BLANK` - default to a blank commit message.
	Squash_merge_commit_title string `json:"squash_merge_commit_title,omitempty"` // The default value for a squash merge commit title: - `PR_TITLE` - default to the pull request's title. - `COMMIT_OR_PR_TITLE` - default to the commit's title (if only one commit) or the pull request's title (when more than one commit).
	Tags_url string `json:"tags_url"`
	Subscription_url string `json:"subscription_url"`
	Disabled bool `json:"disabled"` // Returns whether or not this repository disabled.
	Languages_url string `json:"languages_url"`
	Branches_url string `json:"branches_url"`
	Subscribers_count int `json:"subscribers_count,omitempty"`
	Merge_commit_title string `json:"merge_commit_title,omitempty"` // The default value for a merge commit title. - `PR_TITLE` - default to the pull request's title. - `MERGE_MESSAGE` - default to the classic title for a merge message (e.g., Merge pull request #123 from branch-name).
	Collaborators_url string `json:"collaborators_url"`
	Size int `json:"size"` // The size of the repository, in kilobytes. Size is calculated hourly. When a repository is initially created, the size is 0.
	Full_name string `json:"full_name"`
	Topics []string `json:"topics,omitempty"`
	Releases_url string `json:"releases_url"`
	Milestones_url string `json:"milestones_url"`
	Allow_rebase_merge bool `json:"allow_rebase_merge,omitempty"` // Whether to allow rebase merges for pull requests.
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Custom_properties map[string]interface{} `json:"custom_properties,omitempty"` // The custom properties that were defined for the repository. The keys are the custom property names, and the values are the corresponding custom property values.
	Labels_url string `json:"labels_url"`
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"` // Whether to delete head branches when pull requests are merged
	Homepage string `json:"homepage"`
	Git_tags_url string `json:"git_tags_url"`
	Language string `json:"language"`
	Name string `json:"name"` // The name of the repository.
	Contributors_url string `json:"contributors_url"`
	Trees_url string `json:"trees_url"`
	Template_repository map[string]interface{} `json:"template_repository,omitempty"`
	Downloads_url string `json:"downloads_url"`
	Forks_count int `json:"forks_count"`
	Git_url string `json:"git_url"`
	Compare_url string `json:"compare_url"`
	Created_at string `json:"created_at"`
	Is_template bool `json:"is_template,omitempty"` // Whether this repository acts as a template that can be used to generate new repositories.
	Open_issues int `json:"open_issues"`
	Anonymous_access_enabled bool `json:"anonymous_access_enabled,omitempty"` // Whether anonymous git access is enabled for this repository
	Teams_url string `json:"teams_url"`
	Has_wiki bool `json:"has_wiki"` // Whether the wiki is enabled.
	Keys_url string `json:"keys_url"`
	Private bool `json:"private"` // Whether the repository is private or public.
	Has_discussions bool `json:"has_discussions,omitempty"` // Whether discussions are enabled.
	Description string `json:"description"`
	Git_commits_url string `json:"git_commits_url"`
	Mirror_url string `json:"mirror_url"`
	Network_count int `json:"network_count,omitempty"`
	Html_url string `json:"html_url"`
	Allow_merge_commit bool `json:"allow_merge_commit,omitempty"` // Whether to allow merge commits for pull requests.
	Contents_url string `json:"contents_url"`
	License GeneratedType `json:"license"` // License Simple
	Allow_forking bool `json:"allow_forking,omitempty"` // Whether to allow forking this repo
	Stargazers_url string `json:"stargazers_url"`
	Updated_at string `json:"updated_at"`
	Comments_url string `json:"comments_url"`
	Watchers_count int `json:"watchers_count"`
	Has_projects bool `json:"has_projects"` // Whether projects are enabled.
	Subscribers_url string `json:"subscribers_url"`
	Allow_update_branch bool `json:"allow_update_branch,omitempty"` // Whether or not a pull request head branch that is behind its base branch can always be updated even if it is not required to be up to date before merging.
	Node_id string `json:"node_id"`
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Open_issues_count int `json:"open_issues_count"`
	Hooks_url string `json:"hooks_url"`
	Visibility string `json:"visibility,omitempty"` // The repository visibility: public, private, or internal.
	Fork bool `json:"fork"`
	Has_pages bool `json:"has_pages"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub user.
	Watchers int `json:"watchers"`
	Id int64 `json:"id"` // Unique identifier of the repository
	Deployments_url string `json:"deployments_url"`
	Forks int `json:"forks"`
	Assignees_url string `json:"assignees_url"`
	Archive_url string `json:"archive_url"`
	Use_squash_pr_title_as_default bool `json:"use_squash_pr_title_as_default,omitempty"` // Whether a squash merge commit can use the pull request title as default. **This property is closing down. Please use `squash_merge_commit_title` instead.
	Issue_comment_url string `json:"issue_comment_url"`
	Clone_url string `json:"clone_url"`
	Forks_url string `json:"forks_url"`
	Master_branch string `json:"master_branch,omitempty"`
	Commits_url string `json:"commits_url"`
	Issue_events_url string `json:"issue_events_url"`
	Merge_commit_message string `json:"merge_commit_message,omitempty"` // The default value for a merge commit message. - `PR_TITLE` - default to the pull request's title. - `PR_BODY` - default to the pull request's body. - `BLANK` - default to a blank commit message.
	Ssh_url string `json:"ssh_url"`
	Pushed_at string `json:"pushed_at"`
	Pulls_url string `json:"pulls_url"`
	Issues_url string `json:"issues_url"`
	Blobs_url string `json:"blobs_url"`
	Allow_squash_merge bool `json:"allow_squash_merge,omitempty"` // Whether to allow squash merges for pull requests.
	Merges_url string `json:"merges_url"`
	Starred_at string `json:"starred_at,omitempty"`
	Allow_auto_merge bool `json:"allow_auto_merge,omitempty"` // Whether to allow Auto-merge to be used on pull requests.
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"` // Whether to require contributors to sign off on web-based commits
	Has_downloads bool `json:"has_downloads"` // Whether downloads are enabled.
	Archived bool `json:"archived"` // Whether the repository is archived.
	Notifications_url string `json:"notifications_url"`
	Temp_clone_token string `json:"temp_clone_token,omitempty"`
	Default_branch string `json:"default_branch"` // The default branch of the repository.
	Has_issues bool `json:"has_issues"` // Whether issues are enabled.
	Git_refs_url string `json:"git_refs_url"`
	Statuses_url string `json:"statuses_url"`
	Svn_url string `json:"svn_url"`
	Url string `json:"url"`
	Stargazers_count int `json:"stargazers_count"`
	Events_url string `json:"events_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Milestone Webhooksmilestone `json:"milestone,omitempty"` // A collection of related issues and pull requests.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Created_at string `json:"created_at,omitempty"`
	Id int `json:"id,omitempty"`
	Key string `json:"key"` // The Base64 encoded public key.
	Key_id string `json:"key_id"` // The identifier for the key.
	Title string `json:"title,omitempty"`
	Url string `json:"url,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation Installation `json:"installation"` // Installation
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repositories []map[string]interface{} `json:"repositories,omitempty"` // An array of repository objects that the installation can access.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Requester interface{} `json:"requester,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Expiry string `json:"expiry,omitempty"` // The duration of the interaction restriction. Default: `one_day`.
	Limit string `json:"limit"` // The type of GitHub user that can comment, open issues, or create pull requests while the interaction limit is in effect.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Label Webhookslabel `json:"label"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Details interface{} `json:"details,omitempty"`
	Type string `json:"type,omitempty"` // The location type. Because secrets may be found in different types of resources (ie. code, comments, issues, pull requests, discussions), this field identifies the type of resource where the secret was found.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Total_active_caches_count int `json:"total_active_caches_count"` // The count of active caches across all repositories of an enterprise or an organization.
	Total_active_caches_size_in_bytes int `json:"total_active_caches_size_in_bytes"` // The total size in bytes of all active cache items across all repositories of an enterprise or an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Public_gists int `json:"public_gists,omitempty"`
	Created_at string `json:"created_at,omitempty"`
	Html_url string `json:"html_url"`
	Following int `json:"following,omitempty"`
	Gists_url string `json:"gists_url"`
	Followers_url string `json:"followers_url"`
	Login string `json:"login"`
	Email string `json:"email,omitempty"`
	Company string `json:"company,omitempty"`
	Avatar_url string `json:"avatar_url"`
	Repos_url string `json:"repos_url"`
	Followers int `json:"followers,omitempty"`
	Type string `json:"type"`
	Subscriptions_url string `json:"subscriptions_url"`
	Updated_at string `json:"updated_at,omitempty"`
	Hireable bool `json:"hireable,omitempty"`
	Id int64 `json:"id"`
	Name string `json:"name,omitempty"`
	Score float64 `json:"score"`
	Suspended_at string `json:"suspended_at,omitempty"`
	Starred_url string `json:"starred_url"`
	Gravatar_id string `json:"gravatar_id"`
	Blog string `json:"blog,omitempty"`
	Public_repos int `json:"public_repos,omitempty"`
	Text_matches []map[string]interface{} `json:"text_matches,omitempty"`
	Url string `json:"url"`
	User_view_type string `json:"user_view_type,omitempty"`
	Organizations_url string `json:"organizations_url"`
	Received_events_url string `json:"received_events_url"`
	Following_url string `json:"following_url"`
	Bio string `json:"bio,omitempty"`
	Node_id string `json:"node_id"`
	Events_url string `json:"events_url"`
	Location string `json:"location,omitempty"`
	Site_admin bool `json:"site_admin"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Actor map[string]interface{} `json:"actor"` // The actor who updated the ruleset
	Updated_at string `json:"updated_at"`
	Version_id int `json:"version_id"` // The ID of the previous version of the ruleset
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Runner_label string `json:"runner_label,omitempty"` // The label of the runner to use for code scanning default setup when runner_type is 'labeled'.
	Runner_type string `json:"runner_type,omitempty"` // Whether to use labeled runners or standard GitHub runners.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Project_card Webhooksprojectcard `json:"project_card"`
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Slug string `json:"slug"`
	Notification_setting string `json:"notification_setting,omitempty"`
	Description string `json:"description"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Repositories_url string `json:"repositories_url"`
	Members_url string `json:"members_url"`
	Assignment string `json:"assignment,omitempty"` // Determines if the team has a direct, indirect, or mixed relationship to a role
	Name string `json:"name"`
	Parent GeneratedType `json:"parent"` // Groups of organization members that gives permissions on specified repositories.
	Permission string `json:"permission"`
	Html_url string `json:"html_url"`
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	Privacy string `json:"privacy,omitempty"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Body_text string `json:"body_text,omitempty"`
	Commit_id string `json:"commit_id"` // A commit SHA for the review.
	Html_url string `json:"html_url"`
	Node_id string `json:"node_id"`
	Body string `json:"body"` // The text of the review.
	Pull_request_url string `json:"pull_request_url"`
	_links map[string]interface{} `json:"_links"`
	Body_html string `json:"body_html,omitempty"`
	Event string `json:"event"`
	Id int `json:"id"` // Unique identifier of the review
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	State string `json:"state"`
	Submitted_at string `json:"submitted_at,omitempty"`
	User GeneratedType `json:"user"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Html_url string `json:"html_url"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Starred_at string `json:"starred_at"` // The time the star was created. This is a timestamp in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`. Will be `null` for the `deleted` action.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Passing int `json:"passing"` // The number of students that have passed the assignment.
	Max_teams int `json:"max_teams,omitempty"` // The maximum allowable teams for the assignment.
	Submitted int `json:"submitted"` // The number of students that have submitted the assignment.
	Id int `json:"id"` // Unique identifier of the repository.
	Max_members int `json:"max_members,omitempty"` // The maximum allowable members per team.
	Deadline string `json:"deadline"` // The time at which the assignment is due.
	Editor string `json:"editor"` // The selected editor for the assignment.
	Invitations_enabled bool `json:"invitations_enabled"` // Whether the invitation link is enabled. Visiting an enabled invitation link will accept the assignment.
	Language string `json:"language"` // The programming language used in the assignment.
	Students_are_repo_admins bool `json:"students_are_repo_admins"` // Whether students are admins on created repository on accepted assignment.
	Classroom GeneratedType `json:"classroom"` // A GitHub Classroom classroom
	Title string `json:"title"` // Assignment title.
	Accepted int `json:"accepted"` // The number of students that have accepted the assignment.
	Public_repo bool `json:"public_repo"` // Whether an accepted assignment creates a public repository.
	Type string `json:"type"` // Whether it's a Group Assignment or Individual Assignment.
	Invite_link string `json:"invite_link"` // The link that a student can use to accept the assignment.
	Slug string `json:"slug"` // Sluggified name of the assignment.
	Feedback_pull_requests_enabled bool `json:"feedback_pull_requests_enabled"` // Whether feedback pull request will be created on assignment acceptance.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Requester Webhooksuser `json:"requester,omitempty"`
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation Installation `json:"installation"` // Installation
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repositories []map[string]interface{} `json:"repositories,omitempty"` // An array of repository objects that the installation can access.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue Webhooksissue `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"` // The action that was performed.
	Assignee Webhooksusermannequin `json:"assignee,omitempty"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Notification_setting string `json:"notification_setting,omitempty"` // The notification setting the team has set
	Permission string `json:"permission"` // Permission that the team will have for its repositories
	Slug string `json:"slug"`
	Name string `json:"name"` // Name of the team
	Url string `json:"url"` // URL for the team
	Html_url string `json:"html_url"`
	Ldap_dn string `json:"ldap_dn,omitempty"` // Distinguished Name (DN) that team maps to within LDAP environment
	Repositories_url string `json:"repositories_url"`
	Description string `json:"description"` // Description of the team
	Id int `json:"id"` // Unique identifier of the team
	Privacy string `json:"privacy,omitempty"` // The level of privacy this team should have
	Members_url string `json:"members_url"`
	Node_id string `json:"node_id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Created_at string `json:"created_at"` // The date and time at which the variable was created, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Name string `json:"name"` // The name of the variable.
	Updated_at string `json:"updated_at"` // The date and time at which the variable was last updated, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Value string `json:"value"` // The value of the variable.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Merge_group GeneratedType `json:"merge_group"` // A group of pull requests that the merge queue has grouped together to be merged.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Name string `json:"name"` // The name of the machine.
	Operating_system string `json:"operating_system"` // The operating system of the machine.
	Prebuild_availability string `json:"prebuild_availability"` // Whether a prebuild is currently available when creating a codespace for this machine and repository. If a branch was not specified as a ref, the default branch will be assumed. Value will be "null" if prebuilds are not supported or prebuild availability could not be determined. Value will be "none" if no prebuild is available. Latest values "ready" and "in_progress" indicate the prebuild availability status.
	Storage_in_bytes int `json:"storage_in_bytes"` // How much storage is available to the codespace.
	Cpus int `json:"cpus"` // How many cores are available to the codespace.
	Display_name string `json:"display_name"` // The display name of the machine includes cores, memory, and storage.
	Memory_in_bytes int `json:"memory_in_bytes"` // How much memory is available to the codespace.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Deployment Deployment `json:"deployment,omitempty"` // A request for a specific ref(branch,sha,tag) to be deployed
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Workflow_job map[string]interface{} `json:"workflow_job"`
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
}

// Webhookspullrequest5 represents the Webhookspullrequest5 schema from the OpenAPI specification
type Webhookspullrequest5 struct {
	Node_id string `json:"node_id"`
	Body string `json:"body"`
	Merged bool `json:"merged,omitempty"`
	Comments_url string `json:"comments_url"`
	Active_lock_reason string `json:"active_lock_reason"`
	Updated_at string `json:"updated_at"`
	Head map[string]interface{} `json:"head"`
	Rebaseable bool `json:"rebaseable,omitempty"`
	Changed_files int `json:"changed_files,omitempty"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Additions int `json:"additions,omitempty"`
	Created_at string `json:"created_at"`
	Html_url string `json:"html_url"`
	Locked bool `json:"locked"`
	Merge_commit_sha string `json:"merge_commit_sha"`
	Review_comment_url string `json:"review_comment_url"`
	Number int `json:"number"` // Number uniquely identifying the pull request within its repository.
	Commits int `json:"commits,omitempty"`
	Assignees []map[string]interface{} `json:"assignees"`
	State string `json:"state"` // State of this Pull Request. Either `open` or `closed`.
	Title string `json:"title"` // The title of the pull request.
	Comments int `json:"comments,omitempty"`
	Closed_at string `json:"closed_at"`
	Review_comments_url string `json:"review_comments_url"`
	_links map[string]interface{} `json:"_links"`
	Mergeable_state string `json:"mergeable_state,omitempty"`
	Requested_reviewers []interface{} `json:"requested_reviewers"`
	Requested_teams []map[string]interface{} `json:"requested_teams"`
	Draft bool `json:"draft"` // Indicates whether or not the pull request is a draft.
	Merged_by map[string]interface{} `json:"merged_by,omitempty"`
	Milestone map[string]interface{} `json:"milestone"` // A collection of related issues and pull requests.
	Labels []map[string]interface{} `json:"labels"`
	Assignee map[string]interface{} `json:"assignee"`
	Base map[string]interface{} `json:"base"`
	User map[string]interface{} `json:"user"`
	Url string `json:"url"`
	Commits_url string `json:"commits_url"`
	Issue_url string `json:"issue_url"`
	Maintainer_can_modify bool `json:"maintainer_can_modify,omitempty"` // Indicates whether maintainers can modify the pull request.
	Merged_at string `json:"merged_at"`
	Auto_merge map[string]interface{} `json:"auto_merge"` // The status of auto merging a pull request.
	Deletions int `json:"deletions,omitempty"`
	Mergeable bool `json:"mergeable,omitempty"`
	Diff_url string `json:"diff_url"`
	Statuses_url string `json:"statuses_url"`
	Review_comments int `json:"review_comments,omitempty"`
	Id int `json:"id"`
	Patch_url string `json:"patch_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Check_suite map[string]interface{} `json:"check_suite"` // The [check_suite](https://docs.github.com/rest/checks/suites#get-a-check-suite).
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Release map[string]interface{} `json:"release"` // The [release](https://docs.github.com/rest/releases/releases/#get-a-release) object.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request GeneratedType `json:"pull_request"`
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Rule GeneratedType `json:"rule"`
	Dismissed_reason string `json:"dismissed_reason"` // **Required when the state is dismissed.** The reason for dismissing or closing the alert.
	Tool GeneratedType `json:"tool"`
	Dismissed_at string `json:"dismissed_at"` // The time that the alert was dismissed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Html_url string `json:"html_url"` // The GitHub URL of the alert resource.
	Dismissal_approved_by GeneratedType `json:"dismissal_approved_by,omitempty"` // A GitHub user.
	Number int `json:"number"` // The security alert number.
	Url string `json:"url"` // The REST API URL of the alert resource.
	Instances_url string `json:"instances_url"` // The REST API URL for fetching the list of instances for an alert.
	Updated_at string `json:"updated_at,omitempty"` // The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Dismissed_comment string `json:"dismissed_comment,omitempty"` // The dismissal comment associated with the dismissal of the alert.
	Most_recent_instance GeneratedType `json:"most_recent_instance"`
	State string `json:"state"` // State of a code scanning alert.
	Created_at string `json:"created_at"` // The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Fixed_at string `json:"fixed_at,omitempty"` // The time that the alert was no longer detected and was considered fixed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Dismissed_by GeneratedType `json:"dismissed_by"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Description string `json:"description"` // A detailed description of what the advisory impacts.
	Severity string `json:"severity,omitempty"` // The severity of the advisory. You must choose between setting this field or `cvss_vector_string`.
	Start_private_fork bool `json:"start_private_fork,omitempty"` // Whether to create a temporary private fork of the repository to collaborate on a fix.
	Summary string `json:"summary"` // A short summary of the advisory.
	Vulnerabilities []map[string]interface{} `json:"vulnerabilities,omitempty"` // An array of products affected by the vulnerability detailed in a repository security advisory.
	Cvss_vector_string string `json:"cvss_vector_string,omitempty"` // The CVSS vector that calculates the severity of the advisory. You must choose between setting this field or `severity`.
	Cwe_ids []string `json:"cwe_ids,omitempty"` // A list of Common Weakness Enumeration (CWE) IDs.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Head_sha string `json:"head_sha"` // The SHA of the commit that is being checked.
	External_id string `json:"external_id"`
	Id int `json:"id"` // The id of the check.
	Url string `json:"url"`
	Check_suite GeneratedType `json:"check_suite"` // A suite of checks performed on the code of a given code change
	Conclusion string `json:"conclusion"`
	Html_url string `json:"html_url"`
	Started_at string `json:"started_at"`
	Status string `json:"status"` // The phase of the lifecycle that the check is currently in.
	Pull_requests []GeneratedType `json:"pull_requests"`
	Deployment GeneratedType `json:"deployment,omitempty"` // A deployment created as the result of an Actions check run from a workflow that references an environment
	Details_url string `json:"details_url"`
	Node_id string `json:"node_id"`
	Output map[string]interface{} `json:"output"`
	App Integration `json:"app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Completed_at string `json:"completed_at"`
	Name string `json:"name"` // The name of the check.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Block_creations map[string]interface{} `json:"block_creations,omitempty"`
	Required_signatures map[string]interface{} `json:"required_signatures,omitempty"`
	Required_status_checks GeneratedType `json:"required_status_checks,omitempty"` // Status Check Policy
	Url string `json:"url"`
	Allow_fork_syncing map[string]interface{} `json:"allow_fork_syncing,omitempty"` // Whether users can pull changes from upstream when the branch is locked. Set to `true` to allow fork syncing. Set to `false` to prevent fork syncing.
	Enforce_admins map[string]interface{} `json:"enforce_admins,omitempty"`
	Required_conversation_resolution map[string]interface{} `json:"required_conversation_resolution,omitempty"`
	Required_linear_history map[string]interface{} `json:"required_linear_history,omitempty"`
	Lock_branch map[string]interface{} `json:"lock_branch,omitempty"` // Whether to set the branch as read-only. If this is true, users will not be able to push to the branch.
	Required_pull_request_reviews map[string]interface{} `json:"required_pull_request_reviews,omitempty"`
	Allow_force_pushes map[string]interface{} `json:"allow_force_pushes,omitempty"`
	Restrictions GeneratedType `json:"restrictions,omitempty"` // Branch Restriction Policy
	Allow_deletions map[string]interface{} `json:"allow_deletions,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Type GeneratedType `json:"type"` // The type of issue.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue Webhooksissue `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Tree_id string `json:"tree_id"` // SHA for the commit's tree
	Author map[string]interface{} `json:"author"` // Information about the Git author
	Committer map[string]interface{} `json:"committer"` // Information about the Git committer
	Id string `json:"id"` // SHA for the commit
	Message string `json:"message"` // Message describing the purpose of the commit
	Timestamp string `json:"timestamp"` // Timestamp of the commit
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Repository_ruleset GeneratedType `json:"repository_ruleset"` // A set of rules to apply when specified conditions are met.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Permission string `json:"permission"`
	User GeneratedType `json:"user"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Commits int `json:"commits"`
	Mergeable_state string `json:"mergeable_state"`
	Auto_merge GeneratedType `json:"auto_merge"` // The status of auto merging a pull request.
	Additions int `json:"additions"`
	Draft bool `json:"draft,omitempty"` // Indicates whether or not the pull request is a draft.
	User GeneratedType `json:"user"` // A GitHub user.
	Updated_at string `json:"updated_at"`
	Locked bool `json:"locked"`
	Number int `json:"number"` // Number uniquely identifying the pull request within its repository.
	Title string `json:"title"` // The title of the pull request.
	Diff_url string `json:"diff_url"`
	Maintainer_can_modify bool `json:"maintainer_can_modify"` // Indicates whether maintainers can modify the pull request.
	Requested_teams []GeneratedType `json:"requested_teams,omitempty"`
	Merged_at string `json:"merged_at"`
	Body string `json:"body"`
	Deletions int `json:"deletions"`
	Assignees []GeneratedType `json:"assignees,omitempty"`
	Closed_at string `json:"closed_at"`
	Head map[string]interface{} `json:"head"`
	State string `json:"state"` // State of this Pull Request. Either `open` or `closed`.
	Review_comments_url string `json:"review_comments_url"`
	Assignee GeneratedType `json:"assignee"` // A GitHub user.
	Comments_url string `json:"comments_url"`
	Merged bool `json:"merged"`
	Rebaseable bool `json:"rebaseable,omitempty"`
	Review_comments int `json:"review_comments"`
	Comments int `json:"comments"`
	Url string `json:"url"`
	_links map[string]interface{} `json:"_links"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Labels []map[string]interface{} `json:"labels"`
	Merged_by GeneratedType `json:"merged_by"` // A GitHub user.
	Patch_url string `json:"patch_url"`
	Commits_url string `json:"commits_url"`
	Changed_files int `json:"changed_files"`
	Created_at string `json:"created_at"`
	Id int64 `json:"id"`
	Issue_url string `json:"issue_url"`
	Active_lock_reason string `json:"active_lock_reason,omitempty"`
	Mergeable bool `json:"mergeable"`
	Milestone GeneratedType `json:"milestone"` // A collection of related issues and pull requests.
	Base map[string]interface{} `json:"base"`
	Html_url string `json:"html_url"`
	Node_id string `json:"node_id"`
	Requested_reviewers []GeneratedType `json:"requested_reviewers,omitempty"`
	Merge_commit_sha string `json:"merge_commit_sha"`
	Review_comment_url string `json:"review_comment_url"`
	Statuses_url string `json:"statuses_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Completed int `json:"completed"`
	Percent_completed int `json:"percent_completed"`
	Total int `json:"total"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action,omitempty"`
	Check_run GeneratedType `json:"check_run"` // A check performed on the code of a given code change
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// Collaborator represents the Collaborator schema from the OpenAPI specification
type Collaborator struct {
	Gravatar_id string `json:"gravatar_id"`
	Subscriptions_url string `json:"subscriptions_url"`
	Site_admin bool `json:"site_admin"`
	Followers_url string `json:"followers_url"`
	Login string `json:"login"`
	Following_url string `json:"following_url"`
	Html_url string `json:"html_url"`
	Role_name string `json:"role_name"`
	Url string `json:"url"`
	Name string `json:"name,omitempty"`
	Repos_url string `json:"repos_url"`
	Type string `json:"type"`
	Starred_url string `json:"starred_url"`
	Avatar_url string `json:"avatar_url"`
	Gists_url string `json:"gists_url"`
	User_view_type string `json:"user_view_type,omitempty"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Organizations_url string `json:"organizations_url"`
	Email string `json:"email,omitempty"`
	Events_url string `json:"events_url"`
	Node_id string `json:"node_id"`
	Received_events_url string `json:"received_events_url"`
	Id int64 `json:"id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Comment string `json:"comment"` // Comment associated with the pending deployment protection rule. **Required when state is not provided.**
	Environment_name string `json:"environment_name"` // The name of the environment to approve or reject.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Security_severity_level string `json:"security_severity_level,omitempty"` // The security severity of the alert.
	Description string `json:"description,omitempty"` // A short description of the rule used to detect the alert.
	Name string `json:"name,omitempty"` // The name of the rule used to detect the alert.
	Severity string `json:"severity,omitempty"` // The severity of the alert.
	Full_description string `json:"full_description,omitempty"` // A description of the rule used to detect the alert.
	Help_uri string `json:"help_uri,omitempty"` // A link to the documentation for the rule used to detect the alert.
	Id string `json:"id,omitempty"` // A unique identifier for the rule used to detect the alert.
	Help string `json:"help,omitempty"` // Detailed documentation for the rule as GitHub Flavored Markdown.
	Tags []string `json:"tags,omitempty"` // A set of tags applicable for the rule.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Requester interface{} `json:"requester,omitempty"`
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation Installation `json:"installation"` // Installation
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repositories []map[string]interface{} `json:"repositories,omitempty"` // An array of repository objects that the installation can access.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Issue_title_url string `json:"issue_title_url"` // The API URL to get the issue where the secret was detected.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Security_advisory Webhookssecurityadvisory `json:"security_advisory"` // The details of the security advisory, including summary, description, and severity.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Estimated_paid_storage_for_month int `json:"estimated_paid_storage_for_month"` // Estimated storage space (GB) used in billing cycle.
	Estimated_storage_for_month int `json:"estimated_storage_for_month"` // Estimated sum of free and paid storage space (GB) used in billing cycle.
	Days_left_in_billing_cycle int `json:"days_left_in_billing_cycle"` // Numbers of days left in billing cycle.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id int `json:"id"`
	Assigner GeneratedType `json:"assigner"` // A GitHub user.
	Node_id string `json:"node_id"`
	Url string `json:"url"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Created_at string `json:"created_at"`
	Performed_via_github_app Integration `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Assignee GeneratedType `json:"assignee"` // A GitHub user.
	Commit_id string `json:"commit_id"`
	Commit_url string `json:"commit_url"`
	Event string `json:"event"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Payload string `json:"payload"` // A URL-encoded string of the check_run.completed JSON payload. The decoded payload is a JSON object.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Plan map[string]interface{} `json:"plan,omitempty"`
	Public_gists int `json:"public_gists"`
	Issues_url string `json:"issues_url"`
	Members_can_create_private_pages bool `json:"members_can_create_private_pages,omitempty"`
	Twitter_username string `json:"twitter_username,omitempty"`
	Updated_at string `json:"updated_at"`
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"`
	Public_members_url string `json:"public_members_url"`
	Default_repository_permission string `json:"default_repository_permission,omitempty"`
	Members_can_create_public_repositories bool `json:"members_can_create_public_repositories,omitempty"`
	Avatar_url string `json:"avatar_url"`
	Disk_usage int `json:"disk_usage,omitempty"`
	Type string `json:"type"`
	Hooks_url string `json:"hooks_url"`
	Has_repository_projects bool `json:"has_repository_projects"`
	Company string `json:"company,omitempty"`
	Blog string `json:"blog,omitempty"`
	Collaborators int `json:"collaborators,omitempty"`
	Followers int `json:"followers"`
	Has_organization_projects bool `json:"has_organization_projects"`
	Members_can_create_repositories bool `json:"members_can_create_repositories,omitempty"`
	Following int `json:"following"`
	Archived_at string `json:"archived_at"`
	Description string `json:"description"`
	Email string `json:"email,omitempty"`
	Members_allowed_repository_creation_type string `json:"members_allowed_repository_creation_type,omitempty"`
	Url string `json:"url"`
	Login string `json:"login"`
	Name string `json:"name,omitempty"`
	Created_at string `json:"created_at"`
	Events_url string `json:"events_url"`
	Billing_email string `json:"billing_email,omitempty"`
	Html_url string `json:"html_url"`
	Location string `json:"location,omitempty"`
	Total_private_repos int `json:"total_private_repos,omitempty"`
	Members_url string `json:"members_url"`
	Public_repos int `json:"public_repos"`
	Members_can_create_public_pages bool `json:"members_can_create_public_pages,omitempty"`
	Node_id string `json:"node_id"`
	Is_verified bool `json:"is_verified,omitempty"`
	Private_gists int `json:"private_gists,omitempty"`
	Repos_url string `json:"repos_url"`
	Members_can_create_private_repositories bool `json:"members_can_create_private_repositories,omitempty"`
	Members_can_fork_private_repositories bool `json:"members_can_fork_private_repositories,omitempty"`
	Owned_private_repos int `json:"owned_private_repos,omitempty"`
	Members_can_create_pages bool `json:"members_can_create_pages,omitempty"`
	Members_can_create_internal_repositories bool `json:"members_can_create_internal_repositories,omitempty"`
	Id int `json:"id"`
	Two_factor_requirement_enabled bool `json:"two_factor_requirement_enabled,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Models []map[string]interface{} `json:"models,omitempty"` // List of model metrics for a custom models and the default model.
	Total_engaged_users int `json:"total_engaged_users,omitempty"` // Total number of users who prompted Copilot Chat on github.com at least once.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pages []map[string]interface{} `json:"pages"` // The pages that were updated.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	Type string `json:"type"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // Full Repository
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Changes map[string]interface{} `json:"changes"`
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

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Number int `json:"number"` // The pull request number.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Assignee Webhooksusermannequin `json:"assignee,omitempty"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Expires_at string `json:"expires_at"`
	Limit string `json:"limit"` // The type of GitHub user that can comment, open issues, or create pull requests while the interaction limit is in effect.
	Origin string `json:"origin"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Run_duration_ms int `json:"run_duration_ms,omitempty"`
	Billable map[string]interface{} `json:"billable"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Alert GeneratedType `json:"alert"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Secret_type_display_name string `json:"secret_type_display_name,omitempty"` // User-friendly name for the detected secret, matching the `secret_type`. For a list of built-in patterns, see "[Supported secret scanning patterns](https://docs.github.com/code-security/secret-scanning/introduction/supported-secret-scanning-patterns#supported-secrets)."
	Push_protection_bypass_request_reviewer GeneratedType `json:"push_protection_bypass_request_reviewer,omitempty"` // A GitHub user.
	Is_base64_encoded bool `json:"is_base64_encoded,omitempty"` // A boolean value representing whether or not alert is base64 encoded
	Number int `json:"number,omitempty"` // The security alert number.
	Push_protection_bypassed_at string `json:"push_protection_bypassed_at,omitempty"` // The time that push protection was bypassed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	State string `json:"state,omitempty"` // Sets the state of the secret scanning alert. You must provide `resolution` when you set the state to `resolved`.
	Validity string `json:"validity,omitempty"` // The token status as of the latest validity check.
	Secret string `json:"secret,omitempty"` // The secret that was detected.
	Updated_at string `json:"updated_at,omitempty"` // The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Created_at string `json:"created_at,omitempty"` // The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Locations_url string `json:"locations_url,omitempty"` // The REST API URL of the code locations for this alert.
	Has_more_locations bool `json:"has_more_locations,omitempty"` // A boolean value representing whether or not the token in the alert was detected in more than one location.
	Resolution string `json:"resolution,omitempty"` // **Required when the `state` is `resolved`.** The reason for resolving the alert.
	Resolved_at string `json:"resolved_at,omitempty"` // The time that the alert was resolved in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Push_protection_bypassed_by GeneratedType `json:"push_protection_bypassed_by,omitempty"` // A GitHub user.
	Repository GeneratedType `json:"repository,omitempty"` // A GitHub repository.
	Publicly_leaked bool `json:"publicly_leaked,omitempty"` // Whether the secret was publicly leaked.
	Push_protection_bypassed bool `json:"push_protection_bypassed,omitempty"` // Whether push protection was bypassed for the detected secret.
	Resolution_comment string `json:"resolution_comment,omitempty"` // The comment that was optionally added when this alert was closed
	Resolved_by GeneratedType `json:"resolved_by,omitempty"` // A GitHub user.
	First_location_detected GeneratedType `json:"first_location_detected,omitempty"` // Details on the location where the token was initially detected. This can be a commit, wiki commit, issue, discussion, pull request.
	Url string `json:"url,omitempty"` // The REST API URL of the alert resource.
	Secret_type string `json:"secret_type,omitempty"` // The type of secret that secret scanning detected.
	Multi_repo bool `json:"multi_repo,omitempty"` // Whether the detected secret was found in multiple repositories in the same organization or enterprise.
	Push_protection_bypass_request_html_url string `json:"push_protection_bypass_request_html_url,omitempty"` // The URL to a push protection bypass request.
	Push_protection_bypass_request_reviewer_comment string `json:"push_protection_bypass_request_reviewer_comment,omitempty"` // An optional comment when reviewing a push protection bypass.
	Html_url string `json:"html_url,omitempty"` // The GitHub URL of the alert resource.
	Push_protection_bypass_request_comment string `json:"push_protection_bypass_request_comment,omitempty"` // An optional comment when requesting a push protection bypass.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id int `json:"id"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Node_id string `json:"node_id"`
	Url string `json:"url"`
	Event string `json:"event"`
	Lock_reason string `json:"lock_reason"`
	Commit_url string `json:"commit_url"`
	Created_at string `json:"created_at"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Commit_id string `json:"commit_id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	UsageItems []map[string]interface{} `json:"usageItems,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Ref string `json:"ref"` // The Git reference of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Action string `json:"action"`
	Commit_oid string `json:"commit_oid"` // The commit SHA of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Alert map[string]interface{} `json:"alert"` // The code scanning alert involved in the event.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Include_claim_keys []string `json:"include_claim_keys"` // Array of unique strings. Each claim key can only contain alphanumeric characters and underscores.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Updated_at string `json:"updated_at"`
	Node_id string `json:"node_id"`
	Repository_url string `json:"repository_url"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Creator GeneratedType `json:"creator"` // A GitHub user.
	State string `json:"state"` // The state of the status.
	Description string `json:"description"` // A short description of the status.
	Deployment_url string `json:"deployment_url"`
	Environment_url string `json:"environment_url,omitempty"` // The URL for accessing your environment.
	Created_at string `json:"created_at"`
	Url string `json:"url"`
	Log_url string `json:"log_url,omitempty"` // The URL to associate with this status.
	Environment string `json:"environment,omitempty"` // The environment of the deployment that the status is for.
	Target_url string `json:"target_url"` // Closing down notice: the URL to associate with this status.
	Id int64 `json:"id"`
}

// Topic represents the Topic schema from the OpenAPI specification
type Topic struct {
	Names []string `json:"names"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Workflow Webhooksworkflow `json:"workflow"`
	Workflow_run map[string]interface{} `json:"workflow_run"`
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// Webhooksissue2 represents the Webhooksissue2 schema from the OpenAPI specification
type Webhooksissue2 struct {
	Events_url string `json:"events_url"`
	Id int64 `json:"id"`
	Performed_via_github_app map[string]interface{} `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Draft bool `json:"draft,omitempty"`
	Updated_at string `json:"updated_at"`
	User map[string]interface{} `json:"user"`
	Active_lock_reason string `json:"active_lock_reason"`
	Type GeneratedType `json:"type,omitempty"` // The type of issue.
	Milestone map[string]interface{} `json:"milestone"` // A collection of related issues and pull requests.
	Labels_url string `json:"labels_url"`
	Created_at string `json:"created_at"`
	Sub_issues_summary map[string]interface{} `json:"sub_issues_summary,omitempty"`
	Assignee map[string]interface{} `json:"assignee,omitempty"`
	Number int `json:"number"`
	Title string `json:"title"` // Title of the issue
	Locked bool `json:"locked,omitempty"`
	Comments_url string `json:"comments_url"`
	Labels []map[string]interface{} `json:"labels,omitempty"`
	Url string `json:"url"` // URL for the issue
	Html_url string `json:"html_url"`
	State string `json:"state,omitempty"` // State of the issue; either 'open' or 'closed'
	Pull_request map[string]interface{} `json:"pull_request,omitempty"`
	Timeline_url string `json:"timeline_url,omitempty"`
	Node_id string `json:"node_id"`
	Repository_url string `json:"repository_url"`
	Reactions map[string]interface{} `json:"reactions"`
	Assignees []map[string]interface{} `json:"assignees"`
	Comments int `json:"comments"`
	State_reason string `json:"state_reason,omitempty"`
	Body string `json:"body"` // Contents of the issue
	Closed_at string `json:"closed_at"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
}

// Milestone represents the Milestone schema from the OpenAPI specification
type Milestone struct {
	Description string `json:"description"`
	Node_id string `json:"node_id"`
	Number int `json:"number"` // The number of the milestone.
	Url string `json:"url"`
	Labels_url string `json:"labels_url"`
	Open_issues int `json:"open_issues"`
	Closed_at string `json:"closed_at"`
	Id int `json:"id"`
	Due_on string `json:"due_on"`
	State string `json:"state"` // The state of the milestone.
	Created_at string `json:"created_at"`
	Closed_issues int `json:"closed_issues"`
	Html_url string `json:"html_url"`
	Title string `json:"title"` // The title of the milestone.
	Creator GeneratedType `json:"creator"` // A GitHub user.
	Updated_at string `json:"updated_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Stargazers_count int `json:"stargazers_count"`
	Git_commits_url string `json:"git_commits_url"`
	License GeneratedType `json:"license"` // License Simple
	Homepage string `json:"homepage"`
	Allow_auto_merge bool `json:"allow_auto_merge,omitempty"` // Whether to allow Auto-merge to be used on pull requests.
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Downloads_url string `json:"downloads_url"`
	Stargazers_url string `json:"stargazers_url"`
	Has_projects bool `json:"has_projects"` // Whether projects are enabled.
	Hooks_url string `json:"hooks_url"`
	Issues_url string `json:"issues_url"`
	Releases_url string `json:"releases_url"`
	Default_branch string `json:"default_branch"` // The default branch of the repository.
	Forks int `json:"forks"`
	Has_issues bool `json:"has_issues"` // Whether issues are enabled.
	Statuses_url string `json:"statuses_url"`
	Contributors_url string `json:"contributors_url"`
	Subscribers_count int `json:"subscribers_count,omitempty"`
	Tags_url string `json:"tags_url"`
	Url string `json:"url"`
	Topics []string `json:"topics,omitempty"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Html_url string `json:"html_url"`
	Subscribers_url string `json:"subscribers_url"`
	Trees_url string `json:"trees_url"`
	Pulls_url string `json:"pulls_url"`
	Node_id string `json:"node_id"`
	Contents_url string `json:"contents_url"`
	Allow_merge_commit bool `json:"allow_merge_commit,omitempty"` // Whether to allow merge commits for pull requests.
	Updated_at string `json:"updated_at"`
	Subscription_url string `json:"subscription_url"`
	Allow_rebase_merge bool `json:"allow_rebase_merge,omitempty"` // Whether to allow rebase merges for pull requests.
	Allow_squash_merge bool `json:"allow_squash_merge,omitempty"` // Whether to allow squash merges for pull requests.
	Watchers int `json:"watchers"`
	Git_url string `json:"git_url"`
	Open_issues_count int `json:"open_issues_count"`
	Has_wiki bool `json:"has_wiki"` // Whether the wiki is enabled.
	Mirror_url string `json:"mirror_url"`
	Milestones_url string `json:"milestones_url"`
	Network_count int `json:"network_count,omitempty"`
	Watchers_count int `json:"watchers_count"`
	Ssh_url string `json:"ssh_url"`
	Issue_events_url string `json:"issue_events_url"`
	Description string `json:"description"`
	Name string `json:"name"` // The name of the repository.
	Git_refs_url string `json:"git_refs_url"`
	Has_downloads bool `json:"has_downloads"` // Whether downloads are enabled.
	Is_template bool `json:"is_template,omitempty"` // Whether this repository acts as a template that can be used to generate new repositories.
	Allow_forking bool `json:"allow_forking,omitempty"` // Whether to allow forking this repo
	Git_tags_url string `json:"git_tags_url"`
	Blobs_url string `json:"blobs_url"`
	Full_name string `json:"full_name"`
	Forks_url string `json:"forks_url"`
	Created_at string `json:"created_at"`
	Collaborators_url string `json:"collaborators_url"`
	Keys_url string `json:"keys_url"`
	Id int `json:"id"` // Unique identifier of the repository
	Pushed_at string `json:"pushed_at"`
	Visibility string `json:"visibility,omitempty"` // The repository visibility: public, private, or internal.
	Branches_url string `json:"branches_url"`
	Deployments_url string `json:"deployments_url"`
	Has_pages bool `json:"has_pages"`
	Language string `json:"language"`
	Forks_count int `json:"forks_count"`
	Size int `json:"size"`
	Labels_url string `json:"labels_url"`
	Assignees_url string `json:"assignees_url"`
	Temp_clone_token string `json:"temp_clone_token,omitempty"`
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"` // Whether to delete head branches when pull requests are merged
	Merges_url string `json:"merges_url"`
	Issue_comment_url string `json:"issue_comment_url"`
	Commits_url string `json:"commits_url"`
	Notifications_url string `json:"notifications_url"`
	Open_issues int `json:"open_issues"`
	Archive_url string `json:"archive_url"`
	Role_name string `json:"role_name,omitempty"`
	Compare_url string `json:"compare_url"`
	Comments_url string `json:"comments_url"`
	Events_url string `json:"events_url"`
	Private bool `json:"private"` // Whether the repository is private or public.
	Master_branch string `json:"master_branch,omitempty"`
	Clone_url string `json:"clone_url"`
	Disabled bool `json:"disabled"` // Returns whether or not this repository disabled.
	Teams_url string `json:"teams_url"`
	Archived bool `json:"archived"` // Whether the repository is archived.
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"` // Whether to require contributors to sign off on web-based commits
	Svn_url string `json:"svn_url"`
	Languages_url string `json:"languages_url"`
	Fork bool `json:"fork"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id int `json:"id"` // The unique identifier of the deployment protection rule integration.
	Integration_url string `json:"integration_url"` // The URL for the endpoint to get details about the app.
	Node_id string `json:"node_id"` // The node ID for the deployment protection rule integration.
	Slug string `json:"slug"` // The slugified name of the deployment protection rule integration.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository_property map[string]interface{} `json:"repository_property"`
}

// Root represents the Root schema from the OpenAPI specification
type Root struct {
	Public_gists_url string `json:"public_gists_url"`
	Starred_gists_url string `json:"starred_gists_url"`
	Notifications_url string `json:"notifications_url"`
	Current_user_url string `json:"current_user_url"`
	Topic_search_url string `json:"topic_search_url,omitempty"`
	Authorizations_url string `json:"authorizations_url"`
	Organization_url string `json:"organization_url"`
	Events_url string `json:"events_url"`
	Commit_search_url string `json:"commit_search_url"`
	Keys_url string `json:"keys_url"`
	Organization_teams_url string `json:"organization_teams_url"`
	Repository_search_url string `json:"repository_search_url"`
	Starred_url string `json:"starred_url"`
	User_organizations_url string `json:"user_organizations_url"`
	Issue_search_url string `json:"issue_search_url"`
	Issues_url string `json:"issues_url"`
	Followers_url string `json:"followers_url"`
	Feeds_url string `json:"feeds_url"`
	Current_user_repositories_url string `json:"current_user_repositories_url"`
	Emojis_url string `json:"emojis_url"`
	User_search_url string `json:"user_search_url"`
	Label_search_url string `json:"label_search_url"`
	Repository_url string `json:"repository_url"`
	User_url string `json:"user_url"`
	Code_search_url string `json:"code_search_url"`
	Hub_url string `json:"hub_url,omitempty"`
	Following_url string `json:"following_url"`
	User_repositories_url string `json:"user_repositories_url"`
	Current_user_authorizations_html_url string `json:"current_user_authorizations_html_url"`
	Gists_url string `json:"gists_url"`
	Rate_limit_url string `json:"rate_limit_url"`
	Organization_repositories_url string `json:"organization_repositories_url"`
	Emails_url string `json:"emails_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Inherited_allows_public_repositories bool `json:"inherited_allows_public_repositories,omitempty"`
	Network_configuration_id string `json:"network_configuration_id,omitempty"` // The identifier of a hosted compute network configuration.
	Selected_workflows []string `json:"selected_workflows,omitempty"` // List of workflows the runner group should be allowed to run. This setting will be ignored unless `restricted_to_workflows` is set to `true`.
	Name string `json:"name"`
	Visibility string `json:"visibility"`
	Default bool `json:"default"`
	Inherited bool `json:"inherited"`
	Restricted_to_workflows bool `json:"restricted_to_workflows,omitempty"` // If `true`, the runner group will be restricted to running only the workflows specified in the `selected_workflows` array.
	Hosted_runners_url string `json:"hosted_runners_url,omitempty"`
	Runners_url string `json:"runners_url"`
	Selected_repositories_url string `json:"selected_repositories_url,omitempty"` // Link to the selected repositories resource for this runner group. Not present unless visibility was set to `selected`
	Workflow_restrictions_read_only bool `json:"workflow_restrictions_read_only,omitempty"` // If `true`, the `restricted_to_workflows` and `selected_workflows` fields cannot be modified.
	Allows_public_repositories bool `json:"allows_public_repositories"`
	Id float64 `json:"id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Threat_model string `json:"threat_model,omitempty"` // Threat model to be used for code scanning analysis. Use `remote` to analyze only network sources and `remote_and_local` to include local sources like filesystem access, command-line arguments, database reads, environment variable and standard input.
	Updated_at string `json:"updated_at,omitempty"` // Timestamp of latest configuration update.
	Languages []string `json:"languages,omitempty"` // Languages to be analyzed.
	Query_suite string `json:"query_suite,omitempty"` // CodeQL query suite to be used.
	Runner_label string `json:"runner_label,omitempty"` // Runner label to be used if the runner type is labeled.
	Runner_type string `json:"runner_type,omitempty"` // Runner type to be used.
	Schedule string `json:"schedule,omitempty"` // The frequency of the periodic analysis.
	State string `json:"state,omitempty"` // Code scanning default setup has been configured or not.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Name string `json:"name"` // The name of the secret.
	Selected_repositories_url string `json:"selected_repositories_url,omitempty"`
	Updated_at string `json:"updated_at"`
	Visibility string `json:"visibility"` // Visibility of a secret
	Created_at string `json:"created_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Project Webhooksproject `json:"project"`
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes,omitempty"` // The changes to the project if the action was `edited`.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Cvss map[string]interface{} `json:"cvss"` // Details for the advisory pertaining to the Common Vulnerability Scoring System.
	Ghsa_id string `json:"ghsa_id"` // The unique GitHub Security Advisory ID assigned to the advisory.
	Identifiers []map[string]interface{} `json:"identifiers"` // Values that identify this advisory among security information sources.
	References []map[string]interface{} `json:"references"` // Links to additional advisory information.
	Cwes []map[string]interface{} `json:"cwes"` // Details for the advisory pertaining to Common Weakness Enumeration.
	Epss GeneratedType `json:"epss,omitempty"` // The EPSS scores as calculated by the [Exploit Prediction Scoring System](https://www.first.org/epss).
	Withdrawn_at string `json:"withdrawn_at"` // The time that the advisory was withdrawn in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Summary string `json:"summary"` // A short, plain text summary of the advisory.
	Description string `json:"description"` // A long-form Markdown-supported description of the advisory.
	Cvss_severities GeneratedType `json:"cvss_severities,omitempty"`
	Vulnerabilities []GeneratedType `json:"vulnerabilities"` // Vulnerable version range information for the advisory.
	Severity string `json:"severity"` // The severity of the advisory.
	Published_at string `json:"published_at"` // The time that the advisory was published in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Updated_at string `json:"updated_at"` // The time that the advisory was last modified in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Cve_id string `json:"cve_id"` // The unique CVE ID assigned to the advisory.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Visibility string `json:"visibility"` // Visibility of a variable
	Created_at string `json:"created_at"` // The date and time at which the variable was created, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Name string `json:"name"` // The name of the variable.
	Selected_repositories_url string `json:"selected_repositories_url,omitempty"`
	Updated_at string `json:"updated_at"` // The date and time at which the variable was last updated, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Value string `json:"value"` // The value of the variable.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Commit_url string `json:"commit_url"`
	Event string `json:"event"`
	Id int `json:"id"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Commit_id string `json:"commit_id"`
	Project_card map[string]interface{} `json:"project_card,omitempty"`
	Created_at string `json:"created_at"`
	Node_id string `json:"node_id"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Name string `json:"name"` // The unique package name within its ecosystem.
	Ecosystem string `json:"ecosystem"` // The package's language or package management ecosystem.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Permission string `json:"permission"` // Permission that the team will have for its repositories
	Description string `json:"description"` // Description of the team
	Members_url string `json:"members_url"`
	Privacy string `json:"privacy,omitempty"` // The level of privacy this team should have
	Id int `json:"id"` // Unique identifier of the team
	Repositories_url string `json:"repositories_url"`
	Slug string `json:"slug"`
	Url string `json:"url"` // URL for the team
	Ldap_dn string `json:"ldap_dn,omitempty"` // Distinguished Name (DN) that team maps to within LDAP environment
	Notification_setting string `json:"notification_setting,omitempty"` // The notification setting the team has set
	Html_url string `json:"html_url"`
	Name string `json:"name"` // Name of the team
	Node_id string `json:"node_id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	UsageItems []map[string]interface{} `json:"usageItems,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Updated_at string `json:"updated_at"`
	Visibility string `json:"visibility"` // Visibility of a secret
	Created_at string `json:"created_at"`
	Name string `json:"name"` // The name of the secret.
	Selected_repositories_url string `json:"selected_repositories_url,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enabled bool `json:"enabled"` // Whether GitHub Actions is enabled on the repository.
	Selected_actions_url string `json:"selected_actions_url,omitempty"` // The API URL to use to get or set the actions and reusable workflows that are allowed to run, when `allowed_actions` is set to `selected`.
	Allowed_actions string `json:"allowed_actions,omitempty"` // The permissions policy that controls the actions and reusable workflows that are allowed to run.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Platform string `json:"platform,omitempty"`
	Version_info map[string]interface{} `json:"version_info,omitempty"`
	Dependencies []map[string]interface{} `json:"dependencies,omitempty"`
	Repo string `json:"repo,omitempty"`
	Readme string `json:"readme,omitempty"`
	Commit_oid string `json:"commit_oid,omitempty"`
	Description string `json:"description,omitempty"`
	Homepage string `json:"homepage,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Name string `json:"name,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Node_id string `json:"node_id"`
	Updated_at string `json:"updated_at"`
	Avatar_url string `json:"avatar_url"`
	Description string `json:"description,omitempty"` // A short description of the enterprise.
	Id int `json:"id"` // Unique identifier of the enterprise
	Name string `json:"name"` // The name of the enterprise.
	Slug string `json:"slug"` // The slug url identifier for the enterprise.
	Website_url string `json:"website_url,omitempty"` // The enterprise's website URL.
	Created_at string `json:"created_at"`
	Html_url string `json:"html_url"`
}

// Webhooksmembership represents the Webhooksmembership schema from the OpenAPI specification
type Webhooksmembership struct {
	Organization_url string `json:"organization_url"`
	Role string `json:"role"`
	State string `json:"state"`
	Url string `json:"url"`
	User map[string]interface{} `json:"user"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Pull_request_review_url string `json:"pull_request_review_url"` // The API URL to get the pull request review where the secret was detected.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Created_at string `json:"created_at"`
	Description string `json:"description"`
	Id int `json:"id"` // Unique identifier of the deployment
	Performed_via_github_app GeneratedType `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Production_environment bool `json:"production_environment,omitempty"` // Specifies if the given environment is one that end-users directly interact with. Default: false.
	Statuses_url string `json:"statuses_url"`
	Task string `json:"task"` // Parameter to specify a task to execute
	Environment string `json:"environment"` // Name for the target deployment environment.
	Node_id string `json:"node_id"`
	Repository_url string `json:"repository_url"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	Transient_environment bool `json:"transient_environment,omitempty"` // Specifies if the given environment is will no longer exist at some point in the future. Default: false.
	Original_environment string `json:"original_environment,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Commit map[string]interface{} `json:"commit"`
	Name string `json:"name"`
	Protected bool `json:"protected"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Actor_id int `json:"actor_id,omitempty"` // The ID of the actor that can bypass a ruleset. If `actor_type` is `OrganizationAdmin`, this should be `1`. If `actor_type` is `DeployKey`, this should be null. `OrganizationAdmin` is not applicable for personal repositories.
	Actor_type string `json:"actor_type"` // The type of actor that can bypass a ruleset.
	Bypass_mode string `json:"bypass_mode,omitempty"` // When the specified actor can bypass the ruleset. `pull_request` means that an actor can only bypass rules on pull requests. `pull_request` is not applicable for the `DeployKey` actor type. Also, `pull_request` is only applicable to branch rulesets.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Member Webhooksuser `json:"member"`
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Team Webhooksteam `json:"team"` // Groups of organization members that gives permissions on specified repositories.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Scope string `json:"scope"` // The scope of the membership. Currently, can only be `team`.
	Sender map[string]interface{} `json:"sender"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue Webhooksissue2 `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Language string `json:"language"` // The language of the CodeQL database.
	Updated_at string `json:"updated_at"` // The date and time at which the CodeQL database was last updated, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Url string `json:"url"` // The URL at which to download the CodeQL database. The `Accept` header must be set to the value of the `content_type` property.
	Size int `json:"size"` // The size of the CodeQL database file in bytes.
	Created_at string `json:"created_at"` // The date and time at which the CodeQL database was created, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Id int `json:"id"` // The ID of the CodeQL database.
	Name string `json:"name"` // The name of the CodeQL database.
	Uploader GeneratedType `json:"uploader"` // A GitHub user.
	Commit_oid string `json:"commit_oid,omitempty"` // The commit SHA of the repository at the time the CodeQL database was created.
	Content_type string `json:"content_type"` // The MIME type of the CodeQL database file.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Html_url string `json:"html_url"`
	Key string `json:"key"`
	Name string `json:"name"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Errors []string `json:"errors,omitempty"` // Any errors that ocurred during processing of the delivery.
	Processing_status string `json:"processing_status,omitempty"` // `pending` files have not yet been processed, while `complete` means results from the SARIF have been stored. `failed` files have either not been processed at all, or could only be partially processed.
	Analyses_url string `json:"analyses_url,omitempty"` // The REST API URL for getting the analyses associated with the upload.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Last_active_on string `json:"last_active_on,omitempty"` // The time at which the runner was last used, in ISO 8601 format.
	Public_ips []GeneratedType `json:"public_ips,omitempty"` // The public IP ranges when public IP is enabled for the hosted runners.
	Runner_group_id int `json:"runner_group_id,omitempty"` // The unique identifier of the group that the hosted runner belongs to.
	Id int `json:"id"` // The unique identifier of the hosted runner.
	Maximum_runners int `json:"maximum_runners,omitempty"` // The maximum amount of hosted runners. Runners will not scale automatically above this number. Use this setting to limit your cost.
	Platform string `json:"platform"` // The operating system of the image.
	Public_ip_enabled bool `json:"public_ip_enabled"` // Whether public IP is enabled for the hosted runners.
	Status string `json:"status"` // The status of the runner.
	Name string `json:"name"` // The name of the hosted runner.
	Image_details GeneratedType `json:"image_details"` // Provides details of a hosted runner image
	Machine_size_details GeneratedType `json:"machine_size_details"` // Provides details of a particular machine spec.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Dismissed_reason string `json:"dismissed_reason"` // **Required when the state is dismissed.** The reason for dismissing or closing the alert.
	Rule GeneratedType `json:"rule"`
	Updated_at string `json:"updated_at,omitempty"` // The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Instances_url string `json:"instances_url"` // The REST API URL for fetching the list of instances for an alert.
	Most_recent_instance GeneratedType `json:"most_recent_instance"`
	Dismissal_approved_by GeneratedType `json:"dismissal_approved_by,omitempty"` // A GitHub user.
	Dismissed_at string `json:"dismissed_at"` // The time that the alert was dismissed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Url string `json:"url"` // The REST API URL of the alert resource.
	State string `json:"state"` // State of a code scanning alert.
	Fixed_at string `json:"fixed_at,omitempty"` // The time that the alert was no longer detected and was considered fixed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Dismissed_comment string `json:"dismissed_comment,omitempty"` // The dismissal comment associated with the dismissal of the alert.
	Number int `json:"number"` // The security alert number.
	Html_url string `json:"html_url"` // The GitHub URL of the alert resource.
	Dismissed_by GeneratedType `json:"dismissed_by"` // A GitHub user.
	Tool GeneratedType `json:"tool"`
	Created_at string `json:"created_at"` // The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Documentation_url string `json:"documentation_url,omitempty"`
	Message string `json:"message,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Status string `json:"status,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Html_url string `json:"html_url"`
	Node_id string `json:"node_id"`
	Repository GeneratedType `json:"repository"` // Minimal Repository
	Url string `json:"url"`
	Comments_url string `json:"comments_url"`
	Score float64 `json:"score"`
	Committer GeneratedType `json:"committer"` // Metaproperties for Git author/committer information.
	Parents []map[string]interface{} `json:"parents"`
	Sha string `json:"sha"`
	Text_matches []map[string]interface{} `json:"text_matches,omitempty"`
	Author GeneratedType `json:"author"` // A GitHub user.
	Commit map[string]interface{} `json:"commit"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Name string `json:"name"`
	Permissions map[string]interface{} `json:"permissions"`
	Created_at string `json:"created_at"`
	Url string `json:"url"`
	Body string `json:"body"`
	Number int `json:"number"`
	Owner_url string `json:"owner_url"`
	State string `json:"state"`
	Columns_url string `json:"columns_url"`
	Id int `json:"id"`
	Updated_at string `json:"updated_at"`
	Node_id string `json:"node_id"`
	Html_url string `json:"html_url"`
	Creator GeneratedType `json:"creator"` // A GitHub user.
	Organization_permission string `json:"organization_permission,omitempty"` // The organization permission for this project. Only present when owner is an organization.
	Private bool `json:"private,omitempty"` // Whether the project is private or not. Only present when owner is an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository_selection string `json:"repository_selection"` // Describe whether all repositories have been selected or there's a selection involved
	Requester Webhooksuser `json:"requester"`
	Installation Installation `json:"installation"` // Installation
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repositories_added []map[string]interface{} `json:"repositories_added"` // An array of repository objects, which were added to the installation.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Repositories_removed []map[string]interface{} `json:"repositories_removed"` // An array of repository objects, which were removed from the installation.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Sponsorship Webhookssponsorship `json:"sponsorship"`
	Action string `json:"action"`
	Effective_date string `json:"effective_date,omitempty"` // The `pending_cancellation` and `pending_tier_change` event types will include the date the cancellation or tier change will take effect.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Projects_v2_item GeneratedType `json:"projects_v2_item"` // An item belonging to a project
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Deleted_by GeneratedType `json:"deleted_by"` // A GitHub user.
	Short_description string `json:"short_description"`
	Public bool `json:"public"`
	Title string `json:"title"`
	Updated_at string `json:"updated_at"`
	Closed_at string `json:"closed_at"`
	Created_at string `json:"created_at"`
	Id float64 `json:"id"`
	Deleted_at string `json:"deleted_at"`
	Description string `json:"description"`
	Node_id string `json:"node_id"`
	Number int `json:"number"`
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Creator GeneratedType `json:"creator"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Dismissed_at string `json:"dismissed_at"` // The time that the alert was dismissed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Html_url string `json:"html_url"` // The GitHub URL of the alert resource.
	Dismissed_by GeneratedType `json:"dismissed_by"` // A GitHub user.
	Dismissed_comment string `json:"dismissed_comment,omitempty"` // The dismissal comment associated with the dismissal of the alert.
	Instances_url string `json:"instances_url"` // The REST API URL for fetching the list of instances for an alert.
	Most_recent_instance GeneratedType `json:"most_recent_instance"`
	Updated_at string `json:"updated_at,omitempty"` // The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Tool GeneratedType `json:"tool"`
	Fixed_at string `json:"fixed_at,omitempty"` // The time that the alert was no longer detected and was considered fixed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	State string `json:"state"` // State of a code scanning alert.
	Dismissed_reason string `json:"dismissed_reason"` // **Required when the state is dismissed.** The reason for dismissing or closing the alert.
	Dismissal_approved_by GeneratedType `json:"dismissal_approved_by,omitempty"` // A GitHub user.
	Rule GeneratedType `json:"rule"`
	Url string `json:"url"` // The REST API URL of the alert resource.
	Number int `json:"number"` // The security alert number.
	Repository GeneratedType `json:"repository"` // A GitHub repository.
	Created_at string `json:"created_at"` // The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Permission string `json:"permission"`
	Role_name string `json:"role_name"`
	User GeneratedType `json:"user"` // Collaborator
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Created_at string `json:"created_at"` // The date and time of when the advisory was created, in ISO 8601 format.
	Html_url string `json:"html_url"` // The URL for the advisory.
	State string `json:"state"` // The state of the advisory.
	Collaborating_users []GeneratedType `json:"collaborating_users"` // A list of users that collaborate on the advisory.
	Collaborating_teams []Team `json:"collaborating_teams"` // A list of teams that collaborate on the advisory.
	Url string `json:"url"` // The API URL for the advisory.
	Credits []map[string]interface{} `json:"credits"`
	Cwe_ids []string `json:"cwe_ids"` // A list of only the CWE IDs.
	Summary string `json:"summary"` // A short summary of the advisory.
	Ghsa_id string `json:"ghsa_id"` // The GitHub Security Advisory ID.
	Identifiers []map[string]interface{} `json:"identifiers"`
	Severity string `json:"severity"` // The severity of the advisory.
	Updated_at string `json:"updated_at"` // The date and time of when the advisory was last updated, in ISO 8601 format.
	Withdrawn_at string `json:"withdrawn_at"` // The date and time of when the advisory was withdrawn, in ISO 8601 format.
	Cvss map[string]interface{} `json:"cvss"`
	Cve_id string `json:"cve_id"` // The Common Vulnerabilities and Exposures (CVE) ID.
	Vulnerabilities []GeneratedType `json:"vulnerabilities"`
	Description string `json:"description"` // A detailed description of what the advisory entails.
	Author interface{} `json:"author"` // The author of the advisory.
	Private_fork interface{} `json:"private_fork"` // A temporary private fork of the advisory's repository for collaborating on a fix.
	Published_at string `json:"published_at"` // The date and time of when the advisory was published, in ISO 8601 format.
	Cvss_severities GeneratedType `json:"cvss_severities,omitempty"`
	Credits_detailed []GeneratedType `json:"credits_detailed"`
	Closed_at string `json:"closed_at"` // The date and time of when the advisory was closed, in ISO 8601 format.
	Publisher interface{} `json:"publisher"` // The publisher of the advisory.
	Cwes []map[string]interface{} `json:"cwes"`
	Submission map[string]interface{} `json:"submission"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Alert GeneratedType `json:"alert"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Color string `json:"color"`
	Name string `json:"name"`
	Node_id string `json:"node_id"`
	Description string `json:"description"`
	Score float64 `json:"score"`
	Id int `json:"id"`
	Text_matches []map[string]interface{} `json:"text_matches,omitempty"`
	Url string `json:"url"`
	Default bool `json:"default"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Ref string `json:"ref"` // The Git reference of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Alert map[string]interface{} `json:"alert"` // The code scanning alert involved in the event.
	Commit_oid string `json:"commit_oid"` // The commit SHA of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Html_url string `json:"html_url,omitempty"`
	Location GeneratedType `json:"location,omitempty"` // Describe a region within a file for the alert.
	Ref string `json:"ref,omitempty"` // The Git reference, formatted as `refs/pull/<number>/merge`, `refs/pull/<number>/head`, `refs/heads/<branch name>` or simply `<branch name>`.
	Category string `json:"category,omitempty"` // Identifies the configuration under which the analysis was executed. Used to distinguish between multiple analyses for the same tool and commit, but performed on different languages or different parts of the code.
	Message map[string]interface{} `json:"message,omitempty"`
	State string `json:"state,omitempty"` // State of a code scanning alert.
	Commit_sha string `json:"commit_sha,omitempty"`
	Environment string `json:"environment,omitempty"` // Identifies the variable values associated with the environment in which the analysis that generated this alert instance was performed, such as the language that was analyzed.
	Analysis_key string `json:"analysis_key,omitempty"` // Identifies the configuration under which the analysis was executed. For example, in GitHub Actions this includes the workflow filename and job name.
	Classifications []string `json:"classifications,omitempty"` // Classifications that have been applied to the file that triggered the alert. For example identifying it as documentation, or a generated file.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Completed_at string `json:"completed_at,omitempty"` // The time that the scan was completed. Empty if the scan is running
	Started_at string `json:"started_at,omitempty"` // The time that the scan was started. Empty if the scan is pending
	Status string `json:"status,omitempty"` // The state of the scan. Either "completed", "running", or "pending"
	Type string `json:"type,omitempty"` // The type of scan
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Projects_v2 GeneratedType `json:"projects_v2"` // A projects v2 project
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Hooks []string `json:"hooks,omitempty"`
	Github_enterprise_importer []string `json:"github_enterprise_importer,omitempty"`
	Ssh_keys []string `json:"ssh_keys,omitempty"`
	Copilot []string `json:"copilot,omitempty"`
	Api []string `json:"api,omitempty"`
	Verifiable_password_authentication bool `json:"verifiable_password_authentication"`
	Actions []string `json:"actions,omitempty"`
	Packages []string `json:"packages,omitempty"`
	Importer []string `json:"importer,omitempty"`
	Domains map[string]interface{} `json:"domains,omitempty"`
	Web []string `json:"web,omitempty"`
	Codespaces []string `json:"codespaces,omitempty"`
	Actions_macos []string `json:"actions_macos,omitempty"`
	Pages []string `json:"pages,omitempty"`
	Ssh_key_fingerprints map[string]interface{} `json:"ssh_key_fingerprints,omitempty"`
	Dependabot []string `json:"dependabot,omitempty"`
	Git []string `json:"git,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Project_card interface{} `json:"project_card"`
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes,omitempty"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Pull_request_comment_url string `json:"pull_request_comment_url"` // The API URL to get the pull request comment where the secret was detected.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Type string `json:"type"` // Whether it's a group assignment or individual assignment.
	Submitted int `json:"submitted"` // The number of students that have submitted the assignment.
	Invite_link string `json:"invite_link"` // The link that a student can use to accept the assignment.
	Students_are_repo_admins bool `json:"students_are_repo_admins"` // Whether students are admins on created repository when a student accepts the assignment.
	Editor string `json:"editor"` // The selected editor for the assignment.
	Classroom Classroom `json:"classroom"` // A GitHub Classroom classroom
	Max_members int `json:"max_members"` // The maximum allowable members per team.
	Starter_code_repository GeneratedType `json:"starter_code_repository"` // A GitHub repository view for Classroom
	Max_teams int `json:"max_teams"` // The maximum allowable teams for the assignment.
	Title string `json:"title"` // Assignment title.
	Accepted int `json:"accepted"` // The number of students that have accepted the assignment.
	Deadline string `json:"deadline"` // The time at which the assignment is due.
	Language string `json:"language"` // The programming language used in the assignment.
	Public_repo bool `json:"public_repo"` // Whether an accepted assignment creates a public repository.
	Feedback_pull_requests_enabled bool `json:"feedback_pull_requests_enabled"` // Whether feedback pull request will be created when a student accepts the assignment.
	Slug string `json:"slug"` // Sluggified name of the assignment.
	Invitations_enabled bool `json:"invitations_enabled"` // Whether the invitation link is enabled. Visiting an enabled invitation link will accept the assignment.
	Passing int `json:"passing"` // The number of students that have passed the assignment.
	Id int `json:"id"` // Unique identifier of the repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Region string `json:"region"` // The location of the subnet this network settings resource is configured for.
	Subnet_id string `json:"subnet_id"` // The subnet this network settings resource is configured for.
	Id string `json:"id"` // The unique identifier of the network settings resource.
	Name string `json:"name"` // The name of the network settings resource.
	Network_configuration_id string `json:"network_configuration_id,omitempty"` // The identifier of the network configuration that is using this settings resource.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Fixed_at string `json:"fixed_at"` // The time that the alert was no longer detected and was considered fixed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Number int `json:"number"` // The security alert number.
	Html_url string `json:"html_url"` // The GitHub URL of the alert resource.
	State string `json:"state"` // The state of the Dependabot alert.
	Dismissed_at string `json:"dismissed_at"` // The time that the alert was dismissed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Repository GeneratedType `json:"repository"` // A GitHub repository.
	Security_advisory GeneratedType `json:"security_advisory"` // Details for the GitHub Security Advisory.
	Security_vulnerability GeneratedType `json:"security_vulnerability"` // Details pertaining to one vulnerable version range for the advisory.
	Dismissed_by GeneratedType `json:"dismissed_by"` // A GitHub user.
	Dismissed_reason string `json:"dismissed_reason"` // The reason that the alert was dismissed.
	Dismissed_comment string `json:"dismissed_comment"` // An optional comment associated with the alert's dismissal.
	Created_at string `json:"created_at"` // The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Auto_dismissed_at string `json:"auto_dismissed_at,omitempty"` // The time that the alert was auto-dismissed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Dependency map[string]interface{} `json:"dependency"` // Details for the vulnerable dependency.
	Updated_at string `json:"updated_at"` // The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Url string `json:"url"` // The REST API URL of the alert resource.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Milestone Webhooksmilestone `json:"milestone"` // A collection of related issues and pull requests.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes,omitempty"` // If the action was `edited`, the changes to the rule.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Rule Webhooksrule `json:"rule"` // The branch protection rule. Includes a `name` and all the [branch protection settings](https://docs.github.com/github/administering-a-repository/defining-the-mergeability-of-pull-requests/about-protected-branches#about-branch-protection-settings) applied to branches that match the name. Binary settings are boolean. Multi-level configurations are one of `off`, `non_admins`, or `everyone`. Actor and build lists are arrays of strings.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Dependabot_security_updates map[string]interface{} `json:"dependabot_security_updates,omitempty"` // Enable or disable Dependabot security updates for the repository.
	Secret_scanning map[string]interface{} `json:"secret_scanning,omitempty"`
	Secret_scanning_ai_detection map[string]interface{} `json:"secret_scanning_ai_detection,omitempty"`
	Secret_scanning_non_provider_patterns map[string]interface{} `json:"secret_scanning_non_provider_patterns,omitempty"`
	Secret_scanning_push_protection map[string]interface{} `json:"secret_scanning_push_protection,omitempty"`
	Advanced_security map[string]interface{} `json:"advanced_security,omitempty"`
	Code_security map[string]interface{} `json:"code_security,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Projects_v2_item GeneratedType `json:"projects_v2_item"` // An item belonging to a project
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// Webhooksanswer represents the Webhooksanswer schema from the OpenAPI specification
type Webhooksanswer struct {
	Node_id string `json:"node_id"`
	Reactions map[string]interface{} `json:"reactions,omitempty"`
	Updated_at string `json:"updated_at"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Body string `json:"body"`
	Html_url string `json:"html_url"`
	Id int `json:"id"`
	Parent_id interface{} `json:"parent_id"`
	Repository_url string `json:"repository_url"`
	User map[string]interface{} `json:"user"`
	Child_comment_count int `json:"child_comment_count"`
	Created_at string `json:"created_at"`
	Discussion_id int `json:"discussion_id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Marketplace_purchase Webhooksmarketplacepurchase `json:"marketplace_purchase"`
	Action string `json:"action"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Effective_date string `json:"effective_date"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Previous_marketplace_purchase Webhookspreviousmarketplacepurchase `json:"previous_marketplace_purchase,omitempty"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	Type string `json:"type"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Comment Webhooksreviewcomment `json:"comment"` // The [comment](https://docs.github.com/rest/pulls/comments#get-a-review-comment-for-a-pull-request) itself.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository_id int64 `json:"repository_id"` // The id of the repository associated with this event.
	Status string `json:"status"` // Describes the response returned after attempting the delivery.
	Duration float64 `json:"duration"` // Time spent delivering.
	Installation_id int64 `json:"installation_id"` // The id of the GitHub App installation associated with this event.
	Event string `json:"event"` // The event that triggered the delivery.
	Id int64 `json:"id"` // Unique identifier of the webhook delivery.
	Redelivery bool `json:"redelivery"` // Whether the webhook delivery is a redelivery.
	Status_code int `json:"status_code"` // Status code received when delivery was made.
	Throttled_at string `json:"throttled_at,omitempty"` // Time when the webhook delivery was throttled.
	Action string `json:"action"` // The type of activity for the event that triggered the delivery.
	Delivered_at string `json:"delivered_at"` // Time when the webhook delivery occurred.
	Guid string `json:"guid"` // Unique identifier for the event (shared with all deliveries for all webhooks that subscribe to this event).
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Week int `json:"week"`
	Days []int `json:"days"`
	Total int `json:"total"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Workflow_job_runs []map[string]interface{} `json:"workflow_job_runs,omitempty"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Reviewers []map[string]interface{} `json:"reviewers,omitempty"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Action string `json:"action"`
	Approver Webhooksapprover `json:"approver,omitempty"`
	Since string `json:"since"`
	Workflow_job_run Webhooksworkflowjobrun `json:"workflow_job_run,omitempty"`
	Workflow_run map[string]interface{} `json:"workflow_run"`
	Comment string `json:"comment,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Definition GeneratedType `json:"definition"` // Custom property defined on an organization
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// Event represents the Event schema from the OpenAPI specification
type Event struct {
	Payload map[string]interface{} `json:"payload"`
	Public bool `json:"public"`
	Repo map[string]interface{} `json:"repo"`
	Type string `json:"type"`
	Actor Actor `json:"actor"` // Actor
	Created_at string `json:"created_at"`
	Id string `json:"id"`
	Org Actor `json:"org,omitempty"` // Actor
}

// Label represents the Label schema from the OpenAPI specification
type Label struct {
	Node_id string `json:"node_id"`
	Url string `json:"url"` // URL for the label
	Color string `json:"color"` // 6-character hex code, without the leading #, identifying the color
	Default bool `json:"default"` // Whether this label comes by default in a new repository.
	Description string `json:"description"` // Optional description of the label, such as its purpose.
	Id int64 `json:"id"` // Unique identifier for the label.
	Name string `json:"name"` // The name of the label.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Workflow_job interface{} `json:"workflow_job"`
	Action string `json:"action"`
	Deployment Deployment `json:"deployment,omitempty"` // A request for a specific ref(branch,sha,tag) to be deployed
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Check_suite map[string]interface{} `json:"check_suite"` // The [check_suite](https://docs.github.com/rest/checks/suites#get-a-check-suite).
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Owned_private_repos int `json:"owned_private_repos"`
	User_view_type string `json:"user_view_type,omitempty"`
	Gravatar_id string `json:"gravatar_id"`
	Node_id string `json:"node_id"`
	Followers_url string `json:"followers_url"`
	Public_gists int `json:"public_gists"`
	Updated_at string `json:"updated_at"`
	Repos_url string `json:"repos_url"`
	Url string `json:"url"`
	Private_gists int `json:"private_gists"`
	Events_url string `json:"events_url"`
	Notification_email string `json:"notification_email,omitempty"`
	Created_at string `json:"created_at"`
	Hireable bool `json:"hireable"`
	Ldap_dn string `json:"ldap_dn,omitempty"`
	Organizations_url string `json:"organizations_url"`
	Location string `json:"location"`
	Gists_url string `json:"gists_url"`
	Following int `json:"following"`
	Business_plus bool `json:"business_plus,omitempty"`
	Total_private_repos int `json:"total_private_repos"`
	Blog string `json:"blog"`
	Name string `json:"name"`
	Received_events_url string `json:"received_events_url"`
	Avatar_url string `json:"avatar_url"`
	Bio string `json:"bio"`
	Two_factor_authentication bool `json:"two_factor_authentication"`
	Public_repos int `json:"public_repos"`
	Followers int `json:"followers"`
	Following_url string `json:"following_url"`
	Type string `json:"type"`
	Plan map[string]interface{} `json:"plan,omitempty"`
	Subscriptions_url string `json:"subscriptions_url"`
	Disk_usage int `json:"disk_usage"`
	Email string `json:"email"`
	Site_admin bool `json:"site_admin"`
	Starred_url string `json:"starred_url"`
	Collaborators int `json:"collaborators"`
	Login string `json:"login"`
	Twitter_username string `json:"twitter_username,omitempty"`
	Html_url string `json:"html_url"`
	Id int64 `json:"id"`
	Company string `json:"company"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Url string `json:"url"`
	Id int `json:"id"`
	Assignee GeneratedType `json:"assignee"` // A GitHub user.
	Commit_id string `json:"commit_id"`
	Commit_url string `json:"commit_url"`
	Created_at string `json:"created_at"`
	Node_id string `json:"node_id"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Assigner GeneratedType `json:"assigner"` // A GitHub user.
	Event string `json:"event"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Require_code_owner_reviews bool `json:"require_code_owner_reviews"`
	Require_last_push_approval bool `json:"require_last_push_approval,omitempty"` // Whether the most recent push must be approved by someone other than the person who pushed it.
	Required_approving_review_count int `json:"required_approving_review_count,omitempty"`
	Url string `json:"url,omitempty"`
	Bypass_pull_request_allowances map[string]interface{} `json:"bypass_pull_request_allowances,omitempty"` // Allow specific users, teams, or apps to bypass pull request requirements.
	Dismiss_stale_reviews bool `json:"dismiss_stale_reviews"`
	Dismissal_restrictions map[string]interface{} `json:"dismissal_restrictions,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Comment map[string]interface{} `json:"comment"` // The [comment](https://docs.github.com/rest/issues/comments#get-an-issue-comment) itself.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) the comment belongs to.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Label Webhookslabel `json:"label,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Github_reviewed_at string `json:"github_reviewed_at"` // The date and time of when the advisory was reviewed by GitHub, in ISO 8601 format.
	Withdrawn_at string `json:"withdrawn_at"` // The date and time of when the advisory was withdrawn, in ISO 8601 format.
	Description string `json:"description"` // A detailed description of what the advisory entails.
	Vulnerabilities []Vulnerability `json:"vulnerabilities"` // The products and respective version ranges affected by the advisory.
	References []string `json:"references"`
	Repository_advisory_url string `json:"repository_advisory_url"` // The API URL for the repository advisory.
	Identifiers []map[string]interface{} `json:"identifiers"`
	Type string `json:"type"` // The type of advisory.
	Credits []map[string]interface{} `json:"credits"` // The users who contributed to the advisory.
	Cvss map[string]interface{} `json:"cvss"`
	Epss GeneratedType `json:"epss,omitempty"` // The EPSS scores as calculated by the [Exploit Prediction Scoring System](https://www.first.org/epss).
	Updated_at string `json:"updated_at"` // The date and time of when the advisory was last updated, in ISO 8601 format.
	Published_at string `json:"published_at"` // The date and time of when the advisory was published, in ISO 8601 format.
	Severity string `json:"severity"` // The severity of the advisory.
	Source_code_location string `json:"source_code_location"` // The URL of the advisory's source code.
	Cve_id string `json:"cve_id"` // The Common Vulnerabilities and Exposures (CVE) ID.
	Cvss_severities GeneratedType `json:"cvss_severities,omitempty"`
	Url string `json:"url"` // The API URL for the advisory.
	Nvd_published_at string `json:"nvd_published_at"` // The date and time when the advisory was published in the National Vulnerability Database, in ISO 8601 format. This field is only populated when the advisory is imported from the National Vulnerability Database.
	Ghsa_id string `json:"ghsa_id"` // The GitHub Security Advisory ID.
	Html_url string `json:"html_url"` // The URL for the advisory.
	Summary string `json:"summary"` // A short summary of the advisory.
	Cwes []map[string]interface{} `json:"cwes"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Updated_at string `json:"updated_at"`
	Client_id string `json:"client_id,omitempty"`
	Description string `json:"description"`
	External_url string `json:"external_url"`
	Id int `json:"id"` // Unique identifier of the GitHub app
	Name string `json:"name"` // The name of the GitHub app
	Node_id string `json:"node_id"`
	Owner interface{} `json:"owner"`
	Html_url string `json:"html_url"`
	Slug string `json:"slug,omitempty"` // The slug name of the GitHub app
	Created_at string `json:"created_at"`
	Events []string `json:"events"` // The list of events for the GitHub app. Note that the `installation_target`, `security_advisory`, and `meta` events are not included because they are global events and not specific to an installation.
	Installations_count int `json:"installations_count,omitempty"` // The number of installations associated with the GitHub app. Only returned when the integration is requesting details about itself.
	Permissions map[string]interface{} `json:"permissions"` // The set of permissions for the GitHub app
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Updated_at string `json:"updated_at,omitempty"` // The time the issue type last updated.
	Color string `json:"color,omitempty"` // The color of the issue type.
	Created_at string `json:"created_at,omitempty"` // The time the issue type created.
	Description string `json:"description"` // The description of the issue type.
	Id int `json:"id"` // The unique identifier of the issue type.
	Is_enabled bool `json:"is_enabled,omitempty"` // The enabled state of the issue type.
	Name string `json:"name"` // The name of the issue type.
	Node_id string `json:"node_id"` // The node identifier of the issue type.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request GeneratedType `json:"pull_request"`
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Projects_v2_item GeneratedType `json:"projects_v2_item"` // An item belonging to a project
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes Webhooksprojectchanges `json:"changes"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Release Webhooksrelease `json:"release"` // The [release](https://docs.github.com/rest/releases/releases/#get-a-release) object.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url"`
	Yearly_price_in_cents int `json:"yearly_price_in_cents"`
	Accounts_url string `json:"accounts_url"`
	Bullets []string `json:"bullets"`
	Has_free_trial bool `json:"has_free_trial"`
	Monthly_price_in_cents int `json:"monthly_price_in_cents"`
	Name string `json:"name"`
	Number int `json:"number"`
	Price_model string `json:"price_model"`
	Unit_name string `json:"unit_name"`
	Description string `json:"description"`
	Id int `json:"id"`
	State string `json:"state"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Personal_access_token_request GeneratedType `json:"personal_access_token_request"` // Details of a Personal Access Token Request.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Current_user_can_approve bool `json:"current_user_can_approve"` // Whether the currently authenticated user can approve the deployment
	Environment map[string]interface{} `json:"environment"`
	Reviewers []map[string]interface{} `json:"reviewers"` // The people or teams that may approve jobs that reference the environment. You can list up to six users or teams as reviewers. The reviewers must have at least read access to the repository. Only one of the required reviewers needs to approve the job for it to proceed.
	Wait_timer int `json:"wait_timer"` // The set duration of the wait timer
	Wait_timer_started_at string `json:"wait_timer_started_at"` // The time that the wait timer began.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Build map[string]interface{} `json:"build"` // The [List GitHub Pages builds](https://docs.github.com/rest/pages/pages#list-github-pages-builds) itself.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Id int `json:"id"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Projects_v2_item GeneratedType `json:"projects_v2_item"` // An item belonging to a project
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes interface{} `json:"changes,omitempty"` // The changes made to the item may involve modifications in the item's fields and draft issue body. It includes altered values for text, number, date, single select, and iteration fields, along with the GraphQL node ID of the changed field.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Project_column Webhooksprojectcolumn `json:"project_column"`
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Commit_url string `json:"commit_url"`
	Id int `json:"id"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Commit_id string `json:"commit_id"`
	Event string `json:"event"`
	Url string `json:"url"`
	Created_at string `json:"created_at"`
	Dismissed_review map[string]interface{} `json:"dismissed_review"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Node_id string `json:"node_id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Protected_branches bool `json:"protected_branches"` // Whether only branches with branch protection rules can deploy to this environment. If `protected_branches` is `true`, `custom_branch_policies` must be `false`; if `protected_branches` is `false`, `custom_branch_policies` must be `true`.
	Custom_branch_policies bool `json:"custom_branch_policies"` // Whether only branches that match the specified name patterns can deploy to this environment. If `custom_branch_policies` is `true`, `protected_branches` must be `false`; if `custom_branch_policies` is `false`, `protected_branches` must be `true`.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Marketplace_purchase map[string]interface{} `json:"marketplace_purchase"`
	Organization_billing_email string `json:"organization_billing_email,omitempty"`
	Type string `json:"type"`
	Url string `json:"url"`
	Email string `json:"email,omitempty"`
	Id int `json:"id"`
	Login string `json:"login"`
	Marketplace_pending_change map[string]interface{} `json:"marketplace_pending_change,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Changes Webhookschanges `json:"changes"` // The changes to the comment.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Issue interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) the comment belongs to.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Comment Webhooksissuecomment `json:"comment"` // The [comment](https://docs.github.com/rest/issues/comments#get-an-issue-comment) itself.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Definition map[string]interface{} `json:"definition"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Projects_v2 GeneratedType `json:"projects_v2"` // A projects v2 project
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Rules_count int `json:"rules_count"` // The total number of rules used in the analysis.
	Deletable bool `json:"deletable"`
	Created_at string `json:"created_at"` // The time that the analysis was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Results_count int `json:"results_count"` // The total number of results in the analysis.
	Tool GeneratedType `json:"tool"`
	Ref string `json:"ref"` // The Git reference, formatted as `refs/pull/<number>/merge`, `refs/pull/<number>/head`, `refs/heads/<branch name>` or simply `<branch name>`.
	Warning string `json:"warning"` // Warning generated when processing the analysis
	Sarif_id string `json:"sarif_id"` // An identifier for the upload.
	Environment string `json:"environment"` // Identifies the variable values associated with the environment in which this analysis was performed.
	Url string `json:"url"` // The REST API URL of the analysis resource.
	Category string `json:"category,omitempty"` // Identifies the configuration under which the analysis was executed. Used to distinguish between multiple analyses for the same tool and commit, but performed on different languages or different parts of the code.
	Commit_sha string `json:"commit_sha"` // The SHA of the commit to which the analysis you are uploading relates.
	Error string `json:"error"`
	Analysis_key string `json:"analysis_key"` // Identifies the configuration under which the analysis was executed. For example, in GitHub Actions this includes the workflow filename and job name.
	Id int `json:"id"` // Unique identifier for this analysis.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Author GeneratedType `json:"author"` // A GitHub user.
	Total int `json:"total"`
	Weeks []map[string]interface{} `json:"weeks"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Description string `json:"description,omitempty"` // Short description of the property
	Required bool `json:"required,omitempty"` // Whether the property is required.
	Value_type string `json:"value_type"` // The type of the value for the property
	Values_editable_by string `json:"values_editable_by,omitempty"` // Who can edit the values of the property
	Allowed_values []string `json:"allowed_values,omitempty"` // An ordered list of the allowed values of the property. The property can have up to 200 allowed values.
	Default_value string `json:"default_value,omitempty"` // Default value of the property
}

// Webhooksproject represents the Webhooksproject schema from the OpenAPI specification
type Webhooksproject struct {
	Updated_at string `json:"updated_at"`
	Body string `json:"body"` // Body of the project
	Created_at string `json:"created_at"`
	Name string `json:"name"` // Name of the project
	Url string `json:"url"`
	Html_url string `json:"html_url"`
	Number int `json:"number"`
	Columns_url string `json:"columns_url"`
	Id int `json:"id"`
	State string `json:"state"` // State of the project; either 'open' or 'closed'
	Creator map[string]interface{} `json:"creator"`
	Node_id string `json:"node_id"`
	Owner_url string `json:"owner_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Milestone Milestone `json:"milestone,omitempty"` // A collection of related issues and pull requests.
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request Webhookspullrequest5 `json:"pull_request"`
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sha string `json:"sha"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Created_at string `json:"created_at"`
	Id int `json:"id"`
	Key string `json:"key"`
	Title string `json:"title"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action,omitempty"`
	Check_run GeneratedType `json:"check_run"` // A check performed on the code of a given code change
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Name string `json:"name"`
	Source string `json:"source"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Is_verified bool `json:"is_verified,omitempty"`
	Total_private_repos int `json:"total_private_repos,omitempty"`
	Email string `json:"email,omitempty"`
	Following int `json:"following"`
	Default_repository_permission string `json:"default_repository_permission,omitempty"`
	Members_can_create_private_repositories bool `json:"members_can_create_private_repositories,omitempty"`
	Node_id string `json:"node_id"`
	Members_can_view_dependency_insights bool `json:"members_can_view_dependency_insights,omitempty"`
	Members_allowed_repository_creation_type string `json:"members_allowed_repository_creation_type,omitempty"`
	Members_can_create_pages bool `json:"members_can_create_pages,omitempty"`
	Has_organization_projects bool `json:"has_organization_projects"`
	Members_url string `json:"members_url"`
	Collaborators int `json:"collaborators,omitempty"` // The number of collaborators on private repositories. This field may be null if the number of private repositories is over 50,000.
	Owned_private_repos int `json:"owned_private_repos,omitempty"`
	Members_can_create_repositories bool `json:"members_can_create_repositories,omitempty"`
	Html_url string `json:"html_url"`
	Secret_scanning_push_protection_custom_link string `json:"secret_scanning_push_protection_custom_link,omitempty"` // An optional URL string to display to contributors who are blocked from pushing a secret.
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"`
	Members_can_change_repo_visibility bool `json:"members_can_change_repo_visibility,omitempty"`
	Has_repository_projects bool `json:"has_repository_projects"`
	Followers int `json:"followers"`
	Dependabot_security_updates_enabled_for_new_repositories bool `json:"dependabot_security_updates_enabled_for_new_repositories,omitempty"` // **Endpoint closing down notice.** Please use [code security configurations](https://docs.github.com/rest/code-security/configurations) instead. Whether Dependabot security updates are automatically enabled for new repositories and repositories transferred to this organization. This field is only visible to organization owners or members of a team with the security manager role.
	Billing_email string `json:"billing_email,omitempty"`
	Members_can_fork_private_repositories bool `json:"members_can_fork_private_repositories,omitempty"`
	Advanced_security_enabled_for_new_repositories bool `json:"advanced_security_enabled_for_new_repositories,omitempty"` // **Endpoint closing down notice.** Please use [code security configurations](https://docs.github.com/rest/code-security/configurations) instead. Whether GitHub Advanced Security is enabled for new repositories and repositories transferred to this organization. This field is only visible to organization owners or members of a team with the security manager role.
	Blog string `json:"blog,omitempty"`
	Members_can_create_public_repositories bool `json:"members_can_create_public_repositories,omitempty"`
	Description string `json:"description"`
	Default_repository_branch string `json:"default_repository_branch,omitempty"` // The default branch for repositories created in this organization.
	Secret_scanning_push_protection_custom_link_enabled bool `json:"secret_scanning_push_protection_custom_link_enabled,omitempty"` // Whether a custom link is shown to contributors who are blocked from pushing a secret by push protection.
	Events_url string `json:"events_url"`
	Members_can_delete_issues bool `json:"members_can_delete_issues,omitempty"`
	Members_can_invite_outside_collaborators bool `json:"members_can_invite_outside_collaborators,omitempty"`
	Type string `json:"type"`
	Dependency_graph_enabled_for_new_repositories bool `json:"dependency_graph_enabled_for_new_repositories,omitempty"` // **Endpoint closing down notice.** Please use [code security configurations](https://docs.github.com/rest/code-security/configurations) instead. Whether dependency graph is automatically enabled for new repositories and repositories transferred to this organization. This field is only visible to organization owners or members of a team with the security manager role.
	Members_can_create_internal_repositories bool `json:"members_can_create_internal_repositories,omitempty"`
	Members_can_create_teams bool `json:"members_can_create_teams,omitempty"`
	Members_can_delete_repositories bool `json:"members_can_delete_repositories,omitempty"`
	Created_at string `json:"created_at"`
	Two_factor_requirement_enabled bool `json:"two_factor_requirement_enabled,omitempty"`
	Updated_at string `json:"updated_at"`
	Issues_url string `json:"issues_url"`
	Disk_usage int `json:"disk_usage,omitempty"`
	Location string `json:"location,omitempty"`
	Avatar_url string `json:"avatar_url"`
	Id int `json:"id"`
	Public_repos int `json:"public_repos"`
	Readers_can_create_discussions bool `json:"readers_can_create_discussions,omitempty"`
	Members_can_create_public_pages bool `json:"members_can_create_public_pages,omitempty"`
	Repos_url string `json:"repos_url"`
	Display_commenter_full_name_setting_enabled bool `json:"display_commenter_full_name_setting_enabled,omitempty"`
	Public_gists int `json:"public_gists"`
	Name string `json:"name,omitempty"`
	Company string `json:"company,omitempty"`
	Private_gists int `json:"private_gists,omitempty"`
	Public_members_url string `json:"public_members_url"`
	Hooks_url string `json:"hooks_url"`
	Plan map[string]interface{} `json:"plan,omitempty"`
	Archived_at string `json:"archived_at"`
	Secret_scanning_enabled_for_new_repositories bool `json:"secret_scanning_enabled_for_new_repositories,omitempty"` // **Endpoint closing down notice.** Please use [code security configurations](https://docs.github.com/rest/code-security/configurations) instead. Whether secret scanning is automatically enabled for new repositories and repositories transferred to this organization. This field is only visible to organization owners or members of a team with the security manager role.
	Twitter_username string `json:"twitter_username,omitempty"`
	Dependabot_alerts_enabled_for_new_repositories bool `json:"dependabot_alerts_enabled_for_new_repositories,omitempty"` // **Endpoint closing down notice.** Please use [code security configurations](https://docs.github.com/rest/code-security/configurations) instead. Whether Dependabot alerts are automatically enabled for new repositories and repositories transferred to this organization. This field is only visible to organization owners or members of a team with the security manager role.
	Url string `json:"url"`
	Deploy_keys_enabled_for_repositories bool `json:"deploy_keys_enabled_for_repositories,omitempty"` // Controls whether or not deploy keys may be added and used for repositories in the organization.
	Secret_scanning_push_protection_enabled_for_new_repositories bool `json:"secret_scanning_push_protection_enabled_for_new_repositories,omitempty"` // **Endpoint closing down notice.** Please use [code security configurations](https://docs.github.com/rest/code-security/configurations) instead. Whether secret scanning push protection is automatically enabled for new repositories and repositories transferred to this organization. This field is only visible to organization owners or members of a team with the security manager role.
	Members_can_create_private_pages bool `json:"members_can_create_private_pages,omitempty"`
	Login string `json:"login"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Blocked_user Webhooksuser `json:"blocked_user"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Issue_body_url string `json:"issue_body_url"` // The API URL to get the issue where the secret was detected.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation Installation `json:"installation"` // Installation
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repositories []map[string]interface{} `json:"repositories,omitempty"` // An array of repository objects that the installation can access.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Requester interface{} `json:"requester,omitempty"`
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Provider string `json:"provider"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	All []int `json:"all"`
	Owner []int `json:"owner"`
}

// Installation represents the Installation schema from the OpenAPI specification
type Installation struct {
	App_id int `json:"app_id"`
	Target_id int `json:"target_id"` // The ID of the user or organization this token is being scoped to.
	Has_multiple_single_files bool `json:"has_multiple_single_files,omitempty"`
	Id int `json:"id"` // The ID of the installation.
	Events []string `json:"events"`
	Repositories_url string `json:"repositories_url"`
	Account interface{} `json:"account"`
	Single_file_paths []string `json:"single_file_paths,omitempty"`
	Access_tokens_url string `json:"access_tokens_url"`
	Html_url string `json:"html_url"`
	Permissions GeneratedType `json:"permissions"` // The permissions granted to the user access token.
	Created_at string `json:"created_at"`
	Repository_selection string `json:"repository_selection"` // Describe whether all repositories have been selected or there's a selection involved
	Updated_at string `json:"updated_at"`
	Contact_email string `json:"contact_email,omitempty"`
	App_slug string `json:"app_slug"`
	Suspended_at string `json:"suspended_at"`
	Suspended_by GeneratedType `json:"suspended_by"` // A GitHub user.
	Target_type string `json:"target_type"`
	Single_file_name string `json:"single_file_name"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Assignee GeneratedType `json:"assignee"` // A GitHub user.
	Event string `json:"event"`
	Id int `json:"id"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Commit_id string `json:"commit_id"`
	Commit_url string `json:"commit_url"`
	Url string `json:"url"`
	Created_at string `json:"created_at"`
	Node_id string `json:"node_id"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
}

// License represents the License schema from the OpenAPI specification
type License struct {
	Html_url string `json:"html_url"`
	Implementation string `json:"implementation"`
	Limitations []string `json:"limitations"`
	Spdx_id string `json:"spdx_id"`
	Url string `json:"url"`
	Body string `json:"body"`
	Conditions []string `json:"conditions"`
	Featured bool `json:"featured"`
	Permissions []string `json:"permissions"`
	Description string `json:"description"`
	Key string `json:"key"`
	Name string `json:"name"`
	Node_id string `json:"node_id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Commits int `json:"commits"`
	Mergeable_state string `json:"mergeable_state"`
	Auto_merge GeneratedType `json:"auto_merge"` // The status of auto merging a pull request.
	Additions int `json:"additions"`
	Draft bool `json:"draft,omitempty"` // Indicates whether or not the pull request is a draft.
	User GeneratedType `json:"user"` // A GitHub user.
	Updated_at string `json:"updated_at"`
	Locked bool `json:"locked"`
	Number int `json:"number"` // Number uniquely identifying the pull request within its repository.
	Title string `json:"title"` // The title of the pull request.
	Diff_url string `json:"diff_url"`
	Maintainer_can_modify bool `json:"maintainer_can_modify"` // Indicates whether maintainers can modify the pull request.
	Requested_teams []GeneratedType `json:"requested_teams,omitempty"`
	Merged_at string `json:"merged_at"`
	Body string `json:"body"`
	Deletions int `json:"deletions"`
	Assignees []GeneratedType `json:"assignees,omitempty"`
	Closed_at string `json:"closed_at"`
	Head map[string]interface{} `json:"head"`
	State string `json:"state"` // State of this Pull Request. Either `open` or `closed`.
	Review_comments_url string `json:"review_comments_url"`
	Assignee GeneratedType `json:"assignee"` // A GitHub user.
	Comments_url string `json:"comments_url"`
	Merged bool `json:"merged"`
	Rebaseable bool `json:"rebaseable,omitempty"`
	Review_comments int `json:"review_comments"`
	Comments int `json:"comments"`
	Url string `json:"url"`
	_links map[string]interface{} `json:"_links"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Labels []map[string]interface{} `json:"labels"`
	Merged_by GeneratedType `json:"merged_by"` // A GitHub user.
	Patch_url string `json:"patch_url"`
	Commits_url string `json:"commits_url"`
	Changed_files int `json:"changed_files"`
	Created_at string `json:"created_at"`
	Id int64 `json:"id"`
	Issue_url string `json:"issue_url"`
	Active_lock_reason string `json:"active_lock_reason,omitempty"`
	Mergeable bool `json:"mergeable"`
	Milestone GeneratedType `json:"milestone"` // A collection of related issues and pull requests.
	Base map[string]interface{} `json:"base"`
	Html_url string `json:"html_url"`
	Node_id string `json:"node_id"`
	Requested_reviewers []GeneratedType `json:"requested_reviewers,omitempty"`
	Merge_commit_sha string `json:"merge_commit_sha"`
	Review_comment_url string `json:"review_comment_url"`
	Statuses_url string `json:"statuses_url"`
	Use_squash_pr_title_as_default bool `json:"use_squash_pr_title_as_default,omitempty"` // Whether a squash merge commit can use the pull request title as default. **This property is closing down. Please use `squash_merge_commit_title` instead.**
	Allow_auto_merge bool `json:"allow_auto_merge,omitempty"` // Whether to allow auto-merge for pull requests.
	Allow_update_branch bool `json:"allow_update_branch,omitempty"` // Whether to allow updating the pull request's branch.
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"` // Whether to delete head branches when pull requests are merged.
	Merge_commit_message string `json:"merge_commit_message,omitempty"` // The default value for a merge commit message. - `PR_TITLE` - default to the pull request's title. - `PR_BODY` - default to the pull request's body. - `BLANK` - default to a blank commit message.
	Merge_commit_title string `json:"merge_commit_title,omitempty"` // The default value for a merge commit title. - `PR_TITLE` - default to the pull request's title. - `MERGE_MESSAGE` - default to the classic title for a merge message (e.g., "Merge pull request #123 from branch-name").
	Squash_merge_commit_message string `json:"squash_merge_commit_message,omitempty"` // The default value for a squash merge commit message: - `PR_BODY` - default to the pull request's body. - `COMMIT_MESSAGES` - default to the branch's commit messages. - `BLANK` - default to a blank commit message.
	Squash_merge_commit_title string `json:"squash_merge_commit_title,omitempty"` // The default value for a squash merge commit title: - `PR_TITLE` - default to the pull request's title. - `COMMIT_OR_PR_TITLE` - default to the commit's title (if only one commit) or the pull request's title (when more than one commit).
}

// Discussion represents the Discussion schema from the OpenAPI specification
type Discussion struct {
	Locked bool `json:"locked"`
	Node_id string `json:"node_id"`
	Updated_at string `json:"updated_at"`
	Active_lock_reason string `json:"active_lock_reason"`
	Created_at string `json:"created_at"`
	Id int `json:"id"`
	Number int `json:"number"`
	Title string `json:"title"`
	Html_url string `json:"html_url"`
	State_reason string `json:"state_reason"` // The reason for the current state
	Answer_chosen_by map[string]interface{} `json:"answer_chosen_by"`
	Answer_html_url string `json:"answer_html_url"`
	State string `json:"state"` // The current state of the discussion. `converting` means that the discussion is being converted from an issue. `transferring` means that the discussion is being transferred from another repository.
	Labels []Label `json:"labels,omitempty"`
	Timeline_url string `json:"timeline_url,omitempty"`
	Answer_chosen_at string `json:"answer_chosen_at"`
	Body string `json:"body"`
	Repository_url string `json:"repository_url"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	User map[string]interface{} `json:"user"`
	Reactions map[string]interface{} `json:"reactions,omitempty"`
	Category map[string]interface{} `json:"category"`
	Comments int `json:"comments"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Project_column Webhooksprojectcolumn `json:"project_column"`
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Alert GeneratedType `json:"alert"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Workflow Webhooksworkflow `json:"workflow,omitempty"`
	Deployment map[string]interface{} `json:"deployment"` // The [deployment](https://docs.github.com/rest/deployments/deployments#list-deployments).
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Workflow_run map[string]interface{} `json:"workflow_run,omitempty"`
	Check_run map[string]interface{} `json:"check_run,omitempty"`
	Deployment_status map[string]interface{} `json:"deployment_status"` // The [deployment status](https://docs.github.com/rest/deployments/statuses#list-deployment-statuses).
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Expire_at string `json:"expire_at,omitempty"` // The time that the bypass will expire in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Reason string `json:"reason,omitempty"` // The reason for bypassing push protection.
	Token_type string `json:"token_type,omitempty"` // The token type this bypass is for.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Key string `json:"key"` // The Base64 encoded public key.
	Key_id string `json:"key_id"` // The identifier for the key.
}

// Webhooksworkflow represents the Webhooksworkflow schema from the OpenAPI specification
type Webhooksworkflow struct {
	Node_id string `json:"node_id"`
	Updated_at string `json:"updated_at"`
	Created_at string `json:"created_at"`
	State string `json:"state"`
	Url string `json:"url"`
	Id int `json:"id"`
	Path string `json:"path"`
	Badge_url string `json:"badge_url"`
	Html_url string `json:"html_url"`
	Name string `json:"name"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Node_id string `json:"node_id"`
	Target_date string `json:"target_date,omitempty"`
	Created_at string `json:"created_at"`
	Creator GeneratedType `json:"creator,omitempty"` // A GitHub user.
	Project_node_id string `json:"project_node_id,omitempty"`
	Body string `json:"body,omitempty"` // Body of the status update
	Start_date string `json:"start_date,omitempty"`
	Status string `json:"status,omitempty"`
	Updated_at string `json:"updated_at"`
	Id float64 `json:"id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Count int `json:"count"`
	Referrer string `json:"referrer"`
	Uniques int `json:"uniques"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request Webhookspullrequest5 `json:"pull_request"`
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Milestone Milestone `json:"milestone,omitempty"` // A collection of related issues and pull requests.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Changes map[string]interface{} `json:"changes"` // The changes to the comment if the action was `edited`.
	Pull_request GeneratedType `json:"pull_request"`
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Plan GeneratedType `json:"plan"` // Marketplace Listing Plan
	Unit_count int `json:"unit_count"`
	Updated_at string `json:"updated_at"`
	Account GeneratedType `json:"account"`
	Billing_cycle string `json:"billing_cycle"`
	Free_trial_ends_on string `json:"free_trial_ends_on"`
	Next_billing_date string `json:"next_billing_date"`
	On_free_trial bool `json:"on_free_trial"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Full_description string `json:"full_description,omitempty"` // A description of the rule used to detect the alert.
	Id string `json:"id,omitempty"` // A unique identifier for the rule used to detect the alert.
	Severity string `json:"severity,omitempty"` // The severity of the alert.
	Description string `json:"description,omitempty"` // A short description of the rule used to detect the alert.
	Help_uri string `json:"help_uri,omitempty"` // A link to the documentation for the rule used to detect the alert.
	Help string `json:"help,omitempty"` // Detailed documentation for the rule as GitHub Flavored Markdown.
	Security_severity_level string `json:"security_severity_level,omitempty"` // The security severity of the alert.
	Name string `json:"name,omitempty"` // The name of the rule used to detect the alert.
	Tags []string `json:"tags,omitempty"` // A set of tags applicable for the rule.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sha string `json:"sha"` // SHA for the commit
	Author map[string]interface{} `json:"author"` // Identifying information for the git-user
	Url string `json:"url"`
	Tree map[string]interface{} `json:"tree"`
	Verification map[string]interface{} `json:"verification"`
	Html_url string `json:"html_url"`
	Parents []map[string]interface{} `json:"parents"`
	Committer map[string]interface{} `json:"committer"` // Identifying information for the git-user
	Message string `json:"message"` // Message describing the purpose of the commit
	Node_id string `json:"node_id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Key Webhooksdeploykey `json:"key"` // The [`deploy key`](https://docs.github.com/rest/deploy-keys/deploy-keys#get-a-deploy-key) resource.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Sponsorship Webhookssponsorship `json:"sponsorship"`
	Action string `json:"action"`
	Changes Webhookschanges8 `json:"changes"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Html_url string `json:"html_url"`
	Id int64 `json:"id"` // Unique identifier of the repository invitation.
	Node_id string `json:"node_id"`
	Repository GeneratedType `json:"repository"` // Minimal Repository
	Invitee GeneratedType `json:"invitee"` // A GitHub user.
	Inviter GeneratedType `json:"inviter"` // A GitHub user.
	Permissions string `json:"permissions"` // The permission associated with the invitation.
	Created_at string `json:"created_at"`
	Expired bool `json:"expired,omitempty"` // Whether or not the invitation has expired
	Url string `json:"url"` // URL for the repository invitation
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Assignee Webhooksuser `json:"assignee,omitempty"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue Webhooksissue `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"` // The action that was performed.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	Type string `json:"type"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Membership Webhooksmembership `json:"membership"` // The membership between the user and the organization. Not present when the action is `member_invited`.
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Billable map[string]interface{} `json:"billable"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Updated_at string `json:"updated_at"` // The date and time at which the secret was created, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Visibility string `json:"visibility"` // The type of repositories in the organization that the secret is visible to
	Created_at string `json:"created_at"` // The date and time at which the secret was created, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Name string `json:"name"` // The name of the secret
	Selected_repositories_url string `json:"selected_repositories_url,omitempty"` // The API URL at which the list of repositories this secret is visible to can be retrieved
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Alert GeneratedType `json:"alert"` // A Dependabot alert.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Status string `json:"status"` // Description of the status of the attempted delivery
	Action string `json:"action"` // The type of activity for the event that triggered the delivery.
	Delivered_at string `json:"delivered_at"` // Time when the delivery was delivered.
	Redelivery bool `json:"redelivery"` // Whether the delivery is a redelivery.
	Id int `json:"id"` // Unique identifier of the delivery.
	Response map[string]interface{} `json:"response"`
	Url string `json:"url,omitempty"` // The URL target of the delivery.
	Guid string `json:"guid"` // Unique identifier for the event (shared with all deliveries for all webhooks that subscribe to this event).
	Throttled_at string `json:"throttled_at,omitempty"` // Time when the webhook delivery was throttled.
	Duration float64 `json:"duration"` // Time spent delivering.
	Installation_id int `json:"installation_id"` // The id of the GitHub App installation associated with this event.
	Request map[string]interface{} `json:"request"`
	Status_code int `json:"status_code"` // Status code received when delivery was made.
	Event string `json:"event"` // The event that triggered the delivery.
	Repository_id int `json:"repository_id"` // The id of the repository associated with this event.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes,omitempty"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Member Webhooksuser `json:"member"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	End_column int `json:"end_column"`
	Path string `json:"path"`
	Raw_details string `json:"raw_details"`
	Start_line int `json:"start_line"`
	End_line int `json:"end_line"`
	Start_column int `json:"start_column"`
	Annotation_level string `json:"annotation_level"`
	Blob_href string `json:"blob_href"`
	Message string `json:"message"`
	Title string `json:"title"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Projects_v2 GeneratedType `json:"projects_v2"` // A projects v2 project
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Ref_name map[string]interface{} `json:"ref_name,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Node_id string `json:"node_id"`
	Object map[string]interface{} `json:"object"`
	Ref string `json:"ref"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action,omitempty"`
	Alert GeneratedType `json:"alert"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Location GeneratedType `json:"location"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Alert GeneratedType `json:"alert"` // A Dependabot alert.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Status string `json:"status,omitempty"` // The attachment status of the code security configuration on the repository.
	Repository GeneratedType `json:"repository,omitempty"` // A GitHub repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	First_patched_version map[string]interface{} `json:"first_patched_version"` // Details pertaining to the package version that patches this vulnerability.
	Package GeneratedType `json:"package"` // Details for the vulnerable package.
	Severity string `json:"severity"` // The severity of the vulnerability.
	Vulnerable_version_range string `json:"vulnerable_version_range"` // Conditions that identify vulnerable versions of this vulnerability's package.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Since string `json:"since"`
	Workflow_job_run map[string]interface{} `json:"workflow_job_run"`
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Reviewers []map[string]interface{} `json:"reviewers"`
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Environment string `json:"environment"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Requestor Webhooksuser `json:"requestor"`
	Workflow_run map[string]interface{} `json:"workflow_run"`
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Package map[string]interface{} `json:"package"` // The name of the package affected by the vulnerability.
	Patched_versions string `json:"patched_versions"` // The package version(s) that resolve the vulnerability.
	Vulnerable_functions []string `json:"vulnerable_functions"` // The functions in the package that are affected.
	Vulnerable_version_range string `json:"vulnerable_version_range"` // The range of the package versions affected by the vulnerability.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Registry_type string `json:"registry_type"` // The registry type.
	Selected_repository_ids []int `json:"selected_repository_ids,omitempty"` // An array of repository IDs that can access the organization private registry when `visibility` is set to `selected`.
	Updated_at string `json:"updated_at"`
	Username string `json:"username,omitempty"` // The username to use when authenticating with the private registry.
	Visibility string `json:"visibility"` // Which type of organization repositories have access to the private registry. `selected` means only the repositories specified by `selected_repository_ids` can access the private registry.
	Created_at string `json:"created_at"`
	Name string `json:"name"` // The name of the private registry configuration.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Project Webhooksproject `json:"project"`
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	Type string `json:"type"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Name string `json:"name"` // The name pattern that branches must match in order to deploy to the environment. Wildcard characters will not match `/`. For example, to match branches that begin with `release/` and contain an additional single slash, use `release/*/*`. For more information about pattern matching syntax, see the [Ruby File.fnmatch documentation](https://ruby-doc.org/core-2.5.1/File.html#method-c-fnmatch).
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Description string `json:"description"`
	Url string `json:"url"`
	Id int `json:"id"`
	Issues_url string `json:"issues_url"`
	Login string `json:"login"`
	Members_url string `json:"members_url"`
	Node_id string `json:"node_id"`
	Public_members_url string `json:"public_members_url"`
	Events_url string `json:"events_url"`
	Hooks_url string `json:"hooks_url"`
	Repos_url string `json:"repos_url"`
	Avatar_url string `json:"avatar_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Parent_issue_id float64 `json:"parent_issue_id"` // The ID of the parent issue.
	Parent_issue_repo Repository `json:"parent_issue_repo"` // A repository on GitHub.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Sub_issue_id float64 `json:"sub_issue_id"` // The ID of the sub-issue.
	Parent_issue Issue `json:"parent_issue"` // Issues are a great way to keep track of tasks, enhancements, and bugs for your projects.
	Action string `json:"action"`
	Sub_issue Issue `json:"sub_issue"` // Issues are a great way to keep track of tasks, enhancements, and bugs for your projects.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Comment string `json:"comment,omitempty"` // Optional comment to include with the review.
	Environment_name string `json:"environment_name"` // The name of the environment to approve or reject.
	State string `json:"state"` // Whether to approve or reject deployment to the specified environments.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Comment Webhooksreviewcomment `json:"comment"` // The [comment](https://docs.github.com/rest/pulls/comments#get-a-review-comment-for-a-pull-request) itself.
	Pull_request map[string]interface{} `json:"pull_request"`
	Changes Webhookschanges `json:"changes"` // The changes to the comment.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Content string `json:"content,omitempty"`
	Encoding string `json:"encoding,omitempty"`
	Entries []map[string]interface{} `json:"entries,omitempty"`
	Name string `json:"name"`
	Path string `json:"path"`
	Sha string `json:"sha"`
	Type string `json:"type"`
	Download_url string `json:"download_url"`
	Size int `json:"size"`
	Git_url string `json:"git_url"`
	Html_url string `json:"html_url"`
	Url string `json:"url"`
	_links map[string]interface{} `json:"_links"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Expires_at string `json:"expires_at"`
	Primary_key_id int `json:"primary_key_id"`
	Raw_key string `json:"raw_key"`
	Id int64 `json:"id"`
	Public_key string `json:"public_key"`
	Can_sign bool `json:"can_sign"`
	Can_certify bool `json:"can_certify"`
	Subkeys []map[string]interface{} `json:"subkeys"`
	Emails []map[string]interface{} `json:"emails"`
	Name string `json:"name,omitempty"`
	Created_at string `json:"created_at"`
	Revoked bool `json:"revoked"`
	Can_encrypt_comms bool `json:"can_encrypt_comms"`
	Can_encrypt_storage bool `json:"can_encrypt_storage"`
	Key_id string `json:"key_id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Description string `json:"description,omitempty"` // A detailed description of what the advisory impacts.
	Summary string `json:"summary,omitempty"` // A short summary of the advisory.
	Severity string `json:"severity,omitempty"` // The severity of the advisory. You must choose between setting this field or `cvss_vector_string`.
	Vulnerabilities []map[string]interface{} `json:"vulnerabilities,omitempty"` // A product affected by the vulnerability detailed in a repository security advisory.
	Collaborating_users []string `json:"collaborating_users,omitempty"` // A list of usernames who have been granted write access to the advisory.
	Credits []map[string]interface{} `json:"credits,omitempty"` // A list of users receiving credit for their participation in the security advisory.
	Cve_id string `json:"cve_id,omitempty"` // The Common Vulnerabilities and Exposures (CVE) ID.
	State string `json:"state,omitempty"` // The state of the advisory.
	Collaborating_teams []string `json:"collaborating_teams,omitempty"` // A list of team slugs which have been granted write access to the advisory.
	Cvss_vector_string string `json:"cvss_vector_string,omitempty"` // The CVSS vector that calculates the severity of the advisory. You must choose between setting this field or `severity`.
	Cwe_ids []string `json:"cwe_ids,omitempty"` // A list of Common Weakness Enumeration (CWE) IDs.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Project_card Webhooksprojectcard `json:"project_card"`
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Preferences map[string]interface{} `json:"preferences"`
	Repository GeneratedType `json:"repository"` // Minimal Repository
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
}

// Webhooksteam represents the Webhooksteam schema from the OpenAPI specification
type Webhooksteam struct {
	Members_url string `json:"members_url,omitempty"`
	Privacy string `json:"privacy,omitempty"`
	Permission string `json:"permission,omitempty"` // Permission that the team will have for its repositories
	Slug string `json:"slug,omitempty"`
	Url string `json:"url,omitempty"` // URL for the team
	Deleted bool `json:"deleted,omitempty"`
	Html_url string `json:"html_url,omitempty"`
	Node_id string `json:"node_id,omitempty"`
	Notification_setting string `json:"notification_setting,omitempty"`
	Parent map[string]interface{} `json:"parent,omitempty"`
	Description string `json:"description,omitempty"` // Description of the team
	Name string `json:"name"` // Name of the team
	Repositories_url string `json:"repositories_url,omitempty"`
	Id int `json:"id"` // Unique identifier of the team
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Created_at string `json:"created_at"`
	Ignored bool `json:"ignored"` // Determines if all notifications should be blocked from this repository.
	Reason string `json:"reason"`
	Repository_url string `json:"repository_url"`
	Subscribed bool `json:"subscribed"` // Determines if notifications should be received from this repository.
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue Webhooksissue `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Label Webhookslabel `json:"label,omitempty"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Action string `json:"action"`
	Label Webhookslabel `json:"label,omitempty"`
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request map[string]interface{} `json:"pull_request"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Full_name string `json:"full_name"` // The full, globally unique, name of the repository.
	Id int `json:"id"` // A unique identifier of the repository.
	Name string `json:"name"` // The name of the repository.
	Private bool `json:"private"` // Whether the repository is private.
	Stargazers_count int `json:"stargazers_count"`
	Updated_at string `json:"updated_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Alert map[string]interface{} `json:"alert"` // The security alert of the vulnerable dependency.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Repository_ruleset GeneratedType `json:"repository_ruleset"` // A set of rules to apply when specified conditions are met.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Required_status_checks GeneratedType `json:"required_status_checks,omitempty"` // Protected Branch Required Status Check
	Block_creations map[string]interface{} `json:"block_creations,omitempty"`
	Allow_deletions map[string]interface{} `json:"allow_deletions,omitempty"`
	Name string `json:"name,omitempty"`
	Protection_url string `json:"protection_url,omitempty"`
	Allow_force_pushes map[string]interface{} `json:"allow_force_pushes,omitempty"`
	Allow_fork_syncing map[string]interface{} `json:"allow_fork_syncing,omitempty"` // Whether users can pull changes from upstream when the branch is locked. Set to `true` to allow fork syncing. Set to `false` to prevent fork syncing.
	Required_conversation_resolution map[string]interface{} `json:"required_conversation_resolution,omitempty"`
	Lock_branch map[string]interface{} `json:"lock_branch,omitempty"` // Whether to set the branch as read-only. If this is true, users will not be able to push to the branch.
	Required_linear_history map[string]interface{} `json:"required_linear_history,omitempty"`
	Restrictions GeneratedType `json:"restrictions,omitempty"` // Branch Restriction Policy
	Enforce_admins GeneratedType `json:"enforce_admins,omitempty"` // Protected Branch Admin Enforced
	Required_signatures map[string]interface{} `json:"required_signatures,omitempty"`
	Url string `json:"url,omitempty"`
	Enabled bool `json:"enabled,omitempty"`
	Required_pull_request_reviews GeneratedType `json:"required_pull_request_reviews,omitempty"` // Protected Branch Pull Request Review
}

// Enterprise represents the Enterprise schema from the OpenAPI specification
type Enterprise struct {
	Id int `json:"id"` // Unique identifier of the enterprise
	Node_id string `json:"node_id"`
	Website_url string `json:"website_url,omitempty"` // The enterprise's website URL.
	Name string `json:"name"` // The name of the enterprise.
	Updated_at string `json:"updated_at"`
	Avatar_url string `json:"avatar_url"`
	Description string `json:"description,omitempty"` // A short description of the enterprise.
	Html_url string `json:"html_url"`
	Slug string `json:"slug"` // The slug url identifier for the enterprise.
	Created_at string `json:"created_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Rate_limited_request_count int64 `json:"rate_limited_request_count,omitempty"` // The total number of requests that were rate limited within the queried time period
	Total_request_count int64 `json:"total_request_count,omitempty"` // The total number of requests within the queried time period
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Payload string `json:"payload"` // A URL-encoded string of the ping JSON payload. The decoded payload is a JSON object.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url,omitempty"`
	Checks []map[string]interface{} `json:"checks"`
	Contexts []string `json:"contexts"`
	Contexts_url string `json:"contexts_url,omitempty"`
	Enforcement_level string `json:"enforcement_level,omitempty"`
	Strict bool `json:"strict,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Body_text string `json:"body_text,omitempty"`
	Html_url string `json:"html_url"`
	Node_id string `json:"node_id"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Id int64 `json:"id"` // Unique identifier of the issue comment
	Url string `json:"url"` // URL for the issue comment
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Body string `json:"body,omitempty"` // Contents of the issue comment
	Created_at string `json:"created_at"`
	Issue_url string `json:"issue_url"`
	Reactions GeneratedType `json:"reactions,omitempty"`
	Updated_at string `json:"updated_at"`
	User GeneratedType `json:"user"` // A GitHub user.
	Body_html string `json:"body_html,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Type string `json:"type"`
	Parameters map[string]interface{} `json:"parameters,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Default_level string `json:"default_level,omitempty"` // The default repository access level for Dependabot updates.
	Accessible_repositories []GeneratedType `json:"accessible_repositories,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Checks []map[string]interface{} `json:"checks"`
	Contexts []string `json:"contexts"`
	Contexts_url string `json:"contexts_url"`
	Strict bool `json:"strict"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Node_id string `json:"node_id,omitempty"`
	Comments []GeneratedType `json:"comments,omitempty"`
	Commit_id string `json:"commit_id,omitempty"`
	Event string `json:"event,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository_full_name string `json:"repository_full_name"`
	Repository_id int `json:"repository_id"`
	Repository_name string `json:"repository_name"`
	Properties []GeneratedType `json:"properties"` // List of custom property names and associated values
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes,omitempty"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Projects_v2_status_update GeneratedType `json:"projects_v2_status_update"` // An status update belonging to a project
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Name string `json:"name"` // The name pattern that branches or tags must match in order to deploy to the environment. Wildcard characters will not match `/`. For example, to match branches that begin with `release/` and contain an additional single slash, use `release/*/*`. For more information about pattern matching syntax, see the [Ruby File.fnmatch documentation](https://ruby-doc.org/core-2.5.1/File.html#method-c-fnmatch).
	Type string `json:"type,omitempty"` // Whether this rule targets a branch or tag
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Member Webhooksuser `json:"member"`
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Scope string `json:"scope"` // The scope of the membership. Currently, can only be `team`.
	Sender map[string]interface{} `json:"sender"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Team Webhooksteam `json:"team"` // Groups of organization members that gives permissions on specified repositories.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Alert GeneratedType `json:"alert"` // A Dependabot alert.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// Migration represents the Migration schema from the OpenAPI specification
type Migration struct {
	Exclude_git_data bool `json:"exclude_git_data"`
	Exclude_attachments bool `json:"exclude_attachments"`
	Lock_repositories bool `json:"lock_repositories"`
	Node_id string `json:"node_id"`
	State string `json:"state"`
	Id int64 `json:"id"`
	Url string `json:"url"`
	Exclude_owner_projects bool `json:"exclude_owner_projects"`
	Archive_url string `json:"archive_url,omitempty"`
	Exclude_releases bool `json:"exclude_releases"`
	Guid string `json:"guid"`
	Org_metadata_only bool `json:"org_metadata_only"`
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Repositories []Repository `json:"repositories"` // The repositories included in the migration. Only returned for export migrations.
	Created_at string `json:"created_at"`
	Exclude []string `json:"exclude,omitempty"` // Exclude related items from being returned in the response in order to improve performance of the request. The array can include any of: `"repositories"`.
	Exclude_metadata bool `json:"exclude_metadata"`
	Updated_at string `json:"updated_at"`
}

// Webhookschanges represents the Webhookschanges schema from the OpenAPI specification
type Webhookschanges struct {
	Body map[string]interface{} `json:"body,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Default_workflow_permissions string `json:"default_workflow_permissions,omitempty"` // The default workflow permissions granted to the GITHUB_TOKEN when running workflows.
	Can_approve_pull_request_reviews bool `json:"can_approve_pull_request_reviews,omitempty"` // Whether GitHub Actions can approve pull requests. Enabling this can be a security risk.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Starred_at interface{} `json:"starred_at"` // The time the star was created. This is a timestamp in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`. Will be `null` for the `deleted` action.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// Webhooksreview represents the Webhooksreview schema from the OpenAPI specification
type Webhooksreview struct {
	Node_id string `json:"node_id"`
	Html_url string `json:"html_url"`
	Submitted_at string `json:"submitted_at"`
	User map[string]interface{} `json:"user"`
	Commit_id string `json:"commit_id"` // A commit SHA for the review.
	Id int `json:"id"` // Unique identifier of the review
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Pull_request_url string `json:"pull_request_url"`
	State string `json:"state"`
	_links map[string]interface{} `json:"_links"`
	Body string `json:"body"` // The text of the review.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Key string `json:"key"`
	Title string `json:"title"`
	Url string `json:"url"`
	Verified bool `json:"verified"`
	Added_by string `json:"added_by,omitempty"`
	Enabled bool `json:"enabled,omitempty"`
	Id int `json:"id"`
	Read_only bool `json:"read_only"`
	Created_at string `json:"created_at"`
	Last_used string `json:"last_used,omitempty"`
}

// Webhooksworkflowjobrun represents the Webhooksworkflowjobrun schema from the OpenAPI specification
type Webhooksworkflowjobrun struct {
	Status string `json:"status"`
	Updated_at string `json:"updated_at"`
	Conclusion interface{} `json:"conclusion"`
	Created_at string `json:"created_at"`
	Environment string `json:"environment"`
	Html_url string `json:"html_url"`
	Id int `json:"id"`
	Name interface{} `json:"name"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Updated_at string `json:"updated_at"`
	Body_html string `json:"body_html"`
	Last_edited_at string `json:"last_edited_at"`
	Pinned bool `json:"pinned"` // Whether or not this discussion should be pinned for easy retrieval.
	Body_version string `json:"body_version"` // The current version of the body content. If provided, this update operation will be rejected if the given version does not match the latest version on the server.
	Team_url string `json:"team_url"`
	Author GeneratedType `json:"author"` // A GitHub user.
	Html_url string `json:"html_url"`
	Node_id string `json:"node_id"`
	Body string `json:"body"` // The main text of the discussion.
	Comments_url string `json:"comments_url"`
	Number int `json:"number"` // The unique sequence number of a team discussion.
	Created_at string `json:"created_at"`
	Reactions GeneratedType `json:"reactions,omitempty"`
	Title string `json:"title"` // The title of the discussion.
	Url string `json:"url"`
	Comments_count int `json:"comments_count"`
	Private bool `json:"private"` // Whether or not this discussion should be restricted to team members and organization owners.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Email string `json:"email,omitempty"`
	Id int64 `json:"id"`
	Repos_url string `json:"repos_url"`
	Avatar_url string `json:"avatar_url"`
	Subscriptions_url string `json:"subscriptions_url"`
	Received_events_url string `json:"received_events_url"`
	Type string `json:"type"`
	Following_url string `json:"following_url"`
	Starred_url string `json:"starred_url"`
	Name string `json:"name,omitempty"`
	Login string `json:"login"`
	Html_url string `json:"html_url"`
	Gists_url string `json:"gists_url"`
	Node_id string `json:"node_id"`
	Starred_at string `json:"starred_at,omitempty"`
	User_view_type string `json:"user_view_type,omitempty"`
	Organizations_url string `json:"organizations_url"`
	Site_admin bool `json:"site_admin"`
	Url string `json:"url"`
	Gravatar_id string `json:"gravatar_id"`
	Events_url string `json:"events_url"`
	Followers_url string `json:"followers_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	Type string `json:"type"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Discussion_body_url string `json:"discussion_body_url"` // The URL to the discussion where the secret was detected.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sha string `json:"sha,omitempty"` // SHA of commit with autofix.
	Target_ref string `json:"target_ref,omitempty"` // The Git reference of target branch for the commit. For more information, see "[Git References](https://git-scm.com/book/en/v2/Git-Internals-Git-References)" in the Git documentation.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Alert Webhooksalert `json:"alert"` // The security alert of the vulnerable dependency.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Package map[string]interface{} `json:"package"` // Information about the package.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Type string `json:"type"` // What type of scan was completed
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Started_at string `json:"started_at"` // The time that the alert was resolved in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Action string `json:"action"`
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Completed_at string `json:"completed_at"` // The time that the alert was resolved in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Custom_pattern_name string `json:"custom_pattern_name,omitempty"` // If the scan was triggered by a custom pattern update, this will be the name of the pattern that was updated
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Secret_types []string `json:"secret_types,omitempty"` // List of patterns that were updated. This will be empty for normal backfill scans or custom pattern updates
	Source string `json:"source"` // What type of content was scanned
	Custom_pattern_scope string `json:"custom_pattern_scope,omitempty"` // If the scan was triggered by a custom pattern update, this will be the scope of the pattern that was updated
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Project_column Webhooksprojectcolumn `json:"project_column"`
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sha string `json:"sha"`
	Line_numbers []string `json:"line_numbers,omitempty"`
	Repository GeneratedType `json:"repository"` // Minimal Repository
	Score float64 `json:"score"`
	File_size int `json:"file_size,omitempty"`
	Name string `json:"name"`
	Text_matches []map[string]interface{} `json:"text_matches,omitempty"`
	Url string `json:"url"`
	Git_url string `json:"git_url"`
	Html_url string `json:"html_url"`
	Language string `json:"language,omitempty"`
	Last_modified_at string `json:"last_modified_at,omitempty"`
	Path string `json:"path"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Check_run GeneratedType `json:"check_run"` // A check performed on the code of a given code change
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Requested_action map[string]interface{} `json:"requested_action,omitempty"` // The action requested by the user.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Description string `json:"description,omitempty"`
	Id string `json:"id"`
	Name string `json:"name"`
	Color string `json:"color,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Total_count int `json:"total_count"` // Total number of caches
	Actions_caches []map[string]interface{} `json:"actions_caches"` // Array of caches
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sha string `json:"sha"`
	State string `json:"state"`
	Statuses []GeneratedType `json:"statuses"`
	Total_count int `json:"total_count"`
	Url string `json:"url"`
	Commit_url string `json:"commit_url"`
	Repository GeneratedType `json:"repository"` // Minimal Repository
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Type string `json:"type"`
	Parameters map[string]interface{} `json:"parameters,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Cvss_v3 map[string]interface{} `json:"cvss_v3,omitempty"`
	Cvss_v4 map[string]interface{} `json:"cvss_v4,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Projects_v2 GeneratedType `json:"projects_v2"` // A projects v2 project
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// Email represents the Email schema from the OpenAPI specification
type Email struct {
	Email string `json:"email"`
	Primary bool `json:"primary"`
	Verified bool `json:"verified"`
	Visibility string `json:"visibility"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Comment Webhookscomment `json:"comment"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Commit map[string]interface{} `json:"commit"`
	Content map[string]interface{} `json:"content"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Commit_oid string `json:"commit_oid"` // The commit SHA of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Ref string `json:"ref"` // The Git reference of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Alert map[string]interface{} `json:"alert"` // The code scanning alert involved in the event.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository map[string]interface{} `json:"repository,omitempty"` // A git repository
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Team Webhooksteam1 `json:"team"` // Groups of organization members that gives permissions on specified repositories.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// Import represents the Import schema from the OpenAPI specification
type Import struct {
	Message string `json:"message,omitempty"`
	Push_percent int `json:"push_percent,omitempty"`
	Repository_url string `json:"repository_url"`
	Svn_root string `json:"svn_root,omitempty"`
	Url string `json:"url"`
	Has_large_files bool `json:"has_large_files,omitempty"`
	Project_choices []map[string]interface{} `json:"project_choices,omitempty"`
	Status_text string `json:"status_text,omitempty"`
	Tfvc_project string `json:"tfvc_project,omitempty"`
	Large_files_count int `json:"large_files_count,omitempty"`
	Html_url string `json:"html_url"`
	Commit_count int `json:"commit_count,omitempty"`
	Large_files_size int `json:"large_files_size,omitempty"`
	Use_lfs bool `json:"use_lfs,omitempty"`
	Failed_step string `json:"failed_step,omitempty"`
	Svc_root string `json:"svc_root,omitempty"`
	Authors_count int `json:"authors_count,omitempty"`
	Import_percent int `json:"import_percent,omitempty"`
	Error_message string `json:"error_message,omitempty"`
	Status string `json:"status"`
	Vcs string `json:"vcs"`
	Vcs_url string `json:"vcs_url"` // The URL of the originating repository.
	Authors_url string `json:"authors_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Push_protection_bypass_request_reviewer_comment string `json:"push_protection_bypass_request_reviewer_comment,omitempty"` // An optional comment when reviewing a push protection bypass.
	Resolved_by GeneratedType `json:"resolved_by,omitempty"` // A GitHub user.
	Publicly_leaked bool `json:"publicly_leaked,omitempty"` // Whether the detected secret was publicly leaked.
	Secret_type_display_name string `json:"secret_type_display_name,omitempty"` // User-friendly name for the detected secret, matching the `secret_type`. For a list of built-in patterns, see "[Supported secret scanning patterns](https://docs.github.com/code-security/secret-scanning/introduction/supported-secret-scanning-patterns#supported-secrets)."
	Url string `json:"url,omitempty"` // The REST API URL of the alert resource.
	Push_protection_bypassed bool `json:"push_protection_bypassed,omitempty"` // Whether push protection was bypassed for the detected secret.
	Resolution string `json:"resolution,omitempty"` // The reason for resolving the alert.
	Html_url string `json:"html_url,omitempty"` // The GitHub URL of the alert resource.
	Multi_repo bool `json:"multi_repo,omitempty"` // Whether the detected secret was found in multiple repositories in the same organization or business.
	Push_protection_bypass_request_comment string `json:"push_protection_bypass_request_comment,omitempty"` // An optional comment when requesting a push protection bypass.
	Push_protection_bypassed_at string `json:"push_protection_bypassed_at,omitempty"` // The time that push protection was bypassed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Locations_url string `json:"locations_url,omitempty"` // The REST API URL of the code locations for this alert.
	Push_protection_bypassed_by GeneratedType `json:"push_protection_bypassed_by,omitempty"` // A GitHub user.
	Resolved_at string `json:"resolved_at,omitempty"` // The time that the alert was resolved in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Push_protection_bypass_request_reviewer GeneratedType `json:"push_protection_bypass_request_reviewer,omitempty"` // A GitHub user.
	Updated_at string `json:"updated_at,omitempty"` // The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Number int `json:"number,omitempty"` // The security alert number.
	Resolution_comment string `json:"resolution_comment,omitempty"` // An optional comment to resolve an alert.
	Secret_type string `json:"secret_type,omitempty"` // The type of secret that secret scanning detected.
	Created_at string `json:"created_at,omitempty"` // The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Push_protection_bypass_request_html_url string `json:"push_protection_bypass_request_html_url,omitempty"` // The URL to a push protection bypass request.
	Validity string `json:"validity,omitempty"` // The token status as of the latest validity check.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Key Webhooksdeploykey `json:"key"` // The [`deploy key`](https://docs.github.com/rest/deploy-keys/deploy-keys#get-a-deploy-key) resource.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Path string `json:"path"` // The path to the workflow file
	Ref string `json:"ref,omitempty"` // The ref (branch or tag) of the workflow file to use
	Repository_id int `json:"repository_id"` // The ID of the repository where the workflow is defined
	Sha string `json:"sha,omitempty"` // The commit SHA of the workflow file to use
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repositories []map[string]interface{} `json:"repositories,omitempty"` // An array of repository objects that the installation can access.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Requester interface{} `json:"requester,omitempty"`
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation Installation `json:"installation"` // Installation
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Comment Webhookscomment `json:"comment"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// Webhooksprojectcolumn represents the Webhooksprojectcolumn schema from the OpenAPI specification
type Webhooksprojectcolumn struct {
	After_id int `json:"after_id,omitempty"`
	Project_url string `json:"project_url"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
	Cards_url string `json:"cards_url"`
	Id int `json:"id"` // The unique identifier of the project column
	Name string `json:"name"` // Name of the project column
	Node_id string `json:"node_id"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Views []Traffic `json:"views"`
	Count int `json:"count"`
	Uniques int `json:"uniques"`
}

// Link represents the Link schema from the OpenAPI specification
type Link struct {
	Href string `json:"href"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Team Webhooksteam1 `json:"team"` // Groups of organization members that gives permissions on specified repositories.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Type string `json:"type"`
	Parameters map[string]interface{} `json:"parameters,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Effective_date string `json:"effective_date"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Marketplace_purchase Webhooksmarketplacepurchase `json:"marketplace_purchase"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Previous_marketplace_purchase map[string]interface{} `json:"previous_marketplace_purchase,omitempty"`
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Pull_request_title_url string `json:"pull_request_title_url"` // The API URL to get the pull request where the secret was detected.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id int `json:"id,omitempty"`
	Key string `json:"key"` // The Base64 encoded public key.
	Key_id string `json:"key_id"` // The identifier for the key.
	Title string `json:"title,omitempty"`
	Url string `json:"url,omitempty"`
	Created_at string `json:"created_at,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	User GeneratedType `json:"user"` // A GitHub user.
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Body string `json:"body"` // The comment text.
	Created_at string `json:"created_at"`
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Login string `json:"login"`
	Role string `json:"role"`
	Email string `json:"email"`
	Failed_reason string `json:"failed_reason,omitempty"`
	Node_id string `json:"node_id"`
	Created_at string `json:"created_at"`
	Id int64 `json:"id"`
	Invitation_teams_url string `json:"invitation_teams_url"`
	Team_count int `json:"team_count"`
	Failed_at string `json:"failed_at,omitempty"`
	Invitation_source string `json:"invitation_source,omitempty"`
	Inviter GeneratedType `json:"inviter"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action,omitempty"`
	Environment string `json:"environment,omitempty"` // The name of the environment that has the deployment protection rule.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Deployment_callback_url string `json:"deployment_callback_url,omitempty"` // The URL to review the deployment protection rule.
	Deployment Deployment `json:"deployment,omitempty"` // A request for a specific ref(branch,sha,tag) to be deployed
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Event string `json:"event,omitempty"` // The event that triggered the deployment protection rule.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_requests []GeneratedType `json:"pull_requests,omitempty"`
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Spdx_id string `json:"spdx_id"`
	Url string `json:"url"`
	Html_url string `json:"html_url,omitempty"`
	Key string `json:"key"`
	Name string `json:"name"`
	Node_id string `json:"node_id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Changes map[string]interface{} `json:"changes"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"` // The `event_type` that was specified in the `POST /repos/{owner}/{repo}/dispatches` request body.
	Branch string `json:"branch"`
	Client_payload map[string]interface{} `json:"client_payload"` // The `client_payload` that was specified in the `POST /repos/{owner}/{repo}/dispatches` request body.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Allow_auto_merge bool `json:"allow_auto_merge,omitempty"`
	Collaborators_url string `json:"collaborators_url"`
	Statuses_url string `json:"statuses_url"`
	Has_pages bool `json:"has_pages"`
	Releases_url string `json:"releases_url"`
	Assignees_url string `json:"assignees_url"`
	Merges_url string `json:"merges_url"`
	Teams_url string `json:"teams_url"`
	Issue_comment_url string `json:"issue_comment_url"`
	Commits_url string `json:"commits_url"`
	Has_discussions bool `json:"has_discussions,omitempty"`
	Watchers int `json:"watchers"`
	Branches_url string `json:"branches_url"`
	Tags_url string `json:"tags_url"`
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"`
	Hooks_url string `json:"hooks_url"`
	Id int `json:"id"`
	Visibility string `json:"visibility,omitempty"` // The repository visibility: public, private, or internal.
	Allow_merge_commit bool `json:"allow_merge_commit,omitempty"`
	Default_branch string `json:"default_branch"`
	Topics []string `json:"topics,omitempty"`
	Blobs_url string `json:"blobs_url"`
	Has_projects bool `json:"has_projects"`
	Description string `json:"description"`
	Open_issues int `json:"open_issues"`
	Mirror_url string `json:"mirror_url"`
	Homepage string `json:"homepage"`
	Open_issues_count int `json:"open_issues_count"`
	Node_id string `json:"node_id"`
	Git_url string `json:"git_url"`
	Is_template bool `json:"is_template,omitempty"`
	Allow_rebase_merge bool `json:"allow_rebase_merge,omitempty"`
	Language string `json:"language"`
	Downloads_url string `json:"downloads_url"`
	Ssh_url string `json:"ssh_url"`
	Fork bool `json:"fork"`
	Forks int `json:"forks"`
	Issues_url string `json:"issues_url"`
	Name string `json:"name"`
	Comments_url string `json:"comments_url"`
	Allow_forking bool `json:"allow_forking,omitempty"`
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"`
	Archive_url string `json:"archive_url"`
	Clone_url string `json:"clone_url"`
	Pushed_at string `json:"pushed_at"`
	Languages_url string `json:"languages_url"`
	Keys_url string `json:"keys_url"`
	Updated_at string `json:"updated_at"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	License GeneratedType `json:"license"` // License Simple
	Pulls_url string `json:"pulls_url"`
	Size int `json:"size"`
	Created_at string `json:"created_at"`
	Watchers_count int `json:"watchers_count"`
	Deployments_url string `json:"deployments_url"`
	Notifications_url string `json:"notifications_url"`
	Subscription_url string `json:"subscription_url"`
	Svn_url string `json:"svn_url"`
	Text_matches []map[string]interface{} `json:"text_matches,omitempty"`
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Stargazers_url string `json:"stargazers_url"`
	Git_refs_url string `json:"git_refs_url"`
	Disabled bool `json:"disabled"` // Returns whether or not this repository disabled.
	Has_downloads bool `json:"has_downloads"`
	Subscribers_url string `json:"subscribers_url"`
	Forks_url string `json:"forks_url"`
	Score float64 `json:"score"`
	Contributors_url string `json:"contributors_url"`
	Git_tags_url string `json:"git_tags_url"`
	Full_name string `json:"full_name"`
	Has_issues bool `json:"has_issues"`
	Has_wiki bool `json:"has_wiki"`
	Forks_count int `json:"forks_count"`
	Git_commits_url string `json:"git_commits_url"`
	Stargazers_count int `json:"stargazers_count"`
	Temp_clone_token string `json:"temp_clone_token,omitempty"`
	Trees_url string `json:"trees_url"`
	Html_url string `json:"html_url"`
	Issue_events_url string `json:"issue_events_url"`
	Labels_url string `json:"labels_url"`
	Milestones_url string `json:"milestones_url"`
	Private bool `json:"private"`
	Archived bool `json:"archived"`
	Contents_url string `json:"contents_url"`
	Allow_squash_merge bool `json:"allow_squash_merge,omitempty"`
	Compare_url string `json:"compare_url"`
	Url string `json:"url"`
	Events_url string `json:"events_url"`
	Master_branch string `json:"master_branch,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Total_engaged_users int `json:"total_engaged_users,omitempty"` // Number of users who accepted at least one Copilot code suggestion, across all active editors. Includes both full and partial acceptances.
	Editors []map[string]interface{} `json:"editors,omitempty"`
	Languages []map[string]interface{} `json:"languages,omitempty"` // Code completion metrics for active languages.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id int `json:"id"` // The ID of the installation.
	Node_id string `json:"node_id"` // The global node ID of the installation.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Property_name string `json:"property_name"` // The name of the property
	Value string `json:"value"` // The value assigned to the property
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Zen string `json:"zen,omitempty"` // Random string of GitHub zen.
	Hook map[string]interface{} `json:"hook,omitempty"` // The webhook that is being pinged
	Hook_id int `json:"hook_id,omitempty"` // The ID of the webhook that triggered the ping.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Base_branch string `json:"base_branch,omitempty"`
	Merge_type string `json:"merge_type,omitempty"`
	Message string `json:"message,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Length int `json:"length,omitempty"` // The length of the IP prefix.
	Prefix string `json:"prefix,omitempty"` // The prefix for the public IP.
	Enabled bool `json:"enabled,omitempty"` // Whether public IP is enabled.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization"` // A GitHub user.
	Permissions []string `json:"permissions"` // A list of permissions included in this role.
	Updated_at string `json:"updated_at"` // The date and time the role was last updated.
	Base_role string `json:"base_role,omitempty"` // The system role from which this role inherits permissions.
	Description string `json:"description,omitempty"` // A short description about who this role is for or what permissions it grants.
	Source string `json:"source,omitempty"` // Source answers the question, "where did this role come from?"
	Name string `json:"name"` // The name of the role.
	Created_at string `json:"created_at"` // The date and time the role was created.
	Id int64 `json:"id"` // The unique identifier of the role.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Permissions map[string]interface{} `json:"permissions"` // Permissions requested, categorized by type of permission.
	Token_expires_at string `json:"token_expires_at"` // Date and time when the associated fine-grained personal access token expires.
	Token_last_used_at string `json:"token_last_used_at"` // Date and time when the associated fine-grained personal access token was last used for authentication.
	Token_name string `json:"token_name"` // The name given to the user's token. This field can also be found in an organization's settings page for Active Tokens.
	Id int `json:"id"` // Unique identifier of the fine-grained personal access token grant. The `pat_id` used to get details about an approved fine-grained personal access token.
	Access_granted_at string `json:"access_granted_at"` // Date and time when the fine-grained personal access token was approved to access the organization.
	Repositories_url string `json:"repositories_url"` // URL to the list of repositories the fine-grained personal access token can access. Only follow when `repository_selection` is `subset`.
	Token_id int `json:"token_id"` // Unique identifier of the user's token. This field can also be found in audit log events and the organization's settings for their PAT grants.
	Repository_selection string `json:"repository_selection"` // Type of repository selection requested.
	Token_expired bool `json:"token_expired"` // Whether the associated fine-grained personal access token has expired.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repositories []map[string]interface{} `json:"repositories,omitempty"` // Repositories in which users used Copilot for Pull Requests to generate pull request summaries
	Total_engaged_users int `json:"total_engaged_users,omitempty"` // The number of users who used Copilot for Pull Requests on github.com to generate a pull request summary at least once.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id int64 `json:"id"`
	Number int `json:"number"`
	Url string `json:"url"`
	Base map[string]interface{} `json:"base"`
	Head map[string]interface{} `json:"head"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request GeneratedType `json:"pull_request"`
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// Webhooksreviewcomment represents the Webhooksreviewcomment schema from the OpenAPI specification
type Webhooksreviewcomment struct {
	Html_url string `json:"html_url"` // HTML URL for the pull request review comment.
	Created_at string `json:"created_at"`
	Diff_hunk string `json:"diff_hunk"` // The diff of the line that the comment refers to.
	Original_commit_id string `json:"original_commit_id"` // The SHA of the original commit to which the comment applies.
	Body string `json:"body"` // The text of the comment.
	In_reply_to_id int `json:"in_reply_to_id,omitempty"` // The comment ID to reply to.
	_links map[string]interface{} `json:"_links"`
	Pull_request_review_id int `json:"pull_request_review_id"` // The ID of the pull request review to which the comment belongs.
	Subject_type string `json:"subject_type,omitempty"` // The level at which the comment is targeted, can be a diff line or a file.
	Position int `json:"position"` // The line index in the diff to which the comment applies.
	Original_position int `json:"original_position"` // The index of the original line in the diff to which the comment applies.
	Start_line int `json:"start_line"` // The first line of the range for a multi-line comment.
	Reactions map[string]interface{} `json:"reactions"`
	Side string `json:"side"` // The side of the first line of the range for a multi-line comment.
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Updated_at string `json:"updated_at"`
	Url string `json:"url"` // URL for the pull request review comment
	Original_start_line int `json:"original_start_line"` // The first line of the range for a multi-line comment.
	Pull_request_url string `json:"pull_request_url"` // URL for the pull request that the review comment belongs to.
	Id int `json:"id"` // The ID of the pull request review comment.
	Path string `json:"path"` // The relative path of the file to which the comment applies.
	Start_side string `json:"start_side"` // The side of the first line of the range for a multi-line comment.
	Commit_id string `json:"commit_id"` // The SHA of the commit to which the comment applies.
	Original_line int `json:"original_line"` // The line of the blob to which the comment applies. The last line of the range for a multi-line comment
	Line int `json:"line"` // The line of the blob to which the comment applies. The last line of the range for a multi-line comment
	Node_id string `json:"node_id"` // The node ID of the pull request review comment.
	User map[string]interface{} `json:"user"`
}

// Dependency represents the Dependency schema from the OpenAPI specification
type Dependency struct {
	Package_url string `json:"package_url,omitempty"` // Package-url (PURL) of dependency. See https://github.com/package-url/purl-spec for more details.
	Relationship string `json:"relationship,omitempty"` // A notation of whether a dependency is requested directly by this manifest or is a dependency of another dependency.
	Scope string `json:"scope,omitempty"` // A notation of whether the dependency is required for the primary build artifact (runtime) or is only used for development. Future versions of this specification may allow for more granular scopes.
	Dependencies []string `json:"dependencies,omitempty"` // Array of package-url (PURLs) of direct child dependencies.
	Metadata Metadata `json:"metadata,omitempty"` // User-defined metadata to store domain-specific information limited to 8 keys with scalar values.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Security_advisory map[string]interface{} `json:"security_advisory"` // The details of the security advisory, including summary, description, and severity.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// Hovercard represents the Hovercard schema from the OpenAPI specification
type Hovercard struct {
	Contexts []map[string]interface{} `json:"contexts"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Documentation_url string `json:"documentation_url,omitempty"`
	Message string `json:"message,omitempty"`
	Schemas []string `json:"schemas,omitempty"`
	ScimType string `json:"scimType,omitempty"`
	Status int `json:"status,omitempty"`
	Detail string `json:"detail,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Commit_id string `json:"commit_id"`
	Side string `json:"side,omitempty"` // The side of the first line of the range for a multi-line comment.
	Subject_type string `json:"subject_type,omitempty"` // The level at which the comment is targeted, can be a diff line or a file.
	Start_side string `json:"start_side,omitempty"` // The side of the first line of the range for a multi-line comment.
	Position int `json:"position"`
	In_reply_to_id int `json:"in_reply_to_id,omitempty"`
	Original_position int `json:"original_position"`
	Node_id string `json:"node_id"`
	Line int `json:"line,omitempty"` // The line of the blob to which the comment applies. The last line of the range for a multi-line comment
	Path string `json:"path"`
	Id int64 `json:"id"`
	Body_html string `json:"body_html,omitempty"`
	Original_start_line int `json:"original_start_line,omitempty"` // The original first line of the range for a multi-line comment.
	Updated_at string `json:"updated_at"`
	User GeneratedType `json:"user"` // A GitHub user.
	Original_line int `json:"original_line,omitempty"` // The original line of the blob to which the comment applies. The last line of the range for a multi-line comment
	Body string `json:"body"`
	Body_text string `json:"body_text,omitempty"`
	Created_at string `json:"created_at"`
	Pull_request_review_id int64 `json:"pull_request_review_id"`
	Reactions GeneratedType `json:"reactions,omitempty"`
	_links map[string]interface{} `json:"_links"`
	Original_commit_id string `json:"original_commit_id"`
	Pull_request_url string `json:"pull_request_url"`
	Diff_hunk string `json:"diff_hunk"`
	Start_line int `json:"start_line,omitempty"` // The first line of the range for a multi-line comment.
	Url string `json:"url"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Html_url string `json:"html_url"`
}

// Commit represents the Commit schema from the OpenAPI specification
type Commit struct {
	Sha string `json:"sha"`
	Comments_url string `json:"comments_url"`
	Node_id string `json:"node_id"`
	Stats map[string]interface{} `json:"stats,omitempty"`
	Committer interface{} `json:"committer"`
	Url string `json:"url"`
	Author interface{} `json:"author"`
	Commit map[string]interface{} `json:"commit"`
	Files []GeneratedType `json:"files,omitempty"`
	Html_url string `json:"html_url"`
	Parents []map[string]interface{} `json:"parents"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Created_at string `json:"created_at,omitempty"`
	Enabled bool `json:"enabled,omitempty"`
	Id int `json:"id,omitempty"`
	Pattern string `json:"pattern"`
	Updated_at string `json:"updated_at,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enabled bool `json:"enabled"` // Whether Dependabot security updates are enabled for the repository.
	Paused bool `json:"paused"` // Whether Dependabot security updates are paused for the repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Uniques int `json:"uniques"`
	Clones []Traffic `json:"clones"`
	Count int `json:"count"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repo Repository `json:"repo"` // A repository on GitHub.
	Starred_at string `json:"starred_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Branch string `json:"branch,omitempty"` // Name of the exported branch
	Completed_at string `json:"completed_at,omitempty"` // Completion time of the last export operation
	Export_url string `json:"export_url,omitempty"` // Url for fetching export details
	Html_url string `json:"html_url,omitempty"` // Web url for the exported branch
	Id string `json:"id,omitempty"` // Id for the export details
	Sha string `json:"sha,omitempty"` // Git commit SHA of the exported branch
	State string `json:"state,omitempty"` // State of the latest export
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Personal_access_token_request GeneratedType `json:"personal_access_token_request"` // Details of a Personal Access Token Request.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Type string `json:"type"`
	Parameters map[string]interface{} `json:"parameters,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Changes Webhooksprojectchanges `json:"changes"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Projects_v2_item GeneratedType `json:"projects_v2_item"` // An item belonging to a project
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Created_at string `json:"created_at"`
	Name string `json:"name"` // The name of the secret.
	Updated_at string `json:"updated_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Rule Webhooksrule `json:"rule"` // The branch protection rule. Includes a `name` and all the [branch protection settings](https://docs.github.com/github/administering-a-repository/defining-the-mergeability-of-pull-requests/about-protected-branches#about-branch-protection-settings) applied to branches that match the name. Binary settings are boolean. Multi-level configurations are one of `off`, `non_admins`, or `everyone`. Actor and build lists are arrays of strings.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id int `json:"id"`
	Login string `json:"login"`
	Avatar_url string `json:"avatar_url"`
	Html_url string `json:"html_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Node_id string `json:"node_id,omitempty"`
	Requester GeneratedType `json:"requester"` // A GitHub user.
	Account interface{} `json:"account"`
	Created_at string `json:"created_at"`
	Id int `json:"id"` // Unique identifier of the request installation.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Started_at string `json:"started_at"` // The start time of an autofix in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Status string `json:"status"` // The status of an autofix.
	Description string `json:"description"` // The description of an autofix.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Node_id string `json:"node_id"`
	Url string `json:"url"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Event string `json:"event"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Commit_id string `json:"commit_id"`
	Commit_url string `json:"commit_url"`
	Id int `json:"id"`
	Project_card map[string]interface{} `json:"project_card,omitempty"`
	Created_at string `json:"created_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Security_advisory Webhookssecurityadvisory `json:"security_advisory"` // The details of the security advisory, including summary, description, and severity.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Label map[string]interface{} `json:"label"`
	Node_id string `json:"node_id"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Commit_url string `json:"commit_url"`
	Commit_id string `json:"commit_id"`
	Created_at string `json:"created_at"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Url string `json:"url"`
	Event string `json:"event"`
	Id int `json:"id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Deleted bool `json:"deleted"` // Whether this push deleted the `ref`.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Base_ref string `json:"base_ref"`
	Repository map[string]interface{} `json:"repository"` // A git repository
	After string `json:"after"` // The SHA of the most recent commit on `ref` after the push.
	Commits []map[string]interface{} `json:"commits"` // An array of commit objects describing the pushed commits. (Pushed commits are all commits that are included in the `compare` between the `before` commit and the `after` commit.) The array includes a maximum of 2048 commits. If necessary, you can use the [Commits API](https://docs.github.com/rest/commits) to fetch additional commits.
	Compare string `json:"compare"` // URL that shows the changes in this `ref` update, from the `before` commit to the `after` commit. For a newly created `ref` that is directly based on the default branch, this is the comparison between the head of the default branch and the `after` commit. Otherwise, this shows all commits until the `after` commit.
	Forced bool `json:"forced"` // Whether this push was a force push of the `ref`.
	Head_commit map[string]interface{} `json:"head_commit"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Pusher map[string]interface{} `json:"pusher"` // Metaproperties for Git author/committer information.
	Ref string `json:"ref"` // The full git ref that was pushed. Example: `refs/heads/main` or `refs/tags/v3.14.1`.
	Created bool `json:"created"` // Whether this push created the `ref`.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Before string `json:"before"` // The SHA of the most recent commit on `ref` before the push.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Configuration GeneratedType `json:"configuration,omitempty"` // A code security configuration
	Status string `json:"status,omitempty"` // The attachment status of the code security configuration on the repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Answer Webhooksanswer `json:"answer"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Payload string `json:"payload"` // A URL-encoded string of the check_run.rerequested JSON payload. The decoded payload is a JSON object.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Health_percentage int `json:"health_percentage"`
	Updated_at string `json:"updated_at"`
	Content_reports_enabled bool `json:"content_reports_enabled,omitempty"`
	Description string `json:"description"`
	Documentation string `json:"documentation"`
	Files map[string]interface{} `json:"files"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Group_name string `json:"group_name,omitempty"` // If a group assignment, name of the group the student is in
	Assignment_name string `json:"assignment_name"` // Name of the assignment
	Github_username string `json:"github_username"` // GitHub username of the student
	Points_available int `json:"points_available"` // Number of points available for the assignment
	Starter_code_url string `json:"starter_code_url"` // URL of the starter code for the assignment
	Student_repository_name string `json:"student_repository_name"` // Name of the student's assignment repository
	Submission_timestamp string `json:"submission_timestamp"` // Timestamp of the student's assignment submission
	Points_awarded int `json:"points_awarded"` // Number of points awarded to the student
	Roster_identifier string `json:"roster_identifier"` // Roster identifier of the student
	Student_repository_url string `json:"student_repository_url"` // URL of the student's assignment repository
	Assignment_url string `json:"assignment_url"` // URL of the assignment
}

// Status represents the Status schema from the OpenAPI specification
type Status struct {
	Avatar_url string `json:"avatar_url"`
	State string `json:"state"`
	Target_url string `json:"target_url"`
	Created_at string `json:"created_at"`
	Description string `json:"description"`
	Id int `json:"id"`
	Context string `json:"context"`
	Creator GeneratedType `json:"creator"` // A GitHub user.
	Node_id string `json:"node_id"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Node_id string `json:"node_id"` // The GraphQL identifier of the repository.
	Private bool `json:"private"` // Whether the repository is private.
	Default_branch string `json:"default_branch"` // The default branch for the repository.
	Full_name string `json:"full_name"` // The full, globally unique name of the repository.
	Html_url string `json:"html_url"` // The URL to view the repository on GitHub.com.
	Id int `json:"id"` // A unique identifier of the repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Alert GeneratedType `json:"alert"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Body string `json:"body"`
	Commit_id string `json:"commit_id"`
	Position int `json:"position"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	Html_url string `json:"html_url"`
	Node_id string `json:"node_id"`
	Reactions GeneratedType `json:"reactions,omitempty"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Created_at string `json:"created_at"`
	User GeneratedType `json:"user"` // A GitHub user.
	Id int `json:"id"`
	Line int `json:"line"`
	Path string `json:"path"`
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

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Secret string `json:"secret,omitempty"` // If provided, the `secret` will be used as the `key` to generate the HMAC hex digest value for [delivery signature headers](https://docs.github.com/webhooks/event-payloads/#delivery-headers).
	Url string `json:"url,omitempty"` // The URL to which the payloads will be delivered.
	Content_type string `json:"content_type,omitempty"` // The media type used to serialize the payloads. Supported values include `json` and `form`. The default is `form`.
	Insecure_ssl GeneratedType `json:"insecure_ssl,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Html_url string `json:"html_url"`
	Events_url string `json:"events_url"`
	Following_url string `json:"following_url"`
	Email string `json:"email,omitempty"`
	Gravatar_id string `json:"gravatar_id"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Url string `json:"url"`
	Repos_url string `json:"repos_url"`
	Role_name string `json:"role_name"`
	Starred_url string `json:"starred_url"`
	Name string `json:"name,omitempty"`
	Received_events_url string `json:"received_events_url"`
	Site_admin bool `json:"site_admin"`
	Type string `json:"type"`
	Followers_url string `json:"followers_url"`
	Id int64 `json:"id"`
	Organizations_url string `json:"organizations_url"`
	Subscriptions_url string `json:"subscriptions_url"`
	Login string `json:"login"`
	Node_id string `json:"node_id"`
	Avatar_url string `json:"avatar_url"`
	Gists_url string `json:"gists_url"`
	User_view_type string `json:"user_view_type,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Type string `json:"type"`
}

// Webhooksalert represents the Webhooksalert schema from the OpenAPI specification
type Webhooksalert struct {
	Fixed_in string `json:"fixed_in,omitempty"`
	Node_id string `json:"node_id"`
	External_reference string `json:"external_reference"`
	Number int `json:"number"`
	Affected_range string `json:"affected_range"`
	Dismisser map[string]interface{} `json:"dismisser,omitempty"`
	Severity string `json:"severity"`
	State string `json:"state"`
	External_identifier string `json:"external_identifier"`
	Id int `json:"id"`
	Created_at string `json:"created_at"`
	Ghsa_id string `json:"ghsa_id"`
	Dismissed_at string `json:"dismissed_at,omitempty"`
	Affected_package_name string `json:"affected_package_name"`
	Fixed_at string `json:"fixed_at,omitempty"`
	Dismiss_reason string `json:"dismiss_reason,omitempty"`
	Fix_reason string `json:"fix_reason,omitempty"`
}

// Project represents the Project schema from the OpenAPI specification
type Project struct {
	Body string `json:"body"` // Body of the project
	Id int `json:"id"`
	Owner_url string `json:"owner_url"`
	Updated_at string `json:"updated_at"`
	Creator GeneratedType `json:"creator"` // A GitHub user.
	Html_url string `json:"html_url"`
	Columns_url string `json:"columns_url"`
	Name string `json:"name"` // Name of the project
	State string `json:"state"` // State of the project; either 'open' or 'closed'
	Created_at string `json:"created_at"`
	Number int `json:"number"`
	Organization_permission string `json:"organization_permission,omitempty"` // The baseline permission that all organization members have on this project. Only present if owner is an organization.
	Private bool `json:"private,omitempty"` // Whether or not this project can be seen by everyone. Only present if owner is an organization.
	Node_id string `json:"node_id"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id string `json:"id"` // SHA for the commit
	Message string `json:"message"` // Message describing the purpose of the commit
	Timestamp string `json:"timestamp"` // Timestamp of the commit
	Tree_id string `json:"tree_id"` // SHA for the commit's tree
	Author map[string]interface{} `json:"author"` // Information about the Git author
	Committer map[string]interface{} `json:"committer"` // Information about the Git committer
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Projects_v2_item GeneratedType `json:"projects_v2_item"` // An item belonging to a project
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes,omitempty"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Pull_request_body_url string `json:"pull_request_body_url"` // The API URL to get the pull request where the secret was detected.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Key string `json:"key"` // The Base64 encoded public key.
	Key_id string `json:"key_id"` // The identifier for the key.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	State string `json:"state"` // Whether deployment to the environment(s) was approved or rejected or pending (with comments)
	User GeneratedType `json:"user"` // A GitHub user.
	Comment string `json:"comment"` // The comment submitted with the deployment review
	Environments []map[string]interface{} `json:"environments"` // The list of environments that were approved or rejected
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pusher_type string `json:"pusher_type"` // The pusher type for the event. Can be either `user` or a deploy key.
	Ref string `json:"ref"` // The [`git ref`](https://docs.github.com/rest/git/refs#get-a-reference) resource.
	Ref_type string `json:"ref_type"` // The type of Git ref object deleted in the repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Path string `json:"path"`
	Ref string `json:"ref,omitempty"`
	Sha string `json:"sha"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Title string `json:"title"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	End_column float64 `json:"end_column"` // The column at which the secret ends within the end line when the file is interpreted as 8BIT ASCII
	Path string `json:"path"` // The file path in the repository
	Blob_sha string `json:"blob_sha"` // SHA-1 hash ID of the associated blob
	Commit_sha string `json:"commit_sha"` // SHA-1 hash ID of the associated commit
	Commit_url string `json:"commit_url"` // The API URL to get the associated commit resource
	Start_line float64 `json:"start_line"` // Line number at which the secret starts in the file
	Blob_url string `json:"blob_url"` // The API URL to get the associated blob resource
	End_line float64 `json:"end_line"` // Line number at which the secret ends in the file
	Start_column float64 `json:"start_column"` // The column at which the secret starts within the start line when the file is interpreted as 8BIT ASCII
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Workflow_job map[string]interface{} `json:"workflow_job"`
	Action string `json:"action"`
	Deployment Deployment `json:"deployment,omitempty"` // A request for a specific ref(branch,sha,tag) to be deployed
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Confirm_delete_url string `json:"confirm_delete_url"` // Next deletable analysis in chain, with last analysis deletion confirmation
	Next_analysis_url string `json:"next_analysis_url"` // Next deletable analysis in chain, without last analysis deletion confirmation
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Name string `json:"name"` // The name of the network configuration.
	Network_settings_ids []string `json:"network_settings_ids,omitempty"` // The unique identifier of each network settings in the configuration.
	Compute_service string `json:"compute_service,omitempty"` // The hosted compute service the network configuration supports.
	Created_on string `json:"created_on"` // The time at which the network configuration was created, in ISO 8601 format.
	Id string `json:"id"` // The unique identifier of the network configuration.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Label Webhookslabel `json:"label"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Created_at string `json:"created_at"` // The date and time the campaign was created, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Number int `json:"number"` // The number of the newly created campaign
	Ends_at string `json:"ends_at"` // The date and time the campaign has ended, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Name string `json:"name,omitempty"` // The campaign name
	Updated_at string `json:"updated_at"` // The date and time the campaign was last updated, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Alert_stats map[string]interface{} `json:"alert_stats,omitempty"`
	Contact_link string `json:"contact_link"` // The contact link of the campaign.
	Description string `json:"description"` // The campaign description
	Published_at string `json:"published_at,omitempty"` // The date and time the campaign was published, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	State string `json:"state"` // Indicates whether a campaign is open or closed
	Managers []GeneratedType `json:"managers"` // The campaign managers
	Team_managers []Team `json:"team_managers,omitempty"` // The campaign team managers
	Closed_at string `json:"closed_at,omitempty"` // The date and time the campaign was closed, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ. Will be null if the campaign is still open.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Membership Webhooksmembership `json:"membership,omitempty"` // The membership between the user and the organization. Not present when the action is `member_invited`.
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id int `json:"id"`
	Import_url string `json:"import_url"`
	Name string `json:"name"`
	Remote_id string `json:"remote_id"`
	Remote_name string `json:"remote_name"`
	Url string `json:"url"`
	Email string `json:"email"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Project Webhooksproject `json:"project"`
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// Webhookspreviousmarketplacepurchase represents the Webhookspreviousmarketplacepurchase schema from the OpenAPI specification
type Webhookspreviousmarketplacepurchase struct {
	On_free_trial bool `json:"on_free_trial"`
	Plan map[string]interface{} `json:"plan"`
	Unit_count int `json:"unit_count"`
	Account map[string]interface{} `json:"account"`
	Billing_cycle string `json:"billing_cycle"`
	Free_trial_ends_on interface{} `json:"free_trial_ends_on"`
	Next_billing_date string `json:"next_billing_date,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id int `json:"id,omitempty"` // Unique identifier of the label.
	Name string `json:"name"` // Name of the label.
	Type string `json:"type,omitempty"` // The type of label. Read-only labels are applied automatically when the runner is configured.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Effective_date string `json:"effective_date"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Marketplace_purchase map[string]interface{} `json:"marketplace_purchase"`
	Previous_marketplace_purchase Webhookspreviousmarketplacepurchase `json:"previous_marketplace_purchase,omitempty"`
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Action string `json:"action"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Pushed_at string `json:"pushed_at"`
	Mirror_url string `json:"mirror_url"`
	Name string `json:"name"` // The name of the repository.
	Issue_events_url string `json:"issue_events_url"`
	Releases_url string `json:"releases_url"`
	Allow_merge_commit bool `json:"allow_merge_commit,omitempty"` // Whether to allow merge commits for pull requests.
	Archive_url string `json:"archive_url"`
	Merges_url string `json:"merges_url"`
	Squash_merge_commit_message string `json:"squash_merge_commit_message,omitempty"` // The default value for a squash merge commit message: - `PR_BODY` - default to the pull request's body. - `COMMIT_MESSAGES` - default to the branch's commit messages. - `BLANK` - default to a blank commit message.
	Notifications_url string `json:"notifications_url"`
	Merge_commit_title string `json:"merge_commit_title,omitempty"` // The default value for a merge commit title. - `PR_TITLE` - default to the pull request's title. - `MERGE_MESSAGE` - default to the classic title for a merge message (e.g., Merge pull request #123 from branch-name).
	Squash_merge_commit_title string `json:"squash_merge_commit_title,omitempty"` // The default value for a squash merge commit title: - `PR_TITLE` - default to the pull request's title. - `COMMIT_OR_PR_TITLE` - default to the commit's title (if only one commit) or the pull request's title (when more than one commit).
	Use_squash_pr_title_as_default bool `json:"use_squash_pr_title_as_default,omitempty"` // Whether a squash merge commit can use the pull request title as default. **This property is closing down. Please use `squash_merge_commit_title` instead.
	Starred_at string `json:"starred_at,omitempty"`
	Milestones_url string `json:"milestones_url"`
	Allow_forking bool `json:"allow_forking,omitempty"` // Whether to allow forking this repo
	Ssh_url string `json:"ssh_url"`
	Collaborators_url string `json:"collaborators_url"`
	Id int64 `json:"id"` // Unique identifier of the repository
	Watchers_count int `json:"watchers_count"`
	Disabled bool `json:"disabled"` // Returns whether or not this repository disabled.
	Hooks_url string `json:"hooks_url"`
	Trees_url string `json:"trees_url"`
	Tags_url string `json:"tags_url"`
	Forks int `json:"forks"`
	Teams_url string `json:"teams_url"`
	Allow_update_branch bool `json:"allow_update_branch,omitempty"` // Whether or not a pull request head branch that is behind its base branch can always be updated even if it is not required to be up to date before merging.
	Contributors_url string `json:"contributors_url"`
	Private bool `json:"private"` // Whether the repository is private or public.
	Blobs_url string `json:"blobs_url"`
	Has_projects bool `json:"has_projects"` // Whether projects are enabled.
	Keys_url string `json:"keys_url"`
	Allow_auto_merge bool `json:"allow_auto_merge,omitempty"` // Whether to allow Auto-merge to be used on pull requests.
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"` // Whether to delete head branches when pull requests are merged
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Git_refs_url string `json:"git_refs_url"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Has_pages bool `json:"has_pages"`
	Statuses_url string `json:"statuses_url"`
	Has_issues bool `json:"has_issues"` // Whether issues are enabled.
	Is_template bool `json:"is_template,omitempty"` // Whether this repository acts as a template that can be used to generate new repositories.
	Url string `json:"url"`
	License GeneratedType `json:"license"` // License Simple
	Visibility string `json:"visibility,omitempty"` // The repository visibility: public, private, or internal.
	Forks_url string `json:"forks_url"`
	Size int `json:"size"` // The size of the repository, in kilobytes. Size is calculated hourly. When a repository is initially created, the size is 0.
	Has_wiki bool `json:"has_wiki"` // Whether the wiki is enabled.
	Issues_url string `json:"issues_url"`
	Deployments_url string `json:"deployments_url"`
	Forks_count int `json:"forks_count"`
	Commits_url string `json:"commits_url"`
	Description string `json:"description"`
	Archived bool `json:"archived"` // Whether the repository is archived.
	Html_url string `json:"html_url"`
	Topics []string `json:"topics,omitempty"`
	Stargazers_url string `json:"stargazers_url"`
	Comments_url string `json:"comments_url"`
	Homepage string `json:"homepage"`
	Master_branch string `json:"master_branch,omitempty"`
	Assignees_url string `json:"assignees_url"`
	Events_url string `json:"events_url"`
	Languages_url string `json:"languages_url"`
	Has_downloads bool `json:"has_downloads"` // Whether downloads are enabled.
	Has_discussions bool `json:"has_discussions,omitempty"` // Whether discussions are enabled.
	Code_search_index_status map[string]interface{} `json:"code_search_index_status,omitempty"` // The status of the code search index for this repository
	Issue_comment_url string `json:"issue_comment_url"`
	Subscription_url string `json:"subscription_url"`
	Open_issues int `json:"open_issues"`
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"` // Whether to require contributors to sign off on web-based commits
	Subscribers_url string `json:"subscribers_url"`
	Allow_squash_merge bool `json:"allow_squash_merge,omitempty"` // Whether to allow squash merges for pull requests.
	Pulls_url string `json:"pulls_url"`
	Watchers int `json:"watchers"`
	Default_branch string `json:"default_branch"` // The default branch of the repository.
	Svn_url string `json:"svn_url"`
	Updated_at string `json:"updated_at"`
	Downloads_url string `json:"downloads_url"`
	Contents_url string `json:"contents_url"`
	Merge_commit_message string `json:"merge_commit_message,omitempty"` // The default value for a merge commit message. - `PR_TITLE` - default to the pull request's title. - `PR_BODY` - default to the pull request's body. - `BLANK` - default to a blank commit message.
	Anonymous_access_enabled bool `json:"anonymous_access_enabled,omitempty"` // Whether anonymous git access is enabled for this repository
	Fork bool `json:"fork"`
	Compare_url string `json:"compare_url"`
	Clone_url string `json:"clone_url"`
	Language string `json:"language"`
	Temp_clone_token string `json:"temp_clone_token,omitempty"`
	Git_tags_url string `json:"git_tags_url"`
	Labels_url string `json:"labels_url"`
	Branches_url string `json:"branches_url"`
	Created_at string `json:"created_at"`
	Stargazers_count int `json:"stargazers_count"`
	Full_name string `json:"full_name"`
	Git_commits_url string `json:"git_commits_url"`
	Git_url string `json:"git_url"`
	Node_id string `json:"node_id"`
	Open_issues_count int `json:"open_issues_count"`
	Allow_rebase_merge bool `json:"allow_rebase_merge,omitempty"` // Whether to allow rebase merges for pull requests.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Alert Webhooksalert `json:"alert"` // The security alert of the vulnerable dependency.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	After_sha string `json:"after_sha,omitempty"` // The last commit sha in the push evaluation.
	Repository_name string `json:"repository_name,omitempty"` // The name of the repository without the `.git` extension.
	Result string `json:"result,omitempty"` // The result of the rule evaluations for rules with the `active` enforcement status.
	Before_sha string `json:"before_sha,omitempty"` // The first commit sha before the push evaluation.
	Id int `json:"id,omitempty"` // The unique identifier of the rule insight.
	Pushed_at string `json:"pushed_at,omitempty"`
	Repository_id int `json:"repository_id,omitempty"` // The ID of the repository associated with the rule evaluation.
	Rule_evaluations []map[string]interface{} `json:"rule_evaluations,omitempty"` // Details on the evaluated rules.
	Actor_id int `json:"actor_id,omitempty"` // The number that identifies the user.
	Actor_name string `json:"actor_name,omitempty"` // The handle for the GitHub user account.
	Evaluation_result string `json:"evaluation_result,omitempty"` // The result of the rule evaluations for rules with the `active` and `evaluate` enforcement statuses, demonstrating whether rules would pass or fail if all rules in the rule suite were `active`. Null if no rules with `evaluate` enforcement status were run.
	Ref string `json:"ref,omitempty"` // The ref name that the evaluation ran on.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Milestone Webhooksmilestone3 `json:"milestone"` // A collection of related issues and pull requests.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Labels_url string `json:"labels_url"`
	Node_id string `json:"node_id"`
	Url string `json:"url"`
	Due_on string `json:"due_on"`
	Id int `json:"id"`
	Closed_issues int `json:"closed_issues"`
	Description string `json:"description"`
	Updated_at string `json:"updated_at"`
	Closed_at string `json:"closed_at"`
	Number int `json:"number"` // The number of the milestone.
	Created_at string `json:"created_at"`
	Title string `json:"title"` // The title of the milestone.
	Creator GeneratedType `json:"creator"` // A GitHub user.
	Html_url string `json:"html_url"`
	Open_issues int `json:"open_issues"`
	State string `json:"state"` // The state of the milestone.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Display_name string `json:"display_name"` // Display name for this image.
	Id string `json:"id"` // The ID of the image. Use this ID for the `image` parameter when creating a new larger runner.
	Size_gb int `json:"size_gb"` // Image size in GB.
	Source string `json:"source"` // The image provider.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Created_at string `json:"created_at"`
	Url string `json:"url"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Commit_id string `json:"commit_id"`
	Node_id string `json:"node_id"`
	Commit_url string `json:"commit_url"`
	Event string `json:"event"`
	Id int `json:"id"`
	State_reason string `json:"state_reason,omitempty"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
}

// Environment represents the Environment schema from the OpenAPI specification
type Environment struct {
	Id int64 `json:"id"` // The id of the environment.
	Updated_at string `json:"updated_at"` // The time that the environment was last updated, in ISO 8601 format.
	Name string `json:"name"` // The name of the environment.
	Protection_rules []interface{} `json:"protection_rules,omitempty"` // Built-in deployment protection rules for the environment.
	Deployment_branch_policy GeneratedType `json:"deployment_branch_policy,omitempty"` // The type of deployment branch policy for this environment. To allow all branches to deploy, set to `null`.
	Created_at string `json:"created_at"` // The time that the environment was created, in ISO 8601 format.
	Html_url string `json:"html_url"`
	Node_id string `json:"node_id"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Personal_access_token_request GeneratedType `json:"personal_access_token_request"` // Details of a Personal Access Token Request.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Reviewer GeneratedType `json:"reviewer"` // A required reviewing team
	File_patterns []string `json:"file_patterns"` // Array of file patterns. Pull requests which change matching files must be approved by the specified team. File patterns use the same syntax as `.gitignore` files.
	Minimum_approvals int `json:"minimum_approvals"` // Minimum number of approvals required from the specified team. If set to zero, the team will be added to the pull request but approval is optional.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Contents_url string `json:"contents_url"`
	Deletions int `json:"deletions"`
	Patch string `json:"patch,omitempty"`
	Previous_filename string `json:"previous_filename,omitempty"`
	Sha string `json:"sha"`
	Blob_url string `json:"blob_url"`
	Changes int `json:"changes"`
	Raw_url string `json:"raw_url"`
	Additions int `json:"additions"`
	Filename string `json:"filename"`
	Status string `json:"status"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Dismissed_by GeneratedType `json:"dismissed_by"` // A GitHub user.
	Fixed_at string `json:"fixed_at"` // The time that the alert was no longer detected and was considered fixed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Security_advisory GeneratedType `json:"security_advisory"` // Details for the GitHub Security Advisory.
	State string `json:"state"` // The state of the Dependabot alert.
	Url string `json:"url"` // The REST API URL of the alert resource.
	Dismissed_comment string `json:"dismissed_comment"` // An optional comment associated with the alert's dismissal.
	Dismissed_at string `json:"dismissed_at"` // The time that the alert was dismissed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Created_at string `json:"created_at"` // The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Html_url string `json:"html_url"` // The GitHub URL of the alert resource.
	Dependency map[string]interface{} `json:"dependency"` // Details for the vulnerable dependency.
	Auto_dismissed_at string `json:"auto_dismissed_at,omitempty"` // The time that the alert was auto-dismissed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Security_vulnerability GeneratedType `json:"security_vulnerability"` // Details pertaining to one vulnerable version range for the advisory.
	Dismissed_reason string `json:"dismissed_reason"` // The reason that the alert was dismissed.
	Number int `json:"number"` // The security alert number.
	Updated_at string `json:"updated_at"` // The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Commit_url string `json:"commit_url"`
	Event string `json:"event"`
	Id int `json:"id"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Node_id string `json:"node_id"`
	Commit_id string `json:"commit_id"`
	Created_at string `json:"created_at"`
	Rename map[string]interface{} `json:"rename"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repositories []GeneratedType `json:"repositories"` // A list of repositories that were skipped. This list may not include all repositories that were skipped. This is only available when the repository was found and the user has access to it.
	Repository_count int `json:"repository_count"` // The total number of repositories that were skipped for this reason.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url"`
	Enabled bool `json:"enabled"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Projects_v2_status_update GeneratedType `json:"projects_v2_status_update"` // An status update belonging to a project
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository_id map[string]interface{} `json:"repository_id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Notifications_url string `json:"notifications_url"`
	Name string `json:"name"` // The name of the repository.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub user.
	Created_at string `json:"created_at"`
	Contributors_url string `json:"contributors_url"`
	Merge_commit_title string `json:"merge_commit_title,omitempty"` // The default value for a merge commit title. - `PR_TITLE` - default to the pull request's title. - `MERGE_MESSAGE` - default to the classic title for a merge message (e.g., Merge pull request #123 from branch-name).
	Has_downloads bool `json:"has_downloads"` // Whether downloads are enabled.
	Comments_url string `json:"comments_url"`
	Stargazers_count int `json:"stargazers_count"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	Allow_squash_merge bool `json:"allow_squash_merge,omitempty"` // Whether to allow squash merges for pull requests.
	Languages_url string `json:"languages_url"`
	Events_url string `json:"events_url"`
	Issue_comment_url string `json:"issue_comment_url"`
	Allow_update_branch bool `json:"allow_update_branch,omitempty"` // Whether or not a pull request head branch that is behind its base branch can always be updated even if it is not required to be up to date before merging.
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"` // Whether to delete head branches when pull requests are merged
	Assignees_url string `json:"assignees_url"`
	Master_branch string `json:"master_branch,omitempty"`
	Git_refs_url string `json:"git_refs_url"`
	Blobs_url string `json:"blobs_url"`
	Squash_merge_commit_message string `json:"squash_merge_commit_message,omitempty"` // The default value for a squash merge commit message: - `PR_BODY` - default to the pull request's body. - `COMMIT_MESSAGES` - default to the branch's commit messages. - `BLANK` - default to a blank commit message.
	Use_squash_pr_title_as_default bool `json:"use_squash_pr_title_as_default,omitempty"` // Whether a squash merge commit can use the pull request title as default. **This property is closing down. Please use `squash_merge_commit_title` instead.
	Archive_url string `json:"archive_url"`
	Forks int `json:"forks"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Html_url string `json:"html_url"`
	Node_id string `json:"node_id"`
	Subscribers_count int `json:"subscribers_count,omitempty"`
	Ssh_url string `json:"ssh_url"`
	Subscription_url string `json:"subscription_url"`
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"` // Whether to require contributors to sign off on web-based commits
	Full_name string `json:"full_name"`
	Statuses_url string `json:"statuses_url"`
	Custom_properties map[string]interface{} `json:"custom_properties,omitempty"` // The custom properties that were defined for the repository. The keys are the custom property names, and the values are the corresponding custom property values.
	Has_projects bool `json:"has_projects"` // Whether projects are enabled.
	Language string `json:"language"`
	License GeneratedType `json:"license"` // License Simple
	Releases_url string `json:"releases_url"`
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Topics []string `json:"topics,omitempty"`
	Default_branch string `json:"default_branch"` // The default branch of the repository.
	Pushed_at string `json:"pushed_at"`
	Subscribers_url string `json:"subscribers_url"`
	Commits_url string `json:"commits_url"`
	Open_issues_count int `json:"open_issues_count"`
	Fork bool `json:"fork"`
	Merge_commit_message string `json:"merge_commit_message,omitempty"` // The default value for a merge commit message. - `PR_TITLE` - default to the pull request's title. - `PR_BODY` - default to the pull request's body. - `BLANK` - default to a blank commit message.
	Teams_url string `json:"teams_url"`
	Hooks_url string `json:"hooks_url"`
	Labels_url string `json:"labels_url"`
	Archived bool `json:"archived"` // Whether the repository is archived.
	Description string `json:"description"`
	Starred_at string `json:"starred_at,omitempty"`
	Allow_forking bool `json:"allow_forking,omitempty"` // Whether to allow forking this repo
	Anonymous_access_enabled bool `json:"anonymous_access_enabled,omitempty"` // Whether anonymous git access is enabled for this repository
	Watchers int `json:"watchers"`
	Has_issues bool `json:"has_issues"` // Whether issues are enabled.
	Watchers_count int `json:"watchers_count"`
	Svn_url string `json:"svn_url"`
	Is_template bool `json:"is_template,omitempty"` // Whether this repository acts as a template that can be used to generate new repositories.
	Git_url string `json:"git_url"`
	Issues_url string `json:"issues_url"`
	Disabled bool `json:"disabled"` // Returns whether or not this repository disabled.
	Contents_url string `json:"contents_url"`
	Open_issues int `json:"open_issues"`
	Stargazers_url string `json:"stargazers_url"`
	Allow_rebase_merge bool `json:"allow_rebase_merge,omitempty"` // Whether to allow rebase merges for pull requests.
	Trees_url string `json:"trees_url"`
	Squash_merge_commit_title string `json:"squash_merge_commit_title,omitempty"` // The default value for a squash merge commit title: - `PR_TITLE` - default to the pull request's title. - `COMMIT_OR_PR_TITLE` - default to the commit's title (if only one commit) or the pull request's title (when more than one commit).
	Branches_url string `json:"branches_url"`
	Has_pages bool `json:"has_pages"`
	Forks_url string `json:"forks_url"`
	Git_tags_url string `json:"git_tags_url"`
	Pulls_url string `json:"pulls_url"`
	Id int64 `json:"id"` // Unique identifier of the repository
	Allow_auto_merge bool `json:"allow_auto_merge,omitempty"` // Whether to allow Auto-merge to be used on pull requests.
	Allow_merge_commit bool `json:"allow_merge_commit,omitempty"` // Whether to allow merge commits for pull requests.
	Tags_url string `json:"tags_url"`
	Template_repository map[string]interface{} `json:"template_repository,omitempty"`
	Keys_url string `json:"keys_url"`
	Issue_events_url string `json:"issue_events_url"`
	Private bool `json:"private"` // Whether the repository is private or public.
	Temp_clone_token string `json:"temp_clone_token,omitempty"`
	Collaborators_url string `json:"collaborators_url"`
	Has_discussions bool `json:"has_discussions,omitempty"` // Whether discussions are enabled.
	Clone_url string `json:"clone_url"`
	Deployments_url string `json:"deployments_url"`
	Size int `json:"size"` // The size of the repository, in kilobytes. Size is calculated hourly. When a repository is initially created, the size is 0.
	Downloads_url string `json:"downloads_url"`
	Compare_url string `json:"compare_url"`
	Has_wiki bool `json:"has_wiki"` // Whether the wiki is enabled.
	Mirror_url string `json:"mirror_url"`
	Network_count int `json:"network_count,omitempty"`
	Merges_url string `json:"merges_url"`
	Visibility string `json:"visibility,omitempty"` // The repository visibility: public, private, or internal.
	Milestones_url string `json:"milestones_url"`
	Forks_count int `json:"forks_count"`
	Git_commits_url string `json:"git_commits_url"`
	Homepage string `json:"homepage"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Review_requester GeneratedType `json:"review_requester,omitempty"` // A GitHub user.
	Milestone GeneratedType `json:"milestone,omitempty"` // Issue Event Milestone
	Rename GeneratedType `json:"rename,omitempty"` // Issue Event Rename
	Dismissed_review GeneratedType `json:"dismissed_review,omitempty"`
	Event string `json:"event"`
	Requested_reviewer GeneratedType `json:"requested_reviewer,omitempty"` // A GitHub user.
	Id int64 `json:"id"`
	Assigner GeneratedType `json:"assigner,omitempty"` // A GitHub user.
	Author_association string `json:"author_association,omitempty"` // How the author is associated with the repository.
	Commit_id string `json:"commit_id"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Project_card GeneratedType `json:"project_card,omitempty"` // Issue Event Project Card
	Label GeneratedType `json:"label,omitempty"` // Issue Event Label
	Lock_reason string `json:"lock_reason,omitempty"`
	Url string `json:"url"`
	Created_at string `json:"created_at"`
	Issue GeneratedType `json:"issue,omitempty"` // Issues are a great way to keep track of tasks, enhancements, and bugs for your projects.
	Requested_team Team `json:"requested_team,omitempty"` // Groups of organization members that gives permissions on specified repositories.
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Commit_url string `json:"commit_url"`
	Node_id string `json:"node_id"`
	Assignee GeneratedType `json:"assignee,omitempty"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	Type string `json:"type"`
}

// Webhooksmilestone represents the Webhooksmilestone schema from the OpenAPI specification
type Webhooksmilestone struct {
	Id int `json:"id"`
	Creator map[string]interface{} `json:"creator"`
	Updated_at string `json:"updated_at"`
	Node_id string `json:"node_id"`
	Title string `json:"title"` // The title of the milestone.
	Url string `json:"url"`
	Labels_url string `json:"labels_url"`
	State string `json:"state"` // The state of the milestone.
	Closed_at string `json:"closed_at"`
	Number int `json:"number"` // The number of the milestone.
	Created_at string `json:"created_at"`
	Html_url string `json:"html_url"`
	Due_on string `json:"due_on"`
	Open_issues int `json:"open_issues"`
	Closed_issues int `json:"closed_issues"`
	Description string `json:"description"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Created_at string `json:"created_at"`
	Ignored bool `json:"ignored"`
	Reason string `json:"reason"`
	Repository_url string `json:"repository_url,omitempty"`
	Subscribed bool `json:"subscribed"`
	Thread_url string `json:"thread_url,omitempty"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Branch string `json:"branch"`
	Path string `json:"path"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) the comment belongs to.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Comment Webhooksissuecomment `json:"comment"` // The [comment](https://docs.github.com/rest/issues/comments#get-an-issue-comment) itself.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Teams_url string `json:"teams_url"`
	Url string `json:"url"`
	Users []map[string]interface{} `json:"users"`
	Users_url string `json:"users_url"`
	Apps []map[string]interface{} `json:"apps"`
	Apps_url string `json:"apps_url"`
	Teams []map[string]interface{} `json:"teams"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Teams []Team `json:"teams"`
	Users []GeneratedType `json:"users"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Workflow Webhooksworkflow `json:"workflow"`
	Workflow_run map[string]interface{} `json:"workflow_run"`
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Thread map[string]interface{} `json:"thread"`
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Discussion_comment_url string `json:"discussion_comment_url"` // The API URL to get the discussion comment where the secret was detected.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Node_id string `json:"node_id"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Commit_url string `json:"commit_url"`
	Created_at string `json:"created_at"`
	Url string `json:"url"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Commit_id string `json:"commit_id"`
	Event string `json:"event"`
	Id int `json:"id"`
	Label map[string]interface{} `json:"label"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Name string `json:"name"` // The name of the secret.
	Updated_at string `json:"updated_at"`
	Created_at string `json:"created_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Definition GeneratedType `json:"definition"` // Custom property defined on an organization
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Sponsorship Webhookssponsorship `json:"sponsorship"`
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Review Webhooksreview `json:"review"` // The review that was affected.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Label Webhookslabel `json:"label"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Payload string `json:"payload"` // A URL-encoded string of the secret_scanning_alert_location.created JSON payload. The decoded payload is a JSON object.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Installation GeneratedType `json:"installation"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Changes map[string]interface{} `json:"changes"`
	Target_type string `json:"target_type"`
	Account map[string]interface{} `json:"account"`
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Status string `json:"status"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Alert map[string]interface{} `json:"alert"` // The code scanning alert involved in the event.
	Commit_oid string `json:"commit_oid"` // The commit SHA of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Action string `json:"action"`
	Ref string `json:"ref"` // The Git reference of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	App GeneratedType `json:"app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Created_at string `json:"created_at"`
	Head_branch string `json:"head_branch"`
	Pull_requests []GeneratedType `json:"pull_requests"`
	Url string `json:"url"`
	Check_runs_url string `json:"check_runs_url"`
	Runs_rerequestable bool `json:"runs_rerequestable,omitempty"`
	Node_id string `json:"node_id"`
	After string `json:"after"`
	Head_sha string `json:"head_sha"` // The SHA of the head commit that is being checked.
	Id int64 `json:"id"`
	Status string `json:"status"` // The phase of the lifecycle that the check suite is currently in. Statuses of waiting, requested, and pending are reserved for GitHub Actions check suites.
	Conclusion string `json:"conclusion"`
	Repository GeneratedType `json:"repository"` // Minimal Repository
	Head_commit GeneratedType `json:"head_commit"` // A commit.
	Latest_check_runs_count int `json:"latest_check_runs_count"`
	Rerequestable bool `json:"rerequestable,omitempty"`
	Before string `json:"before"`
	Updated_at string `json:"updated_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Use_default bool `json:"use_default"` // Whether to use the default template or not. If `true`, the `include_claim_keys` field is ignored.
	Include_claim_keys []string `json:"include_claim_keys,omitempty"` // Array of unique strings. Each claim key can only contain alphanumeric characters and underscores.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Errors []map[string]interface{} `json:"errors"`
}

// Tag represents the Tag schema from the OpenAPI specification
type Tag struct {
	Zipball_url string `json:"zipball_url"`
	Commit map[string]interface{} `json:"commit"`
	Name string `json:"name"`
	Node_id string `json:"node_id"`
	Tarball_url string `json:"tarball_url"`
}

// Reaction represents the Reaction schema from the OpenAPI specification
type Reaction struct {
	Created_at string `json:"created_at"`
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	User GeneratedType `json:"user"` // A GitHub user.
	Content string `json:"content"` // The reaction to use
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Payload string `json:"payload"` // A URL-encoded string of the check_run.created JSON payload. The decoded payload is a JSON object.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Issue Webhooksissue `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Type GeneratedType `json:"type"` // The type of issue.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Head_commit GeneratedType `json:"head_commit"` // A commit.
	Logs_url string `json:"logs_url"` // The URL to download the logs for the workflow run.
	Pull_requests []GeneratedType `json:"pull_requests"` // Pull requests that are open with a `head_sha` or `head_branch` that matches the workflow run. The returned pull requests do not necessarily indicate pull requests that triggered the run.
	Head_branch string `json:"head_branch"`
	Workflow_id int `json:"workflow_id"` // The ID of the parent workflow.
	Created_at string `json:"created_at"`
	Jobs_url string `json:"jobs_url"` // The URL to the jobs for the workflow run.
	Url string `json:"url"` // The URL to the workflow run.
	Rerun_url string `json:"rerun_url"` // The URL to rerun the workflow run.
	Check_suite_url string `json:"check_suite_url"` // The URL to the associated check suite.
	Status string `json:"status"`
	Artifacts_url string `json:"artifacts_url"` // The URL to the artifacts for the workflow run.
	Path string `json:"path"` // The full path of the workflow
	Updated_at string `json:"updated_at"`
	Head_repository_id int `json:"head_repository_id,omitempty"`
	Head_sha string `json:"head_sha"` // The SHA of the head commit that points to the version of the workflow being run.
	Actor GeneratedType `json:"actor,omitempty"` // A GitHub user.
	Check_suite_node_id string `json:"check_suite_node_id,omitempty"` // The node ID of the associated check suite.
	Run_number int `json:"run_number"` // The auto incrementing run number for the workflow run.
	Repository GeneratedType `json:"repository"` // Minimal Repository
	Previous_attempt_url string `json:"previous_attempt_url,omitempty"` // The URL to the previous attempted run of this workflow, if one exists.
	Display_title string `json:"display_title"` // The event-specific title associated with the run or the run-name if set, or the value of `run-name` if it is set in the workflow.
	Id int `json:"id"` // The ID of the workflow run.
	Name string `json:"name,omitempty"` // The name of the workflow run.
	Referenced_workflows []GeneratedType `json:"referenced_workflows,omitempty"`
	Event string `json:"event"`
	Node_id string `json:"node_id"`
	Workflow_url string `json:"workflow_url"` // The URL to the workflow.
	Check_suite_id int `json:"check_suite_id,omitempty"` // The ID of the associated check suite.
	Html_url string `json:"html_url"`
	Conclusion string `json:"conclusion"`
	Run_started_at string `json:"run_started_at,omitempty"` // The start time of the latest run. Resets on re-run.
	Head_repository GeneratedType `json:"head_repository"` // Minimal Repository
	Cancel_url string `json:"cancel_url"` // The URL to cancel the workflow run.
	Run_attempt int `json:"run_attempt,omitempty"` // Attempt number of the run, 1 for first attempt and higher if the workflow was re-run.
	Triggering_actor GeneratedType `json:"triggering_actor,omitempty"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pusher_type string `json:"pusher_type"` // The pusher type for the event. Can be either `user` or a deploy key.
	Description string `json:"description"` // The repository's current description.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Ref string `json:"ref"` // The [`git ref`](https://docs.github.com/rest/git/refs#get-a-reference) resource.
	Ref_type string `json:"ref_type"` // The type of Git ref object created in the repository.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Master_branch string `json:"master_branch"` // The name of the repository's default branch (usually `main`).
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"` // The action performed. Can be `created`.
	Comment map[string]interface{} `json:"comment"` // The [commit comment](${externalDocsUpapp/api/description/components/schemas/webhooks/issue-comment-created.yamlrl}/rest/commits/comments#get-a-commit-comment) resource.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository_name map[string]interface{} `json:"repository_name"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repositories_url string `json:"repositories_url"`
	Repository_selection string `json:"repository_selection"` // Describe whether all repositories have been selected or there's a selection involved
	Single_file_name string `json:"single_file_name"`
	Single_file_paths []string `json:"single_file_paths,omitempty"`
	Account GeneratedType `json:"account"` // A GitHub user.
	Has_multiple_single_files bool `json:"has_multiple_single_files,omitempty"`
	Permissions GeneratedType `json:"permissions"` // The permissions granted to the user access token.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Key string `json:"key"`
	Name string `json:"name"`
	Url string `json:"url"`
	Body string `json:"body,omitempty"`
	Html_url string `json:"html_url"`
}

// Feed represents the Feed schema from the OpenAPI specification
type Feed struct {
	User_url string `json:"user_url"`
	Repository_discussions_category_url string `json:"repository_discussions_category_url,omitempty"` // A feed of discussions for a given repository and category.
	Current_user_organization_url string `json:"current_user_organization_url,omitempty"`
	Current_user_organization_urls []string `json:"current_user_organization_urls,omitempty"`
	Current_user_public_url string `json:"current_user_public_url,omitempty"`
	Repository_discussions_url string `json:"repository_discussions_url,omitempty"` // A feed of discussions for a given repository.
	Security_advisories_url string `json:"security_advisories_url,omitempty"`
	_links map[string]interface{} `json:"_links"`
	Current_user_actor_url string `json:"current_user_actor_url,omitempty"`
	Current_user_url string `json:"current_user_url,omitempty"`
	Timeline_url string `json:"timeline_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Project_id int `json:"project_id"`
	Project_url string `json:"project_url"`
	Url string `json:"url"`
	Column_name string `json:"column_name"`
	Id int `json:"id"`
	Previous_column_name string `json:"previous_column_name,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Alert map[string]interface{} `json:"alert"` // The security alert of the vulnerable dependency.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Key string `json:"key"`
	Name string `json:"name"`
	Node_id string `json:"node_id"`
	Spdx_id string `json:"spdx_id"`
	Url string `json:"url"`
	Html_url string `json:"html_url,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Review map[string]interface{} `json:"review"` // The review that was affected.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Documentation_url string `json:"documentation_url,omitempty"`
	Message string `json:"message,omitempty"`
	Status string `json:"status,omitempty"`
	Url string `json:"url,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Release Webhooksrelease1 `json:"release"` // The [release](https://docs.github.com/rest/releases/releases/#get-a-release) object.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Forkee interface{} `json:"forkee"` // The created [`repository`](https://docs.github.com/rest/repos/repos#get-a-repository) resource.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// Repository represents the Repository schema from the OpenAPI specification
type Repository struct {
	Merge_commit_message string `json:"merge_commit_message,omitempty"` // The default value for a merge commit message. - `PR_TITLE` - default to the pull request's title. - `PR_BODY` - default to the pull request's body. - `BLANK` - default to a blank commit message.
	Downloads_url string `json:"downloads_url"`
	Private bool `json:"private"` // Whether the repository is private or public.
	Collaborators_url string `json:"collaborators_url"`
	Homepage string `json:"homepage"`
	Commits_url string `json:"commits_url"`
	Is_template bool `json:"is_template,omitempty"` // Whether this repository acts as a template that can be used to generate new repositories.
	Description string `json:"description"`
	Issue_events_url string `json:"issue_events_url"`
	Events_url string `json:"events_url"`
	Created_at string `json:"created_at"`
	Hooks_url string `json:"hooks_url"`
	Fork bool `json:"fork"`
	Comments_url string `json:"comments_url"`
	Id int64 `json:"id"` // Unique identifier of the repository
	Git_tags_url string `json:"git_tags_url"`
	Forks_count int `json:"forks_count"`
	Keys_url string `json:"keys_url"`
	Allow_forking bool `json:"allow_forking,omitempty"` // Whether to allow forking this repo
	Contributors_url string `json:"contributors_url"`
	Git_url string `json:"git_url"`
	Pulls_url string `json:"pulls_url"`
	Size int `json:"size"` // The size of the repository, in kilobytes. Size is calculated hourly. When a repository is initially created, the size is 0.
	Has_discussions bool `json:"has_discussions,omitempty"` // Whether discussions are enabled.
	Disabled bool `json:"disabled"` // Returns whether or not this repository disabled.
	Has_pages bool `json:"has_pages"`
	Compare_url string `json:"compare_url"`
	Issues_url string `json:"issues_url"`
	Subscribers_url string `json:"subscribers_url"`
	Clone_url string `json:"clone_url"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Assignees_url string `json:"assignees_url"`
	Temp_clone_token string `json:"temp_clone_token,omitempty"`
	Updated_at string `json:"updated_at"`
	Visibility string `json:"visibility,omitempty"` // The repository visibility: public, private, or internal.
	Name string `json:"name"` // The name of the repository.
	Statuses_url string `json:"statuses_url"`
	Svn_url string `json:"svn_url"`
	Code_search_index_status map[string]interface{} `json:"code_search_index_status,omitempty"` // The status of the code search index for this repository
	Allow_squash_merge bool `json:"allow_squash_merge,omitempty"` // Whether to allow squash merges for pull requests.
	Git_commits_url string `json:"git_commits_url"`
	Issue_comment_url string `json:"issue_comment_url"`
	Has_downloads bool `json:"has_downloads"` // Whether downloads are enabled.
	Trees_url string `json:"trees_url"`
	Default_branch string `json:"default_branch"` // The default branch of the repository.
	Full_name string `json:"full_name"`
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"` // Whether to delete head branches when pull requests are merged
	Has_issues bool `json:"has_issues"` // Whether issues are enabled.
	Notifications_url string `json:"notifications_url"`
	Allow_rebase_merge bool `json:"allow_rebase_merge,omitempty"` // Whether to allow rebase merges for pull requests.
	Open_issues int `json:"open_issues"`
	Pushed_at string `json:"pushed_at"`
	Merges_url string `json:"merges_url"`
	Allow_merge_commit bool `json:"allow_merge_commit,omitempty"` // Whether to allow merge commits for pull requests.
	Mirror_url string `json:"mirror_url"`
	Subscription_url string `json:"subscription_url"`
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"` // Whether to require contributors to sign off on web-based commits
	Allow_update_branch bool `json:"allow_update_branch,omitempty"` // Whether or not a pull request head branch that is behind its base branch can always be updated even if it is not required to be up to date before merging.
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Branches_url string `json:"branches_url"`
	Deployments_url string `json:"deployments_url"`
	Blobs_url string `json:"blobs_url"`
	Url string `json:"url"`
	Releases_url string `json:"releases_url"`
	Allow_auto_merge bool `json:"allow_auto_merge,omitempty"` // Whether to allow Auto-merge to be used on pull requests.
	Forks_url string `json:"forks_url"`
	Has_wiki bool `json:"has_wiki"` // Whether the wiki is enabled.
	Master_branch string `json:"master_branch,omitempty"`
	Topics []string `json:"topics,omitempty"`
	License GeneratedType `json:"license"` // License Simple
	Archived bool `json:"archived"` // Whether the repository is archived.
	Stargazers_count int `json:"stargazers_count"`
	Watchers int `json:"watchers"`
	Milestones_url string `json:"milestones_url"`
	Ssh_url string `json:"ssh_url"`
	Archive_url string `json:"archive_url"`
	Labels_url string `json:"labels_url"`
	Use_squash_pr_title_as_default bool `json:"use_squash_pr_title_as_default,omitempty"` // Whether a squash merge commit can use the pull request title as default. **This property is closing down. Please use `squash_merge_commit_title` instead.
	Contents_url string `json:"contents_url"`
	Squash_merge_commit_title string `json:"squash_merge_commit_title,omitempty"` // The default value for a squash merge commit title: - `PR_TITLE` - default to the pull request's title. - `COMMIT_OR_PR_TITLE` - default to the commit's title (if only one commit) or the pull request's title (when more than one commit).
	Html_url string `json:"html_url"`
	Languages_url string `json:"languages_url"`
	Git_refs_url string `json:"git_refs_url"`
	Stargazers_url string `json:"stargazers_url"`
	Watchers_count int `json:"watchers_count"`
	Node_id string `json:"node_id"`
	Open_issues_count int `json:"open_issues_count"`
	Anonymous_access_enabled bool `json:"anonymous_access_enabled,omitempty"` // Whether anonymous git access is enabled for this repository
	Has_projects bool `json:"has_projects"` // Whether projects are enabled.
	Forks int `json:"forks"`
	Merge_commit_title string `json:"merge_commit_title,omitempty"` // The default value for a merge commit title. - `PR_TITLE` - default to the pull request's title. - `MERGE_MESSAGE` - default to the classic title for a merge message (e.g., Merge pull request #123 from branch-name).
	Squash_merge_commit_message string `json:"squash_merge_commit_message,omitempty"` // The default value for a squash merge commit message: - `PR_BODY` - default to the pull request's body. - `COMMIT_MESSAGES` - default to the branch's commit messages. - `BLANK` - default to a blank commit message.
	Language string `json:"language"`
	Tags_url string `json:"tags_url"`
	Starred_at string `json:"starred_at,omitempty"`
	Teams_url string `json:"teams_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Name string `json:"name"` // Name of the project column
	Node_id string `json:"node_id"`
	Project_url string `json:"project_url"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	Cards_url string `json:"cards_url"`
	Created_at string `json:"created_at"`
	Id int `json:"id"` // The unique identifier of the project column
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Node_id string `json:"node_id"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Review_requester GeneratedType `json:"review_requester"` // A GitHub user.
	Commit_url string `json:"commit_url"`
	Requested_team Team `json:"requested_team,omitempty"` // Groups of organization members that gives permissions on specified repositories.
	Created_at string `json:"created_at"`
	Requested_reviewer GeneratedType `json:"requested_reviewer,omitempty"` // A GitHub user.
	Url string `json:"url"`
	Commit_id string `json:"commit_id"`
	Event string `json:"event"`
	Id int `json:"id"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Parent_issue_id float64 `json:"parent_issue_id"` // The ID of the parent issue.
	Sub_issue Issue `json:"sub_issue"` // Issues are a great way to keep track of tasks, enhancements, and bugs for your projects.
	Sub_issue_id float64 `json:"sub_issue_id"` // The ID of the sub-issue.
	Action string `json:"action"`
	Parent_issue_repo Repository `json:"parent_issue_repo"` // A repository on GitHub.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Parent_issue Issue `json:"parent_issue"` // Issues are a great way to keep track of tasks, enhancements, and bugs for your projects.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Definition GeneratedType `json:"definition"` // Custom property defined on an organization
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Check_suite map[string]interface{} `json:"check_suite"` // The [check_suite](https://docs.github.com/rest/checks/suites#get-a-check-suite).
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Body string `json:"body"` // The generated body describing the contents of the release supporting markdown formatting
	Name string `json:"name"` // The generated name of the release
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	+1 int `json:"+1"`
	Confused int `json:"confused"`
	Eyes int `json:"eyes"`
	Laugh int `json:"laugh"`
	Total_count int `json:"total_count"`
	-1 int `json:"-1"`
	Heart int `json:"heart"`
	Hooray int `json:"hooray"`
	Rocket int `json:"rocket"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Git_url string `json:"git_url"`
	Html_url string `json:"html_url"`
	Sha string `json:"sha"`
	Size int `json:"size"`
	Target string `json:"target,omitempty"`
	Url string `json:"url"`
	Path string `json:"path"`
	Submodule_git_url string `json:"submodule_git_url,omitempty"`
	Content string `json:"content"`
	Name string `json:"name"`
	Type string `json:"type"`
	_links map[string]interface{} `json:"_links"`
	Download_url string `json:"download_url"`
	Encoding string `json:"encoding"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Title string `json:"title"`
	Uniques int `json:"uniques"`
	Count int `json:"count"`
	Path string `json:"path"`
}

// Classroom represents the Classroom schema from the OpenAPI specification
type Classroom struct {
	Url string `json:"url"` // The URL of the classroom on GitHub Classroom.
	Archived bool `json:"archived"` // Whether classroom is archived.
	Id int `json:"id"` // Unique identifier of the classroom.
	Name string `json:"name"` // The name of the classroom.
	Organization GeneratedType `json:"organization"` // A GitHub organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Issue Webhooksissue2 `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Target_ref string `json:"target_ref,omitempty"` // The Git reference of target branch for the commit. Branch needs to already exist. For more information, see "[Git References](https://git-scm.com/book/en/v2/Git-Internals-Git-References)" in the Git documentation.
	Message string `json:"message,omitempty"` // Commit message to be used.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Sub_issue_id float64 `json:"sub_issue_id"` // The ID of the sub-issue.
	Sub_issue_repo Repository `json:"sub_issue_repo"` // A repository on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Parent_issue_id float64 `json:"parent_issue_id"` // The ID of the parent issue.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Parent_issue Issue `json:"parent_issue"` // Issues are a great way to keep track of tasks, enhancements, and bugs for your projects.
	Sub_issue Issue `json:"sub_issue"` // Issues are a great way to keep track of tasks, enhancements, and bugs for your projects.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Filename string `json:"filename"`
	Os string `json:"os"`
	Sha256_checksum string `json:"sha256_checksum,omitempty"`
	Temp_download_token string `json:"temp_download_token,omitempty"` // A short lived bearer token used to download the runner, if needed.
	Architecture string `json:"architecture"`
	Download_url string `json:"download_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Reviewers []map[string]interface{} `json:"reviewers,omitempty"`
	Action string `json:"action"`
	Approver Webhooksapprover `json:"approver,omitempty"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Workflow_job_run Webhooksworkflowjobrun `json:"workflow_job_run,omitempty"`
	Comment string `json:"comment,omitempty"`
	Workflow_job_runs []map[string]interface{} `json:"workflow_job_runs,omitempty"`
	Workflow_run map[string]interface{} `json:"workflow_run"`
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Since string `json:"since"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Alerts_threshold string `json:"alerts_threshold"` // The severity level at which code scanning results that raise alerts block a reference update. For more information on alert severity levels, see "[About code scanning alerts](https://docs.github.com/code-security/code-scanning/managing-code-scanning-alerts/about-code-scanning-alerts#about-alert-severity-and-security-severity-levels)."
	Security_alerts_threshold string `json:"security_alerts_threshold"` // The severity level at which code scanning results that raise security alerts block a reference update. For more information on security severity levels, see "[About code scanning alerts](https://docs.github.com/code-security/code-scanning/managing-code-scanning-alerts/about-code-scanning-alerts#about-alert-severity-and-security-severity-levels)."
	Tool string `json:"tool"` // The name of a code scanning tool
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Releases_url string `json:"releases_url"` // A template for the API URL to get information about releases on the repository.
	Html_url string `json:"html_url"` // The URL to view the repository on GitHub.com.
	Statuses_url string `json:"statuses_url"` // A template for the API URL to get information about statuses of a commit.
	Url string `json:"url"` // The URL to get more information about the repository from the GitHub API.
	Branches_url string `json:"branches_url"` // A template for the API URL to get information about branches in the repository.
	Collaborators_url string `json:"collaborators_url"` // A template for the API URL to get information about collaborators of the repository.
	Git_refs_url string `json:"git_refs_url"` // A template for the API URL to get information about Git refs of the repository.
	Contents_url string `json:"contents_url"` // A template for the API URL to get the contents of the repository.
	Stargazers_url string `json:"stargazers_url"` // The API URL to list the stargazers on the repository.
	Comments_url string `json:"comments_url"` // A template for the API URL to get information about comments on the repository.
	Hooks_url string `json:"hooks_url"` // The API URL to list the hooks on the repository.
	Labels_url string `json:"labels_url"` // A template for the API URL to get information about labels of the repository.
	Pulls_url string `json:"pulls_url"` // A template for the API URL to get information about pull requests on the repository.
	Git_commits_url string `json:"git_commits_url"` // A template for the API URL to get information about Git commits of the repository.
	Issues_url string `json:"issues_url"` // A template for the API URL to get information about issues on the repository.
	Compare_url string `json:"compare_url"` // A template for the API URL to compare two commits or refs.
	Fork bool `json:"fork"` // Whether the repository is a fork.
	Languages_url string `json:"languages_url"` // The API URL to get information about the languages of the repository.
	Forks_url string `json:"forks_url"` // The API URL to list the forks of the repository.
	Issue_comment_url string `json:"issue_comment_url"` // A template for the API URL to get information about issue comments on the repository.
	Node_id string `json:"node_id"` // The GraphQL identifier of the repository.
	Subscribers_url string `json:"subscribers_url"` // The API URL to list the subscribers on the repository.
	Name string `json:"name"` // The name of the repository.
	Commits_url string `json:"commits_url"` // A template for the API URL to get information about commits on the repository.
	Keys_url string `json:"keys_url"` // A template for the API URL to get information about deploy keys on the repository.
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Full_name string `json:"full_name"` // The full, globally unique, name of the repository.
	Notifications_url string `json:"notifications_url"` // A template for the API URL to get information about notifications on the repository.
	Archive_url string `json:"archive_url"` // A template for the API URL to download the repository as an archive.
	Description string `json:"description"` // The repository description.
	Git_tags_url string `json:"git_tags_url"` // A template for the API URL to get information about Git tags of the repository.
	Id int64 `json:"id"` // A unique identifier of the repository.
	Assignees_url string `json:"assignees_url"` // A template for the API URL to list the available assignees for issues in the repository.
	Downloads_url string `json:"downloads_url"` // The API URL to list the downloads on the repository.
	Private bool `json:"private"` // Whether the repository is private.
	Deployments_url string `json:"deployments_url"` // The API URL to list the deployments of the repository.
	Milestones_url string `json:"milestones_url"` // A template for the API URL to get information about milestones of the repository.
	Tags_url string `json:"tags_url"` // The API URL to get information about tags on the repository.
	Trees_url string `json:"trees_url"` // A template for the API URL to create or retrieve a raw Git tree of the repository.
	Issue_events_url string `json:"issue_events_url"` // A template for the API URL to get information about issue events on the repository.
	Merges_url string `json:"merges_url"` // The API URL to merge branches in the repository.
	Teams_url string `json:"teams_url"` // The API URL to list the teams on the repository.
	Subscription_url string `json:"subscription_url"` // The API URL to subscribe to notifications for this repository.
	Events_url string `json:"events_url"` // The API URL to list the events of the repository.
	Blobs_url string `json:"blobs_url"` // A template for the API URL to create or retrieve a raw Git blob in the repository.
	Contributors_url string `json:"contributors_url"` // A template for the API URL to list the contributors to the repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Type string `json:"type"`
	Parameters map[string]interface{} `json:"parameters,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Type string `json:"type"` // The type of the reviewer
	Id int `json:"id"` // ID of the reviewer which must review changes to matching files.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enabled bool `json:"enabled"` // Whether the deployment protection rule is enabled for the environment.
	Id int `json:"id"` // The unique identifier for the deployment protection rule.
	Node_id string `json:"node_id"` // The node ID for the deployment protection rule.
	App GeneratedType `json:"app"` // A GitHub App that is providing a custom deployment protection rule.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sbom map[string]interface{} `json:"sbom"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Workflow_run map[string]interface{} `json:"workflow_run"`
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Deployment map[string]interface{} `json:"deployment"` // The [deployment](https://docs.github.com/rest/deployments/deployments#list-deployments).
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Workflow Webhooksworkflow `json:"workflow"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// Integration represents the Integration schema from the OpenAPI specification
type Integration struct {
	Description string `json:"description"`
	Node_id string `json:"node_id"`
	Updated_at string `json:"updated_at"`
	Created_at string `json:"created_at"`
	Events []string `json:"events"` // The list of events for the GitHub app. Note that the `installation_target`, `security_advisory`, and `meta` events are not included because they are global events and not specific to an installation.
	Id int `json:"id"` // Unique identifier of the GitHub app
	Installations_count int `json:"installations_count,omitempty"` // The number of installations associated with the GitHub app. Only returned when the integration is requesting details about itself.
	Client_id string `json:"client_id,omitempty"`
	Html_url string `json:"html_url"`
	Permissions map[string]interface{} `json:"permissions"` // The set of permissions for the GitHub app
	Slug string `json:"slug,omitempty"` // The slug name of the GitHub app
	External_url string `json:"external_url"`
	Name string `json:"name"` // The name of the GitHub app
	Owner interface{} `json:"owner"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Included_gigabytes_bandwidth int `json:"included_gigabytes_bandwidth"` // Free storage space (GB) for GitHub Packages.
	Total_gigabytes_bandwidth_used int `json:"total_gigabytes_bandwidth_used"` // Sum of the free and paid storage space (GB) for GitHuub Packages.
	Total_paid_gigabytes_bandwidth_used int `json:"total_paid_gigabytes_bandwidth_used"` // Total paid storage space (GB) for GitHuub Packages.
}

// Webhookschanges8 represents the Webhookschanges8 schema from the OpenAPI specification
type Webhookschanges8 struct {
	Tier map[string]interface{} `json:"tier"`
}

// Traffic represents the Traffic schema from the OpenAPI specification
type Traffic struct {
	Count int `json:"count"`
	Timestamp string `json:"timestamp"`
	Uniques int `json:"uniques"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Updated_at string `json:"updated_at"`
	Archived_at string `json:"archived_at"`
	Created_at string `json:"created_at"`
	Creator GeneratedType `json:"creator,omitempty"` // A GitHub user.
	Project_node_id string `json:"project_node_id,omitempty"`
	Content_node_id string `json:"content_node_id"`
	Content_type string `json:"content_type"` // The type of content tracked in a project item
	Id float64 `json:"id"`
	Node_id string `json:"node_id,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	Type string `json:"type"`
}

// Webhooksprojectchanges represents the Webhooksprojectchanges schema from the OpenAPI specification
type Webhooksprojectchanges struct {
	Archived_at map[string]interface{} `json:"archived_at,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	User Webhooksuser `json:"user,omitempty"`
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Invitation map[string]interface{} `json:"invitation"` // The invitation for the user or email if the action is `member_invited`.
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Name string `json:"name,omitempty"` // The name pattern that branches or tags must match in order to deploy to the environment.
	Node_id string `json:"node_id,omitempty"`
	Type string `json:"type,omitempty"` // Whether this rule targets a branch or tag.
	Id int `json:"id,omitempty"` // The unique identifier of the branch or tag policy.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Cpu_cores int `json:"cpu_cores"` // The number of cores.
	Id string `json:"id"` // The ID used for the `size` parameter when creating a new runner.
	Memory_gb int `json:"memory_gb"` // The available RAM for the machine spec.
	Storage_gb int `json:"storage_gb"` // The available SSD storage for the machine spec.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Alert GeneratedType `json:"alert"` // A Dependabot alert.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Status string `json:"status"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Rate GeneratedType `json:"rate"`
	Resources map[string]interface{} `json:"resources"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Membership Webhooksmembership `json:"membership,omitempty"` // The membership between the user and the organization. Not present when the action is `member_invited`.
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes Webhookschanges8 `json:"changes"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Sponsorship Webhookssponsorship `json:"sponsorship"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Effective_date string `json:"effective_date,omitempty"` // The `pending_cancellation` and `pending_tier_change` event types will include the date the cancellation or tier change will take effect.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Has_multiple_single_files bool `json:"has_multiple_single_files,omitempty"`
	Permissions GeneratedType `json:"permissions,omitempty"` // The permissions granted to the user access token.
	Repositories []Repository `json:"repositories,omitempty"`
	Repository_selection string `json:"repository_selection,omitempty"`
	Single_file string `json:"single_file,omitempty"`
	Single_file_paths []string `json:"single_file_paths,omitempty"`
	Token string `json:"token"`
	Expires_at string `json:"expires_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Protected bool `json:"protected"`
	Protection GeneratedType `json:"protection"` // Branch Protection
	Protection_url string `json:"protection_url"`
	Required_approving_review_count int `json:"required_approving_review_count,omitempty"`
	_links map[string]interface{} `json:"_links"`
	Commit Commit `json:"commit"` // Commit
	Name string `json:"name"`
	Pattern string `json:"pattern,omitempty"`
}

// Webhookssecurityadvisory represents the Webhookssecurityadvisory schema from the OpenAPI specification
type Webhookssecurityadvisory struct {
	Identifiers []map[string]interface{} `json:"identifiers"`
	Vulnerabilities []map[string]interface{} `json:"vulnerabilities"`
	Cvss_severities GeneratedType `json:"cvss_severities,omitempty"`
	Ghsa_id string `json:"ghsa_id"`
	References []map[string]interface{} `json:"references"`
	Updated_at string `json:"updated_at"`
	Withdrawn_at string `json:"withdrawn_at"`
	Cvss map[string]interface{} `json:"cvss"`
	Severity string `json:"severity"`
	Cwes []map[string]interface{} `json:"cwes"`
	Published_at string `json:"published_at"`
	Summary string `json:"summary"`
	Description string `json:"description"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Rule Webhooksrule `json:"rule"` // The branch protection rule. Includes a `name` and all the [branch protection settings](https://docs.github.com/github/administering-a-repository/defining-the-mergeability-of-pull-requests/about-protected-branches#about-branch-protection-settings) applied to branches that match the name. Binary settings are boolean. Multi-level configurations are one of `off`, `non_admins`, or `everyone`. Actor and build lists are arrays of strings.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Added_this_cycle int `json:"added_this_cycle,omitempty"` // Seats added during the current billing cycle.
	Inactive_this_cycle int `json:"inactive_this_cycle,omitempty"` // The number of seats that have not used Copilot during the current billing cycle.
	Pending_cancellation int `json:"pending_cancellation,omitempty"` // The number of seats that are pending cancellation at the end of the current billing cycle.
	Pending_invitation int `json:"pending_invitation,omitempty"` // The number of users who have been invited to receive a Copilot seat through this organization.
	Total int `json:"total,omitempty"` // The total number of seats being billed for the organization as of the current billing cycle.
	Active_this_cycle int `json:"active_this_cycle,omitempty"` // The number of seats that have used Copilot during the current billing cycle.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Username string `json:"username,omitempty"` // The username to use when authenticating with the private registry.
	Visibility string `json:"visibility"` // Which type of organization repositories have access to the private registry.
	Created_at string `json:"created_at"`
	Name string `json:"name"` // The name of the private registry configuration.
	Registry_type string `json:"registry_type"` // The registry type.
	Updated_at string `json:"updated_at"`
}

// Workflow represents the Workflow schema from the OpenAPI specification
type Workflow struct {
	Html_url string `json:"html_url"`
	Created_at string `json:"created_at"`
	State string `json:"state"`
	Badge_url string `json:"badge_url"`
	Name string `json:"name"`
	Deleted_at string `json:"deleted_at,omitempty"`
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	Path string `json:"path"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"` // The action that was performed.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Name string `json:"name"`
	Active bool `json:"active"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	Deliveries_url string `json:"deliveries_url,omitempty"`
	Ping_url string `json:"ping_url"`
	Config map[string]interface{} `json:"config"`
	Events []string `json:"events"`
	Type string `json:"type"`
	Id int `json:"id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Alert GeneratedType `json:"alert"` // A Dependabot alert.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Permissions_upgraded map[string]interface{} `json:"permissions_upgraded"` // Requested permissions that elevate access for a previously approved request for access, categorized by type of permission.
	Token_name string `json:"token_name"` // The name given to the user's token. This field can also be found in an organization's settings page for Active Tokens.
	Permissions_added map[string]interface{} `json:"permissions_added"` // New requested permissions, categorized by type of permission.
	Permissions_result map[string]interface{} `json:"permissions_result"` // Permissions requested, categorized by type of permission. This field incorporates `permissions_added` and `permissions_upgraded`.
	Repository_count int `json:"repository_count"` // The number of repositories the token is requesting access to. This field is only populated when `repository_selection` is `subset`.
	Token_expires_at string `json:"token_expires_at"` // Date and time when the associated fine-grained personal access token expires.
	Token_last_used_at string `json:"token_last_used_at"` // Date and time when the associated fine-grained personal access token was last used for authentication.
	Created_at string `json:"created_at"` // Date and time when the request for access was created.
	Id int `json:"id"` // Unique identifier of the request for access via fine-grained personal access token. Used as the `pat_request_id` parameter in the list and review API calls.
	Repositories []map[string]interface{} `json:"repositories"` // An array of repository objects the token is requesting access to. This field is only populated when `repository_selection` is `subset`.
	Repository_selection string `json:"repository_selection"` // Type of repository selection requested.
	Token_id int `json:"token_id"` // Unique identifier of the user's token. This field can also be found in audit log events and the organization's settings for their PAT grants.
	Token_expired bool `json:"token_expired"` // Whether the associated fine-grained personal access token has expired.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"` // The changes to the collaborator permissions
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Member Webhooksuser `json:"member"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Base_sha string `json:"base_sha"` // The SHA of the merge group's parent commit.
	Head_commit GeneratedType `json:"head_commit"` // A commit.
	Head_ref string `json:"head_ref"` // The full ref of the merge group.
	Head_sha string `json:"head_sha"` // The SHA of the merge group.
	Base_ref string `json:"base_ref"` // The full ref of the branch the merge group will be merged into.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
	Branches []map[string]interface{} `json:"branches"` // An array of branch objects containing the status' SHA. Each branch contains the given SHA, but the SHA may or may not be the head of the branch. The array includes a maximum of 10 branches.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Commit map[string]interface{} `json:"commit"`
	State string `json:"state"` // The new state. Can be `pending`, `success`, `failure`, or `error`.
	Avatar_url string `json:"avatar_url,omitempty"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Context string `json:"context"`
	Name string `json:"name"`
	Target_url string `json:"target_url"` // The optional link added to the status.
	Sha string `json:"sha"` // The Commit SHA.
	Description string `json:"description"` // The optional human-readable description added to the status.
	Id int `json:"id"` // The unique identifier of the status.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Alert GeneratedType `json:"alert"` // A Dependabot alert.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Limit int `json:"limit"`
	Remaining int `json:"remaining"`
	Reset int `json:"reset"`
	Used int `json:"used"`
}

// Webhooksissue represents the Webhooksissue schema from the OpenAPI specification
type Webhooksissue struct {
	Comments int `json:"comments"`
	Events_url string `json:"events_url"`
	Html_url string `json:"html_url"`
	Created_at string `json:"created_at"`
	Milestone map[string]interface{} `json:"milestone"` // A collection of related issues and pull requests.
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Repository_url string `json:"repository_url"`
	Assignees []map[string]interface{} `json:"assignees"`
	Draft bool `json:"draft,omitempty"`
	Node_id string `json:"node_id"`
	Updated_at string `json:"updated_at"`
	Locked bool `json:"locked,omitempty"`
	Body string `json:"body"` // Contents of the issue
	Active_lock_reason string `json:"active_lock_reason"`
	Number int `json:"number"`
	Reactions map[string]interface{} `json:"reactions"`
	Url string `json:"url"` // URL for the issue
	Labels []map[string]interface{} `json:"labels,omitempty"`
	Id int64 `json:"id"`
	Type GeneratedType `json:"type,omitempty"` // The type of issue.
	User map[string]interface{} `json:"user"`
	Closed_at string `json:"closed_at"`
	Performed_via_github_app map[string]interface{} `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Sub_issues_summary map[string]interface{} `json:"sub_issues_summary,omitempty"`
	Title string `json:"title"` // Title of the issue
	State_reason string `json:"state_reason,omitempty"`
	Comments_url string `json:"comments_url"`
	Assignee map[string]interface{} `json:"assignee,omitempty"`
	Pull_request map[string]interface{} `json:"pull_request,omitempty"`
	State string `json:"state,omitempty"` // State of the issue; either 'open' or 'closed'
	Labels_url string `json:"labels_url"`
	Timeline_url string `json:"timeline_url,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request GeneratedType `json:"pull_request"`
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// Webhooksteam1 represents the Webhooksteam1 schema from the OpenAPI specification
type Webhooksteam1 struct {
	Slug string `json:"slug,omitempty"`
	Url string `json:"url,omitempty"` // URL for the team
	Deleted bool `json:"deleted,omitempty"`
	Html_url string `json:"html_url,omitempty"`
	Id int `json:"id"` // Unique identifier of the team
	Parent map[string]interface{} `json:"parent,omitempty"`
	Permission string `json:"permission,omitempty"` // Permission that the team will have for its repositories
	Name string `json:"name"` // Name of the team
	Repositories_url string `json:"repositories_url,omitempty"`
	Description string `json:"description,omitempty"` // Description of the team
	Notification_setting string `json:"notification_setting,omitempty"` // Whether team members will receive notifications when their team is @mentioned
	Privacy string `json:"privacy,omitempty"`
	Members_url string `json:"members_url,omitempty"`
	Node_id string `json:"node_id,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Auto_merge GeneratedType `json:"auto_merge"` // The status of auto merging a pull request.
	Number int `json:"number"`
	Node_id string `json:"node_id"`
	Labels []map[string]interface{} `json:"labels"`
	Requested_reviewers []GeneratedType `json:"requested_reviewers,omitempty"`
	_links map[string]interface{} `json:"_links"`
	Title string `json:"title"`
	Base map[string]interface{} `json:"base"`
	Url string `json:"url"`
	Commits_url string `json:"commits_url"`
	User GeneratedType `json:"user"` // A GitHub user.
	Html_url string `json:"html_url"`
	Assignees []GeneratedType `json:"assignees,omitempty"`
	Head map[string]interface{} `json:"head"`
	Diff_url string `json:"diff_url"`
	Milestone GeneratedType `json:"milestone"` // A collection of related issues and pull requests.
	Created_at string `json:"created_at"`
	Merged_at string `json:"merged_at"`
	Review_comment_url string `json:"review_comment_url"`
	Statuses_url string `json:"statuses_url"`
	Updated_at string `json:"updated_at"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Assignee GeneratedType `json:"assignee"` // A GitHub user.
	Comments_url string `json:"comments_url"`
	Id int64 `json:"id"`
	Draft bool `json:"draft,omitempty"` // Indicates whether or not the pull request is a draft.
	Review_comments_url string `json:"review_comments_url"`
	Body string `json:"body"`
	Issue_url string `json:"issue_url"`
	State string `json:"state"`
	Requested_teams []Team `json:"requested_teams,omitempty"`
	Patch_url string `json:"patch_url"`
	Locked bool `json:"locked"`
	Merge_commit_sha string `json:"merge_commit_sha"`
	Active_lock_reason string `json:"active_lock_reason,omitempty"`
	Closed_at string `json:"closed_at"`
}

// Webhooksuser represents the Webhooksuser schema from the OpenAPI specification
type Webhooksuser struct {
	Avatar_url string `json:"avatar_url,omitempty"`
	Following_url string `json:"following_url,omitempty"`
	Gravatar_id string `json:"gravatar_id,omitempty"`
	Name string `json:"name,omitempty"`
	Organizations_url string `json:"organizations_url,omitempty"`
	Html_url string `json:"html_url,omitempty"`
	Followers_url string `json:"followers_url,omitempty"`
	Gists_url string `json:"gists_url,omitempty"`
	Deleted bool `json:"deleted,omitempty"`
	Received_events_url string `json:"received_events_url,omitempty"`
	Node_id string `json:"node_id,omitempty"`
	Subscriptions_url string `json:"subscriptions_url,omitempty"`
	Type string `json:"type,omitempty"`
	Id int64 `json:"id"`
	User_view_type string `json:"user_view_type,omitempty"`
	Url string `json:"url,omitempty"`
	Email string `json:"email,omitempty"`
	Site_admin bool `json:"site_admin,omitempty"`
	Login string `json:"login"`
	Events_url string `json:"events_url,omitempty"`
	Repos_url string `json:"repos_url,omitempty"`
	Starred_url string `json:"starred_url,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	After string `json:"after"`
	Before string `json:"before"`
	Number int `json:"number"` // The pull request number.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Projects_v2_item GeneratedType `json:"projects_v2_item"` // An item belonging to a project
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Pending_cancellation_date string `json:"pending_cancellation_date,omitempty"` // The pending cancellation date for the seat, in `YYYY-MM-DD` format. This will be null unless the assignee's Copilot access has been canceled during the current billing cycle. If the seat has been cancelled, this corresponds to the start of the organization's next billing cycle.
	Plan_type string `json:"plan_type,omitempty"` // The Copilot plan of the organization, or the parent enterprise, when applicable.
	Updated_at string `json:"updated_at,omitempty"` // **Closing down notice:** This field is no longer relevant and is closing down. Use the `created_at` field to determine when the assignee was last granted access to GitHub Copilot. Timestamp of when the assignee's GitHub Copilot access was last updated, in ISO 8601 format.
	Created_at string `json:"created_at"` // Timestamp of when the assignee was last granted access to GitHub Copilot, in ISO 8601 format.
	Last_activity_at string `json:"last_activity_at,omitempty"` // Timestamp of user's last GitHub Copilot activity, in ISO 8601 format.
	Assigning_team interface{} `json:"assigning_team,omitempty"` // The team through which the assignee is granted access to GitHub Copilot, if applicable.
	Last_activity_editor string `json:"last_activity_editor,omitempty"` // Last editor that was used by the user for a GitHub Copilot completion.
	Assignee GeneratedType `json:"assignee,omitempty"` // A GitHub user.
}

// Webhooksusermannequin represents the Webhooksusermannequin schema from the OpenAPI specification
type Webhooksusermannequin struct {
	Type string `json:"type,omitempty"`
	Email string `json:"email,omitempty"`
	Login string `json:"login"`
	Gists_url string `json:"gists_url,omitempty"`
	Gravatar_id string `json:"gravatar_id,omitempty"`
	Node_id string `json:"node_id,omitempty"`
	Url string `json:"url,omitempty"`
	Name string `json:"name,omitempty"`
	Repos_url string `json:"repos_url,omitempty"`
	Id int `json:"id"`
	Organizations_url string `json:"organizations_url,omitempty"`
	Avatar_url string `json:"avatar_url,omitempty"`
	Events_url string `json:"events_url,omitempty"`
	Following_url string `json:"following_url,omitempty"`
	Site_admin bool `json:"site_admin,omitempty"`
	Deleted bool `json:"deleted,omitempty"`
	Starred_url string `json:"starred_url,omitempty"`
	Subscriptions_url string `json:"subscriptions_url,omitempty"`
	User_view_type string `json:"user_view_type,omitempty"`
	Followers_url string `json:"followers_url,omitempty"`
	Html_url string `json:"html_url,omitempty"`
	Received_events_url string `json:"received_events_url,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Commit_url string `json:"commit_url"`
	Node_id string `json:"node_id"`
	Project_card map[string]interface{} `json:"project_card,omitempty"`
	Id int `json:"id"`
	Url string `json:"url"`
	Commit_id string `json:"commit_id"`
	Created_at string `json:"created_at"`
	Event string `json:"event"`
	Performed_via_github_app Integration `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Actor GeneratedType `json:"actor"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Number int `json:"number"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Reason string `json:"reason,omitempty"`
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Pull_request map[string]interface{} `json:"pull_request"`
}

// Webhooksapprover represents the Webhooksapprover schema from the OpenAPI specification
type Webhooksapprover struct {
	Site_admin bool `json:"site_admin,omitempty"`
	Login string `json:"login,omitempty"`
	Repos_url string `json:"repos_url,omitempty"`
	Avatar_url string `json:"avatar_url,omitempty"`
	Received_events_url string `json:"received_events_url,omitempty"`
	Gravatar_id string `json:"gravatar_id,omitempty"`
	Events_url string `json:"events_url,omitempty"`
	Organizations_url string `json:"organizations_url,omitempty"`
	User_view_type string `json:"user_view_type,omitempty"`
	Subscriptions_url string `json:"subscriptions_url,omitempty"`
	Type string `json:"type,omitempty"`
	Url string `json:"url,omitempty"`
	Followers_url string `json:"followers_url,omitempty"`
	Id int `json:"id,omitempty"`
	Node_id string `json:"node_id,omitempty"`
	Starred_url string `json:"starred_url,omitempty"`
	Html_url string `json:"html_url,omitempty"`
	Following_url string `json:"following_url,omitempty"`
	Gists_url string `json:"gists_url,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id int `json:"id"` // The ID of the variant analysis.
	Query_pack_url string `json:"query_pack_url"` // The download url for the query pack.
	Skipped_repositories map[string]interface{} `json:"skipped_repositories,omitempty"` // Information about repositories that were skipped from processing. This information is only available to the user that initiated the variant analysis.
	Scanned_repositories []map[string]interface{} `json:"scanned_repositories,omitempty"`
	Status string `json:"status"`
	Updated_at string `json:"updated_at,omitempty"` // The date and time at which the variant analysis was last updated, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Completed_at string `json:"completed_at,omitempty"` // The date and time at which the variant analysis was completed, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ. Will be null if the variant analysis has not yet completed or this information is not available.
	Created_at string `json:"created_at,omitempty"` // The date and time at which the variant analysis was created, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Failure_reason string `json:"failure_reason,omitempty"` // The reason for a failure of the variant analysis. This is only available if the variant analysis has failed.
	Controller_repo GeneratedType `json:"controller_repo"` // A GitHub repository.
	Query_language string `json:"query_language"` // The language targeted by the CodeQL query
	Actions_workflow_run_id int `json:"actions_workflow_run_id,omitempty"` // The GitHub Actions workflow run used to execute this variant analysis. This is only available if the workflow run has started.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Backfill_scans []GeneratedType `json:"backfill_scans,omitempty"`
	Custom_pattern_backfill_scans []interface{} `json:"custom_pattern_backfill_scans,omitempty"`
	Incremental_scans []GeneratedType `json:"incremental_scans,omitempty"`
	Pattern_update_scans []GeneratedType `json:"pattern_update_scans,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Old_answer Webhooksanswer `json:"old_answer"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// Webhooksmilestone3 represents the Webhooksmilestone3 schema from the OpenAPI specification
type Webhooksmilestone3 struct {
	Closed_issues int `json:"closed_issues"`
	Open_issues int `json:"open_issues"`
	Updated_at string `json:"updated_at"`
	Title string `json:"title"` // The title of the milestone.
	Closed_at string `json:"closed_at"`
	State string `json:"state"` // The state of the milestone.
	Creator map[string]interface{} `json:"creator"`
	Created_at string `json:"created_at"`
	Number int `json:"number"` // The number of the milestone.
	Html_url string `json:"html_url"`
	Description string `json:"description"`
	Due_on string `json:"due_on"`
	Url string `json:"url"`
	Labels_url string `json:"labels_url"`
	Id int `json:"id"`
	Node_id string `json:"node_id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Property_values []string `json:"property_values"` // The values to match for the repository property
	Source string `json:"source,omitempty"` // The source of the repository property. Defaults to 'custom' if not specified.
	Name string `json:"name"` // The name of the repository property to target
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Commit_id string `json:"commit_id"`
	Commit_url string `json:"commit_url"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Created_at string `json:"created_at"`
	Milestone map[string]interface{} `json:"milestone"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Id int `json:"id"`
	Url string `json:"url"`
	Event string `json:"event"`
	Node_id string `json:"node_id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Type string `json:"type"`
	Href string `json:"href"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Alt_domain map[string]interface{} `json:"alt_domain,omitempty"`
	Domain map[string]interface{} `json:"domain,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sha string `json:"sha"` // SHA for the commit
	Author map[string]interface{} `json:"author"` // Identifying information for the git-user
	Node_id string `json:"node_id"`
	Url string `json:"url"`
	Event string `json:"event,omitempty"`
	Tree map[string]interface{} `json:"tree"`
	Verification map[string]interface{} `json:"verification"`
	Html_url string `json:"html_url"`
	Parents []map[string]interface{} `json:"parents"`
	Committer map[string]interface{} `json:"committer"` // Identifying information for the git-user
	Message string `json:"message"` // Message describing the purpose of the commit
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Bio string `json:"bio"`
	Url string `json:"url"`
	Subscriptions_url string `json:"subscriptions_url"`
	Node_id string `json:"node_id"`
	Notification_email string `json:"notification_email,omitempty"`
	Following int `json:"following"`
	Blog string `json:"blog"`
	Received_events_url string `json:"received_events_url"`
	Public_repos int `json:"public_repos"`
	Collaborators int `json:"collaborators,omitempty"`
	Disk_usage int `json:"disk_usage,omitempty"`
	Events_url string `json:"events_url"`
	Repos_url string `json:"repos_url"`
	Hireable bool `json:"hireable"`
	Owned_private_repos int `json:"owned_private_repos,omitempty"`
	Starred_url string `json:"starred_url"`
	Followers_url string `json:"followers_url"`
	Site_admin bool `json:"site_admin"`
	Avatar_url string `json:"avatar_url"`
	Plan map[string]interface{} `json:"plan,omitempty"`
	Email string `json:"email"`
	Following_url string `json:"following_url"`
	Html_url string `json:"html_url"`
	Updated_at string `json:"updated_at"`
	Id int64 `json:"id"`
	Private_gists int `json:"private_gists,omitempty"`
	Gists_url string `json:"gists_url"`
	Name string `json:"name"`
	Created_at string `json:"created_at"`
	Company string `json:"company"`
	Followers int `json:"followers"`
	Total_private_repos int `json:"total_private_repos,omitempty"`
	Organizations_url string `json:"organizations_url"`
	Login string `json:"login"`
	Location string `json:"location"`
	Twitter_username string `json:"twitter_username,omitempty"`
	Public_gists int `json:"public_gists"`
	Gravatar_id string `json:"gravatar_id"`
	Type string `json:"type"`
	User_view_type string `json:"user_view_type,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Active_caches_count int `json:"active_caches_count"` // The number of active caches in the repository.
	Active_caches_size_in_bytes int `json:"active_caches_size_in_bytes"` // The sum of the size in bytes of all the active cache items in the repository.
	Full_name string `json:"full_name"` // The repository owner and name for the cache usage being shown.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Registry_package map[string]interface{} `json:"registry_package"`
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Public_ips map[string]interface{} `json:"public_ips"` // Provides details of static public IP limits for GitHub-hosted Hosted Runners
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Members_count int `json:"members_count"`
	Organization GeneratedType `json:"organization"` // Team Organization
	Repos_count int `json:"repos_count"`
	Slug string `json:"slug"`
	Description string `json:"description"`
	Notification_setting string `json:"notification_setting,omitempty"` // The notification setting the team has set
	Privacy string `json:"privacy,omitempty"` // The level of privacy this team should have
	Repositories_url string `json:"repositories_url"`
	Permission string `json:"permission"` // Permission that the team will have for its repositories
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"` // URL for the team
	Id int `json:"id"` // Unique identifier of the team
	Parent GeneratedType `json:"parent,omitempty"` // Groups of organization members that gives permissions on specified repositories.
	Html_url string `json:"html_url"`
	Node_id string `json:"node_id"`
	Ldap_dn string `json:"ldap_dn,omitempty"` // Distinguished Name (DN) that team maps to within LDAP environment
	Members_url string `json:"members_url"`
	Name string `json:"name"` // Name of the team
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes,omitempty"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Dismissal_commit_id string `json:"dismissal_commit_id,omitempty"`
	Dismissal_message string `json:"dismissal_message"`
	Review_id int `json:"review_id"`
	State string `json:"state"`
}
