
import os


basedir = os.path.abspath(os.path.dirname(__file__))


class Config(object):
    SECRET_KEY = os.environ.get('SECRET_KEY') or 'YOULL NEVER GUESS'
    BASE_URL = os.environ.get('BASE_URL')

    SQLALCHEMY_DATABASE_URI = ''
    SQLALCHEMY_TRACK_MODIFICATIONS = False

    SLACK_SIGNING_SECRET = os.environ.get('SLACK_SIGNING_SECRET')
    SLACK_CLIENT_ID = os.environ.get('SLACK_CLIENT_ID')
    SLACK_CLIENT_SECRET = os.environ.get('SLACK_CLIENT_SECRET')
    SLACK_AUTH_TOKEN = os.environ.get('SLACK_AUTH_TOKEN')
    SLACK_BOT_TOKEN = os.environ.get('SLACK_BOT_TOKEN')


    @staticmethod
    def init_app(app):
        pass


class Development(Config):
    DEBUG = True
    DEVELOPMENT = True
    BASE_URL = os.environ.get('NGROK_URL')


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

