from datetime import datetime

def handle(req):
    """handle a request to the function
    Args:
        req (str): request body
    """
    resp = str(datetime.now()) + ": " +  req
    return resp
