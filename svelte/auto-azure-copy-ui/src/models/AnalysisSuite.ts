import type { Analysis } from "./Analysis";

export interface AnalysisSuite {
  id: number;
  suite_name: string;
  analyses: Analysis[];
}
