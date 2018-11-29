#!usr/bin/python3
"""Configurations for Make School x Architect: Phase-0"""

import os


basedir = os.path.abspath(os.path.dirname(__file__))


class Config(object):
    SQLALCHEMY_DATABASE_URI = ''
    DBUSER = ''
    DBHOST = ''
    DBPASS = ''
    DBNAME = ''

    SQLALCHEMY_TRACK_MODIFICATIONS = False

    @staticmethod
    def init_app(app):
        pass


class Development(Config):
    DEBUG = True
    HOST = '0.0.0.0'
    DBUSER = 'meeshbhoombah'
    DBHOST = '127.0.0.1'
    DBPASS = None
    DBNAME = 'phase0dev'


class Production(Config):
    DEBUG = False
    HOST = '0.0.0.0'
    DBUSER = 'meeshbhoombah'
    DBHOST = '127.0.0.1'
    DBPASS = None
    DBNAME = 'phase0dev'


config = {
    'development': Development,
    'prod': Production,

    'default': Development
}

