ARG PHP_VERSION=8.0
{{ range $key, $module := .PhpModules -}}
ARG INSTALL_{{ upper $key }}={{ $module.Enabled }}
{{ end }}

ARG PHP_UPLOAD_MAX_FILESIZE {{ .MaxUploadSize }}
ARG PHP_POST_MAX_SIZE {{ .MaxUploadSize }}

# Backend Builder
FROM composer:2 as php-builder

WORKDIR /app

COPY composer.json composer.lock{{ if eq .AdminGen "Nova" }} auth.json{{ end }} ./
{{- if eq .AdminGen "Nova" }}
COPY nova-components/ nova-components/
{{- end }}

RUN set -x \
    && composer install \
        --ignore-platform-reqs \
        --no-autoloader \
        --no-interaction \
        --no-progress \
        --no-suggest

COPY . ./
RUN set -x \
    && export TELESCOPE_ENABLED=false \
    && composer dump-autoload \
        --classmap-authoritative \
        --no-interaction \
    {{- if eq .AdminGen "Nova" }}
    && php artisan nova:publish \
    {{- end }}
    && php artisan vendor:publish --tag=public

# Frontend Builder
FROM node:lts-alpine as node-builder

WORKDIR /app

RUN set -x \
    && apk add --no-cache \
        autoconf \
        automake \
        bash \
        g++ \
        libc6-compat \
        libjpeg-turbo \
        libjpeg-turbo-dev \
        libpng \
        libpng-dev \
        libtool \
        libwebp \
        libwebp-dev \
        make \
        nasm

COPY artisan package.json package-lock.json webpack.mix.js ./

RUN set -x \
    && npm ci

COPY public/ public/
COPY resources/ resources/

ARG NODE_ENV=production
ARG ENV=local

RUN set -x \
    && if [ "$ENV" != "local" ]; then \
        npm run production \
    ; fi

# Final Image
FROM clevyr/php:$PHP_VERSION

WORKDIR /app

COPY --chown=root docker/app/rootfs /

RUN crontab /etc/cron.d/scheduler

COPY --from=php-builder --chown=82:82 /app .
COPY --from=node-builder --chown=82:82 /app/public public/

CMD ["s6-svscan", "/etc/s6/app"]
