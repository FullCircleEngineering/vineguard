import pendulum
import json
import math

DEVICE_ID = "mock-dev-id"
OUTPUT_FILENAME = "data.json"
INTERVAL_SEC = 15


def gen_data():
    now = pendulum.now(tz="utc")
    rg = pendulum.period(now.subtract(days=7), now)
    output_lst = []
    for i, dt in enumerate(rg.range("seconds", INTERVAL_SEC)):
        # this will yield approx 40K records.
        # let's establish a sin curve with period approx. 24hrs ~= 5760 records
        msg = dict()
        angle = (i * 2 * math.pi) / 5760
        magnitude = 10
        mean = 15
        temp_c = mean + magnitude * math.sin(angle)
        msg.update(
            {
                "BatV": 3.11,
                "TempC": round(temp_c, 1),
                "device_id": DEVICE_ID,
                "time": dt.to_rfc3339_string(),
            }
        )
        output_lst.append(msg)
    with open(OUTPUT_FILENAME, "w") as fp:
        json.dump(output_lst, fp, allow_nan=False, indent=4)


if __name__ == "__main__":
    gen_data()
