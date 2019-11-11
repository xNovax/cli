# Doppler CLI

## Warning: This tool is current pre-release. For the current stable version, please use the [node-cli](https://github.com/DopplerHQ/node-cli).

## Install

TODO: this section

## Development

Build the app:

`go build`

To specify the version at build time:

`go build -ldflags '-X github.com/DopplerHQ/cli/version.ProgramVersion=YOUR_VERSION'`

Or even at run time:

`go run -ldflags '-X github.com/DopplerHQ/cli/version.ProgramVersion=YOUR_VERSION'` main.go YOUR_ARGS

### Generate a GPG key

Store the keys and passphrase in your doppler config

```
gpg --full-generate-key
gpg --list-secret-keys  # copy the key's 40-character ID
gpg --output secret.key --armor --export-secret-key KEY_ID
gpg --output public.key --armor --export KEY_ID
gpg --keyserver pgp.mit.edu --send-key LAST_8_DIGITS_OF_KEY_ID
```

### Test

Test building for all targets:

`goreleaser release --snapshot --skip-publish --rm-dist`


### Release

To release a new version, run:

```
make release VERSION=vX.Y.Z
```

This command will push local changes to Origin, create a new tag, and push the tag to Origin. It will then build and release the doppler binaries.

Note: The release will automatically fail if the tag and HEAD have diverged
`   ⨯ release failed after 0.13s error=git tag v0.0.2 was not made against commit c9c6950d18790c17db11fedae331a226f8f12c6b`

### Help

Issue: `gpg: signing failed: Inappropriate ioctl for device`
Fix: `export GPG_TTY=$(tty)`
