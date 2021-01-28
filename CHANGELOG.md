# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- The `detect-ajax` option to prevent redirection for requests with a `X-Requested-With: XMLHttpRequest` header

### Removed

- The `access-token-duration` in favor of `default-session-duration`, which controls all data and not only `access-token` as the config was supposed to control

### Fixed

- Trigger offline access token when there's `offline_access` in scopes instead of just `offline` to adhere to the [spec](https://openid.net/specs/openid-connect-core-1_0.html#OfflineAccess)
- Multiple requests cancelling each other due to all sharing a dynamic key that's cleaned by the fastest
- Access cookie expiration date same as token expiration causing requests to be denied when there's a valid refresh token
- Wrong expiration dates due to parsing the refresh-token as if it were an id-token

## [2.3.0] - 2018-07-30

### Added

- Ability to use a "any" operation on the roles rather then just "and" with the inclusion of a `require-any-role` [#PR389](https://github.com/gambol99/keycloak-proxy/pull/389)
- The `--enable-request-id` option to inject a request id into the upstream request [#PR392](https://github.com/gambol99/keycloak-proxy/pull/392)
- Ability for the proxy to generate self-signed certificates for use via the `--enable-self-signed-tls` [#PR394](https://github.com/gambol99/keycloak-proxy/pull/394)
- Support for token with multiple audiences in the claims [#PR401](https://github.com/gambol99/keycloak-proxy/pull/401)
- The `--max-idle-connections` and `--max-idle-connections-per-host` settings to support tuning the http connection pool size for performance needs [#PR405](https://github.com/gambol99/keycloak-proxy/pull/405)

### Changed

- The `http-cookie-only` option as default true [#PR397](https://github.com/gambol99/keycloak-proxy/pull/397)

## [2.2.2] - 208-07-03

### Added

- Ability to the add response headers via `--response-headers` [#PR386](https://github.com/gambol99/keycloak-proxy/pull/386)

## [2.2.1] - 2018-05-31

### Fixed

- The logout handler, when logout redirection is enable if no redirect param is given we default to the hostname [#PR375](https://github.com/gambol99/keycloak-proxy/pull/375)

## [2.2.0] - 2018-05-30

### Added

- A `--enable-default-deny` option to make denial by default [#PR320](https://github.com/gambol99/keycloak-proxy/pull/320)
- A `enable-logout-redirect` which redirects the /oauth/logout to the provider [#PR327](https://github.com/gambol99/keycloak-proxy/pull/327)
- Environment variables alternatives for the forwarding username and password [#PR329](https://github.com/gambol99/keycloak-proxy/pull/329)
- Metrics latency metrics for the forwarding proxy and the certificate rotation [#PR325](https://github.com/gambol99/keycloak-proxy/pull/325)
- Spelling check to the tests [#PR322](https://github.com/gambol99/keycloak-proxy/pull/322)
- The X-Auth-Audience to the upstream headers [#PR319](https://github.com/gambol99/keycloak-proxy/pull/319)
- The ability to control the timeout on the initial openid configuration from .well-known/openid-configuration [#PR315](https://github.com/gambol99/keycloak-proxy/pull/315)
- A --preserve-host option to preserve the host header of the proxied request in the upstream request [#PR328](https://github.com/gambol99/keycloak-proxy/pull/328)
- The feature to customize the oauth prefix (defaults to /oauth) [#PR326](https://github.com/gambol99/keycloak-proxy/pull/326)
- Additional metrics covering provider request latency, token breakdown [#PR324](https://github.com/gambol99/keycloak-proxy/pull/324)
- The `--enable-session-cookies` to enable session only cookies [#PR357](https://github.com/gambol99/keycloak-proxy/pull/357)
- A ability to match string arrays claims [#PR364](https://github.com/gambol99/keycloak-proxy/pull/364)
- A warning messaage to indicate disabling the write-timeout when using pprof [#PR370](https://github.com/gambol99/keycloak-proxy/pull/370)
- A method check for the hijacker [#PR302](https://github.com/gambol99/keycloak-proxy/pull/302)

### Changed

- The upstream-keepalive to default to true [#PR321](https://github.com/gambol99/keycloak-proxy/pull/321)
- Force configuration to use the wildcard [#PR338](https://github.com/gambol99/keycloak-proxy/pull/338)
- The docker base image alpine 3.7 [#PR313](https://github.com/gambol99/keycloak-proxy/pull/313)
- Golang version 1.10 [#PR316](https://github.com/gambol99/keycloak-proxy/pull/316)
- Imported the fix to the cache headers from upstream go-oidc [#PR341](https://github.com/gambol99/keycloak-proxy/pull/341)
- Switched from glide to dep lib tool [#PR373](https://github.com/gambol99/keycloak-proxy/pull/373)
- Switched to using SHA256 from MD5 for the token hash [#PR350](https://github.com/gambol99/keycloak-proxy/pull/350)
- Switched to using golang v1.10.2 [#PR374](https://github.com/gambol99/keycloak-proxy/pull/374)
- Cookies session on by default and turning the default denial on [#PR368](https://github.com/gambol99/keycloak-proxy/pull/368)

### Fixed

- Up the redirect_uri to the logout [#PR365](https://github.com/gambol99/keycloak-proxy/pull/365)
- A redirection bug [#PR337](https://github.com/gambol99/keycloak-proxy/pull/337)
- Updated the go-oidc to fix the cache header [issues](https://github.com/gambol99/keycloak-proxy/issues/340)[#PR339](https://github.com/gambol99/keycloak-proxy/pull/339)
- Up the readme indicating we can run without client secret [#PR342](https://github.com/gambol99/keycloak-proxy/pull/342)
- Up the redirect url in the logout handler [#PR345](https://github.com/gambol99/keycloak-proxy/pull/345)
- Switched to using the upstream stream goproxy [#PR349](https://github.com/gambol99/keycloak-proxy/pull/349)
- Removing the unused code [#PR352](https://github.com/gambol99/keycloak-proxy/pull/352)
- Reducing the aggressive timeouts on the upstream [#PR354](https://github.com/gambol99/keycloak-proxy/pull/354)
- The issue with a zero exp claim [#PR355](https://github.com/gambol99/keycloak-proxy/pull/355)

## [2.1.1] - 2018-02-12

### Added

- The groups parameter to the resource, permitting users to use the `groups` claim in the token [#PR301](https://github.com/gambol99/keycloak-proxy/pull/301)

### Changed

- Removed the authors file [#PR299](https://github.com/gambol99/keycloak-proxy/pull/299)

### Fixed

- The custom headers when upgrading to websockets [#PR311](https://github.com/gambol99/keycloak-proxy/pull/311)
- Exception when upgrading to websockets [#PR303](https://github.com/gambol99/keycloak-proxy/pull/303)

## [2.1.0] - 2017-12-21

### Added

- Changed the routing engine from gin to echo
- We now normalize all inbound URI before applying the protection middleware
- The order of the resources are no longer important, the framework will handle the routing [#PR199](https://github.com/gambol99/keycloak-proxy/pull/199)
- Improved the overall spec of the proxy by removing URL inspection and prefix checking [#PR199](https://github.com/gambol99/keycloak-proxy/pull/199)
- Removed the CORS implementation and using the default echo middles, which is more compliant [#PR199](https://github.com/gambol99/keycloak-proxy/pull/199)
- A warning for suspect resource urls not using wildcards [#PR206](https://github.com/gambol99/keycloak-proxy/pull/206)
- A build time to the version tag [#PR212](https://github.com/gambol99/keycloak-proxy/pull/212)
- Coveralls coverage submission to the ci build [#PR215](https://github.com/gambol99/keycloak-proxy/pull/215)
- Spelling code coverage to the ci build [#PR208](https://github.com/gambol99/keycloak-proxy/pull/208)
- Update the encryption to use aes gcm [#PR220](https://github.com/gambol99/keycloak-proxy/pull/220)
- The --enable-encrypted-token option to enable encrypting the access token:wq
- The --skip-client-id option to permit skipping the verification of the auduence against client in token [#PR236](https://github.com/gambol99/keycloak-proxy/pull/236)
- Updated the base image to apline 3.6 in commit [0fdebaf821](https://github.com/gambol99/keycloak-proxy/pull/236/commits/0fdebaf8215e9480896f01ec7ab2ef7caa242da1)
- Moved to use zap for the logging [#PR237](https://github.com/gambol99/keycloak-proxy/pull/237)
- Making the X-Auth-Token optional in the upstream headers via the --enable-token-header [#PR247](https://github.com/gambol99/keycloak-proxy/pull/247)
- The ability to load a CA authority to provide trust on upstream endpoint [#PR248](https://github.com/gambol99/keycloak-proxy/pull/248)
- The ability to set various http server and upstream timeout [#PR268](https://github.com/gambol99/keycloak-proxy/pull/268)
- The `--enable-authorization-cookies` command line option to control upstream cookies [$PR287](https://github.com/gambol99/keycloak-proxy/pull/287)
- Docker image instructions to the readme [#PR204](https://github.com/gambol99/keycloak-proxy/pull/204)
- Unit tests for the debug handlers [#PR223](https://github.com/gambol99/keycloak-proxy/pull/223)

### Changed

- The proxy no longer uses prefixes for resources, if you wish to use wildcard urls you need to specify it, i.e. `--resource=/` becomes `--resource=/*` or `=admin/` becomes `=admin/*` or `/admin*`; a full set of routing details can be found [here](https://echo.labstack.com/guide/routing) [#PR199](https://github.com/gambol99/keycloak-proxy/pull/199)
- Removed the --enable-cors-global option, CORS is now handled the default echo middleware
- Option from log-requests -> enable-logging [#PR199](https://github.com/gambol99/keycloak-proxy/pull/199)
- Option from json-format -> enable-json-logging [#PR199](https://github.com/gambol99/keycloak-proxy/pull/199)

### Fixed

- The parsing of slices for command line arguments (i.e. --cors-origins etc)
- Any accidental proxying on the /oauth or /debug URI
- Removed all references to the underlining web framework in tests
- Adding unit tests for proxy protocol and using the run() method [#PR214](https://github.com/gambol99/keycloak-proxy/pull/214)
- Removed unnecessary commands in the Dockerfile [#PR213](https://github.com/gambol99/keycloak-proxy/pull/213)
- Removed the unrequired testing tools [#PR210](https://github.com/gambol99/keycloak-proxy/pull/210)
- A number of linting errors highlighted by gometalinter [#PR209](https://github.com/gambol99/keycloak-proxy/pull/209)
- The logout handler panic when revocation url is not set [#PR254](https://github.com/gambol99/keycloak-proxy/pull/254)
- The Host header on the forwarding proxy [#PR290](https://github.com/gambol99/keycloak-proxy/pull/290)

## [2.0.7] - 2017-05-18

### Fixed

- Backported Fix to the proxy proxy call [767967c3](https://github.com/gambol99/keycloak-proxy/commit/767967c3499795e3141e74cace5ae3d70f27cf61)

## [2.0.6] - 2017-03-24

### Fixed

- Ensuring we abort all requests to /oauth/ [#PR205](https://github.com/gambol99/keycloak-proxy/pull/205)

## [2.0.5] - 2017-03-22

### Fixed

- We normalize all urls before the protection middleware is applied [#PR202](https://github.com/gambol99/keycloak-proxy/pull/202)

## [2.0.4] - 2017-03-17

### Fixed

- Bug in authentication, which permitted double slashed url entry [#PR200](https://github.com/gambol99/keycloak-proxy/pull/200)

### Added

- Grabbing the revocation-url from the idp config if user override is not specified [#PR193](https://github.com/gambol99/keycloak-proxy/pull/193)

## [2.0.3] - 2017-02-07

### Added

- The PROXY_ENCRYPTION_KEY environment varable [#PR191](https://github.com/gambol99/keycloak-proxy/pull/191)

## [2.0.2] - 2017-02-06

### Added

- The --enable-cors-global to switch on CORs header injects into every response [#PR174](https://github.com/gambol99/keycloak-proxy/pull/174)
- The ability to reload the certificates when the change [#PR178](https://github.com/gambol99/keycloak-proxy/pull/178)
- Removing the requirement of a redirection-url, if none is specified it will use Host header or the X-Forwarded-Host if present [#PR183](https://github.com/gambol99/keycloak-proxy/pull/183)
- A new options to control the access token duration [#PR188](https://github.com/gambol99/keycloak-proxy/pull/188)

### Changed

- Updated the gin dependency to latest version and removed dependency in tests for gin [#PR181](https://github.com/gambol99/keycloak-proxy/pull/181)
- Updated to go-proxy to the latest version [#PR180](https://github.com/gambol99/keycloak-proxy/pull/180)
- The CLI to use reflection of the config struct [#PR176](https://github.com/gambol99/keycloak-proxy/pull/176)
- Updated the docker base image to alpine:3.5 [#PR184](https://github.com/gambol99/keycloak-proxy/pull/184)

### Fixed

- Spelling mistakes [#PR177](https://github.com/gambol99/keycloak-proxy/pull/177)
- The time.Duration flags in the reflection code [#PR173](https://github.com/gambol99/keycloak-proxy/pull/173)
- The environment variable type [#PR176](https://github.com/gambol99/keycloak-proxy/pull/176)
- The refresh tokens, the access token cookie was timing out too quickly ([#PR188](https://github.com/gambol99/keycloak-proxy/pull/188)

## [2.0.1] - 2017-01-09

### Fixed

- The cli option for --resources. Need to start writing tests for the cli options

## [2.0.0] - 2017-01-06

### Added

- The --skip-openid-provider-tls-verify option to bypass the TLS verification for Idp [#PR147](https://github.com/gambol99/keycloak-proxy/pull/147)
- A http service to permit http -> https redirects --enable-https-redirect [#PR126](https://github.com/gambol99/keycloak-proxy/pull/162)
- A pprof debug handler to support profiling the proxy, via --enable-profiling [#PR156](https://github.com/gambol99/keycloak-proxy/pull/156)

### Fixed

- The --headers and --tags command line options, had a typo on the mergeMaps method [#PR142](https://github.com/gambol99/keycloak-proxy/pull/142)
- Cleaned up how the cli command line options are processed [#PR164](https://github.com/gambol99/keycloak-proxy/pull/164)
- Cleaned up the option checking for forwarding proxy tls setting [#PR163](https://github.com/gambol99/keycloak-proxy/pull/163)
- Using timeout rather than multiple attempts for discovery url [#PR153](https://github.com/gambol99/keycloak-proxy/pull/153)
- Updated the go-oidc library with various fixes [#PR159](https://github.com/gambol99/keycloak-proxy/pull/159)

### Changed

- The login handler by default has been switched off, you must enable for --enable-login-handler
- The CORS format in the configuration file
- The command line options scope -> scopes
- The command line options log-json-format -> json-format
- The command line options resource -> resources
- The command line options tags -> tags

## [1.2.8] - 2016-09-30

### Fixed

- A bug in the --cookie-domain options
- Unit test for the cookie-domain options
- Switched to using set rather than add to the headers

## [1.2.7] - 2016-09-16

### Added

- Allow the user to enable or disable adding the Authorization header

### Fixed

- Unit tests for the logout handlers
- Unit tests for the authorization header handling

## [1.2.6] - 2016-08-26

### Added

- Ability to control the http-only cookie option, default to false

### Fixed

- Revocation url bug

## [1.2.5] - 2016-08-25

### Fixed

- The /oauth/login handler to return 401 on failed logins

## [1.2.4] - 2016-08-13

### Added

- Ability to set the forwarding proxy certificates
- Logging for outbound forward signing requests

### Changed

- Removed the --idle-duration option, was never really implemented well

### Fixed

- The expiration of the access token, if no idle-duration is
- The forwarding proxy for SSL
- The bug in the containedSubString method
- Up the config resource definition to use 'uri' not 'url'

## [1.2.3] - 2016-07-18

### Added

- A prometheus metrics endpoint, at present a break down by status_code is provided
- The ability to override the cookie domain from the default host header
- The ability to load a client certificate used by the reverse and forwarding upstream proxies.

### Changed

- The godeps for codegangsta cli to it's renamed version

### Fixed

- The environment variable command line options, the IsSet in cli does not check environment variable setters

## [1.2.2] - 2020-07-05

### Changed

- General Code fix uo
- Removing from dockerfile user and group

## [1.2.1] - 2016-06-21

### Changed

- The dockerfile to create a user and group and not run at root

## [1.2.0] - 2016-06-20

### Changed

- The /oauth/login handler to use post form values rather than query parameter to ensure (to a degree) they are not logged

## [1.1.1] - 2016-06-16

### Fixed

- The configuration bug which required a redirection-url even when redirection was shifted off

## [1.1.0] - 2016-06-11

### Added

- Environment variables to some of the command line options
- The option of a forwarding agent, i.e. you can seat the proxy front of your application, login to keycloak and use the proxy as forwarding agent to sign outbound requests.
- The version information into a header on /oauth/health endpoint
- Removed the need to specify a client-secret, which means to cope with authz only or public endpoints
- Added role url tokenizer, /auth/%role%/ will extract the role element and check the token as it
- Added proxy protocol support for the listening socket (--enable-proxy-protocol=true)
- Added the ability to listen on a unix socket
- A auto build to quay.io on the travis build for master and tags
- A git+sha to the usage

### Changed

- The state parameter (which is used as a redirect) to base64 the value allowing you to use complex urls
- The X-Auth-Subject, it not is the actual subject from the token (makes more sense). X-Auth-UserID will either be the subject id or the preferred username

### Fixed

- The host header to proxy to upstreams outside of the proxy domain - [issue](https://github.com/golang/go/issues/7618)
- Default to gin mode release unless verbose is true
- Removed the gin debug logging for tests and builds
- Removed the default upstream, as it caught people by surprise and some accidentally forwarded to themselves

## [1.0.6] - 2016-05-06

### Fixed

- The logout endpoint, ensuring users sessions are revoked. Note: i've not really tested this against Keycloak and Google. Revocation or logouts seems to have somewhat scattered implementation across providers.

## [1.0.5] - 2016-05-04

### Added

- You can choose the cookie name of the access and refresh token via --cookie-{access,refresh}-name
- An additional option --add-claims to inject custom claims from the token into the authentication headers i.e. --add-claims=given_name would add X-Auth-Given-Name (assumed the claims exists)
- Added the --secure-cookie option to control the 'secure' flag on the cookie

### Changed

- The claims option from 'claims' to 'match-claims' (command line and config)
- Keepalive config option to the same as the command line 'keepalive' -> 'upstream-keepalives'
- The config option from 'upstream' to 'upstream-url', same as command line

## [1.0.4] - 2016-04-30

### Fixed

- The cookie sessions expiration

### Added

- A idle duration configuration option which controls the expiration of access token cookie and thus session. If the session is not used within that period, the session is removed.
- The upstream endpoint has also be a unix socket

### Changed

- The client id in json/yaml config file from clientid -> client-id

## [1.0.2] - 2016-04-22

### Fixed

- Cleaned up a lot of code base to make this simpler
- Elements in the refresh tokens and simplified the controller
- Removed of the code out from methods into functions to reduce the dependencies (unit testing is easier as well)
- How the refresh tokens are implemented, i was somewhat confused between refresh token and offline token
- The encryption key length, must be either 16 or 32 for aes-128/256 selection

### Added

- The ability to store the refresh token in either local boltdb file or a redis service rather than an encrypted cookie (note, the token regardless is encrypted)
- A /oauth/logout endpoint to logout the user
- A /oauth/login (niche requirement) to provide grant_type=password requests

### Changed

- Rename configuration options to conform to their command line equivalents
  - refresh_sessions   -> refresh-sessions
  - discovery_url      -> discovery-url
  - redirection_url    -> redirection-url
  - tls_ca_certificate -> tls-ca-certificate
  - tls_private_key    -> tls-private-key
  - tls_cert           -> tls-cert
  - log_json_format    -> log-json-format
  - log_requests       -> log-requests
  - forbidden_page     -> forbidden-page
  - sign_in_page       -> sign-in-page
  - secret             -> client-secret

## [1.0.1] - 2020-04-10

### Fixed

- The refresh tokens for those provides whom do not use JWT tokens, Google Connect for example

## [1.0.0] - 2016-04-08

### Added

- The /oauth/expiration controller to test for access token expiration
- The /oauth/token as a helper method to display the access token

[Unreleased]: https://github.com/arquivei/louketo-proxy/compare/v2.3.0...HEAD
[2.3.0]: https://github.com/arquivei/louketo-proxy/compare/2.2.2...2.3.0
[2.2.2]: https://github.com/arquivei/louketo-proxy/compare/2.2.1...2.2.2
[2.2.1]: https://github.com/arquivei/louketo-proxy/compare/2.2.0...2.2.1
[2.2.0]: https://github.com/arquivei/louketo-proxy/compare/2.1.1...2.2.0
[2.1.1]: https://github.com/arquivei/louketo-proxy/compare/2.1.0...2.1.1
[2.1.0]: https://github.com/arquivei/louketo-proxy/compare/2.0.7...2.1.0
[2.0.7]: https://github.com/arquivei/louketo-proxy/compare/2.0.6...2.0.7
[2.0.6]: https://github.com/arquivei/louketo-proxy/compare/2.0.5...2.0.6
[2.0.5]: https://github.com/arquivei/louketo-proxy/compare/2.0.4...2.0.5
[2.0.4]: https://github.com/arquivei/louketo-proxy/compare/2.0.3...2.0.4
[2.0.3]: https://github.com/arquivei/louketo-proxy/compare/2.0.1...2.0.3
[2.0.1]: https://github.com/arquivei/louketo-proxy/compare/2.0.0...2.0.1
[2.0.0]: https://github.com/arquivei/louketo-proxy/compare/1.2.8...2.0.0
[1.2.8]: https://github.com/arquivei/louketo-proxy/compare/1.2.7...1.2.8
[1.2.7]: https://github.com/arquivei/louketo-proxy/compare/1.2.6...1.2.7
[1.2.6]: https://github.com/arquivei/louketo-proxy/compare/1.2.5...1.2.6
[1.2.5]: https://github.com/arquivei/louketo-proxy/compare/1.2.4...1.2.5
[1.2.4]: https://github.com/arquivei/louketo-proxy/compare/1.2.3...1.2.4
[1.2.3]: https://github.com/arquivei/louketo-proxy/compare/1.2.2...1.2.3
[1.2.2]: https://github.com/arquivei/louketo-proxy/compare/1.2.1...1.2.2
[1.2.1]: https://github.com/arquivei/louketo-proxy/compare/1.2.0...1.2.1
[1.2.0]: https://github.com/arquivei/louketo-proxy/compare/1.1.1...1.2.0
[1.1.1]: https://github.com/arquivei/louketo-proxy/compare/1.0.6...1.1.1
[1.0.6]: https://github.com/arquivei/louketo-proxy/compare/1.0.5...1.0.6
[1.0.5]: https://github.com/arquivei/louketo-proxy/compare/1.0.4...1.0.5
[1.0.4]: https://github.com/arquivei/louketo-proxy/compare/1.0.2...1.0.4
[1.0.2]: https://github.com/arquivei/louketo-proxy/compare/1.0.1...1.0.2
[1.0.1]: https://github.com/arquivei/louketo-proxy/compare/1.0.0...1.0.1
[1.0.0]: https://github.com/arquivei/louketo-proxy/releases/tag/v1.0.0
