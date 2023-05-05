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

const getAnalysisRunByID = `-- name GetAnalysisRunByID :one
SELECT id, name, anls_id, sample_id, status, created_at
  FROM autocopy.analysis_run
  WHERE id = $1 
`

// GetAnalysisRunByID gets an analysis run by id
func (q *Queries) GetAnalysisRunByID(ctx context.Context, id int) (AnalysisRun, error) {
	row := q.db.QueryRowContext(ctx, getAnalysisRunByID, id)
	var anlsRun AnalysisRun
	err := row.Scan(
		&anlsRun.ID,
		&anlsRun.Name,
		&anlsRun.AnalysisId,
		&anlsRun.OrderSampleID,
		&anlsRun.Status,
		&anlsRun.CreatedAt,
	)
	return anlsRun, err
}

const getAllAnalysisRunByOrderSampleID = `-- name: GetAllAnalysisRunBySampleID :many
SELECT id, name, anls_id, order_sample_id, status, created_at
  FROM autocopy.analysis_run
  WHERE order_sample_id = $1
`

// GetAllAnalysisRunBySampleID returns all analysis runs for a given sample id
func (q *Queries) GetAllAnalysisRunBySampleID(ctx context.Context, id int) ([]AnalysisRun, error) {
	rows, err := q.db.QueryContext(ctx, getAllAnalysisRunByOrderSampleID, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var anlsRuns []AnalysisRun
	for rows.Next() {
		var anlsRun AnalysisRun
		if err = rows.Scan(
			&anlsRun.ID,
			&anlsRun.Name,
			&anlsRun.AnalysisId,
			&anlsRun.OrderSampleID,
			&anlsRun.Status,
			&anlsRun.CreatedAt,
		); err != nil {
			return nil, err
		}
		anlsRuns = append(anlsRuns, anlsRun)
	}
	if err = rows.Close(); err != nil {
		return nil, err
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return anlsRuns, nil
}

const getAnalysisRunOutputsByAnlsRunID = `-- name GetAnalysisRunOutputsByAnlsRunID :many
SELECT id, name, anls_run_id, filesystem_path, location_encoding, created_at
  FROM autocopy.analysis_run_outputs
  WHERE anls_run_id = $1
`

func (q *Queries) GetAnalysisRunOutputsByAnlsRunID(ctx context.Context, id int) ([]AnalysisRunOutputs, error) {
	rows, err := q.db.QueryContext(ctx, getAnalysisRunOutputsByAnlsRunID, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var anlsRunOutputs []AnalysisRunOutputs
	for rows.Next() {
		var anlsRunOutput AnalysisRunOutputs
		if err = rows.Scan(
			&anlsRunOutput.ID,
			&anlsRunOutput.Name,
			&anlsRunOutput.AnalysisRunID,
			&anlsRunOutput.FilesystemPath,
			&anlsRunOutput.LocationEncoding,
			&anlsRunOutput.CreatedAt,
		); err != nil {
			return nil, err
		}
		anlsRunOutputs = append(anlsRunOutputs, anlsRunOutput)
	}
	if err = rows.Close(); err != nil {
		return nil, err
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return anlsRunOutputs, nil
}
