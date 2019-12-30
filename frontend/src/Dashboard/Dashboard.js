import React from 'react';
import { Graph} from "../_components/Graph";
import { TimeRange } from 'pondjs';
import { Jumbotron, Form, Button } from 'react-bootstrap';
const initialTimeRange = new TimeRange([1326309060000, 1329941520000]);

class Dashboard extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            graphOne: Graph,
        };
        this.propValue = "";
    }

    render() {
        return (
            <div>
                <div className="d-lg-flex p-4">
                    {<this.state.graphOne timerange={initialTimeRange} />}
                </div>
                <div className="d-flex p-4">
                        <DeviceManager />
                </div>
            </div>
        )
    }
}

class DeviceManager extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            devIdTmp: "",
            phoneTmp: "",
            registrations: []  // e.g. [{"devId": str, "phone": str}, {"devId": str, "phone": str}, ...]
        }
        this.handleRegistrationSubmit = this.handleRegistrationSubmit.bind(this);
        this.handleFormChange = this.handleFormChange.bind(this);
    }

    componentDidMount() {
        // TODO: initialize our registrations with a call to the backend.  Placeholder for now...
        fetch(process.env.REACT_APP_BACKEND_ADDR)
            .then(res => res.json())
            .then(
                (result) => {this.setState({registrations: result})},
                (error) => {console.log("problem fetching user's devices:", error)}
            )
    }

    handleFormChange = (e) => this.setState({[e.target.name]: e.target.value})

    handleRegistrationSubmit(e) {
        // TODO: other verification steps for registering a new device.  Also, POST a new registration to backend.
        e.preventDefault();
        this.setState({
            registrations: this.state.registrations.concat({"devId": this.state.devIdTmp, "phone": this.state.phoneTmp})
        })
        this.setState({devIdTmp:"", phoneTmp:""})
        document.getElementById("device-registration-form").reset()
    }

    render() {
        return (
            <div>
                <Jumbotron>
                    <h5>Your Alerts:</h5>
                    <ul>
                        {this.state.registrations.map((itm) => <li><b>device:</b> {itm.devId} ... <i>notify:</i> {itm.phone}</li>)}
                    </ul>
                </Jumbotron>
                <Form id="device-registration-form" onSubmit={this.handleRegistrationSubmit}>
                    <Form.Group>
                        <Form.Label>Device Identifier</Form.Label>
                        <Form.Control 
                            placeholder="Enter device ID" 
                            onChange={this.handleFormChange}
                            name="devIdTmp"
                            value={this.state.devIdTmp} 
                        />
                        <Form.Text className="text-muted">Enter your device's unique identifier. e.g. lsn50-1234</Form.Text>
                    </Form.Group>
                    <Form.Group>
                        <Form.Label>Notification Phone Number</Form.Label>
                        <Form.Control 
                            placeholder="Enter phone #" 
                            onChange={this.handleFormChange}
                            name="phoneTmp"
                            value={this.state.phoneTmp} 
                        />
                        <Form.Text className="text-muted">Enter a phone number to receive notifications</Form.Text>
                    </Form.Group>
                    <Button variant="primary" type="submit">Register</Button>
                </Form>
            </div>
        )
    }
}

export { Dashboard }