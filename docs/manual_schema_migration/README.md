# Slack API specs

The OpenAPI definitations here are taken from the [Official Slack API Specs](https://github.com/slackapi/slack-api-specs) on github.

They're then converted to OpenAPI v3 using the online [editor.swagger.io](https://editor.swagger.io/) tool, exported as yaml and then any minor remaining linting issues are fixed.

## Generating the rust client

A rust client is generated from the OpenAPI specs following similar steps as described [here](https://www.twilio.com/docs/openapi/generating-a-rust-client-for-twilios-api#setup)

```sh
# install the client generator
brew install openapi-generator

# Generate client into temp directory
openapi-generator generate -g rust \
  -i api_schema/v3.temp_modifications.openapi.yaml \
  -o ./tmp

# copy useful bit out of the generated client
cp -r ./tmp/src/ ./src/api_client/

mv ./src/api_client/lib.rs ./src/api_client/mod.rs
```

You may still need to:

- Manually update `use crate::*` references
- Remove/hoist parts of `mod.rs` into the crate root
