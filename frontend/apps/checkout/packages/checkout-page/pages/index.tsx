import * as React from "react";
import Web3 from 'web3';
import styles from './index.module.scss';
import deployedContracts from "../../hardhat/deployments/hardhat_contracts.json";
import ConnectWalletButton from '../components/ConnectWalletButton'
import ProductsDisplay from "../components/ProductsDisplay";
import Checkout from "../components/Checkout";
import quboRecurringJson from '../QuboReccuring.sol/QuboRecurring.json';

const Contract = require('web3-eth-contract');

interface TabPanelProps {
  children?: React.ReactNode;
  index: number;
  value: number;
}

const PRODUCTS = [
  {
    name: "Monkey #1000 Subscription",
    price: 100,
    currency: "TCRO / month",
    desc: 'Amazing monkey JPEG'
  },
]
export default function App() {
  const [value, setValue] = React.useState(0);

  const handleChange = (event: React.SyntheticEvent, newValue: number) => {
    setValue(newValue);
  };

  const web3 = new Web3(new Web3.providers.HttpProvider('https://cronos-testnet-3.crypto.org:8545/'));
  const port = 3000
  Contract.setProvider('https://cronos-testnet-3.crypto.org:8545/');
  const contractAddress = '0xBEf683217707f38deC8B6e3374F0626CEc5ad298';
  const contract = new Contract(quboRecurringJson['abi'], contractAddress)
  return (
    <div className={styles.checkoutWrapper}>
				<div className={styles.backgroundWrapper}>
					<div className={styles.background}>
						<div className={styles.background__middle1}/>
						<div className={styles.background__teal2Wrapper}>
							<div className={styles.background__teal2}/>
						</div>
						<div className={styles.background__middle3}/>
						<div className={styles.background__teal4Wrapper}>
							<div className={styles.background__teal4}/>
						</div>
					</div>
				</div>
      <ProductsDisplay products={PRODUCTS}/>
      <Checkout products={PRODUCTS} contract={contract}/>
    </div>
  );
}
