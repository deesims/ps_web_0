// This file is generated by SQLBoiler (https://github.com/vattle/sqlboiler)
// and is meant to be re-generated in place and/or deleted at any time.
// DO NOT EDIT

package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testWorkReviews(t *testing.T) {
	t.Parallel()

	query := WorkReviews(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testWorkReviewsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	workReview := &WorkReview{}
	if err = randomize.Struct(seed, workReview, workReviewDBTypes, true); err != nil {
		t.Errorf("Unable to randomize WorkReview struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = workReview.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = workReview.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := WorkReviews(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testWorkReviewsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	workReview := &WorkReview{}
	if err = randomize.Struct(seed, workReview, workReviewDBTypes, true); err != nil {
		t.Errorf("Unable to randomize WorkReview struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = workReview.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = WorkReviews(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := WorkReviews(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testWorkReviewsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	workReview := &WorkReview{}
	if err = randomize.Struct(seed, workReview, workReviewDBTypes, true); err != nil {
		t.Errorf("Unable to randomize WorkReview struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = workReview.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := WorkReviewSlice{workReview}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := WorkReviews(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testWorkReviewsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	workReview := &WorkReview{}
	if err = randomize.Struct(seed, workReview, workReviewDBTypes, true, workReviewColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WorkReview struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = workReview.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := WorkReviewExists(tx, workReview.WorksForjobjobID, workReview.WorksForUserID)
	if err != nil {
		t.Errorf("Unable to check if WorkReview exists: %s", err)
	}
	if !e {
		t.Errorf("Expected WorkReviewExistsG to return true, but got false.")
	}
}
func testWorkReviewsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	workReview := &WorkReview{}
	if err = randomize.Struct(seed, workReview, workReviewDBTypes, true, workReviewColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WorkReview struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = workReview.Insert(tx); err != nil {
		t.Error(err)
	}

	workReviewFound, err := FindWorkReview(tx, workReview.WorksForjobjobID, workReview.WorksForUserID)
	if err != nil {
		t.Error(err)
	}

	if workReviewFound == nil {
		t.Error("want a record, got nil")
	}
}
func testWorkReviewsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	workReview := &WorkReview{}
	if err = randomize.Struct(seed, workReview, workReviewDBTypes, true, workReviewColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WorkReview struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = workReview.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = WorkReviews(tx).Bind(workReview); err != nil {
		t.Error(err)
	}
}

func testWorkReviewsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	workReview := &WorkReview{}
	if err = randomize.Struct(seed, workReview, workReviewDBTypes, true, workReviewColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WorkReview struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = workReview.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := WorkReviews(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testWorkReviewsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	workReviewOne := &WorkReview{}
	workReviewTwo := &WorkReview{}
	if err = randomize.Struct(seed, workReviewOne, workReviewDBTypes, false, workReviewColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WorkReview struct: %s", err)
	}
	if err = randomize.Struct(seed, workReviewTwo, workReviewDBTypes, false, workReviewColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WorkReview struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = workReviewOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = workReviewTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := WorkReviews(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testWorkReviewsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	workReviewOne := &WorkReview{}
	workReviewTwo := &WorkReview{}
	if err = randomize.Struct(seed, workReviewOne, workReviewDBTypes, false, workReviewColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WorkReview struct: %s", err)
	}
	if err = randomize.Struct(seed, workReviewTwo, workReviewDBTypes, false, workReviewColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WorkReview struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = workReviewOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = workReviewTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := WorkReviews(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func workReviewBeforeInsertHook(e boil.Executor, o *WorkReview) error {
	*o = WorkReview{}
	return nil
}

func workReviewAfterInsertHook(e boil.Executor, o *WorkReview) error {
	*o = WorkReview{}
	return nil
}

func workReviewAfterSelectHook(e boil.Executor, o *WorkReview) error {
	*o = WorkReview{}
	return nil
}

func workReviewBeforeUpdateHook(e boil.Executor, o *WorkReview) error {
	*o = WorkReview{}
	return nil
}

func workReviewAfterUpdateHook(e boil.Executor, o *WorkReview) error {
	*o = WorkReview{}
	return nil
}

func workReviewBeforeDeleteHook(e boil.Executor, o *WorkReview) error {
	*o = WorkReview{}
	return nil
}

func workReviewAfterDeleteHook(e boil.Executor, o *WorkReview) error {
	*o = WorkReview{}
	return nil
}

func workReviewBeforeUpsertHook(e boil.Executor, o *WorkReview) error {
	*o = WorkReview{}
	return nil
}

func workReviewAfterUpsertHook(e boil.Executor, o *WorkReview) error {
	*o = WorkReview{}
	return nil
}

func testWorkReviewsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &WorkReview{}
	o := &WorkReview{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, workReviewDBTypes, false); err != nil {
		t.Errorf("Unable to randomize WorkReview object: %s", err)
	}

	AddWorkReviewHook(boil.BeforeInsertHook, workReviewBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	workReviewBeforeInsertHooks = []WorkReviewHook{}

	AddWorkReviewHook(boil.AfterInsertHook, workReviewAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	workReviewAfterInsertHooks = []WorkReviewHook{}

	AddWorkReviewHook(boil.AfterSelectHook, workReviewAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	workReviewAfterSelectHooks = []WorkReviewHook{}

	AddWorkReviewHook(boil.BeforeUpdateHook, workReviewBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	workReviewBeforeUpdateHooks = []WorkReviewHook{}

	AddWorkReviewHook(boil.AfterUpdateHook, workReviewAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	workReviewAfterUpdateHooks = []WorkReviewHook{}

	AddWorkReviewHook(boil.BeforeDeleteHook, workReviewBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	workReviewBeforeDeleteHooks = []WorkReviewHook{}

	AddWorkReviewHook(boil.AfterDeleteHook, workReviewAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	workReviewAfterDeleteHooks = []WorkReviewHook{}

	AddWorkReviewHook(boil.BeforeUpsertHook, workReviewBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	workReviewBeforeUpsertHooks = []WorkReviewHook{}

	AddWorkReviewHook(boil.AfterUpsertHook, workReviewAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	workReviewAfterUpsertHooks = []WorkReviewHook{}
}
func testWorkReviewsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	workReview := &WorkReview{}
	if err = randomize.Struct(seed, workReview, workReviewDBTypes, true, workReviewColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WorkReview struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = workReview.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := WorkReviews(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testWorkReviewsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	workReview := &WorkReview{}
	if err = randomize.Struct(seed, workReview, workReviewDBTypes, true); err != nil {
		t.Errorf("Unable to randomize WorkReview struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = workReview.Insert(tx, workReviewColumns...); err != nil {
		t.Error(err)
	}

	count, err := WorkReviews(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testWorkReviewsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	workReview := &WorkReview{}
	if err = randomize.Struct(seed, workReview, workReviewDBTypes, true, workReviewColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WorkReview struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = workReview.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = workReview.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testWorkReviewsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	workReview := &WorkReview{}
	if err = randomize.Struct(seed, workReview, workReviewDBTypes, true, workReviewColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WorkReview struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = workReview.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := WorkReviewSlice{workReview}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testWorkReviewsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	workReview := &WorkReview{}
	if err = randomize.Struct(seed, workReview, workReviewDBTypes, true, workReviewColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WorkReview struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = workReview.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := WorkReviews(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	workReviewDBTypes = map[string]string{`Comments`: `character varying`, `Rating`: `integer`, `WorksForUserID`: `numeric`, `WorksForjobjobID`: `numeric`}
	_                 = bytes.MinRead
)

func testWorkReviewsUpdate(t *testing.T) {
	t.Parallel()

	if len(workReviewColumns) == len(workReviewPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	workReview := &WorkReview{}
	if err = randomize.Struct(seed, workReview, workReviewDBTypes, true); err != nil {
		t.Errorf("Unable to randomize WorkReview struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = workReview.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := WorkReviews(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, workReview, workReviewDBTypes, true, workReviewColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WorkReview struct: %s", err)
	}

	if err = workReview.Update(tx); err != nil {
		t.Error(err)
	}
}

func testWorkReviewsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(workReviewColumns) == len(workReviewPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	workReview := &WorkReview{}
	if err = randomize.Struct(seed, workReview, workReviewDBTypes, true); err != nil {
		t.Errorf("Unable to randomize WorkReview struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = workReview.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := WorkReviews(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, workReview, workReviewDBTypes, true, workReviewPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize WorkReview struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(workReviewColumns, workReviewPrimaryKeyColumns) {
		fields = workReviewColumns
	} else {
		fields = strmangle.SetComplement(
			workReviewColumns,
			workReviewPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(workReview))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := WorkReviewSlice{workReview}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testWorkReviewsUpsert(t *testing.T) {
	t.Parallel()

	if len(workReviewColumns) == len(workReviewPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	workReview := WorkReview{}
	if err = randomize.Struct(seed, &workReview, workReviewDBTypes, true); err != nil {
		t.Errorf("Unable to randomize WorkReview struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = workReview.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert WorkReview: %s", err)
	}

	count, err := WorkReviews(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &workReview, workReviewDBTypes, false, workReviewPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize WorkReview struct: %s", err)
	}

	if err = workReview.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert WorkReview: %s", err)
	}

	count, err = WorkReviews(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
