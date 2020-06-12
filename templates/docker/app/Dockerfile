{{ range $key, $module := .Modules.Modules -}}
ARG INSTALL_{{ upper $key }}={{ $module.Enabled }}
{{ end }}
# Backend Builder
FROM clevyr/prestissimo as php-builder

WORKDIR /app

COPY composer.json composer.lock {{ if eq .AdminGen "Nova" }}auth.json{{ end -}} ./
{{- if eq .AdminGen "Nova" }}
COPY nova/ nova/
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

COPY artisan package.json package-lock.json webpack.mix.js ./

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

RUN set -x \
    && npm ci

COPY public/ public/
COPY resources/ resources/

ARG NODE_ENV=production
ARG ENV=local

RUN set -x \
    && if [ "$ENV" != "local" ]; then \
        node_modules/webpack/bin/webpack.js \
            --no-progress \
            --hide-modules \
            --config=node_modules/laravel-mix/setup/webpack.config.js; \
    fi

# Final Image
FROM clevyr/laravel-base-image-builder:7.4

WORKDIR /app
VOLUME /app

COPY --chown=root docker/app/rootfs /

RUN crontab /etc/cron.d/scheduler

COPY --from=php-builder --chown=82:82 /app .
COPY --from=node-builder --chown=82:82 /app/public public/

CMD ["s6-svscan", "/etc/s6/app"]
