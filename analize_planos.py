import sys

from curl_cffi import requests

url = "https://recarga-api.vivo.com.br/recharge-values"

headers = {
    "Accept": "application/json, text/plain, */*",
    "Accept-Language": "en-US,en;q=0.9,pt-BR;q=0.8,pt;q=0.7",
    "Channel": "VIVO_WEB",
    "Origin": "https://recarga.vivo.com.br",
    "Referer": "https://recarga.vivo.com.br/",
}

try:
    response = requests.get(
        url,
        headers=headers,
        impersonate="chrome120",
        timeout=10,
    )
    response.raise_for_status()
except requests.exceptions.RequestException as e:
    print(f"❌ Erro ao buscar planos: {e}")
    sys.exit(1)

planos = response.json()

# Análise: menor custo por dia
melhor_custo_dia = None
menor_custo_dia = float("inf")

# Análise: melhor custo-benefício (considerando duração)
analise_completa = []

print("=" * 70)
print("📊 ANÁLISE DE PLANOS VIVO")
print("=" * 70)
print()

for plano in planos:
    dias = plano["expiresIn"]
    valor = plano["value"]["amount"] / 100
    custo_por_dia = valor / dias

    # Guardar análise completa
    analise_completa.append(
        {
            "valor": valor,
            "dias": dias,
            "custo_dia": custo_por_dia,
            "bonus": plano.get("bonusAmount", {}).get("amount", 0) / 100 if plano.get("bonusAmount") else 0,
            "descricao": plano.get("description", ""),
        },
    )

    # Encontrar menor custo por dia
    if custo_por_dia < menor_custo_dia:
        menor_custo_dia = custo_por_dia
        melhor_custo_dia = {
            "valor": valor,
            "dias": dias,
            "custo_dia": custo_por_dia,
        }

# Ordenar por custo/dia
analise_completa.sort(key=lambda x: x["custo_dia"])

print("🏆 MELHOR CUSTO POR DIA:")
print(f"   Valor: R$ {melhor_custo_dia['valor']:.2f}")
print(f"   Validade: {melhor_custo_dia['dias']} dias")
print(f"   Custo/dia: R$ {melhor_custo_dia['custo_dia']:.4f}")
print()

print("📋 RANKING DE PLANOS (do melhor para o pior custo/dia):")
print("-" * 70)
print(f"{'Valor':<12} {'Dias':<8} {'R$/dia':<12} {'Bônus':<10} {'Descrição':<20}")
print("-" * 70)

for i, plano in enumerate(analise_completa, 1):
    marca = "🥇" if i == 1 else "🥈" if i == 2 else "🥉" if i == 3 else f"{i:2d}."

    bonus_str = f"R$ {plano['bonus']:.2f}" if plano["bonus"] > 0 else "-"
    desc = plano["descricao"][:18] + "..." if len(plano["descricao"]) > 18 else plano["descricao"]

    print(
        f"{marca} R$ {plano['valor']:6.2f}  {plano['dias']:3d}d    R$ {plano['custo_dia']:.4f}   {bonus_str:<10} {desc}"
    )

print("-" * 70)
print()

# Comparação: quanto você economiza escolhendo o melhor
pior_plano = analise_completa[-1]
economia_mensal = (pior_plano["custo_dia"] - melhor_custo_dia["custo_dia"]) * 30

print(f"💰 ECONOMIA MENSAL:")
print(f"   Escolhendo o melhor plano ao invés do pior:")
print(f"   → Economia de R$ {economia_mensal:.2f} por mês")
print(f"   → Economia de R$ {economia_mensal * 12:.2f} por ano")

print()
print("=" * 70)
