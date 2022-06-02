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

## Testing
```
# Run the PHP Tests
docker-compose exec app php artisan test
```
