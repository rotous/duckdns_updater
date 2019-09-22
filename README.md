# DuckDNS Update Daemon

This is a small daemon that can be used to update the IP address of a [DuckDNS](https://www.duckdns.org/) domain.

The following ENV vars can be set to control the daemon:

- DUCKDNS_UPDATE_URL - The URL that will be used to update the IP address. Defaults to `https://www.duckdns.org/update?domains=%s&token=%s&verbose=true`
- DUCKDNS_DEFAULT_INTERVAL - The interval in seconds between the update requests. Defaults to 300 seconds.
- DUCKDNS_DOMAINS - The domains registered at DuckDNS (comma separated list) for which the IP address will be updated. Required.
- DUCKDNS_KEY - The key that was given by DuckDNS to update the domains. Required

