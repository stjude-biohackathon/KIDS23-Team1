import type { SampleFile } from "./SampleFile";

export interface Sample {
  id: number;
  sample_name: string;
  sample_type: string;
  sample_state: string;
  sample_files: SampleFile[];
}