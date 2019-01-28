# Archie
Communicates with Slack to send/recieve DMs and handle Event/Slash commands.

## Getting Started
These instructions will get you a copy of the project up and running on your 
local machine for development and testing purposes. See deployment for notes on 
how to deploy the project on a live system.

### Prerequisites
- Postgres SQL `11`
- `pip` `10.0.1`
- `virtualenv` `16.0.0`
- Git `2.17.1`
- Python `3.7.0`

### Installing
Clone this repository.
```
$ git clone [URL] 
```

Navigate to the current directory.
```
$ cd make/archie/
```

Assuming Postgres is installed, start the server in the background. It is
recommended to use the Desktop application, but the same can be accomplished
via the command-line.
```
$ pg_ctl -D /usr/local/var/postgres start
```
To later stop the server:
```
$ pg_ctl -D /usr/local/var/postgres stop
```

Create a `python3` virtual environment with `virtualenv`. For instructions on 
using Docker, view the `README` in the directory above. 
```
$ virtualenv -p python3 venv
```

Activate the virtual environment and set the **required** environment variables.
```
```

Set up the database using the `click` built-in for Flask. Do `flask --help` to
view all commands. First run the shell then create the database using the
injected variables.
```
$ flask shell
Python 3.7.0 (default, Jun 29 2018, 20:14:27) 
...
>>> db.create_all()
```

Then configure and migrate the database.
```
$ flask deploy
```

## Running Tests
Archie's test suite is located in `/test`.

Archie leverages Flask's built in integration of the 
[click](https://click.palletsprojects.com/en/7.x/)
command-line interface to automate testing. Visit 
[the Flask documentation](http://flask.pocoo.org/docs/0.12/cli/) 
for more information.

Simply run the `test` command.
```
$ flask test
test_app_exists (test_deploy.ServerRunningTestCase) ... 
...
```



