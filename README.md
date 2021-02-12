# go-prime-api
A golang implementation of a prime number API used for recruitment.

## Getting Started

You can run via AWS Lambda and a standard http server. The http server is expected for local development where you 
don't need any special AWS integration.

1. Setup Serverless framework
2. `npm install`
3. Profit

### First Time Only

On a totally clean install to AWS, run the following to configure the domain name and associate the TLS certificate:

1. `sls create_domain`
2. Login to your serverless account and create the application as per the values at the top of `serverless.yml`
2. `sls deploy`

## Deployments

Deployed via the Serverless Framework. There's a makefile to simplify the commands. See getting started for details.

Run the following commands:
```shell
make clean
make build
make deploy
```

### Property Encryption/Decryption

Property security is handled via the use of AWS's Key Management Service (KMS). The following sections discuss how they work,
however to simplify things, two helper scripts set to use the active KMS key id exist in the `./docs` folder. Simply run:

`encrypt.sh <plain_text>`

and

`decrypt.sh <cipher_text>`

