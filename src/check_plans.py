from const import VALUES_URL
from http_utils import get


def best_plain(session):
    response = get(session, VALUES_URL, "Procurando o melhor plano")
    melhor_custo_dia = None
    menor_custo_dia = float("inf")

    for plano in response.json():
        dias = plano["expiresIn"]
        valor = plano["value"]["amount"] / 100
        custo_por_dia = valor / dias

        if custo_por_dia < menor_custo_dia:

            menor_custo_dia = custo_por_dia
            melhor_custo_dia = {
                "valor": valor,
                "dias": dias,
                "custo_dia": custo_por_dia,
                "id": plano["id"],
            }
    return melhor_custo_dia
