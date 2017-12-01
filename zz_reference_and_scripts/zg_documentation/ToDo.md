

- Setup a local cluster of pgxl so you can connnect and talk to it
	- Pull a blank docker script and 



#Drive Space

- On the configuration level you can set the max disk percentage -- This will then be used when checking for disk space issues. Any drives higher than the configured percentage need to be maintained on the server level as a simple struct server_level_issues: { issue_type : 'string, affected_details : 'string', message : '' }


#Ulimits 

- ulimits -a needs to get parsed for a server and 

	
	


#SSH Module

- Currently we are only supporting SSH via passwords, need to add support for keys also
- We do not have a way to tunnel through ssh of another machine, this will be helpful in times where a particular machine has passwordless ssh to other machines.
- We need to be able to sudo as particular users, this will be helpful for checking ulimits for particular users.


