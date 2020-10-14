### ping-identity-exporter

Implementation of https://docs.pingidentity.com/bundle/pingaccess-53/page/nwx1564006726494.html


### Usage

```
./ping-identity-exporter --config config.yaml
```

By default, ping-identity-exporter will search for config.yaml on the same directory as the binary

```yml
listenPort: 9999
insecure: true
timeout: 30
pingAccessHeartbeatEndpoint: https://%s:3000/pa/heartbeat.ping
```


### Current issues


### Contributing

We welcome any contributions. Please fork the project on GitHub and open Pull Requests for any proposed changes.
