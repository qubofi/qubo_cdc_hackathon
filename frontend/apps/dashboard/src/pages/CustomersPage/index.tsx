/* eslint-disable react/state-in-constructor */
// Basic Imports
import React, { useEffect } from 'react';
import axios from 'axios';
import { TabPanel } from 'react-tabs';
import Table from '../../sharedComponents/Table';
import useRequest from '../../sharedComponents/useRequest';
import { CUSTOMERS_TABLE } from '../../DATA/FAKE_DATA';

/* Styles */
import styles from './index.module.scss';

/* Components */
import PageHeader from '../../components/PageHeader';
import CTAButton from '../../components/CTAButton';
import TabsWrapper from '../../components/TabsWrapper';

const CustomersPage = () => {
	const [selectedTabList, setSelectedTabList] = React.useState<number>(0);
	const handleTabChange = (index: number) => {
		setSelectedTabList(index);
	};
	const { data, error, status } = useRequest({
		url: 'http://18.143.189.118:8080/api/',
		method: 'get'
	});
	return (
		<>
			<main className='page-container'>
				<PageHeader
					title='Customers'
					breadcrumb={[{ content: 'Home', to: '/' }, { content: 'Customers' }]}
					callToAction=''
				/>
				<div className={styles.campaignPageHeaderRow}>
					<h2 className={styles.pageTitle}> Customers </h2>
					<div className={styles.buttonsRow}>
						<CTAButton
							colorScheme='brand'
							type='secondary'
							size='tiny'
							to='/brand/campaigns/new'
							disabled
						>
							<i className='bx bxs-download'></i>
							Export
						</CTAButton>
					</div>
				</div>
				<TabsWrapper
					handleTabChange={handleTabChange}
					selectedTab={selectedTabList}
					tabsDisplayList={['All']}
				>
					<TabPanel>
						<Table 
							data={CUSTOMERS_TABLE.data}
							columns={CUSTOMERS_TABLE.columns}
						/>
					</TabPanel>
				</TabsWrapper>
			</main>
		</>
	);
};

export default CustomersPage;
