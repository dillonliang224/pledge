#!/bin/bash

export http_proxy="http://10.10.152.203:4411"
export https_proxy="http://10.10.152.203:4411"
export no_proxy=127.0.0.1,localhost

bazel run //:gazelle