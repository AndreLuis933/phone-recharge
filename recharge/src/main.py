import logging

from check_plans import best_plain
from config import is_docker
from const import AMOUNT_VALUE, PAYMENT_METHOD
from cookies import load_cookies, validate_session
from curl_cffi import requests
from login import fazer_login
from notification import notify_error, notify_success
from payments.credit_card import make_credit_card_payment
from payments.pix import make_pix_payment

logger = logging.getLogger(__name__)

class PlanoEscolidoNaoMelhorError(Exception):
    """Exceção lançada quando o código não é recebido dentro do tempo limite."""


def main():
    session = requests.Session()

    if not (not is_docker and load_cookies(session) and validate_session(session)):
        logger.info("Fazendo login")
        fazer_login(session)

    best = best_plain(session)
    if int(best["valor"]) != AMOUNT_VALUE:
        logger.error("A configuraçao nao esta a mais otimizada")
        logger.info(f"Valor configurado {AMOUNT_VALUE} e esperado {best['valor']}")
        logger.info(best)
        msg = "Nao execultado pois a configuraçao nao é a mais otimizada"
        raise PlanoEscolidoNaoMelhorError(msg)
    if PAYMENT_METHOD == "pix":
        make_pix_payment(session)
    elif PAYMENT_METHOD == "credit_card":
        make_credit_card_payment(session, best["id"])
    else:
        msg = "PAYMENT_METHOD invalid"
        raise ValueError(msg)


if __name__ == "__main__":
    try:
        main()
        notify_success("Job Mensal Das")
    except Exception:
        notify_error("Job Mensal Das")
        raise
