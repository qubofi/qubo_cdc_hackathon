// Base imports
import ReactDOM from 'react-dom';
import BaseRouter from './routes';

// Import css to remove margin
import './index.scss';

const app = (<BaseRouter/>);

ReactDOM.render(app, document.getElementById('root'));
