# Moracle

## Project Summary

Moracle is a simple, yet fast oracle client built using Monad Testnet. It takes pricing data from several different sources and bundles them into a confidence interval. Next, a protocol requests a price by sending a transaction containing a SHA256-encrypted message. The program will send back a transaction, where input data will contain an encrypted version of the requested price. The receiver will then use a private key to decode the message and read the price. 

## To-Do List

- [ ]  Get Pricing Data from Outside APIs
- [ ]  Create Confidence Interval Using Prices
- [ ]  Encrypt the Pricing Data Using SHA256
- [ ]  Connect to Monad Testnet
- [ ]  Create Recipient Client
- [ ]  Request a Price On-Chain
- [ ]  Read the Price
- [ ]  (Continuously) Compare Price Intervals with Other Oracles

## Current Issues

### Lack of “True” Price

In order to calculate the t-score for a Student t distribution, you need the implied population mean, which is the theoretical “real” price of an asset. This creates a “Chicken and the Egg” problem, where the true price is not known to create the confidence interval.

For the time being, I am using the average between the top 4 oracles by Total Value Secured  (TVS).

## Instructions

[In Progress…]
