// Code generated by ent, DO NOT EDIT.

package codecov

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/redhat-appstudio/quality-studio/pkg/storage/ent/db/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.CodeCov {
	return predicate.CodeCov(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.CodeCov {
	return predicate.CodeCov(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.CodeCov {
	return predicate.CodeCov(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.CodeCov {
	return predicate.CodeCov(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.CodeCov {
	return predicate.CodeCov(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.CodeCov {
	return predicate.CodeCov(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.CodeCov {
	return predicate.CodeCov(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.CodeCov {
	return predicate.CodeCov(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.CodeCov {
	return predicate.CodeCov(sql.FieldLTE(FieldID, id))
}

// RepositoryName applies equality check predicate on the "repository_name" field. It's identical to RepositoryNameEQ.
func RepositoryName(v string) predicate.CodeCov {
	return predicate.CodeCov(sql.FieldEQ(FieldRepositoryName, v))
}

// GitOrganization applies equality check predicate on the "git_organization" field. It's identical to GitOrganizationEQ.
func GitOrganization(v string) predicate.CodeCov {
	return predicate.CodeCov(sql.FieldEQ(FieldGitOrganization, v))
}

// CoveragePercentage applies equality check predicate on the "coverage_percentage" field. It's identical to CoveragePercentageEQ.
func CoveragePercentage(v float64) predicate.CodeCov {
	return predicate.CodeCov(sql.FieldEQ(FieldCoveragePercentage, v))
}

// AverageRetestsToMerge applies equality check predicate on the "average_retests_to_merge" field. It's identical to AverageRetestsToMergeEQ.
func AverageRetestsToMerge(v float64) predicate.CodeCov {
	return predicate.CodeCov(sql.FieldEQ(FieldAverageRetestsToMerge, v))
}

// CoverageTrend applies equality check predicate on the "coverage_trend" field. It's identical to CoverageTrendEQ.
func CoverageTrend(v string) predicate.CodeCov {
	return predicate.CodeCov(sql.FieldEQ(FieldCoverageTrend, v))
}

// RepositoryNameEQ applies the EQ predicate on the "repository_name" field.
func RepositoryNameEQ(v string) predicate.CodeCov {
	return predicate.CodeCov(sql.FieldEQ(FieldRepositoryName, v))
}

// RepositoryNameNEQ applies the NEQ predicate on the "repository_name" field.
func RepositoryNameNEQ(v string) predicate.CodeCov {
	return predicate.CodeCov(sql.FieldNEQ(FieldRepositoryName, v))
}

// RepositoryNameIn applies the In predicate on the "repository_name" field.
func RepositoryNameIn(vs ...string) predicate.CodeCov {
	return predicate.CodeCov(sql.FieldIn(FieldRepositoryName, vs...))
}

// RepositoryNameNotIn applies the NotIn predicate on the "repository_name" field.
func RepositoryNameNotIn(vs ...string) predicate.CodeCov {
	return predicate.CodeCov(sql.FieldNotIn(FieldRepositoryName, vs...))
}

// RepositoryNameGT applies the GT predicate on the "repository_name" field.
func RepositoryNameGT(v string) predicate.CodeCov {
	return predicate.CodeCov(sql.FieldGT(FieldRepositoryName, v))
}

// RepositoryNameGTE applies the GTE predicate on the "repository_name" field.
func RepositoryNameGTE(v string) predicate.CodeCov {
	return predicate.CodeCov(sql.FieldGTE(FieldRepositoryName, v))
}

// RepositoryNameLT applies the LT predicate on the "repository_name" field.
func RepositoryNameLT(v string) predicate.CodeCov {
	return predicate.CodeCov(sql.FieldLT(FieldRepositoryName, v))
}

// RepositoryNameLTE applies the LTE predicate on the "repository_name" field.
func RepositoryNameLTE(v string) predicate.CodeCov {
	return predicate.CodeCov(sql.FieldLTE(FieldRepositoryName, v))
}

// RepositoryNameContains applies the Contains predicate on the "repository_name" field.
func RepositoryNameContains(v string) predicate.CodeCov {
	return predicate.CodeCov(sql.FieldContains(FieldRepositoryName, v))
}

// RepositoryNameHasPrefix applies the HasPrefix predicate on the "repository_name" field.
func RepositoryNameHasPrefix(v string) predicate.CodeCov {
	return predicate.CodeCov(sql.FieldHasPrefix(FieldRepositoryName, v))
}

// RepositoryNameHasSuffix applies the HasSuffix predicate on the "repository_name" field.
func RepositoryNameHasSuffix(v string) predicate.CodeCov {
	return predicate.CodeCov(sql.FieldHasSuffix(FieldRepositoryName, v))
}

// RepositoryNameEqualFold applies the EqualFold predicate on the "repository_name" field.
func RepositoryNameEqualFold(v string) predicate.CodeCov {
	return predicate.CodeCov(sql.FieldEqualFold(FieldRepositoryName, v))
}

// RepositoryNameContainsFold applies the ContainsFold predicate on the "repository_name" field.
func RepositoryNameContainsFold(v string) predicate.CodeCov {
	return predicate.CodeCov(sql.FieldContainsFold(FieldRepositoryName, v))
}

// GitOrganizationEQ applies the EQ predicate on the "git_organization" field.
func GitOrganizationEQ(v string) predicate.CodeCov {
	return predicate.CodeCov(sql.FieldEQ(FieldGitOrganization, v))
}

// GitOrganizationNEQ applies the NEQ predicate on the "git_organization" field.
func GitOrganizationNEQ(v string) predicate.CodeCov {
	return predicate.CodeCov(sql.FieldNEQ(FieldGitOrganization, v))
}

// GitOrganizationIn applies the In predicate on the "git_organization" field.
func GitOrganizationIn(vs ...string) predicate.CodeCov {
	return predicate.CodeCov(sql.FieldIn(FieldGitOrganization, vs...))
}

// GitOrganizationNotIn applies the NotIn predicate on the "git_organization" field.
func GitOrganizationNotIn(vs ...string) predicate.CodeCov {
	return predicate.CodeCov(sql.FieldNotIn(FieldGitOrganization, vs...))
}

// GitOrganizationGT applies the GT predicate on the "git_organization" field.
func GitOrganizationGT(v string) predicate.CodeCov {
	return predicate.CodeCov(sql.FieldGT(FieldGitOrganization, v))
}

// GitOrganizationGTE applies the GTE predicate on the "git_organization" field.
func GitOrganizationGTE(v string) predicate.CodeCov {
	return predicate.CodeCov(sql.FieldGTE(FieldGitOrganization, v))
}

// GitOrganizationLT applies the LT predicate on the "git_organization" field.
func GitOrganizationLT(v string) predicate.CodeCov {
	return predicate.CodeCov(sql.FieldLT(FieldGitOrganization, v))
}

// GitOrganizationLTE applies the LTE predicate on the "git_organization" field.
func GitOrganizationLTE(v string) predicate.CodeCov {
	return predicate.CodeCov(sql.FieldLTE(FieldGitOrganization, v))
}

// GitOrganizationContains applies the Contains predicate on the "git_organization" field.
func GitOrganizationContains(v string) predicate.CodeCov {
	return predicate.CodeCov(sql.FieldContains(FieldGitOrganization, v))
}

// GitOrganizationHasPrefix applies the HasPrefix predicate on the "git_organization" field.
func GitOrganizationHasPrefix(v string) predicate.CodeCov {
	return predicate.CodeCov(sql.FieldHasPrefix(FieldGitOrganization, v))
}

// GitOrganizationHasSuffix applies the HasSuffix predicate on the "git_organization" field.
func GitOrganizationHasSuffix(v string) predicate.CodeCov {
	return predicate.CodeCov(sql.FieldHasSuffix(FieldGitOrganization, v))
}

// GitOrganizationEqualFold applies the EqualFold predicate on the "git_organization" field.
func GitOrganizationEqualFold(v string) predicate.CodeCov {
	return predicate.CodeCov(sql.FieldEqualFold(FieldGitOrganization, v))
}

// GitOrganizationContainsFold applies the ContainsFold predicate on the "git_organization" field.
func GitOrganizationContainsFold(v string) predicate.CodeCov {
	return predicate.CodeCov(sql.FieldContainsFold(FieldGitOrganization, v))
}

// CoveragePercentageEQ applies the EQ predicate on the "coverage_percentage" field.
func CoveragePercentageEQ(v float64) predicate.CodeCov {
	return predicate.CodeCov(sql.FieldEQ(FieldCoveragePercentage, v))
}

// CoveragePercentageNEQ applies the NEQ predicate on the "coverage_percentage" field.
func CoveragePercentageNEQ(v float64) predicate.CodeCov {
	return predicate.CodeCov(sql.FieldNEQ(FieldCoveragePercentage, v))
}

// CoveragePercentageIn applies the In predicate on the "coverage_percentage" field.
func CoveragePercentageIn(vs ...float64) predicate.CodeCov {
	return predicate.CodeCov(sql.FieldIn(FieldCoveragePercentage, vs...))
}

// CoveragePercentageNotIn applies the NotIn predicate on the "coverage_percentage" field.
func CoveragePercentageNotIn(vs ...float64) predicate.CodeCov {
	return predicate.CodeCov(sql.FieldNotIn(FieldCoveragePercentage, vs...))
}

// CoveragePercentageGT applies the GT predicate on the "coverage_percentage" field.
func CoveragePercentageGT(v float64) predicate.CodeCov {
	return predicate.CodeCov(sql.FieldGT(FieldCoveragePercentage, v))
}

// CoveragePercentageGTE applies the GTE predicate on the "coverage_percentage" field.
func CoveragePercentageGTE(v float64) predicate.CodeCov {
	return predicate.CodeCov(sql.FieldGTE(FieldCoveragePercentage, v))
}

// CoveragePercentageLT applies the LT predicate on the "coverage_percentage" field.
func CoveragePercentageLT(v float64) predicate.CodeCov {
	return predicate.CodeCov(sql.FieldLT(FieldCoveragePercentage, v))
}

// CoveragePercentageLTE applies the LTE predicate on the "coverage_percentage" field.
func CoveragePercentageLTE(v float64) predicate.CodeCov {
	return predicate.CodeCov(sql.FieldLTE(FieldCoveragePercentage, v))
}

// AverageRetestsToMergeEQ applies the EQ predicate on the "average_retests_to_merge" field.
func AverageRetestsToMergeEQ(v float64) predicate.CodeCov {
	return predicate.CodeCov(sql.FieldEQ(FieldAverageRetestsToMerge, v))
}

// AverageRetestsToMergeNEQ applies the NEQ predicate on the "average_retests_to_merge" field.
func AverageRetestsToMergeNEQ(v float64) predicate.CodeCov {
	return predicate.CodeCov(sql.FieldNEQ(FieldAverageRetestsToMerge, v))
}

// AverageRetestsToMergeIn applies the In predicate on the "average_retests_to_merge" field.
func AverageRetestsToMergeIn(vs ...float64) predicate.CodeCov {
	return predicate.CodeCov(sql.FieldIn(FieldAverageRetestsToMerge, vs...))
}

// AverageRetestsToMergeNotIn applies the NotIn predicate on the "average_retests_to_merge" field.
func AverageRetestsToMergeNotIn(vs ...float64) predicate.CodeCov {
	return predicate.CodeCov(sql.FieldNotIn(FieldAverageRetestsToMerge, vs...))
}

// AverageRetestsToMergeGT applies the GT predicate on the "average_retests_to_merge" field.
func AverageRetestsToMergeGT(v float64) predicate.CodeCov {
	return predicate.CodeCov(sql.FieldGT(FieldAverageRetestsToMerge, v))
}

// AverageRetestsToMergeGTE applies the GTE predicate on the "average_retests_to_merge" field.
func AverageRetestsToMergeGTE(v float64) predicate.CodeCov {
	return predicate.CodeCov(sql.FieldGTE(FieldAverageRetestsToMerge, v))
}

// AverageRetestsToMergeLT applies the LT predicate on the "average_retests_to_merge" field.
func AverageRetestsToMergeLT(v float64) predicate.CodeCov {
	return predicate.CodeCov(sql.FieldLT(FieldAverageRetestsToMerge, v))
}

// AverageRetestsToMergeLTE applies the LTE predicate on the "average_retests_to_merge" field.
func AverageRetestsToMergeLTE(v float64) predicate.CodeCov {
	return predicate.CodeCov(sql.FieldLTE(FieldAverageRetestsToMerge, v))
}

// AverageRetestsToMergeIsNil applies the IsNil predicate on the "average_retests_to_merge" field.
func AverageRetestsToMergeIsNil() predicate.CodeCov {
	return predicate.CodeCov(sql.FieldIsNull(FieldAverageRetestsToMerge))
}

// AverageRetestsToMergeNotNil applies the NotNil predicate on the "average_retests_to_merge" field.
func AverageRetestsToMergeNotNil() predicate.CodeCov {
	return predicate.CodeCov(sql.FieldNotNull(FieldAverageRetestsToMerge))
}

// CoverageTrendEQ applies the EQ predicate on the "coverage_trend" field.
func CoverageTrendEQ(v string) predicate.CodeCov {
	return predicate.CodeCov(sql.FieldEQ(FieldCoverageTrend, v))
}

// CoverageTrendNEQ applies the NEQ predicate on the "coverage_trend" field.
func CoverageTrendNEQ(v string) predicate.CodeCov {
	return predicate.CodeCov(sql.FieldNEQ(FieldCoverageTrend, v))
}

// CoverageTrendIn applies the In predicate on the "coverage_trend" field.
func CoverageTrendIn(vs ...string) predicate.CodeCov {
	return predicate.CodeCov(sql.FieldIn(FieldCoverageTrend, vs...))
}

// CoverageTrendNotIn applies the NotIn predicate on the "coverage_trend" field.
func CoverageTrendNotIn(vs ...string) predicate.CodeCov {
	return predicate.CodeCov(sql.FieldNotIn(FieldCoverageTrend, vs...))
}

// CoverageTrendGT applies the GT predicate on the "coverage_trend" field.
func CoverageTrendGT(v string) predicate.CodeCov {
	return predicate.CodeCov(sql.FieldGT(FieldCoverageTrend, v))
}

// CoverageTrendGTE applies the GTE predicate on the "coverage_trend" field.
func CoverageTrendGTE(v string) predicate.CodeCov {
	return predicate.CodeCov(sql.FieldGTE(FieldCoverageTrend, v))
}

// CoverageTrendLT applies the LT predicate on the "coverage_trend" field.
func CoverageTrendLT(v string) predicate.CodeCov {
	return predicate.CodeCov(sql.FieldLT(FieldCoverageTrend, v))
}

// CoverageTrendLTE applies the LTE predicate on the "coverage_trend" field.
func CoverageTrendLTE(v string) predicate.CodeCov {
	return predicate.CodeCov(sql.FieldLTE(FieldCoverageTrend, v))
}

// CoverageTrendContains applies the Contains predicate on the "coverage_trend" field.
func CoverageTrendContains(v string) predicate.CodeCov {
	return predicate.CodeCov(sql.FieldContains(FieldCoverageTrend, v))
}

// CoverageTrendHasPrefix applies the HasPrefix predicate on the "coverage_trend" field.
func CoverageTrendHasPrefix(v string) predicate.CodeCov {
	return predicate.CodeCov(sql.FieldHasPrefix(FieldCoverageTrend, v))
}

// CoverageTrendHasSuffix applies the HasSuffix predicate on the "coverage_trend" field.
func CoverageTrendHasSuffix(v string) predicate.CodeCov {
	return predicate.CodeCov(sql.FieldHasSuffix(FieldCoverageTrend, v))
}

// CoverageTrendIsNil applies the IsNil predicate on the "coverage_trend" field.
func CoverageTrendIsNil() predicate.CodeCov {
	return predicate.CodeCov(sql.FieldIsNull(FieldCoverageTrend))
}

// CoverageTrendNotNil applies the NotNil predicate on the "coverage_trend" field.
func CoverageTrendNotNil() predicate.CodeCov {
	return predicate.CodeCov(sql.FieldNotNull(FieldCoverageTrend))
}

// CoverageTrendEqualFold applies the EqualFold predicate on the "coverage_trend" field.
func CoverageTrendEqualFold(v string) predicate.CodeCov {
	return predicate.CodeCov(sql.FieldEqualFold(FieldCoverageTrend, v))
}

// CoverageTrendContainsFold applies the ContainsFold predicate on the "coverage_trend" field.
func CoverageTrendContainsFold(v string) predicate.CodeCov {
	return predicate.CodeCov(sql.FieldContainsFold(FieldCoverageTrend, v))
}

// HasCodecov applies the HasEdge predicate on the "codecov" edge.
func HasCodecov() predicate.CodeCov {
	return predicate.CodeCov(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CodecovTable, CodecovColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCodecovWith applies the HasEdge predicate on the "codecov" edge with a given conditions (other predicates).
func HasCodecovWith(preds ...predicate.Repository) predicate.CodeCov {
	return predicate.CodeCov(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CodecovInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CodecovTable, CodecovColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.CodeCov) predicate.CodeCov {
	return predicate.CodeCov(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CodeCov) predicate.CodeCov {
	return predicate.CodeCov(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.CodeCov) predicate.CodeCov {
	return predicate.CodeCov(func(s *sql.Selector) {
		p(s.Not())
	})
}