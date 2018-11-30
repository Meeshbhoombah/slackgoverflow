#!usr/bin/python3
"""Main entrypoint into 'Architect' Flask application."""


import os
from .config import config
from flask import Flask, jsonify
from flask_restful import Resource, Api, reqparse
from flask_sqlalchemy import SQLAlchemy


db = SQLAlchemy()


def create_app(config):
    "Configures the server, db, and routes"
    app = Flask(__name__)

    # Using config object from `config.py`
    server = 'default'

    try:
        server = os.environ.get('FLASK_ENV')
    except:
        pass

    app.config.from_object(config[server]) 
    config[server].init_app(app)


    """ SQL DB CONFIG """
    app.config['SQLALCHEMY_DATABASE_URI'] = 'postgresql://'
    + '{usr}:'.format(app.config['DBUSER'])
    + '{dbpass}'.format(dbpass = app.config['DBPASS'])
    + '@{host}:'.format(host = app.config['DBHOST'])
    + '5432/{db}'.format(db = app.config['DBNAME'])


    """ ROUTES """
    api = Api(app)

    from .user.resource import User

    api.add_resource(User, '/user')

