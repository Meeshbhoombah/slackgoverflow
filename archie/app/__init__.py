#!usr/bin/python3
"""Main entrypoint into 'Architect' Flask application."""


import os
from .config import config
from flask import Flask, jsonify
from flask_restful import Resource, Api, reqparse
from flask_sqlalchemy import SQLAlchemy


######## FLASK CONFIG ########
app = Flask(__name__)

# Using config object from `config.py`
server = 'default'

try:
    server = os.environ.get('FLASK_ENV')
except:
    pass

app.config.from_object(config[server]) 
config[server].init_app(app)


######## SQL DB CONFIG ########
app.config['SQLALCHEMY_DATABASE_URI'] = 'postgresql://{usr}:{dbpass}@{host}:5432/{db}'.format(
    usr = app.config['DBUSER'],
    dbpass = app.config['DBPASS'],
    host = app.config['DBHOST'],
    db = app.config['DBNAME']
)

from .user.model import user_db
user_db.init_app(app)

# man evocation required to create db
with app.app_context():
    user_db.create_all()


######## ROUTES ########
api = Api(app)

from .user.resource import User

api.add_resource(User, '/user')

