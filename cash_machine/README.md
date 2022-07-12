You're writing software for a cash machine (ATM) in the UK.

Assume that machines have an unlimited number of coins & notes in the following denominations:

£50, £20, £10, £5, £2, £1, 50p, 20p, 10p, 5p, 2p, 1p

Write a function breakIntoChange(amount) that takes any non-negative amount in pounds and returns the minimum number of coins and notes representating that amount.

e.g.

breakIntoChange(3.45) // {'£2': 1, '£1': 1, '20p': 2, '5p': 1}

breakIntoChange(160) // {'£50': 3, '£10': 1}

breakIntoChange(0.13) // {'10p': 1, '2p': 1, '1p': 1}