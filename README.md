# Slack specs

This projects goal is to try and create high quality OpenAPI v3 specs of [Slack's Web API](https://api.slack.com/web)

If you are just after the API specs you will find them here:

- [./dist/openapi.manicured.yaml](./dist/openapi.manicured.yaml) - Subset of API operations that have been manually updated/verified
- [./dist/openapi.complete.yaml](./dist/openapi.complete.yaml) - Complete list of API oprations fro the original Slack API specs

The starting point of this project was to manually extract the [OpenAPI v2 specs that slack publish](https://github.com/slackapi/slack-api-specs/tree/master/web-api) which are unsuitable for generating a client with, convert them to OpenAPI v3 yaml files, split each tag/category out into it's own folder and then striving to make manual fixes to them as required.

## Authoring, Verifying and Validating the API specs

This repository uses a spec-first approach, meaning that the spec definitions are authored by contributers, it is NOT generated from code or any external systems/resources.

Authoring such a complicated OpenAPI spec can be laborious, so in order to make the authoring process easier we do a number of things that help faciliate an acceptable authoring experience, these include:

- Splitting the OpenAPI spec definitions across multiple files and compiling them together as a build step using [`conflate`](https://github.com/miracl/conflate)
- Generating/compiling two versions of the Slack API spec into
  - A `complete` one that all the API Operations and components that were imported as part of the original v2 OpenAPI spec
  - A `manicured` one that contains a subset of API operations that have been manually reviewed to be accurate and/or complete enough to include in (may still contain issues).
- Using [`spectral`](https://stoplight.io/open-source/spectral/) to lint the generated OpenAPI v3 specs to ensure they are valid according to the OpenAPI spec and some other desired linting rules (eg preventing unused components)
- Generating client libraries from the API spec in different programming languages to verify that the schemas in the spec are robust enough for tools to generate useful code.
  - `./test-harnesses/go` - Golang codegen using [`deepmap/oapi-codegen`](https://github.com/deepmap/oapi-codegen)
  - `./test-harnesses/rust` - Rust codegen using [`OpenAPITools/openapi-generator`](https://github.com/OpenAPITools/openapi-generator)
- Maintaining `unit` tests use fixtures based on the official slack documentation
  - Validate that code generated from the spec functions
  - WIP: Validate the requests/responses recieved when running the tests pass the validation rules specified by the OpenAPI spec
- Maintaining `smoke` tests that execute against a real Slack workspace instance to
  - Validate that code generated from the spec functions
  - WIP: Validate the requests/responses recieved when running the tests pass the validation rules specified by the OpenAPI spec

## Contributing Quickstart

### Global dependencies

You will need to have the following global dependencies installed on your local machine

- [go](https://go.dev/doc/install)
- [cargo](https://doc.rust-lang.org/cargo/getting-started/installation.html)

### Local config files

```sh
# Copy the example .env.local file (used to override environment variables for local development)
cp docs/examples/.env.local .

# Copy the recommended vscode settings to your workspace config
cp .vscode/settings.recommended.json .vscode/settings.json
```

## Run development scripts

```sh
# Install/setup tool deps
make deps.install

# Generate OpenAPI specs and test harness client libraries
make generate

# Verify OpenAPI specs and test harnesses client librarires using static analysis tools
make verify

# Test OpenAPI specs and test harnesses client librarires and generate reports
make test
```

For more details and before opening pull requests please read the full [./CONTRIBUTING.md](CONTRIBUTING.md)

## Wishlist

- Run the smoke tests against a real slack workspace for PRs
- Add automatic request and response validation for both the `unit` and `smoke` tests. Some options include:
  - [`EXXETA/openapi-cop`](https://github.com/EXXETA/openapi-cop)
  - [`kin-openapi/openapi3`](https://github.com/getkin/kin-openapi) example in [`deepmap/oapi-codegen`](https://github.com/deepmap/oapi-codegen/blob/master/pkg/middleware/oapi_validate.go)
- Add a rust `test-harness`
- Generate a CLI program from the OpenAPI spec and have a `smoke` test suite that uses the generated CLI.

## Progress

| Operation                                                  | in complete | in manicured | Go unit tests | Go smoke tests |
| ---------------------------------------------------------- | ----------- | ------------ | ------------- | -------------- |
| `admin.apps.approve`                                       | ✅          | ❌           | ❌            | ❌             |
| `admin.apps.approved.list`                                 | ✅          | ❌           | ❌            | ❌             |
| `admin.apps.requests.list`                                 | ✅          | ❌           | ❌            | ❌             |
| `admin.apps.restrict`                                      | ✅          | ❌           | ❌            | ❌             |
| `admin.apps.restricted.list`                               | ✅          | ❌           | ❌            | ❌             |
| `admin.conversations.archive`                              | ✅          | ❌           | ❌            | ❌             |
| `admin.conversations.convertToPrivate`                     | ✅          | ❌           | ❌            | ❌             |
| `admin.conversations.create`                               | ✅          | ❌           | ❌            | ❌             |
| `admin.conversations.delete`                               | ✅          | ❌           | ❌            | ❌             |
| `admin.conversations.disconnectShared`                     | ✅          | ❌           | ❌            | ❌             |
| `admin.conversations.ekm.listOriginalConnectedChannelInfo` | ✅          | ❌           | ❌            | ❌             |
| `admin.conversations.getConversationPrefs`                 | ✅          | ❌           | ❌            | ❌             |
| `admin.conversations.getTeams`                             | ✅          | ❌           | ❌            | ❌             |
| `admin.conversations.invite`                               | ✅          | ❌           | ❌            | ❌             |
| `admin.conversations.rename`                               | ✅          | ❌           | ❌            | ❌             |
| `admin.conversations.restrictAccess.addGroup`              | ✅          | ❌           | ❌            | ❌             |
| `admin.conversations.restrictAccess.listGroups`            | ✅          | ❌           | ❌            | ❌             |
| `admin.conversations.restrictAccess.removeGroup`           | ✅          | ❌           | ❌            | ❌             |
| `admin.conversations.search`                               | ✅          | ❌           | ❌            | ❌             |
| `admin.conversations.setConversationPrefs`                 | ✅          | ❌           | ❌            | ❌             |
| `admin.conversations.setTeams`                             | ✅          | ❌           | ❌            | ❌             |
| `admin.conversations.unarchive`                            | ✅          | ❌           | ❌            | ❌             |
| `admin.inviteRequests.approve`                             | ✅          | ❌           | ❌            | ❌             |
| `admin.inviteRequests.approved.list`                       | ✅          | ❌           | ❌            | ❌             |
| `admin.inviteRequests.denied.list`                         | ✅          | ❌           | ❌            | ❌             |
| `admin.inviteRequests.deny`                                | ✅          | ❌           | ❌            | ❌             |
| `admin.inviteRequests.list`                                | ✅          | ❌           | ❌            | ❌             |
| `admin.teams.admins.list`                                  | ✅          | ❌           | ❌            | ❌             |
| `admin.teams.create`                                       | ✅          | ❌           | ❌            | ❌             |
| `admin.teams.list`                                         | ✅          | ❌           | ❌            | ❌             |
| `admin.teams.owners.list`                                  | ✅          | ❌           | ❌            | ❌             |
| `admin.teams.settings.info`                                | ✅          | ❌           | ❌            | ❌             |
| `admin.teams.settings.setDefaultChannels`                  | ✅          | ❌           | ❌            | ❌             |
| `admin.teams.settings.setDescription`                      | ✅          | ❌           | ❌            | ❌             |
| `admin.teams.settings.setDiscoverability`                  | ✅          | ❌           | ❌            | ❌             |
| `admin.teams.settings.setIcon`                             | ✅          | ❌           | ❌            | ❌             |
| `admin.teams.settings.setName`                             | ✅          | ❌           | ❌            | ❌             |
| `admin.usergroups.addChannels`                             | ✅          | ❌           | ❌            | ❌             |
| `admin.usergroups.addTeams`                                | ✅          | ❌           | ❌            | ❌             |
| `admin.usergroups.listChannels`                            | ✅          | ❌           | ❌            | ❌             |
| `admin.usergroups.removeChannels`                          | ✅          | ❌           | ❌            | ❌             |
| `admin.users.assign`                                       | ✅          | ❌           | ❌            | ❌             |
| `admin.users.invite`                                       | ✅          | ❌           | ❌            | ❌             |
| `admin.users.list`                                         | ✅          | ❌           | ❌            | ❌             |
| `admin.users.remove`                                       | ✅          | ❌           | ❌            | ❌             |
| `admin.users.session.invalidate`                           | ✅          | ❌           | ❌            | ❌             |
| `admin.users.session.reset`                                | ✅          | ❌           | ❌            | ❌             |
| `admin.users.setAdmin`                                     | ✅          | ❌           | ❌            | ❌             |
| `admin.users.setExpiration`                                | ✅          | ❌           | ❌            | ❌             |
| `admin.users.setOwner`                                     | ✅          | ❌           | ❌            | ❌             |
| `admin.users.setRegular`                                   | ✅          | ❌           | ❌            | ❌             |
| `api.test`                                                 | ✅          | ❌           | ❌            | ❌             |
| `apps.event.authorizations.list`                           | ✅          | ❌           | ❌            | ❌             |
| `apps.permissions.info`                                    | ✅          | ❌           | ❌            | ❌             |
| `apps.permissions.request`                                 | ✅          | ❌           | ❌            | ❌             |
| `apps.permissions.resources.list`                          | ✅          | ❌           | ❌            | ❌             |
| `apps.permissions.scopes.list`                             | ✅          | ❌           | ❌            | ❌             |
| `apps.permissions.users.list`                              | ✅          | ❌           | ❌            | ❌             |
| `apps.permissions.users.request`                           | ✅          | ❌           | ❌            | ❌             |
| `apps.uninstall`                                           | ✅          | ❌           | ❌            | ❌             |
| `auth.revoke`                                              | ✅          | ❌           | ❌            | ❌             |
| `auth.test`                                                | ✅          | ❌           | ❌            | ❌             |
| `bots.info`                                                | ✅          | ❌           | ❌            | ❌             |
| `calls.add`                                                | ✅          | ❌           | ❌            | ❌             |
| `calls.end`                                                | ✅          | ❌           | ❌            | ❌             |
| `calls.info`                                               | ✅          | ❌           | ❌            | ❌             |
| `calls.participants.add`                                   | ✅          | ❌           | ❌            | ❌             |
| `calls.participants.remove`                                | ✅          | ❌           | ❌            | ❌             |
| `calls.update`                                             | ✅          | ❌           | ❌            | ❌             |
| `chat.delete`                                              | ✅          | ❌           | ❌            | ❌             |
| `chat.deleteScheduledMessage`                              | ✅          | ❌           | ❌            | ❌             |
| `chat.getPermalink`                                        | ✅          | ❌           | ❌            | ❌             |
| `chat.meMessage`                                           | ✅          | ❌           | ❌            | ❌             |
| `chat.postEphemeral`                                       | ✅          | ❌           | ❌            | ❌             |
| `chat.postMessage`                                         | ✅          | ❌           | ❌            | ❌             |
| `chat.scheduleMessage`                                     | ✅          | ❌           | ❌            | ❌             |
| `chat.scheduledMessages.list`                              | ✅          | ❌           | ❌            | ❌             |
| `chat.unfurl`                                              | ✅          | ❌           | ❌            | ❌             |
| `chat.update`                                              | ✅          | ❌           | ❌            | ❌             |
| `conversations.archive`                                    | ✅          | ❌           | ❌            | ❌             |
| `conversations.close`                                      | ✅          | ❌           | ❌            | ❌             |
| `conversations.create`                                     | ✅          | ❌           | ❌            | ❌             |
| `conversations.history`                                    | ✅          | ❌           | ❌            | ❌             |
| `conversations.info`                                       | ✅          | ❌           | ❌            | ❌             |
| `conversations.invite`                                     | ✅          | ❌           | ❌            | ❌             |
| `conversations.join`                                       | ✅          | ❌           | ❌            | ❌             |
| `conversations.kick`                                       | ✅          | ❌           | ❌            | ❌             |
| `conversations.leave`                                      | ✅          | ❌           | ❌            | ❌             |
| `conversations.list`                                       | ✅          | ✅           | ✅            | 🚧             |
| `conversations.mark`                                       | ✅          | ❌           | ❌            | ❌             |
| `conversations.members`                                    | ✅          | ❌           | ❌            | ❌             |
| `conversations.open`                                       | ✅          | ❌           | ❌            | ❌             |
| `conversations.rename`                                     | ✅          | ❌           | ❌            | ❌             |
| `conversations.replies`                                    | ✅          | ❌           | ❌            | ❌             |
| `conversations.setPurpose`                                 | ✅          | ❌           | ❌            | ❌             |
| `conversations.setTopic`                                   | ✅          | ❌           | ❌            | ❌             |
| `conversations.unarchive`                                  | ✅          | ❌           | ❌            | ❌             |
| `dialog.open`                                              | ✅          | ❌           | ❌            | ❌             |
| `dnd.endDnd`                                               | ✅          | ❌           | ❌            | ❌             |
| `dnd.endSnooze`                                            | ✅          | ❌           | ❌            | ❌             |
| `dnd.info`                                                 | ✅          | ❌           | ❌            | ❌             |
| `dnd.setSnooze`                                            | ✅          | ❌           | ❌            | ❌             |
| `dnd.teamInfo`                                             | ✅          | ❌           | ❌            | ❌             |
| `files.comments.delete`                                    | ✅          | ❌           | ❌            | ❌             |
| `files.delete`                                             | ✅          | ❌           | ❌            | ❌             |
| `files.info`                                               | ✅          | ❌           | ❌            | ❌             |
| `files.list`                                               | ✅          | ❌           | ❌            | ❌             |
| `files.remote.add`                                         | ✅          | ❌           | ❌            | ❌             |
| `files.remote.info`                                        | ✅          | ❌           | ❌            | ❌             |
| `files.remote.list`                                        | ✅          | ❌           | ❌            | ❌             |
| `files.remote.remove`                                      | ✅          | ❌           | ❌            | ❌             |
| `files.remote.share`                                       | ✅          | ❌           | ❌            | ❌             |
| `files.remote.update`                                      | ✅          | ❌           | ❌            | ❌             |
| `files.revokePublicURL`                                    | ✅          | ❌           | ❌            | ❌             |
| `files.sharedPublicURL`                                    | ✅          | ❌           | ❌            | ❌             |
| `files.upload`                                             | ✅          | ❌           | ❌            | ❌             |
| `migration.exchange`                                       | ✅          | ❌           | ❌            | ❌             |
| `oauth.access`                                             | ✅          | ❌           | ❌            | ❌             |
| `oauth.token`                                              | ✅          | ❌           | ❌            | ❌             |
| `oauth.v2.access`                                          | ✅          | ❌           | ❌            | ❌             |
| `pins.add`                                                 | ✅          | ❌           | ❌            | ❌             |
| `pins.list`                                                | ✅          | ❌           | ❌            | ❌             |
| `pins.remove`                                              | ✅          | ❌           | ❌            | ❌             |
| `reactions.add`                                            | ✅          | ❌           | ❌            | ❌             |
| `reactions.get`                                            | ✅          | ❌           | ❌            | ❌             |
| `reactions.list`                                           | ✅          | ❌           | ❌            | ❌             |
| `reactions.remove`                                         | ✅          | ❌           | ❌            | ❌             |
| `reminders.add`                                            | ✅          | ❌           | ❌            | ❌             |
| `reminders.complete`                                       | ✅          | ❌           | ❌            | ❌             |
| `reminders.delete`                                         | ✅          | ❌           | ❌            | ❌             |
| `reminders.info`                                           | ✅          | ❌           | ❌            | ❌             |
| `reminders.list`                                           | ✅          | ❌           | ❌            | ❌             |
| `rtm.connect`                                              | ✅          | ❌           | ❌            | ❌             |
| `search.messages`                                          | ✅          | ❌           | ❌            | ❌             |
| `stars.add`                                                | ✅          | ❌           | ❌            | ❌             |
| `stars.list`                                               | ✅          | ❌           | ❌            | ❌             |
| `stars.remove`                                             | ✅          | ❌           | ❌            | ❌             |
| `team.accessLogs`                                          | ✅          | ❌           | ❌            | ❌             |
| `team.billableInfo`                                        | ✅          | ❌           | ❌            | ❌             |
| `team.info`                                                | ✅          | ❌           | ❌            | ❌             |
| `team.integrationLogs`                                     | ✅          | ❌           | ❌            | ❌             |
| `team.profile.get`                                         | ✅          | ❌           | ❌            | ❌             |
| `usergroups.create`                                        | ✅          | ❌           | ❌            | ❌             |
| `usergroups.disable`                                       | ✅          | ❌           | ❌            | ❌             |
| `usergroups.enable`                                        | ✅          | ❌           | ❌            | ❌             |
| `usergroups.list`                                          | ✅          | ❌           | ❌            | ❌             |
| `usergroups.update`                                        | ✅          | ❌           | ❌            | ❌             |
| `usergroups.users.list`                                    | ✅          | ❌           | ❌            | ❌             |
| `usergroups.users.update`                                  | ✅          | ❌           | ❌            | ❌             |
| `users.conversations`                                      | ✅          | ❌           | ❌            | ❌             |
| `users.deletePhoto`                                        | ✅          | ❌           | ❌            | ❌             |
| `users.getPresence`                                        | ✅          | ❌           | ❌            | ❌             |
| `users.identity`                                           | ✅          | ❌           | ❌            | ❌             |
| `users.info`                                               | ✅          | ❌           | ❌            | ❌             |
| `users.list`                                               | ✅          | ❌           | ❌            | ❌             |
| `users.lookupByEmail`                                      | ✅          | ❌           | ❌            | ❌             |
| `users.profile.get`                                        | ✅          | ❌           | ❌            | ❌             |
| `users.profile.set`                                        | ✅          | ❌           | ❌            | ❌             |
| `users.setActive`                                          | ✅          | ❌           | ❌            | ❌             |
| `users.setPhoto`                                           | ✅          | ❌           | ❌            | ❌             |
| `users.setPresence`                                        | ✅          | ❌           | ❌            | ❌             |
| `views.open`                                               | ✅          | ❌           | ❌            | ❌             |
| `views.publish`                                            | ✅          | ❌           | ❌            | ❌             |
| `views.push`                                               | ✅          | ❌           | ❌            | ❌             |
| `views.update`                                             | ✅          | ❌           | ❌            | ❌             |
| `workflows.stepCompleted`                                  | ✅          | ❌           | ❌            | ❌             |
| `workflows.stepFailed`                                     | ✅          | ❌           | ❌            | ❌             |
| `workflows.updateStep`                                     | ✅          | ❌           | ❌            | ❌             |
