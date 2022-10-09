export type Column = {
    dataKey: string;
    name: string;
    type?: 'string' | 'number' | 'date';
};

export interface TableProps {
    columns?: Column[];
    data: {
        [key: string]: string;
    }[];
};