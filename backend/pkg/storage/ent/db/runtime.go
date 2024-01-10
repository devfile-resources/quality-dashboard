// Code generated by ent, DO NOT EDIT.

package db

import (
	"github.com/google/uuid"
	"github.com/redhat-appstudio/quality-studio/pkg/storage/ent/db/bugs"
	"github.com/redhat-appstudio/quality-studio/pkg/storage/ent/db/codecov"
	"github.com/redhat-appstudio/quality-studio/pkg/storage/ent/db/failure"
	"github.com/redhat-appstudio/quality-studio/pkg/storage/ent/db/pullrequests"
	"github.com/redhat-appstudio/quality-studio/pkg/storage/ent/db/repository"
	"github.com/redhat-appstudio/quality-studio/pkg/storage/ent/db/teams"
	"github.com/redhat-appstudio/quality-studio/pkg/storage/ent/db/workflows"
	"github.com/redhat-appstudio/quality-studio/pkg/storage/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	bugsFields := schema.Bugs{}.Fields()
	_ = bugsFields
	// bugsDescJiraKey is the schema descriptor for jira_key field.
	bugsDescJiraKey := bugsFields[1].Descriptor()
	// bugs.JiraKeyValidator is a validator for the "jira_key" field. It is called by the builders before save.
	bugs.JiraKeyValidator = bugsDescJiraKey.Validators[0].(func(string) error)
	// bugsDescResolved is the schema descriptor for resolved field.
	bugsDescResolved := bugsFields[5].Descriptor()
	// bugs.DefaultResolved holds the default value on creation for the resolved field.
	bugs.DefaultResolved = bugsDescResolved.Default.(bool)
	// bugsDescResolutionTime is the schema descriptor for resolution_time field.
	bugsDescResolutionTime := bugsFields[7].Descriptor()
	// bugs.DefaultResolutionTime holds the default value on creation for the resolution_time field.
	bugs.DefaultResolutionTime = bugsDescResolutionTime.Default.(float64)
	// bugsDescID is the schema descriptor for id field.
	bugsDescID := bugsFields[0].Descriptor()
	// bugs.DefaultID holds the default value on creation for the id field.
	bugs.DefaultID = bugsDescID.Default.(func() uuid.UUID)
	codecovFields := schema.CodeCov{}.Fields()
	_ = codecovFields
	// codecovDescRepositoryName is the schema descriptor for repository_name field.
	codecovDescRepositoryName := codecovFields[1].Descriptor()
	// codecov.RepositoryNameValidator is a validator for the "repository_name" field. It is called by the builders before save.
	codecov.RepositoryNameValidator = codecovDescRepositoryName.Validators[0].(func(string) error)
	// codecovDescID is the schema descriptor for id field.
	codecovDescID := codecovFields[0].Descriptor()
	// codecov.DefaultID holds the default value on creation for the id field.
	codecov.DefaultID = codecovDescID.Default.(func() uuid.UUID)
	failureFields := schema.Failure{}.Fields()
	_ = failureFields
	// failureDescJiraKey is the schema descriptor for jira_key field.
	failureDescJiraKey := failureFields[1].Descriptor()
	// failure.JiraKeyValidator is a validator for the "jira_key" field. It is called by the builders before save.
	failure.JiraKeyValidator = failureDescJiraKey.Validators[0].(func(string) error)
	// failureDescID is the schema descriptor for id field.
	failureDescID := failureFields[0].Descriptor()
	// failure.DefaultID holds the default value on creation for the id field.
	failure.DefaultID = failureDescID.Default.(func() uuid.UUID)
	pullrequestsFields := schema.PullRequests{}.Fields()
	_ = pullrequestsFields
	// pullrequestsDescPrID is the schema descriptor for pr_id field.
	pullrequestsDescPrID := pullrequestsFields[0].Descriptor()
	// pullrequests.DefaultPrID holds the default value on creation for the pr_id field.
	pullrequests.DefaultPrID = pullrequestsDescPrID.Default.(func() uuid.UUID)
	repositoryFields := schema.Repository{}.Fields()
	_ = repositoryFields
	// repositoryDescRepositoryName is the schema descriptor for repository_name field.
	repositoryDescRepositoryName := repositoryFields[1].Descriptor()
	// repository.RepositoryNameValidator is a validator for the "repository_name" field. It is called by the builders before save.
	repository.RepositoryNameValidator = repositoryDescRepositoryName.Validators[0].(func(string) error)
	// repositoryDescGitOrganization is the schema descriptor for git_organization field.
	repositoryDescGitOrganization := repositoryFields[2].Descriptor()
	// repository.GitOrganizationValidator is a validator for the "git_organization" field. It is called by the builders before save.
	repository.GitOrganizationValidator = repositoryDescGitOrganization.Validators[0].(func(string) error)
	// repositoryDescDescription is the schema descriptor for description field.
	repositoryDescDescription := repositoryFields[3].Descriptor()
	// repository.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	repository.DescriptionValidator = repositoryDescDescription.Validators[0].(func(string) error)
	// repositoryDescGitURL is the schema descriptor for git_url field.
	repositoryDescGitURL := repositoryFields[4].Descriptor()
	// repository.GitURLValidator is a validator for the "git_url" field. It is called by the builders before save.
	repository.GitURLValidator = repositoryDescGitURL.Validators[0].(func(string) error)
	// repositoryDescID is the schema descriptor for id field.
	repositoryDescID := repositoryFields[0].Descriptor()
	// repository.IDValidator is a validator for the "id" field. It is called by the builders before save.
	repository.IDValidator = func() func(string) error {
		validators := repositoryDescID.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(id string) error {
			for _, fn := range fns {
				if err := fn(id); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	teamsFields := schema.Teams{}.Fields()
	_ = teamsFields
	// teamsDescID is the schema descriptor for id field.
	teamsDescID := teamsFields[0].Descriptor()
	// teams.DefaultID holds the default value on creation for the id field.
	teams.DefaultID = teamsDescID.Default.(func() uuid.UUID)
	workflowsFields := schema.Workflows{}.Fields()
	_ = workflowsFields
	// workflowsDescWorkflowID is the schema descriptor for workflow_id field.
	workflowsDescWorkflowID := workflowsFields[0].Descriptor()
	// workflows.DefaultWorkflowID holds the default value on creation for the workflow_id field.
	workflows.DefaultWorkflowID = workflowsDescWorkflowID.Default.(func() uuid.UUID)
	// workflowsDescWorkflowName is the schema descriptor for workflow_name field.
	workflowsDescWorkflowName := workflowsFields[1].Descriptor()
	// workflows.WorkflowNameValidator is a validator for the "workflow_name" field. It is called by the builders before save.
	workflows.WorkflowNameValidator = workflowsDescWorkflowName.Validators[0].(func(string) error)
}