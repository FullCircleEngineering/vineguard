import { TimeSeries } from "pondjs";

const data = {
    name: "traffic",
    columns: ["time", "value"],
    points: [
        [1570654139000, 52],
        [1570654239000, 18],
        [1570654339000, 26],
        [1570654439000, 93],
    ]
};

const dataTs = new TimeSeries(data);
export {dataTs};