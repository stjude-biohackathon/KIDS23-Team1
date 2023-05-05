// all post and get calls here as methods
// TODO - add pagination where needed for endpoints, may need pagination model
// TODO - replace mockGetSample with actual endpoint call
async function mockGetSamples(): Promise<any> {
    return {
      response: {
        sample: {
          sample_files: [
            {
              ID: 1,
              CreatedAt: '2023-05-05T00:00:00-05:00',
              UpdatedAt: '2023-05-05T00:00:00-05:00',
              DeletedAt: null,
              File_name: 'SJACT001D_bam',
              File_size: 1111111,
              File_location: 'hot',
              File_state: 'hot',
              SampleRefer: 1,
            },
            {
              ID: 2,
              CreatedAt: '2023-05-05T00:00:00-05:00',
              UpdatedAt: '2023-05-05T00:00:00-05:00',
              DeletedAt: null,
              File_name: 'SJACT001D_vcf',
              File_size: 222,
              File_location: 'hot',
              File_state: 'hot',
              SampleRefer: 1,
            },
          ],
          sample_id: 1,
          sample_name: 'SJACT001_D',
          sample_state: 'Ready',
          sample_type: 'Autopsy',
        },
      },
    };
}
