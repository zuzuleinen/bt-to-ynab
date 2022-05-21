# About

If you are using [YNAB](https://www.youneedabudget.com/) you cannot link an account
from [Banca Transilvania](https://www.bancatransilvania.ro/).

`bt-to-ynab` will help you transform a transactions CSV file from BT format to YNAB format.

# Installation

```shell 
go install github.com/zuzuleinen/bt-to-ynab
```

# Usage

* Login into your [BT24](https://www.bt24.ro/)
* Download your transactions history from Conturile mele > Cautare tranzactii in .CSV format
* Use

```shell
bt-to-ynab downloadedTansactionsFromBT.csv > youreNewFileForYNAB.csv
```

* Import `youreNewFileForYNAB.csv` in [YNAB](https://www.youneedabudget.com/)