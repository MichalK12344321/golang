# golang
![BE schematic](resources/be_schem.png)

Requirements:
* The user should be able to start SSH collection using topology (ip/host and port) and credentials (username and password).
* The user should be able to list all collections and check their IDs and statuses.
* The user should be able to download log for finished collection.

Extra requirements:
1. Filtering:
  The user should be able to filter list of all collections by status (eg. 'running')
  The user should be able to mix filters to list all finished collections ('failed' and 'success')
2. Terminating:
  The user should be able to terminate currently running collection.
  Terminated collection status shall be 'terminated'.
3. Printing running collection log:
  The user should be able to get log for running collection.
  The log for running collection shall print information collected until the request.
