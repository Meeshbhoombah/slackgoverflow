
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


    def test_creation_isfail_without_slack_id(self):
        u = User(password = '123')
        with self.assertRaises(IntegrityError):
            db.session.commit()


    def test_default_user_isunregistered(self):
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
        u1 = User(password = '123', slack_id = 'U12345678')
        self.assertNotEqual(u0.hashword, u1.hashword)


    def test_valid_registration_token(self):
        u = User(slack_id = 'W12345678')
        db.session.add(u)
        db.session.commit()
        token = u.generate_registration_token()
        u.username = 'test'
        u.password = '123'
        db.session.commit()
        self.assertTrue(u.confirm_registration(token))


    def test_invalid_registration_token(self):
        u0 = User(slack_id = 'W12345678')
        u1 = User(slack_id = 'U12345678')
        db.session.add(u0)
        db.session.add(u1)
        db.session.commit()
        token = u0.generate_registration_token()
        self.assertFalse(u1.confirm_registration(token))


    def test_expired_registration_token(self):
        u = User(slack_id = 'W12345678')
        db.session.add(u)
        db.session.commit()
        token = u.generate_registration_token(expiration = 1)
        time.sleep(2)
        self.assertFalse(u.confirm_registration(token))
    

    def test_timestamps(self):
        u = User(slack_id = 'W12345679')        
        db.session.add(u)
        db.session.commit()
        self.assertTrue(
            (datetime.utcnow() - u.created_on).total_seconds() < 3)
        self.assertTrue(
            (datetime.utcnow() - u.last_seen).total_seconds() < 3)


    def test_ping(self):
        u = User(slack_id = 'W12345678')
        db.session.add(u)
        db.session.commit()
        time.sleep(2)
        last_seen_before = u.last_seen
        u.ping()
        self.assertTrue(u.last_seen > last_seen_before)

