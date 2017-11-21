package config_prasers

import (
	"errors"
	"github.com/mitchellh/mapstructure"
	"regexp"
	"strings"
)

type pg_conf struct {
	File_contents string

	Data_directory                      string
	Hba_file                            string
	Ident_file                          string
	External_pid_file                   string
	Listen_addresses                    string
	Port                                string
	Max_connections                     string
	Superuser_reserved_connections      string
	Unix_socket_directories             string
	Unix_socket_group                   string
	Unix_socket_permissions             string
	Bonjour                             string
	Bonjour_name                        string
	Authentication_timeout              string
	Ssl                                 string
	Ssl_ciphers                         string
	Ssl_prefer_server_ciphers           string
	Ssl_ecdh_curve                      string
	Ssl_cert_file                       string
	Ssl_key_file                        string
	Ssl_ca_file                         string
	Ssl_crl_file                        string
	Password_encryption                 string
	Db_user_namespace                   string
	Row_security                        string
	Krb_server_keyfile                  string
	Krb_caseins_users                   string
	Tcp_keepalives_idle                 string
	Tcp_keepalives_interval             string
	Tcp_keepalives_count                string
	Shared_buffers                      string
	Huge_pages                          string
	Temp_buffers                        string
	Max_prepared_transactions           string
	Work_mem                            string
	Maintenance_work_mem                string
	Autovacuum_work_mem                 string
	Max_stack_depth                     string
	Dynamic_shared_memory_type          string
	Temp_file_limit                     string
	Max_files_per_process               string
	Shared_preload_libraries            string
	Vacuum_cost_delay                   string
	Vacuum_cost_page_hit                string
	Vacuum_cost_page_miss               string
	Vacuum_cost_page_dirty              string
	Vacuum_cost_limit                   string
	Bgwriter_delay                      string
	Bgwriter_lru_maxpages               string
	Bgwriter_lru_multiplier             string
	Effective_io_concurrency            string
	Max_worker_processes                string
	Shared_queues                       string
	Shared_queue_size                   string
	Wal_level                           string
	Fsync                               string
	Synchronous_commit                  string
	Wal_sync_method                     string
	Full_page_writes                    string
	Wal_compression                     string
	Wal_log_hints                       string
	Wal_buffers                         string
	Wal_writer_delay                    string
	Commit_delay                        string
	Commit_siblings                     string
	Checkpoint_timeout                  string
	Max_wal_size                        string
	Min_wal_size                        string
	Checkpoint_completion_target        string
	Checkpoint_warning                  string
	Archive_mode                        string
	Archive_command                     string
	Archive_timeout                     string
	Max_wal_senders                     string
	Wal_keep_segments                   string
	Wal_sender_timeout                  string
	Max_replication_slots               string
	Track_commit_timestamp              string
	Synchronous_standby_names           string
	Vacuum_defer_cleanup_age            string
	Hot_standby                         string
	Max_standby_archive_delay           string
	Max_standby_streaming_delay         string
	Wal_receiver_status_interval        string
	Hot_standby_feedback                string
	Wal_receiver_timeout                string
	Wal_retrieve_retry_interval         string
	Enable_bitmapscan                   string
	Enable_hashagg                      string
	Enable_hashjoin                     string
	Enable_indexscan                    string
	Enable_indexonlyscan                string
	Enable_material                     string
	Enable_mergejoin                    string
	Enable_nestloop                     string
	Enable_seqscan                      string
	Enable_sort                         string
	Enable_tidscan                      string
	Seq_page_cost                       string
	Random_page_cost                    string
	Cpu_tuple_cost                      string
	Cpu_index_tuple_cost                string
	Cpu_operator_cost                   string
	Network_byte_cost                   string
	Remote_query_cost                   string
	Effective_cache_size                string
	Geqo                                string
	Geqo_threshold                      string
	Geqo_effort                         string
	Geqo_pool_size                      string
	Geqo_generations                    string
	Geqo_selection_bias                 string
	Geqo_seed                           string
	Default_statistics_target           string
	Constraint_exclusion                string
	Cursor_tuple_fraction               string
	From_collapse_limit                 string
	Join_collapse_limit                 string
	Log_destination                     string
	Logging_collector                   string
	Log_directory                       string
	Log_filename                        string
	Log_file_mode                       string
	Log_truncate_on_rotation            string
	Log_rotation_age                    string
	Log_rotation_size                   string
	Syslog_facility                     string
	Syslog_ident                        string
	Event_source                        string
	Client_min_messages                 string
	Log_min_messages                    string
	Log_min_error_statement             string
	Log_min_duration_statement          string
	Debug_print_parse                   string
	Debug_print_rewritten               string
	Debug_print_plan                    string
	Debug_pretty_print                  string
	Log_checkpoints                     string
	Log_connections                     string
	Log_disconnections                  string
	Log_duration                        string
	Log_error_verbosity                 string
	Log_hostname                        string
	Log_line_prefix                     string
	Log_lock_waits                      string
	Log_statement                       string
	Log_replication_commands            string
	Log_temp_files                      string
	Log_timezone                        string
	Cluster_name                        string
	Update_process_title                string
	Track_activities                    string
	Track_counts                        string
	Track_io_timing                     string
	Track_functions                     string
	Track_activity_query_size           string
	Stats_temp_directory                string
	Log_parser_stats                    string
	Log_planner_stats                   string
	Log_executor_stats                  string
	Log_statement_stats                 string
	Autovacuum                          string
	Log_autovacuum_min_duration         string
	Autovacuum_max_workers              string
	Autovacuum_naptime                  string
	Autovacuum_vacuum_threshold         string
	Autovacuum_analyze_threshold        string
	Autovacuum_vacuum_scale_factor      string
	Autovacuum_analyze_scale_factor     string
	Autovacuum_freeze_max_age           string
	Autovacuum_multixact_freeze_max_age string
	Autovacuum_vacuum_cost_delay        string
	Autovacuum_vacuum_cost_limit        string
	Search_path                         string
	Default_tablespace                  string
	Temp_tablespaces                    string
	Check_function_bodies               string
	Default_transaction_isolation       string
	Default_transaction_read_only       string
	Default_transaction_deferrable      string
	Session_replication_role            string
	Statement_timeout                   string
	Lock_timeout                        string
	Vacuum_freeze_min_age               string
	Vacuum_freeze_table_age             string
	Vacuum_multixact_freeze_min_age     string
	Vacuum_multixact_freeze_table_age   string
	Bytea_output                        string
	Xmlbinary                           string
	Xmloption                           string
	Gin_fuzzy_search_limit              string
	Gin_pending_list_limit              string
	Datestyle                           string
	Intervalstyle                       string
	Timezone                            string
	Timezone_abbreviations              string
	Extra_float_digits                  string
	Client_encoding                     string
	Lc_messages                         string
	Lc_monetary                         string
	Lc_numeric                          string
	Lc_time                             string
	Default_text_search_config          string
	Dynamic_library_path                string
	Local_preload_libraries             string
	Session_preload_libraries           string
	Deadlock_timeout                    string
	Max_locks_per_transaction           string
	Max_pred_locks_per_transaction      string
	Array_nulls                         string
	Backslash_quote                     string
	Default_with_oids                   string
	Escape_string_warning               string
	Lo_compat_privileges                string
	Operator_precedence_warning         string
	Quote_all_identifiers               string
	Sql_inheritance                     string
	Standard_conforming_strings         string
	Synchronize_seqscans                string
	Transform_null_equals               string
	Exit_on_error                       string
	Restart_after_crash                 string
	Persistent_datanode_connections     string
	Max_coordinators                    string
	Max_datanodes                       string
	Pgxc_node_name                      string
	Gtm_backup_barrier                  string
	Include_dir                         string
	Include_if_exists                   string
	Include                             string
	Max_pool_size                       string
	Pool_conn_keepalive                 string
	Pool_maintenance_timeout            string
	Pooler_port                         string
	Gtm_host                            string
	Gtm_port                            string
	Kv_pairs                            map[string]string
}

func (pc *pg_conf) set_contents(file_contents string) {

	pc.File_contents = file_contents
}

func (pc *pg_conf) parse() (errs error) {

	if len(pc.File_contents) == 0 {
		return errors.New("No contents set for config")
	}

	re_blank := regexp.MustCompile(`^[ \t]*$`)
	re_comment := regexp.MustCompile(`^[ \t]*#`)
	kv_re := regexp.MustCompile(`[ \t]?(?P<key>\S+)[ |\t]*=[ |\t]?(?P<value>\S+)[ |\t]*#?(?:.*)?`)

	retMap := make(map[string]string)

	for _, v := range strings.Split(pc.File_contents, "\n") {

		v = strings.TrimSpace(v)

		if re_blank.MatchString(v) || re_comment.MatchString(v) {

			continue
		}

		var ik, iv string

		for _, v := range kv_re.FindAllStringSubmatch(v, -1) {
			for i, vv := range v {
				if i != 0 {

					if kv_re.SubexpNames()[i] == "key" {
						ik = strings.ToLower(vv)
					}

					if kv_re.SubexpNames()[i] == "value" {
						iv = vv
					}
				}
			}

			retMap[ik] = iv
		}
	}

	pc.Kv_pairs = retMap

	mapstructure.Decode(retMap, &pc)
	return
}
