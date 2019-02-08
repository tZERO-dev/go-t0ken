[<img src="https://storage.googleapis.com/media.tzero.com/t0ken/logo.png" width="400px" />](https://www.tzero.com/)

###### *Note - This is experimental/alpha code at the moment that is still in active development, use at your own risk.

---

# T0ken CLI for the tZERO Suite of Smart-Contracts.

This project is aimed at developers who want to interact/deploy/experiment with the tZERO T0ken related smart contracts,
but don't want to spend a lot of ramp up time with all of the related tooling.

## Installation

Download a binary for your OS from the [releases](https://github.com/tZERO-dev/go-t0ken/releases) page, and place the binary somewhere within your path.

To build from source, you'll want to have [Go](https://golang.org/) 1.11.5+ installed.  
Next, you should be able to do a `go install`:

```bash
% go get github.com/tzero-dev/go-t0ken/...
```

## Configuration

To avoid passing in repetitious flags for your Ethereum node URL, contract address, etc., you can place them
into `t0ken.yaml`, which will be read from `CWD` for each run.


Here's a minimal configuration:

```yaml
url: http://some_ethereum_node_url:8545
t0ken: 0x5bd5B4e1a2c9B12812795E7217201B78C8C10b78
```

A more extensive configuration, including keystore addresses for performing transactions, and
a `keystoreAddressAliases` section allowing you to alias addresses by name:

eg.
```sh
% t0ken token issueTokens 500 --keystoreAddress 0x0000000000000000000000000000000000000000
  becomes
% t0ken token issueTokens 500 --keystoreAddress issuer
```

```yaml
url: http://some_ethereum_node_url:8545
keystore: ./keystore
keystoreAddress: 0x0000000000000000000000000000000000000000
keystoreAddressAliases:
  issuer: 0x0000000000000000000000000000000000000000
  owner: 0x0000000000000000000000000000000000000000

t0ken: 0x5bd5B4e1a2c9B12812795E7217201B78C8C10b78
```

## License
This project is licensed under the [Apache 2.0](http://www.apache.org/licenses/LICENSE-2.0.html) license.

## Links
 - [tZERO's Website](https://www.tzero.com/)

