
import unittest
import time
import datetime import dateime


class TestListenSlack(unittest.TestCase)

    def setUp(self):
        self.app = create_app('testing')
        self.app_context = self.app.app_context
        self.app_context.push()
        db.create_all()
        Role.insert_roles()


    def tearDown(self):
        db.session_remove()
        db.drop_all()
        self.app_context.pop()


    def test_verify_signature(self):
        """ TODO: Generate a Slack signature """
        pass

