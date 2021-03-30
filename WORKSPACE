load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository", "new_git_repository")
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

git_repository(
    name = "io_bazel_rules_go",
    #tag = "v0.26.0",
    commit = "ddc5747aa3a52e7703f48281d82ea06985d95628",
    remote = "git@repo.jazznetworks.com:vendor/github.com_bazelbuild_rules_go.git",
    shallow_since = "1606940837 -0500",
)

# git_repository(
#     name = "io_bazel_rules_go",
#     #tag = "v0.25.0",
#     commit = "8c8191291ccc7e25a73bf93ae127da7e9ed656db",
#     remote = "git@repo.jazznetworks.com:vendor/github.com_bazelbuild_rules_go.git",
#     shallow_since = "1606940837 -0500",
# )

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains(version = "1.16.1")

http_archive(
    name = "shared_library_external",
    build_file = "shared_library_external.BUILD",
    sha256 = "1766e8f94e08fd073c125f4076f5c939f3c04f4aaddd9ac4030f062809bc0657",
    url = "https://github.com/jfpowell/bazel_linking/raw/main/shared_library.zip"
)