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
	Changes map[string]interface{} `json:"changes,omitempty"`
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
	Line int `json:"line"`
	Node_id string `json:"node_id"`
	Updated_at string `json:"updated_at"`
	Commit_id string `json:"commit_id"`
	Path string `json:"path"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Created_at string `json:"created_at"`
	Position int `json:"position"`
	Reactions GeneratedType `json:"reactions,omitempty"`
	Body string `json:"body"`
	Id int `json:"id"`
	Url string `json:"url"`
	User GeneratedType `json:"user"` // A GitHub user.
	Html_url string `json:"html_url"`
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

// Verification represents the Verification schema from the OpenAPI specification
type Verification struct {
	Verified_at string `json:"verified_at"`
	Payload string `json:"payload"`
	Reason string `json:"reason"`
	Signature string `json:"signature"`
	Verified bool `json:"verified"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Security_advisory map[string]interface{} `json:"security_advisory"` // The details of the security advisory, including summary, description, and severity.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repositories []GeneratedType `json:"repositories"` // A list of repositories that were skipped. This list may not include all repositories that were skipped. This is only available when the repository was found and the user has access to it.
	Repository_count int `json:"repository_count"` // The total number of repositories that were skipped for this reason.
}

// Codespace represents the Codespace schema from the OpenAPI specification
type Codespace struct {
	Git_status map[string]interface{} `json:"git_status"` // Details about the codespace's git repository.
	Recent_folders []string `json:"recent_folders"`
	State string `json:"state"` // State of this codespace.
	Url string `json:"url"` // API URL for this codespace.
	Retention_expires_at string `json:"retention_expires_at,omitempty"` // When a codespace will be auto-deleted based on the "retention_period_minutes" and "last_used_at"
	Billable_owner GeneratedType `json:"billable_owner"` // A GitHub user.
	Pulls_url string `json:"pulls_url"` // API URL for the Pull Request associated with this codespace, if any.
	Retention_period_minutes int `json:"retention_period_minutes,omitempty"` // Duration in minutes after codespace has gone idle in which it will be deleted. Must be integer minutes between 0 and 43200 (30 days).
	Idle_timeout_notice string `json:"idle_timeout_notice,omitempty"` // Text to show user when codespace idle timeout minutes has been overriden by an organization policy
	Display_name string `json:"display_name,omitempty"` // Display name for this codespace.
	Machine GeneratedType `json:"machine"` // A description of the machine powering a codespace.
	Machines_url string `json:"machines_url"` // API URL to access available alternate machine types for this codespace.
	Pending_operation_disabled_reason string `json:"pending_operation_disabled_reason,omitempty"` // Text to show user when codespace is disabled by a pending operation
	Pending_operation bool `json:"pending_operation,omitempty"` // Whether or not a codespace has a pending async operation. This would mean that the codespace is temporarily unavailable. The only thing that you can do with a codespace in this state is delete it.
	Name string `json:"name"` // Automatically generated name of this codespace.
	Last_used_at string `json:"last_used_at"` // Last known time this codespace was started.
	Repository GeneratedType `json:"repository"` // Minimal Repository
	Updated_at string `json:"updated_at"`
	Web_url string `json:"web_url"` // URL to access this codespace on the web.
	Stop_url string `json:"stop_url"` // API URL to stop this codespace.
	Last_known_stop_notice string `json:"last_known_stop_notice,omitempty"` // The text to display to a user when a codespace has been stopped for a potentially actionable reason.
	Publish_url string `json:"publish_url,omitempty"` // API URL to publish this codespace to a new repository.
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Id int64 `json:"id"`
	Devcontainer_path string `json:"devcontainer_path,omitempty"` // Path to devcontainer.json from repo root used to create Codespace.
	Environment_id string `json:"environment_id"` // UUID identifying this codespace's environment.
	Idle_timeout_minutes int `json:"idle_timeout_minutes"` // The number of minutes of inactivity after which this codespace will be automatically stopped.
	Prebuild bool `json:"prebuild"` // Whether the codespace was created from a prebuild.
	Created_at string `json:"created_at"`
	Runtime_constraints map[string]interface{} `json:"runtime_constraints,omitempty"`
	Start_url string `json:"start_url"` // API URL to start this codespace.
	Location string `json:"location"` // The initally assigned location of a new codespace.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Analyses_url string `json:"analyses_url,omitempty"` // The REST API URL for getting the analyses associated with the upload.
	Errors []string `json:"errors,omitempty"` // Any errors that ocurred during processing of the delivery.
	Processing_status string `json:"processing_status,omitempty"` // `pending` files have not yet been processed, while `complete` means results from the SARIF have been stored. `failed` files have either not been processed at all, or could only be partially processed.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Active_this_cycle int `json:"active_this_cycle,omitempty"` // The number of seats that have used Copilot during the current billing cycle.
	Added_this_cycle int `json:"added_this_cycle,omitempty"` // Seats added during the current billing cycle.
	Inactive_this_cycle int `json:"inactive_this_cycle,omitempty"` // The number of seats that have not used Copilot during the current billing cycle.
	Pending_cancellation int `json:"pending_cancellation,omitempty"` // The number of seats that are pending cancellation at the end of the current billing cycle.
	Pending_invitation int `json:"pending_invitation,omitempty"` // The number of users who have been invited to receive a Copilot seat through this organization.
	Total int `json:"total,omitempty"` // The total number of seats being billed for the organization as of the current billing cycle.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Role string `json:"role"` // The role of the user in the team.
	State string `json:"state"` // The state of the user's membership in the team.
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Projects_v2_item GeneratedType `json:"projects_v2_item"` // An item belonging to a project
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes Webhooksprojectchanges `json:"changes"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Projects_v2 GeneratedType `json:"projects_v2"` // A projects v2 project
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Project_column Webhooksprojectcolumn `json:"project_column"`
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Lock_reason string `json:"lock_reason,omitempty"`
	Id int64 `json:"id"`
	Node_id string `json:"node_id"`
	Requested_team Team `json:"requested_team,omitempty"` // Groups of organization members that gives permissions on specified repositories.
	Assignee GeneratedType `json:"assignee,omitempty"` // A GitHub user.
	Milestone GeneratedType `json:"milestone,omitempty"` // Issue Event Milestone
	Project_card GeneratedType `json:"project_card,omitempty"` // Issue Event Project Card
	Url string `json:"url"`
	Dismissed_review GeneratedType `json:"dismissed_review,omitempty"`
	Author_association string `json:"author_association,omitempty"` // How the author is associated with the repository.
	Rename GeneratedType `json:"rename,omitempty"` // Issue Event Rename
	Created_at string `json:"created_at"`
	Assigner GeneratedType `json:"assigner,omitempty"` // A GitHub user.
	Issue GeneratedType `json:"issue,omitempty"` // Issues are a great way to keep track of tasks, enhancements, and bugs for your projects.
	Commit_id string `json:"commit_id"`
	Commit_url string `json:"commit_url"`
	Requested_reviewer GeneratedType `json:"requested_reviewer,omitempty"` // A GitHub user.
	Label GeneratedType `json:"label,omitempty"` // Issue Event Label
	Review_requester GeneratedType `json:"review_requester,omitempty"` // A GitHub user.
	Performed_via_github_app GeneratedType `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Event string `json:"event"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
}

// Webhooksproject represents the Webhooksproject schema from the OpenAPI specification
type Webhooksproject struct {
	Body string `json:"body"` // Body of the project
	Creator map[string]interface{} `json:"creator"`
	Name string `json:"name"` // Name of the project
	Updated_at string `json:"updated_at"`
	Number int `json:"number"`
	Url string `json:"url"`
	Created_at string `json:"created_at"`
	Html_url string `json:"html_url"`
	Node_id string `json:"node_id"`
	State string `json:"state"` // State of the project; either 'open' or 'closed'
	Columns_url string `json:"columns_url"`
	Id int `json:"id"`
	Owner_url string `json:"owner_url"`
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
	Properties []GeneratedType `json:"properties"` // List of custom property names and associated values
	Repository_full_name string `json:"repository_full_name"`
	Repository_id int `json:"repository_id"`
	Repository_name string `json:"repository_name"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Accepted bool `json:"accepted"` // Whether the user has accepted the permissions defined by the devcontainer config
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
	Changes map[string]interface{} `json:"changes"`
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
	Has_discussions bool `json:"has_discussions,omitempty"` // Whether discussions are enabled.
	Allow_update_branch bool `json:"allow_update_branch,omitempty"` // Whether or not a pull request head branch that is behind its base branch can always be updated even if it is not required to be up to date before merging.
	Disabled bool `json:"disabled"` // Returns whether or not this repository disabled.
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"` // Whether to delete head branches when pull requests are merged
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Has_wiki bool `json:"has_wiki"` // Whether the wiki is enabled.
	Pushed_at string `json:"pushed_at"`
	Teams_url string `json:"teams_url"`
	Anonymous_access_enabled bool `json:"anonymous_access_enabled,omitempty"` // Whether anonymous git access is enabled for this repository
	Svn_url string `json:"svn_url"`
	Full_name string `json:"full_name"`
	Custom_properties map[string]interface{} `json:"custom_properties,omitempty"` // The custom properties that were defined for the repository. The keys are the custom property names, and the values are the corresponding custom property values.
	Subscribers_count int `json:"subscribers_count,omitempty"`
	Master_branch string `json:"master_branch,omitempty"`
	Is_template bool `json:"is_template,omitempty"` // Whether this repository acts as a template that can be used to generate new repositories.
	Name string `json:"name"` // The name of the repository.
	Contributors_url string `json:"contributors_url"`
	Merge_commit_title string `json:"merge_commit_title,omitempty"` // The default value for a merge commit title. - `PR_TITLE` - default to the pull request's title. - `MERGE_MESSAGE` - default to the classic title for a merge message (e.g., Merge pull request #123 from branch-name).
	Private bool `json:"private"` // Whether the repository is private or public.
	Allow_merge_commit bool `json:"allow_merge_commit,omitempty"` // Whether to allow merge commits for pull requests.
	Deployments_url string `json:"deployments_url"`
	Updated_at string `json:"updated_at"`
	Git_commits_url string `json:"git_commits_url"`
	Use_squash_pr_title_as_default bool `json:"use_squash_pr_title_as_default,omitempty"` // Whether a squash merge commit can use the pull request title as default. **This property is closing down. Please use `squash_merge_commit_title` instead.
	Created_at string `json:"created_at"`
	Allow_forking bool `json:"allow_forking,omitempty"` // Whether to allow forking this repo
	Releases_url string `json:"releases_url"`
	Fork bool `json:"fork"`
	Description string `json:"description"`
	Labels_url string `json:"labels_url"`
	Tags_url string `json:"tags_url"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub user.
	Branches_url string `json:"branches_url"`
	Statuses_url string `json:"statuses_url"`
	Forks_count int `json:"forks_count"`
	Contents_url string `json:"contents_url"`
	Has_downloads bool `json:"has_downloads"` // Whether downloads are enabled.
	Hooks_url string `json:"hooks_url"`
	Languages_url string `json:"languages_url"`
	Html_url string `json:"html_url"`
	Open_issues_count int `json:"open_issues_count"`
	Allow_squash_merge bool `json:"allow_squash_merge,omitempty"` // Whether to allow squash merges for pull requests.
	Ssh_url string `json:"ssh_url"`
	Url string `json:"url"`
	Forks_url string `json:"forks_url"`
	Merges_url string `json:"merges_url"`
	Subscribers_url string `json:"subscribers_url"`
	Node_id string `json:"node_id"`
	Notifications_url string `json:"notifications_url"`
	Default_branch string `json:"default_branch"` // The default branch of the repository.
	Id int64 `json:"id"` // Unique identifier of the repository
	Pulls_url string `json:"pulls_url"`
	Topics []string `json:"topics,omitempty"`
	Blobs_url string `json:"blobs_url"`
	Git_url string `json:"git_url"`
	Template_repository map[string]interface{} `json:"template_repository,omitempty"`
	Collaborators_url string `json:"collaborators_url"`
	Homepage string `json:"homepage"`
	Issue_comment_url string `json:"issue_comment_url"`
	Has_projects bool `json:"has_projects"` // Whether projects are enabled.
	Milestones_url string `json:"milestones_url"`
	Trees_url string `json:"trees_url"`
	Issue_events_url string `json:"issue_events_url"`
	Events_url string `json:"events_url"`
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"` // Whether to require contributors to sign off on web-based commits
	Mirror_url string `json:"mirror_url"`
	Allow_auto_merge bool `json:"allow_auto_merge,omitempty"` // Whether to allow Auto-merge to be used on pull requests.
	Temp_clone_token string `json:"temp_clone_token,omitempty"`
	Archived bool `json:"archived"` // Whether the repository is archived.
	Language string `json:"language"`
	Clone_url string `json:"clone_url"`
	Subscription_url string `json:"subscription_url"`
	Archive_url string `json:"archive_url"`
	Squash_merge_commit_title string `json:"squash_merge_commit_title,omitempty"` // The default value for a squash merge commit title: - `PR_TITLE` - default to the pull request's title. - `COMMIT_OR_PR_TITLE` - default to the commit's title (if only one commit) or the pull request's title (when more than one commit).
	Git_refs_url string `json:"git_refs_url"`
	Has_issues bool `json:"has_issues"` // Whether issues are enabled.
	Has_pages bool `json:"has_pages"`
	Allow_rebase_merge bool `json:"allow_rebase_merge,omitempty"` // Whether to allow rebase merges for pull requests.
	Merge_commit_message string `json:"merge_commit_message,omitempty"` // The default value for a merge commit message. - `PR_TITLE` - default to the pull request's title. - `PR_BODY` - default to the pull request's body. - `BLANK` - default to a blank commit message.
	Starred_at string `json:"starred_at,omitempty"`
	Size int `json:"size"` // The size of the repository, in kilobytes. Size is calculated hourly. When a repository is initially created, the size is 0.
	Downloads_url string `json:"downloads_url"`
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Squash_merge_commit_message string `json:"squash_merge_commit_message,omitempty"` // The default value for a squash merge commit message: - `PR_BODY` - default to the pull request's body. - `COMMIT_MESSAGES` - default to the branch's commit messages. - `BLANK` - default to a blank commit message.
	Watchers_count int `json:"watchers_count"`
	Forks int `json:"forks"`
	Open_issues int `json:"open_issues"`
	Watchers int `json:"watchers"`
	Stargazers_count int `json:"stargazers_count"`
	Stargazers_url string `json:"stargazers_url"`
	Git_tags_url string `json:"git_tags_url"`
	Visibility string `json:"visibility,omitempty"` // The repository visibility: public, private, or internal.
	License GeneratedType `json:"license"` // License Simple
	Assignees_url string `json:"assignees_url"`
	Network_count int `json:"network_count,omitempty"`
	Keys_url string `json:"keys_url"`
	Comments_url string `json:"comments_url"`
	Commits_url string `json:"commits_url"`
	Issues_url string `json:"issues_url"`
	Compare_url string `json:"compare_url"`
}

// Webhooksanswer represents the Webhooksanswer schema from the OpenAPI specification
type Webhooksanswer struct {
	Html_url string `json:"html_url"`
	Id int `json:"id"`
	Reactions map[string]interface{} `json:"reactions,omitempty"`
	Child_comment_count int `json:"child_comment_count"`
	Node_id string `json:"node_id"`
	Repository_url string `json:"repository_url"`
	Updated_at string `json:"updated_at"`
	User map[string]interface{} `json:"user"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Parent_id interface{} `json:"parent_id"`
	Body string `json:"body"`
	Created_at string `json:"created_at"`
	Discussion_id int `json:"discussion_id"`
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
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Zen string `json:"zen,omitempty"` // Random string of GitHub zen.
	Hook map[string]interface{} `json:"hook,omitempty"` // The webhook that is being pinged
	Hook_id int `json:"hook_id,omitempty"` // The ID of the webhook that triggered the ping.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Passing int `json:"passing"` // The number of students that have passed the assignment.
	Slug string `json:"slug"` // Sluggified name of the assignment.
	Feedback_pull_requests_enabled bool `json:"feedback_pull_requests_enabled"` // Whether feedback pull request will be created when a student accepts the assignment.
	Max_teams int `json:"max_teams"` // The maximum allowable teams for the assignment.
	TypeField string `json:"type"` // Whether it's a group assignment or individual assignment.
	Accepted int `json:"accepted"` // The number of students that have accepted the assignment.
	Starter_code_repository GeneratedType `json:"starter_code_repository"` // A GitHub repository view for Classroom
	Editor string `json:"editor"` // The selected editor for the assignment.
	Students_are_repo_admins bool `json:"students_are_repo_admins"` // Whether students are admins on created repository when a student accepts the assignment.
	Classroom Classroom `json:"classroom"` // A GitHub Classroom classroom
	Max_members int `json:"max_members"` // The maximum allowable members per team.
	Public_repo bool `json:"public_repo"` // Whether an accepted assignment creates a public repository.
	Title string `json:"title"` // Assignment title.
	Deadline string `json:"deadline"` // The time at which the assignment is due.
	Language string `json:"language"` // The programming language used in the assignment.
	Id int `json:"id"` // Unique identifier of the repository.
	Invite_link string `json:"invite_link"` // The link that a student can use to accept the assignment.
	Invitations_enabled bool `json:"invitations_enabled"` // Whether the invitation link is enabled. Visiting an enabled invitation link will accept the assignment.
	Submitted int `json:"submitted"` // The number of students that have submitted the assignment.
}

// Webhooksrelease represents the Webhooksrelease schema from the OpenAPI specification
type Webhooksrelease struct {
	Author map[string]interface{} `json:"author"`
	Draft bool `json:"draft"` // Whether the release is a draft or published
	Body string `json:"body"`
	Prerelease bool `json:"prerelease"` // Whether the release is identified as a prerelease or a full release.
	Created_at string `json:"created_at"`
	Tarball_url string `json:"tarball_url"`
	Id int `json:"id"`
	Published_at string `json:"published_at"`
	Assets []map[string]interface{} `json:"assets"`
	Html_url string `json:"html_url"`
	Target_commitish string `json:"target_commitish"` // Specifies the commitish value that determines where the Git tag is created from.
	Upload_url string `json:"upload_url"`
	Zipball_url string `json:"zipball_url"`
	Name string `json:"name"`
	Assets_url string `json:"assets_url"`
	Url string `json:"url"`
	Discussion_url string `json:"discussion_url,omitempty"`
	Node_id string `json:"node_id"`
	Reactions map[string]interface{} `json:"reactions,omitempty"`
	Tag_name string `json:"tag_name"` // The name of the tag.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Commit_url string `json:"commit_url"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Url string `json:"url"`
	Assignee GeneratedType `json:"assignee"` // A GitHub user.
	Created_at string `json:"created_at"`
	Id int `json:"id"`
	Commit_id string `json:"commit_id"`
	Event string `json:"event"`
	Node_id string `json:"node_id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Actor map[string]interface{} `json:"actor"` // The actor who updated the ruleset
	Updated_at string `json:"updated_at"`
	Version_id int `json:"version_id"` // The ID of the previous version of the ruleset
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Marketplace_purchase map[string]interface{} `json:"marketplace_purchase"`
	Organization_billing_email string `json:"organization_billing_email,omitempty"`
	TypeField string `json:"type"`
	Url string `json:"url"`
	Email string `json:"email,omitempty"`
	Id int `json:"id"`
	Login string `json:"login"`
	Marketplace_pending_change map[string]interface{} `json:"marketplace_pending_change,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Included_minutes int `json:"included_minutes"` // The amount of free GitHub Actions minutes available.
	Minutes_used_breakdown map[string]interface{} `json:"minutes_used_breakdown"`
	Total_minutes_used int `json:"total_minutes_used"` // The sum of the free and paid GitHub Actions minutes used.
	Total_paid_minutes_used int `json:"total_paid_minutes_used"` // The total paid GitHub Actions minutes used.
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
	Updated_at string `json:"updated_at"`
	Created_at string `json:"created_at"`
	Name string `json:"name"` // The name of the secret.
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

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Oid string `json:"oid"` // Full or abbreviated commit hash to reject
	Reason string `json:"reason,omitempty"` // Reason for restriction
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Created_at string `json:"created_at"`
	Last_edited_at string `json:"last_edited_at"`
	Body string `json:"body"` // The main text of the comment.
	Html_url string `json:"html_url"`
	Number int `json:"number"` // The unique sequence number of a team discussion comment.
	Author GeneratedType `json:"author"` // A GitHub user.
	Discussion_url string `json:"discussion_url"`
	Node_id string `json:"node_id"`
	Reactions GeneratedType `json:"reactions,omitempty"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	Body_html string `json:"body_html"`
	Body_version string `json:"body_version"` // The current version of the body content. If provided, this update operation will be rejected if the given version does not match the latest version on the server.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repositories_added []map[string]interface{} `json:"repositories_added"` // An array of repository objects, which were added to the installation.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repositories_removed []map[string]interface{} `json:"repositories_removed"` // An array of repository objects, which were removed from the installation.
	Repository_selection string `json:"repository_selection"` // Describe whether all repositories have been selected or there's a selection involved
	Action string `json:"action"`
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Requester Webhooksuser `json:"requester"`
	Installation Installation `json:"installation"` // Installation
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Commit_oid string `json:"commit_oid"` // The commit SHA of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Ref string `json:"ref"` // The Git reference of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Alert map[string]interface{} `json:"alert"` // The code scanning alert involved in the event.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Id int `json:"id"`
	Commit_id string `json:"commit_id"`
	Commit_url string `json:"commit_url"`
	Event string `json:"event"`
	Created_at string `json:"created_at"`
	Node_id string `json:"node_id"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Rename map[string]interface{} `json:"rename"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Projects_v2_item GeneratedType `json:"projects_v2_item"` // An item belonging to a project
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// Milestone represents the Milestone schema from the OpenAPI specification
type Milestone struct {
	Node_id string `json:"node_id"`
	Updated_at string `json:"updated_at"`
	Title string `json:"title"` // The title of the milestone.
	Creator GeneratedType `json:"creator"` // A GitHub user.
	Id int `json:"id"`
	Url string `json:"url"`
	Labels_url string `json:"labels_url"`
	Due_on string `json:"due_on"`
	Open_issues int `json:"open_issues"`
	Created_at string `json:"created_at"`
	State string `json:"state"` // The state of the milestone.
	Closed_at string `json:"closed_at"`
	Closed_issues int `json:"closed_issues"`
	Description string `json:"description"`
	Number int `json:"number"` // The number of the milestone.
	Html_url string `json:"html_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	TypeField GeneratedType `json:"type"` // The type of issue.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue Webhooksissue `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Total int `json:"total"`
	Completed int `json:"completed"`
	Percent_completed int `json:"percent_completed"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Alert GeneratedType `json:"alert"` // A Dependabot alert.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
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
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue Webhooksissue2 `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
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
	Action string `json:"action"`
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Projects_v2_status_update GeneratedType `json:"projects_v2_status_update"` // An status update belonging to a project
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Answer Webhooksanswer `json:"answer"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Action string `json:"action"`
	Comment Webhookscomment `json:"comment"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Changes map[string]interface{} `json:"changes"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Can_approve_pull_request_reviews bool `json:"can_approve_pull_request_reviews,omitempty"` // Whether GitHub Actions can approve pull requests. Enabling this can be a security risk.
	Default_workflow_permissions string `json:"default_workflow_permissions,omitempty"` // The default workflow permissions granted to the GITHUB_TOKEN when running workflows.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
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
	Has_discussions bool `json:"has_discussions,omitempty"` // Whether discussions are enabled.
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"` // Whether to require contributors to sign off on web-based commits
	Contents_url string `json:"contents_url"`
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"` // Whether to delete head branches when pull requests are merged
	Deployments_url string `json:"deployments_url"`
	Has_issues bool `json:"has_issues"` // Whether issues are enabled.
	Archived bool `json:"archived"` // Whether the repository is archived.
	Commits_url string `json:"commits_url"`
	Merges_url string `json:"merges_url"`
	Language string `json:"language"`
	Ssh_url string `json:"ssh_url"`
	Merge_commit_message string `json:"merge_commit_message,omitempty"` // The default value for a merge commit message. - `PR_TITLE` - default to the pull request's title. - `PR_BODY` - default to the pull request's body. - `BLANK` - default to a blank commit message.
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Subscription_url string `json:"subscription_url"`
	Assignees_url string `json:"assignees_url"`
	Is_template bool `json:"is_template,omitempty"` // Whether this repository acts as a template that can be used to generate new repositories.
	Contributors_url string `json:"contributors_url"`
	Id int64 `json:"id"` // Unique identifier of the repository
	Milestones_url string `json:"milestones_url"`
	Statuses_url string `json:"statuses_url"`
	Homepage string `json:"homepage"`
	Size int `json:"size"` // The size of the repository, in kilobytes. Size is calculated hourly. When a repository is initially created, the size is 0.
	Url string `json:"url"`
	Pushed_at string `json:"pushed_at"`
	Pulls_url string `json:"pulls_url"`
	Visibility string `json:"visibility,omitempty"` // The repository visibility: public, private, or internal.
	Has_pages bool `json:"has_pages"`
	Labels_url string `json:"labels_url"`
	Watchers_count int `json:"watchers_count"`
	Allow_forking bool `json:"allow_forking,omitempty"` // Whether to allow forking this repo
	Squash_merge_commit_message string `json:"squash_merge_commit_message,omitempty"` // The default value for a squash merge commit message: - `PR_BODY` - default to the pull request's body. - `COMMIT_MESSAGES` - default to the branch's commit messages. - `BLANK` - default to a blank commit message.
	Archive_url string `json:"archive_url"`
	Trees_url string `json:"trees_url"`
	Blobs_url string `json:"blobs_url"`
	Forks_url string `json:"forks_url"`
	Issues_url string `json:"issues_url"`
	Git_refs_url string `json:"git_refs_url"`
	Git_commits_url string `json:"git_commits_url"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub user.
	Watchers int `json:"watchers"`
	Releases_url string `json:"releases_url"`
	Downloads_url string `json:"downloads_url"`
	Use_squash_pr_title_as_default bool `json:"use_squash_pr_title_as_default,omitempty"` // Whether a squash merge commit can use the pull request title as default. **This property is closing down. Please use `squash_merge_commit_title` instead.
	Starred_at string `json:"starred_at,omitempty"`
	Topics []string `json:"topics,omitempty"`
	Mirror_url string `json:"mirror_url"`
	Issue_comment_url string `json:"issue_comment_url"`
	Open_issues int `json:"open_issues"`
	Forks int `json:"forks"`
	Languages_url string `json:"languages_url"`
	Temp_clone_token string `json:"temp_clone_token,omitempty"`
	Anonymous_access_enabled bool `json:"anonymous_access_enabled,omitempty"` // Whether anonymous git access is enabled for this repository
	Allow_merge_commit bool `json:"allow_merge_commit,omitempty"` // Whether to allow merge commits for pull requests.
	Description string `json:"description"`
	Html_url string `json:"html_url"`
	Updated_at string `json:"updated_at"`
	Network_count int `json:"network_count,omitempty"`
	Open_issues_count int `json:"open_issues_count"`
	Squash_merge_commit_title string `json:"squash_merge_commit_title,omitempty"` // The default value for a squash merge commit title: - `PR_TITLE` - default to the pull request's title. - `COMMIT_OR_PR_TITLE` - default to the commit's title (if only one commit) or the pull request's title (when more than one commit).
	Svn_url string `json:"svn_url"`
	Allow_update_branch bool `json:"allow_update_branch,omitempty"` // Whether or not a pull request head branch that is behind its base branch can always be updated even if it is not required to be up to date before merging.
	Notifications_url string `json:"notifications_url"`
	Subscribers_url string `json:"subscribers_url"`
	Default_branch string `json:"default_branch"` // The default branch of the repository.
	Tags_url string `json:"tags_url"`
	Collaborators_url string `json:"collaborators_url"`
	Events_url string `json:"events_url"`
	Private bool `json:"private"` // Whether the repository is private or public.
	Template_repository map[string]interface{} `json:"template_repository,omitempty"`
	Git_url string `json:"git_url"`
	Has_projects bool `json:"has_projects"` // Whether projects are enabled.
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Subscribers_count int `json:"subscribers_count,omitempty"`
	Git_tags_url string `json:"git_tags_url"`
	Comments_url string `json:"comments_url"`
	Fork bool `json:"fork"`
	License GeneratedType `json:"license"` // License Simple
	Node_id string `json:"node_id"`
	Issue_events_url string `json:"issue_events_url"`
	Stargazers_count int `json:"stargazers_count"`
	Disabled bool `json:"disabled"` // Returns whether or not this repository disabled.
	Stargazers_url string `json:"stargazers_url"`
	Allow_auto_merge bool `json:"allow_auto_merge,omitempty"` // Whether to allow Auto-merge to be used on pull requests.
	Clone_url string `json:"clone_url"`
	Allow_rebase_merge bool `json:"allow_rebase_merge,omitempty"` // Whether to allow rebase merges for pull requests.
	Teams_url string `json:"teams_url"`
	Allow_squash_merge bool `json:"allow_squash_merge,omitempty"` // Whether to allow squash merges for pull requests.
	Hooks_url string `json:"hooks_url"`
	Created_at string `json:"created_at"`
	Compare_url string `json:"compare_url"`
	Has_wiki bool `json:"has_wiki"` // Whether the wiki is enabled.
	Custom_properties map[string]interface{} `json:"custom_properties,omitempty"` // The custom properties that were defined for the repository. The keys are the custom property names, and the values are the corresponding custom property values.
	Branches_url string `json:"branches_url"`
	Has_downloads bool `json:"has_downloads"` // Whether downloads are enabled.
	Name string `json:"name"` // The name of the repository.
	Full_name string `json:"full_name"`
	Master_branch string `json:"master_branch,omitempty"`
	Forks_count int `json:"forks_count"`
	Keys_url string `json:"keys_url"`
	Merge_commit_title string `json:"merge_commit_title,omitempty"` // The default value for a merge commit title. - `PR_TITLE` - default to the pull request's title. - `MERGE_MESSAGE` - default to the classic title for a merge message (e.g., Merge pull request #123 from branch-name).
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Public_repos int `json:"public_repos,omitempty"`
	Avatar_url string `json:"avatar_url"`
	Repos_url string `json:"repos_url"`
	Starred_url string `json:"starred_url"`
	Organizations_url string `json:"organizations_url"`
	Gists_url string `json:"gists_url"`
	Suspended_at string `json:"suspended_at,omitempty"`
	Gravatar_id string `json:"gravatar_id"`
	Hireable bool `json:"hireable,omitempty"`
	Public_gists int `json:"public_gists,omitempty"`
	Followers_url string `json:"followers_url"`
	Html_url string `json:"html_url"`
	Url string `json:"url"`
	Company string `json:"company,omitempty"`
	Login string `json:"login"`
	Site_admin bool `json:"site_admin"`
	Updated_at string `json:"updated_at,omitempty"`
	Followers int `json:"followers,omitempty"`
	Node_id string `json:"node_id"`
	TypeField string `json:"type"`
	Following int `json:"following,omitempty"`
	Events_url string `json:"events_url"`
	Name string `json:"name,omitempty"`
	Received_events_url string `json:"received_events_url"`
	Following_url string `json:"following_url"`
	Text_matches []map[string]interface{} `json:"text_matches,omitempty"`
	Blog string `json:"blog,omitempty"`
	Location string `json:"location,omitempty"`
	Id int64 `json:"id"`
	Score float64 `json:"score"`
	Bio string `json:"bio,omitempty"`
	Created_at string `json:"created_at,omitempty"`
	Subscriptions_url string `json:"subscriptions_url"`
	User_view_type string `json:"user_view_type,omitempty"`
	Email string `json:"email,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Author GeneratedType `json:"author"` // A GitHub user.
	Total int `json:"total"`
	Weeks []map[string]interface{} `json:"weeks"`
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
	Updated_at string `json:"updated_at"`
	Description string `json:"description,omitempty"`
	Url string `json:"url"`
	Deleted_at string `json:"deleted_at,omitempty"`
	Html_url string `json:"html_url,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Id int `json:"id"` // Unique identifier of the package version.
	Created_at string `json:"created_at"`
	License string `json:"license,omitempty"`
	Name string `json:"name"` // The name of the package version.
	Package_html_url string `json:"package_html_url"`
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
	Id int64 `json:"id"` // The id of the check.
	Status string `json:"status"` // The phase of the lifecycle that the check is currently in. Statuses of waiting, requested, and pending are reserved for GitHub Actions check runs.
	Conclusion string `json:"conclusion"`
	Details_url string `json:"details_url"`
	Output map[string]interface{} `json:"output"`
	Check_suite map[string]interface{} `json:"check_suite"`
	External_id string `json:"external_id"`
	Name string `json:"name"` // The name of the check.
	Pull_requests []GeneratedType `json:"pull_requests"` // Pull requests that are open with a `head_sha` or `head_branch` that matches the check. The returned pull requests do not necessarily indicate pull requests that triggered the check.
	Completed_at string `json:"completed_at"`
	Node_id string `json:"node_id"`
	Head_sha string `json:"head_sha"` // The SHA of the commit that is being checked.
	App GeneratedType `json:"app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Deployment GeneratedType `json:"deployment,omitempty"` // A deployment created as the result of an Actions check run from a workflow that references an environment
	Started_at string `json:"started_at"`
	Url string `json:"url"`
	Html_url string `json:"html_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Name string `json:"name"` // The name of the private registry configuration.
	Registry_type string `json:"registry_type"` // The registry type.
	Selected_repository_ids []int `json:"selected_repository_ids,omitempty"` // An array of repository IDs that can access the organization private registry when `visibility` is set to `selected`.
	Updated_at string `json:"updated_at"`
	Username string `json:"username,omitempty"` // The username to use when authenticating with the private registry.
	Visibility string `json:"visibility"` // Which type of organization repositories have access to the private registry. `selected` means only the repositories specified by `selected_repository_ids` can access the private registry.
	Created_at string `json:"created_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Color string `json:"color"`
	Name string `json:"name"`
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
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Marketplace_purchase Webhooksmarketplacepurchase `json:"marketplace_purchase"`
	Action string `json:"action"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Previous_marketplace_purchase Webhookspreviousmarketplacepurchase `json:"previous_marketplace_purchase,omitempty"`
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Effective_date string `json:"effective_date"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Unit_name string `json:"unit_name"`
	Url string `json:"url"`
	Accounts_url string `json:"accounts_url"`
	Bullets []string `json:"bullets"`
	Id int `json:"id"`
	Monthly_price_in_cents int `json:"monthly_price_in_cents"`
	Number int `json:"number"`
	Price_model string `json:"price_model"`
	Yearly_price_in_cents int `json:"yearly_price_in_cents"`
	Description string `json:"description"`
	Has_free_trial bool `json:"has_free_trial"`
	Name string `json:"name"`
	State string `json:"state"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Ref_type string `json:"ref_type"` // The type of Git ref object created in the repository.
	Pusher_type string `json:"pusher_type"` // The pusher type for the event. Can be either `user` or a deploy key.
	Ref string `json:"ref"` // The [`git ref`](https://docs.github.com/rest/git/refs#get-a-reference) resource.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Description string `json:"description"` // The repository's current description.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Master_branch string `json:"master_branch"` // The name of the repository's default branch (usually `main`).
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Project Webhooksproject `json:"project"`
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Event string `json:"event"`
	Id int `json:"id"`
	Project_card map[string]interface{} `json:"project_card,omitempty"`
	Url string `json:"url"`
	Commit_id string `json:"commit_id"`
	Commit_url string `json:"commit_url"`
	Node_id string `json:"node_id"`
	Performed_via_github_app Integration `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Created_at string `json:"created_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Event string `json:"event"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Html_url string `json:"html_url"`
	Issue_url string `json:"issue_url"`
	Node_id string `json:"node_id"`
	Body string `json:"body,omitempty"` // Contents of the issue comment
	Id int `json:"id"` // Unique identifier of the issue comment
	Performed_via_github_app GeneratedType `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	User GeneratedType `json:"user"` // A GitHub user.
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Body_html string `json:"body_html,omitempty"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
	Body_text string `json:"body_text,omitempty"`
	Reactions GeneratedType `json:"reactions,omitempty"`
	Url string `json:"url"` // URL for the issue comment
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enabled bool `json:"enabled,omitempty"` // Whether public IP is enabled.
	Length int `json:"length,omitempty"` // The length of the IP prefix.
	Prefix string `json:"prefix,omitempty"` // The prefix for the public IP.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Blocked_user Webhooksuser `json:"blocked_user"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Payload string `json:"payload"` // A URL-encoded string of the check_run.rerequested JSON payload. The decoded payload is a JSON object.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Help string `json:"help,omitempty"` // Detailed documentation for the rule as GitHub Flavored Markdown.
	Name string `json:"name,omitempty"` // The name of the rule used to detect the alert.
	Security_severity_level string `json:"security_severity_level,omitempty"` // The security severity of the alert.
	Severity string `json:"severity,omitempty"` // The severity of the alert.
	Full_description string `json:"full_description,omitempty"` // A description of the rule used to detect the alert.
	Help_uri string `json:"help_uri,omitempty"` // A link to the documentation for the rule used to detect the alert.
	Tags []string `json:"tags,omitempty"` // A set of tags applicable for the rule.
	Id string `json:"id,omitempty"` // A unique identifier for the rule used to detect the alert.
	Description string `json:"description,omitempty"` // A short description of the rule used to detect the alert.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sbom map[string]interface{} `json:"sbom"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Billable map[string]interface{} `json:"billable"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	TypeField string `json:"type"`
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
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Commit_id string `json:"commit_id"`
	Created_at string `json:"created_at"`
	Event string `json:"event"`
	Node_id string `json:"node_id"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Commit_url string `json:"commit_url"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Requested_reviewer GeneratedType `json:"requested_reviewer,omitempty"` // A GitHub user.
	Review_requester GeneratedType `json:"review_requester"` // A GitHub user.
	Url string `json:"url"`
	Id int `json:"id"`
	Requested_team Team `json:"requested_team,omitempty"` // Groups of organization members that gives permissions on specified repositories.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Event string `json:"event"`
	Repository GeneratedType `json:"repository"` // Minimal Repository
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
	Id int `json:"id"` // The ID of the workflow run.
	Run_attempt int `json:"run_attempt,omitempty"` // Attempt number of the run, 1 for first attempt and higher if the workflow was re-run.
	Rerun_url string `json:"rerun_url"` // The URL to rerun the workflow run.
	Status string `json:"status"`
	Head_branch string `json:"head_branch"`
	Actor GeneratedType `json:"actor,omitempty"` // A GitHub user.
	Html_url string `json:"html_url"`
	Display_title string `json:"display_title"` // The event-specific title associated with the run or the run-name if set, or the value of `run-name` if it is set in the workflow.
	Logs_url string `json:"logs_url"` // The URL to download the logs for the workflow run.
	Pull_requests []GeneratedType `json:"pull_requests"` // Pull requests that are open with a `head_sha` or `head_branch` that matches the workflow run. The returned pull requests do not necessarily indicate pull requests that triggered the run.
	Name string `json:"name,omitempty"` // The name of the workflow run.
	Jobs_url string `json:"jobs_url"` // The URL to the jobs for the workflow run.
	Check_suite_url string `json:"check_suite_url"` // The URL to the associated check suite.
	Workflow_url string `json:"workflow_url"` // The URL to the workflow.
	Referenced_workflows []GeneratedType `json:"referenced_workflows,omitempty"`
	Artifacts_url string `json:"artifacts_url"` // The URL to the artifacts for the workflow run.
	Run_started_at string `json:"run_started_at,omitempty"` // The start time of the latest run. Resets on re-run.
	Cancel_url string `json:"cancel_url"` // The URL to cancel the workflow run.
	Previous_attempt_url string `json:"previous_attempt_url,omitempty"` // The URL to the previous attempted run of this workflow, if one exists.
	Check_suite_id int `json:"check_suite_id,omitempty"` // The ID of the associated check suite.
	Check_suite_node_id string `json:"check_suite_node_id,omitempty"` // The node ID of the associated check suite.
	Head_repository GeneratedType `json:"head_repository"` // Minimal Repository
	Path string `json:"path"` // The full path of the workflow
	Run_number int `json:"run_number"` // The auto incrementing run number for the workflow run.
	Head_sha string `json:"head_sha"` // The SHA of the head commit that points to the version of the workflow being run.
	Url string `json:"url"` // The URL to the workflow run.
	Workflow_id int `json:"workflow_id"` // The ID of the parent workflow.
	Conclusion string `json:"conclusion"`
	Head_commit GeneratedType `json:"head_commit"` // A commit.
	Node_id string `json:"node_id"`
	Triggering_actor GeneratedType `json:"triggering_actor,omitempty"` // A GitHub user.
	Head_repository_id int `json:"head_repository_id,omitempty"`
}

// Webhookssecurityadvisory represents the Webhookssecurityadvisory schema from the OpenAPI specification
type Webhookssecurityadvisory struct {
	Identifiers []map[string]interface{} `json:"identifiers"`
	Cwes []map[string]interface{} `json:"cwes"`
	Published_at string `json:"published_at"`
	Summary string `json:"summary"`
	Cvss map[string]interface{} `json:"cvss"`
	Ghsa_id string `json:"ghsa_id"`
	References []map[string]interface{} `json:"references"`
	Updated_at string `json:"updated_at"`
	Vulnerabilities []map[string]interface{} `json:"vulnerabilities"`
	Severity string `json:"severity"`
	Withdrawn_at string `json:"withdrawn_at"`
	Cvss_severities GeneratedType `json:"cvss_severities,omitempty"`
	Description string `json:"description"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Summary string `json:"summary"` // A short, plain text summary of the advisory.
	Cvss_severities GeneratedType `json:"cvss_severities,omitempty"`
	Epss GeneratedType `json:"epss,omitempty"` // The EPSS scores as calculated by the [Exploit Prediction Scoring System](https://www.first.org/epss).
	Description string `json:"description"` // A long-form Markdown-supported description of the advisory.
	Severity string `json:"severity"` // The severity of the advisory.
	Cvss map[string]interface{} `json:"cvss"` // Details for the advisory pertaining to the Common Vulnerability Scoring System.
	Ghsa_id string `json:"ghsa_id"` // The unique GitHub Security Advisory ID assigned to the advisory.
	References []map[string]interface{} `json:"references"` // Links to additional advisory information.
	Withdrawn_at string `json:"withdrawn_at"` // The time that the advisory was withdrawn in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Updated_at string `json:"updated_at"` // The time that the advisory was last modified in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Identifiers []map[string]interface{} `json:"identifiers"` // Values that identify this advisory among security information sources.
	Published_at string `json:"published_at"` // The time that the advisory was published in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Cwes []map[string]interface{} `json:"cwes"` // Details for the advisory pertaining to Common Weakness Enumeration.
	Vulnerabilities []GeneratedType `json:"vulnerabilities"` // Vulnerable version range information for the advisory.
	Cve_id string `json:"cve_id"` // The unique CVE ID assigned to the advisory.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Integration_id int `json:"integration_id,omitempty"` // The optional integration ID that this status check must originate from.
	Context string `json:"context"` // The status check context name that must be present on the commit.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Alerts_threshold string `json:"alerts_threshold"` // The severity level at which code scanning results that raise alerts block a reference update. For more information on alert severity levels, see "[About code scanning alerts](https://docs.github.com/code-security/code-scanning/managing-code-scanning-alerts/about-code-scanning-alerts#about-alert-severity-and-security-severity-levels)."
	Security_alerts_threshold string `json:"security_alerts_threshold"` // The severity level at which code scanning results that raise security alerts block a reference update. For more information on security severity levels, see "[About code scanning alerts](https://docs.github.com/code-security/code-scanning/managing-code-scanning-alerts/about-code-scanning-alerts#about-alert-severity-and-security-severity-levels)."
	Tool string `json:"tool"` // The name of a code scanning tool
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	TypeField string `json:"type"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	TypeField string `json:"type"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Effective_date string `json:"effective_date"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Marketplace_purchase Webhooksmarketplacepurchase `json:"marketplace_purchase"`
	Previous_marketplace_purchase Webhookspreviousmarketplacepurchase `json:"previous_marketplace_purchase,omitempty"`
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
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
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Security_advisory Webhookssecurityadvisory `json:"security_advisory"` // The details of the security advisory, including summary, description, and severity.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// Webhookscomment represents the Webhookscomment schema from the OpenAPI specification
type Webhookscomment struct {
	Body string `json:"body"`
	Reactions map[string]interface{} `json:"reactions"`
	Created_at string `json:"created_at"`
	Html_url string `json:"html_url"`
	Parent_id int `json:"parent_id"`
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	User map[string]interface{} `json:"user"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Child_comment_count int `json:"child_comment_count"`
	Discussion_id int `json:"discussion_id"`
	Repository_url string `json:"repository_url"`
	Updated_at string `json:"updated_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Merge_type string `json:"merge_type,omitempty"`
	Message string `json:"message,omitempty"`
	Base_branch string `json:"base_branch,omitempty"`
}

// Webhooksissuecomment represents the Webhooksissuecomment schema from the OpenAPI specification
type Webhooksissuecomment struct {
	Id int64 `json:"id"` // Unique identifier of the issue comment
	Html_url string `json:"html_url"`
	Issue_url string `json:"issue_url"`
	Node_id string `json:"node_id"`
	Updated_at string `json:"updated_at"`
	Performed_via_github_app Integration `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Url string `json:"url"` // URL for the issue comment
	User map[string]interface{} `json:"user"`
	Created_at string `json:"created_at"`
	Reactions map[string]interface{} `json:"reactions"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Body string `json:"body"` // Contents of the issue comment
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
	TypeField string `json:"type"`
	Parameters map[string]interface{} `json:"parameters,omitempty"`
}

// Deployment represents the Deployment schema from the OpenAPI specification
type Deployment struct {
	Node_id string `json:"node_id"`
	Repository_url string `json:"repository_url"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Created_at string `json:"created_at"`
	Production_environment bool `json:"production_environment,omitempty"` // Specifies if the given environment is one that end-users directly interact with. Default: false.
	Ref string `json:"ref"` // The ref to deploy. This can be a branch, tag, or sha.
	Updated_at string `json:"updated_at"`
	Creator GeneratedType `json:"creator"` // A GitHub user.
	Task string `json:"task"` // Parameter to specify a task to execute
	Description string `json:"description"`
	Environment string `json:"environment"` // Name for the target deployment environment.
	Id int64 `json:"id"` // Unique identifier of the deployment
	Sha string `json:"sha"`
	Original_environment string `json:"original_environment,omitempty"`
	Url string `json:"url"`
	Payload interface{} `json:"payload"`
	Transient_environment bool `json:"transient_environment,omitempty"` // Specifies if the given environment is will no longer exist at some point in the future. Default: false.
	Statuses_url string `json:"statuses_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id string `json:"id,omitempty"` // An identifier for the upload.
	Url string `json:"url,omitempty"` // The REST API URL for checking the status of the upload.
}

// Runner represents the Runner schema from the OpenAPI specification
type Runner struct {
	Runner_group_id int `json:"runner_group_id,omitempty"` // The ID of the runner group.
	Status string `json:"status"` // The status of the runner.
	Busy bool `json:"busy"`
	Ephemeral bool `json:"ephemeral,omitempty"`
	Id int `json:"id"` // The ID of the runner.
	Labels []GeneratedType `json:"labels"`
	Name string `json:"name"` // The name of the runner.
	Os string `json:"os"` // The Operating System of the runner.
}

// Release represents the Release schema from the OpenAPI specification
type Release struct {
	Tarball_url string `json:"tarball_url"`
	Author GeneratedType `json:"author"` // A GitHub user.
	Mentions_count int `json:"mentions_count,omitempty"`
	Assets []GeneratedType `json:"assets"`
	Body string `json:"body,omitempty"`
	Target_commitish string `json:"target_commitish"` // Specifies the commitish value that determines where the Git tag is created from.
	Url string `json:"url"`
	Created_at string `json:"created_at"`
	Name string `json:"name"`
	Html_url string `json:"html_url"`
	Upload_url string `json:"upload_url"`
	Id int `json:"id"`
	Body_html string `json:"body_html,omitempty"`
	Draft bool `json:"draft"` // true to create a draft (unpublished) release, false to create a published one.
	Node_id string `json:"node_id"`
	Published_at string `json:"published_at"`
	Tag_name string `json:"tag_name"` // The name of the tag.
	Reactions GeneratedType `json:"reactions,omitempty"`
	Zipball_url string `json:"zipball_url"`
	Body_text string `json:"body_text,omitempty"`
	Assets_url string `json:"assets_url"`
	Prerelease bool `json:"prerelease"` // Whether to identify the release as a prerelease or a full release.
	Discussion_url string `json:"discussion_url,omitempty"` // The URL of the release discussion.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Alert Webhooksalert `json:"alert"` // The security alert of the vulnerable dependency.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Page_url string `json:"page_url"` // The GitHub URL to get the associated wiki page
	Path string `json:"path"` // The file path of the wiki page
	Commit_sha string `json:"commit_sha"` // SHA-1 hash ID of the associated commit
	Commit_url string `json:"commit_url"` // The GitHub URL to get the associated wiki commit
	Start_column float64 `json:"start_column"` // The column at which the secret starts within the start line when the file is interpreted as 8-bit ASCII.
	Start_line float64 `json:"start_line"` // Line number at which the secret starts in the file
	End_column float64 `json:"end_column"` // The column at which the secret ends within the end line when the file is interpreted as 8-bit ASCII.
	Blob_sha string `json:"blob_sha"` // SHA-1 hash ID of the associated blob
	End_line float64 `json:"end_line"` // Line number at which the secret ends in the file
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Description string `json:"description"`
	Id int `json:"id"`
	Url string `json:"url"`
	Hooks_url string `json:"hooks_url"`
	Issues_url string `json:"issues_url"`
	Login string `json:"login"`
	Members_url string `json:"members_url"`
	Node_id string `json:"node_id"`
	Public_members_url string `json:"public_members_url"`
	Events_url string `json:"events_url"`
	Avatar_url string `json:"avatar_url"`
	Repos_url string `json:"repos_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Contexts []string `json:"contexts"`
	Contexts_url string `json:"contexts_url"`
	Strict bool `json:"strict"`
	Url string `json:"url"`
	Checks []map[string]interface{} `json:"checks"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Limit string `json:"limit"` // The type of GitHub user that can comment, open issues, or create pull requests while the interaction limit is in effect.
	Expiry string `json:"expiry,omitempty"` // The duration of the interaction restriction. Default: `one_day`.
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
	Id string `json:"id"` // The ID of the GitHub Pages deployment. This is the Git SHA of the deployed commit.
	Page_url string `json:"page_url"` // The URI to the deployed GitHub Pages.
	Preview_url string `json:"preview_url,omitempty"` // The URI to the deployed GitHub Pages preview.
	Status_url string `json:"status_url"` // The URI to monitor GitHub Pages deployment status.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Changes map[string]interface{} `json:"changes"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // Full Repository
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Ignored bool `json:"ignored"` // Determines if all notifications should be blocked from this repository.
	Reason string `json:"reason"`
	Repository_url string `json:"repository_url"`
	Subscribed bool `json:"subscribed"` // Determines if notifications should be received from this repository.
	Url string `json:"url"`
	Created_at string `json:"created_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Ruleset_source_type string `json:"ruleset_source_type,omitempty"` // The type of source for the ruleset that includes this rule.
	Ruleset_id int `json:"ruleset_id,omitempty"` // The ID of the ruleset that includes this rule.
	Ruleset_source string `json:"ruleset_source,omitempty"` // The name of the source of the ruleset that includes this rule.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Action string `json:"action"`
	Assignee Webhooksuser `json:"assignee"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Status string `json:"status,omitempty"` // The attachment status of the code security configuration on the repository.
	Configuration GeneratedType `json:"configuration,omitempty"` // A code security configuration
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Team Webhooksteam1 `json:"team"` // Groups of organization members that gives permissions on specified repositories.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository map[string]interface{} `json:"repository,omitempty"` // A git repository
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Contexts_url string `json:"contexts_url,omitempty"`
	Enforcement_level string `json:"enforcement_level,omitempty"`
	Strict bool `json:"strict,omitempty"`
	Url string `json:"url,omitempty"`
	Checks []map[string]interface{} `json:"checks"`
	Contexts []string `json:"contexts"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repo Repository `json:"repo"` // A repository on GitHub.
	Starred_at string `json:"starred_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Ref string `json:"ref"` // The Git reference of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Commit_oid string `json:"commit_oid"` // The commit SHA of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Alert map[string]interface{} `json:"alert"` // The code scanning alert involved in the event.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Number int `json:"number"` // The pull request number.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Action string `json:"action"`
	Label Webhookslabel `json:"label,omitempty"`
	Pull_request map[string]interface{} `json:"pull_request"`
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Comment Webhookscomment `json:"comment"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Commit_message string `json:"commit_message"` // Commit message for the merge commit.
	Commit_title string `json:"commit_title"` // Title for the merge commit message.
	Enabled_by GeneratedType `json:"enabled_by"` // A GitHub user.
	Merge_method string `json:"merge_method"` // The merge method to use.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Sponsorship Webhookssponsorship `json:"sponsorship"`
	Effective_date string `json:"effective_date,omitempty"` // The `pending_cancellation` and `pending_tier_change` event types will include the date the cancellation or tier change will take effect.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Changes Webhookschanges8 `json:"changes"`
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	State string `json:"state,omitempty"` // State of the latest export
	Branch string `json:"branch,omitempty"` // Name of the exported branch
	Completed_at string `json:"completed_at,omitempty"` // Completion time of the last export operation
	Export_url string `json:"export_url,omitempty"` // Url for fetching export details
	Html_url string `json:"html_url,omitempty"` // Web url for the exported branch
	Id string `json:"id,omitempty"` // Id for the export details
	Sha string `json:"sha,omitempty"` // Git commit SHA of the exported branch
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Actions []string `json:"actions,omitempty"`
	Pages []string `json:"pages,omitempty"`
	Verifiable_password_authentication bool `json:"verifiable_password_authentication"`
	Web []string `json:"web,omitempty"`
	Ssh_key_fingerprints map[string]interface{} `json:"ssh_key_fingerprints,omitempty"`
	Copilot []string `json:"copilot,omitempty"`
	Packages []string `json:"packages,omitempty"`
	Codespaces []string `json:"codespaces,omitempty"`
	Github_enterprise_importer []string `json:"github_enterprise_importer,omitempty"`
	Ssh_keys []string `json:"ssh_keys,omitempty"`
	Api []string `json:"api,omitempty"`
	Domains map[string]interface{} `json:"domains,omitempty"`
	Hooks []string `json:"hooks,omitempty"`
	Dependabot []string `json:"dependabot,omitempty"`
	Actions_macos []string `json:"actions_macos,omitempty"`
	Git []string `json:"git,omitempty"`
	Importer []string `json:"importer,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Release Webhooksrelease `json:"release"` // The [release](https://docs.github.com/rest/releases/releases/#get-a-release) object.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	TypeField string `json:"type"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Target_type string `json:"target_type,omitempty"` // The type of the code security configuration.
	Secret_scanning string `json:"secret_scanning,omitempty"` // The enablement status of secret scanning
	Url string `json:"url,omitempty"` // The URL of the configuration
	Dependency_graph_autosubmit_action_options map[string]interface{} `json:"dependency_graph_autosubmit_action_options,omitempty"` // Feature options for Automatic dependency submission
	Secret_scanning_delegated_bypass string `json:"secret_scanning_delegated_bypass,omitempty"` // The enablement status of secret scanning delegated bypass
	Secret_scanning_generic_secrets string `json:"secret_scanning_generic_secrets,omitempty"` // The enablement status of Copilot secret scanning
	Code_scanning_delegated_alert_dismissal string `json:"code_scanning_delegated_alert_dismissal,omitempty"` // The enablement status of code scanning delegated alert dismissal
	Code_scanning_options map[string]interface{} `json:"code_scanning_options,omitempty"` // Feature options for code scanning
	Description string `json:"description,omitempty"` // A description of the code security configuration
	Secret_scanning_non_provider_patterns string `json:"secret_scanning_non_provider_patterns,omitempty"` // The enablement status of secret scanning non-provider patterns
	Code_scanning_default_setup_options map[string]interface{} `json:"code_scanning_default_setup_options,omitempty"` // Feature options for code scanning default setup
	Dependabot_alerts string `json:"dependabot_alerts,omitempty"` // The enablement status of Dependabot alerts
	Dependency_graph_autosubmit_action string `json:"dependency_graph_autosubmit_action,omitempty"` // The enablement status of Automatic dependency submission
	Enforcement string `json:"enforcement,omitempty"` // The enforcement status for a security configuration
	Name string `json:"name,omitempty"` // The name of the code security configuration. Must be unique within the organization.
	Dependency_graph string `json:"dependency_graph,omitempty"` // The enablement status of Dependency Graph
	Secret_scanning_push_protection string `json:"secret_scanning_push_protection,omitempty"` // The enablement status of secret scanning push protection
	Updated_at string `json:"updated_at,omitempty"`
	Advanced_security string `json:"advanced_security,omitempty"` // The enablement status of GitHub Advanced Security
	Dependabot_security_updates string `json:"dependabot_security_updates,omitempty"` // The enablement status of Dependabot security updates
	Id int `json:"id,omitempty"` // The ID of the code security configuration
	Secret_scanning_delegated_bypass_options map[string]interface{} `json:"secret_scanning_delegated_bypass_options,omitempty"` // Feature options for secret scanning delegated bypass
	Secret_scanning_validity_checks string `json:"secret_scanning_validity_checks,omitempty"` // The enablement status of secret scanning validity checks
	Code_scanning_default_setup string `json:"code_scanning_default_setup,omitempty"` // The enablement status of code scanning default setup
	Html_url string `json:"html_url,omitempty"` // The URL of the configuration
	Private_vulnerability_reporting string `json:"private_vulnerability_reporting,omitempty"` // The enablement status of private vulnerability reporting
	Created_at string `json:"created_at,omitempty"`
	Secret_scanning_delegated_alert_dismissal string `json:"secret_scanning_delegated_alert_dismissal,omitempty"` // The enablement status of secret scanning delegated alert dismissal
}

// Traffic represents the Traffic schema from the OpenAPI specification
type Traffic struct {
	Count int `json:"count"`
	Timestamp string `json:"timestamp"`
	Uniques int `json:"uniques"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Pushed_at string `json:"pushed_at,omitempty"`
	Forks_url string `json:"forks_url"`
	Deployments_url string `json:"deployments_url"`
	Full_name string `json:"full_name"`
	Keys_url string `json:"keys_url"`
	Watchers_count int `json:"watchers_count,omitempty"`
	Description string `json:"description"`
	Git_commits_url string `json:"git_commits_url"`
	Pulls_url string `json:"pulls_url"`
	Teams_url string `json:"teams_url"`
	Comments_url string `json:"comments_url"`
	Labels_url string `json:"labels_url"`
	Updated_at string `json:"updated_at,omitempty"`
	Svn_url string `json:"svn_url,omitempty"`
	Commits_url string `json:"commits_url"`
	Disabled bool `json:"disabled,omitempty"`
	Clone_url string `json:"clone_url,omitempty"`
	Git_url string `json:"git_url,omitempty"`
	Homepage string `json:"homepage,omitempty"`
	Merges_url string `json:"merges_url"`
	Statuses_url string `json:"statuses_url"`
	Mirror_url string `json:"mirror_url,omitempty"`
	Subscribers_url string `json:"subscribers_url"`
	Branches_url string `json:"branches_url"`
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Custom_properties map[string]interface{} `json:"custom_properties,omitempty"` // The custom properties that were defined for the repository. The keys are the custom property names, and the values are the corresponding custom property values.
	Open_issues_count int `json:"open_issues_count,omitempty"`
	Size int `json:"size,omitempty"` // The size of the repository, in kilobytes. Size is calculated hourly. When a repository is initially created, the size is 0.
	Open_issues int `json:"open_issues,omitempty"`
	Security_and_analysis GeneratedType `json:"security_and_analysis,omitempty"`
	Issues_url string `json:"issues_url"`
	License map[string]interface{} `json:"license,omitempty"`
	Has_downloads bool `json:"has_downloads,omitempty"`
	Downloads_url string `json:"downloads_url"`
	Has_discussions bool `json:"has_discussions,omitempty"`
	Notifications_url string `json:"notifications_url"`
	Issue_events_url string `json:"issue_events_url"`
	Html_url string `json:"html_url"`
	Releases_url string `json:"releases_url"`
	Archive_url string `json:"archive_url"`
	Trees_url string `json:"trees_url"`
	Contents_url string `json:"contents_url"`
	Created_at string `json:"created_at,omitempty"`
	Fork bool `json:"fork"`
	Has_issues bool `json:"has_issues,omitempty"`
	Language string `json:"language,omitempty"`
	Allow_forking bool `json:"allow_forking,omitempty"`
	Temp_clone_token string `json:"temp_clone_token,omitempty"`
	Git_refs_url string `json:"git_refs_url"`
	Stargazers_count int `json:"stargazers_count,omitempty"`
	Watchers int `json:"watchers,omitempty"`
	Topics []string `json:"topics,omitempty"`
	Archived bool `json:"archived,omitempty"`
	Id int64 `json:"id"`
	Events_url string `json:"events_url"`
	Code_of_conduct GeneratedType `json:"code_of_conduct,omitempty"` // Code Of Conduct
	Issue_comment_url string `json:"issue_comment_url"`
	Subscription_url string `json:"subscription_url"`
	Compare_url string `json:"compare_url"`
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"`
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"`
	Contributors_url string `json:"contributors_url"`
	Name string `json:"name"`
	Has_wiki bool `json:"has_wiki,omitempty"`
	Tags_url string `json:"tags_url"`
	Languages_url string `json:"languages_url"`
	Is_template bool `json:"is_template,omitempty"`
	Default_branch string `json:"default_branch,omitempty"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Has_projects bool `json:"has_projects,omitempty"`
	Stargazers_url string `json:"stargazers_url"`
	Milestones_url string `json:"milestones_url"`
	Hooks_url string `json:"hooks_url"`
	Forks int `json:"forks,omitempty"`
	Visibility string `json:"visibility,omitempty"`
	Node_id string `json:"node_id"`
	Blobs_url string `json:"blobs_url"`
	Network_count int `json:"network_count,omitempty"`
	Url string `json:"url"`
	Subscribers_count int `json:"subscribers_count,omitempty"`
	Forks_count int `json:"forks_count,omitempty"`
	Assignees_url string `json:"assignees_url"`
	Git_tags_url string `json:"git_tags_url"`
	Ssh_url string `json:"ssh_url,omitempty"`
	Collaborators_url string `json:"collaborators_url"`
	Has_pages bool `json:"has_pages,omitempty"`
	Role_name string `json:"role_name,omitempty"`
	Private bool `json:"private"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Created_at string `json:"created_at"`
	Name string `json:"name"` // The name of the secret.
	Updated_at string `json:"updated_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Schemas []string `json:"schemas,omitempty"`
	Scimtype string `json:"scimType,omitempty"`
	Status int `json:"status,omitempty"`
	Detail string `json:"detail,omitempty"`
	Documentation_url string `json:"documentation_url,omitempty"`
	Message string `json:"message,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Secret_types []string `json:"secret_types,omitempty"` // List of patterns that were updated. This will be empty for normal backfill scans or custom pattern updates
	Source string `json:"source"` // What type of content was scanned
	Started_at string `json:"started_at"` // The time that the alert was resolved in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	TypeField string `json:"type"` // What type of scan was completed
	Action string `json:"action"`
	Completed_at string `json:"completed_at"` // The time that the alert was resolved in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Custom_pattern_name string `json:"custom_pattern_name,omitempty"` // If the scan was triggered by a custom pattern update, this will be the name of the pattern that was updated
	Custom_pattern_scope string `json:"custom_pattern_scope,omitempty"` // If the scan was triggered by a custom pattern update, this will be the scope of the pattern that was updated
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Links map[string]interface{} `json:"_links"`
	Download_url string `json:"download_url"`
	Html_url string `json:"html_url"`
	Sha string `json:"sha"`
	Url string `json:"url"`
	Git_url string `json:"git_url"`
	Name string `json:"name"`
	Path string `json:"path"`
	Size int `json:"size"`
	TypeField string `json:"type"`
	Submodule_git_url string `json:"submodule_git_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Account interface{} `json:"account"`
	Created_at string `json:"created_at"`
	Id int `json:"id"` // Unique identifier of the request installation.
	Node_id string `json:"node_id,omitempty"`
	Requester GeneratedType `json:"requester"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Avatar_url string `json:"avatar_url"`
	Created_at string `json:"created_at"`
	Name string `json:"name"` // The name of the enterprise.
	Node_id string `json:"node_id"`
	Id int `json:"id"` // Unique identifier of the enterprise
	Website_url string `json:"website_url,omitempty"` // The enterprise's website URL.
	Description string `json:"description,omitempty"` // A short description of the enterprise.
	Slug string `json:"slug"` // The slug url identifier for the enterprise.
	Html_url string `json:"html_url"`
	Updated_at string `json:"updated_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Editors []map[string]interface{} `json:"editors,omitempty"`
	Languages []map[string]interface{} `json:"languages,omitempty"` // Code completion metrics for active languages.
	Total_engaged_users int `json:"total_engaged_users,omitempty"` // Number of users who accepted at least one Copilot code suggestion, across all active editors. Includes both full and partial acceptances.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Deliveries_url string `json:"deliveries_url,omitempty"`
	Events []string `json:"events"`
	Ping_url string `json:"ping_url"`
	TypeField string `json:"type"`
	Name string `json:"name"`
	Active bool `json:"active"`
	Config map[string]interface{} `json:"config"`
	Id int `json:"id"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	Created_at string `json:"created_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url"`
	Commit_url string `json:"commit_url"`
	Milestone map[string]interface{} `json:"milestone"`
	Commit_id string `json:"commit_id"`
	Created_at string `json:"created_at"`
	Node_id string `json:"node_id"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Event string `json:"event"`
	Id int `json:"id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Repository_ruleset GeneratedType `json:"repository_ruleset"` // A set of rules to apply when specified conditions are met.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Status string `json:"status"`
}

// Link represents the Link schema from the OpenAPI specification
type Link struct {
	Href string `json:"href"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Name string `json:"name"` // The name pattern that branches must match in order to deploy to the environment. Wildcard characters will not match `/`. For example, to match branches that begin with `release/` and contain an additional single slash, use `release/*/*`. For more information about pattern matching syntax, see the [Ruby File.fnmatch documentation](https://ruby-doc.org/core-2.5.1/File.html#method-c-fnmatch).
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Hook map[string]interface{} `json:"hook"` // The deleted webhook. This will contain different keys based on the type of webhook it is: repository, organization, business, app, or GitHub Marketplace.
	Hook_id int `json:"hook_id"` // The id of the modified webhook.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Created_at string `json:"created_at"`
	Html_url string `json:"html_url"`
	Assignee GeneratedType `json:"assignee"` // A GitHub user.
	Number int `json:"number"`
	Reactions GeneratedType `json:"reactions,omitempty"`
	Title string `json:"title"`
	Assignees []GeneratedType `json:"assignees,omitempty"`
	Labels_url string `json:"labels_url"`
	Labels []map[string]interface{} `json:"labels"`
	Locked bool `json:"locked"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Score float64 `json:"score"`
	Events_url string `json:"events_url"`
	Draft bool `json:"draft,omitempty"`
	Body_text string `json:"body_text,omitempty"`
	Node_id string `json:"node_id"`
	Body string `json:"body,omitempty"`
	Comments int `json:"comments"`
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Timeline_url string `json:"timeline_url,omitempty"`
	Id int64 `json:"id"`
	Updated_at string `json:"updated_at"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Repository_url string `json:"repository_url"`
	Closed_at string `json:"closed_at"`
	Sub_issues_summary map[string]interface{} `json:"sub_issues_summary,omitempty"`
	Comments_url string `json:"comments_url"`
	Url string `json:"url"`
	Body_html string `json:"body_html,omitempty"`
	Text_matches []map[string]interface{} `json:"text_matches,omitempty"`
	State string `json:"state"`
	User GeneratedType `json:"user"` // A GitHub user.
	Pull_request map[string]interface{} `json:"pull_request,omitempty"`
	State_reason string `json:"state_reason,omitempty"`
	TypeField GeneratedType `json:"type,omitempty"` // The type of issue.
	Milestone GeneratedType `json:"milestone"` // A collection of related issues and pull requests.
	Active_lock_reason string `json:"active_lock_reason,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url"`
	Html_url string `json:"html_url"`
	Description string `json:"description"`
	Parent GeneratedType `json:"parent"` // Groups of organization members that gives permissions on specified repositories.
	Privacy string `json:"privacy,omitempty"`
	Node_id string `json:"node_id"`
	Repositories_url string `json:"repositories_url"`
	Name string `json:"name"`
	Members_url string `json:"members_url"`
	Permission string `json:"permission"`
	Slug string `json:"slug"`
	Notification_setting string `json:"notification_setting,omitempty"`
	Assignment string `json:"assignment,omitempty"` // Determines if the team has a direct, indirect, or mixed relationship to a role
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Id int `json:"id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Merge_group GeneratedType `json:"merge_group"` // A group of pull requests that the merge queue has grouped together to be merged.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// Webhooksworkflowjobrun represents the Webhooksworkflowjobrun schema from the OpenAPI specification
type Webhooksworkflowjobrun struct {
	Id int `json:"id"`
	Name interface{} `json:"name"`
	Status string `json:"status"`
	Updated_at string `json:"updated_at"`
	Conclusion interface{} `json:"conclusion"`
	Created_at string `json:"created_at"`
	Environment string `json:"environment"`
	Html_url string `json:"html_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Comment map[string]interface{} `json:"comment"` // The [commit comment](${externalDocsUpapp/api/description/components/schemas/webhooks/issue-comment-created.yamlrl}/rest/commits/comments#get-a-commit-comment) resource.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"` // The action performed. Can be `created`.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Milestone Webhooksmilestone `json:"milestone"` // A collection of related issues and pull requests.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Payload string `json:"payload"` // A URL-encoded string of the ping JSON payload. The decoded payload is a JSON object.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Reason string `json:"reason"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Number int `json:"number"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Html_url string `json:"html_url"`
	Url string `json:"url"`
	Conclusion string `json:"conclusion"`
	App Integration `json:"app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Status string `json:"status"` // The phase of the lifecycle that the check is currently in.
	Output map[string]interface{} `json:"output"`
	Started_at string `json:"started_at"`
	Node_id string `json:"node_id"`
	Details_url string `json:"details_url"`
	Deployment GeneratedType `json:"deployment,omitempty"` // A deployment created as the result of an Actions check run from a workflow that references an environment
	Check_suite GeneratedType `json:"check_suite"` // A suite of checks performed on the code of a given code change
	Id int `json:"id"` // The id of the check.
	External_id string `json:"external_id"`
	Completed_at string `json:"completed_at"`
	Head_sha string `json:"head_sha"` // The SHA of the commit that is being checked.
	Name string `json:"name"` // The name of the check.
	Pull_requests []GeneratedType `json:"pull_requests"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository_name map[string]interface{} `json:"repository_name"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Object map[string]interface{} `json:"object"`
	Sha string `json:"sha"`
	Tag string `json:"tag"` // Name of the tag
	Tagger map[string]interface{} `json:"tagger"`
	Url string `json:"url"` // URL for the tag
	Verification Verification `json:"verification,omitempty"`
	Message string `json:"message"` // Message describing the purpose of the tag
	Node_id string `json:"node_id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sponsorship Webhookssponsorship `json:"sponsorship"`
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Key Webhooksdeploykey `json:"key"` // The [`deploy key`](https://docs.github.com/rest/deploy-keys/deploy-keys#get-a-deploy-key) resource.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Created_at string `json:"created_at"` // The date and time at which the secret was created, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Name string `json:"name"` // The name of the secret
	Selected_repositories_url string `json:"selected_repositories_url,omitempty"` // The API URL at which the list of repositories this secret is visible to can be retrieved
	Updated_at string `json:"updated_at"` // The date and time at which the secret was created, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Visibility string `json:"visibility"` // The type of repositories in the organization that the secret is visible to
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository_id int `json:"repository_id,omitempty"` // The ID of the repository associated with the rule evaluation.
	Repository_name string `json:"repository_name,omitempty"` // The name of the repository without the `.git` extension.
	Evaluation_result string `json:"evaluation_result,omitempty"` // The result of the rule evaluations for rules with the `active` and `evaluate` enforcement statuses, demonstrating whether rules would pass or fail if all rules in the rule suite were `active`. Null if no rules with `evaluate` enforcement status were run.
	Pushed_at string `json:"pushed_at,omitempty"`
	After_sha string `json:"after_sha,omitempty"` // The last commit sha in the push evaluation.
	Result string `json:"result,omitempty"` // The result of the rule evaluations for rules with the `active` enforcement status.
	Actor_name string `json:"actor_name,omitempty"` // The handle for the GitHub user account.
	Before_sha string `json:"before_sha,omitempty"` // The first commit sha before the push evaluation.
	Ref string `json:"ref,omitempty"` // The ref name that the evaluation ran on.
	Rule_evaluations []map[string]interface{} `json:"rule_evaluations,omitempty"` // Details on the evaluated rules.
	Actor_id int `json:"actor_id,omitempty"` // The number that identifies the user.
	Id int `json:"id,omitempty"` // The unique identifier of the rule insight.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Project Webhooksproject `json:"project"`
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Changes interface{} `json:"changes,omitempty"` // The changes made to the item may involve modifications in the item's fields and draft issue body. It includes altered values for text, number, date, single select, and iteration fields, along with the GraphQL node ID of the changed field.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Projects_v2_item GeneratedType `json:"projects_v2_item"` // An item belonging to a project
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Dependency map[string]interface{} `json:"dependency"` // Details for the vulnerable dependency.
	Dismissed_by GeneratedType `json:"dismissed_by"` // A GitHub user.
	Dismissed_comment string `json:"dismissed_comment"` // An optional comment associated with the alert's dismissal.
	Security_advisory GeneratedType `json:"security_advisory"` // Details for the GitHub Security Advisory.
	State string `json:"state"` // The state of the Dependabot alert.
	Updated_at string `json:"updated_at"` // The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Dismissed_at string `json:"dismissed_at"` // The time that the alert was dismissed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Security_vulnerability GeneratedType `json:"security_vulnerability"` // Details pertaining to one vulnerable version range for the advisory.
	Dismissed_reason string `json:"dismissed_reason"` // The reason that the alert was dismissed.
	Fixed_at string `json:"fixed_at"` // The time that the alert was no longer detected and was considered fixed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Html_url string `json:"html_url"` // The GitHub URL of the alert resource.
	Url string `json:"url"` // The REST API URL of the alert resource.
	Number int `json:"number"` // The security alert number.
	Auto_dismissed_at string `json:"auto_dismissed_at,omitempty"` // The time that the alert was auto-dismissed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Created_at string `json:"created_at"` // The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	TypeField GeneratedType `json:"type"` // The type of issue.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue Webhooksissue `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Avatar_url string `json:"avatar_url"`
	Events_url string `json:"events_url"`
	Login string `json:"login"`
	Members_url string `json:"members_url"`
	Node_id string `json:"node_id"`
	Repos_url string `json:"repos_url"`
	Id int `json:"id"`
	Description string `json:"description"`
	Public_members_url string `json:"public_members_url"`
	Hooks_url string `json:"hooks_url"`
	Issues_url string `json:"issues_url"`
	Url string `json:"url"`
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
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Commit_id string `json:"commit_id"`
	Event string `json:"event"`
	Url string `json:"url"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Created_at string `json:"created_at"`
	State_reason string `json:"state_reason,omitempty"`
	Commit_url string `json:"commit_url"`
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
	Duration int `json:"duration"`
	ErrorField map[string]interface{} `json:"error"`
	Pusher GeneratedType `json:"pusher"` // A GitHub user.
	Status string `json:"status"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	Commit string `json:"commit"`
	Created_at string `json:"created_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Issue_comment_url string `json:"issue_comment_url"` // The API URL to get the issue comment where the secret was detected.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url,omitempty"`
	Bypass_pull_request_allowances map[string]interface{} `json:"bypass_pull_request_allowances,omitempty"` // Allow specific users, teams, or apps to bypass pull request requirements.
	Dismiss_stale_reviews bool `json:"dismiss_stale_reviews"`
	Dismissal_restrictions map[string]interface{} `json:"dismissal_restrictions,omitempty"`
	Require_code_owner_reviews bool `json:"require_code_owner_reviews"`
	Require_last_push_approval bool `json:"require_last_push_approval,omitempty"` // Whether the most recent push must be approved by someone other than the person who pushed it.
	Required_approving_review_count int `json:"required_approving_review_count,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Size_gb int `json:"size_gb"` // Image size in GB.
	Source string `json:"source"` // The image provider.
	Display_name string `json:"display_name"` // Display name for this image.
	Id string `json:"id"` // The ID of the image. Use this ID for the `image` parameter when creating a new larger runner.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Commit_url string `json:"commit_url"`
	Created_at string `json:"created_at"`
	Lock_reason string `json:"lock_reason"`
	Node_id string `json:"node_id"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Commit_id string `json:"commit_id"`
	Id int `json:"id"`
	Event string `json:"event"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Sub_issue_id float64 `json:"sub_issue_id"` // The ID of the sub-issue.
	Sub_issue_repo Repository `json:"sub_issue_repo"` // A repository on GitHub.
	Action string `json:"action"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Parent_issue Issue `json:"parent_issue"` // Issues are a great way to keep track of tasks, enhancements, and bugs for your projects.
	Sub_issue Issue `json:"sub_issue"` // Issues are a great way to keep track of tasks, enhancements, and bugs for your projects.
	Parent_issue_id float64 `json:"parent_issue_id"` // The ID of the parent issue.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	TypeField string `json:"type"` // The type of credit the user is receiving.
	User GeneratedType `json:"user"` // A GitHub user.
	State string `json:"state"` // The state of the user's acceptance of the credit.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Event string `json:"event,omitempty"`
	Node_id string `json:"node_id,omitempty"`
	Comments []GeneratedType `json:"comments,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Accessible_repositories []GeneratedType `json:"accessible_repositories,omitempty"`
	Default_level string `json:"default_level,omitempty"` // The default repository access level for Dependabot updates.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Uniques int `json:"uniques"`
	Views []Traffic `json:"views"`
	Count int `json:"count"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url"`
	Gravatar_id string `json:"gravatar_id"`
	Events_url string `json:"events_url"`
	Organizations_url string `json:"organizations_url"`
	Site_admin bool `json:"site_admin"`
	Starred_url string `json:"starred_url"`
	Received_events_url string `json:"received_events_url"`
	Inherited_from []GeneratedType `json:"inherited_from,omitempty"` // Team the user has gotten the role through
	Starred_at string `json:"starred_at,omitempty"`
	Email string `json:"email,omitempty"`
	Name string `json:"name,omitempty"`
	Node_id string `json:"node_id"`
	Gists_url string `json:"gists_url"`
	Html_url string `json:"html_url"`
	Repos_url string `json:"repos_url"`
	Assignment string `json:"assignment,omitempty"` // Determines if the user has a direct, indirect, or mixed relationship to a role
	Subscriptions_url string `json:"subscriptions_url"`
	Followers_url string `json:"followers_url"`
	Login string `json:"login"`
	Avatar_url string `json:"avatar_url"`
	TypeField string `json:"type"`
	User_view_type string `json:"user_view_type,omitempty"`
	Following_url string `json:"following_url"`
	Id int `json:"id"`
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

// Root represents the Root schema from the OpenAPI specification
type Root struct {
	Public_gists_url string `json:"public_gists_url"`
	Emails_url string `json:"emails_url"`
	Events_url string `json:"events_url"`
	Notifications_url string `json:"notifications_url"`
	Following_url string `json:"following_url"`
	Starred_gists_url string `json:"starred_gists_url"`
	Rate_limit_url string `json:"rate_limit_url"`
	Label_search_url string `json:"label_search_url"`
	Current_user_repositories_url string `json:"current_user_repositories_url"`
	Organization_repositories_url string `json:"organization_repositories_url"`
	Followers_url string `json:"followers_url"`
	User_search_url string `json:"user_search_url"`
	Keys_url string `json:"keys_url"`
	Emojis_url string `json:"emojis_url"`
	Code_search_url string `json:"code_search_url"`
	User_organizations_url string `json:"user_organizations_url"`
	Organization_teams_url string `json:"organization_teams_url"`
	Issue_search_url string `json:"issue_search_url"`
	Repository_search_url string `json:"repository_search_url"`
	Organization_url string `json:"organization_url"`
	User_url string `json:"user_url"`
	Authorizations_url string `json:"authorizations_url"`
	Commit_search_url string `json:"commit_search_url"`
	Issues_url string `json:"issues_url"`
	Hub_url string `json:"hub_url,omitempty"`
	Starred_url string `json:"starred_url"`
	Current_user_url string `json:"current_user_url"`
	Topic_search_url string `json:"topic_search_url,omitempty"`
	Current_user_authorizations_html_url string `json:"current_user_authorizations_html_url"`
	User_repositories_url string `json:"user_repositories_url"`
	Gists_url string `json:"gists_url"`
	Feeds_url string `json:"feeds_url"`
	Repository_url string `json:"repository_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Errors []map[string]interface{} `json:"errors"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Projects_v2_status_update GeneratedType `json:"projects_v2_status_update"` // An status update belonging to a project
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Requested_action map[string]interface{} `json:"requested_action,omitempty"` // The action requested by the user.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Check_run GeneratedType `json:"check_run"` // A check performed on the code of a given code change
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Member Webhooksuser `json:"member"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"` // The changes to the collaborator permissions
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id int `json:"id"` // Unique identifier of the fine-grained personal access token grant. The `pat_id` used to get details about an approved fine-grained personal access token.
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Permissions map[string]interface{} `json:"permissions"` // Permissions requested, categorized by type of permission.
	Token_name string `json:"token_name"` // The name given to the user's token. This field can also be found in an organization's settings page for Active Tokens.
	Access_granted_at string `json:"access_granted_at"` // Date and time when the fine-grained personal access token was approved to access the organization.
	Repository_selection string `json:"repository_selection"` // Type of repository selection requested.
	Token_expires_at string `json:"token_expires_at"` // Date and time when the associated fine-grained personal access token expires.
	Token_id int `json:"token_id"` // Unique identifier of the user's token. This field can also be found in audit log events and the organization's settings for their PAT grants.
	Repositories_url string `json:"repositories_url"` // URL to the list of repositories the fine-grained personal access token can access. Only follow when `repository_selection` is `subset`.
	Token_expired bool `json:"token_expired"` // Whether the associated fine-grained personal access token has expired.
	Token_last_used_at string `json:"token_last_used_at"` // Date and time when the associated fine-grained personal access token was last used for authentication.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Body string `json:"body"` // The text of the review.
	Body_html string `json:"body_html,omitempty"`
	Id int `json:"id"` // Unique identifier of the review
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Html_url string `json:"html_url"`
	Submitted_at string `json:"submitted_at,omitempty"`
	Body_text string `json:"body_text,omitempty"`
	Commit_id string `json:"commit_id"` // A commit SHA for the review.
	Event string `json:"event"`
	Pull_request_url string `json:"pull_request_url"`
	Node_id string `json:"node_id"`
	State string `json:"state"`
	User GeneratedType `json:"user"` // A GitHub user.
	Links map[string]interface{} `json:"_links"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Total_private_repos int `json:"total_private_repos,omitempty"`
	Deploy_keys_enabled_for_repositories bool `json:"deploy_keys_enabled_for_repositories,omitempty"` // Controls whether or not deploy keys may be added and used for repositories in the organization.
	Has_repository_projects bool `json:"has_repository_projects"`
	Members_can_create_repositories bool `json:"members_can_create_repositories,omitempty"`
	Login string `json:"login"`
	Owned_private_repos int `json:"owned_private_repos,omitempty"`
	Dependabot_alerts_enabled_for_new_repositories bool `json:"dependabot_alerts_enabled_for_new_repositories,omitempty"` // **Endpoint closing down notice.** Please use [code security configurations](https://docs.github.com/rest/code-security/configurations) instead. Whether Dependabot alerts are automatically enabled for new repositories and repositories transferred to this organization. This field is only visible to organization owners or members of a team with the security manager role.
	Members_can_create_pages bool `json:"members_can_create_pages,omitempty"`
	Secret_scanning_push_protection_custom_link string `json:"secret_scanning_push_protection_custom_link,omitempty"` // An optional URL string to display to contributors who are blocked from pushing a secret.
	Collaborators int `json:"collaborators,omitempty"` // The number of collaborators on private repositories. This field may be null if the number of private repositories is over 50,000.
	Members_can_create_private_pages bool `json:"members_can_create_private_pages,omitempty"`
	Description string `json:"description"`
	Updated_at string `json:"updated_at"`
	Members_can_create_teams bool `json:"members_can_create_teams,omitempty"`
	Advanced_security_enabled_for_new_repositories bool `json:"advanced_security_enabled_for_new_repositories,omitempty"` // **Endpoint closing down notice.** Please use [code security configurations](https://docs.github.com/rest/code-security/configurations) instead. Whether GitHub Advanced Security is enabled for new repositories and repositories transferred to this organization. This field is only visible to organization owners or members of a team with the security manager role.
	Location string `json:"location,omitempty"`
	Is_verified bool `json:"is_verified,omitempty"`
	Created_at string `json:"created_at"`
	Secret_scanning_push_protection_enabled_for_new_repositories bool `json:"secret_scanning_push_protection_enabled_for_new_repositories,omitempty"` // **Endpoint closing down notice.** Please use [code security configurations](https://docs.github.com/rest/code-security/configurations) instead. Whether secret scanning push protection is automatically enabled for new repositories and repositories transferred to this organization. This field is only visible to organization owners or members of a team with the security manager role.
	Members_can_change_repo_visibility bool `json:"members_can_change_repo_visibility,omitempty"`
	Members_can_view_dependency_insights bool `json:"members_can_view_dependency_insights,omitempty"`
	Secret_scanning_enabled_for_new_repositories bool `json:"secret_scanning_enabled_for_new_repositories,omitempty"` // **Endpoint closing down notice.** Please use [code security configurations](https://docs.github.com/rest/code-security/configurations) instead. Whether secret scanning is automatically enabled for new repositories and repositories transferred to this organization. This field is only visible to organization owners or members of a team with the security manager role.
	Default_repository_branch string `json:"default_repository_branch,omitempty"` // The default branch for repositories created in this organization.
	Public_members_url string `json:"public_members_url"`
	Default_repository_permission string `json:"default_repository_permission,omitempty"`
	Members_can_create_internal_repositories bool `json:"members_can_create_internal_repositories,omitempty"`
	Members_can_create_public_repositories bool `json:"members_can_create_public_repositories,omitempty"`
	Public_repos int `json:"public_repos"`
	Members_url string `json:"members_url"`
	Archived_at string `json:"archived_at"`
	Plan map[string]interface{} `json:"plan,omitempty"`
	Members_can_create_private_repositories bool `json:"members_can_create_private_repositories,omitempty"`
	Readers_can_create_discussions bool `json:"readers_can_create_discussions,omitempty"`
	Email string `json:"email,omitempty"`
	Two_factor_requirement_enabled bool `json:"two_factor_requirement_enabled,omitempty"`
	Members_can_invite_outside_collaborators bool `json:"members_can_invite_outside_collaborators,omitempty"`
	Node_id string `json:"node_id"`
	Repos_url string `json:"repos_url"`
	Billing_email string `json:"billing_email,omitempty"`
	Issues_url string `json:"issues_url"`
	TypeField string `json:"type"`
	Members_can_create_public_pages bool `json:"members_can_create_public_pages,omitempty"`
	Url string `json:"url"`
	Events_url string `json:"events_url"`
	Members_can_delete_repositories bool `json:"members_can_delete_repositories,omitempty"`
	Name string `json:"name,omitempty"`
	Id int `json:"id"`
	Private_gists int `json:"private_gists,omitempty"`
	Members_can_delete_issues bool `json:"members_can_delete_issues,omitempty"`
	Html_url string `json:"html_url"`
	Hooks_url string `json:"hooks_url"`
	Followers int `json:"followers"`
	Avatar_url string `json:"avatar_url"`
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"`
	Secret_scanning_push_protection_custom_link_enabled bool `json:"secret_scanning_push_protection_custom_link_enabled,omitempty"` // Whether a custom link is shown to contributors who are blocked from pushing a secret by push protection.
	Members_can_fork_private_repositories bool `json:"members_can_fork_private_repositories,omitempty"`
	Dependabot_security_updates_enabled_for_new_repositories bool `json:"dependabot_security_updates_enabled_for_new_repositories,omitempty"` // **Endpoint closing down notice.** Please use [code security configurations](https://docs.github.com/rest/code-security/configurations) instead. Whether Dependabot security updates are automatically enabled for new repositories and repositories transferred to this organization. This field is only visible to organization owners or members of a team with the security manager role.
	Display_commenter_full_name_setting_enabled bool `json:"display_commenter_full_name_setting_enabled,omitempty"`
	Members_allowed_repository_creation_type string `json:"members_allowed_repository_creation_type,omitempty"`
	Twitter_username string `json:"twitter_username,omitempty"`
	Following int `json:"following"`
	Public_gists int `json:"public_gists"`
	Blog string `json:"blog,omitempty"`
	Disk_usage int `json:"disk_usage,omitempty"`
	Dependency_graph_enabled_for_new_repositories bool `json:"dependency_graph_enabled_for_new_repositories,omitempty"` // **Endpoint closing down notice.** Please use [code security configurations](https://docs.github.com/rest/code-security/configurations) instead. Whether dependency graph is automatically enabled for new repositories and repositories transferred to this organization. This field is only visible to organization owners or members of a team with the security manager role.
	Has_organization_projects bool `json:"has_organization_projects"`
	Company string `json:"company,omitempty"`
}

// Webhooksteam represents the Webhooksteam schema from the OpenAPI specification
type Webhooksteam struct {
	Node_id string `json:"node_id,omitempty"`
	Privacy string `json:"privacy,omitempty"`
	Html_url string `json:"html_url,omitempty"`
	Members_url string `json:"members_url,omitempty"`
	Notification_setting string `json:"notification_setting,omitempty"`
	Permission string `json:"permission,omitempty"` // Permission that the team will have for its repositories
	Url string `json:"url,omitempty"` // URL for the team
	Deleted bool `json:"deleted,omitempty"`
	Description string `json:"description,omitempty"` // Description of the team
	Parent map[string]interface{} `json:"parent,omitempty"`
	Repositories_url string `json:"repositories_url,omitempty"`
	Slug string `json:"slug,omitempty"`
	Id int `json:"id"` // Unique identifier of the team
	Name string `json:"name"` // Name of the team
}

// Webhooksusermannequin represents the Webhooksusermannequin schema from the OpenAPI specification
type Webhooksusermannequin struct {
	TypeField string `json:"type,omitempty"`
	Gists_url string `json:"gists_url,omitempty"`
	Deleted bool `json:"deleted,omitempty"`
	Events_url string `json:"events_url,omitempty"`
	Email string `json:"email,omitempty"`
	Subscriptions_url string `json:"subscriptions_url,omitempty"`
	Avatar_url string `json:"avatar_url,omitempty"`
	Html_url string `json:"html_url,omitempty"`
	Node_id string `json:"node_id,omitempty"`
	Repos_url string `json:"repos_url,omitempty"`
	Followers_url string `json:"followers_url,omitempty"`
	Starred_url string `json:"starred_url,omitempty"`
	Received_events_url string `json:"received_events_url,omitempty"`
	Id int `json:"id"`
	Url string `json:"url,omitempty"`
	Gravatar_id string `json:"gravatar_id,omitempty"`
	Following_url string `json:"following_url,omitempty"`
	Login string `json:"login"`
	Organizations_url string `json:"organizations_url,omitempty"`
	User_view_type string `json:"user_view_type,omitempty"`
	Name string `json:"name,omitempty"`
	Site_admin bool `json:"site_admin,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Merged bool `json:"merged"`
	Message string `json:"message"`
	Sha string `json:"sha"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Created_at string `json:"created_at"`
	Creator GeneratedType `json:"creator,omitempty"` // A GitHub user.
	Id float64 `json:"id"`
	Archived_at string `json:"archived_at"`
	Content_node_id string `json:"content_node_id"`
	Node_id string `json:"node_id,omitempty"`
	Updated_at string `json:"updated_at"`
	Content_type string `json:"content_type"` // The type of content tracked in a project item
	Project_node_id string `json:"project_node_id,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Active_caches_count int `json:"active_caches_count"` // The number of active caches in the repository.
	Active_caches_size_in_bytes int `json:"active_caches_size_in_bytes"` // The sum of the size in bytes of all the active cache items in the repository.
	Full_name string `json:"full_name"` // The repository owner and name for the cache usage being shown.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Member Webhooksuser `json:"member"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes,omitempty"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Confirm_delete_url string `json:"confirm_delete_url"` // Next deletable analysis in chain, with last analysis deletion confirmation
	Next_analysis_url string `json:"next_analysis_url"` // Next deletable analysis in chain, without last analysis deletion confirmation
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository map[string]interface{} `json:"repository,omitempty"` // A git repository
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Team Webhooksteam1 `json:"team"` // Groups of organization members that gives permissions on specified repositories.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	TypeField string `json:"type"`
}

// Webhooksuser represents the Webhooksuser schema from the OpenAPI specification
type Webhooksuser struct {
	Gists_url string `json:"gists_url,omitempty"`
	Repos_url string `json:"repos_url,omitempty"`
	Html_url string `json:"html_url,omitempty"`
	Email string `json:"email,omitempty"`
	Deleted bool `json:"deleted,omitempty"`
	Gravatar_id string `json:"gravatar_id,omitempty"`
	Name string `json:"name,omitempty"`
	User_view_type string `json:"user_view_type,omitempty"`
	Followers_url string `json:"followers_url,omitempty"`
	Organizations_url string `json:"organizations_url,omitempty"`
	Avatar_url string `json:"avatar_url,omitempty"`
	Id int64 `json:"id"`
	Subscriptions_url string `json:"subscriptions_url,omitempty"`
	Site_admin bool `json:"site_admin,omitempty"`
	TypeField string `json:"type,omitempty"`
	Url string `json:"url,omitempty"`
	Starred_url string `json:"starred_url,omitempty"`
	Login string `json:"login"`
	Node_id string `json:"node_id,omitempty"`
	Received_events_url string `json:"received_events_url,omitempty"`
	Events_url string `json:"events_url,omitempty"`
	Following_url string `json:"following_url,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue Webhooksissue2 `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Requester Webhooksuser `json:"requester,omitempty"`
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation Installation `json:"installation"` // Installation
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repositories []map[string]interface{} `json:"repositories,omitempty"` // An array of repository objects that the installation can access.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Registry_package map[string]interface{} `json:"registry_package"`
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// Page represents the Page schema from the OpenAPI specification
type Page struct {
	Https_certificate GeneratedType `json:"https_certificate,omitempty"`
	Protected_domain_state string `json:"protected_domain_state,omitempty"` // The state if the domain is verified
	Url string `json:"url"` // The API address for accessing this Page resource.
	Pending_domain_unverified_at string `json:"pending_domain_unverified_at,omitempty"` // The timestamp when a pending domain becomes unverified.
	Status string `json:"status"` // The status of the most recent build of the Page.
	Https_enforced bool `json:"https_enforced,omitempty"` // Whether https is enabled on the domain
	Public bool `json:"public"` // Whether the GitHub Pages site is publicly visible. If set to `true`, the site is accessible to anyone on the internet. If set to `false`, the site will only be accessible to users who have at least `read` access to the repository that published the site.
	Custom_404 bool `json:"custom_404"` // Whether the Page has a custom 404 page.
	Source GeneratedType `json:"source,omitempty"`
	Build_type string `json:"build_type,omitempty"` // The process in which the Page will be built.
	Cname string `json:"cname"` // The Pages site's custom domain
	Html_url string `json:"html_url,omitempty"` // The web address the Page can be accessed from.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Links map[string]interface{} `json:"_links"`
	Git_url string `json:"git_url"`
	Sha string `json:"sha"`
	TypeField string `json:"type"`
	Url string `json:"url"`
	Html_url string `json:"html_url"`
	Size int `json:"size"`
	Content string `json:"content,omitempty"`
	Download_url string `json:"download_url"`
	Encoding string `json:"encoding,omitempty"`
	Entries []map[string]interface{} `json:"entries,omitempty"`
	Name string `json:"name"`
	Path string `json:"path"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Starred_at string `json:"starred_at"` // The time the star was created. This is a timestamp in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`. Will be `null` for the `deleted` action.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Actor_type string `json:"actor_type"` // The type of actor that can bypass a ruleset.
	Bypass_mode string `json:"bypass_mode,omitempty"` // When the specified actor can bypass the ruleset. `pull_request` means that an actor can only bypass rules on pull requests. `pull_request` is not applicable for the `DeployKey` actor type. Also, `pull_request` is only applicable to branch rulesets.
	Actor_id int `json:"actor_id,omitempty"` // The ID of the actor that can bypass a ruleset. If `actor_type` is `OrganizationAdmin`, this should be `1`. If `actor_type` is `DeployKey`, this should be null. `OrganizationAdmin` is not applicable for personal repositories.
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
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Alert GeneratedType `json:"alert"` // A Dependabot alert.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	TypeField string `json:"type"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id float64 `json:"id"`
	Inherited_allows_public_repositories bool `json:"inherited_allows_public_repositories,omitempty"`
	Name string `json:"name"`
	Restricted_to_workflows bool `json:"restricted_to_workflows,omitempty"` // If `true`, the runner group will be restricted to running only the workflows specified in the `selected_workflows` array.
	Runners_url string `json:"runners_url"`
	Selected_workflows []string `json:"selected_workflows,omitempty"` // List of workflows the runner group should be allowed to run. This setting will be ignored unless `restricted_to_workflows` is set to `true`.
	Visibility string `json:"visibility"`
	DefaultField bool `json:"default"`
	Inherited bool `json:"inherited"`
	Network_configuration_id string `json:"network_configuration_id,omitempty"` // The identifier of a hosted compute network configuration.
	Selected_repositories_url string `json:"selected_repositories_url,omitempty"` // Link to the selected repositories resource for this runner group. Not present unless visibility was set to `selected`
	Workflow_restrictions_read_only bool `json:"workflow_restrictions_read_only,omitempty"` // If `true`, the `restricted_to_workflows` and `selected_workflows` fields cannot be modified.
	Allows_public_repositories bool `json:"allows_public_repositories"`
	Hosted_runners_url string `json:"hosted_runners_url,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Date string `json:"date,omitempty"`
	Email string `json:"email,omitempty"`
	Name string `json:"name,omitempty"`
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
	Files map[string]interface{} `json:"files"`
	Health_percentage int `json:"health_percentage"`
	Updated_at string `json:"updated_at"`
	Content_reports_enabled bool `json:"content_reports_enabled,omitempty"`
	Description string `json:"description"`
	Documentation string `json:"documentation"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Compute_service string `json:"compute_service,omitempty"` // The hosted compute service the network configuration supports.
	Created_on string `json:"created_on"` // The time at which the network configuration was created, in ISO 8601 format.
	Id string `json:"id"` // The unique identifier of the network configuration.
	Name string `json:"name"` // The name of the network configuration.
	Network_settings_ids []string `json:"network_settings_ids,omitempty"` // The unique identifier of each network settings in the configuration.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Name string `json:"name"`
	Url string `json:"url"`
	Body string `json:"body,omitempty"`
	Html_url string `json:"html_url"`
	Key string `json:"key"`
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
	Url string `json:"url"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Assignee GeneratedType `json:"assignee"` // A GitHub user.
	Commit_id string `json:"commit_id"`
	Created_at string `json:"created_at"`
	Commit_url string `json:"commit_url"`
	Event string `json:"event"`
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	Performed_via_github_app Integration `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Assigner GeneratedType `json:"assigner"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Description string `json:"description"`
	Updated_at string `json:"updated_at"`
	Closed_at string `json:"closed_at"`
	Id int `json:"id"`
	Closed_issues int `json:"closed_issues"`
	Created_at string `json:"created_at"`
	Html_url string `json:"html_url"`
	Node_id string `json:"node_id"`
	Open_issues int `json:"open_issues"`
	Title string `json:"title"` // The title of the milestone.
	Due_on string `json:"due_on"`
	Url string `json:"url"`
	Labels_url string `json:"labels_url"`
	Number int `json:"number"` // The number of the milestone.
	State string `json:"state"` // The state of the milestone.
	Creator GeneratedType `json:"creator"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Web_url string `json:"web_url"` // URL to access this codespace on the web.
	Environment_id string `json:"environment_id"` // UUID identifying this codespace's environment.
	Idle_timeout_minutes int `json:"idle_timeout_minutes"` // The number of minutes of inactivity after which this codespace will be automatically stopped.
	Pulls_url string `json:"pulls_url"` // API URL for the Pull Request associated with this codespace, if any.
	Location string `json:"location"` // The initally assigned location of a new codespace.
	Publish_url string `json:"publish_url,omitempty"` // API URL to publish this codespace to a new repository.
	Display_name string `json:"display_name,omitempty"` // Display name for this codespace.
	Start_url string `json:"start_url"` // API URL to start this codespace.
	Idle_timeout_notice string `json:"idle_timeout_notice,omitempty"` // Text to show user when codespace idle timeout minutes has been overriden by an organization policy
	Pending_operation_disabled_reason string `json:"pending_operation_disabled_reason,omitempty"` // Text to show user when codespace is disabled by a pending operation
	Recent_folders []string `json:"recent_folders"`
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Stop_url string `json:"stop_url"` // API URL to stop this codespace.
	Created_at string `json:"created_at"`
	Machines_url string `json:"machines_url"` // API URL to access available alternate machine types for this codespace.
	Name string `json:"name"` // Automatically generated name of this codespace.
	Retention_period_minutes int `json:"retention_period_minutes,omitempty"` // Duration in minutes after codespace has gone idle in which it will be deleted. Must be integer minutes between 0 and 43200 (30 days).
	Machine GeneratedType `json:"machine"` // A description of the machine powering a codespace.
	Billable_owner GeneratedType `json:"billable_owner"` // A GitHub user.
	Last_used_at string `json:"last_used_at"` // Last known time this codespace was started.
	Url string `json:"url"` // API URL for this codespace.
	Devcontainer_path string `json:"devcontainer_path,omitempty"` // Path to devcontainer.json from repo root used to create Codespace.
	Repository GeneratedType `json:"repository"` // Full Repository
	Runtime_constraints map[string]interface{} `json:"runtime_constraints,omitempty"`
	Pending_operation bool `json:"pending_operation,omitempty"` // Whether or not a codespace has a pending async operation. This would mean that the codespace is temporarily unavailable. The only thing that you can do with a codespace in this state is delete it.
	Id int64 `json:"id"`
	Updated_at string `json:"updated_at"`
	Git_status map[string]interface{} `json:"git_status"` // Details about the codespace's git repository.
	Retention_expires_at string `json:"retention_expires_at,omitempty"` // When a codespace will be auto-deleted based on the "retention_period_minutes" and "last_used_at"
	Prebuild bool `json:"prebuild"` // Whether the codespace was created from a prebuild.
	State string `json:"state"` // State of this codespace.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Name string `json:"name"` // The name of the CodeQL database.
	Size int `json:"size"` // The size of the CodeQL database file in bytes.
	Url string `json:"url"` // The URL at which to download the CodeQL database. The `Accept` header must be set to the value of the `content_type` property.
	Updated_at string `json:"updated_at"` // The date and time at which the CodeQL database was last updated, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Uploader GeneratedType `json:"uploader"` // A GitHub user.
	Content_type string `json:"content_type"` // The MIME type of the CodeQL database file.
	Created_at string `json:"created_at"` // The date and time at which the CodeQL database was created, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Id int `json:"id"` // The ID of the CodeQL database.
	Commit_oid string `json:"commit_oid,omitempty"` // The commit SHA of the repository at the time the CodeQL database was created.
	Language string `json:"language"` // The language of the CodeQL database.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Reason string `json:"reason,omitempty"` // Explains why the merge group is being destroyed. The group could have been merged, removed from the queue (dequeued), or invalidated by an earlier queue entry being dequeued (invalidated).
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Merge_group GeneratedType `json:"merge_group"` // A group of pull requests that the merge queue has grouped together to be merged.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Models []map[string]interface{} `json:"models,omitempty"` // List of model metrics for a custom models and the default model.
	Total_engaged_users int `json:"total_engaged_users,omitempty"` // Total number of users who prompted Copilot Chat on github.com at least once.
}

// Feed represents the Feed schema from the OpenAPI specification
type Feed struct {
	Repository_discussions_url string `json:"repository_discussions_url,omitempty"` // A feed of discussions for a given repository.
	Current_user_organization_url string `json:"current_user_organization_url,omitempty"`
	Timeline_url string `json:"timeline_url"`
	User_url string `json:"user_url"`
	Current_user_public_url string `json:"current_user_public_url,omitempty"`
	Current_user_url string `json:"current_user_url,omitempty"`
	Repository_discussions_category_url string `json:"repository_discussions_category_url,omitempty"` // A feed of discussions for a given repository and category.
	Security_advisories_url string `json:"security_advisories_url,omitempty"`
	Links map[string]interface{} `json:"_links"`
	Current_user_actor_url string `json:"current_user_actor_url,omitempty"`
	Current_user_organization_urls []string `json:"current_user_organization_urls,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Workflow_job interface{} `json:"workflow_job"`
	Action string `json:"action"`
	Deployment Deployment `json:"deployment,omitempty"` // A request for a specific ref(branch,sha,tag) to be deployed
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Total_engaged_users int `json:"total_engaged_users,omitempty"` // Total number of users who prompted Copilot Chat in the IDE.
	Editors []map[string]interface{} `json:"editors,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Rate GeneratedType `json:"rate"`
	Resources map[string]interface{} `json:"resources"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Payload string `json:"payload"` // A URL-encoded string of the check_run.created JSON payload. The decoded payload is a JSON object.
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
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Workflow Webhooksworkflow `json:"workflow"`
	Workflow_run map[string]interface{} `json:"workflow_run"`
	Action string `json:"action"`
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
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Release Webhooksrelease `json:"release"` // The [release](https://docs.github.com/rest/releases/releases/#get-a-release) object.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url"`
	Cards_url string `json:"cards_url"`
	Created_at string `json:"created_at"`
	Id int `json:"id"` // The unique identifier of the project column
	Name string `json:"name"` // Name of the project column
	Node_id string `json:"node_id"`
	Project_url string `json:"project_url"`
	Updated_at string `json:"updated_at"`
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
	Created_at string `json:"created_at"`
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	Url string `json:"url"`
	Label map[string]interface{} `json:"label"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Commit_url string `json:"commit_url"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Commit_id string `json:"commit_id"`
	Event string `json:"event"`
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

// Workflow represents the Workflow schema from the OpenAPI specification
type Workflow struct {
	Html_url string `json:"html_url"`
	Name string `json:"name"`
	State string `json:"state"`
	Url string `json:"url"`
	Badge_url string `json:"badge_url"`
	Deleted_at string `json:"deleted_at,omitempty"`
	Id int `json:"id"`
	Path string `json:"path"`
	Updated_at string `json:"updated_at"`
	Node_id string `json:"node_id"`
	Created_at string `json:"created_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Name string `json:"name"` // Name of the label.
	TypeField string `json:"type,omitempty"` // The type of label. Read-only labels are applied automatically when the runner is configured.
	Id int `json:"id,omitempty"` // Unique identifier of the label.
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
	Href string `json:"href"`
	TypeField string `json:"type"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Issue interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) the comment belongs to.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Comment map[string]interface{} `json:"comment"` // The [comment](https://docs.github.com/rest/issues/comments#get-an-issue-comment) itself.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Guid string `json:"guid,omitempty"` // The GUID of the tool used to generate the code scanning analysis, if provided in the uploaded SARIF data.
	Name string `json:"name,omitempty"` // The name of the tool used to generate the code scanning analysis.
	Version string `json:"version,omitempty"` // The version of the tool used to generate the code scanning analysis.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Name string `json:"name"` // Name of the team
	Html_url string `json:"html_url"`
	Organization GeneratedType `json:"organization"` // Team Organization
	Parent GeneratedType `json:"parent,omitempty"` // Groups of organization members that gives permissions on specified repositories.
	Permission string `json:"permission"` // Permission that the team will have for its repositories
	Slug string `json:"slug"`
	Description string `json:"description"`
	Node_id string `json:"node_id"`
	Id int `json:"id"` // Unique identifier of the team
	Notification_setting string `json:"notification_setting,omitempty"` // The notification setting the team has set
	Privacy string `json:"privacy,omitempty"` // The level of privacy this team should have
	Repositories_url string `json:"repositories_url"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"` // URL for the team
	Repos_count int `json:"repos_count"`
	Ldap_dn string `json:"ldap_dn,omitempty"` // Distinguished Name (DN) that team maps to within LDAP environment
	Members_url string `json:"members_url"`
	Created_at string `json:"created_at"`
	Members_count int `json:"members_count"`
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

// Authorization represents the Authorization schema from the OpenAPI specification
type Authorization struct {
	Token_last_eight string `json:"token_last_eight"`
	App map[string]interface{} `json:"app"`
	Created_at string `json:"created_at"`
	Id int64 `json:"id"`
	Token string `json:"token"`
	Updated_at string `json:"updated_at"`
	Installation GeneratedType `json:"installation,omitempty"`
	Url string `json:"url"`
	Hashed_token string `json:"hashed_token"`
	Note string `json:"note"`
	Expires_at string `json:"expires_at"`
	Fingerprint string `json:"fingerprint"`
	Note_url string `json:"note_url"`
	Scopes []string `json:"scopes"` // A list of scopes that this authorization is in.
	User GeneratedType `json:"user,omitempty"` // A GitHub user.
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
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Marketplace_purchase Webhooksmarketplacepurchase `json:"marketplace_purchase"`
	Action string `json:"action"`
	Effective_date string `json:"effective_date"`
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Previous_marketplace_purchase map[string]interface{} `json:"previous_marketplace_purchase,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Status string `json:"status,omitempty"` // The current status of the deployment.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Publicly_leaked bool `json:"publicly_leaked,omitempty"` // Whether the detected secret was publicly leaked.
	Resolution string `json:"resolution,omitempty"` // The reason for resolving the alert.
	Secret_type_display_name string `json:"secret_type_display_name,omitempty"` // User-friendly name for the detected secret, matching the `secret_type`. For a list of built-in patterns, see "[Supported secret scanning patterns](https://docs.github.com/code-security/secret-scanning/introduction/supported-secret-scanning-patterns#supported-secrets)."
	Validity string `json:"validity,omitempty"` // The token status as of the latest validity check.
	Updated_at string `json:"updated_at,omitempty"` // The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Secret_type string `json:"secret_type,omitempty"` // The type of secret that secret scanning detected.
	Push_protection_bypass_request_comment string `json:"push_protection_bypass_request_comment,omitempty"` // An optional comment when requesting a push protection bypass.
	Push_protection_bypass_request_html_url string `json:"push_protection_bypass_request_html_url,omitempty"` // The URL to a push protection bypass request.
	Html_url string `json:"html_url,omitempty"` // The GitHub URL of the alert resource.
	Locations_url string `json:"locations_url,omitempty"` // The REST API URL of the code locations for this alert.
	Resolution_comment string `json:"resolution_comment,omitempty"` // An optional comment to resolve an alert.
	Multi_repo bool `json:"multi_repo,omitempty"` // Whether the detected secret was found in multiple repositories in the same organization or business.
	Url string `json:"url,omitempty"` // The REST API URL of the alert resource.
	Created_at string `json:"created_at,omitempty"` // The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Push_protection_bypassed_by GeneratedType `json:"push_protection_bypassed_by,omitempty"` // A GitHub user.
	Resolved_by GeneratedType `json:"resolved_by,omitempty"` // A GitHub user.
	Number int `json:"number,omitempty"` // The security alert number.
	Push_protection_bypass_request_reviewer GeneratedType `json:"push_protection_bypass_request_reviewer,omitempty"` // A GitHub user.
	Push_protection_bypassed_at string `json:"push_protection_bypassed_at,omitempty"` // The time that push protection was bypassed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Resolved_at string `json:"resolved_at,omitempty"` // The time that the alert was resolved in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Push_protection_bypass_request_reviewer_comment string `json:"push_protection_bypass_request_reviewer_comment,omitempty"` // An optional comment when reviewing a push protection bypass.
	Push_protection_bypassed bool `json:"push_protection_bypassed,omitempty"` // Whether push protection was bypassed for the detected secret.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Comment string `json:"comment"` // The comment submitted with the deployment review
	Environments []map[string]interface{} `json:"environments"` // The list of environments that were approved or rejected
	State string `json:"state"` // Whether deployment to the environment(s) was approved or rejected or pending (with comments)
	User GeneratedType `json:"user"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Include_claim_keys []string `json:"include_claim_keys,omitempty"` // Array of unique strings. Each claim key can only contain alphanumeric characters and underscores.
	Use_default bool `json:"use_default"` // Whether to use the default template or not. If `true`, the `include_claim_keys` field is ignored.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Projects_v2_item GeneratedType `json:"projects_v2_item"` // An item belonging to a project
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Score float64 `json:"score"`
	Sha string `json:"sha"`
	Commit map[string]interface{} `json:"commit"`
	Committer GeneratedType `json:"committer"` // Metaproperties for Git author/committer information.
	Html_url string `json:"html_url"`
	Parents []map[string]interface{} `json:"parents"`
	Node_id string `json:"node_id"`
	Repository GeneratedType `json:"repository"` // Minimal Repository
	Url string `json:"url"`
	Author GeneratedType `json:"author"` // A GitHub user.
	Text_matches []map[string]interface{} `json:"text_matches,omitempty"`
	Comments_url string `json:"comments_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url"`
	Base map[string]interface{} `json:"base"`
	Head map[string]interface{} `json:"head"`
	Id int64 `json:"id"`
	Number int `json:"number"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Login string `json:"login"`
	Members_url string `json:"members_url"`
	Public_members_url string `json:"public_members_url"`
	Description string `json:"description"`
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	Url string `json:"url"`
	Hooks_url string `json:"hooks_url"`
	Repos_url string `json:"repos_url"`
	Avatar_url string `json:"avatar_url"`
	Events_url string `json:"events_url"`
	Issues_url string `json:"issues_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enabled bool `json:"enabled,omitempty"`
	Lock_branch map[string]interface{} `json:"lock_branch,omitempty"` // Whether to set the branch as read-only. If this is true, users will not be able to push to the branch.
	Allow_deletions map[string]interface{} `json:"allow_deletions,omitempty"`
	Enforce_admins GeneratedType `json:"enforce_admins,omitempty"` // Protected Branch Admin Enforced
	Block_creations map[string]interface{} `json:"block_creations,omitempty"`
	Required_linear_history map[string]interface{} `json:"required_linear_history,omitempty"`
	Required_conversation_resolution map[string]interface{} `json:"required_conversation_resolution,omitempty"`
	Protection_url string `json:"protection_url,omitempty"`
	Required_signatures map[string]interface{} `json:"required_signatures,omitempty"`
	Allow_fork_syncing map[string]interface{} `json:"allow_fork_syncing,omitempty"` // Whether users can pull changes from upstream when the branch is locked. Set to `true` to allow fork syncing. Set to `false` to prevent fork syncing.
	Name string `json:"name,omitempty"`
	Required_status_checks GeneratedType `json:"required_status_checks,omitempty"` // Protected Branch Required Status Check
	Allow_force_pushes map[string]interface{} `json:"allow_force_pushes,omitempty"`
	Url string `json:"url,omitempty"`
	Required_pull_request_reviews GeneratedType `json:"required_pull_request_reviews,omitempty"` // Protected Branch Pull Request Review
	Restrictions GeneratedType `json:"restrictions,omitempty"` // Branch Restriction Policy
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository_ruleset GeneratedType `json:"repository_ruleset"` // A set of rules to apply when specified conditions are met.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Check_run GeneratedType `json:"check_run"` // A check performed on the code of a given code change
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	State string `json:"state"`
	Submitted_at string `json:"submitted_at,omitempty"`
	User GeneratedType `json:"user"` // A GitHub user.
	Pull_request_url string `json:"pull_request_url"`
	Body string `json:"body"` // The text of the review.
	Body_text string `json:"body_text,omitempty"`
	Id int64 `json:"id"` // Unique identifier of the review
	Links map[string]interface{} `json:"_links"`
	Body_html string `json:"body_html,omitempty"`
	Commit_id string `json:"commit_id"` // A commit SHA for the review. If the commit object was garbage collected or forcibly deleted, then it no longer exists in Git and this value will be `null`.
	Html_url string `json:"html_url"`
	Node_id string `json:"node_id"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Pull_request_review_comment_url string `json:"pull_request_review_comment_url"` // The API URL to get the pull request review comment where the secret was detected.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository,omitempty"` // A GitHub repository.
	Status string `json:"status,omitempty"` // The attachment status of the code security configuration on the repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Usageitems []map[string]interface{} `json:"usageItems,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Payload string `json:"payload"` // A URL-encoded string of the check_run.requested_action JSON payload. The decoded payload is a JSON object.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Thread map[string]interface{} `json:"thread"`
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request map[string]interface{} `json:"pull_request"`
}

// Stargazer represents the Stargazer schema from the OpenAPI specification
type Stargazer struct {
	Starred_at string `json:"starred_at"`
	User GeneratedType `json:"user"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pusher_type string `json:"pusher_type"` // The pusher type for the event. Can be either `user` or a deploy key.
	Ref string `json:"ref"` // The [`git ref`](https://docs.github.com/rest/git/refs#get-a-reference) resource.
	Ref_type string `json:"ref_type"` // The type of Git ref object deleted in the repository.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Vulnerable_version_range string `json:"vulnerable_version_range"` // The range of the package versions affected by the vulnerability.
	PackageField map[string]interface{} `json:"package"` // The name of the package affected by the vulnerability.
	Patched_versions string `json:"patched_versions"` // The package version(s) that resolve the vulnerability.
	Vulnerable_functions []string `json:"vulnerable_functions"` // The functions in the package that are affected.
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
	Action string `json:"action"` // The action that was performed.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// Commit represents the Commit schema from the OpenAPI specification
type Commit struct {
	Url string `json:"url"`
	Files []GeneratedType `json:"files,omitempty"`
	Node_id string `json:"node_id"`
	Author interface{} `json:"author"`
	Html_url string `json:"html_url"`
	Sha string `json:"sha"`
	Comments_url string `json:"comments_url"`
	Commit map[string]interface{} `json:"commit"`
	Committer interface{} `json:"committer"`
	Parents []map[string]interface{} `json:"parents"`
	Stats map[string]interface{} `json:"stats,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Comment Webhooksreviewcomment `json:"comment"` // The [comment](https://docs.github.com/rest/pulls/comments#get-a-review-comment-for-a-pull-request) itself.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
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
	Files []GeneratedType `json:"files,omitempty"`
	Behind_by int `json:"behind_by"`
	Diff_url string `json:"diff_url"`
	Html_url string `json:"html_url"`
	Merge_base_commit Commit `json:"merge_base_commit"` // Commit
	Permalink_url string `json:"permalink_url"`
	Base_commit Commit `json:"base_commit"` // Commit
	Patch_url string `json:"patch_url"`
	Status string `json:"status"`
	Total_commits int `json:"total_commits"`
	Url string `json:"url"`
	Ahead_by int `json:"ahead_by"`
	Commits []Commit `json:"commits"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Old_answer Webhooksanswer `json:"old_answer"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Discussion_title_url string `json:"discussion_title_url"` // The URL to the discussion where the secret was detected.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Comments_enabled bool `json:"comments_enabled,omitempty"`
	Truncated bool `json:"truncated,omitempty"`
	Forks_url string `json:"forks_url"`
	History []interface{} `json:"history,omitempty"`
	Owner GeneratedType `json:"owner,omitempty"` // A GitHub user.
	Forks []interface{} `json:"forks,omitempty"`
	Id string `json:"id"`
	Files map[string]interface{} `json:"files"`
	Created_at string `json:"created_at"`
	Url string `json:"url"`
	Comments_url string `json:"comments_url"`
	Git_pull_url string `json:"git_pull_url"`
	Updated_at string `json:"updated_at"`
	Comments int `json:"comments"`
	Git_push_url string `json:"git_push_url"`
	Html_url string `json:"html_url"`
	Commits_url string `json:"commits_url"`
	User GeneratedType `json:"user"` // A GitHub user.
	Description string `json:"description"`
	Node_id string `json:"node_id"`
	Public bool `json:"public"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository_id map[string]interface{} `json:"repository_id"`
}

// Webhookschanges represents the Webhookschanges schema from the OpenAPI specification
type Webhookschanges struct {
	Body map[string]interface{} `json:"body,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
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

// Tag represents the Tag schema from the OpenAPI specification
type Tag struct {
	Zipball_url string `json:"zipball_url"`
	Commit map[string]interface{} `json:"commit"`
	Name string `json:"name"`
	Node_id string `json:"node_id"`
	Tarball_url string `json:"tarball_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Issue_title_url string `json:"issue_title_url"` // The API URL to get the issue where the secret was detected.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Key string `json:"key"`
	Name string `json:"name"`
	Url string `json:"url"`
	Html_url string `json:"html_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Ref string `json:"ref"` // The full git ref that was pushed. Example: `refs/heads/main` or `refs/tags/v3.14.1`.
	Repository map[string]interface{} `json:"repository"` // A git repository
	Before string `json:"before"` // The SHA of the most recent commit on `ref` before the push.
	Created bool `json:"created"` // Whether this push created the `ref`.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Commits []map[string]interface{} `json:"commits"` // An array of commit objects describing the pushed commits. (Pushed commits are all commits that are included in the `compare` between the `before` commit and the `after` commit.) The array includes a maximum of 2048 commits. If necessary, you can use the [Commits API](https://docs.github.com/rest/commits) to fetch additional commits.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Forced bool `json:"forced"` // Whether this push was a force push of the `ref`.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	After string `json:"after"` // The SHA of the most recent commit on `ref` after the push.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Compare string `json:"compare"` // URL that shows the changes in this `ref` update, from the `before` commit to the `after` commit. For a newly created `ref` that is directly based on the default branch, this is the comparison between the head of the default branch and the `after` commit. Otherwise, this shows all commits until the `after` commit.
	Deleted bool `json:"deleted"` // Whether this push deleted the `ref`.
	Pusher map[string]interface{} `json:"pusher"` // Metaproperties for Git author/committer information.
	Base_ref string `json:"base_ref"`
	Head_commit map[string]interface{} `json:"head_commit"`
}

// Installation represents the Installation schema from the OpenAPI specification
type Installation struct {
	Events []string `json:"events"`
	Account interface{} `json:"account"`
	App_slug string `json:"app_slug"`
	Contact_email string `json:"contact_email,omitempty"`
	Suspended_by GeneratedType `json:"suspended_by"` // A GitHub user.
	App_id int `json:"app_id"`
	Has_multiple_single_files bool `json:"has_multiple_single_files,omitempty"`
	Access_tokens_url string `json:"access_tokens_url"`
	Repositories_url string `json:"repositories_url"`
	Suspended_at string `json:"suspended_at"`
	Permissions GeneratedType `json:"permissions"` // The permissions granted to the user access token.
	Updated_at string `json:"updated_at"`
	Created_at string `json:"created_at"`
	Html_url string `json:"html_url"`
	Target_id int `json:"target_id"` // The ID of the user or organization this token is being scoped to.
	Single_file_name string `json:"single_file_name"`
	Id int `json:"id"` // The ID of the installation.
	Repository_selection string `json:"repository_selection"` // Describe whether all repositories have been selected or there's a selection involved
	Target_type string `json:"target_type"`
	Single_file_paths []string `json:"single_file_paths,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Tool GeneratedType `json:"tool"`
	Rules_count int `json:"rules_count"` // The total number of rules used in the analysis.
	Commit_sha string `json:"commit_sha"` // The SHA of the commit to which the analysis you are uploading relates.
	Environment string `json:"environment"` // Identifies the variable values associated with the environment in which this analysis was performed.
	Analysis_key string `json:"analysis_key"` // Identifies the configuration under which the analysis was executed. For example, in GitHub Actions this includes the workflow filename and job name.
	Deletable bool `json:"deletable"`
	Id int `json:"id"` // Unique identifier for this analysis.
	Warning string `json:"warning"` // Warning generated when processing the analysis
	Category string `json:"category,omitempty"` // Identifies the configuration under which the analysis was executed. Used to distinguish between multiple analyses for the same tool and commit, but performed on different languages or different parts of the code.
	Created_at string `json:"created_at"` // The time that the analysis was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Sarif_id string `json:"sarif_id"` // An identifier for the upload.
	ErrorField string `json:"error"`
	Ref string `json:"ref"` // The Git reference, formatted as `refs/pull/<number>/merge`, `refs/pull/<number>/head`, `refs/heads/<branch name>` or simply `<branch name>`.
	Url string `json:"url"` // The REST API URL of the analysis resource.
	Results_count int `json:"results_count"` // The total number of results in the analysis.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url"`
	Html_url string `json:"html_url,omitempty"`
	Key string `json:"key"`
	Name string `json:"name"`
	Node_id string `json:"node_id"`
	Spdx_id string `json:"spdx_id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Color string `json:"color"`
	DefaultField bool `json:"default"`
	Node_id string `json:"node_id"`
	Score float64 `json:"score"`
	Description string `json:"description"`
	Name string `json:"name"`
	Id int `json:"id"`
	Text_matches []map[string]interface{} `json:"text_matches,omitempty"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id int `json:"id"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Url string `json:"url"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Commit_url string `json:"commit_url"`
	Dismissed_review map[string]interface{} `json:"dismissed_review"`
	Created_at string `json:"created_at"`
	Event string `json:"event"`
	Node_id string `json:"node_id"`
	Commit_id string `json:"commit_id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Definition GeneratedType `json:"definition"` // Custom property defined on an organization
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Body string `json:"body"` // The generated body describing the contents of the release supporting markdown formatting
	Name string `json:"name"` // The generated name of the release
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	TypeField string `json:"type"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Key string `json:"key"`
	Created_at string `json:"created_at,omitempty"`
	Id int `json:"id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Heart int `json:"heart"`
	Url string `json:"url"`
	Confused int `json:"confused"`
	Hooray int `json:"hooray"`
	Eyes int `json:"eyes"`
	Laugh int `json:"laugh"`
	Rocket int `json:"rocket"`
	Field1 int `json:"+1"`
	Field1 int `json:"-1"`
	Total_count int `json:"total_count"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Secret string `json:"secret,omitempty"` // If provided, the `secret` will be used as the `key` to generate the HMAC hex digest value for [delivery signature headers](https://docs.github.com/webhooks/event-payloads/#delivery-headers).
	Url string `json:"url,omitempty"` // The URL to which the payloads will be delivered.
	Content_type string `json:"content_type,omitempty"` // The media type used to serialize the payloads. Supported values include `json` and `form`. The default is `form`.
	Insecure_ssl string `json:"insecure_ssl,omitempty"`
}

// Project represents the Project schema from the OpenAPI specification
type Project struct {
	Url string `json:"url"`
	Created_at string `json:"created_at"`
	Body string `json:"body"` // Body of the project
	Id int `json:"id"`
	Private bool `json:"private,omitempty"` // Whether or not this project can be seen by everyone. Only present if owner is an organization.
	Node_id string `json:"node_id"`
	Html_url string `json:"html_url"`
	Organization_permission string `json:"organization_permission,omitempty"` // The baseline permission that all organization members have on this project. Only present if owner is an organization.
	State string `json:"state"` // State of the project; either 'open' or 'closed'
	Owner_url string `json:"owner_url"`
	Creator GeneratedType `json:"creator"` // A GitHub user.
	Number int `json:"number"`
	Columns_url string `json:"columns_url"`
	Updated_at string `json:"updated_at"`
	Name string `json:"name"` // Name of the project
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Comment Webhookscomment `json:"comment"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
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
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Forkee interface{} `json:"forkee"` // The created [`repository`](https://docs.github.com/rest/repos/repos#get-a-repository) resource.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Ref_name map[string]interface{} `json:"ref_name,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Users []map[string]interface{} `json:"users"`
	Users_url string `json:"users_url"`
	Apps []map[string]interface{} `json:"apps"`
	Apps_url string `json:"apps_url"`
	Teams []map[string]interface{} `json:"teams"`
	Teams_url string `json:"teams_url"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Raw_details string `json:"raw_details"`
	Start_line int `json:"start_line"`
	Title string `json:"title"`
	End_column int `json:"end_column"`
	End_line int `json:"end_line"`
	Blob_href string `json:"blob_href"`
	Path string `json:"path"`
	Annotation_level string `json:"annotation_level"`
	Message string `json:"message"`
	Start_column int `json:"start_column"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Parent_issue_id float64 `json:"parent_issue_id"` // The ID of the parent issue.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sub_issue_id float64 `json:"sub_issue_id"` // The ID of the sub-issue.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Sub_issue Issue `json:"sub_issue"` // Issues are a great way to keep track of tasks, enhancements, and bugs for your projects.
	Action string `json:"action"`
	Parent_issue Issue `json:"parent_issue"` // Issues are a great way to keep track of tasks, enhancements, and bugs for your projects.
	Sub_issue_repo Repository `json:"sub_issue_repo"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Commit_id string `json:"commit_id"`
	Created_at string `json:"created_at"`
	Commit_url string `json:"commit_url"`
	Event string `json:"event"`
	Node_id string `json:"node_id"`
	Project_card map[string]interface{} `json:"project_card,omitempty"`
	Url string `json:"url"`
	Id int `json:"id"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Actor GeneratedType `json:"actor"` // A GitHub user.
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
	Alert GeneratedType `json:"alert"` // A Dependabot alert.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Percentage float64 `json:"percentage,omitempty"`
	Percentile float64 `json:"percentile,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Path string `json:"path"`
	Size int `json:"size"`
	TypeField string `json:"type"`
	Download_url string `json:"download_url"`
	Name string `json:"name"`
	Encoding string `json:"encoding"`
	Html_url string `json:"html_url"`
	Links map[string]interface{} `json:"_links"`
	Git_url string `json:"git_url"`
	License GeneratedType `json:"license"` // License Simple
	Sha string `json:"sha"`
	Url string `json:"url"`
	Content string `json:"content"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Pull_request Webhookspullrequest5 `json:"pull_request"`
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Milestone Milestone `json:"milestone,omitempty"` // A collection of related issues and pull requests.
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Definition map[string]interface{} `json:"definition"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Threat_model string `json:"threat_model,omitempty"` // Threat model to be used for code scanning analysis. Use `remote` to analyze only network sources and `remote_and_local` to include local sources like filesystem access, command-line arguments, database reads, environment variable and standard input.
	Languages []string `json:"languages,omitempty"` // CodeQL languages to be analyzed.
	Query_suite string `json:"query_suite,omitempty"` // CodeQL query suite to be used.
	Runner_label string `json:"runner_label,omitempty"` // Runner label to be used if the runner type is labeled.
	Runner_type string `json:"runner_type,omitempty"` // Runner type to be used.
	State string `json:"state,omitempty"` // The desired state of code scanning default setup.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	After string `json:"after,omitempty"`
	Repository GeneratedType `json:"repository,omitempty"` // Minimal Repository
	Updated_at string `json:"updated_at,omitempty"`
	Before string `json:"before,omitempty"`
	Head_branch string `json:"head_branch,omitempty"`
	Id int `json:"id,omitempty"`
	Status string `json:"status,omitempty"`
	App Integration `json:"app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Conclusion string `json:"conclusion,omitempty"`
	Created_at string `json:"created_at,omitempty"`
	Pull_requests []GeneratedType `json:"pull_requests,omitempty"`
	Url string `json:"url,omitempty"`
	Head_sha string `json:"head_sha,omitempty"` // The SHA of the head commit that is being checked.
	Node_id string `json:"node_id,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id int `json:"id"` // The ID of the installation.
	Node_id string `json:"node_id"` // The global node ID of the installation.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Milestone Webhooksmilestone3 `json:"milestone"` // A collection of related issues and pull requests.
}

// Migration represents the Migration schema from the OpenAPI specification
type Migration struct {
	Created_at string `json:"created_at"`
	Repositories []Repository `json:"repositories"` // The repositories included in the migration. Only returned for export migrations.
	Exclude_git_data bool `json:"exclude_git_data"`
	Id int64 `json:"id"`
	Node_id string `json:"node_id"`
	Guid string `json:"guid"`
	Lock_repositories bool `json:"lock_repositories"`
	Archive_url string `json:"archive_url,omitempty"`
	Org_metadata_only bool `json:"org_metadata_only"`
	Exclude_owner_projects bool `json:"exclude_owner_projects"`
	Exclude_attachments bool `json:"exclude_attachments"`
	State string `json:"state"`
	Exclude_metadata bool `json:"exclude_metadata"`
	Url string `json:"url"`
	Exclude_releases bool `json:"exclude_releases"`
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Updated_at string `json:"updated_at"`
	Exclude []string `json:"exclude,omitempty"` // Exclude related items from being returned in the response in order to improve performance of the request. The array can include any of: `"repositories"`.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Allowed_actions string `json:"allowed_actions,omitempty"` // The permissions policy that controls the actions and reusable workflows that are allowed to run.
	Enabled bool `json:"enabled"` // Whether GitHub Actions is enabled on the repository.
	Selected_actions_url string `json:"selected_actions_url,omitempty"` // The API URL to use to get or set the actions and reusable workflows that are allowed to run, when `allowed_actions` is set to `selected`.
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
	Created_at string `json:"created_at"`
	Id int `json:"id"`
	Url string `json:"url"`
	Verified bool `json:"verified"`
	Read_only bool `json:"read_only"`
	Title string `json:"title"`
	Key string `json:"key"`
	Added_by string `json:"added_by,omitempty"`
	Enabled bool `json:"enabled,omitempty"`
	Last_used string `json:"last_used,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Changes map[string]interface{} `json:"changes"`
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Pull_request map[string]interface{} `json:"pull_request"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Review Webhooksreview `json:"review"` // The review that was affected.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Path string `json:"path"`
	Url string `json:"url"`
	Download_url string `json:"download_url"`
	Encoding string `json:"encoding"`
	Git_url string `json:"git_url"`
	Sha string `json:"sha"`
	Content string `json:"content"`
	Html_url string `json:"html_url"`
	Size int `json:"size"`
	Submodule_git_url string `json:"submodule_git_url,omitempty"`
	Target string `json:"target,omitempty"`
	TypeField string `json:"type"`
	Links map[string]interface{} `json:"_links"`
	Name string `json:"name"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Membership Webhooksmembership `json:"membership"` // The membership between the user and the organization. Not present when the action is `member_invited`.
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
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
	Id int `json:"id"` // ID of the reviewer which must review changes to matching files.
	TypeField string `json:"type"` // The type of the reviewer
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Short_description string `json:"short_description"`
	Title string `json:"title"`
	Created_at string `json:"created_at"`
	Deleted_by GeneratedType `json:"deleted_by"` // A GitHub user.
	Number int `json:"number"`
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Updated_at string `json:"updated_at"`
	Closed_at string `json:"closed_at"`
	Deleted_at string `json:"deleted_at"`
	Node_id string `json:"node_id"`
	Public bool `json:"public"`
	Creator GeneratedType `json:"creator"` // A GitHub user.
	Description string `json:"description"`
	Id float64 `json:"id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Vulnerabilities []map[string]interface{} `json:"vulnerabilities,omitempty"` // An array of products affected by the vulnerability detailed in a repository security advisory.
	Cvss_vector_string string `json:"cvss_vector_string,omitempty"` // The CVSS vector that calculates the severity of the advisory. You must choose between setting this field or `severity`.
	Cwe_ids []string `json:"cwe_ids,omitempty"` // A list of Common Weakness Enumeration (CWE) IDs.
	Description string `json:"description"` // A detailed description of what the advisory impacts.
	Severity string `json:"severity,omitempty"` // The severity of the advisory. You must choose between setting this field or `cvss_vector_string`.
	Start_private_fork bool `json:"start_private_fork,omitempty"` // Whether to create a temporary private fork of the repository to collaborate on a fix.
	Summary string `json:"summary"` // A short summary of the advisory.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Private bool `json:"private"` // Whether the repository is private.
	Default_branch string `json:"default_branch"` // The default branch for the repository.
	Full_name string `json:"full_name"` // The full, globally unique name of the repository.
	Html_url string `json:"html_url"` // The URL to view the repository on GitHub.com.
	Id int `json:"id"` // A unique identifier of the repository.
	Node_id string `json:"node_id"` // The GraphQL identifier of the repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Pull_request_body_url string `json:"pull_request_body_url"` // The API URL to get the pull request where the secret was detected.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Teams []Team `json:"teams"`
	Users []GeneratedType `json:"users"`
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
	Repository_property map[string]interface{} `json:"repository_property"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Alert map[string]interface{} `json:"alert"` // The code scanning alert involved in the event.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Action string `json:"action"`
	Commit_oid string `json:"commit_oid"` // The commit SHA of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Ref string `json:"ref"` // The Git reference of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Created_at string `json:"created_at"`
	Team_url string `json:"team_url"`
	Html_url string `json:"html_url"`
	Number int `json:"number"` // The unique sequence number of a team discussion.
	Author GeneratedType `json:"author"` // A GitHub user.
	Body_html string `json:"body_html"`
	Url string `json:"url"`
	Reactions GeneratedType `json:"reactions,omitempty"`
	Comments_url string `json:"comments_url"`
	Title string `json:"title"` // The title of the discussion.
	Body_version string `json:"body_version"` // The current version of the body content. If provided, this update operation will be rejected if the given version does not match the latest version on the server.
	Node_id string `json:"node_id"`
	Updated_at string `json:"updated_at"`
	Pinned bool `json:"pinned"` // Whether or not this discussion should be pinned for easy retrieval.
	Comments_count int `json:"comments_count"`
	Last_edited_at string `json:"last_edited_at"`
	Body string `json:"body"` // The main text of the discussion.
	Private bool `json:"private"` // Whether or not this discussion should be restricted to team members and organization owners.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	TypeField string `json:"type"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Version_id int `json:"version_id"` // The ID of the previous version of the ruleset
	Actor map[string]interface{} `json:"actor"` // The actor who updated the ruleset
	Updated_at string `json:"updated_at"`
	State map[string]interface{} `json:"state"` // The state of the ruleset version
}

// Status represents the Status schema from the OpenAPI specification
type Status struct {
	Id int `json:"id"`
	State string `json:"state"`
	Updated_at string `json:"updated_at"`
	Context string `json:"context"`
	Node_id string `json:"node_id"`
	Avatar_url string `json:"avatar_url"`
	Target_url string `json:"target_url"`
	Creator GeneratedType `json:"creator"` // A GitHub user.
	Url string `json:"url"`
	Created_at string `json:"created_at"`
	Description string `json:"description"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Rule Webhooksrule `json:"rule"` // The branch protection rule. Includes a `name` and all the [branch protection settings](https://docs.github.com/github/administering-a-repository/defining-the-mergeability-of-pull-requests/about-protected-branches#about-branch-protection-settings) applied to branches that match the name. Binary settings are boolean. Multi-level configurations are one of `off`, `non_admins`, or `everyone`. Actor and build lists are arrays of strings.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// Package represents the Package schema from the OpenAPI specification
type Package struct {
	Name string `json:"name"` // The name of the package.
	Package_type string `json:"package_type"`
	Created_at string `json:"created_at"`
	Html_url string `json:"html_url"`
	Id int `json:"id"` // Unique identifier of the package.
	Owner GeneratedType `json:"owner,omitempty"` // A GitHub user.
	Repository GeneratedType `json:"repository,omitempty"` // Minimal Repository
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	Version_count int `json:"version_count"` // The number of versions of the package.
	Visibility string `json:"visibility"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Project_column Webhooksprojectcolumn `json:"project_column"`
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue Webhooksissue2 `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Deployment Deployment `json:"deployment,omitempty"` // A request for a specific ref(branch,sha,tag) to be deployed
	Pull_requests []GeneratedType `json:"pull_requests,omitempty"`
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Event string `json:"event,omitempty"` // The event that triggered the deployment protection rule.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Deployment_callback_url string `json:"deployment_callback_url,omitempty"` // The URL to review the deployment protection rule.
	Environment string `json:"environment,omitempty"` // The name of the environment that has the deployment protection rule.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Action string `json:"action,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Review Webhooksreview `json:"review"` // The review that was affected.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Failure_reason string `json:"failure_reason,omitempty"` // The reason for a failure of the variant analysis. This is only available if the variant analysis has failed.
	Updated_at string `json:"updated_at,omitempty"` // The date and time at which the variant analysis was last updated, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Actions_workflow_run_id int `json:"actions_workflow_run_id,omitempty"` // The GitHub Actions workflow run used to execute this variant analysis. This is only available if the workflow run has started.
	Id int `json:"id"` // The ID of the variant analysis.
	Query_language string `json:"query_language"` // The language targeted by the CodeQL query
	Query_pack_url string `json:"query_pack_url"` // The download url for the query pack.
	Scanned_repositories []map[string]interface{} `json:"scanned_repositories,omitempty"`
	Skipped_repositories map[string]interface{} `json:"skipped_repositories,omitempty"` // Information about repositories that were skipped from processing. This information is only available to the user that initiated the variant analysis.
	Controller_repo GeneratedType `json:"controller_repo"` // A GitHub repository.
	Created_at string `json:"created_at,omitempty"` // The date and time at which the variant analysis was created, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Status string `json:"status"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Completed_at string `json:"completed_at,omitempty"` // The date and time at which the variant analysis was completed, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ. Will be null if the variant analysis has not yet completed or this information is not available.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Color string `json:"color,omitempty"` // The color of the issue type.
	Created_at string `json:"created_at,omitempty"` // The time the issue type created.
	Description string `json:"description"` // The description of the issue type.
	Id int `json:"id"` // The unique identifier of the issue type.
	Is_enabled bool `json:"is_enabled,omitempty"` // The enabled state of the issue type.
	Name string `json:"name"` // The name of the issue type.
	Node_id string `json:"node_id"` // The node identifier of the issue type.
	Updated_at string `json:"updated_at,omitempty"` // The time the issue type last updated.
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
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// Discussion represents the Discussion schema from the OpenAPI specification
type Discussion struct {
	Body string `json:"body"`
	Active_lock_reason string `json:"active_lock_reason"`
	Comments int `json:"comments"`
	Updated_at string `json:"updated_at"`
	Timeline_url string `json:"timeline_url,omitempty"`
	Reactions map[string]interface{} `json:"reactions,omitempty"`
	Repository_url string `json:"repository_url"`
	State_reason string `json:"state_reason"` // The reason for the current state
	Number int `json:"number"`
	Answer_chosen_by map[string]interface{} `json:"answer_chosen_by"`
	Answer_chosen_at string `json:"answer_chosen_at"`
	Answer_html_url string `json:"answer_html_url"`
	Id int `json:"id"`
	Category map[string]interface{} `json:"category"`
	Html_url string `json:"html_url"`
	User map[string]interface{} `json:"user"`
	Locked bool `json:"locked"`
	Node_id string `json:"node_id"`
	Title string `json:"title"`
	Created_at string `json:"created_at"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Labels []Label `json:"labels,omitempty"`
	State string `json:"state"` // The current state of the discussion. `converting` means that the discussion is being converted from an issue. `transferring` means that the discussion is being transferred from another repository.
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
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// Webhookspullrequest5 represents the Webhookspullrequest5 schema from the OpenAPI specification
type Webhookspullrequest5 struct {
	Requested_reviewers []interface{} `json:"requested_reviewers"`
	Review_comments int `json:"review_comments,omitempty"`
	Number int `json:"number"` // Number uniquely identifying the pull request within its repository.
	Rebaseable bool `json:"rebaseable,omitempty"`
	Url string `json:"url"`
	Base map[string]interface{} `json:"base"`
	Merged_at string `json:"merged_at"`
	Id int `json:"id"`
	Diff_url string `json:"diff_url"`
	Mergeable bool `json:"mergeable,omitempty"`
	Merged bool `json:"merged,omitempty"`
	Title string `json:"title"` // The title of the pull request.
	Closed_at string `json:"closed_at"`
	Html_url string `json:"html_url"`
	Maintainer_can_modify bool `json:"maintainer_can_modify,omitempty"` // Indicates whether maintainers can modify the pull request.
	Patch_url string `json:"patch_url"`
	Body string `json:"body"`
	Comments_url string `json:"comments_url"`
	Head map[string]interface{} `json:"head"`
	Issue_url string `json:"issue_url"`
	Locked bool `json:"locked"`
	Statuses_url string `json:"statuses_url"`
	User map[string]interface{} `json:"user"`
	Milestone map[string]interface{} `json:"milestone"` // A collection of related issues and pull requests.
	Assignee map[string]interface{} `json:"assignee"`
	Labels []map[string]interface{} `json:"labels"`
	Node_id string `json:"node_id"`
	Requested_teams []map[string]interface{} `json:"requested_teams"`
	Deletions int `json:"deletions,omitempty"`
	Comments int `json:"comments,omitempty"`
	Created_at string `json:"created_at"`
	Merge_commit_sha string `json:"merge_commit_sha"`
	Updated_at string `json:"updated_at"`
	Mergeable_state string `json:"mergeable_state,omitempty"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Draft bool `json:"draft"` // Indicates whether or not the pull request is a draft.
	State string `json:"state"` // State of this Pull Request. Either `open` or `closed`.
	Links map[string]interface{} `json:"_links"`
	Commits_url string `json:"commits_url"`
	Commits int `json:"commits,omitempty"`
	Review_comments_url string `json:"review_comments_url"`
	Auto_merge map[string]interface{} `json:"auto_merge"` // The status of auto merging a pull request.
	Changed_files int `json:"changed_files,omitempty"`
	Merged_by map[string]interface{} `json:"merged_by,omitempty"`
	Review_comment_url string `json:"review_comment_url"`
	Active_lock_reason string `json:"active_lock_reason"`
	Additions int `json:"additions,omitempty"`
	Assignees []map[string]interface{} `json:"assignees"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Selected_repositories_url string `json:"selected_repositories_url,omitempty"`
	Updated_at string `json:"updated_at"` // The date and time at which the variable was last updated, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Value string `json:"value"` // The value of the variable.
	Visibility string `json:"visibility"` // Visibility of a variable
	Created_at string `json:"created_at"` // The date and time at which the variable was created, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Name string `json:"name"` // The name of the variable.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Severity string `json:"severity,omitempty"` // The severity of the advisory. You must choose between setting this field or `cvss_vector_string`.
	Vulnerabilities []map[string]interface{} `json:"vulnerabilities,omitempty"` // A product affected by the vulnerability detailed in a repository security advisory.
	Credits []map[string]interface{} `json:"credits,omitempty"` // A list of users receiving credit for their participation in the security advisory.
	Cvss_vector_string string `json:"cvss_vector_string,omitempty"` // The CVSS vector that calculates the severity of the advisory. You must choose between setting this field or `severity`.
	State string `json:"state,omitempty"` // The state of the advisory.
	Summary string `json:"summary,omitempty"` // A short summary of the advisory.
	Description string `json:"description,omitempty"` // A detailed description of what the advisory impacts.
	Collaborating_teams []string `json:"collaborating_teams,omitempty"` // A list of team slugs which have been granted write access to the advisory.
	Collaborating_users []string `json:"collaborating_users,omitempty"` // A list of usernames who have been granted write access to the advisory.
	Cve_id string `json:"cve_id,omitempty"` // The Common Vulnerabilities and Exposures (CVE) ID.
	Cwe_ids []string `json:"cwe_ids,omitempty"` // A list of Common Weakness Enumeration (CWE) IDs.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	TypeField string `json:"type"`
	Parameters map[string]interface{} `json:"parameters,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Updated_at string `json:"updated_at"`
	Full_name string `json:"full_name"` // The full, globally unique, name of the repository.
	Id int `json:"id"` // A unique identifier of the repository.
	Name string `json:"name"` // The name of the repository.
	Private bool `json:"private"` // Whether the repository is private.
	Stargazers_count int `json:"stargazers_count"`
}

// Dependency represents the Dependency schema from the OpenAPI specification
type Dependency struct {
	Scope string `json:"scope,omitempty"` // A notation of whether the dependency is required for the primary build artifact (runtime) or is only used for development. Future versions of this specification may allow for more granular scopes.
	Dependencies []string `json:"dependencies,omitempty"` // Array of package-url (PURLs) of direct child dependencies.
	Metadata Metadata `json:"metadata,omitempty"` // User-defined metadata to store domain-specific information limited to 8 keys with scalar values.
	Package_url string `json:"package_url,omitempty"` // Package-url (PURL) of dependency. See https://github.com/package-url/purl-spec for more details.
	Relationship string `json:"relationship,omitempty"` // A notation of whether a dependency is requested directly by this manifest or is a dependency of another dependency.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Pull_requests string `json:"pull_requests,omitempty"` // The level of permission to grant the access token for pull requests and related comments, assignees, labels, milestones, and merges.
	Organization_announcement_banners string `json:"organization_announcement_banners,omitempty"` // The level of permission to grant the access token to view and manage announcement banners for an organization.
	Security_events string `json:"security_events,omitempty"` // The level of permission to grant the access token to view and manage security events like code scanning alerts.
	Organization_administration string `json:"organization_administration,omitempty"` // The level of permission to grant the access token to manage access to an organization.
	Codespaces string `json:"codespaces,omitempty"` // The level of permission to grant the access token to create, edit, delete, and list Codespaces.
	Organization_personal_access_token_requests string `json:"organization_personal_access_token_requests,omitempty"` // The level of permission to grant the access token for viewing and managing fine-grained personal access tokens that have been approved by an organization.
	Email_addresses string `json:"email_addresses,omitempty"` // The level of permission to grant the access token to manage the email addresses belonging to a user.
	Team_discussions string `json:"team_discussions,omitempty"` // The level of permission to grant the access token to manage team discussions and related comments.
	Git_ssh_keys string `json:"git_ssh_keys,omitempty"` // The level of permission to grant the access token to manage git SSH keys.
	Actions string `json:"actions,omitempty"` // The level of permission to grant the access token for GitHub Actions workflows, workflow runs, and artifacts.
	Issues string `json:"issues,omitempty"` // The level of permission to grant the access token for issues and related comments, assignees, labels, and milestones.
	Organization_custom_properties string `json:"organization_custom_properties,omitempty"` // The level of permission to grant the access token for custom property management.
	Repository_projects string `json:"repository_projects,omitempty"` // The level of permission to grant the access token to manage repository projects, columns, and cards.
	Repository_custom_properties string `json:"repository_custom_properties,omitempty"` // The level of permission to grant the access token to view and edit custom properties for a repository, when allowed by the property.
	Dependabot_secrets string `json:"dependabot_secrets,omitempty"` // The level of permission to grant the access token to manage Dependabot secrets.
	Pages string `json:"pages,omitempty"` // The level of permission to grant the access token to retrieve Pages statuses, configuration, and builds, as well as create new builds.
	Single_file string `json:"single_file,omitempty"` // The level of permission to grant the access token to manage just a single file.
	Gpg_keys string `json:"gpg_keys,omitempty"` // The level of permission to grant the access token to view and manage GPG keys belonging to a user.
	Environments string `json:"environments,omitempty"` // The level of permission to grant the access token for managing repository environments.
	Repository_hooks string `json:"repository_hooks,omitempty"` // The level of permission to grant the access token to manage the post-receive hooks for a repository.
	Organization_plan string `json:"organization_plan,omitempty"` // The level of permission to grant the access token for viewing an organization's plan.
	Secret_scanning_alerts string `json:"secret_scanning_alerts,omitempty"` // The level of permission to grant the access token to view and manage secret scanning alerts.
	Organization_custom_org_roles string `json:"organization_custom_org_roles,omitempty"` // The level of permission to grant the access token for custom organization roles management.
	Organization_packages string `json:"organization_packages,omitempty"` // The level of permission to grant the access token for organization packages published to GitHub Packages.
	Organization_self_hosted_runners string `json:"organization_self_hosted_runners,omitempty"` // The level of permission to grant the access token to view and manage GitHub Actions self-hosted runners available to an organization.
	Administration string `json:"administration,omitempty"` // The level of permission to grant the access token for repository creation, deletion, settings, teams, and collaborators creation.
	Contents string `json:"contents,omitempty"` // The level of permission to grant the access token for repository contents, commits, branches, downloads, releases, and merges.
	Followers string `json:"followers,omitempty"` // The level of permission to grant the access token to manage the followers belonging to a user.
	Profile string `json:"profile,omitempty"` // The level of permission to grant the access token to manage the profile settings belonging to a user.
	Secrets string `json:"secrets,omitempty"` // The level of permission to grant the access token to manage repository secrets.
	Deployments string `json:"deployments,omitempty"` // The level of permission to grant the access token for deployments and deployment statuses.
	Statuses string `json:"statuses,omitempty"` // The level of permission to grant the access token for commit statuses.
	Members string `json:"members,omitempty"` // The level of permission to grant the access token for organization teams and members.
	Organization_events string `json:"organization_events,omitempty"` // The level of permission to grant the access token to view events triggered by an activity in an organization.
	Organization_secrets string `json:"organization_secrets,omitempty"` // The level of permission to grant the access token to manage organization secrets.
	Checks string `json:"checks,omitempty"` // The level of permission to grant the access token for checks on code.
	Packages string `json:"packages,omitempty"` // The level of permission to grant the access token for packages published to GitHub Packages.
	Interaction_limits string `json:"interaction_limits,omitempty"` // The level of permission to grant the access token to view and manage interaction limits on a repository.
	Organization_projects string `json:"organization_projects,omitempty"` // The level of permission to grant the access token to manage organization projects and projects public preview (where available).
	Organization_hooks string `json:"organization_hooks,omitempty"` // The level of permission to grant the access token to manage the post-receive hooks for an organization.
	Organization_custom_roles string `json:"organization_custom_roles,omitempty"` // The level of permission to grant the access token for custom repository roles management.
	Starring string `json:"starring,omitempty"` // The level of permission to grant the access token to list and manage repositories a user is starring.
	Vulnerability_alerts string `json:"vulnerability_alerts,omitempty"` // The level of permission to grant the access token to manage Dependabot alerts.
	Workflows string `json:"workflows,omitempty"` // The level of permission to grant the access token to update GitHub Actions workflow files.
	Metadata string `json:"metadata,omitempty"` // The level of permission to grant the access token to search repositories, list collaborators, and access repository metadata.
	Organization_personal_access_tokens string `json:"organization_personal_access_tokens,omitempty"` // The level of permission to grant the access token for viewing and managing fine-grained personal access token requests to an organization.
	Organization_copilot_seat_management string `json:"organization_copilot_seat_management,omitempty"` // The level of permission to grant the access token for managing access to GitHub Copilot for members of an organization with a Copilot Business subscription. This property is in public preview and is subject to change.
	Organization_user_blocking string `json:"organization_user_blocking,omitempty"` // The level of permission to grant the access token to view and manage users blocked by the organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Dismissal_commit_id string `json:"dismissal_commit_id,omitempty"`
	Dismissal_message string `json:"dismissal_message"`
	Review_id int `json:"review_id"`
	State string `json:"state"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Label Webhookslabel `json:"label"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Release Webhooksrelease `json:"release"` // The [release](https://docs.github.com/rest/releases/releases/#get-a-release) object.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sha string `json:"sha"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Query_suite string `json:"query_suite,omitempty"` // CodeQL query suite to be used.
	Runner_label string `json:"runner_label,omitempty"` // Runner label to be used if the runner type is labeled.
	Runner_type string `json:"runner_type,omitempty"` // Runner type to be used.
	Schedule string `json:"schedule,omitempty"` // The frequency of the periodic analysis.
	State string `json:"state,omitempty"` // Code scanning default setup has been configured or not.
	Threat_model string `json:"threat_model,omitempty"` // Threat model to be used for code scanning analysis. Use `remote` to analyze only network sources and `remote_and_local` to include local sources like filesystem access, command-line arguments, database reads, environment variable and standard input.
	Updated_at string `json:"updated_at,omitempty"` // Timestamp of latest configuration update.
	Languages []string `json:"languages,omitempty"` // Languages to be analyzed.
}

// License represents the License schema from the OpenAPI specification
type License struct {
	Permissions []string `json:"permissions"`
	Description string `json:"description"`
	Body string `json:"body"`
	Implementation string `json:"implementation"`
	Spdx_id string `json:"spdx_id"`
	Url string `json:"url"`
	Conditions []string `json:"conditions"`
	Featured bool `json:"featured"`
	Html_url string `json:"html_url"`
	Key string `json:"key"`
	Limitations []string `json:"limitations"`
	Name string `json:"name"`
	Node_id string `json:"node_id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Languages_url string `json:"languages_url"` // The API URL to get information about the languages of the repository.
	Full_name string `json:"full_name"` // The full, globally unique, name of the repository.
	Stargazers_url string `json:"stargazers_url"` // The API URL to list the stargazers on the repository.
	Forks_url string `json:"forks_url"` // The API URL to list the forks of the repository.
	Releases_url string `json:"releases_url"` // A template for the API URL to get information about releases on the repository.
	Git_tags_url string `json:"git_tags_url"` // A template for the API URL to get information about Git tags of the repository.
	Html_url string `json:"html_url"` // The URL to view the repository on GitHub.com.
	Statuses_url string `json:"statuses_url"` // A template for the API URL to get information about statuses of a commit.
	Hooks_url string `json:"hooks_url"` // The API URL to list the hooks on the repository.
	Teams_url string `json:"teams_url"` // The API URL to list the teams on the repository.
	Git_refs_url string `json:"git_refs_url"` // A template for the API URL to get information about Git refs of the repository.
	Url string `json:"url"` // The URL to get more information about the repository from the GitHub API.
	Blobs_url string `json:"blobs_url"` // A template for the API URL to create or retrieve a raw Git blob in the repository.
	Branches_url string `json:"branches_url"` // A template for the API URL to get information about branches in the repository.
	Assignees_url string `json:"assignees_url"` // A template for the API URL to list the available assignees for issues in the repository.
	Labels_url string `json:"labels_url"` // A template for the API URL to get information about labels of the repository.
	Archive_url string `json:"archive_url"` // A template for the API URL to download the repository as an archive.
	Pulls_url string `json:"pulls_url"` // A template for the API URL to get information about pull requests on the repository.
	Deployments_url string `json:"deployments_url"` // The API URL to list the deployments of the repository.
	Notifications_url string `json:"notifications_url"` // A template for the API URL to get information about notifications on the repository.
	Issues_url string `json:"issues_url"` // A template for the API URL to get information about issues on the repository.
	Events_url string `json:"events_url"` // The API URL to list the events of the repository.
	Node_id string `json:"node_id"` // The GraphQL identifier of the repository.
	Compare_url string `json:"compare_url"` // A template for the API URL to compare two commits or refs.
	Contributors_url string `json:"contributors_url"` // A template for the API URL to list the contributors to the repository.
	Keys_url string `json:"keys_url"` // A template for the API URL to get information about deploy keys on the repository.
	Milestones_url string `json:"milestones_url"` // A template for the API URL to get information about milestones of the repository.
	Comments_url string `json:"comments_url"` // A template for the API URL to get information about comments on the repository.
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Downloads_url string `json:"downloads_url"` // The API URL to list the downloads on the repository.
	Id int64 `json:"id"` // A unique identifier of the repository.
	Collaborators_url string `json:"collaborators_url"` // A template for the API URL to get information about collaborators of the repository.
	Subscribers_url string `json:"subscribers_url"` // The API URL to list the subscribers on the repository.
	Trees_url string `json:"trees_url"` // A template for the API URL to create or retrieve a raw Git tree of the repository.
	Description string `json:"description"` // The repository description.
	Issue_comment_url string `json:"issue_comment_url"` // A template for the API URL to get information about issue comments on the repository.
	Contents_url string `json:"contents_url"` // A template for the API URL to get the contents of the repository.
	Issue_events_url string `json:"issue_events_url"` // A template for the API URL to get information about issue events on the repository.
	Subscription_url string `json:"subscription_url"` // The API URL to subscribe to notifications for this repository.
	Tags_url string `json:"tags_url"` // The API URL to get information about tags on the repository.
	Commits_url string `json:"commits_url"` // A template for the API URL to get information about commits on the repository.
	Git_commits_url string `json:"git_commits_url"` // A template for the API URL to get information about Git commits of the repository.
	Merges_url string `json:"merges_url"` // The API URL to merge branches in the repository.
	Name string `json:"name"` // The name of the repository.
	Fork bool `json:"fork"` // Whether the repository is a fork.
	Private bool `json:"private"` // Whether the repository is private.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Assignees []GeneratedType `json:"assignees,omitempty"`
	Comments_url string `json:"comments_url"`
	Created_at string `json:"created_at"`
	Diff_url string `json:"diff_url"`
	Id int64 `json:"id"`
	Title string `json:"title"`
	Requested_teams []Team `json:"requested_teams,omitempty"`
	Head map[string]interface{} `json:"head"`
	Labels []map[string]interface{} `json:"labels"`
	State string `json:"state"`
	Auto_merge GeneratedType `json:"auto_merge"` // The status of auto merging a pull request.
	Commits_url string `json:"commits_url"`
	Merge_commit_sha string `json:"merge_commit_sha"`
	Active_lock_reason string `json:"active_lock_reason,omitempty"`
	Locked bool `json:"locked"`
	Patch_url string `json:"patch_url"`
	User GeneratedType `json:"user"` // A GitHub user.
	Draft bool `json:"draft,omitempty"` // Indicates whether or not the pull request is a draft.
	Closed_at string `json:"closed_at"`
	Body string `json:"body"`
	Issue_url string `json:"issue_url"`
	Statuses_url string `json:"statuses_url"`
	Review_comment_url string `json:"review_comment_url"`
	Review_comments_url string `json:"review_comments_url"`
	Assignee GeneratedType `json:"assignee"` // A GitHub user.
	Base map[string]interface{} `json:"base"`
	Html_url string `json:"html_url"`
	Updated_at string `json:"updated_at"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Merged_at string `json:"merged_at"`
	Url string `json:"url"`
	Links map[string]interface{} `json:"_links"`
	Requested_reviewers []GeneratedType `json:"requested_reviewers,omitempty"`
	Node_id string `json:"node_id"`
	Number int `json:"number"`
	Milestone GeneratedType `json:"milestone"` // A collection of related issues and pull requests.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Environment string `json:"environment"`
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Reviewers []map[string]interface{} `json:"reviewers"`
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Since string `json:"since"`
	Workflow_run map[string]interface{} `json:"workflow_run"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Workflow_job_run map[string]interface{} `json:"workflow_job_run"`
	Action string `json:"action"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Requestor Webhooksuser `json:"requestor"`
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Effective_date string `json:"effective_date,omitempty"` // The `pending_cancellation` and `pending_tier_change` event types will include the date the cancellation or tier change will take effect.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Sponsorship Webhookssponsorship `json:"sponsorship"`
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
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Ref string `json:"ref"` // The Git reference of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Alert map[string]interface{} `json:"alert"` // The code scanning alert involved in the event.
	Commit_oid string `json:"commit_oid"` // The commit SHA of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Changes map[string]interface{} `json:"changes,omitempty"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Repository_ruleset GeneratedType `json:"repository_ruleset"` // A set of rules to apply when specified conditions are met.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Passing bool `json:"passing"` // Whether a submission passed.
	Repository GeneratedType `json:"repository"` // A GitHub repository view for Classroom
	Students []GeneratedType `json:"students"`
	Submitted bool `json:"submitted"` // Whether an accepted assignment has been submitted.
	Assignment GeneratedType `json:"assignment"` // A GitHub Classroom assignment
	Commit_count int `json:"commit_count"` // Count of student commits.
	Grade string `json:"grade"` // Most recent grade.
	Id int `json:"id"` // Unique identifier of the repository.
}

// Webhooksmilestone represents the Webhooksmilestone schema from the OpenAPI specification
type Webhooksmilestone struct {
	Creator map[string]interface{} `json:"creator"`
	Number int `json:"number"` // The number of the milestone.
	Html_url string `json:"html_url"`
	State string `json:"state"` // The state of the milestone.
	Due_on string `json:"due_on"`
	Labels_url string `json:"labels_url"`
	Description string `json:"description"`
	Closed_issues int `json:"closed_issues"`
	Title string `json:"title"` // The title of the milestone.
	Open_issues int `json:"open_issues"`
	Url string `json:"url"`
	Created_at string `json:"created_at"`
	Id int `json:"id"`
	Updated_at string `json:"updated_at"`
	Closed_at string `json:"closed_at"`
	Node_id string `json:"node_id"`
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
	Documentation_url string `json:"documentation_url"`
	Errors []string `json:"errors,omitempty"`
	Message string `json:"message"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Comment string `json:"comment,omitempty"` // Optional comment to include with the review.
	Environment_name string `json:"environment_name"` // The name of the environment to approve or reject.
	State string `json:"state"` // Whether to approve or reject deployment to the specified environments.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Team_managers []Team `json:"team_managers,omitempty"` // The campaign team managers
	Alert_stats map[string]interface{} `json:"alert_stats,omitempty"`
	Contact_link string `json:"contact_link"` // The contact link of the campaign.
	Ends_at string `json:"ends_at"` // The date and time the campaign has ended, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Managers []GeneratedType `json:"managers"` // The campaign managers
	Number int `json:"number"` // The number of the newly created campaign
	Updated_at string `json:"updated_at"` // The date and time the campaign was last updated, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Created_at string `json:"created_at"` // The date and time the campaign was created, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Description string `json:"description"` // The campaign description
	Name string `json:"name,omitempty"` // The campaign name
	Closed_at string `json:"closed_at,omitempty"` // The date and time the campaign was closed, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ. Will be null if the campaign is still open.
	Published_at string `json:"published_at,omitempty"` // The date and time the campaign was published, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	State string `json:"state"` // Indicates whether a campaign is open or closed
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Node_id string `json:"node_id,omitempty"`
	TypeField string `json:"type,omitempty"` // Whether this rule targets a branch or tag.
	Id int `json:"id,omitempty"` // The unique identifier of the branch or tag policy.
	Name string `json:"name,omitempty"` // The name pattern that branches or tags must match in order to deploy to the environment.
}

// Import represents the Import schema from the OpenAPI specification
type Import struct {
	Vcs string `json:"vcs"`
	Large_files_size int `json:"large_files_size,omitempty"`
	Failed_step string `json:"failed_step,omitempty"`
	Tfvc_project string `json:"tfvc_project,omitempty"`
	Use_lfs bool `json:"use_lfs,omitempty"`
	Has_large_files bool `json:"has_large_files,omitempty"`
	Authors_count int `json:"authors_count,omitempty"`
	Large_files_count int `json:"large_files_count,omitempty"`
	Commit_count int `json:"commit_count,omitempty"`
	Status_text string `json:"status_text,omitempty"`
	Authors_url string `json:"authors_url"`
	Status string `json:"status"`
	Push_percent int `json:"push_percent,omitempty"`
	Import_percent int `json:"import_percent,omitempty"`
	Project_choices []map[string]interface{} `json:"project_choices,omitempty"`
	Svn_root string `json:"svn_root,omitempty"`
	Svc_root string `json:"svc_root,omitempty"`
	Vcs_url string `json:"vcs_url"` // The URL of the originating repository.
	Url string `json:"url"`
	Message string `json:"message,omitempty"`
	Html_url string `json:"html_url"`
	Repository_url string `json:"repository_url"`
	Error_message string `json:"error_message,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Subject_type string `json:"subject_type,omitempty"` // The level at which the comment is targeted, can be a diff line or a file.
	Pull_request_review_id int64 `json:"pull_request_review_id"`
	Original_commit_id string `json:"original_commit_id"`
	Start_side string `json:"start_side,omitempty"` // The side of the first line of the range for a multi-line comment.
	Created_at string `json:"created_at"`
	Original_position int `json:"original_position"`
	User GeneratedType `json:"user"` // A GitHub user.
	Links map[string]interface{} `json:"_links"`
	Body string `json:"body"`
	Path string `json:"path"`
	Line int `json:"line,omitempty"` // The line of the blob to which the comment applies. The last line of the range for a multi-line comment
	Id int64 `json:"id"`
	Body_html string `json:"body_html,omitempty"`
	Html_url string `json:"html_url"`
	Position int `json:"position"`
	Start_line int `json:"start_line,omitempty"` // The first line of the range for a multi-line comment.
	Original_line int `json:"original_line,omitempty"` // The original line of the blob to which the comment applies. The last line of the range for a multi-line comment
	Commit_id string `json:"commit_id"`
	Diff_hunk string `json:"diff_hunk"`
	Side string `json:"side,omitempty"` // The side of the first line of the range for a multi-line comment.
	Updated_at string `json:"updated_at"`
	Reactions GeneratedType `json:"reactions,omitempty"`
	Url string `json:"url"`
	Pull_request_url string `json:"pull_request_url"`
	Body_text string `json:"body_text,omitempty"`
	Node_id string `json:"node_id"`
	In_reply_to_id int `json:"in_reply_to_id,omitempty"`
	Original_start_line int `json:"original_start_line,omitempty"` // The original first line of the range for a multi-line comment.
	Author_association string `json:"author_association"` // How the author is associated with the repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Merged_by GeneratedType `json:"merged_by"` // A GitHub user.
	Links map[string]interface{} `json:"_links"`
	Locked bool `json:"locked"`
	Assignees []GeneratedType `json:"assignees,omitempty"`
	Review_comments_url string `json:"review_comments_url"`
	Merged bool `json:"merged"`
	Draft bool `json:"draft,omitempty"` // Indicates whether or not the pull request is a draft.
	Rebaseable bool `json:"rebaseable,omitempty"`
	Auto_merge GeneratedType `json:"auto_merge"` // The status of auto merging a pull request.
	Mergeable bool `json:"mergeable"`
	Closed_at string `json:"closed_at"`
	State string `json:"state"` // State of this Pull Request. Either `open` or `closed`.
	Statuses_url string `json:"statuses_url"`
	Review_comment_url string `json:"review_comment_url"`
	Base map[string]interface{} `json:"base"`
	Comments int `json:"comments"`
	Title string `json:"title"` // The title of the pull request.
	Labels []map[string]interface{} `json:"labels"`
	Merge_commit_sha string `json:"merge_commit_sha"`
	Id int64 `json:"id"`
	Patch_url string `json:"patch_url"`
	Number int `json:"number"` // Number uniquely identifying the pull request within its repository.
	Assignee GeneratedType `json:"assignee"` // A GitHub user.
	Requested_teams []GeneratedType `json:"requested_teams,omitempty"`
	Updated_at string `json:"updated_at"`
	Deletions int `json:"deletions"`
	Diff_url string `json:"diff_url"`
	Commits_url string `json:"commits_url"`
	Milestone GeneratedType `json:"milestone"` // A collection of related issues and pull requests.
	Node_id string `json:"node_id"`
	Additions int `json:"additions"`
	Created_at string `json:"created_at"`
	Mergeable_state string `json:"mergeable_state"`
	Commits int `json:"commits"`
	Html_url string `json:"html_url"`
	Comments_url string `json:"comments_url"`
	Active_lock_reason string `json:"active_lock_reason,omitempty"`
	Url string `json:"url"`
	Maintainer_can_modify bool `json:"maintainer_can_modify"` // Indicates whether maintainers can modify the pull request.
	Issue_url string `json:"issue_url"`
	Changed_files int `json:"changed_files"`
	Review_comments int `json:"review_comments"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Requested_reviewers []GeneratedType `json:"requested_reviewers,omitempty"`
	User GeneratedType `json:"user"` // A GitHub user.
	Head map[string]interface{} `json:"head"`
	Merged_at string `json:"merged_at"`
	Body string `json:"body"`
	Squash_merge_commit_title string `json:"squash_merge_commit_title,omitempty"` // The default value for a squash merge commit title: - `PR_TITLE` - default to the pull request's title. - `COMMIT_OR_PR_TITLE` - default to the commit's title (if only one commit) or the pull request's title (when more than one commit).
	Use_squash_pr_title_as_default bool `json:"use_squash_pr_title_as_default,omitempty"` // Whether a squash merge commit can use the pull request title as default. **This property is closing down. Please use `squash_merge_commit_title` instead.**
	Allow_auto_merge bool `json:"allow_auto_merge,omitempty"` // Whether to allow auto-merge for pull requests.
	Allow_update_branch bool `json:"allow_update_branch,omitempty"` // Whether to allow updating the pull request's branch.
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"` // Whether to delete head branches when pull requests are merged.
	Merge_commit_message string `json:"merge_commit_message,omitempty"` // The default value for a merge commit message. - `PR_TITLE` - default to the pull request's title. - `PR_BODY` - default to the pull request's body. - `BLANK` - default to a blank commit message.
	Merge_commit_title string `json:"merge_commit_title,omitempty"` // The default value for a merge commit title. - `PR_TITLE` - default to the pull request's title. - `MERGE_MESSAGE` - default to the classic title for a merge message (e.g., "Merge pull request #123 from branch-name").
	Squash_merge_commit_message string `json:"squash_merge_commit_message,omitempty"` // The default value for a squash merge commit message: - `PR_BODY` - default to the pull request's body. - `COMMIT_MESSAGES` - default to the branch's commit messages. - `BLANK` - default to a blank commit message.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Commit map[string]interface{} `json:"commit"`
	Name string `json:"name"`
	Protected bool `json:"protected"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Member Webhooksuser `json:"member"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Status string `json:"status"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Allowed_actions string `json:"allowed_actions,omitempty"` // The permissions policy that controls the actions and reusable workflows that are allowed to run.
	Enabled_repositories string `json:"enabled_repositories"` // The policy that controls the repositories in the organization that are allowed to run GitHub Actions.
	Selected_actions_url string `json:"selected_actions_url,omitempty"` // The API URL to use to get or set the actions and reusable workflows that are allowed to run, when `allowed_actions` is set to `selected`.
	Selected_repositories_url string `json:"selected_repositories_url,omitempty"` // The API URL to use to get or set the selected repositories that are allowed to run GitHub Actions, when `enabled_repositories` is set to `selected`.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Personal_access_token_request GeneratedType `json:"personal_access_token_request"` // Details of a Personal Access Token Request.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Privacy string `json:"privacy,omitempty"` // The level of privacy this team should have
	Url string `json:"url"` // URL for the team
	Description string `json:"description"` // Description of the team
	Html_url string `json:"html_url"`
	Id int `json:"id"` // Unique identifier of the team
	Ldap_dn string `json:"ldap_dn,omitempty"` // Distinguished Name (DN) that team maps to within LDAP environment
	Name string `json:"name"` // Name of the team
	Node_id string `json:"node_id"`
	Notification_setting string `json:"notification_setting,omitempty"` // The notification setting the team has set
	Repositories_url string `json:"repositories_url"`
	Slug string `json:"slug"`
	Members_url string `json:"members_url"`
	Permission string `json:"permission"` // Permission that the team will have for its repositories
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Build map[string]interface{} `json:"build"` // The [List GitHub Pages builds](https://docs.github.com/rest/pages/pages#list-github-pages-builds) itself.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Id int `json:"id"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	TypeField string `json:"type"`
	Parameters map[string]interface{} `json:"parameters,omitempty"`
}

// Snapshot represents the Snapshot schema from the OpenAPI specification
type Snapshot struct {
	Version int `json:"version"` // The version of the repository snapshot submission.
	Detector map[string]interface{} `json:"detector"` // A description of the detector used.
	Job map[string]interface{} `json:"job"`
	Manifests map[string]interface{} `json:"manifests,omitempty"` // A collection of package manifests, which are a collection of related dependencies declared in a file or representing a logical group of dependencies.
	Metadata Metadata `json:"metadata,omitempty"` // User-defined metadata to store domain-specific information limited to 8 keys with scalar values.
	Ref string `json:"ref"` // The repository branch that triggered this snapshot.
	Scanned string `json:"scanned"` // The time at which the snapshot was scanned.
	Sha string `json:"sha"` // The commit SHA associated with this dependency snapshot. Maximum length: 40 characters.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Selected_repositories_url string `json:"selected_repositories_url"` // The API URL at which the list of repositories this secret is visible to can be retrieved
	Updated_at string `json:"updated_at"` // The date and time at which the secret was last updated, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Visibility string `json:"visibility"` // The type of repositories in the organization that the secret is visible to
	Created_at string `json:"created_at"` // The date and time at which the secret was created, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Name string `json:"name"` // The name of the secret
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Merges_url string `json:"merges_url"`
	Fork bool `json:"fork"`
	Forks int `json:"forks"`
	Compare_url string `json:"compare_url"`
	Statuses_url string `json:"statuses_url"`
	Trees_url string `json:"trees_url"`
	Issues_url string `json:"issues_url"`
	Mirror_url string `json:"mirror_url"`
	Keys_url string `json:"keys_url"`
	Labels_url string `json:"labels_url"`
	Allow_squash_merge bool `json:"allow_squash_merge,omitempty"`
	Comments_url string `json:"comments_url"`
	Git_commits_url string `json:"git_commits_url"`
	Git_url string `json:"git_url"`
	Allow_rebase_merge bool `json:"allow_rebase_merge,omitempty"`
	Url string `json:"url"`
	Description string `json:"description"`
	Homepage string `json:"homepage"`
	Languages_url string `json:"languages_url"`
	Commits_url string `json:"commits_url"`
	Deployments_url string `json:"deployments_url"`
	Collaborators_url string `json:"collaborators_url"`
	Git_refs_url string `json:"git_refs_url"`
	Branches_url string `json:"branches_url"`
	Has_projects bool `json:"has_projects"`
	Name string `json:"name"`
	Events_url string `json:"events_url"`
	Allow_forking bool `json:"allow_forking,omitempty"`
	Git_tags_url string `json:"git_tags_url"`
	Forks_url string `json:"forks_url"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Contributors_url string `json:"contributors_url"`
	Is_template bool `json:"is_template,omitempty"`
	Pulls_url string `json:"pulls_url"`
	Downloads_url string `json:"downloads_url"`
	Releases_url string `json:"releases_url"`
	Language string `json:"language"`
	Open_issues int `json:"open_issues"`
	Notifications_url string `json:"notifications_url"`
	Created_at string `json:"created_at"`
	Stargazers_url string `json:"stargazers_url"`
	Hooks_url string `json:"hooks_url"`
	Disabled bool `json:"disabled"` // Returns whether or not this repository disabled.
	Master_branch string `json:"master_branch,omitempty"`
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Subscribers_url string `json:"subscribers_url"`
	Issue_events_url string `json:"issue_events_url"`
	Has_wiki bool `json:"has_wiki"`
	Allow_auto_merge bool `json:"allow_auto_merge,omitempty"`
	Teams_url string `json:"teams_url"`
	Temp_clone_token string `json:"temp_clone_token,omitempty"`
	Ssh_url string `json:"ssh_url"`
	Issue_comment_url string `json:"issue_comment_url"`
	Node_id string `json:"node_id"`
	Watchers int `json:"watchers"`
	Milestones_url string `json:"milestones_url"`
	Has_issues bool `json:"has_issues"`
	Open_issues_count int `json:"open_issues_count"`
	License GeneratedType `json:"license"` // License Simple
	Subscription_url string `json:"subscription_url"`
	Svn_url string `json:"svn_url"`
	Html_url string `json:"html_url"`
	Text_matches []map[string]interface{} `json:"text_matches,omitempty"`
	Archived bool `json:"archived"`
	Topics []string `json:"topics,omitempty"`
	Pushed_at string `json:"pushed_at"`
	Blobs_url string `json:"blobs_url"`
	Stargazers_count int `json:"stargazers_count"`
	Private bool `json:"private"`
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"`
	Default_branch string `json:"default_branch"`
	Size int `json:"size"`
	Watchers_count int `json:"watchers_count"`
	Full_name string `json:"full_name"`
	Score float64 `json:"score"`
	Tags_url string `json:"tags_url"`
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"`
	Clone_url string `json:"clone_url"`
	Has_discussions bool `json:"has_discussions,omitempty"`
	Allow_merge_commit bool `json:"allow_merge_commit,omitempty"`
	Id int `json:"id"`
	Archive_url string `json:"archive_url"`
	Has_pages bool `json:"has_pages"`
	Forks_count int `json:"forks_count"`
	Visibility string `json:"visibility,omitempty"` // The repository visibility: public, private, or internal.
	Contents_url string `json:"contents_url"`
	Updated_at string `json:"updated_at"`
	Assignees_url string `json:"assignees_url"`
	Has_downloads bool `json:"has_downloads"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Repository_advisory GeneratedType `json:"repository_advisory"` // A repository security advisory.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// Contributor represents the Contributor schema from the OpenAPI specification
type Contributor struct {
	Gravatar_id string `json:"gravatar_id,omitempty"`
	Starred_url string `json:"starred_url,omitempty"`
	Url string `json:"url,omitempty"`
	Name string `json:"name,omitempty"`
	Organizations_url string `json:"organizations_url,omitempty"`
	Email string `json:"email,omitempty"`
	Events_url string `json:"events_url,omitempty"`
	Repos_url string `json:"repos_url,omitempty"`
	Subscriptions_url string `json:"subscriptions_url,omitempty"`
	Avatar_url string `json:"avatar_url,omitempty"`
	Id int `json:"id,omitempty"`
	Login string `json:"login,omitempty"`
	Contributions int `json:"contributions"`
	Node_id string `json:"node_id,omitempty"`
	TypeField string `json:"type"`
	Gists_url string `json:"gists_url,omitempty"`
	Site_admin bool `json:"site_admin,omitempty"`
	User_view_type string `json:"user_view_type,omitempty"`
	Followers_url string `json:"followers_url,omitempty"`
	Received_events_url string `json:"received_events_url,omitempty"`
	Html_url string `json:"html_url,omitempty"`
	Following_url string `json:"following_url,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Client_id string `json:"client_id,omitempty"`
	External_url string `json:"external_url"`
	Html_url string `json:"html_url"`
	Permissions map[string]interface{} `json:"permissions"` // The set of permissions for the GitHub app
	Created_at string `json:"created_at"`
	Description string `json:"description"`
	Events []string `json:"events"` // The list of events for the GitHub app. Note that the `installation_target`, `security_advisory`, and `meta` events are not included because they are global events and not specific to an installation.
	Id int `json:"id"` // Unique identifier of the GitHub app
	Installations_count int `json:"installations_count,omitempty"` // The number of installations associated with the GitHub app. Only returned when the integration is requesting details about itself.
	Slug string `json:"slug,omitempty"` // The slug name of the GitHub app
	Updated_at string `json:"updated_at"`
	Name string `json:"name"` // The name of the GitHub app
	Node_id string `json:"node_id"`
	Owner interface{} `json:"owner"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Total_engaged_users int `json:"total_engaged_users,omitempty"` // The number of users who used Copilot for Pull Requests on github.com to generate a pull request summary at least once.
	Repositories []map[string]interface{} `json:"repositories,omitempty"` // Repositories in which users used Copilot for Pull Requests to generate pull request summaries
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Pull_request_review_url string `json:"pull_request_review_url"` // The API URL to get the pull request review where the secret was detected.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Editor string `json:"editor"` // The selected editor for the assignment.
	Id int `json:"id"` // Unique identifier of the repository.
	Language string `json:"language"` // The programming language used in the assignment.
	Deadline string `json:"deadline"` // The time at which the assignment is due.
	Feedback_pull_requests_enabled bool `json:"feedback_pull_requests_enabled"` // Whether feedback pull request will be created on assignment acceptance.
	Invite_link string `json:"invite_link"` // The link that a student can use to accept the assignment.
	Max_members int `json:"max_members,omitempty"` // The maximum allowable members per team.
	Submitted int `json:"submitted"` // The number of students that have submitted the assignment.
	Max_teams int `json:"max_teams,omitempty"` // The maximum allowable teams for the assignment.
	Passing int `json:"passing"` // The number of students that have passed the assignment.
	Slug string `json:"slug"` // Sluggified name of the assignment.
	Accepted int `json:"accepted"` // The number of students that have accepted the assignment.
	Students_are_repo_admins bool `json:"students_are_repo_admins"` // Whether students are admins on created repository on accepted assignment.
	Classroom GeneratedType `json:"classroom"` // A GitHub Classroom classroom
	Invitations_enabled bool `json:"invitations_enabled"` // Whether the invitation link is enabled. Visiting an enabled invitation link will accept the assignment.
	Title string `json:"title"` // Assignment title.
	TypeField string `json:"type"` // Whether it's a Group Assignment or Individual Assignment.
	Public_repo bool `json:"public_repo"` // Whether an accepted assignment creates a public repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id int `json:"id"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Node_id string `json:"node_id"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Project_card map[string]interface{} `json:"project_card,omitempty"`
	Url string `json:"url"`
	Event string `json:"event"`
	Commit_id string `json:"commit_id"`
	Commit_url string `json:"commit_url"`
	Created_at string `json:"created_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Start_private_fork bool `json:"start_private_fork,omitempty"` // Whether to create a temporary private fork of the repository to collaborate on a fix.
	Cve_id string `json:"cve_id,omitempty"` // The Common Vulnerabilities and Exposures (CVE) ID.
	Cwe_ids []string `json:"cwe_ids,omitempty"` // A list of Common Weakness Enumeration (CWE) IDs.
	Summary string `json:"summary"` // A short summary of the advisory.
	Vulnerabilities []map[string]interface{} `json:"vulnerabilities"` // A product affected by the vulnerability detailed in a repository security advisory.
	Credits []map[string]interface{} `json:"credits,omitempty"` // A list of users receiving credit for their participation in the security advisory.
	Description string `json:"description"` // A detailed description of what the advisory impacts.
	Severity string `json:"severity,omitempty"` // The severity of the advisory. You must choose between setting this field or `cvss_vector_string`.
	Cvss_vector_string string `json:"cvss_vector_string,omitempty"` // The CVSS vector that calculates the severity of the advisory. You must choose between setting this field or `severity`.
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
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Permissions_upgraded map[string]interface{} `json:"permissions_upgraded"` // Requested permissions that elevate access for a previously approved request for access, categorized by type of permission.
	Repositories []map[string]interface{} `json:"repositories"` // An array of repository objects the token is requesting access to. This field is only populated when `repository_selection` is `subset`.
	Created_at string `json:"created_at"` // Date and time when the request for access was created.
	Repository_selection string `json:"repository_selection"` // Type of repository selection requested.
	Token_last_used_at string `json:"token_last_used_at"` // Date and time when the associated fine-grained personal access token was last used for authentication.
	Id int `json:"id"` // Unique identifier of the request for access via fine-grained personal access token. Used as the `pat_request_id` parameter in the list and review API calls.
	Permissions_added map[string]interface{} `json:"permissions_added"` // New requested permissions, categorized by type of permission.
	Permissions_result map[string]interface{} `json:"permissions_result"` // Permissions requested, categorized by type of permission. This field incorporates `permissions_added` and `permissions_upgraded`.
	Repository_count int `json:"repository_count"` // The number of repositories the token is requesting access to. This field is only populated when `repository_selection` is `subset`.
	Token_id int `json:"token_id"` // Unique identifier of the user's token. This field can also be found in audit log events and the organization's settings for their PAT grants.
	Token_expired bool `json:"token_expired"` // Whether the associated fine-grained personal access token has expired.
	Token_expires_at string `json:"token_expires_at"` // Date and time when the associated fine-grained personal access token expires.
	Token_name string `json:"token_name"` // The name given to the user's token. This field can also be found in an organization's settings page for Active Tokens.
	Owner GeneratedType `json:"owner"` // A GitHub user.
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
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Base_role string `json:"base_role,omitempty"` // The system role from which this role inherits permissions.
	Description string `json:"description,omitempty"` // A short description about who this role is for or what permissions it grants.
	Name string `json:"name"` // The name of the role.
	Updated_at string `json:"updated_at"` // The date and time the role was last updated.
	Permissions []string `json:"permissions"` // A list of permissions included in this role.
	Created_at string `json:"created_at"` // The date and time the role was created.
	Id int64 `json:"id"` // The unique identifier of the role.
	Source string `json:"source,omitempty"` // Source answers the question, "where did this role come from?"
	Organization GeneratedType `json:"organization"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Changes map[string]interface{} `json:"changes"` // The changes to the comment if the action was `edited`.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Number int `json:"number"` // The pull request number.
	Action string `json:"action"`
	Pull_request GeneratedType `json:"pull_request"`
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id string `json:"id"` // The ID of the image. Use this ID for the `image` parameter when creating a new larger runner.
	Platform string `json:"platform"` // The operating system of the image.
	Size_gb int `json:"size_gb"` // Image size in GB.
	Source string `json:"source"` // The image provider.
	Display_name string `json:"display_name"` // Display name for this image.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"` // The changes to the milestone if the action was `edited`.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Milestone Webhooksmilestone `json:"milestone"` // A collection of related issues and pull requests.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Starred_at interface{} `json:"starred_at"` // The time the star was created. This is a timestamp in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`. Will be `null` for the `deleted` action.
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
	Architecture string `json:"architecture"`
	Download_url string `json:"download_url"`
	Filename string `json:"filename"`
	Os string `json:"os"`
	Sha256_checksum string `json:"sha256_checksum,omitempty"`
	Temp_download_token string `json:"temp_download_token,omitempty"` // A short lived bearer token used to download the runner, if needed.
}

// Enterprise represents the Enterprise schema from the OpenAPI specification
type Enterprise struct {
	Id int `json:"id"` // Unique identifier of the enterprise
	Name string `json:"name"` // The name of the enterprise.
	Slug string `json:"slug"` // The slug url identifier for the enterprise.
	Description string `json:"description,omitempty"` // A short description of the enterprise.
	Node_id string `json:"node_id"`
	Avatar_url string `json:"avatar_url"`
	Created_at string `json:"created_at"`
	Website_url string `json:"website_url,omitempty"` // The enterprise's website URL.
	Html_url string `json:"html_url"`
	Updated_at string `json:"updated_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Created_at string `json:"created_at"`
	State string `json:"state"` // The state of the status.
	Environment_url string `json:"environment_url,omitempty"` // The URL for accessing your environment.
	Updated_at string `json:"updated_at"`
	Id int64 `json:"id"`
	Target_url string `json:"target_url"` // Closing down notice: the URL to associate with this status.
	Node_id string `json:"node_id"`
	Description string `json:"description"` // A short description of the status.
	Environment string `json:"environment,omitempty"` // The environment of the deployment that the status is for.
	Creator GeneratedType `json:"creator"` // A GitHub user.
	Log_url string `json:"log_url,omitempty"` // The URL to associate with this status.
	Repository_url string `json:"repository_url"`
	Deployment_url string `json:"deployment_url"`
	Url string `json:"url"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Lock_branch map[string]interface{} `json:"lock_branch,omitempty"` // Whether to set the branch as read-only. If this is true, users will not be able to push to the branch.
	Required_conversation_resolution map[string]interface{} `json:"required_conversation_resolution,omitempty"`
	Required_signatures map[string]interface{} `json:"required_signatures,omitempty"`
	Required_status_checks GeneratedType `json:"required_status_checks,omitempty"` // Status Check Policy
	Allow_deletions map[string]interface{} `json:"allow_deletions,omitempty"`
	Allow_force_pushes map[string]interface{} `json:"allow_force_pushes,omitempty"`
	Block_creations map[string]interface{} `json:"block_creations,omitempty"`
	Restrictions GeneratedType `json:"restrictions,omitempty"` // Branch Restriction Policy
	Url string `json:"url"`
	Required_linear_history map[string]interface{} `json:"required_linear_history,omitempty"`
	Required_pull_request_reviews map[string]interface{} `json:"required_pull_request_reviews,omitempty"`
	Allow_fork_syncing map[string]interface{} `json:"allow_fork_syncing,omitempty"` // Whether users can pull changes from upstream when the branch is locked. Set to `true` to allow fork syncing. Set to `false` to prevent fork syncing.
	Enforce_admins map[string]interface{} `json:"enforce_admins,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	TypeField string `json:"type"`
	Parameters map[string]interface{} `json:"parameters,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"` // The changes to the team if the action was `edited`.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository map[string]interface{} `json:"repository,omitempty"` // A git repository
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Team Webhooksteam1 `json:"team"` // Groups of organization members that gives permissions on specified repositories.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url"`
	User GeneratedType `json:"user"` // A GitHub user.
	Organization GeneratedType `json:"organization"` // A GitHub organization.
	Organization_url string `json:"organization_url"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Role string `json:"role"` // The user's membership type in the organization.
	State string `json:"state"` // The state of the member in the organization. The `pending` state indicates the user has not yet accepted an invitation.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request GeneratedType `json:"pull_request"`
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Column_url string `json:"column_url"`
	Note string `json:"note"`
	Project_url string `json:"project_url"`
	Url string `json:"url"`
	Archived bool `json:"archived,omitempty"` // Whether or not the card is archived
	Created_at string `json:"created_at"`
	Node_id string `json:"node_id"`
	Project_id string `json:"project_id,omitempty"`
	Content_url string `json:"content_url,omitempty"`
	Creator GeneratedType `json:"creator"` // A GitHub user.
	Id int64 `json:"id"` // The project card's ID
	Updated_at string `json:"updated_at"`
	Column_name string `json:"column_name,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Run_url string `json:"run_url,omitempty"` // URL of the corresponding run.
	Run_id int `json:"run_id,omitempty"` // ID of the corresponding run.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Performed_via_github_app GeneratedType `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Production_environment bool `json:"production_environment,omitempty"` // Specifies if the given environment is one that end-users directly interact with. Default: false.
	Statuses_url string `json:"statuses_url"`
	Task string `json:"task"` // Parameter to specify a task to execute
	Transient_environment bool `json:"transient_environment,omitempty"` // Specifies if the given environment is will no longer exist at some point in the future. Default: false.
	Created_at string `json:"created_at"`
	Environment string `json:"environment"` // Name for the target deployment environment.
	Repository_url string `json:"repository_url"`
	Id int `json:"id"` // Unique identifier of the deployment
	Original_environment string `json:"original_environment,omitempty"`
	Updated_at string `json:"updated_at"`
	Description string `json:"description"`
	Node_id string `json:"node_id"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Commit map[string]interface{} `json:"commit"`
	Name string `json:"name"`
	Protected bool `json:"protected"`
	Protection GeneratedType `json:"protection,omitempty"` // Branch Protection
	Protection_url string `json:"protection_url,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Included_gigabytes_bandwidth int `json:"included_gigabytes_bandwidth"` // Free storage space (GB) for GitHub Packages.
	Total_gigabytes_bandwidth_used int `json:"total_gigabytes_bandwidth_used"` // Sum of the free and paid storage space (GB) for GitHuub Packages.
	Total_paid_gigabytes_bandwidth_used int `json:"total_paid_gigabytes_bandwidth_used"` // Total paid storage space (GB) for GitHuub Packages.
}

// Webhookslabel represents the Webhookslabel schema from the OpenAPI specification
type Webhookslabel struct {
	Color string `json:"color"` // 6-character hex code, without the leading #, identifying the color
	DefaultField bool `json:"default"`
	Description string `json:"description"`
	Id int `json:"id"`
	Name string `json:"name"` // The name of the label.
	Node_id string `json:"node_id"`
	Url string `json:"url"` // URL for the label
}

// Manifest represents the Manifest schema from the OpenAPI specification
type Manifest struct {
	Metadata Metadata `json:"metadata,omitempty"` // User-defined metadata to store domain-specific information limited to 8 keys with scalar values.
	Name string `json:"name"` // The name of the manifest.
	Resolved map[string]interface{} `json:"resolved,omitempty"` // A collection of resolved package dependencies.
	File map[string]interface{} `json:"file,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Marketplace_purchase map[string]interface{} `json:"marketplace_purchase"`
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Previous_marketplace_purchase Webhookspreviousmarketplacepurchase `json:"previous_marketplace_purchase,omitempty"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Effective_date string `json:"effective_date"`
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// Reaction represents the Reaction schema from the OpenAPI specification
type Reaction struct {
	User GeneratedType `json:"user"` // A GitHub user.
	Content string `json:"content"` // The reaction to use
	Created_at string `json:"created_at"`
	Id int `json:"id"`
	Node_id string `json:"node_id"`
}

// Thread represents the Thread schema from the OpenAPI specification
type Thread struct {
	Last_read_at string `json:"last_read_at"`
	Reason string `json:"reason"`
	Subject map[string]interface{} `json:"subject"`
	Subscription_url string `json:"subscription_url"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	Unread bool `json:"unread"`
	Id string `json:"id"`
	Repository GeneratedType `json:"repository"` // Minimal Repository
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Body string `json:"body"` // The comment text.
	Created_at string `json:"created_at"`
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	User GeneratedType `json:"user"` // A GitHub user.
	Author_association string `json:"author_association"` // How the author is associated with the repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Vulnerable_version_range string `json:"vulnerable_version_range"` // Conditions that identify vulnerable versions of this vulnerability's package.
	First_patched_version map[string]interface{} `json:"first_patched_version"` // Details pertaining to the package version that patches this vulnerability.
	PackageField GeneratedType `json:"package"` // Details for the vulnerable package.
	Severity string `json:"severity"` // The severity of the vulnerability.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Actor GeneratedType `json:"actor,omitempty"` // A GitHub user.
	Created_at string `json:"created_at"`
	Event string `json:"event"`
	Source map[string]interface{} `json:"source"`
	Updated_at string `json:"updated_at"`
}

// Repository represents the Repository schema from the OpenAPI specification
type Repository struct {
	Watchers_count int `json:"watchers_count"`
	Fork bool `json:"fork"`
	Open_issues int `json:"open_issues"`
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Forks int `json:"forks"`
	Homepage string `json:"homepage"`
	Allow_auto_merge bool `json:"allow_auto_merge,omitempty"` // Whether to allow Auto-merge to be used on pull requests.
	Url string `json:"url"`
	Git_commits_url string `json:"git_commits_url"`
	Id int64 `json:"id"` // Unique identifier of the repository
	Forks_url string `json:"forks_url"`
	Name string `json:"name"` // The name of the repository.
	Pulls_url string `json:"pulls_url"`
	Anonymous_access_enabled bool `json:"anonymous_access_enabled,omitempty"` // Whether anonymous git access is enabled for this repository
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"` // Whether to require contributors to sign off on web-based commits
	Created_at string `json:"created_at"`
	Teams_url string `json:"teams_url"`
	Is_template bool `json:"is_template,omitempty"` // Whether this repository acts as a template that can be used to generate new repositories.
	Subscription_url string `json:"subscription_url"`
	Deployments_url string `json:"deployments_url"`
	Pushed_at string `json:"pushed_at"`
	Archive_url string `json:"archive_url"`
	Size int `json:"size"` // The size of the repository, in kilobytes. Size is calculated hourly. When a repository is initially created, the size is 0.
	Hooks_url string `json:"hooks_url"`
	Contributors_url string `json:"contributors_url"`
	Labels_url string `json:"labels_url"`
	Events_url string `json:"events_url"`
	Merge_commit_message string `json:"merge_commit_message,omitempty"` // The default value for a merge commit message. - `PR_TITLE` - default to the pull request's title. - `PR_BODY` - default to the pull request's body. - `BLANK` - default to a blank commit message.
	Starred_at string `json:"starred_at,omitempty"`
	Statuses_url string `json:"statuses_url"`
	Updated_at string `json:"updated_at"`
	Private bool `json:"private"` // Whether the repository is private or public.
	Keys_url string `json:"keys_url"`
	Allow_forking bool `json:"allow_forking,omitempty"` // Whether to allow forking this repo
	Description string `json:"description"`
	Issue_comment_url string `json:"issue_comment_url"`
	Stargazers_count int `json:"stargazers_count"`
	Allow_update_branch bool `json:"allow_update_branch,omitempty"` // Whether or not a pull request head branch that is behind its base branch can always be updated even if it is not required to be up to date before merging.
	Languages_url string `json:"languages_url"`
	Full_name string `json:"full_name"`
	Has_projects bool `json:"has_projects"` // Whether projects are enabled.
	Watchers int `json:"watchers"`
	Disabled bool `json:"disabled"` // Returns whether or not this repository disabled.
	Archived bool `json:"archived"` // Whether the repository is archived.
	Has_downloads bool `json:"has_downloads"` // Whether downloads are enabled.
	Use_squash_pr_title_as_default bool `json:"use_squash_pr_title_as_default,omitempty"` // Whether a squash merge commit can use the pull request title as default. **This property is closing down. Please use `squash_merge_commit_title` instead.
	Stargazers_url string `json:"stargazers_url"`
	Subscribers_url string `json:"subscribers_url"`
	Collaborators_url string `json:"collaborators_url"`
	Comments_url string `json:"comments_url"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Master_branch string `json:"master_branch,omitempty"`
	Squash_merge_commit_message string `json:"squash_merge_commit_message,omitempty"` // The default value for a squash merge commit message: - `PR_BODY` - default to the pull request's body. - `COMMIT_MESSAGES` - default to the branch's commit messages. - `BLANK` - default to a blank commit message.
	Has_discussions bool `json:"has_discussions,omitempty"` // Whether discussions are enabled.
	Blobs_url string `json:"blobs_url"`
	Forks_count int `json:"forks_count"`
	Merges_url string `json:"merges_url"`
	Contents_url string `json:"contents_url"`
	Assignees_url string `json:"assignees_url"`
	Open_issues_count int `json:"open_issues_count"`
	Ssh_url string `json:"ssh_url"`
	License GeneratedType `json:"license"` // License Simple
	Git_tags_url string `json:"git_tags_url"`
	Language string `json:"language"`
	Temp_clone_token string `json:"temp_clone_token,omitempty"`
	Commits_url string `json:"commits_url"`
	Issues_url string `json:"issues_url"`
	Branches_url string `json:"branches_url"`
	Svn_url string `json:"svn_url"`
	Tags_url string `json:"tags_url"`
	Issue_events_url string `json:"issue_events_url"`
	Has_wiki bool `json:"has_wiki"` // Whether the wiki is enabled.
	Allow_merge_commit bool `json:"allow_merge_commit,omitempty"` // Whether to allow merge commits for pull requests.
	Clone_url string `json:"clone_url"`
	Html_url string `json:"html_url"`
	Merge_commit_title string `json:"merge_commit_title,omitempty"` // The default value for a merge commit title. - `PR_TITLE` - default to the pull request's title. - `MERGE_MESSAGE` - default to the classic title for a merge message (e.g., Merge pull request #123 from branch-name).
	Mirror_url string `json:"mirror_url"`
	Node_id string `json:"node_id"`
	Visibility string `json:"visibility,omitempty"` // The repository visibility: public, private, or internal.
	Has_issues bool `json:"has_issues"` // Whether issues are enabled.
	Releases_url string `json:"releases_url"`
	Has_pages bool `json:"has_pages"`
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"` // Whether to delete head branches when pull requests are merged
	Notifications_url string `json:"notifications_url"`
	Allow_squash_merge bool `json:"allow_squash_merge,omitempty"` // Whether to allow squash merges for pull requests.
	Trees_url string `json:"trees_url"`
	Topics []string `json:"topics,omitempty"`
	Code_search_index_status map[string]interface{} `json:"code_search_index_status,omitempty"` // The status of the code search index for this repository
	Squash_merge_commit_title string `json:"squash_merge_commit_title,omitempty"` // The default value for a squash merge commit title: - `PR_TITLE` - default to the pull request's title. - `COMMIT_OR_PR_TITLE` - default to the commit's title (if only one commit) or the pull request's title (when more than one commit).
	Downloads_url string `json:"downloads_url"`
	Default_branch string `json:"default_branch"` // The default branch of the repository.
	Compare_url string `json:"compare_url"`
	Git_url string `json:"git_url"`
	Milestones_url string `json:"milestones_url"`
	Git_refs_url string `json:"git_refs_url"`
	Allow_rebase_merge bool `json:"allow_rebase_merge,omitempty"` // Whether to allow rebase merges for pull requests.
}

// Webhooksprojectcolumn represents the Webhooksprojectcolumn schema from the OpenAPI specification
type Webhooksprojectcolumn struct {
	Updated_at string `json:"updated_at"`
	Name string `json:"name"` // Name of the project column
	Id int `json:"id"` // The unique identifier of the project column
	Node_id string `json:"node_id"`
	After_id int `json:"after_id,omitempty"`
	Cards_url string `json:"cards_url"`
	Project_url string `json:"project_url"`
	Url string `json:"url"`
	Created_at string `json:"created_at"`
}

// Label represents the Label schema from the OpenAPI specification
type Label struct {
	Name string `json:"name"` // The name of the label.
	Node_id string `json:"node_id"`
	Url string `json:"url"` // URL for the label
	Color string `json:"color"` // 6-character hex code, without the leading #, identifying the color
	DefaultField bool `json:"default"` // Whether this label comes by default in a new repository.
	Description string `json:"description"` // Optional description of the label, such as its purpose.
	Id int64 `json:"id"` // Unique identifier for the label.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Reset int `json:"reset"`
	Used int `json:"used"`
	Limit int `json:"limit"`
	Remaining int `json:"remaining"`
}

// Webhooksworkflow represents the Webhooksworkflow schema from the OpenAPI specification
type Webhooksworkflow struct {
	Url string `json:"url"`
	Badge_url string `json:"badge_url"`
	Node_id string `json:"node_id"`
	Path string `json:"path"`
	Html_url string `json:"html_url"`
	Id int `json:"id"`
	Name string `json:"name"`
	State string `json:"state"`
	Updated_at string `json:"updated_at"`
	Created_at string `json:"created_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Pulls_url string `json:"pulls_url"`
	Open_issues_count int `json:"open_issues_count"`
	Keys_url string `json:"keys_url"`
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"` // Whether to delete head branches when pull requests are merged
	Git_tags_url string `json:"git_tags_url"`
	Merge_commit_message string `json:"merge_commit_message,omitempty"` // The default value for a merge commit message. - `PR_TITLE` - default to the pull request's title. - `PR_BODY` - default to the pull request's body. - `BLANK` - default to a blank commit message.
	Visibility string `json:"visibility,omitempty"` // The repository visibility: public, private, or internal.
	Statuses_url string `json:"statuses_url"`
	Trees_url string `json:"trees_url"`
	Allow_forking bool `json:"allow_forking,omitempty"` // Whether to allow forking this repo
	Mirror_url string `json:"mirror_url"`
	Has_wiki bool `json:"has_wiki"` // Whether the wiki is enabled.
	Merge_commit_title string `json:"merge_commit_title,omitempty"` // The default value for a merge commit title. - `PR_TITLE` - default to the pull request's title. - `MERGE_MESSAGE` - default to the classic title for a merge message (e.g., Merge pull request #123 from branch-name).
	Size int `json:"size"` // The size of the repository, in kilobytes. Size is calculated hourly. When a repository is initially created, the size is 0.
	Disabled bool `json:"disabled"` // Returns whether or not this repository disabled.
	Temp_clone_token string `json:"temp_clone_token,omitempty"`
	Default_branch string `json:"default_branch"` // The default branch of the repository.
	Languages_url string `json:"languages_url"`
	Commits_url string `json:"commits_url"`
	Issues_url string `json:"issues_url"`
	Subscribers_url string `json:"subscribers_url"`
	Is_template bool `json:"is_template,omitempty"` // Whether this repository acts as a template that can be used to generate new repositories.
	Open_issues int `json:"open_issues"`
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"` // Whether to require contributors to sign off on web-based commits
	Anonymous_access_enabled bool `json:"anonymous_access_enabled,omitempty"` // Whether anonymous git access is enabled for this repository
	Language string `json:"language"`
	Teams_url string `json:"teams_url"`
	Milestones_url string `json:"milestones_url"`
	Labels_url string `json:"labels_url"`
	Git_url string `json:"git_url"`
	Created_at string `json:"created_at"`
	Issue_events_url string `json:"issue_events_url"`
	Squash_merge_commit_title string `json:"squash_merge_commit_title,omitempty"` // The default value for a squash merge commit title: - `PR_TITLE` - default to the pull request's title. - `COMMIT_OR_PR_TITLE` - default to the commit's title (if only one commit) or the pull request's title (when more than one commit).
	Updated_at string `json:"updated_at"`
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Allow_auto_merge bool `json:"allow_auto_merge,omitempty"` // Whether to allow Auto-merge to be used on pull requests.
	Squash_merge_commit_message string `json:"squash_merge_commit_message,omitempty"` // The default value for a squash merge commit message: - `PR_BODY` - default to the pull request's body. - `COMMIT_MESSAGES` - default to the branch's commit messages. - `BLANK` - default to a blank commit message.
	Contents_url string `json:"contents_url"`
	Url string `json:"url"`
	Allow_squash_merge bool `json:"allow_squash_merge,omitempty"` // Whether to allow squash merges for pull requests.
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Archive_url string `json:"archive_url"`
	Notifications_url string `json:"notifications_url"`
	Contributors_url string `json:"contributors_url"`
	Git_commits_url string `json:"git_commits_url"`
	Code_search_index_status map[string]interface{} `json:"code_search_index_status,omitempty"` // The status of the code search index for this repository
	Html_url string `json:"html_url"`
	Fork bool `json:"fork"`
	Description string `json:"description"`
	Name string `json:"name"` // The name of the repository.
	Has_pages bool `json:"has_pages"`
	Starred_at string `json:"starred_at,omitempty"`
	Hooks_url string `json:"hooks_url"`
	Issue_comment_url string `json:"issue_comment_url"`
	Allow_update_branch bool `json:"allow_update_branch,omitempty"` // Whether or not a pull request head branch that is behind its base branch can always be updated even if it is not required to be up to date before merging.
	Downloads_url string `json:"downloads_url"`
	Has_downloads bool `json:"has_downloads"` // Whether downloads are enabled.
	Git_refs_url string `json:"git_refs_url"`
	Deployments_url string `json:"deployments_url"`
	Merges_url string `json:"merges_url"`
	Watchers int `json:"watchers"`
	Master_branch string `json:"master_branch,omitempty"`
	Allow_rebase_merge bool `json:"allow_rebase_merge,omitempty"` // Whether to allow rebase merges for pull requests.
	Branches_url string `json:"branches_url"`
	Events_url string `json:"events_url"`
	Compare_url string `json:"compare_url"`
	Assignees_url string `json:"assignees_url"`
	License GeneratedType `json:"license"` // License Simple
	Archived bool `json:"archived"` // Whether the repository is archived.
	Has_discussions bool `json:"has_discussions,omitempty"` // Whether discussions are enabled.
	Forks_url string `json:"forks_url"`
	Releases_url string `json:"releases_url"`
	Collaborators_url string `json:"collaborators_url"`
	Full_name string `json:"full_name"`
	Id int64 `json:"id"` // Unique identifier of the repository
	Node_id string `json:"node_id"`
	Watchers_count int `json:"watchers_count"`
	Tags_url string `json:"tags_url"`
	Private bool `json:"private"` // Whether the repository is private or public.
	Allow_merge_commit bool `json:"allow_merge_commit,omitempty"` // Whether to allow merge commits for pull requests.
	Blobs_url string `json:"blobs_url"`
	Clone_url string `json:"clone_url"`
	Topics []string `json:"topics,omitempty"`
	Forks_count int `json:"forks_count"`
	Homepage string `json:"homepage"`
	Pushed_at string `json:"pushed_at"`
	Svn_url string `json:"svn_url"`
	Forks int `json:"forks"`
	Has_issues bool `json:"has_issues"` // Whether issues are enabled.
	Has_projects bool `json:"has_projects"` // Whether projects are enabled.
	Use_squash_pr_title_as_default bool `json:"use_squash_pr_title_as_default,omitempty"` // Whether a squash merge commit can use the pull request title as default. **This property is closing down. Please use `squash_merge_commit_title` instead.
	Comments_url string `json:"comments_url"`
	Stargazers_url string `json:"stargazers_url"`
	Stargazers_count int `json:"stargazers_count"`
	Subscription_url string `json:"subscription_url"`
	Ssh_url string `json:"ssh_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	TypeField string `json:"type"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Label Webhookslabel `json:"label"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Message string `json:"message"` // Message describing the purpose of the commit
	Tree map[string]interface{} `json:"tree"`
	Event string `json:"event,omitempty"`
	Html_url string `json:"html_url"`
	Parents []map[string]interface{} `json:"parents"`
	Sha string `json:"sha"` // SHA for the commit
	Verification map[string]interface{} `json:"verification"`
	Author map[string]interface{} `json:"author"` // Identifying information for the git-user
	Committer map[string]interface{} `json:"committer"` // Identifying information for the git-user
	Url string `json:"url"`
	Node_id string `json:"node_id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Check_suite map[string]interface{} `json:"check_suite"` // The [check_suite](https://docs.github.com/rest/checks/suites#get-a-check-suite).
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id int `json:"id"`
	Event string `json:"event"`
	Label map[string]interface{} `json:"label"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Node_id string `json:"node_id"`
	Commit_url string `json:"commit_url"`
	Created_at string `json:"created_at"`
	Url string `json:"url"`
	Commit_id string `json:"commit_id"`
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
	Single_file string `json:"single_file,omitempty"`
	Token string `json:"token"` // The token used for authentication
	Expires_at string `json:"expires_at"` // The time this token expires
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Repositories []Repository `json:"repositories,omitempty"` // The repositories this token has access to
	Repository_selection string `json:"repository_selection,omitempty"` // Describe whether all repositories have been selected or there's a selection involved
}

// Webhooksmilestone3 represents the Webhooksmilestone3 schema from the OpenAPI specification
type Webhooksmilestone3 struct {
	Id int `json:"id"`
	Url string `json:"url"`
	Title string `json:"title"` // The title of the milestone.
	Updated_at string `json:"updated_at"`
	Description string `json:"description"`
	Node_id string `json:"node_id"`
	Labels_url string `json:"labels_url"`
	Closed_at string `json:"closed_at"`
	Html_url string `json:"html_url"`
	Created_at string `json:"created_at"`
	Creator map[string]interface{} `json:"creator"`
	State string `json:"state"` // The state of the milestone.
	Closed_issues int `json:"closed_issues"`
	Open_issues int `json:"open_issues"`
	Due_on string `json:"due_on"`
	Number int `json:"number"` // The number of the milestone.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Permission string `json:"permission"`
	User GeneratedType `json:"user"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository"` // A GitHub repository.
	Result_count int `json:"result_count,omitempty"` // The number of results in the case of a successful analysis. This is only available for successful analyses.
	Source_location_prefix string `json:"source_location_prefix,omitempty"` // The source location prefix to use. This is only available for successful analyses.
	Analysis_status string `json:"analysis_status"` // The new status of the CodeQL variant analysis repository task.
	Artifact_size_in_bytes int `json:"artifact_size_in_bytes,omitempty"` // The size of the artifact. This is only available for successful analyses.
	Artifact_url string `json:"artifact_url,omitempty"` // The URL of the artifact. This is only available for successful analyses.
	Database_commit_sha string `json:"database_commit_sha,omitempty"` // The SHA of the commit the CodeQL database was built against. This is only available for successful analyses.
	Failure_message string `json:"failure_message,omitempty"` // The reason of the failure of this repo task. This is only available if the repository task has failed.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Created_at string `json:"created_at"`
	Name string `json:"name"` // The name of the secret.
	Selected_repositories_url string `json:"selected_repositories_url,omitempty"`
	Updated_at string `json:"updated_at"`
	Visibility string `json:"visibility"` // Visibility of a secret
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
	Commit_url string `json:"commit_url"`
	Repository GeneratedType `json:"repository"` // Minimal Repository
	Sha string `json:"sha"`
	State string `json:"state"`
	Statuses []GeneratedType `json:"statuses"`
	Total_count int `json:"total_count"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Download_count int `json:"download_count"`
	Uploader GeneratedType `json:"uploader"` // A GitHub user.
	Url string `json:"url"`
	Label string `json:"label"`
	State string `json:"state"` // State of the release asset.
	Updated_at string `json:"updated_at"`
	Created_at string `json:"created_at"`
	Digest string `json:"digest"`
	Content_type string `json:"content_type"`
	Id int `json:"id"`
	Name string `json:"name"` // The file name of the asset.
	Node_id string `json:"node_id"`
	Size int `json:"size"`
	Browser_download_url string `json:"browser_download_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Pushed_at string `json:"pushed_at,omitempty"`
	Hooks_url string `json:"hooks_url"`
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"`
	Default_branch string `json:"default_branch,omitempty"`
	Created_at string `json:"created_at,omitempty"`
	Open_issues_count int `json:"open_issues_count,omitempty"`
	Events_url string `json:"events_url"`
	Language string `json:"language,omitempty"`
	Allow_forking bool `json:"allow_forking,omitempty"`
	Blobs_url string `json:"blobs_url"`
	Git_commits_url string `json:"git_commits_url"`
	Branches_url string `json:"branches_url"`
	Owner GeneratedType `json:"owner"` // A GitHub user.
	License map[string]interface{} `json:"license,omitempty"`
	Milestones_url string `json:"milestones_url"`
	Watchers int `json:"watchers,omitempty"`
	Ssh_url string `json:"ssh_url,omitempty"`
	Subscribers_count int `json:"subscribers_count,omitempty"`
	Forks_count int `json:"forks_count,omitempty"`
	Is_template bool `json:"is_template,omitempty"`
	Has_discussions bool `json:"has_discussions,omitempty"`
	Languages_url string `json:"languages_url"`
	Merges_url string `json:"merges_url"`
	Has_downloads bool `json:"has_downloads,omitempty"`
	Releases_url string `json:"releases_url"`
	Keys_url string `json:"keys_url"`
	Tags_url string `json:"tags_url"`
	Has_pages bool `json:"has_pages,omitempty"`
	Watchers_count int `json:"watchers_count,omitempty"`
	Git_tags_url string `json:"git_tags_url"`
	Archived bool `json:"archived,omitempty"`
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"`
	Deployments_url string `json:"deployments_url"`
	Teams_url string `json:"teams_url"`
	Temp_clone_token string `json:"temp_clone_token,omitempty"`
	Labels_url string `json:"labels_url"`
	Comments_url string `json:"comments_url"`
	Id int64 `json:"id"`
	Issues_url string `json:"issues_url"`
	Fork bool `json:"fork"`
	Full_name string `json:"full_name"`
	Statuses_url string `json:"statuses_url"`
	Stargazers_count int `json:"stargazers_count,omitempty"`
	Commits_url string `json:"commits_url"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Html_url string `json:"html_url"`
	Has_projects bool `json:"has_projects,omitempty"`
	Topics []string `json:"topics,omitempty"`
	Trees_url string `json:"trees_url"`
	Issue_events_url string `json:"issue_events_url"`
	Size int `json:"size,omitempty"` // The size of the repository, in kilobytes. Size is calculated hourly. When a repository is initially created, the size is 0.
	Svn_url string `json:"svn_url,omitempty"`
	Code_of_conduct GeneratedType `json:"code_of_conduct,omitempty"` // Code Of Conduct
	Contributors_url string `json:"contributors_url"`
	Has_issues bool `json:"has_issues,omitempty"`
	Description string `json:"description"`
	Git_url string `json:"git_url,omitempty"`
	Clone_url string `json:"clone_url,omitempty"`
	Stargazers_url string `json:"stargazers_url"`
	Subscription_url string `json:"subscription_url"`
	Url string `json:"url"`
	Updated_at string `json:"updated_at,omitempty"`
	Collaborators_url string `json:"collaborators_url"`
	Has_wiki bool `json:"has_wiki,omitempty"`
	Custom_properties map[string]interface{} `json:"custom_properties,omitempty"` // The custom properties that were defined for the repository. The keys are the custom property names, and the values are the corresponding custom property values.
	Subscribers_url string `json:"subscribers_url"`
	Forks int `json:"forks,omitempty"`
	Git_refs_url string `json:"git_refs_url"`
	Issue_comment_url string `json:"issue_comment_url"`
	Archive_url string `json:"archive_url"`
	Name string `json:"name"`
	Open_issues int `json:"open_issues,omitempty"`
	Assignees_url string `json:"assignees_url"`
	Notifications_url string `json:"notifications_url"`
	Compare_url string `json:"compare_url"`
	Private bool `json:"private"`
	Security_and_analysis GeneratedType `json:"security_and_analysis,omitempty"`
	Forks_url string `json:"forks_url"`
	Mirror_url string `json:"mirror_url,omitempty"`
	Downloads_url string `json:"downloads_url"`
	Homepage string `json:"homepage,omitempty"`
	Visibility string `json:"visibility,omitempty"`
	Disabled bool `json:"disabled,omitempty"`
	Network_count int `json:"network_count,omitempty"`
	Node_id string `json:"node_id"`
	Role_name string `json:"role_name,omitempty"`
	Pulls_url string `json:"pulls_url"`
	Contents_url string `json:"contents_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Created_at string `json:"created_at"` // The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Number int `json:"number"` // The security alert number.
	State string `json:"state"` // State of a code scanning alert.
	Dismissed_at string `json:"dismissed_at"` // The time that the alert was dismissed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Dismissed_reason string `json:"dismissed_reason"` // **Required when the state is dismissed.** The reason for dismissing or closing the alert.
	Rule GeneratedType `json:"rule"`
	Tool GeneratedType `json:"tool"`
	Dismissed_comment string `json:"dismissed_comment,omitempty"` // The dismissal comment associated with the dismissal of the alert.
	Instances_url string `json:"instances_url"` // The REST API URL for fetching the list of instances for an alert.
	Url string `json:"url"` // The REST API URL of the alert resource.
	Html_url string `json:"html_url"` // The GitHub URL of the alert resource.
	Updated_at string `json:"updated_at,omitempty"` // The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Repository GeneratedType `json:"repository"` // A GitHub repository.
	Fixed_at string `json:"fixed_at,omitempty"` // The time that the alert was no longer detected and was considered fixed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Dismissal_approved_by GeneratedType `json:"dismissal_approved_by,omitempty"` // A GitHub user.
	Dismissed_by GeneratedType `json:"dismissed_by"` // A GitHub user.
	Most_recent_instance GeneratedType `json:"most_recent_instance"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	State string `json:"state"` // State of a code scanning alert.
	Instances_url string `json:"instances_url"` // The REST API URL for fetching the list of instances for an alert.
	Html_url string `json:"html_url"` // The GitHub URL of the alert resource.
	Number int `json:"number"` // The security alert number.
	Dismissal_approved_by GeneratedType `json:"dismissal_approved_by,omitempty"` // A GitHub user.
	Dismissed_by GeneratedType `json:"dismissed_by"` // A GitHub user.
	Updated_at string `json:"updated_at,omitempty"` // The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Dismissed_at string `json:"dismissed_at"` // The time that the alert was dismissed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Dismissed_reason string `json:"dismissed_reason"` // **Required when the state is dismissed.** The reason for dismissing or closing the alert.
	Most_recent_instance GeneratedType `json:"most_recent_instance"`
	Fixed_at string `json:"fixed_at,omitempty"` // The time that the alert was no longer detected and was considered fixed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Created_at string `json:"created_at"` // The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Rule GeneratedType `json:"rule"`
	Url string `json:"url"` // The REST API URL of the alert resource.
	Dismissed_comment string `json:"dismissed_comment,omitempty"` // The dismissal comment associated with the dismissal of the alert.
	Tool GeneratedType `json:"tool"`
}

// Autolink represents the Autolink schema from the OpenAPI specification
type Autolink struct {
	Is_alphanumeric bool `json:"is_alphanumeric"` // Whether this autolink reference matches alphanumeric characters. If false, this autolink reference only matches numeric characters.
	Key_prefix string `json:"key_prefix"` // The prefix of a key that is linkified.
	Url_template string `json:"url_template"` // A template for the target URL that is generated if a key was found.
	Id int `json:"id"`
}

// Classroom represents the Classroom schema from the OpenAPI specification
type Classroom struct {
	Name string `json:"name"` // The name of the classroom.
	Organization GeneratedType `json:"organization"` // A GitHub organization.
	Url string `json:"url"` // The URL of the classroom on GitHub Classroom.
	Archived bool `json:"archived"` // Whether classroom is archived.
	Id int `json:"id"` // Unique identifier of the classroom.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Issue interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) the comment belongs to.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Comment Webhooksissuecomment `json:"comment"` // The [comment](https://docs.github.com/rest/issues/comments#get-an-issue-comment) itself.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Created_at string `json:"created_at"` // The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Dismissed_comment string `json:"dismissed_comment"` // An optional comment associated with the alert's dismissal.
	Dismissed_reason string `json:"dismissed_reason"` // The reason that the alert was dismissed.
	Url string `json:"url"` // The REST API URL of the alert resource.
	Dismissed_by GeneratedType `json:"dismissed_by"` // A GitHub user.
	State string `json:"state"` // The state of the Dependabot alert.
	Auto_dismissed_at string `json:"auto_dismissed_at,omitempty"` // The time that the alert was auto-dismissed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Dependency map[string]interface{} `json:"dependency"` // Details for the vulnerable dependency.
	Updated_at string `json:"updated_at"` // The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Html_url string `json:"html_url"` // The GitHub URL of the alert resource.
	Number int `json:"number"` // The security alert number.
	Security_vulnerability GeneratedType `json:"security_vulnerability"` // Details pertaining to one vulnerable version range for the advisory.
	Dismissed_at string `json:"dismissed_at"` // The time that the alert was dismissed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Fixed_at string `json:"fixed_at"` // The time that the alert was no longer detected and was considered fixed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Repository GeneratedType `json:"repository"` // A GitHub repository.
	Security_advisory GeneratedType `json:"security_advisory"` // Details for the GitHub Security Advisory.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Previous_marketplace_purchase map[string]interface{} `json:"previous_marketplace_purchase,omitempty"`
	Effective_date string `json:"effective_date"`
	Marketplace_purchase Webhooksmarketplacepurchase `json:"marketplace_purchase"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Name string `json:"name"`
	File_size int `json:"file_size,omitempty"`
	Git_url string `json:"git_url"`
	Html_url string `json:"html_url"`
	Language string `json:"language,omitempty"`
	Line_numbers []string `json:"line_numbers,omitempty"`
	Sha string `json:"sha"`
	Text_matches []map[string]interface{} `json:"text_matches,omitempty"`
	Url string `json:"url"`
	Last_modified_at string `json:"last_modified_at,omitempty"`
	Path string `json:"path"`
	Repository GeneratedType `json:"repository"` // Minimal Repository
	Score float64 `json:"score"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Can_approve_pull_request_reviews bool `json:"can_approve_pull_request_reviews"` // Whether GitHub Actions can approve pull requests. Enabling this can be a security risk.
	Default_workflow_permissions string `json:"default_workflow_permissions"` // The default workflow permissions granted to the GITHUB_TOKEN when running workflows.
}

// Webhooksalert represents the Webhooksalert schema from the OpenAPI specification
type Webhooksalert struct {
	Fixed_in string `json:"fixed_in,omitempty"`
	Affected_package_name string `json:"affected_package_name"`
	External_identifier string `json:"external_identifier"`
	Fix_reason string `json:"fix_reason,omitempty"`
	Severity string `json:"severity"`
	Dismissed_at string `json:"dismissed_at,omitempty"`
	Ghsa_id string `json:"ghsa_id"`
	External_reference string `json:"external_reference"`
	Fixed_at string `json:"fixed_at,omitempty"`
	Dismiss_reason string `json:"dismiss_reason,omitempty"`
	Dismisser map[string]interface{} `json:"dismisser,omitempty"`
	Affected_range string `json:"affected_range"`
	Node_id string `json:"node_id"`
	Created_at string `json:"created_at"`
	Id int `json:"id"`
	State string `json:"state"`
	Number int `json:"number"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Origin string `json:"origin"`
	Expires_at string `json:"expires_at"`
	Limit string `json:"limit"` // The type of GitHub user that can comment, open issues, or create pull requests while the interaction limit is in effect.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	After string `json:"after"`
	Before string `json:"before"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Pull_request map[string]interface{} `json:"pull_request"`
	Action string `json:"action"`
	Number int `json:"number"` // The pull request number.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Git_tags_url string `json:"git_tags_url"`
	Size int `json:"size"`
	Ssh_url string `json:"ssh_url"`
	Issue_comment_url string `json:"issue_comment_url"`
	Allow_squash_merge bool `json:"allow_squash_merge,omitempty"` // Whether to allow squash merges for pull requests.
	Issue_events_url string `json:"issue_events_url"`
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Open_issues int `json:"open_issues"`
	Events_url string `json:"events_url"`
	Clone_url string `json:"clone_url"`
	Archive_url string `json:"archive_url"`
	Git_refs_url string `json:"git_refs_url"`
	Has_issues bool `json:"has_issues"` // Whether issues are enabled.
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"` // Whether to require contributors to sign off on web-based commits
	Contributors_url string `json:"contributors_url"`
	Updated_at string `json:"updated_at"`
	Collaborators_url string `json:"collaborators_url"`
	Subscribers_count int `json:"subscribers_count,omitempty"`
	Tags_url string `json:"tags_url"`
	Hooks_url string `json:"hooks_url"`
	Assignees_url string `json:"assignees_url"`
	Watchers int `json:"watchers"`
	Compare_url string `json:"compare_url"`
	Pulls_url string `json:"pulls_url"`
	Url string `json:"url"`
	Contents_url string `json:"contents_url"`
	Archived bool `json:"archived"` // Whether the repository is archived.
	Comments_url string `json:"comments_url"`
	Forks int `json:"forks"`
	Has_pages bool `json:"has_pages"`
	Git_commits_url string `json:"git_commits_url"`
	Stargazers_url string `json:"stargazers_url"`
	Forks_count int `json:"forks_count"`
	Role_name string `json:"role_name,omitempty"`
	Merges_url string `json:"merges_url"`
	Network_count int `json:"network_count,omitempty"`
	Svn_url string `json:"svn_url"`
	Private bool `json:"private"` // Whether the repository is private or public.
	License GeneratedType `json:"license"` // License Simple
	Subscription_url string `json:"subscription_url"`
	Topics []string `json:"topics,omitempty"`
	Description string `json:"description"`
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"` // Whether to delete head branches when pull requests are merged
	Temp_clone_token string `json:"temp_clone_token,omitempty"`
	Node_id string `json:"node_id"`
	Has_wiki bool `json:"has_wiki"` // Whether the wiki is enabled.
	Trees_url string `json:"trees_url"`
	Disabled bool `json:"disabled"` // Returns whether or not this repository disabled.
	Releases_url string `json:"releases_url"`
	Has_downloads bool `json:"has_downloads"` // Whether downloads are enabled.
	Pushed_at string `json:"pushed_at"`
	Keys_url string `json:"keys_url"`
	Has_projects bool `json:"has_projects"` // Whether projects are enabled.
	Commits_url string `json:"commits_url"`
	Created_at string `json:"created_at"`
	Allow_forking bool `json:"allow_forking,omitempty"` // Whether to allow forking this repo
	Allow_auto_merge bool `json:"allow_auto_merge,omitempty"` // Whether to allow Auto-merge to be used on pull requests.
	Allow_rebase_merge bool `json:"allow_rebase_merge,omitempty"` // Whether to allow rebase merges for pull requests.
	Downloads_url string `json:"downloads_url"`
	Id int `json:"id"` // Unique identifier of the repository
	Homepage string `json:"homepage"`
	Statuses_url string `json:"statuses_url"`
	Name string `json:"name"` // The name of the repository.
	Stargazers_count int `json:"stargazers_count"`
	Visibility string `json:"visibility,omitempty"` // The repository visibility: public, private, or internal.
	Open_issues_count int `json:"open_issues_count"`
	Issues_url string `json:"issues_url"`
	Language string `json:"language"`
	Teams_url string `json:"teams_url"`
	Notifications_url string `json:"notifications_url"`
	Deployments_url string `json:"deployments_url"`
	Master_branch string `json:"master_branch,omitempty"`
	Allow_merge_commit bool `json:"allow_merge_commit,omitempty"` // Whether to allow merge commits for pull requests.
	Html_url string `json:"html_url"`
	Subscribers_url string `json:"subscribers_url"`
	Full_name string `json:"full_name"`
	Forks_url string `json:"forks_url"`
	Is_template bool `json:"is_template,omitempty"` // Whether this repository acts as a template that can be used to generate new repositories.
	Blobs_url string `json:"blobs_url"`
	Default_branch string `json:"default_branch"` // The default branch of the repository.
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Languages_url string `json:"languages_url"`
	Milestones_url string `json:"milestones_url"`
	Watchers_count int `json:"watchers_count"`
	Branches_url string `json:"branches_url"`
	Fork bool `json:"fork"`
	Mirror_url string `json:"mirror_url"`
	Git_url string `json:"git_url"`
	Labels_url string `json:"labels_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sub_issue Issue `json:"sub_issue"` // Issues are a great way to keep track of tasks, enhancements, and bugs for your projects.
	Action string `json:"action"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Parent_issue_id float64 `json:"parent_issue_id"` // The ID of the parent issue.
	Parent_issue_repo Repository `json:"parent_issue_repo"` // A repository on GitHub.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sub_issue_id float64 `json:"sub_issue_id"` // The ID of the sub-issue.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Parent_issue Issue `json:"parent_issue"` // Issues are a great way to keep track of tasks, enhancements, and bugs for your projects.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Workflow Webhooksworkflow `json:"workflow"`
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Workflow_run map[string]interface{} `json:"workflow_run"`
	Deployment map[string]interface{} `json:"deployment"` // The [deployment](https://docs.github.com/rest/deployments/deployments#list-deployments).
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Pull_request_title_url string `json:"pull_request_title_url"` // The API URL to get the pull request where the secret was detected.
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
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Label Webhookslabel `json:"label"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Branch string `json:"branch"`
	Path string `json:"path"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	State string `json:"state"`
	Context string `json:"context"`
	Required bool `json:"required,omitempty"`
	Target_url string `json:"target_url"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	Avatar_url string `json:"avatar_url"`
	Created_at string `json:"created_at"`
	Node_id string `json:"node_id"`
	Description string `json:"description"`
	Id int `json:"id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action,omitempty"`
	Alert GeneratedType `json:"alert"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Location GeneratedType `json:"location"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Key string `json:"key"` // The Base64 encoded public key.
	Key_id string `json:"key_id"` // The identifier for the key.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Details interface{} `json:"details,omitempty"`
	TypeField string `json:"type,omitempty"` // The location type. Because secrets may be found in different types of resources (ie. code, comments, issues, pull requests, discussions), this field identifies the type of resource where the secret was found.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repositories_removed []map[string]interface{} `json:"repositories_removed"` // An array of repository objects, which were removed from the installation.
	Repository_selection string `json:"repository_selection"` // Describe whether all repositories have been selected or there's a selection involved
	Requester Webhooksuser `json:"requester"`
	Installation Installation `json:"installation"` // Installation
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Repositories_added []map[string]interface{} `json:"repositories_added"` // An array of repository objects, which were added to the installation.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
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
	Completed_at string `json:"completed_at,omitempty"` // The time that the scan was completed. Empty if the scan is running
	Started_at string `json:"started_at,omitempty"` // The time that the scan was started. Empty if the scan is pending
	Status string `json:"status,omitempty"` // The state of the scan. Either "completed", "running", or "pending"
	TypeField string `json:"type,omitempty"` // The type of scan
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Days []int `json:"days"`
	Total int `json:"total"`
	Week int `json:"week"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Commit_id string `json:"commit_id"`
	Commit_url string `json:"commit_url"`
	Created_at string `json:"created_at"`
	Event string `json:"event"`
	Milestone map[string]interface{} `json:"milestone"`
	Node_id string `json:"node_id"`
	Id int `json:"id"`
	Url string `json:"url"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Description string `json:"description,omitempty"` // A short description of the rule used to detect the alert.
	Full_description string `json:"full_description,omitempty"` // A description of the rule used to detect the alert.
	Id string `json:"id,omitempty"` // A unique identifier for the rule used to detect the alert.
	Security_severity_level string `json:"security_severity_level,omitempty"` // The security severity of the alert.
	Severity string `json:"severity,omitempty"` // The severity of the alert.
	Tags []string `json:"tags,omitempty"` // A set of tags applicable for the rule.
	Name string `json:"name,omitempty"` // The name of the rule used to detect the alert.
	Help string `json:"help,omitempty"` // Detailed documentation for the rule as GitHub Flavored Markdown.
	Help_uri string `json:"help_uri,omitempty"` // A link to the documentation for the rule used to detect the alert.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Key string `json:"key"` // The Base64 encoded public key.
	Key_id string `json:"key_id"` // The identifier for the key.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Title string `json:"title"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Verified_allowed bool `json:"verified_allowed,omitempty"` // Whether actions from GitHub Marketplace verified creators are allowed. Set to `true` to allow all actions by GitHub Marketplace verified creators.
	Github_owned_allowed bool `json:"github_owned_allowed,omitempty"` // Whether GitHub-owned actions are allowed. For example, this includes the actions in the `actions` organization.
	Patterns_allowed []string `json:"patterns_allowed,omitempty"` // Specifies a list of string-matching patterns to allow specific action(s) and reusable workflow(s). Wildcards, tags, and SHAs are allowed. For example, `monalisa/octocat@*`, `monalisa/octocat@v2`, `monalisa/*`. > [!NOTE] > The `patterns_allowed` setting only applies to public repositories.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Value string `json:"value"` // The value assigned to the property
	Property_name string `json:"property_name"` // The name of the property
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
	Label Webhookslabel `json:"label,omitempty"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue Webhooksissue `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Message string `json:"message,omitempty"`
	Status string `json:"status,omitempty"`
	Url string `json:"url,omitempty"`
	Documentation_url string `json:"documentation_url,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Body_text string `json:"body_text,omitempty"`
	Active_lock_reason string `json:"active_lock_reason,omitempty"`
	Closed_by GeneratedType `json:"closed_by,omitempty"` // A GitHub user.
	Pull_request map[string]interface{} `json:"pull_request,omitempty"`
	Id int64 `json:"id"`
	Assignee GeneratedType `json:"assignee"` // A GitHub user.
	Created_at string `json:"created_at"`
	TypeField GeneratedType `json:"type,omitempty"` // The type of issue.
	Title string `json:"title"` // Title of the issue
	Labels []interface{} `json:"labels"` // Labels to associate with this issue; pass one or more label names to replace the set of labels on this issue; send an empty array to clear all labels from the issue; note that the labels are silently dropped for users without push access to the repository
	Reactions GeneratedType `json:"reactions,omitempty"`
	User GeneratedType `json:"user"` // A GitHub user.
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Body_html string `json:"body_html,omitempty"`
	Html_url string `json:"html_url"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	State_reason string `json:"state_reason,omitempty"` // The reason for the current state
	Assignees []GeneratedType `json:"assignees,omitempty"`
	Repository_url string `json:"repository_url"`
	Number int `json:"number"` // Number uniquely identifying the issue within its repository
	Comments_url string `json:"comments_url"`
	Url string `json:"url"` // URL for the issue
	Labels_url string `json:"labels_url"`
	Draft bool `json:"draft,omitempty"`
	Closed_at string `json:"closed_at"`
	Comments int `json:"comments"`
	Milestone GeneratedType `json:"milestone"` // A collection of related issues and pull requests.
	Timeline_url string `json:"timeline_url,omitempty"`
	Sub_issues_summary GeneratedType `json:"sub_issues_summary,omitempty"`
	Updated_at string `json:"updated_at"`
	Node_id string `json:"node_id"`
	Locked bool `json:"locked"`
	State string `json:"state"` // State of the issue; either 'open' or 'closed'
	Events_url string `json:"events_url"`
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Body string `json:"body,omitempty"` // Contents of the issue
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Expires_at string `json:"expires_at"`
	Raw_key string `json:"raw_key"`
	Public_key string `json:"public_key"`
	Can_encrypt_storage bool `json:"can_encrypt_storage"`
	Created_at string `json:"created_at"`
	Id int64 `json:"id"`
	Name string `json:"name,omitempty"`
	Revoked bool `json:"revoked"`
	Emails []map[string]interface{} `json:"emails"`
	Can_encrypt_comms bool `json:"can_encrypt_comms"`
	Primary_key_id int `json:"primary_key_id"`
	Can_certify bool `json:"can_certify"`
	Subkeys []map[string]interface{} `json:"subkeys"`
	Key_id string `json:"key_id"`
	Can_sign bool `json:"can_sign"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Changes Webhookschanges `json:"changes"` // The changes to the comment.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Comment Webhooksreviewcomment `json:"comment"` // The [comment](https://docs.github.com/rest/pulls/comments#get-a-review-comment-for-a-pull-request) itself.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
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
	Invitation_teams_url string `json:"invitation_teams_url"`
	Node_id string `json:"node_id"`
	Role string `json:"role"`
	Invitation_source string `json:"invitation_source,omitempty"`
	Team_count int `json:"team_count"`
	Failed_at string `json:"failed_at,omitempty"`
	Login string `json:"login"`
	Created_at string `json:"created_at"`
	Id int64 `json:"id"`
	Inviter GeneratedType `json:"inviter"` // A GitHub user.
	Email string `json:"email"`
	Failed_reason string `json:"failed_reason,omitempty"`
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
	Memory_gb int `json:"memory_gb"` // The available RAM for the machine spec.
	Storage_gb int `json:"storage_gb"` // The available SSD storage for the machine spec.
	Cpu_cores int `json:"cpu_cores"` // The number of cores.
	Id string `json:"id"` // The ID used for the `size` parameter when creating a new runner.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Include_claim_keys []string `json:"include_claim_keys"` // Array of unique strings. Each claim key can only contain alphanumeric characters and underscores.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Approver Webhooksapprover `json:"approver,omitempty"`
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Since string `json:"since"`
	Workflow_job_run Webhooksworkflowjobrun `json:"workflow_job_run,omitempty"`
	Comment string `json:"comment,omitempty"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Workflow_run map[string]interface{} `json:"workflow_run"`
	Action string `json:"action"`
	Reviewers []map[string]interface{} `json:"reviewers,omitempty"`
	Workflow_job_runs []map[string]interface{} `json:"workflow_job_runs,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Events_url string `json:"events_url"`
	Allow_rebase_merge bool `json:"allow_rebase_merge,omitempty"`
	Fork bool `json:"fork"`
	Master_branch string `json:"master_branch,omitempty"`
	Archived bool `json:"archived"`
	Issue_comment_url string `json:"issue_comment_url"`
	Forks_count int `json:"forks_count"`
	Watchers int `json:"watchers"`
	Hooks_url string `json:"hooks_url"`
	Labels_url string `json:"labels_url"`
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"`
	Stargazers_url string `json:"stargazers_url"`
	Has_projects bool `json:"has_projects"`
	Downloads_url string `json:"downloads_url"`
	Template_repository GeneratedType `json:"template_repository,omitempty"` // A repository on GitHub.
	Merge_commit_title string `json:"merge_commit_title,omitempty"` // The default value for a merge commit title. - `PR_TITLE` - default to the pull request's title. - `MERGE_MESSAGE` - default to the classic title for a merge message (e.g., Merge pull request #123 from branch-name).
	Tags_url string `json:"tags_url"`
	Has_issues bool `json:"has_issues"`
	Subscription_url string `json:"subscription_url"`
	Trees_url string `json:"trees_url"`
	Is_template bool `json:"is_template,omitempty"`
	Allow_update_branch bool `json:"allow_update_branch,omitempty"`
	Has_pages bool `json:"has_pages"`
	Commits_url string `json:"commits_url"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub user.
	Svn_url string `json:"svn_url"`
	Parent Repository `json:"parent,omitempty"` // A repository on GitHub.
	Releases_url string `json:"releases_url"`
	Homepage string `json:"homepage"`
	Issue_events_url string `json:"issue_events_url"`
	Issues_url string `json:"issues_url"`
	Mirror_url string `json:"mirror_url"`
	Stargazers_count int `json:"stargazers_count"`
	Git_refs_url string `json:"git_refs_url"`
	Description string `json:"description"`
	Node_id string `json:"node_id"`
	Merges_url string `json:"merges_url"`
	Teams_url string `json:"teams_url"`
	Has_downloads bool `json:"has_downloads,omitempty"`
	Blobs_url string `json:"blobs_url"`
	Merge_commit_message string `json:"merge_commit_message,omitempty"` // The default value for a merge commit message. - `PR_TITLE` - default to the pull request's title. - `PR_BODY` - default to the pull request's body. - `BLANK` - default to a blank commit message.
	Squash_merge_commit_message string `json:"squash_merge_commit_message,omitempty"` // The default value for a squash merge commit message: - `PR_BODY` - default to the pull request's body. - `COMMIT_MESSAGES` - default to the branch's commit messages. - `BLANK` - default to a blank commit message.
	Subscribers_url string `json:"subscribers_url"`
	Updated_at string `json:"updated_at"`
	Anonymous_access_enabled bool `json:"anonymous_access_enabled,omitempty"` // Whether anonymous git access is allowed.
	Security_and_analysis GeneratedType `json:"security_and_analysis,omitempty"`
	Use_squash_pr_title_as_default bool `json:"use_squash_pr_title_as_default,omitempty"`
	Statuses_url string `json:"statuses_url"`
	License GeneratedType `json:"license"` // License Simple
	Pulls_url string `json:"pulls_url"`
	Topics []string `json:"topics,omitempty"`
	Archive_url string `json:"archive_url"`
	Git_url string `json:"git_url"`
	Name string `json:"name"`
	Collaborators_url string `json:"collaborators_url"`
	Allow_merge_commit bool `json:"allow_merge_commit,omitempty"`
	Allow_auto_merge bool `json:"allow_auto_merge,omitempty"`
	Compare_url string `json:"compare_url"`
	Temp_clone_token string `json:"temp_clone_token,omitempty"`
	Git_commits_url string `json:"git_commits_url"`
	Contents_url string `json:"contents_url"`
	Html_url string `json:"html_url"`
	Language string `json:"language"`
	Allow_forking bool `json:"allow_forking,omitempty"`
	Forks int `json:"forks"`
	Pushed_at string `json:"pushed_at"`
	Squash_merge_commit_title string `json:"squash_merge_commit_title,omitempty"` // The default value for a squash merge commit title: - `PR_TITLE` - default to the pull request's title. - `COMMIT_OR_PR_TITLE` - default to the commit's title (if only one commit) or the pull request's title (when more than one commit).
	Open_issues_count int `json:"open_issues_count"`
	Comments_url string `json:"comments_url"`
	Contributors_url string `json:"contributors_url"`
	Has_wiki bool `json:"has_wiki"`
	Milestones_url string `json:"milestones_url"`
	Network_count int `json:"network_count"`
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Notifications_url string `json:"notifications_url"`
	Git_tags_url string `json:"git_tags_url"`
	Size int `json:"size"` // The size of the repository, in kilobytes. Size is calculated hourly. When a repository is initially created, the size is 0.
	Url string `json:"url"`
	Allow_squash_merge bool `json:"allow_squash_merge,omitempty"`
	Languages_url string `json:"languages_url"`
	Id int64 `json:"id"`
	Watchers_count int `json:"watchers_count"`
	Subscribers_count int `json:"subscribers_count"`
	Private bool `json:"private"`
	Default_branch string `json:"default_branch"`
	Deployments_url string `json:"deployments_url"`
	Open_issues int `json:"open_issues"`
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Ssh_url string `json:"ssh_url"`
	Clone_url string `json:"clone_url"`
	Visibility string `json:"visibility,omitempty"` // The repository visibility: public, private, or internal.
	Assignees_url string `json:"assignees_url"`
	Disabled bool `json:"disabled"` // Returns whether or not this repository disabled.
	Source Repository `json:"source,omitempty"` // A repository on GitHub.
	Created_at string `json:"created_at"`
	Code_of_conduct GeneratedType `json:"code_of_conduct,omitempty"` // Code of Conduct Simple
	Forks_url string `json:"forks_url"`
	Full_name string `json:"full_name"`
	Has_discussions bool `json:"has_discussions"`
	Keys_url string `json:"keys_url"`
	Custom_properties map[string]interface{} `json:"custom_properties,omitempty"` // The custom properties that were defined for the repository. The keys are the custom property names, and the values are the corresponding custom property values.
	Branches_url string `json:"branches_url"`
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
	Runner_label string `json:"runner_label,omitempty"` // The label of the runner to use for code scanning default setup when runner_type is 'labeled'.
	Runner_type string `json:"runner_type,omitempty"` // Whether to use labeled runners or standard GitHub runners.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	TypeField string `json:"type"`
	Parameters map[string]interface{} `json:"parameters,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Name string `json:"name"` // The name pattern that branches or tags must match in order to deploy to the environment. Wildcard characters will not match `/`. For example, to match branches that begin with `release/` and contain an additional single slash, use `release/*/*`. For more information about pattern matching syntax, see the [Ruby File.fnmatch documentation](https://ruby-doc.org/core-2.5.1/File.html#method-c-fnmatch).
	TypeField string `json:"type,omitempty"` // Whether this rule targets a branch or tag
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Number int `json:"number"` // The security alert number.
	Instances_url string `json:"instances_url"` // The REST API URL for fetching the list of instances for an alert.
	Dismissal_approved_by GeneratedType `json:"dismissal_approved_by,omitempty"` // A GitHub user.
	Tool GeneratedType `json:"tool"`
	Url string `json:"url"` // The REST API URL of the alert resource.
	Rule GeneratedType `json:"rule"`
	Most_recent_instance GeneratedType `json:"most_recent_instance"`
	Dismissed_reason string `json:"dismissed_reason"` // **Required when the state is dismissed.** The reason for dismissing or closing the alert.
	Updated_at string `json:"updated_at,omitempty"` // The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Dismissed_comment string `json:"dismissed_comment,omitempty"` // The dismissal comment associated with the dismissal of the alert.
	Dismissed_at string `json:"dismissed_at"` // The time that the alert was dismissed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Dismissed_by GeneratedType `json:"dismissed_by"` // A GitHub user.
	Html_url string `json:"html_url"` // The GitHub URL of the alert resource.
	State string `json:"state"` // State of a code scanning alert.
	Created_at string `json:"created_at"` // The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Fixed_at string `json:"fixed_at,omitempty"` // The time that the alert was no longer detected and was considered fixed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url,omitempty"` // The URL target of the delivery.
	Action string `json:"action"` // The type of activity for the event that triggered the delivery.
	Id int `json:"id"` // Unique identifier of the delivery.
	Redelivery bool `json:"redelivery"` // Whether the delivery is a redelivery.
	Status_code int `json:"status_code"` // Status code received when delivery was made.
	Guid string `json:"guid"` // Unique identifier for the event (shared with all deliveries for all webhooks that subscribe to this event).
	Request map[string]interface{} `json:"request"`
	Repository_id int `json:"repository_id"` // The id of the repository associated with this event.
	Duration float64 `json:"duration"` // Time spent delivering.
	Event string `json:"event"` // The event that triggered the delivery.
	Response map[string]interface{} `json:"response"`
	Installation_id int `json:"installation_id"` // The id of the GitHub App installation associated with this event.
	Status string `json:"status"` // Description of the status of the attempted delivery
	Delivered_at string `json:"delivered_at"` // Time when the delivery was delivered.
	Throttled_at string `json:"throttled_at,omitempty"` // Time when the webhook delivery was throttled.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Pull_request map[string]interface{} `json:"pull_request"`
	Action string `json:"action"`
	Assignee Webhooksusermannequin `json:"assignee,omitempty"`
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	TypeField string `json:"type"`
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
	Archived bool `json:"archived"` // Returns whether classroom is archived or not.
	Id int `json:"id"` // Unique identifier of the classroom.
	Name string `json:"name"` // The name of the classroom.
	Url string `json:"url"` // The url of the classroom on GitHub Classroom.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Reviewer GeneratedType `json:"reviewer"` // A required reviewing team
	File_patterns []string `json:"file_patterns"` // Array of file patterns. Pull requests which change matching files must be approved by the specified team. File patterns use the same syntax as `.gitignore` files.
	Minimum_approvals int `json:"minimum_approvals"` // Minimum number of approvals required from the specified team. If set to zero, the team will be added to the pull request but approval is optional.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation Installation `json:"installation"` // Installation
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repositories []map[string]interface{} `json:"repositories,omitempty"` // An array of repository objects that the installation can access.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Requester interface{} `json:"requester,omitempty"`
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	PackageField map[string]interface{} `json:"package"` // Information about the package.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Public_ips map[string]interface{} `json:"public_ips"` // Provides details of static public IP limits for GitHub-hosted Hosted Runners
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Comment string `json:"comment"` // Comment associated with the pending deployment protection rule. **Required when state is not provided.**
	Environment_name string `json:"environment_name"` // The name of the environment to approve or reject.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Token_id int `json:"token_id"` // Unique identifier of the user's token. This field can also be found in audit log events and the organization's settings for their PAT grants.
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Reason string `json:"reason"` // Reason for requesting access.
	Token_expired bool `json:"token_expired"` // Whether the associated fine-grained personal access token has expired.
	Token_name string `json:"token_name"` // The name given to the user's token. This field can also be found in an organization's settings page for Active Tokens.
	Created_at string `json:"created_at"` // Date and time when the request for access was created.
	Id int `json:"id"` // Unique identifier of the request for access via fine-grained personal access token. The `pat_request_id` used to review PAT requests.
	Permissions map[string]interface{} `json:"permissions"` // Permissions requested, categorized by type of permission.
	Repositories_url string `json:"repositories_url"` // URL to the list of repositories requested to be accessed via fine-grained personal access token. Should only be followed when `repository_selection` is `subset`.
	Token_last_used_at string `json:"token_last_used_at"` // Date and time when the associated fine-grained personal access token was last used for authentication.
	Repository_selection string `json:"repository_selection"` // Type of repository selection requested.
	Token_expires_at string `json:"token_expires_at"` // Date and time when the associated fine-grained personal access token expires.
}

// Webhooksapprover represents the Webhooksapprover schema from the OpenAPI specification
type Webhooksapprover struct {
	Node_id string `json:"node_id,omitempty"`
	Organizations_url string `json:"organizations_url,omitempty"`
	Gists_url string `json:"gists_url,omitempty"`
	Login string `json:"login,omitempty"`
	Url string `json:"url,omitempty"`
	User_view_type string `json:"user_view_type,omitempty"`
	Starred_url string `json:"starred_url,omitempty"`
	Events_url string `json:"events_url,omitempty"`
	Received_events_url string `json:"received_events_url,omitempty"`
	Subscriptions_url string `json:"subscriptions_url,omitempty"`
	Followers_url string `json:"followers_url,omitempty"`
	Id int `json:"id,omitempty"`
	TypeField string `json:"type,omitempty"`
	Following_url string `json:"following_url,omitempty"`
	Gravatar_id string `json:"gravatar_id,omitempty"`
	Avatar_url string `json:"avatar_url,omitempty"`
	Site_admin bool `json:"site_admin,omitempty"`
	Repos_url string `json:"repos_url,omitempty"`
	Html_url string `json:"html_url,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Location GeneratedType `json:"location,omitempty"` // Describe a region within a file for the alert.
	Classifications []string `json:"classifications,omitempty"` // Classifications that have been applied to the file that triggered the alert. For example identifying it as documentation, or a generated file.
	Analysis_key string `json:"analysis_key,omitempty"` // Identifies the configuration under which the analysis was executed. For example, in GitHub Actions this includes the workflow filename and job name.
	Commit_sha string `json:"commit_sha,omitempty"`
	Message map[string]interface{} `json:"message,omitempty"`
	Category string `json:"category,omitempty"` // Identifies the configuration under which the analysis was executed. Used to distinguish between multiple analyses for the same tool and commit, but performed on different languages or different parts of the code.
	Ref string `json:"ref,omitempty"` // The Git reference, formatted as `refs/pull/<number>/merge`, `refs/pull/<number>/head`, `refs/heads/<branch name>` or simply `<branch name>`.
	State string `json:"state,omitempty"` // State of a code scanning alert.
	Environment string `json:"environment,omitempty"` // Identifies the variable values associated with the environment in which the analysis that generated this alert instance was performed, such as the language that was analyzed.
	Html_url string `json:"html_url,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"` // The action that was performed.
	Assignee Webhooksusermannequin `json:"assignee,omitempty"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue Webhooksissue `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// Webhooksteam1 represents the Webhooksteam1 schema from the OpenAPI specification
type Webhooksteam1 struct {
	Parent map[string]interface{} `json:"parent,omitempty"`
	Privacy string `json:"privacy,omitempty"`
	Url string `json:"url,omitempty"` // URL for the team
	Deleted bool `json:"deleted,omitempty"`
	Members_url string `json:"members_url,omitempty"`
	Description string `json:"description,omitempty"` // Description of the team
	Html_url string `json:"html_url,omitempty"`
	Permission string `json:"permission,omitempty"` // Permission that the team will have for its repositories
	Repositories_url string `json:"repositories_url,omitempty"`
	Slug string `json:"slug,omitempty"`
	Id int `json:"id"` // Unique identifier of the team
	Name string `json:"name"` // Name of the team
	Node_id string `json:"node_id,omitempty"`
	Notification_setting string `json:"notification_setting,omitempty"` // Whether team members will receive notifications when their team is @mentioned
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Branches_url string `json:"branches_url"` // A template for the API URL to get information about branches in the repository.
	Git_commits_url string `json:"git_commits_url"` // A template for the API URL to get information about Git commits of the repository.
	Git_refs_url string `json:"git_refs_url"` // A template for the API URL to get information about Git refs of the repository.
	Full_name string `json:"full_name"` // The full, globally unique, name of the repository.
	Comments_url string `json:"comments_url"` // A template for the API URL to get information about comments on the repository.
	Description string `json:"description"` // The repository description.
	Labels_url string `json:"labels_url"` // A template for the API URL to get information about labels of the repository.
	Stargazers_url string `json:"stargazers_url"` // The API URL to list the stargazers on the repository.
	Blobs_url string `json:"blobs_url"` // A template for the API URL to create or retrieve a raw Git blob in the repository.
	Issue_comment_url string `json:"issue_comment_url"` // A template for the API URL to get information about issue comments on the repository.
	Collaborators_url string `json:"collaborators_url"` // A template for the API URL to get information about collaborators of the repository.
	Fork bool `json:"fork"` // Whether the repository is a fork.
	Trees_url string `json:"trees_url"` // A template for the API URL to create or retrieve a raw Git tree of the repository.
	Pulls_url string `json:"pulls_url"` // A template for the API URL to get information about pull requests on the repository.
	Milestones_url string `json:"milestones_url"` // A template for the API URL to get information about milestones of the repository.
	Archive_url string `json:"archive_url"` // A template for the API URL to download the repository as an archive.
	Keys_url string `json:"keys_url"` // A template for the API URL to get information about deploy keys on the repository.
	Url string `json:"url"` // The URL to get more information about the repository from the GitHub API.
	Downloads_url string `json:"downloads_url"` // The API URL to list the downloads on the repository.
	Commits_url string `json:"commits_url"` // A template for the API URL to get information about commits on the repository.
	Git_tags_url string `json:"git_tags_url"` // A template for the API URL to get information about Git tags of the repository.
	Languages_url string `json:"languages_url"` // The API URL to get information about the languages of the repository.
	Forks_url string `json:"forks_url"` // The API URL to list the forks of the repository.
	Subscription_url string `json:"subscription_url"` // The API URL to subscribe to notifications for this repository.
	Contents_url string `json:"contents_url"` // A template for the API URL to get the contents of the repository.
	Merges_url string `json:"merges_url"` // The API URL to merge branches in the repository.
	Tags_url string `json:"tags_url"` // The API URL to get information about tags on the repository.
	Assignees_url string `json:"assignees_url"` // A template for the API URL to list the available assignees for issues in the repository.
	Compare_url string `json:"compare_url"` // A template for the API URL to compare two commits or refs.
	Teams_url string `json:"teams_url"` // The API URL to list the teams on the repository.
	Issues_url string `json:"issues_url"` // A template for the API URL to get information about issues on the repository.
	Private bool `json:"private"` // Whether the repository is private.
	Name string `json:"name"` // The name of the repository.
	Hooks_url string `json:"hooks_url"` // The API URL to list the hooks on the repository.
	Releases_url string `json:"releases_url"` // A template for the API URL to get information about releases on the repository.
	Issue_events_url string `json:"issue_events_url"` // A template for the API URL to get information about issue events on the repository.
	Node_id string `json:"node_id"` // The GraphQL identifier of the repository.
	Subscribers_url string `json:"subscribers_url"` // The API URL to list the subscribers on the repository.
	Statuses_url string `json:"statuses_url"` // A template for the API URL to get information about statuses of a commit.
	Events_url string `json:"events_url"` // The API URL to list the events of the repository.
	Deployments_url string `json:"deployments_url"` // The API URL to list the deployments of the repository.
	Html_url string `json:"html_url"` // The URL to view the repository on GitHub.com.
	Contributors_url string `json:"contributors_url"` // A template for the API URL to list the contributors to the repository.
	Notifications_url string `json:"notifications_url"` // A template for the API URL to get information about notifications on the repository.
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Id int64 `json:"id"` // A unique identifier of the repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Project Webhooksproject `json:"project"`
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Blocked_user Webhooksuser `json:"blocked_user"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Is_enabled bool `json:"is_enabled"` // Whether or not the issue type is enabled at the organization level.
	Name string `json:"name"` // Name of the issue type.
	Color string `json:"color,omitempty"` // Color for the issue type.
	Description string `json:"description,omitempty"` // Description of the issue type.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	State string `json:"state"`
	Description string `json:"description"`
	Domains []string `json:"domains"` // Array of the domain set and its alternate name (if it is configured)
	Expires_at string `json:"expires_at,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Reactions GeneratedType `json:"reactions,omitempty"`
	Path string `json:"path"` // The relative path of the file to which the comment applies.
	Body_html string `json:"body_html,omitempty"`
	In_reply_to_id int `json:"in_reply_to_id,omitempty"` // The comment ID to reply to.
	Original_position int `json:"original_position,omitempty"` // The index of the original line in the diff to which the comment applies. This field is closing down; use `original_line` instead.
	Created_at string `json:"created_at"`
	Pull_request_review_id int64 `json:"pull_request_review_id"` // The ID of the pull request review to which the comment belongs.
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Commit_id string `json:"commit_id"` // The SHA of the commit to which the comment applies.
	Original_commit_id string `json:"original_commit_id"` // The SHA of the original commit to which the comment applies.
	Position int `json:"position,omitempty"` // The line index in the diff to which the comment applies. This field is closing down; use `line` instead.
	Diff_hunk string `json:"diff_hunk"` // The diff of the line that the comment refers to.
	Id int64 `json:"id"` // The ID of the pull request review comment.
	Url string `json:"url"` // URL for the pull request review comment
	Body_text string `json:"body_text,omitempty"`
	Links map[string]interface{} `json:"_links"`
	Line int `json:"line,omitempty"` // The line of the blob to which the comment applies. The last line of the range for a multi-line comment
	Original_line int `json:"original_line,omitempty"` // The line of the blob to which the comment applies. The last line of the range for a multi-line comment
	Subject_type string `json:"subject_type,omitempty"` // The level at which the comment is targeted, can be a diff line or a file.
	Node_id string `json:"node_id"` // The node ID of the pull request review comment.
	Html_url string `json:"html_url"` // HTML URL for the pull request review comment.
	Start_line int `json:"start_line,omitempty"` // The first line of the range for a multi-line comment.
	Updated_at string `json:"updated_at"`
	User GeneratedType `json:"user"` // A GitHub user.
	Pull_request_url string `json:"pull_request_url"` // URL for the pull request that the review comment belongs to.
	Body string `json:"body"` // The text of the comment.
	Side string `json:"side,omitempty"` // The side of the diff to which the comment applies. The side of the last line of the range for a multi-line comment
	Start_side string `json:"start_side,omitempty"` // The side of the first line of the range for a multi-line comment.
	Original_start_line int `json:"original_start_line,omitempty"` // The first line of the range for a multi-line comment.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Ref string `json:"ref"` // The Git reference of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Alert map[string]interface{} `json:"alert"` // The code scanning alert involved in the event.
	Commit_oid string `json:"commit_oid"` // The commit SHA of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	From string `json:"from"`
	To string `json:"to"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Network_configuration_id string `json:"network_configuration_id,omitempty"` // The identifier of the network configuration that is using this settings resource.
	Region string `json:"region"` // The location of the subnet this network settings resource is configured for.
	Subnet_id string `json:"subnet_id"` // The subnet this network settings resource is configured for.
	Id string `json:"id"` // The unique identifier of the network settings resource.
	Name string `json:"name"` // The name of the network settings resource.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Personal_access_token_request GeneratedType `json:"personal_access_token_request"` // Details of a Personal Access Token Request.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Security_advisory Webhookssecurityadvisory `json:"security_advisory"` // The details of the security advisory, including summary, description, and severity.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
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
	Site_admin bool `json:"site_admin"`
	Gists_url string `json:"gists_url"`
	Html_url string `json:"html_url"`
	Starred_at string `json:"starred_at,omitempty"`
	Subscriptions_url string `json:"subscriptions_url"`
	Url string `json:"url"`
	Email string `json:"email,omitempty"`
	User_view_type string `json:"user_view_type,omitempty"`
	Repos_url string `json:"repos_url"`
	Starred_url string `json:"starred_url"`
	Organizations_url string `json:"organizations_url"`
	Avatar_url string `json:"avatar_url"`
	Gravatar_id string `json:"gravatar_id"`
	Followers_url string `json:"followers_url"`
	Following_url string `json:"following_url"`
	Received_events_url string `json:"received_events_url"`
	Node_id string `json:"node_id"`
	Login string `json:"login"`
	Name string `json:"name,omitempty"`
	TypeField string `json:"type"`
	Events_url string `json:"events_url"`
	Id int64 `json:"id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Has_multiple_single_files bool `json:"has_multiple_single_files,omitempty"`
	Permissions GeneratedType `json:"permissions"` // The permissions granted to the user access token.
	Repositories_url string `json:"repositories_url"`
	Repository_selection string `json:"repository_selection"` // Describe whether all repositories have been selected or there's a selection involved
	Single_file_name string `json:"single_file_name"`
	Single_file_paths []string `json:"single_file_paths,omitempty"`
	Account GeneratedType `json:"account"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Projects_v2 GeneratedType `json:"projects_v2"` // A projects v2 project
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Workflow Webhooksworkflow `json:"workflow"`
	Workflow_run map[string]interface{} `json:"workflow_run"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Projects_v2_item GeneratedType `json:"projects_v2_item"` // An item belonging to a project
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Target_ref string `json:"target_ref,omitempty"` // The Git reference of target branch for the commit. For more information, see "[Git References](https://git-scm.com/book/en/v2/Git-Internals-Git-References)" in the Git documentation.
	Sha string `json:"sha,omitempty"` // SHA of commit with autofix.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes,omitempty"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// Webhooksprojectcard represents the Webhooksprojectcard schema from the OpenAPI specification
type Webhooksprojectcard struct {
	Creator map[string]interface{} `json:"creator"`
	Node_id string `json:"node_id"`
	Column_id int `json:"column_id"`
	Id int `json:"id"` // The project card's ID
	Note string `json:"note"`
	Project_url string `json:"project_url"`
	Url string `json:"url"`
	After_id int `json:"after_id,omitempty"`
	Archived bool `json:"archived"` // Whether or not the card is archived
	Updated_at string `json:"updated_at"`
	Column_url string `json:"column_url"`
	Content_url string `json:"content_url,omitempty"`
	Created_at string `json:"created_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Platform_chat string `json:"platform_chat,omitempty"` // The organization policy for allowing or disallowing Copilot features on GitHub.com.
	Public_code_suggestions string `json:"public_code_suggestions"` // The organization policy for allowing or blocking suggestions matching public code (duplication detection filter).
	Seat_breakdown GeneratedType `json:"seat_breakdown"` // The breakdown of Copilot Business seats for the organization.
	Seat_management_setting string `json:"seat_management_setting"` // The mode of assigning new seats.
	Cli string `json:"cli,omitempty"` // The organization policy for allowing or disallowing Copilot in the CLI.
	Ide_chat string `json:"ide_chat,omitempty"` // The organization policy for allowing or disallowing Copilot Chat in the IDE.
	Plan_type string `json:"plan_type,omitempty"` // The Copilot plan of the organization, or the parent enterprise, when applicable.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Raw_url string `json:"raw_url"`
	Status string `json:"status"`
	Additions int `json:"additions"`
	Blob_url string `json:"blob_url"`
	Contents_url string `json:"contents_url"`
	Deletions int `json:"deletions"`
	Filename string `json:"filename"`
	Patch string `json:"patch,omitempty"`
	Sha string `json:"sha"`
	Changes int `json:"changes"`
	Previous_filename string `json:"previous_filename,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Discussion_body_url string `json:"discussion_body_url"` // The URL to the discussion where the secret was detected.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Owner []int `json:"owner"`
	All []int `json:"all"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id int `json:"id,omitempty"`
	Pattern string `json:"pattern"`
	Updated_at string `json:"updated_at,omitempty"`
	Created_at string `json:"created_at,omitempty"`
	Enabled bool `json:"enabled,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url"`
	Node_id string `json:"node_id"`
	Object map[string]interface{} `json:"object"`
	Ref string `json:"ref"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Workflow_job_runs []map[string]interface{} `json:"workflow_job_runs,omitempty"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Reviewers []map[string]interface{} `json:"reviewers,omitempty"`
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Approver Webhooksapprover `json:"approver,omitempty"`
	Since string `json:"since"`
	Workflow_job_run Webhooksworkflowjobrun `json:"workflow_job_run,omitempty"`
	Workflow_run map[string]interface{} `json:"workflow_run"`
	Comment string `json:"comment,omitempty"`
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
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

// Artifact represents the Artifact schema from the OpenAPI specification
type Artifact struct {
	Expired bool `json:"expired"` // Whether or not the artifact has expired.
	Expires_at string `json:"expires_at"`
	Node_id string `json:"node_id"`
	Size_in_bytes int `json:"size_in_bytes"` // The size in bytes of the artifact.
	Workflow_run map[string]interface{} `json:"workflow_run,omitempty"`
	Created_at string `json:"created_at"`
	Name string `json:"name"` // The name of the artifact.
	Updated_at string `json:"updated_at"`
	Archive_download_url string `json:"archive_download_url"`
	Digest string `json:"digest,omitempty"` // The SHA256 digest of the artifact. This field will only be populated on artifacts uploaded with upload-artifact v4 or newer. For older versions, this field will be null.
	Id int `json:"id"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Platform string `json:"platform,omitempty"`
	Name string `json:"name,omitempty"`
	Readme string `json:"readme,omitempty"`
	Commit_oid string `json:"commit_oid,omitempty"`
	Dependencies []map[string]interface{} `json:"dependencies,omitempty"`
	Description string `json:"description,omitempty"`
	Homepage string `json:"homepage,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Repo string `json:"repo,omitempty"`
	Version_info map[string]interface{} `json:"version_info,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pages []map[string]interface{} `json:"pages"` // The pages that were updated.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Reason string `json:"reason"`
	Action string `json:"action"`
	Number int `json:"number"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request map[string]interface{} `json:"pull_request"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Actions_caches []map[string]interface{} `json:"actions_caches"` // Array of caches
	Total_count int `json:"total_count"` // Total number of caches
}

// Team represents the Team schema from the OpenAPI specification
type Team struct {
	Repositories_url string `json:"repositories_url"`
	Description string `json:"description"`
	Html_url string `json:"html_url"`
	Name string `json:"name"`
	Node_id string `json:"node_id"`
	Parent GeneratedType `json:"parent"` // Groups of organization members that gives permissions on specified repositories.
	Url string `json:"url"`
	Notification_setting string `json:"notification_setting,omitempty"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Privacy string `json:"privacy,omitempty"`
	Slug string `json:"slug"`
	Id int `json:"id"`
	Members_url string `json:"members_url"`
	Permission string `json:"permission"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Label Webhookslabel `json:"label,omitempty"`
	Number int `json:"number"` // The pull request number.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Blob_sha string `json:"blob_sha"` // SHA-1 hash ID of the associated blob
	Blob_url string `json:"blob_url"` // The API URL to get the associated blob resource
	Commit_sha string `json:"commit_sha"` // SHA-1 hash ID of the associated commit
	Start_line float64 `json:"start_line"` // Line number at which the secret starts in the file
	End_line float64 `json:"end_line"` // Line number at which the secret ends in the file
	Start_column float64 `json:"start_column"` // The column at which the secret starts within the start line when the file is interpreted as 8BIT ASCII
	Commit_url string `json:"commit_url"` // The API URL to get the associated commit resource
	End_column float64 `json:"end_column"` // The column at which the secret ends within the end line when the file is interpreted as 8BIT ASCII
	Path string `json:"path"` // The file path in the repository
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository_id int `json:"repository_id"` // The ID of the repository where the workflow is defined
	Sha string `json:"sha,omitempty"` // The commit SHA of the workflow file to use
	Path string `json:"path"` // The path to the workflow file
	Ref string `json:"ref,omitempty"` // The ref (branch or tag) of the workflow file to use
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Source string `json:"source"`
	Name string `json:"name"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Rate_limited_request_count int64 `json:"rate_limited_request_count,omitempty"` // The total number of requests that were rate limited within the queried time period
	Total_request_count int64 `json:"total_request_count,omitempty"` // The total number of requests within the queried time period
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Push_protection_bypassed bool `json:"push_protection_bypassed,omitempty"` // Whether push protection was bypassed for the detected secret.
	Repository GeneratedType `json:"repository,omitempty"` // A GitHub repository.
	Resolved_by GeneratedType `json:"resolved_by,omitempty"` // A GitHub user.
	Html_url string `json:"html_url,omitempty"` // The GitHub URL of the alert resource.
	Push_protection_bypassed_by GeneratedType `json:"push_protection_bypassed_by,omitempty"` // A GitHub user.
	Secret_type string `json:"secret_type,omitempty"` // The type of secret that secret scanning detected.
	Has_more_locations bool `json:"has_more_locations,omitempty"` // A boolean value representing whether or not the token in the alert was detected in more than one location.
	Locations_url string `json:"locations_url,omitempty"` // The REST API URL of the code locations for this alert.
	Validity string `json:"validity,omitempty"` // The token status as of the latest validity check.
	Created_at string `json:"created_at,omitempty"` // The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	First_location_detected interface{} `json:"first_location_detected,omitempty"` // Details on the location where the token was initially detected. This can be a commit, wiki commit, issue, discussion, pull request.
	Secret_type_display_name string `json:"secret_type_display_name,omitempty"` // User-friendly name for the detected secret, matching the `secret_type`. For a list of built-in patterns, see "[Supported secret scanning patterns](https://docs.github.com/code-security/secret-scanning/introduction/supported-secret-scanning-patterns#supported-secrets)."
	Number int `json:"number,omitempty"` // The security alert number.
	Push_protection_bypass_request_html_url string `json:"push_protection_bypass_request_html_url,omitempty"` // The URL to a push protection bypass request.
	Updated_at string `json:"updated_at,omitempty"` // The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Publicly_leaked bool `json:"publicly_leaked,omitempty"` // Whether the secret was publicly leaked.
	Resolution string `json:"resolution,omitempty"` // **Required when the `state` is `resolved`.** The reason for resolving the alert.
	Push_protection_bypass_request_reviewer_comment string `json:"push_protection_bypass_request_reviewer_comment,omitempty"` // An optional comment when reviewing a push protection bypass.
	Is_base64_encoded bool `json:"is_base64_encoded,omitempty"` // A boolean value representing whether or not alert is base64 encoded
	Push_protection_bypass_request_reviewer GeneratedType `json:"push_protection_bypass_request_reviewer,omitempty"` // A GitHub user.
	State string `json:"state,omitempty"` // Sets the state of the secret scanning alert. You must provide `resolution` when you set the state to `resolved`.
	Multi_repo bool `json:"multi_repo,omitempty"` // Whether the detected secret was found in multiple repositories in the same organization or enterprise.
	Push_protection_bypass_request_comment string `json:"push_protection_bypass_request_comment,omitempty"` // An optional comment when requesting a push protection bypass.
	Resolved_at string `json:"resolved_at,omitempty"` // The time that the alert was resolved in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Secret string `json:"secret,omitempty"` // The secret that was detected.
	Push_protection_bypassed_at string `json:"push_protection_bypassed_at,omitempty"` // The time that push protection was bypassed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Resolution_comment string `json:"resolution_comment,omitempty"` // The comment that was optionally added when this alert was closed
	Url string `json:"url,omitempty"` // The REST API URL of the alert resource.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Public_repos int `json:"public_repos"`
	Events_url string `json:"events_url"`
	Gravatar_id string `json:"gravatar_id"`
	Twitter_username string `json:"twitter_username,omitempty"`
	TypeField string `json:"type"`
	Repos_url string `json:"repos_url"`
	Hireable bool `json:"hireable"`
	Public_gists int `json:"public_gists"`
	Html_url string `json:"html_url"`
	Collaborators int `json:"collaborators"`
	Two_factor_authentication bool `json:"two_factor_authentication"`
	Followers int `json:"followers"`
	Blog string `json:"blog"`
	Following_url string `json:"following_url"`
	Id int64 `json:"id"`
	Login string `json:"login"`
	Notification_email string `json:"notification_email,omitempty"`
	Location string `json:"location"`
	Email string `json:"email"`
	Starred_url string `json:"starred_url"`
	Subscriptions_url string `json:"subscriptions_url"`
	Updated_at string `json:"updated_at"`
	Avatar_url string `json:"avatar_url"`
	Organizations_url string `json:"organizations_url"`
	Url string `json:"url"`
	Node_id string `json:"node_id"`
	Site_admin bool `json:"site_admin"`
	Company string `json:"company"`
	Private_gists int `json:"private_gists"`
	Received_events_url string `json:"received_events_url"`
	Name string `json:"name"`
	Gists_url string `json:"gists_url"`
	Created_at string `json:"created_at"`
	Business_plus bool `json:"business_plus,omitempty"`
	Following int `json:"following"`
	Owned_private_repos int `json:"owned_private_repos"`
	Plan map[string]interface{} `json:"plan,omitempty"`
	Total_private_repos int `json:"total_private_repos"`
	Bio string `json:"bio"`
	Followers_url string `json:"followers_url"`
	Ldap_dn string `json:"ldap_dn,omitempty"`
	Disk_usage int `json:"disk_usage"`
	User_view_type string `json:"user_view_type,omitempty"`
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
	Sha string `json:"sha"`
	Tree []map[string]interface{} `json:"tree"` // Objects specifying a tree structure
	Truncated bool `json:"truncated"`
	Url string `json:"url,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Label Webhookslabel `json:"label"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Project_card Webhooksprojectcard `json:"project_card"`
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url"`
	Description string `json:"description,omitempty"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	Group_id string `json:"group_id,omitempty"`
	Sync_to_organizations string `json:"sync_to_organizations,omitempty"`
	Group_name string `json:"group_name,omitempty"`
	Html_url string `json:"html_url"`
	Members_url string `json:"members_url"`
	Organization_selection_type string `json:"organization_selection_type,omitempty"`
	Updated_at string `json:"updated_at"`
	Created_at string `json:"created_at"`
	Id int64 `json:"id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Errors []map[string]interface{} `json:"errors,omitempty"`
	Message string `json:"message"`
	Documentation_url string `json:"documentation_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Expire_at string `json:"expire_at,omitempty"` // The time that the bypass will expire in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Reason string `json:"reason,omitempty"` // The reason for bypassing push protection.
	Token_type string `json:"token_type,omitempty"` // The token type this bypass is for.
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
	Sha string `json:"sha"` // SHA for the commit
	Url string `json:"url"`
	Author map[string]interface{} `json:"author"` // Identifying information for the git-user
	Committer map[string]interface{} `json:"committer"` // Identifying information for the git-user
	Html_url string `json:"html_url"`
	Tree map[string]interface{} `json:"tree"`
	Verification map[string]interface{} `json:"verification"`
	Message string `json:"message"` // Message describing the purpose of the commit
	Node_id string `json:"node_id"`
	Parents []map[string]interface{} `json:"parents"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Issue_body_url string `json:"issue_body_url"` // The API URL to get the issue where the secret was detected.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Release Webhooksrelease1 `json:"release"` // The [release](https://docs.github.com/rest/releases/releases/#get-a-release) object.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	TypeField string `json:"type"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Project Webhooksproject `json:"project"`
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes,omitempty"` // The changes to the project if the action was `edited`.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Release map[string]interface{} `json:"release"` // The [release](https://docs.github.com/rest/releases/releases/#get-a-release) object.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Sponsorship Webhookssponsorship `json:"sponsorship"`
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// Key represents the Key schema from the OpenAPI specification
type Key struct {
	Title string `json:"title"`
	Url string `json:"url"`
	Verified bool `json:"verified"`
	Created_at string `json:"created_at"`
	Id int64 `json:"id"`
	Key string `json:"key"`
	Read_only bool `json:"read_only"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Status string `json:"status"`
	Url string `json:"url"`
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

// Language represents the Language schema from the OpenAPI specification
type Language struct {
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url"` // The API URL for the advisory.
	Identifiers []map[string]interface{} `json:"identifiers"`
	Source_code_location string `json:"source_code_location"` // The URL of the advisory's source code.
	Published_at string `json:"published_at"` // The date and time of when the advisory was published, in ISO 8601 format.
	Html_url string `json:"html_url"` // The URL for the advisory.
	Vulnerabilities []Vulnerability `json:"vulnerabilities"` // The products and respective version ranges affected by the advisory.
	Cvss_severities GeneratedType `json:"cvss_severities,omitempty"`
	Summary string `json:"summary"` // A short summary of the advisory.
	Cve_id string `json:"cve_id"` // The Common Vulnerabilities and Exposures (CVE) ID.
	Ghsa_id string `json:"ghsa_id"` // The GitHub Security Advisory ID.
	References []string `json:"references"`
	Repository_advisory_url string `json:"repository_advisory_url"` // The API URL for the repository advisory.
	Credits []map[string]interface{} `json:"credits"` // The users who contributed to the advisory.
	Nvd_published_at string `json:"nvd_published_at"` // The date and time when the advisory was published in the National Vulnerability Database, in ISO 8601 format. This field is only populated when the advisory is imported from the National Vulnerability Database.
	Cwes []map[string]interface{} `json:"cwes"`
	Epss GeneratedType `json:"epss,omitempty"` // The EPSS scores as calculated by the [Exploit Prediction Scoring System](https://www.first.org/epss).
	TypeField string `json:"type"` // The type of advisory.
	Withdrawn_at string `json:"withdrawn_at"` // The date and time of when the advisory was withdrawn, in ISO 8601 format.
	Cvss map[string]interface{} `json:"cvss"`
	Description string `json:"description"` // A detailed description of what the advisory entails.
	Github_reviewed_at string `json:"github_reviewed_at"` // The date and time of when the advisory was reviewed by GitHub, in ISO 8601 format.
	Updated_at string `json:"updated_at"` // The date and time of when the advisory was last updated, in ISO 8601 format.
	Severity string `json:"severity"` // The severity of the advisory.
}

// Webhooksreviewcomment represents the Webhooksreviewcomment schema from the OpenAPI specification
type Webhooksreviewcomment struct {
	Original_start_line int `json:"original_start_line"` // The first line of the range for a multi-line comment.
	Position int `json:"position"` // The line index in the diff to which the comment applies.
	Side string `json:"side"` // The side of the first line of the range for a multi-line comment.
	Updated_at string `json:"updated_at"`
	Diff_hunk string `json:"diff_hunk"` // The diff of the line that the comment refers to.
	Pull_request_review_id int `json:"pull_request_review_id"` // The ID of the pull request review to which the comment belongs.
	Links map[string]interface{} `json:"_links"`
	Body string `json:"body"` // The text of the comment.
	Original_line int `json:"original_line"` // The line of the blob to which the comment applies. The last line of the range for a multi-line comment
	Pull_request_url string `json:"pull_request_url"` // URL for the pull request that the review comment belongs to.
	Start_side string `json:"start_side"` // The side of the first line of the range for a multi-line comment.
	Reactions map[string]interface{} `json:"reactions"`
	Commit_id string `json:"commit_id"` // The SHA of the commit to which the comment applies.
	Line int `json:"line"` // The line of the blob to which the comment applies. The last line of the range for a multi-line comment
	User map[string]interface{} `json:"user"`
	Html_url string `json:"html_url"` // HTML URL for the pull request review comment.
	Original_commit_id string `json:"original_commit_id"` // The SHA of the original commit to which the comment applies.
	Start_line int `json:"start_line"` // The first line of the range for a multi-line comment.
	Subject_type string `json:"subject_type,omitempty"` // The level at which the comment is targeted, can be a diff line or a file.
	Url string `json:"url"` // URL for the pull request review comment
	Created_at string `json:"created_at"`
	In_reply_to_id int `json:"in_reply_to_id,omitempty"` // The comment ID to reply to.
	Original_position int `json:"original_position"` // The index of the original line in the diff to which the comment applies.
	Id int `json:"id"` // The ID of the pull request review comment.
	Path string `json:"path"` // The relative path of the file to which the comment applies.
	Node_id string `json:"node_id"` // The node ID of the pull request review comment.
	Author_association string `json:"author_association"` // How the author is associated with the repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Total_active_users int `json:"total_active_users,omitempty"` // The total number of Copilot users with activity belonging to any Copilot feature, globally, for the given day. Includes passive activity such as receiving a code suggestion, as well as engagement activity such as accepting a code suggestion or prompting chat. Does not include authentication events. Is not limited to the individual features detailed on the endpoint.
	Total_engaged_users int `json:"total_engaged_users,omitempty"` // The total number of Copilot users who engaged with any Copilot feature, for the given day. Examples include but are not limited to accepting a code suggestion, prompting Copilot chat, or triggering a PR Summary. Does not include authentication events. Is not limited to the individual features detailed on the endpoint.
	Copilot_dotcom_chat GeneratedType `json:"copilot_dotcom_chat,omitempty"` // Usage metrics for Copilot Chat in GitHub.com
	Copilot_dotcom_pull_requests GeneratedType `json:"copilot_dotcom_pull_requests,omitempty"` // Usage metrics for Copilot for pull requests.
	Copilot_ide_chat GeneratedType `json:"copilot_ide_chat,omitempty"` // Usage metrics for Copilot Chat in the IDE.
	Copilot_ide_code_completions GeneratedType `json:"copilot_ide_code_completions,omitempty"` // Usage metrics for Copilot editor code completions in the IDE.
	Date string `json:"date"` // The date for which the usage metrics are aggregated, in `YYYY-MM-DD` format.
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

// Environment represents the Environment schema from the OpenAPI specification
type Environment struct {
	Id int64 `json:"id"` // The id of the environment.
	Updated_at string `json:"updated_at"` // The time that the environment was last updated, in ISO 8601 format.
	Protection_rules []interface{} `json:"protection_rules,omitempty"` // Built-in deployment protection rules for the environment.
	Html_url string `json:"html_url"`
	Name string `json:"name"` // The name of the environment.
	Node_id string `json:"node_id"`
	Url string `json:"url"`
	Created_at string `json:"created_at"` // The time that the environment was created, in ISO 8601 format.
	Deployment_branch_policy GeneratedType `json:"deployment_branch_policy,omitempty"` // The type of deployment branch policy for this environment. To allow all branches to deploy, set to `null`.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Last_active_on string `json:"last_active_on,omitempty"` // The time at which the runner was last used, in ISO 8601 format.
	Maximum_runners int `json:"maximum_runners,omitempty"` // The maximum amount of hosted runners. Runners will not scale automatically above this number. Use this setting to limit your cost.
	Name string `json:"name"` // The name of the hosted runner.
	Public_ips []GeneratedType `json:"public_ips,omitempty"` // The public IP ranges when public IP is enabled for the hosted runners.
	Status string `json:"status"` // The status of the runner.
	Machine_size_details GeneratedType `json:"machine_size_details"` // Provides details of a particular machine spec.
	Public_ip_enabled bool `json:"public_ip_enabled"` // Whether public IP is enabled for the hosted runners.
	Id int `json:"id"` // The unique identifier of the hosted runner.
	Image_details GeneratedType `json:"image_details"` // Provides details of a hosted runner image
	Platform string `json:"platform"` // The operating system of the image.
	Runner_group_id int `json:"runner_group_id,omitempty"` // The unique identifier of the group that the hosted runner belongs to.
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
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Check_suite map[string]interface{} `json:"check_suite"` // The [check_suite](https://docs.github.com/rest/checks/suites#get-a-check-suite).
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Commit_oid string `json:"commit_oid"` // The commit SHA of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Ref string `json:"ref"` // The Git reference of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Alert map[string]interface{} `json:"alert"` // The code scanning alert involved in the event.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
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
	Name string `json:"name"` // The name of the private registry configuration.
	Registry_type string `json:"registry_type"` // The registry type.
	Updated_at string `json:"updated_at"`
	Username string `json:"username,omitempty"` // The username to use when authenticating with the private registry.
	Visibility string `json:"visibility"` // Which type of organization repositories have access to the private registry.
	Created_at string `json:"created_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Cve_id string `json:"cve_id"` // The Common Vulnerabilities and Exposures (CVE) ID.
	Cvss_severities GeneratedType `json:"cvss_severities,omitempty"`
	Updated_at string `json:"updated_at"` // The date and time of when the advisory was last updated, in ISO 8601 format.
	Closed_at string `json:"closed_at"` // The date and time of when the advisory was closed, in ISO 8601 format.
	State string `json:"state"` // The state of the advisory.
	Created_at string `json:"created_at"` // The date and time of when the advisory was created, in ISO 8601 format.
	Credits_detailed []GeneratedType `json:"credits_detailed"`
	Identifiers []map[string]interface{} `json:"identifiers"`
	Cwe_ids []string `json:"cwe_ids"` // A list of only the CWE IDs.
	Author interface{} `json:"author"` // The author of the advisory.
	Cwes []map[string]interface{} `json:"cwes"`
	Collaborating_users []GeneratedType `json:"collaborating_users"` // A list of users that collaborate on the advisory.
	Credits []map[string]interface{} `json:"credits"`
	Submission map[string]interface{} `json:"submission"`
	Summary string `json:"summary"` // A short summary of the advisory.
	Vulnerabilities []GeneratedType `json:"vulnerabilities"`
	Withdrawn_at string `json:"withdrawn_at"` // The date and time of when the advisory was withdrawn, in ISO 8601 format.
	Url string `json:"url"` // The API URL for the advisory.
	Cvss map[string]interface{} `json:"cvss"`
	Description string `json:"description"` // A detailed description of what the advisory entails.
	Html_url string `json:"html_url"` // The URL for the advisory.
	Published_at string `json:"published_at"` // The date and time of when the advisory was published, in ISO 8601 format.
	Ghsa_id string `json:"ghsa_id"` // The GitHub Security Advisory ID.
	Severity string `json:"severity"` // The severity of the advisory.
	Collaborating_teams []Team `json:"collaborating_teams"` // A list of teams that collaborate on the advisory.
	Private_fork interface{} `json:"private_fork"` // A temporary private fork of the advisory's repository for collaborating on a fix.
	Publisher interface{} `json:"publisher"` // The publisher of the advisory.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Pull_request_comment_url string `json:"pull_request_comment_url"` // The API URL to get the pull request comment where the secret was detected.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Ecosystem string `json:"ecosystem"` // The package's language or package management ecosystem.
	Name string `json:"name"` // The unique package name within its ecosystem.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Target string `json:"target"`
	Url string `json:"url"`
	Links map[string]interface{} `json:"_links"`
	Download_url string `json:"download_url"`
	Html_url string `json:"html_url"`
	Name string `json:"name"`
	TypeField string `json:"type"`
	Git_url string `json:"git_url"`
	Path string `json:"path"`
	Sha string `json:"sha"`
	Size int `json:"size"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Alt_domain map[string]interface{} `json:"alt_domain,omitempty"`
	Domain map[string]interface{} `json:"domain,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	New_property_values []GeneratedType `json:"new_property_values"` // The new custom property values for the repository.
	Old_property_values []GeneratedType `json:"old_property_values"` // The old custom property values for the repository.
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Changes map[string]interface{} `json:"changes"`
	Installation GeneratedType `json:"installation"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Account map[string]interface{} `json:"account"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Target_type string `json:"target_type"`
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Release Webhooksrelease1 `json:"release"` // The [release](https://docs.github.com/rest/releases/releases/#get-a-release) object.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// Webhooksmembership represents the Webhooksmembership schema from the OpenAPI specification
type Webhooksmembership struct {
	User map[string]interface{} `json:"user"`
	Organization_url string `json:"organization_url"`
	Role string `json:"role"`
	State string `json:"state"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Node_id string `json:"node_id"`
	Project_node_id string `json:"project_node_id,omitempty"`
	Status string `json:"status,omitempty"`
	Body string `json:"body,omitempty"` // Body of the status update
	Creator GeneratedType `json:"creator,omitempty"` // A GitHub user.
	Id float64 `json:"id"`
	Start_date string `json:"start_date,omitempty"`
	Target_date string `json:"target_date,omitempty"`
	Updated_at string `json:"updated_at"`
	Created_at string `json:"created_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository"` // Minimal Repository
	Preferences map[string]interface{} `json:"preferences"`
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
	Created_at string `json:"created_at"`
	Name string `json:"name"` // The name of the secret.
	Selected_repositories_url string `json:"selected_repositories_url,omitempty"`
	Updated_at string `json:"updated_at"`
	Visibility string `json:"visibility"` // Visibility of a secret
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Commit map[string]interface{} `json:"commit"`
	Content map[string]interface{} `json:"content"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Projects_v2_status_update GeneratedType `json:"projects_v2_status_update"` // An status update belonging to a project
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Discussion_comment_url string `json:"discussion_comment_url"` // The API URL to get the discussion comment where the secret was detected.
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
	Throttled_at string `json:"throttled_at,omitempty"` // Time when the webhook delivery was throttled.
	Event string `json:"event"` // The event that triggered the delivery.
	Id int64 `json:"id"` // Unique identifier of the webhook delivery.
	Repository_id int64 `json:"repository_id"` // The id of the repository associated with this event.
	Guid string `json:"guid"` // Unique identifier for the event (shared with all deliveries for all webhooks that subscribe to this event).
	Redelivery bool `json:"redelivery"` // Whether the webhook delivery is a redelivery.
	Duration float64 `json:"duration"` // Time spent delivering.
	Status_code int `json:"status_code"` // Status code received when delivery was made.
	Action string `json:"action"` // The type of activity for the event that triggered the delivery.
	Delivered_at string `json:"delivered_at"` // Time when the webhook delivery occurred.
	Installation_id int64 `json:"installation_id"` // The id of the GitHub App installation associated with this event.
	Status string `json:"status"` // Describes the response returned after attempting the delivery.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Description string `json:"description,omitempty"` // Short description of the property
	Allowed_values []string `json:"allowed_values,omitempty"` // An ordered list of the allowed values of the property. The property can have up to 200 allowed values.
	Default_value string `json:"default_value,omitempty"` // Default value of the property
	Property_name string `json:"property_name"` // The name of the property
	Source_type string `json:"source_type,omitempty"` // The source type of the property
	Required bool `json:"required,omitempty"` // Whether the property is required.
	Url string `json:"url,omitempty"` // The URL that can be used to fetch, update, or delete info about this property via the API.
	Values_editable_by string `json:"values_editable_by,omitempty"` // Who can edit the values of the property
	Value_type string `json:"value_type"` // The type of the value for the property
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	First_location_detected interface{} `json:"first_location_detected,omitempty"` // Details on the location where the token was initially detected. This can be a commit, wiki commit, issue, discussion, pull request.
	Push_protection_bypass_request_comment string `json:"push_protection_bypass_request_comment,omitempty"` // An optional comment when requesting a push protection bypass.
	Push_protection_bypassed_at string `json:"push_protection_bypassed_at,omitempty"` // The time that push protection was bypassed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Created_at string `json:"created_at,omitempty"` // The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Resolved_by GeneratedType `json:"resolved_by,omitempty"` // A GitHub user.
	Secret string `json:"secret,omitempty"` // The secret that was detected.
	Push_protection_bypass_request_reviewer GeneratedType `json:"push_protection_bypass_request_reviewer,omitempty"` // A GitHub user.
	Locations_url string `json:"locations_url,omitempty"` // The REST API URL of the code locations for this alert.
	Publicly_leaked bool `json:"publicly_leaked,omitempty"` // Whether the detected secret was publicly leaked.
	Resolved_at string `json:"resolved_at,omitempty"` // The time that the alert was resolved in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Multi_repo bool `json:"multi_repo,omitempty"` // Whether the detected secret was found in multiple repositories under the same organization or enterprise.
	Push_protection_bypass_request_reviewer_comment string `json:"push_protection_bypass_request_reviewer_comment,omitempty"` // An optional comment when reviewing a push protection bypass.
	Resolution string `json:"resolution,omitempty"` // **Required when the `state` is `resolved`.** The reason for resolving the alert.
	Is_base64_encoded bool `json:"is_base64_encoded,omitempty"` // A boolean value representing whether or not alert is base64 encoded
	Number int `json:"number,omitempty"` // The security alert number.
	Push_protection_bypass_request_html_url string `json:"push_protection_bypass_request_html_url,omitempty"` // The URL to a push protection bypass request.
	Has_more_locations bool `json:"has_more_locations,omitempty"` // A boolean value representing whether or not the token in the alert was detected in more than one location.
	Resolution_comment string `json:"resolution_comment,omitempty"` // An optional comment to resolve an alert.
	Secret_type_display_name string `json:"secret_type_display_name,omitempty"` // User-friendly name for the detected secret, matching the `secret_type`. For a list of built-in patterns, see "[Supported secret scanning patterns](https://docs.github.com/code-security/secret-scanning/introduction/supported-secret-scanning-patterns#supported-secrets)."
	Url string `json:"url,omitempty"` // The REST API URL of the alert resource.
	State string `json:"state,omitempty"` // Sets the state of the secret scanning alert. You must provide `resolution` when you set the state to `resolved`.
	Validity string `json:"validity,omitempty"` // The token status as of the latest validity check.
	Secret_type string `json:"secret_type,omitempty"` // The type of secret that secret scanning detected.
	Html_url string `json:"html_url,omitempty"` // The GitHub URL of the alert resource.
	Push_protection_bypassed_by GeneratedType `json:"push_protection_bypassed_by,omitempty"` // A GitHub user.
	Updated_at string `json:"updated_at,omitempty"` // The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Push_protection_bypassed bool `json:"push_protection_bypassed,omitempty"` // Whether push protection was bypassed for the detected secret.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Privacy string `json:"privacy,omitempty"` // The level of privacy this team should have
	Repositories_url string `json:"repositories_url"`
	Description string `json:"description"` // Description of the team
	Html_url string `json:"html_url"`
	Url string `json:"url"` // URL for the team
	Members_url string `json:"members_url"`
	Notification_setting string `json:"notification_setting,omitempty"` // The notification setting the team has set
	Slug string `json:"slug"`
	Ldap_dn string `json:"ldap_dn,omitempty"` // Distinguished Name (DN) that team maps to within LDAP environment
	Name string `json:"name"` // Name of the team
	Id int `json:"id"` // Unique identifier of the team
	Node_id string `json:"node_id"`
	Permission string `json:"permission"` // Permission that the team will have for its repositories
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Alert map[string]interface{} `json:"alert"` // The security alert of the vulnerable dependency.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Workflow_run map[string]interface{} `json:"workflow_run,omitempty"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Deployment_status map[string]interface{} `json:"deployment_status"` // The [deployment status](https://docs.github.com/rest/deployments/statuses#list-deployment-statuses).
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Workflow Webhooksworkflow `json:"workflow,omitempty"`
	Action string `json:"action"`
	Check_run map[string]interface{} `json:"check_run,omitempty"`
	Deployment map[string]interface{} `json:"deployment"` // The [deployment](https://docs.github.com/rest/deployments/deployments#list-deployments).
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Payload string `json:"payload"` // A URL-encoded string of the secret_scanning_alert_location.created JSON payload. The decoded payload is a JSON object.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Alert map[string]interface{} `json:"alert"` // The security alert of the vulnerable dependency.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// Activity represents the Activity schema from the OpenAPI specification
type Activity struct {
	Node_id string `json:"node_id"`
	Ref string `json:"ref"` // The full Git reference, formatted as `refs/heads/<branch name>`.
	Timestamp string `json:"timestamp"` // The time when the activity occurred.
	Activity_type string `json:"activity_type"` // The type of the activity that was performed.
	Actor GeneratedType `json:"actor"` // A GitHub user.
	After string `json:"after"` // The SHA of the commit after the activity.
	Before string `json:"before"` // The SHA of the commit before the activity.
	Id int `json:"id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Thread map[string]interface{} `json:"thread"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Changes map[string]interface{} `json:"changes"` // The changes to the issue.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Label Webhookslabel `json:"label,omitempty"`
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
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
	Author map[string]interface{} `json:"author"` // Information about the Git author
	Committer map[string]interface{} `json:"committer"` // Information about the Git committer
	Id string `json:"id"` // SHA for the commit
	Message string `json:"message"` // Message describing the purpose of the commit
	Timestamp string `json:"timestamp"` // Timestamp of the commit
	Tree_id string `json:"tree_id"` // SHA for the commit's tree
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Email string `json:"email,omitempty"`
	Two_factor_requirement_enabled bool `json:"two_factor_requirement_enabled,omitempty"`
	Blog string `json:"blog,omitempty"`
	Id int `json:"id"`
	Total_private_repos int `json:"total_private_repos,omitempty"`
	Has_organization_projects bool `json:"has_organization_projects"`
	Url string `json:"url"`
	Members_can_create_public_repositories bool `json:"members_can_create_public_repositories,omitempty"`
	Members_can_create_private_pages bool `json:"members_can_create_private_pages,omitempty"`
	Public_members_url string `json:"public_members_url"`
	Name string `json:"name,omitempty"`
	Followers int `json:"followers"`
	Repos_url string `json:"repos_url"`
	Public_gists int `json:"public_gists"`
	TypeField string `json:"type"`
	Company string `json:"company,omitempty"`
	Plan map[string]interface{} `json:"plan,omitempty"`
	Issues_url string `json:"issues_url"`
	Avatar_url string `json:"avatar_url"`
	Members_can_create_repositories bool `json:"members_can_create_repositories,omitempty"`
	Billing_email string `json:"billing_email,omitempty"`
	Hooks_url string `json:"hooks_url"`
	Members_can_fork_private_repositories bool `json:"members_can_fork_private_repositories,omitempty"`
	Members_url string `json:"members_url"`
	Html_url string `json:"html_url"`
	Following int `json:"following"`
	Description string `json:"description"`
	Members_can_create_public_pages bool `json:"members_can_create_public_pages,omitempty"`
	Private_gists int `json:"private_gists,omitempty"`
	Default_repository_permission string `json:"default_repository_permission,omitempty"`
	Is_verified bool `json:"is_verified,omitempty"`
	Members_can_create_internal_repositories bool `json:"members_can_create_internal_repositories,omitempty"`
	Created_at string `json:"created_at"`
	Collaborators int `json:"collaborators,omitempty"`
	Members_can_create_pages bool `json:"members_can_create_pages,omitempty"`
	Node_id string `json:"node_id"`
	Archived_at string `json:"archived_at"`
	Disk_usage int `json:"disk_usage,omitempty"`
	Login string `json:"login"`
	Twitter_username string `json:"twitter_username,omitempty"`
	Location string `json:"location,omitempty"`
	Members_can_create_private_repositories bool `json:"members_can_create_private_repositories,omitempty"`
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"`
	Has_repository_projects bool `json:"has_repository_projects"`
	Members_allowed_repository_creation_type string `json:"members_allowed_repository_creation_type,omitempty"`
	Events_url string `json:"events_url"`
	Updated_at string `json:"updated_at"`
	Owned_private_repos int `json:"owned_private_repos,omitempty"`
	Public_repos int `json:"public_repos"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Label Webhookslabel `json:"label,omitempty"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	PackageField map[string]interface{} `json:"package"` // Information about the package.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Student_repository_url string `json:"student_repository_url"` // URL of the student's assignment repository
	Github_username string `json:"github_username"` // GitHub username of the student
	Points_awarded int `json:"points_awarded"` // Number of points awarded to the student
	Starter_code_url string `json:"starter_code_url"` // URL of the starter code for the assignment
	Student_repository_name string `json:"student_repository_name"` // Name of the student's assignment repository
	Assignment_name string `json:"assignment_name"` // Name of the assignment
	Group_name string `json:"group_name,omitempty"` // If a group assignment, name of the group the student is in
	Assignment_url string `json:"assignment_url"` // URL of the assignment
	Submission_timestamp string `json:"submission_timestamp"` // Timestamp of the student's assignment submission
	Points_available int `json:"points_available"` // Number of points available for the assignment
	Roster_identifier string `json:"roster_identifier"` // Roster identifier of the student
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

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Referrer string `json:"referrer"`
	Uniques int `json:"uniques"`
	Count int `json:"count"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"` // The `event_type` that was specified in the `POST /repos/{owner}/{repo}/dispatches` request body.
	Branch string `json:"branch"`
	Client_payload map[string]interface{} `json:"client_payload"` // The `client_payload` that was specified in the `POST /repos/{owner}/{repo}/dispatches` request body.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Status string `json:"status,omitempty"`
	Documentation_url string `json:"documentation_url,omitempty"`
	Message string `json:"message,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Incremental_scans []GeneratedType `json:"incremental_scans,omitempty"`
	Pattern_update_scans []GeneratedType `json:"pattern_update_scans,omitempty"`
	Backfill_scans []GeneratedType `json:"backfill_scans,omitempty"`
	Custom_pattern_backfill_scans []interface{} `json:"custom_pattern_backfill_scans,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	TypeField string `json:"type"`
}

// Hook represents the Hook schema from the OpenAPI specification
type Hook struct {
	Name string `json:"name"` // The name of a valid service, use 'web' for a webhook.
	Test_url string `json:"test_url"`
	TypeField string `json:"type"`
	Last_response GeneratedType `json:"last_response"`
	Ping_url string `json:"ping_url"`
	Url string `json:"url"`
	Config GeneratedType `json:"config"` // Configuration object of the webhook
	Created_at string `json:"created_at"`
	Deliveries_url string `json:"deliveries_url,omitempty"`
	Events []string `json:"events"` // Determines what events the hook is triggered for. Default: ['push'].
	Id int `json:"id"` // Unique identifier of the webhook.
	Updated_at string `json:"updated_at"`
	Active bool `json:"active"` // Determines whether the hook is actually triggered on pushes.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Billing_cycle string `json:"billing_cycle"`
	Free_trial_ends_on string `json:"free_trial_ends_on"`
	Next_billing_date string `json:"next_billing_date"`
	On_free_trial bool `json:"on_free_trial"`
	Plan GeneratedType `json:"plan"` // Marketplace Listing Plan
	Unit_count int `json:"unit_count"`
	Updated_at string `json:"updated_at"`
	Account GeneratedType `json:"account"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Days_left_in_billing_cycle int `json:"days_left_in_billing_cycle"` // Numbers of days left in billing cycle.
	Estimated_paid_storage_for_month int `json:"estimated_paid_storage_for_month"` // Estimated storage space (GB) used in billing cycle.
	Estimated_storage_for_month int `json:"estimated_storage_for_month"` // Estimated sum of free and paid storage space (GB) used in billing cycle.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Expired bool `json:"expired,omitempty"` // Whether or not the invitation has expired
	Id int64 `json:"id"` // Unique identifier of the repository invitation.
	Repository GeneratedType `json:"repository"` // Minimal Repository
	Url string `json:"url"` // URL for the repository invitation
	Invitee GeneratedType `json:"invitee"` // A GitHub user.
	Html_url string `json:"html_url"`
	Node_id string `json:"node_id"`
	Permissions string `json:"permissions"` // The permission associated with the invitation.
	Created_at string `json:"created_at"`
	Inviter GeneratedType `json:"inviter"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enabled bool `json:"enabled"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Usageitems []map[string]interface{} `json:"usageItems,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url"`
	Node_id string `json:"node_id"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Assignee GeneratedType `json:"assignee"` // A GitHub user.
	Commit_url string `json:"commit_url"`
	Id int `json:"id"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Commit_id string `json:"commit_id"`
	Created_at string `json:"created_at"`
	Event string `json:"event"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Started_at string `json:"started_at"` // The start time of an autofix in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Status string `json:"status"` // The status of an autofix.
	Description string `json:"description"` // The description of an autofix.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	TypeField string `json:"type"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Wait_timer int `json:"wait_timer"` // The set duration of the wait timer
	Wait_timer_started_at string `json:"wait_timer_started_at"` // The time that the wait timer began.
	Current_user_can_approve bool `json:"current_user_can_approve"` // Whether the currently authenticated user can approve the deployment
	Environment map[string]interface{} `json:"environment"`
	Reviewers []map[string]interface{} `json:"reviewers"` // The people or teams that may approve jobs that reference the environment. You can list up to six users or teams as reviewers. The reviewers must have at least read access to the repository. Only one of the required reviewers needs to approve the job for it to proceed.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Project_card map[string]interface{} `json:"project_card"`
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Check_suite map[string]interface{} `json:"check_suite"` // The [check_suite](https://docs.github.com/rest/checks/suites#get-a-check-suite).
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	TypeField string `json:"type"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Pattern string `json:"pattern,omitempty"`
	Protected bool `json:"protected"`
	Protection GeneratedType `json:"protection"` // Branch Protection
	Protection_url string `json:"protection_url"`
	Required_approving_review_count int `json:"required_approving_review_count,omitempty"`
	Links map[string]interface{} `json:"_links"`
	Commit Commit `json:"commit"` // Commit
	Name string `json:"name"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Change_status map[string]interface{} `json:"change_status,omitempty"`
	Committed_at string `json:"committed_at,omitempty"`
	Url string `json:"url,omitempty"`
	User GeneratedType `json:"user,omitempty"` // A GitHub user.
	Version string `json:"version,omitempty"`
}

// Webhooksrule represents the Webhooksrule schema from the OpenAPI specification
type Webhooksrule struct {
	Pull_request_reviews_enforcement_level string `json:"pull_request_reviews_enforcement_level"`
	Signature_requirement_enforcement_level string `json:"signature_requirement_enforcement_level"`
	Require_code_owner_review bool `json:"require_code_owner_review"`
	Required_approving_review_count int `json:"required_approving_review_count"`
	Allow_force_pushes_enforcement_level string `json:"allow_force_pushes_enforcement_level"`
	Authorized_dismissal_actors_only bool `json:"authorized_dismissal_actors_only"`
	Create_protected bool `json:"create_protected,omitempty"`
	Merge_queue_enforcement_level string `json:"merge_queue_enforcement_level"`
	Name string `json:"name"`
	Authorized_actor_names []string `json:"authorized_actor_names"`
	Repository_id int `json:"repository_id"`
	Required_deployments_enforcement_level string `json:"required_deployments_enforcement_level"`
	Authorized_actors_only bool `json:"authorized_actors_only"`
	Dismiss_stale_reviews_on_push bool `json:"dismiss_stale_reviews_on_push"`
	Require_last_push_approval bool `json:"require_last_push_approval,omitempty"` // Whether the most recent push must be approved by someone other than the person who pushed it
	Admin_enforced bool `json:"admin_enforced"`
	Lock_branch_enforcement_level string `json:"lock_branch_enforcement_level"` // The enforcement level of the branch lock setting. `off` means the branch is not locked, `non_admins` means the branch is read-only for non_admins, and `everyone` means the branch is read-only for everyone.
	Required_status_checks []string `json:"required_status_checks"`
	Required_status_checks_enforcement_level string `json:"required_status_checks_enforcement_level"`
	Created_at string `json:"created_at"`
	Linear_history_requirement_enforcement_level string `json:"linear_history_requirement_enforcement_level"`
	Required_conversation_resolution_level string `json:"required_conversation_resolution_level"`
	Allow_deletions_enforcement_level string `json:"allow_deletions_enforcement_level"`
	Strict_required_status_checks_policy bool `json:"strict_required_status_checks_policy"`
	Updated_at string `json:"updated_at"`
	Id int `json:"id"`
	Ignore_approvals_from_contributors bool `json:"ignore_approvals_from_contributors"`
	Lock_allows_fork_sync bool `json:"lock_allows_fork_sync,omitempty"` // Whether users can pull changes from upstream when the branch is locked. Set to `true` to allow users to pull changes from upstream when the branch is locked. This setting is only applicable for forks.
}

// Issue represents the Issue schema from the OpenAPI specification
type Issue struct {
	Updated_at string `json:"updated_at"`
	Body_text string `json:"body_text,omitempty"`
	Events_url string `json:"events_url"`
	Id int64 `json:"id"`
	Reactions GeneratedType `json:"reactions,omitempty"`
	Comments_url string `json:"comments_url"`
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Url string `json:"url"` // URL for the issue
	Draft bool `json:"draft,omitempty"`
	Labels_url string `json:"labels_url"`
	Closed_at string `json:"closed_at"`
	Sub_issues_summary GeneratedType `json:"sub_issues_summary,omitempty"`
	User GeneratedType `json:"user"` // A GitHub user.
	Node_id string `json:"node_id"`
	Labels []interface{} `json:"labels"` // Labels to associate with this issue; pass one or more label names to replace the set of labels on this issue; send an empty array to clear all labels from the issue; note that the labels are silently dropped for users without push access to the repository
	Body string `json:"body,omitempty"` // Contents of the issue
	Comments int `json:"comments"`
	Assignee GeneratedType `json:"assignee"` // A GitHub user.
	State_reason string `json:"state_reason,omitempty"` // The reason for the current state
	TypeField GeneratedType `json:"type,omitempty"` // The type of issue.
	Closed_by GeneratedType `json:"closed_by,omitempty"` // A GitHub user.
	Html_url string `json:"html_url"`
	Pull_request map[string]interface{} `json:"pull_request,omitempty"`
	Assignees []GeneratedType `json:"assignees,omitempty"`
	Milestone GeneratedType `json:"milestone"` // A collection of related issues and pull requests.
	Created_at string `json:"created_at"`
	Locked bool `json:"locked"`
	Active_lock_reason string `json:"active_lock_reason,omitempty"`
	Number int `json:"number"` // Number uniquely identifying the issue within its repository
	Title string `json:"title"` // Title of the issue
	Repository_url string `json:"repository_url"`
	State string `json:"state"` // State of the issue; either 'open' or 'closed'
	Timeline_url string `json:"timeline_url,omitempty"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Body_html string `json:"body_html,omitempty"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Commit_id string `json:"commit_id"`
	Review_requester GeneratedType `json:"review_requester"` // A GitHub user.
	Url string `json:"url"`
	Commit_url string `json:"commit_url"`
	Created_at string `json:"created_at"`
	Id int `json:"id"`
	Requested_reviewer GeneratedType `json:"requested_reviewer,omitempty"` // A GitHub user.
	Requested_team Team `json:"requested_team,omitempty"` // Groups of organization members that gives permissions on specified repositories.
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Event string `json:"event"`
	Node_id string `json:"node_id"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Workflow_run map[string]interface{} `json:"workflow_run"`
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Workflow Webhooksworkflow `json:"workflow"`
}

// Webhooksissue represents the Webhooksissue schema from the OpenAPI specification
type Webhooksissue struct {
	Repository_url string `json:"repository_url"`
	Assignee map[string]interface{} `json:"assignee,omitempty"`
	Locked bool `json:"locked,omitempty"`
	Id int64 `json:"id"`
	Active_lock_reason string `json:"active_lock_reason"`
	Labels []map[string]interface{} `json:"labels,omitempty"`
	Comments int `json:"comments"`
	Html_url string `json:"html_url"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Draft bool `json:"draft,omitempty"`
	Reactions map[string]interface{} `json:"reactions"`
	State_reason string `json:"state_reason,omitempty"`
	Labels_url string `json:"labels_url"`
	Milestone map[string]interface{} `json:"milestone"` // A collection of related issues and pull requests.
	Performed_via_github_app map[string]interface{} `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	TypeField GeneratedType `json:"type,omitempty"` // The type of issue.
	Events_url string `json:"events_url"`
	Node_id string `json:"node_id"`
	Number int `json:"number"`
	Updated_at string `json:"updated_at"`
	Sub_issues_summary map[string]interface{} `json:"sub_issues_summary,omitempty"`
	User map[string]interface{} `json:"user"`
	Pull_request map[string]interface{} `json:"pull_request,omitempty"`
	Title string `json:"title"` // Title of the issue
	Comments_url string `json:"comments_url"`
	Created_at string `json:"created_at"`
	State string `json:"state,omitempty"` // State of the issue; either 'open' or 'closed'
	Closed_at string `json:"closed_at"`
	Body string `json:"body"` // Contents of the issue
	Timeline_url string `json:"timeline_url,omitempty"`
	Url string `json:"url"` // URL for the issue
	Assignees []map[string]interface{} `json:"assignees"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Release Webhooksrelease `json:"release"` // The [release](https://docs.github.com/rest/releases/releases/#get-a-release) object.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Changes map[string]interface{} `json:"changes,omitempty"` // The changes to the label if the action was `edited`.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Label Webhookslabel `json:"label"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Project_card Webhooksprojectcard `json:"project_card"`
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Name string `json:"name"`
	Url string `json:"url"`
	Html_url string `json:"html_url"`
	Key string `json:"key"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Project_column Webhooksprojectcolumn `json:"project_column"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	End_line int `json:"end_line,omitempty"`
	Path string `json:"path,omitempty"`
	Start_column int `json:"start_column,omitempty"`
	Start_line int `json:"start_line,omitempty"`
	End_column int `json:"end_column,omitempty"`
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
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Assigner GeneratedType `json:"assigner"` // A GitHub user.
	Commit_id string `json:"commit_id"`
	Created_at string `json:"created_at"`
	Node_id string `json:"node_id"`
	Url string `json:"url"`
	Event string `json:"event"`
	Id int `json:"id"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Assignee GeneratedType `json:"assignee"` // A GitHub user.
	Commit_url string `json:"commit_url"`
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
	Version string `json:"version"`
	Change_status map[string]interface{} `json:"change_status"`
	Committed_at string `json:"committed_at"`
	Url string `json:"url"`
	User GeneratedType `json:"user"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Uniques int `json:"uniques"`
	Clones []Traffic `json:"clones"`
	Count int `json:"count"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Custom_branch_policies bool `json:"custom_branch_policies"` // Whether only branches that match the specified name patterns can deploy to this environment. If `custom_branch_policies` is `true`, `protected_branches` must be `false`; if `custom_branch_policies` is `false`, `protected_branches` must be `true`.
	Protected_branches bool `json:"protected_branches"` // Whether only branches with branch protection rules can deploy to this environment. If `protected_branches` is `true`, `custom_branch_policies` must be `false`; if `protected_branches` is `false`, `custom_branch_policies` must be `true`.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Html_url string `json:"html_url"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Email string `json:"email,omitempty"`
	Id int `json:"id"`
	Login string `json:"login"`
	Node_id string `json:"node_id,omitempty"`
	Organization_billing_email string `json:"organization_billing_email,omitempty"`
	TypeField string `json:"type"`
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
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Comment Webhooksissuecomment `json:"comment"` // The [comment](https://docs.github.com/rest/issues/comments#get-an-issue-comment) itself.
	Issue interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) the comment belongs to.
	Changes Webhookschanges `json:"changes"` // The changes to the comment.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Role_name string `json:"role_name"`
	User GeneratedType `json:"user"` // Collaborator
	Permission string `json:"permission"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue Webhooksissue `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"` // The action that was performed.
	Assignee Webhooksuser `json:"assignee,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Ref string `json:"ref"`
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Workflow string `json:"workflow"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Inputs map[string]interface{} `json:"inputs"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Html_url string `json:"html_url,omitempty"`
	Node_id string `json:"node_id,omitempty"`
	Truncated bool `json:"truncated,omitempty"`
	Files map[string]interface{} `json:"files,omitempty"`
	Created_at string `json:"created_at,omitempty"`
	Url string `json:"url,omitempty"`
	Comments int `json:"comments,omitempty"`
	Forks_url string `json:"forks_url,omitempty"`
	Id string `json:"id,omitempty"`
	Updated_at string `json:"updated_at,omitempty"`
	User string `json:"user,omitempty"`
	Fork_of map[string]interface{} `json:"fork_of,omitempty"` // Gist
	Git_push_url string `json:"git_push_url,omitempty"`
	Description string `json:"description,omitempty"`
	History []GeneratedType `json:"history,omitempty"`
	Public bool `json:"public,omitempty"`
	Forks []map[string]interface{} `json:"forks,omitempty"`
	Commits_url string `json:"commits_url,omitempty"`
	Git_pull_url string `json:"git_pull_url,omitempty"`
	Comments_enabled bool `json:"comments_enabled,omitempty"`
	Comments_url string `json:"comments_url,omitempty"`
	Owner GeneratedType `json:"owner,omitempty"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Base_ref string `json:"base_ref"` // The full ref of the branch the merge group will be merged into.
	Base_sha string `json:"base_sha"` // The SHA of the merge group's parent commit.
	Head_commit GeneratedType `json:"head_commit"` // A commit.
	Head_ref string `json:"head_ref"` // The full ref of the merge group.
	Head_sha string `json:"head_sha"` // The SHA of the merge group.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Avatar_url string `json:"avatar_url"`
	Starred_url string `json:"starred_url"`
	Email string `json:"email,omitempty"`
	Following_url string `json:"following_url"`
	Id int64 `json:"id"`
	Login string `json:"login"`
	Events_url string `json:"events_url"`
	Html_url string `json:"html_url"`
	Repos_url string `json:"repos_url"`
	Url string `json:"url"`
	Subscriptions_url string `json:"subscriptions_url"`
	TypeField string `json:"type"`
	Received_events_url string `json:"received_events_url"`
	Followers_url string `json:"followers_url"`
	Name string `json:"name,omitempty"`
	Gravatar_id string `json:"gravatar_id"`
	User_view_type string `json:"user_view_type,omitempty"`
	Site_admin bool `json:"site_admin"`
	Starred_at string `json:"starred_at,omitempty"`
	Gists_url string `json:"gists_url"`
	Organizations_url string `json:"organizations_url"`
	Node_id string `json:"node_id"`
}

// Hovercard represents the Hovercard schema from the OpenAPI specification
type Hovercard struct {
	Contexts []map[string]interface{} `json:"contexts"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Payload string `json:"payload"` // A URL-encoded string of the check_run.completed JSON payload. The decoded payload is a JSON object.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id int `json:"id"`
	Login string `json:"login"`
	Name string `json:"name"`
	Node_id string `json:"node_id"`
	Avatar_url string `json:"avatar_url"`
	Html_url string `json:"html_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url"`
	Provider string `json:"provider"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Membership Webhooksmembership `json:"membership"` // The membership between the user and the organization. Not present when the action is `member_invited`.
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	TypeField string `json:"type"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Assignee GeneratedType `json:"assignee"` // A GitHub user.
	Requested_teams []GeneratedType `json:"requested_teams,omitempty"`
	Updated_at string `json:"updated_at"`
	Deletions int `json:"deletions"`
	Diff_url string `json:"diff_url"`
	Commits_url string `json:"commits_url"`
	Milestone GeneratedType `json:"milestone"` // A collection of related issues and pull requests.
	Node_id string `json:"node_id"`
	Additions int `json:"additions"`
	Created_at string `json:"created_at"`
	Mergeable_state string `json:"mergeable_state"`
	Commits int `json:"commits"`
	Html_url string `json:"html_url"`
	Comments_url string `json:"comments_url"`
	Active_lock_reason string `json:"active_lock_reason,omitempty"`
	Url string `json:"url"`
	Maintainer_can_modify bool `json:"maintainer_can_modify"` // Indicates whether maintainers can modify the pull request.
	Issue_url string `json:"issue_url"`
	Changed_files int `json:"changed_files"`
	Review_comments int `json:"review_comments"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Requested_reviewers []GeneratedType `json:"requested_reviewers,omitempty"`
	User GeneratedType `json:"user"` // A GitHub user.
	Head map[string]interface{} `json:"head"`
	Merged_at string `json:"merged_at"`
	Body string `json:"body"`
	Merged_by GeneratedType `json:"merged_by"` // A GitHub user.
	Links map[string]interface{} `json:"_links"`
	Locked bool `json:"locked"`
	Assignees []GeneratedType `json:"assignees,omitempty"`
	Review_comments_url string `json:"review_comments_url"`
	Merged bool `json:"merged"`
	Draft bool `json:"draft,omitempty"` // Indicates whether or not the pull request is a draft.
	Rebaseable bool `json:"rebaseable,omitempty"`
	Auto_merge GeneratedType `json:"auto_merge"` // The status of auto merging a pull request.
	Mergeable bool `json:"mergeable"`
	Closed_at string `json:"closed_at"`
	State string `json:"state"` // State of this Pull Request. Either `open` or `closed`.
	Statuses_url string `json:"statuses_url"`
	Review_comment_url string `json:"review_comment_url"`
	Base map[string]interface{} `json:"base"`
	Comments int `json:"comments"`
	Title string `json:"title"` // The title of the pull request.
	Labels []map[string]interface{} `json:"labels"`
	Merge_commit_sha string `json:"merge_commit_sha"`
	Id int64 `json:"id"`
	Patch_url string `json:"patch_url"`
	Number int `json:"number"` // Number uniquely identifying the pull request within its repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Values_editable_by string `json:"values_editable_by,omitempty"` // Who can edit the values of the property
	Allowed_values []string `json:"allowed_values,omitempty"` // An ordered list of the allowed values of the property. The property can have up to 200 allowed values.
	Default_value string `json:"default_value,omitempty"` // Default value of the property
	Description string `json:"description,omitempty"` // Short description of the property
	Required bool `json:"required,omitempty"` // Whether the property is required.
	Value_type string `json:"value_type"` // The type of the value for the property
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Display_name string `json:"display_name"` // The display name of the machine includes cores, memory, and storage.
	Memory_in_bytes int `json:"memory_in_bytes"` // How much memory is available to the codespace.
	Name string `json:"name"` // The name of the machine.
	Operating_system string `json:"operating_system"` // The operating system of the machine.
	Prebuild_availability string `json:"prebuild_availability"` // Whether a prebuild is currently available when creating a codespace for this machine and repository. If a branch was not specified as a ref, the default branch will be assumed. Value will be "null" if prebuilds are not supported or prebuild availability could not be determined. Value will be "none" if no prebuild is available. Latest values "ready" and "in_progress" indicate the prebuild availability status.
	Storage_in_bytes int `json:"storage_in_bytes"` // How much storage is available to the codespace.
	Cpus int `json:"cpus"` // How many cores are available to the codespace.
}

// Webhooksmarketplacepurchase represents the Webhooksmarketplacepurchase schema from the OpenAPI specification
type Webhooksmarketplacepurchase struct {
	Billing_cycle string `json:"billing_cycle"`
	Free_trial_ends_on string `json:"free_trial_ends_on"`
	Next_billing_date string `json:"next_billing_date"`
	On_free_trial bool `json:"on_free_trial"`
	Plan map[string]interface{} `json:"plan"`
	Unit_count int `json:"unit_count"`
	Account map[string]interface{} `json:"account"`
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

// Webhooksrelease1 represents the Webhooksrelease1 schema from the OpenAPI specification
type Webhooksrelease1 struct {
	Name string `json:"name"`
	Body string `json:"body"`
	Created_at string `json:"created_at"`
	Tag_name string `json:"tag_name"` // The name of the tag.
	Draft bool `json:"draft"` // Whether the release is a draft or published
	Prerelease bool `json:"prerelease"` // Whether the release is identified as a prerelease or a full release.
	Id int `json:"id"`
	Target_commitish string `json:"target_commitish"` // Specifies the commitish value that determines where the Git tag is created from.
	Discussion_url string `json:"discussion_url,omitempty"`
	Published_at string `json:"published_at"`
	Author map[string]interface{} `json:"author"`
	Node_id string `json:"node_id"`
	Reactions map[string]interface{} `json:"reactions,omitempty"`
	Upload_url string `json:"upload_url"`
	Tarball_url string `json:"tarball_url"`
	Assets []map[string]interface{} `json:"assets"`
	Assets_url string `json:"assets_url"`
	Url string `json:"url"`
	Zipball_url string `json:"zipball_url"`
	Html_url string `json:"html_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Total_active_caches_count int `json:"total_active_caches_count"` // The count of active caches across all repositories of an enterprise or an organization.
	Total_active_caches_size_in_bytes int `json:"total_active_caches_size_in_bytes"` // The total size in bytes of all active cache items across all repositories of an enterprise or an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Current_user_can_bypass string `json:"current_user_can_bypass,omitempty"` // The bypass type of the user making the API request for this ruleset. This field is only returned when querying the repository-level endpoint.
	Id int `json:"id"` // The ID of the ruleset
	Rules []GeneratedType `json:"rules,omitempty"`
	Source_type string `json:"source_type,omitempty"` // The type of the source of the ruleset
	Target string `json:"target,omitempty"` // The target of the ruleset
	Created_at string `json:"created_at,omitempty"`
	Name string `json:"name"` // The name of the ruleset
	Updated_at string `json:"updated_at,omitempty"`
	Conditions interface{} `json:"conditions,omitempty"`
	Node_id string `json:"node_id,omitempty"`
	Source string `json:"source"` // The name of the source
	Links map[string]interface{} `json:"_links,omitempty"`
	Bypass_actors []GeneratedType `json:"bypass_actors,omitempty"` // The actors that can bypass the rules in this ruleset
	Enforcement string `json:"enforcement"` // The enforcement level of the ruleset. `evaluate` allows admins to test rules before enforcing them. Admins can view insights on the Rule Insights page (`evaluate` is only available with GitHub Enterprise).
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Invitation map[string]interface{} `json:"invitation"` // The invitation for the user or email if the action is `member_invited`.
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	User Webhooksuser `json:"user,omitempty"`
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Key Webhooksdeploykey `json:"key"` // The [`deploy key`](https://docs.github.com/rest/deploy-keys/deploy-keys#get-a-deploy-key) resource.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Cvss_v3 map[string]interface{} `json:"cvss_v3,omitempty"`
	Cvss_v4 map[string]interface{} `json:"cvss_v4,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Description string `json:"description"` // The optional human-readable description added to the status.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Created_at string `json:"created_at"`
	Target_url string `json:"target_url"` // The optional link added to the status.
	Updated_at string `json:"updated_at"`
	Id int `json:"id"` // The unique identifier of the status.
	Branches []map[string]interface{} `json:"branches"` // An array of branch objects containing the status' SHA. Each branch contains the given SHA, but the SHA may or may not be the head of the branch. The array includes a maximum of 10 branches.
	Commit map[string]interface{} `json:"commit"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Name string `json:"name"`
	State string `json:"state"` // The new state. Can be `pending`, `success`, `failure`, or `error`.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Context string `json:"context"`
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sha string `json:"sha"` // The Commit SHA.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Avatar_url string `json:"avatar_url,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enabled bool `json:"enabled"` // Whether Dependabot security updates are enabled for the repository.
	Paused bool `json:"paused"` // Whether Dependabot security updates are paused for the repository.
}

// Blob represents the Blob schema from the OpenAPI specification
type Blob struct {
	Highlighted_content string `json:"highlighted_content,omitempty"`
	Node_id string `json:"node_id"`
	Sha string `json:"sha"`
	Size int `json:"size"`
	Url string `json:"url"`
	Content string `json:"content"`
	Encoding string `json:"encoding"`
}

// Vulnerability represents the Vulnerability schema from the OpenAPI specification
type Vulnerability struct {
	Vulnerable_version_range string `json:"vulnerable_version_range"` // The range of the package versions affected by the vulnerability.
	First_patched_version string `json:"first_patched_version"` // The package version that resolves the vulnerability.
	PackageField map[string]interface{} `json:"package"` // The name of the package affected by the vulnerability.
	Vulnerable_functions []string `json:"vulnerable_functions"` // The functions in the package that are affected by the vulnerability.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Count int `json:"count"`
	Path string `json:"path"`
	Title string `json:"title"`
	Uniques int `json:"uniques"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository_url string `json:"repository_url,omitempty"`
	Subscribed bool `json:"subscribed"`
	Thread_url string `json:"thread_url,omitempty"`
	Url string `json:"url"`
	Created_at string `json:"created_at"`
	Ignored bool `json:"ignored"`
	Reason string `json:"reason"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Runs_rerequestable bool `json:"runs_rerequestable,omitempty"`
	Created_at string `json:"created_at"`
	Head_commit GeneratedType `json:"head_commit"` // A commit.
	After string `json:"after"`
	App GeneratedType `json:"app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Head_branch string `json:"head_branch"`
	Status string `json:"status"` // The phase of the lifecycle that the check suite is currently in. Statuses of waiting, requested, and pending are reserved for GitHub Actions check suites.
	Check_runs_url string `json:"check_runs_url"`
	Url string `json:"url"`
	Latest_check_runs_count int `json:"latest_check_runs_count"`
	Repository GeneratedType `json:"repository"` // Minimal Repository
	Rerequestable bool `json:"rerequestable,omitempty"`
	Before string `json:"before"`
	Node_id string `json:"node_id"`
	Pull_requests []GeneratedType `json:"pull_requests"`
	Head_sha string `json:"head_sha"` // The SHA of the head commit that is being checked.
	Id int64 `json:"id"`
	Updated_at string `json:"updated_at"`
	Conclusion string `json:"conclusion"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Body_text string `json:"body_text,omitempty"`
	Html_url string `json:"html_url"`
	Id int64 `json:"id"` // Unique identifier of the issue comment
	Body string `json:"body,omitempty"` // Contents of the issue comment
	Created_at string `json:"created_at"`
	Issue_url string `json:"issue_url"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Body_html string `json:"body_html,omitempty"`
	Node_id string `json:"node_id"`
	Reactions GeneratedType `json:"reactions,omitempty"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"` // URL for the issue comment
	User GeneratedType `json:"user"` // A GitHub user.
	Author_association string `json:"author_association"` // How the author is associated with the repository.
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
	Description string `json:"description,omitempty"`
	Id string `json:"id"`
	Name string `json:"name"`
	Color string `json:"color,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization_permission string `json:"organization_permission,omitempty"` // The organization permission for this project. Only present when owner is an organization.
	Owner_url string `json:"owner_url"`
	Private bool `json:"private,omitempty"` // Whether the project is private or not. Only present when owner is an organization.
	Updated_at string `json:"updated_at"`
	Creator GeneratedType `json:"creator"` // A GitHub user.
	Number int `json:"number"`
	State string `json:"state"`
	Columns_url string `json:"columns_url"`
	Node_id string `json:"node_id"`
	Created_at string `json:"created_at"`
	Permissions map[string]interface{} `json:"permissions"`
	Url string `json:"url"`
	Name string `json:"name"`
	Html_url string `json:"html_url"`
	Body string `json:"body"`
	Id int `json:"id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Billable map[string]interface{} `json:"billable"`
	Run_duration_ms int `json:"run_duration_ms,omitempty"`
}

// Metadata represents the Metadata schema from the OpenAPI specification
type Metadata struct {
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	TypeField string `json:"type"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Parent_issue Issue `json:"parent_issue"` // Issues are a great way to keep track of tasks, enhancements, and bugs for your projects.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sub_issue_id float64 `json:"sub_issue_id"` // The ID of the sub-issue.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Parent_issue_id float64 `json:"parent_issue_id"` // The ID of the parent issue.
	Parent_issue_repo Repository `json:"parent_issue_repo"` // A repository on GitHub.
	Sub_issue Issue `json:"sub_issue"` // Issues are a great way to keep track of tasks, enhancements, and bugs for your projects.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Team Webhooksteam `json:"team"` // Groups of organization members that gives permissions on specified repositories.
	Action string `json:"action"`
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Sender map[string]interface{} `json:"sender"`
	Member Webhooksuser `json:"member"`
	Scope string `json:"scope"` // The scope of the membership. Currently, can only be `team`.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Name string `json:"name"` // The name of the repository property to target
	Property_values []string `json:"property_values"` // The values to match for the repository property
	Source string `json:"source,omitempty"` // The source of the repository property. Defaults to 'custom' if not specified.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Personal_access_token_request GeneratedType `json:"personal_access_token_request"` // Details of a Personal Access Token Request.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Subscriptions_url string `json:"subscriptions_url"`
	Url string `json:"url"`
	Gravatar_id string `json:"gravatar_id"`
	Id int64 `json:"id"`
	Organizations_url string `json:"organizations_url"`
	Role_name string `json:"role_name"`
	TypeField string `json:"type"`
	Followers_url string `json:"followers_url"`
	Html_url string `json:"html_url"`
	Repos_url string `json:"repos_url"`
	Name string `json:"name,omitempty"`
	Node_id string `json:"node_id"`
	Email string `json:"email,omitempty"`
	User_view_type string `json:"user_view_type,omitempty"`
	Avatar_url string `json:"avatar_url"`
	Events_url string `json:"events_url"`
	Starred_url string `json:"starred_url"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Received_events_url string `json:"received_events_url"`
	Site_admin bool `json:"site_admin"`
	Following_url string `json:"following_url"`
	Login string `json:"login"`
	Gists_url string `json:"gists_url"`
}

// Webhookschanges8 represents the Webhookschanges8 schema from the OpenAPI specification
type Webhookschanges8 struct {
	Tier map[string]interface{} `json:"tier"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Projects_v2_item GeneratedType `json:"projects_v2_item"` // An item belonging to a project
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes Webhooksprojectchanges `json:"changes"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
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
	Message string `json:"message,omitempty"` // Commit message to be used.
	Target_ref string `json:"target_ref,omitempty"` // The Git reference of target branch for the commit. Branch needs to already exist. For more information, see "[Git References](https://git-scm.com/book/en/v2/Git-Internals-Git-References)" in the Git documentation.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Site_admin bool `json:"site_admin"`
	Company string `json:"company"`
	Public_repos int `json:"public_repos"`
	Location string `json:"location"`
	Login string `json:"login"`
	Received_events_url string `json:"received_events_url"`
	Twitter_username string `json:"twitter_username,omitempty"`
	Bio string `json:"bio"`
	Url string `json:"url"`
	Collaborators int `json:"collaborators,omitempty"`
	Id int64 `json:"id"`
	Events_url string `json:"events_url"`
	Starred_url string `json:"starred_url"`
	Following int `json:"following"`
	Node_id string `json:"node_id"`
	Hireable bool `json:"hireable"`
	Owned_private_repos int `json:"owned_private_repos,omitempty"`
	Subscriptions_url string `json:"subscriptions_url"`
	Repos_url string `json:"repos_url"`
	Notification_email string `json:"notification_email,omitempty"`
	Gravatar_id string `json:"gravatar_id"`
	Private_gists int `json:"private_gists,omitempty"`
	Blog string `json:"blog"`
	Followers_url string `json:"followers_url"`
	Name string `json:"name"`
	Plan map[string]interface{} `json:"plan,omitempty"`
	Avatar_url string `json:"avatar_url"`
	Followers int `json:"followers"`
	Following_url string `json:"following_url"`
	Public_gists int `json:"public_gists"`
	TypeField string `json:"type"`
	Gists_url string `json:"gists_url"`
	Total_private_repos int `json:"total_private_repos,omitempty"`
	Html_url string `json:"html_url"`
	Email string `json:"email"`
	Organizations_url string `json:"organizations_url"`
	Updated_at string `json:"updated_at"`
	User_view_type string `json:"user_view_type,omitempty"`
	Disk_usage int `json:"disk_usage,omitempty"`
	Created_at string `json:"created_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes,omitempty"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Membership Webhooksmembership `json:"membership,omitempty"` // The membership between the user and the organization. Not present when the action is `member_invited`.
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Number int `json:"number"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// Collaborator represents the Collaborator schema from the OpenAPI specification
type Collaborator struct {
	Subscriptions_url string `json:"subscriptions_url"`
	Url string `json:"url"`
	Node_id string `json:"node_id"`
	Followers_url string `json:"followers_url"`
	TypeField string `json:"type"`
	Email string `json:"email,omitempty"`
	Id int64 `json:"id"`
	Site_admin bool `json:"site_admin"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	User_view_type string `json:"user_view_type,omitempty"`
	Organizations_url string `json:"organizations_url"`
	Following_url string `json:"following_url"`
	Received_events_url string `json:"received_events_url"`
	Html_url string `json:"html_url"`
	Repos_url string `json:"repos_url"`
	Role_name string `json:"role_name"`
	Starred_url string `json:"starred_url"`
	Events_url string `json:"events_url"`
	Gists_url string `json:"gists_url"`
	Login string `json:"login"`
	Avatar_url string `json:"avatar_url"`
	Gravatar_id string `json:"gravatar_id"`
	Name string `json:"name,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Workflow_job map[string]interface{} `json:"workflow_job"`
	Action string `json:"action"`
	Deployment Deployment `json:"deployment,omitempty"` // A request for a specific ref(branch,sha,tag) to be deployed
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repositories []map[string]interface{} `json:"repositories,omitempty"` // An array of repository objects that the installation can access.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Requester interface{} `json:"requester,omitempty"`
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation Installation `json:"installation"` // Installation
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Project_card Webhooksprojectcard `json:"project_card"`
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// Job represents the Job schema from the OpenAPI specification
type Job struct {
	Started_at string `json:"started_at"` // The time that the job started, in ISO 8601 format.
	Head_sha string `json:"head_sha"` // The SHA of the commit that is being run.
	Runner_name string `json:"runner_name"` // The name of the runner to which this job has been assigned. (If a runner hasn't yet been assigned, this will be null.)
	Run_attempt int `json:"run_attempt,omitempty"` // Attempt number of the associated workflow run, 1 for first attempt and higher if the workflow was re-run.
	Id int `json:"id"` // The id of the job.
	Conclusion string `json:"conclusion"` // The outcome of the job.
	Runner_id int `json:"runner_id"` // The ID of the runner to which this job has been assigned. (If a runner hasn't yet been assigned, this will be null.)
	Steps []map[string]interface{} `json:"steps,omitempty"` // Steps in this job.
	Runner_group_name string `json:"runner_group_name"` // The name of the runner group to which this job has been assigned. (If a runner hasn't yet been assigned, this will be null.)
	Name string `json:"name"` // The name of the job.
	Runner_group_id int `json:"runner_group_id"` // The ID of the runner group to which this job has been assigned. (If a runner hasn't yet been assigned, this will be null.)
	Head_branch string `json:"head_branch"` // The name of the current branch.
	Check_run_url string `json:"check_run_url"`
	Url string `json:"url"`
	Workflow_name string `json:"workflow_name"` // The name of the workflow.
	Created_at string `json:"created_at"` // The time that the job created, in ISO 8601 format.
	Run_url string `json:"run_url"`
	Status string `json:"status"` // The phase of the lifecycle that the job is currently in.
	Node_id string `json:"node_id"`
	Run_id int `json:"run_id"` // The id of the associated workflow run.
	Completed_at string `json:"completed_at"` // The time that the job finished, in ISO 8601 format.
	Html_url string `json:"html_url"`
	Labels []string `json:"labels"` // Labels for the workflow job. Specified by the "runs_on" attribute in the action's workflow file.
}

// Webhooksissue2 represents the Webhooksissue2 schema from the OpenAPI specification
type Webhooksissue2 struct {
	Pull_request map[string]interface{} `json:"pull_request,omitempty"`
	Assignee map[string]interface{} `json:"assignee,omitempty"`
	Events_url string `json:"events_url"`
	Comments_url string `json:"comments_url"`
	Number int `json:"number"`
	Performed_via_github_app map[string]interface{} `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Labels []map[string]interface{} `json:"labels,omitempty"`
	Body string `json:"body"` // Contents of the issue
	Labels_url string `json:"labels_url"`
	Milestone map[string]interface{} `json:"milestone"` // A collection of related issues and pull requests.
	Created_at string `json:"created_at"`
	Assignees []map[string]interface{} `json:"assignees"`
	Timeline_url string `json:"timeline_url,omitempty"`
	Updated_at string `json:"updated_at"`
	Comments int `json:"comments"`
	Url string `json:"url"` // URL for the issue
	State string `json:"state,omitempty"` // State of the issue; either 'open' or 'closed'
	Node_id string `json:"node_id"`
	Closed_at string `json:"closed_at"`
	State_reason string `json:"state_reason,omitempty"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Id int64 `json:"id"`
	Reactions map[string]interface{} `json:"reactions"`
	Title string `json:"title"` // Title of the issue
	Repository_url string `json:"repository_url"`
	TypeField GeneratedType `json:"type,omitempty"` // The type of issue.
	Active_lock_reason string `json:"active_lock_reason"`
	Locked bool `json:"locked,omitempty"`
	Html_url string `json:"html_url"`
	Sub_issues_summary map[string]interface{} `json:"sub_issues_summary,omitempty"`
	User map[string]interface{} `json:"user"`
	Draft bool `json:"draft,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Alert Webhooksalert `json:"alert"` // The security alert of the vulnerable dependency.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes,omitempty"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Project_card interface{} `json:"project_card"`
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// Integration represents the Integration schema from the OpenAPI specification
type Integration struct {
	Created_at string `json:"created_at"`
	Installations_count int `json:"installations_count,omitempty"` // The number of installations associated with the GitHub app. Only returned when the integration is requesting details about itself.
	Slug string `json:"slug,omitempty"` // The slug name of the GitHub app
	Updated_at string `json:"updated_at"`
	Events []string `json:"events"` // The list of events for the GitHub app. Note that the `installation_target`, `security_advisory`, and `meta` events are not included because they are global events and not specific to an installation.
	External_url string `json:"external_url"`
	Html_url string `json:"html_url"`
	Owner interface{} `json:"owner"`
	Client_id string `json:"client_id,omitempty"`
	Description string `json:"description"`
	Permissions map[string]interface{} `json:"permissions"` // The set of permissions for the GitHub app
	Id int `json:"id"` // Unique identifier of the GitHub app
	Name string `json:"name"` // The name of the GitHub app
	Node_id string `json:"node_id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Rule Webhooksrule `json:"rule"` // The branch protection rule. Includes a `name` and all the [branch protection settings](https://docs.github.com/github/administering-a-repository/defining-the-mergeability-of-pull-requests/about-protected-branches#about-branch-protection-settings) applied to branches that match the name. Binary settings are boolean. Multi-level configurations are one of `off`, `non_admins`, or `everyone`. Actor and build lists are arrays of strings.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes,omitempty"` // If the action was `edited`, the changes to the rule.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// Webhooksprojectchanges represents the Webhooksprojectchanges schema from the OpenAPI specification
type Webhooksprojectchanges struct {
	Archived_at map[string]interface{} `json:"archived_at,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Path string `json:"path"`
	Ref string `json:"ref,omitempty"`
	Sha string `json:"sha"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Last_activity_editor string `json:"last_activity_editor,omitempty"` // Last editor that was used by the user for a GitHub Copilot completion.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Updated_at string `json:"updated_at,omitempty"` // **Closing down notice:** This field is no longer relevant and is closing down. Use the `created_at` field to determine when the assignee was last granted access to GitHub Copilot. Timestamp of when the assignee's GitHub Copilot access was last updated, in ISO 8601 format.
	Pending_cancellation_date string `json:"pending_cancellation_date,omitempty"` // The pending cancellation date for the seat, in `YYYY-MM-DD` format. This will be null unless the assignee's Copilot access has been canceled during the current billing cycle. If the seat has been cancelled, this corresponds to the start of the organization's next billing cycle.
	Plan_type string `json:"plan_type,omitempty"` // The Copilot plan of the organization, or the parent enterprise, when applicable.
	Assigning_team interface{} `json:"assigning_team,omitempty"` // The team through which the assignee is granted access to GitHub Copilot, if applicable.
	Last_activity_at string `json:"last_activity_at,omitempty"` // Timestamp of user's last GitHub Copilot activity, in ISO 8601 format.
	Assignee GeneratedType `json:"assignee,omitempty"` // A GitHub user.
	Created_at string `json:"created_at"` // Timestamp of when the assignee was last granted access to GitHub Copilot, in ISO 8601 format.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Token string `json:"token"`
	Expires_at string `json:"expires_at"`
	Has_multiple_single_files bool `json:"has_multiple_single_files,omitempty"`
	Permissions GeneratedType `json:"permissions,omitempty"` // The permissions granted to the user access token.
	Repositories []Repository `json:"repositories,omitempty"`
	Repository_selection string `json:"repository_selection,omitempty"`
	Single_file string `json:"single_file,omitempty"`
	Single_file_paths []string `json:"single_file_paths,omitempty"`
}

// Webhooksreview represents the Webhooksreview schema from the OpenAPI specification
type Webhooksreview struct {
	Node_id string `json:"node_id"`
	State string `json:"state"`
	Submitted_at string `json:"submitted_at"`
	Pull_request_url string `json:"pull_request_url"`
	Body string `json:"body"` // The text of the review.
	User map[string]interface{} `json:"user"`
	Links map[string]interface{} `json:"_links"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Commit_id string `json:"commit_id"` // A commit SHA for the review.
	Html_url string `json:"html_url"`
	Id int `json:"id"` // Unique identifier of the review
}

// Actor represents the Actor schema from the OpenAPI specification
type Actor struct {
	Avatar_url string `json:"avatar_url"`
	Display_login string `json:"display_login,omitempty"`
	Gravatar_id string `json:"gravatar_id"`
	Id int `json:"id"`
	Login string `json:"login"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Commit_id string `json:"commit_id"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Project_card map[string]interface{} `json:"project_card,omitempty"`
	Url string `json:"url"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Event string `json:"event"`
	Node_id string `json:"node_id"`
	Created_at string `json:"created_at"`
	Commit_url string `json:"commit_url"`
	Id int `json:"id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Related []map[string]interface{} `json:"related,omitempty"`
	Repository_count int `json:"repository_count,omitempty"`
	Created_by string `json:"created_by"`
	Featured bool `json:"featured"`
	Text_matches []map[string]interface{} `json:"text_matches,omitempty"`
	Logo_url string `json:"logo_url,omitempty"`
	Name string `json:"name"`
	Created_at string `json:"created_at"`
	Display_name string `json:"display_name"`
	Aliases []map[string]interface{} `json:"aliases,omitempty"`
	Released string `json:"released"`
	Score float64 `json:"score"`
	Short_description string `json:"short_description"`
	Description string `json:"description"`
	Updated_at string `json:"updated_at"`
	Curated bool `json:"curated"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Review map[string]interface{} `json:"review"` // The review that was affected.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Access_level string `json:"access_level"` // Defines the level of access that workflows outside of the repository have to actions and reusable workflows within the repository. `none` means the access is only possible from workflows in this repository. `user` level access allows sharing across user owned private repositories only. `organization` level access allows sharing across the organization.
}

// Webhooksdeploykey represents the Webhooksdeploykey schema from the OpenAPI specification
type Webhooksdeploykey struct {
	Read_only bool `json:"read_only"`
	Title string `json:"title"`
	Id int `json:"id"`
	Added_by string `json:"added_by,omitempty"`
	Url string `json:"url"`
	Verified bool `json:"verified"`
	Created_at string `json:"created_at"`
	Enabled bool `json:"enabled,omitempty"`
	Key string `json:"key"`
	Last_used string `json:"last_used,omitempty"`
}

// Topic represents the Topic schema from the OpenAPI specification
type Topic struct {
	Names []string `json:"names"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Path string `json:"path"`
	Ref_name string `json:"ref_name"`
	Size int `json:"size"`
	Oid string `json:"oid"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Reason string `json:"reason,omitempty"`
	Number int `json:"number"`
	Pull_request map[string]interface{} `json:"pull_request"`
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Scope string `json:"scope"` // The scope of the membership. Currently, can only be `team`.
	Sender map[string]interface{} `json:"sender"`
	Team Webhooksteam `json:"team"` // Groups of organization members that gives permissions on specified repositories.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Action string `json:"action"`
	Member Webhooksuser `json:"member"`
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}
