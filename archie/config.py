
import os


basedir = os.path.abspath(os.path.dirname(__file__))


class Config(object):
    SQLALCHEMY_DATABASE_URI = ''
    SQLALCHEMY_TRACK_MODIFICATIONS = False

    @staticmethod
    def init_app(app):
        print("Initalizing app...")


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

