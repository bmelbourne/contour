# Contour Configuration Reference

- [Serve Flags](#serve-flags)
- [Configuration File](#configuration-file)
- [Environment Variables](#environment-variables)
- [Bootstrap Config File](#bootstrap-config-file)

## Overview

There are various ways to configure Contour, flags, the configuration file, as well as environment variables.
Contour has a precedence of configuration for contour serve, meaning anything configured in the config file is overridden by environment vars which are overridden by cli flags.

## Serve Flags

The `contour serve` command is the main command which is used to watch for Kubernetes resource and process them into Envoy configuration which is then streamed to any Envoy via its xDS gRPC connection.
There are a number of flags that can be passed to this command which further configures how Contour operates.
Many of these flags are mirrored in the [Contour Configuration File](#configuration-file).

| Flag Name                                                       | Description                                                                             |
| --------------------------------------------------------------- | --------------------------------------------------------------------------------------- |
| `--config-path`                                                 | Path to base configuration                                                              |
| `--contour-config-name`                                         | Name of the ContourConfiguration resource to use                                        |
| `--incluster`                                                   | Use in cluster configuration                                                            |
| `--kubeconfig=</path/to/file>`                                  | Path to kubeconfig (if not in running inside a cluster)                                 |
| `--xds-address=<ipaddr>`                                        | xDS gRPC API address                                                                    |
| `--xds-port=<port>`                                             | xDS gRPC API port                                                                       |
| `--stats-address=<ipaddr>`                                      | Envoy /stats interface address                                                          |
| `--stats-port=<port>`                                           | Envoy /stats interface port                                                             |
| `--debug-http-address=<address>`                                | Address the debug http endpoint will bind to.                                           |
| `--debug-http-port=<port>`                                      | Port the debug http endpoint will bind to                                               |
| `--http-address=<ipaddr>`                                       | Address the metrics HTTP endpoint will bind to                                          |
| `--http-port=<port>`                                            | Port the metrics HTTP endpoint will bind to.                                            |
| `--health-address=<ipaddr>`                                     | Address the health HTTP endpoint will bind to                                           |
| `--health-port=<port>`                                          | Port the health HTTP endpoint will bind to                                              |
| `--contour-cafile=</path/to/file\|CONTOUR_CERT_FILE>`           | CA bundle file name for serving gRPC with TLS                                           |
| `--contour-cert-file=</path/to/file\|CONTOUR_CERT_FILE>`        | Contour certificate file name for serving gRPC over TLS                                 |
| `--contour-key-file=</path/to/file\|CONTOUR_KEY_FILE>`          | Contour key file name for serving gRPC over TLS                                         |
| `--insecure`                                                    | Allow serving without TLS secured gRPC                                                  |
| `--root-namespaces=<ns,ns>`                                     | Restrict contour to searching these namespaces for root ingress routes                  |
| `--watch-namespaces=<ns,ns>`                                    | Restrict contour to searching these namespaces for all resources                        |
| `--ingress-class-name=<name>`                                   | Contour IngressClass name (comma-separated list allowed)                                |
| `--ingress-status-address=<address>`                            | Address to set in Ingress object status                                                 |
| `--envoy-http-access-log=</path/to/file>`                       | Envoy HTTP access log                                                                   |
| `--envoy-https-access-log=</path/to/file>`                      | Envoy HTTPS access log                                                                  |
| `--envoy-service-http-address=<ipaddr>`                         | Kubernetes Service address for HTTP requests                                            |
| `--envoy-service-https-address=<ipaddr>`                        | Kubernetes Service address for HTTPS requests                                           |
| `--envoy-service-http-port=<port>`                              | Kubernetes Service port for HTTP requests                                               |
| `--envoy-service-https-port=<port>`                             | Kubernetes Service port for HTTPS requests                                              |
| `--envoy-service-name=<name>`                                   | Name of the Envoy service to inspect for Ingress status details.                        |
| `--envoy-service-namespace=<namespace>`                         | Envoy Service Namespace                                                                 |
| `--use-proxy-protocol`                                          | Use PROXY protocol for all listeners                                                    |
| `--accesslog-format=<envoy\|json>`                              | Format for Envoy access logs                                                            |
| `--disable-leader-election`                                     | Disable leader election mechanism                                                       |
| `--disable-feature=<extensionservices\|tlsroutes\|grpcroutes>`  | Do not start an informer for the specified resources. Flag can be given multiple times. |
| `--leader-election-lease-duration`                              | The duration of the leadership lease.                                                   |
| `--leader-election-renew-deadline`                              | The duration leader will retry refreshing leadership before giving up.                  |
| `--leader-election-retry-period`                                | The interval which Contour will attempt to acquire leadership lease.                    |
| `--leader-election-resource-name`                               | The name of the resource (Lease) leader election will lease.                            |
| `--leader-election-resource-namespace`                          | The namespace of the resource (Lease) leader election will lease.                       |
| `-d, --debug`                                                   | Enable debug logging                                                                    |
| `--kubernetes-debug=<log level>`                                | Enable Kubernetes client debug logging                                                  |
| `--log-format=<text\|json>`                                     | Log output format for Contour. Either text (default) or json.                           |
| `--kubernetes-client-qps=<qps>`                                 | QPS allowed for the Kubernetes client.                                                  |
| `--kubernetes-client-burst=<burst>`                             | Burst allowed for the Kubernetes client.                                                |

## Configuration File

A configuration file can be passed to the `--config-path` argument of the `contour serve` command to specify additional configuration to Contour.
In most deployments, this file is passed to Contour via a ConfigMap which is mounted as a volume to the Contour pod.

The Contour configuration file is optional.
In its absence, Contour will operate with reasonable defaults.
Where Contour settings can also be specified with command-line flags, the command-line value takes precedence over the configuration file.

| Field Name                | Type                   | Default                                                                                              | Description                                                                                                                                                                                                                                                                           |
|---------------------------| ---------------------- |------------------------------------------------------------------------------------------------------| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| accesslog-format          | string                 | `envoy`                                                                                              | This key sets the global [access log format][2] for Envoy. Valid options are `envoy` or `json`.                                                                                                                                                                                       |
| accesslog-format-string   | string                 | None                                                                                                 | If present, this specifies custom access log format for Envoy. See [Envoy documentation](https://www.envoyproxy.io/docs/envoy/latest/configuration/observability/access_log/usage) for more information about the syntax. This field only has effect if `accesslog-format` is `envoy` |
| accesslog-level           | string                 | `info`                                                                                               | This field specifies the verbosity level of the access log. Valid options are `info` (default, all requests are logged), `error` (all non-success, i.e. 300+ response code, requests are logged), `critical` (all server error, i.e. 500+ response code, requests are logged) and `disabled`. |
| debug                     | boolean                | `false`                                                                                              | Enables debug logging.                                                                                                                                                                                                                                                                |
| default-http-versions     | string array           | <code style="white-space:nowrap">HTTP/1.1</code> <br> <code style="white-space:nowrap">HTTP/2</code> | This array specifies the HTTP versions that Contour should program Envoy to serve. HTTP versions are specified as strings of the form "HTTP/x", where "x" represents the version number.                                                                                              |
| disableAllowChunkedLength | boolean                | `false`                                                                                              | If this field is true, Contour will disable the RFC-compliant Envoy behavior to strip the `Content-Length` header if `Transfer-Encoding: chunked` is also set. This is an emergency off-switch to revert back to Envoy's default behavior in case of failures.
| compression               | CompressionParameters  |                                                                                                      | The [compression configuration](#compression-parameters).                                                                                                                                                                                                                             |
| disableMergeSlashes       | boolean                | `false`                                                                                              | This field disables Envoy's non-standard merge_slashes path transformation behavior that strips duplicate slashes from request URL paths.
| serverHeaderTransformation       | string                | `overwrite`                                                                                              | This field defines the action to be applied to the Server header on the response path. Values: `overwrite` (default), `append_if_absent`, `pass_through`
| disablePermitInsecure     | boolean                | `false`                                                                                              | If this field is true, Contour will ignore `PermitInsecure` field in HTTPProxy documents.                                                                                                                                                                                             |
| envoy-service-name        | string                 | `envoy`                                                                                              | This sets the service name that will be inspected for address details to be applied to Ingress objects.                                                                                                                                                                               |
| envoy-service-namespace   | string                 | `projectcontour`                                                                                     | This sets the namespace of the service that will be inspected for address details to be applied to Ingress objects. If the `CONTOUR_NAMESPACE` environment variable is present, Contour will populate this field with its value.                                                      |
| ingress-status-address    | string                 | None                                                                                                 | If present, this specifies the address that will be copied into the Ingress status for each Ingress that Contour manages. It is exclusive with `envoy-service-name` and `envoy-service-namespace`.                                                                                    |
| incluster                 | boolean                | `false`                                                                                              | This field specifies that Contour is running in a Kubernetes cluster and should use the in-cluster client access configuration.                                                                                                                                                       |
| json-fields               | string array           | [fields][5]                                                                                          | This is the list the field names to include in the JSON [access log format][2]. This field only has effect if `accesslog-format` is `json`.                                                                                                                                           |
| kubeconfig                | string                 | `$HOME/.kube/config`                                                                                 | Path to a Kubernetes [kubeconfig file][3] for when Contour is executed outside a cluster.                                                                                                                                                                                             |
| kubernetesClientQPS          | float32             |                                                                                                      | QPS allowed for the Kubernetes client.                                                                                                                                                                    |
| kubernetesClientBurst        | int                    |                                                                                                      | Burst allowed for the Kubernetes client.                                                                                                                                                                    |
| policy                    | PolicyConfig           |                                                                                                      | The default [policy configuration](#policy-configuration).                                                                                                                                                                                                                            |
| tls                       | TLS                    |                                                                                                      | The default [TLS configuration](#tls-configuration).                                                                                                                                                                                                                                  |
| timeouts                  | TimeoutConfig          |                                                                                                      | The [timeout configuration](#timeout-configuration).                                                                                                                                                                                                                                  |
| cluster                   | ClusterConfig          |                                                                                                      | The [cluster configuration](#cluster-configuration).                                                                                                                                                                                                                                  |
| network                   | NetworkConfig          |                                                                                                      | The [network configuration](#network-configuration).                                                                                                                                                                                                                                  |
| listener                  | ListenerConfig         |                                                                                                      | The [listener configuration](#listener-configuration).                                                                                                                                                                                                                                |
| gateway                   | GatewayConfig          |                                                                                                      | The [gateway-api Gateway configuration](#gateway-configuration).                                                                                                                                                                                                                      |
| rateLimitService          | RateLimitServiceConfig |                                                                                                      | The [rate limit service configuration](#rate-limit-service-configuration).                                                                                                                                                                                                            |
| enableExternalNameService | boolean                | `false`                                                                                              | Enable ExternalName Service processing. Enabling this has security implications. Please see the [advisory](https://github.com/projectcontour/contour/security/advisories/GHSA-5ph6-qq5x-7jwc) for more details.                                                                       |
| metrics                   | MetricsParameters     |                                                                                                       | The [metrics configuration](#metrics-configuration) |
| featureFlags              | string array           | `[]`                                                                                                 | Defines the toggle to enable new contour features. Available toggles are:  <br/> 1. `useEndpointSlices` - configures contour to fetch endpoint data from k8s endpoint slices.                                                                                                         |
| omEnforcedHealthListener | OMEnforcedHealthListenerConfig | | The [overload manager enforced health listener config](#overload-manager-enforced-health-listener-config) |

### TLS Configuration

The TLS configuration block can be used to configure default values for how
Contour should provision TLS hosts.

| Field Name               | Type     | Default                                                                                                           | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      |
| ------------------------ | -------- | ----------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| minimum-protocol-version | string   | `1.2`                                                                                                             | This field specifies the minimum TLS protocol version that is allowed. Valid options are `1.2` (default) and `1.3`. Any other value defaults to TLS 1.2.
| maximum-protocol-version | string   | `1.3`                                                                                                              | This field specifies the maximum TLS protocol version that is allowed. Valid options are `1.2` and `1.3`. Any other value defaults to TLS 1.3.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| fallback-certificate     |          |                                                                                                                   | [Fallback certificate configuration](#fallback-certificate).                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| envoy-client-certificate |          |                                                                                                                   | [Client certificate configuration for Envoy](#envoy-client-certificate).                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         |
| cipher-suites            | []string | See [config package documentation](https://pkg.go.dev/github.com/projectcontour/contour/pkg/config#pkg-variables) | This field specifies the TLS ciphers to be supported by TLS listeners when negotiating TLS 1.2. This parameter should only be used by advanced users. Note that this is ignored when TLS 1.3 is in use. The set of ciphers that are allowed is a superset of those supported by default in stock, non-FIPS Envoy builds and FIPS builds as specified [here](https://www.envoyproxy.io/docs/envoy/latest/api-v3/extensions/transport_sockets/tls/v3/common.proto#envoy-v3-api-field-extensions-transport-sockets-tls-v3-tlsparameters-cipher-suites). Custom ciphers not accepted by Envoy in a standard build are not supported. |

### Upstream TLS Configuration

The Upstream TLS configuration block can be used to configure default values for how Contour establishes TLS for upstream connections.

| Field Name               | Type     | Default                                                                                                           | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      |
| ------------------------ | -------- | ----------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| minimum-protocol-version | string   | `1.2` | This field specifies the minimum TLS protocol version that is allowed. Valid options are `1.2` (default) and `1.3`. Any other value defaults to TLS 1.2. |
| maximum-protocol-version | string   | `1.3` | This field specifies the maximum TLS protocol version that is allowed. Valid options are `1.2` and `1.3`. Any other value defaults to TLS 1.3. |
| cipher-suites | []string | See [config package documentation](https://pkg.go.dev/github.com/projectcontour/contour/pkg/config#pkg-variables) | This field specifies the TLS ciphers to be supported by TLS listeners when negotiating TLS 1.2. This parameter should only be used by advanced users. Note that this is ignored when TLS 1.3 is in use. The set of ciphers that are allowed is a superset of those supported by default in stock, non-FIPS Envoy builds and FIPS builds as specified [here](https://www.envoyproxy.io/docs/envoy/latest/api-v3/extensions/transport_sockets/tls/v3/common.proto#envoy-v3-api-field-extensions-transport-sockets-tls-v3-tlsparameters-cipher-suites). Custom ciphers not accepted by Envoy in a standard build are not supported. |

### Fallback Certificate

| Field Name | Type   | Default | Description                                                                                     |
| ---------- | ------ | ------- | ----------------------------------------------------------------------------------------------- |
| name       | string | `""`    | This field specifies the name of the Kubernetes secret to use as the fallback certificate.      |
| namespace  | string | `""`    | This field specifies the namespace of the Kubernetes secret to use as the fallback certificate. |


### Envoy Client Certificate

| Field Name | Type   | Default | Description                                                                                                                                                            |
| ---------- | ------ | ------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| name       | string | `""`    | This field specifies the name of the Kubernetes secret to use as the client certificate and private key when establishing TLS connections to the backend service.      |
| namespace  | string | `""`    | This field specifies the namespace of the Kubernetes secret to use as the client certificate and private key when establishing TLS connections to the backend service. |


### Timeout Configuration

The timeout configuration block can be used to configure various timeouts for the proxies. All fields are optional; Contour/Envoy defaults apply if a field is not specified.

| Field Name                       | Type   | Default | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                      |
| -------------------------------- | ------ | ------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| request-timeout                  | string | none*   | This field specifies the default request timeout. Note that this is a timeout for the entire request, not an idle timeout. Must be a [valid Go duration string][4], or omitted or set to `infinity` to disable the timeout entirely. See [the Envoy documentation][12] for more information.<br /><br />_Note: A value of `0s` previously disabled this timeout entirely. This is no longer the case. Use `infinity` or omit this field to disable the timeout._ |
| connection-idle-timeout          | string | `60s`   | This field defines how long the proxy should wait while there are no active requests (for HTTP/1.1) or streams (for HTTP/2) before terminating an HTTP connection. The timeout applies to downstream connections only. Must be a [valid Go duration string][4], or `infinity` to disable the timeout entirely. See [the Envoy documentation][8] for more information.                                                                                            |
| stream-idle-timeout              | string | `5m`*   | This field defines how long the proxy should wait while there is no activity during single request/response (for HTTP/1.1) or stream (for HTTP/2). Timeout will not trigger while HTTP/1.1 connection is idle between two consecutive requests. Must be a [valid Go duration string][4], or `infinity` to disable the timeout entirely. See [the Envoy documentation][9] for more information.                                                                   |
| max-connection-duration          | string | none*   | This field defines the maximum period of time after an HTTP connection has been established from the client to the proxy before it is closed by the proxy, regardless of whether there has been activity or not. Must be a [valid Go duration string][4], or omitted or set to `infinity` for no max duration. See [the Envoy documentation][10] for more information.                                                                                           |
| delayed-close-timeout            | string | `1s`*   | *Note: this is an advanced setting that should not normally need to be tuned.* <br /><br /> This field defines how long envoy will wait, once connection close processing has been initiated, for the downstream peer to close the connection before Envoy closes the socket associated with the connection. Setting this timeout to 'infinity' will disable it.  See [the Envoy documentation][13] for more information.                                        |
| connection-shutdown-grace-period | string | `5s`*   | This field defines how long the proxy will wait between sending an initial GOAWAY frame and a second, final GOAWAY frame when terminating an HTTP/2 connection. During this grace period, the proxy will continue to respond to new streams. After the final GOAWAY frame has been sent, the proxy will refuse new streams. Must be a [valid Go duration string][4]. See [the Envoy documentation][11] for more information.                                     |
| connect-timeout                  | string | `2s`    | This field defines how long the proxy will wait for the upstream connection to be established.

_This is Envoy's default setting value and is not explicitly configured by Contour._

### Cluster Configuration

The cluster configuration block can be used to configure various parameters for Envoy clusters.

| Field Name                        | Type   | Default | Description                                                                                                                                                                     |
|-----------------------------------|--------|---------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| dns-lookup-family                 | string | auto    | This field specifies the dns-lookup-family to use for upstream requests to externalName type Kubernetes services from an HTTPProxy route. Values are: `auto`, `v4`, `v6`, `all` |
| max-requests-per-connection       | int    | none    | This field specifies the maximum requests for upstream connections. If not specified, there is no limit                                                                         |
| circuit-breakers       | [CircuitBreakers](#circuit-breakers)    | none    | This field specifies the default value for [circuit-breaker-annotations](https://projectcontour.io/docs/main/config/annotations/) for services that don't specify them.                                                                    |
| per-connection-buffer-limit-bytes | int    | 1MiB*   | This field specifies the soft limit on size of the cluster’s new connection read and write buffer. If not specified, Envoy defaults of 1MiB apply                               |
| upstream-tls |  UpstreamTLS   |    | [Upstream TLS configuration](#upstream-tls)                            |

_This is Envoy's default setting value and is not explicitly configured by Contour._




### Network Configuration

The network configuration block can be used to configure various parameters network connections.

| Field Name              | Type | Default | Description                                                                                                                                                                                                                                                                                                                                                                                    |
|-------------------------|------|---------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| num-trusted-hops        | int  | 0       | Configures the number of additional ingress proxy hops from the right side of the x-forwarded-for HTTP header to trust.                                                                                                                                                                                                                                                                        |
| admin-port              | int  | 9001    | Configures the Envoy Admin read-only listener on Envoy. Set to `0` to disable.                                                                                                                                                                                                                                                                                                                 |
| strip-trailing-host-dot | bool | false   | Defines if trailing dot of the host should be removed from host/authority header before any processing of request by HTTP filters or routing. This affects the upstream host header. Without setting this option to true, incoming requests with host example.com. will not match against route with domains match set to example.com. See [the Envoy documentation][15] for more information. |

### Listener Configuration

The listener configuration block can be used to configure various parameters for Envoy listener.

| Field Name                        | Type   | Default | Description                                                                                                                                                                                                                                                   |
|-----------------------------------|--------|---------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| connection-balancer               | string | `""`    | This field specifies the listener connection balancer. If the value is `exact`, the listener will use the exact connection balancer to balance connections between threads in a single Envoy process. See [the Envoy documentation][14] for more information. |
| max-requests-per-connection       | int    | none    | This field specifies the maximum requests for downstream connections. If not specified, there is no limit                                                                                                                                                     |
| per-connection-buffer-limit-bytes | int    | 1MiB*   | This field specifies the soft limit on size of the listener’s new connection read and write buffer. If not specified, Envoy defaults of 1MiB apply                                                                                                            |
| socket-options                    | SocketOptions |  | The [Socket Options](#socket-options) for Envoy listeners.                                                                                                                                                                                                    |
| max-requests-per-io-cycle         | int    | none    | Defines the limit on number of HTTP requests that Envoy will process from a single connection in a single I/O cycle. Requests over this limit are processed in subsequent I/O cycles. Can be used as a mitigation for CVE-2023-44487 when abusive traffic is detected. Configures the `http.max_requests_per_io_cycle` Envoy runtime setting. The default value when this is not set is no limit. |
| http2-max-concurrent-streams      | int    | none    | Defines the value for SETTINGS_MAX_CONCURRENT_STREAMS Envoy will advertise in the SETTINGS frame in HTTP/2 connections and the limit for concurrent streams allowed for a peer on a single HTTP/2 connection. It is recommended to not set this lower than 100 but this field can be used to bound resource usage by HTTP/2 connections and mitigate attacks like CVE-2023-44487. The default value when this is not set is unlimited. |

_This is Envoy's default setting value and is not explicitly configured by Contour._

### Gateway Configuration

The gateway configuration block is used to configure which gateway-api Gateway Contour should configure:

| Field Name     | Type           | Default | Description                                                                    |
| -------------- | -------------- | ------- | ------------------------------------------------------------------------------ |
| gatewayRef     | NamespacedName |         | [Gateway namespace and name](#gateway-ref). |

### Gateway Ref

| Field Name | Type   | Default | Description                                                                                     |
| ---------- | ------ | ------- | ----------------------------------------------------------------------------------------------- |
| name       | string | `""`    | This field specifies the name of the specific Gateway to reconcile.                             |
| namespace  | string | `""`    | This field specifies the namespace of the specific Gateway to reconcile.                        |

### Policy Configuration

The Policy configuration block can be used to configure default policy values
that are set if not overridden by the user.

The `request-headers` field is used to rewrite headers on a HTTP request, and
the `response-headers` field is used to rewrite headers on a HTTP response.

| Field Name       | Type         | Default | Description                                                                                       |
| ---------------- | ------------ | ------- | ------------------------------------------------------------------------------------------------- |
| request-headers  | HeaderPolicy | none    | The default request headers set or removed on all service routes if not overridden in the object  |
| response-headers | HeaderPolicy | none    | The default response headers set or removed on all service routes if not overridden in the object |
| applyToIngress   | Boolean      | false   | Whether the global policy should apply to Ingress objects                                         |

#### HeaderPolicy

The `set` field sets an HTTP header value, creating it if it doesn't already exist but not overwriting it if it does.
The `remove` field removes an HTTP header.

| Field Name | Type              | Default | Description                                                                     |
| ---------- | ----------------- | ------- | ------------------------------------------------------------------------------- |
| set        | map[string]string | none    | Map of headers to set on all service routes if not overridden in the object     |
| remove     | []string          | none    | List of headers to remove on all service routes if not overridden in the object |

Note: the values of entries in the `set` and `remove` fields can be overridden in HTTPProxy objects but it is not possible to remove these entries.

### Rate Limit Service Configuration

The rate limit service configuration block is used to configure an optional global rate limit service:

| Field Name                  | Type   | Default | Description                                                                                                                                                                                                                                                                                                            |
|-----------------------------| ------ | ------- |------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| extensionService            | string | <none>  | This field identifies the extension service defining the rate limit service, formatted as <namespace>/<name>.                                                                                                                                                                                                          |
| domain                      | string | contour | This field defines the rate limit domain value to pass to the rate limit service. Acts as a container for a set of rate limit definitions within the RLS.                                                                                                                                                              |
| failOpen                    | bool   | false   | This field defines whether to allow requests to proceed when the rate limit service fails to respond with a valid rate limit decision within the timeout defined on the extension service.                                                                                                                             |
| enableXRateLimitHeaders     | bool   | false   | This field defines whether to include the X-RateLimit headers X-RateLimit-Limit, X-RateLimit-Remaining, and X-RateLimit-Reset (as defined by the IETF Internet-Draft https://tools.ietf.org/id/draft-polli-ratelimit-headers-03.html), on responses to clients when the Rate Limit Service is consulted for a request. |
| enableResourceExhaustedCode | bool   | false   | This field defines whether to translate status code 429 to gRPC RESOURCE_EXHAUSTED instead of UNAVAILABLE.                                                                                                                                                                                                             |

### Metrics Configuration

MetricsParameters holds configurable parameters for Contour and Envoy metrics.

| Field Name  | Type                    | Default | Description                                                          |
| ----------- | ----------------------- | ------- | -------------------------------------------------------------------- |
| contour     | MetricsServerParameters |         | [Metrics Server Parameters](#metrics-server-parameters) for Contour. |
| envoy       | MetricsServerParameters |         | [Metrics Server Parameters](#metrics-server-parameters) for Envoy.   |

### Metrics Server Parameters

MetricsServerParameters holds configurable parameters for Contour and Envoy metrics.
Metrics are served over HTTPS if `server-certificate-path` and `server-key-path` are set.
Metrics and health endpoints cannot have the same port number when metrics are served over HTTPS.

| Field Name              | Type   | Default                      | Description                                                                  |
| ----------------------- | ------ | ---------------------------- | -----------------------------------------------------------------------------|
| address                 | string | 0.0.0.0                      | Address that metrics server will bind to.                                    |
| port                    | int    | 8000 (Contour), 8002 (Envoy) | Port that metrics server will bind to.                                       |
| server-certificate-path | string | none                         | Optional path to the server certificate file.                                |
| server-key-path         | string | none                         | Optional path to the server private key file.                                |
| ca-certificate-path     | string | none                         | Optional path to the CA certificate file used to verify client certificates. |

### Socket Options

| Field Name      | Type   | Default | Description                                                                   |
| --------------- | ------ | ------- | ----------------------------------------------------------------------------- |
| tos             | int    | 0       | Defines the value for IPv4 TOS field (including 6 bit DSCP field) for IP packets originating from Envoy listeners. Single value is applied to all listeners. The value must be in the range 0-255, 0 means socket option is not set. If listeners are bound to IPv6-only addresses, setting this option will cause an error. |
| traffic-class   | int    | 0       | Defines the value for IPv6 Traffic Class field (including 6 bit DSCP field) for IP packets originating from the Envoy listeners. Single value is applied to all listeners. The value must be in the range 0-255, 0 means socket option is not set. If listeners are bound to IPv4-only addresses, setting this option will cause an error. |


### Circuit Breakers

| Field Name      | Type   | Default | Description                                                                   |
| --------------- | ------ | ------- | ----------------------------------------------------------------------------- |
| max-connections | int    | 0       | The maximum number of connections that a single Envoy instance allows to the Kubernetes Service; defaults to 1024. |
| max-pending-requests  | int    | 0       | The maximum number of pending requests that a single Envoy instance allows to the Kubernetes Service; defaults to 1024. |
| max-requests | int    | 0       | The maximum parallel requests a single Envoy instance allows to the Kubernetes Service; defaults to 1024 |
| max-retries  | int    | 0       | The maximum number of parallel retries a single Envoy instance allows to the Kubernetes Service; defaults to 3. This setting only makes sense if the cluster is configured to do retries.|

### Compression Parameters

Sets the compression configuration applied in the compression HTTP filter of the default Listener filters.

| Field Name | Type   | Default | Description             |
|------------|--------|--------|-------------------------|
| algorithm  | string | "gzip" | Compression algorithm. Setting this to `disabled` will make Envoy skip "Accept-Encoding: gzip,deflate" request header and always return uncompressed response. Values:`gzip` (default), `brotli`, `zstd`, `disabled`. |

### Overload Manager Enforced Health Listener Config

OMEnforcedHealthListenerConfig holds the configuration for a health stats listener that enforces limits imposed by the overload manager,
notably max global downstream connections. This listener can be used to fail readiness checks without failing liveness to
load shed connections under heavy load.

By default, this configuration is not enabled.

| Field Name              | Type   | Default | Description                                |
| ----------------------- | ------ | ------- | ------------------------------------------ |
| address                 | string | none    | Address that health listener will bind to. |
| port                    | int    | 0       | Port that health listener will bind to.    |

### Configuration Example

The following is an example ConfigMap with configuration file included:

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: contour
  namespace: projectcontour
data:
  contour.yaml: |
    # specify the gateway-api Gateway Contour should configure
    # gateway:
    #   namespace: projectcontour
    #   name: contour
    #
    # should contour expect to be running inside a k8s cluster
    # incluster: true
    #
    # path to kubeconfig (if not running inside a k8s cluster)
    # kubeconfig: /path/to/.kube/config
    #
    # Disable RFC-compliant behavior to strip "Content-Length" header if
    # "Tranfer-Encoding: chunked" is also set.
    # disableAllowChunkedLength: false
    # Disable HTTPProxy permitInsecure field
    disablePermitInsecure: false
    tls:
    # minimum TLS version that Contour will negotiate
    # minimum-protocol-version: "1.2"
    # TLS ciphers to be supported by Envoy TLS listeners when negotiating
    # TLS 1.2.
    # cipher-suites:
    # - '[ECDHE-ECDSA-AES128-GCM-SHA256|ECDHE-ECDSA-CHACHA20-POLY1305]'
    # - '[ECDHE-RSA-AES128-GCM-SHA256|ECDHE-RSA-CHACHA20-POLY1305]'
    # - 'ECDHE-ECDSA-AES256-GCM-SHA384'
    # - 'ECDHE-RSA-AES256-GCM-SHA384'
    # Defines the Kubernetes name/namespace matching a secret to use
    # as the fallback certificate when requests which don't match the
    # SNI defined for a vhost.
      fallback-certificate:
    #   name: fallback-secret-name
    #   namespace: projectcontour
      envoy-client-certificate:
    #   name: envoy-client-cert-secret-name
    #   namespace: projectcontour
    ### Logging options
    # Default setting
    accesslog-format: envoy
    # The default access log format is defined by Envoy but it can be customized by setting following variable.
    # accesslog-format-string: "...\n"
    # To enable JSON logging in Envoy
    # accesslog-format: json
    # accesslog-level: info
    # The default fields that will be logged are specified below.
    # To customise this list, just add or remove entries.
    # The canonical list is available at
    # https://godoc.org/github.com/projectcontour/contour/internal/envoy#JSONFields
    # json-fields:
    #   - "@timestamp"
    #   - "authority"
    #   - "bytes_received"
    #   - "bytes_sent"
    #   - "downstream_local_address"
    #   - "downstream_remote_address"
    #   - "duration"
    #   - "method"
    #   - "path"
    #   - "protocol"
    #   - "request_id"
    #   - "requested_server_name"
    #   - "response_code"
    #   - "response_flags"
    #   - "uber_trace_id"
    #   - "upstream_cluster"
    #   - "upstream_host"
    #   - "upstream_local_address"
    #   - "upstream_service_time"
    #   - "user_agent"
    #   - "x_forwarded_for"
    #
    # default-http-versions:
    # - "HTTP/2"
    # - "HTTP/1.1"
    #
    # The following shows the default proxy timeout settings.
    # timeouts:
    #   request-timeout: infinity
    #   connection-idle-timeout: 60s
    #   stream-idle-timeout: 5m
    #   max-connection-duration: infinity
    #   connection-shutdown-grace-period: 5s
    #
    # Envoy cluster settings.
    # cluster:
    #   configure the cluster dns lookup family
    #   valid options are: auto (default), v4, v6, all
    #   dns-lookup-family: auto
    #   the maximum requests for upstream connections.
    #   If not specified, there is no limit.
    #   Setting this parameter to 1 will effectively disable keep alive
    #   max-requests-per-connection: 0
    #   the soft limit on size of the cluster’s new connection read and write buffers
    #   per-connection-buffer-limit-bytes: 32768
    #
    # network:
    #   Configure the number of additional ingress proxy hops from the
    #   right side of the x-forwarded-for HTTP header to trust.
    #   num-trusted-hops: 0
    #   Configure the port used to access the Envoy Admin interface.
    #   admin-port: 9001
    #
    # Configure an optional global rate limit service.
    # rateLimitService:
    #   Identifies the extension service defining the rate limit service,
    #   formatted as <namespace>/<name>.
    #   extensionService: projectcontour/ratelimit
    #   Defines the rate limit domain to pass to the rate limit service.
    #   Acts as a container for a set of rate limit definitions within
    #   the RLS.
    #   domain: contour
    #   Defines whether to allow requests to proceed when the rate limit
    #   service fails to respond with a valid rate limit decision within
    #   the timeout defined on the extension service.
    #   failOpen: false
    #   Defines whether to include the X-RateLimit headers X-RateLimit-Limit,
    #   X-RateLimit-Remaining, and X-RateLimit-Reset (as defined by the IETF
    #   Internet-Draft linked below), on responses to clients when the Rate
    #   Limit Service is consulted for a request.
    #   ref. https://tools.ietf.org/id/draft-polli-ratelimit-headers-03.html
    #   enableXRateLimitHeaders: false
    #   Defines whether to translate status code 429 to grpc code RESOURCE_EXHAUSTED
    #   instead of the default UNAVAILABLE
    #   enableResourceExhaustedCode: false
    #
    # Global Policy settings.
    # policy:
    #   # Default headers to set on all requests (unless set/removed on the HTTPProxy object itself)
    #   request-headers:
    #     set:
    #       # example: the hostname of the Envoy instance that proxied the request
    #       X-Envoy-Hostname: %HOSTNAME%
    #       # example: add a l5d-dst-override header to instruct Linkerd what service the request is destined for
    #       l5d-dst-override: %CONTOUR_SERVICE_NAME%.%CONTOUR_NAMESPACE%.svc.cluster.local:%CONTOUR_SERVICE_PORT%
    #   # default headers to set on all responses (unless set/removed on the HTTPProxy object itself)
    #   response-headers:
    #     set:
    #       # example: Envoy flags that provide additional details about the response or connection
    #       X-Envoy-Response-Flags: %RESPONSE_FLAGS%
    #   Whether or not the policy settings should apply to ingress objects
    #   applyToIngress: true
    #
    # metrics:
    #  contour:
    #    address: 0.0.0.0
    #    port: 8000
    #    server-certificate-path: /path/to/server-cert.pem
    #    server-key-path: /path/to/server-private-key.pem
    #    ca-certificate-path: /path/to/root-ca-for-client-validation.pem
    #  envoy:
    #    address: 0.0.0.0
    #    port: 8002
    #    server-certificate-path: /path/to/server-cert.pem
    #    server-key-path: /path/to/server-private-key.pem
    #    ca-certificate-path: /path/to/root-ca-for-client-validation.pem
    #
    # listener:
    #  connection-balancer: exact
    #  socket-options:
    #    tos: 64
    #    traffic-class: 64
    #
    # omEnforcedHealthListener:
    #   address: 0.0.0.0
    #   port: 8003
```

_Note:_ The default example `contour` includes this [file][1] for easy deployment of Contour.

## Environment Variables

### CONTOUR_NAMESPACE

If present, the value of the `CONTOUR_NAMESPACE` environment variable is used as:

1. The value for the `contour bootstrap --namespace` flag unless otherwise specified.
1. The value for the `contour certgen --namespace` flag unless otherwise specified.
1. The value for the `contour serve --envoy-service-namespace` flag unless otherwise specified.
1. The value for the `contour serve --leader-election-resource-namespace` flag unless otherwise specified.

The `CONTOUR_NAMESPACE` environment variable is set via the [Downward API][6] in the Contour [example manifests][7].

## Bootstrap Config File

The bootstrap configuration file is generated by an initContainer in the Envoy daemonset which runs the `contour bootstrap` command to generate the file.
This configuration file configures the Envoy container to connect to Contour and receive configuration via xDS.

The next section outlines all the available flags that can be passed to the `contour bootstrap` command which are used to customize
the configuration file to match the environment in which Envoy is deployed.

### Bootstrap Flags

There are flags that can be passed to `contour bootstrap` that help configure how Envoy
connects to Contour:

| Flag                                   | Default           | Description                                                                                                                                                                                                  |
| -------------------------------------- |-------------------| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| <nobr>--resources-dir</nobr>           | ""                | Directory where resource files will be written.                                                                                                                                                              |
| <nobr>--admin-address</nobr>           | /admin/admin.sock | Path to Envoy admin unix domain socket.                                                                                                                                                                      |
| <nobr>--admin-port (Deprecated)</nobr> | 9001              | Deprecated: Port is now configured as a Contour flag.                                                                                                                                                        |
| <nobr>--xds-address</nobr>             | 127.0.0.1         | Address to connect to Contour xDS server on.                                                                                                                                                                 |
| <nobr>--xds-port</nobr>                | 8001              | Port to connect to Contour xDS server on.                                                                                                                                                                    |
| <nobr>--envoy-cafile</nobr>            | ""                | CA filename for Envoy secure xDS gRPC communication.                                                                                                                                                         |
| <nobr>--envoy-cert-file</nobr>         | ""                | Client certificate filename for Envoy secure xDS gRPC communication.                                                                                                                                         |
| <nobr>--envoy-key-file</nobr>          | ""                | Client key filename for Envoy secure xDS gRPC communication.                                                                                                                                                 |
| <nobr>--namespace</nobr>               | projectcontour    | Namespace the Envoy container will run, also configured via ENV variable "CONTOUR_NAMESPACE". Namespace is used as part of the metric names on static resources defined in the bootstrap configuration file. |
| <nobr>--xds-resource-version</nobr>    | v3                | Currently, the only valid xDS API resource version is `v3`.                                                                                                                                                  |
| <nobr>--dns-lookup-family</nobr>       | auto              | Defines what DNS Resolution Policy to use for Envoy -> Contour cluster name lookup. Either v4, v6, auto or all.                                                                                                   |
| <nobr>--log-format                     | text              | Log output format for Contour. Either text or json. |
| <nobr>--overload-max-heap              | 0                 | Defines the maximum heap memory of the envoy controlled by the overload manager. When the value is greater than 0, the overload manager is enabled, and when envoy reaches 95% of the maximum heap size, it performs a shrink heap operation. When it reaches 98% of the maximum heap size, Envoy Will stop accepting requests. |
| <nobr>--overload-downstream-max-conn              | 0                 | 	OverloadMaxDownstreamConn defines the envoy global downstream connection limit controlled by the overload manager. When the value is greater than 0 the overload manager is enabled and listeners will begin rejecting connections when the the connection threshold is hit. Metrics and health listeners are not subject to the connection limits, however, they still count against the global limit. |


[1]: {{< param github_url>}}/tree/{{< param branch >}}/examples/contour/01-contour-config.yaml
[2]: config/access-logging
[3]: https://kubernetes.io/docs/concepts/configuration/organize-cluster-access-kubeconfig/
[4]: https://golang.org/pkg/time/#ParseDuration
[5]: https://godoc.org/github.com/projectcontour/contour/internal/envoy#DefaultFields
[6]: https://kubernetes.io/docs/tasks/inject-data-application/environment-variable-expose-pod-information/
[7]: {{< param github_url>}}/tree/{{< param branch >}}/examples/contour
[8]: https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/core/v3/protocol.proto#envoy-v3-api-field-config-core-v3-httpprotocoloptions-idle-timeout
[9]: https://www.envoyproxy.io/docs/envoy/latest/api-v3/extensions/filters/network/http_connection_manager/v3/http_connection_manager.proto#envoy-v3-api-field-extensions-filters-network-http-connection-manager-v3-httpconnectionmanager-stream-idle-timeout
[10]: https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/core/v3/protocol.proto#envoy-v3-api-field-config-core-v3-httpprotocoloptions-max-connection-duration
[11]: https://www.envoyproxy.io/docs/envoy/latest/api-v3/extensions/filters/network/http_connection_manager/v3/http_connection_manager.proto#envoy-v3-api-field-extensions-filters-network-http-connection-manager-v3-httpconnectionmanager-drain-timeout
[12]: https://www.envoyproxy.io/docs/envoy/latest/api-v3/extensions/filters/network/http_connection_manager/v3/http_connection_manager.proto#envoy-v3-api-field-extensions-filters-network-http-connection-manager-v3-httpconnectionmanager-request-timeout
[13]: https://www.envoyproxy.io/docs/envoy/latest/api-v3/extensions/filters/network/http_connection_manager/v3/http_connection_manager.proto#envoy-v3-api-field-extensions-filters-network-http-connection-manager-v3-httpconnectionmanager-delayed-close-timeout
[14]: https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/listener/v3/listener.proto#config-listener-v3-listener-connectionbalanceconfig
[15]: https://www.envoyproxy.io/docs/envoy/latest/api-v3/extensions/filters/network/http_connection_manager/v3/http_connection_manager.proto?highlight=strip_trailing_host_dot