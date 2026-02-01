import logging
import os

from const import AMOUNT_VALUE, CHECKOUT_URL, CUSTOMER_URL, MSISDN, RECARGE_URL
from http_utils import get, post

logger = logging.getLogger(__name__)

def make_checkout(session):
    response = get(session, CUSTOMER_URL, "Obeter detalhes para fazer checkout")
    customer = response.json()
    customer.pop("lastReloadAt", None)
    credit_card = next(
        (card for card in customer.get("creditCards", []) if isinstance(card, dict) and "token" in card),
        None,
    )
    if not credit_card:
        logger.exception("Sem cartao de credito cadastrado")
    checkout_payload = {
        "customer": customer,
        "rechargeValue": AMOUNT_VALUE * 100,
        "paymentMethod": "credit",
        "extras": {"targetMsisdn": MSISDN, "scheduleConfig": None, "selectedCard": None, "isScheduled": False},
    }
    response = post(session, CHECKOUT_URL, checkout_payload, "Fazer checkout")
    logger.info("Checkout realizada com sucesso!")
    return credit_card


def make_credit_card_payment(session, id_recharge):
    credit_card = make_checkout(session)
    recarga_payload = {
        "targetMsisdn": MSISDN,
        "rechargeValue": {"id": id_recharge, "value": AMOUNT_VALUE * 100},
        "paymentMethod": {
            "type": "credit",
            "data": {"token": credit_card["token"], "cvv": os.environ["CVV"], "favorite": True},
        },
    }
    response = post(session, RECARGE_URL, recarga_payload, "Fazer pagamento com cartao de credito")
    logger.info("Recarga realizada com sucesso!")
    logger.info(response.json())
