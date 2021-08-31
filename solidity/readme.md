# Generate pawningShop.go

On window, install solcjs and abigen

solcjs --abi ./contracts/PawningShop.sol
solcjs --bin ./contracts/PawningShop.sol
abigen --bin=__contracts_PawningShop_sol_PawningShop.bin --abi=__contracts_PawningShop_sol_PawningShop.abi --pkg=contracts --out=pawningShop.go