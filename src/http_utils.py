from curl_cffi import requests

from const import HEADERS, IMPERSONETE, TIMEOUT_DEFALT


def get(session, url, descricao, params=None):
    try:
        response = session.get(
            url,
            params=params,
            headers=HEADERS,
            impersonate=IMPERSONETE,
            timeout=TIMEOUT_DEFALT,
        )
        response.raise_for_status()

    except requests.exceptions.RequestException as e:
        erro = f"Erro ao {descricao}: {e}"
        if hasattr(e, "response") and e.response is not None:
            erro += f"\nResponse: {e.response.text}"
        raise Exception(erro) from e
    return response


def post(session, url, json, descricao):
    try:
        response = session.post(
            url,
            json=json,
            headers=HEADERS,
            impersonate=IMPERSONETE,
            timeout=TIMEOUT_DEFALT,
        )
        response.raise_for_status()

    except requests.exceptions.RequestException as e:
        erro = f"Erro ao {descricao}: {e}"
        if hasattr(e, "response") and e.response is not None:
            erro += f"\nResponse: {e.response.text}"
        raise Exception(erro) from e
    return response
