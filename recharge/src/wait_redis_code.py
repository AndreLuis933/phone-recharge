import time

import redis


def wait_for_code(timeout=300):
    r = redis.Redis(host="redis", port=6379, decode_responses=True)
    start = time.time()

    while True:
        code = r.get("vivo:code")

        if code:
            r.delete("vivo:code")
            return code

        # Verifica timeout
        if time.time() - start > timeout:
            return None

        time.sleep(0.5)



