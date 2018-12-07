"""Deploys, runs, & tests the `Phase 0` backend via CLI

Requires a `PostgreSQL` db, managed by this module. 
"""


import os
from dotenv import load_dotenv


dotenv_loc = os.path.join(os.path.dirname(__file__), '.env')
if os.path.exists(dotenv_loc) and not os.environ.get('FLASK_ENV'):
    print('Importing environment from .env...')
    load_dotenv(dotenv_loc)


import click
from app import create_app, db
from app.models import Permission, Role, User, Question, Comment, Answer
from flask_migrate import Migrate, upgrade


app = create_app(os.environ.get('FLASK_ENV') or 'default')
migrate = Migrate(app, db)


@app.shell_context_processor
def make_shell_command():
    return dict(db=db, Permission=Permission, Role=Role, User=User, 
                Question=Question, Comment=Comment, Answer=Answer)


@app.cli.command()
def test():
    """Finds `tests/` directory & runs test suite."""
    import unittest

    tests = unittest.TestLoader().discover('tests')
    unittest.TextTestRunner(verbosity = 2).run(tests)


@app.cli.command()
def deploy():
    """Updates database with latest migration(s) & inserts Roles."""
    upgrade()

    Role.insert_roles()


