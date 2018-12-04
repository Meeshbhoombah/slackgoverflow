"""Deploys, runs, & tests the `Phase 0` backend via CLI

Requires a `PostgreSQL` db, managed by this module. 
"""

# TODO: import create_app not working
# TODO: determine if manage.py handles creation of db


""" .ENV FILE """
import os
import sys
from dotenv import load_dotenv


sys.path.append('.')

dotenv_loc = os.path.join(os.path.dirname(__file__), '.env')
if os.path.exists(dotenv_loc) and not os.environ.get('FLASK_ENV'):
    print('Importing environment from .env...')
    load_dotenv(dotenv_loc)


""" CREATE APP """
import click
from app import create_app


app = create_app(os.environ.get('FLASK_ENV') or 'default')


""" CLI OPERATIONS """
@app.cli.command()
def test():
    """Runs Unittests in `tests/`"""
    import unittest

    tests = unittest.TestLoader().discover('tests')
    unittest.TextTestRunner(verbosity = 2).run(tests)

