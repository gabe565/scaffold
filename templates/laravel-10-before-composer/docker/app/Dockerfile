ARG PHP_VERSION=8.1
ARG NODE_VERSION=16
{{- range $key, $module := .PhpModules }}
ARG INSTALL_{{ upper $key }}={{ $module.Enabled }}
{{- end }}
ARG PHP_UPLOAD_MAX_FILESIZE {{ .MaxUploadSize }}
ARG PHP_POST_MAX_SIZE {{ .MaxUploadSize }}


# Backend build
FROM composer:2.0 as php-builder
WORKDIR /app

COPY composer.json composer.lock auth.json ./
{{- if .ComposerDeps.ModuleEnabled "laravel/nova" }}
COPY nova-components/ nova-components/
{{- end }}
RUN composer install \
        --ignore-platform-reqs \
        --no-autoloader \
        --no-interaction \
        --no-progress \
        --no-suggest

COPY . .
RUN set -x \
    && export TELESCOPE_ENABLED=false \
    && composer dump-autoload \
        --classmap-authoritative \
        --no-interaction \
    {{- if .ComposerDeps.ModuleEnabled "laravel/nova" }}
    && php artisan nova:publish \
    {{- end }}
    && php artisan vendor:publish --tag=public


# Frontend build
FROM node:$NODE_VERSION-alpine as node-builder
WORKDIR /app

COPY package.json package-lock.json ./
RUN npm ci

COPY --from=php-builder /app .
ARG NODE_ENV=production
RUN npm run build


# Local image
FROM ghcr.io/clevyr/php:$PHP_VERSION as local-image
WORKDIR /app

COPY --chown=root docker/app/rootfs /
RUN crontab /etc/cron.d/scheduler

CMD ["s6-svscan", "/etc/s6/app"]


# Deployed image
FROM local-image

COPY --from=php-builder --chown=82:82 /app .
COPY --from=node-builder --chown=82:82 /app/public public/
