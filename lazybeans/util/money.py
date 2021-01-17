from typing import List

from decimal import Decimal


def money_fmt(decimal: Decimal, currency: str, decimal_point: str = ".", positive: str = "", negative: str = "-", places=2) -> str:
    q = Decimal(10) ** -places
    sign, digits, exp = decimal.quantize(q).as_tuple()
    digits_str = list(map(str, digits))
    result: List[str] = [currency, " "]

    for _ in range(places):
        result.append(digits_str.pop() if digits_str else "0")
    result.append(decimal_point)
    if not digits_str:
        result.append("0")
    while digits_str:
        result.append(digits_str.pop())
    result.append(negative if sign else positive)

    return "".join(reversed(result))
