
import os


basedir = os.path.abspath(os.path.dirname(__file__))


class Config(object):
    SECRET_KEY = os.environ.get('SECRET_KEY') or 'YOULL NEVER GUESS'

    SQLALCHEMY_DATABASE_URI = ''
    SQLALCHEMY_TRACK_MODIFICATIONS = False

    SLACK_SIGNING_SECRET = os.environ.get('SLACK_SIGNING_SECRET')

    @staticmethod
    def init_app(app):
        pass


class Development(Config):
    DEBUG = True
    DEVELOPMENT = True


class Testing(Config):
    TESTING = True


class Production(Config):
    DEBUG = False


config = {
    'development' : Development,
    'production'  : Production,
    'testing'     : Testing,

    'default'     : Development
}

