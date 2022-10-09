import { TableProps } from "../../../../features/components/Table/types";

export const PRODUCTS_TABLE: TableProps = {
    columns: [
        {
            dataKey: "productName",
            name: "Product Name",
            type: "string",
        },
        {
            dataKey: "productCost",
            name: "Product Cost",
            type: "string",
        },
        {
            dataKey: "productDescription",
            name: "Product Description",
            type: "string",
        },
        {
            dataKey: "productStatus",
            name: "Product Status",
            type: "string",
        }
    ],
    data: [
       {
            productName: "Starter Subscription",
            productCost: "$100 / month",
            productDescription: "This is a starter subscription",
            productStatus: "Active"
       },
       {
            productName: "Premium Subscription",
            productCost: "$200 / month",
            productDescription: "This is a premium subscription",
            productStatus: "Active"
        },
        {
            productName: "Enterprise Subscription",
            productCost: "$300 / month",
            productDescription: "This is an enterprise subscription",
            productStatus: "Active"
        }
    ]    
}

export const CUSTOMERS_TABLE: TableProps = {
    columns: [
        {
            dataKey: "id",
            name: "ID",
            type: "number",
        },
        {
            dataKey: "name",
            name: "Name",
            type: "string",
        },
        {
            dataKey: "walletAddress",
            name: "Wallet Address",
            type: "string",
        }
    ],
    data: [
        {
            id: "1",
            name: "John Doe",
            walletAddress: "0x1234567890",
        },
        {
            id: "2",
            name: "Jane Doe",
            walletAddress: "0x1234567890",
        },
        {
            id: "3",
            name: "John Smith",
            walletAddress: "0x1234567890",
        }
    ]
}

export const ORDERS_TABLE: TableProps = {
    columns: [
        {
            dataKey: "productName",
            name: "Subscription Name",
            type: "string",
        },
        {
            dataKey: "customerAddress",
            name: "Customer Address",
            type: "string",
        },
        {
            dataKey: "subscribedSince",
            name: "Subscribed Since",
            type: "string",
        },
        {
            dataKey: "nextPayment",
            name: "Next Payment",
            type: "string",
        }
    ],
    data: [
       {
            productName: "Starter Subscription",
            customerAddress: "0x22F2DS11330",
            subscribedSince: "2022-08-01",
            nextPayment: "2022-09-01"
         },
            {
            productName: "Premium Subscription",
            customerAddress: "0x123DAS2F9S",
            subscribedSince: "2021-01-06",
            nextPayment: "2021-02-06"
            },
            {
            productName: "Enterprise Subscription",
            customerAddress: "0xD7ABD9N30DN",
            subscribedSince: "2021-05-02",
            nextPayment: "2021-06-02"
            }
    ]    
}


export const EXPIRED_ORDERS_TABLE: TableProps = {
    columns: [
        {
            dataKey: "productName",
            name: "Subscription Name",
            type: "string",
        },
        {
            dataKey: "customerAddress",
            name: "Customer Address",
            type: "string",
        },
        {
            dataKey: "subscribedSince",
            name: "Subscribed Since",
            type: "string",
        },
        {
            dataKey: "nextPayment",
            name: "Next Payment",
            type: "string",
        }
    ],
    data: [
        {
            productName: "Starter Subscription",
            customerAddress: "0x34KNF8ABDLS",
            subscribedSince: "2021-03-17",
            nextPayment: "-"
        },
        {
            productName: "Starter Subscription",
            customerAddress: "0x1446712777",
            subscribedSince: "2021-08-15",
            nextPayment: "-"
        }
    ]    
}