// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testServerStatsMemberPeriods(t *testing.T) {
	t.Parallel()

	query := ServerStatsMemberPeriods()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testServerStatsMemberPeriodsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServerStatsMemberPeriod{}
	if err = randomize.Struct(seed, o, serverStatsMemberPeriodDBTypes, true, serverStatsMemberPeriodColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServerStatsMemberPeriod struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ServerStatsMemberPeriods().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testServerStatsMemberPeriodsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServerStatsMemberPeriod{}
	if err = randomize.Struct(seed, o, serverStatsMemberPeriodDBTypes, true, serverStatsMemberPeriodColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServerStatsMemberPeriod struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := ServerStatsMemberPeriods().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ServerStatsMemberPeriods().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testServerStatsMemberPeriodsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServerStatsMemberPeriod{}
	if err = randomize.Struct(seed, o, serverStatsMemberPeriodDBTypes, true, serverStatsMemberPeriodColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServerStatsMemberPeriod struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ServerStatsMemberPeriodSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ServerStatsMemberPeriods().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testServerStatsMemberPeriodsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServerStatsMemberPeriod{}
	if err = randomize.Struct(seed, o, serverStatsMemberPeriodDBTypes, true, serverStatsMemberPeriodColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServerStatsMemberPeriod struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ServerStatsMemberPeriodExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if ServerStatsMemberPeriod exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ServerStatsMemberPeriodExists to return true, but got false.")
	}
}

func testServerStatsMemberPeriodsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServerStatsMemberPeriod{}
	if err = randomize.Struct(seed, o, serverStatsMemberPeriodDBTypes, true, serverStatsMemberPeriodColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServerStatsMemberPeriod struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	serverStatsMemberPeriodFound, err := FindServerStatsMemberPeriod(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if serverStatsMemberPeriodFound == nil {
		t.Error("want a record, got nil")
	}
}

func testServerStatsMemberPeriodsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServerStatsMemberPeriod{}
	if err = randomize.Struct(seed, o, serverStatsMemberPeriodDBTypes, true, serverStatsMemberPeriodColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServerStatsMemberPeriod struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = ServerStatsMemberPeriods().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testServerStatsMemberPeriodsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServerStatsMemberPeriod{}
	if err = randomize.Struct(seed, o, serverStatsMemberPeriodDBTypes, true, serverStatsMemberPeriodColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServerStatsMemberPeriod struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := ServerStatsMemberPeriods().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testServerStatsMemberPeriodsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	serverStatsMemberPeriodOne := &ServerStatsMemberPeriod{}
	serverStatsMemberPeriodTwo := &ServerStatsMemberPeriod{}
	if err = randomize.Struct(seed, serverStatsMemberPeriodOne, serverStatsMemberPeriodDBTypes, false, serverStatsMemberPeriodColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServerStatsMemberPeriod struct: %s", err)
	}
	if err = randomize.Struct(seed, serverStatsMemberPeriodTwo, serverStatsMemberPeriodDBTypes, false, serverStatsMemberPeriodColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServerStatsMemberPeriod struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = serverStatsMemberPeriodOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = serverStatsMemberPeriodTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ServerStatsMemberPeriods().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testServerStatsMemberPeriodsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	serverStatsMemberPeriodOne := &ServerStatsMemberPeriod{}
	serverStatsMemberPeriodTwo := &ServerStatsMemberPeriod{}
	if err = randomize.Struct(seed, serverStatsMemberPeriodOne, serverStatsMemberPeriodDBTypes, false, serverStatsMemberPeriodColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServerStatsMemberPeriod struct: %s", err)
	}
	if err = randomize.Struct(seed, serverStatsMemberPeriodTwo, serverStatsMemberPeriodDBTypes, false, serverStatsMemberPeriodColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServerStatsMemberPeriod struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = serverStatsMemberPeriodOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = serverStatsMemberPeriodTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServerStatsMemberPeriods().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testServerStatsMemberPeriodsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServerStatsMemberPeriod{}
	if err = randomize.Struct(seed, o, serverStatsMemberPeriodDBTypes, true, serverStatsMemberPeriodColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServerStatsMemberPeriod struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServerStatsMemberPeriods().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testServerStatsMemberPeriodsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServerStatsMemberPeriod{}
	if err = randomize.Struct(seed, o, serverStatsMemberPeriodDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ServerStatsMemberPeriod struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(serverStatsMemberPeriodColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := ServerStatsMemberPeriods().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testServerStatsMemberPeriodsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServerStatsMemberPeriod{}
	if err = randomize.Struct(seed, o, serverStatsMemberPeriodDBTypes, true, serverStatsMemberPeriodColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServerStatsMemberPeriod struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testServerStatsMemberPeriodsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServerStatsMemberPeriod{}
	if err = randomize.Struct(seed, o, serverStatsMemberPeriodDBTypes, true, serverStatsMemberPeriodColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServerStatsMemberPeriod struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ServerStatsMemberPeriodSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testServerStatsMemberPeriodsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServerStatsMemberPeriod{}
	if err = randomize.Struct(seed, o, serverStatsMemberPeriodDBTypes, true, serverStatsMemberPeriodColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServerStatsMemberPeriod struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ServerStatsMemberPeriods().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	serverStatsMemberPeriodDBTypes = map[string]string{`ID`: `bigint`, `GuildID`: `bigint`, `CreatedAt`: `timestamp with time zone`, `NumMembers`: `bigint`, `Joins`: `bigint`, `Leaves`: `bigint`, `MaxOnline`: `bigint`}
	_                              = bytes.MinRead
)

func testServerStatsMemberPeriodsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(serverStatsMemberPeriodPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(serverStatsMemberPeriodColumns) == len(serverStatsMemberPeriodPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ServerStatsMemberPeriod{}
	if err = randomize.Struct(seed, o, serverStatsMemberPeriodDBTypes, true, serverStatsMemberPeriodColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServerStatsMemberPeriod struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServerStatsMemberPeriods().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, serverStatsMemberPeriodDBTypes, true, serverStatsMemberPeriodPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ServerStatsMemberPeriod struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testServerStatsMemberPeriodsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(serverStatsMemberPeriodColumns) == len(serverStatsMemberPeriodPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ServerStatsMemberPeriod{}
	if err = randomize.Struct(seed, o, serverStatsMemberPeriodDBTypes, true, serverStatsMemberPeriodColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServerStatsMemberPeriod struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServerStatsMemberPeriods().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, serverStatsMemberPeriodDBTypes, true, serverStatsMemberPeriodPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ServerStatsMemberPeriod struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(serverStatsMemberPeriodColumns, serverStatsMemberPeriodPrimaryKeyColumns) {
		fields = serverStatsMemberPeriodColumns
	} else {
		fields = strmangle.SetComplement(
			serverStatsMemberPeriodColumns,
			serverStatsMemberPeriodPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := ServerStatsMemberPeriodSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testServerStatsMemberPeriodsUpsert(t *testing.T) {
	t.Parallel()

	if len(serverStatsMemberPeriodColumns) == len(serverStatsMemberPeriodPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := ServerStatsMemberPeriod{}
	if err = randomize.Struct(seed, &o, serverStatsMemberPeriodDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ServerStatsMemberPeriod struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ServerStatsMemberPeriod: %s", err)
	}

	count, err := ServerStatsMemberPeriods().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, serverStatsMemberPeriodDBTypes, false, serverStatsMemberPeriodPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ServerStatsMemberPeriod struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ServerStatsMemberPeriod: %s", err)
	}

	count, err = ServerStatsMemberPeriods().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
