/* @refresh reload */
import { render } from 'solid-js/web'
import { Route, Router } from "@solidjs/router";
import App from './app';
import './index.css'
import Dashboard from './pages/dashboard';
import NotFound from './pages/notfound';

const root = document.getElementById('root')

render(() => (
    <Router root={App}>
        <Route path="/" component={Dashboard} />
        <Route path="*paramName" component={NotFound} />
    </Router>
), root!);
