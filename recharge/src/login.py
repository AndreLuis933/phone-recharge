from curl_cffi import requests

from const import LOGIN_URL, MSISDN, SMS_URL
from cookies import save_cookies
from http_utils import post


def fazer_login(session: requests.Session):
    post(session, SMS_URL, {"msisdn": MSISDN}, "Enviar sms")

    codigo_sms = input("Digite o código do SMS: ")

    login_payload = {
        "channel": "VIVO_WEB",
        "data": codigo_sms,
        "keepLogged": True,
        "msisdn": MSISDN,
        "redirectRoute": "/web/home",
        "type": "sms",
    }

    post(session, LOGIN_URL, login_payload, "Logar com o codigo sms enviado")
    save_cookies(session)
