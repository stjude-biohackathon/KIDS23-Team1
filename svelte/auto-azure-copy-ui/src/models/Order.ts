exportinterface Order {
  id: number;
  order_name: string;
  user_refer: number;
  sample_refer: number;
  analysis_suite_refer: number;
  analysis_run_refer: number;
  analysis_refer: number;
  sample: Sample;
  analysis_suite: AnalysisSuite;
  analysis_run: AnalysisRun;
  analysis: Analysis;
}[D[A