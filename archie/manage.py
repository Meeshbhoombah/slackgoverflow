"""Deploys, runs, & tests the `Phase 0` backend via CLI

Requires a `PostgreSQL` db, managed by this module. 
"""


import os
from dotenv import load_dotenv

dotenv_loc = os.path.join(os.path.dirname(__file__), '.env')
if os.path.exists(dotenv_loc) and not os.environ.get('FLASK_ENV'):
    print('Importing environment from .env...')
    load_dotenv(dotenv_loc)


COV = None
if os.environ.get('FLASK_COVERAGE'):
    import coverage
    COV = coverage.coverage(branch=True, include='app/*')
    COV.start()


import click
import sys
from app import create_app, db
from flask_migrate import Migrate, upgrade
from app.models import Permission, Role, User, Question, Comment, Answer


app = create_app(os.environ.get('FLASK_ENV') or 'default')
migrate = Migrate(app, db)


@app.shell_context_processor
def make_shell_command():
    return dict(db=db, Permission=Permission, Role=Role, User=User, 
                Question=Question, Comment=Comment, Answer=Answer)


@app.cli.command()
@click.option('--coverage/--no-coverage', default=False,
              help='Run tests under code coverage.')
def test(coverage):
    """Run test suite in `/test`"""
    if coverage and not os.environ.get('FLASK_COVERAGE'):
        import subprocess
        os.environ['FLASK_COVERAGE'] = '1'
        sys.exit(subprocess.call(sys.argv))

    import unittest
    tests = unittest.TestLoader().discover('tests')
    unittest.TextTestRunner(verbosity=2).run(tests)
    if COV:
        COV.stop()
        COV.save()
        print('Coverage Summary:')
        COV.report()
        basedir = os.path.abspath(os.path.dirname(__file__))
        covdir = os.path.join(basedir, 'tmp/coverage')
        COV.html_report(directory=covdir)
        print('HTML version: file://%s/index.html' % covdir)
        COV.erase()


@app.cli.command()
def deploy():
    """Updates database with latest migration(s) & inserts Roles."""
    upgrade()

    Role.insert_roles()


