"""Test setup and creation of the Phase-0 Backend"""

import unittest
from flask import current_app
from app import create_app


class ServerRunningTestCase(unittest.TestCase):

    def SetUp(self):
        self.app = create_app('Testing') 
        self.app_context = self.app.app_context()
        self.app_context.push()

    
    def TearDown(self):
        self.app_context.pop()


    def test_app_exists(self):
        self.assetFalse(current_app is None) 


    def test_app_is_testing(self):
        self.assertTrue(current_app.config['TESTING'])
        
