load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "7b9bbe3ea1fccb46dcfa6c3f3e29ba7ec740d8733370e21cdc8937467b4a4349",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.22.4/rules_go-v0.22.4.tar.gz",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.22.4/rules_go-v0.22.4.tar.gz",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_rules_dependencies", "go_register_toolchains", "go_download_sdk")

go_download_sdk(
    name = "go_sdk",
    urls = ["http://mirrors.ustc.edu.cn/golang/{}"],
    sdks = {
        "linux_amd64": (
            "go1.14.2.linux-amd64.tar.gz",
            "6272d6e940ecb71ea5636ddb5fab3933e087c1356173c61f4a803895e947ebb3",
        ),
        "darwin_amd64": (
            "go1.14.2.darwin-amd64.tar.gz",
            "e0db81f890bb253552b3fd783fccbc2cdda02552295cb305e75984eef1c1e2b9",
        ),
        "windows_amd64": (
            "go1.14.2.windows-amd64.zip",
            "1b5a60b3bbaa81106d5ee03499b5734ec093c6a255abf9a6a067f0f497a57916",
        ),
    },
)

go_rules_dependencies()

go_register_toolchains()

http_archive(
    name = "bazel_gazelle",
    urls = [
        "https://storage.googleapis.com/bazel-mirror/github.com/bazelbuild/bazel-gazelle/releases/download/v0.20.0/bazel-gazelle-v0.20.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.20.0/bazel-gazelle-v0.20.0.tar.gz",
    ],
    sha256 = "d8c45ee70ec39a57e7a05e5027c32b1576cc7f16d9dd37135b0eddde45cf1b10",
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()