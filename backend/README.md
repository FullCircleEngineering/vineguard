# backend

The backend server loads the following environment variables:

| Name    | Description                    |
| ----    | -----------                    |
| TTN_KEY | API Key for The Things Network |

When running locally, you must source these environment variables.

Create a file `.secret`, including:
```bash
export TTN_KEY="ttn-account-v2.your-special-values"
```

And run `source ./.secret` before running the server.