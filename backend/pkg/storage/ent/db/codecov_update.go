// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/redhat-appstudio/quality-studio/pkg/storage/ent/db/codecov"
	"github.com/redhat-appstudio/quality-studio/pkg/storage/ent/db/predicate"
	"github.com/redhat-appstudio/quality-studio/pkg/storage/ent/db/repository"
)

// CodeCovUpdate is the builder for updating CodeCov entities.
type CodeCovUpdate struct {
	config
	hooks    []Hook
	mutation *CodeCovMutation
}

// Where appends a list predicates to the CodeCovUpdate builder.
func (ccu *CodeCovUpdate) Where(ps ...predicate.CodeCov) *CodeCovUpdate {
	ccu.mutation.Where(ps...)
	return ccu
}

// SetRepositoryName sets the "repository_name" field.
func (ccu *CodeCovUpdate) SetRepositoryName(s string) *CodeCovUpdate {
	ccu.mutation.SetRepositoryName(s)
	return ccu
}

// SetGitOrganization sets the "git_organization" field.
func (ccu *CodeCovUpdate) SetGitOrganization(s string) *CodeCovUpdate {
	ccu.mutation.SetGitOrganization(s)
	return ccu
}

// SetCoveragePercentage sets the "coverage_percentage" field.
func (ccu *CodeCovUpdate) SetCoveragePercentage(f float64) *CodeCovUpdate {
	ccu.mutation.ResetCoveragePercentage()
	ccu.mutation.SetCoveragePercentage(f)
	return ccu
}

// AddCoveragePercentage adds f to the "coverage_percentage" field.
func (ccu *CodeCovUpdate) AddCoveragePercentage(f float64) *CodeCovUpdate {
	ccu.mutation.AddCoveragePercentage(f)
	return ccu
}

// SetAverageRetestsToMerge sets the "average_retests_to_merge" field.
func (ccu *CodeCovUpdate) SetAverageRetestsToMerge(f float64) *CodeCovUpdate {
	ccu.mutation.ResetAverageRetestsToMerge()
	ccu.mutation.SetAverageRetestsToMerge(f)
	return ccu
}

// SetNillableAverageRetestsToMerge sets the "average_retests_to_merge" field if the given value is not nil.
func (ccu *CodeCovUpdate) SetNillableAverageRetestsToMerge(f *float64) *CodeCovUpdate {
	if f != nil {
		ccu.SetAverageRetestsToMerge(*f)
	}
	return ccu
}

// AddAverageRetestsToMerge adds f to the "average_retests_to_merge" field.
func (ccu *CodeCovUpdate) AddAverageRetestsToMerge(f float64) *CodeCovUpdate {
	ccu.mutation.AddAverageRetestsToMerge(f)
	return ccu
}

// ClearAverageRetestsToMerge clears the value of the "average_retests_to_merge" field.
func (ccu *CodeCovUpdate) ClearAverageRetestsToMerge() *CodeCovUpdate {
	ccu.mutation.ClearAverageRetestsToMerge()
	return ccu
}

// SetCoverageTrend sets the "coverage_trend" field.
func (ccu *CodeCovUpdate) SetCoverageTrend(s string) *CodeCovUpdate {
	ccu.mutation.SetCoverageTrend(s)
	return ccu
}

// SetNillableCoverageTrend sets the "coverage_trend" field if the given value is not nil.
func (ccu *CodeCovUpdate) SetNillableCoverageTrend(s *string) *CodeCovUpdate {
	if s != nil {
		ccu.SetCoverageTrend(*s)
	}
	return ccu
}

// ClearCoverageTrend clears the value of the "coverage_trend" field.
func (ccu *CodeCovUpdate) ClearCoverageTrend() *CodeCovUpdate {
	ccu.mutation.ClearCoverageTrend()
	return ccu
}

// SetCodecovID sets the "codecov" edge to the Repository entity by ID.
func (ccu *CodeCovUpdate) SetCodecovID(id string) *CodeCovUpdate {
	ccu.mutation.SetCodecovID(id)
	return ccu
}

// SetNillableCodecovID sets the "codecov" edge to the Repository entity by ID if the given value is not nil.
func (ccu *CodeCovUpdate) SetNillableCodecovID(id *string) *CodeCovUpdate {
	if id != nil {
		ccu = ccu.SetCodecovID(*id)
	}
	return ccu
}

// SetCodecov sets the "codecov" edge to the Repository entity.
func (ccu *CodeCovUpdate) SetCodecov(r *Repository) *CodeCovUpdate {
	return ccu.SetCodecovID(r.ID)
}

// Mutation returns the CodeCovMutation object of the builder.
func (ccu *CodeCovUpdate) Mutation() *CodeCovMutation {
	return ccu.mutation
}

// ClearCodecov clears the "codecov" edge to the Repository entity.
func (ccu *CodeCovUpdate) ClearCodecov() *CodeCovUpdate {
	ccu.mutation.ClearCodecov()
	return ccu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ccu *CodeCovUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, CodeCovMutation](ctx, ccu.sqlSave, ccu.mutation, ccu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ccu *CodeCovUpdate) SaveX(ctx context.Context) int {
	affected, err := ccu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ccu *CodeCovUpdate) Exec(ctx context.Context) error {
	_, err := ccu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccu *CodeCovUpdate) ExecX(ctx context.Context) {
	if err := ccu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ccu *CodeCovUpdate) check() error {
	if v, ok := ccu.mutation.RepositoryName(); ok {
		if err := codecov.RepositoryNameValidator(v); err != nil {
			return &ValidationError{Name: "repository_name", err: fmt.Errorf(`db: validator failed for field "CodeCov.repository_name": %w`, err)}
		}
	}
	return nil
}

func (ccu *CodeCovUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ccu.check(); err != nil {
		return n, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   codecov.Table,
			Columns: codecov.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: codecov.FieldID,
			},
		},
	}
	if ps := ccu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ccu.mutation.RepositoryName(); ok {
		_spec.SetField(codecov.FieldRepositoryName, field.TypeString, value)
	}
	if value, ok := ccu.mutation.GitOrganization(); ok {
		_spec.SetField(codecov.FieldGitOrganization, field.TypeString, value)
	}
	if value, ok := ccu.mutation.CoveragePercentage(); ok {
		_spec.SetField(codecov.FieldCoveragePercentage, field.TypeFloat64, value)
	}
	if value, ok := ccu.mutation.AddedCoveragePercentage(); ok {
		_spec.AddField(codecov.FieldCoveragePercentage, field.TypeFloat64, value)
	}
	if value, ok := ccu.mutation.AverageRetestsToMerge(); ok {
		_spec.SetField(codecov.FieldAverageRetestsToMerge, field.TypeFloat64, value)
	}
	if value, ok := ccu.mutation.AddedAverageRetestsToMerge(); ok {
		_spec.AddField(codecov.FieldAverageRetestsToMerge, field.TypeFloat64, value)
	}
	if ccu.mutation.AverageRetestsToMergeCleared() {
		_spec.ClearField(codecov.FieldAverageRetestsToMerge, field.TypeFloat64)
	}
	if value, ok := ccu.mutation.CoverageTrend(); ok {
		_spec.SetField(codecov.FieldCoverageTrend, field.TypeString, value)
	}
	if ccu.mutation.CoverageTrendCleared() {
		_spec.ClearField(codecov.FieldCoverageTrend, field.TypeString)
	}
	if ccu.mutation.CodecovCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   codecov.CodecovTable,
			Columns: []string{codecov.CodecovColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: repository.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ccu.mutation.CodecovIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   codecov.CodecovTable,
			Columns: []string{codecov.CodecovColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: repository.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ccu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{codecov.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ccu.mutation.done = true
	return n, nil
}

// CodeCovUpdateOne is the builder for updating a single CodeCov entity.
type CodeCovUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CodeCovMutation
}

// SetRepositoryName sets the "repository_name" field.
func (ccuo *CodeCovUpdateOne) SetRepositoryName(s string) *CodeCovUpdateOne {
	ccuo.mutation.SetRepositoryName(s)
	return ccuo
}

// SetGitOrganization sets the "git_organization" field.
func (ccuo *CodeCovUpdateOne) SetGitOrganization(s string) *CodeCovUpdateOne {
	ccuo.mutation.SetGitOrganization(s)
	return ccuo
}

// SetCoveragePercentage sets the "coverage_percentage" field.
func (ccuo *CodeCovUpdateOne) SetCoveragePercentage(f float64) *CodeCovUpdateOne {
	ccuo.mutation.ResetCoveragePercentage()
	ccuo.mutation.SetCoveragePercentage(f)
	return ccuo
}

// AddCoveragePercentage adds f to the "coverage_percentage" field.
func (ccuo *CodeCovUpdateOne) AddCoveragePercentage(f float64) *CodeCovUpdateOne {
	ccuo.mutation.AddCoveragePercentage(f)
	return ccuo
}

// SetAverageRetestsToMerge sets the "average_retests_to_merge" field.
func (ccuo *CodeCovUpdateOne) SetAverageRetestsToMerge(f float64) *CodeCovUpdateOne {
	ccuo.mutation.ResetAverageRetestsToMerge()
	ccuo.mutation.SetAverageRetestsToMerge(f)
	return ccuo
}

// SetNillableAverageRetestsToMerge sets the "average_retests_to_merge" field if the given value is not nil.
func (ccuo *CodeCovUpdateOne) SetNillableAverageRetestsToMerge(f *float64) *CodeCovUpdateOne {
	if f != nil {
		ccuo.SetAverageRetestsToMerge(*f)
	}
	return ccuo
}

// AddAverageRetestsToMerge adds f to the "average_retests_to_merge" field.
func (ccuo *CodeCovUpdateOne) AddAverageRetestsToMerge(f float64) *CodeCovUpdateOne {
	ccuo.mutation.AddAverageRetestsToMerge(f)
	return ccuo
}

// ClearAverageRetestsToMerge clears the value of the "average_retests_to_merge" field.
func (ccuo *CodeCovUpdateOne) ClearAverageRetestsToMerge() *CodeCovUpdateOne {
	ccuo.mutation.ClearAverageRetestsToMerge()
	return ccuo
}

// SetCoverageTrend sets the "coverage_trend" field.
func (ccuo *CodeCovUpdateOne) SetCoverageTrend(s string) *CodeCovUpdateOne {
	ccuo.mutation.SetCoverageTrend(s)
	return ccuo
}

// SetNillableCoverageTrend sets the "coverage_trend" field if the given value is not nil.
func (ccuo *CodeCovUpdateOne) SetNillableCoverageTrend(s *string) *CodeCovUpdateOne {
	if s != nil {
		ccuo.SetCoverageTrend(*s)
	}
	return ccuo
}

// ClearCoverageTrend clears the value of the "coverage_trend" field.
func (ccuo *CodeCovUpdateOne) ClearCoverageTrend() *CodeCovUpdateOne {
	ccuo.mutation.ClearCoverageTrend()
	return ccuo
}

// SetCodecovID sets the "codecov" edge to the Repository entity by ID.
func (ccuo *CodeCovUpdateOne) SetCodecovID(id string) *CodeCovUpdateOne {
	ccuo.mutation.SetCodecovID(id)
	return ccuo
}

// SetNillableCodecovID sets the "codecov" edge to the Repository entity by ID if the given value is not nil.
func (ccuo *CodeCovUpdateOne) SetNillableCodecovID(id *string) *CodeCovUpdateOne {
	if id != nil {
		ccuo = ccuo.SetCodecovID(*id)
	}
	return ccuo
}

// SetCodecov sets the "codecov" edge to the Repository entity.
func (ccuo *CodeCovUpdateOne) SetCodecov(r *Repository) *CodeCovUpdateOne {
	return ccuo.SetCodecovID(r.ID)
}

// Mutation returns the CodeCovMutation object of the builder.
func (ccuo *CodeCovUpdateOne) Mutation() *CodeCovMutation {
	return ccuo.mutation
}

// ClearCodecov clears the "codecov" edge to the Repository entity.
func (ccuo *CodeCovUpdateOne) ClearCodecov() *CodeCovUpdateOne {
	ccuo.mutation.ClearCodecov()
	return ccuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ccuo *CodeCovUpdateOne) Select(field string, fields ...string) *CodeCovUpdateOne {
	ccuo.fields = append([]string{field}, fields...)
	return ccuo
}

// Save executes the query and returns the updated CodeCov entity.
func (ccuo *CodeCovUpdateOne) Save(ctx context.Context) (*CodeCov, error) {
	return withHooks[*CodeCov, CodeCovMutation](ctx, ccuo.sqlSave, ccuo.mutation, ccuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ccuo *CodeCovUpdateOne) SaveX(ctx context.Context) *CodeCov {
	node, err := ccuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ccuo *CodeCovUpdateOne) Exec(ctx context.Context) error {
	_, err := ccuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccuo *CodeCovUpdateOne) ExecX(ctx context.Context) {
	if err := ccuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ccuo *CodeCovUpdateOne) check() error {
	if v, ok := ccuo.mutation.RepositoryName(); ok {
		if err := codecov.RepositoryNameValidator(v); err != nil {
			return &ValidationError{Name: "repository_name", err: fmt.Errorf(`db: validator failed for field "CodeCov.repository_name": %w`, err)}
		}
	}
	return nil
}

func (ccuo *CodeCovUpdateOne) sqlSave(ctx context.Context) (_node *CodeCov, err error) {
	if err := ccuo.check(); err != nil {
		return _node, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   codecov.Table,
			Columns: codecov.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: codecov.FieldID,
			},
		},
	}
	id, ok := ccuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`db: missing "CodeCov.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ccuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, codecov.FieldID)
		for _, f := range fields {
			if !codecov.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
			}
			if f != codecov.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ccuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ccuo.mutation.RepositoryName(); ok {
		_spec.SetField(codecov.FieldRepositoryName, field.TypeString, value)
	}
	if value, ok := ccuo.mutation.GitOrganization(); ok {
		_spec.SetField(codecov.FieldGitOrganization, field.TypeString, value)
	}
	if value, ok := ccuo.mutation.CoveragePercentage(); ok {
		_spec.SetField(codecov.FieldCoveragePercentage, field.TypeFloat64, value)
	}
	if value, ok := ccuo.mutation.AddedCoveragePercentage(); ok {
		_spec.AddField(codecov.FieldCoveragePercentage, field.TypeFloat64, value)
	}
	if value, ok := ccuo.mutation.AverageRetestsToMerge(); ok {
		_spec.SetField(codecov.FieldAverageRetestsToMerge, field.TypeFloat64, value)
	}
	if value, ok := ccuo.mutation.AddedAverageRetestsToMerge(); ok {
		_spec.AddField(codecov.FieldAverageRetestsToMerge, field.TypeFloat64, value)
	}
	if ccuo.mutation.AverageRetestsToMergeCleared() {
		_spec.ClearField(codecov.FieldAverageRetestsToMerge, field.TypeFloat64)
	}
	if value, ok := ccuo.mutation.CoverageTrend(); ok {
		_spec.SetField(codecov.FieldCoverageTrend, field.TypeString, value)
	}
	if ccuo.mutation.CoverageTrendCleared() {
		_spec.ClearField(codecov.FieldCoverageTrend, field.TypeString)
	}
	if ccuo.mutation.CodecovCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   codecov.CodecovTable,
			Columns: []string{codecov.CodecovColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: repository.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ccuo.mutation.CodecovIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   codecov.CodecovTable,
			Columns: []string{codecov.CodecovColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: repository.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &CodeCov{config: ccuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ccuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{codecov.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ccuo.mutation.done = true
	return _node, nil
}
