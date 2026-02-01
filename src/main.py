import logging

from curl_cffi import requests

import config  # noqa: F401
from check_plans import best_plain
from const import AMOUNT_VALUE, PAYMENT_METHOD
from cookies import load_cookies, validate_session
from login import fazer_login
from payments.credit_card import make_credit_card_payment
from payments.pix import make_pix_payment

logger = logging.getLogger(__name__)


def main():
    session = requests.Session()

    if not (load_cookies(session) and validate_session(session)):
        logger.info("Fazendo login")
        fazer_login(session)

    best = best_plain(session)
    if int(best["valor"] != AMOUNT_VALUE):
        logger.error("A configuraçao nao esta a mais otimizada")
        logger.info(f"Valor configurado {AMOUNT_VALUE} e esperado {best['valor']}")
        logger.info(best)
        return
    if PAYMENT_METHOD == "pix":
        make_pix_payment(session)
    elif PAYMENT_METHOD == "credit_card":
        make_credit_card_payment(session, best["id"])
    else:
        logger.error("metodo invalido")


if __name__ == "__main__":
    main()
