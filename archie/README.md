# Archie
Phase-0 Flask backend.

## Getting Started
These instructions will get you a copy of the project up and running on your 
local machine for development and testing purposes. See deployment for notes on 
how to deploy the project on a live system.

### Prerequisites
- Postgres
- pip
- virtualenv
- Git

### Installing
Clone this repository.
```
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

