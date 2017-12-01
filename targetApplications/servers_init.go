package targetApplications


// This is a file that is used to init servers related health and spec information that may be relevant for the application.

// Initial Recce:
// CPU(nproc), Memory, HDD details as shown on df -h
// Uptime of the server
// Time delta from current server (where this binary is running) -- Config level threshold for time difference check (by default we can take it as 1.5 seconds
// Populate the ulimit -a for the logged in user


// Run the HDD usage check()
// Save the load average values for this server

// HDD Usage check:
// df -h -- update the server level object with the latest details.
// Check if there are any values higher than expected for any of the settings.


// - On the configuration level you can set the max disk percentage -- This will then be used when checking for disk
// space issues. Any drives higher than the configured percentage need to be maintained on the server level as a simple
// struct server_level_issues: { issue_type : 'string, affected_details : 'string', message : '' }


// Get load average details for this server
// Actual Load average of 1, 5 and 15 minutes (uptime)  -- CPU Adjusted Load average : the 3 values divided by the number of CPU cores



