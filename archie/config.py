
import os


basedir = os.path.abspath(os.path.dirname(__file__))


class Config(object):

    @staticmethod
    def init_app(app):
        print("Initalizing app...")


class Development(Config):
    DEBUG = True
    DEVELOPMENT = True
    HOST = '0.0.0.0'


class Production(Config):
    DEBUG = False


config = {
    'development' : Development,
    'production'  : Production,

    'default'     : Development
}

