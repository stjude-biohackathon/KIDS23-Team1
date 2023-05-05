export interface AnalysisRun {
  id: number;
  analysis_run_id: number;
  analysis_run_status: string;
  analysis_run_type: string;
  analysis_run_output: AnalysisRunOutput[];
}