schema_version = 1

project {
  license        = "MPL-2.0"
  copyright_year = 2025
  copyright_holder = "Michael Kosir"

  header_ignore = [
    # examples used within documentation (prose)
    "examples/**",

    # golangci-lint tooling configuration
    ".golangci.yml",

    # GoReleaser tooling configuration
    ".goreleaser.yml",
  ]
}
