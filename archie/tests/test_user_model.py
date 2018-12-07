
import unittest
import time
from datetime import datetime
from app import create_app, db
from app.models import Role, User, Question, Comment, Answer


class UserModelTestCase(unittest.TestCase):

    def setUp(self):
        self.app = create_app('testing')
        self.app_context = self.app.app_context()
        self.app_context.push()
        db.create_all()
        Role.insert_roles()


    def tearDown(self):
        db.session.remove()
        db.drop_all()
        self.app_context.pop()


    def test_password_setter(self):
        u = User(password = '123')
        self.assertTrue(u.hashword is not None)


    def test_password_getter(self):
        u = User(password = '123')
        with self.assertRaises(AttributeError):
            u.password


    def test_password_verification(self):
        u = User(password = '123')
        self.assertTrue(u.verify_password('123'))
        self.assertFalse(u.verify_password('321'))

    
    def test_password_salts_are_random(self):
        u1 = User(password = '123')
        u2 = User(password = '123')
        self.assertNotEqual(u1.hashword, u2.hashword)


