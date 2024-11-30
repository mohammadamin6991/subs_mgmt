
## How to prepare the development environment

To start developing or running in a test environment, you need to use [devcontainers](https://code.visualstudio.com/docs/devcontainers/containers) (You must have docker installed as well). After installing the vscode extension, use CTRL+P and type `open in dev container`

1. Go to `project` directory and run `make up_build`
    ```bash
    cd project
    make up_build
    ```
2. Go to pgadmin at localhost:15432, login to the development server using `postgres` as uesrname and passowrd.
3. Create 3 Databases
    - subs_auth
    - subs_svc_s3
    - subs_mgmt

4. Run database migration using make command

    ```bash
    cd project
    make migrate_clean
    ```

TODO:
- [x] Fix docker network call between development and make-deployed stack
- [x] Install golang-migrate in development container
