package db

import "context"

const getAnalysisSuiteByName = `-- name: GetAnalysisSuiteByName :one
SELECT id, suite_name FROM autocopy.analysis_suite
where suite_name = $1 LIMIT 1;
`

// GetAnalysisSuiteByName gets an analysis suite by name
func (q *Queries) GetAnalysisSuiteByName(ctx context.Context, name string) (AnalysisSuite, error) {
	row := q.db.QueryRowContext(ctx, getAnalysisSuiteByName, name)
	var anlsSuite AnalysisSuite
	err := row.Scan(
		&anlsSuite.ID,
		&anlsSuite.SuiteName,
	)
	return anlsSuite, err
}

const getAnlsTypesByAnlsSuiteId = `-- GetAnlsTypesByAnlsSuiteId :many
SELECT id, anls_type_name, anls_suite_id FROM autocopy.analysis_type
where anls_suide_id = $1
`

// GetAnlsTypesByAnlsSuiteId returns all analysis types for a given suite id
func (q *Queries) GetAnlsTypesByAnlsSuiteId(ctx context.Context, id int) ([]AnalysisType, error) {
	rows, err := q.db.QueryContext(ctx, getAnlsTypesByAnlsSuiteId, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var anlsTypes []AnalysisType
	for rows.Next() {
		var anlsType AnalysisType
		if err = rows.Scan(
			&anlsType.ID,
			&anlsType.AnalysisTypeName,
			&anlsType.AnalysisSuiteID,
		); err != nil {
			return nil, err
		}
		anlsTypes = append(anlsTypes, anlsType)
	}
	if err = rows.Close(); err != nil {
		return nil, err
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return anlsTypes, nil
}
