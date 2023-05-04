set schema 'autocopy';
-- insert analysis suite
insert into analysis_suite (suite_name) values('test_suite');

-- select analysis_suite.id into suite_id from analysis_suite where suite_name = 'test_suite';
-- insert some anls_types
insert into analysis_type (anls_type_name, anls_suite_id) values('many_mod_intake', 1);
insert into analysis_type (anls_type_name, anls_suite_id) values('many_small_classify', 1);
insert into analysis_type (anls_type_name, anls_suite_id) values('one_big_few_small_mapping', 1);
insert into analysis_type (anls_type_name, anls_suite_id) values('one_moderate_coverage', 1);

