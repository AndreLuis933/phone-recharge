import json
from pathlib import Path

from curl_cffi import requests

from const import HEADERS, PLAN_TYPE

COOKIES_FILE = Path("cookies_session.json")


def save_cookies(session: requests.Session):
    cookies_dict = session.cookies.get_dict()
    COOKIES_FILE.write_text(json.dumps(cookies_dict, indent=2))


def load_cookies(session: requests.Session) -> bool:
    if not COOKIES_FILE.exists():
        return False

    try:
        cookies_dict = json.loads(COOKIES_FILE.read_text())
        session.cookies.update(cookies_dict)
    except Exception as e:
        print(f"⚠️  Erro ao carregar cookies: {e}")
        return False
    return True


def validate_session(session: requests.Session) -> bool:
    try:
        response = session.get(
            PLAN_TYPE,
            headers=HEADERS,
            impersonate="chrome120",
            timeout=10,
        )
        is_valid = response.status_code < 400

        if is_valid:
            print(f"✅ Sessão válida (Status: {response.status_code})")
        else:
            print(f"❌ Sessão inválida (Status: {response.status_code})")

    except requests.exceptions.RequestException as e:
        print(f"❌ Erro ao validar sessão: {e}")
        return False
    return is_valid
