# snyk-cr-import

## Description:
Import container images into Snyk from your Container Registry. (ACR, ECR, Artifactory)

# Usage
By default the whole repo will be imported if the `file_path` flag is not specified.
- `--help` - show help and all available flags
- `--token=` - Snyk API token
- `--crId=` - Container Registry ID: This can be found on the Integration page in the Settings area for all integrations that have been configured. https://app.snyk.io/org/*org-name*/manage/integrations Each org will have a unique integration ID
- `--orgId=` - Snyk Target/Destination Organization ID
- `--imageName=` - Name of the container image including the tag: hezro/juice-shop:1.1.0
- `@args` - Optional: You can also read one or more arguments from a file. 

# Examples
You can find usage instructions by running:

```bash
snyk-cr-import --help
```

How to import a container:
```bash
snyk-cr-import --token=<Snyk Token> --crId=<container registry integration id> --orgId=<organization id> --imagehName=<image name and tag>
```

Use the args files to pass in the Snyk Token:
```bash
echo --token=<snyk token> > args
```
```bash
snyk-cr-import @args --crId=<container registry integration id> --orgId=<organization id> --imagehName=<image name and tag>
```
