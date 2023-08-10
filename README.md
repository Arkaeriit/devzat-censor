# Devzat Censor

[Devzat](https://github.com/quackduck/devzat) houses a lot of discussions. To ensure polite and pleasant discussions, a censoring of rude words is nice. This plugins acts as a middleware that passes the messages through [Rustrict](https://github.com/finnbear/rustrict) to remove the worste words.

## Admin usage

The plugin is made for a single-file executable. It is configured with the following environment variable.

|Variable name |Description                    |Default                                                                     |
|--------------|-------------------------------|----------------------------------------------------------------------------|
|`PLUGIN_HOST` |URL of the chat-room interface |`https://devzat.hackclub.com:5556`                                          |
|`PLUGIN_TOKEN`|Authentication token           |Does not defaults to anything. The program panics if the token is not given.|

