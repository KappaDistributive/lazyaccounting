import csv
import os
from pathlib import Path
from typing import List, NoReturn

import click
import yaml

from lazybeans.converters.converter import CsvCAMTEntryConverter, CsvKKBEntryConverter
from lazybeans.data.beancount_transaction import BeancountTransaction
from lazybeans.data.csv_camt import CsvCAMTEntry
from lazybeans.data.csv_kkb import CsvKKBEntry


CONFIG_PATH = Path(os.getenv("LAZYBEANS_CONFIG_PATH") or Path.home() / ".config/lazybeans")


@click.command()
@click.option("--input_file")
def convert_camt(input_file: str) -> None:
    with open(CONFIG_PATH / "csv_camt_entry_converter.yaml") as f:
        config = yaml.safe_load(f)
    converter = CsvCAMTEntryConverter(config)
    transactions: List[BeancountTransaction] = []

    with open(input_file, encoding="Latin-1") as csv_file:
        csv_reader = csv.DictReader(csv_file, delimiter=";", quotechar='"')
        for row in csv_reader:
            transactions.append(converter(CsvCAMTEntry.from_row(row)))

    transactions = sorted(transactions, key=lambda x: x.date)
    for transaction in transactions:
        print(transaction)
        print()


@click.command()
@click.option("--input_file")
def convert_kkb(input_file: str) -> None:
    with open(CONFIG_PATH / "csv_kkb_entry_converter.yaml") as f:
        config = yaml.safe_load(f)
    converter = CsvKKBEntryConverter(config)
    transactions: List[BeancountTransaction] = []

    with open(input_file) as csv_file:
        next(csv_file)
        csv_reader = csv.DictReader(csv_file, delimiter=";", quotechar="\"")
        for row in csv_reader:
            transaction = converter(CsvKKBEntry.from_row(row))
            if transaction is not None:
                transactions.append(transaction)

    transactions = sorted(transactions, key=lambda x: x.date)
    for transaction in transactions:
        print(transaction)
        print()



@click.group()
def cli():
    pass


if __name__ == "__main__":
    cli.add_command(convert_camt)
    cli.add_command(convert_kkb)
    cli()
