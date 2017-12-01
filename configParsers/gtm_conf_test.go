package configParsers

import (
	"testing"
)

const (
	sampleGtmConf = `
#-----------------------------
# GTM Proxy configuration file
#-----------------------------
#
# This file must be placed on gtm working directory
# specified by -D command line option of gtmProxy or gtm_ctl.
# The configuration file name must be "gtmProxy.conf"
#
#
# This file consists of lines of the form
#
#  name = value
#
# (The "=" is optional.) Whitespace may be used.   Comments are
# introduced with "#" anywhere on a line.  The complete list of
# parameter names and allowed values can be found in the
# Postgres-XC documentation.
#
# The commented-out settings shown in this file represent the default
# values.
#
# Re-commenting a setting is NOT sufficient to revert it to the default
# value.
#
# You need to restart the server.

#------------------------------------------------------------------------------
# GENERAL PARAMETERS
#------------------------------------------------------------------------------
nodename = 'one'					# Specifies the node name.
								# (changes requires restart)
#listen_addresses = '*'			# Listen addresses of this GTM.
								# (changes requires restart)
port = 6666					# Port number of this GTM.
								# (changes requires restart)

#------------------------------------------------------------------------------
# GTM PROXY PARAMETERS
#------------------------------------------------------------------------------
#worker_threads = 1				# Number of the worker thread of this
								# GTM proxy
								# (changes requires restart)

#------------------------------------------------------------------------------
# GTM CONNECTION PARAMETERS
#------------------------------------------------------------------------------
# Those parameters are used to connect to a GTM server
gtm_host = 'localhost' 					# Listen address of the active GTM.
								# (changes requires restart)
gtm_port = 6668 					# Port number of the active GTM.
								# (changes requires restart)

#------------------------------------------------------------------------------
# Behavior at GTM communication error
#------------------------------------------------------------------------------
#gtm_connect_retry_interval = 0	# How long (in secs) to wait until the next
								# retry to connect to GTM.
#
#
#------------------------------------------------------------------------------
# Other options
#------------------------------------------------------------------------------
#keepalives_idle = 0			# KeepalivesIdle parameter.
#keepalives_interval = 0		# KeepalivesInterval parameter.
#keepalives_count = 0			# KeepalivesCount internal parameter.
#log_file = 'gtmProxy.log'		# Log file name
#log_min_messages = WARNING		# log_min_messages.  Default WARNING.
							  	# Valid value: DEBUG, DEBUG5, DEBUG4, DEBUG3,
								# DEBUG2, DEBUG1, INFO, NOTICE, WARNING,
								# ERROR, LOG, FATAL, PANIC.

#===========================
# Added at initialization, 20171115_16:39:53
nodename = 'gtm_pxy1'
listen_addresses = '*'
port = 20001
gtm_host = '10.0.1.5'
gtm_port = 8080
worker_threads = 1
gtm_connect_retry_interval = 1
# End of addition
`
)

func TestGTMParse(t *testing.T) {

	var p GTMConfig
	p.FileContents = sampleGtmConf
	p.Parse()

	/*
		nodename = 'gtm_pxy1'
		listen_addresses = '*'
		port = 20001
		gtm_host = '10.0.1.5'
		gtm_port = 8080
		worker_threads = 1
		gtm_connect_retry_interval = 1
	*/

	if p.Nodename != `gtm_pxy1` || p.ListenAddresses != `*` || p.GtmHost != `10.0.1.5` || p.GtmPort != `8080` || p.WorkerThreads != `1` || p.GtmConnectRetryInterval != `1` {

		t.Errorf("GTM Proxy parser got unexpected results -- Returned Object %v", p)
		t.Fail()
	}
}
