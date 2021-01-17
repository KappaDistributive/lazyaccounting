import datetime
from dataclasses import dataclass
from typing import Dict

from decimal import Decimal


@dataclass
class CsvCAMTEntry:
    order_account: str
    booking_date: datetime.date
    value_date: datetime.date
    posting_text: str
    purpose: str
    creditor_id: str
    mandate_reference: str
    customer_reference: str
    collector_reference: str
    direct_debit_amount: str
    reverse_direct_debit_amount: str
    recipient_or_payer: str
    iban: str
    bic: str
    amount: Decimal
    currency: str
    info: str

    @staticmethod
    def from_row(row: Dict[str, str]) -> "CsvCAMTEntry":
        return CsvCAMTEntry(
            order_account=row["Auftragskonto"],
            booking_date=datetime.datetime.strptime(
                row["Buchungstag"], "%d.%m.%y"
            ).date(),
            value_date=datetime.datetime.strptime(
                row["Valutadatum"], "%d.%m.%y"
            ).date(),
            posting_text=row["Buchungstext"],
            purpose=row["Verwendungszweck"],
            creditor_id=row["Glaeubiger ID"],
            mandate_reference=row["Mandatsreferenz"],
            customer_reference=row["Kundenreferenz (End-to-End)"],
            collector_reference=row["Sammlerreferenz"],
            direct_debit_amount=row["Lastschrift Ursprungsbetrag"],
            reverse_direct_debit_amount=row["Auslagenersatz Ruecklastschrift"],
            recipient_or_payer=row["Beguenstigter/Zahlungspflichtiger"],
            iban=row["Kontonummer/IBAN"],
            bic=row["BIC (SWIFT-Code)"],
            amount=Decimal(row["Betrag"].replace(",", ".")),
            currency=row["Waehrung"],
            info=row["Info"],
        )
