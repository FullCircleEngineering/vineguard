import React from 'react';
import { Navbar, Container } from 'react-bootstrap';
import {
    Router,
    Route,
    Redirect
} from 'react-router-dom';
import { createBrowserHistory } from 'history';
import { authenticator } from "../_services/authenticator";
import { Dashboard } from "../Dashboard/Dashboard";
import { LoginPage } from "../LoginPage/LoginPage";
import { GoogleLogout } from 'react-google-login';

const hist = createBrowserHistory();

const CustomRoute = ({ component: Component, ...rest }) =>
    <Route {...rest} render={props => {
        const currentUser = authenticator.currentUser.value;
        if (!currentUser) {
            return <Redirect to={"/login"}/>
        }
        return <Component {...props} />
    }} />
;

class App extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            currentUser : null
        };
    }

    componentDidMount() {
        authenticator.currentUser.subscribe(msg => this.setState({currentUser: msg}));
    }

    logout() {
        authenticator.logout();
    }

    render() {
        return (
            <Router history={hist}>
                <div>
                    {this.state.currentUser &&
                    <Navbar bg="dark" variant="dark">
                        <Navbar.Brand>
                            <img
                                src="https://image.flaticon.com/icons/svg/1412/1412542.svg"
                                width="30"
                                height="30"
                                className="d-inline-block align-top"
                                alt=""
                            /> {"  Vineguard"}
                        </Navbar.Brand>
                        <Container className="justify-content-end">
                        <GoogleLogout
                        clientId={process.env.REACT_APP_GOOGLE_OAUTH_CLIENT_ID}
                        buttonText="Logout"
                        onLogoutSuccess={this.logout}
                        icon="false"
                        >
                        </GoogleLogout>
                        </Container>
                    </Navbar>
                    }
                    <CustomRoute path={"/"} component={Dashboard} />
                    <Route path={"/login"} component={LoginPage} />
                </div>
            </Router>
        )
    }
}

export { App };