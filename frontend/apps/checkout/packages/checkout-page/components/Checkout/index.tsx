import * as React from "react";
import { Button, Divider, Grid, Typography, Link } from "@mui/material";

import web3 from "web3";
import {
  ConnectButton
} from "@rainbow-me/rainbowkit";
import axios from "axios";
import { useEffect, useState } from "react";
import { SnackbarAction, useSnackbar } from "notistack";
import { useQuery, gql } from "@apollo/client";
import styles from "./index.module.scss";
import { useAccount, useSignMessage, useNetwork , usePrepareContractWrite, useContractWrite } from 'wagmi';

// Example GraphQL query for the Storage contract updates
const QUERY = gql`
  query Updates {
    updates(orderBy: timestamp, orderDirection: desc, first: 5) {
      id
      number
      sender
      timestamp
    }
  }
`;

export default function Checkout({ products, contract }) {
  const { enqueueSnackbar, closeSnackbar } = useSnackbar();
  const getSubscription = () => contract.methods.getSubscription(
    "0x24C8Cd92C9F2025a6d293D7706307Bfed6d4d426", 
    "0x24C8Cd92C9F2025a6d293D7706307Bfed6d4d426",
    0
  ).call().then(
    (res: any) => {
      console.log(res)
    }
  );
  // 0x24C8Cd92C9F2025a6d293D7706307Bfed6d4d426
  const { isConnected, address } = useAccount();
  const purchaseSubscription = () => contract.methods.createSubscription(
    1, // Subscription ID
    "0x24C8Cd92C9F2025a6d293D7706307Bfed6d4d426", // Merchant Address
    2, // Subscription Price
    "0x0000000000000000000000000000000000000000", // Token Address
    2, // Subscription Period (e.g. 2 months)
    2630000 // One month
  ).send({from: address}).then((res:any) => {
    console.log(res);
    enqueueSnackbar("SUCCESS???", { variant: 'success' });
  }).catch((error) => {
    console.log(error);
    enqueueSnackbar("ERROR", { variant: 'error' });
  });
  const purchaseSubscriptionDemo = async () => {
    setTimeout(() => {
      enqueueSnackbar("Successfully Subscribed!", { variant: 'success' });
    }, [5000])
  }

  return (
    <div className={styles.checkoutWrapper}>
        <h2>
          Connect Wallet
        </h2>
        <ConnectButton/>
        <Divider component="div" sx={{ m: 1 }} />
        <Button 
          sx={{ m: 1, marginLeft: 0 }}
          variant="contained" 
          onClick={purchaseSubscriptionDemo}
          className={styles.checkoutButton}
        >
          Pay
        </Button>
    </div>
  );
}
