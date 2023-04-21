load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

##############################################################################
# Versions
##############################################################################

BAZEL_GAZELLE_VERSION = "0.24.0"

BAZEL_SKYLIB_VERSION = "1.4.0"

BAZEL_VERSION = "6.1.1"

BUF_VERSION = "1.17.0"

DOTNET_VERSION = "7.0.200"

GAPIC_CSHARP_VERSION = "1.4.11"

GAPIC_GO_VERSION = "0.35.2"

GAPIC_JAVA_VERSION = "2.15.3"

GAPIC_PYTHON_VERSION = "1.9.1"

GAPIC_TYPESCRIPT_VERSION = "3.0.4"

GAX_DOTNET_VERSION = "4.3.1"

GO_VERSION = "1.20.3"

GOOGLE_APIS_VERSION = "b99ef53a44c7f1cd3e0e538119931a0b712f1ace"

GRPC_JAVA_VERSION = "1.53.0"

GRPC_VERSION = "1.47.0"

JAVA_VERSION = "openjdk-17.0.2"

PROTOC_GEN_VALIDATE_VERSION = "0.10.1"

PROTOC_VERSION = "3.21.12"

RULES_DOTNET_VERSION = "0.8.9"

RULES_GAPIC_VERSION = "0.23.1"

RULES_GO_VERSION = "0.38.1"

RULES_JVM_EXTERNAL_VERSION = "4.5"

RULES_PKG_VERSION = "0.7.0"

RULES_PYTHON_VERSION = "0.9.0"

SHELLCHECK_VERSION = "0.7.2"

SHFMT_VERSION = "3.2.4"

load("//buildtools:asdf.bzl", "gen_tool_versions")

gen_tool_versions(
    name = "asdf",
    versions = {
        "bazel": BAZEL_VERSION,
        "buf": BUF_VERSION,
        "dotnet-core": DOTNET_VERSION,
        "golang": GO_VERSION,
        "java": JAVA_VERSION,
        "protoc": PROTOC_VERSION,
        "shellcheck": SHELLCHECK_VERSION,
        "shfmt": SHFMT_VERSION,
    },
)
##############################################################################
# Common
##############################################################################

http_archive(
    name = "go_googleapis",
    sha256 = "ed06aff59a0fcd2b36aeca12934896307e13bcac23e9f15aaf73187fc78555c0",
    strip_prefix = "googleapis-{0}".format(GOOGLE_APIS_VERSION),
    url = "https://github.com/googleapis/googleapis/archive/{0}.zip".format(GOOGLE_APIS_VERSION),
)

load("@go_googleapis//:repository_rules.bzl", "switched_rules_by_language")

switched_rules_by_language(
    name = "com_google_googleapis_imports",
    cc = True,
    csharp = True,
    gapic = True,
    go = True,
    grpc = True,
    java = True,
    nodejs = True,
    php = False,
    python = True,
    ruby = False,
)

_bazel_skylib_sha256 = "f24ab666394232f834f74d19e2ff142b0af17466ea0c69a3f4c276ee75f6efce"

http_archive(
    name = "bazel_skylib",
    sha256 = _bazel_skylib_sha256,
    urls = ["https://github.com/bazelbuild/bazel-skylib/releases/download/{0}/bazel-skylib-{0}.tar.gz".format(BAZEL_SKYLIB_VERSION)],
)

# Protobuf depends on very old version of rules_jvm_external.
# Importing older version of rules_jvm_external first (this is one of the things that protobuf_deps() call
# below will do) breaks the Java client library generation process, so importing the proper version explicitly before calling protobuf_deps().

RULES_JVM_EXTERNAL_SHA = "b17d7388feb9bfa7f2fa09031b32707df529f26c91ab9e5d909eb1676badd9a6"

http_archive(
    name = "rules_jvm_external",
    sha256 = RULES_JVM_EXTERNAL_SHA,
    strip_prefix = "rules_jvm_external-{0}".format(RULES_JVM_EXTERNAL_VERSION),
    url = "https://github.com/bazelbuild/rules_jvm_external/archive/{0}.zip".format(RULES_JVM_EXTERNAL_VERSION),
)

load("@rules_jvm_external//:repositories.bzl", "rules_jvm_external_deps")

rules_jvm_external_deps()

load("@rules_jvm_external//:setup.bzl", "rules_jvm_external_setup")

rules_jvm_external_setup()

# Python rules should go early in the dependencies list, otherwise a wrong
# version of the library will be selected as a transitive dependency of gRPC.
http_archive(
    name = "rules_python",
    sha256 = "5fa3c738d33acca3b97622a13a741129f67ef43f5fdfcec63b29374cc0574c29",
    strip_prefix = "rules_python-{0}".format(RULES_PYTHON_VERSION),
    url = "https://github.com/bazelbuild/rules_python/archive/{0}.tar.gz".format(RULES_PYTHON_VERSION),
)

http_archive(
    name = "rules_pkg",
    sha256 = "8a298e832762eda1830597d64fe7db58178aa84cd5926d76d5b744d6558941c2",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_pkg/releases/download/{0}/rules_pkg-{0}.tar.gz".format(RULES_PKG_VERSION),
        "https://github.com/bazelbuild/rules_pkg/releases/download/{0}/rules_pkg-{0}.tar.gz".format(RULES_PKG_VERSION),
    ],
)

load("@rules_pkg//:deps.bzl", "rules_pkg_dependencies")

rules_pkg_dependencies()

http_archive(
    name = "com_google_protobuf",
    sha256 = "930c2c3b5ecc6c9c12615cf5ad93f1cd6e12d0aba862b572e076259970ac3a53",
    strip_prefix = "protobuf-{0}".format(PROTOC_VERSION),
    urls = ["https://github.com/protocolbuffers/protobuf/archive/v{0}.tar.gz".format(PROTOC_VERSION)],
)

load("@com_google_protobuf//:protobuf_deps.bzl", "PROTOBUF_MAVEN_ARTIFACTS", "protobuf_deps")

protobuf_deps()

http_archive(
    name = "rules_proto",
    sha256 = "602e7161d9195e50246177e7c55b2f39950a9cf7366f74ed5f22fd45750cd208",
    strip_prefix = "rules_proto-97d8af4dc474595af3900dd85cb3a29ad28cc313",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_proto/archive/97d8af4dc474595af3900dd85cb3a29ad28cc313.tar.gz",
        "https://github.com/bazelbuild/rules_proto/archive/97d8af4dc474595af3900dd85cb3a29ad28cc313.tar.gz",
    ],
)

load("@rules_proto//proto:repositories.bzl", "rules_proto_dependencies", "rules_proto_toolchains")

rules_proto_dependencies()

rules_proto_toolchains()

##############################################################################
# Go
##############################################################################

# rules_gapic also depends on rules_go, so it must come after our own dependency on rules_go.
# It must also come before gapic-generator-go so as to ensure that it does not bring in an old
# version of rules_gapic.

_rules_gapic_sha256 = "cda71a5e50daa31bdf7c1bbc9196cea21adb3daea97e2a28dc9569f03c2a4f52"

http_archive(
    name = "rules_gapic",
    sha256 = _rules_gapic_sha256,
    strip_prefix = "rules_gapic-{0}".format(RULES_GAPIC_VERSION),
    urls = ["https://github.com/googleapis/rules_gapic/archive/v{0}.tar.gz".format(RULES_GAPIC_VERSION)],
)

# This must be above the download of gRPC (in C++ section) and
# rules_gapic_repositories because both depend on rules_go and we need to manage
# our version of rules_go explicitly rather than depend on the version those
# bring in transitively.

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "dd926a88a564a9246713a9c00b35315f54cbd46b31a26d5d8fb264c07045f05d",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v{0}/rules_go-v{0}.zip".format(RULES_GO_VERSION),
        "https://github.com/bazelbuild/rules_go/releases/download/v{0}/rules_go-v{0}.zip".format(RULES_GO_VERSION),
    ],
)

# Gazelle dependency version should match gazelle dependency expected by gRPC

http_archive(
    name = "bazel_gazelle",
    sha256 = "de69a09dc70417580aabf20a28619bb3ef60d038470c7cf8442fafcf627c21cb",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v{0}/bazel-gazelle-v{0}.tar.gz".format(BAZEL_GAZELLE_VERSION),
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v{0}/bazel-gazelle-v{0}.tar.gz".format(BAZEL_GAZELLE_VERSION),
    ],
)

# Until this project is migrated to consume the new subdirectory of generated
# types e.g. longrunningpb, we must define our own version of longrunning here.
# @unused
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

http_archive(
    name = "com_googleapis_gapic_generator_go",
    sha256 = "d9fa55ef3bc14e1c3c15870ef1080a29e6af2b996e37a1d7043f35a61aa1e869",
    strip_prefix = "gapic-generator-go-{0}".format(GAPIC_GO_VERSION),
    urls = ["https://github.com/googleapis/gapic-generator-go/archive/v{0}.tar.gz".format(GAPIC_GO_VERSION)],
)

load("@com_googleapis_gapic_generator_go//:repositories.bzl", "com_googleapis_gapic_generator_go_repositories")
load("//:repositories.bzl", "go_dependencies")

# gazelle:repository_macro repositories.bzl%go_dependencies
go_dependencies()

com_googleapis_gapic_generator_go_repositories()

http_archive(
    name = "com_envoyproxy_protoc_gen_validate",
    sha256 = "884f7166893d4869d9e86c171777c11e51b138a6ec170e1d8eba8f091a9ef85a",
    strip_prefix = "protoc-gen-validate-{0}".format(PROTOC_GEN_VALIDATE_VERSION),
    urls = [
        "https://github.com/bufbuild/protoc-gen-validate/archive/refs/tags/v{0}.tar.gz".format(PROTOC_GEN_VALIDATE_VERSION),
    ],
)

# rules_go and gazelle dependencies are loaded after gapic-generator-go
# dependencies to ensure that they do not override any of the go_repository
# dependencies of gapic-generator-go.
load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_register_toolchains(version = GO_VERSION)

go_rules_dependencies()

gazelle_dependencies()

load("@rules_gapic//:repositories.bzl", "rules_gapic_repositories")

rules_gapic_repositories()

##############################################################################
# C++
##############################################################################
# C++ must go before everything else, since the key dependencies (protobuf and gRPC)
# are C++ repositories and they have the highest chance to have the correct versions of the
# transitive dependencies (for those dependencies which are shared by many other repositories
# imported in this workspace).
#
# Note, even though protobuf and gRPC are mostly written in C++, they are used to generate stuff
# for most of the other languages as well, so they can be considered as the core cross-language
# dependencies.

# Import boringssl explicitly to override what gRPC imports as its dependency.
# Boringssl build fails on gcc12 without this fix:
# https://github.com/google/boringssl/commit/8462a367bb57e9524c3d8eca9c62733c63a63cf4,
# which is present only in the newest version of boringssl, not the one imported
# by gRPC. Remove this import once gRPC depends on a newer version.
http_archive(
    name = "boringssl",
    sha256 = "b460f8673f3393e58ce506e9cdde7f2c3b2575b075f214cb819fb57d809f052b",
    strip_prefix = "boringssl-bb41bc007079982da419c0ec3186e510cbcf09d0",
    urls = [
        "https://github.com/google/boringssl/archive/bb41bc007079982da419c0ec3186e510cbcf09d0.zip",
    ],
)

_grpc_sha256 = "edf25f4db6c841853b7a29d61b0980b516dc31a1b6cdc399bcf24c1446a4a249"

http_archive(
    name = "com_github_grpc_grpc",
    sha256 = _grpc_sha256,
    strip_prefix = "grpc-{0}".format(GRPC_VERSION),
    urls = ["https://github.com/grpc/grpc/archive/v{0}.zip".format(GRPC_VERSION)],
)

load("@com_github_grpc_grpc//bazel:grpc_deps.bzl", "grpc_deps")

grpc_deps()

# gRPC enforces a specific version of Go toolchain which conflicts with our build.
# All the relevant parts of grpc_extra_deps() are imported in this  WORKSPACE file
# explicitly, that is why we do not call grpc_extra_deps() here.

load("@build_bazel_rules_apple//apple:repositories.bzl", "apple_rules_dependencies")

apple_rules_dependencies()

load("@build_bazel_apple_support//lib:repositories.bzl", "apple_support_dependencies")

apple_support_dependencies()

# Starting in protobuf 3.19, protobuf project started to provide
# PROTOBUF_MAVEN_ARTIFACTS variable so that Bazel users can resolve their
# dependencies through maven_install.
# https://github.com/protocolbuffers/protobuf/issues/9132

load("@rules_jvm_external//:defs.bzl", "maven_install")

maven_install(
    artifacts = PROTOBUF_MAVEN_ARTIFACTS,
    generate_compat_repositories = True,
    repositories = [
        "https://repo.maven.apache.org/maven2/",
    ],
)

maven_install(
    artifacts = [
        "com.google.api:gapic-generator-java:{0}".format(GAPIC_JAVA_VERSION),
    ],
    #Update this False for local development
    fail_on_missing_checksum = True,
    repositories = [
        "m2Local",
        "https://repo.maven.apache.org/maven2/",
    ],
)

http_archive(
    name = "gapic_generator_java",
    sha256 = "752a930d4d0f6c287265eaa513ff8341b5fbb2aa60ec90a3d3e92934225a79d2",
    strip_prefix = "gapic-generator-java-{0}".format(GAPIC_JAVA_VERSION),
    urls = ["https://github.com/googleapis/gapic-generator-java/archive/v{0}.zip".format(GAPIC_JAVA_VERSION)],
)

http_archive(
    name = "io_grpc_grpc_java",
    sha256 = "fd0a649d03a8da06746814f414fb4d36c1b2f34af2aad4e19ae43f7c4bd6f15e",
    strip_prefix = "grpc-java-{0}".format(GRPC_JAVA_VERSION),
    urls = ["https://github.com/grpc/grpc-java/archive/refs/tags/v{0}.tar.gz".format(GRPC_JAVA_VERSION)],
)

# gax-java is part of gapic-generator-java repository
http_archive(
    name = "com_google_api_gax_java",
    sha256 = "752a930d4d0f6c287265eaa513ff8341b5fbb2aa60ec90a3d3e92934225a79d2",
    strip_prefix = "gapic-generator-java-{0}/gax-java".format(GAPIC_JAVA_VERSION),
    urls = ["https://github.com/googleapis/gapic-generator-java/archive/v{0}.zip".format(GAPIC_JAVA_VERSION)],
)

load("@com_google_api_gax_java//:repository_rules.bzl", "com_google_api_gax_java_properties")

com_google_api_gax_java_properties(
    name = "com_google_api_gax_java_properties",
    file = "@com_google_api_gax_java//:dependencies.properties",
)

load("@com_google_api_gax_java//:repositories.bzl", "com_google_api_gax_java_repositories")

com_google_api_gax_java_repositories()

load("@io_grpc_grpc_java//:repositories.bzl", "grpc_java_repositories")

grpc_java_repositories()

##############################################################################
# Python
##############################################################################
load("@rules_gapic//python:py_gapic_repositories.bzl", "py_gapic_repositories")

py_gapic_repositories()

load("@rules_python//python:pip.bzl", "pip_install")

pip_install()

http_archive(
    name = "gapic_generator_python",
    sha256 = "a9fa11a9fe9783b07fa1affc9b3a3514ccf51029244c2814266dd777ea04ca75",
    strip_prefix = "gapic-generator-python-{0}".format(GAPIC_PYTHON_VERSION),
    urls = ["https://github.com/googleapis/gapic-generator-python/archive/v{0}.zip".format(GAPIC_PYTHON_VERSION)],
)

load(
    "@gapic_generator_python//:repositories.bzl",
    "gapic_generator_python",
    "gapic_generator_register_toolchains",
)

gapic_generator_python()

gapic_generator_register_toolchains()

##############################################################################
# TypeScript
##############################################################################

_gapic_generator_typescript_sha256 = "24e7d2e36930f31825c74b4d29d58a80b7e292f372e5789cdcce3887de222793"

### TypeScript generator
http_archive(
    name = "gapic_generator_typescript",
    sha256 = _gapic_generator_typescript_sha256,
    strip_prefix = "gapic-generator-typescript-{0}".format(GAPIC_TYPESCRIPT_VERSION),
    urls = ["https://github.com/googleapis/gapic-generator-typescript/archive/v{0}.tar.gz".format(GAPIC_TYPESCRIPT_VERSION)],
)

load("@gapic_generator_typescript//:repositories.bzl", "NODE_VERSION", "gapic_generator_typescript_repositories")

gapic_generator_typescript_repositories()

load("@aspect_rules_js//js:repositories.bzl", "rules_js_dependencies")

rules_js_dependencies()

load("@aspect_rules_ts//ts:repositories.bzl", "rules_ts_dependencies")

rules_ts_dependencies(
    ts_version_from = "@gapic_generator_typescript//:package.json",
)

load("@rules_nodejs//nodejs:repositories.bzl", "nodejs_register_toolchains")

nodejs_register_toolchains(
    name = "nodejs",
    node_version = NODE_VERSION,
)

load("@aspect_rules_js//npm:npm_import.bzl", "npm_translate_lock", "pnpm_repository")

npm_translate_lock(
    name = "npm",
    data = ["@gapic_generator_typescript//:package.json"],
    pnpm_lock = "@gapic_generator_typescript//:pnpm-lock.yaml",
)

load("@npm//:repositories.bzl", "npm_repositories")

npm_repositories()

pnpm_repository(name = "pnpm")

##############################################################################
# C#
##############################################################################

http_archive(
    name = "rules_dotnet",
    sha256 = "2650540b29ef1b31b665305bb13497d25d5f565bde459b3e614474177783c7e0",
    strip_prefix = "rules_dotnet-{0}".format(RULES_DOTNET_VERSION),
    url = "https://github.com/bazelbuild/rules_dotnet/releases/download/v{0}/rules_dotnet-v{0}.tar.gz".format(RULES_DOTNET_VERSION),
)

load("@rules_dotnet//dotnet:repositories.bzl", "dotnet_register_toolchains", "rules_dotnet_dependencies")
load("@rules_dotnet//dotnet:rules_dotnet_nuget_packages.bzl", "rules_dotnet_nuget_packages")
load("@rules_dotnet//dotnet:paket2bazel_dependencies.bzl", "paket2bazel_dependencies")

# Required to access the C#-specific common resources config.

_gax_dotnet_sha256 = "f3684a6c352012b511b2f49707788a78a31f601ea10447d21ef225874f7f4d23"

http_archive(
    name = "gax_dotnet",
    build_file_content = "exports_files([\"Google.Api.Gax/ResourceNames/CommonResourcesConfig.json\"])",
    sha256 = _gax_dotnet_sha256,
    strip_prefix = "gax-dotnet-Google.Api.Gax-{0}".format(GAX_DOTNET_VERSION),
    urls = ["https://github.com/googleapis/gax-dotnet/archive/refs/tags/Google.Api.Gax-{0}.tar.gz".format(GAX_DOTNET_VERSION)],
)

rules_dotnet_dependencies()

# Here you can specify the version of the .NET SDK to use.
dotnet_register_toolchains("dotnet", DOTNET_VERSION)

rules_dotnet_nuget_packages()

paket2bazel_dependencies()

load("//:paket.bzl", "paket")

paket()

_gapic_generator_csharp_sha256 = "40bb2ecf1e540df8f58bdca15c48e3da6fbdddc9c5786421b858222fb4e25202"

http_archive(
    name = "gapic_generator_csharp",
    sha256 = _gapic_generator_csharp_sha256,
    strip_prefix = "gapic-generator-csharp-{0}".format(GAPIC_CSHARP_VERSION),
    urls = ["https://github.com/googleapis/gapic-generator-csharp/archive/refs/tags/v{0}.tar.gz".format(GAPIC_CSHARP_VERSION)],
)

load("@gapic_generator_csharp//:repositories.bzl", "gapic_generator_csharp_repositories")

gapic_generator_csharp_repositories()
