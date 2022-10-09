import * as React from 'react';
import Nav from '../Nav';
import BreadCrumb from '../BreadCrumb';
import './index.scss';

interface Props {
	title?: string;
	callToAction?: any;
	breadcrumb?: any;
}

const PageHeader = (props: Props) => {
	const { title, callToAction, breadcrumb } = props;

	return (
		<>
			<title>Qubo - {title}</title>
 
			<div className="page-header">
				<p> Qubo </p>
				<Nav />
				{breadcrumb && breadcrumb.length && (
					<BreadCrumb sections={breadcrumb} />
				)}
				<div className="page-header__left">
					<p>{callToAction}</p>
				</div>
			</div>
		</>
	);
};

export default PageHeader;
