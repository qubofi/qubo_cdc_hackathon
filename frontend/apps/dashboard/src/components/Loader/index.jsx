import React from 'react';
import CircularProgress from '@material-ui/core/CircularProgress';

/** Container, renders home page.
 * @memberof Container
 * @param {Component} Component1, Component1.
 * @param {ReduxAction} Action1, Action1.
 */
const Loader = () => {
	return <CircularProgress color="primary" />;	
}

export default Loader;
