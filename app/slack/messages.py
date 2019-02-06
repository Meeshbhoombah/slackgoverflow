
def onboard(member_id):
    msg = "Welcome to #devp2p! Create an account to start asking/answering questions."
    payload = [
        {
            "fallback": "Something went wrong. Please rejoin the channel.",
            "color": "#000000",
            "title": "Sign Up with Architect :hammer:",
            "title_link": create_account,
            "footer":"Sign up takes < 3 minutes. Start earning Drops now!"
        }
    ]   

    return msg, payload


