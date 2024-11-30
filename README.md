
## How to prepare the development environment

To start developing or running in a test environment, you need to use [devcontainers](https://code.visualstudio.com/docs/devcontainers/containers) (You must have docker installed as well). After installing the vscode extension, use CTRL+P and type `open in dev container`

1. Go to `project` directory and run `make up_build`
    ```bash
    cd project
    make up_build
    ```
2. Go to pgadmin at localhost:15432 and login using this detail
    ```yaml
    user: admin@admin.com
    password: admin
    ```

3. Add a server with this detail

    ```yaml
    host: localhost
    port: 5432
    username: postgres
    password: postgres
    ```

4. Create 3 Databases
    - subs_auth
    - subs_svc_s3
    - subs_mgmt

5. Run database migration using make command

    ```bash
    cd project
    make migrate_clean
    ```

TODO:
- [x] Fix docker network call between development and make-deployed stack
- [x] Install golang-migrate in development container
