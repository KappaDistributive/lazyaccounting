import datetime
from dataclasses import dataclass
from typing import Dict, Optional

from decimal import Decimal


@dataclass
class CsvKKBEntry:
    credit_card_number: str
    transaction_date: datetime.date
    booking_date: datetime.date
    vendor_name: str
    sales_category: str
    amount_foreign_currency: Optional[Decimal]
    unit_foreign_currency: Optional[str]
    converion_rate: Optional[Decimal]
    amount_in_euro: Optional[Decimal]
    amazon_points: Decimal
    prime_points: Decimal

    @staticmethod
    def from_row(row: Dict[str, str]) -> "CsvKKBEntry":
        return CsvKKBEntry(
            credit_card_number=row["Kreditkartennummer (teilmaskiert)"],
            transaction_date=datetime.datetime.strptime(
                row["Transaktionsdatum"], "%d.%m.%Y"
            ).date(),
            booking_date=datetime.datetime.strptime(
                row["Buchungsdatum"], "%d.%m.%Y"
            ).date(),
            vendor_name=row["Händler (Name, Stadt & Land)"],
            sales_category=row["Umsatzkategorie"],
            amount_foreign_currency=Decimal(
                row["Betrag in Fremdwährung"].replace(",", ".")
            ) if row["Betrag in Fremdwährung"] != "" else None,
            unit_foreign_currency=row["Einheit Fremdwährung"]
            if row["Betrag in Fremdwährung"] != "" else None,
            converion_rate=Decimal(
                row["Umrechnungskurs"].replace(",", ".")
            ) if row["Umrechnungskurs"] != "" else None,
            amount_in_euro=Decimal(
                row["Betrag in Euro"].replace(",", ".")
            ) if row["Betrag in Euro"] != "" else None,
            amazon_points=Decimal(
                row["Amazon Punkte"] if row["Amazon Punkte"] != "" else 0
            ),
            prime_points=Decimal(
                row["Prime Punkte"] if row["Prime Punkte"] != "" else 0
            ),
        )
