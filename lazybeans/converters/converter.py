import datetime
import re
from typing import Any, Dict, List, Optional, Tuple

from decimal import Decimal

from lazybeans.data.beancount_transaction import BeancountTransaction
from lazybeans.data.csv_camt import CsvCAMTEntry
from lazybeans.data.csv_kkb import CsvKKBEntry
from lazybeans.util.logging import init_logger


class CsvCAMTEntryConverter:
    def __init__(self, config: Dict[str, Any]):
        self.config = config
        self.logger = init_logger(self.__class__.__name__)

    def __call__(self, csv_camt_entry: CsvCAMTEntry) -> BeancountTransaction:
        date: datetime.date = csv_camt_entry.booking_date
        payee: Optional[str] = None
        narration: str = f"{csv_camt_entry.purpose};{csv_camt_entry.posting_text}"
        complete_entry, postings = self.calculate_postings(csv_camt_entry)
        flag = "*" if complete_entry else "!"

        return BeancountTransaction(
            date=date,
            flag=flag,
            payee=payee,
            narration=narration,
            postings=postings
        )

    def calculate_account(
        self,
        csv_camt_entry: CsvCAMTEntry
    ) -> Tuple[bool, str, str, Optional[Dict[str, Any]]]:
        for rule in self.config["calculate_account"]:
            field = getattr(csv_camt_entry, rule["regex_field"])
            regex_operation = rule.get("regex_operation", "n/a")
            if rule["regex_operation"] == "match":
                if re.match(rule["regex"], field):
                    return (
                        rule["complete"],
                        rule["source_account"],
                        rule.get("target_account", None) or self.config.get("default_target_account", "<tbd>"),
                        rule
                    )

            else:
                raise NotImplementedError(f"regex_operation `{regex_operation}` is not implemented.")
        return False, "<tbd>", self.config.get("default_target_account", "<tbd>"), None

    def calculate_postings(
        self,
        csv_camt_entry: CsvCAMTEntry
    ) -> Tuple[bool, List[Tuple[str, Optional[Tuple[Decimal, str]]]]]:
        complete_entry, source_account, target_account, _ = self.calculate_account(csv_camt_entry)
        amount_display = (csv_camt_entry.amount, csv_camt_entry.currency)
        postings: List[Tuple[str, Optional[Tuple[Decimal, str]]]] = [
            (source_account, (-1 * amount_display[0], amount_display[1])),
            (target_account, amount_display),
        ]
        return complete_entry, postings


class CsvKKBEntryConverter:
    def __init__(self, config: Dict[str, Any]):
        self.config = config
        self.logger = init_logger(self.__class__.__name__)


    def __call__(self, csv_kkb_entry: CsvKKBEntry) -> Optional[BeancountTransaction]:
        date: datetime.date = csv_kkb_entry.booking_date
        payee: Optional[str] = None
        narration: str = f"{csv_kkb_entry.sales_category}:{csv_kkb_entry.vendor_name}"
        complete_entry, postings = self.calculate_postings(csv_kkb_entry)
        flag = "*" if complete_entry else "!"

        return BeancountTransaction(
            date=date,
            flag=flag,
            payee=payee,
            narration=narration,
            postings=postings
        )

    def calculate_account(
        self,
        csv_kkb_entry: CsvKKBEntry
    ) -> Tuple[bool, str, str, Optional[Dict[str, Any]]]:
        for rule in self.config["calculate_account"]:
            field = getattr(csv_kkb_entry, rule["regex_field"])
            regex_operation = rule.get("regex_operation", "n/a")
            if rule["regex_operation"] == "match":
                if re.match(rule["regex"], field):
                    return (
                        rule["complete"],
                        rule["source_account"],
                        rule.get("target_account", None) or self.config.get("default_target_account", "<tbd>"),
                        rule
                    )
            else:
                raise NotImplementedError(f"regex_operation `{regex_operation}` is not implemented.")
        return False, "<tbd>", self.config.get("default_target_account", "<tbd>"), None

    def calculate_postings(
        self,
        csv_kkb_entry: CsvKKBEntry
    ) -> Tuple[bool, List[Tuple[str, Optional[Tuple[Decimal, str]]]]]:
        complete_entry, source_account, target_account, rule = self.calculate_account(csv_kkb_entry)
        if rule is not None :
            amount_display: Optional[Tuple[Decimal, str]] = None
            if rule["amount_field"] == "amount_foreign_currency":
                if (csv_kkb_entry.amount_foreign_currency is not None and
                    csv_kkb_entry.unit_foreign_currency is not None):
                    amount_display = (csv_kkb_entry.amount_foreign_currency, csv_kkb_entry.unit_foreign_currency)
            else:
                if csv_kkb_entry.amount_in_euro is not None:
                    amount_display = (csv_kkb_entry.amount_in_euro, "EUR")

            postings: List[Tuple[str, Optional[Tuple[Decimal, str]]]] = [
                (source_account, amount_display),
                (target_account, (-1 * amount_display[0], amount_display[1]) if amount_display is not None else None),
            ]

            return complete_entry, postings
        else:
            self.logger.error("Failed to process KKB Entry:")
            self.logger.error(csv_kkb_entry)
            return False, []
