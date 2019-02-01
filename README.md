[![Build Status](https://travis-ci.org/warrensbox/gitapp-installer.svg?branch=master)](https://travis-ci.org/warrensbox/gitapp-installer)
[![Go Report Card](https://goreportcard.com/badge/github.com/warrensbox/gitapp-installer)](https://goreportcard.com/report/github.com/warrensbox/gitapp-installer)
[![CircleCI](https://circleci.com/gh/warrensbox/gitapp-installer/tree/master.svg?style=shield&circle-token=d74b0de145c45b1d0da97f817363c77350e1a121)](https://circleci.com/gh/warrensbox/gitapp-installer)

# Terragrunt Switcher 

<img style="text-allign:center" src="https://s3.us-east-2.amazonaws.com/kepler-images/warrensbox/tgswitch/smallerlogo.png" alt="drawing" width="110" height="140"/>


The `tgswitch` command line tool lets you switch between different versions of [terragrunt](https://www.terragrunt.io/). 
If you do not have a particular version of terragrunt installed, `tgswitch` will download the version you desire.
The installation is minimal and easy. 
Once installed, simply select the version you require from the dropdown and start using terragrunt. 


See installation guide here: [tgswitch installation](https://warrensbox.github.io/gitapp-installer/)

## Installation

`tgswitch` is available for MacOS and Linux based operating systems.

### Homebrew

Installation for MacOS is the easiest with Homebrew. [If you do not have homebrew installed, click here](https://brew.sh/). 


```ruby
brew install warrensbox/tap/tgswitch
```

### Linux

Installation for other linux operation systems.

```sh
curl -L https://raw.githubusercontent.com/warrensbox/gitapp-installer/release/install.sh | bash
```

### Install from source

Alternatively, you can install the binary from source [here](https://github.com/warrensbox/gitapp-installer/releases) 

## How to use:
### Use dropdown menu to select version
<img src="https://s3.us-east-2.amazonaws.com/kepler-images/warrensbox/tgswitch/tgswitch.gif" alt="drawing" style="width: 180px;"/>

1.  You can switch between different versions of terragrunt by typing the command `tgswitch` on your terminal. 
2.  Select the version of terragrunt you require by using the up and down arrow.
3.  Hit **Enter** to select the desired version.

The most recently selected versions are presented at the top of the dropdown.

### Supply version on command line
<img src="https://s3.us-east-2.amazonaws.com/kepler-images/warrensbox/tgswitch/tgswitch-v4.gif" alt="drawing" style="width: 170px;"/>

1. You can also supply the desired version as an argument on the command line.
2. For example, `tgswitch 0.10.7` for version 0.10.7 of terragrunt.
3. Hit **Enter** to switch version.

## Additional Info

See how to *upgrade*, *uninstall*, *troubleshoot* here:[More info](https://warrensbox.github.io/gitapp-installer/additional)


## Issues

Please open  *issues* here:  [New Issue](https://github.com/warrensbox/gitapp-installer/issues)
