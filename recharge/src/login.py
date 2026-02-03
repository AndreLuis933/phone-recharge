from const import LOGIN_URL, MSISDN, SMS_URL
from cookies import save_cookies
from curl_cffi import requests
from http_utils import post
from wait_redis_code import wait_for_code


class CodigoNaoRecebidoError(Exception):
    """Exceção lançada quando o código não é recebido dentro do tempo limite."""


def fazer_login(session: requests.Session):
    post(session, SMS_URL, {"msisdn": MSISDN}, "Enviar sms")

    code = wait_for_code()
    if code is None:
        msg = "Codigo nao chegou dentro de 5 minutos"
        raise CodigoNaoRecebidoError(msg)

    login_payload = {
        "channel": "VIVO_WEB",
        "data": code,
        "keepLogged": True,
        "msisdn": MSISDN,
        "redirectRoute": "/web/home",
        "type": "sms",
    }

    post(session, LOGIN_URL, login_payload, "Logar com o codigo sms enviado")
    save_cookies(session)
