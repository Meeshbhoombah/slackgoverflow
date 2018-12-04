"""Creates an instance of the Phase-0 server"""


import os
from config import config
from flask import Flask
from flask_sqlalchemy import SQLAlchemy


db = SQLAlchemy()


def create_app(server):
    """Configures the server, db, and routes"""
    app = Flask(__name__)

    app.config.from_object(config[server] or 'default') 
    config[server].init_app(app)

    try:
        dbpass = os.environ['DBPASS']
    except KeyError:
        dbpass = ''

    app.config.update(
        SQLALCHEMY_DATABASE_URI = 'postgresql+psycopg2://{}:{}@{}/{}'.format(
            os.environ['DBUSER'],
            dbpass,
            os.environ['DBHOST'],
            os.environ['DBNAME']
        )
    )

    db.init_app(app)

    return app

