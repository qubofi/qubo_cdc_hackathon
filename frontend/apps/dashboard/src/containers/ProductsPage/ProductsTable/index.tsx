import React from 'react';
import { Table } from 'rsuite';
import { MEDIUM_WIDTH, SMALL_WIDTH } from '../../../DATA/table';

/* Styles */
import styles from './index.module.scss';
import './tableReset.scss';
import 'rsuite/dist/styles/rsuite-default.css';

interface Props {
	campaigns: any[];
}

const CustomersTable = (props: Props) => {
	const { campaigns } = props;
	return (
		<section className={styles.campaignTableWrapper}>
			<Table
				height={400}
				id="campaignTable"
				data={campaigns}
				hover={false}
				onRowClick={(data: any) => {
					console.log(data);
				}}
			>
				<Table.Column
					width={SMALL_WIDTH}
					fixed
					verticalAlign="middle"
					align="left"
				>
					<Table.HeaderCell>Id</Table.HeaderCell>
					<Table.Cell dataKey="id" />
				</Table.Column>

				<Table.Column width={MEDIUM_WIDTH} verticalAlign="middle" align="left" flexGrow={1}>
					<Table.HeaderCell>Wallet Address</Table.HeaderCell>
					<Table.Cell dataKey="wallet_addr" />
				</Table.Column>
			</Table>
		</section>
	);
};

export default CustomersTable;
