import type { Analysis } from "./Analysis";
import type { AnalysisRun } from "./AnalysisRun";
import type { AnalysisSuite } from "./AnalysisSuite";
import type { Sample } from "./Sample";

export interface Order {
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
}