from abc import ABC, abstractmethod

from lazybeans.data.beancount_transaction import BeancountTransaction
from lazybeans.data.csv_camt import CsvCAMTEntry


class CsvCAMTEntryConverter(ABC):
    @abstractmethod
    def __call__(self, csv_camt_entry: CsvCAMTEntry) -> BeancountTransaction:
        pass
