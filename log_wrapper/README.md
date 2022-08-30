# Log Wrapper
### This project execute a python script and continually read the stdout and stderr from the python script to process and register in a SQLite database for logging purposes.

---
## How to use this project:
1. Build the binary:
```bash
go build .
```
2. Execute the binary providing the following argumments:
```bash
./log_wrapper -database-folder=<path> -database-name=<name> -python-script=<file.py>
```
3. By default the log_wrapper is executed without verbose, but you can activate using the -v flag. This will show the SQL insert queries for logs and errors in the SQLite database.
4. ```bash
./log_wrapper -database-folder=<path> -database-name=<name> -python-script=<file.py> -v
```
