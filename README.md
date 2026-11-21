# iot-contract

Shared Go contract (envelope, vocabulary, inventory, intentions, results, configuration) between iot-gateway-server and iot-backend.

## Messages

| Kind | Direction | Topic | Body |
|---|---|---|---|
| `inventory` | up | `…/gateways/{gw}/inventory` | `Inventory` |
| `config` | down | `…/gateways/{gw}/config` | `ConfigDocument` (one kind, always whole) |
| `config_report` | up | `…/gateways/{gw}/config_report` | `ConfigReport` (every kind the gateway supports) |

The envelope is used in both directions; downward `boot_id` and `seq` are empty.
Config items are versioned by the backend; the gateway stores and echoes versions and never computes them.
Display colours are names (`red`, `green`, `yellow`, `purple`, `blue`, `light_yellow`, `white`); firmware codes exist only in the gateway.
