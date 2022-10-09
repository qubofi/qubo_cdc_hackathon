import React from 'react';
import "@celo-tools/use-contractkit/lib/styles.css";
import "@rainbow-me/rainbowkit/styles.css";
import Head from "next/head";
import { SnackbarProvider } from "notistack";
import { ApolloProvider } from "@apollo/client";
import client from "@/apollo-client";
import {
  Chain,
  getDefaultWallets,
  RainbowKitProvider,
} from "@rainbow-me/rainbowkit";
import { chain, configureChains, createClient, WagmiConfig } from "wagmi";
import { jsonRpcProvider } from 'wagmi/providers/jsonRpc';
const Contract = require('web3-eth-contract');

function MyApp({ Component, pageProps, router }: AppProps): React.ReactElement {
  const avalancheChain: Chain = {
    id: 338,
    name: 'Cronos Testnet',
    network: 'cronos',
    iconUrl: 'https://example.com/icon.svg',
    iconBackground: '#fff',
    nativeCurrency: {
      decimals: 18,
      name: 'Cronos',
      symbol: 'TCRO',
    },
    rpcUrls: {
      default: 'https://evm-t3.cronos.org	',
    },
    blockExplorers: {
      default: { name: 'SnowTrace', url: 'https://snowtrace.io' },
      etherscan: { name: 'SnowTrace', url: 'https://snowtrace.io' },
    },
    testnet: true,
  };
  
  const { provider, chains } = configureChains(
    [avalancheChain],
    [jsonRpcProvider({ rpc: chain => ({ http: chain.rpcUrls.default }) })]
  );
  
  const { connectors } = getDefaultWallets({
    appName: "My RainbowKit App",
    chains
  });
  
  const wagmiClient = createClient({
    autoConnect: true,
    connectors,
    provider
  });
  
  return (
    <>
      <Head>
        <title>Qubo</title>
        <meta name="description" content="Celo DApp Starter" />
        <meta name="viewport" content="initial-scale=1, width=device-width" />
        <link
          rel="stylesheet"
          href="https://fonts.googleapis.com/css?family=Roboto:300,400,500,700&display=swap"
        />
      </Head>
      <WagmiConfig client={wagmiClient}>
        <RainbowKitProvider chains={chains}>
          <SnackbarProvider
            maxSnack={3}
            anchorOrigin={{
              vertical: "bottom",
              horizontal: "right",
            }}
          >
            <ApolloProvider client={client}>
              <div suppressHydrationWarning>
                {typeof window === "undefined" ? null : (
                  <Component {...pageProps} />
                )}
              </div>
            </ApolloProvider>
          </SnackbarProvider>
        </RainbowKitProvider>
      </WagmiConfig>
    </>
  );
}

export default MyApp;
