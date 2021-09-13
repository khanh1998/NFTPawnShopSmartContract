# NFT Pawning Shop
This application is for people who own some NFT token, they need money but don't want to sell their's token, they can list the token in our application as collateral, and other users can give them a loan. So, the borrower got the money, their token is locked in our smart contract. When the time of repayment comes in, the borrower has to pay the original money plus interest to the lender, if they are not, then the token will be transferred to the lender.
# Run project
## 1. Install Metamask
1.1 Install Metamask extension in Chrome [here](https://metamask.io/download.html)\
1.2 After that, create an account.\
Remember to keep the mnemonic of your accounts in secret.\
Mnemonic contains 12 words and it look like: `jealous expect hundred young unlock disagree major siren surge acoustic machine catalog`\
We gonna need this mnemonic when create workspace in Ganache
## 2. Install Ganache or Ganache-cli
2.1 Install Ganache [here](https://www.trufflesuite.com/ganache) [or Ganache CLI]\
2.2 Start Ganche UI\
![Ganache UI](/images/ganache-ui.png)\
2.3 Choose `New Workspace`\
2.4 Fill in the information\
![](/images/ganache-setting-mnemonic.png)\
Just fill workspace name, and then switch to tab `Account & Key` and fill in the **mnemonic** you got from **step 1**.\
2.5 `Save workspace`\
Now you got an new workspace in Ganache, and all the accounts you see on the UI are the same with accounts in Metamask wallet.\
![](/images/ganache-workspace-info.png)

Remember the `Network ID` and `RPC Server`, because we gonna need it later.\
So this Ganache server is running on host `http://127.0.0.1` and port `7545`.

## 3. Connect Metamask to Ganache
3.1. Click on Metamask icon in Chrome\
3.2. Click button to show all availabe networks\
![](/images/metamask-networks.png)
So, our Ganache network is just like Ethereum but it run on local machine, now we want to add our network to Metamask.\
3.2. Click on `Custom RPC`\
3.3. Fill in information like bellow\
![](/images/metamask-network-info.png)
3.4. Finally connect to Ganache
## 4. Deploy smart contract to Ganache
4.1. Go to `/solidity` folter\
4.2. Run command `npm install -g truffle @truffle/hdwallet-provider dotenv`\
4.3. Run command `truffle migrate --network ganache --reset`\
4.4. You can view some information about the smart contracts in `Contracts` tab.\
Do you see the address of the `PawningContract`, keep it, because we need it in the later step.\
![](/images/ganache-contracts.png)
## 5. Install Go, NodeJS
Go [here](https://golang.org/doc/install)
## 6. Run api
6.1 Go to `/api` folder\
6.2 Run command `go mod download` to install packages\
6.3 Update mongodb uri in `app.env`\
Add file `app.env` to `/api` folder, contains bellow content:
>MONGODB_URI=mongodb+srv://username:password@cluster.v5cg7.azure.mongodb.net/databaseName?retryWrites=true&w=majority
>HOST=localhost:4000
>DATABASE_NAME=cooking_recipe
>SYMMETRIC_KEY=this is my secret symmetric keya
>TOKEN_DURATION=15m

6.4 Run command `go run .`
## 7. Run event listener
7.1 Go to `/event_listener` folder\
7.2 Run command `go mod download` to install packages\
7.3 Update address of `PawningShop` contract to `dev.env`, because your contract address is changed when deployed\
Add a new `dev.env`to `/event_listener` folder, contains bellow content:
>API_HOST=pawningshop:4000
>PAWN_PATH=/pawns
>BID_PATH=/bids
>BID_PAWN_PATH=/bids-pawns
>NOTIFY_HOST=http://localhost:7789
>NOTIFICATION_PATH=/notifications
>NETWORK_ADDRESS=ws://localhost:8545
>CONTRACT_ADDRESS=0xF8eC32B1884F17275aEed75DFd877DAd54ab2Ce4

CONTRACT_ADDRESS is address of `PawningShop` contract.

7.4 Because this project using `go-ethereum` library, and it requires `gcc`, `g++` and `make`. In Ubuntu, these tools are available by default, but in Windows you have to install those tools by yourself.
7.5 Run command `ENV=DEV go run .`

## 8. Run notify
8.1 Go to `/notify` folder
8.2 Run command `npm install`
8.3 Run command `npm run dev`

## 9. Run UI
9.1 Go to `/client` folder\
9.2 Run command `npm install`\
9.3 Run command `npm run serve`
