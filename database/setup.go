package database

import "fmt"

func (db *DB) CreateTables() {
	tableStatements := []string{
		CreateInstructorTableSQL,
		CreateDepartmentTableSQL,
		CreateStudentsTableSQL,
		CreateCoursesTableSQL,
		CreateEnrollmentsTable,
	}

	for _, stmt := range tableStatements {
		if _, err := db.Exec(stmt); err != nil {
			fmt.Printf("failed to execute statement: %s, error: %v\n", stmt, err)
		}
	}
}
