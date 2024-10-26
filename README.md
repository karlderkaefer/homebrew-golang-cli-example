# genderize

This is a simple example project to demonstrate how to structure a Go project.
This cli just predicts the gender of a given name.
This is just used as useless example for CI purposes.

## Installation

Since the repo is public you actually don't need a token to install the formula.
This just shows a way how to install formulas from private repos.

```bash
brew install gh
gh auth login -s repo,read:org
export HOMEBREW_GITHUB_API_TOKEN="$(gh auth token)"
brew tap karlderkaefer/homebrew-tap
brew install genderize
```

Upgrade

```bash
export HOMEBREW_GITHUB_API_TOKEN="$(gh auth token)"
brew upgrade karlderkaefer/homebrew-tap/genderize
```

or use in-built autoupgade function

```bash
genderize upgrade
# Warning: karlderkaefer/tap/genderize 1.0.2 already installed
# Upgrade completed successfully.
```

## Usage

```bash
genderize version
# genderize 1.0.14, commit c49bf05bf183dbfbc7ed4ab319cabcf296d34b85, built at 2024-09-08T12:45:28Z
genderize peter  
# Checking URL https://api.genderize.io/?name=peter
# The predicted gender for the name 'peter' is 'male' with a probability of 1.00
```

## Deinstallation

```bash
brew remove genderize
brew untap karlderkaefer/homebrew-tap
```

## Steps to distribute releases from private repos

This sections provides a step-by-step guide how to distribute go releases from private repos. We goreleaser because it can manage the process of updating the brew formulars.

1. Create a private repo for homebrew formular. Popular name is `homebrew-tap`.
2. Add the [custom downloader script](https://github.com/karlderkaefer/homebrew-tap/blob/5af99309e2a83f3445068fd91e9ce7cead34f0d1/lib/custom_downloader.rb#L5) to your homebrew repo.
3. Create a go releaser config file [.goreleaser.yml](./.goreleaser.yml) in your go project.
4. When running `goreleaser release` the release will be created and uploaded to the github release page. Additionally the brew formular will be updated in the external repo.