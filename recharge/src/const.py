import os

HEADERS = {
    "Accept": "application/json, text/plain, */*",
    "Accept-Language": "en-US,en;q=0.9",
    "Content-Type": "application/json",
    "Channel": "VIVO_WEB",
    "Origin": "https://recarga.vivo.com.br",
    "Referer": "https://recarga.vivo.com.br/",
}
TIMEOUT_DEFALT = 10
IMPERSONETE = "chrome120"

MSISDN = os.environ["MSISDN"]
AMOUNT_VALUE = int(os.environ["AMOUNT_VALUE"])
PAYMENT_METHOD = os.environ["PAYMENT_METHOD"]

SMS_URL = "https://recarga-api.vivo.com.br/sms-tokens"
LOGIN_URL = "https://recarga-api.vivo.com.br/sessions/"
VALUES_URL = "https://recarga-api.vivo.com.br/customer/recharge-values"
RECARGE_URL = f"https://recarga-api.vivo.com.br/customers/{MSISDN}/recharges"
CHECKOUT_URL = "https://recarga-api.vivo.com.br/customer/events/checkout"
PLAN_TYPE = f"https://recarga-api.vivo.com.br/customers/{MSISDN}/plan-type"
CUSTOMER_URL = "https://recarga-api.vivo.com.br/customer"
