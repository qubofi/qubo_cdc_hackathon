import React from 'react';
import { Table } from 'rsuite';

/* Styles */
import styles from './index.module.scss';
import './tableReset.scss';
import 'rsuite/dist/styles/rsuite-default.css';
import { TableProps, Column } from './types';

const CustomersTable = (props: TableProps) => {
	const { columns = [], data } = props;
	let tableColumns: Column[] = columns;
	if (columns?.length === 0) {
		tableColumns = Object.keys(data?.[0]).map((key) => {
			return {
				name: key,
				dataKey: key.replace(' ', '-'),
				type: 'string'
			};
		}) || [];
	}

	return (
		<section className={styles.campaignTableWrapper}>
			<Table
				height={400}
				id="campaignTable"
				data={data}
				hover={false}
				onRowClick={(data: any) => {
					console.log(data);
				}}
			>
				{
					tableColumns.map((col: Column) => {
					return (
						<Table.Column
							verticalAlign="middle"
							align="left"
							flexGrow={1}
						>
							<Table.HeaderCell>{col.name}</Table.HeaderCell>
							<Table.Cell dataKey={col.dataKey} />
						</Table.Column>);
					})
				}
			</Table>
		</section>
	);
};

export default CustomersTable;
