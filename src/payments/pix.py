import logging

from const import AMOUNT_VALUE, MSISDN, RECARGE_URL
from http_utils import post

logger = logging.getLogger(__name__)
recarga_payload = {
    "externalId": None,
    "paymentMethod": {"type": "pix"},
    "rechargeValue": {"value": AMOUNT_VALUE * 100},
    "targetMsisdn": MSISDN,
}


def make_pix_payment(session):
    import qrcode  # noqa: PLC0415
    response = post(session, RECARGE_URL, recarga_payload, "Fazer pagamento com pix")
    json_response = response.json()
    pix_code = json_response["pixCode"]
    logger.info("Qr code gerado")
    logger.info(response.json())

    img = qrcode.make(pix_code)
    img.save("pix_qrcode.png")
