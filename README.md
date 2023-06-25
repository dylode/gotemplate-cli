# Go Template CLI (gotcli)

Go Template CLI (gotcli) is a compact tool designed to create files utilizing [Go templates](https://pkg.go.dev/text/template), a powerful feature provided by the Go programming language. By providing a template file and JSON or YAML data, gotcli generates customized output files. Go Template CLI uses [Sprig](https://masterminds.github.io/sprig/) to provide common template functions. 

An example:

`frr.conf.gotemplate`:

```
frr version 8.5.2_git
frr defaults traditional
hostname dev-frr
no ipv6 forwarding
service integrated-vtysh-config
!
router bgp 65100
{{- range (index . 0).Containers }}
    {{- if (.Name | contains "worker") }}
    neighbor {{ .IPv4Address | splitList "/" | first }} remote-as 65200
    {{- end }}
{{- end }}
exit
!
```

`docker network inspect my-network`:

```json
[
    {
        "Name": "my-network",
        "Id": "e7493fab5f6391a569aff3431c60eafd56701ae822641ec7acffdb8f4b0137a9",
        "Created": "2023-06-24T14:27:33.335557509Z",
        "Scope": "local",
        "Driver": "bridge",
        "EnableIPv6": true,
        "IPAM": {
            "Driver": "default",
            "Options": {},
            "Config": [
                {
                    "Subnet": "172.18.0.0/16",
                    "Gateway": "172.18.0.1"
                },
                {
                    "Subnet": "fc00:f853:ccd:e793::/64",
                    "Gateway": "fc00:f853:ccd:e793::1"
                }
            ]
        },
        "Internal": false,
        "Attachable": false,
        "Ingress": false,
        "ConfigFrom": {
            "Network": ""
        },
        "ConfigOnly": false,
        "Containers": {
            "2e11d5011c619fa9d384d4a62bf54ec283c447f00a29560f7c52a8db408925d2": {
                "Name": "kubernetes-base-cluster-worker2",
                "EndpointID": "b664e82888629f00068dc490bc398bcc929ac29e469a8643f2bea9073930730b",
                "MacAddress": "02:42:ac:12:00:04",
                "IPv4Address": "172.18.0.4/16",
                "IPv6Address": "fc00:f853:ccd:e793::4/64"
            },
            "562d262a9ec7a9c939322cd5a7e9bd761886885e3d7ffbfcae3478d22cf51425": {
                "Name": "kubernetes-base-cluster-worker",
                "EndpointID": "2e0dd6a0efcf6b2fb3b94663481999af30863d96ec7335a68ce769434b49eb5f",
                "MacAddress": "02:42:ac:12:00:05",
                "IPv4Address": "172.18.0.5/16",
                "IPv6Address": "fc00:f853:ccd:e793::5/64"
            },
            "8da6d2e3c18574dae60cd44de5fc094f9d9f91170ee67b467811f0047bac98ac": {
                "Name": "kubernetes-base-cluster-worker3",
                "EndpointID": "a01eafe92a7c5f764d5e89bd3d73135d5c8fe21b0312d1abaf1dc89d7a564de0",
                "MacAddress": "02:42:ac:12:00:03",
                "IPv4Address": "172.18.0.3/16",
                "IPv6Address": "fc00:f853:ccd:e793::3/64"
            },
            "f367240956e3dd1c5ec7caa968609f0760b2dd0c4394427eb2c71ab8ea60b9b7": {
                "Name": "kubernetes-base-cluster-control-plane",
                "EndpointID": "8781faae64095bda992a3a9b8b3d725da33c5b27ae6e7aae6247a02eb267925a",
                "MacAddress": "02:42:ac:12:00:02",
                "IPv4Address": "172.18.0.2/16",
                "IPv6Address": "fc00:f853:ccd:e793::2/64"
            }
        },
        "Options": {
            "com.docker.network.bridge.enable_ip_masquerade": "true",
            "com.docker.network.driver.mtu": "1500"
        },
        "Labels": {}
    }
]
```

`gotcli render -j "$(docker network inspect my-network)" "$(cat frr.conf.gotemplate)"`:

```
frr version 8.5.2_git
frr defaults traditional
hostname dev-frr
no ipv6 forwarding
service integrated-vtysh-config
!
router bgp 65100
    neighbor 172.18.0.4 remote-as 65200
    neighbor 172.18.0.5 remote-as 65200
    neighbor 172.18.0.3 remote-as 65200
exit
!
```

# How to install

Either build the binary yourself by cloning the Git repo and run `go build` or use the prebuilt binary:

```bash
wget https://github.com/dylode/gotemplate-cli/releases/download/v.0.1.0/gotcli
sudo chmod +x gotcli
sudo mv gotcli /usr/local/bin/
```