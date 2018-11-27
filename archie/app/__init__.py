#!usr/bin/python3
"""Main entrypoint into 'Simple SMS Client' Flask and SQL application.

Demonstates functionality of the simple blockchain and smart contracts in a 
tangible way.

License: MIT
Website:
"""

import os
from .config import config
from flask import Flask, jsonify
from flask_restful import Resource, Api, reqparse
from flask_sqlalchemy import SQLAlchemy

"""CONFIG"""
app = Flask(__name__)

# Using config object from `config.py`
server = 'default'

try:
    os.environ.get('FLASK_ENV')
except:
    pass

app.config.from_object(config[server]) 
config[server].init_app(app)

api = Api(app)

"""CONFIGURE DATABASE"""
app.config['SQLALCHEMY_DATABASE_URI'] = 'postgresql://{usr}:{dbpass}@{host}:5432/{db}'.format(
    usr = app.config['DBUSER'],
    dbpass = app.config['DBPASS'],
    host = app.config['DBHOST'],
    db = app.config['DBNAME']
)

# MODELS
from .usr.model import user_db
user_db.init_app(app)

with app.app_context():
    user_db.create_all()


"""ROUTES"""
from .twilio.resource import Client
from .usr.resource import User

api.add_resource(User, '/user')
api.add_resource(Client, '/twilio')


