# Zebi Scaper

This is the repository made with my go temple gogoleplate.

##Todo

- [ ] Instagram followers scraping 
- [ ] Send rapport by email


## Requirements

To run this project you'll need to have docker and yarn installed on your machine.
If you want to develop on this project it's recommended to have `golang` installed on your machine.

## Project setup

First you need to created your `.env` file (you can use the .env.dist file).

```sh
  docker-compose up --build
```

### RSA keys (needed to generate and read tokens)

You will need 2 files placed at the root of your project : `public.pem` and `private.pem`.

```sh
    # use the following password: private_key
    openssl genrsa -des3 -out private.pem 2048
    openssl rsa -in private.pem -outform PEM -pubout -out public.pem
```
# Commit naming conventions

If you want to contribute to the project you'll need to name your commits according to the following convention :

    type(action): description.

The following types are available :

-   **feat** : Feature commit.
-   **fix** : Fixing a bug.
-   **update** : Updating code or dependencies.
-   **revert** : Removing changes.
-   **doc** : Add changes to the documentation.
-   **refacto** : Refactoring code.
-   **build** : Modifications linked to the infrastructure.
-   **param** : Modifications on the config.

## Branch naming convention

The branch should have a name that reflects it's purpose.

The convention is to prefix the branch name with `feature-`. All the words must be separated by a `-`.

Branch name example : `feature-user-authentication`

## Routes

the available routes are the following :

| route                       | method | description                       | params          | response                |
| --------------------------- | ------ | --------------------------------- | --------------- | ----------------------- |
| `/hello`                    | GET    | check if the api is available     | none            | `Hello world!`          |
