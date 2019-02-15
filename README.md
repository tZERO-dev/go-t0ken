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

#### `mainnet` configuration

```yaml
url: http://some_ethereum_node_url:8545
keystore: ./keystore
keystoreAddress: 0x0
keystoreAddressAliases:
  admin: 0x0
  broker: 0x0
  custodian: 0x0
  custodial: 0x0
  issuer: 0x0
  owner: 0x0
  tzero: 0x0

storage: 0x2d1477dd9c494e8758ec8d03f9f8b838ce394414
custodianRegistry: 0x2963488e2a140ca324e086ab8f89b5d533f1081d
brokerDealerRegistry: 0x3ecb8f0d127e22d436b26fccad4f38d7f5b91ee9
investorRegistry: 0x857f6a42634a14847cc4e0226f36906f0a77cee3
t0ken: 0x7337a2423b982e00c060d90710656714751f068e
t0kenCompliance: 0x5bd5B4e1a2c9B12812795E7217201B78C8C10b78
```

## License
This project is licensed under the [Apache 2.0](http://www.apache.org/licenses/LICENSE-2.0.html) license.

## Links
 - [tZERO's Website](https://www.tzero.com/)

