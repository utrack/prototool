# Uber Protobuf Style Guide V2

This is the second version of the Uber Protobuf Style Guide.

## Example

See the [uber](uber) directory for an example of all concepts explained in this Style Guide.

## Package Naming

Some conventions:

- A **package name** as the full package name, i.e. `foo.bar.baz.v1`.
- A **package sub-name** as a part of a package name, ie `foo`, `bar`, or `baz`.
- A **package version** is the last package sub-name that specifies the version,
  i.e. `v1`, `v1beta1`, or `v2`.

Package sub-names should be short and descriptive, and can use abbreviations if necesary.
Package sub-names should only include characters in the range `[a-z0-9]`, i.e always lowercase
and with only letter and digit characters. If names get too long or have underscores, the
generated stubs in certain languages are less than idiomatic.

As illustrative examples, the following are not acceptable package names.

```proto
// Note that specifying multiple packages is not valid Protobuf, however
// we do this here for brevity.

// The package sub-name credit_card_analysis is not short, and contains underscores.
package uber.finance.credit_card_analysis.v1;
// The package sub-name creditcardanalysis is longer than desired.
package uber.finance.creditcardanalysis.v1;
```

The following are acceptable package names.

```proto
// Each package sub-name is short and to the point.
package uber.trips.watch.v1;
// Grouping by finance and then payment is acceptable.
package uber.finance.payment.v1;
// Cca is for credit card analysis.
package uber.finance.cca.v1;
```

## Package Versioning

The last package sub-name should be a major version of the package, or the major version
followed by the beta version, specified as `vMAJOR` or `vMAJORbetaBETA`, where `MAJOR` and `BETA`
are both greater than 0. The following are examples of acceptable package names.

```proto
package foo.bar.v1beta1;
package foo.bar.v1beta2;
package foo.bar.v1;
package foo.bar.v2beta1;
package foo.bar.v2;
package bar.v2;
```

As illustrative examples, the following are not acceptable package names.

```proto
// No version.
package foo.bar;
// Major version is not greater than 0.
package foo.bar.v0;
// Beta version is not greater than 0.
package foo.bar.v1beta0;
```

Packages with only a major version are considered **stable** packages, and packages with a major
and beta version are considered **beta** packages.

Breaking changes should never be made in stable packages, and stable packages should never depend
on beta packages. Both wire-incompatible and source-code-incompatible changes are considered
breaking changes. The following are the list of changes currently understood to be breaking.

- Deleting or renaming a package.
- Deleting or renaming an enum, enum value, message, message field, service, or service method.
- Changing the type of a message field.
- Changing the tag of a message field.
- Changing the label of a message field, i.e. optional, repeated, required.
- Moving a message field currently in a oneof out of the oneof.
- Moving a message field currently not in a oneof into the oneof.
- Changing the function signature of a method.
- Changing the stream value of a method request or response.

Beta packages should be used with extreme caution, and are not recommended.

Instead of making a breaking change, rely on deprecation of types.

```proto
// Note that all enums, messages, services, and service methods require
// sentence comments, and each service must be in a separate file, as
// outlined below, however we omit this here for brevity.

enum Foo {
  option deprecated = true;
  FOO_INVALID = 0;
  FOO_ONE = 1;
}

enum Bar {
  BAR_INVALID = 0;
  BAR_ONE = 1 [deprecated = true];
  BAR_TWO = 2;
}

message Baz {
  option deprecated = true;
  int64 one = 1;
}

message Bat {
  int64 one = 1 [deprecated = true];
  int64 two = 2;
}

service BamAPI {
  option deprecated = true;
  rpc Hello(HelloRequest) returns (HelloResponse) {}
}

service BanAPI {
  rpc (GoodbyeRequest) returns (GoodbyeResponse) {
    option deprecated = true;
  }
}
```

If you really want to make a breaking change, or just want to clean up a package, make a new
version of the package by incrementing the major version and copy your definitions as
necessary. For example, copy `foo.bar.v1` to `foo.bar.v2`, and do any cleanups required.
This is not a breaking change as `foo.bar.v2` is a new package. Of course, you are responsible
for the migration of your callers.

## Directory Structure

Files should be stored in a directory structure that matches their package sub-names. All files
in a given directory should be in the same package.

The following is an example of this in practice.

```
.
└── uber
    ├── finance
    │   ├── cca
    │   │   └── v1
    │   │       └── cca.proto // package uber.finance.cca.v1
    │   └── payment
    │       ├── v1
    │       │   └── payment.proto // package uber.payment.v1
    │       └── v1beta1
    │           └── payment.proto // package uber.payment.v1beta1
    └── trips
        └── watch
            ├── v1
            │   ├── trip_watcher_api.proto // package uber.trips.watch.v1
            │   └── watch.proto // package uber.trips.watch.v1
            └── v2
                ├── trip_watcher_api.proto // package uber.trips.watch.v2
                └── watch.proto // pacakge uber.trips.watch.v2
```
