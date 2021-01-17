import datetime

from dataclasses import dataclass
from typing import List, Optional, Tuple

from decimal import Decimal

from lazybeans.util.money import money_fmt


@dataclass
class BeancountTransaction:
    # meta: Meta
    date: datetime.date
    flag: str
    payee: Optional[str]
    narration: str
    # tags: Set
    # links: Set
    postings: List[Tuple[str, Optional[Tuple[Decimal, str]]]]

    def __str__(self) -> str:
        string_representation: str = ""
        string_representation += datetime.datetime.strftime(self.date, "%Y-%m-%d") + " "
        string_representation += self.flag + ' "'
        string_representation += self.narration + '"\n'

        for posting in self.postings:
            account, value = posting
            posting_representation: str = "  " + account
            if value is not None:
                amount, currency = value
                value_representation = money_fmt(amount, currency)
                posting_representation += (
                    73 - len(posting_representation) - len(value_representation)
                ) * " "
                posting_representation += value_representation
            string_representation += posting_representation + "\n"

        return string_representation
