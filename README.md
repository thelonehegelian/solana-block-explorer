# Solana Block Explorer (Written in Go)

**Description:**

The Solana Block Explorer is a simple utility written in Go that allows users to explore the Solana blockchain. It offers basic functionality for fetching and displaying information about Solana blocks.

**Usage:**

To run the program, execute the following command:

```shell
go run main.go
```

Upon execution, the program generates two CSV files, `blocks.csv` and `transactions.csv`, containing blockchain data.

**Features:**

- **Block Table:** The generated CSV files contain information in the following format:

  ```
  BlockHeight/Number | MinedTimeStamp | BlockHash | PreviousBlockhash | TransactionCount | Epoch(?) | BlockMiner/Validator(?)
  ```

- **Limited Data Retrieval:** The program is designed to work within API timeout constraints, so the number of blocks and transactions printed and saved is limited. Users should be aware of these limitations when using the tool.

**TODO:**

- **Enhance Data Retrieval:** Consider improving the program's data retrieval capabilities to overcome API timeout limitations and retrieve a larger set of data.
- **User Interface:** Develop a more user-friendly interface to make it easier for users to interact with and explore the Solana blockchain.
- **Filtering and Search:** Implement features that allow users to filter and search for specific blocks or transactions based on criteria of interest.
- **Data Export Options:** Provide options for users to export data in formats other than CSV, such as JSON or Excel.
- **Error Handling:** Enhance error handling to gracefully handle network issues or API-related errors.

The Solana Block Explorer project serves as a basic tool for exploring Solana blockchain data, and further enhancements can make it even more useful for users interested in analyzing Solana blockchain information.
