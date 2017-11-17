package postgresql_conf

type pg_conf struct {

	file_contents						string

	data_directory                      string
	hba_file                            string
	ident_file                          string
	external_pid_file                   string
	listen_addresses                    string
	port                                string
	max_connections                     string
	superuser_reserved_connections      string
	unix_socket_directories             string
	unix_socket_group                   string
	unix_socket_permissions             string
	bonjour                             string
	bonjour_name                        string
	authentication_timeout              string
	ssl                                 string
	ssl_ciphers                         string
	ssl_prefer_server_ciphers           string
	ssl_ecdh_curve                      string
	ssl_cert_file                       string
	ssl_key_file                        string
	ssl_ca_file                         string
	ssl_crl_file                        string
	password_encryption                 string
	db_user_namespace                   string
	row_security                        string
	krb_server_keyfile                  string
	krb_caseins_users                   string
	tcp_keepalives_idle                 string
	tcp_keepalives_interval             string
	tcp_keepalives_count                string
	shared_buffers                      string
	huge_pages                          string
	temp_buffers                        string
	max_prepared_transactions           string
	work_mem                            string
	maintenance_work_mem                string
	autovacuum_work_mem                 string
	max_stack_depth                     string
	dynamic_shared_memory_type          string
	temp_file_limit                     string
	max_files_per_process               string
	shared_preload_libraries            string
	vacuum_cost_delay                   string
	vacuum_cost_page_hit                string
	vacuum_cost_page_miss               string
	vacuum_cost_page_dirty              string
	vacuum_cost_limit                   string
	bgwriter_delay                      string
	bgwriter_lru_maxpages               string
	bgwriter_lru_multiplier             string
	effective_io_concurrency            string
	max_worker_processes                string
	shared_queues                       string
	shared_queue_size                   string
	wal_level                           string
	fsync                               string
	synchronous_commit                  string
	wal_sync_method                     string
	full_page_writes                    string
	wal_compression                     string
	wal_log_hints                       string
	wal_buffers                         string
	wal_writer_delay                    string
	commit_delay                        string
	commit_siblings                     string
	checkpoint_timeout                  string
	max_wal_size                        string
	min_wal_size                        string
	checkpoint_completion_target        string
	checkpoint_warning                  string
	archive_mode                        string
	archive_command                     string
	archive_timeout                     string
	max_wal_senders                     string
	wal_keep_segments                   string
	wal_sender_timeout                  string
	max_replication_slots               string
	track_commit_timestamp              string
	synchronous_standby_names           string
	vacuum_defer_cleanup_age            string
	hot_standby                         string
	max_standby_archive_delay           string
	max_standby_streaming_delay         string
	wal_receiver_status_interval        string
	hot_standby_feedback                string
	wal_receiver_timeout                string
	wal_retrieve_retry_interval         string
	enable_bitmapscan                   string
	enable_hashagg                      string
	enable_hashjoin                     string
	enable_indexscan                    string
	enable_indexonlyscan                string
	enable_material                     string
	enable_mergejoin                    string
	enable_nestloop                     string
	enable_seqscan                      string
	enable_sort                         string
	enable_tidscan                      string
	seq_page_cost                       string
	random_page_cost                    string
	cpu_tuple_cost                      string
	cpu_index_tuple_cost                string
	cpu_operator_cost                   string
	network_byte_cost                   string
	remote_query_cost                   string
	effective_cache_size                string
	geqo                                string
	geqo_threshold                      string
	geqo_effort                         string
	geqo_pool_size                      string
	geqo_generations                    string
	geqo_selection_bias                 string
	geqo_seed                           string
	default_statistics_target           string
	constraint_exclusion                string
	cursor_tuple_fraction               string
	from_collapse_limit                 string
	join_collapse_limit                 string
	log_destination                     string
	logging_collector                   string
	log_directory                       string
	log_filename                        string
	log_file_mode                       string
	log_truncate_on_rotation            string
	log_rotation_age                    string
	log_rotation_size                   string
	syslog_facility                     string
	syslog_ident                        string
	event_source                        string
	client_min_messages                 string
	log_min_messages                    string
	log_min_error_statement             string
	log_min_duration_statement          string
	debug_print_parse                   string
	debug_print_rewritten               string
	debug_print_plan                    string
	debug_pretty_print                  string
	log_checkpoints                     string
	log_connections                     string
	log_disconnections                  string
	log_duration                        string
	log_error_verbosity                 string
	log_hostname                        string
	log_line_prefix                     string
	log_lock_waits                      string
	log_statement                       string
	log_replication_commands            string
	log_temp_files                      string
	log_timezone                        string
	cluster_name                        string
	update_process_title                string
	track_activities                    string
	track_counts                        string
	track_io_timing                     string
	track_functions                     string
	track_activity_query_size           string
	stats_temp_directory                string
	log_parser_stats                    string
	log_planner_stats                   string
	log_executor_stats                  string
	log_statement_stats                 string
	autovacuum                          string
	log_autovacuum_min_duration         string
	autovacuum_max_workers              string
	autovacuum_naptime                  string
	autovacuum_vacuum_threshold         string
	autovacuum_analyze_threshold        string
	autovacuum_vacuum_scale_factor      string
	autovacuum_analyze_scale_factor     string
	autovacuum_freeze_max_age           string
	autovacuum_multixact_freeze_max_age string
	autovacuum_vacuum_cost_delay        string
	autovacuum_vacuum_cost_limit        string
	search_path                         string
	default_tablespace                  string
	temp_tablespaces                    string
	check_function_bodies               string
	default_transaction_isolation       string
	default_transaction_read_only       string
	default_transaction_deferrable      string
	session_replication_role            string
	statement_timeout                   string
	lock_timeout                        string
	vacuum_freeze_min_age               string
	vacuum_freeze_table_age             string
	vacuum_multixact_freeze_min_age     string
	vacuum_multixact_freeze_table_age   string
	bytea_output                        string
	xmlbinary                           string
	xmloption                           string
	gin_fuzzy_search_limit              string
	gin_pending_list_limit              string
	datestyle                           string
	intervalstyle                       string
	timezone                            string
	timezone_abbreviations              string
	extra_float_digits                  string
	client_encoding                     string
	lc_messages                         string
	lc_monetary                         string
	lc_numeric                          string
	lc_time                             string
	default_text_search_config          string
	dynamic_library_path                string
	local_preload_libraries             string
	session_preload_libraries           string
	deadlock_timeout                    string
	max_locks_per_transaction           string
	max_pred_locks_per_transaction      string
	array_nulls                         string
	backslash_quote                     string
	default_with_oids                   string
	escape_string_warning               string
	lo_compat_privileges                string
	operator_precedence_warning         string
	quote_all_identifiers               string
	sql_inheritance                     string
	standard_conforming_strings         string
	synchronize_seqscans                string
	transform_null_equals               string
	exit_on_error                       string
	restart_after_crash                 string
	persistent_datanode_connections     string
	max_coordinators                    string
	max_datanodes                       string
	pgxc_node_name                      string
	gtm_backup_barrier                  string
	include_dir                         string
	include_if_exists                   string
	include                             string
	max_pool_size                       string
	pool_conn_keepalive                 string
	pool_maintenance_timeout            string
	pooler_port                         string
	gtm_host                            string
	gtm_port                            string
}


func (pc *pg_conf) set_contents(file_contents string){

	pc.file_contents = file_contents
}

func (pc *pg_conf) parse() (errs error) {

	//if len(pc.file_contents) == 0 {
	//	return errors.New("No contents set for config")
	//}

	//re_blank := regexp.MustCompile(`^[ \t]*$`)
	//re_comment := regexp.MustCompile(`^[ \t]*#`)
	//kv_re := regexp.MustCompile(`[ \t]*(?P<key>\S+)[ |\t]*=[ |\t]?(?P<value>\S+)[ |\t]*#?(?:.*)?`)

	//var p pg_conf
	//
	//p["file_contents"] = "a"


}






