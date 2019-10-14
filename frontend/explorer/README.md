# Explorer

A block explorer for EuroChain

## Run it yourself

### Prerequisites
* Caddyserver
* EuroChain daemon (`eurochaind`)


Make sure you have `eurochaind` (the EuroChain daemon) running with the explorer module enabled:
`eurochaind -M cgte`

Now start caddy from the `caddy` folder of this repository:
`caddy -conf Caddyfile.local`
and browse to http://localhost:2015
