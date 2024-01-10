// Code generated by ent, DO NOT EDIT.

package teams

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the teams type in the database.
	Label = "teams"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "team_id"
	// FieldTeamName holds the string denoting the team_name field in the database.
	FieldTeamName = "team_name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldJiraKeys holds the string denoting the jira_keys field in the database.
	FieldJiraKeys = "jira_keys"
	// EdgeRepositories holds the string denoting the repositories edge name in mutations.
	EdgeRepositories = "repositories"
	// EdgeBugs holds the string denoting the bugs edge name in mutations.
	EdgeBugs = "bugs"
	// EdgeFailures holds the string denoting the failures edge name in mutations.
	EdgeFailures = "failures"
	// RepositoryFieldID holds the string denoting the ID field of the Repository.
	RepositoryFieldID = "id"
	// BugsFieldID holds the string denoting the ID field of the Bugs.
	BugsFieldID = "id"
	// FailureFieldID holds the string denoting the ID field of the Failure.
	FailureFieldID = "id"
	// Table holds the table name of the teams in the database.
	Table = "teams"
	// RepositoriesTable is the table that holds the repositories relation/edge.
	RepositoriesTable = "repositories"
	// RepositoriesInverseTable is the table name for the Repository entity.
	// It exists in this package in order to avoid circular dependency with the "repository" package.
	RepositoriesInverseTable = "repositories"
	// RepositoriesColumn is the table column denoting the repositories relation/edge.
	RepositoriesColumn = "teams_repositories"
	// BugsTable is the table that holds the bugs relation/edge.
	BugsTable = "bugs"
	// BugsInverseTable is the table name for the Bugs entity.
	// It exists in this package in order to avoid circular dependency with the "bugs" package.
	BugsInverseTable = "bugs"
	// BugsColumn is the table column denoting the bugs relation/edge.
	BugsColumn = "teams_bugs"
	// FailuresTable is the table that holds the failures relation/edge.
	FailuresTable = "failures"
	// FailuresInverseTable is the table name for the Failure entity.
	// It exists in this package in order to avoid circular dependency with the "failure" package.
	FailuresInverseTable = "failures"
	// FailuresColumn is the table column denoting the failures relation/edge.
	FailuresColumn = "teams_failures"
)

// Columns holds all SQL columns for teams fields.
var Columns = []string{
	FieldID,
	FieldTeamName,
	FieldDescription,
	FieldJiraKeys,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
