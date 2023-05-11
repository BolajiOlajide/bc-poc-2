# BC-POC-2

* Generate an installation token.
* Run the POC with the command

```sh
go run main.go -token <TOKEN> -repo <REPO>
```

* Make sure the github app has access to the repository.
* Confirm the pull request was created successfully.
* You can optionally pass in the `-branch` flag to specify a custom branch name, the default branch name is `bc-poc`

