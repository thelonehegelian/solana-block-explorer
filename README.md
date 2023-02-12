Initial Plan: 
---
1. Get the latest Epoch
2. Get All the blocks in the latest Epoch
3. Get all the transactions in a single Block
   1. Get input and output of transactions (functions invoked?)
   2. Smart Contracts they interacted with
4. Find unique addresses?
5. 

Notes:

- First block in an Epoch does not have transactions

Block Table
---
BlockHeight/Number | MinedTimeStamp | BlockHash | PreviousBlockhash | TransactionCount | Epoch(?) | BlockMiner/Validator(?)


Database Structure: 
--- 
**Transactions table:**
- transaction_id: a unique identifier for the transaction
- block_id: the id of the block in which the transaction was included
- timestamp: the time at which the transaction was included in the block
- from_address: the address sending the transaction
- to_address: the address receiving the transaction
- value: the amount of tokens or data being transferred in the transaction
- fee: the fee associated with the transaction
- data: any additional data associated with the transaction

**Blocks table:**
- block_id: a unique identifier for the block
- block_hash: the hash of the block
- previous_block_hash: the hash of the previous block
- timestamp: the time at which the block was created
- validator: the address of the validator who added the block to the blockchain

**Addresses table:**
- address: a unique identifier for the address
- balance: the current balance of the address
- transaction_count: the number of transactions associated with the address

**Tokens table:**
- token_id: a unique identifier for the token
- symbol: the symbol or abbreviated name of the token
- name: the full name of the token
- total_supply: the total number of tokens in existence
- decimals: the number of decimal places for the token


Transaction relationships: You could model the relationships between transactions by creating nodes for each transaction and connecting them based on the inputs and outputs of the transactions.

Address relationships: You could model the relationships between addresses by creating nodes for each address and connecting them based on the transactions that they are involved in.

Token relationships: If your Solana graph network includes token transfers, you could model the relationships between tokens by creating nodes for each token and connecting them based on the transfers of the tokens.

Block relationships: You could model the relationships between blocks by creating nodes for each block and connecting them based on the previous block hashes.