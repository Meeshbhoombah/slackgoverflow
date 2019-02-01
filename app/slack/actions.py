
from datetime import datetime
from slackclient import SlackClient


class Handler(object):

    def __init__(self, event_type):
        handlers = {
                'member_joined_channel' : self.member_joined_channel
        }

        self.handler = handlers[event_type]

        self.sc = SlackClient(token = current_app.config['SLACK_AUTH_TOKEN'])
        self.sb = SlackClient(token = current_app.config['SLACK_BOT_TOKEN'])



    def member_joined_channel(event_data):
        """Handles `member_joined_channel` event when User joins the channel`"""
         try:
            assert event_data['event']['channel'] == 'CEET14B25'
        except (AssertionError):
            return make_response("Not the #devp2p channel.", 404)

        # Get slack_id for User who triggered event
        member_id = event_data['event']['user']

        # Check if preexisting user
        u = User.query.filter_by(_slack_id = member_id).first()

        if u is None:
            # first time
            u = User(slack_id = member_id)

            db.session.add(u)
            db.session.commit()

            token = u.generate_registration_token() 
            create_account = current_app.config['BASE_URL'] + '/create/' + token

            msg = "Welcome to #devp2p! Create an account to start asking/answering questions."
            msg_attachments = [
                {
                    "fallback": "Something went wrong. Please rejoin the channel.",
                    "color": "#000000",
                    "title": "Sign Up with Architect :hammer:",
                    "title_link": create_account,
                    "footer":"Sign up takes < 3 minutes (not to mention, Unlocks Drops)."
                }
            ]
           
            sc.api_call(
                "chat.postMessage",
                channel = u._slack_id,
                as_user = True,
                text = msg,
                attachments = msg_attachments
            )

        else:
            u.pong()

            if u.last_seen.day < datetime.today().day:
                msg = 'Welcome back!'

                if u.username:
                    msg += u.username

                sc.api_call(
                    "chat.postMessage",
                    channel = u._slack_id,
                    as_user = True,
                    text = msg
                )

        response = make_response("Success.", 200)
        response.headers['X-Slack-Powered-By'] = 'Architect'
        return response
        

