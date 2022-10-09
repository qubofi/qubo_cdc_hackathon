// Basic Imports
import React from 'react';
import styles from './index.module.scss';
import history from '../../utils/History';

const Nav = () => {
	const handleChangePage = (newPage: string) => {
		history.push(newPage);
	};

	return (
		<section className={styles.brandNav}>
			<button
				className={`
					${styles.navOption}
					${history.location.pathname === '/' && styles.navOption_selected}
				`}
				onClick={() => handleChangePage('/')}
			>
				Orders
			</button>
			<button
				className={`
					${styles.navOption}
					${history.location.pathname === '/products' && styles.navOption_selected}
				`}
				onClick={() => handleChangePage('/products')}
			>
				Products
			</button>
			<button
				className={`
					${styles.navOption}
					${history.location.pathname === '/customers' && styles.navOption_selected}
				`}
				onClick={() => handleChangePage('/customers')}
			>
				Customers
			</button>
		</section>
	);
};

export default Nav;
