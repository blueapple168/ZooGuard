
# Features

## All Nodes:

### `ls` of the logs folder
	- Check the logs folder, show the ls of the files with date and size.
		- Standard strings that will be removed from logs as grep -v --- these are for banners that will just pollute logs -- This will be as []string in the config file.
		- Good to have, size of log per day for that machine -- You can do a wc -l for this if you want
		- This will be a good comparision measure across machines, if a particular machine is generating more logs per day then it needs to be investigated why.


### Drive space taken by the log folder for this node

### Drive space taken by the log folder for all nodes comparision

### Uptime of this database process, when it was started

### Logs will need to be combined into a single file if we want to serve queries on a single day basis

### Download the whole log file for the server for a particular day


### Run a grep on the logs for a particular day to search for keywords
	- Optionally specify a -v option to exclude stuff


### Do a wc -l to do a count on the number of instances of that search term for that day
	
	
### View the last 500 lines of a log file for a particular day
	- User can specify the number of lines they want to see from the list
	- User can give -v option to remove stuff from the tail command


### Have an error master:
	- Keyword name
	- Expected error file where it will be found
	- The regular expression for checking if the error exists
	- The description of this error message, what it means and how to go about fixing it.
	- Severity of this error -- Does this need to be fixed immediately or can wait for a maintenance timout window.


### On each server type run a log analysis and check for known type errors and list them separately 
	- For each file type that is being checked, check for specific known errors on the server.
	- Maintain an object of :
		- Date that error occured
		- How many instances of the error were found on the server
		- Error keyword


### Drive Space
	- List of servers that have df -h higher than 65% certain percentage
	- List of servers that have df -h higher than user specified percentage


### OS Level Params
	- Get OS level params for all the nodes
	- Specify a key value combination for the os param that needs to be in place for a particular role
		- The value can be a computed value from an OS command -- so Memory amount or Number of files
		- The value can either be a string = match
		- The value can be an integer with a <, >, <=, >= match 
		(What you will be checking here is stuff like noproc, nofiles)
	- On the Server level have roles added based on what you have instaled on it, all servers will share one role, 'all'
	- List the servers that do not pass the OS level test.
	- Remember, these need to be run as the user that actually runs postgres


### Performance and Health Tests
	- These are scripted tests you can run on a server, this is like 'fsync' test -- you can specify on what servers you want to run this on and you will be able to see the results of it.
	- Get a count of the number of threads used by GTM on all the GTM and Proxy servers
	- Do a check of time across all the servers and make sure it is running the same time -- if you have deviation list the servers that are different or list all the servers.


### Specified the desired configuration you want to ensure based on the role
	- This will be role specific with the value that you expect for that server
	- There is a special role called 'all' -- This applies to all the servers.
	- If something is specified in 'all' and specific role, then the specific role will win and that configuration will be the one that will be returned for the expected configuration.
	- Methods to be able to query based on the role name, keyword name 


### Parse postgres config files
	- Parse these files and fill them for a given server:
		- postgresql.conf 
			- Remove the comments
			- Get all the unique possible values for the postgresql.com
			- Make an object of the possible values
				- This will have the keyword for the value
				- Description of what it means
				- Does a change on this value require a restart?
		- pg_hba.conf


### Test the configuration files
	- Show when the configuration files were modified last
	- Get the configuration files of all the servers
	- Run the test across each machine
	- Get the all and the machine role specific configuration specified
	- Check if the configuration matches the one specified on the role. If not fail the configuration test.
	- After the defined tests have been run, the custom tests will be run:
		- The connection pool settings test needs to be run -- Check if the configuration is correct for this role server.


























