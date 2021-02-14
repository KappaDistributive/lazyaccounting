from abc import ABC, abstractmethod

from lazybeans.data.beancount_transaction import BeancountTransaction
from lazybeans.data.csv_camt import CsvCAMTEntry
from lazybeans.data.csv_kkb import CsvKKBEntry


class CsvCAMTEntryConverter(ABC):
    @abstractmethod
    def __call__(self, csv_camt_entry: CsvCAMTEntry) -> BeancountTransaction:
        pass


class CsvKKBEntryConverter(ABC):
    @abstractmethod
    def __call__(self, csv_kkb_entry: CsvKKBEntry) -> BeancountTransaction:
        pass
