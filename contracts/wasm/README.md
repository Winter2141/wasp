## Smart Contracts

Note that most example smart contracts only show the concepts involved in
creating them but should not be taken as fully functional code yet.

Sample smart contracts:

- dividend

  Allows for setting up automatic dividend dispersal to a number of recipients
  according to distribution factors per recipient. Any tokens sent to this
  contract will automatically be divided accordingly over these recipients.

- donatewithfeedback

  Allows for donations and registers feedback associated with the donation. The
  contract owner can at any point decide to withdraw donated funds from the
  contract.

- erc20

  Experimental implementation of an ERC20 smart contract as first introduced by
  Ethereum.

- fairauction

  Allows an auctioneer to auction a number of tokens. The contract owner takes a
  small fee. The contract guarantees that the tokens will be sent to the highest
  bidder, and that the losing bidders will be completely refunded. Everyone
  involved stakes their tokens, so there is no possibility for anyone to cheat.

- fairroulette

  A simple betting contract. Betters can bet on a random color and after a
  predefined time period the contract will automatically pay the total bet
  amount proportionally to the bet size of the winners.

- helloworld

  ISC version of the ubiquitous "Hello, world!" program.

- inccounter

  A simple test contract. All it does is increment a counter value. It is also
  used to test basic ISC capabilities, like persisting state, batching
  requests, and sending (time-locked) requests from a contract.

- testcore

  Helper smart contract to test the ISC core functionality.

- testwasmlib

  Helper smart contract to test the WasmLib functionality.

- tokenregistry

  Mints and registers colored tokens in a token registry.

### How to create your own Rust smart contracts

Prerequisites:

* install the latest Rust tools, you can
  [find them here](https://www.rust-lang.org/tools/install).
* When installing under Windows the Rust installation program may tell you that
  you need the Visual Studio C++ Build Tools, which you can
  [download here](https://visualstudio.microsoft.com/visual-cpp-build-tools/).
  Note that you only need to install the C++ build tools, which is the top-left
  selection.
* install Wasm-pack, which can be
  [downloaded here](https://rustwasm.github.io/wasm-pack/).

Building a Rust smart contract is very simple when using the Rust plugin in any
IntelliJ based development environment. Open the _contracts/wasm_ sub folder in
your IntelliJ, which then provides you with the Rust workspace.

The easiest way to create a new contract is to copy the _helloworld_ folder to a
properly named new folder within the _rust_ sub folder. Next, change the fields
in the first section of the new folder's _cargo.toml_ file to match your
preferences. Make sure the package name equals the folder name. Finally, add the
new folder to the workspace in the _cargo.toml_ in the _contracts/wasm_ folder.

To build the new smart contract select _Run->Edit Configurations_. Add a new
configuration based on the _wasmpack_ template, type the _name_ of the new
configuration, type the _command_ `build`, and select the new folder as the
_working directory_. You can now run this configuration to compile the smart
contract directly to Wasm. Once compilation is successful you will find the
resulting Wasm file in the _pkg_ sub folder of the new folder.

