
import unittest
import time
from datetime import datetime
from app import create_app, db
from app.models import Role, User, Question, Comment, Answer
from sqlalchemy.exc import IntegrityError


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


    def test_user_creation_without_slack_id(self):
        u = User(password = '123')
        with self.assertRaises(IntegrityError):
            db.session.commit()


    def test_default_user_is_unregistered(self):
        u = User(slack_id = 'U12345678')
        default_role = Role.query.filter_by(default = True).first()
        self.assertEqual(u.role_id, default_role.id)


    def test_slack_setter(self):
        u0 = User(slack_id = 'U12345678')
        self.assertIsNotNone(u0.slack_id)
        u1 = User(slack_id = 'W12345678')
        self.assertIsNotNone(u1.slack_id)

        with self.assertRaises(AttributeError):
            u2 = User(slack_id = 'A12345678')

        with self.assertRaises(AttributeError):
            u3 = User(slack_id = 'W123456789')


    def test_password_setter(self):
        u = User(password = '123', slack_id = 'W12345678')
        self.assertIsNotNone(u.hashword)


    def test_password_getter(self):
        u = User(password = '123', slack_id = 'W12345678')
        with self.assertRaises(AttributeError):
            error = u.password


    def test_password_verification(self):
        u = User(password = '123', slack_id = 'W12345678')
        self.assertTrue(u.verify_password('123'))
        self.assertFalse(u.verify_password('321'))

    
    def test_password_salts_are_random(self):
        u0 = User(password = '123', slack_id = 'W12345678')
        u1 = User(password = '123', slack_id = 'W12345678')
        self.assertNotEqual(u0.hashword, u1.hashword)

