package main

type Record struct {
	ID string
	Name string
}

func ReadWhateverFromDatabase(db *DB) ([]Record, error) {
	stmt := `
select id, name from my_table_name
`
	rows, err := db.Db.Query(stmt)
	if err != nil {
		return nil, err
	}
	var ret []Record
	for rows.Next() {
		var r Record
		if err := rows.Scan(
			&r.ID, &r.Name,
			); err != nil {
			return nil, err
		}
		ret = append(ret, r)
	}
	return ret, nil
}