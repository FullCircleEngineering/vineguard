import React from 'react';
import {
    Charts,
    ChartContainer,
    ChartRow,
    YAxis,
    LineChart,
    Resizable
} from "react-timeseries-charts";
import { dataTs } from "../_data/data";
import { Container } from 'react-bootstrap';

class Graph extends React.Component {

    constructor(props) {
        super(props);
        this.state = {timerange: this.props.timerange};
        this.handleTimeRangeChange = this.handleTimeRangeChange.bind(this);
    }

    handleTimeRangeChange = tmrg => {this.setState({timerange: tmrg})};

    componentDidMount() {
        if (!this.state.timerange) {
            this.setState({timerange: dataTs.timerange()})
        }
    }

    render () {
        return (
            <Container>
                <Resizable>
                    <ChartContainer
                        timeRange={this.state.timerange}
                        onTimeRangeChanged={this.handleTimeRangeChange}
                        format="%m/%d %H:%M"
                        enablePanZoom={true}
                        showGrid={true}
                    >
                        <ChartRow height="200">
                            <YAxis id="axis1" label="value" min={18} max={100} width="60" type="linear"/>
                            <Charts>
                                <LineChart axis="axis1" series={dataTs} column={["value"]}/>
                            </Charts>
                        </ChartRow>
                    </ChartContainer>
                </Resizable>
            </Container>
        );
    }
}

export { Graph};
