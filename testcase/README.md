Basic setting is enough to run the most of the tcs. Some tcs need additional setting.

## Basic setting
We'll use local-deploy as a basic setting.

Git clone kaiaspray somewhere else.
```
cd $HOME
git clone https://github.com/kaiachain/kaiaspray.git
cd local-deploy
```
Folow the README.md to setup the local network.

## Auction, Gasless TCs

To enable the auction & gasless module, we need:
* All CNs have enough balance (> 10 KAIA)
* All CNs (kcnd.conf) open the auction RPC namespace.
  `RPC_API=auction,db,eth,klay, ...`
* All CNs configure balance check 0.
  `ADDITIONAL="--gasless.balance-check-level 0"`

To enable the auction tcs, we need:
* to setup key as the owner of the registry
* to setup auction target tx type list among the next list
  ` --auctionTargetTxTypeList="VT,SC,rSC,GAA,GAS,rGAA,rGAS" `
