import React from 'react';
import { authenticator } from "../_services/authenticator";
import { GoogleLogin } from 'react-google-login';
import { Jumbotron } from "react-bootstrap";
class LoginPage extends React.Component {
    constructor(props) {
        super(props);
        this.authenticate = this.authenticate.bind(this);
        if (authenticator.currentUser.value) {
            this.props.history.push("/");
        }
    }

    authenticate(googleResp) {
        authenticator.setUserEmail(googleResp.profileObj.email);
        authenticator.setUserToken(googleResp.tokenId)
        this.props.history.push("/");
        // this.echoGoogleAuth(googleResp)
    }

    echoGoogleAuth = (response) => {
        console.log(response);
    }

    render() {
        return (
            <div className="d-flex">
            <div className="d-sm-inline-flex p-4 mr-auto ml-auto">
                <Jumbotron>
                    <img src="https://image.flaticon.com/icons/svg/1412/1412542.svg" width="50" alt="" /> 
                    <h2 className="display-4">Vineguard</h2>
                    <hr className="my-2" />
                    <GoogleLogin
                    clientId={process.env.REACT_APP_GOOGLE_OAUTH_CLIENT_ID}
                    buttonText="Login"
                    onSuccess={this.authenticate}
                    cookiePolicy={'single_host_origin'}
                    />
                </Jumbotron>
            </div>
            </div>
        );
    }
}

export { LoginPage }
