#
# Builder image
#

FROM golang:1.14.4-alpine AS build-env
ARG SOURCE=*

ADD $SOURCE /src/
WORKDIR /src/

RUN apk update && apk add make git

# Unpack any tars, then try to execute a Makefile, but if the SOURCE url is
# just a tar of binaries, then there probably won't be one. Using multiple RUN
# commands to ensure any errors are caught.
RUN find . -name '*.tar.gz' -type f | xargs -rn1 tar -xzf
RUN if [ -f Makefile ]; then make; fi
RUN cp "$(find . -name 'louketo-proxy' -type f -print -quit)" /louketo-proxy

#
# Actual image
#

FROM alpine:3.12

LABEL Name=louketo-proxy \
      Release=https://github.com/louketo/louketo-proxy \
      Url=https://github.com/louketo/louketo-proxy \
      Help=https://github.com/louketo/louketo-proxy/issues

WORKDIR "/opt/louketo"

COPY --from=build-env /louketo-proxy ./
RUN chmod +x louketo-proxy

ENTRYPOINT [ "/opt/louketo/louketo-proxy" ]
