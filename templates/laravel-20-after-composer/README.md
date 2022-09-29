{{ .AppSlug }}
===

## To Run

Run the following:
```
cp ~/.composer/auth.json .
cp .env.example .env
docker-compose up
```
**Trusted Certificate issues?**

##### *macOS*

If you are running having trusted certificate issues execute the following command to add local development trusted certs to your local keychain.

```
security add-trusted-cert -r trustRoot -k ~/Library/Keychains/login.keychain-db ~/.config/caddy/caddy/pki/authorities/local/root.crt
```

You may also need to navigate to the following URLs to accept the cert for the
first time:
```
https://clevyr.run
https://hot.clevyr.run
```

**Running in Firefox and still having issues?** 

If you are having issues with loading try setting the following flag.

The following flag toggles the feature that prevents certificate authorities (CAs) 

1. Type **about:config** in the address bar and press `` Return``
2. Type **enterprise** in the *Search* field.
3. *Toggle* the preference **security.enterprise_roots.enabled**

## Linting
```
# PHP Linting
docker-compose exec app ./vendor/bin/phpstan analyse

# JS Linting
docker-compose exec hot npm run lint
docker-compose exec hot npm run stylelint-lint
```

## [Laravel Pint](https://laravel.com/docs/9.x/pint)
Laravel Pint is an opinionated PHP code style fixer for minimalists.
```
docker-compose exec app ./vendor/bin/pint
```

## Testing
```
# Run the PHP Pest / UnitTest Tests
docker-compose exec app php artisan test

# Run the Laravel Dusk (i.e. browser-driven) tests
# NOTE: You have to disable the Vite hot reloader to run these.
docker-compose stop hot && docker-compose exec app php artisan pest:dusk && docker-compose start hot
```
