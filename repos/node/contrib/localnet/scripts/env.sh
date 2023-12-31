# these are the connector (MPI) contract addresses
export GOERLI_MPI_ADDRESS=0x00007d0BA516a2bA02D77907d3a1348C1187Ae62
export GOERLILOCALNET_MPI_ADDRESS=0x00007d0BA516a2bA02D77907d3a1348C1187Ae62
export BSCTESTNET_MPI_ADDRESS=0x000054d3A0Bc83Ec7808F52fCdC28A96c89F6C5c
export MUMBAI_MPI_ADDRESS=0x000054d3A0Bc83Ec7808F52fCdC28A96c89F6C5c
export ROPSTEN_MPI_ADDRESS=0x000054d3A0Bc83Ec7808F52fCdC28A96c89F6C5c
export BAOBAB_MPI_ADDRESS=0x000054d3A0Bc83Ec7808F52fCdC28A96c89F6C5c

# uncomment if you want to use pool price as oracle
#unset DUMMY_PRICE
#export GOERLI_POOL_ADDRESS=V2:0xb3a16c2b68bbb0111ebd27871a5934b949837d95:ETHZETA
#export BSCTESTNET_POOL_ADDRESS=V2:0x775fbbceffcfc5fccd999df2924039d086199222:ZETAETH
#export MUMBAI_POOL_ADDRESS=V3:0x0F65426268D78e563eaDEaa865e92C4B948dD8A1:ZETAETH
#export ROPSTEN_POOL_ADDRESS=V2:0x3b45806771fa4508f11ec1601240e81f577a9fd1:ZETAETH
export DUMMY_PRICE=yes

# this disables keygen instruction from zetacore.
export DISABLE_TSS_KEYGEN=yes

# use your own quicknode API links please; free plan is more than enough
export BSCTESTNET_ENDPOINT=https://data-seed-prebsc-1-s1.binance.org:8545
export GOERLI_ENDPOINT=https://goerli.infura.io/v3/50b6673dc48443e59047246df462902c
export GOERLILOCALNET_ENDPOINT=https://goerli.infura.io/v3/50b6673dc48443e59047246df462902c
export ROPSTEN_ENDPOINT=https://ropsten.infura.io/v3/50b6673dc48443e59047246df462902c
export MUMBAI_ENDPOINT=https://polygon-mumbai.infura.io/v3/50b6673dc48443e59047246df462902c
export BAOBAB_ENDPOINT=https://klaytn-baobab-rpc.allthatnode.com:8551
