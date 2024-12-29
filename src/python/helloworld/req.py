import requests

def get_status_code():
    response = requests.get("https://github.com/idlebot/monorepo")
    return "Hello " + str(response.status_code)