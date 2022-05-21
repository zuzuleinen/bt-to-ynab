# About

If you are using [YNAB](https://www.youneedabudget.com/) you cannot link an account
from [Banca Transilvania](https://www.bancatransilvania.ro/).

`bt-to-ynab` will help you transform a transactions CSV file from **BT** format to **YNAB** format.

# Installation

Go should be [installed and set up](https://golang.org/doc/install) on your system.

Install `bt-to-ynab`:

```shell 
go install github.com/zuzuleinen/bt-to-ynab@latest
```

# Usage

* Login into your [BT24](https://www.bt24.ro/)
* Download your account transactions from `Conturile mele` > `Cautare tranzactii` in a .CSV format
* Use

```shell
bt-to-ynab downloadedTansactionsFromBT.csv > youreNewFileForYNAB.csv
```

* Import `youreNewFileForYNAB.csv` in [YNAB](https://www.youneedabudget.com/)

## Questions or suggestions

I wrote this library in an attempt to avoid cleaning my apartment. Error handling is missing on purpose and this repo
should not be used as a reference for good code.

If you encounter a problem feel free to [open](https://github.com/zuzuleinen/bt-to-ynab/issues/new) an issue or send me
an e-mail at **andrey.boar[at]gmail.com**