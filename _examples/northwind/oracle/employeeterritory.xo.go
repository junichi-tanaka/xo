package oracle

// Code generated by xo. DO NOT EDIT.

import (
	"context"
)

// EmployeeTerritory represents a row from 'northwind.employee_territories'.
type EmployeeTerritory struct {
	EmployeeID  int    `json:"employee_id"`  // employee_id
	TerritoryID string `json:"territory_id"` // territory_id
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the EmployeeTerritory exists in the database.
func (et *EmployeeTerritory) Exists() bool {
	return et._exists
}

// Deleted returns true when the EmployeeTerritory has been marked for deletion from
// the database.
func (et *EmployeeTerritory) Deleted() bool {
	return et._deleted
}

// Insert inserts the EmployeeTerritory to the database.
func (et *EmployeeTerritory) Insert(ctx context.Context, db DB) error {
	switch {
	case et._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case et._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	const sqlstr = `INSERT INTO northwind.employee_territories (` +
		`employee_id, territory_id` +
		`) VALUES (` +
		`:1, :2` +
		`)`
	// run
	logf(sqlstr, et.EmployeeID, et.TerritoryID)
	if _, err := db.ExecContext(ctx, sqlstr, et.EmployeeID, et.TerritoryID); err != nil {
		return logerror(err)
	}
	// set exists
	et._exists = true
	return nil
}

// ------ NOTE: Update statements omitted due to lack of fields other than primary key ------

// Delete deletes the EmployeeTerritory from the database.
func (et *EmployeeTerritory) Delete(ctx context.Context, db DB) error {
	switch {
	case !et._exists: // doesn't exist
		return nil
	case et._deleted: // deleted
		return nil
	}
	// delete with composite primary key
	const sqlstr = `DELETE FROM northwind.employee_territories ` +
		`WHERE employee_id = :1 AND territory_id = :2`
	// run
	logf(sqlstr, et.EmployeeID, et.TerritoryID)
	if _, err := db.ExecContext(ctx, sqlstr, et.EmployeeID, et.TerritoryID); err != nil {
		return logerror(err)
	}
	// set deleted
	et._deleted = true
	return nil
}

// EmployeeTerritoryByEmployeeIDTerritoryID retrieves a row from 'northwind.employee_territories' as a EmployeeTerritory.
//
// Generated from index 'employee_territories_pkey'.
func EmployeeTerritoryByEmployeeIDTerritoryID(ctx context.Context, db DB, employeeID int, territoryID string) (*EmployeeTerritory, error) {
	// query
	const sqlstr = `SELECT ` +
		`employee_id, territory_id ` +
		`FROM northwind.employee_territories ` +
		`WHERE employee_id = :1 AND territory_id = :2`
	// run
	logf(sqlstr, employeeID, territoryID)
	et := EmployeeTerritory{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, employeeID, territoryID).Scan(&et.EmployeeID, &et.TerritoryID); err != nil {
		return nil, logerror(err)
	}
	return &et, nil
}

// Employee returns the Employee associated with the EmployeeTerritory's EmployeeID (employee_id).
//
// Generated from foreign key 'employee_territories_employee_id_fkey'.
func (et *EmployeeTerritory) Employee(ctx context.Context, db DB) (*Employee, error) {
	return EmployeeByEmployeeID(ctx, db, et.EmployeeID)
}

// Territory returns the Territory associated with the EmployeeTerritory's TerritoryID (territory_id).
//
// Generated from foreign key 'employee_territories_territory_id_fkey'.
func (et *EmployeeTerritory) Territory(ctx context.Context, db DB) (*Territory, error) {
	return TerritoryByTerritoryID(ctx, db, et.TerritoryID)
}
