// Lib imports
const Web3 = require('web3');
const express = require('express');
const Contract = require('web3-eth-contract');
const quboRecurringJson = require('./QuboReccuring.sol/QuboRecurring.json')

// Initialisation
const app = express()
// const web3 = new Web3("https://cronos-testnet-3.crypto.org:8545/" || "ws://localhost:8545");
const web3 = new Web3(new Web3.providers.HttpProvider('https://cronos-testnet-3.crypto.org:8545/'));
const port = 3000

// set provider for all later instances to use
Contract.setProvider('https://cronos-testnet-3.crypto.org:8545/');

const contractAddress = '0xAEa7c87000aF04815e3786C87D22c09Fa9466B51';
const contract = new Contract(quboRecurringJson['abi'], contractAddress);

app.get('/', (req, res) => {
  res.send(contract);
})

app.post('/getSubscription', (req, res) => {
    contract.methods.getSubscription(req.query.customer, req.query.merchant, req.query.id).call((err, response)=> res.send(response))
})

app.post('/getSubscriptionReceipts', (req, res) => {
    contract.methods.getSubscriptionReceipts(req.query.address).call((err, response)=> res.send(response))
})

app.post('/subscriptionTimeRemaining', (req, res) => {
    contract.methods.subscriptionTimeRemaining(req.query.value).call((err, response)=> res.send(response))
})

app.post('/cancelSubscription', (req, res) => {
    contract.methods.cancelSubscription(req.query.customer, req.query.merchant, req.query.id).call((err, response)=> res.send(response))
})

app.post('/createSubscription', (req, res) => {
    contract.methods.createSubscription(req.query.id, req.query.merchant, req.query.subscriptionPrice, req.query.tokenAddress, req.query.subscriptionPeriod, req.query.periodLength).call((err, response)=> res.send(response))
})

app.post('/executePayment', (req, res) => {
    contract.methods.executePayment(req.query.customer, req.query.id).call((err, response)=> res.send(response))
})


app.listen(port, () => {
  console.log(`Example app listening on port ${port}`)
})