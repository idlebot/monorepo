load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

##############################################################################
# Common
##############################################################################

http_archive(
    name = "go_googleapis",
    sha256 = "ed06aff59a0fcd2b36aeca12934896307e13bcac23e9f15aaf73187fc78555c0",
    strip_prefix = "googleapis-b99ef53a44c7f1cd3e0e538119931a0b712f1ace",
    url = "https://github.com/googleapis/googleapis/archive/b99ef53a44c7f1cd3e0e538119931a0b712f1ace.zip",
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

_bazel_skylib_version = "1.4.0"

_bazel_skylib_sha256 = "f24ab666394232f834f74d19e2ff142b0af17466ea0c69a3f4c276ee75f6efce"

http_archive(
    name = "bazel_skylib",
    sha256 = _bazel_skylib_sha256,
    urls = ["https://github.com/bazelbuild/bazel-skylib/releases/download/{0}/bazel-skylib-{0}.tar.gz".format(_bazel_skylib_version)],
)

# Protobuf depends on very old version of rules_jvm_external.
# Importing older version of rules_jvm_external first (this is one of the things that protobuf_deps() call
# below will do) breaks the Java client library generation process, so importing the proper version explicitly before calling protobuf_deps().
RULES_JVM_EXTERNAL_TAG = "4.5"

RULES_JVM_EXTERNAL_SHA = "b17d7388feb9bfa7f2fa09031b32707df529f26c91ab9e5d909eb1676badd9a6"

http_archive(
    name = "rules_jvm_external",
    sha256 = RULES_JVM_EXTERNAL_SHA,
    strip_prefix = "rules_jvm_external-%s" % RULES_JVM_EXTERNAL_TAG,
    url = "https://github.com/bazelbuild/rules_jvm_external/archive/%s.zip" % RULES_JVM_EXTERNAL_TAG,
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
    strip_prefix = "rules_python-0.9.0",
    url = "https://github.com/bazelbuild/rules_python/archive/0.9.0.tar.gz",
)

http_archive(
    name = "rules_pkg",
    sha256 = "8a298e832762eda1830597d64fe7db58178aa84cd5926d76d5b744d6558941c2",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_pkg/releases/download/0.7.0/rules_pkg-0.7.0.tar.gz",
        "https://github.com/bazelbuild/rules_pkg/releases/download/0.7.0/rules_pkg-0.7.0.tar.gz",
    ],
)

load("@rules_pkg//:deps.bzl", "rules_pkg_dependencies")

rules_pkg_dependencies()

http_archive(
    name = "com_google_protobuf",
    sha256 = "930c2c3b5ecc6c9c12615cf5ad93f1cd6e12d0aba862b572e076259970ac3a53",
    strip_prefix = "protobuf-3.21.12",
    urls = ["https://github.com/protocolbuffers/protobuf/archive/v3.21.12.tar.gz"],
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
_rules_gapic_version = "0.23.1"

_rules_gapic_sha256 = "cda71a5e50daa31bdf7c1bbc9196cea21adb3daea97e2a28dc9569f03c2a4f52"

http_archive(
    name = "rules_gapic",
    sha256 = _rules_gapic_sha256,
    strip_prefix = "rules_gapic-%s" % _rules_gapic_version,
    urls = ["https://github.com/googleapis/rules_gapic/archive/v%s.tar.gz" % _rules_gapic_version],
)

# This must be above the download of gRPC (in C++ section) and
# rules_gapic_repositories because both depend on rules_go and we need to manage
# our version of rules_go explicitly rather than depend on the version those
# bring in transitively.
_io_bazel_rules_go_version = "0.38.1"

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "dd926a88a564a9246713a9c00b35315f54cbd46b31a26d5d8fb264c07045f05d",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v{0}/rules_go-v{0}.zip".format(_io_bazel_rules_go_version),
        "https://github.com/bazelbuild/rules_go/releases/download/v{0}/rules_go-v{0}.zip".format(_io_bazel_rules_go_version),
    ],
)

# Gazelle dependency version should match gazelle dependency expected by gRPC
_bazel_gazelle_version = "0.24.0"

http_archive(
    name = "bazel_gazelle",
    sha256 = "de69a09dc70417580aabf20a28619bb3ef60d038470c7cf8442fafcf627c21cb",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v{0}/bazel-gazelle-v{0}.tar.gz".format(_bazel_gazelle_version),
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v{0}/bazel-gazelle-v{0}.tar.gz".format(_bazel_gazelle_version),
    ],
)

# Until this project is migrated to consume the new subdirectory of generated
# types e.g. longrunningpb, we must define our own version of longrunning here.
# @unused
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

_gapic_generator_go_version = "0.35.2"

http_archive(
    name = "com_googleapis_gapic_generator_go",
    sha256 = "d9fa55ef3bc14e1c3c15870ef1080a29e6af2b996e37a1d7043f35a61aa1e869",
    strip_prefix = "gapic-generator-go-%s" % _gapic_generator_go_version,
    urls = ["https://github.com/googleapis/gapic-generator-go/archive/v%s.tar.gz" % _gapic_generator_go_version],
)

load("@com_googleapis_gapic_generator_go//:repositories.bzl", "com_googleapis_gapic_generator_go_repositories")
load("//:repositories.bzl", "go_dependencies")

# gazelle:repository_macro repositories.bzl%go_dependencies
go_dependencies()

com_googleapis_gapic_generator_go_repositories()

http_archive(
    name = "com_envoyproxy_protoc_gen_validate",
    sha256 = "884f7166893d4869d9e86c171777c11e51b138a6ec170e1d8eba8f091a9ef85a",
    strip_prefix = "protoc-gen-validate-0.10.1",
    urls = [
        "https://github.com/bufbuild/protoc-gen-validate/archive/refs/tags/v0.10.1.tar.gz",
    ],
)

# rules_go and gazelle dependencies are loaded after gapic-generator-go
# dependencies to ensure that they do not override any of the go_repository
# dependencies of gapic-generator-go.
load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_register_toolchains(version = "1.20.2")

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

_grpc_version = "1.47.0"

_grpc_sha256 = "edf25f4db6c841853b7a29d61b0980b516dc31a1b6cdc399bcf24c1446a4a249"

http_archive(
    name = "com_github_grpc_grpc",
    sha256 = _grpc_sha256,
    strip_prefix = "grpc-%s" % _grpc_version,
    urls = ["https://github.com/grpc/grpc/archive/v%s.zip" % _grpc_version],
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

_gapic_generator_java_version = "2.15.3"

maven_install(
    artifacts = [
        "com.google.api:gapic-generator-java:" + _gapic_generator_java_version,
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
    strip_prefix = "gapic-generator-java-%s" % _gapic_generator_java_version,
    urls = ["https://github.com/googleapis/gapic-generator-java/archive/v%s.zip" % _gapic_generator_java_version],
)

_io_grpc_grpc_java_version = "1.53.0"

http_archive(
    name = "io_grpc_grpc_java",
    sha256 = "fd0a649d03a8da06746814f414fb4d36c1b2f34af2aad4e19ae43f7c4bd6f15e",
    strip_prefix = "grpc-java-%s" % _io_grpc_grpc_java_version,
    urls = ["https://github.com/grpc/grpc-java/archive/refs/tags/v%s.tar.gz" % _io_grpc_grpc_java_version],
)

# gax-java is part of gapic-generator-java repository
http_archive(
    name = "com_google_api_gax_java",
    sha256 = "752a930d4d0f6c287265eaa513ff8341b5fbb2aa60ec90a3d3e92934225a79d2",
    strip_prefix = "gapic-generator-java-%s/gax-java" % _gapic_generator_java_version,
    urls = ["https://github.com/googleapis/gapic-generator-java/archive/v%s.zip" % _gapic_generator_java_version],
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

_gapic_generator_python_version = "1.9.1"

http_archive(
    name = "gapic_generator_python",
    sha256 = "a9fa11a9fe9783b07fa1affc9b3a3514ccf51029244c2814266dd777ea04ca75",
    strip_prefix = "gapic-generator-python-%s" % _gapic_generator_python_version,
    urls = ["https://github.com/googleapis/gapic-generator-python/archive/v%s.zip" % _gapic_generator_python_version],
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

_gapic_generator_typescript_version = "3.0.4"

_gapic_generator_typescript_sha256 = "24e7d2e36930f31825c74b4d29d58a80b7e292f372e5789cdcce3887de222793"

### TypeScript generator
http_archive(
    name = "gapic_generator_typescript",
    sha256 = _gapic_generator_typescript_sha256,
    strip_prefix = "gapic-generator-typescript-%s" % _gapic_generator_typescript_version,
    urls = ["https://github.com/googleapis/gapic-generator-typescript/archive/v%s.tar.gz" % _gapic_generator_typescript_version],
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
    strip_prefix = "rules_dotnet-0.8.9",
    url = "https://github.com/bazelbuild/rules_dotnet/releases/download/v0.8.9/rules_dotnet-v0.8.9.tar.gz",
)

load(
    "@rules_dotnet//dotnet:repositories.bzl",
    "dotnet_register_toolchains",
    "rules_dotnet_dependencies",
)

rules_dotnet_dependencies()

# Here you can specify the version of the .NET SDK to use.
dotnet_register_toolchains("dotnet", "7.0.101")

load("@rules_dotnet//dotnet:rules_dotnet_nuget_packages.bzl", "rules_dotnet_nuget_packages")

rules_dotnet_nuget_packages()

# Required to access the C#-specific common resources config.
_gax_dotnet_version = "Google.Api.Gax-4.3.1"

_gax_dotnet_sha256 = "f3684a6c352012b511b2f49707788a78a31f601ea10447d21ef225874f7f4d23"

http_archive(
    name = "gax_dotnet",
    build_file_content = "exports_files([\"Google.Api.Gax/ResourceNames/CommonResourcesConfig.json\"])",
    sha256 = _gax_dotnet_sha256,
    strip_prefix = "gax-dotnet-%s" % _gax_dotnet_version,
    urls = ["https://github.com/googleapis/gax-dotnet/archive/refs/tags/%s.tar.gz" % _gax_dotnet_version],
)

_gapic_generator_csharp_version = "1.4.11"

_gapic_generator_csharp_sha256 = "40bb2ecf1e540df8f58bdca15c48e3da6fbdddc9c5786421b858222fb4e25202"

http_archive(
    name = "gapic_generator_csharp",
    sha256 = _gapic_generator_csharp_sha256,
    strip_prefix = "gapic-generator-csharp-%s" % _gapic_generator_csharp_version,
    urls = ["https://github.com/googleapis/gapic-generator-csharp/archive/refs/tags/v%s.tar.gz" % _gapic_generator_csharp_version],
)

load("@gapic_generator_csharp//:repositories.bzl", "gapic_generator_csharp_repositories")

gapic_generator_csharp_repositories()

load(
    "@rules_dotnet//dotnet:defs.bzl",
    "nuget_repo",
)

nuget_repo(
    name = "dotnet_deps",
    packages = [
        (
            "Expecto",
            "9.0.4",
            "sha512-k0TT6pNIyzDaJD0ZxHDhNU0UmmWZlum2XFfHTGrkApQ+JUdjcoBqKOACXrSkfiLVYsD8Ww768eeAiKPP3QYetw==",
            [
                "FSharp.Core",
                "Mono.Cecil",
            ],
            [],
        ),
        ("FSharp.Core", "6.0.3", "sha512-aDyKHiVFMwXWJrfW90iAeKyvw/lN+x98DPfx4oXke9Qnl4dz1sOi8KT2iczGeunqyWXh7nm+XUJ18i/0P3pZYw==", [], []),
        (
            "FSharp.Data",
            "5.0.2",
            "sha512-BlDokqEWMysUMedhZzaREUPrhAbj8VRUEXjUrd85fzH63XaxppqjEYtpjQLnQcwkyWI71bzr3cfzYgaAANQLAQ==",
            ["FSharp.Core"],
            [],
        ),
        (
            "Microsoft.AspNetCore.App.Ref",
            "6.0.8",
            "sha512-yLy7tFshfGLJRCFdlmOv8YOlJ4J5IfE88bnqiulxsJzhgEQNfbPQLpxbvmjCO3Zg0tdBLAS4B8QYWoojkOkWLg==",
            [],
            [
                "Microsoft.Extensions.Caching.Abstractions|6.0.0",
                "Microsoft.Extensions.Caching.Memory|6.0.0",
                "Microsoft.Extensions.Configuration.Abstractions|6.0.0",
                "Microsoft.Extensions.Configuration.Binder|6.0.0",
                "Microsoft.Extensions.Configuration.CommandLine|6.0.0",
                "Microsoft.Extensions.Configuration|6.0.0",
                "Microsoft.Extensions.Configuration.EnvironmentVariables|6.0.0",
                "Microsoft.Extensions.Configuration.FileExtensions|6.0.0",
                "Microsoft.Extensions.Configuration.Ini|6.0.0",
                "Microsoft.Extensions.Configuration.Json|6.0.0",
                "Microsoft.Extensions.Configuration.UserSecrets|6.0.0",
                "Microsoft.Extensions.Configuration.Xml|6.0.0",
                "Microsoft.Extensions.DependencyInjection.Abstractions|6.0.0",
                "Microsoft.Extensions.DependencyInjection|6.0.0",
                "Microsoft.Extensions.FileProviders.Abstractions|6.0.0",
                "Microsoft.Extensions.FileProviders.Composite|6.0.0",
                "Microsoft.Extensions.FileProviders.Physical|6.0.0",
                "Microsoft.Extensions.FileSystemGlobbing|6.0.0",
                "Microsoft.Extensions.Hosting.Abstractions|6.0.0",
                "Microsoft.Extensions.Hosting|6.0.0",
                "Microsoft.Extensions.Http|6.0.0",
                "Microsoft.Extensions.Logging.Abstractions|6.0.0",
                "Microsoft.Extensions.Logging.Configuration|6.0.0",
                "Microsoft.Extensions.Logging.Console|6.0.0",
                "Microsoft.Extensions.Logging.Debug|6.0.0",
                "Microsoft.Extensions.Logging|6.0.0",
                "Microsoft.Extensions.Logging.EventLog|6.0.0",
                "Microsoft.Extensions.Logging.EventSource|6.0.0",
                "Microsoft.Extensions.Logging.TraceSource|6.0.0",
                "Microsoft.Extensions.Options.ConfigurationExtensions|6.0.0",
                "Microsoft.Extensions.Options.DataAnnotations|6.0.0",
                "Microsoft.Extensions.Options|6.0.0",
                "Microsoft.Extensions.Primitives|6.0.0",
                "System.Diagnostics.EventLog|6.0.0",
                "System.IO.Pipelines|6.0.0",
                "System.Security.Cryptography.Xml|6.0.0",
                "Microsoft.AspNetCore.Antiforgery|6.0.0",
                "Microsoft.AspNetCore.Authentication.Abstractions|6.0.0",
                "Microsoft.AspNetCore.Authentication.Cookies|6.0.0",
                "Microsoft.AspNetCore.Authentication.Core|6.0.0",
                "Microsoft.AspNetCore.Authentication|6.0.0",
                "Microsoft.AspNetCore.Authentication.OAuth|6.0.0",
                "Microsoft.AspNetCore.Authorization|6.0.0",
                "Microsoft.AspNetCore.Authorization.Policy|6.0.0",
                "Microsoft.AspNetCore.Components.Authorization|6.0.0",
                "Microsoft.AspNetCore.Components|6.0.0",
                "Microsoft.AspNetCore.Components.Forms|6.0.0",
                "Microsoft.AspNetCore.Components.Server|6.0.0",
                "Microsoft.AspNetCore.Components.Web|6.0.0",
                "Microsoft.AspNetCore.Connections.Abstractions|6.0.0",
                "Microsoft.AspNetCore.CookiePolicy|6.0.0",
                "Microsoft.AspNetCore.Cors|6.0.0",
                "Microsoft.AspNetCore.Cryptography.Internal|6.0.0",
                "Microsoft.AspNetCore.Cryptography.KeyDerivation|6.0.0",
                "Microsoft.AspNetCore.DataProtection.Abstractions|6.0.0",
                "Microsoft.AspNetCore.DataProtection|6.0.0",
                "Microsoft.AspNetCore.DataProtection.Extensions|6.0.0",
                "Microsoft.AspNetCore.Diagnostics.Abstractions|6.0.0",
                "Microsoft.AspNetCore.Diagnostics|6.0.0",
                "Microsoft.AspNetCore.Diagnostics.HealthChecks|6.0.0",
                "Microsoft.AspNetCore|6.0.0",
                "Microsoft.AspNetCore.HostFiltering|6.0.0",
                "Microsoft.AspNetCore.Hosting.Abstractions|6.0.0",
                "Microsoft.AspNetCore.Hosting|6.0.0",
                "Microsoft.AspNetCore.Hosting.Server.Abstractions|6.0.0",
                "Microsoft.AspNetCore.Html.Abstractions|6.0.0",
                "Microsoft.AspNetCore.Http.Abstractions|6.0.0",
                "Microsoft.AspNetCore.Http.Connections.Common|6.0.0",
                "Microsoft.AspNetCore.Http.Connections|6.0.0",
                "Microsoft.AspNetCore.Http|6.0.0",
                "Microsoft.AspNetCore.Http.Extensions|6.0.0",
                "Microsoft.AspNetCore.Http.Features|6.0.0",
                "Microsoft.AspNetCore.Http.Results|6.0.0",
                "Microsoft.AspNetCore.HttpLogging|6.0.0",
                "Microsoft.AspNetCore.HttpOverrides|6.0.0",
                "Microsoft.AspNetCore.HttpsPolicy|6.0.0",
                "Microsoft.AspNetCore.Identity|6.0.0",
                "Microsoft.AspNetCore.Localization|6.0.0",
                "Microsoft.AspNetCore.Localization.Routing|6.0.0",
                "Microsoft.AspNetCore.Metadata|6.0.0",
                "Microsoft.AspNetCore.Mvc.Abstractions|6.0.0",
                "Microsoft.AspNetCore.Mvc.ApiExplorer|6.0.0",
                "Microsoft.AspNetCore.Mvc.Core|6.0.0",
                "Microsoft.AspNetCore.Mvc.Cors|6.0.0",
                "Microsoft.AspNetCore.Mvc.DataAnnotations|6.0.0",
                "Microsoft.AspNetCore.Mvc|6.0.0",
                "Microsoft.AspNetCore.Mvc.Formatters.Json|6.0.0",
                "Microsoft.AspNetCore.Mvc.Formatters.Xml|6.0.0",
                "Microsoft.AspNetCore.Mvc.Localization|6.0.0",
                "Microsoft.AspNetCore.Mvc.Razor|6.0.0",
                "Microsoft.AspNetCore.Mvc.RazorPages|6.0.0",
                "Microsoft.AspNetCore.Mvc.TagHelpers|6.0.0",
                "Microsoft.AspNetCore.Mvc.ViewFeatures|6.0.0",
                "Microsoft.AspNetCore.Razor|6.0.0",
                "Microsoft.AspNetCore.Razor.Runtime|6.0.0",
                "Microsoft.AspNetCore.ResponseCaching.Abstractions|6.0.0",
                "Microsoft.AspNetCore.ResponseCaching|6.0.0",
                "Microsoft.AspNetCore.ResponseCompression|6.0.0",
                "Microsoft.AspNetCore.Rewrite|6.0.0",
                "Microsoft.AspNetCore.Routing.Abstractions|6.0.0",
                "Microsoft.AspNetCore.Routing|6.0.0",
                "Microsoft.AspNetCore.Server.HttpSys|6.0.0",
                "Microsoft.AspNetCore.Server.IIS|6.0.0",
                "Microsoft.AspNetCore.Server.IISIntegration|6.0.0",
                "Microsoft.AspNetCore.Server.Kestrel.Core|6.0.0",
                "Microsoft.AspNetCore.Server.Kestrel|6.0.0",
                "Microsoft.AspNetCore.Server.Kestrel.Transport.Quic|6.0.0",
                "Microsoft.AspNetCore.Server.Kestrel.Transport.Sockets|6.0.0",
                "Microsoft.AspNetCore.Session|6.0.0",
                "Microsoft.AspNetCore.SignalR.Common|6.0.0",
                "Microsoft.AspNetCore.SignalR.Core|6.0.0",
                "Microsoft.AspNetCore.SignalR|6.0.0",
                "Microsoft.AspNetCore.SignalR.Protocols.Json|6.0.0",
                "Microsoft.AspNetCore.StaticFiles|6.0.0",
                "Microsoft.AspNetCore.WebSockets|6.0.0",
                "Microsoft.AspNetCore.WebUtilities|6.0.0",
                "Microsoft.Extensions.Configuration.KeyPerFile|6.0.0",
                "Microsoft.Extensions.Diagnostics.HealthChecks.Abstractions|6.0.0",
                "Microsoft.Extensions.Diagnostics.HealthChecks|6.0.0",
                "Microsoft.Extensions.Features|6.0.0",
                "Microsoft.Extensions.FileProviders.Embedded|6.0.0",
                "Microsoft.Extensions.Identity.Core|6.0.0",
                "Microsoft.Extensions.Identity.Stores|6.0.0",
                "Microsoft.Extensions.Localization.Abstractions|6.0.0",
                "Microsoft.Extensions.Localization|6.0.0",
                "Microsoft.Extensions.ObjectPool|6.0.0",
                "Microsoft.Extensions.WebEncoders|6.0.0",
                "Microsoft.JSInterop|6.0.0",
                "Microsoft.Net.Http.Headers|6.0.0",
            ],
        ),
        ("Microsoft.AspNetCore.App.Runtime.linux-x64", "6.0.8", "sha512-3Hig5sP4ALm0aaB3cYCdhmW0a6SgT23ReaP5oYOZ9p1fQoQy4fHeLlU2LXQTXgJDopd3sQZCaWg639rJCYppiQ==", [], []),
        ("Microsoft.AspNetCore.App.Runtime.osx-x64", "6.0.8", "sha512-AQHu61cati6QzemklVlevQgChYJ3+msUUnXVDh51cEHhFEO/HBLKFWTiS1A49jnLBFpNUY98jPJMauyKIrh4jQ==", [], []),
        ("Microsoft.AspNetCore.App.Runtime.win-x64", "6.0.8", "sha512-fSuPkgA89T57pmGx2g6pcMSizT49ABL43d6s8Vp0PCzPjrme7UBISHATM9zP45Sq6GUhTZe2892wj7NmPa0wBA==", [], []),
        (
            "Microsoft.NETCore.App.Ref",
            "6.0.8",
            "sha512-TcZWOpmw+hWGQANrK0YWS3oHvtxdkn5A5JB284IdgXNvQ4rGABOPK8u52qB2bATbpSy3DbiMdobRxgAB2/mcJQ==",
            [],
            [
                "Microsoft.CSharp|4.4.0",
                "Microsoft.Win32.Primitives|4.3.0",
                "Microsoft.Win32.Registry|4.4.0",
                "runtime.debian.8-x64.runtime.native.System|4.3.0",
                "runtime.debian.8-x64.runtime.native.System.IO.Compression|4.3.0",
                "runtime.debian.8-x64.runtime.native.System.Net.Http|4.3.0",
                "runtime.debian.8-x64.runtime.native.System.Net.Security|4.3.0",
                "runtime.debian.8-x64.runtime.native.System.Security.Cryptography|4.3.0",
                "runtime.debian.8-x64.runtime.native.System.Security.Cryptography.OpenSsl|4.3.0",
                "runtime.fedora.23-x64.runtime.native.System|4.3.0",
                "runtime.fedora.23-x64.runtime.native.System.IO.Compression|4.3.0",
                "runtime.fedora.23-x64.runtime.native.System.Net.Http|4.3.0",
                "runtime.fedora.23-x64.runtime.native.System.Net.Security|4.3.0",
                "runtime.fedora.23-x64.runtime.native.System.Security.Cryptography|4.3.0",
                "runtime.fedora.23-x64.runtime.native.System.Security.Cryptography.OpenSsl|4.3.0",
                "runtime.fedora.24-x64.runtime.native.System|4.3.0",
                "runtime.fedora.24-x64.runtime.native.System.IO.Compression|4.3.0",
                "runtime.fedora.24-x64.runtime.native.System.Net.Http|4.3.0",
                "runtime.fedora.24-x64.runtime.native.System.Net.Security|4.3.0",
                "runtime.fedora.24-x64.runtime.native.System.Security.Cryptography|4.3.0",
                "runtime.fedora.24-x64.runtime.native.System.Security.Cryptography.OpenSsl|4.3.0",
                "runtime.opensuse.13.2-x64.runtime.native.System|4.3.0",
                "runtime.opensuse.13.2-x64.runtime.native.System.IO.Compression|4.3.0",
                "runtime.opensuse.13.2-x64.runtime.native.System.Net.Http|4.3.0",
                "runtime.opensuse.13.2-x64.runtime.native.System.Net.Security|4.3.0",
                "runtime.opensuse.13.2-x64.runtime.native.System.Security.Cryptography|4.3.0",
                "runtime.opensuse.13.2-x64.runtime.native.System.Security.Cryptography.OpenSsl|4.3.0",
                "runtime.opensuse.42.1-x64.runtime.native.System|4.3.0",
                "runtime.opensuse.42.1-x64.runtime.native.System.IO.Compression|4.3.0",
                "runtime.opensuse.42.1-x64.runtime.native.System.Net.Http|4.3.0",
                "runtime.opensuse.42.1-x64.runtime.native.System.Net.Security|4.3.0",
                "runtime.opensuse.42.1-x64.runtime.native.System.Security.Cryptography|4.3.0",
                "runtime.opensuse.42.1-x64.runtime.native.System.Security.Cryptography.OpenSsl|4.3.0",
                "runtime.osx.10.10-x64.runtime.native.System|4.3.0",
                "runtime.osx.10.10-x64.runtime.native.System.IO.Compression|4.3.0",
                "runtime.osx.10.10-x64.runtime.native.System.Net.Http|4.3.0",
                "runtime.osx.10.10-x64.runtime.native.System.Net.Security|4.3.0",
                "runtime.osx.10.10-x64.runtime.native.System.Security.Cryptography|4.3.0",
                "runtime.osx.10.10-x64.runtime.native.System.Security.Cryptography.Apple|4.3.0",
                "runtime.osx.10.10-x64.runtime.native.System.Security.Cryptography.OpenSsl|4.3.0",
                "runtime.rhel.7-x64.runtime.native.System|4.3.0",
                "runtime.rhel.7-x64.runtime.native.System.IO.Compression|4.3.0",
                "runtime.rhel.7-x64.runtime.native.System.Net.Http|4.3.0",
                "runtime.rhel.7-x64.runtime.native.System.Net.Security|4.3.0",
                "runtime.rhel.7-x64.runtime.native.System.Security.Cryptography|4.3.0",
                "runtime.rhel.7-x64.runtime.native.System.Security.Cryptography.OpenSsl|4.3.0",
                "runtime.ubuntu.14.04-x64.runtime.native.System|4.3.0",
                "runtime.ubuntu.14.04-x64.runtime.native.System.IO.Compression|4.3.0",
                "runtime.ubuntu.14.04-x64.runtime.native.System.Net.Http|4.3.0",
                "runtime.ubuntu.14.04-x64.runtime.native.System.Net.Security|4.3.0",
                "runtime.ubuntu.14.04-x64.runtime.native.System.Security.Cryptography|4.3.0",
                "runtime.ubuntu.14.04-x64.runtime.native.System.Security.Cryptography.OpenSsl|4.3.0",
                "runtime.ubuntu.16.04-x64.runtime.native.System|4.3.0",
                "runtime.ubuntu.16.04-x64.runtime.native.System.IO.Compression|4.3.0",
                "runtime.ubuntu.16.04-x64.runtime.native.System.Net.Http|4.3.0",
                "runtime.ubuntu.16.04-x64.runtime.native.System.Net.Security|4.3.0",
                "runtime.ubuntu.16.04-x64.runtime.native.System.Security.Cryptography|4.3.0",
                "runtime.ubuntu.16.04-x64.runtime.native.System.Security.Cryptography.OpenSsl|4.3.0",
                "runtime.ubuntu.16.10-x64.runtime.native.System|4.3.0",
                "runtime.ubuntu.16.10-x64.runtime.native.System.IO.Compression|4.3.0",
                "runtime.ubuntu.16.10-x64.runtime.native.System.Net.Http|4.3.0",
                "runtime.ubuntu.16.10-x64.runtime.native.System.Net.Security|4.3.0",
                "runtime.ubuntu.16.10-x64.runtime.native.System.Security.Cryptography|4.3.0",
                "runtime.ubuntu.16.10-x64.runtime.native.System.Security.Cryptography.OpenSsl|4.3.0",
                "System.AppContext|4.3.0",
                "System.Buffers|4.4.0",
                "System.Collections|4.3.0",
                "System.Collections.Concurrent|4.3.0",
                "System.Collections.Immutable|1.4.0",
                "System.Collections.NonGeneric|4.3.0",
                "System.Collections.Specialized|4.3.0",
                "System.ComponentModel|4.3.0",
                "System.ComponentModel.EventBasedAsync|4.3.0",
                "System.ComponentModel.Primitives|4.3.0",
                "System.ComponentModel.TypeConverter|4.3.0",
                "System.Console|4.3.0",
                "System.Data.Common|4.3.0",
                "System.Diagnostics.Contracts|4.3.0",
                "System.Diagnostics.Debug|4.3.0",
                "System.Diagnostics.DiagnosticSource|4.4.0",
                "System.Diagnostics.FileVersionInfo|4.3.0",
                "System.Diagnostics.Process|4.3.0",
                "System.Diagnostics.StackTrace|4.3.0",
                "System.Diagnostics.TextWriterTraceListener|4.3.0",
                "System.Diagnostics.Tools|4.3.0",
                "System.Diagnostics.TraceSource|4.3.0",
                "System.Diagnostics.Tracing|4.3.0",
                "System.Dynamic.Runtime|4.3.0",
                "System.Globalization|4.3.0",
                "System.Globalization.Calendars|4.3.0",
                "System.Globalization.Extensions|4.3.0",
                "System.IO|4.3.0",
                "System.IO.Compression|4.3.0",
                "System.IO.Compression.ZipFile|4.3.0",
                "System.IO.FileSystem|4.3.0",
                "System.IO.FileSystem.AccessControl|4.4.0",
                "System.IO.FileSystem.DriveInfo|4.3.0",
                "System.IO.FileSystem.Primitives|4.3.0",
                "System.IO.FileSystem.Watcher|4.3.0",
                "System.IO.IsolatedStorage|4.3.0",
                "System.IO.MemoryMappedFiles|4.3.0",
                "System.IO.Pipes|4.3.0",
                "System.IO.UnmanagedMemoryStream|4.3.0",
                "System.Linq|4.3.0",
                "System.Linq.Expressions|4.3.0",
                "System.Linq.Queryable|4.3.0",
                "System.Net.Http|4.3.0",
                "System.Net.NameResolution|4.3.0",
                "System.Net.Primitives|4.3.0",
                "System.Net.Requests|4.3.0",
                "System.Net.Security|4.3.0",
                "System.Net.Sockets|4.3.0",
                "System.Net.WebHeaderCollection|4.3.0",
                "System.ObjectModel|4.3.0",
                "System.Private.DataContractSerialization|4.3.0",
                "System.Reflection|4.3.0",
                "System.Reflection.Emit|4.3.0",
                "System.Reflection.Emit.ILGeneration|4.3.0",
                "System.Reflection.Emit.Lightweight|4.3.0",
                "System.Reflection.Extensions|4.3.0",
                "System.Reflection.Metadata|1.5.0",
                "System.Reflection.Primitives|4.3.0",
                "System.Reflection.TypeExtensions|4.3.0",
                "System.Resources.ResourceManager|4.3.0",
                "System.Runtime|4.3.0",
                "System.Runtime.Extensions|4.3.0",
                "System.Runtime.Handles|4.3.0",
                "System.Runtime.InteropServices|4.3.0",
                "System.Runtime.InteropServices.RuntimeInformation|4.3.0",
                "System.Runtime.Loader|4.3.0",
                "System.Runtime.Numerics|4.3.0",
                "System.Runtime.Serialization.Formatters|4.3.0",
                "System.Runtime.Serialization.Json|4.3.0",
                "System.Runtime.Serialization.Primitives|4.3.0",
                "System.Security.AccessControl|4.4.0",
                "System.Security.Claims|4.3.0",
                "System.Security.Cryptography.Algorithms|4.3.0",
                "System.Security.Cryptography.Cng|4.4.0",
                "System.Security.Cryptography.Csp|4.3.0",
                "System.Security.Cryptography.Encoding|4.3.0",
                "System.Security.Cryptography.OpenSsl|4.4.0",
                "System.Security.Cryptography.Primitives|4.3.0",
                "System.Security.Cryptography.X509Certificates|4.3.0",
                "System.Security.Cryptography.Xml|4.4.0",
                "System.Security.Principal|4.3.0",
                "System.Security.Principal.Windows|4.4.0",
                "System.Text.Encoding|4.3.0",
                "System.Text.Encoding.Extensions|4.3.0",
                "System.Text.RegularExpressions|4.3.0",
                "System.Threading|4.3.0",
                "System.Threading.Overlapped|4.3.0",
                "System.Threading.Tasks|4.3.0",
                "System.Threading.Tasks.Extensions|4.3.0",
                "System.Threading.Tasks.Parallel|4.3.0",
                "System.Threading.Thread|4.3.0",
                "System.Threading.ThreadPool|4.3.0",
                "System.Threading.Timer|4.3.0",
                "System.ValueTuple|4.3.0",
                "System.Xml.ReaderWriter|4.3.0",
                "System.Xml.XDocument|4.3.0",
                "System.Xml.XmlDocument|4.3.0",
                "System.Xml.XmlSerializer|4.3.0",
                "System.Xml.XPath|4.3.0",
                "System.Xml.XPath.XDocument|4.3.0",
            ],
        ),
        ("Microsoft.NETCore.App.Runtime.linux-x64", "6.0.8", "sha512-cjVzAUiYxPv949mXl0IbwzSRq0xBTGcW3N619CUcCwe35Ma1C1Tg1nh75Xc+OEn5+eAMW/S66dy+kQhdc277tA==", [], []),
        ("Microsoft.NETCore.App.Runtime.osx-x64", "6.0.8", "sha512-RDOy3pzl0sutv5U3JAx23JWiw2UCoHAPNsCo35TA8MU2DM+LMDXN/lxi2cslot6GfFsxe0cYhclkEocHa2xMPQ==", [], []),
        ("Microsoft.NETCore.App.Runtime.win-x64", "6.0.8", "sha512-pgpxzvQPZzBPD1lWulgRO/aafBhSBLhqH+SrBD+sYSIu7eswlxE5icW/r8o60fNFKYVg0CFvrnmCut5YpTT27Q==", [], []),
        ("Mono.Cecil", "0.11.4", "sha512-CnjwUMmFHnScNG8e/4DRZQQX67H5ajekRDudmZ6Fy1jCLhyH1jjzbQCOEFhBLa2NjPWQpMF+RHdBJY8a7GgmlA==", [], []),
    ],
)
