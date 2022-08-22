{{ .AppSlug }}
===

## To Run

Run the following:
```
cp ~/.composer/auth.json .
cp .env.example .env
docker-compose up
```

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
./vendor/bin/pint
```

## Testing
```
# Run the PHP Pest / UnitTest Tests
docker-compose exec app php artisan test

# Run the Laravel Dusk (i.e. browser-driven) tests
# NOTE: You have to disable the Vite hot reloader to run these.
docker-compose stop hot && docker-compose exec app php artisan pest:dusk && docker-compose start hot
```
