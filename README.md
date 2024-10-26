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
brew tap karlderkaefer/homebrew-tap/genderize
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
brew untap karlderkaefer/homebrew-tap/genderize
```
