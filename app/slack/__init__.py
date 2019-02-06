
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
        try:
            assert event_data['event']['channel'] == current_app.config['CHANNEL_ID']
        except (AssertionError):
            return make_response("Not the #slackoverflow channel", 404)

        member_id = event_data['event']['user']

        u = User.query.filter_by(_slack_id = member_id).first()

        # If user not found in db they have not joined `slackoverflow`
        if u is None:
            u = User(slack_id = member_id)

            db.session.add(u)
            db.session.commit()

            token = u.generate_registration_token() 
            create_account = current_app.config['BASE_URL'] + '/create/' + token

            from .messages import onboard
            msg, payload = onboard(member_id)

            self.sc.api_call(
                "chat.postMessage",
                channel = u.slack_id,
                as_user = True,
                text = msg,
                attachments = msg_attachments
            )

        else:
            # User has been in the `slackoverflow` channel
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
        

